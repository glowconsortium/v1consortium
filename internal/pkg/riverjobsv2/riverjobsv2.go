package riverjobsv2

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/rivertype"
)

// WorkflowContext represents the data that flows between workflow steps
type WorkflowContext map[string]interface{}

// Get retrieves a value from the context
func (wc WorkflowContext) Get(key string) (interface{}, bool) {
	val, exists := wc[key]
	return val, exists
}

// GetString retrieves a string value from the context
func (wc WorkflowContext) GetString(key string) (string, bool) {
	if val, exists := wc[key]; exists {
		if str, ok := val.(string); ok {
			return str, true
		}
	}
	return "", false
}

// Set adds or updates a value in the context
func (wc WorkflowContext) Set(key string, value interface{}) {
	wc[key] = value
}

// StepResult represents the result of executing a workflow step
type StepResult struct {
	Success      bool            `json:"success"`
	Data         WorkflowContext `json:"data"`          // Data to merge into workflow context
	NextStep     string          `json:"next_step"`     // Override next step (optional)
	ShouldRetry  bool            `json:"should_retry"`  // Whether to retry on failure
	ErrorMessage string          `json:"error_message"` // Error details if failed
}

// StepFunc is the function signature for executing a workflow step
type StepFunc func(ctx context.Context, input map[string]interface{}, workflowCtx WorkflowContext) (*StepResult, error)

// Step defines a workflow step with its execution function
type Step struct {
	Name       string        `json:"name"`
	Execute    StepFunc      `json:"-"`           // The actual execution function
	IsOptional bool          `json:"is_optional"` // Can be skipped on failure
	MaxRetries int           `json:"max_retries"` // Maximum retry attempts
	RetryDelay time.Duration `json:"retry_delay"` // Delay between retries
	Queue      string        `json:"queue"`       // Queue name for this step
	Timeout    time.Duration `json:"timeout"`     // Step timeout
}

// Workflow defines a complete workflow with its steps and flow logic
type Workflow struct {
	Name         string                                                        `json:"name"`
	Steps        []Step                                                        `json:"steps"`
	StepFlow     map[string]string                                             `json:"step_flow"` // stepName -> nextStepName
	FirstStep    string                                                        `json:"first_step"`
	ValidateFunc func(ctx context.Context, input map[string]interface{}) error `json:"-"`
}

// GetStep returns a step by name
func (w *Workflow) GetStep(name string) (*Step, bool) {
	for _, step := range w.Steps {
		if step.Name == name {
			return &step, true
		}
	}
	return nil, false
}

// GetNextStep returns the next step name after the current step
func (w *Workflow) GetNextStep(currentStep string) string {
	if nextStep, exists := w.StepFlow[currentStep]; exists {
		return nextStep
	}
	return "" // End of workflow
}

// Validate validates the workflow input
func (w *Workflow) Validate(ctx context.Context, input map[string]interface{}) error {
	if w.ValidateFunc != nil {
		return w.ValidateFunc(ctx, input)
	}
	return nil
}

// WorkflowArgs represents the River job arguments for workflow execution
type WorkflowArgs struct {
	WorkflowID   string                 `json:"workflow_id"`
	WorkflowName string                 `json:"workflow_name"`
	CurrentStep  string                 `json:"current_step"`
	Input        map[string]interface{} `json:"input"`         // Original workflow input
	Context      WorkflowContext        `json:"context"`       // Accumulated context data
	AttemptCount int                    `json:"attempt_count"` // Current attempt for this step
	InputHash    string                 `json:"input_hash"`    // Hash of input for deduplication
}

// WorkflowStatus represents the current state of a workflow
type WorkflowStatus struct {
	WorkflowID   string                 `json:"workflow_id"`
	WorkflowName string                 `json:"workflow_name"`
	Status       string                 `json:"status"` // "running", "completed", "failed"
	CurrentStep  string                 `json:"current_step"`
	Input        map[string]interface{} `json:"input"`
	Context      WorkflowContext        `json:"context"`
	StartedAt    time.Time              `json:"started_at"`
	CompletedAt  *time.Time             `json:"completed_at,omitempty"`
	ErrorMessage string                 `json:"error_message,omitempty"`
}

func (args WorkflowArgs) Kind() string {
	// Always return the same kind for all workflow jobs since
	// we have one WorkflowExecutor handling all workflow types
	return "workflow_execution"
}

// GenerateWorkflowID creates a new UUID for workflows
func GenerateWorkflowID() string {
	return uuid.New().String()
}

// GenerateInputHash creates a hash of the input for deduplication
func GenerateInputHash(workflowName string, input map[string]interface{}) (string, error) {
	// Create a consistent representation for hashing
	hashData := map[string]interface{}{
		"workflow_name": workflowName,
		"input":         input,
	}

	jsonBytes, err := json.Marshal(hashData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal input for hashing: %w", err)
	}

	hash := sha256.Sum256(jsonBytes)
	return hex.EncodeToString(hash[:]), nil
}

// WorkflowStore interface for workflow persistence and deduplication
type WorkflowStore interface {
	// SaveWorkflowStatus saves or updates workflow status
	SaveWorkflowStatus(ctx context.Context, status *WorkflowStatus) error

	// GetWorkflowStatus retrieves workflow status by ID
	GetWorkflowStatus(ctx context.Context, workflowID string) (*WorkflowStatus, error)

	// GetRunningWorkflowByHash finds a running workflow with the same input hash
	GetRunningWorkflowByHash(ctx context.Context, workflowName, inputHash string) (*WorkflowStatus, error)

	// MarkWorkflowCompleted marks a workflow as completed
	MarkWorkflowCompleted(ctx context.Context, workflowID string, finalContext WorkflowContext) error

	// MarkWorkflowFailed marks a workflow as failed
	MarkWorkflowFailed(ctx context.Context, workflowID string, errorMessage string) error
}

// WorkflowExecutor executes workflow steps
type WorkflowExecutor struct {
	river.WorkerDefaults[WorkflowArgs]
	workflows   map[string]*Workflow
	riverClient interface{}   // Will be set during initialization
	store       WorkflowStore // For workflow persistence and deduplication
}

// NewWorkflowExecutor creates a new workflow executor
func NewWorkflowExecutor(store WorkflowStore) *WorkflowExecutor {
	return &WorkflowExecutor{
		workflows: make(map[string]*Workflow),
		store:     store,
	}
}

// SetClient sets the River client for scheduling next steps
func (we *WorkflowExecutor) SetClient(client interface{}) {
	we.riverClient = client
}

// RegisterWorkflow registers a workflow with the executor
func (we *WorkflowExecutor) RegisterWorkflow(workflow *Workflow) {
	we.workflows[workflow.Name] = workflow
}

// Work executes a workflow step
func (we *WorkflowExecutor) Work(ctx context.Context, job *river.Job[WorkflowArgs]) error {
	args := job.Args

	// Get the workflow
	workflow, exists := we.workflows[args.WorkflowName]
	if !exists {
		return river.JobCancel(fmt.Errorf("workflow not found: %s", args.WorkflowName))
	}

	// Get the current step - if empty, use first step
	currentStep := args.CurrentStep
	if currentStep == "" {
		currentStep = workflow.FirstStep
		// Update workflow status to running for first step
		if we.store != nil {
			status := &WorkflowStatus{
				WorkflowID:   args.WorkflowID,
				WorkflowName: args.WorkflowName,
				Status:       "running",
				CurrentStep:  currentStep,
				Input:        args.Input,
				Context:      args.Context,
				StartedAt:    time.Now(),
			}
			we.store.SaveWorkflowStatus(ctx, status)
		}
	}

	step, exists := workflow.GetStep(currentStep)
	if !exists {
		if we.store != nil {
			we.store.MarkWorkflowFailed(ctx, args.WorkflowID, fmt.Sprintf("step not found: %s", currentStep))
		}
		return river.JobCancel(fmt.Errorf("step not found: %s", currentStep))
	}

	// Execute the step
	result, err := step.Execute(ctx, args.Input, args.Context)
	if err != nil {
		return we.handleStepError(ctx, workflow, step, args, currentStep, err)
	}

	// Merge step result data into workflow context
	if result.Data != nil {
		for k, v := range result.Data {
			args.Context.Set(k, v)
		}
	}

	// Update workflow status with current progress
	if we.store != nil {
		status := &WorkflowStatus{
			WorkflowID:   args.WorkflowID,
			WorkflowName: args.WorkflowName,
			Status:       "running",
			CurrentStep:  currentStep,
			Input:        args.Input,
			Context:      args.Context,
			StartedAt:    time.Now(), // This should be preserved from initial creation
		}
		we.store.SaveWorkflowStatus(ctx, status)
	}

	// Determine next step
	nextStep := result.NextStep
	if nextStep == "" {
		nextStep = workflow.GetNextStep(currentStep)
	}

	// Schedule next step or complete workflow
	if nextStep != "" {
		return we.scheduleNextStep(ctx, args, nextStep)
	}

	// Workflow completed successfully
	if we.store != nil {
		we.store.MarkWorkflowCompleted(ctx, args.WorkflowID, args.Context)
	}
	return nil
} // handleStepError handles step execution errors
func (we *WorkflowExecutor) handleStepError(ctx context.Context, workflow *Workflow, step *Step, args WorkflowArgs, currentStep string, err error) error {
	// Check if we should retry
	if args.AttemptCount < step.MaxRetries {
		// Return error to trigger River's built-in retry mechanism
		// We'll track attempts in the args
		return fmt.Errorf("step %s failed (attempt %d/%d): %w", step.Name, args.AttemptCount+1, step.MaxRetries, err)
	}

	// Check if step is optional
	if step.IsOptional {
		// Skip to next step
		nextStep := workflow.GetNextStep(currentStep)
		if nextStep != "" {
			return we.scheduleNextStep(ctx, args, nextStep)
		}
		return nil // Workflow completed despite optional step failure
	}

	// Step failed and is not optional - fail the workflow
	if we.store != nil {
		we.store.MarkWorkflowFailed(ctx, args.WorkflowID, fmt.Sprintf("step %s failed: %v", step.Name, err))
	}
	return err
}

// scheduleNextStep schedules the next step in the workflow
func (we *WorkflowExecutor) scheduleNextStep(ctx context.Context, currentArgs WorkflowArgs, nextStep string) error {
	if we.riverClient == nil {
		return fmt.Errorf("river client not set in workflow executor")
	}

	// Create args for the next step
	nextArgs := WorkflowArgs{
		WorkflowID:   currentArgs.WorkflowID,
		WorkflowName: currentArgs.WorkflowName,
		CurrentStep:  nextStep,
		Input:        currentArgs.Input,
		Context:      currentArgs.Context,
		AttemptCount: 0, // Reset attempt count for new step
		InputHash:    currentArgs.InputHash,
	}

	// Use type assertion to call Insert - handle the specific pgx.Tx River client type
	switch riverClient := we.riverClient.(type) {
	case *river.Client[pgx.Tx]:
		_, err := riverClient.Insert(ctx, nextArgs, nil)
		if err != nil {
			return fmt.Errorf("failed to schedule next step %s: %w", nextStep, err)
		}
		return nil
	case interface {
		Insert(context.Context, WorkflowArgs, *river.InsertOpts) (*rivertype.JobInsertResult, error)
	}:
		_, err := riverClient.Insert(ctx, nextArgs, nil)
		if err != nil {
			return fmt.Errorf("failed to schedule next step %s: %w", nextStep, err)
		}
		return nil
	default:
		return fmt.Errorf("invalid river client type in workflow executor - got type %T", we.riverClient)
	}
}

// StartWorkflowResult represents the result of starting a workflow
type StartWorkflowResult struct {
	WorkflowID     string `json:"workflow_id"`
	IsNewWorkflow  bool   `json:"is_new_workflow"` // true if new, false if existing found
	ExistingStatus string `json:"existing_status"` // status of existing workflow if found
}

// StartWorkflow initiates a new workflow execution with deduplication
func StartWorkflow(ctx context.Context, client interface{}, store WorkflowStore, workflowName string, input map[string]interface{}) (*StartWorkflowResult, error) {
	// Generate input hash for deduplication
	inputHash, err := GenerateInputHash(workflowName, input)
	if err != nil {
		return nil, fmt.Errorf("failed to generate input hash: %w", err)
	}

	// Check if a workflow with the same input is already running
	if store != nil {
		existingWorkflow, err := store.GetRunningWorkflowByHash(ctx, workflowName, inputHash)
		if err != nil {
			return nil, fmt.Errorf("failed to check for existing workflow: %w", err)
		}

		if existingWorkflow != nil {
			return &StartWorkflowResult{
				WorkflowID:     existingWorkflow.WorkflowID,
				IsNewWorkflow:  false,
				ExistingStatus: existingWorkflow.Status,
			}, nil
		}
	}

	// Generate new workflow ID
	workflowID := GenerateWorkflowID()

	args := WorkflowArgs{
		WorkflowID:   workflowID,
		WorkflowName: workflowName,
		CurrentStep:  "", // Will be determined by workflow's FirstStep
		Input:        input,
		Context:      make(WorkflowContext),
		AttemptCount: 0,
		InputHash:    inputHash,
	}

	// Use type assertion to call Insert - handle the specific pgx.Tx River client type
	switch riverClient := client.(type) {
	case *river.Client[pgx.Tx]:
		_, err := riverClient.Insert(ctx, args, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to insert workflow job: %w", err)
		}
		return &StartWorkflowResult{
			WorkflowID:    workflowID,
			IsNewWorkflow: true,
		}, nil
	case interface {
		Insert(context.Context, WorkflowArgs, *river.InsertOpts) (*rivertype.JobInsertResult, error)
	}:
		_, err := riverClient.Insert(ctx, args, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to insert workflow job: %w", err)
		}
		return &StartWorkflowResult{
			WorkflowID:    workflowID,
			IsNewWorkflow: true,
		}, nil
	default:
		return nil, fmt.Errorf("invalid river client provided - got type %T", client)
	}
}

// GetWorkflowStatus retrieves the status of a workflow
func GetWorkflowStatus(ctx context.Context, store WorkflowStore, workflowID string) (*WorkflowStatus, error) {
	if store == nil {
		return nil, fmt.Errorf("workflow store not provided")
	}

	return store.GetWorkflowStatus(ctx, workflowID)
}
