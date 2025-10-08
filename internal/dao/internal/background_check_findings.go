// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BackgroundCheckFindingsDao is the data access object for the table background_check_findings.
type BackgroundCheckFindingsDao struct {
	table    string                         // table is the underlying table name of the DAO.
	group    string                         // group is the database configuration group name of the current DAO.
	columns  BackgroundCheckFindingsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler             // handlers for customized model modification.
}

// BackgroundCheckFindingsColumns defines and stores column names for the table background_check_findings.
type BackgroundCheckFindingsColumns struct {
	Id                               string //
	BackgroundCheckId                string //
	FindingType                      string //
	Severity                         string //
	Description                      string //
	DateOfRecord                     string //
	Jurisdiction                     string //
	CaseNumber                       string //
	Disposition                      string //
	JobRelated                       string //
	Disqualifying                    string //
	RequiresIndividualizedAssessment string //
	CreatedAt                        string //
}

// backgroundCheckFindingsColumns holds the columns for the table background_check_findings.
var backgroundCheckFindingsColumns = BackgroundCheckFindingsColumns{
	Id:                               "id",
	BackgroundCheckId:                "background_check_id",
	FindingType:                      "finding_type",
	Severity:                         "severity",
	Description:                      "description",
	DateOfRecord:                     "date_of_record",
	Jurisdiction:                     "jurisdiction",
	CaseNumber:                       "case_number",
	Disposition:                      "disposition",
	JobRelated:                       "job_related",
	Disqualifying:                    "disqualifying",
	RequiresIndividualizedAssessment: "requires_individualized_assessment",
	CreatedAt:                        "created_at",
}

// NewBackgroundCheckFindingsDao creates and returns a new DAO object for table data access.
func NewBackgroundCheckFindingsDao(handlers ...gdb.ModelHandler) *BackgroundCheckFindingsDao {
	return &BackgroundCheckFindingsDao{
		group:    "default",
		table:    "background_check_findings",
		columns:  backgroundCheckFindingsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BackgroundCheckFindingsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BackgroundCheckFindingsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BackgroundCheckFindingsDao) Columns() BackgroundCheckFindingsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BackgroundCheckFindingsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BackgroundCheckFindingsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BackgroundCheckFindingsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
