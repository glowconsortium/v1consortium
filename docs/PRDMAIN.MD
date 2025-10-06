# V1 Consortium - Compliance & Screening Services Platform
## Product Requirements Document

**Version:** 1.1  
**Date:** October 6, 2025  
**Author:** V1 Consortium Team

## 1. Executive Summary

This Product Requirements Document (PRD) outlines the functional and non-functional requirements for V1 Consortium, a comprehensive compliance and screening services SaaS platform. The system enables motor carriers, transportation companies, healthcare facilities, construction companies, and other organizations to manage drug testing, Motor Vehicle Record (MVR) checks, DOT physicals, random testing programs, and regulatory compliance documentation.

The platform provides:
* **Public-Facing Site** for marketing, signup, pricing, and subscription purchase
* **Carrier Portal** for managing drivers, ordering compliance services, conducting random testing, and downloading DOT/FMCSA-compliant certificates
* **Complete compliance management system** with automated workflows and third-party integrations

## 2. Goals & Objectives

* Provide carriers with a **complete compliance management system**
* Automate **drug tests, DOT physicals, MVR checks, and random test selection**
* Allow carriers to **download required DOT/FMCSA certificates and compliance docs**
* Support **annual subscription (Jan–Dec cycle)** with renewal enforcement
* Provide **integrations with third-party compliance/testing APIs**
* Ensure regulatory compliance across DOT and non-DOT requirements
* Deliver intuitive user experience for all stakeholder types

## 3. Target Users & User Classes

### 3.1 Primary Users
* **Motor carriers, owner-operators enrolled in consortiums**
* **Transportation companies requiring DOT compliance**
* **Healthcare facilities, construction companies** (non-DOT compliance)

### 3.2 Secondary Users
* **Compliance officers, safety managers**
* **HR managers and designated employer representatives**

### 3.3 User Roles and Permissions

#### 3.3.1 Internal Users
- **Internal Super User (internal_su)**: System administrators with full access to all features and client organizations
- **Internal Administrator (internal_admin)**: Staff managing client accounts and compliance programs
- **Internal Support (internal_support)**: Customer service representatives providing client support

#### 3.3.2 External Users
- **Client Administrator (client_admin)**: Organization administrators managing their compliance programs
- **DER (Designated Employer Representative)**: Manages drug testing programs and compliance
- **Safety Manager (safety_manager)**: Oversees safety programs and compliance reporting
- **HR Manager (hr_manager)**: Manages employee records and screening requirements
- **Employee (employee)**: Individual subject to testing and compliance requirements

#### 3.3.3 Provider Users
- **MRO (Medical Review Officer)**: Reviews and validates drug test results
- **SAP (Substance Abuse Professional)**: Provides evaluation and treatment recommendations
- **Medical Examiner (medical_examiner)**: Conducts DOT physical examinations

## 4. Core Features

### 4.1 Authentication & Subscription Management

#### 4.1.1 Authentication
* Supabase-based authentication (email/password, OAuth options)
* Multi-factor authentication (MFA)
* Role-based access control
* Secure session management

#### 4.1.2 Subscription Tiers
- **Starter**
  - Up to 25 employees
  - Basic drug testing program
  - MVR monitoring
  - Basic reporting
  - Email support

- **Professional**
  - Up to 100 employees
  - Full drug testing program
  - MVR monitoring
  - DOT physical tracking
  - Advanced reporting
  - Phone support

- **Enterprise**
  - Up to 500 employees
  - All compliance services
  - Custom reporting
  - Priority support
  - Training modules

- **Custom**
  - Unlimited employees
  - Full service management
  - Dedicated account manager
  - Custom integrations
  - On-site support

#### 4.1.3 Billing & Payments
* Annual subscription model (Jan–Dec cycle)
* Stripe integration for billing + receipts
* Grace period for renewals
* Automated renewal enforcement
* Custom pricing negotiations

### 4.2 Organization & User Management

#### 4.2.1 Organization Types
- **Internal**: V1 Consortium administrative organization
- **Client**: Customer organizations using compliance services
- **Provider**: Testing facilities, labs, and medical examiners

#### 4.2.2 Organization Profile Management
- Company information and industry classification
- DOT/Non-DOT status designation
- USDOT and MC number tracking
- Contact information and addresses
- Compliance requirements configuration
- Subscription and status management

#### 4.2.3 User Management
- User registration and approval workflow
- Role assignment and permissions
- Profile management and certification tracking
- Account activation/deactivation
- Password reset and security management

### 4.3 Drug & Alcohol Testing Program Management

#### 4.3.1 Testing Types
* Pre-employment testing
* Random testing
* Post-accident testing
* Reasonable suspicion testing
* Return-to-duty testing
* Follow-up testing

#### 4.3.2 Testing Programs
* DOT drug testing (5-panel)
* DOT alcohol testing
* Non-DOT drug testing (customizable panels)
* Non-DOT alcohol testing
* Expanded drug panels

#### 4.3.3 **Random Testing Management** (Enhanced)
* Create/Manage random testing pools (all drivers enrolled)
* Automated random selection of drivers (quarterly/monthly basis)
* Random selection algorithms with pool balancing
* Notifications to carriers & selected drivers
* Audit logs of selections to meet FMCSA requirements
* Testing frequency configuration and optimization

#### 4.3.4 Third-Party Integration
* Order tests through third-party APIs (Quest Diagnostics, etc.)
* Retrieve results automatically via API polling
* Temporal workflows for async test ordering and result processing
* Laboratory information system integration
* Real-time data synchronization

### 4.4 Motor Vehicle Record (MVR) Management

#### 4.4.1 MVR Services
* Initial MVR pulls for new employees
* Continuous monitoring and automated ordering
* Violation detection and alerts
* Driver qualification file management
* State-specific requirement compliance

#### 4.4.2 MVR Integration
* Integration with state DMV systems
* Automated violation flagging
* Driver disqualification tracking
* Compliance status updates

### 4.5 DOT Physical Management

#### 4.5.1 Physical Examination Services
* DOT medical examination scheduling
* Medical certificate tracking and storage
* Exemption management
* Medical review processes

#### 4.5.2 DOT Physical Functions
* Schedule examinations with certified examiners
* Track medical certificates and expiration dates
* Generate renewal reminders
* Manage medical exemptions
* Coordinate with certified medical examiners

### 4.6 **Reporting & Certificate Management** (Enhanced)

#### 4.6.1 Compliance Dashboard
* Real-time compliance status (per carrier + per driver)
* Violation tracking and alerts
* Regulatory change notifications
* Audit preparation tools

#### 4.6.2 **Downloadable Certificates & Documents**
* **Consortium Membership Certificate** (proof of enrollment)
* **Test Completion Certificates** (per driver/test)
* **Random Selection Reports** (FMCSA audit requirement)
* **DOT Annual Reports** and regulatory filings
* **Compliance History Reports** in PDF/CSV format
* **Driver Qualification Files**
* **Medical Certificates and Documentation**

#### 4.6.3 Report Types
* Drug testing summary reports
* Random testing audit reports
* MVR violation reports
* Compliance status reports
* Custom analytics dashboards
* Regulatory filing assistance

### 4.7 Background Check Services

#### 4.7.1 Background Check Types
* Criminal history checks
* Employment verification
* Reference checks
* Education verification
* Professional license verification
* Credit checks (where permitted)

#### 4.7.2 Background Check Management
* Order and track background checks
* Receive and review results
* Manage adverse action processes
* Store and organize reports
* Generate compliance documentation

### 4.8 Document Management & Storage

#### 4.8.1 Document Types
* Test results and lab reports
* MVR reports and violation history
* DOT medical certificates
* Training certificates and records
* Compliance documentation
* Policy documents and procedures

#### 4.8.2 Document Management Functions
* Secure upload and storage via Supabase Storage
* Organize by category and employee
* HIPAA-compliant document access
* Retention policy management
* Document search and retrieval
* Audit trail maintenance

### 4.9 Communication & Notifications

#### 4.9.1 Notification System
* Email notifications for testing requirements
* SMS alerts for urgent compliance matters
* In-app messaging and alerts
* Automated renewal reminders
* Regulatory update notifications

#### 4.9.2 Communication Management
* Notification preference configuration
* Escalation procedures
* Multi-channel communication support
* Template management for standard communications

### 4.10 Mobile Interface & Field Operations

#### 4.10.1 Mobile Optimization
* Responsive design for all device types
* Progressive Web App capabilities
* Mobile-optimized workflows
* Offline capability for critical functions

#### 4.10.2 Field Operation Support
* GPS location services for testing facilities
* Camera integration for document capture
* Mobile employee portal access
* Real-time status updates

## 5. Technical Architecture

### 5.1 Technology Stack

#### 5.1.1 Frontend
* **SvelteKit** with TypeScript
* Responsive design framework
* Progressive Web App capabilities
* Modern UI/UX principles
* Accessibility compliance (WCAG 2.1)

#### 5.1.2 Backend
* **SvelteKit server routes** with TypeScript
* **Go (GoFrame framework)** for complex business logic
* **Temporal workflows** for async operations (test ordering, result polling, random selection scheduling)
* RESTful API architecture
* Microservices design patterns

#### 5.1.3 Database & Storage
* **PostgreSQL via Supabase**
* HIPAA-compliant data storage
* **Supabase Storage** for document management
* Efficient indexing and query optimization
* Automated backup and recovery

#### 5.1.4 Authentication & Security
* **Supabase Auth** with multi-factor authentication
* JWT-based authentication
* Role-based access control (RBAC)
* Session management and security
* Comprehensive audit logging

### 5.2 Third-Party Integrations

#### 5.2.1 Testing & Compliance APIs
* **Quest Diagnostics API** for drug testing
* **MVR provider APIs** for driving record monitoring
* **FMCSA Clearinghouse API** for violation reporting
* **DOT physical provider APIs** for medical examinations

#### 5.2.2 Business Integrations
* **Stripe** for subscription billing and payment processing
* HR management system APIs
* Payroll system integrations
* Document generation libraries (PDF creation)

### 5.3 Workflow Management

#### 5.3.1 Temporal Workflows
* **Test ordering workflows** with retry logic
* **Result polling and processing** automation
* **Random selection scheduling** and execution
* **Notification and reminder** workflows
* **Compliance monitoring** and alerting

## 6. User Flows

### 6.1 Carrier Onboarding Flow (Enhanced)
1. **Signup & Subscription**: Sign up and select/pay for subscription tier
2. **Organization Setup**: Complete organization profile and compliance configuration
3. **Driver Enrollment**: Add drivers to consortium and compliance programs
4. **Service Activation**: Order initial tests (drug/DOT/MVR) and configure random testing
5. **Ongoing Operations**: System conducts automated random driver selection (quarterly)
6. **Certificate Management**: Download consortium certificate and compliance documents
7. **Continuous Monitoring**: Receive alerts, manage violations, maintain compliance

### 6.2 Drug Testing Process Flow
1. **Test Initiation**: System generates random selection or scheduled test requirement
2. **Notification**: Employee receives testing notification via preferred communication method
3. **Test Coordination**: Employee reports to designated testing facility
4. **Sample Collection**: Test conducted and samples sent to laboratory
5. **Result Processing**: MRO reviews results if necessary, validates findings
6. **Result Reporting**: Results reported to DER with appropriate documentation
7. **Follow-up Actions**: Return-to-duty or violation procedures initiated as needed

### 6.3 Random Testing Management Flow
1. **Pool Configuration**: Define random testing pools and frequency requirements
2. **Automated Selection**: Temporal workflow executes random selection algorithm
3. **Notification Distribution**: Selected employees and managers receive notifications
4. **Testing Coordination**: Schedule and track completion of required tests
5. **Audit Documentation**: Generate and store selection reports for compliance
6. **Compliance Reporting**: Provide quarterly and annual random testing reports

### 6.4 Compliance Certificate Generation Flow
1. **Status Verification**: System validates current compliance status
2. **Document Generation**: Create certificates using predefined templates
3. **Digital Signature**: Apply secure digital signatures and timestamps
4. **Storage & Access**: Store certificates in document management system
5. **Download & Distribution**: Provide secure download links to authorized users
6. **Audit Trail**: Maintain logs of certificate generation and access

## 7. Feature Entitlement Matrix

| Feature | Starter | Professional | Enterprise | Custom |
|---------|---------|-------------|------------|--------|
| Employee Limit | 25 | 100 | 500 | Unlimited |
| Drug Testing | Basic | Full Program | Full Program | Full Program |
| Alcohol Testing | ✗ | ✓ | ✓ | ✓ |
| MVR Monitoring | ✓ | ✓ | ✓ | ✓ |
| DOT Physicals | ✗ | ✓ | ✓ | ✓ |
| Background Checks | ✗ | ✗ | ✓ | ✓ |
| Random Pool Management | ✓ | ✓ | ✓ | ✓ |
| Certificate Downloads | Basic | ✓ | ✓ | ✓ |
| Basic Reporting | ✓ | ✓ | ✓ | ✓ |
| Advanced Analytics | ✗ | ✓ | ✓ | ✓ |
| Custom Reports | ✗ | ✗ | ✓ | ✓ |
| Training Modules | ✗ | ✗ | ✓ | ✓ |
| API Access | ✗ | ✗ | ✓ | ✓ |
| Temporal Workflows | ✗ | ✓ | ✓ | ✓ |
| Dedicated Support | ✗ | ✗ | ✗ | ✓ |

## 8. Non-Functional Requirements

### 8.1 Security & Compliance
* **HIPAA compliance** for medical information handling
* **SOC 2 Type II compliance** for security controls
* **Data encryption** in transit and at rest
* **Role-based access control** with principle of least privilege
* **Comprehensive audit trail** for all system actions
* **Multi-factor authentication** for enhanced security
* **Regular security assessments** and penetration testing

### 8.2 Performance & Reliability
* **Page load time** < 2 seconds for optimal user experience
* **Concurrent user support** for 5000+ users
* **Database response time** < 500ms for queries
* **System uptime** 99.9% availability SLA
* **Automated backup and recovery** procedures
* **Load balancing** and horizontal scaling capabilities

### 8.3 Regulatory Compliance
* **DOT Part 40 compliance** for drug and alcohol testing
* **FMCSA regulations adherence** for motor carriers
* **State-specific requirements** compliance
* **Industry standards compliance** across sectors
* **Regular regulatory updates** and system adaptations

### 8.4 Usability & Accessibility
* **Intuitive user interface** design
* **WCAG 2.1 accessibility compliance**
* **Mobile-responsive design** for all devices
* **Multi-language support** capabilities
* **Comprehensive user training resources**

## 9. MVP Scope & Implementation Phases

### 9.1 Phase 1 - MVP (Minimum Viable Product)
* **Core Authentication**: Supabase auth with basic role management
* **Subscription Management**: Stripe integration with basic tiers
* **Organization Setup**: Basic carrier dashboard and driver management
* **Drug Testing**: Order tests via Quest API and view results
* **Basic Reporting**: Generate simple compliance status reports
* **Certificate Downloads**: Basic consortium membership certificate
* **Random Testing**: Basic random selection system

### 9.2 Phase 2 - Enhanced Features
* **Advanced Random Testing**: Temporal workflow integration
* **MVR Monitoring**: Automated MVR pulls and violation tracking
* **DOT Physicals**: Scheduling and certificate management
* **Advanced Certificates**: Full compliance document generation
* **Mobile Optimization**: Responsive design enhancements

### 9.3 Phase 3 - Enterprise Features
* **Background Checks**: Full background screening capabilities
* **Advanced Analytics**: Custom dashboards and reporting
* **API Development**: Third-party integration capabilities
* **Training Modules**: Compliance education platform
* **Advanced Workflows**: Complex Temporal workflow implementations

## 10. Future Enhancements

### 10.1 Technology Enhancements
* **AI-powered compliance predictions** and risk assessment
* **Blockchain integration** for document verification
* **IoT device integration** for workplace monitoring
* **Machine learning** for violation pattern detection
* **Advanced analytics** with predictive modeling

### 10.2 Feature Expansions
* **Native mobile applications** for iOS/Android
* **Telemedicine integration** for remote DOT physicals
* **Wearable device monitoring** for continuous compliance
* **International compliance standards** support
* **Advanced automation** and workflow optimization

### 10.3 Integration Expansions
* **Multi-vendor test support** beyond Quest Diagnostics
* **Enhanced HR system integrations**
* **Payroll system connectivity**
* **Fleet management system APIs**
* **Insurance carrier integrations**

## 11. System Architecture Diagrams

### 11.1 High-Level Architecture
```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   SvelteKit     │    │   Go/GoFrame    │    │   Temporal      │
│   Frontend      │◄──►│   Backend API   │◄──►│   Workflows     │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                        │                        │
         ▼                        ▼                        ▼
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Supabase      │    │   Third-Party   │    │   Document      │
│   (DB + Auth)   │    │   APIs          │    │   Storage       │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

### 11.2 Data Flow Architecture
* **User Interface** → SvelteKit frontend with TypeScript
* **API Layer** → Go/GoFrame backend for business logic
* **Workflow Engine** → Temporal for async operations
* **Data Storage** → Supabase PostgreSQL database
* **Document Storage** → Supabase Storage with encryption
* **External APIs** → Quest Diagnostics, MVR providers, FMCSA

## 12. Conclusion

V1 Consortium represents a comprehensive solution for transportation and workplace compliance management, combining modern web technologies with robust business workflows. The platform addresses critical needs in the compliance industry while providing scalable, secure, and user-friendly tools for organizations of all sizes.

The integration of SvelteKit, Go/GoFrame, Temporal workflows, and Supabase provides a solid foundation for current requirements while enabling future growth and feature expansion. The emphasis on regulatory compliance, security, and user experience ensures the platform will meet the evolving needs of the compliance industry.

---

## Appendices

### A.1 Glossary
- **TPA**: Third Party Administrator - manages compliance programs for multiple employers
- **DER**: Designated Employer Representative - manages drug and alcohol programs
- **MRO**: Medical Review Officer - physician who reviews drug test results
- **SAP**: Substance Abuse Professional - evaluates employees who violate drug policies
- **DOT**: Department of Transportation - federal agency regulating transportation safety
- **MVR**: Motor Vehicle Record - driving history report
- **Clearinghouse**: FMCSA database of drug and alcohol violations
- **FMCSA**: Federal Motor Carrier Safety Administration
- **CFR**: Code of Federal Regulations

### A.2 Regulatory References
- 49 CFR Part 40 - DOT drug and alcohol testing procedures
- 49 CFR Part 382 - FMCSA drug and alcohol regulations
- HIPAA Privacy Rule
- Fair Credit Reporting Act (FCRA)
- State-specific transportation regulations

### A.3 Technical References
- SvelteKit Documentation
- GoFrame Framework Documentation
- Temporal Workflow Documentation
- Supabase Platform Documentation
- Quest Diagnostics API Documentation
