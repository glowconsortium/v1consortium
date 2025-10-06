// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BackgroundCheckFindings is the golang structure of table background_check_findings for DAO operations like Where/Data.
type BackgroundCheckFindings struct {
	g.Meta                           `orm:"table:background_check_findings, do:true"`
	Id                               interface{} //
	BackgroundCheckId                interface{} //
	FindingType                      interface{} //
	Severity                         interface{} //
	Description                      interface{} //
	DateOfRecord                     *gtime.Time //
	Jurisdiction                     interface{} //
	CaseNumber                       interface{} //
	Disposition                      interface{} //
	JobRelated                       interface{} //
	Disqualifying                    interface{} //
	RequiresIndividualizedAssessment interface{} //
	CreatedAt                        *gtime.Time //
}
