// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TrainsDao is the data access object for the table trains.
type TrainsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TrainsColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TrainsColumns defines and stores column names for the table trains.
type TrainsColumns struct {
	Id             string //
	TrainNo        string //
	TrainType      string //
	StartStationId string //
	EndStationId   string //
	DepartureTime  string //
	ArrivalTime    string //
	Duration       string //
	CreatedAt      string //
}

// trainsColumns holds the columns for the table trains.
var trainsColumns = TrainsColumns{
	Id:             "id",
	TrainNo:        "train_no",
	TrainType:      "train_type",
	StartStationId: "start_station_id",
	EndStationId:   "end_station_id",
	DepartureTime:  "departure_time",
	ArrivalTime:    "arrival_time",
	Duration:       "duration",
	CreatedAt:      "created_at",
}

// NewTrainsDao creates and returns a new DAO object for table data access.
func NewTrainsDao(handlers ...gdb.ModelHandler) *TrainsDao {
	return &TrainsDao{
		group:    "default",
		table:    "trains",
		columns:  trainsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TrainsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TrainsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TrainsDao) Columns() TrainsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TrainsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TrainsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TrainsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
