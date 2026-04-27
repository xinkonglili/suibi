// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RaiseHandTime is the golang structure for frame raise_hand_time.
type RaiseHandTime struct {
	Id            uint64      `json:"id"            orm:"id"              description:""`       //
	RaiseHandId   uint64      `json:"raiseHandId"   orm:"raise_hand_id"   description:"举手策略ID"` // 举手策略ID
	StartTime     *gtime.Time `json:"startTime"     orm:"start_time"      description:"开始时间"`   // 开始时间
	EndTime       *gtime.Time `json:"endTime"       orm:"end_time"        description:"结束时间"`   // 结束时间
	CreatedBy     uint64      `json:"createdBy"     orm:"created_by"      description:"创建人ID"`  // 创建人ID
	CreatedByName string      `json:"createdByName" orm:"created_by_name" description:"创建人名称"`  // 创建人名称
	UpdatedBy     uint64      `json:"updatedBy"     orm:"updated_by"      description:"更新人ID"`  // 更新人ID
	UpdatedByName string      `json:"updatedByName" orm:"updated_by_name" description:"更新人名称"`  // 更新人名称
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:""`       //
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"      description:""`       //
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"      description:""`       //
}
