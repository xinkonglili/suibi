package xmap

import (
	"fmt"
	consts2 "goframeP/consts"
	"goframeP/frame/xcity"
	"net/url"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type LatAndLngToAddrResult struct {
	CityName string `json:"city_name"` // 城市名称
	CityCode string `json:"city_code"` // 城市代码
	AdCode   string `json:"ad_code"`   // 行政区划代码
	Address  string `json:"address"`   // 以行政区划+道路+门牌号等信息组成的标准格式化地址
}

// LatAndLngToAddr 坐标转换为地址 lat:纬度 lng:经度 GCJ02坐标系
func (a *Amap) LatAndLngToAddr(mp MapPoint) (addr *LatAndLngToAddrResult) {
	apiUrl, _ := url.Parse("/v3/geocode/regeo")
	addr = &LatAndLngToAddrResult{}
	redisKey := consts2.RedisAmapLatLngKey + fmt.Sprintf("%s_%s", mp.Latitude, mp.Longitude)
	//var cacheExpiration int64 = 3600 // 缓存1小时

	// 先检查本地内存缓存
	if cacheData, err := consts2.GCtx.Get(consts2.SwitchCTX, redisKey); err == nil && cacheData != nil {
		if err = gjson.Unmarshal([]byte(cacheData.String()), addr); err == nil {
			return addr
		} else {
			g.Log().Warningf(a.Ctx, "获取本地缓存城市经纬度失败: key: %s, err: %v", redisKey, err)
		}
	}
	//// 检查Redis缓存
	//if cacheData, err := redis.Get(a.Ctx, redisKey); err == nil && cacheData != "" {
	//	// 如果Redis中有数据，同步到本地缓存
	//	if cacheErr := consts.GCtx.Set(consts.SwitchCTX, redisKey, cacheData, time.Duration(900)*time.Second); cacheErr != nil {
	//		g.Log().Warningf(a.Ctx, "同步城市经纬度到本地内存失败: key: %s, err: %v", redisKey, cacheErr)
	//	} else if err = gjson.Unmarshal([]byte(cacheData), addr); err != nil {
	//		g.Log().Warningf(a.Ctx, "解析Redis缓存城市经纬度失败: key: %v, json: %v, err: %v", redisKey, cacheData, err)
	//	} else {
	//		return addr
	//	}
	//} else if err != nil {
	//	g.Log().Warningf(a.Ctx, "获取Redis缓存城市经纬度失败: key: %s, err: %v", redisKey, err)
	//}

	amapKey, err := g.Cfg().Get(a.Ctx, "amap_key")
	if err != nil {
		g.Log().Warningf(a.Ctx, "[amap] [LatAndLngToAddr] ERR: 高德地图API KEY未配置")
		return nil
	}

	// 调用高德地图API
	params := url.Values{
		"key":        {amapKey.String()},
		"location":   {mp.ToString()},
		"extensions": {"base"},
	}
	apiUrl.RawQuery = params.Encode()

	// 使用 GoFrame 自带 HTTP 客户端发送请求
	client := g.Client()
	client.SetPrefix("https://restapi.amap.com") // 高德地图 API 基础地址

	resp, err := client.Get(a.Ctx, apiUrl.String())
	if err != nil {
		g.Log().Errorf(a.Ctx, "[amap] [LatAndLngToAddr] 高德地图API请求失败, lat:%v, lng:%v, ERR: %v", mp.Latitude, mp.Longitude, err.Error())
		return nil
	}
	defer resp.Close()

	respBody := resp.ReadAllString()
	if respBody == "" {
		g.Log().Errorf(a.Ctx, "[amap] [LatAndLngToAddr] ERR: 高德地图API返回数据为空, lat:%v, lng:%v", mp.Latitude, mp.Longitude)
		return nil
	}
	respJson := gjson.New(respBody)
	if respJson.Get("status").Int() != 1 {
		g.Log().Errorf(a.Ctx, "[amap] [LatAndLngToAddr] ERR: 高德地图API返回错误, lat:%v, lng:%v, 原始响应: %v", mp.Latitude, mp.Longitude, respBody)
		return nil
	}

	// 解析地址信息
	address := respJson.Get("regeocode.formatted_address").String()
	// 解析行政区划代码
	adCode := respJson.Get("regeocode.addressComponent.adcode").String()
	if adCode == "" {
		g.Log().Warningf(a.Ctx, "[amap] [LatAndLngToAddr] ERR: 高德地图API返回数据缺少行政区划代码, lat:%v, lng:%v", mp.Latitude, mp.Longitude)
		return nil
	}
	if len(adCode) > 6 {
		adCode = adCode[len(adCode)-6:]
	}
	cityName := respJson.Get("regeocode.addressComponent.city").String()
	if cityName == "" || cityName == "[]" {
		// 直辖市没有city字段，使用province字段代替
		cityName = respJson.Get("regeocode.addressComponent.province").String()
	}
	var cityCode string
	if len(adCode) >= 4 {
		cityCode = adCode[:4] + "00"
	}

	addr = &LatAndLngToAddrResult{
		CityName: cityName,
		CityCode: xcity.ToSwitchCityCode(cityCode),
		AdCode:   adCode,
		Address:  address,
	}
	//
	//// 缓存经纬度和地址信息到Redis和本地内存
	//if cacheData, err := gjson.Marshal(addr); err == nil {
	//	// 同时更新Redis和本地缓存
	//	if _, redisErr := redis.Set(a.Ctx, redisKey, string(cacheData), cacheExpiration); redisErr != nil {
	//		g.Log().Warningf(a.Ctx, "缓存城市经纬度到Redis失败: key: %s, err: %v", redisKey, redisErr)
	//	}
	//	if cacheErr := consts.GCtx.Set(consts.SwitchCTX, redisKey, string(cacheData), time.Duration(900)*time.Second); cacheErr != nil {
	//		g.Log().Warningf(a.Ctx, "缓存城市代码到本地内存失败: key: %s, err: %v", redisKey, cacheErr)
	//	}
	//}

	return
}
