// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RandomSelectionMembers is the golang structure of table random_selection_members for DAO operations like Where/Data.
type RandomSelectionMembers struct {
	g.Meta         `orm:"table:random_selection_members, do:true"`
	Id             interface{} //
	SelectionId    interface{} //
	UserId         interface{} //
	TestId         interface{} //
	SelectionOrder interface{} //
}
