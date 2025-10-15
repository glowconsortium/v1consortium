package dbosworkflow

import "context"

type OnboardingSteps struct{}

// SendVerificationEmail sends email verification link (idempotent)
func (s *OnboardingSteps) SendVerificationEmail(ctx context.Context, userID, email, firstName, token string) error {
	// Call email service (Sendgrid, SES, etc.)
	// Implementation would use your email provider
	verificationLink := "https://yourapp.com/verify/" + token

	// Pseudo-code for email sending
	// return emailService.Send(EmailRequest{
	// 	To: email,
	// 	Subject: "Verify your email",
	// 	Template: "verification",
	// 	Data: map[string]string{
	// 		"firstName": firstName,
	// 		"link": verificationLink,
	// 	},
	// })

	_ = verificationLink // Avoid unused variable error
	return nil
}

// CreateStripeSubscription creates subscription in Stripe (idempotent)
func (s *OnboardingSteps) CreateStripeSubscription(ctx context.Context, input SubscribeWorkflowInput) (string, error) {
	// Call Stripe API to create subscription
	// Returns Stripe subscription ID

	// Pseudo-code:
	// subscription, err := stripe.Subscriptions.Create(&stripe.SubscriptionParams{
	// 	Customer: customerID,
	// 	Items: []*stripe.SubscriptionItemsParams{
	// 		{Price: getPriceID(input.Tier, input.BillingInterval)},
	// 	},
	// 	PaymentMethod: input.PaymentMethodID,
	// })
	// return subscription.ID, err

	return "sub_mock_" + input.UserID, nil
}

// VerifyStripePayment verifies payment was successful (idempotent)
func (s *OnboardingSteps) VerifyStripePayment(ctx context.Context, subscriptionID string) (bool, error) {
	// Check Stripe subscription status
	// Pseudo-code:
	// sub, err := stripe.Subscriptions.Get(subscriptionID, nil)
	// if err != nil {
	// 	return false, err
	// }
	// return sub.Status == "active" || sub.Status == "trialing", nil

	return true, nil
}

// ProvisionDashboardAccess sets up user dashboard access (idempotent)
func (s *OnboardingSteps) ProvisionDashboardAccess(ctx context.Context, userID, companyID string) error {
	// Create default workspace, permissions, resources, etc.
	// This could call internal services or APIs
	return nil
}

// SendWelcomeEmail sends welcome email (idempotent)
func (s *OnboardingSteps) SendWelcomeEmail(ctx context.Context, email, firstName, companyName string) error {
	// Send welcome email with getting started guide
	return nil
}

// InitializeDefaultSettings creates default settings (idempotent)
func (s *OnboardingSteps) InitializeDefaultSettings(ctx context.Context, userID, companyID string) error {
	// Create default preferences, settings, templates, etc.
	return nil
}

// NotifyAdminNewSignup sends notification to admin team (idempotent)
func (s *OnboardingSteps) NotifyAdminNewSignup(ctx context.Context, userID, companyID string) error {
	// Send Slack notification, email to sales, etc.
	return nil
}
