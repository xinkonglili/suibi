// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package status

import (
	"context"

	"goframeP/frame/api/status/v1"
)

type IStatusV1 interface {
	TagStatusCreate(ctx context.Context, req *v1.TagStatusCreateReq) (res *v1.TagStatusCreateRes, err error)
	TagStatusDelete(ctx context.Context, req *v1.TagStatusDeleteReq) (res *v1.TagStatusDeleteRes, err error)
	TagStatusUpdate(ctx context.Context, req *v1.TagStatusUpdateReq) (res *v1.TagStatusUpdateRes, err error)
	TagStatusGetOne(ctx context.Context, req *v1.TagStatusGetOneReq) (res *v1.TagStatusGetOneRes, err error)
	TagStatusGetList(ctx context.Context, req *v1.TagStatusGetListReq) (res *v1.TagStatusGetListRes, err error)
}
