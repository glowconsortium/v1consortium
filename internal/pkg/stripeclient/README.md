# V1 Consortium Stripe Client Package

A comprehensive Stripe integration package built specifically for V1 Consortium's transportation compliance platform. This package provides full subscription management, payment processing, and customer billing capabilities using the Stripe Go SDK v83.

## Overview

This package provides a modular, type-safe wrapper around Stripe's API with specific features for:
- Annual subscription management with 4 tiers (Starter, Professional, Enterprise, Custom)
- Customer management with compliance-specific metadata
- Payment method storage and management
- Invoice management and automation
- Webhook handling for real-time events
- Checkout session creation for payment collection
- **Comprehensive error handling with known error types**
- V1 Consortium-specific business logic utilities

## Package Structure

```
internal/pkg/stripeclient/
├── types.go           # All data structures and type definitions
├── client.go          # Main client initialization and configuration
├── errors.go          # Known error definitions and error wrapping
├── customers.go       # Customer CRUD operations
├── payment_methods.go # Payment method and setup intent management
├── subscriptions.go   # Subscription lifecycle management
├── invoices.go        # Invoice management and operations
├── webhooks.go        # Webhook validation and handling
├── checkout.go        # Checkout session creation
├── utilities.go       # V1 Consortium-specific business logic
└── example_usage.go   # Examples of error handling patterns
```

## Error Handling

This package provides comprehensive error handling with specific known errors for common Stripe operations. All client methods return wrapped errors that can be checked against known error types.

### Known Error Types

#### Customer Errors
- `ErrCustomerNotFound` - Customer does not exist
- `ErrCustomerAlreadyExists` - Customer already exists with this email
- `ErrInvalidCustomerData` - Invalid customer data provided
- `ErrCustomerDeleted` - Customer has been deleted

#### Payment Method Errors
- `ErrPaymentMethodNotFound` - Payment method does not exist
- `ErrPaymentMethodInvalid` - Payment method data is invalid
- `ErrPaymentMethodDeclined` - Payment method was declined
- `ErrPaymentMethodExpired` - Payment method has expired
- `ErrPaymentMethodInsufficientFunds` - Insufficient funds
- `ErrPaymentMethodNotAttached` - Payment method not attached to customer

#### Subscription Errors
- `ErrSubscriptionNotFound` - Subscription does not exist
- `ErrSubscriptionCanceled` - Subscription has been canceled
- `ErrSubscriptionIncomplete` - Subscription is incomplete
- `ErrSubscriptionPastDue` - Subscription is past due
- `ErrInvalidSubscriptionTier` - Invalid subscription tier
- `ErrTrialAlreadyUsed` - Trial period has already been used

#### Invoice Errors
- `ErrInvoiceNotFound` - Invoice does not exist
- `ErrInvoiceAlreadyPaid` - Invoice has already been paid
- `ErrInvoiceVoid` - Invoice has been voided
- `ErrInvoicePaymentFailed` - Invoice payment failed

#### Webhook Errors
- `ErrInvalidWebhookSignature` - Invalid webhook signature
- `ErrWebhookEventNotSupported` - Webhook event type not supported
- `ErrWebhookProcessingFailed` - Webhook processing failed

#### General Errors
- `ErrInvalidAPIKey` - Invalid API key
- `ErrRateLimitExceeded` - Rate limit exceeded
- `ErrNetworkError` - Network error
- `ErrInternalServerError` - Internal server error
- `ErrInvalidRequest` - Invalid request
- `ErrPermissionDenied` - Permission denied

### Error Handling Functions

```go
// Check if error is a specific known error
if stripeclient.IsKnownError(err, stripeclient.ErrCustomerNotFound) {
    // Handle customer not found
}

// Check if error is any Stripe error
if stripeclient.IsStripeError(err) {
    code := stripeclient.GetErrorCode(err)
    requestID := stripeclient.GetRequestID(err)
}

// Wrap raw Stripe errors
wrappedErr := stripeclient.WrapStripeError(rawStripeErr)
```

### Error Handling Example

```go
customer, err := client.CreateCustomer(ctx, customerData)
if err != nil {
    // Check for specific known errors
    if stripeclient.IsKnownError(err, stripeclient.ErrCustomerAlreadyExists) {
        log.Printf("Customer already exists: %v", err)
        // Handle existing customer logic
    } else if stripeclient.IsKnownError(err, stripeclient.ErrInvalidCustomerData) {
        log.Printf("Invalid customer data: %v", err)
        // Return validation error to user
    } else if stripeclient.IsKnownError(err, stripeclient.ErrRateLimitExceeded) {
        log.Printf("Rate limit exceeded, please retry later: %v", err)
        // Implement retry logic
    } else if stripeclient.IsStripeError(err) {
        // This is a Stripe error but not one of our specific known errors
        log.Printf("Stripe error occurred: %v", err)
        log.Printf("Error code: %s", stripeclient.GetErrorCode(err))
        log.Printf("Request ID: %s", stripeclient.GetRequestID(err))
    } else {
        // General error
        log.Printf("Unexpected error: %v", err)
    }
    return
}
```

## Configuration

### Environment Variables
```bash
STRIPE_SECRET_KEY=sk_test_...  # Your Stripe secret key
STRIPE_WEBHOOK_SECRET=whsec_...  # Your webhook endpoint secret
```

### Initialize Client
```go
import "your-project/internal/pkg/stripeclient"

config := &stripeclient.Config{
    SecretKey:     "sk_test_...",
    WebhookSecret: "whsec_...",
    MaxRetries:    3,
}

client := stripeclient.NewStripeClient(config)
```

## Core Features

### Customer Management
```go
// Create a new customer
customer, err := client.CreateCustomer(ctx, &stripeclient.CustomerData{
    Email: "john@company.com",
    Name:  "John Doe",
    Metadata: map[string]string{
        "company_size": "50",
        "industry": "transportation",
    },
})

// Update customer information
customer, err = client.UpdateCustomer(ctx, "cus_...", &stripeclient.CustomerData{
    Name: "John Smith",
})

// Get customer details
customer, err = client.GetCustomer(ctx, "cus_...")
```

### Subscription Management
```go
// Create annual subscription
subscription, err := client.CreateSubscription(
    ctx, 
    customerID, 
    "price_professional_annual", 
    100, // employee count
    map[string]string{"tier": "professional"},
)

// Cancel subscription
err = client.CancelSubscription(ctx, "sub_...", true) // cancel at period end

// Update subscription
subscription, err = client.UpdateSubscription(ctx, "sub_...", updateParams)
```

### Payment Methods
```go
// Create setup intent for saving payment method
setupIntent, err := client.CreateSetupIntent(ctx, customerID, "card")

// List customer's payment methods
paymentMethods, err := client.ListCustomerPaymentMethods(ctx, customerID)

// Set default payment method
err = client.SetDefaultPaymentMethod(ctx, customerID, "pm_...")
```

### Invoice Management
```go
// Get specific invoice
invoice, err := client.GetInvoice(ctx, "in_...")

// List customer invoices
invoices, err := client.ListCustomerInvoices(ctx, customerID, 12)

// Create and send invoice
invoice, err = client.CreateInvoice(ctx, customerID, "Monthly services", true)
err = client.SendInvoice(ctx, invoice.ID)

// Get upcoming invoice preview
upcoming, err := client.GetUpcomingInvoice(ctx, subscriptionID)
```

### Checkout Sessions
```go
// Create one-time payment session
session, err := client.CreateCheckoutSession(
    ctx,
    customerID,
    "price_...",
    "https://yoursite.com/success",
    "https://yoursite.com/cancel",
)

// Create subscription checkout session
session, err := client.CreateSubscriptionCheckoutSession(
    ctx,
    customerID,
    "price_...",
    successURL,
    cancelURL,
    14, // trial days
)
```

## V1 Consortium Business Logic

### Complete Customer Setup
```go
// Create customer with subscription in one call
customer, subscription, err := client.CreateCompleteCustomerWithSubscription(
    ctx,
    "john@company.com",
    "John Doe",
    stripeclient.TierProfessional,
    100, // employee count
    "pm_...", // payment method ID
)
```

### Subscription Plans
```go
// Get all available plans
plans := stripeclient.GetDefaultSubscriptionPlans()

// Get price ID for tier
priceID := stripeclient.GetDefaultPriceForTier(stripeclient.TierEnterprise)
```

### Comprehensive Billing Info
```go
// Get complete customer billing overview
billingInfo, err := client.GetCustomerUsageAndBilling(ctx, customerID)
// Returns customer, subscriptions, invoices, and payment methods
```

## Webhook Integration

### Setup Webhook Handler
```go
func webhookHandler(w http.ResponseWriter, r *http.Request) {
    event, err := client.ParseWebhookFromRequest(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    switch event.Type {
    case "customer.subscription.updated":
        err = client.HandleCustomerSubscriptionUpdated(event)
    case "invoice.payment_succeeded":
        err = client.HandleInvoicePaymentSucceeded(event)
    case "invoice.payment_failed":
        err = client.HandleInvoicePaymentFailed(event)
    }

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
```

## Subscription Tiers

The package supports 4 predefined subscription tiers:

1. **Starter** (`TierStarter`)
   - Up to 50 employees
   - Basic compliance tracking
   - Email support

2. **Professional** (`TierProfessional`)
   - Up to 200 employees  
   - Advanced compliance automation
   - Drug & alcohol testing integration
   - Priority support

3. **Enterprise** (`TierEnterprise`)
   - Up to 1000 employees
   - Custom compliance workflows
   - API access
   - Dedicated success manager

4. **Custom** (`TierCustom`)
   - Unlimited employees
   - Custom development
   - On-premise options
   - 24/7 dedicated support

## Error Handling

All functions return detailed errors with context:

```go
subscription, err := client.CreateSubscription(ctx, customerID, priceID, employeeCount, metadata)
if err != nil {
    if strings.Contains(err.Error(), "customer not found") {
        // Handle missing customer
    } else if strings.Contains(err.Error(), "payment method required") {
        // Handle missing payment method
    }
    return fmt.Errorf("subscription creation failed: %w", err)
}
```

## Annual Billing Model

The package is designed around V1 Consortium's annual billing cycle:
- Subscriptions run January 1 - December 31
- Automatic renewals with prorated adjustments
- Employee count tracking and tier validation
- Mid-year upgrades with proper proration

## Testing

To test the package:
1. Use Stripe test keys (starting with `sk_test_`)
2. Use test webhook endpoints
3. Test with Stripe's test card numbers
4. Use the Stripe CLI for webhook testing

## Security Considerations

- API keys are never logged or exposed
- Webhook signatures are always validated
- Customer data includes only necessary information
- PCI compliance handled by Stripe (no card data stored)

## Dependencies

- `github.com/stripe/stripe-go/v83` - Official Stripe Go SDK
- Standard Go libraries for HTTP and JSON handling

This package provides everything needed for V1 Consortium's subscription billing and payment processing requirements while maintaining flexibility for future enhancements.