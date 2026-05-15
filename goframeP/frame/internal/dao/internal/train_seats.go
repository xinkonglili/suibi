// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TrainSeatsDao is the data access object for the table train_seats.
type TrainSeatsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TrainSeatsColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TrainSeatsColumns defines and stores column names for the table train_seats.
type TrainSeatsColumns struct {
	Id             string //
	TrainId        string //
	SeatTypeId     string //
	Date           string //
	TotalCount     string //
	AvailableCount string //
	Price          string //
	CreatedAt      string //
	UpdatedAt      string //
}

// trainSeatsColumns holds the columns for the table train_seats.
var trainSeatsColumns = TrainSeatsColumns{
	Id:             "id",
	TrainId:        "train_id",
	SeatTypeId:     "seat_type_id",
	Date:           "date",
	TotalCount:     "total_count",
	AvailableCount: "available_count",
	Price:          "price",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewTrainSeatsDao creates and returns a new DAO object for table data access.
func NewTrainSeatsDao(handlers ...gdb.ModelHandler) *TrainSeatsDao {
	return &TrainSeatsDao{
		group:    "default",
		table:    "train_seats",
		columns:  trainSeatsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TrainSeatsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TrainSeatsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TrainSeatsDao) Columns() TrainSeatsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TrainSeatsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TrainSeatsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TrainSeatsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
