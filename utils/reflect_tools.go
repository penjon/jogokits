package utils

import (
	"reflect"
	"strconv"
)

func SetValue(field *reflect.Value, value string) error {
	if nil != field {
		switch field.Type().Kind() {
		case reflect.String:
			field.SetString(value)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if v, e := strconv.ParseInt(value, 10, 64); e != nil {
				return e
			} else {
				field.SetInt(v)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if v, e := strconv.ParseUint(value, 10, 64); e != nil {
				return e
			} else {
				field.SetUint(v)
			}
		case reflect.Float32, reflect.Float64:
			if v, e := strconv.ParseFloat(value, 64); e != nil {
				return e
			} else {
				field.SetFloat(v)
			}
		case reflect.Bool:
			v, _ := strconv.ParseBool(value)
			field.SetBool(v)
		}
	}
	return nil
}
