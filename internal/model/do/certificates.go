// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Certificates is the golang structure of table certificates for DAO operations like Where/Data.
type Certificates struct {
	g.Meta             `orm:"table:certificates, do:true"`
	Id                 interface{} //
	OrganizationId     interface{} //
	UserId             interface{} //
	CertificateType    interface{} //
	Title              interface{} //
	Description        interface{} //
	CertificateNumber  interface{} //
	IssueDate          *gtime.Time //
	ExpirationDate     *gtime.Time //
	TestId             interface{} //
	PhysicalId         interface{} //
	CertificateUrl     interface{} //
	TemplateUsed       interface{} //
	IsDigitallySigned  interface{} //
	SignatureHash      interface{} //
	SignatureTimestamp *gtime.Time //
	DownloadCount      interface{} //
	LastDownloadedAt   *gtime.Time //
	LastDownloadedBy   interface{} //
	IsRevoked          interface{} //
	RevokedAt          *gtime.Time //
	RevokedBy          interface{} //
	RevocationReason   interface{} //
	CreatedAt          *gtime.Time //
	UpdatedAt          *gtime.Time //
}
