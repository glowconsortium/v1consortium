// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Notifications is the golang structure of table notifications for DAO operations like Where/Data.
type Notifications struct {
	g.Meta            `orm:"table:notifications, do:true"`
	Id                interface{} //
	OrganizationId    interface{} //
	UserId            interface{} //
	Title             interface{} //
	Message           interface{} //
	NotificationType  interface{} //
	Priority          interface{} //
	SentAt            *gtime.Time //
	DeliveredAt       *gtime.Time //
	ReadAt            *gtime.Time //
	EmailAddress      interface{} //
	PhoneNumber       interface{} //
	ExternalMessageId interface{} //
	TestId            interface{} //
	MvrReportId       interface{} //
	PhysicalId        interface{} //
	WorkflowId        interface{} //
	DeliveryAttempts  interface{} //
	LastAttemptAt     *gtime.Time //
	DeliveryError     interface{} //
	CreatedAt         *gtime.Time //
	UpdatedAt         *gtime.Time //
}
