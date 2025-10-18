package steps

import (
	"context"
	"fmt"
	"time"
	"v1consortium/internal/pkg/riverjobs"

	"github.com/gogf/gf/v2/frame/g"
)

// SendVerificationStep sends email verification to the user
type SendVerificationStep struct {
	*BaseStep
}

// NewSendVerificationStep creates a new verification step
func NewSendVerificationStep() *SendVerificationStep {
	return &SendVerificationStep{
		BaseStep: NewBaseStep(
			"send_verification",
			riverjobs.QueueNotification,
			1*time.Minute,
			riverjobs.DefaultRetryPolicy(),
		),
	}
}

// Execute sends the verification email
func (s *SendVerificationStep) Execute(ctx context.Context, input riverjobs.StepInput) (riverjobs.StepResult, error) {
	g.Log().Info(ctx, "Starting verification email send", g.Map{
		"workflow_id": input.WorkflowID,
		"step":        input.StepName,
	})

	// Get user data
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

	firstName, _ := input.WorkflowContext["first_name"]

	// Send verification email
	err := s.sendVerificationEmail(ctx, userID.(string), email.(string), firstName)
	if err != nil {
		return riverjobs.StepResult{
				Success:     false,
				Error:       err.Error(),
				ShouldRetry: true, // Email sending is usually retryable
			}, riverjobs.NewStepError(
				riverjobs.ErrorTypeExternal,
				fmt.Sprintf("failed to send verification email: %v", err),
				true,
				err,
			)
	}

	g.Log().Info(ctx, "Verification email sent successfully", g.Map{
		"workflow_id": input.WorkflowID,
		"user_id":     userID,
		"email":       email,
	})

	return riverjobs.StepResult{
		Success: true,
		OutputData: map[string]interface{}{
			"verification_sent": true,
			"sent_at":           time.Now().Unix(),
		},
		NextState: riverjobs.StatePendingVerification,
	}, nil
}

// sendVerificationEmail sends the verification email
func (s *SendVerificationStep) sendVerificationEmail(ctx context.Context, userID, email string, firstName interface{}) error {
	g.Log().Info(ctx, "Sending verification email", g.Map{
		"user_id":    userID,
		"email":      email,
		"first_name": firstName,
	})

	// Mock implementation - in real code this would send email via service
	// Simulate occasional email service failures
	// if rand.Intn(100) < 10 {
	// 	return fmt.Errorf("email service temporarily unavailable")
	// }

	return nil
}

// IsRetryable determines if email sending errors should be retried
func (s *SendVerificationStep) IsRetryable(err error) bool {
	if stepError, ok := err.(*riverjobs.StepError); ok {
		// Email service errors are usually retryable
		return stepError.Type == riverjobs.ErrorTypeExternal || stepError.Type == riverjobs.ErrorTypeNetwork
	}
	return true
}

// Compensate doesn't do anything for verification emails
func (s *SendVerificationStep) Compensate(ctx context.Context, input riverjobs.StepInput) error {
	// No compensation needed for verification emails
	g.Log().Info(ctx, "No compensation needed for verification email step", g.Map{
		"workflow_id": input.WorkflowID,
	})
	return nil
}
