// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Order is the golang structure for table order.
type Order struct {
	Id                  int64       `json:"id"                  orm:"id"                     description:"主键ID"`                                                                                          // 主键ID
	OrderNumber         string      `json:"orderNumber"         orm:"order_number"           description:"switch主单号"`                                                                                     // switch主单号
	ToUpOrderNumber     string      `json:"toUpOrderNumber"     orm:"to_up_order_number"     description:"switch给上游子单号"`                                                                                  // switch给上游子单号
	ToDownOrderNumber   string      `json:"toDownOrderNumber"   orm:"to_down_order_number"   description:"switch给下游子单号"`                                                                                  // switch给下游子单号
	UpPartnerOrder      string      `json:"upPartnerOrder"      orm:"up_partner_order"       description:"上游单号"`                                                                                          // 上游单号
	UpPartnerSubOrder   string      `json:"upPartnerSubOrder"   orm:"up_partner_sub_order"   description:"上游子单号"`                                                                                         // 上游子单号
	DownPartnerOrder    string      `json:"downPartnerOrder"    orm:"down_partner_order"     description:"下游单号"`                                                                                          // 下游单号
	DownPartnerSubOrder string      `json:"downPartnerSubOrder" orm:"down_partner_sub_order" description:"下游子单号"`                                                                                         // 下游子单号
	OrderStatus         int         `json:"orderStatus"         orm:"order_status"           description:"订单状态 10: 未接单（正在派单）、20:已接单、30:订单进行中、40:订单完成、50订单取消"`                                             // 订单状态 10: 未接单（正在派单）、20:已接单、30:订单进行中、40:订单完成、50订单取消
	FareStatus          int         `json:"fareStatus"          orm:"fare_status"            description:"费用状态 10:未计费、11:司机已举手、30:计费中、31:待支付（计费结束）、40:已支付"`                                               // 费用状态 10:未计费、11:司机已举手、30:计费中、31:待支付（计费结束）、40:已支付
	TripStatus          int         `json:"tripStatus"          orm:"trip_status"            description:"行程状态: 10-无司机(待司机应答) 11-司机已举手(等待pk结果) 20-已接单(待出发) 21-出发接客中(预约单) 22-到达上车点 30-行程中 31-行程结束 70-改派中"` // 行程状态: 10-无司机(待司机应答) 11-司机已举手(等待pk结果) 20-已接单(待出发) 21-出发接客中(预约单) 22-到达上车点 30-行程中 31-行程结束 70-改派中
	UpPartnerId         int64       `json:"upPartnerId"         orm:"up_partner_id"          description:"上游ID"`                                                                                          // 上游ID
	UpPartnerCode       string      `json:"upPartnerCode"       orm:"up_partner_code"        description:"上游code，冗余"`                                                                                     // 上游code，冗余
	DownPartnerId       int64       `json:"downPartnerId"       orm:"down_partner_id"        description:"下游ID"`                                                                                          // 下游ID
	DownPartnerCode     string      `json:"downPartnerCode"     orm:"down_partner_code"      description:"下游code，冗余"`                                                                                     // 下游code，冗余
	UserPhone           string      `json:"userPhone"           orm:"user_phone"             description:"下单人手机号（默认为乘客虚拟号，调用失败将使用真实号码)"`                                                                  // 下单人手机号（默认为乘客虚拟号，调用失败将使用真实号码)
	PassengerPhone      string      `json:"passengerPhone"      orm:"passenger_phone"        description:"乘客手机号"`                                                                                         // 乘客手机号
	PassengerHidePhone  string      `json:"passengerHidePhone"  orm:"passenger_hide_phone"   description:"乘客手机号"`                                                                                         // 乘客手机号
	PassengePhoneSuffix string      `json:"passengePhoneSuffix" orm:"passenge_phone_suffix"  description:"虚拟号"`                                                                                           // 虚拟号
	OrderTime           *gtime.Time `json:"orderTime"           orm:"order_time"             description:"下单时间"`                                                                                          // 下单时间
	OrderFare           float64     `json:"orderFare"           orm:"order_fare"             description:"上游订单费用（下游金额）"`                                                                                  // 上游订单费用（下游金额）
	PayFare             float64     `json:"payFare"             orm:"pay_fare"               description:"商家实收金额（上游金额）"`                                                                                  // 商家实收金额（上游金额）
	UserPayFare         float64     `json:"userPayFare"         orm:"user_pay_fare"          description:"用户支付费用（用户实付）"`                                                                                  // 用户支付费用（用户实付）
	EstimateFare        float64     `json:"estimateFare"        orm:"estimate_fare"          description:"预估费用"`                                                                                          // 预估费用
	OrderType           int         `json:"orderType"           orm:"order_type"             description:"订单类型 1:实时单 2:预约单"`                                                                              // 订单类型 1:实时单 2:预约单
	ValuateType         int         `json:"valuateType"         orm:"valuate_type"           description:"计价类型 1:行后支付 2:一口价 3:上游一口价"`                                                                     // 计价类型 1:行后支付 2:一口价 3:上游一口价
	ProductType         int         `json:"productType"         orm:"product_type"           description:"产品类型 1: 专车 2:快车"`                                                                               // 产品类型 1: 专车 2:快车
	ServiceType         int         `json:"serviceType"         orm:"service_type"           description:"运力类型 1:特惠 2:经济 3:舒适 4:优享 5:商务 6:豪华 7:快车 8:出租车 9:特惠省钱版"`                                         // 运力类型 1:特惠 2:经济 3:舒适 4:优享 5:商务 6:豪华 7:快车 8:出租车 9:特惠省钱版
	Remark              string      `json:"remark"              orm:"remark"                 description:"说明"`                                                                                            // 说明
	CreatedAt           *gtime.Time `json:"createdAt"           orm:"created_at"             description:"创建时间"`                                                                                          // 创建时间
	UpdatedAt           *gtime.Time `json:"updatedAt"           orm:"updated_at"             description:"更新时间"`                                                                                          // 更新时间
	DeletedAt           *gtime.Time `json:"deletedAt"           orm:"deleted_at"             description:"更新时间"`                                                                                          // 更新时间
	PayTime             *gtime.Time `json:"payTime"             orm:"pay_time"               description:"订单支付时间"`                                                                                        // 订单支付时间
	YyOrderNumber       string      `json:"yyOrderNumber"       orm:"yy_order_number"        description:"约约订单号"`                                                                                         // 约约订单号
	InvoiceAmount       float64     `json:"invoiceAmount"       orm:"invoice_amount"         description:"申请开票的金额"`                                                                                       // 申请开票的金额
	TenantOrderId       string      `json:"tenantOrderId"       orm:"tenant_order_id"        description:"渠道订单号"`                                                                                         // 渠道订单号
	TenantId            string      `json:"tenantId"            orm:"tenant_id"              description:"租户ID"`                                                                                          // 租户ID
	SettlementFare      float64     `json:"settlementFare"      orm:"settlement_fare"        description:"结算金额"`                                                                                          // 结算金额
	ServiceBrandId      string      `json:"serviceBrandId"      orm:"service_brand_id"       description:"运力品牌id"`                                                                                        // 运力品牌id
	UpUserId            string      `json:"upUserId"            orm:"up_user_id"             description:"上游用户ID"`                                                                                        // 上游用户ID
}
