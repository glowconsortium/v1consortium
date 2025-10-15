package dbosworkflow

import "v1consortium/internal/consts"

type OnboardingWorkflowInput struct {
	UserID          string
	Email           string
	FirstName       string
	LastName        string
	Provider        consts.AuthProvider
	CompanyName     string
	CompanySize     string
	CompanyIndustry string
}

type SubscribeWorkflowInput struct {
	WorkflowID      string
	UserID          string
	CompanyID       string
	Tier            consts.SubscriptionTier
	PaymentMethodID string
	BillingInterval string
}
