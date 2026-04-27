// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RaiseHandTimeDao is the data access object for the frame raise_hand_time.
type RaiseHandTimeDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  RaiseHandTimeColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// RaiseHandTimeColumns defines and stores column names for the frame raise_hand_time.
type RaiseHandTimeColumns struct {
	Id            string //
	RaiseHandId   string // 举手策略ID
	StartTime     string // 开始时间
	EndTime       string // 结束时间
	CreatedBy     string // 创建人ID
	CreatedByName string // 创建人名称
	UpdatedBy     string // 更新人ID
	UpdatedByName string // 更新人名称
	CreatedAt     string //
	UpdatedAt     string //
	DeletedAt     string //
}

// raiseHandTimeColumns holds the columns for the frame raise_hand_time.
var raiseHandTimeColumns = RaiseHandTimeColumns{
	Id:            "id",
	RaiseHandId:   "raise_hand_id",
	StartTime:     "start_time",
	EndTime:       "end_time",
	CreatedBy:     "created_by",
	CreatedByName: "created_by_name",
	UpdatedBy:     "updated_by",
	UpdatedByName: "updated_by_name",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewRaiseHandTimeDao creates and returns a new DAO object for frame data access.
func NewRaiseHandTimeDao(handlers ...gdb.ModelHandler) *RaiseHandTimeDao {
	return &RaiseHandTimeDao{
		group:    "default",
		table:    "raise_hand_time",
		columns:  raiseHandTimeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RaiseHandTimeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the frame name of the current DAO.
func (dao *RaiseHandTimeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RaiseHandTimeDao) Columns() RaiseHandTimeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RaiseHandTimeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RaiseHandTimeDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *RaiseHandTimeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
