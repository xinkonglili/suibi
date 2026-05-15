package train

import (
	"context"

	"goframeP/frame/api/train/v1"
	"goframeP/frame/internal/service"
)

func (c *ControllerV1) TrainOrderGetOne(ctx context.Context, req *v1.TrainOrderGetOneReq) (res *v1.TrainOrderGetOneRes, err error) {
	order, err := service.NewTrainOrderService(ctx).GetTrainOrderById(req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.TrainOrderGetOneRes{
		TrainOrder: order,
	}, nil
}
