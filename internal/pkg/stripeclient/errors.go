package stripeclient

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/stripe/stripe-go/v83"
)

// Known error variables for common Stripe operations
var (
	// Customer errors
	ErrCustomerNotFound      = errors.New("customer not found")
	ErrCustomerAlreadyExists = errors.New("customer already exists")
	ErrInvalidCustomerData   = errors.New("invalid customer data")
	ErrCustomerDeleted       = errors.New("customer has been deleted")

	// Payment Method errors
	ErrPaymentMethodNotFound          = errors.New("payment method not found")
	ErrPaymentMethodInvalid           = errors.New("payment method is invalid")
	ErrPaymentMethodDeclined          = errors.New("payment method was declined")
	ErrPaymentMethodExpired           = errors.New("payment method has expired")
	ErrPaymentMethodInsufficientFunds = errors.New("insufficient funds")
	ErrPaymentMethodNotAttached       = errors.New("payment method not attached to customer")

	// Subscription errors
	ErrSubscriptionNotFound    = errors.New("subscription not found")
	ErrSubscriptionCanceled    = errors.New("subscription has been canceled")
	ErrSubscriptionIncomplete  = errors.New("subscription is incomplete")
	ErrSubscriptionPastDue     = errors.New("subscription is past due")
	ErrInvalidSubscriptionTier = errors.New("invalid subscription tier")
	ErrTrialAlreadyUsed        = errors.New("trial period has already been used")

	// Invoice errors
	ErrInvoiceNotFound      = errors.New("invoice not found")
	ErrInvoiceAlreadyPaid   = errors.New("invoice has already been paid")
	ErrInvoiceVoid          = errors.New("invoice has been voided")
	ErrInvoicePaymentFailed = errors.New("invoice payment failed")

	// Webhook errors
	ErrInvalidWebhookSignature  = errors.New("invalid webhook signature")
	ErrWebhookEventNotSupported = errors.New("webhook event type not supported")
	ErrWebhookProcessingFailed  = errors.New("webhook processing failed")

	// General errors
	ErrInvalidAPIKey       = errors.New("invalid API key")
	ErrRateLimitExceeded   = errors.New("rate limit exceeded")
	ErrNetworkError        = errors.New("network error")
	ErrInternalServerError = errors.New("internal server error")
	ErrInvalidRequest      = errors.New("invalid request")
	ErrPermissionDenied    = errors.New("permission denied")
)

// StripeError wraps Stripe errors with additional context
type StripeError struct {
	Code       string
	Message    string
	Type       string
	HTTPStatus int
	RequestID  string
	Underlying error
}

func (e *StripeError) Error() string {
	if e.RequestID != "" {
		return fmt.Sprintf("stripe error [%s] %s (request: %s): %s", e.Code, e.Type, e.RequestID, e.Message)
	}
	return fmt.Sprintf("stripe error [%s] %s: %s", e.Code, e.Type, e.Message)
}

func (e *StripeError) Unwrap() error {
	return e.Underlying
}

// IsStripeError checks if an error is a StripeError
func IsStripeError(err error) bool {
	var stripeErr *StripeError
	return errors.As(err, &stripeErr)
}

// WrapStripeError converts a Stripe error to our custom error type with known error mapping
func WrapStripeError(err error) error {
	if err == nil {
		return nil
	}

	var stripeErr *stripe.Error
	if !errors.As(err, &stripeErr) {
		// Not a Stripe error, return as-is wrapped in a generic error
		return &StripeError{
			Code:       "unknown",
			Message:    err.Error(),
			Type:       "unknown_error",
			HTTPStatus: 0,
			Underlying: err,
		}
	}

	wrappedErr := &StripeError{
		Code:       string(stripeErr.Code),
		Message:    stripeErr.Msg,
		Type:       string(stripeErr.Type),
		HTTPStatus: stripeErr.HTTPStatusCode,
		RequestID:  stripeErr.RequestID,
		Underlying: err,
	}

	// Map specific Stripe error codes to our known errors
	switch stripeErr.Code {
	// Customer errors
	case "resource_missing":
		if stripeErr.Type == "invalid_request_error" {
			return fmt.Errorf("%w: %s", ErrCustomerNotFound, wrappedErr.Error())
		}
	case "email_invalid":
		return fmt.Errorf("%w: %s", ErrInvalidCustomerData, wrappedErr.Error())

	// Payment method errors
	case "card_declined":
		return fmt.Errorf("%w: %s", ErrPaymentMethodDeclined, wrappedErr.Error())
	case "expired_card":
		return fmt.Errorf("%w: %s", ErrPaymentMethodExpired, wrappedErr.Error())
	case "insufficient_funds":
		return fmt.Errorf("%w: %s", ErrPaymentMethodInsufficientFunds, wrappedErr.Error())
	case "invalid_cvc", "invalid_expiry_month", "invalid_expiry_year", "invalid_number":
		return fmt.Errorf("%w: %s", ErrPaymentMethodInvalid, wrappedErr.Error())

	// API errors
	case "invalid_api_key":
		return fmt.Errorf("%w: %s", ErrInvalidAPIKey, wrappedErr.Error())
	case "rate_limit":
		return fmt.Errorf("%w: %s", ErrRateLimitExceeded, wrappedErr.Error())

	// Permission errors
	case "account_invalid", "platform_account_required":
		return fmt.Errorf("%w: %s", ErrPermissionDenied, wrappedErr.Error())
	}

	// Map by HTTP status code for generic cases
	switch stripeErr.HTTPStatusCode {
	case http.StatusNotFound:
		switch stripeErr.Type {
		case "invalid_request_error":
			return fmt.Errorf("%w: %s", ErrCustomerNotFound, wrappedErr.Error())
		}
	case http.StatusTooManyRequests:
		return fmt.Errorf("%w: %s", ErrRateLimitExceeded, wrappedErr.Error())
	case http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable:
		return fmt.Errorf("%w: %s", ErrInternalServerError, wrappedErr.Error())
	case http.StatusBadRequest:
		return fmt.Errorf("%w: %s", ErrInvalidRequest, wrappedErr.Error())
	case http.StatusUnauthorized, http.StatusForbidden:
		return fmt.Errorf("%w: %s", ErrPermissionDenied, wrappedErr.Error())
	}

	// Return the wrapped error if no specific mapping found
	return wrappedErr
}

// IsKnownError checks if an error is one of our known error types
func IsKnownError(err error, target error) bool {
	return errors.Is(err, target)
}

// GetErrorCode extracts the Stripe error code from an error
func GetErrorCode(err error) string {
	var stripeErr *StripeError
	if errors.As(err, &stripeErr) {
		return stripeErr.Code
	}
	return ""
}

// GetRequestID extracts the Stripe request ID from an error
func GetRequestID(err error) string {
	var stripeErr *StripeError
	if errors.As(err, &stripeErr) {
		return stripeErr.RequestID
	}
	return ""
}
