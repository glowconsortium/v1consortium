package dbosworkflow

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
)

// WorkflowService provides the main interface for executing workflows
type WorkflowService struct {
	workflows *OnboardingWorkflows
	db        *sql.DB
}

// NewWorkflowService creates a new workflow service instance
func NewWorkflowService(db *sql.DB) *WorkflowService {
	transactions := &OnboardingTransactions{DB: db}
	steps := &OnboardingSteps{}
	workflows := &OnboardingWorkflows{
		Transactions: transactions,
		Steps:        steps,
	}

	return &WorkflowService{
		workflows: workflows,
		db:        db,
	}
}

// StartSignupWorkflow initiates a new user signup workflow
func (s *WorkflowService) StartSignupWorkflow(ctx context.Context, input OnboardingWorkflowInput) (string, error) {
	g.Log().Infof(ctx, "Starting signup workflow for user: %s", input.Email)

	workflowID, err := s.workflows.SignupWorkflow(ctx, input)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to start signup workflow: %v", err)
		return "", fmt.Errorf("failed to start signup workflow: %w", err)
	}

	g.Log().Infof(ctx, "Signup workflow started successfully with ID: %s", workflowID)
	return workflowID, nil
}

// CompleteEmailVerification completes the email verification step
func (s *WorkflowService) CompleteEmailVerification(ctx context.Context, workflowID, userID string) error {
	g.Log().Infof(ctx, "Completing email verification for workflow: %s, user: %s", workflowID, userID)

	err := s.workflows.EmailVerificationWorkflow(ctx, workflowID, userID)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to complete email verification: %v", err)
		return fmt.Errorf("failed to complete email verification: %w", err)
	}

	g.Log().Infof(ctx, "Email verification completed successfully for workflow: %s", workflowID)
	return nil
}

// StartSubscriptionWorkflow initiates the subscription creation workflow
func (s *WorkflowService) StartSubscriptionWorkflow(ctx context.Context, input SubscribeWorkflowInput) (string, error) {
	g.Log().Infof(ctx, "Starting subscription workflow for user: %s", input.UserID)

	subscriptionID, err := s.workflows.SubscriptionWorkflow(ctx, input)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to start subscription workflow: %v", err)
		return "", fmt.Errorf("failed to start subscription workflow: %w", err)
	}

	g.Log().Infof(ctx, "Subscription workflow completed successfully with subscription ID: %s", subscriptionID)
	return subscriptionID, nil
}

// GetWorkflowState retrieves the current state of a workflow
func (s *WorkflowService) GetWorkflowState(ctx context.Context, workflowID string) (*OnboardingWorkflowState, error) {
	state, err := s.workflows.Transactions.GetWorkflowState(ctx, workflowID)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to get workflow state for ID %s: %v", workflowID, err)
		return nil, fmt.Errorf("failed to get workflow state: %w", err)
	}

	return state, nil
}

// ListPendingWorkflows returns workflows that need attention
func (s *WorkflowService) ListPendingWorkflows(ctx context.Context) ([]OnboardingWorkflowState, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT workflow_id, user_id, company_id, status, email_verified, 
		       subscription_id, current_step, error_message, created_at, updated_at
		FROM onboarding_workflow_states 
		WHERE status IN ('pending', 'email_verified', 'failed')
		ORDER BY created_at DESC
		LIMIT 50
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query pending workflows: %w", err)
	}
	defer rows.Close()

	var workflows []OnboardingWorkflowState
	for rows.Next() {
		var w OnboardingWorkflowState
		var subscriptionID sql.NullString
		var errorMessage sql.NullString

		err := rows.Scan(
			&w.WorkflowID, &w.UserID, &w.CompanyID, &w.Status,
			&w.EmailVerified, &subscriptionID, &w.CurrentStep,
			&errorMessage, &w.CreatedAt, &w.UpdatedAt,
		)
		if err != nil {
			g.Log().Errorf(ctx, "Failed to scan workflow row: %v", err)
			continue
		}

		if subscriptionID.Valid {
			w.SubscriptionID = subscriptionID.String
		}
		if errorMessage.Valid {
			w.ErrorMessage = errorMessage.String
		}

		workflows = append(workflows, w)
	}

	return workflows, nil
}

// RetryFailedWorkflow attempts to retry a failed workflow
func (s *WorkflowService) RetryFailedWorkflow(ctx context.Context, workflowID string) error {
	// Get current state
	state, err := s.GetWorkflowState(ctx, workflowID)
	if err != nil {
		return err
	}

	if state.Status != StatusFailed {
		return fmt.Errorf("workflow %s is not in failed state, current status: %s", workflowID, state.Status)
	}

	// Reset the workflow to the previous step
	previousStep := "pending"
	if state.EmailVerified {
		previousStep = "awaiting_subscription"
	}

	err = s.workflows.Transactions.UpdateOnboardingWorkflowState(ctx, OnboardingWorkflowState{
		WorkflowID:  workflowID,
		Status:      StatusPending,
		CurrentStep: previousStep,
		// Clear error message
		ErrorMessage: "",
	})

	if err != nil {
		return fmt.Errorf("failed to reset workflow state: %w", err)
	}

	g.Log().Infof(ctx, "Workflow %s reset for retry", workflowID)
	return nil
}

// GetWorkflowMetrics returns basic metrics about workflow execution
func (s *WorkflowService) GetWorkflowMetrics(ctx context.Context) (map[string]int, error) {
	query := `
		SELECT status, COUNT(*) as count
		FROM onboarding_workflow_states 
		WHERE created_at > NOW() - INTERVAL '24 hours'
		GROUP BY status
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query workflow metrics: %w", err)
	}
	defer rows.Close()

	metrics := make(map[string]int)
	for rows.Next() {
		var status string
		var count int
		if err := rows.Scan(&status, &count); err != nil {
			continue
		}
		metrics[status] = count
	}

	return metrics, nil
}

// CleanupOldWorkflows removes old completed workflow states
func (s *WorkflowService) CleanupOldWorkflows(ctx context.Context, olderThanDays int) (int, error) {
	query := `
		DELETE FROM onboarding_workflow_states 
		WHERE status = $1 
		AND updated_at < NOW() - INTERVAL '%d days'
	`

	result, err := s.db.ExecContext(ctx, fmt.Sprintf(query, olderThanDays), StatusOnboardingComplete)
	if err != nil {
		return 0, fmt.Errorf("failed to cleanup old workflows: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	g.Log().Infof(ctx, "Cleaned up %d old workflow records", rowsAffected)
	return int(rowsAffected), nil
}
