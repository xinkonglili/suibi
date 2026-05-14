// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RaiseHandDate is the golang structure for table raise_hand_date.
type RaiseHandDate struct {
	Id            uint64      `json:"id"            orm:"id"              description:""`       //
	RaiseHandId   uint64      `json:"raiseHandId"   orm:"raise_hand_id"   description:"举手策略ID"` // 举手策略ID
	StartDate     *gtime.Time `json:"startDate"     orm:"start_date"      description:"开始日期"`   // 开始日期
	EndDate       *gtime.Time `json:"endDate"       orm:"end_date"        description:"结束日期"`   // 结束日期
	CreatedBy     uint64      `json:"createdBy"     orm:"created_by"      description:"创建人ID"`  // 创建人ID
	CreatedByName string      `json:"createdByName" orm:"created_by_name" description:"创建人名称"`  // 创建人名称
	UpdatedBy     uint64      `json:"updatedBy"     orm:"updated_by"      description:"更新人ID"`  // 更新人ID
	UpdatedByName string      `json:"updatedByName" orm:"updated_by_name" description:"更新人名称"`  // 更新人名称
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:""`       //
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"      description:""`       //
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"      description:""`       //
}
