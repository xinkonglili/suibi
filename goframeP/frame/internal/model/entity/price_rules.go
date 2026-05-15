// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PriceRules is the golang structure for table price_rules.
type PriceRules struct {
	Id          int64       `json:"id"          orm:"id"            description:""`       //
	TrainTypeId int         `json:"trainTypeId" orm:"train_type_id" description:"列车类型ID"` // 列车类型ID
	SeatTypeId  int         `json:"seatTypeId"  orm:"seat_type_id"  description:"座位类型ID"` // 座位类型ID
	BasePrice   float64     `json:"basePrice"   orm:"base_price"    description:"基础票价"`   // 基础票价
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"    description:""`       //
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"    description:""`       //
}
