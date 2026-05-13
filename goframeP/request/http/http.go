package http

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func Get(ctx context.Context, url string) {
	response, err := g.Client().Get(ctx, url, nil)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	defer response.Close()

	/*
		GoFrame Response 常用方法:
		- GetCookie(ctx, key string) string       // 获取指定cookie
		- GetCookieMap(ctx) map[string]string      // 获取所有cookie
		- ReadAll() []byte                         // 读取响应字节
		- ReadAllString() string                   // 读取响应字符串
		- RawRequest() string                      // 原始请求
		- RawResponse() string                     // 原始响应
	*/

	// 示例1: 获取所有Cookie
	allCookies := response.GetCookieMap()
	g.Log().Info(ctx, "所有Cookie:", allCookies)

	// 遍历所有Cookie的key和value
	for key, value := range allCookies {
		g.Log().Infof(ctx, "Cookie名称: %s, 值: %s", key, value)
	}

	// 示例2: 获取指定的Cookie (需要先知道cookie的key名称)
	// sessionId := response.GetCookie(ctx, "session_id")
	// token := response.GetCookie(ctx, "token")

	g.Log().Info(ctx, "响应内容:", response.ReadAllString())
}

//
