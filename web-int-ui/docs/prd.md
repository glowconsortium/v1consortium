# Product Requirements Document (PRD)

**Project Name:** Carrier Compliance SaaS
**Version:** v1.1
**Author:** System Architect
**Date:** August 21, 2025u want to extend the PRD to cover **random drug/alcohol testing program management** (required by DOT/FMCSA consortium rules) and **regulatory document downloads** (certificates, compliance docs, audit-ready files).

Here’s the **updated PRD** with those additions:

---

# Product Requirements Document (PRD)

**Project Name:** Carrier Compliance SaaS
**Version:** v1.1
**Author:** \[Your Name]
**Date:** \[Today’s Date]

---

## 1. Executive Summary

This SaaS platform enables motor carriers to manage CDL drivers, ensuring compliance with DOT/FMCSA requirements for drug & alcohol testing, DOT physicals, and MVR checks.

The system offers:

* **Public-Facing Site** for marketing, signup, pricing, and subscription purchase.
* **Carrier Portal** for managing drivers, ordering compliance services, conducting random testing, and downloading DOT/FMSCA-compliant certificates.

---

## 2. Goals & Objectives

* Provide carriers with a **complete compliance management system**.
* Automate **drug tests, DOT physicals, MVR checks, and random test selection**.
* Allow carriers to **download required DOT/FMCSA certificates and compliance docs**.
* Support **annual subscription (Jan–Dec cycle)** with renewal enforcement.
* Provide **integrations with third-party compliance/testing APIs**.

---

## 3. Target Users

* **Primary:** Motor carriers, owner-operators enrolled in consortiums.
* **Secondary:** Compliance officers, safety managers.
* **Admins:** Internal compliance staff.

---

## 4. Core Features

### 4.1 Authentication & Subscription

* Supabase-based auth (email/password, OAuth options).
* Annual subscription model (Jan–Dec).
* Stripe (or similar) for billing + receipts.
* Grace period for renewals.

### 4.2 Carrier & Driver Management

* Add/Edit/Delete driver profiles.
* Upload supporting docs (CDL, medical cert, etc.).
* Track compliance status per driver.

### 4.3 Compliance Services

* **Drug & Alcohol Testing**

  * Order through third-party APIs (Quest Diagnostics, etc.).
  * Retrieve results automatically.
* **DOT Physicals**

  * Schedule/order exams.
  * Store and track medical certs.
* **Motor Vehicle Reports (MVRs)**

  * Request and store MVRs.

### 4.4 **Random Testing Management** (NEW)

* Create/Manage random testing pools (all drivers enrolled).
* Automated random selection of drivers (quarterly/monthly basis).
* Notifications to carriers & selected drivers.
* Audit logs of selections to meet FMCSA requirements.

### 4.5 Reporting & Certificates (NEW)

* Compliance dashboard (per carrier + per driver).
* Downloadable certificates:

  * **Consortium Membership Certificate** (proof of enrollment).
  * **Test Completion Certificates** (per driver/test).
  * **Random Selection Reports** (FMCSA audit requirement).
* Export compliance history in **PDF/CSV** format.

### 4.6 Third-Party Integrations

* Token-based API integrations with external carrier management systems.
* Temporal workflows to handle async test ordering, polling, and results.

### 4.7 Admin Panel

* Manage carriers, subscriptions, payments.
* Oversee random testing selections and pools.
* Generate compliance/audit reports.
* Support tools for customer service.

---

## 5. Tech Stack

* **Frontend:** SvelteKit
* **Backend:** Go (GoFrame framework)
* **Workflows:** Temporal (e.g., scheduling random tests, polling APIs)
* **Database:** Supabase (Postgres + Auth)
* **Payments:** Stripe (annual billing)
* **Third-Party APIs:** Quest Diagnostics, MVR vendors, DOT physical providers

---

## 6. User Flows

### 6.1 Carrier Flow (with random testing & certificates)

1. Sign up & pay subscription.
2. Add drivers to consortium.
3. Order tests (drug/DOT/MVR).
4. System automatically conducts **random driver selection** (quarterly).
5. Carrier downloads **consortium certificate** + compliance documents.

### 6.2 Admin Flow

1. Monitor carrier subscriptions.
2. Verify random selection pool logic.
3. Generate compliance docs for audits.
4. Support carriers with missing data.

---

## 7. System Architecture

### 7.1 High-Level Components

* **Frontend (SvelteKit)** → Carrier/Admin dashboards.
* **Backend (Go + GoFrame)** → API & business logic.
* **Temporal** → Orchestration for test ordering, result polling, random test scheduling.
* **Supabase** → Database + authentication.
* **Stripe** → Subscription billing.
* **Third-Party APIs** → Quest Diagnostics, MVR providers, DOT physical exam providers.
* **Document Generator** → PDF certificate/report generator (e.g., Go PDF lib, reportlab).

---

## 8. MVP Scope

* Subscription (Stripe + Supabase).
* Carrier dashboard + driver management.
* Order drug tests via Quest API.
* View results.
* Generate/download **consortium membership certificate**.
* Random test selection system (basic version).

---

## 9. Future Enhancements

* Multi-vendor test support.
* Mobile app.
* Advanced analytics + compliance scoring.
* Driver portal for scheduling tests.
* Automated reminders (email/SMS).

---

## 10. System Architecture

For detailed system architecture diagrams and technical specifications, refer to [system.md](./system.md).

### 10.1 High-Level Components

* **Frontend (SvelteKit)** → Carrier/Admin dashboards.
* **Backend (Go + GoFrame)** → API & business logic.
* **Temporal** → Orchestration for test ordering, result polling, random test scheduling.
* **Supabase** → Database + authentication.
* **Stripe** → Subscription billing.
* **Third-Party APIs** → Quest Diagnostics, MVR providers, DOT physical exam providers.
* **Document Generator** → PDF certificate/report generator (e.g., Go PDF lib, reportlab).

---
