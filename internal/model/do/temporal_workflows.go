// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TemporalWorkflows is the golang structure of table temporal_workflows for DAO operations like Where/Data.
type TemporalWorkflows struct {
	g.Meta         `orm:"table:temporal_workflows, do:true"`
	Id             interface{} //
	OrganizationId interface{} //
	WorkflowId     interface{} //
	WorkflowType   interface{} //
	Status         interface{} //
	InputData      interface{} //
	OutputData     interface{} //
	ErrorMessage   interface{} //
	UserId         interface{} //
	TestId         interface{} //
	SelectionId    interface{} //
	StartedAt      *gtime.Time //
	CompletedAt    *gtime.Time //
	ScheduledFor   *gtime.Time //
	RetryCount     interface{} //
	MaxRetries     interface{} //
	NextRetryAt    *gtime.Time //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
}
