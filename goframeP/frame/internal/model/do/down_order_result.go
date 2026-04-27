// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DownOrderResult is the golang structure of frame down_order_result for DAO operations like Where/Data.
type DownOrderResult struct {
	g.Meta                `orm:"frame:down_order_result, do:true"`
	Id                    any         // 主键ID
	OrderId               any         // 主单id
	OrderNumber           any         // switch单号
	ToDownOrderNumber     any         // switch给下游子单号
	ToUpOrderNumber       any         // switch给上游子单号
	PartnerId             any         // 业务方ID
	TenantOrderId         any         // 下游渠道ID
	TenantId              any         // 租户ID
	BizCode               any         // 业务方编码
	PartnerOrderNumber    any         // 业务方单号
	PartnerSubOrderNumber any         // 业务方子单号
	EstimateMileage       any         // 预估公里数
	EstimateDuration      any         // 预估耗时
	EstimateFare          any         // 预估费用
	Status                any         // 状态 0: 未知，1:创建订单，2:pk成功 3:订单取消 4:已举手
	Remark                any         // 说明
	CreatedAt             *gtime.Time // 创建时间
	UpdatedAt             *gtime.Time // 更新时间
	DeletedAt             *gtime.Time // 更新时间
}
