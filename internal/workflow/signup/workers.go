package signup

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"v1consortium/internal/pkg/riverjobs"

	"github.com/jackc/pgx/v5"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/rivertype"
)

// ValidateStepWorker handles user signup validation
type ValidateStepWorker struct {
	WorkflowManager *riverjobs.WorkflowManager
	RiverClient     *river.Client[pgx.Tx]
	Logger          *slog.Logger
}

// Work implements the River worker interface
func (w *ValidateStepWorker) Work(ctx context.Context, job *river.Job[riverjobs.ValidateStepArgs]) error {
	args := job.Args

	w.Logger.Info("Starting signup validation", "workflow_id", args.WorkflowID)

	// Extract signup data
	signupData := args.SignupData

	// Validate required fields
	if err := w.validateRequiredFields(signupData); err != nil {
		return riverjobs.NewStepError(riverjobs.ErrorTypeValidation, err.Error(), false, err)
	}

	// Validate email format
	if err := w.validateEmail(signupData); err != nil {
		return riverjobs.NewStepError(riverjobs.ErrorTypeValidation, err.Error(), false, err)
	}

	// Validate password strength
	if err := w.validatePassword(signupData); err != nil {
		return riverjobs.NewStepError(riverjobs.ErrorTypeValidation, err.Error(), false, err)
	}

	// Update step status and prepare for next step
	outputData := map[string]interface{}{
		"validated_data":       signupData,
		"validation_timestamp": time.Now().Unix(),
	}

	// Update workflow step status
	err := w.WorkflowManager.UpdateStepStatus(ctx, args.WorkflowID, "validate", riverjobs.StepStatusCompleted, outputData, nil)
	if err != nil {
		return fmt.Errorf("failed to update step status: %w", err)
	}

	// Enqueue next step
	nextArgs := riverjobs.CreateUserStepArgs{
		BaseJobArgs: riverjobs.BaseJobArgs{
			WorkflowID:   args.WorkflowID,
			WorkflowType: args.WorkflowType,
			StepName:     "create_user",
			OrgID:        args.OrgID,
			UserID:       args.UserID,
		},
		UserData: outputData,
	}

	_, err = w.RiverClient.Insert(ctx, nextArgs, nil)
	if err != nil {
		return fmt.Errorf("failed to enqueue next step: %w", err)
	}

	w.Logger.Info("Signup validation completed", "workflow_id", args.WorkflowID)
	return nil
}

// NextRetry implements the River worker interface
func (w *ValidateStepWorker) NextRetry(job *river.Job[riverjobs.ValidateStepArgs]) time.Time {
	return time.Now().Add(time.Duration(job.Attempt) * time.Minute)
}

// Timeout implements the River worker interface
func (w *ValidateStepWorker) Timeout(job *river.Job[riverjobs.ValidateStepArgs]) time.Duration {
	return 2 * time.Minute
}

// Middleware implements the River worker interface
func (w *ValidateStepWorker) Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware {
	return []rivertype.WorkerMiddleware{}
}

// validateRequiredFields checks that all required fields are present
func (w *ValidateStepWorker) validateRequiredFields(data map[string]interface{}) error {
	required := []string{"email", "password", "first_name", "last_name", "organization_name"}

	for _, field := range required {
		if value, exists := data[field]; !exists || value == "" {
			return fmt.Errorf("required field missing or empty: %s", field)
		}
	}

	return nil
}

// validateEmail checks email format
func (w *ValidateStepWorker) validateEmail(data map[string]interface{}) error {
	email, ok := data["email"].(string)
	if !ok {
		return errors.New("email must be a string")
	}

	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return errors.New("invalid email format")
	}

	return nil
}

// validatePassword checks password strength
func (w *ValidateStepWorker) validatePassword(data map[string]interface{}) error {
	password, ok := data["password"].(string)
	if !ok {
		return errors.New("password must be a string")
	}

	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	return nil
}

// CreateUserStepWorker handles user creation in Supabase
type CreateUserStepWorker struct {
	WorkflowManager *riverjobs.WorkflowManager
	RiverClient     *river.Client[pgx.Tx]
	Logger          *slog.Logger
}

// Work implements the River worker interface
func (w *CreateUserStepWorker) Work(ctx context.Context, job *river.Job[riverjobs.CreateUserStepArgs]) error {
	args := job.Args

	w.Logger.Info("Starting user creation", "workflow_id", args.WorkflowID)

	// TODO: Implement actual user creation logic with Supabase
	// For now, simulate the process
	userData := args.UserData

	// Simulate user creation
	userID := fmt.Sprintf("user_%d", time.Now().UnixNano())

	outputData := map[string]interface{}{
		"user_id":    userID,
		"user_data":  userData,
		"created_at": time.Now().Unix(),
	}

	// Update step status
	err := w.WorkflowManager.UpdateStepStatus(ctx, args.WorkflowID, "create_user", riverjobs.StepStatusCompleted, outputData, nil)
	if err != nil {
		return fmt.Errorf("failed to update step status: %w", err)
	}

	// Enqueue next step
	nextArgs := riverjobs.CreateOrganizationStepArgs{
		BaseJobArgs: riverjobs.BaseJobArgs{
			WorkflowID:   args.WorkflowID,
			WorkflowType: args.WorkflowType,
			StepName:     "create_organization",
			OrgID:        args.OrgID,
			UserID:       userID,
		},
		OrgData: outputData,
	}

	_, err = w.RiverClient.Insert(ctx, nextArgs, nil)
	if err != nil {
		return fmt.Errorf("failed to enqueue next step: %w", err)
	}

	w.Logger.Info("User creation completed", "workflow_id", args.WorkflowID, "user_id", userID)
	return nil
}

// NextRetry implements the River worker interface
func (w *CreateUserStepWorker) NextRetry(job *river.Job[riverjobs.CreateUserStepArgs]) time.Time {
	return time.Now().Add(time.Duration(job.Attempt) * time.Minute)
}

// Timeout implements the River worker interface
func (w *CreateUserStepWorker) Timeout(job *river.Job[riverjobs.CreateUserStepArgs]) time.Duration {
	return 5 * time.Minute
}

// Middleware implements the River worker interface
func (w *CreateUserStepWorker) Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware {
	return []rivertype.WorkerMiddleware{}
}

// CreateOrganizationStepWorker handles organization creation
type CreateOrganizationStepWorker struct {
	WorkflowManager *riverjobs.WorkflowManager
	RiverClient     *river.Client[pgx.Tx]
	Logger          *slog.Logger
}

// Work implements the River worker interface
func (w *CreateOrganizationStepWorker) Work(ctx context.Context, job *river.Job[riverjobs.CreateOrganizationStepArgs]) error {
	args := job.Args

	w.Logger.Info("Starting organization creation", "workflow_id", args.WorkflowID)

	// TODO: Implement actual organization creation logic
	// For now, simulate the process
	orgData := args.OrgData

	// Simulate organization creation
	orgID := fmt.Sprintf("org_%d", time.Now().UnixNano())

	outputData := map[string]interface{}{
		"org_id":     orgID,
		"org_data":   orgData,
		"created_at": time.Now().Unix(),
	}

	// Update step status
	err := w.WorkflowManager.UpdateStepStatus(ctx, args.WorkflowID, "create_organization", riverjobs.StepStatusCompleted, outputData, nil)
	if err != nil {
		return fmt.Errorf("failed to update step status: %w", err)
	}

	// Enqueue next step
	nextArgs := riverjobs.SetupStripeStepArgs{
		BaseJobArgs: riverjobs.BaseJobArgs{
			WorkflowID:   args.WorkflowID,
			WorkflowType: args.WorkflowType,
			StepName:     "setup_stripe",
			OrgID:        orgID,
			UserID:       args.UserID,
		},
		StripeData: outputData,
	}

	_, err = w.RiverClient.Insert(ctx, nextArgs, nil)
	if err != nil {
		return fmt.Errorf("failed to enqueue next step: %w", err)
	}

	w.Logger.Info("Organization creation completed", "workflow_id", args.WorkflowID, "org_id", orgID)
	return nil
}

// NextRetry implements the River worker interface
func (w *CreateOrganizationStepWorker) NextRetry(job *river.Job[riverjobs.CreateOrganizationStepArgs]) time.Time {
	return time.Now().Add(time.Duration(job.Attempt) * time.Minute)
}

// Timeout implements the River worker interface
func (w *CreateOrganizationStepWorker) Timeout(job *river.Job[riverjobs.CreateOrganizationStepArgs]) time.Duration {
	return 3 * time.Minute
}

// Middleware implements the River worker interface
func (w *CreateOrganizationStepWorker) Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware {
	return []rivertype.WorkerMiddleware{}
}

// SetupStripeStepWorker handles Stripe customer and subscription setup
type SetupStripeStepWorker struct {
	WorkflowManager *riverjobs.WorkflowManager
	RiverClient     *river.Client[pgx.Tx]
	Logger          *slog.Logger
}

// Work implements the River worker interface
func (w *SetupStripeStepWorker) Work(ctx context.Context, job *river.Job[riverjobs.SetupStripeStepArgs]) error {
	args := job.Args

	w.Logger.Info("Starting Stripe setup", "workflow_id", args.WorkflowID)

	// TODO: Implement actual Stripe integration
	// For now, simulate the process
	stripeData := args.StripeData

	// Simulate Stripe customer creation
	customerID := fmt.Sprintf("cus_%d", time.Now().UnixNano())

	outputData := map[string]interface{}{
		"stripe_customer_id": customerID,
		"stripe_data":        stripeData,
		"created_at":         time.Now().Unix(),
	}

	// Update step status
	err := w.WorkflowManager.UpdateStepStatus(ctx, args.WorkflowID, "setup_stripe", riverjobs.StepStatusCompleted, outputData, nil)
	if err != nil {
		return fmt.Errorf("failed to update step status: %w", err)
	}

	// Enqueue next step
	nextArgs := riverjobs.SendVerificationStepArgs{
		BaseJobArgs: riverjobs.BaseJobArgs{
			WorkflowID:   args.WorkflowID,
			WorkflowType: args.WorkflowType,
			StepName:     "send_verification",
			OrgID:        args.OrgID,
			UserID:       args.UserID,
		},
		VerificationData: outputData,
	}

	_, err = w.RiverClient.Insert(ctx, nextArgs, nil)
	if err != nil {
		return fmt.Errorf("failed to enqueue next step: %w", err)
	}

	w.Logger.Info("Stripe setup completed", "workflow_id", args.WorkflowID, "customer_id", customerID)
	return nil
}

// NextRetry implements the River worker interface
func (w *SetupStripeStepWorker) NextRetry(job *river.Job[riverjobs.SetupStripeStepArgs]) time.Time {
	return time.Now().Add(time.Duration(job.Attempt) * time.Minute)
}

// Timeout implements the River worker interface
func (w *SetupStripeStepWorker) Timeout(job *river.Job[riverjobs.SetupStripeStepArgs]) time.Duration {
	return 10 * time.Minute // Longer timeout for external service
}

// Middleware implements the River worker interface
func (w *SetupStripeStepWorker) Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware {
	return []rivertype.WorkerMiddleware{}
}

// SendVerificationStepWorker handles sending verification email
type SendVerificationStepWorker struct {
	WorkflowManager *riverjobs.WorkflowManager
	RiverClient     *river.Client[pgx.Tx]
	Logger          *slog.Logger
}

// Work implements the River worker interface
func (w *SendVerificationStepWorker) Work(ctx context.Context, job *river.Job[riverjobs.SendVerificationStepArgs]) error {
	args := job.Args

	w.Logger.Info("Starting verification email sending", "workflow_id", args.WorkflowID)

	// TODO: Implement actual email sending logic
	// For now, simulate the process
	verificationData := args.VerificationData

	// Simulate sending verification email
	verificationToken := fmt.Sprintf("token_%d", time.Now().UnixNano())

	outputData := map[string]interface{}{
		"verification_token": verificationToken,
		"verification_data":  verificationData,
		"sent_at":            time.Now().Unix(),
	}

	// Update step status
	err := w.WorkflowManager.UpdateStepStatus(ctx, args.WorkflowID, "send_verification", riverjobs.StepStatusCompleted, outputData, nil)
	if err != nil {
		return fmt.Errorf("failed to update step status: %w", err)
	}

	// Update workflow status to pending verification
	err = w.WorkflowManager.UpdateWorkflowStatus(ctx, args.WorkflowID, string(riverjobs.StatePendingVerification), nil)
	if err != nil {
		return fmt.Errorf("failed to update workflow status: %w", err)
	}

	w.Logger.Info("Verification email sent", "workflow_id", args.WorkflowID, "token", verificationToken)
	return nil
}

// NextRetry implements the River worker interface
func (w *SendVerificationStepWorker) NextRetry(job *river.Job[riverjobs.SendVerificationStepArgs]) time.Time {
	return time.Now().Add(time.Duration(job.Attempt) * time.Minute)
}

// Timeout implements the River worker interface
func (w *SendVerificationStepWorker) Timeout(job *river.Job[riverjobs.SendVerificationStepArgs]) time.Duration {
	return 3 * time.Minute
}

// Middleware implements the River worker interface
func (w *SendVerificationStepWorker) Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware {
	return []rivertype.WorkerMiddleware{}
}
