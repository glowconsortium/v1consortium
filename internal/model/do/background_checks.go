// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BackgroundChecks is the golang structure of table background_checks for DAO operations like Where/Data.
type BackgroundChecks struct {
	g.Meta                    `orm:"table:background_checks, do:true"`
	Id                        interface{} //
	OrganizationId            interface{} //
	UserId                    interface{} //
	CheckType                 interface{} //
	Status                    interface{} //
	OrderedDate               *gtime.Time //
	OrderedBy                 interface{} //
	ExternalOrderId           interface{} //
	ProviderName              interface{} //
	CompletedDate             *gtime.Time //
	ReportDate                *gtime.Time //
	OverallResult             interface{} //
	RequiresReview            interface{} //
	AdverseActionRequired     interface{} //
	FcraDisclosureSent        interface{} //
	FcraAuthorizationReceived interface{} //
	PreAdverseActionSent      interface{} //
	AdverseActionSent         interface{} //
	Notes                     interface{} //
	CreatedAt                 *gtime.Time //
	UpdatedAt                 *gtime.Time //
}
