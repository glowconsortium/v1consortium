// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MvrViolations is the golang structure of table mvr_violations for DAO operations like Where/Data.
type MvrViolations struct {
	g.Meta                       `orm:"table:mvr_violations, do:true"`
	Id                           interface{} //
	MvrReportId                  interface{} //
	ViolationDate                *gtime.Time //
	ViolationCode                interface{} //
	ViolationDescription         interface{} //
	ViolationType                interface{} //
	Severity                     interface{} //
	ConvictionDate               *gtime.Time //
	FineAmount                   interface{} //
	Points                       interface{} //
	State                        interface{} //
	CourtName                    interface{} //
	CaseNumber                   interface{} //
	Disqualifying                interface{} //
	RequiresEmployerNotification interface{} //
	AffectsCdl                   interface{} //
	CreatedAt                    *gtime.Time //
}
