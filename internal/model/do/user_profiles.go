// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserProfiles is the golang structure of table user_profiles for DAO operations like Where/Data.
type UserProfiles struct {
	g.Meta                       `orm:"table:user_profiles, do:true"`
	Id                           interface{} //
	OrganizationId               interface{} //
	Role                         interface{} //
	EmployeeId                   interface{} //
	FirstName                    interface{} //
	LastName                     interface{} //
	Email                        interface{} //
	Phone                        interface{} //
	DateOfBirth                  *gtime.Time //
	SsnLastFour                  interface{} //
	HireDate                     *gtime.Time //
	TerminationDate              *gtime.Time //
	IsActive                     interface{} //
	IsSystemUser                 interface{} //
	StripeCustomerId             interface{} //
	RequiresDotTesting           interface{} //
	RequiresNonDotTesting        interface{} //
	CdlNumber                    interface{} //
	CdlState                     interface{} //
	CdlExpirationDate            *gtime.Time //
	JobTitle                     interface{} //
	Department                   interface{} //
	SupervisorId                 interface{} //
	EmergencyContactName         interface{} //
	EmergencyContactPhone        interface{} //
	EmergencyContactRelationship interface{} //
	CreatedAt                    *gtime.Time //
	UpdatedAt                    *gtime.Time //
	LastLoginAt                  *gtime.Time //
	Settings                     interface{} //
}
