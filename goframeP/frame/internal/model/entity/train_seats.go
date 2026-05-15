// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TrainSeats is the golang structure for table train_seats.
type TrainSeats struct {
	Id             int64       `json:"id"             orm:"id"              description:""` //
	TrainId        int64       `json:"trainId"        orm:"train_id"        description:""` //
	SeatTypeId     int         `json:"seatTypeId"     orm:"seat_type_id"    description:""` //
	Date           *gtime.Time `json:"date"           orm:"date"            description:""` //
	TotalCount     int         `json:"totalCount"     orm:"total_count"     description:""` //
	AvailableCount int         `json:"availableCount" orm:"available_count" description:""` //
	Price          float64     `json:"price"          orm:"price"           description:""` //
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:""` //
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:""` //
}
