package train

import (
	"context"
	"goframeP/frame/internal/model/entity"
	"goframeP/frame/internal/service"

	"goframeP/frame/api/train/v1"
)

func (c *ControllerV1) TrainOrderCreate(ctx context.Context, req *v1.TrainOrderCreateReq) (res *v1.TrainOrderCreateRes, err error) {
	data := entity.TrainOrder{
		UserId:   req.UserId,
		OrderId:  req.OrderId,
		UserName: req.UserName,
	}
	id, err := service.NewTrainOrderService(ctx).CreateTrainOrder(data)
	if err != nil {
		return nil, err
	}
	return &v1.TrainOrderCreateRes{
		Id: id,
	}, nil
}
