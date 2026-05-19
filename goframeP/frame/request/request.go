package request

import (
	"context"
)

type Options struct {
	Method  string
	Timeout int // 请求超时时间，单位：毫秒
	Headers map[string]string
}

type XClient interface {
	GetBizCode() string
	Do(ctx context.Context, url string, data any, options ...Options) (string, error)
	SetDoBizCode(doBizCode string) XClient
}

var (
	registry = make(map[string]XClient)
)

func register(c XClient) {
	bizCode := c.GetBizCode()
	registry[bizCode] = c
}

func Get(bizCode string) XClient {
	if a, ok := registry[bizCode]; ok {
		return a
	}
	return nil
}
