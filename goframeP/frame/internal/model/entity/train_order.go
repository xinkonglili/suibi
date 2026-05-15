// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TrainOrder is the golang structure for table train_order.
type TrainOrder struct {
	Id       int64  `json:"id"       orm:"id"        description:""`     //
	UserId   int    `json:"userId"   orm:"user_id"   description:"用户id"` // 用户id
	OrderId  int    `json:"orderId"  orm:"order_id"  description:""`     //
	UserName string `json:"userName" orm:"user_name" description:""`     //
}
