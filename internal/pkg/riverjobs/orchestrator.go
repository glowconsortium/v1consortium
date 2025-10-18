package riverjobs

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/riverqueue/river"
)

// WorkflowOrchestrator orchestrates multi-step workflows
type WorkflowOrchestrator struct {
	RiverManager     *RiverManager
	WorkflowManager  *WorkflowManager
	ClientManager    *ClientManager
	Logger           *slog.Logger
	workflowRegistry map[string]WorkflowDefinition
	stepRegistry     map[string]StepDefinition
}

// NewWorkflowOrchestrator creates a new workflow orchestrator
func NewWorkflowOrchestrator(riverManager *RiverManager, workflowManager *WorkflowManager, logger *slog.Logger) *WorkflowOrchestrator {
	clientManager := NewClientManager(riverManager.Client, logger)

	return &WorkflowOrchestrator{
		RiverManager:     riverManager,
		WorkflowManager:  workflowManager,
		ClientManager:    clientManager,
		Logger:           logger,
		workflowRegistry: make(map[string]WorkflowDefinition),
		stepRegistry:     make(map[string]StepDefinition),
	}
}

// StartWorkflow initiates a new workflow
func (o *WorkflowOrchestrator) StartWorkflow(ctx context.Context, workflowType string, initialArgs JobArgs, protocol string) (string, error) {
	// Extract base args for workflow execution
	baseArgs, ok := initialArgs.(BaseJobArgs)
	if !ok {
		return "", fmt.Errorf("initial args must embed BaseJobArgs")
	}

	// Generate content hash for deduplication
	argsJSON, err := json.Marshal(initialArgs)
	if err != nil {
		return "", fmt.Errorf("failed to marshal job args: %w", err)
	}
	argsHashBytes := sha256.Sum256(argsJSON)
	argsHash := hex.EncodeToString(argsHashBytes[:])

	// Check if same data already exists for this workflow type and org
	existingWorkflowID, err := o.WorkflowManager.CheckWorkflowExists(ctx, workflowType, baseArgs.OrgID, argsHash)
	if err != nil {
		return "", fmt.Errorf("failed to check existing workflow: %w", err)
	}
	if existingWorkflowID != "" {
		return existingWorkflowID, fmt.Errorf("workflow with same data already exists: %s", existingWorkflowID)
	}

	// Generate unique workflow ID
	workflowID := fmt.Sprintf("%s_%d_%s", workflowType, time.Now().UnixNano(), baseArgs.OrgID)

	// Create workflow execution record
	execution := &WorkflowExecution{
		WorkflowType: workflowType,
		WorkflowID:   workflowID,
		OrgID:        baseArgs.OrgID,
		Status:       StatusPending,
		TotalSteps:   1, // Will be updated as workflow progresses
		Context:      make(map[string]interface{}),
		Metadata:     make(map[string]interface{}),
		ArgsHash:     argsHash,
		StartedAt:    time.Now(),
	}

	if baseArgs.UserID != "" {
		execution.UserID = &baseArgs.UserID
	}

	err = o.WorkflowManager.CreateWorkflowExecution(ctx, execution)
	if err != nil {
		return "", fmt.Errorf("failed to create workflow execution: %w", err)
	}

	// Get appropriate client and enqueue first job
	client := o.ClientManager.GetClient(protocol)
	_, err = client.Insert(ctx, initialArgs, &river.InsertOpts{
		Priority: 1,
		Queue:    river.QueueDefault,
	})

	if err != nil {
		// Update workflow status to failed
		errorMsg := err.Error()
		o.WorkflowManager.UpdateWorkflowStatus(ctx, workflowID, StatusFailed, &errorMsg)
		return "", fmt.Errorf("failed to enqueue initial job: %w", err)
	}

	// Update workflow status to running
	o.WorkflowManager.UpdateWorkflowStatus(ctx, workflowID, StatusRunning, nil)

	o.Logger.Info("workflow started", "workflow_id", workflowID, "type", workflowType, "protocol", protocol)
	return workflowID, nil
}

// GetWorkflowStatus retrieves the current status of a workflow
func (o *WorkflowOrchestrator) GetWorkflowStatus(ctx context.Context, workflowID string) (*WorkflowExecution, error) {
	return o.WorkflowManager.GetWorkflowExecution(ctx, workflowID)
}

// GetWorkflowSteps retrieves all steps for a workflow
func (o *WorkflowOrchestrator) GetWorkflowSteps(ctx context.Context, workflowID string) ([]WorkflowStep, error) {
	return o.WorkflowManager.GetWorkflowSteps(ctx, workflowID)
}

// CompleteWorkflow marks a workflow as completed
func (o *WorkflowOrchestrator) CompleteWorkflow(ctx context.Context, workflowID string) error {
	err := o.WorkflowManager.CompleteWorkflow(ctx, workflowID)
	if err != nil {
		return fmt.Errorf("failed to complete workflow: %w", err)
	}

	o.Logger.Info("workflow completed", "workflow_id", workflowID)
	return nil
}

// FailWorkflow marks a workflow as failed
func (o *WorkflowOrchestrator) FailWorkflow(ctx context.Context, workflowID string, errorMsg string) error {
	err := o.WorkflowManager.FailWorkflow(ctx, workflowID, errorMsg)
	if err != nil {
		return fmt.Errorf("failed to fail workflow: %w", err)
	}

	o.Logger.Error("workflow failed", "workflow_id", workflowID, "error", errorMsg)
	return nil
}

// CancelWorkflow marks a workflow as cancelled
func (o *WorkflowOrchestrator) CancelWorkflow(ctx context.Context, workflowID string) error {
	err := o.WorkflowManager.CancelWorkflow(ctx, workflowID)
	if err != nil {
		return fmt.Errorf("failed to cancel workflow: %w", err)
	}

	o.Logger.Info("workflow cancelled", "workflow_id", workflowID)
	return nil
}

// RestartWorkflow restarts a failed or cancelled workflow
func (o *WorkflowOrchestrator) RestartWorkflow(ctx context.Context, workflowID string, fromStep string) error {
	// Get current workflow execution
	execution, err := o.WorkflowManager.GetWorkflowExecution(ctx, workflowID)
	if err != nil {
		return fmt.Errorf("failed to get workflow execution: %w", err)
	}

	// Only allow restart for failed or cancelled workflows
	if execution.Status != StatusFailed && execution.Status != StatusCancelled {
		return fmt.Errorf("workflow can only be restarted from failed or cancelled state, current status: %s", execution.Status)
	}

	// Update workflow status to running
	err = o.WorkflowManager.UpdateWorkflowStatus(ctx, workflowID, StatusRunning, nil)
	if err != nil {
		return fmt.Errorf("failed to update workflow status: %w", err)
	}

	// TODO: Implement logic to restart from specific step
	// This would involve re-enqueuing the appropriate job

	o.Logger.Info("workflow restarted", "workflow_id", workflowID, "from_step", fromStep)
	return nil
}

// SetRPCClient sets the RPC client for the orchestrator
func (o *WorkflowOrchestrator) SetRPCClient(client RPCJobClient) {
	o.ClientManager.SetRPCClient(client)
}

// GetWorkflowProgress calculates the progress of a workflow
func (o *WorkflowOrchestrator) GetWorkflowProgress(ctx context.Context, workflowID string) (float64, error) {
	steps, err := o.WorkflowManager.GetWorkflowSteps(ctx, workflowID)
	if err != nil {
		return 0, fmt.Errorf("failed to get workflow steps: %w", err)
	}

	if len(steps) == 0 {
		return 0, nil
	}

	completedSteps := 0
	for _, step := range steps {
		if step.Status == StepStatusCompleted {
			completedSteps++
		}
	}

	return float64(completedSteps) / float64(len(steps)), nil
}

// GetWorkflowSummary returns a comprehensive summary of a workflow
func (o *WorkflowOrchestrator) GetWorkflowSummary(ctx context.Context, workflowID string) (*WorkflowSummary, error) {
	execution, err := o.WorkflowManager.GetWorkflowExecution(ctx, workflowID)
	if err != nil {
		return nil, fmt.Errorf("failed to get workflow execution: %w", err)
	}

	steps, err := o.WorkflowManager.GetWorkflowSteps(ctx, workflowID)
	if err != nil {
		return nil, fmt.Errorf("failed to get workflow steps: %w", err)
	}

	progress, _ := o.GetWorkflowProgress(ctx, workflowID)

	summary := &WorkflowSummary{
		Execution: *execution,
		Steps:     steps,
		Progress:  progress,
	}

	return summary, nil
}

// WorkflowSummary provides a complete view of a workflow
type WorkflowSummary struct {
	Execution WorkflowExecution `json:"execution"`
	Steps     []WorkflowStep    `json:"steps"`
	Progress  float64           `json:"progress"`
}

// RegisterWorkflow registers a workflow definition
func (o *WorkflowOrchestrator) RegisterWorkflow(definition WorkflowDefinition) {
	o.workflowRegistry[definition.GetName()] = definition

	// Register all steps in this workflow
	for _, step := range definition.GetSteps() {
		o.stepRegistry[step.GetName()] = step
	}

	o.Logger.Info("registered workflow", "type", definition.GetName(), "steps", len(definition.GetSteps()))
}

// RegisterStep registers an individual step definition
func (o *WorkflowOrchestrator) RegisterStep(step StepDefinition) {
	o.stepRegistry[step.GetName()] = step
	o.Logger.Info("registered step", "name", step.GetName())
}

// StartWorkflowWithDefinition starts a workflow using a registered workflow definition
func (o *WorkflowOrchestrator) StartWorkflowWithDefinition(ctx context.Context, workflowType string, input interface{}, orgID, userID string) (string, error) {
	// Get workflow definition
	definition, exists := o.workflowRegistry[workflowType]
	if !exists {
		return "", fmt.Errorf("workflow definition not found: %s", workflowType)
	}

	// Validate input
	if err := definition.ValidateInput(ctx, input); err != nil {
		return "", fmt.Errorf("input validation failed: %w", err)
	}

	// Generate workflow ID
	workflowID := fmt.Sprintf("%s_%d_%s", workflowType, time.Now().UnixNano(), orgID)

	// Create workflow execution record
	execution := &WorkflowExecution{
		WorkflowType: workflowType,
		WorkflowID:   workflowID,
		OrgID:        orgID,
		Status:       string(definition.GetInitialState()),
		TotalSteps:   len(definition.GetSteps()),
		Context:      make(map[string]interface{}),
		Metadata:     make(map[string]interface{}),
		StartedAt:    time.Now(),
	}

	if userID != "" {
		execution.UserID = &userID
	}

	// Store input in context
	inputJSON, _ := json.Marshal(input)
	var inputMap map[string]interface{}
	json.Unmarshal(inputJSON, &inputMap)
	execution.Context["input"] = inputMap

	// Create workflow execution
	if err := o.WorkflowManager.CreateWorkflowExecution(ctx, execution); err != nil {
		return "", fmt.Errorf("failed to create workflow execution: %w", err)
	}

	// Start first step
	steps := definition.GetSteps()
	if len(steps) > 0 {
		if err := o.executeNextStep(ctx, workflowID, steps[0], inputMap); err != nil {
			o.FailWorkflow(ctx, workflowID, err.Error())
			return "", fmt.Errorf("failed to start first step: %w", err)
		}
	}

	o.Logger.Info("started workflow with definition", "workflow_id", workflowID, "type", workflowType)
	return workflowID, nil
}

// ExecuteStep executes a specific step within a workflow
func (o *WorkflowOrchestrator) ExecuteStep(ctx context.Context, workflowID, stepName string, input StepInput) (StepResult, error) {
	// Get step definition
	stepDef, exists := o.stepRegistry[stepName]
	if !exists {
		return StepResult{}, fmt.Errorf("step definition not found: %s", stepName)
	}

	// Get workflow definition for context
	execution, err := o.WorkflowManager.GetWorkflowExecution(ctx, workflowID)
	if err != nil {
		return StepResult{}, fmt.Errorf("failed to get workflow execution: %w", err)
	}

	workflowDef, exists := o.workflowRegistry[execution.WorkflowType]
	if !exists {
		return StepResult{}, fmt.Errorf("workflow definition not found: %s", execution.WorkflowType)
	}

	// Update step status to running
	o.WorkflowManager.UpdateStepStatus(ctx, workflowID, stepName, StepStatusRunning, nil, nil)

	// Execute the step
	result, err := stepDef.Execute(ctx, input)
	if err != nil {
		// Handle step failure
		stepError := &StepError{
			Type:      ErrorTypeBusiness,
			Message:   err.Error(),
			Retryable: stepDef.IsRetryable(err),
			Cause:     err,
		}

		// Get failure action from workflow definition
		failureAction, actionErr := workflowDef.HandleStepFailure(WorkflowState(execution.Status), stepDef, stepError)
		if actionErr != nil {
			o.Logger.Error("failed to get failure action", "error", actionErr)
			failureAction = FailureAction{Type: FailureFail, FailWorkflow: true}
		}

		// Update step status to failed
		errorMsg := stepError.Error()
		o.WorkflowManager.UpdateStepStatus(ctx, workflowID, stepName, StepStatusFailed, nil, &errorMsg)

		// Handle based on failure action
		switch failureAction.Type {
		case FailureRetry:
			if failureAction.RetryAfter > 0 {
				// Schedule retry (would need additional logic for delayed retry)
				return StepResult{ShouldRetry: true}, stepError
			}
			return StepResult{ShouldRetry: true}, stepError
		case FailureCompensate:
			// Start compensation workflow
			if err := o.startCompensation(ctx, workflowID, stepName, workflowDef); err != nil {
				o.Logger.Error("failed to start compensation", "error", err)
			}
			return StepResult{Success: false}, stepError
		case FailureFail:
			if failureAction.FailWorkflow {
				o.FailWorkflow(ctx, workflowID, stepError.Error())
			}
			return StepResult{Success: false}, stepError
		}
	}

	// Update step status to completed
	o.WorkflowManager.UpdateStepStatus(ctx, workflowID, stepName, StepStatusCompleted, result.OutputData, nil)

	// Determine next action
	nextAction, err := workflowDef.GetNextStep(WorkflowState(execution.Status), result)
	if err != nil {
		return result, fmt.Errorf("failed to determine next step: %w", err)
	}

	// Handle next action
	if err := o.handleNextAction(ctx, workflowID, nextAction, result.OutputData, workflowDef); err != nil {
		return result, fmt.Errorf("failed to handle next action: %w", err)
	}

	return result, nil
}

// executeNextStep executes the next step in a workflow
func (o *WorkflowOrchestrator) executeNextStep(ctx context.Context, workflowID string, step StepDefinition, context map[string]interface{}) error {
	// Create step input
	stepInput := StepInput{
		WorkflowID:      workflowID,
		StepName:        step.GetName(),
		WorkflowContext: context,
		StepContext:     make(map[string]interface{}),
		Metadata:        make(map[string]interface{}),
	}

	// Create step record
	_, err := o.WorkflowManager.CreateWorkflowStep(ctx, workflowID, step.GetName(), stepInput.StepContext)
	if err != nil {
		return fmt.Errorf("failed to create step record: %w", err)
	}

	// Execute step (in real implementation, this would enqueue a job)
	_, err = o.ExecuteStep(ctx, workflowID, step.GetName(), stepInput)
	return err
}

// handleNextAction processes the next action after step completion
func (o *WorkflowOrchestrator) handleNextAction(ctx context.Context, workflowID string, action NextAction, outputData map[string]interface{}, workflowDef WorkflowDefinition) error {
	switch action.Type {
	case ActionComplete:
		return o.CompleteWorkflow(ctx, workflowID)
	case ActionFail:
		return o.FailWorkflow(ctx, workflowID, "workflow failed due to step result")
	case ActionContinue:
		if action.NextStep != "" {
			// Find next step definition
			steps := workflowDef.GetSteps()
			for _, step := range steps {
				if step.GetName() == action.NextStep {
					return o.executeNextStep(ctx, workflowID, step, outputData)
				}
			}
			return fmt.Errorf("next step not found: %s", action.NextStep)
		}
	case ActionWait:
		// Update workflow to waiting state
		return o.WorkflowManager.UpdateWorkflowStatus(ctx, workflowID, string(action.NewState), nil)
	}

	return nil
}

// startCompensation starts the compensation workflow
func (o *WorkflowOrchestrator) startCompensation(ctx context.Context, workflowID, failedStep string, workflowDef WorkflowDefinition) error {
	// Update workflow status to compensating
	err := o.WorkflowManager.UpdateWorkflowStatus(ctx, workflowID, string(StateCompensating), nil)
	if err != nil {
		return fmt.Errorf("failed to update workflow status to compensating: %w", err)
	}

	// Get compensation steps
	compensationSteps := workflowDef.GetCompensationSteps(failedStep)

	// Execute compensation steps in reverse order
	for i := len(compensationSteps) - 1; i >= 0; i-- {
		step := compensationSteps[i]
		stepInput := StepInput{
			WorkflowID:      workflowID,
			StepName:        step.GetName(),
			WorkflowContext: make(map[string]interface{}),
			StepContext:     map[string]interface{}{"compensation": true},
		}

		if err := step.Compensate(ctx, stepInput); err != nil {
			o.Logger.Error("compensation step failed", "step", step.GetName(), "error", err)
			// Continue with other compensation steps even if one fails
		}
	}

	return nil
}
