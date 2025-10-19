package steps

import (
	"context"
	"fmt"
	"time"
	"v1consortium/internal/pkg/riverjobs"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
)

// CreateUserStep creates a user account in Supabase and local database
type CreateUserStep struct {
	*BaseStep
}

// NewCreateUserStep creates a new user creation step
func NewCreateUserStep() *CreateUserStep {
	return &CreateUserStep{
		BaseStep: NewBaseStep(
			"create_user",
			riverjobs.QueueDefault,
			2*time.Minute,
			riverjobs.DefaultRetryPolicy(),
		),
	}
}

// Execute creates the user account
func (s *CreateUserStep) Execute(ctx context.Context, input riverjobs.StepInput) (riverjobs.StepResult, error) {
	g.Log().Info(ctx, "Starting user creation", g.Map{
		"workflow_id": input.WorkflowID,
		"step":        input.StepName,
	})

	// Get validated input from previous step
	validatedInput, exists := input.WorkflowContext["validated_input"]
	if !exists {
		return riverjobs.StepResult{
				Success: false,
				Error:   "validated input not found",
			}, riverjobs.NewStepError(
				riverjobs.ErrorTypeBusiness,
				"validated input not found in workflow context",
				false,
				nil,
			)
	}

	inputData := validatedInput.(map[string]interface{})
	email := inputData["email"].(string)
	password := inputData["password"].(string)
	firstName := inputData["first_name"].(string)
	lastName := inputData["last_name"].(string)
	role := inputData["role"].(string)

	// Create user in Supabase Auth
	userID, err := s.createSupabaseUser(ctx, email, password, firstName, lastName)
	if err != nil {
		return riverjobs.StepResult{
				Success: false,
				Error:   err.Error(),
			}, riverjobs.NewStepError(
				riverjobs.ErrorTypeExternal,
				fmt.Sprintf("failed to create Supabase user: %v", err),
				true, // Supabase errors are usually retryable
				err,
			)
	}

	// Create user profile in local database
	err = s.createUserProfile(ctx, userID, email, firstName, lastName, role)
	if err != nil {
		// Compensate by deleting Supabase user
		s.deleteSupabaseUser(ctx, userID)

		return riverjobs.StepResult{
				Success: false,
				Error:   err.Error(),
			}, riverjobs.NewStepError(
				riverjobs.ErrorTypeDatabase,
				fmt.Sprintf("failed to create user profile: %v", err),
				true,
				err,
			)
	}

	g.Log().Info(ctx, "User creation completed successfully", g.Map{
		"workflow_id": input.WorkflowID,
		"user_id":     userID,
		"email":       email,
	})

	return riverjobs.StepResult{
		Success: true,
		OutputData: map[string]interface{}{
			"user_id":    userID,
			"email":      email,
			"first_name": firstName,
			"last_name":  lastName,
			"role":       role,
		},
		NextState: riverjobs.StateCreatingOrg,
	}, nil
}

// createSupabaseUser creates a user in Supabase Auth
func (s *CreateUserStep) createSupabaseUser(ctx context.Context, email, password, firstName, lastName string) (string, error) {
	// Mock implementation - in real code this would call Supabase
	userID := fmt.Sprintf("user_%s", guid.S())

	g.Log().Info(ctx, "Created Supabase user", g.Map{
		"user_id": userID,
		"email":   email,
	})

	// Simulate potential network error (5% chance)
	// if rand.Intn(100) < 5 {
	// 	return "", fmt.Errorf("supabase network error")
	// }

	return userID, nil
}

// createUserProfile creates user profile in local database
func (s *CreateUserStep) createUserProfile(ctx context.Context, userID, email, firstName, lastName, role string) error {
	// Mock implementation - in real code this would insert into database
	g.Log().Info(ctx, "Created user profile in database", g.Map{
		"user_id":    userID,
		"email":      email,
		"first_name": firstName,
		"last_name":  lastName,
		"role":       role,
	})

	// Simulate potential database error (2% chance)
	// if rand.Intn(100) < 2 {
	// 	return fmt.Errorf("database connection error")
	// }

	return nil
}

// deleteSupabaseUser removes user from Supabase (compensation)
func (s *CreateUserStep) deleteSupabaseUser(ctx context.Context, userID string) error {
	g.Log().Info(ctx, "Deleting Supabase user for compensation", g.Map{
		"user_id": userID,
	})

	// Mock implementation
	return nil
}

// Compensate removes the created user
func (s *CreateUserStep) Compensate(ctx context.Context, input riverjobs.StepInput) error {
	g.Log().Info(ctx, "Starting user creation compensation", g.Map{
		"workflow_id": input.WorkflowID,
	})

	// Get user ID from workflow context
	if outputData, exists := input.WorkflowContext["user_id"]; exists {
		userID := outputData.(string)

		// Delete from Supabase
		if err := s.deleteSupabaseUser(ctx, userID); err != nil {
			g.Log().Error(ctx, "Failed to delete Supabase user during compensation", g.Map{
				"user_id": userID,
				"error":   err,
			})
		}

		// Delete from local database
		if err := s.deleteUserProfile(ctx, userID); err != nil {
			g.Log().Error(ctx, "Failed to delete user profile during compensation", g.Map{
				"user_id": userID,
				"error":   err,
			})
		}
	}

	return nil
}

// deleteUserProfile removes user profile from database
func (s *CreateUserStep) deleteUserProfile(ctx context.Context, userID string) error {
	g.Log().Info(ctx, "Deleting user profile from database", g.Map{
		"user_id": userID,
	})

	// Mock implementation
	return nil
}

// IsRetryable determines if user creation errors should be retried
func (s *CreateUserStep) IsRetryable(err error) bool {
	if stepError, ok := err.(*riverjobs.StepError); ok {
		// Network and database errors are retryable
		return stepError.Type == riverjobs.ErrorTypeNetwork ||
			stepError.Type == riverjobs.ErrorTypeDatabase ||
			stepError.Type == riverjobs.ErrorTypeExternal
	}
	return true
}
