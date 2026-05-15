// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Trains is the golang structure for table trains.
type Trains struct {
	Id             int64       `json:"id"             orm:"id"               description:""` //
	TrainNo        string      `json:"trainNo"        orm:"train_no"         description:""` //
	TrainType      string      `json:"trainType"      orm:"train_type"       description:""` //
	StartStationId int         `json:"startStationId" orm:"start_station_id" description:""` //
	EndStationId   int         `json:"endStationId"   orm:"end_station_id"   description:""` //
	DepartureTime  *gtime.Time `json:"departureTime"  orm:"departure_time"   description:""` //
	ArrivalTime    *gtime.Time `json:"arrivalTime"    orm:"arrival_time"     description:""` //
	Duration       int         `json:"duration"       orm:"duration"         description:""` //
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:""` //
}
