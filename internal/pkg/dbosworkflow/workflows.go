package dbosworkflow

import (
	"context"
	"fmt"
	"time"

	"v1consortium/internal/consts"
)

// Type definitions for the workflow
type User struct {
	ID            string
	Email         string
	PasswordHash  string
	FirstName     string
	LastName      string
	Provider      consts.AuthProvider
	ProviderID    string
	EmailVerified bool
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Company struct {
	ID       string
	Name     string
	Size     string
	Industry string
	OwnerID  string
}

type Subscription struct {
	ID              string
	UserID          string
	CompanyID       string
	Tier            consts.SubscriptionTier
	Status          string
	BillingInterval string
	PaymentProvider string
	ProviderSubID   string
}

type OnboardingWorkflowState struct {
	WorkflowID     string
	UserID         string
	CompanyID      string
	Status         string
	EmailVerified  bool
	SubscriptionID string
	CurrentStep    string
	ErrorMessage   string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// Constants
const (
	StatusPending            = "pending"
	StatusEmailVerified      = "email_verified"
	StatusSubscribed         = "subscribed"
	StatusOnboardingComplete = "onboarding_complete"
	StatusFailed             = "failed"
	AuthProviderEmail        = "email"
)

type OnboardingWorkflows struct {
	Transactions *OnboardingTransactions
	Steps        *OnboardingSteps
}

// SignupWorkflow handles the initial signup and company creation
func (w *OnboardingWorkflows) SignupWorkflow(ctx context.Context, input OnboardingWorkflowInput) (string, error) {
	// Generate a workflow ID
	workflowID := fmt.Sprintf("signup_%d", time.Now().UnixNano())

	// Step 1: Create user record
	user := User{
		Email:         input.Email,
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		Provider:      input.Provider,
		EmailVerified: input.Provider != AuthProviderEmail, // Social logins are pre-verified
		Status:        StatusPending,
	}

	userID, err := w.Transactions.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}

	// Step 2: Create company record
	company := Company{
		Name:     input.CompanyName,
		Size:     input.CompanySize,
		Industry: input.CompanyIndustry,
		OwnerID:  userID,
	}

	companyID, err := w.Transactions.CreateCompany(ctx, company)
	if err != nil {
		return "", err
	}

	// Step 3: Link user to company
	err = w.Transactions.LinkUserToCompany(ctx, userID, companyID)
	if err != nil {
		return "", err
	}

	// Step 4: Create workflow state record
	err = w.Transactions.UpdateOnboardingWorkflowState(ctx, OnboardingWorkflowState{
		WorkflowID:    workflowID,
		UserID:        userID,
		CompanyID:     companyID,
		Status:        StatusPending,
		EmailVerified: user.EmailVerified,
		CurrentStep:   "awaiting_verification",
	})
	if err != nil {
		return "", err
	}

	// Step 5: Send verification email (if needed)
	if input.Provider == AuthProviderEmail {
		err = w.Steps.SendVerificationEmail(ctx, userID, input.Email, input.FirstName, workflowID)
		if err != nil {
			// Log error but don't fail workflow - user can request resend
			// We'll handle this gracefully without failing the workflow
		}
	}

	return workflowID, nil
}

// EmailVerificationWorkflow handles email verification
func (w *OnboardingWorkflows) EmailVerificationWorkflow(ctx context.Context, workflowID, userID string) error {
	// Update user as verified
	err := w.Transactions.UpdateUserEmailVerified(ctx, userID)
	if err != nil {
		return err
	}

	// Update workflow state
	err = w.Transactions.UpdateOnboardingWorkflowState(ctx, OnboardingWorkflowState{
		WorkflowID:    workflowID,
		Status:        StatusEmailVerified,
		EmailVerified: true,
		CurrentStep:   "awaiting_subscription",
	})

	return err
}

// SubscriptionWorkflow handles subscription creation and payment
func (w *OnboardingWorkflows) SubscriptionWorkflow(ctx context.Context, input SubscribeWorkflowInput) (string, error) {
	// Step 1: Create subscription in Stripe
	stripeSubID, err := w.Steps.CreateStripeSubscription(ctx, input)
	if err != nil {
		// Update workflow state with error
		_ = w.Transactions.UpdateOnboardingWorkflowState(ctx, OnboardingWorkflowState{
			WorkflowID:   input.WorkflowID,
			Status:       StatusFailed,
			CurrentStep:  "subscription_failed",
			ErrorMessage: "Failed to create subscription: " + err.Error(),
		})
		return "", err
	}

	// Step 2: Verify payment
	paymentVerified, err := w.Steps.VerifyStripePayment(ctx, stripeSubID)
	if err != nil || !paymentVerified {
		errMsg := "Payment verification failed"
		if err != nil {
			errMsg = err.Error()
		}
		_ = w.Transactions.UpdateOnboardingWorkflowState(ctx, OnboardingWorkflowState{
			WorkflowID:   input.WorkflowID,
			Status:       StatusFailed,
			CurrentStep:  "payment_failed",
			ErrorMessage: errMsg,
		})
		return "", err
	}

	// Step 3: Create subscription record in database
	subscription := Subscription{
		UserID:          input.UserID,
		CompanyID:       input.CompanyID,
		Tier:            input.Tier,
		Status:          "active",
		BillingInterval: input.BillingInterval,
		PaymentProvider: "stripe",
		ProviderSubID:   stripeSubID,
	}

	subscriptionID, err := w.Transactions.CreateSubscription(ctx, subscription)
	if err != nil {
		return "", err
	}

	// Step 4: Activate subscription (updates both subscription and user status)
	err = w.Transactions.ActivateSubscription(ctx, subscriptionID, input.UserID)
	if err != nil {
		return "", err
	}

	// Step 5: Update workflow state
	err = w.Transactions.UpdateOnboardingWorkflowState(ctx, OnboardingWorkflowState{
		WorkflowID:     input.WorkflowID,
		Status:         StatusSubscribed,
		SubscriptionID: subscriptionID,
		CurrentStep:    "provisioning_dashboard",
	})
	if err != nil {
		return "", err
	}

	// Step 6: Provision dashboard access
	err = w.Steps.ProvisionDashboardAccess(ctx, input.UserID, input.CompanyID)
	if err != nil {
		// Log error but don't fail workflow
		// TODO: Add proper logging
	}

	// Step 7: Initialize default settings
	err = w.Steps.InitializeDefaultSettings(ctx, input.UserID, input.CompanyID)
	if err != nil {
		// Log error but don't fail workflow
		// TODO: Add proper logging
	}

	// Step 8: Send welcome email
	// For now, we'll skip complex user/company lookup and just send with basic info
	_ = w.Steps.SendWelcomeEmail(ctx, input.UserID, "", "")

	// Step 9: Notify admin of new signup
	_ = w.Steps.NotifyAdminNewSignup(ctx, input.UserID, input.CompanyID)

	// Step 10: Mark onboarding complete
	err = w.Transactions.UpdateOnboardingWorkflowState(ctx, OnboardingWorkflowState{
		WorkflowID:     input.WorkflowID,
		Status:         StatusOnboardingComplete,
		SubscriptionID: subscriptionID,
		CurrentStep:    "complete",
	})
	if err != nil {
		return "", err
	}

	return subscriptionID, nil
}
