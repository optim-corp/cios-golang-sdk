package check

import (
	"reflect"
)

func IsBool(val interface{}) bool {
	return val == true || val == false
}

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	ref := reflect.ValueOf(i)
	return ref.Kind() == reflect.Ptr && ref.IsNil()
}

func IsPtr(in interface{}) bool {
	return reflect.ValueOf(in).Kind() == reflect.Ptr
}

func IsNumber(in interface{}) bool {
	switch in.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	}
	return false
}
