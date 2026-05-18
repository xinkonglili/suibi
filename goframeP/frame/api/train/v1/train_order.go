package v1

import (
	"goframeP/frame/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type TrainOrderCreateReq struct {
	g.Meta   `path:"/train-order" method:"post" tags:"TrainOrder" summary:"Create train order"`
	UserId   int    `json:"userId" v:"required" dc:"用户id"`                 // 用户id
	OrderId  int    `json:"orderId" v:"required" dc:"订单id"`                // 订单id
	UserName string `json:"userName" v:"required|length:1,50" dc:"用户姓名"` // 用户姓名
}

type TrainOrderCreateRes struct {
	Id int64 `json:"id" dc:"订单ID"`
}

type TrainOrderDeleteReq struct {
	g.Meta `path:"/train-order/{id}" method:"delete" tags:"TrainOrder" summary:"Delete train order"`
	Id     int64 `v:"required" dc:"订单ID"`
}

type TrainOrderDeleteRes struct{}

type TrainOrderUpdateReq struct {
	g.Meta   `path:"/train-order/{id}" method:"put" tags:"TrainOrder" summary:"Update train order"`
	Id       int64   `v:"required" dc:"订单ID"`
	UserId   *int    `json:"userId" dc:"用户id"`
	OrderId  *int    `json:"orderId" dc:"订单id"`
	UserName *string `json:"userName" v:"length:1,50" dc:"用户姓名"`
}

type TrainOrderUpdateRes struct{}

type TrainOrderGetOneReq struct {
	g.Meta `path:"/train-order/{id}" method:"get" tags:"TrainOrder" summary:"Get one train order"`
	Id     int64 `v:"required" dc:"订单ID"`
}

type TrainOrderGetOneRes struct {
	*entity.TrainOrder `dc:"train order"`
}

type TrainOrderGetListReq struct {
	g.Meta   `path:"/train-order" method:"get" tags:"TrainOrder" summary:"Get train order list"`
	UserId   *int   `json:"userId" dc:"用户id"`
	UserName string `json:"userName" dc:"用户姓名（模糊查询）"`
	Page     int    `json:"page" d:"1" dc:"页码"`
	PageSize int    `json:"pageSize" d:"10" dc:"每页数量"`
}

type TrainOrderGetListRes struct {
	List  []*entity.TrainOrder `json:"list" dc:"订单列表"`
	Total int                  `json:"total" dc:"总数"`
	Page  int                  `json:"page" dc:"当前页码"`
}

type TrainOrderPayNotifyReq struct {
	g.Meta  `path:"/train-order/pay-notify" method:"post" tags:"TrainOrder" summary:"支付回调"`
	OrderId string `v:"required" dc:"订单ID"`
}

type TrainOrderPayNotifyRes struct {
	Msg string `json:"msg"`
}
