// Package stripeclient provides examples of how to use the known errors
// This file demonstrates error handling patterns for the V1 Consortium Stripe client
package stripeclient

import (
	"context"
	"fmt"
	"log"
)

// ExampleUsage demonstrates how to use the known errors in the Stripe client
func ExampleUsage() {
	// Initialize the Stripe client
	config := &Config{
		SecretKey:         "sk_test_...", // Replace with your test secret key
		PublishableKey:    "pk_test_...", // Replace with your test publishable key
		WebhookSecret:     "whsec_...",   // Replace with your webhook secret
		DefaultCurrency:   "usd",
		TestMode:          true,
		MaxNetworkRetries: 3,
		APIVersion:        "2023-10-16",
	}

	client := NewStripeClient(config)

	// Example 1: Customer operation with error handling
	ctx := context.Background()

	customerData := &CustomerData{
		Email:       "test@example.com",
		Name:        "Test User",
		Description: "Test customer for V1 Consortium",
		Phone:       "+1234567890",
		Metadata:    map[string]string{"source": "example"},
	}

	customer, err := client.CreateCustomer(ctx, customerData)
	if err != nil {
		// Check for specific known errors
		if IsKnownError(err, ErrCustomerAlreadyExists) {
			log.Printf("Customer already exists: %v", err)
		} else if IsKnownError(err, ErrInvalidCustomerData) {
			log.Printf("Invalid customer data: %v", err)
		} else if IsKnownError(err, ErrRateLimitExceeded) {
			log.Printf("Rate limit exceeded, please retry later: %v", err)
		} else if IsStripeError(err) {
			// This is a Stripe error but not one of our specific known errors
			log.Printf("Stripe error occurred: %v", err)
			log.Printf("Error code: %s", GetErrorCode(err))
			log.Printf("Request ID: %s", GetRequestID(err))
		} else {
			// General error
			log.Printf("Unexpected error: %v", err)
		}
		return
	}

	fmt.Printf("Created customer: %s (%s)\n", customer.Name, customer.ID)

	// Example 2: Payment method operation with error handling
	_, err = client.CreateSetupIntent(ctx, customer.ID, "off_session", map[string]string{
		"purpose": "save_payment_method",
	})
	if err != nil {
		if IsKnownError(err, ErrCustomerNotFound) {
			log.Printf("Customer not found: %v", err)
		} else if IsKnownError(err, ErrPaymentMethodInvalid) {
			log.Printf("Invalid payment method: %v", err)
		} else {
			log.Printf("Error creating setup intent: %v", err)
		}
		return
	}

	// Example 3: Subscription operation with error handling
	subscriptionData, err := client.CreateSubscription(ctx, customer.ID, "price_...", 14, map[string]string{
		"tier": "starter",
	})
	if err != nil {
		if IsKnownError(err, ErrTrialAlreadyUsed) {
			log.Printf("Trial period already used: %v", err)
		} else if IsKnownError(err, ErrInvalidSubscriptionTier) {
			log.Printf("Invalid subscription tier: %v", err)
		} else if IsKnownError(err, ErrPaymentMethodDeclined) {
			log.Printf("Payment method declined: %v", err)
		} else {
			log.Printf("Error creating subscription: %v", err)
		}
		return
	}

	fmt.Printf("Created subscription: %s (Status: %s)\n", subscriptionData.ID, subscriptionData.Status)

	// Example 4: Invoice operations with error handling
	invoices, err := client.ListCustomerInvoices(ctx, customer.ID, 10)
	if err != nil {
		if IsKnownError(err, ErrCustomerNotFound) {
			log.Printf("Customer not found: %v", err)
		} else {
			log.Printf("Error listing invoices: %v", err)
		}
		return
	}

	fmt.Printf("Found %d invoices for customer\n", len(invoices))

	// Example 5: Webhook validation with error handling
	webhookPayload := []byte(`{"id": "evt_test_webhook", "object": "event"}`)
	webhookSignature := "test_signature"

	_, err = client.ValidateWebhook(webhookPayload, webhookSignature)
	if err != nil {
		if IsKnownError(err, ErrInvalidWebhookSignature) {
			log.Printf("Invalid webhook signature: %v", err)
		} else if IsKnownError(err, ErrWebhookEventNotSupported) {
			log.Printf("Webhook event not supported: %v", err)
		} else {
			log.Printf("Error validating webhook: %v", err)
		}
		return
	}

	fmt.Println("All operations completed successfully!")
}

// ExampleSubscriptionService shows error handling in a service layer
type ExampleSubscriptionService struct {
	stripeClient *StripeClient
}

func (s *ExampleSubscriptionService) CreateUserSubscription(ctx context.Context, userEmail, tier string) error {
	// First check if customer exists
	customerData := &CustomerData{
		Email: userEmail,
		Name:  "User Name", // This would come from your user data
	}

	customer, err := s.stripeClient.CreateCustomer(ctx, customerData)
	if err != nil {
		// Handle specific customer creation errors
		if IsKnownError(err, ErrCustomerAlreadyExists) {
			// Customer exists, retrieve instead
			return fmt.Errorf("user already has a Stripe customer record")
		} else if IsKnownError(err, ErrInvalidCustomerData) {
			return fmt.Errorf("invalid user data provided: %w", err)
		} else if IsKnownError(err, ErrRateLimitExceeded) {
			return fmt.Errorf("service temporarily unavailable, please try again later")
		}
		return fmt.Errorf("failed to create customer: %w", err)
	}

	// Create subscription based on tier
	var priceID string
	switch tier {
	case "starter":
		priceID = "price_starter_monthly"
	case "professional":
		priceID = "price_professional_monthly"
	case "enterprise":
		priceID = "price_enterprise_monthly"
	default:
		return fmt.Errorf("invalid subscription tier: %s", tier)
	}

	_, err = s.stripeClient.CreateSubscription(ctx, customer.ID, priceID, 14, map[string]string{
		"tier": tier,
	})
	if err != nil {
		// Handle subscription-specific errors
		if IsKnownError(err, ErrPaymentMethodDeclined) {
			return fmt.Errorf("payment method was declined, please try a different payment method")
		} else if IsKnownError(err, ErrPaymentMethodInsufficientFunds) {
			return fmt.Errorf("insufficient funds, please use a different payment method")
		} else if IsKnownError(err, ErrTrialAlreadyUsed) {
			return fmt.Errorf("trial period has already been used for this customer")
		}
		return fmt.Errorf("failed to create subscription: %w", err)
	}

	return nil
}
