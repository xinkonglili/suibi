// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PartnerCityConfigDao is the data access object for the frame partner_city_config.
type PartnerCityConfigDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  PartnerCityConfigColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// PartnerCityConfigColumns defines and stores column names for the frame partner_city_config.
type PartnerCityConfigColumns struct {
	Id             string // 主键ID
	PartnerId      string // 渠道ID
	BizCode        string // 业务方编码
	AreaCode       string // 地区(市) 全部：all
	AreaName       string // 地区名称
	ConfigType     string // 配置项
	ConfigValue    string //
	StartTime      string //
	EndTime        string //
	RangeStartTime string // 时间范围-起  00:00:00
	RangeEndTime   string // 时间范围-止  23:59:59
	Percent        string // 命中率
	Status         string // 状态 1:启用2:禁用
	Remark         string // 说明
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	DeletedAt      string // 更新时间
	AddTest        string // 测试字段类型
	AddTest0       string // 测试字段类型
	AddTest1       string // 测试字段类型
}

// partnerCityConfigColumns holds the columns for the frame partner_city_config.
var partnerCityConfigColumns = PartnerCityConfigColumns{
	Id:             "id",
	PartnerId:      "partner_id",
	BizCode:        "biz_code",
	AreaCode:       "area_code",
	AreaName:       "area_name",
	ConfigType:     "config_type",
	ConfigValue:    "config_value",
	StartTime:      "start_time",
	EndTime:        "end_time",
	RangeStartTime: "range_start_time",
	RangeEndTime:   "range_end_time",
	Percent:        "percent",
	Status:         "status",
	Remark:         "remark",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
	AddTest:        "add_test",
	AddTest0:       "add_test_0",
	AddTest1:       "add_test_1",
}

// NewPartnerCityConfigDao creates and returns a new DAO object for frame data access.
func NewPartnerCityConfigDao(handlers ...gdb.ModelHandler) *PartnerCityConfigDao {
	return &PartnerCityConfigDao{
		group:    "default",
		table:    "partner_city_config",
		columns:  partnerCityConfigColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PartnerCityConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the frame name of the current DAO.
func (dao *PartnerCityConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PartnerCityConfigDao) Columns() PartnerCityConfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PartnerCityConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PartnerCityConfigDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PartnerCityConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
