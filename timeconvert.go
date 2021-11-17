package timeconvert

import (
	"reflect"
	"time"
)

func TimeConv(items ...interface{}) time.Time {
	res := time.Time{}
	for _, item := range items {
		res = convertToTime(item)
	}
	return res
}

func convertToTime(data interface{}) time.Time {
	if data == nil {
		return time.Time{}
	}
	v := reflect.ValueOf(data)
	v = reflect.Indirect(v)
	switch v.Type().Kind() {
	case reflect.String:
		{
			layoutFormat := "2006-01-02 15:04:05"
			strconv, _ := time.Parse(layoutFormat, v.String())
			return strconv

		}
	case reflect.Int:
		{
			tm := time.Unix(v.Int(), 0)
			return tm

		}
	case reflect.Struct:
		{
			t, _ := data.(time.Time)
			return t
		}
	default:
		return time.Time{}
	}
}
