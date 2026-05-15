// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Stations is the golang structure for table stations.
type Stations struct {
	Id        int         `json:"id"        orm:"id"         description:""` //
	Name      string      `json:"name"      orm:"name"       description:""` //
	City      string      `json:"city"      orm:"city"       description:""` //
	Code      string      `json:"code"      orm:"code"       description:""` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""` //
	Latitude  float64     `json:"latitude"  orm:"latitude"   description:""` //
	Longitude float64     `json:"longitude" orm:"longitude"  description:""` //
}
