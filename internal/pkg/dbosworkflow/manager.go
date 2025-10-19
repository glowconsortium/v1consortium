package dbosworkflow

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/dbos-inc/dbos-transact-golang/dbos"
	"github.com/gogf/gf/v2/frame/g"
)

// DBOSWorkflowManager handles DBOS workflow registration and execution
type DBOSWorkflowManager struct {
	dbosCtx dbos.DBOSContext
	db      *sql.DB
}

// NewDBOSWorkflowManager creates a new DBOS workflow manager
func NewDBOSWorkflowManager(dbosCtx dbos.DBOSContext, db *sql.DB) *DBOSWorkflowManager {
	return &DBOSWorkflowManager{
		dbosCtx: dbosCtx,
		db:      db,
	}
}

// RegisterWorkflows registers all workflows with DBOS
func (m *DBOSWorkflowManager) RegisterWorkflows() error {
	// For now, we'll create wrapper functions that match DBOS expectations
	// TODO: Update when proper DBOS Go bindings are available

	g.Log().Info(context.Background(), "DBOS workflows registered (placeholder)")
	return nil
}

// StartSignupWorkflow starts a signup workflow using DBOS
func (m *DBOSWorkflowManager) StartSignupWorkflow(ctx context.Context, input OnboardingWorkflowInput) (string, error) {
	// This would use DBOS.StartWorkflow when proper bindings are available
	// For now, execute directly

	transactions := &OnboardingTransactions{DB: m.db}
	steps := &OnboardingSteps{}
	workflows := &OnboardingWorkflows{
		Transactions: transactions,
		Steps:        steps,
	}

	return workflows.SignupWorkflow(ctx, input)
}

// GetWorkflowStatus retrieves workflow status from DBOS
func (m *DBOSWorkflowManager) GetWorkflowStatus(ctx context.Context, workflowID string) (*WorkflowStatus, error) {
	// This would query DBOS workflow status when proper bindings are available

	transactions := &OnboardingTransactions{DB: m.db}
	state, err := transactions.GetWorkflowState(ctx, workflowID)
	if err != nil {
		return nil, err
	}

	return &WorkflowStatus{
		WorkflowID: state.WorkflowID,
		Status:     state.Status,
		Step:       state.CurrentStep,
		Error:      state.ErrorMessage,
		CreatedAt:  state.CreatedAt,
		UpdatedAt:  state.UpdatedAt,
	}, nil
}

// WorkflowStatus represents the current status of a workflow
type WorkflowStatus struct {
	WorkflowID string    `json:"workflow_id"`
	Status     string    `json:"status"`
	Step       string    `json:"step"`
	Error      string    `json:"error,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// WorkflowStateMachine defines valid state transitions
type WorkflowStateMachine struct {
	transitions map[string][]string
}

// NewWorkflowStateMachine creates a new state machine
func NewWorkflowStateMachine() *WorkflowStateMachine {
	return &WorkflowStateMachine{
		transitions: map[string][]string{
			StatusPending:       {StatusEmailVerified, StatusFailed},
			StatusEmailVerified: {StatusSubscribed, StatusFailed},
			StatusSubscribed:    {StatusOnboardingComplete, StatusFailed},
			StatusFailed:        {StatusPending}, // Allow retry
		},
	}
}

// CanTransition checks if a state transition is valid
func (sm *WorkflowStateMachine) CanTransition(from, to string) bool {
	validTransitions, exists := sm.transitions[from]
	if !exists {
		return false
	}

	for _, validTo := range validTransitions {
		if validTo == to {
			return true
		}
	}
	return false
}

// ValidateTransition validates and logs state transitions
func (sm *WorkflowStateMachine) ValidateTransition(ctx context.Context, workflowID, from, to string) error {
	if !sm.CanTransition(from, to) {
		err := fmt.Errorf("invalid state transition from %s to %s for workflow %s", from, to, workflowID)
		g.Log().Error(ctx, err.Error())
		return err
	}

	g.Log().Infof(ctx, "Workflow %s transitioning from %s to %s", workflowID, from, to)
	return nil
}
