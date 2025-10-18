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

	riverjobusersignup "v1consortium/internal/logic/riverjobUsersignup"
	"v1consortium/internal/pkg/riverjobs"
	"v1consortium/internal/service"
	signupworkflow "v1consortium/internal/workflow/signup"
)

var (
	RiverWorker = gcmd.Command{
		Name:  "river_worker",
		Usage: "river_worker",
		Brief: "start river job worker",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Info(ctx, "V1 Consortium River Worker starting...")

			// Initialize logger
			logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}))

			// Get database URL
			dbURL := getRiverDBURL(ctx)
			if dbURL == "" {
				g.Log().Fatal(ctx, "Database URL not found in environment variables")
				return fmt.Errorf("database URL not configured")
			}

			// Create database connection pool
			dbPool, err := pgxpool.New(ctx, dbURL)
			if err != nil {
				g.Log().Fatalf(ctx, "Failed to create database pool: %v", err)
				return err
			}
			defer dbPool.Close()

			// Test database connection
			if err := dbPool.Ping(ctx); err != nil {
				g.Log().Fatalf(ctx, "Failed to ping database: %v", err)
				return err
			}

			g.Log().Info(ctx, "Database connection established")

			// Load River configuration from config file
			riverConfig := loadRiverConfig(ctx, dbURL)

			// Initialize River manager with loaded config
			riverManager, err := riverjobs.NewRiverManager(ctx, riverConfig, logger)
			if err != nil {
				g.Log().Fatalf(ctx, "Failed to create River manager: %v", err)
				return err
			}

			// Initialize workflow manager
			workflowManager := riverjobs.NewWorkflowManager(dbPool, logger)

			// Initialize orchestrator
			orchestrator := riverjobs.NewWorkflowOrchestrator(riverManager, workflowManager, logger)

			g.Log().Info(ctx, "River components initialized")

			// Register workers
			if err := registerRiverWorkers(riverManager, workflowManager, logger); err != nil {
				g.Log().Fatalf(ctx, "Failed to register workers: %v", err)
				return err
			}

			g.Log().Info(ctx, "Workers registered successfully")

			// Start River job processing
			if err := riverManager.Start(ctx); err != nil {
				g.Log().Fatalf(ctx, "Failed to start River client: %v", err)
				return err
			}

			// Setup background monitoring jobs with configuration
			setupRiverBackgroundJobs(ctx, orchestrator, workflowManager, dbPool, riverConfig) // Register shutdown handler
			gproc.AddSigHandlerShutdown(func(sig os.Signal) {
				g.Log().Infof(ctx, "Received signal %s, shutting down gracefully...", sig)

				// Stop River client
				stopCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
				defer cancel()

				if err := riverManager.Stop(stopCtx); err != nil {
					g.Log().Errorf(ctx, "Error stopping River client: %v", err)
				}

				g.Log().Info(ctx, "River worker shutdown complete")
			})

			g.Log().Info(ctx, "River worker ready and processing jobs")

			// Block listening for the shutdown signal
			g.Listen()
			return nil
		},
	}
)

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

// registerRiverWorkers registers all workflow workers with River
func registerRiverWorkers(riverManager *riverjobs.RiverManager, workflowManager *riverjobs.WorkflowManager, logger *slog.Logger) error {
	// Initialize worker base for all workers
	workerBase := riverjobs.WorkerBase{
		WorkflowManager: workflowManager,
		RiverClient:     riverManager.Client,
		Logger:          logger,
	}

	// Create workflow orchestrator
	orchestrator := riverjobs.NewWorkflowOrchestrator(riverManager, workflowManager, logger)

	// Register workflow definitions
	signupWorkflow := signupworkflow.NewSignupWorkflow()
	orchestrator.RegisterWorkflow(signupWorkflow)
	logger.Info("Registered SignupWorkflow definition")

	// Register individual step workers for the signup workflow
	validateWorker := &signupworkflow.ValidateStepWorker{
		WorkflowManager: workflowManager,
		RiverClient:     riverManager.Client,
		Logger:          logger,
	}
	river.AddWorker[riverjobs.ValidateStepArgs](riverManager.Workers, validateWorker)
	logger.Info("Registered ValidateStepWorker")

	createUserWorker := &signupworkflow.CreateUserStepWorker{
		WorkflowManager: workflowManager,
		RiverClient:     riverManager.Client,
		Logger:          logger,
	}
	river.AddWorker[riverjobs.CreateUserStepArgs](riverManager.Workers, createUserWorker)
	logger.Info("Registered CreateUserStepWorker")

	createOrgStepWorker := &signupworkflow.CreateOrganizationStepWorker{
		WorkflowManager: workflowManager,
		RiverClient:     riverManager.Client,
		Logger:          logger,
	}
	river.AddWorker[riverjobs.CreateOrganizationStepArgs](riverManager.Workers, createOrgStepWorker)
	logger.Info("Registered CreateOrganizationStepWorker")

	setupStripeWorker := &signupworkflow.SetupStripeStepWorker{
		WorkflowManager: workflowManager,
		RiverClient:     riverManager.Client,
		Logger:          logger,
	}
	river.AddWorker[riverjobs.SetupStripeStepArgs](riverManager.Workers, setupStripeWorker)
	logger.Info("Registered SetupStripeStepWorker")

	sendVerificationWorker := &signupworkflow.SendVerificationStepWorker{
		WorkflowManager: workflowManager,
		RiverClient:     riverManager.Client,
		Logger:          logger,
	}
	river.AddWorker[riverjobs.SendVerificationStepArgs](riverManager.Workers, sendVerificationWorker)
	logger.Info("Registered SendVerificationStepWorker")

	// Keep the legacy UserSignupWorker for backward compatibility
	userSignupWorker := riverjobusersignup.NewUserSignupWorker(&workerBase, orchestrator)
	service.RegisterUserSignupWorker(userSignupWorker)
	river.AddWorker[riverjobs.UserSignupArgs](riverManager.Workers, service.UserSignupWorker())
	logger.Info("Registered UserSignupWorker (legacy)")

	// Register other existing workers
	createOrgWorker := &riverjobs.CreateOrganizationWorker{
		WorkerBase: workerBase,
	}
	river.AddWorker[riverjobs.CreateOrganizationArgs](riverManager.Workers, createOrgWorker)
	logger.Info("Registered CreateOrganizationWorker")

	processSubWorker := &riverjobs.ProcessSubscriptionWorker{
		WorkerBase: workerBase,
	}
	river.AddWorker[riverjobs.ProcessSubscriptionArgs](riverManager.Workers, processSubWorker)
	logger.Info("Registered ProcessSubscriptionWorker")

	// Drug Test Workflow Workers
	drugTestOrderWorker := &riverjobs.DrugTestOrderWorker{
		WorkerBase: workerBase,
	}
	river.AddWorker[riverjobs.DrugTestOrderArgs](riverManager.Workers, drugTestOrderWorker)
	logger.Info("Registered DrugTestOrderWorker")

	sendNotificationWorker := &riverjobs.SendTestNotificationWorker{
		WorkerBase: workerBase,
	}
	river.AddWorker[riverjobs.SendTestNotificationArgs](riverManager.Workers, sendNotificationWorker)
	logger.Info("Registered SendTestNotificationWorker")

	logger.Info("All workflow workers registered successfully")
	return nil
}

// setupRiverBackgroundJobs configures monitoring and maintenance jobs
func setupRiverBackgroundJobs(ctx context.Context, orchestrator *riverjobs.WorkflowOrchestrator, workflowManager *riverjobs.WorkflowManager, dbPool *pgxpool.Pool, config *riverjobs.Config) {
	// Monitor stuck workflows
	if config.BackgroundJobs.StuckWorkflowMonitor.Enabled {
		_, err := gcron.Add(ctx, config.BackgroundJobs.StuckWorkflowMonitor.CronExpression, func(ctx context.Context) {
			thresholdHours := config.BackgroundJobs.StuckWorkflowMonitor.StuckThresholdHours
			if thresholdHours == 0 {
				thresholdHours = 1 // Default to 1 hour
			}
			if err := monitorStuckWorkflows(ctx, dbPool, thresholdHours); err != nil {
				g.Log().Errorf(ctx, "Error monitoring stuck workflows: %v", err)
			}
		})
		if err != nil {
			g.Log().Warningf(ctx, "Failed to add stuck workflow monitoring cron: %v", err)
		} else {
			g.Log().Infof(ctx, "Scheduled stuck workflow monitor: %s", config.BackgroundJobs.StuckWorkflowMonitor.CronExpression)
		}
	}

	// Cleanup completed workflows
	if config.BackgroundJobs.WorkflowCleanup.Enabled {
		_, err := gcron.Add(ctx, config.BackgroundJobs.WorkflowCleanup.CronExpression, func(ctx context.Context) {
			retentionDays := config.BackgroundJobs.WorkflowCleanup.RetentionDays
			if retentionDays == 0 {
				retentionDays = 30 // Default to 30 days
			}
			cleaned, err := cleanupOldWorkflows(ctx, dbPool, retentionDays)
			if err != nil {
				g.Log().Errorf(ctx, "Error cleaning up workflows: %v", err)
			} else if cleaned > 0 {
				g.Log().Infof(ctx, "Cleaned up %d old workflows", cleaned)
			}
		})
		if err != nil {
			g.Log().Warningf(ctx, "Failed to add workflow cleanup cron: %v", err)
		} else {
			g.Log().Infof(ctx, "Scheduled workflow cleanup: %s (retention: %d days)",
				config.BackgroundJobs.WorkflowCleanup.CronExpression,
				config.BackgroundJobs.WorkflowCleanup.RetentionDays)
		}
	}

	// Collect workflow metrics
	if config.BackgroundJobs.MetricsCollection.Enabled {
		_, err := gcron.Add(ctx, config.BackgroundJobs.MetricsCollection.CronExpression, func(ctx context.Context) {
			if err := collectWorkflowMetrics(ctx, dbPool); err != nil {
				g.Log().Errorf(ctx, "Error collecting workflow metrics: %v", err)
			}
		})
		if err != nil {
			g.Log().Warningf(ctx, "Failed to add metrics collection cron: %v", err)
		} else {
			g.Log().Infof(ctx, "Scheduled metrics collection: %s", config.BackgroundJobs.MetricsCollection.CronExpression)
		}
	}

	// Health check
	if config.BackgroundJobs.HealthCheck.Enabled {
		_, err := gcron.Add(ctx, config.BackgroundJobs.HealthCheck.CronExpression, func(ctx context.Context) {
			g.Log().Debug(ctx, "River worker health check - running")
		})
		if err != nil {
			g.Log().Warningf(ctx, "Failed to add health check cron: %v", err)
		} else {
			g.Log().Infof(ctx, "Scheduled health check: %s", config.BackgroundJobs.HealthCheck.CronExpression)
		}
	}

	g.Log().Info(ctx, "River background monitoring jobs configured successfully")
}

// monitorStuckWorkflows identifies and handles workflows that appear to be stuck
func monitorStuckWorkflows(ctx context.Context, dbPool *pgxpool.Pool, thresholdHours int) error {
	// Query for workflows that have been running for more than the threshold
	query := `
		SELECT workflow_id, workflow_type, started_at, current_step
		FROM workflow_executions 
		WHERE status = $1 
		AND started_at < NOW() - INTERVAL '%d hour'
	`

	rows, err := dbPool.Query(ctx, fmt.Sprintf(query, thresholdHours), riverjobs.StatusRunning)
	if err != nil {
		return fmt.Errorf("failed to query stuck workflows: %w", err)
	}
	defer rows.Close()

	var stuckCount int
	for rows.Next() {
		var workflowID, workflowType string
		var currentStep *string
		var startedAt time.Time

		if err := rows.Scan(&workflowID, &workflowType, &startedAt, &currentStep); err != nil {
			g.Log().Errorf(ctx, "Error scanning stuck workflow row: %v", err)
			continue
		}

		stepName := "unknown"
		if currentStep != nil {
			stepName = *currentStep
		}

		g.Log().Warningf(ctx, "Detected stuck workflow: %s (type: %s, step: %s, started: %v)",
			workflowID, workflowType, stepName, startedAt)

		stuckCount++
	}

	if stuckCount > 0 {
		g.Log().Warningf(ctx, "Found %d stuck workflows (threshold: %d hours)", stuckCount, thresholdHours)
	}

	return nil
}

// cleanupOldWorkflows removes completed workflows older than the specified days
func cleanupOldWorkflows(ctx context.Context, dbPool *pgxpool.Pool, retentionDays int) (int, error) {
	// Delete completed workflows older than retention period
	query := `
		DELETE FROM workflow_executions 
		WHERE status IN ($1, $2) 
		AND completed_at < NOW() - INTERVAL '%d days'
	`

	result, err := dbPool.Exec(ctx,
		fmt.Sprintf(query, retentionDays),
		riverjobs.StatusCompleted,
		riverjobs.StatusFailed)
	if err != nil {
		return 0, fmt.Errorf("failed to cleanup old workflows: %w", err)
	}

	return int(result.RowsAffected()), nil
}

// collectWorkflowMetrics gathers and logs workflow statistics
func collectWorkflowMetrics(ctx context.Context, dbPool *pgxpool.Pool) error {
	query := `
		SELECT 
			status,
			COUNT(*) as count,
			AVG(EXTRACT(EPOCH FROM (COALESCE(completed_at, NOW()) - started_at))) as avg_duration_seconds
		FROM workflow_executions 
		WHERE created_at > NOW() - INTERVAL '24 hours'
		GROUP BY status
	`

	rows, err := dbPool.Query(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to collect workflow metrics: %w", err)
	}
	defer rows.Close()

	metrics := make(map[string]map[string]interface{})

	for rows.Next() {
		var status string
		var count int
		var avgDuration *float64

		if err := rows.Scan(&status, &count, &avgDuration); err != nil {
			g.Log().Errorf(ctx, "Error scanning metrics row: %v", err)
			continue
		}

		duration := 0.0
		if avgDuration != nil {
			duration = *avgDuration
		}

		metrics[status] = map[string]interface{}{
			"count":        count,
			"avg_duration": duration,
		}
	}

	g.Log().Infof(ctx, "Workflow metrics (24h): %+v", metrics)
	return nil
}

// loadRiverConfig loads River configuration from config file and environment
func loadRiverConfig(ctx context.Context, dbURL string) *riverjobs.Config {
	cfg := g.Cfg()

	// Load basic configuration with defaults
	config := &riverjobs.Config{
		DatabaseURL:  dbURL,
		PollInterval: 5 * time.Second,
	}

	// Load poll interval from config
	if pollInterval := cfg.MustGet(ctx, "river.workers.pollInterval", "5s").String(); pollInterval != "" {
		if duration, err := time.ParseDuration(pollInterval); err == nil {
			config.PollInterval = duration
		}
	}

	// Load queue configurations
	config.Queues = make(map[string]river.QueueConfig)

	// Set default queue configurations
	defaultQueues := map[string]int{
		river.QueueDefault:          100,
		riverjobs.QueueCritical:     50,
		riverjobs.QueueScheduled:    25,
		riverjobs.QueueNotification: 25,
		riverjobs.QueueExternal:     10,
	}

	// Apply defaults first
	for queueName, defaultWorkers := range defaultQueues {
		config.Queues[queueName] = river.QueueConfig{
			MaxWorkers: defaultWorkers,
		}
	}

	// Override with config file values if present
	if cfg.Available(ctx, "river.workers.queues") {
		queueConfigs := cfg.MustGet(ctx, "river.workers.queues").MapStrVar()

		for queueName, queueConfig := range queueConfigs {
			queueMap := queueConfig.Map()
			if maxWorkers, exists := queueMap["maxWorkers"]; exists {
				if workers := maxWorkers.(float64); workers > 0 {
					config.Queues[queueName] = river.QueueConfig{
						MaxWorkers: int(workers),
					}
				}
			}
		}
	}

	// Load background job configurations
	config.BackgroundJobs = riverjobs.BackgroundJobConfig{
		StuckWorkflowMonitor: riverjobs.CronJobConfig{
			Enabled:             cfg.MustGet(ctx, "river.backgroundJobs.stuckWorkflowMonitor.enabled", true).Bool(),
			CronExpression:      cfg.MustGet(ctx, "river.backgroundJobs.stuckWorkflowMonitor.cronExpression", "0 */5 * * * *").String(),
			StuckThresholdHours: cfg.MustGet(ctx, "river.backgroundJobs.stuckWorkflowMonitor.stuckThresholdHours", 1).Int(),
		},
		WorkflowCleanup: riverjobs.CronJobConfig{
			Enabled:        cfg.MustGet(ctx, "river.backgroundJobs.workflowCleanup.enabled", true).Bool(),
			CronExpression: cfg.MustGet(ctx, "river.backgroundJobs.workflowCleanup.cronExpression", "0 0 3 * * *").String(),
			RetentionDays:  cfg.MustGet(ctx, "river.backgroundJobs.workflowCleanup.retentionDays", 30).Int(),
		},
		MetricsCollection: riverjobs.CronJobConfig{
			Enabled:        cfg.MustGet(ctx, "river.backgroundJobs.metricsCollection.enabled", true).Bool(),
			CronExpression: cfg.MustGet(ctx, "river.backgroundJobs.metricsCollection.cronExpression", "0 */10 * * * *").String(),
		},
		HealthCheck: riverjobs.CronJobConfig{
			Enabled:        cfg.MustGet(ctx, "river.backgroundJobs.healthCheck.enabled", true).Bool(),
			CronExpression: cfg.MustGet(ctx, "river.backgroundJobs.healthCheck.cronExpression", "0 * * * * *").String(),
		},
	}

	// Load workflow configurations
	config.WorkflowDefaults = riverjobs.WorkflowConfig{
		DefaultTimeout: cfg.MustGet(ctx, "river.workflows.defaultTimeout", "5m").String(),
		MaxRetries:     cfg.MustGet(ctx, "river.workflows.maxRetries", 3).Int(),
	}

	g.Log().Infof(ctx, "Loaded River configuration: %d queues, poll interval: %v",
		len(config.Queues), config.PollInterval)

	return config
}

// Global variables to hold River dependencies for the API server
var (
	globalRiverManager    *riverjobs.RiverManager
	globalWorkflowManager *riverjobs.WorkflowManager
	globalOrchestrator    *riverjobs.WorkflowOrchestrator
	globalDBPool          *pgxpool.Pool
)

func setupRiverDependentServices(ctx context.Context) error {

	// Initialize logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	// Get database URL
	dbURL := getRiverDBURL(ctx)
	if dbURL == "" {
		g.Log().Fatal(ctx, "Database URL not found in environment variables")
		return fmt.Errorf("database URL not configured")
	}

	// Create database connection pool - keep it alive for the server lifetime
	dbPool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		g.Log().Fatalf(ctx, "Failed to create database pool: %v", err)
		return err
	}
	globalDBPool = dbPool // Store globally to keep it alive

	// Test database connection
	if err := dbPool.Ping(ctx); err != nil {
		g.Log().Fatalf(ctx, "Failed to ping database: %v", err)
		return err
	}

	g.Log().Info(ctx, "Database connection established")

	// Load River configuration from config file
	riverConfig := loadRiverConfig(ctx, dbURL)

	// Initialize River manager with loaded config
	riverManager, err := riverjobs.NewRiverManager(ctx, riverConfig, logger)
	if err != nil {
		g.Log().Fatalf(ctx, "Failed to create River manager: %v", err)
		return err
	}
	globalRiverManager = riverManager

	// Initialize workflow manager
	workflowManager := riverjobs.NewWorkflowManager(dbPool, logger)
	globalWorkflowManager = workflowManager

	workerBase := riverjobs.WorkerBase{
		WorkflowManager: workflowManager,
		RiverClient:     riverManager.Client,
		Logger:          logger,
	}

	// Initialize orchestrator
	orchestrator := riverjobs.NewWorkflowOrchestrator(riverManager, workflowManager, logger)
	globalOrchestrator = orchestrator

	// Register workflow definitions with orchestrator
	signupWorkflow := signupworkflow.NewSignupWorkflow()
	orchestrator.RegisterWorkflow(signupWorkflow)
	logger.Info("Registered SignupWorkflow definition in API server")

	// Register step workers with River client (required for enqueueing jobs)
	// These workers won't actually process jobs in the API server
	validateWorker := &signupworkflow.ValidateStepWorker{
		WorkflowManager: workflowManager,
		RiverClient:     riverManager.Client,
		Logger:          logger,
	}
	river.AddWorker[riverjobs.ValidateStepArgs](riverManager.Workers, validateWorker)

	createUserWorker := &signupworkflow.CreateUserStepWorker{
		WorkflowManager: workflowManager,
		RiverClient:     riverManager.Client,
		Logger:          logger,
	}
	river.AddWorker[riverjobs.CreateUserStepArgs](riverManager.Workers, createUserWorker)

	createOrgStepWorker := &signupworkflow.CreateOrganizationStepWorker{
		WorkflowManager: workflowManager,
		RiverClient:     riverManager.Client,
		Logger:          logger,
	}
	river.AddWorker[riverjobs.CreateOrganizationStepArgs](riverManager.Workers, createOrgStepWorker)

	setupStripeWorker := &signupworkflow.SetupStripeStepWorker{
		WorkflowManager: workflowManager,
		RiverClient:     riverManager.Client,
		Logger:          logger,
	}
	river.AddWorker[riverjobs.SetupStripeStepArgs](riverManager.Workers, setupStripeWorker)

	sendVerificationWorker := &signupworkflow.SendVerificationStepWorker{
		WorkflowManager: workflowManager,
		RiverClient:     riverManager.Client,
		Logger:          logger,
	}
	river.AddWorker[riverjobs.SendVerificationStepArgs](riverManager.Workers, sendVerificationWorker)

	logger.Info("Registered step workers in API server River client for enqueueing")

	// Create UserSignup worker for API server use (without starting worker process)
	userSignupWorker := riverjobusersignup.NewUserSignupWorker(&workerBase, orchestrator)

	// Register the worker so it can be used by the API controllers
	service.RegisterUserSignupWorker(userSignupWorker)

	// Register the River client for API server use
	service.RegisterRiverClient(riverManager.Client)

	g.Log().Info(ctx, "River dependencies initialized for API server")
	return nil
}

// GetGlobalRiverClient returns the global River client for API server use
func GetGlobalRiverClient() *river.Client[pgx.Tx] {
	if globalRiverManager != nil {
		return globalRiverManager.Client
	}
	return nil
}

// CleanupRiverDependencies should be called during server shutdown
func CleanupRiverDependencies() {
	if globalDBPool != nil {
		globalDBPool.Close()
		g.Log().Info(context.Background(), "Database pool closed")
	}
}
