// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RandomSelections is the golang structure for table random_selections.
type RandomSelections struct {
	Id                 string      `json:"id"                 orm:"id"                  description:""` //
	PoolId             string      `json:"poolId"             orm:"pool_id"             description:""` //
	SelectionDate      *gtime.Time `json:"selectionDate"      orm:"selection_date"      description:""` //
	SelectionPeriod    string      `json:"selectionPeriod"    orm:"selection_period"    description:""` //
	TotalPoolSize      int         `json:"totalPoolSize"      orm:"total_pool_size"     description:""` //
	RequiredSelections int         `json:"requiredSelections" orm:"required_selections" description:""` //
	SelectionAlgorithm string      `json:"selectionAlgorithm" orm:"selection_algorithm" description:""` //
	SelectionSeed      string      `json:"selectionSeed"      orm:"selection_seed"      description:""` //
	Notes              string      `json:"notes"              orm:"notes"               description:""` //
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"          description:""` //
	CreatedBy          string      `json:"createdBy"          orm:"created_by"          description:""` //
}
