// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"v1consortium/internal/pkg/stripeclient"
)

type (
	IStripeService interface {
		CreateCustomer(ctx context.Context, data *stripeclient.CustomerData) (*stripeclient.CustomerData, error)
		GetStripeClientConfig(ctx context.Context) *stripeclient.Config
	}
)

var (
	localStripeService IStripeService
)

func StripeService() IStripeService {
	if localStripeService == nil {
		panic("implement not found for interface IStripeService, forgot register?")
	}
	return localStripeService
}

func RegisterStripeService(i IStripeService) {
	localStripeService = i
}
