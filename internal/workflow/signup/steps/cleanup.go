package steps

import (
	"context"
	"time"
	"v1consortium/internal/pkg/riverjobs"

	"github.com/gogf/gf/v2/frame/g"
)

// CleanupUserStep removes user data during compensation
type CleanupUserStep struct {
	*BaseStep
}

// NewCleanupUserStep creates a new cleanup user step
func NewCleanupUserStep() *CleanupUserStep {
	return &CleanupUserStep{
		BaseStep: NewBaseStep(
			"cleanup_user",
			riverjobs.QueueDefault,
			1*time.Minute,
			riverjobs.NoRetryPolicy(), // Cleanup steps shouldn't retry
		),
	}
}

// Execute removes user data
func (s *CleanupUserStep) Execute(ctx context.Context, input riverjobs.StepInput) (riverjobs.StepResult, error) {
	g.Log().Info(ctx, "Starting user cleanup", g.Map{
		"workflow_id": input.WorkflowID,
	})

	// Get user ID from context
	if userID, exists := input.WorkflowContext["user_id"]; exists {
		s.deleteUser(ctx, userID.(string))
	}

	return riverjobs.StepResult{
		Success: true,
		OutputData: map[string]interface{}{
			"user_cleaned_up": true,
		},
	}, nil
}

// deleteUser removes user from all systems
func (s *CleanupUserStep) deleteUser(ctx context.Context, userID string) {
	g.Log().Info(ctx, "Cleaning up user", g.Map{"user_id": userID})
	// Mock cleanup implementation
}

// CleanupOrganizationStep removes organization data during compensation
type CleanupOrganizationStep struct {
	*BaseStep
}

// NewCleanupOrganizationStep creates a new cleanup organization step
func NewCleanupOrganizationStep() *CleanupOrganizationStep {
	return &CleanupOrganizationStep{
		BaseStep: NewBaseStep(
			"cleanup_organization",
			riverjobs.QueueDefault,
			1*time.Minute,
			riverjobs.NoRetryPolicy(),
		),
	}
}

// Execute removes organization data
func (s *CleanupOrganizationStep) Execute(ctx context.Context, input riverjobs.StepInput) (riverjobs.StepResult, error) {
	g.Log().Info(ctx, "Starting organization cleanup", g.Map{
		"workflow_id": input.WorkflowID,
	})

	// Get org ID from context
	if orgID, exists := input.WorkflowContext["org_id"]; exists {
		s.deleteOrganization(ctx, orgID.(string))
	}

	return riverjobs.StepResult{
		Success: true,
		OutputData: map[string]interface{}{
			"organization_cleaned_up": true,
		},
	}, nil
}

// deleteOrganization removes organization from database
func (s *CleanupOrganizationStep) deleteOrganization(ctx context.Context, orgID string) {
	g.Log().Info(ctx, "Cleaning up organization", g.Map{"org_id": orgID})
	// Mock cleanup implementation
}
