package service

import (
	"context"
	"time"
)

type PushService struct {
	ctx context.Context
}

// PushRecord 推送记录
type PushRecord struct {
	Id          int64     `json:"id"`
	BusinessId  string    `json:"business_id"`
	Payload     string    `json:"payload"`
	Status      string    `json:"status"` // pending, success, failed, retrying
	RetryCount  int       `json:"retry_count"`
	MaxRetry    int       `json:"max_retry"`
	NextRetryAt time.Time `json:"next_retry_at"`
	ErrorMsg    string    `json:"error_msg,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

/*func NewPushService(ctx context.Context) *PushService {
	return &PushService{ctx: ctx}
}

// PushWithRetry 带重试机制的推送
func (s *PushService) PushWithRetry(data interface{}) error {
	//maxRetries := 3

	//err := gretry.Try(s.ctx, func(ctx context.Context) error {
	//	// 执行推送逻辑
	//	return s.doPush(data)
	//}, gretry.Option{
	//	Max:   maxRetries,
	//	Delay: 2 * time.Second, // 每次重试间隔2秒
	//})

	//if err != nil {
	//	g.Log().Error(s.ctx, fmt.Sprintf("推送失败，已重试%d次: %v", maxRetries, err))
	//	// 保存到失败队列，等待人工处理
	//	s.saveToFailedQueue(data, err)
	//	return err
	//}

	g.Log().Info(s.ctx, "推送成功")
	return nil
}

// doPush 执行实际推送
func (s *PushService) doPush(data interface{}) error {
	// 这里实现你的推送逻辑
	// 例如：调用下游API
	client := g.Client()
	client.SetTimeout(10 * time.Second)

	resp, err := client.Post(s.ctx, "http://downstream-api/push", data)
	if err != nil {
		return fmt.Errorf("推送请求失败: %w", err)
	}
	defer resp.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("推送返回错误状态码: %d", resp.StatusCode)
	}

	return nil
}

// saveToFailedQueue 保存失败记录到队列
func (s *PushService) saveToFailedQueue(data interface{}, err error) {
	// 保存到数据库或消息队列的死信队列
	record := PushRecord{
		BusinessId: generateBusinessId(),
		Payload:    fmt.Sprintf("%+v", data),
		Status:     "failed",
		RetryCount: 3,
		MaxRetry:   3,
		ErrorMsg:   err.Error(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// 保存到数据库
	// dao.PushRecord.Insert(s.ctx, record)

	g.Log().Warning(s.ctx, fmt.Sprintf("推送失败记录已保存: %+v", record))
}

// RetryFailedPushes 重试失败的推送（定时任务）
func (s *PushService) RetryFailedPushes() {
	// 从数据库中查询需要重试的记录
	// records := dao.PushRecord.Where("status", "failed").Where("next_retry_at < ?", time.Now()).All()

	// for _, record := range records {
	//     s.retrySingle(record)
	// }
}

// retrySingle 重试单条记录
func (s *PushService) retrySingle(record PushRecord) {
	if record.RetryCount >= record.MaxRetry {
		// 超过最大重试次数，标记为需要人工介入
		s.markNeedManualIntervention(record)
		return
	}

	err := s.doPush(record.Payload)
	if err != nil {
		// 更新重试次数和下次重试时间
		record.RetryCount++
		record.NextRetryAt = time.Now().Add(time.Duration(record.RetryCount*5) * time.Minute)
		record.ErrorMsg = err.Error()
		// dao.PushRecord.Update(record)
	} else {
		// 推送成功
		record.Status = "success"
		// dao.PushRecord.Update(record)
	}
}

// markNeedManualIntervention 标记需要人工介入
//func (s *PushService) markNeedManualIntervention(record PushRecord) {
//	record.Status = "need_manual"
//	// dao.PushRecord.Update(record)
//
//	// 发送告警通知相关人员
//	g.Log().Alert(s.ctx, fmt.Sprintf("【需要人工介入】推送记录 %s 已超过最大重试次数", record.BusinessId))
//}

// GetFailedRecords 获取失败记录列表（用于管理后台展示）
func (s *PushService) GetFailedRecords(page, pageSize int) ([]PushRecord, int, error) {
	// 从数据库查询
	// records, total := dao.PushRecord.Where("status IN ?", []string{"failed", "need_manual"}).Page(page, pageSize).AllAndCount()

	return []PushRecord{}, 0, nil
}

// ManualRetry 人工重试
func (s *PushService) ManualRetry(recordId int64) error {
	// record := dao.PushRecord.FindOne(recordId)
	// if record == nil {
	//     return errors.New("记录不存在")
	// }

	// 重置重试次数
	// record.RetryCount = 0
	// record.Status = "pending"
	// dao.PushRecord.Update(record)

	// 立即执行重试
	// go s.retrySingle(*record)

	return nil
}

func generateBusinessId() string {
	return fmt.Sprintf("push_%d", time.Now().UnixNano())
}
*/
