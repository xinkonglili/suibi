// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PartnerCityConfig is the golang structure for table partner_city_config.
type PartnerCityConfig struct {
	Id             int64       `json:"id"             orm:"id"               description:"主键ID"`             // 主键ID
	PartnerId      int64       `json:"partnerId"      orm:"partner_id"       description:"渠道ID"`             // 渠道ID
	BizCode        string      `json:"bizCode"        orm:"biz_code"         description:"业务方编码"`            // 业务方编码
	AreaCode       string      `json:"areaCode"       orm:"area_code"        description:"地区(市) 全部：all"`     // 地区(市) 全部：all
	AreaName       string      `json:"areaName"       orm:"area_name"        description:"地区名称"`             // 地区名称
	ConfigType     string      `json:"configType"     orm:"config_type"      description:"配置项"`              // 配置项
	ConfigValue    string      `json:"configValue"    orm:"config_value"     description:""`                 //
	StartTime      *gtime.Time `json:"startTime"      orm:"start_time"       description:""`                 //
	EndTime        *gtime.Time `json:"endTime"        orm:"end_time"         description:""`                 //
	RangeStartTime string      `json:"rangeStartTime" orm:"range_start_time" description:"时间范围-起  00:00:00"` // 时间范围-起  00:00:00
	RangeEndTime   string      `json:"rangeEndTime"   orm:"range_end_time"   description:"时间范围-止  23:59:59"` // 时间范围-止  23:59:59
	Percent        uint        `json:"percent"        orm:"percent"          description:"命中率"`              // 命中率
	Status         int         `json:"status"         orm:"status"           description:"状态 1:启用2:禁用"`      // 状态 1:启用2:禁用
	Remark         string      `json:"remark"         orm:"remark"           description:"说明"`               // 说明
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:"创建时间"`             // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:"更新时间"`             // 更新时间
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"       description:"更新时间"`             // 更新时间
	AddTest        string      `json:"addTest"        orm:"add_test"         description:"测试字段类型"`           // 测试字段类型
	AddTest0       string      `json:"addTest0"       orm:"add_test_0"       description:"测试字段类型"`           // 测试字段类型
	AddTest1       string      `json:"addTest1"       orm:"add_test_1"       description:"测试字段类型"`           // 测试字段类型
}
