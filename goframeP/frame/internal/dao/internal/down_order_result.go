// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DownOrderResultDao is the data access object for the table down_order_result.
type DownOrderResultDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  DownOrderResultColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// DownOrderResultColumns defines and stores column names for the table down_order_result.
type DownOrderResultColumns struct {
	Id                    string // 主键ID
	OrderId               string // 主单id
	OrderNumber           string // switch单号
	ToDownOrderNumber     string // switch给下游子单号
	ToUpOrderNumber       string // switch给上游子单号
	PartnerId             string // 业务方ID
	TenantOrderId         string // 下游渠道ID
	TenantId              string // 租户ID
	BizCode               string // 业务方编码
	PartnerOrderNumber    string // 业务方单号
	PartnerSubOrderNumber string // 业务方子单号
	EstimateMileage       string // 预估公里数
	EstimateDuration      string // 预估耗时
	EstimateFare          string // 预估费用
	Status                string // 状态 0: 未知，1:创建订单，2:pk成功 3:订单取消 4:已举手
	Remark                string // 说明
	CreatedAt             string // 创建时间
	UpdatedAt             string // 更新时间
	DeletedAt             string // 更新时间
}

// downOrderResultColumns holds the columns for the table down_order_result.
var downOrderResultColumns = DownOrderResultColumns{
	Id:                    "id",
	OrderId:               "order_id",
	OrderNumber:           "order_number",
	ToDownOrderNumber:     "to_down_order_number",
	ToUpOrderNumber:       "to_up_order_number",
	PartnerId:             "partner_id",
	TenantOrderId:         "tenant_order_id",
	TenantId:              "tenant_id",
	BizCode:               "biz_code",
	PartnerOrderNumber:    "partner_order_number",
	PartnerSubOrderNumber: "partner_sub_order_number",
	EstimateMileage:       "estimate_mileage",
	EstimateDuration:      "estimate_duration",
	EstimateFare:          "estimate_fare",
	Status:                "status",
	Remark:                "remark",
	CreatedAt:             "created_at",
	UpdatedAt:             "updated_at",
	DeletedAt:             "deleted_at",
}

// NewDownOrderResultDao creates and returns a new DAO object for table data access.
func NewDownOrderResultDao(handlers ...gdb.ModelHandler) *DownOrderResultDao {
	return &DownOrderResultDao{
		group:    "default",
		table:    "down_order_result",
		columns:  downOrderResultColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DownOrderResultDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DownOrderResultDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DownOrderResultDao) Columns() DownOrderResultColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DownOrderResultDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DownOrderResultDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *DownOrderResultDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
