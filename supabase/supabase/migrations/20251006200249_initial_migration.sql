-- V1 Consortium Database Schema
-- Comprehensive compliance and screening services platform
-- Date: October 6, 2025

-- Enable necessary extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "citext";

-- Create enums for various types
CREATE TYPE organization_type AS ENUM ('internal', 'client', 'provider');
CREATE TYPE user_role AS ENUM (
    'internal_su', 'internal_admin', 'internal_support',
    'client_admin', 'der', 'safety_manager', 'hr_manager', 'employee',
    'mro', 'sap', 'medical_examiner'
);
CREATE TYPE subscription_tier AS ENUM ('starter', 'professional', 'enterprise', 'custom');
CREATE TYPE subscription_status AS ENUM ('active', 'cancelled', 'suspended', 'expired', 'pending');
CREATE TYPE test_type AS ENUM (
    'pre_employment', 'random', 'post_accident', 'reasonable_suspicion', 
    'return_to_duty', 'follow_up'
);
CREATE TYPE test_category AS ENUM ('drug', 'alcohol', 'drug_alcohol');
CREATE TYPE test_status AS ENUM ('ordered', 'scheduled', 'in_progress', 'completed', 'cancelled', 'no_show');
CREATE TYPE test_result AS ENUM ('negative', 'positive', 'refusal', 'adulterated', 'substituted', 'invalid');
CREATE TYPE mvr_status AS ENUM ('ordered', 'received', 'reviewed', 'flagged');
CREATE TYPE physical_status AS ENUM ('scheduled', 'completed', 'failed', 'expired', 'pending_review');
CREATE TYPE background_check_type AS ENUM (
    'criminal_history', 'employment_verification', 'reference_check',
    'education_verification', 'license_verification', 'credit_check'
);
CREATE TYPE background_check_status AS ENUM ('ordered', 'in_progress', 'completed', 'requires_review');
CREATE TYPE document_type AS ENUM (
    'test_result', 'mvr_report', 'medical_certificate', 'training_certificate',
    'compliance_document', 'policy_document', 'background_report'
);
CREATE TYPE notification_type AS ENUM ('email', 'sms', 'in_app', 'phone');
CREATE TYPE notification_priority AS ENUM ('low', 'normal', 'high', 'urgent');
CREATE TYPE workflow_status AS ENUM ('pending', 'running', 'completed', 'failed', 'cancelled');

-- =============================================
-- CORE TABLES
-- =============================================

-- Organizations table (multi-tenant base)
CREATE TABLE organizations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    type organization_type NOT NULL,
    usdot_number VARCHAR(20),
    mc_number VARCHAR(20),
    industry VARCHAR(100),
    is_dot_regulated BOOLEAN DEFAULT false,
    address_line1 VARCHAR(255),
    address_line2 VARCHAR(255),
    city VARCHAR(100),
    state VARCHAR(10),
    zip_code VARCHAR(20),
    country VARCHAR(10) DEFAULT 'US',
    phone VARCHAR(20),
    email VARCHAR(255),
    website VARCHAR(255),
    tax_id VARCHAR(50),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    is_active BOOLEAN DEFAULT true,
    settings JSONB DEFAULT '{}'::jsonb,
    
    CONSTRAINT organizations_email_check CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);

-- Subscription plans
CREATE TABLE subscription_plans (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    tier subscription_tier NOT NULL,
    description TEXT,
    max_employees INTEGER,
    annual_price DECIMAL(10,2),
    features JSONB DEFAULT '{}'::jsonb,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Organization subscriptions
CREATE TABLE organization_subscriptions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    plan_id UUID NOT NULL REFERENCES subscription_plans(id),
    status subscription_status DEFAULT 'pending',
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    auto_renew BOOLEAN DEFAULT true,
    stripe_subscription_id VARCHAR(255),
    stripe_customer_id VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    
    CONSTRAINT valid_subscription_period CHECK (end_date > start_date)
);

-- Users table (extends Supabase auth.users)
CREATE TABLE user_profiles (
    id UUID PRIMARY KEY REFERENCES auth.users(id) ON DELETE CASCADE,
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    role user_role NOT NULL,
    employee_id VARCHAR(50),
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email CITEXT NOT NULL,
    phone VARCHAR(20),
    date_of_birth DATE,
    ssn_last_four VARCHAR(4),
    hire_date DATE,
    termination_date DATE,
    is_active BOOLEAN DEFAULT true,
    requires_dot_testing BOOLEAN DEFAULT false,
    requires_non_dot_testing BOOLEAN DEFAULT false,
    cdl_number VARCHAR(50),
    cdl_state VARCHAR(10),
    cdl_expiration_date DATE,
    job_title VARCHAR(100),
    department VARCHAR(100),
    supervisor_id UUID REFERENCES user_profiles(id),
    emergency_contact_name VARCHAR(255),
    emergency_contact_phone VARCHAR(20),
    emergency_contact_relationship VARCHAR(50),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    last_login_at TIMESTAMPTZ,
    
    CONSTRAINT user_profiles_email_check CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    CONSTRAINT ssn_last_four_check CHECK (ssn_last_four ~ '^[0-9]{4}$' OR ssn_last_four IS NULL)
);

-- =============================================
-- TESTING PROGRAM TABLES
-- =============================================

-- Testing programs configuration
CREATE TABLE testing_programs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    is_dot_program BOOLEAN DEFAULT false,
    drug_panel_type VARCHAR(50) DEFAULT '5-panel',
    alcohol_testing_enabled BOOLEAN DEFAULT false,
    random_testing_enabled BOOLEAN DEFAULT true,
    random_testing_rate DECIMAL(5,2) DEFAULT 50.00, -- percentage
    testing_frequency VARCHAR(20) DEFAULT 'quarterly', -- monthly, quarterly, annually
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Random testing pools
CREATE TABLE random_testing_pools (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    program_id UUID NOT NULL REFERENCES testing_programs(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Pool memberships (many-to-many between users and pools)
CREATE TABLE pool_memberships (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    pool_id UUID NOT NULL REFERENCES random_testing_pools(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES user_profiles(id) ON DELETE CASCADE,
    joined_at TIMESTAMPTZ DEFAULT NOW(),
    left_at TIMESTAMPTZ,
    is_active BOOLEAN DEFAULT true,
    
    UNIQUE(pool_id, user_id)
);

-- Random selections (audit trail)
CREATE TABLE random_selections (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    pool_id UUID NOT NULL REFERENCES random_testing_pools(id) ON DELETE CASCADE,
    selection_date DATE NOT NULL,
    selection_period VARCHAR(20), -- Q1-2025, Jan-2025, etc.
    total_pool_size INTEGER NOT NULL,
    required_selections INTEGER NOT NULL,
    selection_algorithm VARCHAR(50) DEFAULT 'random',
    selection_seed VARCHAR(255), -- for reproducibility
    notes TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    created_by UUID NOT NULL REFERENCES user_profiles(id)
);

-- Drug and alcohol tests
CREATE TABLE drug_alcohol_tests (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES user_profiles(id) ON DELETE CASCADE,
    program_id UUID REFERENCES testing_programs(id),
    selection_id UUID REFERENCES random_selections(id), -- if from random selection
    test_type test_type NOT NULL,
    test_category test_category NOT NULL,
    status test_status DEFAULT 'ordered',
    result test_result,
    is_dot_test BOOLEAN DEFAULT false,
    
    -- Test ordering information
    ordered_date TIMESTAMPTZ DEFAULT NOW(),
    ordered_by UUID NOT NULL REFERENCES user_profiles(id),
    due_date DATE,
    
    -- Third-party integration data
    external_order_id VARCHAR(255), -- Quest Diagnostics order ID
    external_facility_id VARCHAR(255),
    facility_name VARCHAR(255),
    facility_address TEXT,
    
    -- Test completion information
    collection_date TIMESTAMPTZ,
    collected_by VARCHAR(255),
    lab_id VARCHAR(255),
    lab_accession_number VARCHAR(255),
    
    -- Results information
    result_date TIMESTAMPTZ,
    result_received_date TIMESTAMPTZ,
    mro_review_required BOOLEAN DEFAULT false,
    mro_id UUID REFERENCES user_profiles(id),
    mro_review_date TIMESTAMPTZ,
    mro_notes TEXT,
    
    -- Follow-up actions
    requires_immediate_removal BOOLEAN DEFAULT false,
    return_to_duty_required BOOLEAN DEFAULT false,
    follow_up_tests_required INTEGER DEFAULT 0,
    
    -- Metadata
    notes TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Selected users for each random selection
CREATE TABLE random_selection_members (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    selection_id UUID NOT NULL REFERENCES random_selections(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES user_profiles(id) ON DELETE CASCADE,
    test_id UUID REFERENCES drug_alcohol_tests(id), -- now references the table created above
    selection_order INTEGER, -- 1st selected, 2nd selected, etc.
    
    UNIQUE(selection_id, user_id)
);

-- =============================================
-- MVR (Motor Vehicle Record) TABLES
-- =============================================

-- MVR reports
CREATE TABLE mvr_reports (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES user_profiles(id) ON DELETE CASCADE,
    status mvr_status DEFAULT 'ordered',
    
    -- Order information
    ordered_date TIMESTAMPTZ DEFAULT NOW(),
    ordered_by UUID NOT NULL REFERENCES user_profiles(id),
    
    -- Driver information
    license_number VARCHAR(50) NOT NULL,
    license_state VARCHAR(10) NOT NULL,
    
    -- External provider data
    external_order_id VARCHAR(255),
    provider_name VARCHAR(100),
    
    -- Report data
    report_date DATE,
    report_received_date TIMESTAMPTZ,
    raw_report_data JSONB,
    
    -- Violation summary
    total_violations INTEGER DEFAULT 0,
    major_violations INTEGER DEFAULT 0,
    minor_violations INTEGER DEFAULT 0,
    license_status VARCHAR(50),
    license_expiration_date DATE,
    
    -- Review information
    reviewed_by UUID REFERENCES user_profiles(id),
    reviewed_date TIMESTAMPTZ,
    requires_action BOOLEAN DEFAULT false,
    action_notes TEXT,
    
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Individual MVR violations
CREATE TABLE mvr_violations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    mvr_report_id UUID NOT NULL REFERENCES mvr_reports(id) ON DELETE CASCADE,
    violation_date DATE,
    violation_code VARCHAR(20),
    violation_description TEXT,
    violation_type VARCHAR(50), -- speeding, DUI, reckless driving, etc.
    severity VARCHAR(20), -- minor, major, serious
    conviction_date DATE,
    fine_amount DECIMAL(10,2),
    points INTEGER,
    state VARCHAR(10),
    court_name VARCHAR(255),
    case_number VARCHAR(100),
    
    -- Compliance impact
    disqualifying BOOLEAN DEFAULT false,
    requires_employer_notification BOOLEAN DEFAULT false,
    affects_cdl BOOLEAN DEFAULT false,
    
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- =============================================
-- DOT PHYSICAL TABLES
-- =============================================

-- DOT physical examinations
CREATE TABLE dot_physicals (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES user_profiles(id) ON DELETE CASCADE,
    status physical_status DEFAULT 'scheduled',
    
    -- Scheduling information
    scheduled_date TIMESTAMPTZ,
    scheduled_by UUID REFERENCES user_profiles(id),
    
    -- Medical examiner information
    examiner_id UUID REFERENCES user_profiles(id),
    examiner_name VARCHAR(255),
    examiner_license_number VARCHAR(100),
    examiner_registry_number VARCHAR(100),
    clinic_name VARCHAR(255),
    clinic_address TEXT,
    clinic_phone VARCHAR(20),
    
    -- Examination results
    examination_date DATE,
    certificate_number VARCHAR(100),
    certificate_issue_date DATE,
    certificate_expiration_date DATE,
    
    -- Medical status
    medical_qualification VARCHAR(50), -- qualified, disqualified, etc.
    restrictions TEXT,
    exemptions TEXT,
    
    -- Medical conditions tracking
    requires_monitoring BOOLEAN DEFAULT false,
    monitoring_requirements TEXT,
    next_required_date DATE,
    
    -- Certificate management
    certificate_url VARCHAR(500), -- Supabase storage URL
    certificate_uploaded_at TIMESTAMPTZ,
    
    notes TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- =============================================
-- BACKGROUND CHECK TABLES
-- =============================================

-- Background check orders
CREATE TABLE background_checks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES user_profiles(id) ON DELETE CASCADE,
    check_type background_check_type NOT NULL,
    status background_check_status DEFAULT 'ordered',
    
    -- Order information
    ordered_date TIMESTAMPTZ DEFAULT NOW(),
    ordered_by UUID NOT NULL REFERENCES user_profiles(id),
    
    -- External provider information
    external_order_id VARCHAR(255),
    provider_name VARCHAR(100),
    
    -- Completion information
    completed_date TIMESTAMPTZ,
    report_date DATE,
    
    -- Results summary
    overall_result VARCHAR(50), -- clear, consider, not_clear
    requires_review BOOLEAN DEFAULT false,
    adverse_action_required BOOLEAN DEFAULT false,
    
    -- FCRA compliance
    fcra_disclosure_sent BOOLEAN DEFAULT false,
    fcra_authorization_received BOOLEAN DEFAULT false,
    pre_adverse_action_sent BOOLEAN DEFAULT false,
    adverse_action_sent BOOLEAN DEFAULT false,
    
    notes TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Background check findings
CREATE TABLE background_check_findings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    background_check_id UUID NOT NULL REFERENCES background_checks(id) ON DELETE CASCADE,
    finding_type VARCHAR(100), -- criminal_record, employment_gap, etc.
    severity VARCHAR(20), -- low, medium, high
    description TEXT,
    date_of_record DATE,
    jurisdiction VARCHAR(100),
    case_number VARCHAR(100),
    disposition VARCHAR(255),
    
    -- Impact assessment
    job_related BOOLEAN,
    disqualifying BOOLEAN DEFAULT false,
    requires_individualized_assessment BOOLEAN DEFAULT false,
    
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- =============================================
-- DOCUMENT MANAGEMENT TABLES
-- =============================================

-- Document storage and management
CREATE TABLE documents (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    user_id UUID REFERENCES user_profiles(id), -- NULL for organization-level docs
    document_type document_type NOT NULL,
    
    -- Document metadata
    title VARCHAR(255) NOT NULL,
    description TEXT,
    file_name VARCHAR(255) NOT NULL,
    file_size BIGINT,
    mime_type VARCHAR(100),
    
    -- Storage information
    storage_path VARCHAR(500) NOT NULL, -- Supabase storage path
    storage_bucket VARCHAR(100) DEFAULT 'documents',
    
    -- Related entity references
    test_id UUID REFERENCES drug_alcohol_tests(id),
    mvr_report_id UUID REFERENCES mvr_reports(id),
    physical_id UUID REFERENCES dot_physicals(id),
    background_check_id UUID REFERENCES background_checks(id),
    
    -- Document lifecycle
    uploaded_by UUID NOT NULL REFERENCES user_profiles(id),
    uploaded_at TIMESTAMPTZ DEFAULT NOW(),
    
    -- Security and compliance
    is_confidential BOOLEAN DEFAULT true,
    is_hipaa_protected BOOLEAN DEFAULT false,
    retention_period_years INTEGER,
    auto_delete_date DATE,
    
    -- Access tracking
    download_count INTEGER DEFAULT 0,
    last_accessed_at TIMESTAMPTZ,
    last_accessed_by UUID REFERENCES user_profiles(id),
    
    -- Versioning
    version INTEGER DEFAULT 1,
    parent_document_id UUID REFERENCES documents(id),
    is_current_version BOOLEAN DEFAULT true,
    
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- =============================================
-- WORKFLOW AND AUTOMATION TABLES
-- =============================================

-- Temporal workflow tracking
CREATE TABLE temporal_workflows (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    workflow_id VARCHAR(255) NOT NULL, -- Temporal workflow ID
    workflow_type VARCHAR(100) NOT NULL, -- test_ordering, random_selection, etc.
    status workflow_status DEFAULT 'pending',
    
    -- Workflow metadata
    input_data JSONB,
    output_data JSONB,
    error_message TEXT,
    
    -- Related entities
    user_id UUID REFERENCES user_profiles(id),
    test_id UUID REFERENCES drug_alcohol_tests(id),
    selection_id UUID REFERENCES random_selections(id),
    
    -- Timing
    started_at TIMESTAMPTZ DEFAULT NOW(),
    completed_at TIMESTAMPTZ,
    scheduled_for TIMESTAMPTZ,
    
    -- Retry logic
    retry_count INTEGER DEFAULT 0,
    max_retries INTEGER DEFAULT 3,
    next_retry_at TIMESTAMPTZ,
    
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Notification system
CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    user_id UUID REFERENCES user_profiles(id), -- NULL for organization-wide notifications
    
    -- Notification content
    title VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    notification_type notification_type DEFAULT 'in_app',
    priority notification_priority DEFAULT 'normal',
    
    -- Delivery information
    sent_at TIMESTAMPTZ,
    delivered_at TIMESTAMPTZ,
    read_at TIMESTAMPTZ,
    
    -- Channel-specific data
    email_address VARCHAR(255),
    phone_number VARCHAR(20),
    external_message_id VARCHAR(255), -- For tracking with email/SMS providers
    
    -- Related entities
    test_id UUID REFERENCES drug_alcohol_tests(id),
    mvr_report_id UUID REFERENCES mvr_reports(id),
    physical_id UUID REFERENCES dot_physicals(id),
    
    -- Workflow integration
    workflow_id UUID REFERENCES temporal_workflows(id),
    
    -- Status tracking
    delivery_attempts INTEGER DEFAULT 0,
    last_attempt_at TIMESTAMPTZ,
    delivery_error TEXT,
    
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- =============================================
-- AUDIT AND COMPLIANCE TABLES
-- =============================================

-- System audit trail
CREATE TABLE audit_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID REFERENCES organizations(id),
    user_id UUID REFERENCES user_profiles(id),
    
    -- Action details
    action VARCHAR(100) NOT NULL, -- create, update, delete, view, download, etc.
    entity_type VARCHAR(100) NOT NULL, -- user, test, document, etc.
    entity_id UUID,
    
    -- Change tracking
    old_values JSONB,
    new_values JSONB,
    
    -- Request context
    ip_address INET,
    user_agent TEXT,
    request_id VARCHAR(255),
    
    -- Compliance requirements
    retention_required BOOLEAN DEFAULT true,
    hipaa_log BOOLEAN DEFAULT false,
    
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- Compliance status tracking
CREATE TABLE compliance_status (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES user_profiles(id) ON DELETE CASCADE,
    
    -- Overall compliance status
    is_compliant BOOLEAN DEFAULT false,
    compliance_percentage DECIMAL(5,2) DEFAULT 0.00,
    last_updated TIMESTAMPTZ DEFAULT NOW(),
    
    -- Drug testing compliance
    drug_testing_current BOOLEAN DEFAULT false,
    last_drug_test_date DATE,
    next_drug_test_due DATE,
    
    -- MVR compliance
    mvr_current BOOLEAN DEFAULT false,
    last_mvr_date DATE,
    next_mvr_due DATE,
    
    -- DOT physical compliance
    physical_current BOOLEAN DEFAULT false,
    medical_cert_expiration_date DATE,
    
    -- Background check compliance
    background_check_current BOOLEAN DEFAULT false,
    last_background_check_date DATE,
    
    -- Training compliance
    training_current BOOLEAN DEFAULT false,
    last_training_date DATE,
    
    -- Risk factors
    violations_count INTEGER DEFAULT 0,
    high_risk_flags INTEGER DEFAULT 0,
    
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    
    -- Unique constraint: one compliance status per user
    UNIQUE(user_id)
);

-- =============================================
-- REPORTING AND ANALYTICS TABLES
-- =============================================

-- Saved reports and analytics
CREATE TABLE saved_reports (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    created_by UUID NOT NULL REFERENCES user_profiles(id),
    
    -- Report metadata
    name VARCHAR(255) NOT NULL,
    description TEXT,
    report_type VARCHAR(100) NOT NULL, -- compliance_summary, random_selection, violations, etc.
    
    -- Report configuration
    parameters JSONB DEFAULT '{}'::jsonb,
    filters JSONB DEFAULT '{}'::jsonb,
    
    -- Report output
    generated_at TIMESTAMPTZ,
    file_path VARCHAR(500), -- Path to generated report file
    file_format VARCHAR(20), -- pdf, csv, xlsx
    
    -- Scheduling
    is_scheduled BOOLEAN DEFAULT false,
    schedule_frequency VARCHAR(50), -- daily, weekly, monthly, quarterly
    next_run_date TIMESTAMPTZ,
    
    -- Access control
    is_public BOOLEAN DEFAULT false,
    shared_with UUID[], -- Array of user IDs
    
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Certificate generation tracking
CREATE TABLE certificates (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    user_id UUID REFERENCES user_profiles(id), -- NULL for organization certificates
    
    -- Certificate metadata
    certificate_type VARCHAR(100) NOT NULL, -- consortium_membership, test_completion, etc.
    title VARCHAR(255) NOT NULL,
    description TEXT,
    
    -- Certificate data
    certificate_number VARCHAR(100) UNIQUE NOT NULL,
    issue_date DATE NOT NULL,
    expiration_date DATE,
    
    -- Related entities
    test_id UUID REFERENCES drug_alcohol_tests(id),
    physical_id UUID REFERENCES dot_physicals(id),
    
    -- File storage
    certificate_url VARCHAR(500), -- Supabase storage URL
    template_used VARCHAR(100),
    
    -- Digital signature
    is_digitally_signed BOOLEAN DEFAULT true,
    signature_hash VARCHAR(255),
    signature_timestamp TIMESTAMPTZ,
    
    -- Access tracking
    download_count INTEGER DEFAULT 0,
    last_downloaded_at TIMESTAMPTZ,
    last_downloaded_by UUID REFERENCES user_profiles(id),
    
    -- Status
    is_revoked BOOLEAN DEFAULT false,
    revoked_at TIMESTAMPTZ,
    revoked_by UUID REFERENCES user_profiles(id),
    revocation_reason TEXT,
    
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- =============================================
-- INDEXES FOR PERFORMANCE OPTIMIZATION
-- =============================================

-- Organizations indexes
CREATE INDEX idx_organizations_type ON organizations(type);
CREATE INDEX idx_organizations_usdot ON organizations(usdot_number) WHERE usdot_number IS NOT NULL;
CREATE INDEX idx_organizations_active ON organizations(is_active);

-- User profiles indexes
CREATE INDEX idx_user_profiles_org ON user_profiles(organization_id);
CREATE INDEX idx_user_profiles_role ON user_profiles(role);
CREATE INDEX idx_user_profiles_active ON user_profiles(is_active);
CREATE INDEX idx_user_profiles_email ON user_profiles(email);
CREATE INDEX idx_user_profiles_supervisor ON user_profiles(supervisor_id);
CREATE INDEX idx_user_profiles_cdl ON user_profiles(cdl_number, cdl_state) WHERE cdl_number IS NOT NULL;

-- Subscription indexes
CREATE INDEX idx_org_subscriptions_org ON organization_subscriptions(organization_id);
CREATE INDEX idx_org_subscriptions_status ON organization_subscriptions(status);
CREATE INDEX idx_org_subscriptions_dates ON organization_subscriptions(start_date, end_date);
CREATE INDEX idx_org_subscriptions_stripe ON organization_subscriptions(stripe_subscription_id) WHERE stripe_subscription_id IS NOT NULL;

-- Testing program indexes
CREATE INDEX idx_testing_programs_org ON testing_programs(organization_id);
CREATE INDEX idx_random_pools_org ON random_testing_pools(organization_id);
CREATE INDEX idx_random_pools_program ON random_testing_pools(program_id);
CREATE INDEX idx_pool_memberships_pool ON pool_memberships(pool_id);
CREATE INDEX idx_pool_memberships_user ON pool_memberships(user_id);
CREATE INDEX idx_pool_memberships_active ON pool_memberships(is_active);

-- Random selection indexes
CREATE INDEX idx_random_selections_pool ON random_selections(pool_id);
CREATE INDEX idx_random_selections_date ON random_selections(selection_date);
CREATE INDEX idx_random_selection_members_selection ON random_selection_members(selection_id);
CREATE INDEX idx_random_selection_members_user ON random_selection_members(user_id);

-- Drug/alcohol test indexes
CREATE INDEX idx_drug_tests_org ON drug_alcohol_tests(organization_id);
CREATE INDEX idx_drug_tests_user ON drug_alcohol_tests(user_id);
CREATE INDEX idx_drug_tests_status ON drug_alcohol_tests(status);
CREATE INDEX idx_drug_tests_type ON drug_alcohol_tests(test_type);
CREATE INDEX idx_drug_tests_ordered_date ON drug_alcohol_tests(ordered_date);
CREATE INDEX idx_drug_tests_due_date ON drug_alcohol_tests(due_date) WHERE due_date IS NOT NULL;
CREATE INDEX idx_drug_tests_external_id ON drug_alcohol_tests(external_order_id) WHERE external_order_id IS NOT NULL;
CREATE INDEX idx_drug_tests_selection ON drug_alcohol_tests(selection_id) WHERE selection_id IS NOT NULL;
CREATE INDEX idx_drug_tests_dot ON drug_alcohol_tests(is_dot_test);

-- MVR indexes
CREATE INDEX idx_mvr_reports_org ON mvr_reports(organization_id);
CREATE INDEX idx_mvr_reports_user ON mvr_reports(user_id);
CREATE INDEX idx_mvr_reports_status ON mvr_reports(status);
CREATE INDEX idx_mvr_reports_ordered_date ON mvr_reports(ordered_date);
CREATE INDEX idx_mvr_reports_license ON mvr_reports(license_number, license_state);
CREATE INDEX idx_mvr_violations_report ON mvr_violations(mvr_report_id);
CREATE INDEX idx_mvr_violations_date ON mvr_violations(violation_date);
CREATE INDEX idx_mvr_violations_type ON mvr_violations(violation_type);

-- DOT physical indexes
CREATE INDEX idx_dot_physicals_org ON dot_physicals(organization_id);
CREATE INDEX idx_dot_physicals_user ON dot_physicals(user_id);
CREATE INDEX idx_dot_physicals_status ON dot_physicals(status);
CREATE INDEX idx_dot_physicals_scheduled ON dot_physicals(scheduled_date) WHERE scheduled_date IS NOT NULL;
CREATE INDEX idx_dot_physicals_expiration ON dot_physicals(certificate_expiration_date) WHERE certificate_expiration_date IS NOT NULL;
CREATE INDEX idx_dot_physicals_examiner ON dot_physicals(examiner_id) WHERE examiner_id IS NOT NULL;

-- Background check indexes
CREATE INDEX idx_background_checks_org ON background_checks(organization_id);
CREATE INDEX idx_background_checks_user ON background_checks(user_id);
CREATE INDEX idx_background_checks_type ON background_checks(check_type);
CREATE INDEX idx_background_checks_status ON background_checks(status);
CREATE INDEX idx_background_checks_ordered ON background_checks(ordered_date);
CREATE INDEX idx_background_findings_check ON background_check_findings(background_check_id);

-- Document indexes
CREATE INDEX idx_documents_org ON documents(organization_id);
CREATE INDEX idx_documents_user ON documents(user_id) WHERE user_id IS NOT NULL;
CREATE INDEX idx_documents_type ON documents(document_type);
CREATE INDEX idx_documents_uploaded ON documents(uploaded_at);
CREATE INDEX idx_documents_test ON documents(test_id) WHERE test_id IS NOT NULL;
CREATE INDEX idx_documents_mvr ON documents(mvr_report_id) WHERE mvr_report_id IS NOT NULL;
CREATE INDEX idx_documents_physical ON documents(physical_id) WHERE physical_id IS NOT NULL;
CREATE INDEX idx_documents_background ON documents(background_check_id) WHERE background_check_id IS NOT NULL;

-- Workflow indexes
CREATE INDEX idx_temporal_workflows_org ON temporal_workflows(organization_id);
CREATE INDEX idx_temporal_workflows_id ON temporal_workflows(workflow_id);
CREATE INDEX idx_temporal_workflows_type ON temporal_workflows(workflow_type);
CREATE INDEX idx_temporal_workflows_status ON temporal_workflows(status);
CREATE INDEX idx_temporal_workflows_scheduled ON temporal_workflows(scheduled_for) WHERE scheduled_for IS NOT NULL;

-- Notification indexes
CREATE INDEX idx_notifications_org ON notifications(organization_id);
CREATE INDEX idx_notifications_user ON notifications(user_id) WHERE user_id IS NOT NULL;
CREATE INDEX idx_notifications_type ON notifications(notification_type);
CREATE INDEX idx_notifications_sent ON notifications(sent_at) WHERE sent_at IS NOT NULL;
CREATE INDEX idx_notifications_read ON notifications(read_at) WHERE read_at IS NOT NULL;

-- Audit log indexes
CREATE INDEX idx_audit_logs_org ON audit_logs(organization_id) WHERE organization_id IS NOT NULL;
CREATE INDEX idx_audit_logs_user ON audit_logs(user_id) WHERE user_id IS NOT NULL;
CREATE INDEX idx_audit_logs_action ON audit_logs(action);
CREATE INDEX idx_audit_logs_entity ON audit_logs(entity_type, entity_id);
CREATE INDEX idx_audit_logs_created ON audit_logs(created_at);

-- Compliance status indexes
CREATE INDEX idx_compliance_status_org ON compliance_status(organization_id);
CREATE INDEX idx_compliance_status_user ON compliance_status(user_id);
CREATE INDEX idx_compliance_status_compliant ON compliance_status(is_compliant);
CREATE INDEX idx_compliance_status_updated ON compliance_status(last_updated);

-- Report and certificate indexes
CREATE INDEX idx_saved_reports_org ON saved_reports(organization_id);
CREATE INDEX idx_saved_reports_creator ON saved_reports(created_by);
CREATE INDEX idx_saved_reports_type ON saved_reports(report_type);
CREATE INDEX idx_certificates_org ON certificates(organization_id);
CREATE INDEX idx_certificates_user ON certificates(user_id) WHERE user_id IS NOT NULL;
CREATE INDEX idx_certificates_type ON certificates(certificate_type);
CREATE INDEX idx_certificates_number ON certificates(certificate_number);
CREATE INDEX idx_certificates_expiration ON certificates(expiration_date) WHERE expiration_date IS NOT NULL;

-- =============================================
-- TRIGGERS FOR AUTOMATIC TIMESTAMP UPDATES
-- =============================================

-- Function to update the updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Apply the trigger to tables with updated_at columns
CREATE TRIGGER update_organizations_updated_at BEFORE UPDATE ON organizations
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_subscription_plans_updated_at BEFORE UPDATE ON subscription_plans
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_organization_subscriptions_updated_at BEFORE UPDATE ON organization_subscriptions
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_user_profiles_updated_at BEFORE UPDATE ON user_profiles
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_testing_programs_updated_at BEFORE UPDATE ON testing_programs
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_random_testing_pools_updated_at BEFORE UPDATE ON random_testing_pools
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_drug_alcohol_tests_updated_at BEFORE UPDATE ON drug_alcohol_tests
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_mvr_reports_updated_at BEFORE UPDATE ON mvr_reports
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_dot_physicals_updated_at BEFORE UPDATE ON dot_physicals
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_background_checks_updated_at BEFORE UPDATE ON background_checks
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_documents_updated_at BEFORE UPDATE ON documents
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_temporal_workflows_updated_at BEFORE UPDATE ON temporal_workflows
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_notifications_updated_at BEFORE UPDATE ON notifications
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_compliance_status_updated_at BEFORE UPDATE ON compliance_status
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_saved_reports_updated_at BEFORE UPDATE ON saved_reports
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_certificates_updated_at BEFORE UPDATE ON certificates
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- =============================================
-- ROW LEVEL SECURITY (RLS) POLICIES
-- =============================================

-- Enable RLS on all tables
ALTER TABLE organizations ENABLE ROW LEVEL SECURITY;
ALTER TABLE subscription_plans ENABLE ROW LEVEL SECURITY;
ALTER TABLE organization_subscriptions ENABLE ROW LEVEL SECURITY;
ALTER TABLE user_profiles ENABLE ROW LEVEL SECURITY;
ALTER TABLE testing_programs ENABLE ROW LEVEL SECURITY;
ALTER TABLE random_testing_pools ENABLE ROW LEVEL SECURITY;
ALTER TABLE pool_memberships ENABLE ROW LEVEL SECURITY;
ALTER TABLE random_selections ENABLE ROW LEVEL SECURITY;
ALTER TABLE random_selection_members ENABLE ROW LEVEL SECURITY;
ALTER TABLE drug_alcohol_tests ENABLE ROW LEVEL SECURITY;
ALTER TABLE mvr_reports ENABLE ROW LEVEL SECURITY;
ALTER TABLE mvr_violations ENABLE ROW LEVEL SECURITY;
ALTER TABLE dot_physicals ENABLE ROW LEVEL SECURITY;
ALTER TABLE background_checks ENABLE ROW LEVEL SECURITY;
ALTER TABLE background_check_findings ENABLE ROW LEVEL SECURITY;
ALTER TABLE documents ENABLE ROW LEVEL SECURITY;
ALTER TABLE temporal_workflows ENABLE ROW LEVEL SECURITY;
ALTER TABLE notifications ENABLE ROW LEVEL SECURITY;
ALTER TABLE audit_logs ENABLE ROW LEVEL SECURITY;
ALTER TABLE compliance_status ENABLE ROW LEVEL SECURITY;
ALTER TABLE saved_reports ENABLE ROW LEVEL SECURITY;
ALTER TABLE certificates ENABLE ROW LEVEL SECURITY;

-- Helper function to get current user's organization ID
CREATE OR REPLACE FUNCTION current_user_organization_id()
RETURNS UUID AS $$
  SELECT organization_id FROM user_profiles WHERE id = auth.uid();
$$ LANGUAGE SQL STABLE;

-- Helper function to check if current user has internal role
CREATE OR REPLACE FUNCTION is_internal_user()
RETURNS BOOLEAN AS $$
  SELECT role IN ('internal_su', 'internal_admin', 'internal_support') 
  FROM user_profiles WHERE id = auth.uid();
$$ LANGUAGE SQL STABLE;

-- Organizations RLS policies
CREATE POLICY "Users can view their own organization" ON organizations
    FOR SELECT USING (
        id = current_user_organization_id() OR 
        is_internal_user()
    );

CREATE POLICY "Internal users can manage all organizations" ON organizations
    FOR ALL USING (is_internal_user());

CREATE POLICY "Client admins can update their organization" ON organizations
    FOR UPDATE USING (
        id = current_user_organization_id() AND
        EXISTS (
            SELECT 1 FROM user_profiles 
            WHERE id = auth.uid() AND role = 'client_admin'
        )
    );

-- Subscription plans RLS policies (public read, internal write)
CREATE POLICY "Everyone can view subscription plans" ON subscription_plans
    FOR SELECT USING (is_active = true);

CREATE POLICY "Internal users can manage subscription plans" ON subscription_plans
    FOR ALL USING (is_internal_user());

-- Organization subscriptions RLS policies
CREATE POLICY "Users can view their organization's subscriptions" ON organization_subscriptions
    FOR SELECT USING (
        organization_id = current_user_organization_id() OR
        is_internal_user()
    );

CREATE POLICY "Internal users can manage all subscriptions" ON organization_subscriptions
    FOR ALL USING (is_internal_user());

-- User profiles RLS policies
CREATE POLICY "Users can view profiles in their organization" ON user_profiles
    FOR SELECT USING (
        organization_id = current_user_organization_id() OR
        is_internal_user()
    );

CREATE POLICY "Users can update their own profile" ON user_profiles
    FOR UPDATE USING (id = auth.uid());

CREATE POLICY "Admins can manage users in their organization" ON user_profiles
    FOR ALL USING (
        (organization_id = current_user_organization_id() AND
         EXISTS (
             SELECT 1 FROM user_profiles 
             WHERE id = auth.uid() AND role IN ('client_admin', 'der', 'hr_manager')
         )) OR
        is_internal_user()
    );

-- Testing programs RLS policies
CREATE POLICY "Users can view programs in their organization" ON testing_programs
    FOR SELECT USING (
        organization_id = current_user_organization_id() OR
        is_internal_user()
    );

CREATE POLICY "Authorized users can manage testing programs" ON testing_programs
    FOR ALL USING (
        (organization_id = current_user_organization_id() AND
         EXISTS (
             SELECT 1 FROM user_profiles 
             WHERE id = auth.uid() AND role IN ('client_admin', 'der', 'safety_manager')
         )) OR
        is_internal_user()
    );

-- Random testing pools RLS policies
CREATE POLICY "Users can view pools in their organization" ON random_testing_pools
    FOR SELECT USING (
        organization_id = current_user_organization_id() OR
        is_internal_user()
    );

CREATE POLICY "Authorized users can manage pools" ON random_testing_pools
    FOR ALL USING (
        (organization_id = current_user_organization_id() AND
         EXISTS (
             SELECT 1 FROM user_profiles 
             WHERE id = auth.uid() AND role IN ('client_admin', 'der', 'safety_manager')
         )) OR
        is_internal_user()
    );

-- Pool memberships RLS policies
CREATE POLICY "Users can view memberships in their organization" ON pool_memberships
    FOR SELECT USING (
        EXISTS (
            SELECT 1 FROM random_testing_pools p
            WHERE p.id = pool_id AND (
                p.organization_id = current_user_organization_id() OR
                is_internal_user()
            )
        )
    );

CREATE POLICY "Authorized users can manage pool memberships" ON pool_memberships
    FOR ALL USING (
        EXISTS (
            SELECT 1 FROM random_testing_pools p
            WHERE p.id = pool_id AND (
                (p.organization_id = current_user_organization_id() AND
                 EXISTS (
                     SELECT 1 FROM user_profiles 
                     WHERE id = auth.uid() AND role IN ('client_admin', 'der', 'safety_manager')
                 )) OR
                is_internal_user()
            )
        )
    );

-- Drug/alcohol tests RLS policies
CREATE POLICY "Users can view tests in their organization or their own tests" ON drug_alcohol_tests
    FOR SELECT USING (
        organization_id = current_user_organization_id() OR
        user_id = auth.uid() OR
        is_internal_user()
    );

CREATE POLICY "Authorized users can manage tests" ON drug_alcohol_tests
    FOR ALL USING (
        (organization_id = current_user_organization_id() AND
         EXISTS (
             SELECT 1 FROM user_profiles 
             WHERE id = auth.uid() AND role IN ('client_admin', 'der', 'safety_manager', 'mro')
         )) OR
        is_internal_user()
    );

-- MVR reports RLS policies
CREATE POLICY "Users can view MVR reports in their organization" ON mvr_reports
    FOR SELECT USING (
        organization_id = current_user_organization_id() OR
        user_id = auth.uid() OR
        is_internal_user()
    );

CREATE POLICY "Authorized users can manage MVR reports" ON mvr_reports
    FOR ALL USING (
        (organization_id = current_user_organization_id() AND
         EXISTS (
             SELECT 1 FROM user_profiles 
             WHERE id = auth.uid() AND role IN ('client_admin', 'der', 'safety_manager', 'hr_manager')
         )) OR
        is_internal_user()
    );

-- DOT physicals RLS policies
CREATE POLICY "Users can view physicals in their organization" ON dot_physicals
    FOR SELECT USING (
        organization_id = current_user_organization_id() OR
        user_id = auth.uid() OR
        examiner_id = auth.uid() OR
        is_internal_user()
    );

CREATE POLICY "Authorized users can manage physicals" ON dot_physicals
    FOR ALL USING (
        (organization_id = current_user_organization_id() AND
         EXISTS (
             SELECT 1 FROM user_profiles 
             WHERE id = auth.uid() AND role IN ('client_admin', 'der', 'safety_manager', 'hr_manager')
         )) OR
        examiner_id = auth.uid() OR
        is_internal_user()
    );

-- Background checks RLS policies
CREATE POLICY "Users can view background checks in their organization" ON background_checks
    FOR SELECT USING (
        organization_id = current_user_organization_id() OR
        is_internal_user()
    );

CREATE POLICY "Authorized users can manage background checks" ON background_checks
    FOR ALL USING (
        (organization_id = current_user_organization_id() AND
         EXISTS (
             SELECT 1 FROM user_profiles 
             WHERE id = auth.uid() AND role IN ('client_admin', 'hr_manager')
         )) OR
        is_internal_user()
    );

-- Documents RLS policies
CREATE POLICY "Users can view documents in their organization" ON documents
    FOR SELECT USING (
        organization_id = current_user_organization_id() OR
        user_id = auth.uid() OR
        is_internal_user()
    );

CREATE POLICY "Users can upload documents" ON documents
    FOR INSERT WITH CHECK (
        organization_id = current_user_organization_id() OR
        is_internal_user()
    );

-- Notifications RLS policies
CREATE POLICY "Users can view their own notifications" ON notifications
    FOR SELECT USING (
        user_id = auth.uid() OR
        (user_id IS NULL AND organization_id = current_user_organization_id()) OR
        is_internal_user()
    );

CREATE POLICY "Users can update their own notifications" ON notifications
    FOR UPDATE USING (user_id = auth.uid());

-- Audit logs RLS policies (read-only for most users)
CREATE POLICY "Internal users can view all audit logs" ON audit_logs
    FOR SELECT USING (is_internal_user());

CREATE POLICY "Admins can view their organization's audit logs" ON audit_logs
    FOR SELECT USING (
        organization_id = current_user_organization_id() AND
        EXISTS (
            SELECT 1 FROM user_profiles 
            WHERE id = auth.uid() AND role IN ('client_admin', 'der')
        )
    );

-- Compliance status RLS policies
CREATE POLICY "Users can view compliance status in their organization" ON compliance_status
    FOR SELECT USING (
        organization_id = current_user_organization_id() OR
        user_id = auth.uid() OR
        is_internal_user()
    );

-- Certificates RLS policies
CREATE POLICY "Users can view certificates in their organization" ON certificates
    FOR SELECT USING (
        organization_id = current_user_organization_id() OR
        user_id = auth.uid() OR
        is_internal_user()
    );

-- =============================================
-- INITIAL DATA SETUP
-- =============================================

-- Insert subscription plans
INSERT INTO subscription_plans (name, tier, description, max_employees, annual_price, features) VALUES
('Starter Plan', 'starter', 'Basic compliance management for small operations', 25, 2400.00, 
 '{"drug_testing": "basic", "mvr_monitoring": true, "dot_physicals": false, "support": "email"}'::jsonb),
('Professional Plan', 'professional', 'Full compliance program for medium operations', 100, 8400.00,
 '{"drug_testing": "full", "mvr_monitoring": true, "dot_physicals": true, "alcohol_testing": true, "support": "phone"}'::jsonb),
('Enterprise Plan', 'enterprise', 'Advanced compliance with custom reporting', 500, 24000.00,
 '{"drug_testing": "full", "mvr_monitoring": true, "dot_physicals": true, "background_checks": true, "custom_reports": true, "training": true, "support": "priority"}'::jsonb),
('Custom Plan', 'custom', 'Tailored solution for large organizations', NULL, NULL,
 '{"everything": true, "dedicated_support": true, "custom_integrations": true}'::jsonb);

-- Insert internal organization
INSERT INTO organizations (name, type, email, phone, is_active) VALUES
('V1 Consortium', 'internal', 'admin@v1consortium.com', '555-0100', true);

-- =============================================
-- UTILITY FUNCTIONS
-- =============================================

-- Function to calculate compliance percentage
CREATE OR REPLACE FUNCTION calculate_compliance_percentage(user_profile_id UUID)
RETURNS DECIMAL(5,2) AS $$
DECLARE
    total_requirements INTEGER := 0;
    met_requirements INTEGER := 0;
    compliance_pct DECIMAL(5,2);
    user_rec RECORD;
BEGIN
    -- Get user information
    SELECT * INTO user_rec FROM user_profiles WHERE id = user_profile_id;
    
    IF NOT FOUND THEN
        RETURN 0.00;
    END IF;
    
    -- Check drug testing requirements
    IF user_rec.requires_dot_testing OR user_rec.requires_non_dot_testing THEN
        total_requirements := total_requirements + 1;
        
        -- Check if user has current drug test
        IF EXISTS (
            SELECT 1 FROM drug_alcohol_tests 
            WHERE user_id = user_profile_id 
            AND test_category IN ('drug', 'drug_alcohol')
            AND result = 'negative'
            AND collection_date > CURRENT_DATE - INTERVAL '12 months'
        ) THEN
            met_requirements := met_requirements + 1;
        END IF;
    END IF;
    
    -- Check MVR requirements (for CDL holders)
    IF user_rec.cdl_number IS NOT NULL THEN
        total_requirements := total_requirements + 1;
        
        -- Check if user has current MVR
        IF EXISTS (
            SELECT 1 FROM mvr_reports 
            WHERE user_id = user_profile_id 
            AND status = 'reviewed'
            AND report_date > CURRENT_DATE - INTERVAL '12 months'
        ) THEN
            met_requirements := met_requirements + 1;
        END IF;
    END IF;
    
    -- Check DOT physical requirements
    IF user_rec.requires_dot_testing THEN
        total_requirements := total_requirements + 1;
        
        -- Check if user has current DOT physical
        IF EXISTS (
            SELECT 1 FROM dot_physicals 
            WHERE user_id = user_profile_id 
            AND status = 'completed'
            AND certificate_expiration_date > CURRENT_DATE
        ) THEN
            met_requirements := met_requirements + 1;
        END IF;
    END IF;
    
    -- Calculate percentage
    IF total_requirements = 0 THEN
        compliance_pct := 100.00;
    ELSE
        compliance_pct := (met_requirements::DECIMAL / total_requirements::DECIMAL) * 100.00;
    END IF;
    
    -- Update compliance status table
    INSERT INTO compliance_status (
        organization_id, user_id, is_compliant, compliance_percentage, last_updated
    ) VALUES (
        user_rec.organization_id, user_profile_id, 
        compliance_pct = 100.00, compliance_pct, NOW()
    )
    ON CONFLICT (user_id) DO UPDATE SET
        is_compliant = compliance_pct = 100.00,
        compliance_percentage = compliance_pct,
        last_updated = NOW();
    
    RETURN compliance_pct;
END;
$$ LANGUAGE plpgsql;