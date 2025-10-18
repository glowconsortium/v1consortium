// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"v1consortium/internal/model/do"
	"v1consortium/internal/model/entity"
)

type (
	IOrganizationService interface {
		GetOrganization(ctx context.Context, orgnanizationid string) (*entity.Organizations, error)
		CreateOrganization(ctx context.Context, createdata *do.Organizations) (*entity.Organizations, error)
		UpdateOrganization(ctx context.Context, id string, updatedata *do.Organizations) (*entity.Organizations, error)
		DeleteOrganization(ctx context.Context, id string) error
		ListOrganizations(ctx context.Context, offset int, limit int) ([]*entity.Organizations, error)
		DeactivateOrganization(ctx context.Context, id string) error
		CreateOrganizationSubscription(ctx context.Context, createdata *do.OrganizationSubscriptions) (*entity.OrganizationSubscriptions, error)
		GetOrganizationSubscription(ctx context.Context, id string) (*entity.OrganizationSubscriptions, error)
		GetOrganizationSubscriptionByOrganizationID(ctx context.Context, organizationID string) (*entity.OrganizationSubscriptions, error)
		ListOrganizationSubscriptions(ctx context.Context, offset int, limit int) ([]*entity.OrganizationSubscriptions, error)
		DeactivateOrganizationSubscription(ctx context.Context, id string) error
		UpdateOrganizationSubscription(ctx context.Context, id string, updatedata *do.OrganizationSubscriptions) (*entity.OrganizationSubscriptions, error)
		// get plan by tier
		GetPlanByTier(ctx context.Context, tier string) (*entity.SubscriptionPlans, error)
	}
)

var (
	localOrganizationService IOrganizationService
)

func OrganizationService() IOrganizationService {
	if localOrganizationService == nil {
		panic("implement not found for interface IOrganizationService, forgot register?")
	}
	return localOrganizationService
}

func RegisterOrganizationService(i IOrganizationService) {
	localOrganizationService = i
}
