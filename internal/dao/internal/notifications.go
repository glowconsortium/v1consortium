// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NotificationsDao is the data access object for the table notifications.
type NotificationsDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  NotificationsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// NotificationsColumns defines and stores column names for the table notifications.
type NotificationsColumns struct {
	Id                string //
	OrganizationId    string //
	UserId            string //
	Title             string //
	Message           string //
	NotificationType  string //
	Priority          string //
	SentAt            string //
	DeliveredAt       string //
	ReadAt            string //
	EmailAddress      string //
	PhoneNumber       string //
	ExternalMessageId string //
	TestId            string //
	MvrReportId       string //
	PhysicalId        string //
	WorkflowId        string //
	DeliveryAttempts  string //
	LastAttemptAt     string //
	DeliveryError     string //
	CreatedAt         string //
	UpdatedAt         string //
}

// notificationsColumns holds the columns for the table notifications.
var notificationsColumns = NotificationsColumns{
	Id:                "id",
	OrganizationId:    "organization_id",
	UserId:            "user_id",
	Title:             "title",
	Message:           "message",
	NotificationType:  "notification_type",
	Priority:          "priority",
	SentAt:            "sent_at",
	DeliveredAt:       "delivered_at",
	ReadAt:            "read_at",
	EmailAddress:      "email_address",
	PhoneNumber:       "phone_number",
	ExternalMessageId: "external_message_id",
	TestId:            "test_id",
	MvrReportId:       "mvr_report_id",
	PhysicalId:        "physical_id",
	WorkflowId:        "workflow_id",
	DeliveryAttempts:  "delivery_attempts",
	LastAttemptAt:     "last_attempt_at",
	DeliveryError:     "delivery_error",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

// NewNotificationsDao creates and returns a new DAO object for table data access.
func NewNotificationsDao(handlers ...gdb.ModelHandler) *NotificationsDao {
	return &NotificationsDao{
		group:    "default",
		table:    "notifications",
		columns:  notificationsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *NotificationsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *NotificationsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *NotificationsDao) Columns() NotificationsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *NotificationsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *NotificationsDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *NotificationsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
