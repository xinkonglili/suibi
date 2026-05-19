package request

import (
	"context"
	"goframeP/common/xhttp"
	"goframeP/common/xsign/txmap"
	"goframeP/consts"
	"net/http"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// 腾讯API连接器

func init() {
	register(&TxApiClient{BizCode: consts.TencentMapBizCode})
}

type TxApiClient struct {
	BizCode   string
	doBizCode string
}

func (c *TxApiClient) GetBizCode() string {
	return c.BizCode
}

func (c *TxApiClient) SetDoBizCode(doBizCode string) XClient {
	c.doBizCode = doBizCode
	return c
}

func (c *TxApiClient) Do(ctx context.Context, url string, data any, options ...Options) (resp string, err error) {
	host, _ := g.Cfg().Get(ctx, "tencent_map_host")
	key, _ := g.Cfg().Get(ctx, "tencent_map_key")
	postData := data.(map[string]any)
	postData["key"] = key.String()
	if postData["sig"], err = txmap.CalculateSign(ctx, url, postData); err != nil {
		return "", gerror.Wrap(err, "[txmap] calculate sign error")
	}
	method := http.MethodGet
	var curl = xhttp.New(host.String()).SetTitle("请求腾讯")
	if len(options) > 0 {
		if options[0].Timeout > 0 {
			curl = curl.SetTimeout(options[0].Timeout)
		}
		if !g.IsEmpty(options[0].Headers) {
			for k, v := range options[0].Headers {
				curl.SetHeader(k, v)
			}
		}
		method = options[0].Method
	}

	if method == http.MethodGet {
		resp, err = curl.SetHeader("Content-Type", "application/x-www-form-urlencoded").Get(ctx, url, postData)
	} else if method == http.MethodPost {
		resp, err = curl.Post(ctx, url, postData)
	} else {
		return "", gerror.New("tencent curl method not support: " + method)
	}

	if err != nil {
		return "", gerror.Wrap(err, "do tencent curl error")
	} else if curl.GetResponseCode() != http.StatusOK {
		return "", gerror.New("tencent curl response code not 200, code: " + gconv.String(curl.GetResponseCode()))
	}
	return resp, nil
}
