// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SavedReportsDao is the data access object for the table saved_reports.
type SavedReportsDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  SavedReportsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// SavedReportsColumns defines and stores column names for the table saved_reports.
type SavedReportsColumns struct {
	Id                string //
	OrganizationId    string //
	CreatedBy         string //
	Name              string //
	Description       string //
	ReportType        string //
	Parameters        string //
	Filters           string //
	GeneratedAt       string //
	FilePath          string //
	FileFormat        string //
	IsScheduled       string //
	ScheduleFrequency string //
	NextRunDate       string //
	IsPublic          string //
	SharedWith        string //
	CreatedAt         string //
	UpdatedAt         string //
}

// savedReportsColumns holds the columns for the table saved_reports.
var savedReportsColumns = SavedReportsColumns{
	Id:                "id",
	OrganizationId:    "organization_id",
	CreatedBy:         "created_by",
	Name:              "name",
	Description:       "description",
	ReportType:        "report_type",
	Parameters:        "parameters",
	Filters:           "filters",
	GeneratedAt:       "generated_at",
	FilePath:          "file_path",
	FileFormat:        "file_format",
	IsScheduled:       "is_scheduled",
	ScheduleFrequency: "schedule_frequency",
	NextRunDate:       "next_run_date",
	IsPublic:          "is_public",
	SharedWith:        "shared_with",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

// NewSavedReportsDao creates and returns a new DAO object for table data access.
func NewSavedReportsDao(handlers ...gdb.ModelHandler) *SavedReportsDao {
	return &SavedReportsDao{
		group:    "default",
		table:    "saved_reports",
		columns:  savedReportsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SavedReportsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SavedReportsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SavedReportsDao) Columns() SavedReportsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SavedReportsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SavedReportsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SavedReportsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
