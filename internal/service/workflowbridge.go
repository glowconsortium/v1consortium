// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"v1consortium/internal/pkg/riverjobs"
)

type (
	IWorkflowBridge interface {
		// StartWorkflow starts a new workflow instance
		StartWorkflow(ctx context.Context, workflowName string, input map[string]interface{}, orgID string, userID string) (string, error)
		// StartUserSignupWorkflow starts the user signup workflow with deduplication
		StartUserSignupWorkflow(ctx context.Context, input map[string]interface{}, orgID string, userID string) (string, error)
		// StartDriverOnboardingWorkflow starts the driver onboarding workflow with deduplication
		StartDriverOnboardingWorkflow(ctx context.Context, input map[string]interface{}, orgID string, userID string) (string, error)
		// StartComplianceCheckWorkflow starts the compliance check workflow with deduplication
		StartComplianceCheckWorkflow(ctx context.Context, input map[string]interface{}, orgID string, userID string) (string, error)
		// GetWorkflowStatus gets the current status of a workflow
		GetWorkflowStatus(ctx context.Context, workflowID string) (*riverjobs.WorkflowExecution, error)
	}
)

var (
	localWorkflowBridge IWorkflowBridge
)

func WorkflowBridge() IWorkflowBridge {
	if localWorkflowBridge == nil {
		panic("implement not found for interface IWorkflowBridge, forgot register?")
	}
	return localWorkflowBridge
}

func RegisterWorkflowBridge(i IWorkflowBridge) {
	localWorkflowBridge = i
}
