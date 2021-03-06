/*
 * Collection utility of ApiCreateSeriesImageRequest Struct
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

type ApiCreateSeriesImageRequestStream []ApiCreateSeriesImageRequest

func ApiCreateSeriesImageRequestStreamOf(arg ...ApiCreateSeriesImageRequest) ApiCreateSeriesImageRequestStream {
	return arg
}
func ApiCreateSeriesImageRequestStreamFrom(arg []ApiCreateSeriesImageRequest) ApiCreateSeriesImageRequestStream {
	return arg
}
func CreateApiCreateSeriesImageRequestStream(arg ...ApiCreateSeriesImageRequest) *ApiCreateSeriesImageRequestStream {
	tmp := ApiCreateSeriesImageRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiCreateSeriesImageRequestStream(arg []ApiCreateSeriesImageRequest) *ApiCreateSeriesImageRequestStream {
	tmp := ApiCreateSeriesImageRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiCreateSeriesImageRequestStream) Add(arg ApiCreateSeriesImageRequest) *ApiCreateSeriesImageRequestStream {
	return self.AddAll(arg)
}
func (self *ApiCreateSeriesImageRequestStream) AddAll(arg ...ApiCreateSeriesImageRequest) *ApiCreateSeriesImageRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiCreateSeriesImageRequestStream) AddSafe(arg *ApiCreateSeriesImageRequest) *ApiCreateSeriesImageRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiCreateSeriesImageRequestStream) Aggregate(fn func(ApiCreateSeriesImageRequest, ApiCreateSeriesImageRequest) ApiCreateSeriesImageRequest) *ApiCreateSeriesImageRequestStream {
	result := ApiCreateSeriesImageRequestStreamOf()
	self.ForEach(func(v ApiCreateSeriesImageRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiCreateSeriesImageRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiCreateSeriesImageRequestStream) AllMatch(fn func(ApiCreateSeriesImageRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiCreateSeriesImageRequestStream) AnyMatch(fn func(ApiCreateSeriesImageRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiCreateSeriesImageRequestStream) Clone() *ApiCreateSeriesImageRequestStream {
	temp := make([]ApiCreateSeriesImageRequest, self.Len())
	copy(temp, *self)
	return (*ApiCreateSeriesImageRequestStream)(&temp)
}
func (self *ApiCreateSeriesImageRequestStream) Copy() *ApiCreateSeriesImageRequestStream {
	return self.Clone()
}
func (self *ApiCreateSeriesImageRequestStream) Concat(arg []ApiCreateSeriesImageRequest) *ApiCreateSeriesImageRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiCreateSeriesImageRequestStream) Contains(arg ApiCreateSeriesImageRequest) bool {
	return self.FindIndex(func(_arg ApiCreateSeriesImageRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiCreateSeriesImageRequestStream) Clean() *ApiCreateSeriesImageRequestStream {
	*self = ApiCreateSeriesImageRequestStreamOf()
	return self
}
func (self *ApiCreateSeriesImageRequestStream) Delete(index int) *ApiCreateSeriesImageRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiCreateSeriesImageRequestStream) DeleteRange(startIndex, endIndex int) *ApiCreateSeriesImageRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiCreateSeriesImageRequestStream) Distinct() *ApiCreateSeriesImageRequestStream {
	caches := map[string]bool{}
	result := ApiCreateSeriesImageRequestStreamOf()
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
func (self *ApiCreateSeriesImageRequestStream) Each(fn func(ApiCreateSeriesImageRequest)) *ApiCreateSeriesImageRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiCreateSeriesImageRequestStream) EachRight(fn func(ApiCreateSeriesImageRequest)) *ApiCreateSeriesImageRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiCreateSeriesImageRequestStream) Equals(arr []ApiCreateSeriesImageRequest) bool {
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
func (self *ApiCreateSeriesImageRequestStream) Filter(fn func(ApiCreateSeriesImageRequest, int) bool) *ApiCreateSeriesImageRequestStream {
	result := ApiCreateSeriesImageRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiCreateSeriesImageRequestStream) FilterSlim(fn func(ApiCreateSeriesImageRequest, int) bool) *ApiCreateSeriesImageRequestStream {
	result := ApiCreateSeriesImageRequestStreamOf()
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
func (self *ApiCreateSeriesImageRequestStream) Find(fn func(ApiCreateSeriesImageRequest, int) bool) *ApiCreateSeriesImageRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiCreateSeriesImageRequestStream) FindOr(fn func(ApiCreateSeriesImageRequest, int) bool, or ApiCreateSeriesImageRequest) ApiCreateSeriesImageRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiCreateSeriesImageRequestStream) FindIndex(fn func(ApiCreateSeriesImageRequest, int) bool) int {
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
func (self *ApiCreateSeriesImageRequestStream) First() *ApiCreateSeriesImageRequest {
	return self.Get(0)
}
func (self *ApiCreateSeriesImageRequestStream) FirstOr(arg ApiCreateSeriesImageRequest) ApiCreateSeriesImageRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiCreateSeriesImageRequestStream) ForEach(fn func(ApiCreateSeriesImageRequest, int)) *ApiCreateSeriesImageRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiCreateSeriesImageRequestStream) ForEachRight(fn func(ApiCreateSeriesImageRequest, int)) *ApiCreateSeriesImageRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiCreateSeriesImageRequestStream) GroupBy(fn func(ApiCreateSeriesImageRequest, int) string) map[string][]ApiCreateSeriesImageRequest {
	m := map[string][]ApiCreateSeriesImageRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiCreateSeriesImageRequestStream) GroupByValues(fn func(ApiCreateSeriesImageRequest, int) string) [][]ApiCreateSeriesImageRequest {
	var tmp [][]ApiCreateSeriesImageRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiCreateSeriesImageRequestStream) IndexOf(arg ApiCreateSeriesImageRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiCreateSeriesImageRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiCreateSeriesImageRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiCreateSeriesImageRequestStream) Last() *ApiCreateSeriesImageRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiCreateSeriesImageRequestStream) LastOr(arg ApiCreateSeriesImageRequest) ApiCreateSeriesImageRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiCreateSeriesImageRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiCreateSeriesImageRequestStream) Limit(limit int) *ApiCreateSeriesImageRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiCreateSeriesImageRequestStream) Map(fn func(ApiCreateSeriesImageRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateSeriesImageRequestStream) Map2Int(fn func(ApiCreateSeriesImageRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateSeriesImageRequestStream) Map2Int32(fn func(ApiCreateSeriesImageRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateSeriesImageRequestStream) Map2Int64(fn func(ApiCreateSeriesImageRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateSeriesImageRequestStream) Map2Float32(fn func(ApiCreateSeriesImageRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateSeriesImageRequestStream) Map2Float64(fn func(ApiCreateSeriesImageRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateSeriesImageRequestStream) Map2Bool(fn func(ApiCreateSeriesImageRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateSeriesImageRequestStream) Map2Bytes(fn func(ApiCreateSeriesImageRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateSeriesImageRequestStream) Map2String(fn func(ApiCreateSeriesImageRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateSeriesImageRequestStream) Max(fn func(ApiCreateSeriesImageRequest, int) float64) *ApiCreateSeriesImageRequest {
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
func (self *ApiCreateSeriesImageRequestStream) Min(fn func(ApiCreateSeriesImageRequest, int) float64) *ApiCreateSeriesImageRequest {
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
func (self *ApiCreateSeriesImageRequestStream) NoneMatch(fn func(ApiCreateSeriesImageRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiCreateSeriesImageRequestStream) Get(index int) *ApiCreateSeriesImageRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiCreateSeriesImageRequestStream) GetOr(index int, arg ApiCreateSeriesImageRequest) ApiCreateSeriesImageRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiCreateSeriesImageRequestStream) Peek(fn func(*ApiCreateSeriesImageRequest, int)) *ApiCreateSeriesImageRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiCreateSeriesImageRequestStream) Reduce(fn func(ApiCreateSeriesImageRequest, ApiCreateSeriesImageRequest, int) ApiCreateSeriesImageRequest) *ApiCreateSeriesImageRequestStream {
	return self.ReduceInit(fn, ApiCreateSeriesImageRequest{})
}
func (self *ApiCreateSeriesImageRequestStream) ReduceInit(fn func(ApiCreateSeriesImageRequest, ApiCreateSeriesImageRequest, int) ApiCreateSeriesImageRequest, initialValue ApiCreateSeriesImageRequest) *ApiCreateSeriesImageRequestStream {
	result := ApiCreateSeriesImageRequestStreamOf()
	self.ForEach(func(v ApiCreateSeriesImageRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiCreateSeriesImageRequestStream) ReduceInterface(fn func(interface{}, ApiCreateSeriesImageRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiCreateSeriesImageRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiCreateSeriesImageRequestStream) ReduceString(fn func(string, ApiCreateSeriesImageRequest, int) string) []string {
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
func (self *ApiCreateSeriesImageRequestStream) ReduceInt(fn func(int, ApiCreateSeriesImageRequest, int) int) []int {
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
func (self *ApiCreateSeriesImageRequestStream) ReduceInt32(fn func(int32, ApiCreateSeriesImageRequest, int) int32) []int32 {
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
func (self *ApiCreateSeriesImageRequestStream) ReduceInt64(fn func(int64, ApiCreateSeriesImageRequest, int) int64) []int64 {
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
func (self *ApiCreateSeriesImageRequestStream) ReduceFloat32(fn func(float32, ApiCreateSeriesImageRequest, int) float32) []float32 {
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
func (self *ApiCreateSeriesImageRequestStream) ReduceFloat64(fn func(float64, ApiCreateSeriesImageRequest, int) float64) []float64 {
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
func (self *ApiCreateSeriesImageRequestStream) ReduceBool(fn func(bool, ApiCreateSeriesImageRequest, int) bool) []bool {
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
func (self *ApiCreateSeriesImageRequestStream) Reverse() *ApiCreateSeriesImageRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiCreateSeriesImageRequestStream) Replace(fn func(ApiCreateSeriesImageRequest, int) ApiCreateSeriesImageRequest) *ApiCreateSeriesImageRequestStream {
	return self.ForEach(func(v ApiCreateSeriesImageRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiCreateSeriesImageRequestStream) Select(fn func(ApiCreateSeriesImageRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiCreateSeriesImageRequestStream) Set(index int, val ApiCreateSeriesImageRequest) *ApiCreateSeriesImageRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiCreateSeriesImageRequestStream) Skip(skip int) *ApiCreateSeriesImageRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiCreateSeriesImageRequestStream) SkippingEach(fn func(ApiCreateSeriesImageRequest, int) int) *ApiCreateSeriesImageRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiCreateSeriesImageRequestStream) Slice(startIndex, n int) *ApiCreateSeriesImageRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiCreateSeriesImageRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiCreateSeriesImageRequestStream) Sort(fn func(i, j int) bool) *ApiCreateSeriesImageRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiCreateSeriesImageRequestStream) Tail() *ApiCreateSeriesImageRequest {
	return self.Last()
}
func (self *ApiCreateSeriesImageRequestStream) TailOr(arg ApiCreateSeriesImageRequest) ApiCreateSeriesImageRequest {
	return self.LastOr(arg)
}
func (self *ApiCreateSeriesImageRequestStream) ToList() []ApiCreateSeriesImageRequest {
	return self.Val()
}
func (self *ApiCreateSeriesImageRequestStream) Unique() *ApiCreateSeriesImageRequestStream {
	return self.Distinct()
}
func (self *ApiCreateSeriesImageRequestStream) Val() []ApiCreateSeriesImageRequest {
	if self == nil {
		return []ApiCreateSeriesImageRequest{}
	}
	return *self.Copy()
}
func (self *ApiCreateSeriesImageRequestStream) While(fn func(ApiCreateSeriesImageRequest, int) bool) *ApiCreateSeriesImageRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiCreateSeriesImageRequestStream) Where(fn func(ApiCreateSeriesImageRequest) bool) *ApiCreateSeriesImageRequestStream {
	result := ApiCreateSeriesImageRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiCreateSeriesImageRequestStream) WhereSlim(fn func(ApiCreateSeriesImageRequest) bool) *ApiCreateSeriesImageRequestStream {
	result := ApiCreateSeriesImageRequestStreamOf()
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
