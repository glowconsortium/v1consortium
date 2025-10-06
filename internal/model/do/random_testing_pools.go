// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RandomTestingPools is the golang structure of table random_testing_pools for DAO operations like Where/Data.
type RandomTestingPools struct {
	g.Meta         `orm:"table:random_testing_pools, do:true"`
	Id             interface{} //
	OrganizationId interface{} //
	ProgramId      interface{} //
	Name           interface{} //
	Description    interface{} //
	IsActive       interface{} //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
}
