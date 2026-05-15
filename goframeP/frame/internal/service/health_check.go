package service

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
)

type HealthCheckService struct {
	ctx context.Context
}

// DownstreamServiceStatus 下游服务状态
type DownstreamServiceStatus struct {
	ServiceName string    `json:"service_name"`
	IsHealthy   bool      `json:"is_healthy"`
	LastCheckAt time.Time `json:"last_check_at"`
	ErrorMsg    string    `json:"error_msg,omitempty"`
}

func NewHealthCheckService(ctx context.Context) *HealthCheckService {
	return &HealthCheckService{ctx: ctx}
}

// StartHealthCheck 启动定时健康检查
func (s *HealthCheckService) StartHealthCheck() error {
	// 每30秒检查一次下游服务状态
	_, err := gcron.AddSingleton(s.ctx, "@every 30s", func(ctx context.Context) {
		s.checkDownstreamServices()
	})
	if err != nil {
		return err
	}
	g.Log().Info(s.ctx, "健康检查定时任务已启动")
	return nil
}

// checkDownstreamServices 检查下游服务状态
func (s *HealthCheckService) checkDownstreamServices() {
	services := []string{
		"http://downstream-service-1/health",
		"http://downstream-service-2/health",
	}

	for _, serviceURL := range services {
		status := s.checkSingleService(serviceURL)
		s.saveStatus(status)

		if !status.IsHealthy {
			g.Log().Warning(s.ctx, fmt.Sprintf("下游服务异常: %s, 错误: %s", status.ServiceName, status.ErrorMsg))
			// 可以发送告警通知
			s.sendAlert(status)
		}
	}
}

// checkSingleService 检查单个服务
func (s *HealthCheckService) checkSingleService(url string) DownstreamServiceStatus {
	status := DownstreamServiceStatus{
		ServiceName: url,
		LastCheckAt: time.Now(),
	}

	// 发起HTTP请求检查健康状态
	client := g.Client()
	client.SetTimeout(5 * time.Second)
	
	resp, err := client.Get(s.ctx, url)
	if err != nil {
		status.IsHealthy = false
		status.ErrorMsg = fmt.Sprintf("连接失败: %v", err)
		return status
	}
	defer resp.Close()

	if resp.StatusCode == 200 {
		status.IsHealthy = true
	} else {
		status.IsHealthy = false
		status.ErrorMsg = fmt.Sprintf("HTTP状态码: %d", resp.StatusCode)
	}

	return status
}

// saveStatus 保存状态到数据库或缓存
func (s *HealthCheckService) saveStatus(status DownstreamServiceStatus) {
	// 这里可以保存到数据库或Redis
	// 示例：保存到Redis
	key := fmt.Sprintf("health_check:%s", status.ServiceName)
	g.Redis().Set(s.ctx, key, status, 5*time.Minute)
}

// sendAlert 发送告警
func (s *HealthCheckService) sendAlert(status DownstreamServiceStatus) {
	// 可以实现发送邮件、短信、钉钉、企业微信等告警
	g.Log().Error(s.ctx, fmt.Sprintf("【告警】下游服务异常: %+v", status))
	
	// 示例：发送到钉钉
	// s.sendDingTalkAlert(status)
}

// GetServiceStatus 获取服务状态
func (s *HealthCheckService) GetServiceStatus(serviceName string) (*DownstreamServiceStatus, error) {
	key := fmt.Sprintf("health_check:%s", serviceName)
	var status DownstreamServiceStatus
	
	err := g.Redis().GetVar(s.ctx, key).Scan(&status)
	if err != nil {
		return nil, err
	}
	
	return &status, nil
}
