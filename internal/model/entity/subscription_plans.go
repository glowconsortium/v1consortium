// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SubscriptionPlans is the golang structure for table subscription_plans.
type SubscriptionPlans struct {
	Id           string      `json:"id"           orm:"id"            description:""` //
	Name         string      `json:"name"         orm:"name"          description:""` //
	Tier         string      `json:"tier"         orm:"tier"          description:""` //
	Description  string      `json:"description"  orm:"description"   description:""` //
	MaxEmployees int         `json:"maxEmployees" orm:"max_employees" description:""` //
	AnnualPrice  float64     `json:"annualPrice"  orm:"annual_price"  description:""` //
	Features     string      `json:"features"     orm:"features"      description:""` //
	IsActive     bool        `json:"isActive"     orm:"is_active"     description:""` //
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:""` //
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:""` //
}
