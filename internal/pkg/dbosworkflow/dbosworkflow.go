package dbosworkflow

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
)

// WorkflowExecutor handles custom workflow execution
type WorkflowExecutor struct {
	workflows *OnboardingWorkflows
	db        *sql.DB
}

// NewWorkflowExecutor creates a new workflow executor instance
func NewWorkflowExecutor(db *sql.DB) *WorkflowExecutor {
	transactions := &OnboardingTransactions{DB: db}
	steps := &OnboardingSteps{}
	workflows := &OnboardingWorkflows{
		Transactions: transactions,
		Steps:        steps,
	}

	return &WorkflowExecutor{
		workflows: workflows,
		db:        db,
	}
}

// ProcessPendingWorkflows checks for and processes pending workflows
func (w *WorkflowExecutor) ProcessPendingWorkflows(ctx context.Context) error {
	g.Log().Debug(ctx, "Processing pending workflows...")

	// Query for pending workflows
	rows, err := w.db.QueryContext(ctx, `
		SELECT workflow_id, user_id, company_id, status, current_step 
		FROM onboarding_workflow_states 
		WHERE status IN ('pending', 'email_verified') 
		AND updated_at < NOW() - INTERVAL '1 minute'
		LIMIT 10
	`)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to query pending workflows: %v", err)
		return fmt.Errorf("failed to query pending workflows: %w", err)
	}
	defer rows.Close()

	var processed int
	for rows.Next() {
		var workflowID, userID, companyID, status, currentStep string
		if err := rows.Scan(&workflowID, &userID, &companyID, &status, &currentStep); err != nil {
			g.Log().Errorf(ctx, "Failed to scan workflow row: %v", err)
			continue
		}

		g.Log().Debugf(ctx, "Processing workflow: %s, status: %s, step: %s", workflowID, status, currentStep)

		if err := w.processWorkflow(ctx, workflowID, userID, companyID, status, currentStep); err != nil {
			g.Log().Errorf(ctx, "Failed to process workflow %s: %v", workflowID, err)
		} else {
			processed++
		}
	}

	if processed > 0 {
		g.Log().Infof(ctx, "Processed %d pending workflows", processed)
	}

	return nil
}

// processWorkflow handles individual workflow processing
func (w *WorkflowExecutor) processWorkflow(ctx context.Context, workflowID, userID, companyID, status, currentStep string) error {
	g.Log().Debugf(ctx, "Processing individual workflow: %s", workflowID)

	switch status {
	case "pending":
		return w.processPendingWorkflow(ctx, workflowID, userID, companyID, currentStep)
	case "email_verified":
		return w.processEmailVerifiedWorkflow(ctx, workflowID, userID, companyID, currentStep)
	default:
		g.Log().Debugf(ctx, "Workflow %s has status %s, no processing needed", workflowID, status)
		return nil
	}
}

// processPendingWorkflow handles workflows in pending status
func (w *WorkflowExecutor) processPendingWorkflow(ctx context.Context, workflowID, userID, companyID, currentStep string) error {
	switch currentStep {
	case "awaiting_verification":
		g.Log().Debugf(ctx, "Workflow %s waiting for email verification", workflowID)
		// TODO: Check if user should be auto-verified (e.g., OAuth users)
		// TODO: Send reminder emails for pending verifications
		return nil
	case "verification_timeout":
		g.Log().Debugf(ctx, "Workflow %s verification timeout - sending reminder", workflowID)
		// TODO: Send verification reminder email
		return nil
	default:
		g.Log().Debugf(ctx, "Unknown pending step %s for workflow %s", currentStep, workflowID)
		return nil
	}
}

// processEmailVerifiedWorkflow handles workflows where email is verified
func (w *WorkflowExecutor) processEmailVerifiedWorkflow(ctx context.Context, workflowID, userID, companyID, currentStep string) error {
	switch currentStep {
	case "awaiting_subscription":
		g.Log().Debugf(ctx, "Workflow %s waiting for subscription", workflowID)
		// TODO: Check for completed subscriptions
		// TODO: Send subscription reminder emails
		return nil
	case "subscription_timeout":
		g.Log().Debugf(ctx, "Workflow %s subscription timeout - sending reminder", workflowID)
		// TODO: Send subscription reminder email
		return nil
	default:
		g.Log().Debugf(ctx, "Unknown email verified step %s for workflow %s", currentStep, workflowID)
		return nil
	}
}

// ProcessFailedWorkflows checks for and handles failed workflows
func (w *WorkflowExecutor) ProcessFailedWorkflows(ctx context.Context) error {
	g.Log().Debug(ctx, "Processing failed workflows...")

	rows, err := w.db.QueryContext(ctx, `
		SELECT workflow_id, user_id, company_id, status, current_step, error_message
		FROM onboarding_workflow_states 
		WHERE status = 'failed' 
		AND updated_at < NOW() - INTERVAL '5 minutes'
		LIMIT 5
	`)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to query failed workflows: %v", err)
		return fmt.Errorf("failed to query failed workflows: %w", err)
	}
	defer rows.Close()

	var processed int
	for rows.Next() {
		var workflowID, userID, companyID, status, currentStep, errorMessage string
		if err := rows.Scan(&workflowID, &userID, &companyID, &status, &currentStep, &errorMessage); err != nil {
			g.Log().Errorf(ctx, "Failed to scan failed workflow row: %v", err)
			continue
		}

		g.Log().Infof(ctx, "Processing failed workflow: %s, error: %s", workflowID, errorMessage)

		if err := w.handleFailedWorkflow(ctx, workflowID, userID, companyID, currentStep, errorMessage); err != nil {
			g.Log().Errorf(ctx, "Failed to handle failed workflow %s: %v", workflowID, err)
		} else {
			processed++
		}
	}

	if processed > 0 {
		g.Log().Infof(ctx, "Processed %d failed workflows", processed)
	}

	return nil
}

// handleFailedWorkflow processes individual failed workflows
func (w *WorkflowExecutor) handleFailedWorkflow(ctx context.Context, workflowID, userID, companyID, currentStep, errorMessage string) error {
	g.Log().Debugf(ctx, "Handling failed workflow: %s, step: %s", workflowID, currentStep)

	// TODO: Implement retry logic based on failure type
	// TODO: Send notification to admin team
	// TODO: Log to monitoring system
	// TODO: Determine if workflow should be auto-retried or marked for manual review

	switch currentStep {
	case "subscription_failed":
		g.Log().Debugf(ctx, "Workflow %s failed at subscription step", workflowID)
		// TODO: Check if payment method is still valid
		// TODO: Retry subscription creation
		return nil
	case "payment_failed":
		g.Log().Debugf(ctx, "Workflow %s failed at payment step", workflowID)
		// TODO: Send payment failure notification to user
		// TODO: Provide alternative payment options
		return nil
	default:
		g.Log().Debugf(ctx, "Unknown failure step %s for workflow %s", currentStep, workflowID)
		return nil
	}
}

// GetWorkflowMetrics returns basic metrics about workflow execution
func (w *WorkflowExecutor) GetWorkflowMetrics(ctx context.Context) (map[string]int, error) {
	g.Log().Debug(ctx, "Collecting workflow metrics...")

	query := `
		SELECT status, COUNT(*) as count
		FROM onboarding_workflow_states 
		WHERE created_at > NOW() - INTERVAL '24 hours'
		GROUP BY status
	`

	rows, err := w.db.QueryContext(ctx, query)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to query workflow metrics: %v", err)
		return nil, fmt.Errorf("failed to query workflow metrics: %w", err)
	}
	defer rows.Close()

	metrics := make(map[string]int)
	for rows.Next() {
		var status string
		var count int
		if err := rows.Scan(&status, &count); err != nil {
			g.Log().Errorf(ctx, "Failed to scan metrics row: %v", err)
			continue
		}
		metrics[status] = count
	}

	g.Log().Debugf(ctx, "Workflow metrics collected: %+v", metrics)
	return metrics, nil
}

// CleanupCompletedWorkflows removes old completed workflow states
func (w *WorkflowExecutor) CleanupCompletedWorkflows(ctx context.Context, olderThanDays int) (int, error) {
	g.Log().Debugf(ctx, "Cleaning up workflows older than %d days...", olderThanDays)

	query := `
		DELETE FROM onboarding_workflow_states 
		WHERE status = $1 
		AND updated_at < NOW() - INTERVAL '%d days'
	`

	result, err := w.db.ExecContext(ctx, fmt.Sprintf(query, olderThanDays), "onboarding_complete")
	if err != nil {
		g.Log().Errorf(ctx, "Failed to cleanup old workflows: %v", err)
		return 0, fmt.Errorf("failed to cleanup old workflows: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	g.Log().Infof(ctx, "Cleaned up %d old workflow records", rowsAffected)
	return int(rowsAffected), nil
}
