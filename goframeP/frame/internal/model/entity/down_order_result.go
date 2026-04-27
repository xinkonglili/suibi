// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DownOrderResult is the golang structure for frame down_order_result.
type DownOrderResult struct {
	Id                    int64       `json:"id"                    orm:"id"                       description:"主键ID"`                                // 主键ID
	OrderId               int64       `json:"orderId"               orm:"order_id"                 description:"主单id"`                                // 主单id
	OrderNumber           string      `json:"orderNumber"           orm:"order_number"             description:"switch单号"`                            // switch单号
	ToDownOrderNumber     string      `json:"toDownOrderNumber"     orm:"to_down_order_number"     description:"switch给下游子单号"`                        // switch给下游子单号
	ToUpOrderNumber       string      `json:"toUpOrderNumber"       orm:"to_up_order_number"       description:"switch给上游子单号"`                        // switch给上游子单号
	PartnerId             int64       `json:"partnerId"             orm:"partner_id"               description:"业务方ID"`                               // 业务方ID
	TenantOrderId         string      `json:"tenantOrderId"         orm:"tenant_order_id"          description:"下游渠道ID"`                              // 下游渠道ID
	TenantId              string      `json:"tenantId"              orm:"tenant_id"                description:"租户ID"`                                // 租户ID
	BizCode               string      `json:"bizCode"               orm:"biz_code"                 description:"业务方编码"`                               // 业务方编码
	PartnerOrderNumber    string      `json:"partnerOrderNumber"    orm:"partner_order_number"     description:"业务方单号"`                               // 业务方单号
	PartnerSubOrderNumber string      `json:"partnerSubOrderNumber" orm:"partner_sub_order_number" description:"业务方子单号"`                              // 业务方子单号
	EstimateMileage       int         `json:"estimateMileage"       orm:"estimate_mileage"         description:"预估公里数"`                               // 预估公里数
	EstimateDuration      int         `json:"estimateDuration"      orm:"estimate_duration"        description:"预估耗时"`                                // 预估耗时
	EstimateFare          float64     `json:"estimateFare"          orm:"estimate_fare"            description:"预估费用"`                                // 预估费用
	Status                int         `json:"status"                orm:"status"                   description:"状态 0: 未知，1:创建订单，2:pk成功 3:订单取消 4:已举手"` // 状态 0: 未知，1:创建订单，2:pk成功 3:订单取消 4:已举手
	Remark                string      `json:"remark"                orm:"remark"                   description:"说明"`                                  // 说明
	CreatedAt             *gtime.Time `json:"createdAt"             orm:"created_at"               description:"创建时间"`                                // 创建时间
	UpdatedAt             *gtime.Time `json:"updatedAt"             orm:"updated_at"               description:"更新时间"`                                // 更新时间
	DeletedAt             *gtime.Time `json:"deletedAt"             orm:"deleted_at"               description:"更新时间"`                                // 更新时间
}
