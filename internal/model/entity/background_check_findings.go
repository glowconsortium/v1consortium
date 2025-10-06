// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BackgroundCheckFindings is the golang structure for table background_check_findings.
type BackgroundCheckFindings struct {
	Id                               string      `json:"id"                               orm:"id"                                 description:""` //
	BackgroundCheckId                string      `json:"backgroundCheckId"                orm:"background_check_id"                description:""` //
	FindingType                      string      `json:"findingType"                      orm:"finding_type"                       description:""` //
	Severity                         string      `json:"severity"                         orm:"severity"                           description:""` //
	Description                      string      `json:"description"                      orm:"description"                        description:""` //
	DateOfRecord                     *gtime.Time `json:"dateOfRecord"                     orm:"date_of_record"                     description:""` //
	Jurisdiction                     string      `json:"jurisdiction"                     orm:"jurisdiction"                       description:""` //
	CaseNumber                       string      `json:"caseNumber"                       orm:"case_number"                        description:""` //
	Disposition                      string      `json:"disposition"                      orm:"disposition"                        description:""` //
	JobRelated                       bool        `json:"jobRelated"                       orm:"job_related"                        description:""` //
	Disqualifying                    bool        `json:"disqualifying"                    orm:"disqualifying"                      description:""` //
	RequiresIndividualizedAssessment bool        `json:"requiresIndividualizedAssessment" orm:"requires_individualized_assessment" description:""` //
	CreatedAt                        *gtime.Time `json:"createdAt"                        orm:"created_at"                         description:""` //
}
