// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PoolMemberships is the golang structure for table pool_memberships.
type PoolMemberships struct {
	Id       string      `json:"id"       orm:"id"        description:""` //
	PoolId   string      `json:"poolId"   orm:"pool_id"   description:""` //
	UserId   string      `json:"userId"   orm:"user_id"   description:""` //
	JoinedAt *gtime.Time `json:"joinedAt" orm:"joined_at" description:""` //
	LeftAt   *gtime.Time `json:"leftAt"   orm:"left_at"   description:""` //
	IsActive bool        `json:"isActive" orm:"is_active" description:""` //
}
