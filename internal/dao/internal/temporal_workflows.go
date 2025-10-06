// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TemporalWorkflowsDao is the data access object for the table temporal_workflows.
type TemporalWorkflowsDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  TemporalWorkflowsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// TemporalWorkflowsColumns defines and stores column names for the table temporal_workflows.
type TemporalWorkflowsColumns struct {
	Id             string //
	OrganizationId string //
	WorkflowId     string //
	WorkflowType   string //
	Status         string //
	InputData      string //
	OutputData     string //
	ErrorMessage   string //
	UserId         string //
	TestId         string //
	SelectionId    string //
	StartedAt      string //
	CompletedAt    string //
	ScheduledFor   string //
	RetryCount     string //
	MaxRetries     string //
	NextRetryAt    string //
	CreatedAt      string //
	UpdatedAt      string //
}

// temporalWorkflowsColumns holds the columns for the table temporal_workflows.
var temporalWorkflowsColumns = TemporalWorkflowsColumns{
	Id:             "id",
	OrganizationId: "organization_id",
	WorkflowId:     "workflow_id",
	WorkflowType:   "workflow_type",
	Status:         "status",
	InputData:      "input_data",
	OutputData:     "output_data",
	ErrorMessage:   "error_message",
	UserId:         "user_id",
	TestId:         "test_id",
	SelectionId:    "selection_id",
	StartedAt:      "started_at",
	CompletedAt:    "completed_at",
	ScheduledFor:   "scheduled_for",
	RetryCount:     "retry_count",
	MaxRetries:     "max_retries",
	NextRetryAt:    "next_retry_at",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewTemporalWorkflowsDao creates and returns a new DAO object for table data access.
func NewTemporalWorkflowsDao(handlers ...gdb.ModelHandler) *TemporalWorkflowsDao {
	return &TemporalWorkflowsDao{
		group:    "default",
		table:    "temporal_workflows",
		columns:  temporalWorkflowsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TemporalWorkflowsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TemporalWorkflowsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TemporalWorkflowsDao) Columns() TemporalWorkflowsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TemporalWorkflowsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TemporalWorkflowsDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TemporalWorkflowsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
