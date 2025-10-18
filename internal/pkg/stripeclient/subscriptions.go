package stripeclient

import (
	"context"
	"time"

	"github.com/stripe/stripe-go/v83"
)

// CreateSubscription creates a new subscription for a customer
func (sc *StripeClient) CreateSubscription(ctx context.Context, customerID string, priceID string, trialPeriodDays int64, metadata map[string]string) (*SubscriptionData, error) {
	params := &stripe.SubscriptionCreateParams{
		Customer: stripe.String(customerID),
		Items: []*stripe.SubscriptionCreateItemParams{
			{
				Price: stripe.String(priceID),
			},
		},
		Metadata:        metadata,
		PaymentBehavior: stripe.String("default_incomplete"),
		PaymentSettings: &stripe.SubscriptionCreatePaymentSettingsParams{
			SaveDefaultPaymentMethod: stripe.String("on_subscription"),
		},
		Expand: stripe.StringSlice([]string{"latest_invoice.payment_intent"}),
	}

	if trialPeriodDays > 0 {
		params.TrialPeriodDays = stripe.Int64(trialPeriodDays)
	}

	sub, err := sc.client.V1Subscriptions.Create(ctx, params)
	if err != nil {
		return nil, WrapStripeError(err)
	}

	return sc.convertToSubscriptionData(sub), nil
}

// GetSubscription retrieves a subscription by ID
func (sc *StripeClient) GetSubscription(ctx context.Context, subscriptionID string) (*SubscriptionData, error) {
	sub, err := sc.client.V1Subscriptions.Retrieve(ctx, subscriptionID, nil)
	if err != nil {
		return nil, WrapStripeError(err)
	}

	return sc.convertToSubscriptionData(sub), nil
}

// UpdateSubscription updates an existing subscription
func (sc *StripeClient) UpdateSubscription(ctx context.Context, subscriptionID string, params *stripe.SubscriptionUpdateParams) (*SubscriptionData, error) {
	sub, err := sc.client.V1Subscriptions.Update(ctx, subscriptionID, params)
	if err != nil {
		return nil, WrapStripeError(err)
	}

	return sc.convertToSubscriptionData(sub), nil
}

// CancelSubscription cancels a subscription
func (sc *StripeClient) CancelSubscription(ctx context.Context, subscriptionID string, cancelAtPeriodEnd bool) (*SubscriptionData, error) {
	if cancelAtPeriodEnd {
		params := &stripe.SubscriptionUpdateParams{
			CancelAtPeriodEnd: stripe.Bool(true),
		}
		sub, err := sc.client.V1Subscriptions.Update(ctx, subscriptionID, params)
		if err != nil {
			return nil, WrapStripeError(err)
		}
		return sc.convertToSubscriptionData(sub), nil
	} else {
		cancelParams := &stripe.SubscriptionCancelParams{}
		sub, err := sc.client.V1Subscriptions.Cancel(ctx, subscriptionID, cancelParams)
		if err != nil {
			return nil, WrapStripeError(err)
		}
		return sc.convertToSubscriptionData(sub), nil
	}
}

// ListCustomerSubscriptions lists all subscriptions for a customer
func (sc *StripeClient) ListCustomerSubscriptions(ctx context.Context, customerID string) ([]*SubscriptionData, error) {
	params := &stripe.SubscriptionListParams{
		Customer: stripe.String(customerID),
	}

	var subscriptions []*SubscriptionData

	for sub, err := range sc.client.V1Subscriptions.List(ctx, params) {
		if err != nil {
			return nil, WrapStripeError(err)
		}
		subscriptions = append(subscriptions, sc.convertToSubscriptionData(sub))
	}

	return subscriptions, nil
}

// convertToSubscriptionData converts a Stripe subscription to our internal data structure
func (sc *StripeClient) convertToSubscriptionData(sub *stripe.Subscription) *SubscriptionData {
	var items []*SubscriptionItemData
	if sub.Items != nil {
		for _, item := range sub.Items.Data {
			items = append(items, &SubscriptionItemData{
				ID:       item.ID,
				PriceID:  item.Price.ID,
				Quantity: item.Quantity,
			})
		}
	}

	var defaultPaymentMethodID string
	if sub.DefaultPaymentMethod != nil {
		defaultPaymentMethodID = sub.DefaultPaymentMethod.ID
	}

	var trialEnd *time.Time
	if sub.TrialEnd != 0 {
		t := time.Unix(sub.TrialEnd, 0)
		trialEnd = &t
	}

	var customerID string
	if sub.Customer != nil {
		customerID = sub.Customer.ID
	}

	return &SubscriptionData{
		ID:                     sub.ID,
		CustomerID:             customerID,
		Status:                 string(sub.Status),
		CurrentPeriodStart:     calculatePeriodStart(sub),
		CurrentPeriodEnd:       calculatePeriodEnd(sub),
		CancelAtPeriodEnd:      sub.CancelAtPeriodEnd,
		TrialEnd:               trialEnd,
		DefaultPaymentMethodID: defaultPaymentMethodID,
		Items:                  items,
		Metadata:               sub.Metadata,
	}
}

// calculatePeriodStart determines the current period start date for a subscription
func calculatePeriodStart(sub *stripe.Subscription) time.Time {
	// Use the subscription start date if available
	if sub.StartDate > 0 {
		return time.Unix(sub.StartDate, 0)
	}

	// Fall back to billing cycle anchor
	if sub.BillingCycleAnchor > 0 {
		return time.Unix(sub.BillingCycleAnchor, 0)
	}

	// Ultimate fallback to creation date
	return time.Unix(sub.Created, 0)
}

// calculatePeriodEnd determines the current period end date for a subscription
func calculatePeriodEnd(sub *stripe.Subscription) time.Time {
	// If the subscription is ended, use that date
	if sub.EndedAt > 0 {
		return time.Unix(sub.EndedAt, 0)
	}

	// If there's a cancel date, use that
	if sub.CancelAt > 0 {
		return time.Unix(sub.CancelAt, 0)
	}

	// For V1 Consortium annual subscriptions, calculate next year from start
	periodStart := calculatePeriodStart(sub)
	return periodStart.AddDate(1, 0, 0) // Add 1 year for annual subscription
}
