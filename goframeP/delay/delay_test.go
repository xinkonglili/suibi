package delay

import (
	"context"
	"errors"
	"fmt"
	"goframeP/delay/common"
	"goframeP/delay/schedule"
	"log"
	"strings"
	"sync"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	maxRetry       = 50            // 最大重试50次
	chunkSize      = 20            // 每批处理20个订单
	lockExpiration = 6             // 分布式锁6秒过期
	lockKeyPrefix  = "order:lock:" // 锁前缀
)

// OrderCron 订单定时任务处理器
type OrderCron struct {
	ctx        context.Context
	workerPool chan struct{}
}

func (c *OrderCron) StartOrderWorker() {
	// 从Redis有序集合中取出最多800个待处理的订单
	res, err := g.Redis().ZRange(context.Background(), "key", 0, 800, gredis.ZRangeOption{
		// WithScores: true,
	})
	if err != nil {
		return
	}

	//members, err := redis.ZRange(c.ctx, schedule.CacheOrderDetailKey, 0, 800)
	//if err != nil {
	//	log.Println("Fetch error:", err)
	//	return
	//}

	// 分批次处理，每批20个
	memberChunk, err := Chunk(res.Strings(), chunkSize)
	if err != nil {
		log.Println("Chunk error:", err)
		return
	}

	var wg sync.WaitGroup
	for _, memberArr := range memberChunk {
		wg.Add(1)
		c.workerPool <- struct{}{} // 工作池限流

		go func(arr []string) {
			defer func() {
				<-c.workerPool
				wg.Done()
			}()
			c.ProcessOrderGroup(arr)
		}(memberArr)
	}

	wg.Wait()
	close(c.workerPool)
}

// 处理订单组
func (c *OrderCron) ProcessOrderGroup(memberArr []string) {
	//timeNow := time.Now().Unix()
	//for _, memberKey := range memberArr {
	//	_, timeAt, err := SplitMemberScore(memberKey)
	//	if err != nil {
	//		common.ErrLogAndMsg(c.ctx, "解析订单标识失败:"+err.Error())
	//		continue
	//	}
	//	// 检查是否到达执行时间
	//	if timeNow < int64(timeAt) {
	//		continue
	//	}
	//
	//	c.processOrderDetail(memberKey)
	//}
}

// 处理订单详情
func (c *OrderCron) processOrderDetail(memberKey string) {
	orderNo, retry, err := SplitOrderIdentifier(memberKey)
	if err != nil {
		common.ErrLogAndMsg(c.ctx, "解析订单标识失败:"+err.Error())
		return
	}

	if err := c.checkOrderExecutionStatus(orderNo); err != nil {
		common.ErrLogAndMsg(c.ctx, "检查订单执行状态失败:"+err.Error())
		return
	}

	//lockKey := lockKeyPrefix + orderNo
	//lockToken := guid.S()
	//if err := c.acquireLock(lockKey, lockToken); err != nil {
	//	common.ErrLogAndMsg(c.ctx, "获取分布式锁失败:"+err.Error())
	//	return
	//}
	//defer c.releaseLock(lockKey, lockToken)
	//
	//if err := schedule.OrderScheduleRemove(c.ctx, memberKey); err != nil {
	//	common.ErrLogAndMsg(c.ctx, "移除订单调度失败:"+err.Error())
	//	return
	//}

	c.handleOrderQuery(orderNo, retry)
}

func (c *OrderCron) checkOrderExecutionStatus(orderNo string) error {
	cacheKey := schedule.CacheOrderExecutePrefix + orderNo
	exeStatus, _ := g.Redis().Get(c.ctx, cacheKey)
	if exeStatus.String() != "" {
		return fmt.Errorf("订单已在处理中")
	}
	return nil
}

func (c *OrderCron) releaseLock(key string, value string) {
	//把锁释放，使用lua脚本实现
	script := `
		if redis.call("GET", KEYS[1]) == ARGV[1] then
			return redis.call("DEL", KEYS[1])
		else
			return 0
		end`

	//script string, numKeys int64, keys []string, args []an

	keys := []string{key}
	args := []interface{}{value} // 过期时间作为参数传递

	_, err := g.Redis().Eval(c.ctx, script, int64(len(keys)), keys, args)
	if err != nil {

	}
}

func (c *OrderCron) handleOrderQuery(no string, retry int) {
	return
}

func (c *OrderCron) acquireLock(key string, token string) interface{} {
	return nil
}

func SplitOrderIdentifier(member string) (orderNo string, retry int, err error) {
	parts := strings.Split(member, ":")
	if len(parts) < 2 {
		return "", 0, errors.New("SplitOrderIdentifier length error")
	}
	orderNo = parts[0]
	retry = gconv.Int(parts[1])
	return
}

// 一维数组转二维数组
func Chunk[T any, Slice ~[]T](collection Slice, size int) ([]Slice, error) {
	if size <= 0 {
		return nil, errors.New("second parameter must be greater than 0")
	}

	chunksNum := len(collection) / size
	if len(collection)%size != 0 {
		chunksNum += 1
	}

	result := make([]Slice, 0, chunksNum)

	for i := 0; i < chunksNum; i++ {
		last := (i + 1) * size
		if last > len(collection) {
			last = len(collection)
		}

		// Copy chunk in a new slice, to prevent memory leak and free memory from initial collection.
		newSlice := make(Slice, last-i*size)
		copy(newSlice, collection[i*size:last])
		result = append(result, newSlice)
	}

	return result, nil
}
func ZRange(ctx context.Context, key string, min, max int64) ([]string, error) {
	res, err := g.Redis().ZRange(ctx, key, min, max, gredis.ZRangeOption{
		// WithScores: true,
	})
	if err != nil {
		return nil, err
	}
	return res.Strings(), nil
}
