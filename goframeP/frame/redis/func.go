package redis

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

var RedisPrefix = "Redis"
var RedisOnce sync.Once

func RedisInit() {
	g.Log(RedisPrefix).SetPrefix(RedisPrefix)
}

func getKey(key string) string {
	RedisOnce.Do(RedisInit)
	return key
}

/*
典型使用场景：
限流/频率控制，比如：
某个用户 1 分钟内最多请求 100 次
每次请求调用此函数，key的值累加 1，ttl 设为 60 秒
返回值超过100就拒绝请求
每次调用都刷新 ttl，保证窗口是滑动的*/
// 原子性地对 Redis 中的某个 key的值进行累加，并同时刷新它的过期时间。
func IncrByAndUpdateTtl(ctx context.Context, key string, increment, ttl int64) (int64, error) {
	script := `
			local num = redis.call('INCRBY', KEYS[1], tonumber(ARGV[1]))
			redis.call('EXPIRE', KEYS[1], tonumber(ARGV[2]))
			return num
			`
	if num, err := Eval(ctx, script, []string{key}, []any{increment, ttl}); err != nil {
		return 0, err
	} else {
		return num.(int64), nil //做断言
	}
}

func Eval(ctx context.Context, script string, keys []string, args []interface{}) (res interface{}, err error) {
	gVal, err := DoRedisEval(ctx, script, keys, args) //实际执行 Redis EVAL命令，传入脚本keys、args
	if err != nil {
		return
	}
	return gVal.Int64(), err
}

func DoRedisEval(ctx context.Context, script string, keys []string, args []interface{}) (res *gvar.Var, err error) {
	for i := 0; i < len(keys); i++ {
		keys[i] = getKey(keys[i])
	}

	//timePoint := time.Now()
	//defer func(ctx context.Context, timePoint time.Time, keys []string, args []interface{}) {
	//	go func(ctx context.Context, timePoint time.Time, keys []string, args []interface{}) {
	//		keysJson, _ := json.Marshal(keys)
	//		argsJson, _ := json.Marshal(args)
	//		g.Log(RedisPrefix).Debug(ctx, fmt.Sprintf("[ %d ms]", time.Since(timePoint).Milliseconds())+" Eval "+string(keysJson)+string(argsJson))
	//	}(ctx, timePoint, keys, args)
	//}(ctx, timePoint, keys, args)

	return g.Redis().Eval(ctx, script, int64(len(keys)), keys, args)
}

func ZRange(ctx context.Context, key string, min, max int64) ([]string, error) {
	res, err := g.Redis().ZRange(ctx, getKey(key), min, max, gredis.ZRangeOption{
		// WithScores: true,
	})
	if err != nil {
		return nil, err
	}
	return res.Strings(), nil
}

func ZRem(ctx context.Context, key string, member any) (int64, error) {
	return g.Redis().ZRem(ctx, getKey(key), member)
}

func Rpush(ctx context.Context, key string, member any) (int64, error) {
	return g.Redis().RPush(ctx, getKey(key), member)
}

func LPop(ctx context.Context, key string) (string, error) {
	pop, err := g.Redis().LPop(ctx, getKey(key))
	if err != nil {
		return "", err
	}
	return pop.String(), nil
}
