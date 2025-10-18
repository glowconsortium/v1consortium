package riverjobusersignup

import (
	"context"
	"fmt"
	"time"
	"v1consortium/internal/pkg/riverjobs"
	"v1consortium/internal/workflow/signup"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/rivertype"
)

// sUserSignupWorker orchestrates the signup workflow using the step-based approach
type sUserSignupWorker struct {
	riverjobs.WorkerBase
	orchestrator   *riverjobs.WorkflowOrchestrator
	signupWorkflow *signup.SignupWorkflow
}

func NewUserSignupWorker(wb *riverjobs.WorkerBase, orchestrator *riverjobs.WorkflowOrchestrator) *sUserSignupWorker {
	signupWorkflow := signup.NewSignupWorkflow()

	// Register the workflow with the orchestrator
	orchestrator.RegisterWorkflow(signupWorkflow)

	return &sUserSignupWorker{
		WorkerBase:     *wb,
		orchestrator:   orchestrator,
		signupWorkflow: signupWorkflow,
	}
}

func (s *sUserSignupWorker) Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware {
	return []rivertype.WorkerMiddleware{}
}

func (w *sUserSignupWorker) NextRetry(job *river.Job[riverjobs.UserSignupArgs]) time.Time {
	return time.Now().Add(time.Duration(job.Attempt) * time.Minute)
}

func (w *sUserSignupWorker) Timeout(job *river.Job[riverjobs.UserSignupArgs]) time.Duration {
	return 5 * time.Minute
}

func (w *sUserSignupWorker) Work(ctx context.Context, job *river.Job[riverjobs.UserSignupArgs]) error {
	w.LogJobStart(ctx, job.Args.WorkflowID, job.Args.StepName, job.Args)

	// This worker now delegates to the workflow orchestrator
	// The step name in the job args determines which step to execute
	stepInput := riverjobs.StepInput{
		WorkflowID:      job.Args.WorkflowID,
		StepName:        job.Args.StepName,
		WorkflowContext: make(map[string]interface{}),
		StepContext: map[string]interface{}{
			"email":             job.Args.Email,
			"password":          job.Args.Password,
			"first_name":        job.Args.FirstName,
			"last_name":         job.Args.LastName,
			"role":              job.Args.Role,
			"organization_data": job.Args.OrganizationData,
		},
		Metadata: make(map[string]interface{}),
	}

	// Get current workflow context
	execution, err := w.orchestrator.GetWorkflowStatus(ctx, job.Args.WorkflowID)
	if err != nil {
		return fmt.Errorf("failed to get workflow status: %w", err)
	}

	stepInput.WorkflowContext = execution.Context

	// Execute the step through the orchestrator
	result, err := w.orchestrator.ExecuteStep(ctx, job.Args.WorkflowID, job.Args.StepName, stepInput)
	if err != nil {
		return fmt.Errorf("step execution failed: %w", err)
	}

	// Log the result
	g.Log().Info(ctx, "Step completed", g.Map{
		"workflow_id": job.Args.WorkflowID,
		"step_name":   job.Args.StepName,
		"success":     result.Success,
	})

	return nil
}

func (w *sUserSignupWorker) FailWorkflow(ctx context.Context, workflowID, stepName string, err error) {
	errorMsg := err.Error()
	w.orchestrator.FailWorkflow(ctx, workflowID, errorMsg)
}

func (w *sUserSignupWorker) StartNewFlow(ctx context.Context, args riverjobs.UserSignupArgs, opts *river.InsertOpts) (string, error) {
	// Create signup input from args
	signupInput := signup.SignupInput{
		Email:            args.Email,
		Password:         args.Password,
		FirstName:        args.FirstName,
		LastName:         args.LastName,
		Role:             args.Role,
		OrganizationData: args.OrganizationData,
	}

	// Start user signup workflow using the workflow definition
	workflowID, err := w.orchestrator.StartWorkflowWithDefinition(
		ctx,
		"user_signup",
		signupInput,
		args.OrgID,
		args.UserID,
	)
	if err != nil {
		return "", fmt.Errorf("failed to start signup workflow: %w", err)
	}

	g.Log().Info(ctx, "Started signup workflow", g.Map{
		"workflow_id": workflowID,
		"email":       args.Email,
		"org_id":      args.OrgID,
	})

	return workflowID, nil
}

func (w *sUserSignupWorker) GetSignupStatus(ctx context.Context, workflowID string) (*riverjobs.WorkflowExecution, error) {
	execution, err := w.orchestrator.GetWorkflowStatus(ctx, workflowID)
	if err != nil {
		return nil, fmt.Errorf("failed to get workflow status: %w", err)
	}

	return execution, nil

}

// The old implementation methods have been removed and replaced by
// the step-based workflow system. Each step is now implemented as
// a separate step in the internal/workflow/signup/steps/ package.
