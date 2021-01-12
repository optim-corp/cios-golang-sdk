/*
 * Collection utility of ApiUpdateRouteRequest Struct
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

type ApiUpdateRouteRequestStream []ApiUpdateRouteRequest

func ApiUpdateRouteRequestStreamOf(arg ...ApiUpdateRouteRequest) ApiUpdateRouteRequestStream {
	return arg
}
func ApiUpdateRouteRequestStreamFrom(arg []ApiUpdateRouteRequest) ApiUpdateRouteRequestStream {
	return arg
}
func CreateApiUpdateRouteRequestStream(arg ...ApiUpdateRouteRequest) *ApiUpdateRouteRequestStream {
	tmp := ApiUpdateRouteRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiUpdateRouteRequestStream(arg []ApiUpdateRouteRequest) *ApiUpdateRouteRequestStream {
	tmp := ApiUpdateRouteRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiUpdateRouteRequestStream) Add(arg ApiUpdateRouteRequest) *ApiUpdateRouteRequestStream {
	return self.AddAll(arg)
}
func (self *ApiUpdateRouteRequestStream) AddAll(arg ...ApiUpdateRouteRequest) *ApiUpdateRouteRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiUpdateRouteRequestStream) AddSafe(arg *ApiUpdateRouteRequest) *ApiUpdateRouteRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiUpdateRouteRequestStream) Aggregate(fn func(ApiUpdateRouteRequest, ApiUpdateRouteRequest) ApiUpdateRouteRequest) *ApiUpdateRouteRequestStream {
	result := ApiUpdateRouteRequestStreamOf()
	self.ForEach(func(v ApiUpdateRouteRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiUpdateRouteRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiUpdateRouteRequestStream) AllMatch(fn func(ApiUpdateRouteRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiUpdateRouteRequestStream) AnyMatch(fn func(ApiUpdateRouteRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiUpdateRouteRequestStream) Clone() *ApiUpdateRouteRequestStream {
	temp := make([]ApiUpdateRouteRequest, self.Len())
	copy(temp, *self)
	return (*ApiUpdateRouteRequestStream)(&temp)
}
func (self *ApiUpdateRouteRequestStream) Copy() *ApiUpdateRouteRequestStream {
	return self.Clone()
}
func (self *ApiUpdateRouteRequestStream) Concat(arg []ApiUpdateRouteRequest) *ApiUpdateRouteRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiUpdateRouteRequestStream) Contains(arg ApiUpdateRouteRequest) bool {
	return self.FindIndex(func(_arg ApiUpdateRouteRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiUpdateRouteRequestStream) Clean() *ApiUpdateRouteRequestStream {
	*self = ApiUpdateRouteRequestStreamOf()
	return self
}
func (self *ApiUpdateRouteRequestStream) Delete(index int) *ApiUpdateRouteRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiUpdateRouteRequestStream) DeleteRange(startIndex, endIndex int) *ApiUpdateRouteRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiUpdateRouteRequestStream) Distinct() *ApiUpdateRouteRequestStream {
	caches := map[string]bool{}
	result := ApiUpdateRouteRequestStreamOf()
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
func (self *ApiUpdateRouteRequestStream) Each(fn func(ApiUpdateRouteRequest)) *ApiUpdateRouteRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiUpdateRouteRequestStream) EachRight(fn func(ApiUpdateRouteRequest)) *ApiUpdateRouteRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiUpdateRouteRequestStream) Equals(arr []ApiUpdateRouteRequest) bool {
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
func (self *ApiUpdateRouteRequestStream) Filter(fn func(ApiUpdateRouteRequest, int) bool) *ApiUpdateRouteRequestStream {
	result := ApiUpdateRouteRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiUpdateRouteRequestStream) FilterSlim(fn func(ApiUpdateRouteRequest, int) bool) *ApiUpdateRouteRequestStream {
	result := ApiUpdateRouteRequestStreamOf()
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
func (self *ApiUpdateRouteRequestStream) Find(fn func(ApiUpdateRouteRequest, int) bool) *ApiUpdateRouteRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiUpdateRouteRequestStream) FindOr(fn func(ApiUpdateRouteRequest, int) bool, or ApiUpdateRouteRequest) ApiUpdateRouteRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiUpdateRouteRequestStream) FindIndex(fn func(ApiUpdateRouteRequest, int) bool) int {
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
func (self *ApiUpdateRouteRequestStream) First() *ApiUpdateRouteRequest {
	return self.Get(0)
}
func (self *ApiUpdateRouteRequestStream) FirstOr(arg ApiUpdateRouteRequest) ApiUpdateRouteRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiUpdateRouteRequestStream) ForEach(fn func(ApiUpdateRouteRequest, int)) *ApiUpdateRouteRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiUpdateRouteRequestStream) ForEachRight(fn func(ApiUpdateRouteRequest, int)) *ApiUpdateRouteRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiUpdateRouteRequestStream) GroupBy(fn func(ApiUpdateRouteRequest, int) string) map[string][]ApiUpdateRouteRequest {
	m := map[string][]ApiUpdateRouteRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiUpdateRouteRequestStream) GroupByValues(fn func(ApiUpdateRouteRequest, int) string) [][]ApiUpdateRouteRequest {
	var tmp [][]ApiUpdateRouteRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiUpdateRouteRequestStream) IndexOf(arg ApiUpdateRouteRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiUpdateRouteRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiUpdateRouteRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiUpdateRouteRequestStream) Last() *ApiUpdateRouteRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiUpdateRouteRequestStream) LastOr(arg ApiUpdateRouteRequest) ApiUpdateRouteRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiUpdateRouteRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiUpdateRouteRequestStream) Limit(limit int) *ApiUpdateRouteRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiUpdateRouteRequestStream) Map(fn func(ApiUpdateRouteRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateRouteRequestStream) Map2Int(fn func(ApiUpdateRouteRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateRouteRequestStream) Map2Int32(fn func(ApiUpdateRouteRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateRouteRequestStream) Map2Int64(fn func(ApiUpdateRouteRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateRouteRequestStream) Map2Float32(fn func(ApiUpdateRouteRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateRouteRequestStream) Map2Float64(fn func(ApiUpdateRouteRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateRouteRequestStream) Map2Bool(fn func(ApiUpdateRouteRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateRouteRequestStream) Map2Bytes(fn func(ApiUpdateRouteRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateRouteRequestStream) Map2String(fn func(ApiUpdateRouteRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateRouteRequestStream) Max(fn func(ApiUpdateRouteRequest, int) float64) *ApiUpdateRouteRequest {
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
func (self *ApiUpdateRouteRequestStream) Min(fn func(ApiUpdateRouteRequest, int) float64) *ApiUpdateRouteRequest {
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
func (self *ApiUpdateRouteRequestStream) NoneMatch(fn func(ApiUpdateRouteRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiUpdateRouteRequestStream) Get(index int) *ApiUpdateRouteRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiUpdateRouteRequestStream) GetOr(index int, arg ApiUpdateRouteRequest) ApiUpdateRouteRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiUpdateRouteRequestStream) Peek(fn func(*ApiUpdateRouteRequest, int)) *ApiUpdateRouteRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiUpdateRouteRequestStream) Reduce(fn func(ApiUpdateRouteRequest, ApiUpdateRouteRequest, int) ApiUpdateRouteRequest) *ApiUpdateRouteRequestStream {
	return self.ReduceInit(fn, ApiUpdateRouteRequest{})
}
func (self *ApiUpdateRouteRequestStream) ReduceInit(fn func(ApiUpdateRouteRequest, ApiUpdateRouteRequest, int) ApiUpdateRouteRequest, initialValue ApiUpdateRouteRequest) *ApiUpdateRouteRequestStream {
	result := ApiUpdateRouteRequestStreamOf()
	self.ForEach(func(v ApiUpdateRouteRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiUpdateRouteRequestStream) ReduceInterface(fn func(interface{}, ApiUpdateRouteRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiUpdateRouteRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiUpdateRouteRequestStream) ReduceString(fn func(string, ApiUpdateRouteRequest, int) string) []string {
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
func (self *ApiUpdateRouteRequestStream) ReduceInt(fn func(int, ApiUpdateRouteRequest, int) int) []int {
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
func (self *ApiUpdateRouteRequestStream) ReduceInt32(fn func(int32, ApiUpdateRouteRequest, int) int32) []int32 {
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
func (self *ApiUpdateRouteRequestStream) ReduceInt64(fn func(int64, ApiUpdateRouteRequest, int) int64) []int64 {
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
func (self *ApiUpdateRouteRequestStream) ReduceFloat32(fn func(float32, ApiUpdateRouteRequest, int) float32) []float32 {
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
func (self *ApiUpdateRouteRequestStream) ReduceFloat64(fn func(float64, ApiUpdateRouteRequest, int) float64) []float64 {
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
func (self *ApiUpdateRouteRequestStream) ReduceBool(fn func(bool, ApiUpdateRouteRequest, int) bool) []bool {
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
func (self *ApiUpdateRouteRequestStream) Reverse() *ApiUpdateRouteRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiUpdateRouteRequestStream) Replace(fn func(ApiUpdateRouteRequest, int) ApiUpdateRouteRequest) *ApiUpdateRouteRequestStream {
	return self.ForEach(func(v ApiUpdateRouteRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiUpdateRouteRequestStream) Select(fn func(ApiUpdateRouteRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiUpdateRouteRequestStream) Set(index int, val ApiUpdateRouteRequest) *ApiUpdateRouteRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiUpdateRouteRequestStream) Skip(skip int) *ApiUpdateRouteRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiUpdateRouteRequestStream) SkippingEach(fn func(ApiUpdateRouteRequest, int) int) *ApiUpdateRouteRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiUpdateRouteRequestStream) Slice(startIndex, n int) *ApiUpdateRouteRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiUpdateRouteRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiUpdateRouteRequestStream) Sort(fn func(i, j int) bool) *ApiUpdateRouteRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiUpdateRouteRequestStream) Tail() *ApiUpdateRouteRequest {
	return self.Last()
}
func (self *ApiUpdateRouteRequestStream) TailOr(arg ApiUpdateRouteRequest) ApiUpdateRouteRequest {
	return self.LastOr(arg)
}
func (self *ApiUpdateRouteRequestStream) ToList() []ApiUpdateRouteRequest {
	return self.Val()
}
func (self *ApiUpdateRouteRequestStream) Unique() *ApiUpdateRouteRequestStream {
	return self.Distinct()
}
func (self *ApiUpdateRouteRequestStream) Val() []ApiUpdateRouteRequest {
	if self == nil {
		return []ApiUpdateRouteRequest{}
	}
	return *self.Copy()
}
func (self *ApiUpdateRouteRequestStream) While(fn func(ApiUpdateRouteRequest, int) bool) *ApiUpdateRouteRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiUpdateRouteRequestStream) Where(fn func(ApiUpdateRouteRequest) bool) *ApiUpdateRouteRequestStream {
	result := ApiUpdateRouteRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiUpdateRouteRequestStream) WhereSlim(fn func(ApiUpdateRouteRequest) bool) *ApiUpdateRouteRequestStream {
	result := ApiUpdateRouteRequestStreamOf()
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
