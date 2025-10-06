// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TestingPrograms is the golang structure for table testing_programs.
type TestingPrograms struct {
	Id                    string      `json:"id"                    orm:"id"                      description:""` //
	OrganizationId        string      `json:"organizationId"        orm:"organization_id"         description:""` //
	Name                  string      `json:"name"                  orm:"name"                    description:""` //
	IsDotProgram          bool        `json:"isDotProgram"          orm:"is_dot_program"          description:""` //
	DrugPanelType         string      `json:"drugPanelType"         orm:"drug_panel_type"         description:""` //
	AlcoholTestingEnabled bool        `json:"alcoholTestingEnabled" orm:"alcohol_testing_enabled" description:""` //
	RandomTestingEnabled  bool        `json:"randomTestingEnabled"  orm:"random_testing_enabled"  description:""` //
	RandomTestingRate     float64     `json:"randomTestingRate"     orm:"random_testing_rate"     description:""` //
	TestingFrequency      string      `json:"testingFrequency"      orm:"testing_frequency"       description:""` //
	IsActive              bool        `json:"isActive"              orm:"is_active"               description:""` //
	CreatedAt             *gtime.Time `json:"createdAt"             orm:"created_at"              description:""` //
	UpdatedAt             *gtime.Time `json:"updatedAt"             orm:"updated_at"              description:""` //
}
