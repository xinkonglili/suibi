package service

import (
	"context"
	v1 "goframeP/frame/api/status/v1"
	"goframeP/frame/internal/dao"
	"goframeP/frame/internal/model/entity"
)

type service struct {
	ctx context.Context
	// 可注入缓存等依赖
}

func NewService(ctx context.Context) *service {
	return &service{ctx: ctx}
}

func (s *service) CreateTagStatus(req v1.TagStatusCreateReq) (int64, error) {
	temp := entity.TagStatus{
		Status:  req.Status,
		Tag:     req.Tag,
		Kaiguan: req.Kaiguan,
	}
	id, err := dao.TagStatus.Create(s.ctx, temp)
	if err != nil {
		return 0, err
	}
	return id, nil
}
