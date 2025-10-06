// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DotPhysicals is the golang structure for table dot_physicals.
type DotPhysicals struct {
	Id                        string      `json:"id"                        orm:"id"                          description:""` //
	OrganizationId            string      `json:"organizationId"            orm:"organization_id"             description:""` //
	UserId                    string      `json:"userId"                    orm:"user_id"                     description:""` //
	Status                    string      `json:"status"                    orm:"status"                      description:""` //
	ScheduledDate             *gtime.Time `json:"scheduledDate"             orm:"scheduled_date"              description:""` //
	ScheduledBy               string      `json:"scheduledBy"               orm:"scheduled_by"                description:""` //
	ExaminerId                string      `json:"examinerId"                orm:"examiner_id"                 description:""` //
	ExaminerName              string      `json:"examinerName"              orm:"examiner_name"               description:""` //
	ExaminerLicenseNumber     string      `json:"examinerLicenseNumber"     orm:"examiner_license_number"     description:""` //
	ExaminerRegistryNumber    string      `json:"examinerRegistryNumber"    orm:"examiner_registry_number"    description:""` //
	ClinicName                string      `json:"clinicName"                orm:"clinic_name"                 description:""` //
	ClinicAddress             string      `json:"clinicAddress"             orm:"clinic_address"              description:""` //
	ClinicPhone               string      `json:"clinicPhone"               orm:"clinic_phone"                description:""` //
	ExaminationDate           *gtime.Time `json:"examinationDate"           orm:"examination_date"            description:""` //
	CertificateNumber         string      `json:"certificateNumber"         orm:"certificate_number"          description:""` //
	CertificateIssueDate      *gtime.Time `json:"certificateIssueDate"      orm:"certificate_issue_date"      description:""` //
	CertificateExpirationDate *gtime.Time `json:"certificateExpirationDate" orm:"certificate_expiration_date" description:""` //
	MedicalQualification      string      `json:"medicalQualification"      orm:"medical_qualification"       description:""` //
	Restrictions              string      `json:"restrictions"              orm:"restrictions"                description:""` //
	Exemptions                string      `json:"exemptions"                orm:"exemptions"                  description:""` //
	RequiresMonitoring        bool        `json:"requiresMonitoring"        orm:"requires_monitoring"         description:""` //
	MonitoringRequirements    string      `json:"monitoringRequirements"    orm:"monitoring_requirements"     description:""` //
	NextRequiredDate          *gtime.Time `json:"nextRequiredDate"          orm:"next_required_date"          description:""` //
	CertificateUrl            string      `json:"certificateUrl"            orm:"certificate_url"             description:""` //
	CertificateUploadedAt     *gtime.Time `json:"certificateUploadedAt"     orm:"certificate_uploaded_at"     description:""` //
	Notes                     string      `json:"notes"                     orm:"notes"                       description:""` //
	CreatedAt                 *gtime.Time `json:"createdAt"                 orm:"created_at"                  description:""` //
	UpdatedAt                 *gtime.Time `json:"updatedAt"                 orm:"updated_at"                  description:""` //
}
