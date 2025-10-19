package steps

import (
	"context"
	"fmt"
	"v1consortium/internal/pkg/riverjobs"

	"github.com/gogf/gf/v2/frame/g"
)

// CreateUserStepWorker creates a user account
type CreateUserStepWorker struct {
	*riverjobs.BaseStepWorker
}

// NewCreateUserStepWorker creates a new user creation step worker
func NewCreateUserStepWorker(workflowManager *riverjobs.SimpleWorkflowManager) *CreateUserStepWorker {
	worker := &CreateUserStepWorker{}
	worker.BaseStepWorker = riverjobs.NewBaseStepWorker("create_user", workflowManager, worker)
	return worker
}

// Execute creates the user account
func (w *CreateUserStepWorker) Execute(ctx context.Context, args riverjobs.StepArgs) (map[string]interface{}, error) {
	g.Log().Info(ctx, "Starting user creation", g.Map{
		"workflow_id": args.WorkflowID,
		"step":        args.StepName,
	})

	// Extract user data from workflow input
	email, ok := args.WorkflowInput["email"].(string)
	if !ok || email == "" {
		return nil, fmt.Errorf("email is required for user creation")
	}

	firstName, _ := args.WorkflowInput["first_name"].(string)
	lastName, _ := args.WorkflowInput["last_name"].(string)
	password, _ := args.WorkflowInput["password"].(string)
	role, _ := args.WorkflowInput["role"].(string)

	if password == "" {
		return nil, fmt.Errorf("password is required for user creation")
	}

	// TODO: Implement actual user creation logic
	// This would typically involve:
	// 1. Creating user in Supabase Auth
	// 2. Creating user profile in local database
	// 3. Setting up user permissions

	g.Log().Info(ctx, "User creation simulated", g.Map{
		"email":      email,
		"first_name": firstName,
		"last_name":  lastName,
		"role":       role,
	})

	// Return user data for next steps
	return map[string]interface{}{
		"user_id":    "temp-user-id-123", // TODO: Replace with actual user ID
		"email":      email,
		"first_name": firstName,
		"last_name":  lastName,
		"role":       role,
		"created":    true,
	}, nil
}
