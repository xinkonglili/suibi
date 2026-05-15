package train

import (
	"context"

	"goframeP/frame/api/train/v1"
	"goframeP/frame/internal/service"
)

func (c *ControllerV1) TrainOrderDelete(ctx context.Context, req *v1.TrainOrderDeleteReq) (res *v1.TrainOrderDeleteRes, err error) {
	err = service.NewTrainOrderService(ctx).DeleteTrainOrder(req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.TrainOrderDeleteRes{}, nil
}
