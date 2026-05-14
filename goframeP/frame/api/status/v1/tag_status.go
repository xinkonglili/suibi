package v1

import (
	"goframeP/frame/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 创建标签状态
type TagStatusCreateReq struct {
	g.Meta  `path:"/tag-status" method:"post" tags:"TagStatus" summary:"Create tag status"`
	Kaiguan int    `json:"kaiguan" v:"required" dc:"开关"`
	Status  string `json:"status" dc:"状态"`
	Tag     string `json:"tag" v:"required" dc:"标签"`
}
type TagStatusCreateRes struct {
	Id int64 `json:"id" dc:"id"`
}

// 删除标签状态
type TagStatusDeleteReq struct {
	g.Meta `path:"/tag-status/{id}" method:"delete" tags:"TagStatus" summary:"Delete tag status"`
	Id     int64 `v:"required" dc:"id"`
}
type TagStatusDeleteRes struct{}

// 更新标签状态
type TagStatusUpdateReq struct {
	g.Meta  `path:"/tag-status/{id}" method:"put" tags:"TagStatus" summary:"Update tag status"`
	Id      int64   `v:"required" dc:"id"`
	Kaiguan *int    `json:"kaiguan" dc:"开关"`
	Status  *string `json:"status" dc:"状态"`
	Tag     *string `json:"tag" dc:"标签"`
}
type TagStatusUpdateRes struct{}

// 获取单个标签状态
type TagStatusGetOneReq struct {
	g.Meta `path:"/tag-status/{id}" method:"get" tags:"TagStatus" summary:"Get one tag status"`
	Id     int64 `v:"required" dc:"id"`
}
type TagStatusGetOneRes struct {
	*entity.TagStatus `dc:"tag status"`
}

// 获取标签状态列表
type TagStatusGetListReq struct {
	g.Meta `path:"/tag-status" method:"get" tags:"TagStatus" summary:"Get tag status list"`
}
type TagStatusGetListRes struct {
	List []*entity.TagStatus `json:"list" dc:"tag status list"`
}
