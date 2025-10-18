package stripeclient

import (
	"time"

	"github.com/stripe/stripe-go/v83"
)

// SubscriptionTier represents the V1 Consortium subscription tiers
type SubscriptionTier string

const (
	TierStarter      SubscriptionTier = "starter"
	TierProfessional SubscriptionTier = "professional"
	TierEnterprise   SubscriptionTier = "enterprise"
	TierCustom       SubscriptionTier = "custom"
)

// Config holds configuration for the Stripe client
type Config struct {
	SecretKey         string
	PublishableKey    string
	WebhookSecret     string
	DefaultCurrency   string
	TestMode          bool
	MaxNetworkRetries int64
	APIVersion        string
}

// CustomerData represents customer information for Stripe
type CustomerData struct {
	ID          string
	Email       string
	Name        string
	Description string
	Phone       string
	Metadata    map[string]string
	Address     *stripe.Address
}

// PaymentMethodData represents payment method information
type PaymentMethodData struct {
	ID         string
	Type       string
	CustomerID string
	Card       *stripe.PaymentMethodCard
	IsDefault  bool
	Metadata   map[string]string
}

// SetupIntentData represents setup intent information for saving payment methods
type SetupIntentData struct {
	ID              string
	ClientSecret    string
	Status          string
	CustomerID      string
	PaymentMethodID string
	Usage           string
	Metadata        map[string]string
}

// SubscriptionData represents subscription information
type SubscriptionData struct {
	ID                     string
	CustomerID             string
	Status                 string
	CurrentPeriodStart     time.Time
	CurrentPeriodEnd       time.Time
	CancelAtPeriodEnd      bool
	TrialEnd               *time.Time
	DefaultPaymentMethodID string
	Items                  []*SubscriptionItemData
	Metadata               map[string]string
}

// SubscriptionItemData represents subscription item details
type SubscriptionItemData struct {
	ID       string
	PriceID  string
	Quantity int64
}

// InvoiceData represents invoice information
type InvoiceData struct {
	ID               string
	CustomerID       string
	SubscriptionID   string
	Status           string
	AmountTotal      int64
	AmountDue        int64
	AmountPaid       int64
	Currency         string
	DueDate          *time.Time
	PaidAt           *time.Time
	PaymentMethodID  string
	HostedInvoiceURL string
	InvoicePDF       string
	Metadata         map[string]string
}

// InvoiceLineItemData represents an invoice line item
type InvoiceLineItemData struct {
	Amount      int64
	Currency    string
	Description string
	Quantity    int64
}

// CheckoutSessionData represents a Stripe checkout session
type CheckoutSessionData struct {
	ID            string `json:"id"`
	CustomerID    string `json:"customer_id"`
	PaymentStatus string `json:"payment_status"`
	URL           string `json:"url"`
	SuccessURL    string `json:"success_url"`
	CancelURL     string `json:"cancel_url"`
	ExpiresAt     int64  `json:"expires_at"`
	AmountTotal   int64  `json:"amount_total"`
	Currency      string `json:"currency"`
}

// WebhookEvent represents a Stripe webhook event
type WebhookEvent struct {
	ID      string                 `json:"id"`
	Type    string                 `json:"type"`
	Data    map[string]interface{} `json:"data"`
	Created time.Time              `json:"created"`
}

// SubscriptionPlan represents a V1 Consortium subscription plan
type SubscriptionPlan struct {
	Tier         SubscriptionTier `json:"tier"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	MaxEmployees int64            `json:"max_employees"`
	Features     []string         `json:"features"`
}
