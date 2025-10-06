// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BackgroundChecks is the golang structure for table background_checks.
type BackgroundChecks struct {
	Id                        string      `json:"id"                        orm:"id"                          description:""` //
	OrganizationId            string      `json:"organizationId"            orm:"organization_id"             description:""` //
	UserId                    string      `json:"userId"                    orm:"user_id"                     description:""` //
	CheckType                 string      `json:"checkType"                 orm:"check_type"                  description:""` //
	Status                    string      `json:"status"                    orm:"status"                      description:""` //
	OrderedDate               *gtime.Time `json:"orderedDate"               orm:"ordered_date"                description:""` //
	OrderedBy                 string      `json:"orderedBy"                 orm:"ordered_by"                  description:""` //
	ExternalOrderId           string      `json:"externalOrderId"           orm:"external_order_id"           description:""` //
	ProviderName              string      `json:"providerName"              orm:"provider_name"               description:""` //
	CompletedDate             *gtime.Time `json:"completedDate"             orm:"completed_date"              description:""` //
	ReportDate                *gtime.Time `json:"reportDate"                orm:"report_date"                 description:""` //
	OverallResult             string      `json:"overallResult"             orm:"overall_result"              description:""` //
	RequiresReview            bool        `json:"requiresReview"            orm:"requires_review"             description:""` //
	AdverseActionRequired     bool        `json:"adverseActionRequired"     orm:"adverse_action_required"     description:""` //
	FcraDisclosureSent        bool        `json:"fcraDisclosureSent"        orm:"fcra_disclosure_sent"        description:""` //
	FcraAuthorizationReceived bool        `json:"fcraAuthorizationReceived" orm:"fcra_authorization_received" description:""` //
	PreAdverseActionSent      bool        `json:"preAdverseActionSent"      orm:"pre_adverse_action_sent"     description:""` //
	AdverseActionSent         bool        `json:"adverseActionSent"         orm:"adverse_action_sent"         description:""` //
	Notes                     string      `json:"notes"                     orm:"notes"                       description:""` //
	CreatedAt                 *gtime.Time `json:"createdAt"                 orm:"created_at"                  description:""` //
	UpdatedAt                 *gtime.Time `json:"updatedAt"                 orm:"updated_at"                  description:""` //
}
