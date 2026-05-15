package cmd

import (
	"context"
	"goframeP/frame/internal/controller/status"
	"goframeP/frame/internal/controller/train"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 启动健康检查定时任务
			//healthCheckService := service.NewHealthCheckService(ctx)
			//if err := healthCheckService.StartHealthCheck(); err != nil {
			//	g.Log().Error(ctx, "启动健康检查失败:", err)
			//}

			s := g.Server()

			// 基础健康检查接口（用于负载均衡器/容器编排）
			//s.Group("/health", func(group *ghttp.RouterGroup) {
			//	group.GET("/", func(r *ghttp.Request) {
			//		// 快速检查关键依赖
			//		health := map[string]interface{}{
			//			"status":  "ok",
			//			"version": "1.0.0",
			//			"time":    time.Now().Format("2006-01-02 15:04:05"),
			//		}
			//
			//		//// 检查数据库连接（轻量级）
			//		//if err := g.DB().Ping(); err != nil {
			//		//	health["status"] = "degraded" // 降级状态
			//		//	health["database"] = "disconnected"
			//		//	r.Response.WriteHeader(503) // Service Unavailable
			//		//} else {
			//		//	health["database"] = "connected"
			//		//}
			//
			//		r.Response.WriteJsonExit(health)
			//	})
			//})

			// 详细健康检查接口（用于管理后台/监控系统）
			//s.Group("/api/health", func(group *ghttp.RouterGroup) {
			//	group.GET("/detail", func(r *ghttp.Request) {
			//		// 可以返回更详细的健康信息
			//		r.Response.WriteJsonExit(map[string]interface{}{
			//			"status": "ok",
			//			"checks": map[string]interface{}{
			//				"database": "connected",
			//				"redis":    "connected",
			//			},
			//		})
			//	})
			//})

			s.Group("/ss/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					status.NewV1(),
				)
			})

			s.Group("/order/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					train.NewV1(),
				)
			})

			s.Run()
			return nil
		},
	}
)
