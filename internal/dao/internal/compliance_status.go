// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ComplianceStatusDao is the data access object for the table compliance_status.
type ComplianceStatusDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  ComplianceStatusColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// ComplianceStatusColumns defines and stores column names for the table compliance_status.
type ComplianceStatusColumns struct {
	Id                        string //
	OrganizationId            string //
	UserId                    string //
	IsCompliant               string //
	CompliancePercentage      string //
	LastUpdated               string //
	DrugTestingCurrent        string //
	LastDrugTestDate          string //
	NextDrugTestDue           string //
	MvrCurrent                string //
	LastMvrDate               string //
	NextMvrDue                string //
	PhysicalCurrent           string //
	MedicalCertExpirationDate string //
	BackgroundCheckCurrent    string //
	LastBackgroundCheckDate   string //
	TrainingCurrent           string //
	LastTrainingDate          string //
	ViolationsCount           string //
	HighRiskFlags             string //
	CreatedAt                 string //
	UpdatedAt                 string //
}

// complianceStatusColumns holds the columns for the table compliance_status.
var complianceStatusColumns = ComplianceStatusColumns{
	Id:                        "id",
	OrganizationId:            "organization_id",
	UserId:                    "user_id",
	IsCompliant:               "is_compliant",
	CompliancePercentage:      "compliance_percentage",
	LastUpdated:               "last_updated",
	DrugTestingCurrent:        "drug_testing_current",
	LastDrugTestDate:          "last_drug_test_date",
	NextDrugTestDue:           "next_drug_test_due",
	MvrCurrent:                "mvr_current",
	LastMvrDate:               "last_mvr_date",
	NextMvrDue:                "next_mvr_due",
	PhysicalCurrent:           "physical_current",
	MedicalCertExpirationDate: "medical_cert_expiration_date",
	BackgroundCheckCurrent:    "background_check_current",
	LastBackgroundCheckDate:   "last_background_check_date",
	TrainingCurrent:           "training_current",
	LastTrainingDate:          "last_training_date",
	ViolationsCount:           "violations_count",
	HighRiskFlags:             "high_risk_flags",
	CreatedAt:                 "created_at",
	UpdatedAt:                 "updated_at",
}

// NewComplianceStatusDao creates and returns a new DAO object for table data access.
func NewComplianceStatusDao(handlers ...gdb.ModelHandler) *ComplianceStatusDao {
	return &ComplianceStatusDao{
		group:    "default",
		table:    "compliance_status",
		columns:  complianceStatusColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ComplianceStatusDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ComplianceStatusDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ComplianceStatusDao) Columns() ComplianceStatusColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ComplianceStatusDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ComplianceStatusDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ComplianceStatusDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
