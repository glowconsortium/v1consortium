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

	"v1consortium/internal/logic/workflowbridge"
	"v1consortium/internal/pkg/riverjobs"
	"v1consortium/internal/service"
	signupworkflow "v1consortium/internal/workflow/signup"
	signupsteps "v1consortium/internal/workflow/signup/steps"
)

type RiverComponents struct {
	RiverManager          *riverjobs.RiverManager
	SimpleWorkflowManager *riverjobs.SimpleWorkflowManager
	DBPool                *pgxpool.Pool
	Logger                *slog.Logger
}

var (
	RiverWorker = gcmd.Command{
		Name:  "river_worker",
		Usage: "river_worker",
		Brief: "start river job worker",
		Func: func(ctx context.Context, parser *gcmd.Parser) error {
			g.Log().Info(ctx, "V1 Consortium River Worker starting...")

			components, err := initializeRiverComponents(ctx)
			if err != nil {
				return err
			}
			defer components.DBPool.Close()

			// Register workers
			if err := registerAllWorkers(components); err != nil {
				g.Log().Fatalf(ctx, "Failed to register workers: %v", err)
				return err
			}

			// Start River job processing
			if err := components.RiverManager.Start(ctx); err != nil {
				g.Log().Fatalf(ctx, "Failed to start River client: %v", err)
				return err
			}

			// Setup background monitoring jobs
			config := loadRiverConfig(ctx, getRiverDBURL(ctx))
			setupRiverBackgroundJobs(ctx, components.SimpleWorkflowManager, components.DBPool, config)

			// Setup graceful shutdown
			setupShutdownHandler(ctx, components.RiverManager)

			g.Log().Info(ctx, "River worker ready and processing jobs")
			g.Listen()
			return nil
		},
	}
)

// initializeRiverComponents sets up all River-related components
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

	// Load River configuration
	riverConfig := loadRiverConfig(ctx, dbURL)

	// Initialize River manager
	riverManager, err := riverjobs.NewRiverManager(ctx, riverConfig, logger)
	if err != nil {
		dbPool.Close()
		return nil, fmt.Errorf("failed to create River manager: %w", err)
	}

	// Initialize simple workflow manager
	simpleWorkflowManager := riverjobs.NewSimpleWorkflowManager(riverManager, dbPool)

	g.Log().Info(ctx, "River components initialized")

	return &RiverComponents{
		RiverManager:          riverManager,
		SimpleWorkflowManager: simpleWorkflowManager,
		DBPool:                dbPool,
		Logger:                logger,
	}, nil
}

// registerAllWorkers registers all workflow workers with River
func registerAllWorkers(components *RiverComponents) error {
	// Register workflow definition
	signupWorkflow := signupworkflow.NewSignupWorkflowDefinition()
	components.SimpleWorkflowManager.RegisterWorkflow(signupWorkflow)
	components.Logger.Info("Registered SignupWorkflow definition")

	// Register signup step workers (using StepArgs)
	// Validate step: register adapter with concrete arg type so kind is unique
	validateWorker := signupsteps.NewValidateStepWorker(components.SimpleWorkflowManager)
	validateAdapter := &riverjobs.ValidateAdapter{Impl: validateWorker, Wm: components.SimpleWorkflowManager}
	river.AddWorker[riverjobs.ValidateStepArgs](components.RiverManager.Workers, validateAdapter)
	components.Logger.Info("Registered ValidateStepWorker")

	// Create user
	createUserWorker := signupsteps.NewCreateUserStepWorker(components.SimpleWorkflowManager)
	createUserAdapter := &riverjobs.CreateUserAdapter{Impl: createUserWorker, Wm: components.SimpleWorkflowManager}
	river.AddWorker[riverjobs.CreateUserStepArgs](components.RiverManager.Workers, createUserAdapter)
	components.Logger.Info("Registered CreateUserStepWorker")

	// Create organization
	createOrgWorker := signupsteps.NewCreateOrganizationStepWorker(components.SimpleWorkflowManager)
	createOrgAdapter := &riverjobs.CreateOrgAdapter{Impl: createOrgWorker, Wm: components.SimpleWorkflowManager}
	river.AddWorker[riverjobs.CreateOrganizationStepArgs](components.RiverManager.Workers, createOrgAdapter)
	components.Logger.Info("Registered CreateOrganizationStepWorker")

	// Stripe setup
	setupStripeWorker := signupsteps.NewSetupStripeStepWorker(components.SimpleWorkflowManager)
	setupStripeAdapter := &riverjobs.SetupStripeAdapter{Impl: setupStripeWorker, Wm: components.SimpleWorkflowManager}
	river.AddWorker[riverjobs.SetupStripeStepArgs](components.RiverManager.Workers, setupStripeAdapter)
	components.Logger.Info("Registered SetupStripeStepWorker")

	// Send verification
	sendVerificationWorker := signupsteps.NewSendVerificationStepWorker(components.SimpleWorkflowManager)
	sendVerificationAdapter := &riverjobs.SendVerificationAdapter{Impl: sendVerificationWorker, Wm: components.SimpleWorkflowManager}
	river.AddWorker[riverjobs.SendVerificationStepArgs](components.RiverManager.Workers, sendVerificationAdapter)
	components.Logger.Info("Registered SendVerificationStepWorker")

	components.Logger.Info("All workflow workers registered successfully")
	return nil
}

// registerOtherWorkers registers other existing workers
func registerOtherWorkers(components *RiverComponents, workerBase riverjobs.WorkerBase) {
	workers := []struct {
		name   string
		worker interface{}
	}{
		{"CreateOrganizationWorker", &riverjobs.CreateOrganizationWorker{WorkerBase: workerBase}},
		{"ProcessSubscriptionWorker", &riverjobs.ProcessSubscriptionWorker{WorkerBase: workerBase}},
		{"DrugTestOrderWorker", &riverjobs.DrugTestOrderWorker{WorkerBase: workerBase}},
		{"SendTestNotificationWorker", &riverjobs.SendTestNotificationWorker{WorkerBase: workerBase}},
	}

	for _, worker := range workers {
		switch w := worker.worker.(type) {
		case *riverjobs.CreateOrganizationWorker:
			river.AddWorker[riverjobs.CreateOrganizationArgs](components.RiverManager.Workers, w)
		case *riverjobs.ProcessSubscriptionWorker:
			river.AddWorker[riverjobs.ProcessSubscriptionArgs](components.RiverManager.Workers, w)
		case *riverjobs.DrugTestOrderWorker:
			river.AddWorker[riverjobs.DrugTestOrderArgs](components.RiverManager.Workers, w)
		case *riverjobs.SendTestNotificationWorker:
			river.AddWorker[riverjobs.SendTestNotificationArgs](components.RiverManager.Workers, w)
		}
		components.Logger.Info("Registered " + worker.name)
	}
}

// setupShutdownHandler configures graceful shutdown
func setupShutdownHandler(ctx context.Context, riverManager *riverjobs.RiverManager) {
	gproc.AddSigHandlerShutdown(func(sig os.Signal) {
		g.Log().Infof(ctx, "Received signal %s, shutting down gracefully...", sig)

		stopCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := riverManager.Stop(stopCtx); err != nil {
			g.Log().Errorf(ctx, "Error stopping River client: %v", err)
		}

		g.Log().Info(ctx, "River worker shutdown complete")
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

// setupRiverBackgroundJobs configures monitoring and maintenance jobs
func setupRiverBackgroundJobs(ctx context.Context, simpleWorkflowManager *riverjobs.SimpleWorkflowManager, dbPool *pgxpool.Pool, config *riverjobs.Config) {
	type backgroundJob struct {
		name     string
		enabled  bool
		cronExpr string
		jobFunc  func(context.Context)
	}

	// Define all background jobs
	jobs := []backgroundJob{
		{
			name:     "stuck workflow monitor",
			enabled:  config.BackgroundJobs.StuckWorkflowMonitor.Enabled,
			cronExpr: config.BackgroundJobs.StuckWorkflowMonitor.CronExpression,
			jobFunc: func(ctx context.Context) {
				threshold := config.BackgroundJobs.StuckWorkflowMonitor.StuckThresholdHours
				if threshold == 0 {
					threshold = 1
				}
				if err := monitorStuckWorkflows(ctx, dbPool, threshold); err != nil {
					g.Log().Errorf(ctx, "Error monitoring stuck workflows: %v", err)
				}
			},
		},
		{
			name:     "workflow cleanup",
			enabled:  config.BackgroundJobs.WorkflowCleanup.Enabled,
			cronExpr: config.BackgroundJobs.WorkflowCleanup.CronExpression,
			jobFunc: func(ctx context.Context) {
				retention := config.BackgroundJobs.WorkflowCleanup.RetentionDays
				if retention == 0 {
					retention = 30
				}
				cleaned, err := cleanupOldWorkflows(ctx, dbPool, retention)
				if err != nil {
					g.Log().Errorf(ctx, "Error cleaning up workflows: %v", err)
				} else if cleaned > 0 {
					g.Log().Infof(ctx, "Cleaned up %d old workflows", cleaned)
				}
			},
		},
		{
			name:     "metrics collection",
			enabled:  config.BackgroundJobs.MetricsCollection.Enabled,
			cronExpr: config.BackgroundJobs.MetricsCollection.CronExpression,
			jobFunc: func(ctx context.Context) {
				if err := collectWorkflowMetrics(ctx, dbPool); err != nil {
					g.Log().Errorf(ctx, "Error collecting workflow metrics: %v", err)
				}
			},
		},
		{
			name:     "health check",
			enabled:  config.BackgroundJobs.HealthCheck.Enabled,
			cronExpr: config.BackgroundJobs.HealthCheck.CronExpression,
			jobFunc: func(ctx context.Context) {
				g.Log().Debug(ctx, "River worker health check - running")
			},
		},
	}

	// Schedule enabled jobs
	for _, job := range jobs {
		if job.enabled {
			if _, err := gcron.Add(ctx, job.cronExpr, job.jobFunc); err != nil {
				g.Log().Warningf(ctx, "Failed to add %s cron: %v", job.name, err)
			} else {
				g.Log().Infof(ctx, "Scheduled %s: %s", job.name, job.cronExpr)
			}
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
	defaultTimeoutStr := cfg.MustGet(ctx, "river.workflows.defaultTimeout", "5m").String()
	defaultTimeout, err := time.ParseDuration(defaultTimeoutStr)
	if err != nil {
		g.Log().Warningf(ctx, "Invalid default timeout '%s', using 5 minutes", defaultTimeoutStr)
		defaultTimeout = 5 * time.Minute
	}

	config.WorkflowDefaults = riverjobs.WorkflowConfig{
		DefaultTimeout: defaultTimeout,
		MaxRetries:     cfg.MustGet(ctx, "river.workflows.maxRetries", 3).Int(),
	}

	g.Log().Infof(ctx, "Loaded River configuration: %d queues, poll interval: %v",
		len(config.Queues), config.PollInterval)

	return config
}

// Global variables to hold River dependencies for the API server
var (
	globalComponents *RiverComponents
)

// setupRiverDependentServices initializes River components for API server use
func setupRiverDependentServices(ctx context.Context) error {
	components, err := initializeRiverComponents(ctx)
	if err != nil {
		return err
	}

	// Register all workflows - IMPORTANT: This must happen before registering the bridge
	if err := registerAllWorkers(components); err != nil {
		return fmt.Errorf("failed to register workflows: %w", err)
	}

	// Store globally for API server use
	globalComponents = components

	// Register workflow service for API controllers
	workflowService := riverjobs.NewWorkflowService(components.SimpleWorkflowManager)
	// TODO: Create service.RegisterWorkflowService function
	_ = workflowService // Temporarily unused
	service.RegisterRiverClient(components.RiverManager.Client)

	workflowbrigde := workflowbridge.NewWorkflowBridge(components.SimpleWorkflowManager)
	service.RegisterWorkflowBridge(workflowbrigde)
	g.Log().Info(ctx, "River dependencies initialized for API server")
	return nil
}

// GetGlobalRiverClient returns the global River client for API server use
func GetGlobalRiverClient() *river.Client[pgx.Tx] {
	if globalComponents != nil && globalComponents.RiverManager != nil {
		return globalComponents.RiverManager.Client
	}
	return nil
}

// CleanupRiverDependencies should be called during server shutdown
func CleanupRiverDependencies() {
	if globalComponents != nil && globalComponents.DBPool != nil {
		globalComponents.DBPool.Close()
		g.Log().Info(context.Background(), "Database pool closed")
	}
}
