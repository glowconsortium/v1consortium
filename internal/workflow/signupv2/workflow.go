package signupv2

import (
	"context"
	"fmt"
	"time"
	"v1consortium/internal/pkg/riverjobsv2"

	"github.com/gogf/gf/v2/frame/g"
)

// SignupInput represents the input for the signup workflow
type SignupInput struct {
	Email            string                 `json:"email"`
	Password         string                 `json:"password"`
	FirstName        string                 `json:"first_name"`
	LastName         string                 `json:"last_name"`
	Role             string                 `json:"role"`
	OrganizationData map[string]interface{} `json:"organization_data"`
	Metadata         map[string]interface{} `json:"metadata"`
}

// NewSignupWorkflow creates the complete signup workflow with all steps and their functions
func NewSignupWorkflow() *riverjobsv2.Workflow {
	return &riverjobsv2.Workflow{
		Name:      "signup",
		FirstStep: "validate",
		Steps: []riverjobsv2.Step{
			{
				Name:       "validate",
				Execute:    validateStep,
				IsOptional: false,
				MaxRetries: 3,
				RetryDelay: time.Second * 5,
				Queue:      "default",
				Timeout:    time.Minute * 2,
			},
			{
				Name:       "create_user",
				Execute:    createUserStep,
				IsOptional: false,
				MaxRetries: 3,
				RetryDelay: time.Second * 10,
				Queue:      "default",
				Timeout:    time.Minute * 5,
			},
			{
				Name:       "create_organization",
				Execute:    createOrganizationStep,
				IsOptional: false,
				MaxRetries: 3,
				RetryDelay: time.Second * 10,
				Queue:      "default",
				Timeout:    time.Minute * 5,
			},
			{
				Name:       "setup_stripe",
				Execute:    setupStripeStep,
				IsOptional: true, // Optional step - can fail without stopping workflow
				MaxRetries: 5,
				RetryDelay: time.Second * 15,
				Queue:      "external", // Different queue for external service calls
				Timeout:    time.Minute * 10,
			},
			{
				Name:       "send_verification",
				Execute:    sendVerificationStep,
				IsOptional: true, // Optional step
				MaxRetries: 3,
				RetryDelay: time.Second * 5,
				Queue:      "notifications",
				Timeout:    time.Minute * 2,
			},
		},
		StepFlow: map[string]string{
			"validate":            "create_user",
			"create_user":         "create_organization",
			"create_organization": "setup_stripe",
			"setup_stripe":        "send_verification",
			"send_verification":   "", // End of workflow
		},
		ValidateFunc: validateWorkflowInput,
	}
}

// validateWorkflowInput validates the initial signup input
func validateWorkflowInput(ctx context.Context, input map[string]interface{}) error {
	// Extract signup input
	email, ok := input["email"].(string)
	if !ok || email == "" {
		return fmt.Errorf("email is required")
	}

	password, ok := input["password"].(string)
	if !ok || len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}

	firstName, ok := input["first_name"].(string)
	if !ok || len(firstName) < 2 {
		return fmt.Errorf("first_name must be at least 2 characters")
	}

	lastName, ok := input["last_name"].(string)
	if !ok || len(lastName) < 2 {
		return fmt.Errorf("last_name must be at least 2 characters")
	}

	role, ok := input["role"].(string)
	if !ok || (role != "driver" && role != "admin" && role != "owner") {
		return fmt.Errorf("role must be one of: driver, admin, owner")
	}

	return nil
}

// validateStep validates the signup data and prepares for user creation
func validateStep(ctx context.Context, input map[string]interface{}, workflowCtx riverjobsv2.WorkflowContext) (*riverjobsv2.StepResult, error) {
	// Basic validation (already done by ValidateFunc, but we can add more here)
	email, _ := input["email"].(string)

	// Check if email already exists (mock implementation)
	if emailExists(ctx, email) {
		return &riverjobsv2.StepResult{
			Success:      false,
			ErrorMessage: "email already exists",
		}, fmt.Errorf("email %s already exists", email)
	}

	// Add validation timestamp to context
	result := &riverjobsv2.StepResult{
		Success: true,
		Data: riverjobsv2.WorkflowContext{
			"validation_completed_at": time.Now(),
			"validated_email":         email,
		},
	}

	g.Log().Infof(ctx, "Validation completed successfully for email %s", email)
	return result, nil
}

// createUserStep creates the user account
func createUserStep(ctx context.Context, input map[string]interface{}, workflowCtx riverjobsv2.WorkflowContext) (*riverjobsv2.StepResult, error) {
	email, _ := input["email"].(string)
	password, _ := input["password"].(string)
	firstName, _ := input["first_name"].(string)
	lastName, _ := input["last_name"].(string)
	role, _ := input["role"].(string)

	// Mock user creation
	userID, err := createUser(ctx, email, password, firstName, lastName, role)
	if err != nil {
		return &riverjobsv2.StepResult{
			Success:      false,
			ErrorMessage: fmt.Sprintf("failed to create user: %v", err),
		}, err
	}

	g.Log().Infof(ctx, "User created successfully with ID %s for email %s", userID, email)
	return &riverjobsv2.StepResult{
		Success: true,
		Data: riverjobsv2.WorkflowContext{
			"user_id":         userID,
			"user_created_at": time.Now(),
		},
	}, nil
}

// createOrganizationStep creates the organization and associates the user
func createOrganizationStep(ctx context.Context, input map[string]interface{}, workflowCtx riverjobsv2.WorkflowContext) (*riverjobsv2.StepResult, error) {
	userID, exists := workflowCtx.GetString("user_id")
	if !exists {
		return &riverjobsv2.StepResult{
			Success:      false,
			ErrorMessage: "user_id not found in workflow context",
		}, fmt.Errorf("user_id not found in context")
	}

	// Get organization data from input
	orgData, _ := input["organization_data"].(map[string]interface{})

	// Mock organization creation
	orgID, err := createOrganization(ctx, userID, orgData)
	if err != nil {
		return &riverjobsv2.StepResult{
			Success:      false,
			ErrorMessage: fmt.Sprintf("failed to create organization: %v", err),
		}, err
	}

	g.Log().Infof(ctx, "Organization created successfully with ID %s for user %s", orgID, userID)
	return &riverjobsv2.StepResult{
		Success: true,
		Data: riverjobsv2.WorkflowContext{
			"organization_id":         orgID,
			"organization_created_at": time.Now(),
		},
	}, nil
}

// setupStripeStep sets up Stripe customer and subscription (optional)
func setupStripeStep(ctx context.Context, input map[string]interface{}, workflowCtx riverjobsv2.WorkflowContext) (*riverjobsv2.StepResult, error) {
	userID, _ := workflowCtx.GetString("user_id")
	orgID, _ := workflowCtx.GetString("organization_id")
	email, _ := input["email"].(string)

	// Mock Stripe setup
	stripeCustomerID, err := setupStripeCustomer(ctx, userID, orgID, email)
	if err != nil {
		// Since this is optional, we log the error but don't fail the workflow
		return &riverjobsv2.StepResult{
			Success:      false,
			ErrorMessage: fmt.Sprintf("failed to setup Stripe: %v", err),
			ShouldRetry:  true, // Retry external service failures
		}, err
	}

	g.Log().Infof(ctx, "Stripe customer setup successfully with ID %s for user %s (org: %s)", stripeCustomerID, userID, orgID)
	return &riverjobsv2.StepResult{
		Success: true,
		Data: riverjobsv2.WorkflowContext{
			"stripe_customer_id": stripeCustomerID,
			"stripe_setup_at":    time.Now(),
		},
	}, nil
}

// sendVerificationStep sends email verification (optional)
func sendVerificationStep(ctx context.Context, input map[string]interface{}, workflowCtx riverjobsv2.WorkflowContext) (*riverjobsv2.StepResult, error) {
	userID, _ := workflowCtx.GetString("user_id")
	email, _ := input["email"].(string)

	// Mock email sending
	err := sendVerificationEmail(ctx, userID, email)
	if err != nil {
		return &riverjobsv2.StepResult{
			Success:      false,
			ErrorMessage: fmt.Sprintf("failed to send verification email: %v", err),
			ShouldRetry:  true,
		}, err
	}

	g.Log().Infof(ctx, "Verification email sent successfully to %s for user %s", email, userID)
	return &riverjobsv2.StepResult{
		Success: true,
		Data: riverjobsv2.WorkflowContext{
			"verification_email_sent": true,
			"verification_sent_at":    time.Now(),
		},
	}, nil
}

// Mock implementation functions (replace with real implementations)

func emailExists(ctx context.Context, email string) bool {
	// Mock implementation - replace with real database check
	mockExistingEmails := []string{"existing@example.com", "test@duplicate.com"}
	for _, existing := range mockExistingEmails {
		if email == existing {
			return true
		}
	}
	return false
}

func createUser(ctx context.Context, email, password, firstName, lastName, role string) (string, error) {
	// Mock implementation - replace with real user creation
	if email == "fail@example.com" {
		return "", fmt.Errorf("mock user creation failure")
	}

	// Generate mock user ID
	userID := fmt.Sprintf("user_%d", time.Now().UnixNano())
	return userID, nil
}

func createOrganization(ctx context.Context, userID string, orgData map[string]interface{}) (string, error) {
	// Mock implementation - replace with real organization creation
	if userID == "fail_org_user" {
		return "", fmt.Errorf("mock organization creation failure")
	}

	// Generate mock organization ID
	orgID := fmt.Sprintf("org_%d", time.Now().UnixNano())
	return orgID, nil
}

func setupStripeCustomer(ctx context.Context, userID, orgID, email string) (string, error) {
	// Mock implementation - replace with real Stripe integration
	if email == "stripe_fail@example.com" {
		return "", fmt.Errorf("mock Stripe setup failure")
	}

	// Generate mock Stripe customer ID
	customerID := fmt.Sprintf("cus_%d", time.Now().UnixNano())
	return customerID, nil
}

func sendVerificationEmail(ctx context.Context, userID, email string) error {
	// Mock implementation - replace with real email sending
	if email == "email_fail@example.com" {
		return fmt.Errorf("mock email sending failure")
	}

	g.Log().Infof(ctx, "Sent verification email to %s for user %s", email, userID)

	return nil
}
