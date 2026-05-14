// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrderDao is the data access object for the table order.
type OrderDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  OrderColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// OrderColumns defines and stores column names for the table order.
type OrderColumns struct {
	Id                  string // 主键ID
	OrderNumber         string // switch主单号
	ToUpOrderNumber     string // switch给上游子单号
	ToDownOrderNumber   string // switch给下游子单号
	UpPartnerOrder      string // 上游单号
	UpPartnerSubOrder   string // 上游子单号
	DownPartnerOrder    string // 下游单号
	DownPartnerSubOrder string // 下游子单号
	OrderStatus         string // 订单状态 10: 未接单（正在派单）、20:已接单、30:订单进行中、40:订单完成、50订单取消
	FareStatus          string // 费用状态 10:未计费、11:司机已举手、30:计费中、31:待支付（计费结束）、40:已支付
	TripStatus          string // 行程状态: 10-无司机(待司机应答) 11-司机已举手(等待pk结果) 20-已接单(待出发) 21-出发接客中(预约单) 22-到达上车点 30-行程中 31-行程结束 70-改派中
	UpPartnerId         string // 上游ID
	UpPartnerCode       string // 上游code，冗余
	DownPartnerId       string // 下游ID
	DownPartnerCode     string // 下游code，冗余
	UserPhone           string // 下单人手机号（默认为乘客虚拟号，调用失败将使用真实号码)
	PassengerPhone      string // 乘客手机号
	PassengerHidePhone  string // 乘客手机号
	PassengePhoneSuffix string // 虚拟号
	OrderTime           string // 下单时间
	OrderFare           string // 上游订单费用（下游金额）
	PayFare             string // 商家实收金额（上游金额）
	UserPayFare         string // 用户支付费用（用户实付）
	EstimateFare        string // 预估费用
	OrderType           string // 订单类型 1:实时单 2:预约单
	ValuateType         string // 计价类型 1:行后支付 2:一口价 3:上游一口价
	ProductType         string // 产品类型 1: 专车 2:快车
	ServiceType         string // 运力类型 1:特惠 2:经济 3:舒适 4:优享 5:商务 6:豪华 7:快车 8:出租车 9:特惠省钱版
	Remark              string // 说明
	CreatedAt           string // 创建时间
	UpdatedAt           string // 更新时间
	DeletedAt           string // 更新时间
	PayTime             string // 订单支付时间
	YyOrderNumber       string // 约约订单号
	InvoiceAmount       string // 申请开票的金额
	TenantOrderId       string // 渠道订单号
	TenantId            string // 租户ID
	SettlementFare      string // 结算金额
	ServiceBrandId      string // 运力品牌id
	UpUserId            string // 上游用户ID
}

// orderColumns holds the columns for the table order.
var orderColumns = OrderColumns{
	Id:                  "id",
	OrderNumber:         "order_number",
	ToUpOrderNumber:     "to_up_order_number",
	ToDownOrderNumber:   "to_down_order_number",
	UpPartnerOrder:      "up_partner_order",
	UpPartnerSubOrder:   "up_partner_sub_order",
	DownPartnerOrder:    "down_partner_order",
	DownPartnerSubOrder: "down_partner_sub_order",
	OrderStatus:         "order_status",
	FareStatus:          "fare_status",
	TripStatus:          "trip_status",
	UpPartnerId:         "up_partner_id",
	UpPartnerCode:       "up_partner_code",
	DownPartnerId:       "down_partner_id",
	DownPartnerCode:     "down_partner_code",
	UserPhone:           "user_phone",
	PassengerPhone:      "passenger_phone",
	PassengerHidePhone:  "passenger_hide_phone",
	PassengePhoneSuffix: "passenge_phone_suffix",
	OrderTime:           "order_time",
	OrderFare:           "order_fare",
	PayFare:             "pay_fare",
	UserPayFare:         "user_pay_fare",
	EstimateFare:        "estimate_fare",
	OrderType:           "order_type",
	ValuateType:         "valuate_type",
	ProductType:         "product_type",
	ServiceType:         "service_type",
	Remark:              "remark",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	DeletedAt:           "deleted_at",
	PayTime:             "pay_time",
	YyOrderNumber:       "yy_order_number",
	InvoiceAmount:       "invoice_amount",
	TenantOrderId:       "tenant_order_id",
	TenantId:            "tenant_id",
	SettlementFare:      "settlement_fare",
	ServiceBrandId:      "service_brand_id",
	UpUserId:            "up_user_id",
}

// NewOrderDao creates and returns a new DAO object for table data access.
func NewOrderDao(handlers ...gdb.ModelHandler) *OrderDao {
	return &OrderDao{
		group:    "default",
		table:    "order",
		columns:  orderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OrderDao) Columns() OrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OrderDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *OrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
