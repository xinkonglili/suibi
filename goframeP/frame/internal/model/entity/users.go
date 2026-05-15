// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Id        int64       `json:"id"        orm:"id"         description:""`               //
	Username  string      `json:"username"  orm:"username"   description:""`               //
	Password  string      `json:"password"  orm:"password"   description:""`               //
	Phone     string      `json:"phone"     orm:"phone"      description:""`               //
	Email     string      `json:"email"     orm:"email"      description:""`               //
	RealName  string      `json:"realName"  orm:"real_name"  description:""`               //
	IdCard    string      `json:"idCard"    orm:"id_card"    description:""`               //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`               //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`               //
	Sex       int         `json:"sex"       orm:"sex"        description:"'0:未知 1:男 2:女'"` // '0:未知 1:男 2:女'
}
