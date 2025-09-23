# V1 Consortium - Compliance & Screening Services Platform
## Product Requirements Document

**Version:** 1.0  
**Date:** September 23, 2025  
**Author:** Claude

## 1. Introduction

### 1.1 Purpose
This Product Requirements Document (PRD) outlines the functional and non-functional requirements for V1 Consortium, a comprehensive compliance and screening services SaaS platform built with SvelteKit, TypeScript, and Supabase. The system enables organizations to manage drug testing, Motor Vehicle Record (MVR) checks, DOT physicals, and other compliance screening services for both DOT and non-DOT clients.

### 1.2 Scope
V1 Consortium will provide a complete solution for transportation companies, healthcare facilities, construction companies, and other organizations requiring employee screening and compliance management. The platform will manage drug testing programs, MVR monitoring, DOT physical scheduling, background checks, and compliance reporting with various user roles and subscription tiers.

### 1.3 Definitions, Acronyms, and Abbreviations
- **DOT** - Department of Transportation
- **MVR** - Motor Vehicle Record
- **DER** - Designated Employer Representative
- **MRO** - Medical Review Officer
- **SAP** - Substance Abuse Professional
- **TPA** - Third Party Administrator
- **C/TPA** - Consortium/Third Party Administrator
- **FMCSA** - Federal Motor Carrier Safety Administration
- **CFR** - Code of Federal Regulations
- **DOT Physical** - Department of Transportation Medical Examination
- **Random Pool** - Group of employees subject to random drug testing

## 2. Product Overview

### 2.1 Product Perspective
V1 Consortium is a new web-based SaaS application designed to serve as a comprehensive Third Party Administrator (TPA) for compliance and screening services. It operates as a standalone system with integrations to testing laboratories, MVR providers, medical examination facilities, and regulatory reporting systems.

### 2.2 Product Features
- Consortium management for DOT and non-DOT clients
- Drug and alcohol testing program administration
- MVR monitoring and reporting
- DOT physical scheduling and tracking
- Background check coordination
- Compliance reporting and dashboard
- Client and employee management
- Violation tracking and follow-up
- Random testing pool management
- Educational resources and training
- Mobile-optimized interface for field operations

### 2.3 User Classes and Characteristics

#### 2.3.1 Internal Users
- **Internal Super User (internal_su)**: System administrators with full access to all features and client organizations
- **Internal Administrator (internal_admin)**: Staff managing client accounts and compliance programs
- **Internal Support (internal_support)**: Customer service representatives providing client support

#### 2.3.2 External Users
- **Client Administrator (client_admin)**: Organization administrators managing their compliance programs
- **DER (Designated Employer Representative)**: Manages drug testing programs and compliance
- **Safety Manager (safety_manager)**: Oversees safety programs and compliance reporting
- **HR Manager (hr_manager)**: Manages employee records and screening requirements
- **Employee (employee)**: Individual subject to testing and compliance requirements

#### 2.3.3 Provider Users
- **MRO (Medical Review Officer)**: Reviews and validates drug test results
- **SAP (Substance Abuse Professional)**: Provides evaluation and treatment recommendations
- **Medical Examiner (medical_examiner)**: Conducts DOT physical examinations

### 2.4 Operating Environment
- Web browsers (Chrome, Firefox, Safari, Edge)
- Mobile devices (responsive design with mobile-optimized views)
- Backend: SvelteKit server with Supabase
- Database: PostgreSQL via Supabase

### 2.5 Design and Implementation Constraints
- Built with SvelteKit and TypeScript
- Utilizes Supabase for database, authentication, and storage
- HIPAA compliant design for medical information
- DOT regulatory compliance requirements
- Responsive design to support desktop and mobile devices
- RESTful API architecture

### 2.6 Assumptions and Dependencies
- Reliable internet connectivity for users
- Supabase services availability
- Third-party integrations for testing labs and MVR providers
- Compliance with federal and state regulations

## 3. Specific Requirements

### 3.1 Organization Management

#### 3.1.1 Organization Types
- **Internal**: V1 Consortium administrative organization
- **Client**: Customer organizations using compliance services
- **Provider**: Testing facilities, labs, and medical examiners

#### 3.1.2 Organization Profile
- Company name
- Type (internal, client, provider)
- Industry classification
- DOT/Non-DOT status
- Contact information
- Physical address
- USDOT number (if applicable)
- MC number (if applicable)
- Subscription details
- Compliance requirements
- Created date
- Status (active, inactive, trial)

#### 3.1.3 Organization Management Functions
- CRUD operations for organizations
- Activate/deactivate organizations
- Manage subscription status
- Assign compliance programs
- Monitor regulatory requirements
- Generate compliance certificates

### 3.2 User Management

#### 3.2.1 User Profile
- Username
- Email
- Password (securely stored)
- Full name
- Title/Role
- Contact information
- Organization affiliation
- User role
- Certifications
- Status (active, inactive)
- Profile picture
- Created date
- Last login date

#### 3.2.2 User Roles and Permissions
- **Internal Super User (internal_su)**
  - Manage all client organizations
  - Configure system settings
  - Generate reports across all clients
  - Manage subscription tiers
  - Access all compliance data

- **Internal Administrator (internal_admin)**
  - Manage assigned client accounts
  - Configure compliance programs
  - Generate client reports
  - Provide client support

- **Internal Support (internal_support)**
  - View client information
  - Provide customer support
  - Generate basic reports
  - Limited administrative capabilities

- **Client Administrator (client_admin)**
  - Manage organization profile
  - Invite and manage users
  - Configure compliance programs
  - Access organization reports
  - Manage employee roster

- **DER (Designated Employer Representative)**
  - Manage drug testing programs
  - Review test results
  - Handle compliance violations
  - Generate testing reports
  - Coordinate with MRO

- **Safety Manager (safety_manager)**
  - Monitor compliance status
  - Generate safety reports
  - Manage training programs
  - Track violations and corrective actions

- **HR Manager (hr_manager)**
  - Manage employee records
  - Schedule testing and physicals
  - Track compliance requirements
  - Generate HR reports

- **Employee (employee)**
  - View personal compliance status
  - Schedule appointments
  - Access training materials
  - Update personal information

#### 3.2.3 User Management Functions
- Registration and approval workflow
- Role assignment and permissions
- Password reset and security
- Account activation/deactivation
- User profile management
- Training and certification tracking

### 3.3 Subscription Management

#### 3.3.1 Subscription Tiers
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

#### 3.3.2 Subscription Management Functions
- Subscribe/upgrade/downgrade plans
- Track usage against limits
- Process payments
- Generate invoices
- Manage renewal cycles
- Custom pricing negotiations

### 3.4 Drug Testing Program Management

#### 3.4.1 Testing Types
- Pre-employment testing
- Random testing
- Post-accident testing
- Reasonable suspicion testing
- Return-to-duty testing
- Follow-up testing

#### 3.4.2 Testing Programs
- DOT drug testing (5-panel)
- DOT alcohol testing
- Non-DOT drug testing (customizable panels)
- Non-DOT alcohol testing
- Expanded drug panels

#### 3.4.3 Random Pool Management
- Employee pool creation and management
- Random selection algorithms
- Testing frequency configuration
- Pool balancing and optimization
- Quarterly reporting

#### 3.4.4 Drug Testing Functions
- Schedule and coordinate testing
- Manage testing locations
- Track test results
- MRO review coordination
- Violation reporting
- Return-to-duty process management
- Generate DOT reporting forms

### 3.5 MVR Monitoring

#### 3.5.1 MVR Services
- Initial MVR pulls
- Continuous monitoring
- Violation alerts
- Driver qualification files
- State-specific requirements

#### 3.5.2 MVR Management Functions
- Automated MVR ordering
- Violation detection and alerts
- Driver disqualification tracking
- Compliance reporting
- Integration with state DMV systems

### 3.6 DOT Physical Management

#### 3.6.1 Physical Examination Services
- DOT medical examinations
- Medical certificate tracking
- Exemption management
- Medical review processes

#### 3.6.2 DOT Physical Functions
- Schedule examinations
- Track medical certificates
- Monitor expiration dates
- Generate renewal reminders
- Manage medical exemptions
- Coordinate with certified medical examiners

### 3.7 Background Check Services

#### 3.7.1 Background Check Types
- Criminal history checks
- Employment verification
- Reference checks
- Education verification
- Professional license verification
- Credit checks (where permitted)

#### 3.7.2 Background Check Functions
- Order and track background checks
- Receive and review results
- Manage adverse action processes
- Store and organize reports
- Generate compliance documentation

### 3.8 Compliance Management

#### 3.8.1 Compliance Programs
- DOT compliance programs
- Non-DOT workplace policies
- Industry-specific requirements
- State and local regulations

#### 3.8.2 Compliance Functions
- Track compliance status
- Generate violation reports
- Manage corrective actions
- Monitor regulatory changes
- Provide compliance alerts
- Generate audit reports

### 3.9 Reporting and Analytics

#### 3.9.1 Report Types
- Drug testing summary reports
- Random testing reports
- MVR violation reports
- Compliance status reports
- DOT annual reports
- Custom analytics dashboards

#### 3.9.2 Reporting Functions
- Generate predefined reports
- Create custom reports
- Export reports (PDF, Excel)
- Schedule automated reports
- Dashboard visualizations
- Regulatory filing assistance

### 3.10 Document Management

#### 3.10.1 Document Types
- Test results and lab reports
- MVR reports
- DOT medical certificates
- Training certificates
- Compliance documentation
- Policy documents

#### 3.10.2 Document Management Functions
- Upload and store documents
- Organize by category and employee
- Secure document access
- Retention policy management
- Document search and retrieval
- Audit trail maintenance

### 3.11 Communication and Notifications

#### 3.11.1 Notification Types
- Testing notifications
- Compliance alerts
- Violation notifications
- Renewal reminders
- Regulatory updates

#### 3.11.2 Communication Functions
- Email notifications
- SMS alerts
- In-app messaging
- Notification preferences
- Escalation procedures
- Automated reminders

### 3.12 Mobile Interface

#### 3.12.1 Mobile Features
- Employee portal access
- Testing scheduling
- Compliance status viewing
- Document upload
- Notification management

#### 3.12.2 Mobile Functions
- Responsive design
- Mobile-optimized workflows
- Offline capability for critical functions
- GPS location services
- Camera integration for documents

### 3.13 Integration Capabilities

#### 3.13.1 Integration Types
- Laboratory information systems
- MVR provider APIs
- Medical examiner systems
- FMCSA Clearinghouse
- HR management systems
- Payroll systems

#### 3.13.2 Integration Functions
- Real-time data synchronization
- Automated result importing
- Regulatory reporting automation
- Third-party system connectivity
- API management

## 4. Feature Entitlement Matrix

| Feature | Starter | Professional | Enterprise | Custom |
|---------|---------|-------------|------------|--------|
| Employee Limit | 25 | 100 | 500 | Unlimited |
| Drug Testing | Basic | Full Program | Full Program | Full Program |
| Alcohol Testing | ✗ | ✓ | ✓ | ✓ |
| MVR Monitoring | ✓ | ✓ | ✓ | ✓ |
| DOT Physicals | ✗ | ✓ | ✓ | ✓ |
| Background Checks | ✗ | ✗ | ✓ | ✓ |
| Random Pool Management | ✓ | ✓ | ✓ | ✓ |
| Basic Reporting | ✓ | ✓ | ✓ | ✓ |
| Advanced Analytics | ✗ | ✓ | ✓ | ✓ |
| Custom Reports | ✗ | ✗ | ✓ | ✓ |
| Training Modules | ✗ | ✗ | ✓ | ✓ |
| API Access | ✗ | ✗ | ✓ | ✓ |
| Dedicated Support | ✗ | ✗ | ✗ | ✓ |

## 5. Non-Functional Requirements

### 5.1 Security and Compliance
- HIPAA compliance for medical information
- SOC 2 Type II compliance
- Data encryption in transit and at rest
- Role-based access control
- Audit trail for all actions
- Secure authentication (MFA)
- Regular security assessments

### 5.2 Performance
- Page load time < 2 seconds
- Support for 5000+ concurrent users
- Database response time < 500ms
- 99.9% system uptime
- Automated backup and recovery

### 5.3 Regulatory Compliance
- DOT Part 40 compliance
- FMCSA regulations adherence
- State-specific requirements
- Industry standards compliance
- Regular regulatory updates

### 5.4 Usability
- Intuitive user interface
- Accessibility compliance (WCAG 2.1)
- Mobile-responsive design
- Multi-language support
- User training resources

### 5.5 Scalability
- Horizontal scaling capabilities
- Database optimization
- Efficient resource utilization
- Load balancing
- Caching strategies

## 6. User Flows

### 6.1 Client Onboarding
1. Client requests services and selects subscription
2. V1 Consortium creates client organization
3. Client administrator receives access credentials
4. Client completes organization profile and compliance setup
5. Employees are enrolled in compliance programs
6. Testing and monitoring begins

### 6.2 Drug Testing Process
1. System generates random selection or scheduled test
2. Employee receives testing notification
3. Employee reports to testing facility
4. Test is conducted and results sent to lab
5. MRO reviews results if necessary
6. Results are reported to DER
7. Appropriate follow-up actions taken

### 6.3 MVR Monitoring
1. System automatically orders MVR for employee
2. MVR results are received and analyzed
3. Violations are flagged and reported
4. DER is notified of disqualifying violations
5. Corrective actions are coordinated
6. Compliance status is updated

### 6.4 DOT Physical Process
1. Employee physical expiration is monitored
2. Renewal reminders are sent
3. Physical appointment is scheduled
4. Examination is conducted by certified examiner
5. Medical certificate is issued and tracked
6. Compliance status is updated

## 7. Technical Requirements

### 7.1 Frontend
- SvelteKit with TypeScript
- Responsive design framework
- Progressive Web App capabilities
- Modern UI/UX principles
- Accessibility features

### 7.2 Backend
- SvelteKit server routes
- TypeScript
- RESTful API design
- Microservices architecture
- Proper error handling and logging

### 7.3 Database
- PostgreSQL via Supabase
- HIPAA-compliant data storage
- Efficient indexing strategy
- Data retention policies
- Backup and recovery procedures

### 7.4 Authentication and Authorization
- Supabase Auth with MFA
- Role-based access control
- JWT-based authentication
- Session management
- Audit logging

### 7.5 Storage
- Supabase Storage
- Document versioning
- Secure file handling
- Storage encryption
- Retention management

### 7.6 Integrations
- Laboratory API connections
- MVR provider integrations
- FMCSA Clearinghouse API
- Third-party system APIs
- Webhook support

## 8. Future Considerations

### 8.1 Potential Enhancements
- AI-powered compliance predictions
- Advanced analytics and reporting
- Mobile applications for iOS/Android
- Blockchain for document verification
- IoT integration for workplace monitoring
- Telemedicine integration
- Wearable device monitoring

### 8.2 Regulatory Adaptations
- New DOT regulations
- State-specific requirements
- International compliance standards
- Industry-specific regulations
- Technology advancement integration

## 9. Appendices

### 9.1 Glossary
- **TPA**: Third Party Administrator - manages compliance programs for multiple employers
- **DER**: Designated Employer Representative - manages drug and alcohol programs
- **MRO**: Medical Review Officer - physician who reviews drug test results
- **SAP**: Substance Abuse Professional - evaluates employees who violate drug policies
- **DOT**: Department of Transportation - federal agency regulating transportation safety
- **MVR**: Motor Vehicle Record - driving history report
- **Clearinghouse**: FMCSA database of drug and alcohol violations

### 9.2 Regulatory References
- 49 CFR Part 40 - DOT drug and alcohol testing procedures
- 49 CFR Part 382 - FMCSA drug and alcohol regulations
- HIPAA Privacy Rule
- Fair Credit Reporting Act (FCRA)
- State-specific transportation regulations