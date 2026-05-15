// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TrainOrderDao is the data access object for the table train_order.
type TrainOrderDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TrainOrderColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TrainOrderColumns defines and stores column names for the table train_order.
type TrainOrderColumns struct {
	Id       string //
	UserId   string // 用户id
	OrderId  string //
	UserName string //
}

// trainOrderColumns holds the columns for the table train_order.
var trainOrderColumns = TrainOrderColumns{
	Id:       "id",
	UserId:   "user_id",
	OrderId:  "order_id",
	UserName: "user_name",
}

// NewTrainOrderDao creates and returns a new DAO object for table data access.
func NewTrainOrderDao(handlers ...gdb.ModelHandler) *TrainOrderDao {
	return &TrainOrderDao{
		group:    "default",
		table:    "train_order",
		columns:  trainOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TrainOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TrainOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TrainOrderDao) Columns() TrainOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TrainOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TrainOrderDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TrainOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
