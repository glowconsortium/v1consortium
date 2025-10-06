// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DotPhysicals is the golang structure of table dot_physicals for DAO operations like Where/Data.
type DotPhysicals struct {
	g.Meta                    `orm:"table:dot_physicals, do:true"`
	Id                        interface{} //
	OrganizationId            interface{} //
	UserId                    interface{} //
	Status                    interface{} //
	ScheduledDate             *gtime.Time //
	ScheduledBy               interface{} //
	ExaminerId                interface{} //
	ExaminerName              interface{} //
	ExaminerLicenseNumber     interface{} //
	ExaminerRegistryNumber    interface{} //
	ClinicName                interface{} //
	ClinicAddress             interface{} //
	ClinicPhone               interface{} //
	ExaminationDate           *gtime.Time //
	CertificateNumber         interface{} //
	CertificateIssueDate      *gtime.Time //
	CertificateExpirationDate *gtime.Time //
	MedicalQualification      interface{} //
	Restrictions              interface{} //
	Exemptions                interface{} //
	RequiresMonitoring        interface{} //
	MonitoringRequirements    interface{} //
	NextRequiredDate          *gtime.Time //
	CertificateUrl            interface{} //
	CertificateUploadedAt     *gtime.Time //
	Notes                     interface{} //
	CreatedAt                 *gtime.Time //
	UpdatedAt                 *gtime.Time //
}
