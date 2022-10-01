package converter

import (
	"reflect"
	"strconv"
)

// Func to convert interface{} to int32
func InterfaceToFloat64(data interface{}) float64 {
	typ := reflect.TypeOf(data)
	val := reflect.ValueOf(data)
	switch typ.Kind() {
	case reflect.String:
		dataInt, _ := strconv.ParseInt(val.String(), 10, 64)
		return float64(dataInt)
	case reflect.Float32:
		return float64(data.(float32))
	case reflect.Float64:
		return data.(float64)
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
		return float64(val.Int())
	default:
		return float64(0)
	}
}
