// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package train

import (
	"context"

	"goframeP/frame/api/train/v1"
)

type ITrainV1 interface {
	TrainOrderCreate(ctx context.Context, req *v1.TrainOrderCreateReq) (res *v1.TrainOrderCreateRes, err error)
	TrainOrderDelete(ctx context.Context, req *v1.TrainOrderDeleteReq) (res *v1.TrainOrderDeleteRes, err error)
	TrainOrderUpdate(ctx context.Context, req *v1.TrainOrderUpdateReq) (res *v1.TrainOrderUpdateRes, err error)
	TrainOrderGetOne(ctx context.Context, req *v1.TrainOrderGetOneReq) (res *v1.TrainOrderGetOneRes, err error)
	TrainOrderGetList(ctx context.Context, req *v1.TrainOrderGetListReq) (res *v1.TrainOrderGetListRes, err error)
	TrainOrderPayNotify(ctx context.Context, req *v1.TrainOrderPayNotifyReq) (res *v1.TrainOrderPayNotifyRes, err error)
}
