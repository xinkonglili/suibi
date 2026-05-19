package request

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"goframeP/common/xapp/alitrip_app"
	"goframeP/common/xhttp"
	"goframeP/common/xsign/alitrip"
	"goframeP/consts/const_alitrip"
	"net/http"
	"net/url"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// 通用连接器

func init() {
	register(&AlitripClient{BizCode: const_alitrip.BizCode})
	register(&AlitripClient{BizCode: const_alitrip.BizCodeXh})
	register(&AlitripClient{BizCode: const_alitrip.BizCodeJz})
}

type AlitripClient struct {
	BizCode   string
	doBizCode string
}

func (c *AlitripClient) GetBizCode() string {
	return c.BizCode
}

func (c *AlitripClient) SetDoBizCode(doBizCode string) XClient {
	return c
}

func (c *AlitripClient) Do(ctx context.Context, uri string, data any, options ...Options) (resp string, err error) {
	apiUrl, _ := url.Parse(uri)
	host, err := g.Cfg().Get(ctx, "alitrip_host")
	if err != nil {
		return "", gerror.Wrap(err, "[request] ["+c.BizCode+"] alitrip_host为空")
	}

	var curl = xhttp.New(host.String()).SetTitle("请求"+c.BizCode).SetHeader("Content-Type", "application/json; charset=utf-8")

	// 业务参数
	postData := gconv.Map(data)
	appInfo, err := alitrip_app.ToAppInfo(ctx, gconv.String(postData["client_id"]))
	if err != nil {
		return "", gerror.Wrap(err, "[request] ["+c.BizCode+"] appInfo")
	}
	appKey := gconv.String(appInfo["app_key"])
	if appKey == "" {
		return "", gerror.Wrap(err, "[request] ["+c.BizCode+"] appKey为空")
	}
	appSecret := gconv.String(appInfo["app_secret"])
	if appSecret == "" {
		return "", gerror.Wrap(err, "[request] ["+c.BizCode+"] appSecret为空")
	}

	postJson, err := JsonEncodeLikePHP(data)
	if err != nil {
		return "", gerror.Wrap(err, "[request] ["+c.BizCode+"] postJson marshal error")
	}

	// 公共参数补全
	commMap := gmap.NewStrAnyMap()
	commMap.Set("app_key", appKey)
	commMap.Set("v", const_alitrip.Version)
	commMap.Set("format", const_alitrip.RespFormat)
	commMap.Set("sign_method", const_alitrip.SignMethod)
	commMap.Set("method", apiUrl.Query().Get("method"))
	commMap.Set("timestamp", gtime.Now().Format(gtime.Datetime()))
	// commMap.Set("simplify", const_alitrip.RespSimplify)
	commMap.Set("partner_id", const_alitrip.PartnerId)
	// commMap.Set("session", "") // todo check

	queryData := commMap.Map()

	// 签名参数 = 公共参数 + 业务参数
	signMap := commMap.Clone()
	signMap.Set("param", string(postJson))

	var waitSignStr string
	if queryData["sign"], waitSignStr, err = alitrip.CalculateSign(ctx, signMap.Map(), "", appSecret); err != nil {
		g.Log().Infof(ctx, "待签名参数: %s", waitSignStr)
		return "", gerror.Wrap(err, "[request] ["+c.BizCode+"] sign计算失败")
	}

	// 设置请求头
	curl.SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	// 开始调用
	queryParams := apiUrl.Query()
	for k, v := range queryData {
		queryParams.Set(k, gconv.String(v))
	}

	apiUrl.RawQuery = queryParams.Encode()
	resp, err = curl.Post(ctx, apiUrl.String(), map[string]any{"param": string(postJson)})
	if err != nil {
		return "", gerror.Wrap(err, "[request] ["+c.BizCode+"] do alitrip curl error")
	} else if curl.GetResponseCode() != http.StatusOK {
		return "", gerror.New("[request] [" + c.BizCode + "] curl response code not 200, code: " + gconv.String(curl.GetResponseCode()))
	}

	return resp, nil
}

func JsonEncodeLikePHP(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return escapeUnicode(string(b)), nil
}

// 仅把中文转成 \uXXXX，其它字符保持不变
func escapeUnicode(s string) string {
	var buf bytes.Buffer
	for _, r := range s {
		if r > 127 {
			fmt.Fprintf(&buf, `\u%04x`, r)
		} else {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}
