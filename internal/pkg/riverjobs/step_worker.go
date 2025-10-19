package riverjobs

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/rivertype"
)

// StepWorker is the base interface for all workflow step workers
type StepWorker interface {
	// Execute performs the step's business logic
	Execute(ctx context.Context, args StepArgs) (map[string]interface{}, error)

	// GetStepName returns the name of this step
	GetStepName() string

	// NextRetry returns when to retry this step on failure
	NextRetry(job *river.Job[StepArgs]) time.Time

	// Timeout returns the maximum execution time for this step
	Timeout(job *river.Job[StepArgs]) time.Duration
}

// BaseStepWorker provides common functionality for all step workers
type BaseStepWorker struct {
	stepName        string
	workflowManager *SimpleWorkflowManager
	stepImpl        StepWorker // Reference to the actual step implementation
}

// NewBaseStepWorker creates a new base step worker
func NewBaseStepWorker(stepName string, workflowManager *SimpleWorkflowManager, stepImpl StepWorker) *BaseStepWorker {
	return &BaseStepWorker{
		stepName:        stepName,
		workflowManager: workflowManager,
		stepImpl:        stepImpl,
	}
}

// GetStepName returns the step name
func (w *BaseStepWorker) GetStepName() string {
	return w.stepName
}

// NextRetry provides default retry logic
func (w *BaseStepWorker) NextRetry(job *river.Job[StepArgs]) time.Time {
	return time.Now().Add(time.Duration(job.Attempt) * time.Minute)
}

// Timeout provides default timeout
func (w *BaseStepWorker) Timeout(job *river.Job[StepArgs]) time.Duration {
	return 5 * time.Minute
}

// Middleware provides default middleware (empty)
func (w *BaseStepWorker) Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware {
	return []rivertype.WorkerMiddleware{}
}

// Work is the River worker implementation that calls the step's Execute method
func (w *BaseStepWorker) Work(ctx context.Context, job *river.Job[StepArgs]) error {
	args := job.Args

	g.Log().Info(ctx, "Starting step execution", g.Map{
		"workflow_id":   args.WorkflowID,
		"workflow_type": args.WorkflowType,
		"step_name":     args.StepName,
		"org_id":        args.OrgID,
	})

	// Execute the step-specific logic
	stepOutput, err := w.stepImpl.Execute(ctx, args)
	if err != nil {
		g.Log().Error(ctx, "Step execution failed", g.Map{
			"workflow_id": args.WorkflowID,
			"step_name":   args.StepName,
			"error":       err,
		})

		// Mark workflow as failed
		failErr := w.workflowManager.FailWorkflow(ctx, args.WorkflowID, err.Error())
		if failErr != nil {
			g.Log().Error(ctx, "Failed to mark workflow as failed", g.Map{
				"workflow_id": args.WorkflowID,
				"error":       failErr,
			})
		}

		return err
	}

	// Enqueue next steps
	err = w.workflowManager.EnqueueNextSteps(
		ctx,
		args.WorkflowID,
		args.WorkflowType,
		args.StepName,
		args.WorkflowInput,
		stepOutput,
		args.OrgID,
		args.UserID,
	)

	if err != nil {
		g.Log().Error(ctx, "Failed to enqueue next steps", g.Map{
			"workflow_id": args.WorkflowID,
			"step_name":   args.StepName,
			"error":       err,
		})
		return err
	}

	g.Log().Info(ctx, "Step execution completed", g.Map{
		"workflow_id": args.WorkflowID,
		"step_name":   args.StepName,
		"has_output":  len(stepOutput) > 0,
	})

	return nil
}

// runStep is a shared helper that executes a StepWorker given StepArgs
func runStep(ctx context.Context, wm *SimpleWorkflowManager, stepImpl StepWorker, args StepArgs) error {
	g.Log().Info(ctx, "Starting step execution (adapter)", g.Map{
		"workflow_id":   args.WorkflowID,
		"workflow_type": args.WorkflowType,
		"step_name":     args.StepName,
		"org_id":        args.OrgID,
	})

	stepOutput, err := stepImpl.Execute(ctx, args)
	if err != nil {
		g.Log().Error(ctx, "Step execution failed (adapter)", g.Map{
			"workflow_id": args.WorkflowID,
			"step_name":   args.StepName,
			"error":       err,
		})

		// Mark workflow as failed
		if failErr := wm.FailWorkflow(ctx, args.WorkflowID, err.Error()); failErr != nil {
			g.Log().Error(ctx, "Failed to mark workflow as failed (adapter)", g.Map{
				"workflow_id": args.WorkflowID,
				"error":       failErr,
			})
		}
		return err
	}

	// Enqueue next steps
	if enqErr := wm.EnqueueNextSteps(ctx, args.WorkflowID, args.WorkflowType, args.StepName, args.WorkflowInput, stepOutput, args.OrgID, args.UserID); enqErr != nil {
		g.Log().Error(ctx, "Failed to enqueue next steps (adapter)", g.Map{
			"workflow_id": args.WorkflowID,
			"step_name":   args.StepName,
			"error":       enqErr,
		})
		return enqErr
	}

	g.Log().Info(ctx, "Step execution completed (adapter)", g.Map{
		"workflow_id": args.WorkflowID,
		"step_name":   args.StepName,
		"has_output":  len(stepOutput) > 0,
	})

	return nil
}

// Adapters for concrete River arg types. Each converts the specific arg type
// into StepArgs and calls runStep.
type validateAdapter struct {
	impl StepWorker
	wm   *SimpleWorkflowManager
}

func (a *validateAdapter) Work(ctx context.Context, job *river.Job[ValidateStepArgs]) error {
	ja := job.Args
	args := StepArgs{
		WorkflowID:    ja.WorkflowID,
		WorkflowType:  ja.WorkflowType,
		StepName:      ja.StepName,
		OrgID:         ja.OrgID,
		UserID:        ja.UserID,
		WorkflowInput: ja.SignupData,
		StepInput:     nil,
	}
	return runStep(ctx, a.wm, a.impl, args)
}

type createUserAdapter struct {
	impl StepWorker
	wm   *SimpleWorkflowManager
}

func (a *createUserAdapter) Work(ctx context.Context, job *river.Job[CreateUserStepArgs]) error {
	ja := job.Args
	// map UserData into WorkflowInput so step implementations that expect WorkflowInput work
	args := StepArgs{
		WorkflowID:    ja.WorkflowID,
		WorkflowType:  ja.WorkflowType,
		StepName:      ja.StepName,
		OrgID:         ja.OrgID,
		UserID:        ja.UserID,
		WorkflowInput: ja.UserData,
		StepInput:     nil,
	}
	return runStep(ctx, a.wm, a.impl, args)
}

type createOrgAdapter struct {
	impl StepWorker
	wm   *SimpleWorkflowManager
}

func (a *createOrgAdapter) Work(ctx context.Context, job *river.Job[CreateOrganizationStepArgs]) error {
	ja := job.Args
	args := StepArgs{
		WorkflowID:    ja.WorkflowID,
		WorkflowType:  ja.WorkflowType,
		StepName:      ja.StepName,
		OrgID:         ja.OrgID,
		UserID:        ja.UserID,
		WorkflowInput: ja.OrgData,
		StepInput:     nil,
	}
	return runStep(ctx, a.wm, a.impl, args)
}

type setupStripeAdapter struct {
	impl StepWorker
	wm   *SimpleWorkflowManager
}

func (a *setupStripeAdapter) Work(ctx context.Context, job *river.Job[SetupStripeStepArgs]) error {
	ja := job.Args
	args := StepArgs{
		WorkflowID:    ja.WorkflowID,
		WorkflowType:  ja.WorkflowType,
		StepName:      ja.StepName,
		OrgID:         ja.OrgID,
		UserID:        ja.UserID,
		WorkflowInput: ja.StripeData,
		StepInput:     nil,
	}
	return runStep(ctx, a.wm, a.impl, args)
}

type sendVerificationAdapter struct {
	impl StepWorker
	wm   *SimpleWorkflowManager
}

func (a *sendVerificationAdapter) Work(ctx context.Context, job *river.Job[SendVerificationStepArgs]) error {
	ja := job.Args
	args := StepArgs{
		WorkflowID:    ja.WorkflowID,
		WorkflowType:  ja.WorkflowType,
		StepName:      ja.StepName,
		OrgID:         ja.OrgID,
		UserID:        ja.UserID,
		WorkflowInput: ja.VerificationData,
		StepInput:     nil,
	}
	return runStep(ctx, a.wm, a.impl, args)
}

// Exported adapter types with fields matching caller expectations
type ValidateAdapter struct {
	Impl StepWorker
	Wm   *SimpleWorkflowManager
}

func (a *ValidateAdapter) Work(ctx context.Context, job *river.Job[ValidateStepArgs]) error {
	return (&validateAdapter{impl: a.Impl, wm: a.Wm}).Work(ctx, job)
}
func (a *ValidateAdapter) Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware {
	return []rivertype.WorkerMiddleware{}
}
func (a *ValidateAdapter) NextRetry(job *river.Job[ValidateStepArgs]) time.Time {
	return time.Now().Add(time.Duration(job.Attempt) * time.Minute)
}
func (a *ValidateAdapter) Timeout(job *river.Job[ValidateStepArgs]) time.Duration {
	return 5 * time.Minute
}

type CreateUserAdapter struct {
	Impl StepWorker
	Wm   *SimpleWorkflowManager
}

func (a *CreateUserAdapter) Work(ctx context.Context, job *river.Job[CreateUserStepArgs]) error {
	return (&createUserAdapter{impl: a.Impl, wm: a.Wm}).Work(ctx, job)
}
func (a *CreateUserAdapter) Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware {
	return []rivertype.WorkerMiddleware{}
}
func (a *CreateUserAdapter) NextRetry(job *river.Job[CreateUserStepArgs]) time.Time {
	return time.Now().Add(time.Duration(job.Attempt) * time.Minute)
}
func (a *CreateUserAdapter) Timeout(job *river.Job[CreateUserStepArgs]) time.Duration {
	return 5 * time.Minute
}

type CreateOrgAdapter struct {
	Impl StepWorker
	Wm   *SimpleWorkflowManager
}

func (a *CreateOrgAdapter) Work(ctx context.Context, job *river.Job[CreateOrganizationStepArgs]) error {
	return (&createOrgAdapter{impl: a.Impl, wm: a.Wm}).Work(ctx, job)
}
func (a *CreateOrgAdapter) Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware {
	return []rivertype.WorkerMiddleware{}
}
func (a *CreateOrgAdapter) NextRetry(job *river.Job[CreateOrganizationStepArgs]) time.Time {
	return time.Now().Add(time.Duration(job.Attempt) * time.Minute)
}
func (a *CreateOrgAdapter) Timeout(job *river.Job[CreateOrganizationStepArgs]) time.Duration {
	return 5 * time.Minute
}

type SetupStripeAdapter struct {
	Impl StepWorker
	Wm   *SimpleWorkflowManager
}

func (a *SetupStripeAdapter) Work(ctx context.Context, job *river.Job[SetupStripeStepArgs]) error {
	return (&setupStripeAdapter{impl: a.Impl, wm: a.Wm}).Work(ctx, job)
}
func (a *SetupStripeAdapter) Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware {
	return []rivertype.WorkerMiddleware{}
}
func (a *SetupStripeAdapter) NextRetry(job *river.Job[SetupStripeStepArgs]) time.Time {
	return time.Now().Add(time.Duration(job.Attempt) * time.Minute)
}
func (a *SetupStripeAdapter) Timeout(job *river.Job[SetupStripeStepArgs]) time.Duration {
	return 5 * time.Minute
}

type SendVerificationAdapter struct {
	Impl StepWorker
	Wm   *SimpleWorkflowManager
}

func (a *SendVerificationAdapter) Work(ctx context.Context, job *river.Job[SendVerificationStepArgs]) error {
	return (&sendVerificationAdapter{impl: a.Impl, wm: a.Wm}).Work(ctx, job)
}
func (a *SendVerificationAdapter) Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware {
	return []rivertype.WorkerMiddleware{}
}
func (a *SendVerificationAdapter) NextRetry(job *river.Job[SendVerificationStepArgs]) time.Time {
	return time.Now().Add(time.Duration(job.Attempt) * time.Minute)
}
func (a *SendVerificationAdapter) Timeout(job *river.Job[SendVerificationStepArgs]) time.Duration {
	return 5 * time.Minute
}
