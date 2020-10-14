/*
 * @Descripttion:
 * @Author: Sun
 * @Date: 2020-10-14 17:30:38
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2020-10-14 17:41:21
 */
package utils

import (
	"reflect"
	mapset"github.com/deckarep/golang-set"
)

// StructToMap 结构体转Map
func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// StructToMap 结构体转Map
func StructToMapFilterField(obj interface{},fields... string) map[string]interface{} {
	set := newSet(fields...)
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		if set.Contains(t.Field(i).Name){
			continue
		}
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func newSet(fields... string)mapset.Set{
	set := mapset.NewSet()
	for _,v := range fields{
		set.Add(v)
	}
	return set
}
