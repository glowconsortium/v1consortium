package steps

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"
	"v1consortium/internal/consts"
	"v1consortium/internal/pkg/riverjobs"

	"github.com/gogf/gf/v2/frame/g"
)

// ValidateStep validates the signup input data
type ValidateStep struct {
	*BaseStep
}

// NewValidateStep creates a new validation step
func NewValidateStep() *ValidateStep {
	return &ValidateStep{
		BaseStep: NewBaseStep(
			"validate",
			riverjobs.QueueDefault,
			30*time.Second,
			riverjobs.NoRetryPolicy(), // Validation failures should not retry
		),
	}
}

// Execute validates the signup input
func (s *ValidateStep) Execute(ctx context.Context, input riverjobs.StepInput) (riverjobs.StepResult, error) {
	g.Log().Info(ctx, "Starting signup validation", g.Map{
		"workflow_id": input.WorkflowID,
		"step":        input.StepName,
	})

	// Get input data from workflow context
	inputData, exists := input.WorkflowContext["input"]
	if !exists {
		return riverjobs.StepResult{
				Success: false,
				Error:   "input data not found in workflow context",
			}, riverjobs.NewStepError(
				riverjobs.ErrorTypeValidation,
				"input data not found in workflow context",
				false,
				nil,
			)
	}

	// Convert to map for validation
	inputMap, ok := inputData.(map[string]interface{})
	if !ok {
		return riverjobs.StepResult{
				Success: false,
				Error:   "invalid input data format",
			}, riverjobs.NewStepError(
				riverjobs.ErrorTypeValidation,
				"invalid input data format",
				false,
				nil,
			)
	}

	// Validate required fields manually
	if err := s.validateRequiredFields(inputMap); err != nil {
		return riverjobs.StepResult{
				Success: false,
				Error:   err.Error(),
			}, riverjobs.NewStepError(
				riverjobs.ErrorTypeValidation,
				fmt.Sprintf("validation failed: %v", err),
				false,
				err,
			)
	}

	// Additional business validation
	email := inputMap["email"].(string)
	if email == "" {
		return riverjobs.StepResult{
				Success: false,
				Error:   "email cannot be empty",
			}, riverjobs.NewStepError(
				riverjobs.ErrorTypeValidation,
				"email cannot be empty",
				false,
				nil,
			)
	}

	// Check if email already exists (this would be a database call in real implementation)
	if s.emailExists(ctx, email) {
		return riverjobs.StepResult{
				Success: false,
				Error:   "email already exists",
			}, riverjobs.NewStepError(
				riverjobs.ErrorTypeValidation,
				"email already exists",
				false,
				nil,
			)
	}

	g.Log().Info(ctx, "Signup validation completed successfully", g.Map{
		"workflow_id": input.WorkflowID,
		"email":       email,
	})

	return riverjobs.StepResult{
		Success: true,
		OutputData: map[string]interface{}{
			"validated_input": inputMap,
			"email":           email,
		},
		NextState: riverjobs.StateCreatingUser,
	}, nil
}

// emailExists checks if the email already exists in the system
func (s *ValidateStep) emailExists(ctx context.Context, email string) bool {
	// Mock implementation - in real code this would query the database
	// For now, let's assume no email exists to allow the flow to continue
	return false
}

// validateRequiredFields validates the required input fields
func (s *ValidateStep) validateRequiredFields(inputMap map[string]interface{}) error {
	// Check required fields
	requiredFields := []string{"email", "password", "first_name", "last_name", "role"}
	for _, field := range requiredFields {
		value, exists := inputMap[field]
		if !exists {
			return fmt.Errorf("field '%s' is required", field)
		}

		str, ok := value.(string)
		if !ok {
			return fmt.Errorf("field '%s' must be a string", field)
		}

		str = strings.TrimSpace(str)
		if str == "" {
			return fmt.Errorf("field '%s' cannot be empty", field)
		}

		// Update the trimmed value
		inputMap[field] = str
	}

	// Validate email format
	email := inputMap["email"].(string)
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}

	// Validate password length
	password := inputMap["password"].(string)
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	// Validate name lengths
	firstName := inputMap["first_name"].(string)
	if len(firstName) < 2 {
		return fmt.Errorf("first_name must be at least 2 characters long")
	}

	lastName := inputMap["last_name"].(string)
	if len(lastName) < 2 {
		return fmt.Errorf("last_name must be at least 2 characters long")
	}

	// Validate role
	role := inputMap["role"].(string)
	validRoles := map[string]bool{
		string(consts.RoleClientAdmin): true,
		string(consts.RoleEmployee):    true,
		string(consts.RoleHRManager):   true,
	}
	if !validRoles[role] {
		validList := []string{
			string(consts.RoleClientAdmin),
			string(consts.RoleEmployee),
			string(consts.RoleHRManager),
		}
		return fmt.Errorf("invalid role '%s', must be one of: %s", role, strings.Join(validList, ", "))
	}

	return nil
}

// IsRetryable determines if validation errors should be retried
func (s *ValidateStep) IsRetryable(err error) bool {
	// Validation errors are never retryable
	return false
}
