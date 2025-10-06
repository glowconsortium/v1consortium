// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PoolMemberships is the golang structure of table pool_memberships for DAO operations like Where/Data.
type PoolMemberships struct {
	g.Meta   `orm:"table:pool_memberships, do:true"`
	Id       interface{} //
	PoolId   interface{} //
	UserId   interface{} //
	JoinedAt *gtime.Time //
	LeftAt   *gtime.Time //
	IsActive interface{} //
}
