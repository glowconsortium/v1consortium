// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MvrViolations is the golang structure for table mvr_violations.
type MvrViolations struct {
	Id                           string      `json:"id"                           orm:"id"                             description:""` //
	MvrReportId                  string      `json:"mvrReportId"                  orm:"mvr_report_id"                  description:""` //
	ViolationDate                *gtime.Time `json:"violationDate"                orm:"violation_date"                 description:""` //
	ViolationCode                string      `json:"violationCode"                orm:"violation_code"                 description:""` //
	ViolationDescription         string      `json:"violationDescription"         orm:"violation_description"          description:""` //
	ViolationType                string      `json:"violationType"                orm:"violation_type"                 description:""` //
	Severity                     string      `json:"severity"                     orm:"severity"                       description:""` //
	ConvictionDate               *gtime.Time `json:"convictionDate"               orm:"conviction_date"                description:""` //
	FineAmount                   float64     `json:"fineAmount"                   orm:"fine_amount"                    description:""` //
	Points                       int         `json:"points"                       orm:"points"                         description:""` //
	State                        string      `json:"state"                        orm:"state"                          description:""` //
	CourtName                    string      `json:"courtName"                    orm:"court_name"                     description:""` //
	CaseNumber                   string      `json:"caseNumber"                   orm:"case_number"                    description:""` //
	Disqualifying                bool        `json:"disqualifying"                orm:"disqualifying"                  description:""` //
	RequiresEmployerNotification bool        `json:"requiresEmployerNotification" orm:"requires_employer_notification" description:""` //
	AffectsCdl                   bool        `json:"affectsCdl"                   orm:"affects_cdl"                    description:""` //
	CreatedAt                    *gtime.Time `json:"createdAt"                    orm:"created_at"                     description:""` //
}
