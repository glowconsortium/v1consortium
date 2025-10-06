// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SubscriptionPlans is the golang structure of table subscription_plans for DAO operations like Where/Data.
type SubscriptionPlans struct {
	g.Meta       `orm:"table:subscription_plans, do:true"`
	Id           interface{} //
	Name         interface{} //
	Tier         interface{} //
	Description  interface{} //
	MaxEmployees interface{} //
	AnnualPrice  interface{} //
	Features     interface{} //
	IsActive     interface{} //
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
