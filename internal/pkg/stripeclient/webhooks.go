package stripeclient

import (
	"net/http"
	"time"

	"github.com/stripe/stripe-go/v83/webhook"
)

// ValidateWebhook validates a Stripe webhook payload and signature
func (sc *StripeClient) ValidateWebhook(payload []byte, signatureHeader string) (*WebhookEvent, error) {
	if sc.config.WebhookSecret == "" {
		return nil, ErrInvalidWebhookSignature
	}

	event, err := webhook.ConstructEvent(payload, signatureHeader, sc.config.WebhookSecret)
	if err != nil {
		return nil, WrapStripeError(err)
	}

	return &WebhookEvent{
		ID:      event.ID,
		Type:    string(event.Type),
		Data:    make(map[string]interface{}), // Simplified for now
		Created: time.Unix(event.Created, 0),
	}, nil
}

// ParseWebhookFromRequest extracts and validates a webhook from an HTTP request
func (sc *StripeClient) ParseWebhookFromRequest(req *http.Request) (*WebhookEvent, error) {
	payload, err := getRequestBody(req)
	if err != nil {
		return nil, err
	}

	signatureHeader := req.Header.Get("Stripe-Signature")
	if signatureHeader == "" {
		return nil, ErrInvalidWebhookSignature
	}

	return sc.ValidateWebhook(payload, signatureHeader)
}

// Helper function to read request body
func getRequestBody(req *http.Request) ([]byte, error) {
	defer req.Body.Close()

	// Read up to 1MB of data
	const maxBodySize = 1024 * 1024
	body := make([]byte, maxBodySize)
	n, err := req.Body.Read(body)
	if err != nil && err.Error() != "EOF" {
		return nil, err
	}

	return body[:n], nil
}

// Common webhook event handlers
func (sc *StripeClient) HandleCustomerSubscriptionUpdated(event *WebhookEvent) error {
	// Implementation depends on your business logic
	// This is a placeholder for subscription update handling
	return nil
}

func (sc *StripeClient) HandleInvoicePaymentSucceeded(event *WebhookEvent) error {
	// Implementation depends on your business logic
	// This is a placeholder for successful payment handling
	return nil
}

func (sc *StripeClient) HandleInvoicePaymentFailed(event *WebhookEvent) error {
	// Implementation depends on your business logic
	// This is a placeholder for failed payment handling
	return nil
}
