package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type RequestBody struct {
	MainOrderId     string `json:"mainOrderId"`
	SubOrderId      string `json:"subOrderId"`
	CoopMainOrderId string `json:"coopMainOrderId"`
	CoopSubOrderId  string `json:"coopSubOrderId"`
	QueryPosFlag    bool   `json:"queryPosFlag"`
	QueryNavFlag    bool   `json:"queryNavFlag"`
	RouteType       string `json:"routeType"`
	InterlinkFlag   bool   `json:"interlinkFlag"`
}

func main() {
	requests := []struct {
		url  string
		body RequestBody
	}{
		{
			url: "https://switch/v1/order/orderDetail?bizCode=ao&timestamp=1778046181&sign=7E8A867D9AF9ECE4",
			body: RequestBody{
				MainOrderId:     "202661",
				SubOrderId:      "20260000161_SY_2_CC_T_2",
				CoopMainOrderId: "X02000161",
				CoopSubOrderId:  "1519510902000",
				QueryPosFlag:    false,
				QueryNavFlag:    false,
				RouteType:       "0",
				InterlinkFlag:   false,
			},
		},
	}

	client := &http.Client{Timeout: 10 * time.Second}
	successCount := 0
	failCount := 0

	for idx, req := range requests {
		bodyBytes, _ := json.Marshal(req.body)
		resp, err := client.Post(req.url, "application/json", bytes.NewBuffer(bodyBytes))
		if err != nil {
			failCount++
			fmt.Printf("[%d]  请求异常: mainOrderId=%s, err=%v\n", idx+1, req.body.MainOrderId, err)
			continue
		}
		respBody, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode == 200 {
			successCount++
			fmt.Printf("[%d]  成功: mainOrderId=%s, status=%d, resp=%s\n", idx+1, req.body.MainOrderId, resp.StatusCode, string(respBody))
		} else {
			failCount++
			fmt.Printf("[%d]  失败: mainOrderId=%s, status=%d, resp=%s\n", idx+1, req.body.MainOrderId, resp.StatusCode, string(respBody))
		}
	}

	fmt.Printf("\n========== 执行完毕: 共%d条, 成功%d条, 失败%d条 ==========\n", len(requests), successCount, failCount)
}
