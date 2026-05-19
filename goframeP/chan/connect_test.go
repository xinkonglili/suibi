package _chan

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/guid"
)

/*g.Client()
↓
GoFrame gclient
↓
Go 标准库 http.Client
↓
http.Transport  ← 真正管连接池的地方
↓
TCP 连接池 [ conn1, conn2, conn3... ]*/

var a = func(r *ghttp.Request) {
	r.Response.Write("user world")
}

var b = func(r *ghttp.Request) {
	r.Response.WriteJson(
		map[string]string{
			"code": "0",
			"msg":  "success",
		},
	)
}

func NewV1() interface{} {
	return b
}

func TestServeFrame(t *testing.T) {
	var cmdMain = gcmd.Command{
		Name:          "",
		Usage:         "",
		Brief:         "",
		Description:   "",
		Arguments:     nil,
		FuncWithValue: nil,
		HelpFunc:      nil,
		Examples:      "",
		Additional:    "",
		Strict:        false,
		CaseSensitive: false,
		Config:        "",
	}

	cmdMain.Func = func(ctx context.Context, parser *gcmd.Parser) (err error) {
		s := g.Server(guid.S())
		RegisterHealthRoutes(s)
		RegisterNoPoolRoutes(s)
		RegisterHasPoolRoutes(s)
		s.Run()
		return nil
	}

	cmdMain.Run(context.Background())
}

func TestConnectPool(t *testing.T) {

	// 每次新建，连接池没用
	for i := 0; i < 100; i++ {
		g.Client().Post(context.Background(), "http://localhost:8080/noPool/", g.Map{
			"id":   10000,
			"name": "john no pool",
		}) // 100次请求，可能建100个TCP连接
	}

	// 复用实例，连接池生效
	client := g.Client()
	for i := 0; i < 100; i++ {
		client.Post(context.Background(), "http://localhost:8080/hasPool/", g.Map{
			"id":   10001,
			"name": "john has pool",
		}) // 100次请求，可能只用3~5个TCP连接
	}
}

// RegisterHealthRoutes 注册健康检查路由
func RegisterHealthRoutes(s *ghttp.Server) {
	s.Group("/health", func(group *ghttp.RouterGroup) {
		group.GET("/", func(r *ghttp.Request) {
			r.Response.Write("health")
		})
	})
}

// RegisterNoPoolRoutes 注册无连接池测试路由
func RegisterNoPoolRoutes(s *ghttp.Server) {
	s.Group("/noPool", func(group *ghttp.RouterGroup) {
		group.GET("/", func(r *ghttp.Request) {
			r.Response.Write("no pool")
		})
	})
}

// RegisterHasPoolRoutes 注册有连接池测试路由
func RegisterHasPoolRoutes(s *ghttp.Server) {
	s.Group("/hasPool", func(group *ghttp.RouterGroup) {
		group.GET("/", func(r *ghttp.Request) {
			r.Response.Write("has pool")
		})
	})
}

// RegisterSwitchV1Routes 注册 Switch V1 版本路由
func RegisterSwitchV1Routes(s *ghttp.Server) {
	s.Group("/switch/v1", func(group *ghttp.RouterGroup) {
		group.Middleware(a)
		group.Bind(
			NewV1(),
		)
	})
}
