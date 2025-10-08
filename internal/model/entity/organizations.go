// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Organizations is the golang structure for table organizations.
type Organizations struct {
	Id             string      `json:"id"             orm:"id"               description:""` //
	Name           string      `json:"name"           orm:"name"             description:""` //
	Type           string      `json:"type"           orm:"type"             description:""` //
	UsdotNumber    string      `json:"usdotNumber"    orm:"usdot_number"     description:""` //
	McNumber       string      `json:"mcNumber"       orm:"mc_number"        description:""` //
	Industry       string      `json:"industry"       orm:"industry"         description:""` //
	IsDotRegulated bool        `json:"isDotRegulated" orm:"is_dot_regulated" description:""` //
	AddressLine1   string      `json:"addressLine1"   orm:"address_line1"    description:""` //
	AddressLine2   string      `json:"addressLine2"   orm:"address_line2"    description:""` //
	City           string      `json:"city"           orm:"city"             description:""` //
	State          string      `json:"state"          orm:"state"            description:""` //
	ZipCode        string      `json:"zipCode"        orm:"zip_code"         description:""` //
	Country        string      `json:"country"        orm:"country"          description:""` //
	Phone          string      `json:"phone"          orm:"phone"            description:""` //
	Email          string      `json:"email"          orm:"email"            description:""` //
	Website        string      `json:"website"        orm:"website"          description:""` //
	TaxId          string      `json:"taxId"          orm:"tax_id"           description:""` //
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:""` //
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:""` //
	IsActive       bool        `json:"isActive"       orm:"is_active"        description:""` //
	Settings       string      `json:"settings"       orm:"settings"         description:""` //
}
