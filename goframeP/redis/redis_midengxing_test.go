package redis

import (
	"context"
	"fmt"
	"goframeP/delay/schedule"
	"testing"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func TestRedisSetNX(t *testing.T) {
	var mainOrderId = "XH202605060000000161"
	cacheKey := fmt.Sprintf("%s%s", schedule.CacheOrderExecutePrefix, mainOrderId)
	ctx := context.Background()
	// SetNX: 只有 key 不存在时才写入，返回是否成功
	success, err := g.Redis().SetNX(ctx, cacheKey, mainOrderId)
	if err != nil {
		t.Error(ctx, "SetNX", err)
	}
	if !success {
		// key 已存在，说明有并发请求正在处理，直接返回
		return gerror.NewCode(consts.SwitchCode{OutCode: 1429, OutMessage: "请求处理中，请勿重复提交"})
	}
	defer func() {
		g.Redis().Del(ctx, cacheKey)
	}()

	t.Log(ctx, "订单执行锁", "key", cacheKey, "value", mainOrderId)
}
