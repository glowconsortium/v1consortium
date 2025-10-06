// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Documents is the golang structure of table documents for DAO operations like Where/Data.
type Documents struct {
	g.Meta               `orm:"table:documents, do:true"`
	Id                   interface{} //
	OrganizationId       interface{} //
	UserId               interface{} //
	DocumentType         interface{} //
	Title                interface{} //
	Description          interface{} //
	FileName             interface{} //
	FileSize             interface{} //
	MimeType             interface{} //
	StoragePath          interface{} //
	StorageBucket        interface{} //
	TestId               interface{} //
	MvrReportId          interface{} //
	PhysicalId           interface{} //
	BackgroundCheckId    interface{} //
	UploadedBy           interface{} //
	UploadedAt           *gtime.Time //
	IsConfidential       interface{} //
	IsHipaaProtected     interface{} //
	RetentionPeriodYears interface{} //
	AutoDeleteDate       *gtime.Time //
	DownloadCount        interface{} //
	LastAccessedAt       *gtime.Time //
	LastAccessedBy       interface{} //
	Version              interface{} //
	ParentDocumentId     interface{} //
	IsCurrentVersion     interface{} //
	CreatedAt            *gtime.Time //
	UpdatedAt            *gtime.Time //
}
