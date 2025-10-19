package workflowbridge

import (
	"context"
	"v1consortium/internal/pkg/riverjobs"
	"v1consortium/internal/service"
)

type sWorkflowBridge struct {
	manager *riverjobs.SimpleWorkflowManager
}

func NewWorkflowBridge(manager *riverjobs.SimpleWorkflowManager) service.IWorkflowBridge {
	return &sWorkflowBridge{
		manager: manager,
	}
}

// StartWorkflow starts a new workflow instance
func (wb *sWorkflowBridge) StartWorkflow(ctx context.Context, workflowName string, input map[string]interface{}, orgID, userID string) (string, error) {
	return wb.manager.StartWorkflow(ctx, workflowName, input, orgID, userID)
}

// StartUserSignupWorkflow starts the user signup workflow with deduplication
func (wb *sWorkflowBridge) StartUserSignupWorkflow(ctx context.Context, input map[string]interface{}, orgID, userID string) (string, error) {
	return wb.manager.StartWorkflow(ctx, "user_signup", input, orgID, userID)
}

// StartDriverOnboardingWorkflow starts the driver onboarding workflow with deduplication
func (wb *sWorkflowBridge) StartDriverOnboardingWorkflow(ctx context.Context, input map[string]interface{}, orgID, userID string) (string, error) {
	return wb.manager.StartWorkflow(ctx, "driver_onboarding", input, orgID, userID)
}

// StartComplianceCheckWorkflow starts the compliance check workflow with deduplication
func (wb *sWorkflowBridge) StartComplianceCheckWorkflow(ctx context.Context, input map[string]interface{}, orgID, userID string) (string, error) {
	return wb.manager.StartWorkflow(ctx, "compliance_check", input, orgID, userID)
}

// GetWorkflowStatus gets the current status of a workflow
func (wb *sWorkflowBridge) GetWorkflowStatus(ctx context.Context, workflowID string) (*riverjobs.WorkflowExecution, error) {
	execution, err := wb.manager.GetWorkflowStatus(ctx, workflowID)
	if err != nil {
		return nil, err
	}
	return execution, nil
}
