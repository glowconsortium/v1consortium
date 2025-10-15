package consts

type AuthProvider string

const (
	AuthProviderEmail  AuthProvider = "email"
	AuthProviderGoogle AuthProvider = "google"
	AuthProviderGithub AuthProvider = "github"
)

type SubscriptionTier string

const (
	TierStarter      SubscriptionTier = "starter"
	TierProfessional SubscriptionTier = "professional"
	TierEnterprise   SubscriptionTier = "enterprise"
)

type OnboardingStatus string

const (
	StatusPending            OnboardingStatus = "pending"
	StatusEmailVerified      OnboardingStatus = "email_verified"
	StatusSubscribed         OnboardingStatus = "subscribed"
	StatusOnboardingComplete OnboardingStatus = "complete"
	StatusFailed             OnboardingStatus = "failed"
)
