// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DrugAlcoholTests is the golang structure for table drug_alcohol_tests.
type DrugAlcoholTests struct {
	Id                       string      `json:"id"                       orm:"id"                         description:""` //
	OrganizationId           string      `json:"organizationId"           orm:"organization_id"            description:""` //
	UserId                   string      `json:"userId"                   orm:"user_id"                    description:""` //
	ProgramId                string      `json:"programId"                orm:"program_id"                 description:""` //
	SelectionId              string      `json:"selectionId"              orm:"selection_id"               description:""` //
	TestType                 string      `json:"testType"                 orm:"test_type"                  description:""` //
	TestCategory             string      `json:"testCategory"             orm:"test_category"              description:""` //
	Status                   string      `json:"status"                   orm:"status"                     description:""` //
	Result                   string      `json:"result"                   orm:"result"                     description:""` //
	IsDotTest                bool        `json:"isDotTest"                orm:"is_dot_test"                description:""` //
	OrderedDate              *gtime.Time `json:"orderedDate"              orm:"ordered_date"               description:""` //
	OrderedBy                string      `json:"orderedBy"                orm:"ordered_by"                 description:""` //
	DueDate                  *gtime.Time `json:"dueDate"                  orm:"due_date"                   description:""` //
	ExternalOrderId          string      `json:"externalOrderId"          orm:"external_order_id"          description:""` //
	ExternalFacilityId       string      `json:"externalFacilityId"       orm:"external_facility_id"       description:""` //
	FacilityName             string      `json:"facilityName"             orm:"facility_name"              description:""` //
	FacilityAddress          string      `json:"facilityAddress"          orm:"facility_address"           description:""` //
	CollectionDate           *gtime.Time `json:"collectionDate"           orm:"collection_date"            description:""` //
	CollectedBy              string      `json:"collectedBy"              orm:"collected_by"               description:""` //
	LabId                    string      `json:"labId"                    orm:"lab_id"                     description:""` //
	LabAccessionNumber       string      `json:"labAccessionNumber"       orm:"lab_accession_number"       description:""` //
	ResultDate               *gtime.Time `json:"resultDate"               orm:"result_date"                description:""` //
	ResultReceivedDate       *gtime.Time `json:"resultReceivedDate"       orm:"result_received_date"       description:""` //
	MroReviewRequired        bool        `json:"mroReviewRequired"        orm:"mro_review_required"        description:""` //
	MroId                    string      `json:"mroId"                    orm:"mro_id"                     description:""` //
	MroReviewDate            *gtime.Time `json:"mroReviewDate"            orm:"mro_review_date"            description:""` //
	MroNotes                 string      `json:"mroNotes"                 orm:"mro_notes"                  description:""` //
	RequiresImmediateRemoval bool        `json:"requiresImmediateRemoval" orm:"requires_immediate_removal" description:""` //
	ReturnToDutyRequired     bool        `json:"returnToDutyRequired"     orm:"return_to_duty_required"    description:""` //
	FollowUpTestsRequired    int         `json:"followUpTestsRequired"    orm:"follow_up_tests_required"   description:""` //
	Notes                    string      `json:"notes"                    orm:"notes"                      description:""` //
	CreatedAt                *gtime.Time `json:"createdAt"                orm:"created_at"                 description:""` //
	UpdatedAt                *gtime.Time `json:"updatedAt"                orm:"updated_at"                 description:""` //
}
