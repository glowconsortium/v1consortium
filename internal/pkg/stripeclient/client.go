package stripeclient

import (
	"github.com/stripe/stripe-go/v83"
)

// StripeClient wraps the Stripe client with our business logic
type StripeClient struct {
	client *stripe.Client
	config *Config
}

// NewStripeClient creates a new Stripe client instance
func NewStripeClient(config *Config) *StripeClient {
	// Set up Stripe configuration
	stripe.Key = config.SecretKey

	// Configure backend settings
	backendConfig := &stripe.BackendConfig{
		MaxNetworkRetries: stripe.Int64(config.MaxNetworkRetries),
	}

	if config.MaxNetworkRetries == 0 {
		backendConfig.MaxNetworkRetries = stripe.Int64(2) // Default retries
	}

	backends := stripe.NewBackends(nil)
	if backendConfig.MaxNetworkRetries != nil {
		backends = stripe.NewBackendsWithConfig(backendConfig)
	}

	client := stripe.NewClient(config.SecretKey, stripe.WithBackends(backends))

	return &StripeClient{
		client: client,
		config: config,
	}
}
