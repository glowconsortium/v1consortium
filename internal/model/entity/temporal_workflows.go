// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TemporalWorkflows is the golang structure for table temporal_workflows.
type TemporalWorkflows struct {
	Id             string      `json:"id"             orm:"id"              description:""` //
	OrganizationId string      `json:"organizationId" orm:"organization_id" description:""` //
	WorkflowId     string      `json:"workflowId"     orm:"workflow_id"     description:""` //
	WorkflowType   string      `json:"workflowType"   orm:"workflow_type"   description:""` //
	Status         string      `json:"status"         orm:"status"          description:""` //
	InputData      string      `json:"inputData"      orm:"input_data"      description:""` //
	OutputData     string      `json:"outputData"     orm:"output_data"     description:""` //
	ErrorMessage   string      `json:"errorMessage"   orm:"error_message"   description:""` //
	UserId         string      `json:"userId"         orm:"user_id"         description:""` //
	TestId         string      `json:"testId"         orm:"test_id"         description:""` //
	SelectionId    string      `json:"selectionId"    orm:"selection_id"    description:""` //
	StartedAt      *gtime.Time `json:"startedAt"      orm:"started_at"      description:""` //
	CompletedAt    *gtime.Time `json:"completedAt"    orm:"completed_at"    description:""` //
	ScheduledFor   *gtime.Time `json:"scheduledFor"   orm:"scheduled_for"   description:""` //
	RetryCount     int         `json:"retryCount"     orm:"retry_count"     description:""` //
	MaxRetries     int         `json:"maxRetries"     orm:"max_retries"     description:""` //
	NextRetryAt    *gtime.Time `json:"nextRetryAt"    orm:"next_retry_at"   description:""` //
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:""` //
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:""` //
}
