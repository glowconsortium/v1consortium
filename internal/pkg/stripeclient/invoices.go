package stripeclient

import (
	"context"
	"time"

	"github.com/stripe/stripe-go/v83"
)

// GetInvoice retrieves a specific invoice by ID
func (sc *StripeClient) GetInvoice(ctx context.Context, invoiceID string) (*InvoiceData, error) {
	invoice, err := sc.client.V1Invoices.Retrieve(ctx, invoiceID, nil)
	if err != nil {
		return nil, WrapStripeError(err)
	}

	return convertToInvoiceData(invoice), nil
}

// ListCustomerInvoices retrieves all invoices for a customer
func (sc *StripeClient) ListCustomerInvoices(ctx context.Context, customerID string, limit int64) ([]*InvoiceData, error) {
	params := &stripe.InvoiceListParams{
		Customer: stripe.String(customerID),
	}
	params.Limit = stripe.Int64(limit)

	iter := sc.client.V1Invoices.List(ctx, params)
	var invoices []*InvoiceData

	// Use the All() method and range over it
	for invoice, err := range iter {
		if err != nil {
			return nil, WrapStripeError(err)
		}
		invoices = append(invoices, convertToInvoiceData(invoice))
	}

	return invoices, nil
}

// CreateInvoice creates a new invoice for a customer
func (sc *StripeClient) CreateInvoice(ctx context.Context, customerID string, description string, autoAdvance bool) (*InvoiceData, error) {
	params := &stripe.InvoiceCreateParams{
		Customer:    stripe.String(customerID),
		Description: stripe.String(description),
		AutoAdvance: stripe.Bool(autoAdvance),
	}

	invoice, err := sc.client.V1Invoices.Create(ctx, params)
	if err != nil {
		return nil, WrapStripeError(err)
	}

	return convertToInvoiceData(invoice), nil
}

// FinalizeInvoice finalizes a draft invoice
func (sc *StripeClient) FinalizeInvoice(ctx context.Context, invoiceID string) (*InvoiceData, error) {
	params := &stripe.InvoiceFinalizeInvoiceParams{}

	invoice, err := sc.client.V1Invoices.FinalizeInvoice(ctx, invoiceID, params)
	if err != nil {
		return nil, WrapStripeError(err)
	}

	return convertToInvoiceData(invoice), nil
}

// PayInvoice pays an invoice using the customer's default payment method
func (sc *StripeClient) PayInvoice(ctx context.Context, invoiceID string, paymentMethodID string) (*InvoiceData, error) {
	params := &stripe.InvoicePayParams{}
	if paymentMethodID != "" {
		params.PaymentMethod = stripe.String(paymentMethodID)
	}

	invoice, err := sc.client.V1Invoices.Pay(ctx, invoiceID, params)
	if err != nil {
		return nil, WrapStripeError(err)
	}

	return convertToInvoiceData(invoice), nil
}

// VoidInvoice voids an invoice (mark as uncollectible)
func (sc *StripeClient) VoidInvoice(ctx context.Context, invoiceID string) (*InvoiceData, error) {
	params := &stripe.InvoiceVoidInvoiceParams{}

	invoice, err := sc.client.V1Invoices.VoidInvoice(ctx, invoiceID, params)
	if err != nil {
		return nil, WrapStripeError(err)
	}

	return convertToInvoiceData(invoice), nil
}

// SendInvoice sends an invoice to the customer
func (sc *StripeClient) SendInvoice(ctx context.Context, invoiceID string) (*InvoiceData, error) {
	params := &stripe.InvoiceSendInvoiceParams{}

	invoice, err := sc.client.V1Invoices.SendInvoice(ctx, invoiceID, params)
	if err != nil {
		return nil, WrapStripeError(err)
	}

	return convertToInvoiceData(invoice), nil
}

// GetUpcomingInvoice previews an upcoming invoice for a subscription or customer
func (sc *StripeClient) GetUpcomingInvoice(ctx context.Context, subscriptionID string) (*InvoiceData, error) {
	params := &stripe.InvoiceCreatePreviewParams{
		Subscription: stripe.String(subscriptionID),
	}

	invoice, err := sc.client.V1Invoices.CreatePreview(ctx, params)
	if err != nil {
		return nil, WrapStripeError(err)
	}

	return convertToInvoiceData(invoice), nil
}

// CreateInvoiceWithLineItems creates an invoice with specific line items
func (sc *StripeClient) CreateInvoiceWithLineItems(ctx context.Context, customerID string, lineItems []*InvoiceLineItemData) (*InvoiceData, error) {
	// First create the invoice
	invoice, err := sc.CreateInvoice(ctx, customerID, "V1 Consortium Services", false)
	if err != nil {
		return nil, err
	}

	// Add line items to the invoice
	for _, item := range lineItems {
		itemParams := &stripe.InvoiceItemCreateParams{
			Customer:    stripe.String(customerID),
			Invoice:     stripe.String(invoice.ID),
			Amount:      stripe.Int64(item.Amount),
			Currency:    stripe.String(string(item.Currency)),
			Description: stripe.String(item.Description),
		}

		if item.Quantity > 0 {
			itemParams.Quantity = stripe.Int64(item.Quantity)
		}

		_, err := sc.client.V1InvoiceItems.Create(ctx, itemParams)
		if err != nil {
			return nil, WrapStripeError(err)
		}
	}

	// Return the updated invoice
	return sc.GetInvoice(ctx, invoice.ID)
}

// convertToInvoiceData converts a Stripe Invoice to our InvoiceData struct
func convertToInvoiceData(invoice *stripe.Invoice) *InvoiceData {
	var customerID string
	if invoice.Customer != nil {
		customerID = invoice.Customer.ID
	}

	// Note: Invoice doesn't have direct Subscription field in v83
	// If we need subscription info, we need to get it from the subscription itself
	var subscriptionID string
	// For now, leave empty since there's no direct Subscription field on Invoice

	// Convert unix timestamps to time.Time
	var dueDate, paidAt *time.Time
	if invoice.DueDate > 0 {
		t := time.Unix(invoice.DueDate, 0)
		dueDate = &t
	}
	if invoice.StatusTransitions != nil && invoice.StatusTransitions.PaidAt > 0 {
		t := time.Unix(invoice.StatusTransitions.PaidAt, 0)
		paidAt = &t
	}

	return &InvoiceData{
		ID:               invoice.ID,
		CustomerID:       customerID,
		SubscriptionID:   subscriptionID,
		Status:           string(invoice.Status),
		AmountTotal:      invoice.Total,
		AmountDue:        invoice.AmountDue,
		AmountPaid:       invoice.AmountPaid,
		Currency:         string(invoice.Currency),
		DueDate:          dueDate,
		PaidAt:           paidAt,
		HostedInvoiceURL: invoice.HostedInvoiceURL,
		InvoicePDF:       invoice.InvoicePDF,
		Metadata:         invoice.Metadata,
	}
}
