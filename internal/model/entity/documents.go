// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Documents is the golang structure for table documents.
type Documents struct {
	Id                   string      `json:"id"                   orm:"id"                     description:""` //
	OrganizationId       string      `json:"organizationId"       orm:"organization_id"        description:""` //
	UserId               string      `json:"userId"               orm:"user_id"                description:""` //
	DocumentType         string      `json:"documentType"         orm:"document_type"          description:""` //
	Title                string      `json:"title"                orm:"title"                  description:""` //
	Description          string      `json:"description"          orm:"description"            description:""` //
	FileName             string      `json:"fileName"             orm:"file_name"              description:""` //
	FileSize             int64       `json:"fileSize"             orm:"file_size"              description:""` //
	MimeType             string      `json:"mimeType"             orm:"mime_type"              description:""` //
	StoragePath          string      `json:"storagePath"          orm:"storage_path"           description:""` //
	StorageBucket        string      `json:"storageBucket"        orm:"storage_bucket"         description:""` //
	TestId               string      `json:"testId"               orm:"test_id"                description:""` //
	MvrReportId          string      `json:"mvrReportId"          orm:"mvr_report_id"          description:""` //
	PhysicalId           string      `json:"physicalId"           orm:"physical_id"            description:""` //
	BackgroundCheckId    string      `json:"backgroundCheckId"    orm:"background_check_id"    description:""` //
	UploadedBy           string      `json:"uploadedBy"           orm:"uploaded_by"            description:""` //
	UploadedAt           *gtime.Time `json:"uploadedAt"           orm:"uploaded_at"            description:""` //
	IsConfidential       bool        `json:"isConfidential"       orm:"is_confidential"        description:""` //
	IsHipaaProtected     bool        `json:"isHipaaProtected"     orm:"is_hipaa_protected"     description:""` //
	RetentionPeriodYears int         `json:"retentionPeriodYears" orm:"retention_period_years" description:""` //
	AutoDeleteDate       *gtime.Time `json:"autoDeleteDate"       orm:"auto_delete_date"       description:""` //
	DownloadCount        int         `json:"downloadCount"        orm:"download_count"         description:""` //
	LastAccessedAt       *gtime.Time `json:"lastAccessedAt"       orm:"last_accessed_at"       description:""` //
	LastAccessedBy       string      `json:"lastAccessedBy"       orm:"last_accessed_by"       description:""` //
	Version              int         `json:"version"              orm:"version"                description:""` //
	ParentDocumentId     string      `json:"parentDocumentId"     orm:"parent_document_id"     description:""` //
	IsCurrentVersion     bool        `json:"isCurrentVersion"     orm:"is_current_version"     description:""` //
	CreatedAt            *gtime.Time `json:"createdAt"            orm:"created_at"             description:""` //
	UpdatedAt            *gtime.Time `json:"updatedAt"            orm:"updated_at"             description:""` //
}
