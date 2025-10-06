// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MvrReports is the golang structure of table mvr_reports for DAO operations like Where/Data.
type MvrReports struct {
	g.Meta                `orm:"table:mvr_reports, do:true"`
	Id                    interface{} //
	OrganizationId        interface{} //
	UserId                interface{} //
	Status                interface{} //
	OrderedDate           *gtime.Time //
	OrderedBy             interface{} //
	LicenseNumber         interface{} //
	LicenseState          interface{} //
	ExternalOrderId       interface{} //
	ProviderName          interface{} //
	ReportDate            *gtime.Time //
	ReportReceivedDate    *gtime.Time //
	RawReportData         interface{} //
	TotalViolations       interface{} //
	MajorViolations       interface{} //
	MinorViolations       interface{} //
	LicenseStatus         interface{} //
	LicenseExpirationDate *gtime.Time //
	ReviewedBy            interface{} //
	ReviewedDate          *gtime.Time //
	RequiresAction        interface{} //
	ActionNotes           interface{} //
	CreatedAt             *gtime.Time //
	UpdatedAt             *gtime.Time //
}
