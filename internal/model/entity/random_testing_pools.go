// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RandomTestingPools is the golang structure for table random_testing_pools.
type RandomTestingPools struct {
	Id             string      `json:"id"             orm:"id"              description:""` //
	OrganizationId string      `json:"organizationId" orm:"organization_id" description:""` //
	ProgramId      string      `json:"programId"      orm:"program_id"      description:""` //
	Name           string      `json:"name"           orm:"name"            description:""` //
	Description    string      `json:"description"    orm:"description"     description:""` //
	IsActive       bool        `json:"isActive"       orm:"is_active"       description:""` //
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:""` //
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:""` //
}
