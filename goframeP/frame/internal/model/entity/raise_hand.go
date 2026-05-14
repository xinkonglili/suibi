// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RaiseHand is the golang structure for table raise_hand.
type RaiseHand struct {
	Id              uint64      `json:"id"              orm:"id"               description:""`               //
	ChannelId       uint64      `json:"channelId"       orm:"channel_id"       description:"渠道ID"`           // 渠道ID
	ChannelCode     string      `json:"channelCode"     orm:"channel_code"     description:"渠道code"`         // 渠道code
	ProviderId      uint64      `json:"providerId"      orm:"provider_id"      description:"服务方ID"`          // 服务方ID
	ProviderCode    string      `json:"providerCode"    orm:"provider_code"    description:"服务方code"`        // 服务方code
	CityId          uint64      `json:"cityId"          orm:"city_id"          description:"城市ID"`           // 城市ID
	CityName        string      `json:"cityName"        orm:"city_name"        description:"城市名称"`           // 城市名称
	WaitingTime     uint        `json:"waitingTime"     orm:"waiting_time"     description:"等待时长(单位：毫秒)"`    // 等待时长(单位：毫秒)
	OptimalDistance uint        `json:"optimalDistance" orm:"optimal_distance" description:"最优接驾距离(单位：米)"`   // 最优接驾距离(单位：米)
	EffectiveTime   uint        `json:"effectiveTime"   orm:"effective_time"   description:"生效时间 1-全天 2-指定"` // 生效时间 1-全天 2-指定
	EffectiveDate   uint        `json:"effectiveDate"   orm:"effective_date"   description:"生效日期 1-全天 2-指定"` // 生效日期 1-全天 2-指定
	Status          uint        `json:"status"          orm:"status"           description:"状态 1-开启 2-关闭"`   // 状态 1-开启 2-关闭
	CreatedBy       uint64      `json:"createdBy"       orm:"created_by"       description:"创建人ID"`          // 创建人ID
	CreatedByName   string      `json:"createdByName"   orm:"created_by_name"  description:"创建人名称"`          // 创建人名称
	UpdatedBy       uint64      `json:"updatedBy"       orm:"updated_by"       description:"更新人ID"`          // 更新人ID
	UpdatedByName   string      `json:"updatedByName"   orm:"updated_by_name"  description:"更新人名称"`          // 更新人名称
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       description:""`               //
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       description:""`               //
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"       description:""`               //
}
