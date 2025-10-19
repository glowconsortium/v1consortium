// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserProfiles is the golang structure for table user_profiles.
type UserProfiles struct {
	Id                           string      `json:"id"                           orm:"id"                             description:""` //
	OrganizationId               string      `json:"organizationId"               orm:"organization_id"                description:""` //
	Role                         string      `json:"role"                         orm:"role"                           description:""` //
	EmployeeId                   string      `json:"employeeId"                   orm:"employee_id"                    description:""` //
	FirstName                    string      `json:"firstName"                    orm:"first_name"                     description:""` //
	LastName                     string      `json:"lastName"                     orm:"last_name"                      description:""` //
	Email                        string      `json:"email"                        orm:"email"                          description:""` //
	Phone                        string      `json:"phone"                        orm:"phone"                          description:""` //
	DateOfBirth                  *gtime.Time `json:"dateOfBirth"                  orm:"date_of_birth"                  description:""` //
	SsnLastFour                  string      `json:"ssnLastFour"                  orm:"ssn_last_four"                  description:""` //
	HireDate                     *gtime.Time `json:"hireDate"                     orm:"hire_date"                      description:""` //
	TerminationDate              *gtime.Time `json:"terminationDate"              orm:"termination_date"               description:""` //
	IsActive                     bool        `json:"isActive"                     orm:"is_active"                      description:""` //
	IsSystemUser                 bool        `json:"isSystemUser"                 orm:"is_system_user"                 description:""` //
	StripeCustomerId             string      `json:"stripeCustomerId"             orm:"stripe_customer_id"             description:""` //
	RequiresDotTesting           bool        `json:"requiresDotTesting"           orm:"requires_dot_testing"           description:""` //
	RequiresNonDotTesting        bool        `json:"requiresNonDotTesting"        orm:"requires_non_dot_testing"       description:""` //
	CdlNumber                    string      `json:"cdlNumber"                    orm:"cdl_number"                     description:""` //
	CdlState                     string      `json:"cdlState"                     orm:"cdl_state"                      description:""` //
	CdlExpirationDate            *gtime.Time `json:"cdlExpirationDate"            orm:"cdl_expiration_date"            description:""` //
	JobTitle                     string      `json:"jobTitle"                     orm:"job_title"                      description:""` //
	Department                   string      `json:"department"                   orm:"department"                     description:""` //
	SupervisorId                 string      `json:"supervisorId"                 orm:"supervisor_id"                  description:""` //
	EmergencyContactName         string      `json:"emergencyContactName"         orm:"emergency_contact_name"         description:""` //
	EmergencyContactPhone        string      `json:"emergencyContactPhone"        orm:"emergency_contact_phone"        description:""` //
	EmergencyContactRelationship string      `json:"emergencyContactRelationship" orm:"emergency_contact_relationship" description:""` //
	CreatedAt                    *gtime.Time `json:"createdAt"                    orm:"created_at"                     description:""` //
	UpdatedAt                    *gtime.Time `json:"updatedAt"                    orm:"updated_at"                     description:""` //
	LastLoginAt                  *gtime.Time `json:"lastLoginAt"                  orm:"last_login_at"                  description:""` //
	Settings                     string      `json:"settings"                     orm:"settings"                       description:""` //
}
