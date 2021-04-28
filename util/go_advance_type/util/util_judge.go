package util

import "reflect"

type (
	Judge struct {
		_Value interface{}
		flag   bool
	}
	Result struct {
		value interface{}
	}
)

func Is(flag bool) Judge {
	return Judge{flag: flag}
}
func (judge Judge) True(val interface{}) Judge {
	if judge.flag == true {
		return judge.set(val)
	}
	return judge
}
func (judge Judge) Yes(val interface{}) Judge {
	if judge.flag == true {
		return judge.set(val)
	}
	return judge
}
func (judge Judge) Ok(val interface{}) Judge {
	if judge.flag == true {
		return judge.set(val)
	}
	return judge
}
func (judge Judge) T(val interface{}) Judge {
	if judge.flag == true {
		return judge.set(val)
	}
	return judge
}
func (judge Judge) False(val interface{}) Judge {
	if judge.flag == false {
		return judge.set(val)
	}
	return judge
}
func (judge Judge) F(val interface{}) Judge {
	if judge.flag == false {
		return judge.set(val)
	}
	return judge
}

func (judge Judge) No(val interface{}) Judge {
	if judge.flag == false {
		return judge.set(val)
	}
	return judge
}
func (judge Judge) Value() Result {
	return Result{value: judge._Value}
}
func (judge Judge) set(val interface{}) Judge {
	if ref := reflect.ValueOf(val); ref.Kind() == reflect.Ptr {
		judge._Value = ref.Interface()
	} else {
		judge._Value = val
	}
	return judge
}
func (r Result) Interface() interface{} {
	return r.value
}
func (r Result) AsInt() int {
	if _, ok := r.value.(int); ok {
		return r.value.(int)
	} else {
		panic("miss match uint8")
	}
}

func (r Result) AsUint8() uint8 {
	if _, ok := r.value.(uint8); ok {
		return r.value.(uint8)
	} else {
		panic("miss match uint8")
	}
}

func (r Result) AsUint16() uint16 {
	if _, ok := r.value.(uint16); ok {
		return r.value.(uint16)
	} else {
		panic("miss match unit16")
	}
}
func (r Result) AsUint32() uint32 {
	if _, ok := r.value.(uint32); ok {
		return r.value.(uint32)
	} else {
		panic("miss match uint32")
	}
}
func (r Result) AsUint64() uint64 {
	if _, ok := r.value.(uint64); ok {
		return r.value.(uint64)
	} else {
		panic("miss match uint64")
	}
}
func (r Result) AsInt64() int64 {
	if _, ok := r.value.(int64); ok {
		return r.value.(int64)
	} else {
		panic("miss match int64")
	}
}
func (r Result) AsInt32() int32 {
	if _, ok := r.value.(int32); ok {
		return r.value.(int32)
	} else {
		panic("miss match int32")
	}
}
func (r Result) AsInt16() int16 {
	if _, ok := r.value.(int16); ok {
		return r.value.(int16)
	} else {
		panic("miss match int16")
	}
}
func (r Result) AsInt8() int8 {
	if _, ok := r.value.(int8); ok {
		return r.value.(int8)
	} else {
		panic("miss match int8")
	}
}
func (r Result) AsFloat32() float32 {
	if _, ok := r.value.(float32); ok {
		return r.value.(float32)
	} else {
		panic("miss match float32")
	}
}
func (r Result) AsFloat64() float64 {
	if _, ok := r.value.(float64); ok {
		return r.value.(float64)
	} else {
		panic("miss match float32")
	}
}
func (r Result) AsBytes() []byte {
	if _, ok := r.value.([]byte); ok {
		return r.value.([]byte)
	} else {
		panic("miss match []byte")
	}
}
func (r Result) AsString() string {
	if _, ok := r.value.(string); ok {
		return r.value.(string)
	} else {
		panic("miss match string")
	}
}
