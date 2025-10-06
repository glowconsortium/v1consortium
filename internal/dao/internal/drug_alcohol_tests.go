// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DrugAlcoholTestsDao is the data access object for the table drug_alcohol_tests.
type DrugAlcoholTestsDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  DrugAlcoholTestsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// DrugAlcoholTestsColumns defines and stores column names for the table drug_alcohol_tests.
type DrugAlcoholTestsColumns struct {
	Id                       string //
	OrganizationId           string //
	UserId                   string //
	ProgramId                string //
	SelectionId              string //
	TestType                 string //
	TestCategory             string //
	Status                   string //
	Result                   string //
	IsDotTest                string //
	OrderedDate              string //
	OrderedBy                string //
	DueDate                  string //
	ExternalOrderId          string //
	ExternalFacilityId       string //
	FacilityName             string //
	FacilityAddress          string //
	CollectionDate           string //
	CollectedBy              string //
	LabId                    string //
	LabAccessionNumber       string //
	ResultDate               string //
	ResultReceivedDate       string //
	MroReviewRequired        string //
	MroId                    string //
	MroReviewDate            string //
	MroNotes                 string //
	RequiresImmediateRemoval string //
	ReturnToDutyRequired     string //
	FollowUpTestsRequired    string //
	Notes                    string //
	CreatedAt                string //
	UpdatedAt                string //
}

// drugAlcoholTestsColumns holds the columns for the table drug_alcohol_tests.
var drugAlcoholTestsColumns = DrugAlcoholTestsColumns{
	Id:                       "id",
	OrganizationId:           "organization_id",
	UserId:                   "user_id",
	ProgramId:                "program_id",
	SelectionId:              "selection_id",
	TestType:                 "test_type",
	TestCategory:             "test_category",
	Status:                   "status",
	Result:                   "result",
	IsDotTest:                "is_dot_test",
	OrderedDate:              "ordered_date",
	OrderedBy:                "ordered_by",
	DueDate:                  "due_date",
	ExternalOrderId:          "external_order_id",
	ExternalFacilityId:       "external_facility_id",
	FacilityName:             "facility_name",
	FacilityAddress:          "facility_address",
	CollectionDate:           "collection_date",
	CollectedBy:              "collected_by",
	LabId:                    "lab_id",
	LabAccessionNumber:       "lab_accession_number",
	ResultDate:               "result_date",
	ResultReceivedDate:       "result_received_date",
	MroReviewRequired:        "mro_review_required",
	MroId:                    "mro_id",
	MroReviewDate:            "mro_review_date",
	MroNotes:                 "mro_notes",
	RequiresImmediateRemoval: "requires_immediate_removal",
	ReturnToDutyRequired:     "return_to_duty_required",
	FollowUpTestsRequired:    "follow_up_tests_required",
	Notes:                    "notes",
	CreatedAt:                "created_at",
	UpdatedAt:                "updated_at",
}

// NewDrugAlcoholTestsDao creates and returns a new DAO object for table data access.
func NewDrugAlcoholTestsDao(handlers ...gdb.ModelHandler) *DrugAlcoholTestsDao {
	return &DrugAlcoholTestsDao{
		group:    "default",
		table:    "drug_alcohol_tests",
		columns:  drugAlcoholTestsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DrugAlcoholTestsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DrugAlcoholTestsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DrugAlcoholTestsDao) Columns() DrugAlcoholTestsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DrugAlcoholTestsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DrugAlcoholTestsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *DrugAlcoholTestsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
