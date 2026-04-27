// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RaiseHandDate is the golang structure of frame raise_hand_date for DAO operations like Where/Data.
type RaiseHandDate struct {
	g.Meta        `orm:"frame:raise_hand_date, do:true"`
	Id            any         //
	RaiseHandId   any         // 举手策略ID
	StartDate     *gtime.Time // 开始日期
	EndDate       *gtime.Time // 结束日期
	CreatedBy     any         // 创建人ID
	CreatedByName any         // 创建人名称
	UpdatedBy     any         // 更新人ID
	UpdatedByName any         // 更新人名称
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
	DeletedAt     *gtime.Time //
}
