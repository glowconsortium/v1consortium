package signup

import (
	"context"
	"fmt"
	"strings"
	"time"
	"v1consortium/internal/consts"
	"v1consortium/internal/pkg/riverjobs"
	"v1consortium/internal/workflow/signup/steps"

	"github.com/gogf/gf/v2/frame/g"
)

// SignupWorkflow implements the user signup workflow
type SignupWorkflow struct {
	stepRegistry map[string]riverjobs.StepDefinition
}

// SignupInput represents the input for the signup workflow
type SignupInput struct {
	Email            string                 `json:"email" v:"required|email"`
	Password         string                 `json:"password" v:"required|min-length:8"`
	FirstName        string                 `json:"first_name" v:"required|min-length:2"`
	LastName         string                 `json:"last_name" v:"required|min-length:2"`
	Role             string                 `json:"role" v:"required|in:driver,admin,owner"`
	OrganizationData map[string]interface{} `json:"organization_data"`
	Metadata         map[string]interface{} `json:"metadata"`
}

// NewSignupWorkflow creates a new signup workflow
func NewSignupWorkflow() *SignupWorkflow {
	workflow := &SignupWorkflow{
		stepRegistry: make(map[string]riverjobs.StepDefinition),
	}

	// Register all steps
	workflow.stepRegistry["validate"] = steps.NewValidateStep()
	workflow.stepRegistry["create_user"] = steps.NewCreateUserStep()
	workflow.stepRegistry["create_organization"] = steps.NewCreateOrganizationStep()
	workflow.stepRegistry["setup_stripe"] = steps.NewSetupStripeStep()
	workflow.stepRegistry["send_verification"] = steps.NewSendVerificationStep()

	return workflow
}

// GetName returns the workflow name
func (w *SignupWorkflow) GetName() string {
	return "user_signup"
}

// GetSteps returns all steps in the workflow
func (w *SignupWorkflow) GetSteps() []riverjobs.StepDefinition {
	return []riverjobs.StepDefinition{
		w.stepRegistry["validate"],
		w.stepRegistry["create_user"],
		w.stepRegistry["create_organization"],
		w.stepRegistry["setup_stripe"],
		w.stepRegistry["send_verification"],
	}
}

// GetInitialState returns the starting state
func (w *SignupWorkflow) GetInitialState() riverjobs.WorkflowState {
	return riverjobs.StateValidating
}

// ValidateInput validates the workflow input
func (w *SignupWorkflow) ValidateInput(ctx context.Context, input interface{}) error {
	signupInput, ok := input.(SignupInput)
	if !ok {
		if inputMap, ok := input.(map[string]interface{}); ok {
			// Convert map to struct
			if err := g.Try(ctx, func(ctx context.Context) {
				g.Validator().Data(inputMap).Run(ctx)
			}); err != nil {
				return fmt.Errorf("input validation failed: %w", err)
			}
		} else {
			return fmt.Errorf("invalid input type, expected SignupInput")
		}
	}

	// Additional business validation
	if signupInput.Email == "" {
		return fmt.Errorf("email is required")
	}

	if len(signupInput.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}

	// validRoles := map[string]bool{"driver": true, "admin": true, "owner": true}
	// if !validRoles[signupInput.Role] {
	// 	return fmt.Errorf("invalid role: %s", signupInput.Role)
	// }
	// Validate role

	validRoles := map[string]bool{
		string(consts.RoleClientAdmin): true,
		string(consts.RoleEmployee):    true,
		string(consts.RoleHRManager):   true,
	}
	if !validRoles[signupInput.Role] {
		validList := []string{
			string(consts.RoleClientAdmin),
			string(consts.RoleEmployee),
			string(consts.RoleHRManager),
		}
		return fmt.Errorf("invalid role '%s', must be one of: %s", signupInput.Role, strings.Join(validList, ", "))
	}
	return nil
}

// GetNextStep determines the next step based on current state and step result
func (w *SignupWorkflow) GetNextStep(currentState riverjobs.WorkflowState, stepResult riverjobs.StepResult) (riverjobs.NextAction, error) {
	if !stepResult.Success {
		return w.handleStepFailure(currentState, stepResult)
	}

	switch currentState {
	case riverjobs.StateValidating:
		return riverjobs.NextAction{
			Type:     riverjobs.ActionContinue,
			NextStep: "create_user",
			NewState: riverjobs.StateCreatingUser,
		}, nil

	case riverjobs.StateCreatingUser:
		return riverjobs.NextAction{
			Type:     riverjobs.ActionContinue,
			NextStep: "create_organization",
			NewState: riverjobs.StateCreatingOrg,
		}, nil

	case riverjobs.StateCreatingOrg:
		return riverjobs.NextAction{
			Type:     riverjobs.ActionContinue,
			NextStep: "setup_stripe",
			NewState: riverjobs.StateSettingUpPayment,
		}, nil

	case riverjobs.StateSettingUpPayment:
		return riverjobs.NextAction{
			Type:     riverjobs.ActionContinue,
			NextStep: "send_verification",
			NewState: riverjobs.StateSendingVerification,
		}, nil

	case riverjobs.StateSendingVerification:
		return riverjobs.NextAction{
			Type:     riverjobs.ActionComplete,
			NewState: riverjobs.StatePendingVerification,
		}, nil

	default:
		return riverjobs.NextAction{}, fmt.Errorf("unknown state: %s", currentState)
	}
}

// handleStepFailure determines next action for failed steps
func (w *SignupWorkflow) handleStepFailure(currentState riverjobs.WorkflowState, stepResult riverjobs.StepResult) (riverjobs.NextAction, error) {
	switch currentState {
	case riverjobs.StateValidating:
		return riverjobs.NextAction{
			Type:     riverjobs.ActionFail,
			NewState: riverjobs.StateValidationFailed,
		}, nil

	case riverjobs.StateCreatingUser:
		return riverjobs.NextAction{
			Type:     riverjobs.ActionFail,
			NewState: riverjobs.StateUserCreationFailed,
		}, nil

	case riverjobs.StateCreatingOrg:
		// If org creation fails, we need to cleanup the user
		return riverjobs.NextAction{
			Type:     riverjobs.ActionCompensate,
			NewState: riverjobs.StateOrgCreationFailed,
		}, nil

	case riverjobs.StateSettingUpPayment:
		// If payment setup fails, cleanup org and user
		return riverjobs.NextAction{
			Type:     riverjobs.ActionCompensate,
			NewState: riverjobs.StatePaymentFailed,
		}, nil

	case riverjobs.StateSendingVerification:
		// If verification fails, we can retry or mark as needing manual intervention
		if stepResult.ShouldRetry {
			return riverjobs.NextAction{
				Type:  riverjobs.ActionRetry,
				Delay: 5 * time.Minute,
			}, nil
		}
		return riverjobs.NextAction{
			Type:     riverjobs.ActionFail,
			NewState: riverjobs.StateVerificationFailed,
		}, nil

	default:
		return riverjobs.NextAction{
			Type:     riverjobs.ActionFail,
			NewState: riverjobs.StateFailed,
		}, nil
	}
}

// HandleStepFailure determines how to handle step failures
func (w *SignupWorkflow) HandleStepFailure(state riverjobs.WorkflowState, step riverjobs.StepDefinition, err error) (riverjobs.FailureAction, error) {
	stepError, ok := err.(*riverjobs.StepError)
	if !ok {
		// Convert regular error to step error
		stepError = riverjobs.NewStepError(riverjobs.ErrorTypeBusiness, err.Error(), false, err)
	}

	switch step.GetName() {
	case "validate":
		// Validation errors are not retryable
		return riverjobs.FailureAction{
			Type:         riverjobs.FailureFail,
			FailWorkflow: true,
			NewState:     riverjobs.StateValidationFailed,
		}, nil

	case "create_user":
		if stepError.Type == riverjobs.ErrorTypeNetwork || stepError.Type == riverjobs.ErrorTypeDatabase {
			return riverjobs.FailureAction{
				Type:       riverjobs.FailureRetry,
				RetryAfter: 1 * time.Minute,
				MaxRetries: 3,
			}, nil
		}
		return riverjobs.FailureAction{
			Type:         riverjobs.FailureFail,
			FailWorkflow: true,
			NewState:     riverjobs.StateUserCreationFailed,
		}, nil

	case "create_organization":
		if stepError.Type == riverjobs.ErrorTypeNetwork || stepError.Type == riverjobs.ErrorTypeDatabase {
			return riverjobs.FailureAction{
				Type:       riverjobs.FailureRetry,
				RetryAfter: 1 * time.Minute,
				MaxRetries: 3,
			}, nil
		}
		// Business logic errors require compensation
		return riverjobs.FailureAction{
			Type:         riverjobs.FailureCompensate,
			Compensate:   true,
			FailWorkflow: true,
			NewState:     riverjobs.StateOrgCreationFailed,
		}, nil

	case "setup_stripe":
		if stepError.Type == riverjobs.ErrorTypeNetwork || stepError.Type == riverjobs.ErrorTypeExternal {
			return riverjobs.FailureAction{
				Type:       riverjobs.FailureRetry,
				RetryAfter: 2 * time.Minute,
				MaxRetries: 5,
			}, nil
		}
		return riverjobs.FailureAction{
			Type:         riverjobs.FailureCompensate,
			Compensate:   true,
			FailWorkflow: true,
			NewState:     riverjobs.StatePaymentFailed,
		}, nil

	case "send_verification":
		// Verification email failures are usually retryable
		return riverjobs.FailureAction{
			Type:       riverjobs.FailureRetry,
			RetryAfter: 5 * time.Minute,
			MaxRetries: 3,
		}, nil

	default:
		return riverjobs.FailureAction{
			Type:         riverjobs.FailureFail,
			FailWorkflow: true,
		}, nil
	}
}

// GetCompensationSteps returns steps for rollback
func (w *SignupWorkflow) GetCompensationSteps(failedStep string) []riverjobs.StepDefinition {
	switch failedStep {
	case "create_organization":
		// If org creation failed, cleanup user
		return []riverjobs.StepDefinition{
			steps.NewCleanupUserStep(),
		}

	case "setup_stripe":
		// If payment setup failed, cleanup org and user
		return []riverjobs.StepDefinition{
			steps.NewCleanupOrganizationStep(),
			steps.NewCleanupUserStep(),
		}

	case "send_verification":
		// For verification failures, we might not want to cleanup everything
		// Just mark as needing manual intervention
		return []riverjobs.StepDefinition{}

	default:
		return []riverjobs.StepDefinition{}
	}
}
