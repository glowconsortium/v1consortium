// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ComplianceStatus is the golang structure for table compliance_status.
type ComplianceStatus struct {
	Id                        string      `json:"id"                        orm:"id"                           description:""` //
	OrganizationId            string      `json:"organizationId"            orm:"organization_id"              description:""` //
	UserId                    string      `json:"userId"                    orm:"user_id"                      description:""` //
	IsCompliant               bool        `json:"isCompliant"               orm:"is_compliant"                 description:""` //
	CompliancePercentage      float64     `json:"compliancePercentage"      orm:"compliance_percentage"        description:""` //
	LastUpdated               *gtime.Time `json:"lastUpdated"               orm:"last_updated"                 description:""` //
	DrugTestingCurrent        bool        `json:"drugTestingCurrent"        orm:"drug_testing_current"         description:""` //
	LastDrugTestDate          *gtime.Time `json:"lastDrugTestDate"          orm:"last_drug_test_date"          description:""` //
	NextDrugTestDue           *gtime.Time `json:"nextDrugTestDue"           orm:"next_drug_test_due"           description:""` //
	MvrCurrent                bool        `json:"mvrCurrent"                orm:"mvr_current"                  description:""` //
	LastMvrDate               *gtime.Time `json:"lastMvrDate"               orm:"last_mvr_date"                description:""` //
	NextMvrDue                *gtime.Time `json:"nextMvrDue"                orm:"next_mvr_due"                 description:""` //
	PhysicalCurrent           bool        `json:"physicalCurrent"           orm:"physical_current"             description:""` //
	MedicalCertExpirationDate *gtime.Time `json:"medicalCertExpirationDate" orm:"medical_cert_expiration_date" description:""` //
	BackgroundCheckCurrent    bool        `json:"backgroundCheckCurrent"    orm:"background_check_current"     description:""` //
	LastBackgroundCheckDate   *gtime.Time `json:"lastBackgroundCheckDate"   orm:"last_background_check_date"   description:""` //
	TrainingCurrent           bool        `json:"trainingCurrent"           orm:"training_current"             description:""` //
	LastTrainingDate          *gtime.Time `json:"lastTrainingDate"          orm:"last_training_date"           description:""` //
	ViolationsCount           int         `json:"violationsCount"           orm:"violations_count"             description:""` //
	HighRiskFlags             int         `json:"highRiskFlags"             orm:"high_risk_flags"              description:""` //
	CreatedAt                 *gtime.Time `json:"createdAt"                 orm:"created_at"                   description:""` //
	UpdatedAt                 *gtime.Time `json:"updatedAt"                 orm:"updated_at"                   description:""` //
}
