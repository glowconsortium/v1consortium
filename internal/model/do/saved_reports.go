// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SavedReports is the golang structure of table saved_reports for DAO operations like Where/Data.
type SavedReports struct {
	g.Meta            `orm:"table:saved_reports, do:true"`
	Id                interface{} //
	OrganizationId    interface{} //
	CreatedBy         interface{} //
	Name              interface{} //
	Description       interface{} //
	ReportType        interface{} //
	Parameters        interface{} //
	Filters           interface{} //
	GeneratedAt       *gtime.Time //
	FilePath          interface{} //
	FileFormat        interface{} //
	IsScheduled       interface{} //
	ScheduleFrequency interface{} //
	NextRunDate       *gtime.Time //
	IsPublic          interface{} //
	SharedWith        interface{} //
	CreatedAt         *gtime.Time //
	UpdatedAt         *gtime.Time //
}
