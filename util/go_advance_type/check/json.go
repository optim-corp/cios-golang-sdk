package check

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"

	jsoniter "github.com/json-iterator/go"
)

func AreEqualJson(data1, data2 interface{}) bool {
	if reflect.DeepEqual(data1, data2) {
		return true
	}
	var _data1, _data2 interface{}
	switch v := data1.(type) {
	case io.Reader:
		if json.NewDecoder(v).Decode(&_data1) != nil {
			return false
		}
	case string:
		if unmarshal([]byte(v), &_data1) != nil {
			return false
		}
	case []byte:
		if unmarshal(v, &_data1) != nil {
			return false
		}
	default:
		if byts, err := marshal(data1); err != nil {
			return false
		} else if unmarshal(byts, &_data1) != nil {
			return false
		}
	}

	switch v := data2.(type) {
	case io.Reader:
		if json.NewDecoder(v).Decode(&_data2) != nil {
			return false
		}
	case string:
		if unmarshal([]byte(v), &_data2) != nil {
			return false
		}
	case []byte:
		if unmarshal(v, &_data2) != nil {
			return false
		}
	default:
		if byts, err := marshal(data2); err != nil {
			return false
		} else if unmarshal(byts, &_data2) != nil {
			return false
		}
	}
	return reflect.DeepEqual(_data1, _data2)
}

func unmarshal(data interface{}, target interface{}) error {
	if !IsPtr(target) {
		return errors.New(fmt.Sprintf("not pointer %v", target))
	}
	switch v := data.(type) {
	case reflect.Value:
		return unmarshal(v.Interface(), target)
	case []byte:
		return jsoniter.Unmarshal(v, target)
	case string:
		return jsoniter.Unmarshal([]byte(v), target)
	default:
		byts, err := marshal(data)
		if err != nil {
			return err
		}
		return unmarshal(byts, target)
	}
}

func marshal(data interface{}) ([]byte, error) {
	return jsoniter.Marshal(data)
}
