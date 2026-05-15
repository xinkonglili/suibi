package train

import (
	"context"

	"goframeP/frame/api/train/v1"
	"goframeP/frame/internal/dao"
	"goframeP/frame/internal/model/entity"
)

func (c *ControllerV1) TrainOrderGetList(ctx context.Context, req *v1.TrainOrderGetListReq) (res *v1.TrainOrderGetListRes, err error) {
	res = &v1.TrainOrderGetListRes{
		Page: req.Page,
	}
	
	// 构建查询条件
	model := dao.TrainOrder.Ctx(ctx)
	
	// 按 userId 筛选
	if req.UserId != nil {
		model = model.Where("user_id", *req.UserId)
	}
	
	// 按 userName 模糊查询
	if req.UserName != "" {
		model = model.WhereLike("user_name", "%"+req.UserName+"%")
	}
	
	// 获取总数
	total, err := model.Count()
	if err != nil {
		return nil, err
	}
	res.Total = total
	
	// 分页查询
	var list []*entity.TrainOrder
	err = model.Page(req.Page, req.PageSize).Scan(&list)
	if err != nil {
		return nil, err
	}
	res.List = list
	
	return res, nil
}
