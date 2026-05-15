// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Stations is the golang structure of table stations for DAO operations like Where/Data.
type Stations struct {
	g.Meta    `orm:"table:stations, do:true"`
	Id        any         //
	Name      any         //
	City      any         //
	Code      any         //
	CreatedAt *gtime.Time //
	Latitude  any         //
	Longitude any         //
}
