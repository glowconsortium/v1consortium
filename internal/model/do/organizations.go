// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Organizations is the golang structure of table organizations for DAO operations like Where/Data.
type Organizations struct {
	g.Meta         `orm:"table:organizations, do:true"`
	Id             interface{} //
	Name           interface{} //
	Type           interface{} //
	UsdotNumber    interface{} //
	McNumber       interface{} //
	Industry       interface{} //
	IsDotRegulated interface{} //
	AddressLine1   interface{} //
	AddressLine2   interface{} //
	City           interface{} //
	State          interface{} //
	ZipCode        interface{} //
	Country        interface{} //
	Phone          interface{} //
	Email          interface{} //
	Website        interface{} //
	TaxId          interface{} //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	IsActive       interface{} //
	Settings       interface{} //
}
