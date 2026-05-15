package train

import (
	"context"

	"goframeP/frame/api/train/v1"
	"goframeP/frame/internal/model/entity"
	"goframeP/frame/internal/service"
)

func (c *ControllerV1) TrainOrderUpdate(ctx context.Context, req *v1.TrainOrderUpdateReq) (res *v1.TrainOrderUpdateRes, err error) {
	data := entity.TrainOrder{
		Id:       req.Id,
		UserId:   *req.UserId,
		OrderId:  *req.OrderId,
		UserName: *req.UserName,
	}
	err = service.NewTrainOrderService(ctx).UpdateTrainOrder(data)
	if err != nil {
		return nil, err
	}
	return &v1.TrainOrderUpdateRes{}, nil
}
