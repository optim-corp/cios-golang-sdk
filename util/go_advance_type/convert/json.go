package convert

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"

	"github.com/optim-corp/cios-golang-sdk/util/go_advance_type/check"

	jsoniter "github.com/json-iterator/go"
)

func MustIndentJson(data interface{}) (result string) {
	result, _ = IndentJson(data)
	return
}
func MustCompactJson(data interface{}) (result string) {
	result, _ = CompactJson(data)
	return
}
func MustMarshalJson(data interface{}) (result []byte) {
	result, _ = MarshalJson(data)
	return
}
func MustJson(data interface{}) (result string) {
	result, _ = Json(data)
	return
}
func Json(data interface{}) (string, error) {
	j, err := MarshalJson(data)
	return MustStr(j), err
}
func MarshalJson(data interface{}) ([]byte, error) {
	return jsoniter.Marshal(data)
}
func UnMarshalJson(data interface{}, target interface{}) error {
	if !check.IsPtr(target) {
		return errors.New(fmt.Sprintf("not pointer %v", target))
	}
	switch v := data.(type) {
	case reflect.Value:
		return UnMarshalJson(v.Interface(), target)
	case []byte:
		return jsoniter.Unmarshal(v, target)
	case string:
		return jsoniter.Unmarshal([]byte(v), target)
	case io.Reader:
		if byts, err := ioutil.ReadAll(v); err != nil {
			return err
		} else {
			return jsoniter.Unmarshal(byts, target)
		}
	case io.ReadCloser:
		if byts, err := ioutil.ReadAll(v); err != nil {
			return err
		} else {
			return jsoniter.Unmarshal(byts, target)
		}
	default:
		byts, err := MarshalJson(data)
		if err != nil {
			return err
		}
		return UnMarshalJson(byts, target)
	}
}

func IndentJson(data interface{}) (result string, err error) {
	var buf bytes.Buffer
	switch v := data.(type) {
	case reflect.Value:
		return IndentJson(v.Interface())
	case []byte:
		if err = json.Indent(&buf, v, "", "  "); err == nil {
			result = buf.String()
		}
	case string:
		if err = json.Indent(&buf, []byte(v), "", "  "); err == nil {
			result = buf.String()
		}
	default:
		if byts, _err := MarshalJson(data); _err != nil {
			err = _err
		} else if err = json.Indent(&buf, byts, "", "  "); err == nil {
			result = buf.String()
		}
	}
	return
}
func CompactJson(data interface{}) (result string, err error) {
	var buf bytes.Buffer
	switch v := data.(type) {
	case reflect.Value:
		return CompactJson(v.Interface())
	case []byte:
		if err = json.Compact(&buf, v); err == nil {
			result = buf.String()
		}
	case string:
		if err = json.Compact(&buf, []byte(v)); err == nil {
			result = buf.String()
		}
	default:
		if byts, _err := MarshalJson(data); _err != nil {
			err = _err
		} else if err = json.Compact(&buf, byts); err == nil {
			result = buf.String()
		}
	}
	return
}
