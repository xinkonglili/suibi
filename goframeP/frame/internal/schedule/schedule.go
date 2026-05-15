package schedule

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

const CacheOrderDetailKey = "order:order_poll_queue"

var ScheduleAfterDur = 30 * time.Second

// 增加调度任务
func OrderScheduleAdd(ctx context.Context, orderNo string) error {
	// 订单查询定时任务开关状态检测
	val, err2 := g.Cfg().Get(ctx, "schedule_config.start_sync_order_worker")
	if err2 != nil {
		return err2
	}
	if val.Int() != 1 {
		return nil
	}

	score := time.Now().Add(ScheduleAfterDur).Unix()
	member := fmt.Sprintf("%s:0:%d", orderNo, score)
	_, err := g.Redis().ZAdd(ctx, CacheOrderDetailKey, &gredis.ZAddOption{}, gredis.ZAddMember{
		Member: member,
		Score:  float64(score),
	})
	if err != nil {
		return err
	}

	return nil
}
