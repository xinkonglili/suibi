// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Orders is the golang structure for table orders.
type Orders struct {
	Id              int64       `json:"id"              orm:"id"                description:""` //
	OrderNo         string      `json:"orderNo"         orm:"order_no"          description:""` //
	UserId          int64       `json:"userId"          orm:"user_id"           description:""` //
	TrainId         int64       `json:"trainId"         orm:"train_id"          description:""` //
	SeatTypeId      int         `json:"seatTypeId"      orm:"seat_type_id"      description:""` //
	Date            *gtime.Time `json:"date"            orm:"date"              description:""` //
	PassengerName   string      `json:"passengerName"   orm:"passenger_name"    description:""` //
	PassengerIdCard string      `json:"passengerIdCard" orm:"passenger_id_card" description:""` //
	Price           float64     `json:"price"           orm:"price"             description:""` //
	Status          string      `json:"status"          orm:"status"            description:""` //
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"        description:""` //
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"        description:""` //
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"        description:""` //
}
