package riverjobs

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/riverqueue/river"
)

// WorkflowOrchestrator orchestrates multi-step workflows
type WorkflowOrchestrator struct {
	RiverManager    *RiverManager
	WorkflowManager *WorkflowManager
	ClientManager   *ClientManager
	Logger          *slog.Logger
}

// NewWorkflowOrchestrator creates a new workflow orchestrator
func NewWorkflowOrchestrator(riverManager *RiverManager, workflowManager *WorkflowManager, logger *slog.Logger) *WorkflowOrchestrator {
	clientManager := NewClientManager(riverManager.Client, logger)

	return &WorkflowOrchestrator{
		RiverManager:    riverManager,
		WorkflowManager: workflowManager,
		ClientManager:   clientManager,
		Logger:          logger,
	}
}

// StartWorkflow initiates a new workflow
func (o *WorkflowOrchestrator) StartWorkflow(ctx context.Context, workflowType, workflowID string, initialArgs JobArgs, protocol string) error {
	// Extract base args for workflow execution
	baseArgs, ok := initialArgs.(BaseJobArgs)
	if !ok {
		return fmt.Errorf("initial args must embed BaseJobArgs")
	}

	// Create workflow execution record
	execution := &WorkflowExecution{
		WorkflowType: workflowType,
		WorkflowID:   workflowID,
		OrgID:        baseArgs.OrgID,
		Status:       StatusPending,
		TotalSteps:   1, // Will be updated as workflow progresses
		Context:      make(map[string]interface{}),
		Metadata:     make(map[string]interface{}),
		StartedAt:    time.Now(),
	}

	if baseArgs.UserID != "" {
		execution.UserID = &baseArgs.UserID
	}

	err := o.WorkflowManager.CreateWorkflowExecution(ctx, execution)
	if err != nil {
		return fmt.Errorf("failed to create workflow execution: %w", err)
	}

	// Get appropriate client and enqueue first job
	client := o.ClientManager.GetClient(protocol)
	_, err = client.Insert(ctx, initialArgs, &river.InsertOpts{
		Priority: 1,
		Queue:    river.QueueDefault,
	})

	if err != nil {
		// Update workflow status to failed
		errorMsg := err.Error()
		o.WorkflowManager.UpdateWorkflowStatus(ctx, workflowID, StatusFailed, &errorMsg)
		return fmt.Errorf("failed to enqueue initial job: %w", err)
	}

	// Update workflow status to running
	o.WorkflowManager.UpdateWorkflowStatus(ctx, workflowID, StatusRunning, nil)

	o.Logger.Info("workflow started", "workflow_id", workflowID, "type", workflowType, "protocol", protocol)
	return nil
}

// GetWorkflowStatus retrieves the current status of a workflow
func (o *WorkflowOrchestrator) GetWorkflowStatus(ctx context.Context, workflowID string) (*WorkflowExecution, error) {
	return o.WorkflowManager.GetWorkflowExecution(ctx, workflowID)
}

// GetWorkflowSteps retrieves all steps for a workflow
func (o *WorkflowOrchestrator) GetWorkflowSteps(ctx context.Context, workflowID string) ([]WorkflowStep, error) {
	return o.WorkflowManager.GetWorkflowSteps(ctx, workflowID)
}

// CompleteWorkflow marks a workflow as completed
func (o *WorkflowOrchestrator) CompleteWorkflow(ctx context.Context, workflowID string) error {
	err := o.WorkflowManager.CompleteWorkflow(ctx, workflowID)
	if err != nil {
		return fmt.Errorf("failed to complete workflow: %w", err)
	}

	o.Logger.Info("workflow completed", "workflow_id", workflowID)
	return nil
}

// FailWorkflow marks a workflow as failed
func (o *WorkflowOrchestrator) FailWorkflow(ctx context.Context, workflowID string, errorMsg string) error {
	err := o.WorkflowManager.FailWorkflow(ctx, workflowID, errorMsg)
	if err != nil {
		return fmt.Errorf("failed to fail workflow: %w", err)
	}

	o.Logger.Error("workflow failed", "workflow_id", workflowID, "error", errorMsg)
	return nil
}

// CancelWorkflow marks a workflow as cancelled
func (o *WorkflowOrchestrator) CancelWorkflow(ctx context.Context, workflowID string) error {
	err := o.WorkflowManager.CancelWorkflow(ctx, workflowID)
	if err != nil {
		return fmt.Errorf("failed to cancel workflow: %w", err)
	}

	o.Logger.Info("workflow cancelled", "workflow_id", workflowID)
	return nil
}

// RestartWorkflow restarts a failed or cancelled workflow
func (o *WorkflowOrchestrator) RestartWorkflow(ctx context.Context, workflowID string, fromStep string) error {
	// Get current workflow execution
	execution, err := o.WorkflowManager.GetWorkflowExecution(ctx, workflowID)
	if err != nil {
		return fmt.Errorf("failed to get workflow execution: %w", err)
	}

	// Only allow restart for failed or cancelled workflows
	if execution.Status != StatusFailed && execution.Status != StatusCancelled {
		return fmt.Errorf("workflow can only be restarted from failed or cancelled state, current status: %s", execution.Status)
	}

	// Update workflow status to running
	err = o.WorkflowManager.UpdateWorkflowStatus(ctx, workflowID, StatusRunning, nil)
	if err != nil {
		return fmt.Errorf("failed to update workflow status: %w", err)
	}

	// TODO: Implement logic to restart from specific step
	// This would involve re-enqueuing the appropriate job

	o.Logger.Info("workflow restarted", "workflow_id", workflowID, "from_step", fromStep)
	return nil
}

// SetRPCClient sets the RPC client for the orchestrator
func (o *WorkflowOrchestrator) SetRPCClient(client RPCJobClient) {
	o.ClientManager.SetRPCClient(client)
}

// GetWorkflowProgress calculates the progress of a workflow
func (o *WorkflowOrchestrator) GetWorkflowProgress(ctx context.Context, workflowID string) (float64, error) {
	steps, err := o.WorkflowManager.GetWorkflowSteps(ctx, workflowID)
	if err != nil {
		return 0, fmt.Errorf("failed to get workflow steps: %w", err)
	}

	if len(steps) == 0 {
		return 0, nil
	}

	completedSteps := 0
	for _, step := range steps {
		if step.Status == StepStatusCompleted {
			completedSteps++
		}
	}

	return float64(completedSteps) / float64(len(steps)), nil
}

// GetWorkflowSummary returns a comprehensive summary of a workflow
func (o *WorkflowOrchestrator) GetWorkflowSummary(ctx context.Context, workflowID string) (*WorkflowSummary, error) {
	execution, err := o.WorkflowManager.GetWorkflowExecution(ctx, workflowID)
	if err != nil {
		return nil, fmt.Errorf("failed to get workflow execution: %w", err)
	}

	steps, err := o.WorkflowManager.GetWorkflowSteps(ctx, workflowID)
	if err != nil {
		return nil, fmt.Errorf("failed to get workflow steps: %w", err)
	}

	progress, _ := o.GetWorkflowProgress(ctx, workflowID)

	summary := &WorkflowSummary{
		Execution: *execution,
		Steps:     steps,
		Progress:  progress,
	}

	return summary, nil
}

// WorkflowSummary provides a complete view of a workflow
type WorkflowSummary struct {
	Execution WorkflowExecution `json:"execution"`
	Steps     []WorkflowStep    `json:"steps"`
	Progress  float64           `json:"progress"`
}
