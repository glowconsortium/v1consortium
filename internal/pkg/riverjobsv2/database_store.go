package riverjobsv2

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DatabaseWorkflowStore implements WorkflowStore using the workflow_executions table
type DatabaseWorkflowStore struct {
	db *pgxpool.Pool // Direct pgxpool connection
	// Alternative: use GoFrame's gdb.DB if you prefer
	// gdb gdb.DB
}

// NewDatabaseWorkflowStore creates a new database-backed workflow store
func NewDatabaseWorkflowStore(db *pgxpool.Pool) *DatabaseWorkflowStore {
	return &DatabaseWorkflowStore{
		db: db,
	}
}

// SaveWorkflowStatus saves or updates workflow status in the database
func (s *DatabaseWorkflowStore) SaveWorkflowStatus(ctx context.Context, status *WorkflowStatus) error {
	query := `
		INSERT INTO workflow_executions (
			workflow_id, workflow_type, status, context, current_step, 
			started_at, completed_at, org_id, user_id, error_message, args_hash
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (workflow_id) DO UPDATE SET
			status = EXCLUDED.status,
			context = EXCLUDED.context,
			current_step = EXCLUDED.current_step,
			completed_at = EXCLUDED.completed_at,
			error_message = EXCLUDED.error_message,
			updated_at = NOW()
	`

	// Extract org_id and user_id from input if available
	var orgID, userID interface{}
	if status.Input != nil {
		if oid, exists := status.Input["org_id"]; exists {
			orgID = oid
		}
		if uid, exists := status.Input["user_id"]; exists {
			userID = uid
		}
	}

	// Generate args hash for deduplication
	argsHash := ""
	if status.Input != nil {
		hash, err := GenerateInputHash(status.WorkflowName, status.Input)
		if err == nil {
			argsHash = hash
		}
	}

	_, err := s.db.Exec(ctx, query,
		status.WorkflowID,
		status.WorkflowName,
		status.Status,
		status.Context,
		status.CurrentStep,
		status.StartedAt,
		status.CompletedAt,
		orgID,
		userID,
		status.ErrorMessage,
		argsHash,
	)

	return err
}

// GetWorkflowStatus retrieves workflow status by ID from the database
func (s *DatabaseWorkflowStore) GetWorkflowStatus(ctx context.Context, workflowID string) (*WorkflowStatus, error) {
	query := `
		SELECT workflow_id, workflow_type, status, context, current_step,
			   started_at, completed_at, created_at, error_message
		FROM workflow_executions 
		WHERE workflow_id = $1
	`

	var status WorkflowStatus
	var createdAt time.Time
	var contextBytes []byte
	err := s.db.QueryRow(ctx, query, workflowID).Scan(
		&status.WorkflowID,
		&status.WorkflowName,
		&status.Status,
		&contextBytes,
		&status.CurrentStep,
		&status.StartedAt,
		&status.CompletedAt,
		&createdAt,
		&status.ErrorMessage,
	)

	if err != nil {
		if err == sql.ErrNoRows || err == pgx.ErrNoRows {
			return nil, fmt.Errorf("workflow not found: %s", workflowID)
		}
		return nil, fmt.Errorf("failed to get workflow status: %w", err)
	}

	// Parse JSONB context
	if len(contextBytes) > 0 {
		if err := json.Unmarshal(contextBytes, &status.Context); err != nil {
			return nil, fmt.Errorf("failed to parse workflow context: %w", err)
		}
	} else {
		status.Context = make(WorkflowContext)
	}

	if err != nil {
		if err == sql.ErrNoRows || err == pgx.ErrNoRows {
			return nil, fmt.Errorf("workflow not found: %s", workflowID)
		}
		return nil, fmt.Errorf("failed to get workflow status: %w", err)
	}

	// Use created_at as started_at if started_at is null
	if status.StartedAt.IsZero() {
		status.StartedAt = createdAt
	}

	return &status, nil
}

// GetRunningWorkflowByHash finds a running workflow with the same input hash
func (s *DatabaseWorkflowStore) GetRunningWorkflowByHash(ctx context.Context, workflowName, inputHash string) (*WorkflowStatus, error) {
	query := `
		SELECT workflow_id, workflow_type, status, context, current_step,
			   started_at, completed_at, created_at, error_message
		FROM workflow_executions 
		WHERE workflow_type = $1 AND args_hash = $2 AND status IN ('pending', 'running')
		ORDER BY created_at DESC
		LIMIT 1
	`

	var status WorkflowStatus
	var createdAt time.Time
	var contextBytes []byte
	err := s.db.QueryRow(ctx, query, workflowName, inputHash).Scan(
		&status.WorkflowID,
		&status.WorkflowName,
		&status.Status,
		&contextBytes,
		&status.CurrentStep,
		&status.StartedAt,
		&status.CompletedAt,
		&createdAt,
		&status.ErrorMessage,
	)

	if err != nil {
		if err == sql.ErrNoRows || err == pgx.ErrNoRows {
			return nil, nil // No running workflow found (not an error)
		}
		return nil, fmt.Errorf("failed to check for existing workflow: %w", err)
	}

	// Parse JSONB context
	if len(contextBytes) > 0 {
		if err := json.Unmarshal(contextBytes, &status.Context); err != nil {
			return nil, fmt.Errorf("failed to parse workflow context: %w", err)
		}
	} else {
		status.Context = make(WorkflowContext)
	}

	// Use created_at as started_at if started_at is null
	if status.StartedAt.IsZero() {
		status.StartedAt = createdAt
	}

	return &status, nil
}

// MarkWorkflowCompleted marks a workflow as completed in the database
func (s *DatabaseWorkflowStore) MarkWorkflowCompleted(ctx context.Context, workflowID string, finalContext WorkflowContext) error {
	query := `
		UPDATE workflow_executions 
		SET status = 'completed', 
			context = $2, 
			completed_at = NOW(),
			updated_at = NOW()
		WHERE workflow_id = $1
	`

	result, err := s.db.Exec(ctx, query, workflowID, finalContext)
	if err != nil {
		return fmt.Errorf("failed to mark workflow as completed: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("workflow not found: %s", workflowID)
	}

	return nil
}

// MarkWorkflowFailed marks a workflow as failed in the database
func (s *DatabaseWorkflowStore) MarkWorkflowFailed(ctx context.Context, workflowID string, errorMessage string) error {
	query := `
		UPDATE workflow_executions 
		SET status = 'failed', 
			error_message = $2, 
			completed_at = NOW(),
			updated_at = NOW()
		WHERE workflow_id = $1
	`

	result, err := s.db.Exec(ctx, query, workflowID, errorMessage)
	if err != nil {
		return fmt.Errorf("failed to mark workflow as failed: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("workflow not found: %s", workflowID)
	}

	return nil
}

// GoFrameWorkflowStore is an alternative implementation using GoFrame's gdb
// Uncomment and use this if you prefer GoFrame's database abstraction
/*
type GoFrameWorkflowStore struct {
	db gdb.DB
}

func NewGoFrameWorkflowStore(db gdb.DB) *GoFrameWorkflowStore {
	return &GoFrameWorkflowStore{db: db}
}

func (s *GoFrameWorkflowStore) SaveWorkflowStatus(ctx context.Context, status *WorkflowStatus) error {
	// Extract org_id and user_id from input if available
	var orgID, userID interface{}
	if status.Input != nil {
		if oid, exists := status.Input["org_id"]; exists {
			orgID = oid
		}
		if uid, exists := status.Input["user_id"]; exists {
			userID = uid
		}
	}

	// Generate args hash for deduplication
	argsHash := ""
	if status.Input != nil {
		hash, err := GenerateInputHash(status.WorkflowName, status.Input)
		if err == nil {
			argsHash = hash
		}
	}

	data := g.Map{
		"workflow_id":   status.WorkflowID,
		"workflow_type": status.WorkflowName,
		"status":        status.Status,
		"context":       status.Context,
		"current_step":  status.CurrentStep,
		"started_at":    status.StartedAt,
		"completed_at":  status.CompletedAt,
		"org_id":        orgID,
		"user_id":       userID,
		"error_message": status.ErrorMessage,
		"args_hash":     argsHash,
	}

	_, err := s.db.Model("workflow_executions").Ctx(ctx).
		Data(data).
		OnConflict("workflow_id").
		Save()

	return err
}

// ... implement other methods similarly using GoFrame's gdb
*/
