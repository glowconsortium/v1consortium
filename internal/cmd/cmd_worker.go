package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/dbos-inc/dbos-transact-golang/dbos"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gproc"
	_ "github.com/lib/pq" // PostgreSQL driver

	"v1consortium/internal/pkg/dbosworkflow"
)

var (
	DBOSWorker = gcmd.Command{
		Name:  "dbos_worker",
		Usage: "dbos_worker",
		Brief: "start dbos worker",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Info(ctx, "V1 Consortium DBOS Worker starting...")

			// Initialize database connection
			dbURL := getDBURL(ctx)

			if dbURL == "" {
				g.Log().Fatal(ctx, "Database URL not found in environment variables")
				return fmt.Errorf("database URL not configured")
			}

			// Connect to database
			db, err := sql.Open("postgres", dbURL)
			if err != nil {
				g.Log().Fatalf(ctx, "Failed to connect to database: %v", err)
				return err
			}
			defer db.Close()

			// Configure database connection pool
			db.SetMaxOpenConns(25)
			db.SetMaxIdleConns(5)
			db.SetConnMaxLifetime(5 * time.Minute)

			// Test database connection
			if err := db.Ping(); err != nil {
				g.Log().Fatalf(ctx, "Failed to ping database: %v", err)
				return err
			}

			g.Log().Info(ctx, "Database connection established")

			// Initialize DBOS context
			dbosContext, err := dbos.NewDBOSContext(context.Background(), dbos.Config{
				AppName:     "v1consortium-worker",
				DatabaseURL: dbURL,
			})
			if err != nil {
				g.Log().Warningf(ctx, "Initializing DBOS failed: %v", err)
				panic(err)
			}

			// Initialize workflow manager
			workflowManager := dbosworkflow.NewDBOSWorkflowManager(dbosContext, db)

			// Register workflows with DBOS
			if err := workflowManager.RegisterWorkflows(); err != nil {
				g.Log().Fatalf(ctx, "Failed to register workflows: %v", err)
				return err
			}

			// Initialize workflow executor for background processing
			workflowExecutor := dbosworkflow.NewWorkflowExecutor(db)

			g.Log().Info(ctx, "Workflow components initialized")

			// Launch DBOS runtime
			err = dbos.Launch(dbosContext)
			if err != nil {
				panic(fmt.Sprintf("Launching DBOS failed: %v", err))
			}
			defer func() {
				g.Log().Info(ctx, "Shutting down DBOS...")
				dbos.Shutdown(dbosContext, 10*time.Second)
			}()

			g.Log().Info(ctx, "DBOS runtime launched successfully")

			// Setup background jobs
			setupBackgroundJobs(ctx, workflowExecutor)

			// Register shutdown handler
			gproc.AddSigHandlerShutdown(func(sig os.Signal) {
				g.Log().Infof(ctx, "Received signal %s, shutting down gracefully...", sig)
			})

			g.Log().Info(ctx, "DBOS worker ready and processing workflows")

			// Block listening for the shutdown signal
			g.Listen()
			return nil
		},
	}
)

// get db url
func getDBURL(ctx context.Context) string {
	dbURL := os.Getenv("DBOS_SYSTEM_DATABASE_URL")
	if dbURL == "" {
		dbURL = os.Getenv("DATABASE_URL")
	}

	if dbURL == "" {
		dbURL = g.Cfg().MustGet(ctx, "dbos.systemDatabaseUrl").String()
	}

	if dbURL == "" {
		g.Log().Fatal(ctx, "Database URL not found in environment variables or config")
		panic("Database URL not configured")
	}
	return dbURL
}

// setupBackgroundJobs configures all background cron jobs
func setupBackgroundJobs(ctx context.Context, executor *dbosworkflow.WorkflowExecutor) {
	// Process pending workflows every 10 seconds
	_, err := gcron.Add(ctx, "*/10 * * * * *", func(ctx context.Context) {
		if err := executor.ProcessPendingWorkflows(ctx); err != nil {
			g.Log().Errorf(ctx, "Error processing pending workflows: %v", err)
		}
	})
	if err != nil {
		g.Log().Warningf(ctx, "Failed to add pending workflow processing cron: %v", err)
	}

	// Process failed workflows every 30 seconds
	_, err = gcron.Add(ctx, "*/30 * * * * *", func(ctx context.Context) {
		if err := executor.ProcessFailedWorkflows(ctx); err != nil {
			g.Log().Errorf(ctx, "Error processing failed workflows: %v", err)
		}
	})
	if err != nil {
		g.Log().Warningf(ctx, "Failed to add failed workflow processing cron: %v", err)
	}

	// Cleanup completed workflows daily at 2 AM
	_, err = gcron.Add(ctx, "0 0 2 * * *", func(ctx context.Context) {
		cleaned, err := executor.CleanupCompletedWorkflows(ctx, 30) // 30 days retention
		if err != nil {
			g.Log().Errorf(ctx, "Error cleaning up workflows: %v", err)
		} else if cleaned > 0 {
			g.Log().Infof(ctx, "Cleaned up %d old workflows", cleaned)
		}
	})
	if err != nil {
		g.Log().Warningf(ctx, "Failed to add workflow cleanup cron: %v", err)
	}

	// Collect and log metrics every 5 minutes
	_, err = gcron.Add(ctx, "0 */5 * * * *", func(ctx context.Context) {
		metrics, err := executor.GetWorkflowMetrics(ctx)
		if err != nil {
			g.Log().Errorf(ctx, "Error collecting workflow metrics: %v", err)
		} else {
			g.Log().Infof(ctx, "Workflow metrics: %+v", metrics)
		}
	})
	if err != nil {
		g.Log().Warningf(ctx, "Failed to add metrics collection cron: %v", err)
	}

	// Health check every 30 seconds
	_, err = gcron.Add(ctx, "*/30 * * * * *", func(ctx context.Context) {
		g.Log().Debug(ctx, "DBOS worker health check - running")
	})
	if err != nil {
		g.Log().Warningf(ctx, "Failed to add health check cron: %v", err)
	}

	g.Log().Info(ctx, "Background jobs configured successfully")
}
