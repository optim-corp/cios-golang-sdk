/*
 * Collection utility of ApiUpdateDeviceModelRequest Struct
 *
 * Generated by: Go Streamer
 */

package cios

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

type ApiUpdateDeviceModelRequestStream []ApiUpdateDeviceModelRequest

func ApiUpdateDeviceModelRequestStreamOf(arg ...ApiUpdateDeviceModelRequest) ApiUpdateDeviceModelRequestStream {
	return arg
}
func ApiUpdateDeviceModelRequestStreamFrom(arg []ApiUpdateDeviceModelRequest) ApiUpdateDeviceModelRequestStream {
	return arg
}
func CreateApiUpdateDeviceModelRequestStream(arg ...ApiUpdateDeviceModelRequest) *ApiUpdateDeviceModelRequestStream {
	tmp := ApiUpdateDeviceModelRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiUpdateDeviceModelRequestStream(arg []ApiUpdateDeviceModelRequest) *ApiUpdateDeviceModelRequestStream {
	tmp := ApiUpdateDeviceModelRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiUpdateDeviceModelRequestStream) Add(arg ApiUpdateDeviceModelRequest) *ApiUpdateDeviceModelRequestStream {
	return self.AddAll(arg)
}
func (self *ApiUpdateDeviceModelRequestStream) AddAll(arg ...ApiUpdateDeviceModelRequest) *ApiUpdateDeviceModelRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiUpdateDeviceModelRequestStream) AddSafe(arg *ApiUpdateDeviceModelRequest) *ApiUpdateDeviceModelRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiUpdateDeviceModelRequestStream) Aggregate(fn func(ApiUpdateDeviceModelRequest, ApiUpdateDeviceModelRequest) ApiUpdateDeviceModelRequest) *ApiUpdateDeviceModelRequestStream {
	result := ApiUpdateDeviceModelRequestStreamOf()
	self.ForEach(func(v ApiUpdateDeviceModelRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiUpdateDeviceModelRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiUpdateDeviceModelRequestStream) AllMatch(fn func(ApiUpdateDeviceModelRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiUpdateDeviceModelRequestStream) AnyMatch(fn func(ApiUpdateDeviceModelRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiUpdateDeviceModelRequestStream) Clone() *ApiUpdateDeviceModelRequestStream {
	temp := make([]ApiUpdateDeviceModelRequest, self.Len())
	copy(temp, *self)
	return (*ApiUpdateDeviceModelRequestStream)(&temp)
}
func (self *ApiUpdateDeviceModelRequestStream) Copy() *ApiUpdateDeviceModelRequestStream {
	return self.Clone()
}
func (self *ApiUpdateDeviceModelRequestStream) Concat(arg []ApiUpdateDeviceModelRequest) *ApiUpdateDeviceModelRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiUpdateDeviceModelRequestStream) Contains(arg ApiUpdateDeviceModelRequest) bool {
	return self.FindIndex(func(_arg ApiUpdateDeviceModelRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiUpdateDeviceModelRequestStream) Clean() *ApiUpdateDeviceModelRequestStream {
	*self = ApiUpdateDeviceModelRequestStreamOf()
	return self
}
func (self *ApiUpdateDeviceModelRequestStream) Delete(index int) *ApiUpdateDeviceModelRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiUpdateDeviceModelRequestStream) DeleteRange(startIndex, endIndex int) *ApiUpdateDeviceModelRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiUpdateDeviceModelRequestStream) Distinct() *ApiUpdateDeviceModelRequestStream {
	caches := map[string]bool{}
	result := ApiUpdateDeviceModelRequestStreamOf()
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
func (self *ApiUpdateDeviceModelRequestStream) Each(fn func(ApiUpdateDeviceModelRequest)) *ApiUpdateDeviceModelRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiUpdateDeviceModelRequestStream) EachRight(fn func(ApiUpdateDeviceModelRequest)) *ApiUpdateDeviceModelRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiUpdateDeviceModelRequestStream) Equals(arr []ApiUpdateDeviceModelRequest) bool {
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
func (self *ApiUpdateDeviceModelRequestStream) Filter(fn func(ApiUpdateDeviceModelRequest, int) bool) *ApiUpdateDeviceModelRequestStream {
	result := ApiUpdateDeviceModelRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiUpdateDeviceModelRequestStream) FilterSlim(fn func(ApiUpdateDeviceModelRequest, int) bool) *ApiUpdateDeviceModelRequestStream {
	result := ApiUpdateDeviceModelRequestStreamOf()
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
func (self *ApiUpdateDeviceModelRequestStream) Find(fn func(ApiUpdateDeviceModelRequest, int) bool) *ApiUpdateDeviceModelRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiUpdateDeviceModelRequestStream) FindOr(fn func(ApiUpdateDeviceModelRequest, int) bool, or ApiUpdateDeviceModelRequest) ApiUpdateDeviceModelRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiUpdateDeviceModelRequestStream) FindIndex(fn func(ApiUpdateDeviceModelRequest, int) bool) int {
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
func (self *ApiUpdateDeviceModelRequestStream) First() *ApiUpdateDeviceModelRequest {
	return self.Get(0)
}
func (self *ApiUpdateDeviceModelRequestStream) FirstOr(arg ApiUpdateDeviceModelRequest) ApiUpdateDeviceModelRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiUpdateDeviceModelRequestStream) ForEach(fn func(ApiUpdateDeviceModelRequest, int)) *ApiUpdateDeviceModelRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiUpdateDeviceModelRequestStream) ForEachRight(fn func(ApiUpdateDeviceModelRequest, int)) *ApiUpdateDeviceModelRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiUpdateDeviceModelRequestStream) GroupBy(fn func(ApiUpdateDeviceModelRequest, int) string) map[string][]ApiUpdateDeviceModelRequest {
	m := map[string][]ApiUpdateDeviceModelRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiUpdateDeviceModelRequestStream) GroupByValues(fn func(ApiUpdateDeviceModelRequest, int) string) [][]ApiUpdateDeviceModelRequest {
	var tmp [][]ApiUpdateDeviceModelRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiUpdateDeviceModelRequestStream) IndexOf(arg ApiUpdateDeviceModelRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiUpdateDeviceModelRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiUpdateDeviceModelRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiUpdateDeviceModelRequestStream) Last() *ApiUpdateDeviceModelRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiUpdateDeviceModelRequestStream) LastOr(arg ApiUpdateDeviceModelRequest) ApiUpdateDeviceModelRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiUpdateDeviceModelRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiUpdateDeviceModelRequestStream) Limit(limit int) *ApiUpdateDeviceModelRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiUpdateDeviceModelRequestStream) Map(fn func(ApiUpdateDeviceModelRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateDeviceModelRequestStream) Map2Int(fn func(ApiUpdateDeviceModelRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateDeviceModelRequestStream) Map2Int32(fn func(ApiUpdateDeviceModelRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateDeviceModelRequestStream) Map2Int64(fn func(ApiUpdateDeviceModelRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateDeviceModelRequestStream) Map2Float32(fn func(ApiUpdateDeviceModelRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateDeviceModelRequestStream) Map2Float64(fn func(ApiUpdateDeviceModelRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateDeviceModelRequestStream) Map2Bool(fn func(ApiUpdateDeviceModelRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateDeviceModelRequestStream) Map2Bytes(fn func(ApiUpdateDeviceModelRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateDeviceModelRequestStream) Map2String(fn func(ApiUpdateDeviceModelRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateDeviceModelRequestStream) Max(fn func(ApiUpdateDeviceModelRequest, int) float64) *ApiUpdateDeviceModelRequest {
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
func (self *ApiUpdateDeviceModelRequestStream) Min(fn func(ApiUpdateDeviceModelRequest, int) float64) *ApiUpdateDeviceModelRequest {
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
func (self *ApiUpdateDeviceModelRequestStream) NoneMatch(fn func(ApiUpdateDeviceModelRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiUpdateDeviceModelRequestStream) Get(index int) *ApiUpdateDeviceModelRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiUpdateDeviceModelRequestStream) GetOr(index int, arg ApiUpdateDeviceModelRequest) ApiUpdateDeviceModelRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiUpdateDeviceModelRequestStream) Peek(fn func(*ApiUpdateDeviceModelRequest, int)) *ApiUpdateDeviceModelRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiUpdateDeviceModelRequestStream) Reduce(fn func(ApiUpdateDeviceModelRequest, ApiUpdateDeviceModelRequest, int) ApiUpdateDeviceModelRequest) *ApiUpdateDeviceModelRequestStream {
	return self.ReduceInit(fn, ApiUpdateDeviceModelRequest{})
}
func (self *ApiUpdateDeviceModelRequestStream) ReduceInit(fn func(ApiUpdateDeviceModelRequest, ApiUpdateDeviceModelRequest, int) ApiUpdateDeviceModelRequest, initialValue ApiUpdateDeviceModelRequest) *ApiUpdateDeviceModelRequestStream {
	result := ApiUpdateDeviceModelRequestStreamOf()
	self.ForEach(func(v ApiUpdateDeviceModelRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiUpdateDeviceModelRequestStream) ReduceInterface(fn func(interface{}, ApiUpdateDeviceModelRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiUpdateDeviceModelRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiUpdateDeviceModelRequestStream) ReduceString(fn func(string, ApiUpdateDeviceModelRequest, int) string) []string {
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
func (self *ApiUpdateDeviceModelRequestStream) ReduceInt(fn func(int, ApiUpdateDeviceModelRequest, int) int) []int {
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
func (self *ApiUpdateDeviceModelRequestStream) ReduceInt32(fn func(int32, ApiUpdateDeviceModelRequest, int) int32) []int32 {
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
func (self *ApiUpdateDeviceModelRequestStream) ReduceInt64(fn func(int64, ApiUpdateDeviceModelRequest, int) int64) []int64 {
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
func (self *ApiUpdateDeviceModelRequestStream) ReduceFloat32(fn func(float32, ApiUpdateDeviceModelRequest, int) float32) []float32 {
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
func (self *ApiUpdateDeviceModelRequestStream) ReduceFloat64(fn func(float64, ApiUpdateDeviceModelRequest, int) float64) []float64 {
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
func (self *ApiUpdateDeviceModelRequestStream) ReduceBool(fn func(bool, ApiUpdateDeviceModelRequest, int) bool) []bool {
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
func (self *ApiUpdateDeviceModelRequestStream) Reverse() *ApiUpdateDeviceModelRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiUpdateDeviceModelRequestStream) Replace(fn func(ApiUpdateDeviceModelRequest, int) ApiUpdateDeviceModelRequest) *ApiUpdateDeviceModelRequestStream {
	return self.ForEach(func(v ApiUpdateDeviceModelRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiUpdateDeviceModelRequestStream) Select(fn func(ApiUpdateDeviceModelRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiUpdateDeviceModelRequestStream) Set(index int, val ApiUpdateDeviceModelRequest) *ApiUpdateDeviceModelRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiUpdateDeviceModelRequestStream) Skip(skip int) *ApiUpdateDeviceModelRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiUpdateDeviceModelRequestStream) SkippingEach(fn func(ApiUpdateDeviceModelRequest, int) int) *ApiUpdateDeviceModelRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiUpdateDeviceModelRequestStream) Slice(startIndex, n int) *ApiUpdateDeviceModelRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiUpdateDeviceModelRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiUpdateDeviceModelRequestStream) Sort(fn func(i, j int) bool) *ApiUpdateDeviceModelRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiUpdateDeviceModelRequestStream) Tail() *ApiUpdateDeviceModelRequest {
	return self.Last()
}
func (self *ApiUpdateDeviceModelRequestStream) TailOr(arg ApiUpdateDeviceModelRequest) ApiUpdateDeviceModelRequest {
	return self.LastOr(arg)
}
func (self *ApiUpdateDeviceModelRequestStream) ToList() []ApiUpdateDeviceModelRequest {
	return self.Val()
}
func (self *ApiUpdateDeviceModelRequestStream) Unique() *ApiUpdateDeviceModelRequestStream {
	return self.Distinct()
}
func (self *ApiUpdateDeviceModelRequestStream) Val() []ApiUpdateDeviceModelRequest {
	if self == nil {
		return []ApiUpdateDeviceModelRequest{}
	}
	return *self.Copy()
}
func (self *ApiUpdateDeviceModelRequestStream) While(fn func(ApiUpdateDeviceModelRequest, int) bool) *ApiUpdateDeviceModelRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiUpdateDeviceModelRequestStream) Where(fn func(ApiUpdateDeviceModelRequest) bool) *ApiUpdateDeviceModelRequestStream {
	result := ApiUpdateDeviceModelRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiUpdateDeviceModelRequestStream) WhereSlim(fn func(ApiUpdateDeviceModelRequest) bool) *ApiUpdateDeviceModelRequestStream {
	result := ApiUpdateDeviceModelRequestStreamOf()
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
