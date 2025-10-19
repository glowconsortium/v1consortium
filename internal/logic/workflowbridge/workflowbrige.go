package workflowbridge

import (
	"context"
	"fmt"
	"v1consortium/internal/pkg/riverjobsv2"
	"v1consortium/internal/service"
)

type sWorkflowBridge struct {
	executor    *riverjobsv2.WorkflowExecutor
	riverClient interface{}               // River client for starting workflows
	store       riverjobsv2.WorkflowStore // Store for workflow persistence and deduplication
}

func NewWorkflowBridge(executor *riverjobsv2.WorkflowExecutor, riverClient interface{}, store riverjobsv2.WorkflowStore) service.IWorkflowBridge {
	return &sWorkflowBridge{
		executor:    executor,
		riverClient: riverClient,
		store:       store,
	}
}

// StartWorkflow starts a new workflow instance with deduplication
func (wb *sWorkflowBridge) StartWorkflow(ctx context.Context, workflowName string, input map[string]interface{}, orgID, userID string) (string, error) {
	// Add org and user IDs to input for workflow context
	if orgID != "" {
		input["org_id"] = orgID
	}
	if userID != "" {
		input["user_id"] = userID
	}

	// Start the workflow using riverjobsv2 with deduplication and UUID generation
	result, err := riverjobsv2.StartWorkflow(ctx, wb.riverClient, wb.store, workflowName, input)
	if err != nil {
		return "", fmt.Errorf("failed to start workflow %s: %w", workflowName, err)
	}

	return result.WorkflowID, nil
}

// StartUserSignupWorkflow starts the user signup workflow with deduplication
func (wb *sWorkflowBridge) StartUserSignupWorkflow(ctx context.Context, input map[string]interface{}, orgID, userID string) (string, error) {
	return wb.StartWorkflow(ctx, "signup", input, orgID, userID)
}

// StartDriverOnboardingWorkflow starts the driver onboarding workflow with deduplication
func (wb *sWorkflowBridge) StartDriverOnboardingWorkflow(ctx context.Context, input map[string]interface{}, orgID, userID string) (string, error) {
	return wb.StartWorkflow(ctx, "driver_onboarding", input, orgID, userID)
}

// StartComplianceCheckWorkflow starts the compliance check workflow with deduplication
func (wb *sWorkflowBridge) StartComplianceCheckWorkflow(ctx context.Context, input map[string]interface{}, orgID, userID string) (string, error) {
	return wb.StartWorkflow(ctx, "compliance_check", input, orgID, userID)
}

// GetWorkflowStatus gets the current status of a workflow using the workflow store
func (wb *sWorkflowBridge) GetWorkflowStatus(ctx context.Context, workflowID string) (*service.WorkflowExecution, error) {
	// Get workflow status from store
	status, err := riverjobsv2.GetWorkflowStatus(ctx, wb.store, workflowID)
	if err != nil {
		return nil, fmt.Errorf("failed to get workflow status: %w", err)
	}

	// Convert riverjobsv2.WorkflowStatus to service.WorkflowExecution
	execution := &service.WorkflowExecution{
		ID:           status.WorkflowID,
		WorkflowID:   status.WorkflowID,
		WorkflowType: status.WorkflowName,
		Status:       status.Status,
		Context:      status.Context,
		CreatedAt:    status.StartedAt,
		ErrorMessage: nil,
	}

	// Set current step if available
	if status.CurrentStep != "" {
		execution.CurrentStep = &status.CurrentStep
	}

	// Set completion time if workflow is completed
	if status.CompletedAt != nil {
		execution.CompletedAt = status.CompletedAt
	}

	// Set error message if workflow failed
	if status.ErrorMessage != "" {
		execution.ErrorMessage = &status.ErrorMessage
	}

	return execution, nil
}
