package riverjobs

import (
	"context"
	"time"
)

// WorkflowDefinition defines the structure and behavior of a workflow
type WorkflowDefinition interface {
	// GetName returns the workflow type name
	GetName() string

	// GetSteps returns all steps in this workflow in execution order
	GetSteps() []StepDefinition

	// GetInitialState returns the starting state for new workflows
	GetInitialState() WorkflowState

	// GetNextStep determines the next step based on current state and step result
	GetNextStep(currentState WorkflowState, stepResult StepResult) (NextAction, error)

	// HandleStepFailure determines how to handle a step failure
	HandleStepFailure(state WorkflowState, step StepDefinition, err error) (FailureAction, error)

	// ValidateInput validates the initial workflow input
	ValidateInput(ctx context.Context, input interface{}) error

	// GetCompensationSteps returns steps to run when rolling back
	GetCompensationSteps(failedStep string) []StepDefinition
}

// StepDefinition defines an individual step within a workflow
type StepDefinition interface {
	// GetName returns the step name
	GetName() string

	// GetQueue returns the queue this step should run in
	GetQueue() string

	// GetTimeout returns the maximum execution time for this step
	GetTimeout() time.Duration

	// GetRetryPolicy returns the retry configuration for this step
	GetRetryPolicy() RetryPolicy

	// Execute runs the step with the given input and returns result
	Execute(ctx context.Context, input StepInput) (StepResult, error)

	// Compensate runs compensation logic when rolling back
	Compensate(ctx context.Context, input StepInput) error

	// IsRetryable determines if an error is worth retrying
	IsRetryable(err error) bool

	// GetPreconditions returns conditions that must be met before running
	GetPreconditions() []Precondition
}

// WorkflowState represents the current state of a workflow
type WorkflowState string

const (
	// Common workflow states
	StateInitializing WorkflowState = "initializing"
	StatePending      WorkflowState = "pending"
	StateRunning      WorkflowState = "running"
	StateCompleted    WorkflowState = "completed"
	StateFailed       WorkflowState = "failed"
	StateCancelled    WorkflowState = "cancelled"
	StateCompensating WorkflowState = "compensating"

	// User signup specific states
	StateValidating          WorkflowState = "validating"
	StateCreatingUser        WorkflowState = "creating_user"
	StateCreatingOrg         WorkflowState = "creating_org"
	StateSettingUpPayment    WorkflowState = "setting_up_payment"
	StateSendingVerification WorkflowState = "sending_verification"
	StatePendingVerification WorkflowState = "pending_verification"
	StateVerified            WorkflowState = "verified"

	// Failure states
	StateValidationFailed   WorkflowState = "validation_failed"
	StateUserCreationFailed WorkflowState = "user_creation_failed"
	StateOrgCreationFailed  WorkflowState = "org_creation_failed"
	StatePaymentFailed      WorkflowState = "payment_failed"
	StateVerificationFailed WorkflowState = "verification_failed"
)

// NextAction defines what should happen after a step completes
type NextAction struct {
	Type      NextActionType `json:"type"`
	NextStep  string         `json:"next_step,omitempty"`
	NewState  WorkflowState  `json:"new_state,omitempty"`
	Delay     time.Duration  `json:"delay,omitempty"`
	Condition string         `json:"condition,omitempty"`
}

// NextActionType defines the types of actions that can be taken
type NextActionType string

const (
	ActionContinue   NextActionType = "continue"   // Move to next step
	ActionComplete   NextActionType = "complete"   // Workflow is done
	ActionFail       NextActionType = "fail"       // Workflow failed
	ActionWait       NextActionType = "wait"       // Wait for external event
	ActionRetry      NextActionType = "retry"      // Retry current step
	ActionCompensate NextActionType = "compensate" // Start rollback
	ActionSkip       NextActionType = "skip"       // Skip to another step
)

// FailureAction defines how to handle step failures
type FailureAction struct {
	Type         FailureActionType `json:"type"`
	RetryAfter   time.Duration     `json:"retry_after,omitempty"`
	MaxRetries   int               `json:"max_retries,omitempty"`
	Compensate   bool              `json:"compensate,omitempty"`
	FailWorkflow bool              `json:"fail_workflow,omitempty"`
	NewState     WorkflowState     `json:"new_state,omitempty"`
}

// FailureActionType defines types of failure handling
type FailureActionType string

const (
	FailureRetry      FailureActionType = "retry"
	FailureCompensate FailureActionType = "compensate"
	FailureFail       FailureActionType = "fail"
	FailureIgnore     FailureActionType = "ignore"
)

// StepInput contains input data for a step
type StepInput struct {
	WorkflowID      string                 `json:"workflow_id"`
	StepName        string                 `json:"step_name"`
	WorkflowContext map[string]interface{} `json:"workflow_context"`
	StepContext     map[string]interface{} `json:"step_context"`
	Metadata        map[string]interface{} `json:"metadata"`
}

// StepResult contains the result of step execution
type StepResult struct {
	Success     bool                   `json:"success"`
	OutputData  map[string]interface{} `json:"output_data"`
	NextState   WorkflowState          `json:"next_state,omitempty"`
	Error       string                 `json:"error,omitempty"`
	ShouldRetry bool                   `json:"should_retry"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// RetryPolicy defines retry behavior for a step
type RetryPolicy struct {
	MaxRetries      int           `json:"max_retries"`
	InitialInterval time.Duration `json:"initial_interval"`
	MaxInterval     time.Duration `json:"max_interval"`
	Multiplier      float64       `json:"multiplier"`
	RandomizeDelay  bool          `json:"randomize_delay"`
}

// Precondition defines a condition that must be met before step execution
type Precondition struct {
	Name      string `json:"name"`
	Condition string `json:"condition"`
	ErrorMsg  string `json:"error_msg"`
}

// DefaultRetryPolicy returns a sensible default retry policy
func DefaultRetryPolicy() RetryPolicy {
	return RetryPolicy{
		MaxRetries:      3,
		InitialInterval: 1 * time.Minute,
		MaxInterval:     10 * time.Minute,
		Multiplier:      2.0,
		RandomizeDelay:  true,
	}
}

// AggressiveRetryPolicy returns a retry policy for critical operations
func AggressiveRetryPolicy() RetryPolicy {
	return RetryPolicy{
		MaxRetries:      5,
		InitialInterval: 30 * time.Second,
		MaxInterval:     5 * time.Minute,
		Multiplier:      1.5,
		RandomizeDelay:  true,
	}
}

// NoRetryPolicy returns a policy that doesn't retry
func NoRetryPolicy() RetryPolicy {
	return RetryPolicy{
		MaxRetries:      0,
		InitialInterval: 0,
		MaxInterval:     0,
		Multiplier:      1.0,
		RandomizeDelay:  false,
	}
}
