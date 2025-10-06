// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MvrReports is the golang structure for table mvr_reports.
type MvrReports struct {
	Id                    string      `json:"id"                    orm:"id"                      description:""` //
	OrganizationId        string      `json:"organizationId"        orm:"organization_id"         description:""` //
	UserId                string      `json:"userId"                orm:"user_id"                 description:""` //
	Status                string      `json:"status"                orm:"status"                  description:""` //
	OrderedDate           *gtime.Time `json:"orderedDate"           orm:"ordered_date"            description:""` //
	OrderedBy             string      `json:"orderedBy"             orm:"ordered_by"              description:""` //
	LicenseNumber         string      `json:"licenseNumber"         orm:"license_number"          description:""` //
	LicenseState          string      `json:"licenseState"          orm:"license_state"           description:""` //
	ExternalOrderId       string      `json:"externalOrderId"       orm:"external_order_id"       description:""` //
	ProviderName          string      `json:"providerName"          orm:"provider_name"           description:""` //
	ReportDate            *gtime.Time `json:"reportDate"            orm:"report_date"             description:""` //
	ReportReceivedDate    *gtime.Time `json:"reportReceivedDate"    orm:"report_received_date"    description:""` //
	RawReportData         string      `json:"rawReportData"         orm:"raw_report_data"         description:""` //
	TotalViolations       int         `json:"totalViolations"       orm:"total_violations"        description:""` //
	MajorViolations       int         `json:"majorViolations"       orm:"major_violations"        description:""` //
	MinorViolations       int         `json:"minorViolations"       orm:"minor_violations"        description:""` //
	LicenseStatus         string      `json:"licenseStatus"         orm:"license_status"          description:""` //
	LicenseExpirationDate *gtime.Time `json:"licenseExpirationDate" orm:"license_expiration_date" description:""` //
	ReviewedBy            string      `json:"reviewedBy"            orm:"reviewed_by"             description:""` //
	ReviewedDate          *gtime.Time `json:"reviewedDate"          orm:"reviewed_date"           description:""` //
	RequiresAction        bool        `json:"requiresAction"        orm:"requires_action"         description:""` //
	ActionNotes           string      `json:"actionNotes"           orm:"action_notes"            description:""` //
	CreatedAt             *gtime.Time `json:"createdAt"             orm:"created_at"              description:""` //
	UpdatedAt             *gtime.Time `json:"updatedAt"             orm:"updated_at"              description:""` //
}
