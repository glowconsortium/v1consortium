// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Notifications is the golang structure for table notifications.
type Notifications struct {
	Id                string      `json:"id"                orm:"id"                  description:""` //
	OrganizationId    string      `json:"organizationId"    orm:"organization_id"     description:""` //
	UserId            string      `json:"userId"            orm:"user_id"             description:""` //
	Title             string      `json:"title"             orm:"title"               description:""` //
	Message           string      `json:"message"           orm:"message"             description:""` //
	NotificationType  string      `json:"notificationType"  orm:"notification_type"   description:""` //
	Priority          string      `json:"priority"          orm:"priority"            description:""` //
	SentAt            *gtime.Time `json:"sentAt"            orm:"sent_at"             description:""` //
	DeliveredAt       *gtime.Time `json:"deliveredAt"       orm:"delivered_at"        description:""` //
	ReadAt            *gtime.Time `json:"readAt"            orm:"read_at"             description:""` //
	EmailAddress      string      `json:"emailAddress"      orm:"email_address"       description:""` //
	PhoneNumber       string      `json:"phoneNumber"       orm:"phone_number"        description:""` //
	ExternalMessageId string      `json:"externalMessageId" orm:"external_message_id" description:""` //
	TestId            string      `json:"testId"            orm:"test_id"             description:""` //
	MvrReportId       string      `json:"mvrReportId"       orm:"mvr_report_id"       description:""` //
	PhysicalId        string      `json:"physicalId"        orm:"physical_id"         description:""` //
	WorkflowId        string      `json:"workflowId"        orm:"workflow_id"         description:""` //
	DeliveryAttempts  int         `json:"deliveryAttempts"  orm:"delivery_attempts"   description:""` //
	LastAttemptAt     *gtime.Time `json:"lastAttemptAt"     orm:"last_attempt_at"     description:""` //
	DeliveryError     string      `json:"deliveryError"     orm:"delivery_error"      description:""` //
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"          description:""` //
	UpdatedAt         *gtime.Time `json:"updatedAt"         orm:"updated_at"          description:""` //
}
