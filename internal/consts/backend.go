package consts

type AuthProvider string

const (
	AuthProviderEmail  AuthProvider = "email"
	AuthProviderGoogle AuthProvider = "google"
	AuthProviderGithub AuthProvider = "github"
)

// Organization Types
type OrganizationType string

const (
	OrgTypeInternal OrganizationType = "internal"
	OrgTypeClient   OrganizationType = "client"
	OrgTypeProvider OrganizationType = "provider"
)

// User Roles
type UserRole string

const (
	RoleInternalSU      UserRole = "internal_su"
	RoleInternalAdmin   UserRole = "internal_admin"
	RoleInternalSupport UserRole = "internal_support"
	RoleClientAdmin     UserRole = "client_admin"
	RoleDER             UserRole = "der"
	RoleSafetyManager   UserRole = "safety_manager"
	RoleHRManager       UserRole = "hr_manager"
	RoleEmployee        UserRole = "employee"
	RoleMRO             UserRole = "mro"
	RoleSAP             UserRole = "sap"
	RoleMedicalExaminer UserRole = "medical_examiner"
)

// Subscription Tiers
type SubscriptionTier string

const (
	TierStarter      SubscriptionTier = "starter"
	TierProfessional SubscriptionTier = "professional"
	TierEnterprise   SubscriptionTier = "enterprise"
	TierCustom       SubscriptionTier = "custom"
)

// Subscription Status
type SubscriptionStatus string

const (
	SubStatusActive    SubscriptionStatus = "active"
	SubStatusCancelled SubscriptionStatus = "cancelled"
	SubStatusSuspended SubscriptionStatus = "suspended"
	SubStatusExpired   SubscriptionStatus = "expired"
	SubStatusPending   SubscriptionStatus = "pending"
)

// Test Types
type TestType string

const (
	TestTypePreEmployment       TestType = "pre_employment"
	TestTypeRandom              TestType = "random"
	TestTypePostAccident        TestType = "post_accident"
	TestTypeReasonableSuspicion TestType = "reasonable_suspicion"
	TestTypeReturnToDuty        TestType = "return_to_duty"
	TestTypeFollowUp            TestType = "follow_up"
)

// Test Categories
type TestCategory string

const (
	TestCategoryDrug        TestCategory = "drug"
	TestCategoryAlcohol     TestCategory = "alcohol"
	TestCategoryDrugAlcohol TestCategory = "drug_alcohol"
)

// Test Status
type TestStatus string

const (
	TestStatusOrdered    TestStatus = "ordered"
	TestStatusScheduled  TestStatus = "scheduled"
	TestStatusInProgress TestStatus = "in_progress"
	TestStatusCompleted  TestStatus = "completed"
	TestStatusCancelled  TestStatus = "cancelled"
	TestStatusNoShow     TestStatus = "no_show"
)

// Test Results
type TestResult string

const (
	TestResultNegative    TestResult = "negative"
	TestResultPositive    TestResult = "positive"
	TestResultRefusal     TestResult = "refusal"
	TestResultAdulterated TestResult = "adulterated"
	TestResultSubstituted TestResult = "substituted"
	TestResultInvalid     TestResult = "invalid"
)

// MVR Status
type MVRStatus string

const (
	MVRStatusOrdered  MVRStatus = "ordered"
	MVRStatusReceived MVRStatus = "received"
	MVRStatusReviewed MVRStatus = "reviewed"
	MVRStatusFlagged  MVRStatus = "flagged"
)

// Physical Status
type PhysicalStatus string

const (
	PhysicalStatusScheduled     PhysicalStatus = "scheduled"
	PhysicalStatusCompleted     PhysicalStatus = "completed"
	PhysicalStatusFailed        PhysicalStatus = "failed"
	PhysicalStatusExpired       PhysicalStatus = "expired"
	PhysicalStatusPendingReview PhysicalStatus = "pending_review"
)

// Background Check Types
type BackgroundCheckType string

const (
	BackgroundCriminalHistory        BackgroundCheckType = "criminal_history"
	BackgroundEmploymentVerification BackgroundCheckType = "employment_verification"
	BackgroundReferenceCheck         BackgroundCheckType = "reference_check"
	BackgroundEducationVerification  BackgroundCheckType = "education_verification"
	BackgroundLicenseVerification    BackgroundCheckType = "license_verification"
	BackgroundCreditCheck            BackgroundCheckType = "credit_check"
)

// Background Check Status
type BackgroundCheckStatus string

const (
	BackgroundStatusOrdered        BackgroundCheckStatus = "ordered"
	BackgroundStatusInProgress     BackgroundCheckStatus = "in_progress"
	BackgroundStatusCompleted      BackgroundCheckStatus = "completed"
	BackgroundStatusRequiresReview BackgroundCheckStatus = "requires_review"
)

// Document Types
type DocumentType string

const (
	DocTypeTestResult          DocumentType = "test_result"
	DocTypeMVRReport           DocumentType = "mvr_report"
	DocTypeMedicalCertificate  DocumentType = "medical_certificate"
	DocTypeTrainingCertificate DocumentType = "training_certificate"
	DocTypeComplianceDocument  DocumentType = "compliance_document"
	DocTypePolicyDocument      DocumentType = "policy_document"
	DocTypeBackgroundReport    DocumentType = "background_report"
)

// Notification Types
type NotificationType string

const (
	NotificationEmail NotificationType = "email"
	NotificationSMS   NotificationType = "sms"
	NotificationInApp NotificationType = "in_app"
	NotificationPhone NotificationType = "phone"
)

// Notification Priorities
type NotificationPriority string

const (
	NotificationPriorityLow    NotificationPriority = "low"
	NotificationPriorityNormal NotificationPriority = "normal"
	NotificationPriorityHigh   NotificationPriority = "high"
	NotificationPriorityUrgent NotificationPriority = "urgent"
)

// Workflow Status
type WorkflowStatus string

const (
	WorkflowStatusPending   WorkflowStatus = "pending"
	WorkflowStatusRunning   WorkflowStatus = "running"
	WorkflowStatusCompleted WorkflowStatus = "completed"
	WorkflowStatusFailed    WorkflowStatus = "failed"
	WorkflowStatusCancelled WorkflowStatus = "cancelled"
)

type OnboardingStatus string

const (
	StatusPending            OnboardingStatus = "pending"
	StatusEmailVerified      OnboardingStatus = "email_verified"
	StatusSubscribed         OnboardingStatus = "subscribed"
	StatusOnboardingComplete OnboardingStatus = "complete"
	StatusFailed             OnboardingStatus = "failed"
)
