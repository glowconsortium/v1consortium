// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserProfilesDao is the data access object for the table user_profiles.
type UserProfilesDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  UserProfilesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// UserProfilesColumns defines and stores column names for the table user_profiles.
type UserProfilesColumns struct {
	Id                           string //
	OrganizationId               string //
	Role                         string //
	EmployeeId                   string //
	FirstName                    string //
	LastName                     string //
	Email                        string //
	Phone                        string //
	DateOfBirth                  string //
	SsnLastFour                  string //
	HireDate                     string //
	TerminationDate              string //
	IsActive                     string //
	RequiresDotTesting           string //
	RequiresNonDotTesting        string //
	CdlNumber                    string //
	CdlState                     string //
	CdlExpirationDate            string //
	JobTitle                     string //
	Department                   string //
	SupervisorId                 string //
	EmergencyContactName         string //
	EmergencyContactPhone        string //
	EmergencyContactRelationship string //
	CreatedAt                    string //
	UpdatedAt                    string //
	LastLoginAt                  string //
}

// userProfilesColumns holds the columns for the table user_profiles.
var userProfilesColumns = UserProfilesColumns{
	Id:                           "id",
	OrganizationId:               "organization_id",
	Role:                         "role",
	EmployeeId:                   "employee_id",
	FirstName:                    "first_name",
	LastName:                     "last_name",
	Email:                        "email",
	Phone:                        "phone",
	DateOfBirth:                  "date_of_birth",
	SsnLastFour:                  "ssn_last_four",
	HireDate:                     "hire_date",
	TerminationDate:              "termination_date",
	IsActive:                     "is_active",
	RequiresDotTesting:           "requires_dot_testing",
	RequiresNonDotTesting:        "requires_non_dot_testing",
	CdlNumber:                    "cdl_number",
	CdlState:                     "cdl_state",
	CdlExpirationDate:            "cdl_expiration_date",
	JobTitle:                     "job_title",
	Department:                   "department",
	SupervisorId:                 "supervisor_id",
	EmergencyContactName:         "emergency_contact_name",
	EmergencyContactPhone:        "emergency_contact_phone",
	EmergencyContactRelationship: "emergency_contact_relationship",
	CreatedAt:                    "created_at",
	UpdatedAt:                    "updated_at",
	LastLoginAt:                  "last_login_at",
}

// NewUserProfilesDao creates and returns a new DAO object for table data access.
func NewUserProfilesDao(handlers ...gdb.ModelHandler) *UserProfilesDao {
	return &UserProfilesDao{
		group:    "default",
		table:    "user_profiles",
		columns:  userProfilesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserProfilesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserProfilesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserProfilesDao) Columns() UserProfilesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserProfilesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserProfilesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserProfilesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
