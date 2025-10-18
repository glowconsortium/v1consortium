package stripeservice

import (
	"context"
	"v1consortium/internal/pkg/stripeclient"
	"v1consortium/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type sStripeService struct{}

func init() {
	service.RegisterStripeService(new())
}

func new() service.IStripeService {
	return &sStripeService{}
}

func (s *sStripeService) CreateCustomer(ctx context.Context, data *stripeclient.CustomerData) (*stripeclient.CustomerData, error) {
	stripeC := stripeclient.NewStripeClient(s.GetStripeClientConfig(ctx))
	return stripeC.CreateCustomer(ctx, data)
}

func (s *sStripeService) GetStripeClientConfig(ctx context.Context) *stripeclient.Config {
	secretKey := g.Cfg().MustGet(ctx, "stripe.secretKey").String()
	maxNetworkRetries := g.Cfg().MustGet(ctx, "stripe.maxNetworkRetries").Int64()
	webhookSecret := g.Cfg().MustGet(ctx, "stripe.webhookSecret").String()
	defaultConcurrency := g.Cfg().MustGet(ctx, "stripe.defaultCurrency").String()
	TestMode := g.Cfg().MustGet(ctx, "stripe.testMode").Bool()
	ApiVersion := g.Cfg().MustGet(ctx, "stripe.apiVersion").String()
	return &stripeclient.Config{
		SecretKey:         secretKey,
		MaxNetworkRetries: maxNetworkRetries,
		WebhookSecret:     webhookSecret,
		DefaultCurrency:   defaultConcurrency,
		TestMode:          TestMode,
		APIVersion:        ApiVersion,
	}
}
