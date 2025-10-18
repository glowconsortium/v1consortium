package riverjobs

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/rivertype"
)

// WorkerBase provides common functionality for all workflow workers
type WorkerBase struct {
	WorkflowManager *WorkflowManager
	RiverClient     *river.Client[pgx.Tx]
	Logger          *slog.Logger
}

// Middleware implements the River Worker interface requirement
func (w *WorkerBase) Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware {
	return []rivertype.WorkerMiddleware{}
}

// NextRetry implements the River Worker interface requirement
func (w *WorkerBase) NextRetry(job *rivertype.JobRow) time.Time {
	// Default exponential backoff retry strategy
	return time.Now().Add(time.Duration(job.Attempt) * time.Minute)
}

// Timeout implements the River Worker interface requirement
func (w *WorkerBase) Timeout(job *rivertype.JobRow) time.Duration {
	// Default timeout of 5 minutes
	return 5 * time.Minute
}

// NewWorkerBase creates a new worker base
func NewWorkerBase(wm *WorkflowManager, client *river.Client[pgx.Tx], logger *slog.Logger) *WorkerBase {
	return &WorkerBase{
		WorkflowManager: wm,
		RiverClient:     client,
		Logger:          logger,
	}
}

// EnqueueNextJob enqueues the next job in the workflow
func (w *WorkerBase) EnqueueNextJob(ctx context.Context, args JobArgs, opts *river.InsertOpts) (*rivertype.JobInsertResult, error) {
	if opts == nil {
		opts = &river.InsertOpts{
			Priority: 1,
			Queue:    river.QueueDefault,
		}
	}

	// Create workflow step record
	stepID, err := w.WorkflowManager.CreateWorkflowStep(ctx, args.GetWorkflowID(), args.GetStepName(), args)
	if err != nil {
		return nil, fmt.Errorf("failed to create workflow step: %w", err)
	}

	// Insert job
	result, err := w.RiverClient.Insert(ctx, args, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to insert job: %w", err)
	}

	// Update step with job ID
	err = w.WorkflowManager.UpdateWorkflowStepJobID(ctx, stepID, result.Job.ID)
	if err != nil {
		w.Logger.Warn("failed to update step job ID", "error", err, "step_id", stepID, "job_id", result.Job.ID)
	}

	return result, nil
}

// EnqueueNextJobTx enqueues the next job in a transaction
func (w *WorkerBase) EnqueueNextJobTx(ctx context.Context, tx pgx.Tx, args JobArgs, opts *river.InsertOpts) (*rivertype.JobInsertResult, error) {
	if opts == nil {
		opts = &river.InsertOpts{
			Priority: 1,
			Queue:    river.QueueDefault,
		}
	}

	// Insert job in transaction
	result, err := w.RiverClient.InsertTx(ctx, tx, args, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to insert job in transaction: %w", err)
	}

	return result, nil
}

// UpdateWorkflowStep updates the status of a workflow step
func (w *WorkerBase) UpdateWorkflowStep(ctx context.Context, workflowID, stepName, status string) error {
	return w.WorkflowManager.UpdateWorkflowStep(ctx, workflowID, stepName, status)
}

// UpdateWorkflowContext updates the workflow context with new data
func (w *WorkerBase) UpdateWorkflowContext(ctx context.Context, workflowID string, contextUpdate map[string]interface{}) error {
	return w.WorkflowManager.UpdateWorkflowContext(ctx, workflowID, contextUpdate)
}

// GetWorkflowExecution retrieves the current workflow execution
func (w *WorkerBase) GetWorkflowExecution(ctx context.Context, workflowID string) (*WorkflowExecution, error) {
	return w.WorkflowManager.GetWorkflowExecution(ctx, workflowID)
}

// CompleteWorkflowStep marks a workflow step as completed
func (w *WorkerBase) CompleteWorkflowStep(ctx context.Context, workflowID, stepName string, outputData map[string]interface{}) error {
	// Update step status to completed
	err := w.WorkflowManager.UpdateWorkflowStep(ctx, workflowID, stepName, StepStatusCompleted)
	if err != nil {
		return fmt.Errorf("failed to update step status: %w", err)
	}

	// Update output data if provided
	if outputData != nil {
		// TODO: Add method to update step output data in WorkflowManager
		w.Logger.Info("step completed with output data", "workflow_id", workflowID, "step", stepName, "output", outputData)
	}

	return nil
}

// FailWorkflowStep marks a workflow step as failed
func (w *WorkerBase) FailWorkflowStep(ctx context.Context, workflowID, stepName string, errorMsg string) error {
	err := w.WorkflowManager.UpdateWorkflowStep(ctx, workflowID, stepName, StepStatusFailed)
	if err != nil {
		return fmt.Errorf("failed to update step status: %w", err)
	}

	w.Logger.Error("step failed", "workflow_id", workflowID, "step", stepName, "error", errorMsg)
	return nil
}

// StartWorkflowStep marks a workflow step as running
func (w *WorkerBase) StartWorkflowStep(ctx context.Context, workflowID, stepName string) error {
	return w.WorkflowManager.UpdateWorkflowStep(ctx, workflowID, stepName, StepStatusRunning)
}

// ExecuteWithWorkflowTracking executes a function with automatic workflow step tracking
func (w *WorkerBase) ExecuteWithWorkflowTracking(
	ctx context.Context,
	workflowID, stepName string,
	fn func(ctx context.Context) (map[string]interface{}, error),
) error {
	// Mark step as running
	if err := w.StartWorkflowStep(ctx, workflowID, stepName); err != nil {
		return fmt.Errorf("failed to start workflow step: %w", err)
	}

	// Execute the function
	outputData, err := fn(ctx)
	if err != nil {
		// Mark step as failed
		if failErr := w.FailWorkflowStep(ctx, workflowID, stepName, err.Error()); failErr != nil {
			w.Logger.Error("failed to mark step as failed", "error", failErr)
		}
		return err
	}

	// Mark step as completed
	if err := w.CompleteWorkflowStep(ctx, workflowID, stepName, outputData); err != nil {
		return fmt.Errorf("failed to complete workflow step: %w", err)
	}

	return nil
}

// ShouldRetry determines if a job should be retried based on the error and attempt count
func (w *WorkerBase) ShouldRetry(err error, attempt int, maxAttempts int) bool {
	if attempt >= maxAttempts {
		return false
	}

	// Add custom retry logic here based on error types
	// For now, retry most errors except for specific cases

	return true
}

// LogJobStart logs the start of a job execution
func (w *WorkerBase) LogJobStart(ctx context.Context, workflowID, stepName string, args interface{}) {
	w.Logger.Info("job started",
		"workflow_id", workflowID,
		"step", stepName,
		"args", args,
	)
}

// LogJobComplete logs the completion of a job execution
func (w *WorkerBase) LogJobComplete(ctx context.Context, workflowID, stepName string, duration string) {
	w.Logger.Info("job completed",
		"workflow_id", workflowID,
		"step", stepName,
		"duration", duration,
	)
}

// LogJobError logs an error during job execution
func (w *WorkerBase) LogJobError(ctx context.Context, workflowID, stepName string, err error, attempt int) {
	w.Logger.Error("job error",
		"workflow_id", workflowID,
		"step", stepName,
		"error", err.Error(),
		"attempt", attempt,
	)
}
