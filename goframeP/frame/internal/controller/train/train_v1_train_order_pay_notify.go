package train

import (
	"context"
	"goframeP/frame/internal/model/entity"
	"goframeP/frame/internal/service"

	"goframeP/frame/api/train/v1"
)

func (c *ControllerV1) TrainOrderPayNotify(ctx context.Context, req *v1.TrainOrderPayNotifyReq) (res *v1.TrainOrderPayNotifyRes, err error) {
	temp := &entity.TrainOrder{
		OrderId: req.OrderId,
	}
	result, err := service.PayHandle(ctx, temp)
	if err != nil {
		return nil, err
	}
	return result, nil
}
