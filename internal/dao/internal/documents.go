// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DocumentsDao is the data access object for the table documents.
type DocumentsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  DocumentsColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// DocumentsColumns defines and stores column names for the table documents.
type DocumentsColumns struct {
	Id                   string //
	OrganizationId       string //
	UserId               string //
	DocumentType         string //
	Title                string //
	Description          string //
	FileName             string //
	FileSize             string //
	MimeType             string //
	StoragePath          string //
	StorageBucket        string //
	TestId               string //
	MvrReportId          string //
	PhysicalId           string //
	BackgroundCheckId    string //
	UploadedBy           string //
	UploadedAt           string //
	IsConfidential       string //
	IsHipaaProtected     string //
	RetentionPeriodYears string //
	AutoDeleteDate       string //
	DownloadCount        string //
	LastAccessedAt       string //
	LastAccessedBy       string //
	Version              string //
	ParentDocumentId     string //
	IsCurrentVersion     string //
	CreatedAt            string //
	UpdatedAt            string //
}

// documentsColumns holds the columns for the table documents.
var documentsColumns = DocumentsColumns{
	Id:                   "id",
	OrganizationId:       "organization_id",
	UserId:               "user_id",
	DocumentType:         "document_type",
	Title:                "title",
	Description:          "description",
	FileName:             "file_name",
	FileSize:             "file_size",
	MimeType:             "mime_type",
	StoragePath:          "storage_path",
	StorageBucket:        "storage_bucket",
	TestId:               "test_id",
	MvrReportId:          "mvr_report_id",
	PhysicalId:           "physical_id",
	BackgroundCheckId:    "background_check_id",
	UploadedBy:           "uploaded_by",
	UploadedAt:           "uploaded_at",
	IsConfidential:       "is_confidential",
	IsHipaaProtected:     "is_hipaa_protected",
	RetentionPeriodYears: "retention_period_years",
	AutoDeleteDate:       "auto_delete_date",
	DownloadCount:        "download_count",
	LastAccessedAt:       "last_accessed_at",
	LastAccessedBy:       "last_accessed_by",
	Version:              "version",
	ParentDocumentId:     "parent_document_id",
	IsCurrentVersion:     "is_current_version",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
}

// NewDocumentsDao creates and returns a new DAO object for table data access.
func NewDocumentsDao(handlers ...gdb.ModelHandler) *DocumentsDao {
	return &DocumentsDao{
		group:    "default",
		table:    "documents",
		columns:  documentsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DocumentsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DocumentsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DocumentsDao) Columns() DocumentsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DocumentsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DocumentsDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *DocumentsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
