// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrganizationSubscriptions is the golang structure for table organization_subscriptions.
type OrganizationSubscriptions struct {
	Id                   string      `json:"id"                   orm:"id"                     description:""` //
	OrganizationId       string      `json:"organizationId"       orm:"organization_id"        description:""` //
	PlanId               string      `json:"planId"               orm:"plan_id"                description:""` //
	Status               string      `json:"status"               orm:"status"                 description:""` //
	StartDate            *gtime.Time `json:"startDate"            orm:"start_date"             description:""` //
	EndDate              *gtime.Time `json:"endDate"              orm:"end_date"               description:""` //
	AutoRenew            bool        `json:"autoRenew"            orm:"auto_renew"             description:""` //
	StripeSubscriptionId string      `json:"stripeSubscriptionId" orm:"stripe_subscription_id" description:""` //
	StripeCustomerId     string      `json:"stripeCustomerId"     orm:"stripe_customer_id"     description:""` //
	CreatedAt            *gtime.Time `json:"createdAt"            orm:"created_at"             description:""` //
	UpdatedAt            *gtime.Time `json:"updatedAt"            orm:"updated_at"             description:""` //
}
