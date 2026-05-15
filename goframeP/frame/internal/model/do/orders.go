// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Orders is the golang structure of table orders for DAO operations like Where/Data.
type Orders struct {
	g.Meta          `orm:"table:orders, do:true"`
	Id              any         //
	OrderNo         any         //
	UserId          any         //
	TrainId         any         //
	SeatTypeId      any         //
	Date            *gtime.Time //
	PassengerName   any         //
	PassengerIdCard any         //
	Price           any         //
	Status          any         //
	CreatedAt       *gtime.Time //
	UpdatedAt       *gtime.Time //
	DeletedAt       *gtime.Time //
}
