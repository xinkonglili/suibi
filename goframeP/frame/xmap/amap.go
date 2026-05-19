package xmap

import (
	"context"
	"fmt"
	"strings"
)

type Amap struct {
	Ctx context.Context
}

type MapPoint struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func NewAmap(ctx context.Context) *Amap {
	return &Amap{Ctx: ctx}
}

func NewMapPoint[T float64 | string](lat T, lng T) MapPoint {
	return MapPoint{
		Latitude:  fmt.Sprintf("%v", lat),
		Longitude: fmt.Sprintf("%v", lng),
	}
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
