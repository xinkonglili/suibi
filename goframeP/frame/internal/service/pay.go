package service

import (
	"context"
	v1 "goframeP/frame/api/train/v1"
	"goframeP/frame/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

func PayHandle(ctx context.Context, req *entity.TrainOrder) (result *v1.TrainOrderPayNotifyRes, err error) {
	orderService := NewTrainOrderService(ctx)
	// 参数验证和初始化
	checkResult, err := orderService.PayOrderCheckInit()
	if err != nil || !checkResult {
		return result, err
	}
	// 判断是否已支付过，如果已支付，直接返回就好
	if req.FareStatus == 40 { //10 --未计费,30-计费中，31-待支付，40-已支付
		g.Log().Info(ctx, "订单已支付:", req.OrderId)
		return &v1.TrainOrderPayNotifyRes{
			Msg: "订单已支付",
		}, nil
	}

	// 未支付，获取下redis锁，看看是否有并发的支付请求
	redisKey := "RedisPayNotifyLockKey" + req.OrderId
	if payNotifyLock, err := g.Redis().SetNX(ctx, redisKey, 10); err != nil {
		g.Log().Error(ctx, "支付通知获取锁失败:", err.Error())
		return result, err
	} else if !payNotifyLock {
		// 未获取到锁，说明有并发的重复支付请求
		g.Log().Info(ctx, "重复的支付通知:", req.OrderId)
		return &v1.TrainOrderPayNotifyRes{
			Msg: "订单已支付",
		}, nil
	}
	defer g.Redis().Del(ctx, redisKey)

	// 获取到锁，处理支付请求
	return orderService.PayOrder(req)
}
