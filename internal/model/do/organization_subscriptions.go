// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrganizationSubscriptions is the golang structure of table organization_subscriptions for DAO operations like Where/Data.
type OrganizationSubscriptions struct {
	g.Meta               `orm:"table:organization_subscriptions, do:true"`
	Id                   interface{} //
	OrganizationId       interface{} //
	PlanId               interface{} //
	Status               interface{} //
	StartDate            *gtime.Time //
	EndDate              *gtime.Time //
	AutoRenew            interface{} //
	StripeSubscriptionId interface{} //
	StripeCustomerId     interface{} //
	CreatedAt            *gtime.Time //
	UpdatedAt            *gtime.Time //
}
