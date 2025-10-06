// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrganizationsDao is the data access object for the table organizations.
type OrganizationsDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  OrganizationsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// OrganizationsColumns defines and stores column names for the table organizations.
type OrganizationsColumns struct {
	Id             string //
	Name           string //
	Type           string //
	UsdotNumber    string //
	McNumber       string //
	Industry       string //
	IsDotRegulated string //
	AddressLine1   string //
	AddressLine2   string //
	City           string //
	State          string //
	ZipCode        string //
	Country        string //
	Phone          string //
	Email          string //
	Website        string //
	TaxId          string //
	CreatedAt      string //
	UpdatedAt      string //
	IsActive       string //
	Settings       string //
}

// organizationsColumns holds the columns for the table organizations.
var organizationsColumns = OrganizationsColumns{
	Id:             "id",
	Name:           "name",
	Type:           "type",
	UsdotNumber:    "usdot_number",
	McNumber:       "mc_number",
	Industry:       "industry",
	IsDotRegulated: "is_dot_regulated",
	AddressLine1:   "address_line1",
	AddressLine2:   "address_line2",
	City:           "city",
	State:          "state",
	ZipCode:        "zip_code",
	Country:        "country",
	Phone:          "phone",
	Email:          "email",
	Website:        "website",
	TaxId:          "tax_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	IsActive:       "is_active",
	Settings:       "settings",
}

// NewOrganizationsDao creates and returns a new DAO object for table data access.
func NewOrganizationsDao(handlers ...gdb.ModelHandler) *OrganizationsDao {
	return &OrganizationsDao{
		group:    "default",
		table:    "organizations",
		columns:  organizationsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OrganizationsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OrganizationsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OrganizationsDao) Columns() OrganizationsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OrganizationsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OrganizationsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OrganizationsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
