// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PriceRules is the golang structure of table price_rules for DAO operations like Where/Data.
type PriceRules struct {
	g.Meta      `orm:"table:price_rules, do:true"`
	Id          any         //
	TrainTypeId any         // 列车类型ID
	SeatTypeId  any         // 座位类型ID
	BasePrice   any         // 基础票价
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
