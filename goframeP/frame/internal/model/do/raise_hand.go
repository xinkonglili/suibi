// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RaiseHand is the golang structure of frame raise_hand for DAO operations like Where/Data.
type RaiseHand struct {
	g.Meta          `orm:"frame:raise_hand, do:true"`
	Id              any         //
	ChannelId       any         // 渠道ID
	ChannelCode     any         // 渠道code
	ProviderId      any         // 服务方ID
	ProviderCode    any         // 服务方code
	CityId          any         // 城市ID
	CityName        any         // 城市名称
	WaitingTime     any         // 等待时长(单位：毫秒)
	OptimalDistance any         // 最优接驾距离(单位：米)
	EffectiveTime   any         // 生效时间 1-全天 2-指定
	EffectiveDate   any         // 生效日期 1-全天 2-指定
	Status          any         // 状态 1-开启 2-关闭
	CreatedBy       any         // 创建人ID
	CreatedByName   any         // 创建人名称
	UpdatedBy       any         // 更新人ID
	UpdatedByName   any         // 更新人名称
	CreatedAt       *gtime.Time //
	UpdatedAt       *gtime.Time //
	DeletedAt       *gtime.Time //
}
