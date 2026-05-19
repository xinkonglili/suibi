package xmap

import (
	"fmt"
	consts2 "goframeP/consts"
	"goframeP/frame/request"
	"net/url"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type DistanceResult struct {
	Distance string `json:"distance"` // 米
	Duration string `json:"duration"` // 秒
}

// Distance 坐标转换为距离 lat:纬度 lng:经度 GCJ02坐标系
func (a *Amap) Distance(from, to MapPoint) *DistanceResult {
	apiUrl, _ := url.Parse("/v3/distance")
	redisKey := consts2.RedisAmapDistanceKey + fmt.Sprintf("%s_%s", from.ToString(), to.ToString())
	distanceRes := &DistanceResult{}
	//var cacheExpiration int64 = 3600 // 缓存1小时

	// 先检查本地内存缓存
	if cacheData, err := consts2.GCtx.Get(consts2.SwitchCTX, redisKey); err == nil && cacheData != nil {
		if err = gjson.Unmarshal([]byte(cacheData.String()), distanceRes); err == nil {
			return distanceRes
		} else {
			g.Log().Warningf(a.Ctx, "获取本地缓存经纬度距离失败: key: %s, err: %v", redisKey, err)
		}
	}
	// 检查Redis缓存
	//if cacheData, err := redis.Get(a.Ctx, redisKey); err == nil && cacheData != "" {
	//	// 如果Redis中有数据，同步到本地缓存
	//	if cacheErr := consts.GCtx.Set(consts.SwitchCTX, redisKey, cacheData, time.Duration(900)*time.Second); cacheErr != nil {
	//		g.Log().Warningf(a.Ctx, "同步经纬度距离到本地内存失败: key: %s, err: %v", redisKey, cacheErr)
	//	} else if err = gjson.Unmarshal([]byte(cacheData), distanceRes); err != nil {
	//		g.Log().Warningf(a.Ctx, "解析Redis缓存经纬度距离失败: key: %v, json: %v, err: %v", redisKey, cacheData, err)
	//	} else {
	//		return distanceRes
	//	}
	//} else if err != nil {
	//	g.Log().Warningf(a.Ctx, "获取Redis缓存经纬度距离失败: key: %s, err: %v", redisKey, err)
	//}

	amapKey, err := g.Cfg().Get(a.Ctx, "amap_key")
	if err != nil {
		g.Log().Warningf(a.Ctx, "[amap] [Distance] ERR: 高德地图API KEY未配置")
		return nil
	}

	// 调用高德地图API
	params := url.Values{
		"key":         {amapKey.String()},
		"origins":     {from.ToString()},
		"destination": {to.ToString()},
		"type":        {"1"}, // 0=直线距离，1=驾车距离，3=步行距离
	}
	apiUrl.RawQuery = params.Encode()
	resp, err := request.Get("amap").Do(a.Ctx, apiUrl.String(), nil)
	if err != nil {
		g.Log().Errorf(a.Ctx, "[amap] [Distance] 高德地图API请求失败, from:%v, to:%v, ERR: %v", from, to, err.Error())
		return nil
	} else if resp == "" {
		g.Log().Errorf(a.Ctx, "[amap] [Distance] ERR: 高德地图API返回数据为空, from:%v, to:%v, ERR: %v", from, to, err)
		return nil
	}
	respJson := gjson.New(resp)
	if respJson.Get("status").Int() != 1 {
		g.Log().Errorf(a.Ctx, "[amap] [Distance] ERR: 高德地图API返回错误, from:%v, to:%v, msg: %v", from, to, respJson.Get("message").String())
		return nil
	}

	results := respJson.Get("results").Array()
	if len(results) == 0 {
		g.Log().Warningf(a.Ctx, "[amap] [Distance] ERR: 高德地图API返回数据缺少距离结果, from:%v, to:%v", from, to)
		return nil
	}
	firstResult := gconv.MapStrStr(results[0])
	distanceRes = &DistanceResult{
		Distance: firstResult["distance"],
		Duration: firstResult["duration"],
	}

	// 缓存经纬度和地址信息到Redis和本地内存
	//if cacheData, err := gjson.Marshal(distanceRes); err == nil {
	//	// 同时更新Redis和本地缓存
	//	if _, redisErr := redis.Set(a.Ctx, redisKey, string(cacheData), cacheExpiration); redisErr != nil {
	//		g.Log().Warningf(a.Ctx, "缓存经纬度距离到Redis失败: key: %s, err: %v", redisKey, redisErr)
	//	}
	//	if cacheErr := consts.GCtx.Set(consts.SwitchCTX, redisKey, string(cacheData), time.Duration(900)*time.Second); cacheErr != nil {
	//		g.Log().Warningf(a.Ctx, "缓存经纬度距离到本地内存失败: key: %s, err: %v", redisKey, cacheErr)
	//	}
	//}

	return distanceRes
}
