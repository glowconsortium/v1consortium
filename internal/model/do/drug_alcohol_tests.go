// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DrugAlcoholTests is the golang structure of table drug_alcohol_tests for DAO operations like Where/Data.
type DrugAlcoholTests struct {
	g.Meta                   `orm:"table:drug_alcohol_tests, do:true"`
	Id                       interface{} //
	OrganizationId           interface{} //
	UserId                   interface{} //
	ProgramId                interface{} //
	SelectionId              interface{} //
	TestType                 interface{} //
	TestCategory             interface{} //
	Status                   interface{} //
	Result                   interface{} //
	IsDotTest                interface{} //
	OrderedDate              *gtime.Time //
	OrderedBy                interface{} //
	DueDate                  *gtime.Time //
	ExternalOrderId          interface{} //
	ExternalFacilityId       interface{} //
	FacilityName             interface{} //
	FacilityAddress          interface{} //
	CollectionDate           *gtime.Time //
	CollectedBy              interface{} //
	LabId                    interface{} //
	LabAccessionNumber       interface{} //
	ResultDate               *gtime.Time //
	ResultReceivedDate       *gtime.Time //
	MroReviewRequired        interface{} //
	MroId                    interface{} //
	MroReviewDate            *gtime.Time //
	MroNotes                 interface{} //
	RequiresImmediateRemoval interface{} //
	ReturnToDutyRequired     interface{} //
	FollowUpTestsRequired    interface{} //
	Notes                    interface{} //
	CreatedAt                *gtime.Time //
	UpdatedAt                *gtime.Time //
}
