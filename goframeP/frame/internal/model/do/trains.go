// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Trains is the golang structure of table trains for DAO operations like Where/Data.
type Trains struct {
	g.Meta         `orm:"table:trains, do:true"`
	Id             any         //
	TrainNo        any         //
	TrainType      any         //
	StartStationId any         //
	EndStationId   any         //
	DepartureTime  *gtime.Time //
	ArrivalTime    *gtime.Time //
	Duration       any         //
	CreatedAt      *gtime.Time //
}
