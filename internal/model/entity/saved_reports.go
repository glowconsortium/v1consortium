// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SavedReports is the golang structure for table saved_reports.
type SavedReports struct {
	Id                string      `json:"id"                orm:"id"                 description:""` //
	OrganizationId    string      `json:"organizationId"    orm:"organization_id"    description:""` //
	CreatedBy         string      `json:"createdBy"         orm:"created_by"         description:""` //
	Name              string      `json:"name"              orm:"name"               description:""` //
	Description       string      `json:"description"       orm:"description"        description:""` //
	ReportType        string      `json:"reportType"        orm:"report_type"        description:""` //
	Parameters        string      `json:"parameters"        orm:"parameters"         description:""` //
	Filters           string      `json:"filters"           orm:"filters"            description:""` //
	GeneratedAt       *gtime.Time `json:"generatedAt"       orm:"generated_at"       description:""` //
	FilePath          string      `json:"filePath"          orm:"file_path"          description:""` //
	FileFormat        string      `json:"fileFormat"        orm:"file_format"        description:""` //
	IsScheduled       bool        `json:"isScheduled"       orm:"is_scheduled"       description:""` //
	ScheduleFrequency string      `json:"scheduleFrequency" orm:"schedule_frequency" description:""` //
	NextRunDate       *gtime.Time `json:"nextRunDate"       orm:"next_run_date"      description:""` //
	IsPublic          bool        `json:"isPublic"          orm:"is_public"          description:""` //
	SharedWith        string      `json:"sharedWith"        orm:"shared_with"        description:""` //
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"         description:""` //
	UpdatedAt         *gtime.Time `json:"updatedAt"         orm:"updated_at"         description:""` //
}
