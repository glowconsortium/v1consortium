package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"

	"v1consortium/internal/logic/workflowbridge"
	"v1consortium/internal/pkg/riverjobsv2"
	"v1consortium/internal/service"
	signupv2 "v1consortium/internal/workflow/signupv2"
)

type RiverComponents struct {
	WorkflowExecutor *riverjobsv2.WorkflowExecutor
	RiverClient      *river.Client[pgx.Tx]     // Fixed: River client uses pgx.Tx, not pgxpool.Tx
	WorkflowStore    riverjobsv2.WorkflowStore // Database store for workflow persistence
	DBPool           *pgxpool.Pool
	Logger           *slog.Logger
}

var (
	RiverWorkerV2 = gcmd.Command{
		Name:  "river_worker",
		Usage: "river_worker",
		Brief: "start river job worker (v2)",
		Func: func(ctx context.Context, parser *gcmd.Parser) error {
			g.Log().Info(ctx, "V1 Consortium River Worker (v2) starting...")

			components, err := initializeRiverComponents(ctx)
			if err != nil {
				return err
			}
			defer components.DBPool.Close()

			// Register workflows
			if err := registerAllWorkers(components); err != nil {
				return fmt.Errorf("failed to register workflows: %w", err)
			}

			// Start River job processing
			if err := components.RiverClient.Start(ctx); err != nil {
				return fmt.Errorf("failed to start River client: %w", err)
			}

			// Setup background monitoring jobs
			setupRiverBackgroundJobs(ctx, components)

			// Setup graceful shutdown
			setupShutdownHandler(ctx, components.RiverClient)

			g.Log().Info(ctx, "River worker (v2) ready and processing jobs")
			g.Listen()
			return nil
		},
	}
)

// initializeRiverComponents sets up all River-related components for v2
func initializeRiverComponents(ctx context.Context) (*RiverComponents, error) {
	// Initialize logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	// Get database URL
	dbURL := getRiverDBURL(ctx)
	if dbURL == "" {
		return nil, fmt.Errorf("database URL not configured")
	}

	// Create database connection pool
	dbPool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create database pool: %w", err)
	}

	// Test database connection
	if err := dbPool.Ping(ctx); err != nil {
		dbPool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	g.Log().Info(ctx, "Database connection established")

	// Create database workflow store
	workflowStore := riverjobsv2.NewDatabaseWorkflowStore(dbPool)

	// Create workflow executor with database store
	workflowExecutor := riverjobsv2.NewWorkflowExecutor(workflowStore)

	// Create Workers bundle and register the workflow executor
	workers := river.NewWorkers()
	river.AddWorker[riverjobsv2.WorkflowArgs](workers, workflowExecutor)

	// Create River client with the workers bundle
	riverClient, err := river.NewClient(riverpgxv5.New(dbPool), &river.Config{
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {MaxWorkers: 100},
			"critical":         {MaxWorkers: 50},
			"external":         {MaxWorkers: 10},
			"notifications":    {MaxWorkers: 25},
		},
		Workers: workers,
	})
	if err != nil {
		dbPool.Close()
		return nil, fmt.Errorf("failed to create River client: %w", err)
	}

	// Set client reference in executor for scheduling next steps
	workflowExecutor.SetClient(riverClient)

	g.Log().Info(ctx, "River components (v2) initialized")

	return &RiverComponents{
		WorkflowExecutor: workflowExecutor,
		RiverClient:      riverClient,
		WorkflowStore:    workflowStore,
		DBPool:           dbPool,
		Logger:           logger,
	}, nil
}

// registerAllWorkers registers all workflows with the executor
func registerAllWorkers(components *RiverComponents) error {
	// Register the signup workflow with the executor
	signupWorkflow := signupv2.NewSignupWorkflow()
	components.WorkflowExecutor.RegisterWorkflow(signupWorkflow)
	components.Logger.Info("Registered SignupWorkflow (v2)")

	// Note: WorkflowExecutor is already registered with River during client creation
	components.Logger.Info("WorkflowExecutor registered with River during client initialization")

	return nil
}

// setupRiverBackgroundJobs configures monitoring and maintenance jobs for v2
func setupRiverBackgroundJobs(ctx context.Context, components *RiverComponents) {
	type backgroundJob struct {
		name     string
		enabled  bool
		cronExpr string
		jobFunc  func(context.Context)
	}

	// Define background jobs (simplified for v2)
	jobs := []backgroundJob{
		{
			name:     "health check",
			enabled:  true,
			cronExpr: "0 * * * * *", // Every minute
			jobFunc: func(ctx context.Context) {
				g.Log().Debug(ctx, "River worker (v2) health check - running")
			},
		},
		{
			name:     "metrics collection",
			enabled:  true,
			cronExpr: "0 */5 * * * *", // Every 5 minutes
			jobFunc: func(ctx context.Context) {
				// Collect basic metrics from River
				g.Log().Debug(ctx, "Collecting River (v2) metrics")
			},
		},
	}

	// Schedule enabled jobs
	for _, job := range jobs {
		if job.enabled {
			if _, err := gcron.AddSingleton(ctx, job.cronExpr, func(ctx context.Context) {
				defer func() {
					if err := recover(); err != nil {
						g.Log().Errorf(ctx, "Background job %s panicked: %v", job.name, err)
					}
				}()
				job.jobFunc(ctx)
			}, job.name); err != nil {
				g.Log().Errorf(ctx, "Failed to schedule background job %s: %v", job.name, err)
			}
		}
	}

	g.Log().Info(ctx, "River (v2) background monitoring jobs configured successfully")
}

// setupShutdownHandler configures graceful shutdown for v2
func setupShutdownHandler(ctx context.Context, riverClient *river.Client[pgx.Tx]) {
	gproc.AddSigHandlerShutdown(func(sig os.Signal) {
		g.Log().Infof(ctx, "Received signal %s, shutting down gracefully...", sig)

		stopCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := riverClient.Stop(stopCtx); err != nil {
			g.Log().Errorf(ctx, "Error stopping River client: %v", err)
		} else {
			g.Log().Info(ctx, "River client stopped successfully")
		}

		g.Log().Info(ctx, "River worker (v2) shutdown complete")
	})
}

// getRiverDBURL gets the database URL for River
func getRiverDBURL(ctx context.Context) string {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = os.Getenv("RIVER_DATABASE_URL")
	}

	if dbURL == "" {
		dbURL = g.Cfg().MustGet(ctx, "river.databaseUrl").String()
	}

	if dbURL == "" {
		g.Log().Fatal(ctx, "Database URL not found in environment variables or config")
		panic("Database URL not configured")
	}
	return dbURL
}

// Global variables to hold River dependencies for the API server (v2)
var (
	globalComponents *RiverComponents
)

// setupRiverDependentServices initializes River components for API server use (v2)
func setupRiverDependentServices(ctx context.Context) error {
	components, err := initializeRiverComponents(ctx)
	if err != nil {
		return fmt.Errorf("failed to initialize River components (v2): %w", err)
	}

	// Register all workflows
	if err := registerAllWorkers(components); err != nil {
		return fmt.Errorf("failed to register workflows (v2): %w", err)
	}

	// Store globally for API server use
	globalComponents = components

	// Start River job processing for the API server
	if err := components.RiverClient.Start(ctx); err != nil {
		return fmt.Errorf("failed to start River client for API server: %w", err)
	}

	// Setup background monitoring jobs for the API server
	setupRiverBackgroundJobs(ctx, components)

	// Register workflow bridge for API controllers
	workflowBridge := workflowbridge.NewWorkflowBridge(components.WorkflowExecutor, components.RiverClient, components.WorkflowStore)
	service.RegisterWorkflowBridge(workflowBridge)

	// Register River client (if needed by other services)
	service.RegisterRiverClient(components.RiverClient)

	g.Log().Info(ctx, "River dependencies (v2) initialized for API server")
	return nil
}

// GetGlobalRiverClient returns the global River client for API server use (v2)
func GetGlobalRiverClient() *river.Client[pgx.Tx] {
	if globalComponents != nil && globalComponents.RiverClient != nil {
		return globalComponents.RiverClient
	}
	return nil
}

// CleanupRiverDependencies should be called during server shutdown (v2)
func CleanupRiverDependencies() {
	if globalComponents != nil {
		// Stop River client gracefully
		if globalComponents.RiverClient != nil {
			stopCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			if err := globalComponents.RiverClient.Stop(stopCtx); err != nil {
				g.Log().Errorf(context.Background(), "Error stopping River client during cleanup: %v", err)
			} else {
				g.Log().Info(context.Background(), "River client stopped successfully during cleanup")
			}
		}

		// Close database pool
		if globalComponents.DBPool != nil {
			globalComponents.DBPool.Close()
		}

		g.Log().Info(context.Background(), "River dependencies (v2) cleaned up")
	}
}
