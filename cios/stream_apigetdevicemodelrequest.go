/*
 * Collection utility of ApiGetDeviceModelRequest Struct
 *
 * Generated by: Go Streamer (https://gitlab.tokyo.optim.co.jp/mrohchi/tool/go-struct-stream-generator)
 * Author https://gitlab.tokyo.optim.co.jp/kazuhiro.seida
 */

package cios

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

type ApiGetDeviceModelRequestStream []ApiGetDeviceModelRequest

func ApiGetDeviceModelRequestStreamOf(arg ...ApiGetDeviceModelRequest) ApiGetDeviceModelRequestStream {
	return arg
}
func ApiGetDeviceModelRequestStreamFrom(arg []ApiGetDeviceModelRequest) ApiGetDeviceModelRequestStream {
	return arg
}
func CreateApiGetDeviceModelRequestStream(arg ...ApiGetDeviceModelRequest) *ApiGetDeviceModelRequestStream {
	tmp := ApiGetDeviceModelRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetDeviceModelRequestStream(arg []ApiGetDeviceModelRequest) *ApiGetDeviceModelRequestStream {
	tmp := ApiGetDeviceModelRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetDeviceModelRequestStream) Add(arg ApiGetDeviceModelRequest) *ApiGetDeviceModelRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetDeviceModelRequestStream) AddAll(arg ...ApiGetDeviceModelRequest) *ApiGetDeviceModelRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetDeviceModelRequestStream) AddSafe(arg *ApiGetDeviceModelRequest) *ApiGetDeviceModelRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetDeviceModelRequestStream) Aggregate(fn func(ApiGetDeviceModelRequest, ApiGetDeviceModelRequest) ApiGetDeviceModelRequest) *ApiGetDeviceModelRequestStream {
	result := ApiGetDeviceModelRequestStreamOf()
	self.ForEach(func(v ApiGetDeviceModelRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetDeviceModelRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDeviceModelRequestStream) AllMatch(fn func(ApiGetDeviceModelRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetDeviceModelRequestStream) AnyMatch(fn func(ApiGetDeviceModelRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetDeviceModelRequestStream) Clone() *ApiGetDeviceModelRequestStream {
	temp := make([]ApiGetDeviceModelRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetDeviceModelRequestStream)(&temp)
}
func (self *ApiGetDeviceModelRequestStream) Copy() *ApiGetDeviceModelRequestStream {
	return self.Clone()
}
func (self *ApiGetDeviceModelRequestStream) Concat(arg []ApiGetDeviceModelRequest) *ApiGetDeviceModelRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetDeviceModelRequestStream) Contains(arg ApiGetDeviceModelRequest) bool {
	return self.FindIndex(func(_arg ApiGetDeviceModelRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetDeviceModelRequestStream) Clean() *ApiGetDeviceModelRequestStream {
	*self = ApiGetDeviceModelRequestStreamOf()
	return self
}
func (self *ApiGetDeviceModelRequestStream) Delete(index int) *ApiGetDeviceModelRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetDeviceModelRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetDeviceModelRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetDeviceModelRequestStream) Distinct() *ApiGetDeviceModelRequestStream {
	caches := map[string]bool{}
	result := ApiGetDeviceModelRequestStreamOf()
	for _, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if !f {
				result = append(result, v)
			}
		} else if caches[key] = true; !f {
			result = append(result, v)
		}

	}
	*self = result
	return self
}
func (self *ApiGetDeviceModelRequestStream) Each(fn func(ApiGetDeviceModelRequest)) *ApiGetDeviceModelRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetDeviceModelRequestStream) EachRight(fn func(ApiGetDeviceModelRequest)) *ApiGetDeviceModelRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetDeviceModelRequestStream) Equals(arr []ApiGetDeviceModelRequest) bool {
	if (*self == nil) != (arr == nil) || len(*self) != len(arr) {
		return false
	}
	for i := range *self {
		if !reflect.DeepEqual((*self)[i], arr[i]) {
			return false
		}
	}
	return true
}
func (self *ApiGetDeviceModelRequestStream) Filter(fn func(ApiGetDeviceModelRequest, int) bool) *ApiGetDeviceModelRequestStream {
	result := ApiGetDeviceModelRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDeviceModelRequestStream) FilterSlim(fn func(ApiGetDeviceModelRequest, int) bool) *ApiGetDeviceModelRequestStream {
	result := ApiGetDeviceModelRequestStreamOf()
	caches := map[string]bool{}
	for i, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if f {
				result.Add(v)
			}
		} else if caches[key] = fn(v, i); caches[key] {
			result.Add(v)

		}
	}
	*self = result
	return self
}
func (self *ApiGetDeviceModelRequestStream) Find(fn func(ApiGetDeviceModelRequest, int) bool) *ApiGetDeviceModelRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetDeviceModelRequestStream) FindOr(fn func(ApiGetDeviceModelRequest, int) bool, or ApiGetDeviceModelRequest) ApiGetDeviceModelRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetDeviceModelRequestStream) FindIndex(fn func(ApiGetDeviceModelRequest, int) bool) int {
	if self == nil {
		return -1
	}
	for i, v := range *self {
		if fn(v, i) {
			return i
		}
	}
	return -1
}
func (self *ApiGetDeviceModelRequestStream) First() *ApiGetDeviceModelRequest {
	return self.Get(0)
}
func (self *ApiGetDeviceModelRequestStream) FirstOr(arg ApiGetDeviceModelRequest) ApiGetDeviceModelRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDeviceModelRequestStream) ForEach(fn func(ApiGetDeviceModelRequest, int)) *ApiGetDeviceModelRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetDeviceModelRequestStream) ForEachRight(fn func(ApiGetDeviceModelRequest, int)) *ApiGetDeviceModelRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetDeviceModelRequestStream) GroupBy(fn func(ApiGetDeviceModelRequest, int) string) map[string][]ApiGetDeviceModelRequest {
	m := map[string][]ApiGetDeviceModelRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetDeviceModelRequestStream) GroupByValues(fn func(ApiGetDeviceModelRequest, int) string) [][]ApiGetDeviceModelRequest {
	var tmp [][]ApiGetDeviceModelRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetDeviceModelRequestStream) IndexOf(arg ApiGetDeviceModelRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetDeviceModelRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetDeviceModelRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetDeviceModelRequestStream) Last() *ApiGetDeviceModelRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetDeviceModelRequestStream) LastOr(arg ApiGetDeviceModelRequest) ApiGetDeviceModelRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDeviceModelRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetDeviceModelRequestStream) Limit(limit int) *ApiGetDeviceModelRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetDeviceModelRequestStream) Map(fn func(ApiGetDeviceModelRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceModelRequestStream) Map2Int(fn func(ApiGetDeviceModelRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceModelRequestStream) Map2Int32(fn func(ApiGetDeviceModelRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceModelRequestStream) Map2Int64(fn func(ApiGetDeviceModelRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceModelRequestStream) Map2Float32(fn func(ApiGetDeviceModelRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceModelRequestStream) Map2Float64(fn func(ApiGetDeviceModelRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceModelRequestStream) Map2Bool(fn func(ApiGetDeviceModelRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceModelRequestStream) Map2Bytes(fn func(ApiGetDeviceModelRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceModelRequestStream) Map2String(fn func(ApiGetDeviceModelRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceModelRequestStream) Max(fn func(ApiGetDeviceModelRequest, int) float64) *ApiGetDeviceModelRequest {
	f := self.Get(0)
	if f == nil {
		return nil
	}
	m := fn(*f, 0)
	index := 0
	for i := 1; i < self.Len(); i++ {
		v := fn(*self.Get(i), i)
		m = math.Max(m, v)
		if m == v {
			index = i
		}
	}
	return self.Get(index)
}
func (self *ApiGetDeviceModelRequestStream) Min(fn func(ApiGetDeviceModelRequest, int) float64) *ApiGetDeviceModelRequest {
	f := self.Get(0)
	if f == nil {
		return nil
	}
	m := fn(*f, 0)
	index := 0
	for i := 1; i < self.Len(); i++ {
		v := fn(*self.Get(i), i)
		m = math.Min(m, v)
		if m == v {
			index = i
		}
	}
	return self.Get(index)
}
func (self *ApiGetDeviceModelRequestStream) NoneMatch(fn func(ApiGetDeviceModelRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetDeviceModelRequestStream) Get(index int) *ApiGetDeviceModelRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetDeviceModelRequestStream) GetOr(index int, arg ApiGetDeviceModelRequest) ApiGetDeviceModelRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDeviceModelRequestStream) Peek(fn func(*ApiGetDeviceModelRequest, int)) *ApiGetDeviceModelRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetDeviceModelRequestStream) Reduce(fn func(ApiGetDeviceModelRequest, ApiGetDeviceModelRequest, int) ApiGetDeviceModelRequest) *ApiGetDeviceModelRequestStream {
	return self.ReduceInit(fn, ApiGetDeviceModelRequest{})
}
func (self *ApiGetDeviceModelRequestStream) ReduceInit(fn func(ApiGetDeviceModelRequest, ApiGetDeviceModelRequest, int) ApiGetDeviceModelRequest, initialValue ApiGetDeviceModelRequest) *ApiGetDeviceModelRequestStream {
	result := ApiGetDeviceModelRequestStreamOf()
	self.ForEach(func(v ApiGetDeviceModelRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDeviceModelRequestStream) ReduceInterface(fn func(interface{}, ApiGetDeviceModelRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetDeviceModelRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetDeviceModelRequestStream) ReduceString(fn func(string, ApiGetDeviceModelRequest, int) string) []string {
	result := []string{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn("", v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetDeviceModelRequestStream) ReduceInt(fn func(int, ApiGetDeviceModelRequest, int) int) []int {
	result := []int{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetDeviceModelRequestStream) ReduceInt32(fn func(int32, ApiGetDeviceModelRequest, int) int32) []int32 {
	result := []int32{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetDeviceModelRequestStream) ReduceInt64(fn func(int64, ApiGetDeviceModelRequest, int) int64) []int64 {
	result := []int64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetDeviceModelRequestStream) ReduceFloat32(fn func(float32, ApiGetDeviceModelRequest, int) float32) []float32 {
	result := []float32{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetDeviceModelRequestStream) ReduceFloat64(fn func(float64, ApiGetDeviceModelRequest, int) float64) []float64 {
	result := []float64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetDeviceModelRequestStream) ReduceBool(fn func(bool, ApiGetDeviceModelRequest, int) bool) []bool {
	result := []bool{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(false, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetDeviceModelRequestStream) Reverse() *ApiGetDeviceModelRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetDeviceModelRequestStream) Replace(fn func(ApiGetDeviceModelRequest, int) ApiGetDeviceModelRequest) *ApiGetDeviceModelRequestStream {
	return self.ForEach(func(v ApiGetDeviceModelRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetDeviceModelRequestStream) Select(fn func(ApiGetDeviceModelRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetDeviceModelRequestStream) Set(index int, val ApiGetDeviceModelRequest) *ApiGetDeviceModelRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetDeviceModelRequestStream) Skip(skip int) *ApiGetDeviceModelRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetDeviceModelRequestStream) SkippingEach(fn func(ApiGetDeviceModelRequest, int) int) *ApiGetDeviceModelRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetDeviceModelRequestStream) Slice(startIndex, n int) *ApiGetDeviceModelRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetDeviceModelRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetDeviceModelRequestStream) Sort(fn func(i, j int) bool) *ApiGetDeviceModelRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetDeviceModelRequestStream) Tail() *ApiGetDeviceModelRequest {
	return self.Last()
}
func (self *ApiGetDeviceModelRequestStream) TailOr(arg ApiGetDeviceModelRequest) ApiGetDeviceModelRequest {
	return self.LastOr(arg)
}
func (self *ApiGetDeviceModelRequestStream) ToList() []ApiGetDeviceModelRequest {
	return self.Val()
}
func (self *ApiGetDeviceModelRequestStream) Unique() *ApiGetDeviceModelRequestStream {
	return self.Distinct()
}
func (self *ApiGetDeviceModelRequestStream) Val() []ApiGetDeviceModelRequest {
	if self == nil {
		return []ApiGetDeviceModelRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetDeviceModelRequestStream) While(fn func(ApiGetDeviceModelRequest, int) bool) *ApiGetDeviceModelRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetDeviceModelRequestStream) Where(fn func(ApiGetDeviceModelRequest) bool) *ApiGetDeviceModelRequestStream {
	result := ApiGetDeviceModelRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDeviceModelRequestStream) WhereSlim(fn func(ApiGetDeviceModelRequest) bool) *ApiGetDeviceModelRequestStream {
	result := ApiGetDeviceModelRequestStreamOf()
	caches := map[string]bool{}
	for _, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if f {
				result.Add(v)
			}
		} else if caches[key] = fn(v); caches[key] {
			result.Add(v)
		}
	}
	*self = result
	return self
}
