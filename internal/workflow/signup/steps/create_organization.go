package steps

import (
	"context"
	"fmt"
	"time"
	"v1consortium/internal/pkg/riverjobs"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
)

// CreateOrganizationStep creates an organization record
type CreateOrganizationStep struct {
	*BaseStep
}

// NewCreateOrganizationStep creates a new organization creation step
func NewCreateOrganizationStep() *CreateOrganizationStep {
	return &CreateOrganizationStep{
		BaseStep: NewBaseStep(
			"create_organization",
			riverjobs.QueueDefault,
			2*time.Minute,
			riverjobs.DefaultRetryPolicy(),
		),
	}
}

// Execute creates the organization
func (s *CreateOrganizationStep) Execute(ctx context.Context, input riverjobs.StepInput) (riverjobs.StepResult, error) {
	g.Log().Info(ctx, "Starting organization creation", g.Map{
		"workflow_id": input.WorkflowID,
		"step":        input.StepName,
	})

	// Get user data from previous step
	userID, exists := input.WorkflowContext["user_id"]
	if !exists {
		return riverjobs.StepResult{
				Success: false,
				Error:   "user_id not found in context",
			}, riverjobs.NewStepError(
				riverjobs.ErrorTypeBusiness,
				"user_id not found in workflow context",
				false,
				nil,
			)
	}

	// Mock organization creation
	orgID := fmt.Sprintf("org_%s", guid.S())
	orgName := fmt.Sprintf("Organization for %s", userID)

	// Create organization record
	err := s.createOrganization(ctx, orgID, orgName, userID.(string))
	if err != nil {
		return riverjobs.StepResult{
				Success: false,
				Error:   err.Error(),
			}, riverjobs.NewStepError(
				riverjobs.ErrorTypeDatabase,
				fmt.Sprintf("failed to create organization: %v", err),
				true,
				err,
			)
	}

	g.Log().Info(ctx, "Organization creation completed", g.Map{
		"workflow_id": input.WorkflowID,
		"org_id":      orgID,
		"org_name":    orgName,
	})

	return riverjobs.StepResult{
		Success: true,
		OutputData: map[string]interface{}{
			"org_id":   orgID,
			"org_name": orgName,
		},
		NextState: riverjobs.StateSettingUpPayment,
	}, nil
}

// createOrganization creates organization in database
func (s *CreateOrganizationStep) createOrganization(ctx context.Context, orgID, orgName, userID string) error {
	g.Log().Info(ctx, "Creating organization in database", g.Map{
		"org_id":   orgID,
		"org_name": orgName,
		"user_id":  userID,
	})
	// Mock implementation
	return nil
}

// Compensate removes the created organization
func (s *CreateOrganizationStep) Compensate(ctx context.Context, input riverjobs.StepInput) error {
	g.Log().Info(ctx, "Starting organization compensation", g.Map{
		"workflow_id": input.WorkflowID,
	})

	if orgID, exists := input.WorkflowContext["org_id"]; exists {
		err := s.deleteOrganization(ctx, orgID.(string))
		if err != nil {
			g.Log().Error(ctx, "Failed to delete organization during compensation", g.Map{
				"org_id": orgID,
				"error":  err,
			})
		}
	}

	return nil
}

// deleteOrganization removes organization from database
func (s *CreateOrganizationStep) deleteOrganization(ctx context.Context, orgID string) error {
	g.Log().Info(ctx, "Deleting organization", g.Map{"org_id": orgID})
	// Mock implementation
	return nil
}
