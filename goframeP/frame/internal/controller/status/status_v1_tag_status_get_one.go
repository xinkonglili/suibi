package status

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"goframeP/frame/api/status/v1"
)

func (c *ControllerV1) TagStatusGetOne(ctx context.Context, req *v1.TagStatusGetOneReq) (res *v1.TagStatusGetOneRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
