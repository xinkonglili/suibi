// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PartnerRelationship is the golang structure for frame partner_relationship.
type PartnerRelationship struct {
	Id              int64       `json:"id"              orm:"id"                description:"主键ID"`                  // 主键ID
	UpPartnerId     int64       `json:"upPartnerId"     orm:"up_partner_id"     description:"上游流量平台ID"`              // 上游流量平台ID
	UpPartnerCode   string      `json:"upPartnerCode"   orm:"up_partner_code"   description:""`                      //
	DownPartnerId   int64       `json:"downPartnerId"   orm:"down_partner_id"   description:"下游运力公司ID"`              // 下游运力公司ID
	DownPartnerCode string      `json:"downPartnerCode" orm:"down_partner_code" description:""`                      //
	AreaCode        string      `json:"areaCode"        orm:"area_code"         description:"地区(市code)，全部“all”"`     // 地区(市code)，全部“all”
	ServiceType     int         `json:"serviceType"     orm:"service_type"      description:"服务类型：0:全部 1:实时单 2:预约单"` // 服务类型：0:全部 1:实时单 2:预约单
	ServiceTypeJson string      `json:"serviceTypeJson" orm:"service_type_json" description:"Switch车型映射上游车型"`        // Switch车型映射上游车型
	Status          int         `json:"status"          orm:"status"            description:"状态(0:未知 1:启用 2:禁用)"`    // 状态(0:未知 1:启用 2:禁用)
	Remark          string      `json:"remark"          orm:"remark"            description:"说明"`                    // 说明
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"        description:"创建时间"`                  // 创建时间
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"        description:"更新时间"`                  // 更新时间
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"        description:"更新时间"`                  // 更新时间
}
