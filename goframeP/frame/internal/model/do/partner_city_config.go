// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PartnerCityConfig is the golang structure of table partner_city_config for DAO operations like Where/Data.
type PartnerCityConfig struct {
	g.Meta         `orm:"table:partner_city_config, do:true"`
	Id             any         // 主键ID
	PartnerId      any         // 渠道ID
	BizCode        any         // 业务方编码
	AreaCode       any         // 地区(市) 全部：all
	AreaName       any         // 地区名称
	ConfigType     any         // 配置项
	ConfigValue    any         //
	StartTime      *gtime.Time //
	EndTime        *gtime.Time //
	RangeStartTime any         // 时间范围-起  00:00:00
	RangeEndTime   any         // 时间范围-止  23:59:59
	Percent        any         // 命中率
	Status         any         // 状态 1:启用2:禁用
	Remark         any         // 说明
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 更新时间
	AddTest        any         // 测试字段类型
	AddTest0       any         // 测试字段类型
	AddTest1       any         // 测试字段类型
}
