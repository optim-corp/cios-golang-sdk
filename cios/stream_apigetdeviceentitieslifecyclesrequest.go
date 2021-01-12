/*
 * Collection utility of ApiGetDeviceEntitiesLifecyclesRequest Struct
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

type ApiGetDeviceEntitiesLifecyclesRequestStream []ApiGetDeviceEntitiesLifecyclesRequest

func ApiGetDeviceEntitiesLifecyclesRequestStreamOf(arg ...ApiGetDeviceEntitiesLifecyclesRequest) ApiGetDeviceEntitiesLifecyclesRequestStream {
	return arg
}
func ApiGetDeviceEntitiesLifecyclesRequestStreamFrom(arg []ApiGetDeviceEntitiesLifecyclesRequest) ApiGetDeviceEntitiesLifecyclesRequestStream {
	return arg
}
func CreateApiGetDeviceEntitiesLifecyclesRequestStream(arg ...ApiGetDeviceEntitiesLifecyclesRequest) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	tmp := ApiGetDeviceEntitiesLifecyclesRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetDeviceEntitiesLifecyclesRequestStream(arg []ApiGetDeviceEntitiesLifecyclesRequest) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	tmp := ApiGetDeviceEntitiesLifecyclesRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Add(arg ApiGetDeviceEntitiesLifecyclesRequest) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) AddAll(arg ...ApiGetDeviceEntitiesLifecyclesRequest) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) AddSafe(arg *ApiGetDeviceEntitiesLifecyclesRequest) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Aggregate(fn func(ApiGetDeviceEntitiesLifecyclesRequest, ApiGetDeviceEntitiesLifecyclesRequest) ApiGetDeviceEntitiesLifecyclesRequest) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	result := ApiGetDeviceEntitiesLifecyclesRequestStreamOf()
	self.ForEach(func(v ApiGetDeviceEntitiesLifecyclesRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetDeviceEntitiesLifecyclesRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) AllMatch(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) AnyMatch(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Clone() *ApiGetDeviceEntitiesLifecyclesRequestStream {
	temp := make([]ApiGetDeviceEntitiesLifecyclesRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetDeviceEntitiesLifecyclesRequestStream)(&temp)
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Copy() *ApiGetDeviceEntitiesLifecyclesRequestStream {
	return self.Clone()
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Concat(arg []ApiGetDeviceEntitiesLifecyclesRequest) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Contains(arg ApiGetDeviceEntitiesLifecyclesRequest) bool {
	return self.FindIndex(func(_arg ApiGetDeviceEntitiesLifecyclesRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Clean() *ApiGetDeviceEntitiesLifecyclesRequestStream {
	*self = ApiGetDeviceEntitiesLifecyclesRequestStreamOf()
	return self
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Delete(index int) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Distinct() *ApiGetDeviceEntitiesLifecyclesRequestStream {
	caches := map[string]bool{}
	result := ApiGetDeviceEntitiesLifecyclesRequestStreamOf()
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
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Each(fn func(ApiGetDeviceEntitiesLifecyclesRequest)) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) EachRight(fn func(ApiGetDeviceEntitiesLifecyclesRequest)) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Equals(arr []ApiGetDeviceEntitiesLifecyclesRequest) bool {
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
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Filter(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) bool) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	result := ApiGetDeviceEntitiesLifecyclesRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) FilterSlim(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) bool) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	result := ApiGetDeviceEntitiesLifecyclesRequestStreamOf()
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
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Find(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) bool) *ApiGetDeviceEntitiesLifecyclesRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) FindOr(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) bool, or ApiGetDeviceEntitiesLifecyclesRequest) ApiGetDeviceEntitiesLifecyclesRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) FindIndex(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) bool) int {
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
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) First() *ApiGetDeviceEntitiesLifecyclesRequest {
	return self.Get(0)
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) FirstOr(arg ApiGetDeviceEntitiesLifecyclesRequest) ApiGetDeviceEntitiesLifecyclesRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) ForEach(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int)) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) ForEachRight(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int)) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) GroupBy(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) string) map[string][]ApiGetDeviceEntitiesLifecyclesRequest {
	m := map[string][]ApiGetDeviceEntitiesLifecyclesRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) GroupByValues(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) string) [][]ApiGetDeviceEntitiesLifecyclesRequest {
	var tmp [][]ApiGetDeviceEntitiesLifecyclesRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) IndexOf(arg ApiGetDeviceEntitiesLifecyclesRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Last() *ApiGetDeviceEntitiesLifecyclesRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) LastOr(arg ApiGetDeviceEntitiesLifecyclesRequest) ApiGetDeviceEntitiesLifecyclesRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Limit(limit int) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Map(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Map2Int(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Map2Int32(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Map2Int64(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Map2Float32(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Map2Float64(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Map2Bool(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Map2Bytes(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Map2String(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Max(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) float64) *ApiGetDeviceEntitiesLifecyclesRequest {
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
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Min(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) float64) *ApiGetDeviceEntitiesLifecyclesRequest {
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
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) NoneMatch(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Get(index int) *ApiGetDeviceEntitiesLifecyclesRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) GetOr(index int, arg ApiGetDeviceEntitiesLifecyclesRequest) ApiGetDeviceEntitiesLifecyclesRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Peek(fn func(*ApiGetDeviceEntitiesLifecyclesRequest, int)) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Reduce(fn func(ApiGetDeviceEntitiesLifecyclesRequest, ApiGetDeviceEntitiesLifecyclesRequest, int) ApiGetDeviceEntitiesLifecyclesRequest) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	return self.ReduceInit(fn, ApiGetDeviceEntitiesLifecyclesRequest{})
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) ReduceInit(fn func(ApiGetDeviceEntitiesLifecyclesRequest, ApiGetDeviceEntitiesLifecyclesRequest, int) ApiGetDeviceEntitiesLifecyclesRequest, initialValue ApiGetDeviceEntitiesLifecyclesRequest) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	result := ApiGetDeviceEntitiesLifecyclesRequestStreamOf()
	self.ForEach(func(v ApiGetDeviceEntitiesLifecyclesRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) ReduceInterface(fn func(interface{}, ApiGetDeviceEntitiesLifecyclesRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetDeviceEntitiesLifecyclesRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) ReduceString(fn func(string, ApiGetDeviceEntitiesLifecyclesRequest, int) string) []string {
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
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) ReduceInt(fn func(int, ApiGetDeviceEntitiesLifecyclesRequest, int) int) []int {
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
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) ReduceInt32(fn func(int32, ApiGetDeviceEntitiesLifecyclesRequest, int) int32) []int32 {
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
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) ReduceInt64(fn func(int64, ApiGetDeviceEntitiesLifecyclesRequest, int) int64) []int64 {
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
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) ReduceFloat32(fn func(float32, ApiGetDeviceEntitiesLifecyclesRequest, int) float32) []float32 {
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
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) ReduceFloat64(fn func(float64, ApiGetDeviceEntitiesLifecyclesRequest, int) float64) []float64 {
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
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) ReduceBool(fn func(bool, ApiGetDeviceEntitiesLifecyclesRequest, int) bool) []bool {
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
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Reverse() *ApiGetDeviceEntitiesLifecyclesRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Replace(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) ApiGetDeviceEntitiesLifecyclesRequest) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	return self.ForEach(func(v ApiGetDeviceEntitiesLifecyclesRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Select(fn func(ApiGetDeviceEntitiesLifecyclesRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Set(index int, val ApiGetDeviceEntitiesLifecyclesRequest) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Skip(skip int) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) SkippingEach(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) int) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Slice(startIndex, n int) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetDeviceEntitiesLifecyclesRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Sort(fn func(i, j int) bool) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Tail() *ApiGetDeviceEntitiesLifecyclesRequest {
	return self.Last()
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) TailOr(arg ApiGetDeviceEntitiesLifecyclesRequest) ApiGetDeviceEntitiesLifecyclesRequest {
	return self.LastOr(arg)
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) ToList() []ApiGetDeviceEntitiesLifecyclesRequest {
	return self.Val()
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Unique() *ApiGetDeviceEntitiesLifecyclesRequestStream {
	return self.Distinct()
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Val() []ApiGetDeviceEntitiesLifecyclesRequest {
	if self == nil {
		return []ApiGetDeviceEntitiesLifecyclesRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) While(fn func(ApiGetDeviceEntitiesLifecyclesRequest, int) bool) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) Where(fn func(ApiGetDeviceEntitiesLifecyclesRequest) bool) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	result := ApiGetDeviceEntitiesLifecyclesRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDeviceEntitiesLifecyclesRequestStream) WhereSlim(fn func(ApiGetDeviceEntitiesLifecyclesRequest) bool) *ApiGetDeviceEntitiesLifecyclesRequestStream {
	result := ApiGetDeviceEntitiesLifecyclesRequestStreamOf()
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
