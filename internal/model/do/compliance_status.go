// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ComplianceStatus is the golang structure of table compliance_status for DAO operations like Where/Data.
type ComplianceStatus struct {
	g.Meta                    `orm:"table:compliance_status, do:true"`
	Id                        interface{} //
	OrganizationId            interface{} //
	UserId                    interface{} //
	IsCompliant               interface{} //
	CompliancePercentage      interface{} //
	LastUpdated               *gtime.Time //
	DrugTestingCurrent        interface{} //
	LastDrugTestDate          *gtime.Time //
	NextDrugTestDue           *gtime.Time //
	MvrCurrent                interface{} //
	LastMvrDate               *gtime.Time //
	NextMvrDue                *gtime.Time //
	PhysicalCurrent           interface{} //
	MedicalCertExpirationDate *gtime.Time //
	BackgroundCheckCurrent    interface{} //
	LastBackgroundCheckDate   *gtime.Time //
	TrainingCurrent           interface{} //
	LastTrainingDate          *gtime.Time //
	ViolationsCount           interface{} //
	HighRiskFlags             interface{} //
	CreatedAt                 *gtime.Time //
	UpdatedAt                 *gtime.Time //
}
