/*
 * Collection utility of ApiDeleteDeviceEntitiesComponentRequest Struct
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

type ApiDeleteDeviceEntitiesComponentRequestStream []ApiDeleteDeviceEntitiesComponentRequest

func ApiDeleteDeviceEntitiesComponentRequestStreamOf(arg ...ApiDeleteDeviceEntitiesComponentRequest) ApiDeleteDeviceEntitiesComponentRequestStream {
	return arg
}
func ApiDeleteDeviceEntitiesComponentRequestStreamFrom(arg []ApiDeleteDeviceEntitiesComponentRequest) ApiDeleteDeviceEntitiesComponentRequestStream {
	return arg
}
func CreateApiDeleteDeviceEntitiesComponentRequestStream(arg ...ApiDeleteDeviceEntitiesComponentRequest) *ApiDeleteDeviceEntitiesComponentRequestStream {
	tmp := ApiDeleteDeviceEntitiesComponentRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiDeleteDeviceEntitiesComponentRequestStream(arg []ApiDeleteDeviceEntitiesComponentRequest) *ApiDeleteDeviceEntitiesComponentRequestStream {
	tmp := ApiDeleteDeviceEntitiesComponentRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Add(arg ApiDeleteDeviceEntitiesComponentRequest) *ApiDeleteDeviceEntitiesComponentRequestStream {
	return self.AddAll(arg)
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) AddAll(arg ...ApiDeleteDeviceEntitiesComponentRequest) *ApiDeleteDeviceEntitiesComponentRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) AddSafe(arg *ApiDeleteDeviceEntitiesComponentRequest) *ApiDeleteDeviceEntitiesComponentRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Aggregate(fn func(ApiDeleteDeviceEntitiesComponentRequest, ApiDeleteDeviceEntitiesComponentRequest) ApiDeleteDeviceEntitiesComponentRequest) *ApiDeleteDeviceEntitiesComponentRequestStream {
	result := ApiDeleteDeviceEntitiesComponentRequestStreamOf()
	self.ForEach(func(v ApiDeleteDeviceEntitiesComponentRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiDeleteDeviceEntitiesComponentRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) AllMatch(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) AnyMatch(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Clone() *ApiDeleteDeviceEntitiesComponentRequestStream {
	temp := make([]ApiDeleteDeviceEntitiesComponentRequest, self.Len())
	copy(temp, *self)
	return (*ApiDeleteDeviceEntitiesComponentRequestStream)(&temp)
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Copy() *ApiDeleteDeviceEntitiesComponentRequestStream {
	return self.Clone()
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Concat(arg []ApiDeleteDeviceEntitiesComponentRequest) *ApiDeleteDeviceEntitiesComponentRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Contains(arg ApiDeleteDeviceEntitiesComponentRequest) bool {
	return self.FindIndex(func(_arg ApiDeleteDeviceEntitiesComponentRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Clean() *ApiDeleteDeviceEntitiesComponentRequestStream {
	*self = ApiDeleteDeviceEntitiesComponentRequestStreamOf()
	return self
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Delete(index int) *ApiDeleteDeviceEntitiesComponentRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) DeleteRange(startIndex, endIndex int) *ApiDeleteDeviceEntitiesComponentRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Distinct() *ApiDeleteDeviceEntitiesComponentRequestStream {
	caches := map[string]bool{}
	result := ApiDeleteDeviceEntitiesComponentRequestStreamOf()
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
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Each(fn func(ApiDeleteDeviceEntitiesComponentRequest)) *ApiDeleteDeviceEntitiesComponentRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) EachRight(fn func(ApiDeleteDeviceEntitiesComponentRequest)) *ApiDeleteDeviceEntitiesComponentRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Equals(arr []ApiDeleteDeviceEntitiesComponentRequest) bool {
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
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Filter(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) bool) *ApiDeleteDeviceEntitiesComponentRequestStream {
	result := ApiDeleteDeviceEntitiesComponentRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) FilterSlim(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) bool) *ApiDeleteDeviceEntitiesComponentRequestStream {
	result := ApiDeleteDeviceEntitiesComponentRequestStreamOf()
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
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Find(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) bool) *ApiDeleteDeviceEntitiesComponentRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) FindOr(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) bool, or ApiDeleteDeviceEntitiesComponentRequest) ApiDeleteDeviceEntitiesComponentRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) FindIndex(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) bool) int {
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
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) First() *ApiDeleteDeviceEntitiesComponentRequest {
	return self.Get(0)
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) FirstOr(arg ApiDeleteDeviceEntitiesComponentRequest) ApiDeleteDeviceEntitiesComponentRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) ForEach(fn func(ApiDeleteDeviceEntitiesComponentRequest, int)) *ApiDeleteDeviceEntitiesComponentRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) ForEachRight(fn func(ApiDeleteDeviceEntitiesComponentRequest, int)) *ApiDeleteDeviceEntitiesComponentRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) GroupBy(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) string) map[string][]ApiDeleteDeviceEntitiesComponentRequest {
	m := map[string][]ApiDeleteDeviceEntitiesComponentRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) GroupByValues(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) string) [][]ApiDeleteDeviceEntitiesComponentRequest {
	var tmp [][]ApiDeleteDeviceEntitiesComponentRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) IndexOf(arg ApiDeleteDeviceEntitiesComponentRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Last() *ApiDeleteDeviceEntitiesComponentRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) LastOr(arg ApiDeleteDeviceEntitiesComponentRequest) ApiDeleteDeviceEntitiesComponentRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Limit(limit int) *ApiDeleteDeviceEntitiesComponentRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Map(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Map2Int(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Map2Int32(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Map2Int64(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Map2Float32(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Map2Float64(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Map2Bool(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Map2Bytes(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Map2String(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Max(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) float64) *ApiDeleteDeviceEntitiesComponentRequest {
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
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Min(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) float64) *ApiDeleteDeviceEntitiesComponentRequest {
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
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) NoneMatch(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Get(index int) *ApiDeleteDeviceEntitiesComponentRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) GetOr(index int, arg ApiDeleteDeviceEntitiesComponentRequest) ApiDeleteDeviceEntitiesComponentRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Peek(fn func(*ApiDeleteDeviceEntitiesComponentRequest, int)) *ApiDeleteDeviceEntitiesComponentRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Reduce(fn func(ApiDeleteDeviceEntitiesComponentRequest, ApiDeleteDeviceEntitiesComponentRequest, int) ApiDeleteDeviceEntitiesComponentRequest) *ApiDeleteDeviceEntitiesComponentRequestStream {
	return self.ReduceInit(fn, ApiDeleteDeviceEntitiesComponentRequest{})
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) ReduceInit(fn func(ApiDeleteDeviceEntitiesComponentRequest, ApiDeleteDeviceEntitiesComponentRequest, int) ApiDeleteDeviceEntitiesComponentRequest, initialValue ApiDeleteDeviceEntitiesComponentRequest) *ApiDeleteDeviceEntitiesComponentRequestStream {
	result := ApiDeleteDeviceEntitiesComponentRequestStreamOf()
	self.ForEach(func(v ApiDeleteDeviceEntitiesComponentRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) ReduceInterface(fn func(interface{}, ApiDeleteDeviceEntitiesComponentRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiDeleteDeviceEntitiesComponentRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) ReduceString(fn func(string, ApiDeleteDeviceEntitiesComponentRequest, int) string) []string {
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
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) ReduceInt(fn func(int, ApiDeleteDeviceEntitiesComponentRequest, int) int) []int {
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
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) ReduceInt32(fn func(int32, ApiDeleteDeviceEntitiesComponentRequest, int) int32) []int32 {
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
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) ReduceInt64(fn func(int64, ApiDeleteDeviceEntitiesComponentRequest, int) int64) []int64 {
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
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) ReduceFloat32(fn func(float32, ApiDeleteDeviceEntitiesComponentRequest, int) float32) []float32 {
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
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) ReduceFloat64(fn func(float64, ApiDeleteDeviceEntitiesComponentRequest, int) float64) []float64 {
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
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) ReduceBool(fn func(bool, ApiDeleteDeviceEntitiesComponentRequest, int) bool) []bool {
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
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Reverse() *ApiDeleteDeviceEntitiesComponentRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Replace(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) ApiDeleteDeviceEntitiesComponentRequest) *ApiDeleteDeviceEntitiesComponentRequestStream {
	return self.ForEach(func(v ApiDeleteDeviceEntitiesComponentRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Select(fn func(ApiDeleteDeviceEntitiesComponentRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Set(index int, val ApiDeleteDeviceEntitiesComponentRequest) *ApiDeleteDeviceEntitiesComponentRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Skip(skip int) *ApiDeleteDeviceEntitiesComponentRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) SkippingEach(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) int) *ApiDeleteDeviceEntitiesComponentRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Slice(startIndex, n int) *ApiDeleteDeviceEntitiesComponentRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiDeleteDeviceEntitiesComponentRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Sort(fn func(i, j int) bool) *ApiDeleteDeviceEntitiesComponentRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Tail() *ApiDeleteDeviceEntitiesComponentRequest {
	return self.Last()
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) TailOr(arg ApiDeleteDeviceEntitiesComponentRequest) ApiDeleteDeviceEntitiesComponentRequest {
	return self.LastOr(arg)
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) ToList() []ApiDeleteDeviceEntitiesComponentRequest {
	return self.Val()
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Unique() *ApiDeleteDeviceEntitiesComponentRequestStream {
	return self.Distinct()
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Val() []ApiDeleteDeviceEntitiesComponentRequest {
	if self == nil {
		return []ApiDeleteDeviceEntitiesComponentRequest{}
	}
	return *self.Copy()
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) While(fn func(ApiDeleteDeviceEntitiesComponentRequest, int) bool) *ApiDeleteDeviceEntitiesComponentRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) Where(fn func(ApiDeleteDeviceEntitiesComponentRequest) bool) *ApiDeleteDeviceEntitiesComponentRequestStream {
	result := ApiDeleteDeviceEntitiesComponentRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiDeleteDeviceEntitiesComponentRequestStream) WhereSlim(fn func(ApiDeleteDeviceEntitiesComponentRequest) bool) *ApiDeleteDeviceEntitiesComponentRequestStream {
	result := ApiDeleteDeviceEntitiesComponentRequestStreamOf()
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
