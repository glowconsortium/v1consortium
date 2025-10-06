// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TestingProgramsDao is the data access object for the table testing_programs.
type TestingProgramsDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  TestingProgramsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// TestingProgramsColumns defines and stores column names for the table testing_programs.
type TestingProgramsColumns struct {
	Id                    string //
	OrganizationId        string //
	Name                  string //
	IsDotProgram          string //
	DrugPanelType         string //
	AlcoholTestingEnabled string //
	RandomTestingEnabled  string //
	RandomTestingRate     string //
	TestingFrequency      string //
	IsActive              string //
	CreatedAt             string //
	UpdatedAt             string //
}

// testingProgramsColumns holds the columns for the table testing_programs.
var testingProgramsColumns = TestingProgramsColumns{
	Id:                    "id",
	OrganizationId:        "organization_id",
	Name:                  "name",
	IsDotProgram:          "is_dot_program",
	DrugPanelType:         "drug_panel_type",
	AlcoholTestingEnabled: "alcohol_testing_enabled",
	RandomTestingEnabled:  "random_testing_enabled",
	RandomTestingRate:     "random_testing_rate",
	TestingFrequency:      "testing_frequency",
	IsActive:              "is_active",
	CreatedAt:             "created_at",
	UpdatedAt:             "updated_at",
}

// NewTestingProgramsDao creates and returns a new DAO object for table data access.
func NewTestingProgramsDao(handlers ...gdb.ModelHandler) *TestingProgramsDao {
	return &TestingProgramsDao{
		group:    "default",
		table:    "testing_programs",
		columns:  testingProgramsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TestingProgramsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TestingProgramsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TestingProgramsDao) Columns() TestingProgramsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TestingProgramsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TestingProgramsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TestingProgramsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
