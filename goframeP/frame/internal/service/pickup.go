package service

import (
	"context"
	v1 "goframeP/frame/api/train/v1"
	"goframeP/frame/internal/dao"
	"goframeP/frame/internal/model/entity"
)

type TrainOrderService struct {
	ctx context.Context
}

func NewTrainOrderService(ctx context.Context) *TrainOrderService {
	return &TrainOrderService{ctx: ctx}
}

// CreateTrainOrder 创建火车订单
func (s *TrainOrderService) CreateTrainOrder(data entity.TrainOrder) (int64, error) {
	id, err := dao.TrainOrder.Create(s.ctx, data)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetTrainOrderById 根据ID获取订单
func (s *TrainOrderService) GetTrainOrderById(id int64) (*entity.TrainOrder, error) {
	var order entity.TrainOrder
	err := dao.TrainOrder.Ctx(s.ctx).Where("id", id).Scan(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// UpdateTrainOrder 更新订单
func (s *TrainOrderService) UpdateTrainOrder(data entity.TrainOrder) error {
	_, err := dao.TrainOrder.Ctx(s.ctx).Data(data).Where("id", data.Id).Update()
	return err
}

// DeleteTrainOrder 删除订单
func (s *TrainOrderService) DeleteTrainOrder(id int64) error {
	_, err := dao.TrainOrder.Ctx(s.ctx).Where("id", id).Delete()
	return err
}

func (s *service) PickUp() {

}

// PayOrderCheckInit
func (s *TrainOrderService) PayOrderCheckInit() (bool, error) {
	return true, nil

}

// PayOrder
func (s *TrainOrderService) PayOrder(req *entity.TrainOrder) (result *v1.TrainOrderPayNotifyRes, err error) {
	//todo 待实现
	return &v1.TrainOrderPayNotifyRes{}, nil
}
