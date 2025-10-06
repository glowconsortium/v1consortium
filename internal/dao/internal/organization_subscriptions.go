// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrganizationSubscriptionsDao is the data access object for the table organization_subscriptions.
type OrganizationSubscriptionsDao struct {
	table    string                           // table is the underlying table name of the DAO.
	group    string                           // group is the database configuration group name of the current DAO.
	columns  OrganizationSubscriptionsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler               // handlers for customized model modification.
}

// OrganizationSubscriptionsColumns defines and stores column names for the table organization_subscriptions.
type OrganizationSubscriptionsColumns struct {
	Id                   string //
	OrganizationId       string //
	PlanId               string //
	Status               string //
	StartDate            string //
	EndDate              string //
	AutoRenew            string //
	StripeSubscriptionId string //
	StripeCustomerId     string //
	CreatedAt            string //
	UpdatedAt            string //
}

// organizationSubscriptionsColumns holds the columns for the table organization_subscriptions.
var organizationSubscriptionsColumns = OrganizationSubscriptionsColumns{
	Id:                   "id",
	OrganizationId:       "organization_id",
	PlanId:               "plan_id",
	Status:               "status",
	StartDate:            "start_date",
	EndDate:              "end_date",
	AutoRenew:            "auto_renew",
	StripeSubscriptionId: "stripe_subscription_id",
	StripeCustomerId:     "stripe_customer_id",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
}

// NewOrganizationSubscriptionsDao creates and returns a new DAO object for table data access.
func NewOrganizationSubscriptionsDao(handlers ...gdb.ModelHandler) *OrganizationSubscriptionsDao {
	return &OrganizationSubscriptionsDao{
		group:    "default",
		table:    "organization_subscriptions",
		columns:  organizationSubscriptionsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OrganizationSubscriptionsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OrganizationSubscriptionsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OrganizationSubscriptionsDao) Columns() OrganizationSubscriptionsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OrganizationSubscriptionsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OrganizationSubscriptionsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OrganizationSubscriptionsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
