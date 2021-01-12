/*
 * Collection utility of ApiCreateVideoStreamsPlayRequest Struct
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

type ApiCreateVideoStreamsPlayRequestStream []ApiCreateVideoStreamsPlayRequest

func ApiCreateVideoStreamsPlayRequestStreamOf(arg ...ApiCreateVideoStreamsPlayRequest) ApiCreateVideoStreamsPlayRequestStream {
	return arg
}
func ApiCreateVideoStreamsPlayRequestStreamFrom(arg []ApiCreateVideoStreamsPlayRequest) ApiCreateVideoStreamsPlayRequestStream {
	return arg
}
func CreateApiCreateVideoStreamsPlayRequestStream(arg ...ApiCreateVideoStreamsPlayRequest) *ApiCreateVideoStreamsPlayRequestStream {
	tmp := ApiCreateVideoStreamsPlayRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiCreateVideoStreamsPlayRequestStream(arg []ApiCreateVideoStreamsPlayRequest) *ApiCreateVideoStreamsPlayRequestStream {
	tmp := ApiCreateVideoStreamsPlayRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiCreateVideoStreamsPlayRequestStream) Add(arg ApiCreateVideoStreamsPlayRequest) *ApiCreateVideoStreamsPlayRequestStream {
	return self.AddAll(arg)
}
func (self *ApiCreateVideoStreamsPlayRequestStream) AddAll(arg ...ApiCreateVideoStreamsPlayRequest) *ApiCreateVideoStreamsPlayRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiCreateVideoStreamsPlayRequestStream) AddSafe(arg *ApiCreateVideoStreamsPlayRequest) *ApiCreateVideoStreamsPlayRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Aggregate(fn func(ApiCreateVideoStreamsPlayRequest, ApiCreateVideoStreamsPlayRequest) ApiCreateVideoStreamsPlayRequest) *ApiCreateVideoStreamsPlayRequestStream {
	result := ApiCreateVideoStreamsPlayRequestStreamOf()
	self.ForEach(func(v ApiCreateVideoStreamsPlayRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiCreateVideoStreamsPlayRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiCreateVideoStreamsPlayRequestStream) AllMatch(fn func(ApiCreateVideoStreamsPlayRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiCreateVideoStreamsPlayRequestStream) AnyMatch(fn func(ApiCreateVideoStreamsPlayRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Clone() *ApiCreateVideoStreamsPlayRequestStream {
	temp := make([]ApiCreateVideoStreamsPlayRequest, self.Len())
	copy(temp, *self)
	return (*ApiCreateVideoStreamsPlayRequestStream)(&temp)
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Copy() *ApiCreateVideoStreamsPlayRequestStream {
	return self.Clone()
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Concat(arg []ApiCreateVideoStreamsPlayRequest) *ApiCreateVideoStreamsPlayRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Contains(arg ApiCreateVideoStreamsPlayRequest) bool {
	return self.FindIndex(func(_arg ApiCreateVideoStreamsPlayRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Clean() *ApiCreateVideoStreamsPlayRequestStream {
	*self = ApiCreateVideoStreamsPlayRequestStreamOf()
	return self
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Delete(index int) *ApiCreateVideoStreamsPlayRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiCreateVideoStreamsPlayRequestStream) DeleteRange(startIndex, endIndex int) *ApiCreateVideoStreamsPlayRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Distinct() *ApiCreateVideoStreamsPlayRequestStream {
	caches := map[string]bool{}
	result := ApiCreateVideoStreamsPlayRequestStreamOf()
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
func (self *ApiCreateVideoStreamsPlayRequestStream) Each(fn func(ApiCreateVideoStreamsPlayRequest)) *ApiCreateVideoStreamsPlayRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiCreateVideoStreamsPlayRequestStream) EachRight(fn func(ApiCreateVideoStreamsPlayRequest)) *ApiCreateVideoStreamsPlayRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Equals(arr []ApiCreateVideoStreamsPlayRequest) bool {
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
func (self *ApiCreateVideoStreamsPlayRequestStream) Filter(fn func(ApiCreateVideoStreamsPlayRequest, int) bool) *ApiCreateVideoStreamsPlayRequestStream {
	result := ApiCreateVideoStreamsPlayRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiCreateVideoStreamsPlayRequestStream) FilterSlim(fn func(ApiCreateVideoStreamsPlayRequest, int) bool) *ApiCreateVideoStreamsPlayRequestStream {
	result := ApiCreateVideoStreamsPlayRequestStreamOf()
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
func (self *ApiCreateVideoStreamsPlayRequestStream) Find(fn func(ApiCreateVideoStreamsPlayRequest, int) bool) *ApiCreateVideoStreamsPlayRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiCreateVideoStreamsPlayRequestStream) FindOr(fn func(ApiCreateVideoStreamsPlayRequest, int) bool, or ApiCreateVideoStreamsPlayRequest) ApiCreateVideoStreamsPlayRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiCreateVideoStreamsPlayRequestStream) FindIndex(fn func(ApiCreateVideoStreamsPlayRequest, int) bool) int {
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
func (self *ApiCreateVideoStreamsPlayRequestStream) First() *ApiCreateVideoStreamsPlayRequest {
	return self.Get(0)
}
func (self *ApiCreateVideoStreamsPlayRequestStream) FirstOr(arg ApiCreateVideoStreamsPlayRequest) ApiCreateVideoStreamsPlayRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiCreateVideoStreamsPlayRequestStream) ForEach(fn func(ApiCreateVideoStreamsPlayRequest, int)) *ApiCreateVideoStreamsPlayRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiCreateVideoStreamsPlayRequestStream) ForEachRight(fn func(ApiCreateVideoStreamsPlayRequest, int)) *ApiCreateVideoStreamsPlayRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiCreateVideoStreamsPlayRequestStream) GroupBy(fn func(ApiCreateVideoStreamsPlayRequest, int) string) map[string][]ApiCreateVideoStreamsPlayRequest {
	m := map[string][]ApiCreateVideoStreamsPlayRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiCreateVideoStreamsPlayRequestStream) GroupByValues(fn func(ApiCreateVideoStreamsPlayRequest, int) string) [][]ApiCreateVideoStreamsPlayRequest {
	var tmp [][]ApiCreateVideoStreamsPlayRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiCreateVideoStreamsPlayRequestStream) IndexOf(arg ApiCreateVideoStreamsPlayRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiCreateVideoStreamsPlayRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiCreateVideoStreamsPlayRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Last() *ApiCreateVideoStreamsPlayRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiCreateVideoStreamsPlayRequestStream) LastOr(arg ApiCreateVideoStreamsPlayRequest) ApiCreateVideoStreamsPlayRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Limit(limit int) *ApiCreateVideoStreamsPlayRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiCreateVideoStreamsPlayRequestStream) Map(fn func(ApiCreateVideoStreamsPlayRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Map2Int(fn func(ApiCreateVideoStreamsPlayRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Map2Int32(fn func(ApiCreateVideoStreamsPlayRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Map2Int64(fn func(ApiCreateVideoStreamsPlayRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Map2Float32(fn func(ApiCreateVideoStreamsPlayRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Map2Float64(fn func(ApiCreateVideoStreamsPlayRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Map2Bool(fn func(ApiCreateVideoStreamsPlayRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Map2Bytes(fn func(ApiCreateVideoStreamsPlayRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Map2String(fn func(ApiCreateVideoStreamsPlayRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Max(fn func(ApiCreateVideoStreamsPlayRequest, int) float64) *ApiCreateVideoStreamsPlayRequest {
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
func (self *ApiCreateVideoStreamsPlayRequestStream) Min(fn func(ApiCreateVideoStreamsPlayRequest, int) float64) *ApiCreateVideoStreamsPlayRequest {
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
func (self *ApiCreateVideoStreamsPlayRequestStream) NoneMatch(fn func(ApiCreateVideoStreamsPlayRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Get(index int) *ApiCreateVideoStreamsPlayRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiCreateVideoStreamsPlayRequestStream) GetOr(index int, arg ApiCreateVideoStreamsPlayRequest) ApiCreateVideoStreamsPlayRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Peek(fn func(*ApiCreateVideoStreamsPlayRequest, int)) *ApiCreateVideoStreamsPlayRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiCreateVideoStreamsPlayRequestStream) Reduce(fn func(ApiCreateVideoStreamsPlayRequest, ApiCreateVideoStreamsPlayRequest, int) ApiCreateVideoStreamsPlayRequest) *ApiCreateVideoStreamsPlayRequestStream {
	return self.ReduceInit(fn, ApiCreateVideoStreamsPlayRequest{})
}
func (self *ApiCreateVideoStreamsPlayRequestStream) ReduceInit(fn func(ApiCreateVideoStreamsPlayRequest, ApiCreateVideoStreamsPlayRequest, int) ApiCreateVideoStreamsPlayRequest, initialValue ApiCreateVideoStreamsPlayRequest) *ApiCreateVideoStreamsPlayRequestStream {
	result := ApiCreateVideoStreamsPlayRequestStreamOf()
	self.ForEach(func(v ApiCreateVideoStreamsPlayRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiCreateVideoStreamsPlayRequestStream) ReduceInterface(fn func(interface{}, ApiCreateVideoStreamsPlayRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiCreateVideoStreamsPlayRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiCreateVideoStreamsPlayRequestStream) ReduceString(fn func(string, ApiCreateVideoStreamsPlayRequest, int) string) []string {
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
func (self *ApiCreateVideoStreamsPlayRequestStream) ReduceInt(fn func(int, ApiCreateVideoStreamsPlayRequest, int) int) []int {
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
func (self *ApiCreateVideoStreamsPlayRequestStream) ReduceInt32(fn func(int32, ApiCreateVideoStreamsPlayRequest, int) int32) []int32 {
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
func (self *ApiCreateVideoStreamsPlayRequestStream) ReduceInt64(fn func(int64, ApiCreateVideoStreamsPlayRequest, int) int64) []int64 {
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
func (self *ApiCreateVideoStreamsPlayRequestStream) ReduceFloat32(fn func(float32, ApiCreateVideoStreamsPlayRequest, int) float32) []float32 {
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
func (self *ApiCreateVideoStreamsPlayRequestStream) ReduceFloat64(fn func(float64, ApiCreateVideoStreamsPlayRequest, int) float64) []float64 {
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
func (self *ApiCreateVideoStreamsPlayRequestStream) ReduceBool(fn func(bool, ApiCreateVideoStreamsPlayRequest, int) bool) []bool {
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
func (self *ApiCreateVideoStreamsPlayRequestStream) Reverse() *ApiCreateVideoStreamsPlayRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Replace(fn func(ApiCreateVideoStreamsPlayRequest, int) ApiCreateVideoStreamsPlayRequest) *ApiCreateVideoStreamsPlayRequestStream {
	return self.ForEach(func(v ApiCreateVideoStreamsPlayRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Select(fn func(ApiCreateVideoStreamsPlayRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Set(index int, val ApiCreateVideoStreamsPlayRequest) *ApiCreateVideoStreamsPlayRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Skip(skip int) *ApiCreateVideoStreamsPlayRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiCreateVideoStreamsPlayRequestStream) SkippingEach(fn func(ApiCreateVideoStreamsPlayRequest, int) int) *ApiCreateVideoStreamsPlayRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Slice(startIndex, n int) *ApiCreateVideoStreamsPlayRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiCreateVideoStreamsPlayRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Sort(fn func(i, j int) bool) *ApiCreateVideoStreamsPlayRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiCreateVideoStreamsPlayRequestStream) Tail() *ApiCreateVideoStreamsPlayRequest {
	return self.Last()
}
func (self *ApiCreateVideoStreamsPlayRequestStream) TailOr(arg ApiCreateVideoStreamsPlayRequest) ApiCreateVideoStreamsPlayRequest {
	return self.LastOr(arg)
}
func (self *ApiCreateVideoStreamsPlayRequestStream) ToList() []ApiCreateVideoStreamsPlayRequest {
	return self.Val()
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Unique() *ApiCreateVideoStreamsPlayRequestStream {
	return self.Distinct()
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Val() []ApiCreateVideoStreamsPlayRequest {
	if self == nil {
		return []ApiCreateVideoStreamsPlayRequest{}
	}
	return *self.Copy()
}
func (self *ApiCreateVideoStreamsPlayRequestStream) While(fn func(ApiCreateVideoStreamsPlayRequest, int) bool) *ApiCreateVideoStreamsPlayRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiCreateVideoStreamsPlayRequestStream) Where(fn func(ApiCreateVideoStreamsPlayRequest) bool) *ApiCreateVideoStreamsPlayRequestStream {
	result := ApiCreateVideoStreamsPlayRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiCreateVideoStreamsPlayRequestStream) WhereSlim(fn func(ApiCreateVideoStreamsPlayRequest) bool) *ApiCreateVideoStreamsPlayRequestStream {
	result := ApiCreateVideoStreamsPlayRequestStreamOf()
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
