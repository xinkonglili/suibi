// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RaiseHandDao is the data access object for the frame raise_hand.
type RaiseHandDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  RaiseHandColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// RaiseHandColumns defines and stores column names for the frame raise_hand.
type RaiseHandColumns struct {
	Id              string //
	ChannelId       string // 渠道ID
	ChannelCode     string // 渠道code
	ProviderId      string // 服务方ID
	ProviderCode    string // 服务方code
	CityId          string // 城市ID
	CityName        string // 城市名称
	WaitingTime     string // 等待时长(单位：毫秒)
	OptimalDistance string // 最优接驾距离(单位：米)
	EffectiveTime   string // 生效时间 1-全天 2-指定
	EffectiveDate   string // 生效日期 1-全天 2-指定
	Status          string // 状态 1-开启 2-关闭
	CreatedBy       string // 创建人ID
	CreatedByName   string // 创建人名称
	UpdatedBy       string // 更新人ID
	UpdatedByName   string // 更新人名称
	CreatedAt       string //
	UpdatedAt       string //
	DeletedAt       string //
}

// raiseHandColumns holds the columns for the frame raise_hand.
var raiseHandColumns = RaiseHandColumns{
	Id:              "id",
	ChannelId:       "channel_id",
	ChannelCode:     "channel_code",
	ProviderId:      "provider_id",
	ProviderCode:    "provider_code",
	CityId:          "city_id",
	CityName:        "city_name",
	WaitingTime:     "waiting_time",
	OptimalDistance: "optimal_distance",
	EffectiveTime:   "effective_time",
	EffectiveDate:   "effective_date",
	Status:          "status",
	CreatedBy:       "created_by",
	CreatedByName:   "created_by_name",
	UpdatedBy:       "updated_by",
	UpdatedByName:   "updated_by_name",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
}

// NewRaiseHandDao creates and returns a new DAO object for frame data access.
func NewRaiseHandDao(handlers ...gdb.ModelHandler) *RaiseHandDao {
	return &RaiseHandDao{
		group:    "default",
		table:    "raise_hand",
		columns:  raiseHandColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RaiseHandDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the frame name of the current DAO.
func (dao *RaiseHandDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RaiseHandDao) Columns() RaiseHandColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RaiseHandDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RaiseHandDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *RaiseHandDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
