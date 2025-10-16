package riverjobs

import (
	"time"

	"github.com/riverqueue/river"
)

// JobArgs interface that all workflow jobs must implement
type JobArgs interface {
	river.JobArgs
	GetWorkflowID() string
	GetWorkflowType() string
	GetStepName() string
}

// BaseJobArgs provides common workflow job functionality
type BaseJobArgs struct {
	WorkflowID   string `json:"workflow_id"`
	WorkflowType string `json:"workflow_type"`
	StepName     string `json:"step_name"`
	OrgID        string `json:"org_id"`
	UserID       string `json:"user_id,omitempty"`
}

func (b BaseJobArgs) GetWorkflowID() string   { return b.WorkflowID }
func (b BaseJobArgs) GetWorkflowType() string { return b.WorkflowType }
func (b BaseJobArgs) GetStepName() string     { return b.StepName }
func (b BaseJobArgs) Kind() string            { return b.WorkflowType + "_" + b.StepName }

// WorkflowExecution represents a workflow execution state
type WorkflowExecution struct {
	ID           string                 `json:"id" db:"workflow_id"`
	WorkflowType string                 `json:"workflow_type" db:"workflow_type"`
	WorkflowID   string                 `json:"workflow_id" db:"workflow_id"`
	OrgID        string                 `json:"org_id" db:"org_id"`
	UserID       *string                `json:"user_id" db:"user_id"`
	Status       string                 `json:"status" db:"status"`
	CurrentStep  *string                `json:"current_step" db:"current_step"`
	TotalSteps   int                    `json:"total_steps" db:"total_steps"`
	Context      map[string]interface{} `json:"context" db:"context"`
	Metadata     map[string]interface{} `json:"metadata" db:"metadata"`
	StartedAt    time.Time              `json:"started_at" db:"started_at"`
	CompletedAt  *time.Time             `json:"completed_at" db:"completed_at"`
	ErrorMessage *string                `json:"error_message" db:"error_message"`
	RetryCount   int                    `json:"retry_count" db:"retry_count"`
	CreatedAt    time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at" db:"updated_at"`
}

// WorkflowStep represents an individual step in a workflow
type WorkflowStep struct {
	ID           string                 `json:"step_id" db:"step_id"`
	WorkflowID   string                 `json:"workflow_id" db:"workflow_id"`
	StepName     string                 `json:"step_name" db:"step_name"`
	StepOrder    int                    `json:"step_order" db:"step_order"`
	RiverJobID   *int64                 `json:"river_job_id" db:"river_job_id"`
	QueueName    *string                `json:"queue_name" db:"queue_name"`
	Status       string                 `json:"status" db:"status"`
	InputData    map[string]interface{} `json:"input_data" db:"input_data"`
	OutputData   map[string]interface{} `json:"output_data" db:"output_data"`
	StartedAt    *time.Time             `json:"started_at" db:"started_at"`
	CompletedAt  *time.Time             `json:"completed_at" db:"completed_at"`
	CreatedAt    time.Time              `json:"created_at" db:"created_at"`
	ErrorMessage *string                `json:"error_message" db:"error_message"`
	RetryCount   int                    `json:"retry_count" db:"retry_count"`
	MaxRetries   int                    `json:"max_retries" db:"max_retries"`
}

// Config holds configuration for River
type Config struct {
	DatabaseURL      string
	Queues           map[string]river.QueueConfig
	PollInterval     time.Duration
	BackgroundJobs   BackgroundJobConfig
	WorkflowDefaults WorkflowConfig
}

// BackgroundJobConfig holds configuration for background jobs
type BackgroundJobConfig struct {
	StuckWorkflowMonitor CronJobConfig
	WorkflowCleanup      CronJobConfig
	MetricsCollection    CronJobConfig
	HealthCheck          CronJobConfig
}

// CronJobConfig holds configuration for a single cron job
type CronJobConfig struct {
	Enabled             bool
	CronExpression      string
	StuckThresholdHours int // For stuck workflow monitor
	RetentionDays       int // For workflow cleanup
}

// WorkflowConfig holds default workflow configuration
type WorkflowConfig struct {
	DefaultTimeout string
	MaxRetries     int
}

// Common queue names
const (
	QueueDefault      = river.QueueDefault
	QueueCritical     = "critical"
	QueueScheduled    = "scheduled"
	QueueNotification = "notification"
	QueueExternal     = "external"
)

// Workflow status constants
const (
	StatusPending   = "pending"
	StatusRunning   = "running"
	StatusCompleted = "completed"
	StatusFailed    = "failed"
	StatusCancelled = "cancelled"
)

// Step status constants
const (
	StepStatusPending   = "pending"
	StepStatusRunning   = "running"
	StepStatusCompleted = "completed"
	StepStatusFailed    = "failed"
	StepStatusSkipped   = "skipped"
)
