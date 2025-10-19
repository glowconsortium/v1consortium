package riverjobs

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
)

// SimpleWorkflowManager handles workflow lifecycle with a choreography approach
type SimpleWorkflowManager struct {
	riverManager *RiverManager
	dbPool       *pgxpool.Pool
	workflows    map[string]*WorkflowDefinition // workflow_type -> definition
}

// WorkflowDefinition defines a workflow's steps and configuration
type WorkflowDefinition struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Steps       []StepDefinition `json:"steps"`
	Config      WorkflowConfig   `json:"config"`
}

// StepDefinition defines a single step in a workflow
type StepDefinition struct {
	Name       string        `json:"name"`
	Queue      string        `json:"queue"`
	Timeout    time.Duration `json:"timeout"`
	MaxRetries int           `json:"max_retries"`
	NextSteps  []string      `json:"next_steps"`  // Next steps to enqueue on success
	ErrorSteps []string      `json:"error_steps"` // Steps to enqueue on error
	Required   bool          `json:"required"`    // If false, workflow can continue if this step fails
}

// WorkflowConfig contains workflow-level configuration
type WorkflowConfig struct {
	MaxRetries     int           `json:"max_retries"`
	DefaultQueue   string        `json:"default_queue"`
	DefaultTimeout time.Duration `json:"default_timeout"`
}

// NewSimpleWorkflowManager creates a new simple workflow manager
func NewSimpleWorkflowManager(riverManager *RiverManager, dbPool *pgxpool.Pool) *SimpleWorkflowManager {
	return &SimpleWorkflowManager{
		riverManager: riverManager,
		dbPool:       dbPool,
		workflows:    make(map[string]*WorkflowDefinition),
	}
}

// RegisterWorkflow registers a workflow definition
func (wm *SimpleWorkflowManager) RegisterWorkflow(workflow *WorkflowDefinition) {
	wm.workflows[workflow.Name] = workflow
	g.Log().Info(context.Background(), "Registered workflow", g.Map{
		"workflow_name": workflow.Name,
		"step_count":    len(workflow.Steps),
	})
}

// StartWorkflow starts a new workflow by enqueueing the first step
func (wm *SimpleWorkflowManager) StartWorkflow(ctx context.Context, workflowType string, input map[string]interface{}, orgID, userID string) (string, error) {
	workflow, exists := wm.workflows[workflowType]
	if !exists {
		// Log available workflows for debugging
		availableWorkflows := make([]string, 0, len(wm.workflows))
		for name := range wm.workflows {
			availableWorkflows = append(availableWorkflows, name)
		}
		g.Log().Error(ctx, "Workflow type not found", g.Map{
			"requested_workflow":  workflowType,
			"available_workflows": availableWorkflows,
		})
		return "", fmt.Errorf("workflow type %s not found", workflowType)
	}

	if len(workflow.Steps) == 0 {
		return "", fmt.Errorf("workflow %s has no steps defined", workflowType)
	}

	// Check for duplicate workflow (prevent enqueueing same job that is not failed)
	inputBytes, _ := json.Marshal(input)
	argsHash := fmt.Sprintf("%x", md5.Sum(inputBytes))

	existingWorkflowID, err := wm.checkForDuplicateWorkflow(ctx, workflowType, orgID, userID, argsHash)
	if err != nil {
		return "", fmt.Errorf("failed to check for duplicate workflow: %w", err)
	}
	if existingWorkflowID != "" {
		g.Log().Info(ctx, "Workflow already running", g.Map{
			"existing_workflow_id": existingWorkflowID,
			"workflow_type":        workflowType,
			"org_id":               orgID,
		})
		return existingWorkflowID, nil
	}

	// Generate workflow ID using proper UUID
	workflowID := uuid.New().String()

	// Create workflow execution record
	err = wm.createWorkflowExecution(ctx, workflowID, workflowType, orgID, userID, input)
	if err != nil {
		return "", fmt.Errorf("failed to create workflow execution: %w", err)
	}

	// Enqueue first step
	firstStep := workflow.Steps[0]
	err = wm.enqueueStep(ctx, workflowID, workflowType, firstStep.Name, input, nil, orgID, userID)
	if err != nil {
		return "", fmt.Errorf("failed to enqueue first step: %w", err)
	}

	g.Log().Info(ctx, "Started workflow", g.Map{
		"workflow_id":   workflowID,
		"workflow_type": workflowType,
		"first_step":    firstStep.Name,
		"org_id":        orgID,
	})

	return workflowID, nil
}

// EnqueueNextSteps enqueues the next steps in a workflow (called by step workers)
func (wm *SimpleWorkflowManager) EnqueueNextSteps(ctx context.Context, workflowID, workflowType, currentStep string, workflowInput map[string]interface{}, stepOutput map[string]interface{}, orgID, userID string) error {
	workflow, exists := wm.workflows[workflowType]
	if !exists {
		return fmt.Errorf("workflow type %s not found", workflowType)
	}

	// Find current step definition
	var currentStepDef *StepDefinition
	for _, step := range workflow.Steps {
		if step.Name == currentStep {
			currentStepDef = &step
			break
		}
	}

	if currentStepDef == nil {
		return fmt.Errorf("step %s not found in workflow %s", currentStep, workflowType)
	}

	// Update workflow context with step output
	if stepOutput != nil {
		// Merge step output into workflow input for next steps
		for key, value := range stepOutput {
			workflowInput[key] = value
		}
	}

	// Update workflow state
	err := wm.updateWorkflowStep(ctx, workflowID, currentStep, workflowInput)
	if err != nil {
		return fmt.Errorf("failed to update workflow state: %w", err)
	}

	// Enqueue next steps
	for _, nextStepName := range currentStepDef.NextSteps {
		err = wm.enqueueStep(ctx, workflowID, workflowType, nextStepName, workflowInput, stepOutput, orgID, userID)
		if err != nil {
			g.Log().Error(ctx, "Failed to enqueue next step", g.Map{
				"workflow_id": workflowID,
				"step":        nextStepName,
				"error":       err,
			})
			// Continue with other steps
		}
	}

	// If no next steps, mark workflow as completed
	if len(currentStepDef.NextSteps) == 0 {
		err = wm.completeWorkflow(ctx, workflowID)
		if err != nil {
			g.Log().Error(ctx, "Failed to mark workflow as completed", g.Map{
				"workflow_id": workflowID,
				"error":       err,
			})
		}
	}

	return nil
}

// FailWorkflow marks a workflow as failed
func (wm *SimpleWorkflowManager) FailWorkflow(ctx context.Context, workflowID, errorMessage string) error {
	return wm.failWorkflowExecution(ctx, workflowID, errorMessage)
}

// GetWorkflowStatus returns the current status of a workflow
func (wm *SimpleWorkflowManager) GetWorkflowStatus(ctx context.Context, workflowID string) (*WorkflowExecution, error) {
	return wm.getWorkflowExecution(ctx, workflowID)
}

// enqueueStep enqueues a single workflow step
func (wm *SimpleWorkflowManager) enqueueStep(ctx context.Context, workflowID, workflowType, stepName string, workflowInput, stepInput map[string]interface{}, orgID, userID string) error {
	// Get step definition for queue configuration
	workflow := wm.workflows[workflowType]
	var stepDef *StepDefinition
	for _, step := range workflow.Steps {
		if step.Name == stepName {
			stepDef = &step
			break
		}
	}

	queue := "default"
	if stepDef != nil && stepDef.Queue != "" {
		queue = stepDef.Queue
	}

	// Create concrete args based on step name to match adapter expectations
	var err error
	baseArgs := BaseJobArgs{
		WorkflowID:   workflowID,
		WorkflowType: workflowType,
		StepName:     stepName,
		OrgID:        orgID,
		UserID:       userID,
	}

	switch stepName {
	case "validate":
		args := ValidateStepArgs{
			BaseJobArgs: baseArgs,
			SignupData:  workflowInput,
		}
		_, err = wm.riverManager.Client.Insert(ctx, args, &river.InsertOpts{Queue: queue})
	case "create_user":
		args := CreateUserStepArgs{
			BaseJobArgs: baseArgs,
			UserData:    workflowInput,
		}
		_, err = wm.riverManager.Client.Insert(ctx, args, &river.InsertOpts{Queue: queue})
	case "create_organization":
		args := CreateOrganizationStepArgs{
			BaseJobArgs: baseArgs,
			OrgData:     workflowInput,
		}
		_, err = wm.riverManager.Client.Insert(ctx, args, &river.InsertOpts{Queue: queue})
	case "setup_stripe":
		args := SetupStripeStepArgs{
			BaseJobArgs: baseArgs,
			StripeData:  workflowInput,
		}
		_, err = wm.riverManager.Client.Insert(ctx, args, &river.InsertOpts{Queue: queue})
	case "send_verification":
		args := SendVerificationStepArgs{
			BaseJobArgs:      baseArgs,
			VerificationData: workflowInput,
		}
		_, err = wm.riverManager.Client.Insert(ctx, args, &river.InsertOpts{Queue: queue})
	default:
		// Fallback to generic StepArgs for unknown steps
		stepArgs := StepArgs{
			WorkflowID:    workflowID,
			WorkflowType:  workflowType,
			StepName:      stepName,
			OrgID:         orgID,
			UserID:        userID,
			WorkflowInput: workflowInput,
			StepInput:     stepInput,
		}
		_, err = wm.riverManager.Client.Insert(ctx, stepArgs, &river.InsertOpts{Queue: queue})
	}

	return err
}

// Database operations for workflow state management
func (wm *SimpleWorkflowManager) createWorkflowExecution(ctx context.Context, workflowID, workflowType, orgID, userID string, input map[string]interface{}) error {
	inputBytes, _ := json.Marshal(input)
	argsHash := fmt.Sprintf("%x", md5.Sum(inputBytes))

	query := `
		INSERT INTO workflow_executions 
		(workflow_id, workflow_type, org_id, user_id, status, context, args_hash, started_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, 'running', $5, $6, NOW(), NOW(), NOW())
	`

	contextBytes, _ := json.Marshal(input)

	var orgIDPtr *string
	if orgID != "" {
		orgIDPtr = &orgID
	}

	var userIDPtr *string
	if userID != "" {
		userIDPtr = &userID
	}

	_, err := wm.dbPool.Exec(ctx, query, workflowID, workflowType, orgIDPtr, userIDPtr, contextBytes, argsHash)
	return err
}

// checkForDuplicateWorkflow checks if there's already a running or pending workflow with the same parameters
func (wm *SimpleWorkflowManager) checkForDuplicateWorkflow(ctx context.Context, workflowType, orgID, userID, argsHash string) (string, error) {
	var query string
	var args []interface{}

	// Build query based on what parameters are provided
	baseQuery := `
		SELECT workflow_id FROM workflow_executions 
		WHERE workflow_type = $1 
		AND args_hash = $2 
		AND status IN ('running', 'pending')
	`

	args = []interface{}{workflowType, argsHash}
	paramIndex := 3

	// Add org_id condition if provided
	if orgID != "" {
		baseQuery += fmt.Sprintf(" AND org_id = $%d", paramIndex)
		args = append(args, orgID)
		paramIndex++
	} else {
		baseQuery += " AND org_id IS NULL"
	}

	// Add user_id condition if provided
	if userID != "" {
		baseQuery += fmt.Sprintf(" AND user_id = $%d", paramIndex)
		args = append(args, userID)
	} else {
		baseQuery += " AND user_id IS NULL"
	}

	query = baseQuery + " ORDER BY created_at DESC LIMIT 1"

	var workflowID string
	err := wm.dbPool.QueryRow(ctx, query, args...).Scan(&workflowID)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return "", nil // No duplicate found
		}
		return "", err
	}

	return workflowID, nil
}

func (wm *SimpleWorkflowManager) updateWorkflowStep(ctx context.Context, workflowID, currentStep string, context map[string]interface{}) error {
	contextBytes, _ := json.Marshal(context)

	query := `
		UPDATE workflow_executions 
		SET current_step = $1, context = $2, updated_at = NOW()
		WHERE workflow_id = $3
	`

	_, err := wm.dbPool.Exec(ctx, query, currentStep, contextBytes, workflowID)
	return err
}

func (wm *SimpleWorkflowManager) completeWorkflow(ctx context.Context, workflowID string) error {
	query := `
		UPDATE workflow_executions 
		SET status = 'completed', completed_at = NOW(), updated_at = NOW()
		WHERE workflow_id = $1
	`

	_, err := wm.dbPool.Exec(ctx, query, workflowID)
	return err
}

func (wm *SimpleWorkflowManager) failWorkflowExecution(ctx context.Context, workflowID, errorMessage string) error {
	query := `
		UPDATE workflow_executions 
		SET status = 'failed', error_message = $1, completed_at = NOW(), updated_at = NOW()
		WHERE workflow_id = $2
	`

	_, err := wm.dbPool.Exec(ctx, query, errorMessage, workflowID)
	return err
}

func (wm *SimpleWorkflowManager) getWorkflowExecution(ctx context.Context, workflowID string) (*WorkflowExecution, error) {
	query := `
		SELECT workflow_id, workflow_type, org_id, user_id, status, current_step, 
		       context, args_hash, started_at, completed_at, error_message, 
		       retry_count, created_at, updated_at
		FROM workflow_executions 
		WHERE workflow_id = $1
	`

	var execution WorkflowExecution
	var contextBytes []byte

	var orgID, userID sql.NullString

	err := wm.dbPool.QueryRow(ctx, query, workflowID).Scan(
		&execution.WorkflowID, &execution.WorkflowType, &orgID, &userID,
		&execution.Status, &execution.CurrentStep, &contextBytes,
		&execution.ArgsHash, &execution.StartedAt, &execution.CompletedAt,
		&execution.ErrorMessage, &execution.RetryCount, &execution.CreatedAt, &execution.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	// Handle nullable fields
	if orgID.Valid {
		execution.OrgID = orgID.String
	}
	if userID.Valid {
		execution.UserID = &userID.String
	}

	// Unmarshal JSON fields
	if len(contextBytes) > 0 {
		json.Unmarshal(contextBytes, &execution.Context)
	}

	return &execution, nil
}
