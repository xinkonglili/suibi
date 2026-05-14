// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TagStatusDao is the data access object for the table tag_status.
type TagStatusDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TagStatusColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TagStatusColumns defines and stores column names for the table tag_status.
type TagStatusColumns struct {
	Kaiguan string //
	Status  string //
	Tag     string //
	Id      string //
}

// tagStatusColumns holds the columns for the table tag_status.
var tagStatusColumns = TagStatusColumns{
	Kaiguan: "kaiguan",
	Status:  "status",
	Tag:     "tag",
	Id:      "id",
}

// NewTagStatusDao creates and returns a new DAO object for table data access.
func NewTagStatusDao(handlers ...gdb.ModelHandler) *TagStatusDao {
	return &TagStatusDao{
		group:    "default",
		table:    "tag_status",
		columns:  tagStatusColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TagStatusDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TagStatusDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TagStatusDao) Columns() TagStatusColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TagStatusDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TagStatusDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TagStatusDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
