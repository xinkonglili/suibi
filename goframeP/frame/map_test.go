package main

import (
	"context"
	"fmt"
	"goframeP/frame/consts"
	"goframeP/frame/request"
	"goframeP/frame/xcity"
	"goframeP/frame/xmap"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func TestMap(t *testing.T) {
	amp := xmap.NewAmap(gctx.New())
	// 使用有效的北京坐标（天安门附近）
	endAddrInfo := amp.LatAndLngToAddr(xmap.NewMapPoint("39.9042", "116.4074"))
	if endAddrInfo == nil {
		t.Log("未获取到地址信息（请检查 amap_key 配置或网络连接）")
		return
	}
	t.Logf("获取到地址: %+v", endAddrInfo)
}

func TestLatLngZeroZero(t *testing.T) {
	ctx := context.Background()

	apiUrl, _ := url.Parse("https://restapi.amap.com/v3/geocode/regeo")
	params := url.Values{}
	params.Set("key", "e0916263d2b0ed241096e88d744b7564")
	params.Set("location", "0.0,0.0")
	params.Set("extensions", "base")
	apiUrl.RawQuery = params.Encode()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, apiUrl.String(), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("原始返回:", string(body))

	respJson := gjson.New(string(body))
	fmt.Println("status :", respJson.Get("status").Int())
	fmt.Println("info   :", respJson.Get("info").String())
	fmt.Println("message:", respJson.Get("message").String())

	if respJson.Get("status").Int() == 1 {
		fmt.Println("formatted_address:", respJson.Get("regeocode.formatted_address").String())
		fmt.Println("city             :", respJson.Get("regeocode.addressComponent.city").String())
		fmt.Println("country          :", respJson.Get("regeocode.addressComponent.country").String())
	}

	address := respJson.Get("regeocode.formatted_address").String()
	if address == "[]" {
		fmt.Println("[amap] [address] ERR: 高德地图API返回数据缺少行政区划代码")
	} else {
		fmt.Println("address:", address)
	}

	// 解析行政区划代码
	adCode := respJson.Get("regeocode.addressComponent.adcode").String()
	if adCode == "[]" {
		fmt.Println("[amap] [LatAndLngToAddr] ERR: 高德地图API返回数据缺少行政区划代码")
	} else {
		fmt.Println("adCode:", adCode)
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

	var endAddrInfo *LatAndLngToAddrResult
	endAddrInfo = &LatAndLngToAddrResult{
		CityName: cityName,
		CityCode: xcity.ToSwitchCityCode(cityCode),
		AdCode:   adCode,
		Address:  address,
	}
	fmt.Println("endAddrInfo:", endAddrInfo)
	if endAddrInfo == nil {
		fmt.Println("[高德] 高德地图接口获取城市编码失败")
	}

}

type Amap struct {
	Ctx context.Context
}

type MapPoint struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type LatAndLngToAddrResult struct {
	CityName string `json:"city_name"` // 城市名称
	CityCode string `json:"city_code"` // 城市代码
	AdCode   string `json:"ad_code"`   // 行政区划代码
	Address  string `json:"address"`   // 以行政区划+道路+门牌号等信息组成的标准格式化地址
}

func TestMapE(t *testing.T) {
	signMap := make(map[string]any)
	signMap["key"] = "jininlie091626PP3d2b0ed24jj1096e88d744b7564920"
	signMap["location"] = "39.9042,116.4074"
	signMap["extensions"] = "base"
	s := gjson.New(signMap).String()
	fmt.Println(s)
	jsonString := gjson.New(signMap).MustToJsonString()
	fmt.Println(jsonString)

	jsonBytes, err := sonic.Marshal(signMap)
	if err != nil {
		g.Log().Errorf(nil, "[juliang] [MapToSignString] JSON序列化失败: %v, signMap: %+v", err, signMap)
	} else {
		fmt.Println(jsonBytes)
	}

	badData := map[string]interface{}{
		"func": func() {}, // 函数无法序列化
	}
	s1 := gjson.New(badData).String()
	fmt.Println(s1)
	jsonString1 := gjson.New(badData).MustToJsonString() //这种会panic
	fmt.Println(jsonString1)
}

// LatAndLngToAddr 坐标转换为地址 lat:纬度 lng:经度 GCJ02坐标系
func (a *Amap) LatAndLngToAddr(mp MapPoint) (addr *LatAndLngToAddrResult) {
	apiUrl, _ := url.Parse("/v3/geocode/regeo")
	addr = &LatAndLngToAddrResult{}
	redisKey := consts.RedisAmapLatLngKey + fmt.Sprintf("%s_%s", mp.Latitude, mp.Longitude)
	//var cacheExpiration int64 = 3600 // 缓存1小时

	// 先检查本地内存缓存
	if cacheData, err := consts.GCtx.Get(consts.SwitchCTX, redisKey); err == nil && cacheData != nil {
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
	resp, err := request.Get("amap").Do(a.Ctx, apiUrl.String(), nil)
	if err != nil {
		g.Log().Errorf(a.Ctx, "[amap] [LatAndLngToAddr] 高德地图API请求失败, lat:%v, lng:%v, ERR: %v", mp.Latitude, mp.Longitude, err.Error())
		return nil
	} else if resp == "" {
		g.Log().Errorf(a.Ctx, "[amap] [LatAndLngToAddr] ERR: 高德地图API返回数据为空, lat:%v, lng:%v", mp.Latitude, mp.Longitude)
		return nil
	}
	respJson := gjson.New(resp)
	if respJson.Get("status").Int() != 1 {
		g.Log().Errorf(a.Ctx, "[amap] [LatAndLngToAddr] ERR: 高德地图API返回错误, lat:%v, lng:%v, msg: %v", mp.Latitude, mp.Longitude, respJson.Get("message").String())
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

	// 缓存经纬度和地址信息到Redis和本地内存
	//if cacheData, err := gjson.Marshal(addr); err == nil {
	//	// 同时更新Redis和本地缓存
	//	//if _, redisErr := redis.Set(a.Ctx, redisKey, string(cacheData), cacheExpiration); redisErr != nil {
	//	//	g.Log().Warningf(a.Ctx, "缓存城市经纬度到Redis失败: key: %s, err: %v", redisKey, redisErr)
	//	//}
	//	//if cacheErr := consts.GCtx.Set(consts.SwitchCTX, redisKey, string(cacheData), time.Duration(900)*time.Second); cacheErr != nil {
	//	//	g.Log().Warningf(a.Ctx, "缓存城市代码到本地内存失败: key: %s, err: %v", redisKey, cacheErr)
	//	//}
	//}

	return
}
func (ap *MapPoint) ToString() string {
	return fmt.Sprintf("%v,%v", ap.GetLng(), ap.GetLat())
}

// 纬度小数点不超过6位（0.1米精度）
func (ap *MapPoint) GetLat() string {
	parts := strings.SplitN(ap.Latitude, ".", 2)
	if len(parts) != 2 || len(parts[1]) <= 6 {
		return ap.Latitude
	}
	return fmt.Sprintf("%s.%s", parts[0], parts[1][:6])
}

// 经度小数点不超过6位（0.1米精度）
func (ap *MapPoint) GetLng() string {
	parts := strings.SplitN(ap.Longitude, ".", 2)
	if len(parts) != 2 || len(parts[1]) <= 6 {
		return ap.Longitude
	}
	return fmt.Sprintf("%s.%s", parts[0], parts[1][:6])
}

//func Set(ctx context.Context, key string, value string, t int64) (result bool, err error) {
//	gVal, err := g.Redis().Set(ctx, getKey(key), value)
//	if err != nil {
//		return result, err
//	}
//
//	//timePoint := time.Now()
//	//defer func(ctx context.Context, timePoint time.Time, key, value string) {
//	//	go func(ctx context.Context, timePoint time.Time, key, value string) {
//	//		g.Log(RedisPrefix).Debug(ctx, fmt.Sprintf("[ %d ms]", time.Since(timePoint).Milliseconds())+" Set "+key+" "+value)
//	//	}(ctx, timePoint, key, value)
//	//}(ctx, timePoint, key, value)
//
//	g.Redis().Expire(ctx, getKey(key), t)
//	return gVal.Bool(), err
//}

//func getKey(key string) string {
//	RedisOnce.Do(RedisInit)
//	return consts.RedisPrefix + key
//}
