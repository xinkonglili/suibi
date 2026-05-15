package v1

import (
	"goframeP/frame/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type UserCreateReq struct {
	g.Meta `path:"/user" method:"post" tags:"User" summary:"Create user"`
	Name   string `v:"required|length:3,10" dc:"user name"`   //这里加入了参数校验的逻辑：v:"required|length:3,10
	Age    uint   `v:"required|between:18,200" dc:"user age"` //这里加入了参数校验的逻辑：v:"required|between:18,200"
}
type UserCreateRes struct {
	Id int64 `json:"id" dc:"user id"`
}

type UserDeleteReq struct {
	g.Meta `path:"/user/{id}" method:"delete" tags:"User" summary:"Delete user"`
	Id     int64 `v:"required" dc:"user id"`
}
type UserDeleteRes struct{}

// Status marks user status.
type Status int

const (
	StatusOK       Status = 0 // User is OK.
	StatusDisabled Status = 1 // User is disabled.
)

type UserUpdateReq struct {
	g.Meta `path:"/user/{id}" method:"put" tags:"User" summary:"Update user"`
	Id     int64   `v:"required" dc:"user id"`
	Name   *string `v:"length:3,10" dc:"user name"`
	Age    *uint   `v:"between:18,200" dc:"user age"`
	Status *Status `v:"in:0,1" dc:"user status"`
}
type UserUpdateRes struct{}

type UserGetOneReq struct {
	g.Meta `path:"/user/{id}" method:"get" tags:"User" summary:"Get one user"`
	Id     int64 `v:"required" dc:"user id"`
}
type UserGetOneRes struct {
	*entity.User `dc:"user"`
}

type UserGetListReq struct {
	g.Meta `path:"/user" method:"get" tags:"User" summary:"Get users"`
	Age    *uint   `v:"between:18,200" dc:"user age"`
	Status *Status `v:"in:0,1" dc:"user status"`
}
type UserGetListRes struct {
	List []*entity.User `json:"list" dc:"user list"`
}
