package status

import (
	"context"
	"goframeP/frame/internal/service"

	"goframeP/frame/api/status/v1"
)

func (c *ControllerV1) TagStatusCreate(ctx context.Context, req *v1.TagStatusCreateReq) (res *v1.TagStatusCreateRes, err error) {
	id, err := service.NewService(ctx).CreateTagStatus(*req)
	if err != nil {
		return nil, err
	}
	return &v1.TagStatusCreateRes{
		Id: id,
	}, nil
}
