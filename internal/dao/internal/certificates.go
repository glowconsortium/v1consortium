// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CertificatesDao is the data access object for the table certificates.
type CertificatesDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  CertificatesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// CertificatesColumns defines and stores column names for the table certificates.
type CertificatesColumns struct {
	Id                 string //
	OrganizationId     string //
	UserId             string //
	CertificateType    string //
	Title              string //
	Description        string //
	CertificateNumber  string //
	IssueDate          string //
	ExpirationDate     string //
	TestId             string //
	PhysicalId         string //
	CertificateUrl     string //
	TemplateUsed       string //
	IsDigitallySigned  string //
	SignatureHash      string //
	SignatureTimestamp string //
	DownloadCount      string //
	LastDownloadedAt   string //
	LastDownloadedBy   string //
	IsRevoked          string //
	RevokedAt          string //
	RevokedBy          string //
	RevocationReason   string //
	CreatedAt          string //
	UpdatedAt          string //
}

// certificatesColumns holds the columns for the table certificates.
var certificatesColumns = CertificatesColumns{
	Id:                 "id",
	OrganizationId:     "organization_id",
	UserId:             "user_id",
	CertificateType:    "certificate_type",
	Title:              "title",
	Description:        "description",
	CertificateNumber:  "certificate_number",
	IssueDate:          "issue_date",
	ExpirationDate:     "expiration_date",
	TestId:             "test_id",
	PhysicalId:         "physical_id",
	CertificateUrl:     "certificate_url",
	TemplateUsed:       "template_used",
	IsDigitallySigned:  "is_digitally_signed",
	SignatureHash:      "signature_hash",
	SignatureTimestamp: "signature_timestamp",
	DownloadCount:      "download_count",
	LastDownloadedAt:   "last_downloaded_at",
	LastDownloadedBy:   "last_downloaded_by",
	IsRevoked:          "is_revoked",
	RevokedAt:          "revoked_at",
	RevokedBy:          "revoked_by",
	RevocationReason:   "revocation_reason",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewCertificatesDao creates and returns a new DAO object for table data access.
func NewCertificatesDao(handlers ...gdb.ModelHandler) *CertificatesDao {
	return &CertificatesDao{
		group:    "default",
		table:    "certificates",
		columns:  certificatesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CertificatesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CertificatesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CertificatesDao) Columns() CertificatesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CertificatesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CertificatesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CertificatesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
