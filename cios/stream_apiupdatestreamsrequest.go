/*
 * Collection utility of ApiUpdateStreamsRequest Struct
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

type ApiUpdateStreamsRequestStream []ApiUpdateStreamsRequest

func ApiUpdateStreamsRequestStreamOf(arg ...ApiUpdateStreamsRequest) ApiUpdateStreamsRequestStream {
	return arg
}
func ApiUpdateStreamsRequestStreamFrom(arg []ApiUpdateStreamsRequest) ApiUpdateStreamsRequestStream {
	return arg
}
func CreateApiUpdateStreamsRequestStream(arg ...ApiUpdateStreamsRequest) *ApiUpdateStreamsRequestStream {
	tmp := ApiUpdateStreamsRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiUpdateStreamsRequestStream(arg []ApiUpdateStreamsRequest) *ApiUpdateStreamsRequestStream {
	tmp := ApiUpdateStreamsRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiUpdateStreamsRequestStream) Add(arg ApiUpdateStreamsRequest) *ApiUpdateStreamsRequestStream {
	return self.AddAll(arg)
}
func (self *ApiUpdateStreamsRequestStream) AddAll(arg ...ApiUpdateStreamsRequest) *ApiUpdateStreamsRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiUpdateStreamsRequestStream) AddSafe(arg *ApiUpdateStreamsRequest) *ApiUpdateStreamsRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiUpdateStreamsRequestStream) Aggregate(fn func(ApiUpdateStreamsRequest, ApiUpdateStreamsRequest) ApiUpdateStreamsRequest) *ApiUpdateStreamsRequestStream {
	result := ApiUpdateStreamsRequestStreamOf()
	self.ForEach(func(v ApiUpdateStreamsRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiUpdateStreamsRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiUpdateStreamsRequestStream) AllMatch(fn func(ApiUpdateStreamsRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiUpdateStreamsRequestStream) AnyMatch(fn func(ApiUpdateStreamsRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiUpdateStreamsRequestStream) Clone() *ApiUpdateStreamsRequestStream {
	temp := make([]ApiUpdateStreamsRequest, self.Len())
	copy(temp, *self)
	return (*ApiUpdateStreamsRequestStream)(&temp)
}
func (self *ApiUpdateStreamsRequestStream) Copy() *ApiUpdateStreamsRequestStream {
	return self.Clone()
}
func (self *ApiUpdateStreamsRequestStream) Concat(arg []ApiUpdateStreamsRequest) *ApiUpdateStreamsRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiUpdateStreamsRequestStream) Contains(arg ApiUpdateStreamsRequest) bool {
	return self.FindIndex(func(_arg ApiUpdateStreamsRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiUpdateStreamsRequestStream) Clean() *ApiUpdateStreamsRequestStream {
	*self = ApiUpdateStreamsRequestStreamOf()
	return self
}
func (self *ApiUpdateStreamsRequestStream) Delete(index int) *ApiUpdateStreamsRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiUpdateStreamsRequestStream) DeleteRange(startIndex, endIndex int) *ApiUpdateStreamsRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiUpdateStreamsRequestStream) Distinct() *ApiUpdateStreamsRequestStream {
	caches := map[string]bool{}
	result := ApiUpdateStreamsRequestStreamOf()
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
func (self *ApiUpdateStreamsRequestStream) Each(fn func(ApiUpdateStreamsRequest)) *ApiUpdateStreamsRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiUpdateStreamsRequestStream) EachRight(fn func(ApiUpdateStreamsRequest)) *ApiUpdateStreamsRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiUpdateStreamsRequestStream) Equals(arr []ApiUpdateStreamsRequest) bool {
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
func (self *ApiUpdateStreamsRequestStream) Filter(fn func(ApiUpdateStreamsRequest, int) bool) *ApiUpdateStreamsRequestStream {
	result := ApiUpdateStreamsRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiUpdateStreamsRequestStream) FilterSlim(fn func(ApiUpdateStreamsRequest, int) bool) *ApiUpdateStreamsRequestStream {
	result := ApiUpdateStreamsRequestStreamOf()
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
func (self *ApiUpdateStreamsRequestStream) Find(fn func(ApiUpdateStreamsRequest, int) bool) *ApiUpdateStreamsRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiUpdateStreamsRequestStream) FindOr(fn func(ApiUpdateStreamsRequest, int) bool, or ApiUpdateStreamsRequest) ApiUpdateStreamsRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiUpdateStreamsRequestStream) FindIndex(fn func(ApiUpdateStreamsRequest, int) bool) int {
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
func (self *ApiUpdateStreamsRequestStream) First() *ApiUpdateStreamsRequest {
	return self.Get(0)
}
func (self *ApiUpdateStreamsRequestStream) FirstOr(arg ApiUpdateStreamsRequest) ApiUpdateStreamsRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiUpdateStreamsRequestStream) ForEach(fn func(ApiUpdateStreamsRequest, int)) *ApiUpdateStreamsRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiUpdateStreamsRequestStream) ForEachRight(fn func(ApiUpdateStreamsRequest, int)) *ApiUpdateStreamsRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiUpdateStreamsRequestStream) GroupBy(fn func(ApiUpdateStreamsRequest, int) string) map[string][]ApiUpdateStreamsRequest {
	m := map[string][]ApiUpdateStreamsRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiUpdateStreamsRequestStream) GroupByValues(fn func(ApiUpdateStreamsRequest, int) string) [][]ApiUpdateStreamsRequest {
	var tmp [][]ApiUpdateStreamsRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiUpdateStreamsRequestStream) IndexOf(arg ApiUpdateStreamsRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiUpdateStreamsRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiUpdateStreamsRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiUpdateStreamsRequestStream) Last() *ApiUpdateStreamsRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiUpdateStreamsRequestStream) LastOr(arg ApiUpdateStreamsRequest) ApiUpdateStreamsRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiUpdateStreamsRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiUpdateStreamsRequestStream) Limit(limit int) *ApiUpdateStreamsRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiUpdateStreamsRequestStream) Map(fn func(ApiUpdateStreamsRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateStreamsRequestStream) Map2Int(fn func(ApiUpdateStreamsRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateStreamsRequestStream) Map2Int32(fn func(ApiUpdateStreamsRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateStreamsRequestStream) Map2Int64(fn func(ApiUpdateStreamsRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateStreamsRequestStream) Map2Float32(fn func(ApiUpdateStreamsRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateStreamsRequestStream) Map2Float64(fn func(ApiUpdateStreamsRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateStreamsRequestStream) Map2Bool(fn func(ApiUpdateStreamsRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateStreamsRequestStream) Map2Bytes(fn func(ApiUpdateStreamsRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateStreamsRequestStream) Map2String(fn func(ApiUpdateStreamsRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiUpdateStreamsRequestStream) Max(fn func(ApiUpdateStreamsRequest, int) float64) *ApiUpdateStreamsRequest {
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
func (self *ApiUpdateStreamsRequestStream) Min(fn func(ApiUpdateStreamsRequest, int) float64) *ApiUpdateStreamsRequest {
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
func (self *ApiUpdateStreamsRequestStream) NoneMatch(fn func(ApiUpdateStreamsRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiUpdateStreamsRequestStream) Get(index int) *ApiUpdateStreamsRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiUpdateStreamsRequestStream) GetOr(index int, arg ApiUpdateStreamsRequest) ApiUpdateStreamsRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiUpdateStreamsRequestStream) Peek(fn func(*ApiUpdateStreamsRequest, int)) *ApiUpdateStreamsRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiUpdateStreamsRequestStream) Reduce(fn func(ApiUpdateStreamsRequest, ApiUpdateStreamsRequest, int) ApiUpdateStreamsRequest) *ApiUpdateStreamsRequestStream {
	return self.ReduceInit(fn, ApiUpdateStreamsRequest{})
}
func (self *ApiUpdateStreamsRequestStream) ReduceInit(fn func(ApiUpdateStreamsRequest, ApiUpdateStreamsRequest, int) ApiUpdateStreamsRequest, initialValue ApiUpdateStreamsRequest) *ApiUpdateStreamsRequestStream {
	result := ApiUpdateStreamsRequestStreamOf()
	self.ForEach(func(v ApiUpdateStreamsRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiUpdateStreamsRequestStream) ReduceInterface(fn func(interface{}, ApiUpdateStreamsRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiUpdateStreamsRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiUpdateStreamsRequestStream) ReduceString(fn func(string, ApiUpdateStreamsRequest, int) string) []string {
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
func (self *ApiUpdateStreamsRequestStream) ReduceInt(fn func(int, ApiUpdateStreamsRequest, int) int) []int {
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
func (self *ApiUpdateStreamsRequestStream) ReduceInt32(fn func(int32, ApiUpdateStreamsRequest, int) int32) []int32 {
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
func (self *ApiUpdateStreamsRequestStream) ReduceInt64(fn func(int64, ApiUpdateStreamsRequest, int) int64) []int64 {
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
func (self *ApiUpdateStreamsRequestStream) ReduceFloat32(fn func(float32, ApiUpdateStreamsRequest, int) float32) []float32 {
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
func (self *ApiUpdateStreamsRequestStream) ReduceFloat64(fn func(float64, ApiUpdateStreamsRequest, int) float64) []float64 {
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
func (self *ApiUpdateStreamsRequestStream) ReduceBool(fn func(bool, ApiUpdateStreamsRequest, int) bool) []bool {
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
func (self *ApiUpdateStreamsRequestStream) Reverse() *ApiUpdateStreamsRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiUpdateStreamsRequestStream) Replace(fn func(ApiUpdateStreamsRequest, int) ApiUpdateStreamsRequest) *ApiUpdateStreamsRequestStream {
	return self.ForEach(func(v ApiUpdateStreamsRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiUpdateStreamsRequestStream) Select(fn func(ApiUpdateStreamsRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiUpdateStreamsRequestStream) Set(index int, val ApiUpdateStreamsRequest) *ApiUpdateStreamsRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiUpdateStreamsRequestStream) Skip(skip int) *ApiUpdateStreamsRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiUpdateStreamsRequestStream) SkippingEach(fn func(ApiUpdateStreamsRequest, int) int) *ApiUpdateStreamsRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiUpdateStreamsRequestStream) Slice(startIndex, n int) *ApiUpdateStreamsRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiUpdateStreamsRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiUpdateStreamsRequestStream) Sort(fn func(i, j int) bool) *ApiUpdateStreamsRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiUpdateStreamsRequestStream) Tail() *ApiUpdateStreamsRequest {
	return self.Last()
}
func (self *ApiUpdateStreamsRequestStream) TailOr(arg ApiUpdateStreamsRequest) ApiUpdateStreamsRequest {
	return self.LastOr(arg)
}
func (self *ApiUpdateStreamsRequestStream) ToList() []ApiUpdateStreamsRequest {
	return self.Val()
}
func (self *ApiUpdateStreamsRequestStream) Unique() *ApiUpdateStreamsRequestStream {
	return self.Distinct()
}
func (self *ApiUpdateStreamsRequestStream) Val() []ApiUpdateStreamsRequest {
	if self == nil {
		return []ApiUpdateStreamsRequest{}
	}
	return *self.Copy()
}
func (self *ApiUpdateStreamsRequestStream) While(fn func(ApiUpdateStreamsRequest, int) bool) *ApiUpdateStreamsRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiUpdateStreamsRequestStream) Where(fn func(ApiUpdateStreamsRequest) bool) *ApiUpdateStreamsRequestStream {
	result := ApiUpdateStreamsRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiUpdateStreamsRequestStream) WhereSlim(fn func(ApiUpdateStreamsRequest) bool) *ApiUpdateStreamsRequestStream {
	result := ApiUpdateStreamsRequestStreamOf()
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
