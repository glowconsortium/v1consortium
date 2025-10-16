package riverjobs

import (
	"context"
	"fmt"
	"time"

	"github.com/riverqueue/river"
	"github.com/riverqueue/river/rivertype"
)

// Example workflow implementations for V1 Consortium

// UserSignupWorkflow - Handles user signup and subscription
type UserSignupArgs struct {
	BaseJobArgs
	Email            string                 `json:"email"`
	Password         string                 `json:"password"`
	OrganizationData map[string]interface{} `json:"organization_data"`
	SubscriptionPlan string                 `json:"subscription_plan"`
	PaymentMethodID  string                 `json:"payment_method_id"`
}

func (args UserSignupArgs) Kind() string { return "user_signup" }

type UserSignupWorker struct {
	WorkerBase
}

// Middleware implements the River Worker interface
func (w *UserSignupWorker) Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware {
	return []rivertype.WorkerMiddleware{}
}

// NextRetry implements the River Worker interface
func (w *UserSignupWorker) NextRetry(job *river.Job[UserSignupArgs]) time.Time {
	return time.Now().Add(time.Duration(job.Attempt) * time.Minute)
}

// Timeout implements the River Worker interface
func (w *UserSignupWorker) Timeout(job *river.Job[UserSignupArgs]) time.Duration {
	return 5 * time.Minute
}

func (w *UserSignupWorker) Work(ctx context.Context, job *river.Job[UserSignupArgs]) error {
	w.LogJobStart(ctx, job.Args.WorkflowID, job.Args.StepName, job.Args)
	startTime := time.Now()

	return w.ExecuteWithWorkflowTracking(ctx, job.Args.WorkflowID, job.Args.StepName, func(ctx context.Context) (map[string]interface{}, error) {
		// 1. Validate signup data
		if err := w.validateSignupData(job.Args); err != nil {
			return nil, fmt.Errorf("validation failed: %w", err)
		}

		// 2. Create user account (would call Supabase)
		userID, err := w.createUserAccount(ctx, job.Args.Email, job.Args.Password)
		if err != nil {
			return nil, fmt.Errorf("failed to create user: %w", err)
		}

		// 3. Update workflow context with user ID
		contextUpdate := map[string]interface{}{
			"user_id": userID,
			"email":   job.Args.Email,
		}

		if err := w.UpdateWorkflowContext(ctx, job.Args.WorkflowID, contextUpdate); err != nil {
			return nil, fmt.Errorf("failed to update context: %w", err)
		}

		// 4. Enqueue next job (create organization)
		nextArgs := CreateOrganizationArgs{
			BaseJobArgs: BaseJobArgs{
				WorkflowID:   job.Args.WorkflowID,
				WorkflowType: job.Args.WorkflowType,
				StepName:     "create_organization",
				OrgID:        job.Args.OrgID,
				UserID:       userID,
			},
			UserID:           userID,
			OrganizationData: job.Args.OrganizationData,
		}

		_, err = w.EnqueueNextJob(ctx, nextArgs, &river.InsertOpts{
			Queue:    QueueDefault,
			Priority: 1,
		})

		if err != nil {
			return nil, fmt.Errorf("failed to enqueue next job: %w", err)
		}

		w.LogJobComplete(ctx, job.Args.WorkflowID, job.Args.StepName, time.Since(startTime).String())

		return map[string]interface{}{
			"user_id": userID,
			"status":  "user_created",
		}, nil
	})
}

func (w *UserSignupWorker) validateSignupData(args UserSignupArgs) error {
	if args.Email == "" {
		return fmt.Errorf("email is required")
	}
	if args.Password == "" {
		return fmt.Errorf("password is required")
	}
	// Add more validation logic
	return nil
}

func (w *UserSignupWorker) createUserAccount(ctx context.Context, email, password string) (string, error) {
	// Mock implementation - would call Supabase Auth
	userID := "user_" + fmt.Sprintf("%d", time.Now().Unix())
	return userID, nil
}

// CreateOrganizationArgs for organization creation step
type CreateOrganizationArgs struct {
	BaseJobArgs
	UserID           string                 `json:"user_id"`
	OrganizationData map[string]interface{} `json:"organization_data"`
}

func (args CreateOrganizationArgs) Kind() string { return "create_organization" }

type CreateOrganizationWorker struct {
	WorkerBase
}

// Middleware implements the River Worker interface
func (w *CreateOrganizationWorker) Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware {
	return []rivertype.WorkerMiddleware{}
}

// NextRetry implements the River Worker interface
func (w *CreateOrganizationWorker) NextRetry(job *river.Job[CreateOrganizationArgs]) time.Time {
	return time.Now().Add(time.Duration(job.Attempt) * time.Minute)
}

// Timeout implements the River Worker interface
func (w *CreateOrganizationWorker) Timeout(job *river.Job[CreateOrganizationArgs]) time.Duration {
	return 5 * time.Minute
}

func (w *CreateOrganizationWorker) Work(ctx context.Context, job *river.Job[CreateOrganizationArgs]) error {
	w.LogJobStart(ctx, job.Args.WorkflowID, job.Args.StepName, job.Args)
	startTime := time.Now()

	return w.ExecuteWithWorkflowTracking(ctx, job.Args.WorkflowID, job.Args.StepName, func(ctx context.Context) (map[string]interface{}, error) {
		// 1. Create organization
		orgID, err := w.createOrganization(ctx, job.Args.UserID, job.Args.OrganizationData)
		if err != nil {
			return nil, fmt.Errorf("failed to create organization: %w", err)
		}

		// 2. Update workflow context
		contextUpdate := map[string]interface{}{
			"organization_id": orgID,
		}

		if err := w.UpdateWorkflowContext(ctx, job.Args.WorkflowID, contextUpdate); err != nil {
			return nil, fmt.Errorf("failed to update context: %w", err)
		}

		// 3. Enqueue subscription processing job
		nextArgs := ProcessSubscriptionArgs{
			BaseJobArgs: BaseJobArgs{
				WorkflowID:   job.Args.WorkflowID,
				WorkflowType: job.Args.WorkflowType,
				StepName:     "process_subscription",
				OrgID:        orgID,
				UserID:       job.Args.UserID,
			},
			OrganizationID: orgID,
			UserID:         job.Args.UserID,
		}

		_, err = w.EnqueueNextJob(ctx, nextArgs, &river.InsertOpts{
			Queue:    QueueCritical, // High priority for subscription
			Priority: 2,
		})

		if err != nil {
			return nil, fmt.Errorf("failed to enqueue subscription job: %w", err)
		}

		w.LogJobComplete(ctx, job.Args.WorkflowID, job.Args.StepName, time.Since(startTime).String())

		return map[string]interface{}{
			"organization_id": orgID,
			"status":          "organization_created",
		}, nil
	})
}

func (w *CreateOrganizationWorker) createOrganization(ctx context.Context, userID string, orgData map[string]interface{}) (string, error) {
	// Mock implementation - would call database
	orgID := "org_" + fmt.Sprintf("%d", time.Now().Unix())
	return orgID, nil
}

// ProcessSubscriptionArgs for subscription processing
type ProcessSubscriptionArgs struct {
	BaseJobArgs
	OrganizationID string `json:"organization_id"`
	UserID         string `json:"user_id"`
}

func (args ProcessSubscriptionArgs) Kind() string { return "process_subscription" }

type ProcessSubscriptionWorker struct {
	WorkerBase
}

// Middleware implements the River Worker interface
func (w *ProcessSubscriptionWorker) Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware {
	return []rivertype.WorkerMiddleware{}
}

// NextRetry implements the River Worker interface
func (w *ProcessSubscriptionWorker) NextRetry(job *river.Job[ProcessSubscriptionArgs]) time.Time {
	return time.Now().Add(time.Duration(job.Attempt) * time.Minute)
}

// Timeout implements the River Worker interface
func (w *ProcessSubscriptionWorker) Timeout(job *river.Job[ProcessSubscriptionArgs]) time.Duration {
	return 5 * time.Minute
}

func (w *ProcessSubscriptionWorker) Work(ctx context.Context, job *river.Job[ProcessSubscriptionArgs]) error {
	w.LogJobStart(ctx, job.Args.WorkflowID, job.Args.StepName, job.Args)
	startTime := time.Now()

	return w.ExecuteWithWorkflowTracking(ctx, job.Args.WorkflowID, job.Args.StepName, func(ctx context.Context) (map[string]interface{}, error) {
		// 1. Process subscription with Stripe
		subscriptionID, err := w.processSubscription(ctx, job.Args.OrganizationID, job.Args.UserID)
		if err != nil {
			return nil, fmt.Errorf("failed to process subscription: %w", err)
		}

		// 2. Update workflow context
		contextUpdate := map[string]interface{}{
			"subscription_id":     subscriptionID,
			"subscription_status": "active",
		}

		if err := w.UpdateWorkflowContext(ctx, job.Args.WorkflowID, contextUpdate); err != nil {
			return nil, fmt.Errorf("failed to update context: %w", err)
		}

		// 3. Complete the workflow
		execution, err := w.GetWorkflowExecution(ctx, job.Args.WorkflowID)
		if err != nil {
			return nil, fmt.Errorf("failed to get workflow execution: %w", err)
		}

		err = w.WorkflowManager.CompleteWorkflow(ctx, execution.WorkflowID)
		if err != nil {
			return nil, fmt.Errorf("failed to complete workflow: %w", err)
		}

		w.LogJobComplete(ctx, job.Args.WorkflowID, job.Args.StepName, time.Since(startTime).String())

		return map[string]interface{}{
			"subscription_id": subscriptionID,
			"status":          "workflow_completed",
		}, nil
	})
}

func (w *ProcessSubscriptionWorker) processSubscription(ctx context.Context, orgID, userID string) (string, error) {
	// Mock implementation - would call Stripe API
	subscriptionID := "sub_" + fmt.Sprintf("%d", time.Now().Unix())
	return subscriptionID, nil
}

// Drug Test Workflow Example
type DrugTestOrderArgs struct {
	BaseJobArgs
	EmployeeID string `json:"employee_id"`
	TestType   string `json:"test_type"`
	FacilityID string `json:"facility_id,omitempty"`
	RushOrder  bool   `json:"rush_order"`
}

func (args DrugTestOrderArgs) Kind() string { return "drug_test_order" }

type DrugTestOrderWorker struct {
	WorkerBase
}

// Middleware implements the River Worker interface
func (w *DrugTestOrderWorker) Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware {
	return []rivertype.WorkerMiddleware{}
}

// NextRetry implements the River Worker interface
func (w *DrugTestOrderWorker) NextRetry(job *river.Job[DrugTestOrderArgs]) time.Time {
	return time.Now().Add(time.Duration(job.Attempt) * time.Minute)
}

// Timeout implements the River Worker interface
func (w *DrugTestOrderWorker) Timeout(job *river.Job[DrugTestOrderArgs]) time.Duration {
	return 5 * time.Minute
}

func (w *DrugTestOrderWorker) Work(ctx context.Context, job *river.Job[DrugTestOrderArgs]) error {
	w.LogJobStart(ctx, job.Args.WorkflowID, job.Args.StepName, job.Args)
	startTime := time.Now()

	return w.ExecuteWithWorkflowTracking(ctx, job.Args.WorkflowID, job.Args.StepName, func(ctx context.Context) (map[string]interface{}, error) {
		// 1. Validate test request
		if err := w.validateTestRequest(job.Args); err != nil {
			return nil, fmt.Errorf("validation failed: %w", err)
		}

		// 2. Select testing facility if not specified
		facilityID := job.Args.FacilityID
		if facilityID == "" {
			facility, err := w.selectOptimalFacility(ctx, job.Args.OrgID)
			if err != nil {
				return nil, fmt.Errorf("failed to select facility: %w", err)
			}
			facilityID = facility
		}

		// 3. Create external order (Quest Diagnostics)
		externalOrderID, err := w.createExternalOrder(ctx, job.Args, facilityID)
		if err != nil {
			return nil, fmt.Errorf("failed to create external order: %w", err)
		}

		// 4. Update workflow context
		contextUpdate := map[string]interface{}{
			"external_order_id": externalOrderID,
			"facility_id":       facilityID,
			"test_status":       "ordered",
		}

		if err := w.UpdateWorkflowContext(ctx, job.Args.WorkflowID, contextUpdate); err != nil {
			return nil, fmt.Errorf("failed to update context: %w", err)
		}

		// 5. Schedule notification job
		notificationArgs := SendTestNotificationArgs{
			BaseJobArgs: BaseJobArgs{
				WorkflowID:   job.Args.WorkflowID,
				WorkflowType: job.Args.WorkflowType,
				StepName:     "send_notification",
				OrgID:        job.Args.OrgID,
				UserID:       job.Args.UserID,
			},
			EmployeeID:      job.Args.EmployeeID,
			ExternalOrderID: externalOrderID,
			FacilityID:      facilityID,
		}

		_, err = w.EnqueueNextJob(ctx, notificationArgs, &river.InsertOpts{
			Queue:    QueueNotification,
			Priority: 1,
		})

		if err != nil {
			return nil, fmt.Errorf("failed to enqueue notification job: %w", err)
		}

		w.LogJobComplete(ctx, job.Args.WorkflowID, job.Args.StepName, time.Since(startTime).String())

		return map[string]interface{}{
			"external_order_id": externalOrderID,
			"facility_id":       facilityID,
			"status":            "test_ordered",
		}, nil
	})
}

func (w *DrugTestOrderWorker) validateTestRequest(args DrugTestOrderArgs) error {
	if args.EmployeeID == "" {
		return fmt.Errorf("employee_id is required")
	}
	if args.TestType == "" {
		return fmt.Errorf("test_type is required")
	}
	return nil
}

func (w *DrugTestOrderWorker) selectOptimalFacility(ctx context.Context, orgID string) (string, error) {
	// Mock implementation - would query facilities database
	return "facility_123", nil
}

func (w *DrugTestOrderWorker) createExternalOrder(ctx context.Context, args DrugTestOrderArgs, facilityID string) (string, error) {
	// Mock implementation - would call Quest Diagnostics API
	orderID := "quest_" + fmt.Sprintf("%d", time.Now().Unix())
	return orderID, nil
}

// Send Test Notification Args
type SendTestNotificationArgs struct {
	BaseJobArgs
	EmployeeID      string `json:"employee_id"`
	ExternalOrderID string `json:"external_order_id"`
	FacilityID      string `json:"facility_id"`
}

func (args SendTestNotificationArgs) Kind() string { return "send_test_notification" }

type SendTestNotificationWorker struct {
	WorkerBase
}

// Middleware implements the River Worker interface
func (w *SendTestNotificationWorker) Middleware(job *rivertype.JobRow) []rivertype.WorkerMiddleware {
	return []rivertype.WorkerMiddleware{}
}

// NextRetry implements the River Worker interface
func (w *SendTestNotificationWorker) NextRetry(job *river.Job[SendTestNotificationArgs]) time.Time {
	return time.Now().Add(time.Duration(job.Attempt) * time.Minute)
}

// Timeout implements the River Worker interface
func (w *SendTestNotificationWorker) Timeout(job *river.Job[SendTestNotificationArgs]) time.Duration {
	return 3 * time.Minute // Shorter timeout for notifications
}

func (w *SendTestNotificationWorker) Work(ctx context.Context, job *river.Job[SendTestNotificationArgs]) error {
	w.LogJobStart(ctx, job.Args.WorkflowID, job.Args.StepName, job.Args)

	return w.ExecuteWithWorkflowTracking(ctx, job.Args.WorkflowID, job.Args.StepName, func(ctx context.Context) (map[string]interface{}, error) {
		// Send notification to employee and supervisor
		err := w.sendNotifications(ctx, job.Args)
		if err != nil {
			return nil, fmt.Errorf("failed to send notifications: %w", err)
		}

		// Complete workflow
		err = w.WorkflowManager.CompleteWorkflow(ctx, job.Args.WorkflowID)
		if err != nil {
			return nil, fmt.Errorf("failed to complete workflow: %w", err)
		}

		return map[string]interface{}{
			"notifications_sent": true,
			"status":             "workflow_completed",
		}, nil
	})
}

func (w *SendTestNotificationWorker) sendNotifications(ctx context.Context, args SendTestNotificationArgs) error {
	// Mock implementation - would send emails/SMS
	return nil
}
