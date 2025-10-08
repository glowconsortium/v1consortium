# V1 Consortium - Compliance & Screening Services Platform

A comprehensive SaaS platform for transportation compliance management, drug testing, MVR monitoring, DOT physicals, and background checks.

## Technology Stack

- **Backend**: Go (GoFrame framework) + SvelteKit server routes
- **Frontend**: SvelteKit with TypeScript
- **Database**: PostgreSQL via Supabase with Row Level Security (RLS)
- **Storage**: Supabase Storage for HIPAA-compliant document management
- **Authentication**: Supabase Auth with multi-factor authentication
- **Workflows**: Temporal for async operations
- **Payments**: Stripe integration
- **API**: Protocol Buffers (gRPC) with REST gateway

## Quick Start
- https://goframe.org/quick

## Database Architecture

### Multi-Tenant Design
The platform uses a sophisticated multi-tenant architecture with Row Level Security (RLS) policies ensuring complete data isolation between organizations:

- **Organizations**: Base tenant entity with support for internal, client, and provider types
- **User Profiles**: Extended Supabase auth with 11 distinct roles and granular permissions
- **Compliance Modules**: Drug testing, MVR monitoring, DOT physicals, background checks
- **Document Management**: HIPAA-compliant storage with versioning and retention policies
- **Audit Trail**: Comprehensive logging for all system actions and data changes

### Key Database Features
- **HIPAA Compliance**: Sensitive medical data properly secured and audited
- **Regulatory Compliance**: Schema supports DOT Part 40, FMCSA, and state requirements
- **Performance**: Optimized indexing and query patterns for multi-tenant access
- **Flexibility**: JSONB fields for extensible configuration without schema changes

For detailed database documentation, see: [`supabase/supabase/readme.md`](supabase/supabase/readme.md)

## Protocol Buffers Architecture

### Service Organization
The protobuf definitions are organized into clear service boundaries supporting microservices architecture:

#### Authentication & Authorization (`v1/auth/`)
- **`auth.proto`**: Core authentication (login, register, MFA, password management)
- **`session.proto`**: Session and token lifecycle management
- **`authorization.proto`**: Role-based permissions with 11 user roles and granular permissions

#### Business Services (`v1/services/`)
- **`organization.proto`**: Multi-tenant organization and user management
- **`drug_testing.proto`**: DOT/non-DOT testing programs with random selection algorithms
- **`mvr.proto`**: Motor Vehicle Record monitoring with continuous compliance tracking
- **`dot_physical.proto`**: Medical examinations and certificate management
- **`background_check.proto`**: FCRA-compliant background screening with adverse action workflows
- **`compliance.proto`**: Real-time compliance status, certificate generation, and reporting
- **`notification.proto`**: Multi-channel communication with template management
- **`document.proto`**: HIPAA-compliant document storage with sharing and retention policies
- **`workflow.proto`**: Temporal workflow orchestration for async operations

#### API Gateway (`v1/gateway/`)
- **`gateway.proto`**: Core gateway functionality (auth, rate limiting, routing)
- **`config.proto`**: Dynamic route configuration and middleware management
- **`monitoring.proto`**: Metrics collection, performance monitoring, and error tracking

#### Data Entities (`v1/pbentity/`)
Auto-generated from database schema including all compliance entities:
- Organizations, User Profiles, Subscription Plans
- Testing Programs, Drug Tests, Random Pools & Selections
- MVR Reports & Violations, DOT Physicals, Background Checks
- Documents, Certificates, Notifications, Audit Logs
- Compliance Status, Temporal Workflows

### Key Features
- **Type Safety**: Strong typing across all service boundaries
- **Versioning**: Proper API versioning with backward compatibility
- **Documentation**: Self-documenting with comprehensive field descriptions
- **Code Generation**: Automated Go and TypeScript client generation

## Workflow Orchestration

### Temporal Integration
The platform leverages Temporal for reliable workflow orchestration:

#### Core Workflows
1. **Drug Test Ordering**: Async test ordering with external lab APIs (Quest Diagnostics, LabCorp)
2. **Result Processing**: Automated result polling and MRO review workflows
3. **Random Selection**: FMCSA-compliant random testing with audit trails
4. **MVR Monitoring**: Continuous driving record monitoring with violation alerts
5. **Background Checks**: Multi-provider screening with FCRA compliance workflows
6. **Notification Delivery**: Multi-channel communication with retry logic
7. **Compliance Monitoring**: Real-time compliance calculations and alerting

#### Workflow Features
- **Reliability**: Automatic retries with exponential backoff
- **Monitoring**: Real-time workflow status and performance metrics
- **Audit Trail**: Complete execution history for compliance requirements
- **Error Handling**: Comprehensive error recovery and notification
- **Scaling**: Horizontal scaling for high-volume operations

For detailed workflow documentation, see: [`workflows.md`](workflows.md)

## Compliance & Security

### Regulatory Compliance
- **DOT Part 40**: Drug and alcohol testing procedures
- **FMCSA**: Motor carrier safety regulations and random testing requirements
- **HIPAA**: Medical information protection and access controls
- **FCRA**: Fair Credit Reporting Act compliance for background checks
- **SOC 2 Type II**: Security controls and audit compliance

### Security Features
- **Multi-Factor Authentication**: TOTP, SMS, and email verification
- **Role-Based Access Control**: 11 distinct user roles with granular permissions
- **Data Encryption**: End-to-end encryption in transit and at rest
- **Audit Logging**: Comprehensive tracking of all system actions
- **IP Filtering**: Configurable access controls and geo-restrictions

## API Documentation

### REST API
- Base URL: `/api/v1`
- Authentication: JWT tokens with refresh rotation
- Rate Limiting: Configurable per-user and per-organization limits
- Documentation: Auto-generated OpenAPI/Swagger specs

### gRPC Services
- Protocol: HTTP/2 with TLS
- Load Balancing: Round-robin with health checks
- Circuit Breakers: Automatic failure detection and recovery
- Monitoring: Detailed metrics and distributed tracing

## Development

### Prerequisites
- Go 1.21+
- Node.js 18+
- PostgreSQL 15+
- Protocol Buffers compiler

### Setup
```bash
# Clone repository
git clone <repository-url>
cd v1consortium

# Install dependencies
go mod download
npm install

# Generate protobuf code
make pb

# Generate database entities
make dao

# Start development server
make run
```

### Project Structure
```
├── api/                    # Generated protobuf Go code
├── internal/               # Go business logic
├── manifest/protobuf/      # Protocol buffer definitions
├── supabase/              # Database schema and migrations
├── web-ui/                # SvelteKit frontend
├── docs/                  # Documentation
└── workflows.md           # Detailed workflow documentation
```