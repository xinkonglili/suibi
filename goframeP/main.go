package main

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var ctx = context.Background()

func main() {
	//chan：
	//1、注意：chan 默认是阻塞的，如果chan没有被消费，那么写入chan的数据就会阻塞，直到chan被消费
	//2、注意有缓冲区，如果缓冲区已满，那么写入chan的数据就会阻塞，直到缓冲区有空间
	//3、向已关闭的chan写入数据会panic，读数据不会 panic
	ch := make(chan int, 10)

	fmt.Println("test -11")
	ch <- 1
	fmt.Println("test -1")
	ch <- 2
	fmt.Println("test 0")
	ch <- 3
	ch <- 4
	close(ch)
	ch <- 5
	for c := range ch {
		fmt.Println(c)
	}

	// 1. 初始化连接
	//rdb := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "", // 密码
	//	DB:       0,  // 数据库编号
	//})
	//
	//// 2. 判断 Key 是否存在
	//key := "my_key"
	//exists, err := rdb.Exists(ctx, key).Result()
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 3. 返回结果处理
	//if exists > 0 {
	//	fmt.Println("Key 存在")
	//} else {
	//	fmt.Println("Key 不存在")
	//}
}

func GetSliceMap(ctx context.Context, key string) (confData []map[string]any, err error) {
	exists, err := Exists(ctx, key)
	if err != nil {
		return
	}
	if exists > 0 {

	}
	return GetSliceMap(ctx, key)
}

func Exists(ctx context.Context, keys ...string) (int64, error) {
	//这里如果key不存在返回什么？那err是什么情况才会返回，什么错误？
	v, err := g.Redis().Exists(ctx, keys...)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func GetKey(ctx context.Context, key string) (string, error) {
	v, err := g.Redis().Get(ctx, key)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

func SwitchSignVerify(r *ghttp.Request) {
	// 公共参数校验
	//bizCode := r.GetQuery("bizCode").String()
	//timestamp := r.GetQuery("timestamp").String()
	//reqSign := r.GetQuery("sign").String()
	//if bizCode == "" || timestamp == "" || reqSign == "" {
	//	r.Response.WriteJson(response.SwitchResponse{
	//		Code:    gerror.Code(errcode.SwitchPublicParamsError).Code(),
	//		Msg:     errcode.SwitchPublicParamsError.Error(),
	//		Data:    nil,
	//		TraceId: gctx.CtxId(r.Context()),
	//	})
	//	return
	//}
	//if bizCode != consts.SwitchBizCode {
	//	//r.Response.WriteJson(response.SwitchResponse{
	//	//	Code:    gerror.Code(errcode.SwitchAuthError).Code(),
	//	//	Msg:     errcode.SwitchAuthError.Error(),
	//	//	Data:    nil,
	//	//	TraceId: gctx.CtxId(r.Context()),
	//	//})
	//	return
	//}

	// 一段测试代码
	//if r.GetQuery("debug").String() != "1" {
	//	r.Response.WriteJson(response.SwitchResponse{
	//		Code:    gerror.Code(errcode.SwitchAuthError).Code(),
	//		Msg:     "当前为调试模式",
	//		Data:    nil,
	//		TraceId: gctx.CtxId(r.Context()),
	//	})
	//	return
	//}
	//
	//if lock, _ := g.Redis().Get(r.Context(), consts.RedisSwitchSignVerify); lock.String() == "1" {
	//	// 关闭签名
	//	r.Middleware.Next()
	//	return
	//}
	//
	//// 签名校验
	//calculateSign, waitSignStr := xsign.CalculateSwitchSign(r.Context(), r.GetBodyString(), bizCode, timestamp, consts.SwitchSecret)
	//if reqSign != calculateSign {
	//	g.Log().Errorf(r.Context(), "签名验证失败, 签名字串: %v, 传入签名: %s, 计算签名: %s", waitSignStr, reqSign, calculateSign)
	//	r.Response.WriteJson(response.SwitchResponse{
	//		Code:    gerror.Code(errcode.SwitchSignError).Code(),
	//		Msg:     errcode.SwitchSignError.Error(),
	//		Data:    nil,
	//		TraceId: gctx.CtxId(r.Context()),
	//	})
	//	return
	//}

	r.Middleware.Next()
}
