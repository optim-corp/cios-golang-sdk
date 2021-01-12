/*
 * Collection utility of ApiGetDeviceEntitiesLifecycleRequest Struct
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

type ApiGetDeviceEntitiesLifecycleRequestStream []ApiGetDeviceEntitiesLifecycleRequest

func ApiGetDeviceEntitiesLifecycleRequestStreamOf(arg ...ApiGetDeviceEntitiesLifecycleRequest) ApiGetDeviceEntitiesLifecycleRequestStream {
	return arg
}
func ApiGetDeviceEntitiesLifecycleRequestStreamFrom(arg []ApiGetDeviceEntitiesLifecycleRequest) ApiGetDeviceEntitiesLifecycleRequestStream {
	return arg
}
func CreateApiGetDeviceEntitiesLifecycleRequestStream(arg ...ApiGetDeviceEntitiesLifecycleRequest) *ApiGetDeviceEntitiesLifecycleRequestStream {
	tmp := ApiGetDeviceEntitiesLifecycleRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetDeviceEntitiesLifecycleRequestStream(arg []ApiGetDeviceEntitiesLifecycleRequest) *ApiGetDeviceEntitiesLifecycleRequestStream {
	tmp := ApiGetDeviceEntitiesLifecycleRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Add(arg ApiGetDeviceEntitiesLifecycleRequest) *ApiGetDeviceEntitiesLifecycleRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) AddAll(arg ...ApiGetDeviceEntitiesLifecycleRequest) *ApiGetDeviceEntitiesLifecycleRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) AddSafe(arg *ApiGetDeviceEntitiesLifecycleRequest) *ApiGetDeviceEntitiesLifecycleRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Aggregate(fn func(ApiGetDeviceEntitiesLifecycleRequest, ApiGetDeviceEntitiesLifecycleRequest) ApiGetDeviceEntitiesLifecycleRequest) *ApiGetDeviceEntitiesLifecycleRequestStream {
	result := ApiGetDeviceEntitiesLifecycleRequestStreamOf()
	self.ForEach(func(v ApiGetDeviceEntitiesLifecycleRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetDeviceEntitiesLifecycleRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) AllMatch(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) AnyMatch(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Clone() *ApiGetDeviceEntitiesLifecycleRequestStream {
	temp := make([]ApiGetDeviceEntitiesLifecycleRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetDeviceEntitiesLifecycleRequestStream)(&temp)
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Copy() *ApiGetDeviceEntitiesLifecycleRequestStream {
	return self.Clone()
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Concat(arg []ApiGetDeviceEntitiesLifecycleRequest) *ApiGetDeviceEntitiesLifecycleRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Contains(arg ApiGetDeviceEntitiesLifecycleRequest) bool {
	return self.FindIndex(func(_arg ApiGetDeviceEntitiesLifecycleRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Clean() *ApiGetDeviceEntitiesLifecycleRequestStream {
	*self = ApiGetDeviceEntitiesLifecycleRequestStreamOf()
	return self
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Delete(index int) *ApiGetDeviceEntitiesLifecycleRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetDeviceEntitiesLifecycleRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Distinct() *ApiGetDeviceEntitiesLifecycleRequestStream {
	caches := map[string]bool{}
	result := ApiGetDeviceEntitiesLifecycleRequestStreamOf()
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
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Each(fn func(ApiGetDeviceEntitiesLifecycleRequest)) *ApiGetDeviceEntitiesLifecycleRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) EachRight(fn func(ApiGetDeviceEntitiesLifecycleRequest)) *ApiGetDeviceEntitiesLifecycleRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Equals(arr []ApiGetDeviceEntitiesLifecycleRequest) bool {
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
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Filter(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) bool) *ApiGetDeviceEntitiesLifecycleRequestStream {
	result := ApiGetDeviceEntitiesLifecycleRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) FilterSlim(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) bool) *ApiGetDeviceEntitiesLifecycleRequestStream {
	result := ApiGetDeviceEntitiesLifecycleRequestStreamOf()
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
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Find(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) bool) *ApiGetDeviceEntitiesLifecycleRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) FindOr(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) bool, or ApiGetDeviceEntitiesLifecycleRequest) ApiGetDeviceEntitiesLifecycleRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) FindIndex(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) bool) int {
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
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) First() *ApiGetDeviceEntitiesLifecycleRequest {
	return self.Get(0)
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) FirstOr(arg ApiGetDeviceEntitiesLifecycleRequest) ApiGetDeviceEntitiesLifecycleRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) ForEach(fn func(ApiGetDeviceEntitiesLifecycleRequest, int)) *ApiGetDeviceEntitiesLifecycleRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) ForEachRight(fn func(ApiGetDeviceEntitiesLifecycleRequest, int)) *ApiGetDeviceEntitiesLifecycleRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) GroupBy(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) string) map[string][]ApiGetDeviceEntitiesLifecycleRequest {
	m := map[string][]ApiGetDeviceEntitiesLifecycleRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) GroupByValues(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) string) [][]ApiGetDeviceEntitiesLifecycleRequest {
	var tmp [][]ApiGetDeviceEntitiesLifecycleRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) IndexOf(arg ApiGetDeviceEntitiesLifecycleRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Last() *ApiGetDeviceEntitiesLifecycleRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) LastOr(arg ApiGetDeviceEntitiesLifecycleRequest) ApiGetDeviceEntitiesLifecycleRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Limit(limit int) *ApiGetDeviceEntitiesLifecycleRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Map(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Map2Int(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Map2Int32(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Map2Int64(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Map2Float32(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Map2Float64(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Map2Bool(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Map2Bytes(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Map2String(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Max(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) float64) *ApiGetDeviceEntitiesLifecycleRequest {
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
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Min(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) float64) *ApiGetDeviceEntitiesLifecycleRequest {
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
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) NoneMatch(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Get(index int) *ApiGetDeviceEntitiesLifecycleRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) GetOr(index int, arg ApiGetDeviceEntitiesLifecycleRequest) ApiGetDeviceEntitiesLifecycleRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Peek(fn func(*ApiGetDeviceEntitiesLifecycleRequest, int)) *ApiGetDeviceEntitiesLifecycleRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Reduce(fn func(ApiGetDeviceEntitiesLifecycleRequest, ApiGetDeviceEntitiesLifecycleRequest, int) ApiGetDeviceEntitiesLifecycleRequest) *ApiGetDeviceEntitiesLifecycleRequestStream {
	return self.ReduceInit(fn, ApiGetDeviceEntitiesLifecycleRequest{})
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) ReduceInit(fn func(ApiGetDeviceEntitiesLifecycleRequest, ApiGetDeviceEntitiesLifecycleRequest, int) ApiGetDeviceEntitiesLifecycleRequest, initialValue ApiGetDeviceEntitiesLifecycleRequest) *ApiGetDeviceEntitiesLifecycleRequestStream {
	result := ApiGetDeviceEntitiesLifecycleRequestStreamOf()
	self.ForEach(func(v ApiGetDeviceEntitiesLifecycleRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) ReduceInterface(fn func(interface{}, ApiGetDeviceEntitiesLifecycleRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetDeviceEntitiesLifecycleRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) ReduceString(fn func(string, ApiGetDeviceEntitiesLifecycleRequest, int) string) []string {
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
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) ReduceInt(fn func(int, ApiGetDeviceEntitiesLifecycleRequest, int) int) []int {
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
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) ReduceInt32(fn func(int32, ApiGetDeviceEntitiesLifecycleRequest, int) int32) []int32 {
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
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) ReduceInt64(fn func(int64, ApiGetDeviceEntitiesLifecycleRequest, int) int64) []int64 {
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
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) ReduceFloat32(fn func(float32, ApiGetDeviceEntitiesLifecycleRequest, int) float32) []float32 {
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
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) ReduceFloat64(fn func(float64, ApiGetDeviceEntitiesLifecycleRequest, int) float64) []float64 {
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
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) ReduceBool(fn func(bool, ApiGetDeviceEntitiesLifecycleRequest, int) bool) []bool {
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
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Reverse() *ApiGetDeviceEntitiesLifecycleRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Replace(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) ApiGetDeviceEntitiesLifecycleRequest) *ApiGetDeviceEntitiesLifecycleRequestStream {
	return self.ForEach(func(v ApiGetDeviceEntitiesLifecycleRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Select(fn func(ApiGetDeviceEntitiesLifecycleRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Set(index int, val ApiGetDeviceEntitiesLifecycleRequest) *ApiGetDeviceEntitiesLifecycleRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Skip(skip int) *ApiGetDeviceEntitiesLifecycleRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) SkippingEach(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) int) *ApiGetDeviceEntitiesLifecycleRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Slice(startIndex, n int) *ApiGetDeviceEntitiesLifecycleRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetDeviceEntitiesLifecycleRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Sort(fn func(i, j int) bool) *ApiGetDeviceEntitiesLifecycleRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Tail() *ApiGetDeviceEntitiesLifecycleRequest {
	return self.Last()
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) TailOr(arg ApiGetDeviceEntitiesLifecycleRequest) ApiGetDeviceEntitiesLifecycleRequest {
	return self.LastOr(arg)
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) ToList() []ApiGetDeviceEntitiesLifecycleRequest {
	return self.Val()
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Unique() *ApiGetDeviceEntitiesLifecycleRequestStream {
	return self.Distinct()
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Val() []ApiGetDeviceEntitiesLifecycleRequest {
	if self == nil {
		return []ApiGetDeviceEntitiesLifecycleRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) While(fn func(ApiGetDeviceEntitiesLifecycleRequest, int) bool) *ApiGetDeviceEntitiesLifecycleRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) Where(fn func(ApiGetDeviceEntitiesLifecycleRequest) bool) *ApiGetDeviceEntitiesLifecycleRequestStream {
	result := ApiGetDeviceEntitiesLifecycleRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDeviceEntitiesLifecycleRequestStream) WhereSlim(fn func(ApiGetDeviceEntitiesLifecycleRequest) bool) *ApiGetDeviceEntitiesLifecycleRequestStream {
	result := ApiGetDeviceEntitiesLifecycleRequestStreamOf()
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
