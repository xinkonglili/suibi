package status

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"goframeP/frame/api/status/v1"
)

func (c *ControllerV1) TagStatusUpdate(ctx context.Context, req *v1.TagStatusUpdateReq) (res *v1.TagStatusUpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
