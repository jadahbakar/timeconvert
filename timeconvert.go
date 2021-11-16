package timeconvert

import (
	"fmt"
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
	timeType := reflect.TypeOf(time.Time{})
	if data == nil {
		return time.Time{}
	}
	v := reflect.ValueOf(data)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(timeType) {
		return time.Time{}
	}
	if _, ok := v.Interface().(time.Time); ok {
		fmt.Println("Ok")
		res := v.Convert(timeType)
		return res.Interface().(time.Time).UTC().Format()
	} else {
		return time.Time{}
	}

	// return res
	// timeType := reflect.TypeOf(time.Time{})
	// v := reflect.ValueOf(data)
	// v = reflect.Indirect(v)
	// if !v.Type().ConvertibleTo(timeType) {
	// 	return time.Time{}
	// }
	// res := v.Convert(timeType)
	// return res.Interface().(time.Time)
	// return time.Time{}
}
