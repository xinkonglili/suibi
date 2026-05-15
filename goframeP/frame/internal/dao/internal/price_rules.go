// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PriceRulesDao is the data access object for the table price_rules.
type PriceRulesDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PriceRulesColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PriceRulesColumns defines and stores column names for the table price_rules.
type PriceRulesColumns struct {
	Id          string //
	TrainTypeId string // 列车类型ID
	SeatTypeId  string // 座位类型ID
	BasePrice   string // 基础票价
	CreatedAt   string //
	UpdatedAt   string //
}

// priceRulesColumns holds the columns for the table price_rules.
var priceRulesColumns = PriceRulesColumns{
	Id:          "id",
	TrainTypeId: "train_type_id",
	SeatTypeId:  "seat_type_id",
	BasePrice:   "base_price",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewPriceRulesDao creates and returns a new DAO object for table data access.
func NewPriceRulesDao(handlers ...gdb.ModelHandler) *PriceRulesDao {
	return &PriceRulesDao{
		group:    "default",
		table:    "price_rules",
		columns:  priceRulesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PriceRulesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PriceRulesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PriceRulesDao) Columns() PriceRulesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PriceRulesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PriceRulesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PriceRulesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
