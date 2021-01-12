/*
 * Collection utility of ApiDeleteSeriesLatestRequest Struct
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

type ApiDeleteSeriesLatestRequestStream []ApiDeleteSeriesLatestRequest

func ApiDeleteSeriesLatestRequestStreamOf(arg ...ApiDeleteSeriesLatestRequest) ApiDeleteSeriesLatestRequestStream {
	return arg
}
func ApiDeleteSeriesLatestRequestStreamFrom(arg []ApiDeleteSeriesLatestRequest) ApiDeleteSeriesLatestRequestStream {
	return arg
}
func CreateApiDeleteSeriesLatestRequestStream(arg ...ApiDeleteSeriesLatestRequest) *ApiDeleteSeriesLatestRequestStream {
	tmp := ApiDeleteSeriesLatestRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiDeleteSeriesLatestRequestStream(arg []ApiDeleteSeriesLatestRequest) *ApiDeleteSeriesLatestRequestStream {
	tmp := ApiDeleteSeriesLatestRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiDeleteSeriesLatestRequestStream) Add(arg ApiDeleteSeriesLatestRequest) *ApiDeleteSeriesLatestRequestStream {
	return self.AddAll(arg)
}
func (self *ApiDeleteSeriesLatestRequestStream) AddAll(arg ...ApiDeleteSeriesLatestRequest) *ApiDeleteSeriesLatestRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiDeleteSeriesLatestRequestStream) AddSafe(arg *ApiDeleteSeriesLatestRequest) *ApiDeleteSeriesLatestRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiDeleteSeriesLatestRequestStream) Aggregate(fn func(ApiDeleteSeriesLatestRequest, ApiDeleteSeriesLatestRequest) ApiDeleteSeriesLatestRequest) *ApiDeleteSeriesLatestRequestStream {
	result := ApiDeleteSeriesLatestRequestStreamOf()
	self.ForEach(func(v ApiDeleteSeriesLatestRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiDeleteSeriesLatestRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiDeleteSeriesLatestRequestStream) AllMatch(fn func(ApiDeleteSeriesLatestRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiDeleteSeriesLatestRequestStream) AnyMatch(fn func(ApiDeleteSeriesLatestRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiDeleteSeriesLatestRequestStream) Clone() *ApiDeleteSeriesLatestRequestStream {
	temp := make([]ApiDeleteSeriesLatestRequest, self.Len())
	copy(temp, *self)
	return (*ApiDeleteSeriesLatestRequestStream)(&temp)
}
func (self *ApiDeleteSeriesLatestRequestStream) Copy() *ApiDeleteSeriesLatestRequestStream {
	return self.Clone()
}
func (self *ApiDeleteSeriesLatestRequestStream) Concat(arg []ApiDeleteSeriesLatestRequest) *ApiDeleteSeriesLatestRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiDeleteSeriesLatestRequestStream) Contains(arg ApiDeleteSeriesLatestRequest) bool {
	return self.FindIndex(func(_arg ApiDeleteSeriesLatestRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiDeleteSeriesLatestRequestStream) Clean() *ApiDeleteSeriesLatestRequestStream {
	*self = ApiDeleteSeriesLatestRequestStreamOf()
	return self
}
func (self *ApiDeleteSeriesLatestRequestStream) Delete(index int) *ApiDeleteSeriesLatestRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiDeleteSeriesLatestRequestStream) DeleteRange(startIndex, endIndex int) *ApiDeleteSeriesLatestRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiDeleteSeriesLatestRequestStream) Distinct() *ApiDeleteSeriesLatestRequestStream {
	caches := map[string]bool{}
	result := ApiDeleteSeriesLatestRequestStreamOf()
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
func (self *ApiDeleteSeriesLatestRequestStream) Each(fn func(ApiDeleteSeriesLatestRequest)) *ApiDeleteSeriesLatestRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiDeleteSeriesLatestRequestStream) EachRight(fn func(ApiDeleteSeriesLatestRequest)) *ApiDeleteSeriesLatestRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiDeleteSeriesLatestRequestStream) Equals(arr []ApiDeleteSeriesLatestRequest) bool {
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
func (self *ApiDeleteSeriesLatestRequestStream) Filter(fn func(ApiDeleteSeriesLatestRequest, int) bool) *ApiDeleteSeriesLatestRequestStream {
	result := ApiDeleteSeriesLatestRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiDeleteSeriesLatestRequestStream) FilterSlim(fn func(ApiDeleteSeriesLatestRequest, int) bool) *ApiDeleteSeriesLatestRequestStream {
	result := ApiDeleteSeriesLatestRequestStreamOf()
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
func (self *ApiDeleteSeriesLatestRequestStream) Find(fn func(ApiDeleteSeriesLatestRequest, int) bool) *ApiDeleteSeriesLatestRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiDeleteSeriesLatestRequestStream) FindOr(fn func(ApiDeleteSeriesLatestRequest, int) bool, or ApiDeleteSeriesLatestRequest) ApiDeleteSeriesLatestRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiDeleteSeriesLatestRequestStream) FindIndex(fn func(ApiDeleteSeriesLatestRequest, int) bool) int {
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
func (self *ApiDeleteSeriesLatestRequestStream) First() *ApiDeleteSeriesLatestRequest {
	return self.Get(0)
}
func (self *ApiDeleteSeriesLatestRequestStream) FirstOr(arg ApiDeleteSeriesLatestRequest) ApiDeleteSeriesLatestRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiDeleteSeriesLatestRequestStream) ForEach(fn func(ApiDeleteSeriesLatestRequest, int)) *ApiDeleteSeriesLatestRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiDeleteSeriesLatestRequestStream) ForEachRight(fn func(ApiDeleteSeriesLatestRequest, int)) *ApiDeleteSeriesLatestRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiDeleteSeriesLatestRequestStream) GroupBy(fn func(ApiDeleteSeriesLatestRequest, int) string) map[string][]ApiDeleteSeriesLatestRequest {
	m := map[string][]ApiDeleteSeriesLatestRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiDeleteSeriesLatestRequestStream) GroupByValues(fn func(ApiDeleteSeriesLatestRequest, int) string) [][]ApiDeleteSeriesLatestRequest {
	var tmp [][]ApiDeleteSeriesLatestRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiDeleteSeriesLatestRequestStream) IndexOf(arg ApiDeleteSeriesLatestRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiDeleteSeriesLatestRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiDeleteSeriesLatestRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiDeleteSeriesLatestRequestStream) Last() *ApiDeleteSeriesLatestRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiDeleteSeriesLatestRequestStream) LastOr(arg ApiDeleteSeriesLatestRequest) ApiDeleteSeriesLatestRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiDeleteSeriesLatestRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiDeleteSeriesLatestRequestStream) Limit(limit int) *ApiDeleteSeriesLatestRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiDeleteSeriesLatestRequestStream) Map(fn func(ApiDeleteSeriesLatestRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteSeriesLatestRequestStream) Map2Int(fn func(ApiDeleteSeriesLatestRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteSeriesLatestRequestStream) Map2Int32(fn func(ApiDeleteSeriesLatestRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteSeriesLatestRequestStream) Map2Int64(fn func(ApiDeleteSeriesLatestRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteSeriesLatestRequestStream) Map2Float32(fn func(ApiDeleteSeriesLatestRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteSeriesLatestRequestStream) Map2Float64(fn func(ApiDeleteSeriesLatestRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteSeriesLatestRequestStream) Map2Bool(fn func(ApiDeleteSeriesLatestRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteSeriesLatestRequestStream) Map2Bytes(fn func(ApiDeleteSeriesLatestRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteSeriesLatestRequestStream) Map2String(fn func(ApiDeleteSeriesLatestRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiDeleteSeriesLatestRequestStream) Max(fn func(ApiDeleteSeriesLatestRequest, int) float64) *ApiDeleteSeriesLatestRequest {
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
func (self *ApiDeleteSeriesLatestRequestStream) Min(fn func(ApiDeleteSeriesLatestRequest, int) float64) *ApiDeleteSeriesLatestRequest {
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
func (self *ApiDeleteSeriesLatestRequestStream) NoneMatch(fn func(ApiDeleteSeriesLatestRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiDeleteSeriesLatestRequestStream) Get(index int) *ApiDeleteSeriesLatestRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiDeleteSeriesLatestRequestStream) GetOr(index int, arg ApiDeleteSeriesLatestRequest) ApiDeleteSeriesLatestRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiDeleteSeriesLatestRequestStream) Peek(fn func(*ApiDeleteSeriesLatestRequest, int)) *ApiDeleteSeriesLatestRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiDeleteSeriesLatestRequestStream) Reduce(fn func(ApiDeleteSeriesLatestRequest, ApiDeleteSeriesLatestRequest, int) ApiDeleteSeriesLatestRequest) *ApiDeleteSeriesLatestRequestStream {
	return self.ReduceInit(fn, ApiDeleteSeriesLatestRequest{})
}
func (self *ApiDeleteSeriesLatestRequestStream) ReduceInit(fn func(ApiDeleteSeriesLatestRequest, ApiDeleteSeriesLatestRequest, int) ApiDeleteSeriesLatestRequest, initialValue ApiDeleteSeriesLatestRequest) *ApiDeleteSeriesLatestRequestStream {
	result := ApiDeleteSeriesLatestRequestStreamOf()
	self.ForEach(func(v ApiDeleteSeriesLatestRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiDeleteSeriesLatestRequestStream) ReduceInterface(fn func(interface{}, ApiDeleteSeriesLatestRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiDeleteSeriesLatestRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiDeleteSeriesLatestRequestStream) ReduceString(fn func(string, ApiDeleteSeriesLatestRequest, int) string) []string {
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
func (self *ApiDeleteSeriesLatestRequestStream) ReduceInt(fn func(int, ApiDeleteSeriesLatestRequest, int) int) []int {
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
func (self *ApiDeleteSeriesLatestRequestStream) ReduceInt32(fn func(int32, ApiDeleteSeriesLatestRequest, int) int32) []int32 {
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
func (self *ApiDeleteSeriesLatestRequestStream) ReduceInt64(fn func(int64, ApiDeleteSeriesLatestRequest, int) int64) []int64 {
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
func (self *ApiDeleteSeriesLatestRequestStream) ReduceFloat32(fn func(float32, ApiDeleteSeriesLatestRequest, int) float32) []float32 {
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
func (self *ApiDeleteSeriesLatestRequestStream) ReduceFloat64(fn func(float64, ApiDeleteSeriesLatestRequest, int) float64) []float64 {
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
func (self *ApiDeleteSeriesLatestRequestStream) ReduceBool(fn func(bool, ApiDeleteSeriesLatestRequest, int) bool) []bool {
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
func (self *ApiDeleteSeriesLatestRequestStream) Reverse() *ApiDeleteSeriesLatestRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiDeleteSeriesLatestRequestStream) Replace(fn func(ApiDeleteSeriesLatestRequest, int) ApiDeleteSeriesLatestRequest) *ApiDeleteSeriesLatestRequestStream {
	return self.ForEach(func(v ApiDeleteSeriesLatestRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiDeleteSeriesLatestRequestStream) Select(fn func(ApiDeleteSeriesLatestRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiDeleteSeriesLatestRequestStream) Set(index int, val ApiDeleteSeriesLatestRequest) *ApiDeleteSeriesLatestRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiDeleteSeriesLatestRequestStream) Skip(skip int) *ApiDeleteSeriesLatestRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiDeleteSeriesLatestRequestStream) SkippingEach(fn func(ApiDeleteSeriesLatestRequest, int) int) *ApiDeleteSeriesLatestRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiDeleteSeriesLatestRequestStream) Slice(startIndex, n int) *ApiDeleteSeriesLatestRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiDeleteSeriesLatestRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiDeleteSeriesLatestRequestStream) Sort(fn func(i, j int) bool) *ApiDeleteSeriesLatestRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiDeleteSeriesLatestRequestStream) Tail() *ApiDeleteSeriesLatestRequest {
	return self.Last()
}
func (self *ApiDeleteSeriesLatestRequestStream) TailOr(arg ApiDeleteSeriesLatestRequest) ApiDeleteSeriesLatestRequest {
	return self.LastOr(arg)
}
func (self *ApiDeleteSeriesLatestRequestStream) ToList() []ApiDeleteSeriesLatestRequest {
	return self.Val()
}
func (self *ApiDeleteSeriesLatestRequestStream) Unique() *ApiDeleteSeriesLatestRequestStream {
	return self.Distinct()
}
func (self *ApiDeleteSeriesLatestRequestStream) Val() []ApiDeleteSeriesLatestRequest {
	if self == nil {
		return []ApiDeleteSeriesLatestRequest{}
	}
	return *self.Copy()
}
func (self *ApiDeleteSeriesLatestRequestStream) While(fn func(ApiDeleteSeriesLatestRequest, int) bool) *ApiDeleteSeriesLatestRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiDeleteSeriesLatestRequestStream) Where(fn func(ApiDeleteSeriesLatestRequest) bool) *ApiDeleteSeriesLatestRequestStream {
	result := ApiDeleteSeriesLatestRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiDeleteSeriesLatestRequestStream) WhereSlim(fn func(ApiDeleteSeriesLatestRequest) bool) *ApiDeleteSeriesLatestRequestStream {
	result := ApiDeleteSeriesLatestRequestStreamOf()
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
