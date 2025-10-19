package riverjobs

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
)

// WorkflowService provides high-level workflow operations
type WorkflowService struct {
	workflowManager *SimpleWorkflowManager
}

// NewWorkflowService creates a new workflow service
func NewWorkflowService(workflowManager *SimpleWorkflowManager) *WorkflowService {
	return &WorkflowService{
		workflowManager: workflowManager,
	}
}

// StartSignupWorkflow starts a user signup workflow
func (ws *WorkflowService) StartSignupWorkflow(ctx context.Context, input map[string]interface{}, orgID, userID string) (string, error) {
	// Validate required fields
	if input["email"] == nil || input["email"].(string) == "" {
		return "", fmt.Errorf("email is required")
	}

	g.Log().Info(ctx, "Starting signup workflow", g.Map{
		"email":  input["email"],
		"org_id": orgID,
	})

	// Start the workflow
	workflowID, err := ws.workflowManager.StartWorkflow(ctx, "user_signup", input, orgID, userID)
	if err != nil {
		return "", fmt.Errorf("failed to start signup workflow: %w", err)
	}

	g.Log().Info(ctx, "Signup workflow started", g.Map{
		"workflow_id": workflowID,
		"email":       input["email"],
	})

	return workflowID, nil
}

// GetWorkflowStatus returns the current status of a workflow
func (ws *WorkflowService) GetWorkflowStatus(ctx context.Context, workflowID string) (*WorkflowExecution, error) {
	return ws.workflowManager.GetWorkflowStatus(ctx, workflowID)
}

// Example usage:
/*
func (c *Controller) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	// Prepare signup input
	input := map[string]interface{}{
		"input": map[string]interface{}{
			"email":      req.Email,
			"password":   req.Password,
			"first_name": req.FirstName,
			"last_name":  req.LastName,
			"role":       req.Role,
		},
	}

	// Start workflow
	workflowID, err := workflowService.StartSignupWorkflow(ctx, input, req.OrgID, "")
	if err != nil {
		return nil, err
	}

	return &api.CreateUserResponse{
		WorkflowID: workflowID,
		Message:    "User signup initiated",
	}, nil
}
*/
