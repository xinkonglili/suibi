package hello

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"goframeP/frame/api/hello/v1"
)

func (c *ControllerV1) UserGetOne(ctx context.Context, req *v1.UserGetOneReq) (res *v1.UserGetOneRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
