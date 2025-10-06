// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RandomSelections is the golang structure of table random_selections for DAO operations like Where/Data.
type RandomSelections struct {
	g.Meta             `orm:"table:random_selections, do:true"`
	Id                 interface{} //
	PoolId             interface{} //
	SelectionDate      *gtime.Time //
	SelectionPeriod    interface{} //
	TotalPoolSize      interface{} //
	RequiredSelections interface{} //
	SelectionAlgorithm interface{} //
	SelectionSeed      interface{} //
	Notes              interface{} //
	CreatedAt          *gtime.Time //
	CreatedBy          interface{} //
}
