// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BackgroundChecksDao is the data access object for the table background_checks.
type BackgroundChecksDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  BackgroundChecksColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// BackgroundChecksColumns defines and stores column names for the table background_checks.
type BackgroundChecksColumns struct {
	Id                        string //
	OrganizationId            string //
	UserId                    string //
	CheckType                 string //
	Status                    string //
	OrderedDate               string //
	OrderedBy                 string //
	ExternalOrderId           string //
	ProviderName              string //
	CompletedDate             string //
	ReportDate                string //
	OverallResult             string //
	RequiresReview            string //
	AdverseActionRequired     string //
	FcraDisclosureSent        string //
	FcraAuthorizationReceived string //
	PreAdverseActionSent      string //
	AdverseActionSent         string //
	Notes                     string //
	CreatedAt                 string //
	UpdatedAt                 string //
}

// backgroundChecksColumns holds the columns for the table background_checks.
var backgroundChecksColumns = BackgroundChecksColumns{
	Id:                        "id",
	OrganizationId:            "organization_id",
	UserId:                    "user_id",
	CheckType:                 "check_type",
	Status:                    "status",
	OrderedDate:               "ordered_date",
	OrderedBy:                 "ordered_by",
	ExternalOrderId:           "external_order_id",
	ProviderName:              "provider_name",
	CompletedDate:             "completed_date",
	ReportDate:                "report_date",
	OverallResult:             "overall_result",
	RequiresReview:            "requires_review",
	AdverseActionRequired:     "adverse_action_required",
	FcraDisclosureSent:        "fcra_disclosure_sent",
	FcraAuthorizationReceived: "fcra_authorization_received",
	PreAdverseActionSent:      "pre_adverse_action_sent",
	AdverseActionSent:         "adverse_action_sent",
	Notes:                     "notes",
	CreatedAt:                 "created_at",
	UpdatedAt:                 "updated_at",
}

// NewBackgroundChecksDao creates and returns a new DAO object for table data access.
func NewBackgroundChecksDao(handlers ...gdb.ModelHandler) *BackgroundChecksDao {
	return &BackgroundChecksDao{
		group:    "default",
		table:    "background_checks",
		columns:  backgroundChecksColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BackgroundChecksDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BackgroundChecksDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BackgroundChecksDao) Columns() BackgroundChecksColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BackgroundChecksDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BackgroundChecksDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BackgroundChecksDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
