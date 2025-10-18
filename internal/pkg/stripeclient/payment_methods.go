package stripeclient

import (
	"context"

	"github.com/stripe/stripe-go/v83"
)

// CreateSetupIntent creates a setup intent for saving payment methods
func (sc *StripeClient) CreateSetupIntent(ctx context.Context, customerID string, usage string, metadata map[string]string) (*SetupIntentData, error) {
	params := &stripe.SetupIntentCreateParams{
		Customer: stripe.String(customerID),
		Usage:    stripe.String(usage), // "on_session" or "off_session"
		Metadata: metadata,
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
	}

	setupIntent, err := sc.client.V1SetupIntents.Create(ctx, params)
	if err != nil {
		return nil, WrapStripeError(err)
	}

	return &SetupIntentData{
		ID:           setupIntent.ID,
		ClientSecret: setupIntent.ClientSecret,
		Status:       string(setupIntent.Status),
		CustomerID:   customerID,
		Usage:        usage,
		Metadata:     metadata,
	}, nil
}

// ConfirmSetupIntent confirms a setup intent
func (sc *StripeClient) ConfirmSetupIntent(ctx context.Context, setupIntentID string, paymentMethodID string) (*SetupIntentData, error) {
	params := &stripe.SetupIntentConfirmParams{
		PaymentMethod: stripe.String(paymentMethodID),
	}

	setupIntent, err := sc.client.V1SetupIntents.Confirm(ctx, setupIntentID, params)
	if err != nil {
		return nil, WrapStripeError(err)
	}

	var customerID string
	if setupIntent.Customer != nil {
		customerID = setupIntent.Customer.ID
	}

	var paymentMethodIDResult string
	if setupIntent.PaymentMethod != nil {
		paymentMethodIDResult = setupIntent.PaymentMethod.ID
	}

	return &SetupIntentData{
		ID:              setupIntent.ID,
		ClientSecret:    setupIntent.ClientSecret,
		Status:          string(setupIntent.Status),
		CustomerID:      customerID,
		PaymentMethodID: paymentMethodIDResult,
		Metadata:        setupIntent.Metadata,
	}, nil
}

// ListCustomerPaymentMethods lists all payment methods for a customer
func (sc *StripeClient) ListCustomerPaymentMethods(ctx context.Context, customerID string) ([]*PaymentMethodData, error) {
	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String(customerID),
		Type:     stripe.String("card"),
	}

	var paymentMethods []*PaymentMethodData

	for pm, err := range sc.client.V1PaymentMethods.List(ctx, params) {
		if err != nil {
			return nil, WrapStripeError(err)
		}

		var customerIDResult string
		if pm.Customer != nil {
			customerIDResult = pm.Customer.ID
		}

		paymentMethods = append(paymentMethods, &PaymentMethodData{
			ID:         pm.ID,
			Type:       string(pm.Type),
			CustomerID: customerIDResult,
			Card:       pm.Card,
			Metadata:   pm.Metadata,
		})
	}

	return paymentMethods, nil
}

// DetachPaymentMethod detaches a payment method from a customer
func (sc *StripeClient) DetachPaymentMethod(ctx context.Context, paymentMethodID string) error {
	_, err := sc.client.V1PaymentMethods.Detach(ctx, paymentMethodID, nil)
	if err != nil {
		return WrapStripeError(err)
	}
	return nil
}

// SetDefaultPaymentMethod sets a payment method as the customer's default
func (sc *StripeClient) SetDefaultPaymentMethod(ctx context.Context, customerID, paymentMethodID string) error {
	params := &stripe.CustomerUpdateParams{
		InvoiceSettings: &stripe.CustomerUpdateInvoiceSettingsParams{
			DefaultPaymentMethod: stripe.String(paymentMethodID),
		},
	}

	_, err := sc.client.V1Customers.Update(ctx, customerID, params)
	if err != nil {
		return WrapStripeError(err)
	}
	return nil
}
