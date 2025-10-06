// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RandomSelectionsDao is the data access object for the table random_selections.
type RandomSelectionsDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  RandomSelectionsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// RandomSelectionsColumns defines and stores column names for the table random_selections.
type RandomSelectionsColumns struct {
	Id                 string //
	PoolId             string //
	SelectionDate      string //
	SelectionPeriod    string //
	TotalPoolSize      string //
	RequiredSelections string //
	SelectionAlgorithm string //
	SelectionSeed      string //
	Notes              string //
	CreatedAt          string //
	CreatedBy          string //
}

// randomSelectionsColumns holds the columns for the table random_selections.
var randomSelectionsColumns = RandomSelectionsColumns{
	Id:                 "id",
	PoolId:             "pool_id",
	SelectionDate:      "selection_date",
	SelectionPeriod:    "selection_period",
	TotalPoolSize:      "total_pool_size",
	RequiredSelections: "required_selections",
	SelectionAlgorithm: "selection_algorithm",
	SelectionSeed:      "selection_seed",
	Notes:              "notes",
	CreatedAt:          "created_at",
	CreatedBy:          "created_by",
}

// NewRandomSelectionsDao creates and returns a new DAO object for table data access.
func NewRandomSelectionsDao(handlers ...gdb.ModelHandler) *RandomSelectionsDao {
	return &RandomSelectionsDao{
		group:    "default",
		table:    "random_selections",
		columns:  randomSelectionsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RandomSelectionsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RandomSelectionsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RandomSelectionsDao) Columns() RandomSelectionsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RandomSelectionsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RandomSelectionsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *RandomSelectionsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
