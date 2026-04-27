// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// User is the golang structure of frame user for DAO operations like Where/Data.
type User struct {
	g.Meta `orm:"frame:user, do:true"`
	Id     any // user id
	Name   any // user name
	Status any // user status
	Age    any // user age
}
