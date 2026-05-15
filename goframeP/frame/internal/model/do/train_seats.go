// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TrainSeats is the golang structure of table train_seats for DAO operations like Where/Data.
type TrainSeats struct {
	g.Meta         `orm:"table:train_seats, do:true"`
	Id             any         //
	TrainId        any         //
	SeatTypeId     any         //
	Date           *gtime.Time //
	TotalCount     any         //
	AvailableCount any         //
	Price          any         //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
}
