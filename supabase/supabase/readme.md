# V1 Consortium Database Schema Documentation

## Overview

This document explains the database design decisions and architecture for the V1 Consortium compliance and screening services platform. The schema is designed to support a comprehensive SaaS platform for transportation compliance management, drug testing, MVR monitoring, DOT physicals, and background checks.

## Architecture Decisions

### Technology Stack
- **Database**: PostgreSQL via Supabase
- **Authentication**: Supabase Auth with Row Level Security (RLS)
- **Storage**: Supabase Storage for document management
- **Extensions**: UUID generation (`uuid-ossp`) and case-insensitive text (`citext`)

### Design Principles

1. **Multi-tenant Architecture**: All data is isolated by organization using RLS policies
2. **HIPAA Compliance**: Sensitive medical data is properly secured and audited
3. **Regulatory Compliance**: Schema supports DOT Part 40, FMCSA, and state requirements
4. **Scalability**: Proper indexing and normalization for performance
5. **Audit Trail**: Comprehensive logging for compliance and debugging
6. **Flexibility**: JSONB fields for extensible configuration and metadata

## Core Entity Relationships

### Organizations (Multi-Tenant Base)
```
organizations (1) ──→ (M) user_profiles
organizations (1) ──→ (M) organization_subscriptions
organizations (1) ──→ (M) testing_programs
organizations (1) ──→ (M) all compliance records
```

**Key Design Decisions:**
- `organization_type` enum distinguishes internal, client, and provider organizations
- Support for both DOT and non-DOT operations via `is_dot_regulated` flag
- JSONB `settings` field allows flexible configuration without schema changes
- USDOT and MC numbers for regulatory identification

### User Management & Authentication
```
auth.users (Supabase) ──→ user_profiles (extended)
user_profiles (1) ──→ (M) pool_memberships
user_profiles (1) ──→ (M) drug_alcohol_tests
user_profiles (1) ──→ (M) mvr_reports
user_profiles (1) ──→ (M) dot_physicals
```

**Key Design Decisions:**
- Extends Supabase auth.users with detailed profile information
- Role-based permissions with granular user roles (11 distinct roles)
- Self-referencing supervisor relationship for organizational hierarchy
- CDL tracking with state-specific information
- Emergency contact information for safety compliance
- Supports both DOT and non-DOT testing requirements per user

### Subscription & Billing Management
```
subscription_plans (1) ──→ (M) organization_subscriptions
organization_subscriptions ──→ organizations
```

**Key Design Decisions:**
- Tiered subscription model (Starter, Professional, Enterprise, Custom)
- JSONB features field allows flexible plan configuration
- Annual subscription cycle (Jan-Dec) as per PRD requirements
- Stripe integration fields for payment processing
- Support for custom pricing negotiations

## Compliance Modules

### Drug & Alcohol Testing Program
```
testing_programs (1) ──→ (M) random_testing_pools
random_testing_pools (1) ──→ (M) pool_memberships (M) ──→ (1) user_profiles
random_selections (1) ──→ (M) random_selection_members
drug_alcohol_tests ──→ random_selections (for random tests)
drug_alcohol_tests ──→ user_profiles (subject)
```

**Key Design Decisions:**
- Separate configuration for DOT vs non-DOT testing programs
- Flexible drug panel types (5-panel, expanded, custom)
- Random testing pools with automated selection tracking
- Complete audit trail for random selections (FMCSA requirement)
- Support for all test types: pre-employment, random, post-accident, etc.
- MRO (Medical Review Officer) workflow integration
- External API integration fields for Quest Diagnostics and other providers

### Motor Vehicle Records (MVR)
```
user_profiles (1) ──→ (M) mvr_reports
mvr_reports (1) ──→ (M) mvr_violations
```

**Key Design Decisions:**
- State-specific license tracking
- Automated violation parsing and categorization
- Severity classification (minor, major, serious)
- CDL impact tracking for commercial drivers
- JSONB storage for raw provider data
- Compliance action tracking and notifications

### DOT Physical Examinations
```
user_profiles (1) ──→ (M) dot_physicals
user_profiles (medical_examiner) ──→ (M) dot_physicals (as examiner)
```

**Key Design Decisions:**
- Medical examiner relationship tracking
- Certificate management with expiration monitoring
- Medical qualification status and restrictions
- Exemption and monitoring requirement tracking
- Integration with certified medical examiner registry

### Background Checks
```
user_profiles (1) ──→ (M) background_checks
background_checks (1) ──→ (M) background_check_findings
```

**Key Design Decisions:**
- Multiple check types (criminal, employment, education, etc.)
- FCRA compliance workflow tracking
- Individualized assessment support for adverse actions
- Jurisdiction and case tracking for legal requirements
- Job-relatedness evaluation for compliance

## Document Management

### Secure Document Storage
```
documents ──→ organizations (multi-tenant)
documents ──→ user_profiles (optional, for personal docs)
documents ──→ [test_id, mvr_report_id, physical_id, background_check_id] (related entities)
```

**Key Design Decisions:**
- Supabase Storage integration for secure file storage
- Document versioning with parent-child relationships
- HIPAA protection flags for medical documents
- Retention policy automation with auto-delete dates
- Access tracking for audit compliance
- Confidentiality levels and download monitoring

## Workflow & Automation

### Temporal Workflow Integration
```
temporal_workflows ──→ organizations
temporal_workflows ──→ [user_id, test_id, selection_id] (related entities)
```

**Key Design Decisions:**
- Async workflow tracking for test ordering, result polling, random selections
- Retry logic with configurable maximum attempts
- Workflow status monitoring and error handling
- Input/output data storage for debugging and audit
- Integration with Temporal workflow engine

### Notification System
```
notifications ──→ organizations
notifications ──→ user_profiles (optional, for personal notifications)
notifications ──→ [test_id, mvr_report_id, physical_id] (related entities)
```

**Key Design Decisions:**
- Multi-channel delivery (email, SMS, in-app, phone)
- Priority levels for urgent compliance matters
- Delivery tracking and retry logic
- External provider integration (email/SMS services)
- Read receipt and engagement tracking

## Security & Compliance

### Row Level Security (RLS) Implementation
**Multi-tenant Data Isolation:**
- All tables implement organization-based RLS policies
- Helper functions for current user organization lookup
- Internal user access across all organizations
- Role-based access within organizations

**Key Security Features:**
- Supabase Auth integration with JWT tokens
- Role-based permissions (11 distinct user roles)
- HIPAA-compliant data access controls
- Comprehensive audit logging
- IP address and user agent tracking

### Audit & Compliance Tracking
```
audit_logs ──→ organizations (optional)
audit_logs ──→ user_profiles (optional)
compliance_status ──→ organizations
compliance_status ──→ user_profiles
```

**Key Design Decisions:**
- Immutable audit trail for all system actions
- HIPAA logging flags for medical data access
- Change tracking with old/new value comparison
- Real-time compliance percentage calculation
- Automated compliance status updates

## Performance Optimizations

### Indexing Strategy
- **Primary Access Patterns**: Organization-based queries (multi-tenant)
- **Temporal Queries**: Date-based indexes for compliance reporting
- **Status Filtering**: Enum-based indexes for workflow states
- **External Integration**: Unique indexes on external provider IDs
- **Compliance Monitoring**: Expiration date indexes for alerts

### Query Optimization
- Partial indexes on nullable foreign keys
- Composite indexes for common query patterns
- JSONB indexing for flexible metadata queries
- Efficient RLS policy implementation

## Data Types & Constraints

### Enum Types for Data Consistency
- **User Roles**: 11 distinct roles with clear permissions
- **Test Types**: All DOT-compliant test categories
- **Status Tracking**: Workflow states for each compliance module
- **Document Types**: Categorized for proper handling and retention

### Validation Constraints
- Email format validation using regex patterns
- SSN last-four digit validation
- Date range validation for subscriptions and certificates
- Referential integrity with cascading deletes where appropriate

## Utility Functions

### Compliance Calculation
- `calculate_compliance_percentage()`: Real-time compliance scoring
- Automatic compliance status updates
- Multi-factor compliance evaluation (testing, MVR, physicals)

### Helper Functions
- `auth.current_user_organization_id()`: Multi-tenant context
- `auth.is_internal_user()`: Role-based access control
- `update_updated_at_column()`: Automatic timestamp management

## Migration Considerations

### Initial Data Setup
- Subscription plans with feature matrices
- Internal organization for system administration
- Baseline compliance configurations

### Future Extensibility
- JSONB fields for configuration flexibility
- Versioned document management
- Temporal workflow integration points
- External API integration fields

## Compliance Requirements Addressed

### DOT/FMCSA Regulations
- 49 CFR Part 40 drug and alcohol testing procedures
- 49 CFR Part 382 FMCSA drug and alcohol regulations
- FMCSA Clearinghouse integration support
- Random testing audit trail requirements

### Data Protection
- HIPAA compliance for medical information
- FCRA compliance for background checks
- SOC 2 Type II security controls support
- Comprehensive audit logging

### Regulatory Reporting
- Random selection documentation
- Compliance certificate generation
- Violation tracking and reporting
- Annual compliance summaries

## Conclusion

This database schema provides a robust foundation for the V1 Consortium compliance platform, balancing regulatory requirements, security needs, and operational efficiency. The multi-tenant architecture with comprehensive RLS policies ensures data isolation while maintaining the flexibility needed for a diverse client base in the transportation and compliance industry.

The design emphasizes auditability, scalability, and regulatory compliance while providing the flexibility needed for future enhancements and integrations. The use of modern PostgreSQL features (JSONB, enums, partial indexes) ensures optimal performance and maintainability.