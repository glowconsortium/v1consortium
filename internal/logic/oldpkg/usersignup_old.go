package riverjobusersignup

import (
	"context"
	"fmt"
	"time"
	"v1consortium/internal/model/do"
	"v1consortium/internal/pkg/riverjobs"
	"v1consortium/internal/pkg/stripeclient"
	"v1consortium/internal/service"
	"v1consortium/internal/workflow/signup"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/rivertype"
)

// sUserSignupWorker now orchestrates the signup workflow using the step-based approach
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

func (w *sUserSignupWorker) validateSignupData(args riverjobs.UserSignupArgs) error {
	if args.Email == "" {
		return fmt.Errorf("email is required")
	}
	if args.Password == "" {
		return fmt.Errorf("password is required")
	}
	if args.OrganizationData == nil {
		return fmt.Errorf("organization data is required")
	}
	if name, ok := args.OrganizationData["companyName"].(string); !ok || name == "" {
		return fmt.Errorf("company name is required")
	}
	if isDot, ok := args.OrganizationData["isDotCompany"].(bool); !ok {
		return fmt.Errorf("isDotCompany flag is required")
	} else if isDot {
		if dotNum, ok := args.OrganizationData["dotNumber"].(string); !ok || dotNum == "" {
			return fmt.Errorf("DOT number is required for .com companies")
		}
	}

	if args.FirstName == "" {
		return fmt.Errorf("first name is required")
	}
	if args.LastName == "" {
		return fmt.Errorf("last name is required")
	}
	// Add more validation logic
	return nil
}

func (w *sUserSignupWorker) createUserAccount(ctx context.Context, email, password string, metadata map[string]interface{}) (string, error) {
	// Mock implementation - would call Supabase Auth
	//userID := "user_" + fmt.Sprintf("%d", time.Now().Unix())

	resp, err := service.SupabaseService().SignUp(ctx, email, password, metadata)
	if err != nil {
		return "", fmt.Errorf("failed to create user account: %w", err)
	}
	return resp.User.ID.String(), nil
}

func (w *sUserSignupWorker) createUserProfileInDB(ctx context.Context, userID, email, firstName, lastName, role, organizationId string) error {
	// Mock implementation - would insert into user_profiles table

	err := service.Auth().CreateUserProfile(ctx, userID, &do.UserProfiles{
		Email:          email,
		FirstName:      firstName,
		LastName:       lastName,
		Role:           role,
		OrganizationId: organizationId,
	})
	if err != nil {
		return fmt.Errorf("failed to create user profile in DB: %w", err)
	}
	g.Log().Infof(ctx, "Created user profile in DB for user %s\n", userID)
	return nil
}

func (w *sUserSignupWorker) createOrganization(ctx context.Context, organizationData map[string]interface{}) (string, error) {
	create_data := do.Organizations{
		Id:             guid.S(),
		Name:           organizationData["companyName"].(string),
		IsDotRegulated: organizationData["isDotCompany"].(bool),
		UsdotNumber:    organizationData["dotNumber"].(string),
		Type:           organizationData["organizationType"].(string),
	}

	org, err := service.OrganizationService().CreateOrganization(ctx, &create_data)
	if err != nil {
		return "", fmt.Errorf("failed to create organization: %w", err)
	}
	return org.Id, nil
}

func (w *sUserSignupWorker) createStripeCustomer(ctx context.Context, userID, email, organizationID, name string) (string, error) {

	stripeCustomerID, err := service.StripeService().CreateCustomer(ctx, &stripeclient.CustomerData{
		Email:       email,
		Name:        name,
		Description: fmt.Sprintf("user_id: %s, organinzation_id: %s", userID, organizationID),
		Metadata: map[string]string{
			"user_id":         userID,
			"organization_id": organizationID,
		},
	})

	if err != nil {
		return "", fmt.Errorf("failed to create stripe customer: %w", err)
	}
	return stripeCustomerID.ID, nil
}

func (w *sUserSignupWorker) createOrganizationSubscription(ctx context.Context, stripeCustomerID, organizationID string) (string, error) {

	subscription, err := service.OrganizationService().CreateOrganizationSubscription(ctx, &do.OrganizationSubscriptions{
		Id:               guid.S(),
		OrganizationId:   organizationID,
		StripeCustomerId: stripeCustomerID,
		PlanId:           w.getDefaultPlanIdForOrganizationType(ctx, "client"),
		Status:           "pending",
		StartDate:        gtime.Now(),
		EndDate:          gtime.NewFromTime(time.Now().AddDate(0, 1, 0)), // 1 month trial
	})

	if err != nil {
		return "", fmt.Errorf("failed to create organization subscription: %w", err)
	}
	return subscription.Id, nil
}

func (w *sUserSignupWorker) getDefaultPlanIdForOrganizationType(ctx context.Context, orgType string) string {
	// Mock implementation - would map organization type to plan ID

	var tier string

	switch orgType {
	case "":
		tier = "custom"
	case "internal":
		tier = "custom"
	case "client":
		tier = "starter"
	case "provider":
		tier = "professional"
	default:
		tier = "custom"
	}

	plan, err := service.OrganizationService().GetPlanByTier(ctx, tier)

	if err != nil || plan == nil {
		return "848c0966-8027-4a30-952e-aeb093b76979"
	}
	return plan.Id
}

func (w *sUserSignupWorker) updateUserProfile(ctx context.Context, userID, stripeCustomerID string) error {
	// Mock implementation - would update user_profiles table
	err := service.Auth().UpdateUserProfile(ctx, userID, map[string]interface{}{
		"stripe_customer_id": stripeCustomerID,
	})
	if err != nil {
		return fmt.Errorf("failed to update user profile: %w", err)
	}
	g.Log().Infof(ctx, "Updated user profile %s with stripe customer ID %s\n", userID, stripeCustomerID)
	return nil
}

func (w *sUserSignupWorker) sendVerificationEmail(ctx context.Context, email, userID string) error {
	// Mock implementation - would integrate with email service
	g.Log().Infof(ctx, "Sent verification email to %s for user %s\n", email, userID)
	return nil
}
