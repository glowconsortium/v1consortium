package stripeclient

import (
	"context"

	"github.com/stripe/stripe-go/v83"
)

// CreateCheckoutSession creates a Stripe Checkout session for one-time payments
func (sc *StripeClient) CreateCheckoutSession(ctx context.Context, customerID string, priceID string, successURL, cancelURL string) (*CheckoutSessionData, error) {
	params := &stripe.CheckoutSessionCreateParams{
		Customer: stripe.String(customerID),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems: []*stripe.CheckoutSessionCreateLineItemParams{
			{
				Price:    stripe.String(priceID),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(successURL),
		CancelURL:  stripe.String(cancelURL),
	}

	session, err := sc.client.V1CheckoutSessions.Create(ctx, params)
	if err != nil {
		return nil, WrapStripeError(err)
	}

	return &CheckoutSessionData{
		ID:            session.ID,
		CustomerID:    customerID,
		PaymentStatus: string(session.PaymentStatus),
		URL:           session.URL,
		SuccessURL:    successURL,
		CancelURL:     cancelURL,
		ExpiresAt:     session.ExpiresAt,
		AmountTotal:   session.AmountTotal,
		Currency:      string(session.Currency),
	}, nil
}

// CreateSubscriptionCheckoutSession creates a Stripe Checkout session for subscription payments
func (sc *StripeClient) CreateSubscriptionCheckoutSession(ctx context.Context, customerID string, priceID string, successURL, cancelURL string, trialPeriodDays int64) (*CheckoutSessionData, error) {
	params := &stripe.CheckoutSessionCreateParams{
		Customer: stripe.String(customerID),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems: []*stripe.CheckoutSessionCreateLineItemParams{
			{
				Price:    stripe.String(priceID),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		SuccessURL: stripe.String(successURL),
		CancelURL:  stripe.String(cancelURL),
	}

	// Add trial period if specified
	if trialPeriodDays > 0 {
		params.SubscriptionData = &stripe.CheckoutSessionCreateSubscriptionDataParams{
			TrialPeriodDays: stripe.Int64(trialPeriodDays),
		}
	}

	session, err := sc.client.V1CheckoutSessions.Create(ctx, params)
	if err != nil {
		return nil, WrapStripeError(err)
	}

	return &CheckoutSessionData{
		ID:            session.ID,
		CustomerID:    customerID,
		PaymentStatus: string(session.PaymentStatus),
		URL:           session.URL,
		SuccessURL:    successURL,
		CancelURL:     cancelURL,
		ExpiresAt:     session.ExpiresAt,
		AmountTotal:   session.AmountTotal,
		Currency:      string(session.Currency),
	}, nil
}

// GetCheckoutSession retrieves a checkout session by ID
func (sc *StripeClient) GetCheckoutSession(ctx context.Context, sessionID string) (*CheckoutSessionData, error) {
	session, err := sc.client.V1CheckoutSessions.Retrieve(ctx, sessionID, nil)
	if err != nil {
		return nil, WrapStripeError(err)
	}

	var customerID string
	if session.Customer != nil {
		customerID = session.Customer.ID
	}

	return &CheckoutSessionData{
		ID:            session.ID,
		CustomerID:    customerID,
		PaymentStatus: string(session.PaymentStatus),
		URL:           session.URL,
		ExpiresAt:     session.ExpiresAt,
		AmountTotal:   session.AmountTotal,
		Currency:      string(session.Currency),
	}, nil
}
