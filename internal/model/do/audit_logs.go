// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuditLogs is the golang structure of table audit_logs for DAO operations like Where/Data.
type AuditLogs struct {
	g.Meta            `orm:"table:audit_logs, do:true"`
	Id                interface{} //
	OrganizationId    interface{} //
	UserId            interface{} //
	Action            interface{} //
	EntityType        interface{} //
	EntityId          interface{} //
	OldValues         interface{} //
	NewValues         interface{} //
	IpAddress         interface{} //
	UserAgent         interface{} //
	RequestId         interface{} //
	RetentionRequired interface{} //
	HipaaLog          interface{} //
	CreatedAt         *gtime.Time //
}
