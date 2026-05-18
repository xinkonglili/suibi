// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TrainOrder is the golang structure of table train_order for DAO operations like Where/Data.
type TrainOrder struct {
	g.Meta     `orm:"table:train_order, do:true"`
	Id         any //
	UserId     any // 用户id
	OrderId    any //
	UserName   any //
	FareStatus any //
}
