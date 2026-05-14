// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PartnerRelationship is the golang structure of table partner_relationship for DAO operations like Where/Data.
type PartnerRelationship struct {
	g.Meta          `orm:"table:partner_relationship, do:true"`
	Id              any         // 主键ID
	UpPartnerId     any         // 上游流量平台ID
	UpPartnerCode   any         //
	DownPartnerId   any         // 下游运力公司ID
	DownPartnerCode any         //
	AreaCode        any         // 地区(市code)，全部“all”
	ServiceType     any         // 服务类型：0:全部 1:实时单 2:预约单
	ServiceTypeJson any         // Switch车型映射上游车型
	Status          any         // 状态(0:未知 1:启用 2:禁用)
	Remark          any         // 说明
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
	DeletedAt       *gtime.Time // 更新时间
}
