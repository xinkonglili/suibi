package common

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
)

// NilData 空返回值，对象形式{}
var NilData = map[int]struct{}{}

func ErrLogAndMsg(ctx context.Context, logErrs ...interface{}) {
	g.Log().Error(ctx, logErrs...)
}

func Log(ctx context.Context, title string, data ...interface{}) {
	var newData string
	for _, item := range data {
		newItem, _ := InterfaceToString(item)
		newData += "|" + newItem
	}
	g.Log().Info(ctx, title+":"+newData)
}

func InterfaceToString(param interface{}) (string, error) {
	thisString := ""
	switch param.(type) {
	case string:
		if value, ok := param.(string); ok {
			thisString = value
		}
	case float64:
		if value, ok := param.(float64); ok {
			thisString = strconv.FormatFloat(value, 'f', -1, 64)
		}
	case int:
		if value, ok := param.(int); ok {
			thisString = strconv.Itoa(value)
		}
	case int64:
		if value, ok := param.(int64); ok {
			thisString = strconv.FormatInt(value, 10)
		}
	case float32:
		if value, ok := param.(float32); ok {
			thisString = strconv.FormatFloat(float64(value), 'f', -1, 32)
		}
	case bool:
		if value, ok := param.(bool); ok {
			if value == true {
				thisString = "true"
			} else {
				thisString = "false"
			}
		}
	case json.Number:
		if value, ok := param.(json.Number); ok {
			thisString = value.String()
		}
	default:
		return "", errors.New("unknow type")
	}

	return thisString, nil
}
