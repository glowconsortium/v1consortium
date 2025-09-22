# FleetOps - Fleet Management System
## Product Requirements Document

**Version:** 1.0  
**Date:** April 23, 2025  
**Author:** Claude

## 1. Introduction

### 1.1 Purpose
This Product Requirements Document (PRD) outlines the functional and non-functional requirements for FleetOps, a comprehensive fleet management SaaS platform built with SvelteKit, TypeScript, and Supabase. The system enables organizations to manage carriers, trucks, trailers, drivers, and dispatchers with various user roles and subscription tiers.

### 1.2 Scope
FleetOps will provide a complete solution for transportation and logistics companies to manage their fleet operations, including carrier management, equipment tracking, driver assignment, trip recording, and performance analytics. The system will support both internal administrative users and external customer organizations with different subscription tiers.

### 1.3 Definitions, Acronyms, and Abbreviations
- **SU** - Super User
- **BOL** - Bill of Lading
- **ELD** - Electronic Logging Device
- **Carrier** - A transportation company that owns and operates trucks
- **Equipment** - A truck, trailer, or truck-trailer combination used for transportation
- **Lane** - A table that defines operational routes for equipment with timing specifications
- **Trip** - A single transportation journey from origin to destination
- **Geopoint** - A geographic coordinate (latitude and longitude)

## 2. Product Overview

### 2.1 Product Perspective
FleetOps is a new web-based SaaS application designed to serve the transportation and logistics industry. It will operate as a standalone system with potential integrations with external systems like email and ELD APIs.

### 2.2 Product Features
- Organization management (internal, external, machines)
- User management with role-based access control
- Carrier and equipment management
- Lane and trip management
- Subscription and feature access control
- Document management
- In-app messaging
- Analytics and reporting
- Mobile-optimized driver interface

### 2.3 User Classes and Characteristics

#### 2.3.1 Internal Users
- **Internal Super User (internal_su)**: System administrators with full access to all features and organizations
- **Internal User (internal_user)**: Support staff with restricted administrative access

#### 2.3.2 External Users
- **External Super User (external_su)**: Organization administrators with full access to their organization
- **External Manager (external_manager)**: Manages carriers and overall operations
- **External Dispatcher (external_dispatcher)**: Assigns drivers and manages trips
- **External Driver (external_driver)**: Performs trips and updates status
- **External Carrier Owner (external_carrier_owner)**: Owns and manages carriers

#### 2.3.3 Machine Users
- Automated services and integrations (placeholder for future expansion)

### 2.4 Operating Environment
- Web browsers (Chrome, Firefox, Safari, Edge)
- Mobile devices (responsive design with mobile-optimized views for drivers)
- Backend: SvelteKit server with Supabase
- Database: PostgreSQL via Supabase

### 2.5 Design and Implementation Constraints
- Built with SvelteKit and TypeScript
- Utilizes Supabase for database, authentication, and storage
- Responsive design to support desktop and mobile devices
- RESTful API architecture

### 2.6 Assumptions and Dependencies
- Reliable internet connectivity for users
- Supabase services availability
- Modern web browser support

## 3. Specific Requirements

### 3.1 Organization Management

#### 3.1.1 Organization Types
- **Internal**: Administrative organization for system management
- **External**: Customer organizations using the system to manage their fleet
- **Machines**: Service integrations and automated processes (placeholder)

#### 3.1.2 Organization Profile
- Name
- Type (internal, external, machines)
- Contact information
- Address
- Logo
- Subscription details (for external organizations)
- Feature entitlements
- Created date
- Status (active, inactive, trial)

#### 3.1.3 Organization Management Functions
- Create, read, update, delete (CRUD) operations for organizations
- Activate/deactivate organizations
- Manage subscription status
- Assign features based on subscription tier
- Monitor usage against subscription limits

### 3.2 User Management

#### 3.2.1 User Profile
- Username
- Email
- Password (securely stored)
- Full name
- Contact information
- Organization affiliation
- Role
- Status (active, inactive)
- Profile picture
- Created date
- Last login date

#### 3.2.2 User Roles and Permissions
- **Internal Super User (internal_su)**
  - Manage all organizations
  - Create/manage internal users
  - Manage all subscription tiers
  - Configure system settings
  - Generate trial periods and promo codes
  - Access all analytics and reports

- **Internal User (internal_user)**
  - View organizations and users
  - Provide support to external organizations
  - Generate reports
  - Limited administrative capabilities

- **External Super User (external_su)**
  - Manage organization profile
  - Invite and manage users within their organization
  - Assign roles to users
  - Access all features available to their subscription tier
  - View organization-wide analytics

- **External Manager (external_manager)**
  - Manage carriers assigned to them
  - View analytics for their carriers
  - Create and manage channels

- **External Dispatcher (external_dispatcher)**
  - Assign drivers to trips
  - Monitor trip progress
  - Update trip status

- **External Driver (external_driver)**
  - View assigned trips
  - Update trip status
  - Upload trip documents
  - View rate sheets

- **External Carrier Owner (external_carrier_owner)**
  - Manage specific carriers
  - View carrier analytics
  - Manage carrier integrations

#### 3.2.3 User Management Functions
- Registration request workflow
- User approval process
- User invitation system
- Role assignment
- Password reset
- Account activation/deactivation
- User profile management

### 3.3 Subscription Management

#### 3.3.1 Subscription Tiers
- **Free**
  - 3 users
  - 1 carrier
  - 2 trucks
  - 2 trailers
  - 2 features (configurable)
  - Limited storage

- **Bronze**
  - 10 users
  - 5 carriers
  - 10 trucks
  - 10 trailers
  - 5 features (configurable)
  - Moderate storage

- **Silver**
  - 50 users
  - 20 carriers
  - 50 trucks
  - 50 trailers
  - More features (configurable)
  - Larger storage

- **Platinum**
  - Unlimited users
  - Unlimited carriers
  - Unlimited trucks
  - Unlimited trailers
  - All features
  - Maximum storage

#### 3.3.2 Trial Management
- Configurable trial periods (days)
- Promo code generation
- Automatic trial expiration notifications
- Trial-to-paid conversion workflow

#### 3.3.3 Subscription Management Functions
- Subscribe/upgrade/downgrade subscription
- Track usage against subscription limits
- Process payments (online)
- Record manual payments (offline)
- Generate invoices
- Send payment reminders
- Manage subscription renewal

### 3.4 Carrier Management

#### 3.4.1 Carrier Profile
- Name
- MC/DOT number
- Contact information
- Address
- Status (active, inactive)
- Associated organization
- Assigned manager
- Integration settings (email, ELD)
- Created date

#### 3.4.2 Carrier Management Functions
- CRUD operations for carriers
- Assign managers to carriers
- Configure carrier integrations
- View carrier analytics
- Manage carrier documentation

### 3.5 Equipment Management

#### 3.5.1 Equipment Types
- Truck
- Trailer
- Equipment wrapper (truck + trailer + driver + dispatcher combination)

#### 3.5.2 Equipment Profile
- Type (truck, trailer)
- Identification number
- Make/model
- Year
- VIN
- License plate
- Status (active, inactive, maintenance)
- Associated carrier
- Assigned driver (for trucks)
- Assigned dispatcher
- Created date

#### 3.5.3 Equipment Management Functions
- CRUD operations for equipment
- Assign drivers to trucks
- Assign dispatchers to equipment
- Track equipment status
- Manage equipment documentation
- Record maintenance events
- Track expenses

### 3.6 Lane and Trip Management

#### 3.6.1 Lane Definition
- Name
- Description
- Route information
- Weekly timeframe
- Associated equipment
- Notes/comments
- Status (active, inactive)

#### 3.6.2 Trip Profile
- Associated lane
- Start location (city, state, geopoint)
- End location (city, state, geopoint)
- Start date/time
- End date/time
- Status (pending, in progress, completed, cancelled)
- Associated truck
- Associated driver
- Associated dispatcher
- Broker information
- Amount/payment details
- Notes
- Documents (BOL, lumber, etc.)

#### 3.6.3 Lane and Trip Management Functions
- CRUD operations for lanes
- CRUD operations for trips
- Route planning
- Trip status updates
- Document upload
- Trip analytics

### 3.7 Document Management

#### 3.7.1 Document Types
- BOL (Bill of Lading)
- Proof of delivery
- Insurance documents
- Vehicle registration
- Driver's license
- Maintenance records
- Expense receipts
- Rate sheets

#### 3.7.2 Document Management Functions
- Upload/download documents
- Document categorization
- Document search
- Storage quota management based on subscription tier
- Document sharing between users

### 3.8 Messaging and Communication

#### 3.8.1 Messaging Types
- Direct messages between users
- Channel messages (group conversations)
- System notifications

#### 3.8.2 Messaging Functions
- Send/receive direct messages
- Create and manage channels
- Invite users to channels
- Notification preferences
- Message moderation (for internal users)
- Message search

### 3.9 Analytics and Reporting

#### 3.9.1 Analytics Types
- Equipment performance metrics
- Driver performance metrics
- Carrier performance metrics
- Trip analytics
- Financial analytics
- Subscription usage analytics

#### 3.9.2 Reporting Functions
- Generate predefined reports
- Create custom reports
- Export reports (CSV, PDF)
- Schedule automated reports
- Dashboard visualizations

### 3.10 Mobile-Optimized Driver Interface

#### 3.10.1 Driver Dashboard
- View assigned trips
- View past trips
- View ongoing trips
- Update trip status
- View rate sheets
- Upload documents (BOL, etc.)
- Messaging access

#### 3.10.2 Driver-Specific Functions
- Trip navigation assistance
- Quick status updates
- Document scanning/upload
- Driver checklist

### 3.11 Integration Capabilities

#### 3.11.1 Integration Types
- Email integration
- ELD API integration
- Future integrations (placeholder)

#### 3.11.2 Integration Functions
- Configure integration settings
- Authenticate with external systems
- Sync data between systems
- Monitor integration status

## 4. Feature Entitlement Matrix

| Feature | Free | Bronze | Silver | Platinum |
|---------|------|--------|--------|----------|
| Organization Management | ✓ | ✓ | ✓ | ✓ |
| User Management | ✓ (3 users) | ✓ (10 users) | ✓ (50 users) | ✓ (Unlimited) |
| Carrier Management | ✓ (1 carrier) | ✓ (5 carriers) | ✓ (20 carriers) | ✓ (Unlimited) |
| Equipment Management | ✓ (2 trucks, 2 trailers) | ✓ (10 trucks, 10 trailers) | ✓ (50 trucks, 50 trailers) | ✓ (Unlimited) |
| Lane Management | ✓ | ✓ | ✓ | ✓ |
| Trip Management | ✓ | ✓ | ✓ | ✓ |
| Basic Reporting | ✓ | ✓ | ✓ | ✓ |
| Advanced Analytics | ✗ | ✓ | ✓ | ✓ |
| Document Management | ✓ (Limited storage) | ✓ (Moderate storage) | ✓ (Large storage) | ✓ (Maximum storage) |
| Messaging | ✓ | ✓ | ✓ | ✓ |
| Channels | ✗ | ✓ (Limited) | ✓ | ✓ |
| Mobile Driver Interface | ✓ | ✓ | ✓ | ✓ |
| Email Integration | ✗ | ✓ | ✓ | ✓ |
| ELD Integration | ✗ | ✗ | ✓ | ✓ |
| Custom Features | ✓ (2 max) | ✓ (5 max) | ✓ (10 max) | ✓ (All) |

## 5. Non-Functional Requirements

### 5.1 Performance
- Page load time < 2 seconds for standard operations
- Support for 1000+ concurrent users
- Database query response time < 500ms
- Real-time updates for messaging and status changes

### 5.2 Security
- Role-based access control
- Data encryption in transit and at rest
- Secure authentication and authorization
- Session management
- Regular security audits
- GDPR compliance

### 5.3 Reliability
- System uptime > 99.9%
- Data backup and recovery procedures
- Error handling and logging
- Graceful degradation under high load

### 5.4 Usability
- Intuitive user interface
- Responsive design for all devices
- Consistent design language
- Accessibility compliance (WCAG 2.1)
- Multi-language support

### 5.5 Scalability
- Horizontal scaling for increased load
- Database sharding capabilities
- Efficient resource utilization
- Caching strategies

## 6. User Flows

### 6.1 User Registration and Onboarding
1. External user requests account with subscription type
2. System captures organization and user details
3. Internal user reviews and approves request
4. System creates organization and super user account
5. External super user receives welcome email
6. External super user completes organization profile
7. External super user invites additional users

### 6.2 Carrier and Equipment Setup
1. External super user creates carrier
2. External super user adds trucks and trailers
3. External super user assigns managers to carriers
4. External super user or manager adds drivers
5. Dispatchers assign drivers to trucks

### 6.3 Trip Management
1. Dispatcher creates lane for equipment
2. Dispatcher creates trip within lane
3. Driver receives trip assignment
4. Driver updates trip status
5. Driver uploads documents
6. Dispatcher reviews completed trip
7. System updates analytics

### 6.4 Subscription Management
1. External organization selects subscription tier
2. System provisions features based on tier
3. System monitors usage against limits
4. System notifies when approaching limits
5. External organization manages payment
6. System renews or modifies subscription

## 7. Technical Requirements

### 7.1 Frontend
- SvelteKit with TypeScript
- Responsive design
- Progressive Web App capabilities
- Modern UI/UX principles

### 7.2 Backend
- SvelteKit server routes
- TypeScript
- RESTful API design
- Proper error handling

### 7.3 Database
- PostgreSQL via Supabase
- Efficient schema design
- Proper indexing
- Data validation

### 7.4 Authentication and Authorization
- Supabase Auth
- JWT-based authentication
- Role-based access control
- Multi-factor authentication (optional)

### 7.5 Storage
- Supabase Storage
- Efficient document storage and retrieval
- Storage quota management

### 7.6 Realtime Features
- Supabase Realtime
- WebSocket connections for messaging
- Real-time status updates

## 8. Future Considerations

### 8.1 Potential Future Features
- AI-driven route optimization
- Predictive maintenance
- Advanced analytics dashboard
- Mobile app for drivers
- Fuel management system
- Integration with accounting software
- Inventory management
- Customer portal

### 8.2 Scalability Considerations
- Multi-region deployment
- Content delivery network integration
- Database sharding strategy
- Microservices architecture evolution

## 9. Appendices

### 9.1 Glossary
- **Carrier**: A transportation company that owns and operates trucks
- **Equipment**: A truck, trailer, or truck-trailer combination used for transportation
- **Lane**: A table that defines operational routes for equipment with timing specifications
- **Trip**: A single transportation journey from origin to destination
- **BOL**: Bill of Lading, a document acknowledging receipt of cargo for shipment
- **ELD**: Electronic Logging Device, used to record driver's hours of service

### 9.2 References
- Transport industry standards and regulations
- Supabase documentation
- SvelteKit documentation
- TypeScript documentation