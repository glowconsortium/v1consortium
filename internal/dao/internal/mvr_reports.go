// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MvrReportsDao is the data access object for the table mvr_reports.
type MvrReportsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MvrReportsColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MvrReportsColumns defines and stores column names for the table mvr_reports.
type MvrReportsColumns struct {
	Id                    string //
	OrganizationId        string //
	UserId                string //
	Status                string //
	OrderedDate           string //
	OrderedBy             string //
	LicenseNumber         string //
	LicenseState          string //
	ExternalOrderId       string //
	ProviderName          string //
	ReportDate            string //
	ReportReceivedDate    string //
	RawReportData         string //
	TotalViolations       string //
	MajorViolations       string //
	MinorViolations       string //
	LicenseStatus         string //
	LicenseExpirationDate string //
	ReviewedBy            string //
	ReviewedDate          string //
	RequiresAction        string //
	ActionNotes           string //
	CreatedAt             string //
	UpdatedAt             string //
}

// mvrReportsColumns holds the columns for the table mvr_reports.
var mvrReportsColumns = MvrReportsColumns{
	Id:                    "id",
	OrganizationId:        "organization_id",
	UserId:                "user_id",
	Status:                "status",
	OrderedDate:           "ordered_date",
	OrderedBy:             "ordered_by",
	LicenseNumber:         "license_number",
	LicenseState:          "license_state",
	ExternalOrderId:       "external_order_id",
	ProviderName:          "provider_name",
	ReportDate:            "report_date",
	ReportReceivedDate:    "report_received_date",
	RawReportData:         "raw_report_data",
	TotalViolations:       "total_violations",
	MajorViolations:       "major_violations",
	MinorViolations:       "minor_violations",
	LicenseStatus:         "license_status",
	LicenseExpirationDate: "license_expiration_date",
	ReviewedBy:            "reviewed_by",
	ReviewedDate:          "reviewed_date",
	RequiresAction:        "requires_action",
	ActionNotes:           "action_notes",
	CreatedAt:             "created_at",
	UpdatedAt:             "updated_at",
}

// NewMvrReportsDao creates and returns a new DAO object for table data access.
func NewMvrReportsDao(handlers ...gdb.ModelHandler) *MvrReportsDao {
	return &MvrReportsDao{
		group:    "default",
		table:    "mvr_reports",
		columns:  mvrReportsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MvrReportsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MvrReportsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MvrReportsDao) Columns() MvrReportsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MvrReportsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MvrReportsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MvrReportsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
