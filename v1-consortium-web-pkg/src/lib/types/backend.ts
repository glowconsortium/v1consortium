// Organization and User Management Enums

export enum OrganizationType {
  INTERNAL = 'internal',
  CLIENT = 'client',
  PROVIDER = 'provider'
}

export enum UserRole {
  // Internal roles
  INTERNAL_SU = 'internal_su',
  INTERNAL_ADMIN = 'internal_admin',
  INTERNAL_SUPPORT = 'internal_support',
  
  // Client roles
  CLIENT_ADMIN = 'client_admin',
  DER = 'der',
  SAFETY_MANAGER = 'safety_manager',
  HR_MANAGER = 'hr_manager',
  EMPLOYEE = 'employee',
  
  // Provider roles
  MRO = 'mro',
  SAP = 'sap',
  MEDICAL_EXAMINER = 'medical_examiner'
}

// Subscription Management Enums

export enum SubscriptionTier {
  STARTER = 'starter',
  PROFESSIONAL = 'professional',
  ENTERPRISE = 'enterprise',
  CUSTOM = 'custom'
}

export enum SubscriptionStatus {
  ACTIVE = 'active',
  CANCELLED = 'cancelled',
  SUSPENDED = 'suspended',
  EXPIRED = 'expired',
  PENDING = 'pending'
}

// Testing and Compliance Enums

export enum TestType {
  PRE_EMPLOYMENT = 'pre_employment',
  RANDOM = 'random',
  POST_ACCIDENT = 'post_accident',
  REASONABLE_SUSPICION = 'reasonable_suspicion',
  RETURN_TO_DUTY = 'return_to_duty',
  FOLLOW_UP = 'follow_up'
}

export enum TestCategory {
  DRUG = 'drug',
  ALCOHOL = 'alcohol',
  DRUG_ALCOHOL = 'drug_alcohol'
}

export enum TestStatus {
  ORDERED = 'ordered',
  SCHEDULED = 'scheduled',
  IN_PROGRESS = 'in_progress',
  COMPLETED = 'completed',
  CANCELLED = 'cancelled',
  NO_SHOW = 'no_show'
}

export enum TestResult {
  NEGATIVE = 'negative',
  POSITIVE = 'positive',
  REFUSAL = 'refusal',
  ADULTERATED = 'adulterated',
  SUBSTITUTED = 'substituted',
  INVALID = 'invalid'
}

// Motor Vehicle Record Enums

export enum MVRStatus {
  ORDERED = 'ordered',
  RECEIVED = 'received',
  REVIEWED = 'reviewed',
  FLAGGED = 'flagged'
}

// Physical and Medical Enums

export enum PhysicalStatus {
  SCHEDULED = 'scheduled',
  COMPLETED = 'completed',
  FAILED = 'failed',
  EXPIRED = 'expired',
  PENDING_REVIEW = 'pending_review'
}

// Background Check Enums

export enum BackgroundCheckType {
  CRIMINAL_HISTORY = 'criminal_history',
  EMPLOYMENT_VERIFICATION = 'employment_verification',
  REFERENCE_CHECK = 'reference_check',
  EDUCATION_VERIFICATION = 'education_verification',
  LICENSE_VERIFICATION = 'license_verification',
  CREDIT_CHECK = 'credit_check'
}

export enum BackgroundCheckStatus {
  ORDERED = 'ordered',
  IN_PROGRESS = 'in_progress',
  COMPLETED = 'completed',
  REQUIRES_REVIEW = 'requires_review'
}

// Document Management Enums

export enum DocumentType {
  TEST_RESULT = 'test_result',
  MVR_REPORT = 'mvr_report',
  MEDICAL_CERTIFICATE = 'medical_certificate',
  TRAINING_CERTIFICATE = 'training_certificate',
  COMPLIANCE_DOCUMENT = 'compliance_document',
  POLICY_DOCUMENT = 'policy_document',
  BACKGROUND_REPORT = 'background_report'
}

// Notification Enums

export enum NotificationType {
  EMAIL = 'email',
  SMS = 'sms',
  IN_APP = 'in_app',
  PHONE = 'phone'
}

export enum NotificationPriority {
  LOW = 'low',
  NORMAL = 'normal',
  HIGH = 'high',
  URGENT = 'urgent'
}

// Workflow Management Enums

export enum WorkflowStatus {
  PENDING = 'pending',
  RUNNING = 'running',
  COMPLETED = 'completed',
  FAILED = 'failed',
  CANCELLED = 'cancelled'
}
