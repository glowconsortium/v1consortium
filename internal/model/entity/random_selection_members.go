// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RandomSelectionMembers is the golang structure for table random_selection_members.
type RandomSelectionMembers struct {
	Id             string `json:"id"             orm:"id"              description:""` //
	SelectionId    string `json:"selectionId"    orm:"selection_id"    description:""` //
	UserId         string `json:"userId"         orm:"user_id"         description:""` //
	TestId         string `json:"testId"         orm:"test_id"         description:""` //
	SelectionOrder int    `json:"selectionOrder" orm:"selection_order" description:""` //
}
