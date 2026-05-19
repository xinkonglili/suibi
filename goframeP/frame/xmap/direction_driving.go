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

type DirectionDrivingResult struct {
	Distance int64 `json:"distance"` // 米
	Duration int64 `json:"duration"` // 秒
}

// DirectionDriving 驾车路径规划 GCJ02坐标系
func (a *Amap) DirectionDriving(from, to MapPoint) *DirectionDrivingResult {
	apiUrl, _ := url.Parse("/v3/direction/driving?parameters")
	redisKey := consts2.RedisAmapDirectionDrivingKey + fmt.Sprintf("%s_%s", from.ToString(), to.ToString())
	res := &DirectionDrivingResult{}
	//var cacheExpiration int64 = 600 // 缓存5分钟

	// 先检查本地内存缓存
	if cacheData, err := consts2.GCtx.Get(consts2.SwitchCTX, redisKey); err == nil && cacheData != nil {
		if err = gjson.Unmarshal([]byte(cacheData.String()), res); err == nil {
			return res
		} else {
			g.Log().Warningf(a.Ctx, "获取本地缓存驾车路径规划失败: key: %s, err: %v", redisKey, err)
		}
	}
	//// 检查Redis缓存
	//if cacheData, err := redis.Get(a.Ctx, redisKey); err == nil && cacheData != "" {
	//	// 如果Redis中有数据，同步到本地缓存
	//	if cacheErr := consts.GCtx.Set(consts.SwitchCTX, redisKey, cacheData, time.Duration(900)*time.Second); cacheErr != nil {
	//		g.Log().Warningf(a.Ctx, "同步驾车路径规划到本地内存失败: key: %s, err: %v", redisKey, cacheErr)
	//	} else if err = gjson.Unmarshal([]byte(cacheData), res); err != nil {
	//		g.Log().Warningf(a.Ctx, "解析Redis缓存驾车路径规划失败: key: %v, json: %v, err: %v", redisKey, cacheData, err)
	//	} else {
	//		return res
	//	}
	//} else if err != nil {
	//	g.Log().Warningf(a.Ctx, "获取Redis缓存驾车路径规划失败: key: %s, err: %v", redisKey, err)
	//}

	amapKey, err := g.Cfg().Get(a.Ctx, "amap_key")
	if err != nil {
		g.Log().Warningf(a.Ctx, "[amap] [DirectionDriving] ERR: 高德地图API KEY未配置")
		return nil
	}

	// 调用高德地图API
	params := url.Values{
		"key":         {amapKey.String()},
		"origin":      {from.ToString()},
		"destination": {to.ToString()},
		"strategy":    {"0"},    // 驾车选择策略 0-速度优先
		"extensions":  {"base"}, // 返回结果控制 base-基础信息，all-全部信息
		"nosteps":     {"1"},    // 是否返回路段信息 0-返回，1-不返回
		"output":      {"json"},
	}
	apiUrl.RawQuery = params.Encode()
	resp, err := request.Get("amap").Do(a.Ctx, apiUrl.String(), nil)
	if err != nil {
		g.Log().Errorf(a.Ctx, "[amap] [DirectionDriving] 高德地图API请求失败, from:%v, to:%v, ERR: %v", from, to, err.Error())
		return nil
	} else if resp == "" {
		g.Log().Errorf(a.Ctx, "[amap] [DirectionDriving] ERR: 高德地图API返回数据为空, from:%v, to:%v, ERR: %v", from, to, err)
		return nil
	}
	respJson := gjson.New(resp)
	if respJson.Get("status").Int() != 1 {
		g.Log().Errorf(a.Ctx, "[amap] [DirectionDriving] ERR: 高德地图API返回错误, from:%v, to:%v, msg: %v", from, to, respJson.Get("message").String())
		return nil
	}

	paths := respJson.Get("route.paths").Array()
	if len(paths) == 0 {
		g.Log().Warningf(a.Ctx, "[amap] [DirectionDriving] ERR: 高德地图API返回数据缺少距离结果, from:%v, to:%v", from, to)
		return nil
	}
	firstResult := gconv.MapStrStr(paths[0])
	res = &DirectionDrivingResult{
		Distance: gconv.Int64(firstResult["distance"]),
		Duration: gconv.Int64(firstResult["duration"]),
	}

	// 缓存经纬度和地址信息到Redis和本地内存
	//if cacheData, err := gjson.Marshal(res); err == nil {
	//	// 同时更新Redis和本地缓存
	//	if _, redisErr := redis.Set(a.Ctx, redisKey, string(cacheData), cacheExpiration); redisErr != nil {
	//		g.Log().Warningf(a.Ctx, "缓存驾车路径规划到Redis失败: key: %s, err: %v", redisKey, redisErr)
	//	}
	//	if cacheErr := consts.GCtx.Set(consts.SwitchCTX, redisKey, string(cacheData), time.Duration(900)*time.Second); cacheErr != nil {
	//		g.Log().Warningf(a.Ctx, "缓存驾车路径规划到本地内存失败: key: %s, err: %v", redisKey, cacheErr)
	//	}
	//}

	return res
}
