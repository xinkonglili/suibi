// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RaiseHandTime is the golang structure of table raise_hand_time for DAO operations like Where/Data.
type RaiseHandTime struct {
	g.Meta        `orm:"table:raise_hand_time, do:true"`
	Id            any         //
	RaiseHandId   any         // 举手策略ID
	StartTime     *gtime.Time // 开始时间
	EndTime       *gtime.Time // 结束时间
	CreatedBy     any         // 创建人ID
	CreatedByName any         // 创建人名称
	UpdatedBy     any         // 更新人ID
	UpdatedByName any         // 更新人名称
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
	DeletedAt     *gtime.Time //
}
