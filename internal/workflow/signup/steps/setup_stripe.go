package steps

import (
	"context"
	"fmt"
	"time"
	"v1consortium/internal/pkg/riverjobs"

	"github.com/gogf/gf/v2/frame/g"
)

// SetupStripeStep sets up Stripe customer and subscription
type SetupStripeStep struct {
	*BaseStep
}

// NewSetupStripeStep creates a new Stripe setup step
func NewSetupStripeStep() *SetupStripeStep {
	return &SetupStripeStep{
		BaseStep: NewBaseStep(
			"setup_stripe",
			riverjobs.QueueExternal,
			3*time.Minute,
			riverjobs.AggressiveRetryPolicy(),
		),
	}
}

// Execute sets up Stripe customer and subscription
func (s *SetupStripeStep) Execute(ctx context.Context, input riverjobs.StepInput) (riverjobs.StepResult, error) {
	g.Log().Info(ctx, "Starting Stripe setup", g.Map{
		"workflow_id": input.WorkflowID,
		"step":        input.StepName,
	})

	// Get organization data
	orgID, exists := input.WorkflowContext["org_id"]
	if !exists {
		return riverjobs.StepResult{
				Success: false,
				Error:   "org_id not found in context",
			}, riverjobs.NewStepError(
				riverjobs.ErrorTypeBusiness,
				"org_id not found in workflow context",
				false,
				nil,
			)
	}

	email, exists := input.WorkflowContext["email"]
	if !exists {
		return riverjobs.StepResult{
				Success: false,
				Error:   "email not found in context",
			}, riverjobs.NewStepError(
				riverjobs.ErrorTypeBusiness,
				"email not found in workflow context",
				false,
				nil,
			)
	}

	// Create Stripe customer
	customerID, err := s.createStripeCustomer(ctx, email.(string), orgID.(string))
	if err != nil {
		return riverjobs.StepResult{
				Success: false,
				Error:   err.Error(),
			}, riverjobs.NewStepError(
				riverjobs.ErrorTypeExternal,
				fmt.Sprintf("failed to create Stripe customer: %v", err),
				true,
				err,
			)
	}

	// Create default subscription
	subscriptionID, err := s.createStripeSubscription(ctx, customerID)
	if err != nil {
		// Compensate by deleting customer
		s.deleteStripeCustomer(ctx, customerID)

		return riverjobs.StepResult{
				Success: false,
				Error:   err.Error(),
			}, riverjobs.NewStepError(
				riverjobs.ErrorTypeExternal,
				fmt.Sprintf("failed to create Stripe subscription: %v", err),
				true,
				err,
			)
	}

	g.Log().Info(ctx, "Stripe setup completed", g.Map{
		"workflow_id":     input.WorkflowID,
		"customer_id":     customerID,
		"subscription_id": subscriptionID,
	})

	return riverjobs.StepResult{
		Success: true,
		OutputData: map[string]interface{}{
			"stripe_customer_id":     customerID,
			"stripe_subscription_id": subscriptionID,
		},
		NextState: riverjobs.StateSendingVerification,
	}, nil
}

// createStripeCustomer creates a Stripe customer
func (s *SetupStripeStep) createStripeCustomer(ctx context.Context, email, orgID string) (string, error) {
	g.Log().Info(ctx, "Creating Stripe customer", g.Map{
		"email":  email,
		"org_id": orgID,
	})

	// Mock implementation
	customerID := fmt.Sprintf("cus_%s", orgID)
	return customerID, nil
}

// createStripeSubscription creates a Stripe subscription
func (s *SetupStripeStep) createStripeSubscription(ctx context.Context, customerID string) (string, error) {
	g.Log().Info(ctx, "Creating Stripe subscription", g.Map{
		"customer_id": customerID,
	})

	// Mock implementation
	subscriptionID := fmt.Sprintf("sub_%s", customerID)
	return subscriptionID, nil
}

// deleteStripeCustomer removes Stripe customer
func (s *SetupStripeStep) deleteStripeCustomer(ctx context.Context, customerID string) error {
	g.Log().Info(ctx, "Deleting Stripe customer", g.Map{
		"customer_id": customerID,
	})
	// Mock implementation
	return nil
}

// Compensate removes Stripe customer and subscription
func (s *SetupStripeStep) Compensate(ctx context.Context, input riverjobs.StepInput) error {
	g.Log().Info(ctx, "Starting Stripe compensation", g.Map{
		"workflow_id": input.WorkflowID,
	})

	if customerID, exists := input.WorkflowContext["stripe_customer_id"]; exists {
		err := s.deleteStripeCustomer(ctx, customerID.(string))
		if err != nil {
			g.Log().Error(ctx, "Failed to delete Stripe customer during compensation", g.Map{
				"customer_id": customerID,
				"error":       err,
			})
		}
	}

	return nil
}
