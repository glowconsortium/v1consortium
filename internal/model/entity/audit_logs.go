// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AuditLogs is the golang structure for table audit_logs.
type AuditLogs struct {
	Id                string      `json:"id"                orm:"id"                 description:""` //
	OrganizationId    string      `json:"organizationId"    orm:"organization_id"    description:""` //
	UserId            string      `json:"userId"            orm:"user_id"            description:""` //
	Action            string      `json:"action"            orm:"action"             description:""` //
	EntityType        string      `json:"entityType"        orm:"entity_type"        description:""` //
	EntityId          string      `json:"entityId"          orm:"entity_id"          description:""` //
	OldValues         string      `json:"oldValues"         orm:"old_values"         description:""` //
	NewValues         string      `json:"newValues"         orm:"new_values"         description:""` //
	IpAddress         string      `json:"ipAddress"         orm:"ip_address"         description:""` //
	UserAgent         string      `json:"userAgent"         orm:"user_agent"         description:""` //
	RequestId         string      `json:"requestId"         orm:"request_id"         description:""` //
	RetentionRequired bool        `json:"retentionRequired" orm:"retention_required" description:""` //
	HipaaLog          bool        `json:"hipaaLog"          orm:"hipaa_log"          description:""` //
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"         description:""` //
}
