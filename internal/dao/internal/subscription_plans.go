// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SubscriptionPlansDao is the data access object for the table subscription_plans.
type SubscriptionPlansDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  SubscriptionPlansColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// SubscriptionPlansColumns defines and stores column names for the table subscription_plans.
type SubscriptionPlansColumns struct {
	Id           string //
	Name         string //
	Tier         string //
	Description  string //
	MaxEmployees string //
	AnnualPrice  string //
	Features     string //
	IsActive     string //
	CreatedAt    string //
	UpdatedAt    string //
}

// subscriptionPlansColumns holds the columns for the table subscription_plans.
var subscriptionPlansColumns = SubscriptionPlansColumns{
	Id:           "id",
	Name:         "name",
	Tier:         "tier",
	Description:  "description",
	MaxEmployees: "max_employees",
	AnnualPrice:  "annual_price",
	Features:     "features",
	IsActive:     "is_active",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewSubscriptionPlansDao creates and returns a new DAO object for table data access.
func NewSubscriptionPlansDao(handlers ...gdb.ModelHandler) *SubscriptionPlansDao {
	return &SubscriptionPlansDao{
		group:    "default",
		table:    "subscription_plans",
		columns:  subscriptionPlansColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SubscriptionPlansDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SubscriptionPlansDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SubscriptionPlansDao) Columns() SubscriptionPlansColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SubscriptionPlansDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SubscriptionPlansDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SubscriptionPlansDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
