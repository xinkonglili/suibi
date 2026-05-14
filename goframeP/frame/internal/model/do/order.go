// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Order is the golang structure of table order for DAO operations like Where/Data.
type Order struct {
	g.Meta              `orm:"table:order, do:true"`
	Id                  any         // 主键ID
	OrderNumber         any         // switch主单号
	ToUpOrderNumber     any         // switch给上游子单号
	ToDownOrderNumber   any         // switch给下游子单号
	UpPartnerOrder      any         // 上游单号
	UpPartnerSubOrder   any         // 上游子单号
	DownPartnerOrder    any         // 下游单号
	DownPartnerSubOrder any         // 下游子单号
	OrderStatus         any         // 订单状态 10: 未接单（正在派单）、20:已接单、30:订单进行中、40:订单完成、50订单取消
	FareStatus          any         // 费用状态 10:未计费、11:司机已举手、30:计费中、31:待支付（计费结束）、40:已支付
	TripStatus          any         // 行程状态: 10-无司机(待司机应答) 11-司机已举手(等待pk结果) 20-已接单(待出发) 21-出发接客中(预约单) 22-到达上车点 30-行程中 31-行程结束 70-改派中
	UpPartnerId         any         // 上游ID
	UpPartnerCode       any         // 上游code，冗余
	DownPartnerId       any         // 下游ID
	DownPartnerCode     any         // 下游code，冗余
	UserPhone           any         // 下单人手机号（默认为乘客虚拟号，调用失败将使用真实号码)
	PassengerPhone      any         // 乘客手机号
	PassengerHidePhone  any         // 乘客手机号
	PassengePhoneSuffix any         // 虚拟号
	OrderTime           *gtime.Time // 下单时间
	OrderFare           any         // 上游订单费用（下游金额）
	PayFare             any         // 商家实收金额（上游金额）
	UserPayFare         any         // 用户支付费用（用户实付）
	EstimateFare        any         // 预估费用
	OrderType           any         // 订单类型 1:实时单 2:预约单
	ValuateType         any         // 计价类型 1:行后支付 2:一口价 3:上游一口价
	ProductType         any         // 产品类型 1: 专车 2:快车
	ServiceType         any         // 运力类型 1:特惠 2:经济 3:舒适 4:优享 5:商务 6:豪华 7:快车 8:出租车 9:特惠省钱版
	Remark              any         // 说明
	CreatedAt           *gtime.Time // 创建时间
	UpdatedAt           *gtime.Time // 更新时间
	DeletedAt           *gtime.Time // 更新时间
	PayTime             *gtime.Time // 订单支付时间
	YyOrderNumber       any         // 约约订单号
	InvoiceAmount       any         // 申请开票的金额
	TenantOrderId       any         // 渠道订单号
	TenantId            any         // 租户ID
	SettlementFare      any         // 结算金额
	ServiceBrandId      any         // 运力品牌id
	UpUserId            any         // 上游用户ID
}
