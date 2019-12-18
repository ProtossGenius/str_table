package str_table

import (
	"fmt"
	"reflect"
)

type FStructAnalysis func(name string, val []interface{}) TableItf

func MapArrayToTable(name string, val []interface{}) TableItf {
	if val == nil || len(val) == 0 {
		return nil
	}
	first := val[0].(map[string]string)
	nameArr := []string{}
	for n, _ := range first {
		nameArr = append(nameArr, n)
	}
	res := NewTable(name, nameArr)
	for _, tmp := range val {
		mp := tmp.(map[string]string)
		tableRow := []string{}
		for _, n := range nameArr {
			tableRow = append(tableRow, mp[n])
		}
		res.AddRow(tableRow)
	}
	return res
}

func StructToTable(name string, val []interface{}) TableItf {
	if val == nil || len(val) == 0 {
		return nil
	}
	mapList := []interface{}{}
	for _, v := range val {
		t := reflect.TypeOf(v).Elem()
		vv := reflect.ValueOf(v).Elem()
		mp := map[string]string{}
		for i := 0; i < t.NumField(); i++ {
			filed := t.Field(i)
			fieldName := t.Name()
			fieldName = filed.Tag.Get("str_table")
			value := fmt.Sprintf("%v", vv.Field(i))
			mp[fieldName] = value
		}
		mapList = append(mapList, mp)
	}
	return MapArrayToTable(name, mapList)
}
