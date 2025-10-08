// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DotPhysicalsDao is the data access object for the table dot_physicals.
type DotPhysicalsDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  DotPhysicalsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// DotPhysicalsColumns defines and stores column names for the table dot_physicals.
type DotPhysicalsColumns struct {
	Id                        string //
	OrganizationId            string //
	UserId                    string //
	Status                    string //
	ScheduledDate             string //
	ScheduledBy               string //
	ExaminerId                string //
	ExaminerName              string //
	ExaminerLicenseNumber     string //
	ExaminerRegistryNumber    string //
	ClinicName                string //
	ClinicAddress             string //
	ClinicPhone               string //
	ExaminationDate           string //
	CertificateNumber         string //
	CertificateIssueDate      string //
	CertificateExpirationDate string //
	MedicalQualification      string //
	Restrictions              string //
	Exemptions                string //
	RequiresMonitoring        string //
	MonitoringRequirements    string //
	NextRequiredDate          string //
	CertificateUrl            string //
	CertificateUploadedAt     string //
	Notes                     string //
	CreatedAt                 string //
	UpdatedAt                 string //
}

// dotPhysicalsColumns holds the columns for the table dot_physicals.
var dotPhysicalsColumns = DotPhysicalsColumns{
	Id:                        "id",
	OrganizationId:            "organization_id",
	UserId:                    "user_id",
	Status:                    "status",
	ScheduledDate:             "scheduled_date",
	ScheduledBy:               "scheduled_by",
	ExaminerId:                "examiner_id",
	ExaminerName:              "examiner_name",
	ExaminerLicenseNumber:     "examiner_license_number",
	ExaminerRegistryNumber:    "examiner_registry_number",
	ClinicName:                "clinic_name",
	ClinicAddress:             "clinic_address",
	ClinicPhone:               "clinic_phone",
	ExaminationDate:           "examination_date",
	CertificateNumber:         "certificate_number",
	CertificateIssueDate:      "certificate_issue_date",
	CertificateExpirationDate: "certificate_expiration_date",
	MedicalQualification:      "medical_qualification",
	Restrictions:              "restrictions",
	Exemptions:                "exemptions",
	RequiresMonitoring:        "requires_monitoring",
	MonitoringRequirements:    "monitoring_requirements",
	NextRequiredDate:          "next_required_date",
	CertificateUrl:            "certificate_url",
	CertificateUploadedAt:     "certificate_uploaded_at",
	Notes:                     "notes",
	CreatedAt:                 "created_at",
	UpdatedAt:                 "updated_at",
}

// NewDotPhysicalsDao creates and returns a new DAO object for table data access.
func NewDotPhysicalsDao(handlers ...gdb.ModelHandler) *DotPhysicalsDao {
	return &DotPhysicalsDao{
		group:    "default",
		table:    "dot_physicals",
		columns:  dotPhysicalsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DotPhysicalsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DotPhysicalsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DotPhysicalsDao) Columns() DotPhysicalsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DotPhysicalsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DotPhysicalsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *DotPhysicalsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
