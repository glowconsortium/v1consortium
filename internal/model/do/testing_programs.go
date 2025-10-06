// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TestingPrograms is the golang structure of table testing_programs for DAO operations like Where/Data.
type TestingPrograms struct {
	g.Meta                `orm:"table:testing_programs, do:true"`
	Id                    interface{} //
	OrganizationId        interface{} //
	Name                  interface{} //
	IsDotProgram          interface{} //
	DrugPanelType         interface{} //
	AlcoholTestingEnabled interface{} //
	RandomTestingEnabled  interface{} //
	RandomTestingRate     interface{} //
	TestingFrequency      interface{} //
	IsActive              interface{} //
	CreatedAt             *gtime.Time //
	UpdatedAt             *gtime.Time //
}
