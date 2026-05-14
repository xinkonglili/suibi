// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PartnerRelationshipDao is the data access object for the table partner_relationship.
type PartnerRelationshipDao struct {
	table    string                     // table is the underlying table name of the DAO.
	group    string                     // group is the database configuration group name of the current DAO.
	columns  PartnerRelationshipColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler         // handlers for customized model modification.
}

// PartnerRelationshipColumns defines and stores column names for the table partner_relationship.
type PartnerRelationshipColumns struct {
	Id              string // 主键ID
	UpPartnerId     string // 上游流量平台ID
	UpPartnerCode   string //
	DownPartnerId   string // 下游运力公司ID
	DownPartnerCode string //
	AreaCode        string // 地区(市code)，全部“all”
	ServiceType     string // 服务类型：0:全部 1:实时单 2:预约单
	ServiceTypeJson string // Switch车型映射上游车型
	Status          string // 状态(0:未知 1:启用 2:禁用)
	Remark          string // 说明
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
	DeletedAt       string // 更新时间
}

// partnerRelationshipColumns holds the columns for the table partner_relationship.
var partnerRelationshipColumns = PartnerRelationshipColumns{
	Id:              "id",
	UpPartnerId:     "up_partner_id",
	UpPartnerCode:   "up_partner_code",
	DownPartnerId:   "down_partner_id",
	DownPartnerCode: "down_partner_code",
	AreaCode:        "area_code",
	ServiceType:     "service_type",
	ServiceTypeJson: "service_type_json",
	Status:          "status",
	Remark:          "remark",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
}

// NewPartnerRelationshipDao creates and returns a new DAO object for table data access.
func NewPartnerRelationshipDao(handlers ...gdb.ModelHandler) *PartnerRelationshipDao {
	return &PartnerRelationshipDao{
		group:    "default",
		table:    "partner_relationship",
		columns:  partnerRelationshipColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PartnerRelationshipDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PartnerRelationshipDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PartnerRelationshipDao) Columns() PartnerRelationshipColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PartnerRelationshipDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PartnerRelationshipDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PartnerRelationshipDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
