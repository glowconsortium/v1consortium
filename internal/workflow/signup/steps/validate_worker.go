package steps

import (
	"context"
	"fmt"
	"v1consortium/internal/pkg/riverjobs"

	"github.com/gogf/gf/v2/frame/g"
)

// ValidateStepWorker validates the signup input data
type ValidateStepWorker struct {
	*riverjobs.BaseStepWorker
}

// NewValidateStepWorker creates a new validation step worker
func NewValidateStepWorker(workflowManager *riverjobs.SimpleWorkflowManager) *ValidateStepWorker {
	worker := &ValidateStepWorker{}
	worker.BaseStepWorker = riverjobs.NewBaseStepWorker("validate", workflowManager, worker)
	return worker
}

// Execute validates the signup input
func (w *ValidateStepWorker) Execute(ctx context.Context, args riverjobs.StepArgs) (map[string]interface{}, error) {
	g.Log().Info(ctx, "Starting signup validation", g.Map{
		"workflow_id": args.WorkflowID,
		"step":        args.StepName,
	})

	// Basic validation - check if input exists
	if args.WorkflowInput == nil {
		return nil, fmt.Errorf("workflow input is missing")
	}

	// Validate email exists (data is now passed directly in WorkflowInput)
	email, hasEmail := args.WorkflowInput["email"].(string)
	if !hasEmail || email == "" {
		return nil, fmt.Errorf("email is required")
	}

	// Validate other required fields
	firstName, _ := args.WorkflowInput["first_name"].(string)
	lastName, _ := args.WorkflowInput["last_name"].(string)
	password, _ := args.WorkflowInput["password"].(string)

	if firstName == "" {
		return nil, fmt.Errorf("first_name is required")
	}
	if lastName == "" {
		return nil, fmt.Errorf("last_name is required")
	}
	if password == "" {
		return nil, fmt.Errorf("password is required")
	}

	g.Log().Info(ctx, "Signup validation completed successfully", g.Map{
		"workflow_id": args.WorkflowID,
		"email":       email,
		"first_name":  firstName,
		"last_name":   lastName,
	})

	// Return validated input for next steps (pass through the original input)
	return args.WorkflowInput, nil
}

// GetStepName returns the step name
func (w *ValidateStepWorker) GetStepName() string {
	return "validate"
}
