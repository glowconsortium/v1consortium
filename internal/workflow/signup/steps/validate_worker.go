package steps

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"v1consortium/internal/pkg/riverjobs"

	"github.com/riverqueue/river"
)

// ValidateStepWorker handles user signup validation
type ValidateStepWorker struct {
	riverjobs.WorkerBase
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

	w.Logger.Info("Signup validation completed", "workflow_id", args.WorkflowID)
	return nil
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
