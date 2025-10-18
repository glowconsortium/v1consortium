package riverjobs

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

// WorkflowManager handles workflow state and coordination
type WorkflowManager struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

// NewWorkflowManager creates a new workflow manager
func NewWorkflowManager(db *pgxpool.Pool, logger *slog.Logger) *WorkflowManager {
	return &WorkflowManager{
		db:     db,
		logger: logger,
	}
}

// CheckWorkflowExists checks if a workflow with the same data already exists
func (wm *WorkflowManager) CheckWorkflowExists(ctx context.Context, workflowType, orgID, argsHash string) (string, error) {
	query := `
		SELECT workflow_id 
		FROM workflow_executions 
		WHERE workflow_type = $1 
		  AND org_id = $2 
		  AND args_hash = $3
		  AND status IN ('pending', 'running')
		ORDER BY created_at DESC 
		LIMIT 1
	`

	var existingWorkflowID string
	err := wm.db.QueryRow(ctx, query, workflowType, orgID, argsHash).Scan(&existingWorkflowID)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return "", nil
		}
		return "", fmt.Errorf("failed to check existing workflow: %w", err)
	}

	return existingWorkflowID, nil
}

// CreateWorkflowExecution creates a new workflow execution record
func (wm *WorkflowManager) CreateWorkflowExecution(ctx context.Context, execution *WorkflowExecution) error {
	query := `
		INSERT INTO workflow_executions (
			workflow_id, workflow_type, org_id, user_id, status,
			context, args_hash, started_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING workflow_id
	`

	contextJSON, _ := json.Marshal(execution.Context)

	return wm.db.QueryRow(ctx, query,
		execution.WorkflowID, execution.WorkflowType, execution.OrgID,
		execution.UserID, execution.Status, contextJSON, execution.ArgsHash, execution.StartedAt,
	).Scan(&execution.WorkflowID)
}

// GetWorkflowExecution retrieves a workflow execution by workflow ID
func (wm *WorkflowManager) GetWorkflowExecution(ctx context.Context, workflowID string) (*WorkflowExecution, error) {
	var execution WorkflowExecution
	query := `
		SELECT workflow_id, workflow_type, org_id, user_id, status, 
		       current_step, context, args_hash, started_at, completed_at,
		       error_message, retry_count, created_at, updated_at
		FROM workflow_executions 
		WHERE workflow_id = $1
	`

	var contextJSON []byte
	err := wm.db.QueryRow(ctx, query, workflowID).Scan(
		&execution.WorkflowID, &execution.WorkflowType,
		&execution.OrgID, &execution.UserID, &execution.Status,
		&execution.CurrentStep, &contextJSON, &execution.ArgsHash,
		&execution.StartedAt, &execution.CompletedAt, &execution.ErrorMessage,
		&execution.RetryCount, &execution.CreatedAt, &execution.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get workflow execution: %w", err)
	}

	if len(contextJSON) > 0 {
		json.Unmarshal(contextJSON, &execution.Context)
	}

	return &execution, nil
}

// UpdateWorkflowStatus updates the workflow status
func (wm *WorkflowManager) UpdateWorkflowStatus(ctx context.Context, workflowID, status string, errorMsg *string) error {
	query := `
		UPDATE workflow_executions 
		SET status = $2, error_message = $3, updated_at = NOW()
		WHERE workflow_id = $1
	`

	_, err := wm.db.Exec(ctx, query, workflowID, status, errorMsg)
	return err
}

// UpdateWorkflowContext updates the workflow context with new data
func (wm *WorkflowManager) UpdateWorkflowContext(ctx context.Context, workflowID string, contextUpdate map[string]interface{}) error {
	query := `
		UPDATE workflow_executions 
		SET context = context || $2, updated_at = NOW()
		WHERE workflow_id = $1
	`
	contextJSON, err := json.Marshal(contextUpdate)
	if err != nil {
		return fmt.Errorf("failed to marshal context: %w", err)
	}

	_, err = wm.db.Exec(ctx, query, workflowID, contextJSON)
	if err != nil {
		wm.logger.Error("failed to update workflow context", "error", err, "workflow_id", workflowID)
	}
	return err
}

// CreateWorkflowStep creates a new workflow step record
func (wm *WorkflowManager) CreateWorkflowStep(ctx context.Context, workflowID, stepName string, inputData interface{}) (string, error) {
	var stepID string
	query := `
		INSERT INTO workflow_steps (workflow_id, step_name, status, input_data)
		VALUES ($1, $2, 'pending', $3)
		RETURNING step_id
	`

	inputJSON, _ := json.Marshal(inputData)
	err := wm.db.QueryRow(ctx, query, workflowID, stepName, inputJSON).Scan(&stepID)
	return stepID, err
}

// UpdateWorkflowStep updates the status of a workflow step
func (wm *WorkflowManager) UpdateWorkflowStep(ctx context.Context, workflowID, stepName, status string) error {
	query := `
		UPDATE workflow_steps 
		SET status = $3
		WHERE workflow_id = $1 AND step_name = $2
	`
	_, err := wm.db.Exec(ctx, query, workflowID, stepName, status)
	if err != nil {
		wm.logger.Error("failed to update workflow step", "error", err, "workflow_id", workflowID, "step", stepName)
	}
	return err
}

// UpdateWorkflowStepJobID updates the job ID for a workflow step
func (wm *WorkflowManager) UpdateWorkflowStepJobID(ctx context.Context, stepID string, jobID int64) error {
	query := `UPDATE workflow_steps SET river_job_id = $2 WHERE step_id = $1`
	_, err := wm.db.Exec(ctx, query, stepID, jobID)
	return err
}

// GetWorkflowSteps retrieves all steps for a workflow
func (wm *WorkflowManager) GetWorkflowSteps(ctx context.Context, workflowID string) ([]WorkflowStep, error) {
	query := `
		SELECT step_id, workflow_id, step_name, step_order, river_job_id, queue_name, status,
		       input_data, output_data, started_at, completed_at,
		       error_message, retry_count, max_retries, created_at
		FROM workflow_steps
		WHERE workflow_id = $1
		ORDER BY step_order, created_at
	`

	rows, err := wm.db.Query(ctx, query, workflowID)
	if err != nil {
		return nil, fmt.Errorf("failed to get workflow steps: %w", err)
	}
	defer rows.Close()

	var steps []WorkflowStep
	for rows.Next() {
		var step WorkflowStep
		var inputJSON, outputJSON []byte

		err := rows.Scan(
			&step.ID, &step.WorkflowID, &step.StepName, &step.StepOrder, &step.RiverJobID, &step.QueueName,
			&step.Status, &inputJSON, &outputJSON, &step.StartedAt,
			&step.CompletedAt, &step.ErrorMessage, &step.RetryCount, &step.MaxRetries, &step.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan workflow step: %w", err)
		}

		if len(inputJSON) > 0 {
			json.Unmarshal(inputJSON, &step.InputData)
		}
		if len(outputJSON) > 0 {
			json.Unmarshal(outputJSON, &step.OutputData)
		}

		steps = append(steps, step)
	}

	return steps, nil
}

// CompleteWorkflow marks a workflow as completed
func (wm *WorkflowManager) CompleteWorkflow(ctx context.Context, workflowID string) error {
	query := `
		UPDATE workflow_executions 
		SET status = $2, completed_at = NOW(), updated_at = NOW()
		WHERE workflow_id = $1
	`

	_, err := wm.db.Exec(ctx, query, workflowID, StatusCompleted)
	return err
}

// FailWorkflow marks a workflow as failed
func (wm *WorkflowManager) FailWorkflow(ctx context.Context, workflowID string, errorMsg string) error {
	query := `
		UPDATE workflow_executions 
		SET status = $2, error_message = $3, updated_at = NOW()
		WHERE workflow_id = $1
	`

	_, err := wm.db.Exec(ctx, query, workflowID, StatusFailed, errorMsg)
	return err
}

// CancelWorkflow marks a workflow as cancelled
func (wm *WorkflowManager) CancelWorkflow(ctx context.Context, workflowID string) error {
	query := `
		UPDATE workflow_executions 
		SET status = $2, updated_at = NOW()
		WHERE workflow_id = $1
	`

	_, err := wm.db.Exec(ctx, query, workflowID, StatusCancelled)
	return err
}

// UpdateStepStatus updates the status of a workflow step
func (wm *WorkflowManager) UpdateStepStatus(ctx context.Context, workflowID, stepName, status string, outputData map[string]interface{}, errorMsg *string) error {
	var outputJSON []byte
	if outputData != nil {
		var err error
		outputJSON, err = json.Marshal(outputData)
		if err != nil {
			return fmt.Errorf("failed to marshal output data: %w", err)
		}
	}

	var query string
	var args []interface{}

	if status == StepStatusRunning {
		query = `
			UPDATE workflow_steps 
			SET status = $3, started_at = NOW(), updated_at = NOW()
			WHERE workflow_id = $1 AND step_name = $2
		`
		args = []interface{}{workflowID, stepName, status}
	} else if status == StepStatusCompleted {
		query = `
			UPDATE workflow_steps 
			SET status = $3, output_data = $4, completed_at = NOW(), updated_at = NOW()
			WHERE workflow_id = $1 AND step_name = $2
		`
		args = []interface{}{workflowID, stepName, status, outputJSON}
	} else if status == StepStatusFailed {
		query = `
			UPDATE workflow_steps 
			SET status = $3, error_message = $4, updated_at = NOW()
			WHERE workflow_id = $1 AND step_name = $2
		`
		args = []interface{}{workflowID, stepName, status, errorMsg}
	} else {
		query = `
			UPDATE workflow_steps 
			SET status = $3, updated_at = NOW()
			WHERE workflow_id = $1 AND step_name = $2
		`
		args = []interface{}{workflowID, stepName, status}
	}

	_, err := wm.db.Exec(ctx, query, args...)
	return err
}
