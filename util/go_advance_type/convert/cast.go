package convert

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"

	. "github.com/optim-corp/cios-golang-sdk/util/go_advance_type/util"
)

var (
	_err    = errors.New("missing type")
	uintErr = errors.New("cant cant to uint")
)

func SafeInt64(in interface{}, defaultNumber int64) int64 {
	result, err := Int64(in)
	if err != nil {
		result = defaultNumber
	}
	return result
}
func SafeInt32(in interface{}, defaultNumber int32) int32 {
	result, err := Int32(in)
	if err != nil {
		result = defaultNumber
	}
	return result
}
func SafeInt16(in interface{}, defaultNumber int16) int16 {
	result, err := Int16(in)
	if err != nil {
		result = defaultNumber
	}
	return result
}
func SafeInt8(in interface{}, defaultNumber int8) int8 {
	result, err := Int8(in)
	if err != nil {
		result = defaultNumber
	}
	return result
}
func SafeInt(in interface{}, defaultNumber int) int {
	result, err := Int(in)
	if err != nil {
		result = defaultNumber
	}
	return result
}

func SafeFloat32(in interface{}, defaultNumber float32) float32 {
	result, err := Float32(in)
	if err != nil {
		result = defaultNumber
	}
	return result
}
func SafeFloat64(in interface{}, defaultNumber float64) float64 {
	result, err := Float64(in)
	if err != nil {
		result = defaultNumber
	}
	return result
}
func SafeUint(in interface{}, defaultNumber uint) uint {
	result, err := Uint(in)
	if err != nil {
		result = defaultNumber
	}
	return result
}
func SafeUint64(in interface{}, defaultNumber uint64) uint64 {
	result, err := Uint64(in)
	if err != nil {
		result = defaultNumber
	}
	return result
}
func SafeUint32(in interface{}, defaultNumber uint32) uint32 {
	result, err := Uint32(in)
	if err != nil {
		result = defaultNumber
	}
	return result
}
func SafeUint16(in interface{}, defaultNumber uint16) uint16 {
	result, err := Uint16(in)
	if err != nil {
		result = defaultNumber
	}
	return result
}
func SafeUint8(in interface{}, defaultNumber uint8) uint8 {
	result, err := Uint8(in)
	if err != nil {
		result = defaultNumber
	}
	return result
}

func SafeStr(in interface{}, defaultStr string) string {
	result, err := Str(in)
	if err != nil {
		result = defaultStr
	}
	return result
}
func SafeBool(in interface{}, defaultVal bool) bool {
	result, err := Bool(in)
	if err != nil {
		result = defaultVal
	}
	return result
}

func MustInt64(in interface{}) (result int64) {
	result, _ = Int64(in)
	return
}
func MustInt32(in interface{}) (result int32) {
	result, _ = Int32(in)
	return
}
func MustInt16(in interface{}) (result int16) {
	result, _ = Int16(in)
	return
}
func MustInt8(in interface{}) (result int8) {
	result, _ = Int8(in)
	return
}
func MustInt(in interface{}) (result int) {
	result, _ = Int(in)
	return
}
func MustFloat64(in interface{}) (result float64) {
	result, _ = Float64(in)
	return
}
func MustFloat32(in interface{}) (result float32) {
	result, _ = Float32(in)
	return
}
func MustStr(in interface{}) (result string) {
	result, _ = Str(in)
	return
}
func MustBool(in interface{}) (result bool) {
	result, _ = Bool(in)
	return
}
func MustUint(in interface{}) (result uint) {
	result, _ = Uint(in)
	return
}
func MustUint64(in interface{}) (result uint64) {
	result, _ = Uint64(in)
	return
}
func MustUint32(in interface{}) (result uint32) {
	result, _ = Uint32(in)
	return
}
func MustUint16(in interface{}) (result uint16) {
	result, _ = Uint16(in)
	return
}
func MustUint8(in interface{}) (result uint8) {
	result, _ = Uint8(in)
	return
}

func Int(in interface{}) (int, error) {
	switch v := in.(type) {
	case nil:
		return 0, nil
	case float32:
		return int(v), nil
	case float64:
		return int(v), nil
	case int:
		return v, nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case uint:
		return int(v), nil
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint32:
		return int(v), nil
	case uint64:
		return int(v), nil
	case bool:
		return Is(v).Ok(1).No(0).Value().AsInt(), nil
	case string:
		if i, err := strconv.ParseInt(v, 0, 64); err != nil {
			return 0, err
		} else {
			return int(i), nil
		}
	case []byte:
		if i, err := strconv.ParseInt(string(v), 0, 64); err != nil {
			return 0, err
		} else {
			return int(i), nil
		}
	case reflect.Value:
		return Int(v.Interface())
	}
	if r := reflect.ValueOf(in); r.Kind() == reflect.Ptr {
		if r.IsNil() {
			return 0, nil
		}
		return Int(r.Elem().Interface())
	}
	return 0, _err
}
func Int8(in interface{}) (int8, error) {
	switch v := in.(type) {
	case nil:
		return 0, nil
	case float32:
		return int8(v), nil
	case float64:
		return int8(v), nil
	case int:
		return int8(v), nil
	case int8:
		return int8(v), nil
	case int16:
		return int8(v), nil
	case int32:
		return int8(v), nil
	case int64:
		return int8(v), nil
	case uint:
		return int8(v), nil
	case uint8:
		return int8(v), nil
	case uint16:
		return int8(v), nil
	case uint32:
		return int8(v), nil
	case uint64:
		return int8(v), nil
	case bool:
		return int8(Is(v).Ok(1).No(0).Value().AsInt()), nil
	case string:
		if i, err := strconv.ParseInt(v, 0, 64); err != nil {
			return 0, err
		} else {
			return int8(i), nil
		}
	case []byte:
		if i, err := strconv.ParseInt(string(v), 0, 64); err != nil {
			return 0, err
		} else {
			return int8(i), nil
		}
	case reflect.Value:
		return Int8(v.Interface())
	}
	if r := reflect.ValueOf(in); r.Kind() == reflect.Ptr {
		if r.IsNil() {
			return 0, nil
		}
		return Int8(r.Elem().Interface())
	}
	return 0, _err
}
func Int16(in interface{}) (int16, error) {
	switch v := in.(type) {
	case nil:
		return 0, nil
	case float32:
		return int16(v), nil
	case float64:
		return int16(v), nil
	case int:
		return int16(v), nil
	case int8:
		return int16(v), nil
	case int16:
		return v, nil
	case int32:
		return int16(v), nil
	case int64:
		return int16(v), nil
	case uint:
		return int16(v), nil
	case uint8:
		return int16(v), nil
	case uint16:
		return int16(v), nil
	case uint32:
		return int16(v), nil
	case uint64:
		return int16(v), nil
	case bool:
		return int16(Is(v).Ok(1).No(0).Value().AsInt()), nil
	case string:
		if i, err := strconv.ParseInt(v, 0, 64); err != nil {
			return 0, err
		} else {
			return int16(i), nil
		}
	case []byte:
		if i, err := strconv.ParseInt(string(v), 0, 64); err != nil {
			return 0, err
		} else {
			return int16(i), nil
		}
	case reflect.Value:
		return Int16(v.Interface())
	}
	if r := reflect.ValueOf(in); r.Kind() == reflect.Ptr {
		if r.IsNil() {
			return 0, nil
		}
		return Int16(r.Elem().Interface())
	}
	return 0, _err
}
func Int32(in interface{}) (int32, error) {
	switch v := in.(type) {
	case nil:
		return 0, nil
	case float32:
		return int32(v), nil
	case float64:
		return int32(v), nil
	case int:
		return int32(v), nil
	case int8:
		return int32(v), nil
	case int16:
		return int32(v), nil
	case int32:
		return v, nil
	case int64:
		return int32(v), nil
	case uint:
		return int32(v), nil
	case uint8:
		return int32(v), nil
	case uint16:
		return int32(v), nil
	case uint32:
		return int32(v), nil
	case uint64:
		return int32(v), nil
	case bool:
		return int32(Is(v).Ok(1).No(0).Value().AsInt()), nil
	case string:
		if i, err := strconv.ParseInt(v, 10, 32); err != nil {
			return 0, err
		} else {
			return int32(i), err
		}
	case []byte:
		if i, err := strconv.ParseInt(string(v), 10, 32); err != nil {
			return 0, err
		} else {
			return int32(i), err
		}
	case reflect.Value:
		return Int32(v.Interface())
	}
	if r := reflect.ValueOf(in); r.Kind() == reflect.Ptr {
		if r.IsNil() {
			return 0, nil
		}

		return Int32(r.Elem().Interface())
	}
	return 0, _err
}
func Int64(in interface{}) (int64, error) {
	switch v := in.(type) {
	case nil:
		return 0, nil
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case uint:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case bool:
		return int64(Is(v).Ok(1).No(0).Value().AsInt()), nil
	case string:
		if i, err := strconv.ParseInt(v, 0, 64); err != nil {
			return 0, err
		} else {
			return i, nil
		}
	case []byte:
		if i, err := strconv.ParseInt(string(v), 0, 64); err != nil {
			return 0, err
		} else {
			return i, nil
		}
	case reflect.Value:
		return Int64(v.Interface())
	}
	if r := reflect.ValueOf(in); r.Kind() == reflect.Ptr {
		if r.IsNil() {
			return 0, nil
		}
		return Int64(r.Elem().Interface())
	}
	return 0, _err

}
func Str(in interface{}) (string, error) {
	switch v := in.(type) {
	case reflect.Value:
		return Str(v.Interface())
	case nil:
		return "", nil
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case int:
		return strconv.FormatInt(int64(v), 10), nil
	case int8:
		return strconv.FormatInt(int64(v), 10), nil
	case int16:
		return strconv.FormatInt(int64(v), 10), nil
	case int32:
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case uint:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint64:
		return strconv.FormatUint(v, 10), nil
	case string:
		return v, nil
	case bool:
		return strconv.FormatBool(v), nil
	case []rune:
		return string(v), nil
	case []byte:
		return *(*string)(unsafe.Pointer(&v)), nil
	case fmt.Stringer:
		return v.String(), nil
	case fmt.GoStringer:
		return v.GoString(), nil
	case io.Reader:
		if b, err := ioutil.ReadAll(v); err != nil {
			return "", err
		} else {
			return *(*string)(unsafe.Pointer(&b)), nil
		}
	case strconv.NumError:
		return v.Error(), v.Err
	}
	if _, ok := in.(error); ok {
		return in.(error).Error(), nil
	}
	if r := reflect.ValueOf(in); r.Kind() == reflect.Ptr {
		if r.IsNil() {
			return "", nil
		}
		return Str(r.Elem().Interface())
	}
	return fmt.Sprintf("%v", in), nil
}
func Float32(in interface{}) (float32, error) {
	switch v := in.(type) {
	case nil:
		return 0, nil
	case reflect.Value:
		return Float32(v.Interface())
	case float32:
		return v, nil
	case float64:
		return float32(v), nil
	case int:
		return float32(v), nil
	case int8:
		return float32(v), nil
	case int16:
		return float32(v), nil
	case int32:
		return float32(v), nil
	case int64:
		return float32(v), nil
	case uint:
		return float32(v), nil
	case uint8:
		return float32(v), nil
	case uint16:
		return float32(v), nil
	case uint32:
		return float32(v), nil
	case uint64:
		return float32(v), nil
	case bool:
		return float32(Is(v).Ok(1).No(0).Value().AsInt()), nil
	case string:
		if i, err := strconv.ParseFloat(v, 32); err != nil {
			return 0, err
		} else {
			return float32(i), nil
		}
	case []byte:
		if i, err := strconv.ParseFloat(string(v), 32); err != nil {
			return 0, err
		} else {
			return float32(i), nil
		}
	}
	if r := reflect.ValueOf(in); r.Kind() == reflect.Ptr {
		if r.IsNil() {
			return 0, nil
		}
		return Float32(r.Elem().Interface())
	}
	return 0, _err
}
func Float64(in interface{}) (float64, error) {
	switch v := in.(type) {
	case nil:
		return 0, nil
	case reflect.Value:
		return Float64(v.Interface())
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case bool:
		return float64(Is(v).Ok(1).No(0).Value().AsInt()), nil
	case string:
		return strconv.ParseFloat(v, 64)
	case []byte:
		return strconv.ParseFloat(string(v), 64)
	}
	if r := reflect.ValueOf(in); r.Kind() == reflect.Ptr {
		if r.IsNil() {
			return 0, nil
		}
		return Float64(r.Elem().Interface())
	}
	return 0, _err
}
func Duration(in interface{}) (time.Duration, error) {
	switch v := in.(type) {
	case nil:
		return 0, nil
	case reflect.Value:
		return Duration(v.Interface())
	case time.Duration:
		return v, nil
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
		return time.Duration(MustInt64(v)), nil
	case float32, float64:
		return time.Duration(MustFloat64(v)), nil
	case string:
		v = strings.ReplaceAll(v, "ns", "")
		v = strings.ReplaceAll(v, "ms", "")
		v = strings.ReplaceAll(v, "s", "")
		t, err := Int64(v)
		return time.Duration(t), err
	default:
		return 0, _err
	}
}
func Bool(in interface{}) (bool, error) {
	switch v := in.(type) {
	case nil:
		return false, nil
	case reflect.Value:
		return Bool(v.Interface())
	case int:
		return v > 0, nil
	case int8:
		return v > 0, nil
	case int16:
		return v > 0, nil
	case int32:
		return v > 0, nil
	case int64:
		return v > 0, nil
	case uint:
		return v > 0, nil
	case uint8:
		return v > 0, nil
	case uint16:
		return v > 0, nil
	case uint32:
		return v > 0, nil
	case uint64:
		return v > 0, nil
	case float32:
		return v > 0, nil
	case float64:
		return v > 0, nil
	case []byte:
		return strconv.ParseBool(MustStr(v))
	case string:
		return strconv.ParseBool(v)
	case bool:
		return v, nil
	}
	if r := reflect.ValueOf(in); r.Kind() == reflect.Ptr {
		if r.IsNil() {
			return false, nil
		}
		return Bool(r.Elem().Interface())
	}
	return false, _err
}
func Uint(in interface{}) (uint, error) {
	switch v := in.(type) {
	case reflect.Value:
		return Uint(v.Interface())
	case string:
		if i, err := strconv.ParseUint(v, 0, 0); err == nil {
			return uint(i), nil
		}
		return 0, uintErr
	case int:
		if v < 0 {
			return 0, uintErr
		}
		return uint(v), nil
	case int64:
		if v < 0 {
			return 0, uintErr
		}
		return uint(v), nil
	case int32:
		if v < 0 {
			return 0, uintErr
		}
		return uint(v), nil
	case int16:
		if v < 0 {
			return 0, uintErr
		}
		return uint(v), nil
	case int8:
		if v < 0 {
			return 0, uintErr
		}
		return uint(v), nil
	case float64:
		if v < 0 {
			return 0, uintErr
		}
		return uint(v), nil
	case float32:
		if v < 0 {
			return 0, uintErr
		}
		return uint(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case uint:
		return v, nil
	case uint64:
		return uint(v), nil
	case uint32:
		return uint(v), nil
	case uint16:
		return uint(v), nil
	case uint8:
		return uint(v), nil
	case nil:
		return 0, nil
	}
	if r := reflect.ValueOf(in); r.Kind() == reflect.Ptr {
		if r.IsNil() {
			return 0, nil
		}
		return Uint(r.Elem().Interface())
	}
	return 0, uintErr
}
func Uint64(in interface{}) (uint64, error) {
	switch v := in.(type) {
	case reflect.Value:
		return Uint64(v.Interface())
	case string:
		if i, err := strconv.ParseUint(v, 0, 64); err == nil {
			return i, nil
		}
		return 0, uintErr
	case int:
		if v < 0 {
			return 0, uintErr
		}
		return uint64(v), nil
	case int64:
		if v < 0 {
			return 0, uintErr
		}
		return uint64(v), nil
	case int32:
		if v < 0 {
			return 0, uintErr
		}
		return uint64(v), nil
	case int16:
		if v < 0 {
			return 0, uintErr
		}
		return uint64(v), nil
	case int8:
		if v < 0 {
			return 0, uintErr
		}
		return uint64(v), nil
	case float32:
		if v < 0 {
			return 0, uintErr
		}
		return uint64(v), nil
	case float64:
		if v < 0 {
			return 0, uintErr
		}
		return uint64(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case uint:
		return uint64(v), nil
	case uint64:
		return v, nil
	case uint32:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case nil:
		return 0, nil
	}
	if r := reflect.ValueOf(in); r.Kind() == reflect.Ptr {
		if r.IsNil() {
			return 0, nil
		}
		return Uint64(r.Elem().Interface())
	}
	return 0, uintErr
}
func Uint32(in interface{}) (uint32, error) {
	switch v := in.(type) {
	case reflect.Value:
		return Uint32(v.Interface())
	case string:
		if i, err := strconv.ParseUint(v, 0, 32); err == nil {
			return uint32(i), nil
		}
		return 0, uintErr
	case int:
		if v < 0 {
			return 0, uintErr
		}
		return uint32(v), nil
	case int64:
		if v < 0 {
			return 0, uintErr
		}
		return uint32(v), nil
	case int32:
		if v < 0 {
			return 0, uintErr
		}
		return uint32(v), nil
	case int16:
		if v < 0 {
			return 0, uintErr
		}
		return uint32(v), nil
	case int8:
		if v < 0 {
			return 0, uintErr
		}
		return uint32(v), nil
	case float64:
		if v < 0 {
			return 0, uintErr
		}
		return uint32(v), nil
	case float32:
		if v < 0 {
			return 0, uintErr
		}
		return uint32(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case uint:
		return uint32(v), nil
	case uint64:
		return uint32(v), nil
	case uint32:
		return v, nil
	case uint16:
		return uint32(v), nil
	case uint8:
		return uint32(v), nil
	case nil:
		return 0, nil
	}
	if r := reflect.ValueOf(in); r.Kind() == reflect.Ptr {
		if r.IsNil() {
			return 0, nil
		}
		return Uint32(r.Elem().Interface())
	}
	return 0, uintErr
}
func Uint16(in interface{}) (uint16, error) {
	switch v := in.(type) {
	case nil:
		return 0, nil
	case reflect.Value:
		return Uint16(v.Interface())
	case string:
		if i, err := strconv.ParseUint(v, 0, 16); err == nil {
			return uint16(i), nil
		}
		return 0, uintErr
	case int:
		if v < 0 {
			return 0, uintErr
		}
		return uint16(v), nil
	case int64:
		if v < 0 {
			return 0, uintErr
		}
		return uint16(v), nil
	case int32:
		if v < 0 {
			return 0, uintErr
		}
		return uint16(v), nil
	case int16:
		if v < 0 {
			return 0, uintErr
		}
		return uint16(v), nil
	case int8:
		if v < 0 {
			return 0, uintErr
		}
		return uint16(v), nil
	case float64:
		if v < 0 {
			return 0, uintErr
		}
		return uint16(v), nil
	case float32:
		if v < 0 {
			return 0, uintErr
		}
		return uint16(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case uint:
		return uint16(v), nil
	case uint64:
		return uint16(v), nil
	case uint32:
		return uint16(v), nil
	case uint16:
		return v, nil
	case uint8:
		return uint16(v), nil
	}
	if r := reflect.ValueOf(in); r.Kind() == reflect.Ptr {
		if r.IsNil() {
			return 0, nil
		}
		return Uint16(r.Elem().Interface())
	}
	return 0, uintErr
}
func Uint8(in interface{}) (uint8, error) {
	switch v := in.(type) {
	case reflect.Value:
		return Uint8(v.Interface())
	case string:
		i, err := strconv.ParseUint(v, 0, 8)
		if err == nil {
			return uint8(i), nil
		}
		return 0, uintErr
	case int:
		if v < 0 {
			return 0, uintErr
		}
		return uint8(v), nil
	case int64:
		if v < 0 {
			return 0, uintErr
		}
		return uint8(v), nil
	case int32:
		if v < 0 {
			return 0, uintErr
		}
		return uint8(v), nil
	case int16:
		if v < 0 {
			return 0, uintErr
		}
		return uint8(v), nil
	case int8:
		if v < 0 {
			return 0, uintErr
		}
		return uint8(v), nil
	case float64:
		if v < 0 {
			return 0, uintErr
		}
		return uint8(v), nil
	case float32:
		if v < 0 {
			return 0, uintErr
		}
		return uint8(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case uint:
		return uint8(v), nil
	case uint64:
		return uint8(v), nil
	case uint32:
		return uint8(v), nil
	case uint16:
		return uint8(v), nil
	case uint8:
		return v, nil
	case nil:
		return 0, nil
	}
	if r := reflect.ValueOf(in); r.Kind() == reflect.Ptr {
		if r.IsNil() {
			return 0, nil
		}
		return Uint8(r.Elem().Interface())
	}
	return 0, uintErr
}

func StringPtr(s string) *string     { return &s }
func BoolPtr(b bool) *bool           { return &b }
func IntPtr(i int) *int              { return &i }
func Int32Ptr(i int32) *int32        { return &i }
func Int64Ptr(i int64) *int64        { return &i }
func Float32Ptr(i float32) *float32  { return &i }
func Float64Ptr(i float64) *float64  { return &i }
func ByteSlicePtr(b []byte) *[]byte  { return &b }
func TimePtr(v time.Time) *time.Time { return &v }
