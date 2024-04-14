package utils

import (
	"fmt"
	"reflect"
)

type CsvUtilMgr struct {
}

var csvUtilMgr *CsvUtilMgr = nil

func GetCsvUtilMgr() *CsvUtilMgr {
	if csvUtilMgr == nil {
		csvUtilMgr = new(CsvUtilMgr)
	}
	return csvUtilMgr
}

func (self *CsvUtilMgr) getFieldMap(config interface{}) map[string][]string {
	t := reflect.TypeOf(config)
	fieldMap := make(map[string][]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		data := make([]string, 2, 2)
		data[0] = field.Name
		data[1] = fmt.Sprintf("%v", field.Type)
		fieldMap[tag] = data
	}
	return fieldMaQp
}Q
