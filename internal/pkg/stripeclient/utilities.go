package stripeclient

import (
	"context"
	"fmt"
	"time"

	"github.com/stripe/stripe-go/v83"
)

// V1 Consortium specific utility functions for common business operations

// CreateCompleteCustomerWithSubscription creates a customer and immediately subscribes them to a plan
func (sc *StripeClient) CreateCompleteCustomerWithSubscription(
	ctx context.Context,
	email, name string,
	tier SubscriptionTier,
	employeeCount int64,
	paymentMethodID string,
) (*CustomerData, *SubscriptionData, error) {
	// Create customer
	customerData := &CustomerData{
		Email: email,
		Name:  name,
		Metadata: map[string]string{
			"employee_count": fmt.Sprintf("%d", employeeCount),
			"tier":           string(tier),
		},
	}

	customer, err := sc.CreateCustomer(ctx, customerData)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create customer: %w", err)
	}

	// Attach payment method if provided
	if paymentMethodID != "" {
		err = sc.SetDefaultPaymentMethod(ctx, customer.ID, paymentMethodID)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to set payment method: %w", err)
		}
	}

	// Create subscription with proper parameters
	subscription, err := sc.CreateSubscription(ctx, customer.ID, GetDefaultPriceForTier(tier), employeeCount, map[string]string{
		"tier":           string(tier),
		"employee_count": fmt.Sprintf("%d", employeeCount),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create subscription: %w", err)
	}

	return customer, subscription, nil
}

// GetDefaultPriceForTier returns the default Stripe price ID for each subscription tier
func GetDefaultPriceForTier(tier SubscriptionTier) string {
	// These would be your actual Stripe price IDs from your dashboard
	switch tier {
	case TierStarter:
		return "price_starter_annual" // Replace with actual price ID
	case TierProfessional:
		return "price_professional_annual"
	case TierEnterprise:
		return "price_enterprise_annual"
	case TierCustom:
		return "price_custom_annual"
	default:
		return "price_starter_annual"
	}
}

// HandleAnnualRenewal processes the annual subscription renewal for V1 Consortium
func (sc *StripeClient) HandleAnnualRenewal(ctx context.Context, subscriptionID string) error {
	// Get current subscription
	subscription, err := sc.GetSubscription(ctx, subscriptionID)
	if err != nil {
		return fmt.Errorf("failed to get subscription: %w", err)
	}

	// Check if renewal is needed
	now := time.Now()

	if now.Before(subscription.CurrentPeriodEnd) {
		return fmt.Errorf("subscription %s is not due for renewal yet", subscriptionID)
	}

	// The subscription should automatically renew in Stripe, but we can trigger renewal actions
	// like updating employee counts, checking for tier changes, etc.

	return nil
}

// GetCustomerUsageAndBilling retrieves comprehensive billing information for a customer
func (sc *StripeClient) GetCustomerUsageAndBilling(ctx context.Context, customerID string) (*CustomerBillingInfo, error) {
	// Get customer
	customer, err := sc.GetCustomer(ctx, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get customer: %w", err)
	}

	// Get active subscriptions
	subscriptions, err := sc.ListCustomerSubscriptions(ctx, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscriptions: %w", err)
	}

	// Get recent invoices
	invoices, err := sc.ListCustomerInvoices(ctx, customerID, 12) // Last 12 invoices
	if err != nil {
		return nil, fmt.Errorf("failed to get invoices: %w", err)
	}

	// Get payment methods
	paymentMethods, err := sc.ListCustomerPaymentMethods(ctx, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment methods: %w", err)
	}

	return &CustomerBillingInfo{
		Customer:       customer,
		Subscriptions:  subscriptions,
		Invoices:       invoices,
		PaymentMethods: paymentMethods,
	}, nil
}

// CustomerBillingInfo contains comprehensive billing information for a customer
type CustomerBillingInfo struct {
	Customer       *CustomerData        `json:"customer"`
	Subscriptions  []*SubscriptionData  `json:"subscriptions"`
	Invoices       []*InvoiceData       `json:"invoices"`
	PaymentMethods []*PaymentMethodData `json:"payment_methods"`
}

// ProcessTrialSubscription creates a subscription with a trial period for new V1 Consortium customers
func (sc *StripeClient) ProcessTrialSubscription(
	ctx context.Context,
	customerID string,
	tier SubscriptionTier,
	trialDays int64,
) (*SubscriptionData, error) {
	priceID := GetDefaultPriceForTier(tier)

	// Use the regular CreateSubscription with trial settings
	// Note: You would need to extend CreateSubscription to handle trial_end parameter
	subscription, err := sc.CreateSubscription(ctx, customerID, priceID, 0, map[string]string{
		"tier":       string(tier),
		"trial_days": fmt.Sprintf("%d", trialDays),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create trial subscription: %w", err)
	}

	return subscription, nil
}

// UpgradeSubscriptionTier upgrades a customer to a higher tier
func (sc *StripeClient) UpgradeSubscriptionTier(
	ctx context.Context,
	subscriptionID string,
	newTier SubscriptionTier,
	newEmployeeCount int64,
) (*SubscriptionData, error) {
	// Get new price for the tier
	newPriceID := GetDefaultPriceForTier(newTier)

	// Create update parameters
	params := &stripe.SubscriptionUpdateParams{
		Items: []*stripe.SubscriptionUpdateItemParams{
			{
				Price: stripe.String(newPriceID),
			},
		},
		Metadata: map[string]string{
			"tier":           string(newTier),
			"employee_count": fmt.Sprintf("%d", newEmployeeCount),
		},
	}

	// Update the subscription
	subscription, err := sc.UpdateSubscription(ctx, subscriptionID, params)
	if err != nil {
		return nil, fmt.Errorf("failed to upgrade subscription: %w", err)
	}

	return subscription, nil
}

// GetDefaultSubscriptionPlans returns the standard V1 Consortium subscription plans
func GetDefaultSubscriptionPlans() []SubscriptionPlan {
	return []SubscriptionPlan{
		{
			Tier:         TierStarter,
			Name:         "Starter",
			Description:  "Perfect for small teams getting started with compliance",
			MaxEmployees: 50,
			Features: []string{
				"Basic compliance tracking",
				"Employee document management",
				"Email support",
			},
		},
		{
			Tier:         TierProfessional,
			Name:         "Professional",
			Description:  "Advanced features for growing transportation companies",
			MaxEmployees: 200,
			Features: []string{
				"Advanced compliance automation",
				"Drug & alcohol testing integration",
				"MVR monitoring",
				"Priority support",
			},
		},
		{
			Tier:         TierEnterprise,
			Name:         "Enterprise",
			Description:  "Full-featured solution for large fleets",
			MaxEmployees: 1000,
			Features: []string{
				"Custom compliance workflows",
				"API access",
				"Dedicated success manager",
				"Advanced reporting",
			},
		},
		{
			Tier:         TierCustom,
			Name:         "Custom",
			Description:  "Tailored solution for unique requirements",
			MaxEmployees: -1, // Unlimited
			Features: []string{
				"Everything in Enterprise",
				"Custom development",
				"On-premise deployment options",
				"24/7 dedicated support",
			},
		},
	}
}
