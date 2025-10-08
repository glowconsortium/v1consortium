// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Certificates is the golang structure for table certificates.
type Certificates struct {
	Id                 string      `json:"id"                 orm:"id"                  description:""` //
	OrganizationId     string      `json:"organizationId"     orm:"organization_id"     description:""` //
	UserId             string      `json:"userId"             orm:"user_id"             description:""` //
	CertificateType    string      `json:"certificateType"    orm:"certificate_type"    description:""` //
	Title              string      `json:"title"              orm:"title"               description:""` //
	Description        string      `json:"description"        orm:"description"         description:""` //
	CertificateNumber  string      `json:"certificateNumber"  orm:"certificate_number"  description:""` //
	IssueDate          *gtime.Time `json:"issueDate"          orm:"issue_date"          description:""` //
	ExpirationDate     *gtime.Time `json:"expirationDate"     orm:"expiration_date"     description:""` //
	TestId             string      `json:"testId"             orm:"test_id"             description:""` //
	PhysicalId         string      `json:"physicalId"         orm:"physical_id"         description:""` //
	CertificateUrl     string      `json:"certificateUrl"     orm:"certificate_url"     description:""` //
	TemplateUsed       string      `json:"templateUsed"       orm:"template_used"       description:""` //
	IsDigitallySigned  bool        `json:"isDigitallySigned"  orm:"is_digitally_signed" description:""` //
	SignatureHash      string      `json:"signatureHash"      orm:"signature_hash"      description:""` //
	SignatureTimestamp *gtime.Time `json:"signatureTimestamp" orm:"signature_timestamp" description:""` //
	DownloadCount      int         `json:"downloadCount"      orm:"download_count"      description:""` //
	LastDownloadedAt   *gtime.Time `json:"lastDownloadedAt"   orm:"last_downloaded_at"  description:""` //
	LastDownloadedBy   string      `json:"lastDownloadedBy"   orm:"last_downloaded_by"  description:""` //
	IsRevoked          bool        `json:"isRevoked"          orm:"is_revoked"          description:""` //
	RevokedAt          *gtime.Time `json:"revokedAt"          orm:"revoked_at"          description:""` //
	RevokedBy          string      `json:"revokedBy"          orm:"revoked_by"          description:""` //
	RevocationReason   string      `json:"revocationReason"   orm:"revocation_reason"   description:""` //
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"          description:""` //
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"          description:""` //
}
