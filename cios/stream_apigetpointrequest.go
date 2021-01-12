/*
 * Collection utility of ApiGetPointRequest Struct
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

type ApiGetPointRequestStream []ApiGetPointRequest

func ApiGetPointRequestStreamOf(arg ...ApiGetPointRequest) ApiGetPointRequestStream {
	return arg
}
func ApiGetPointRequestStreamFrom(arg []ApiGetPointRequest) ApiGetPointRequestStream {
	return arg
}
func CreateApiGetPointRequestStream(arg ...ApiGetPointRequest) *ApiGetPointRequestStream {
	tmp := ApiGetPointRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetPointRequestStream(arg []ApiGetPointRequest) *ApiGetPointRequestStream {
	tmp := ApiGetPointRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetPointRequestStream) Add(arg ApiGetPointRequest) *ApiGetPointRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetPointRequestStream) AddAll(arg ...ApiGetPointRequest) *ApiGetPointRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetPointRequestStream) AddSafe(arg *ApiGetPointRequest) *ApiGetPointRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetPointRequestStream) Aggregate(fn func(ApiGetPointRequest, ApiGetPointRequest) ApiGetPointRequest) *ApiGetPointRequestStream {
	result := ApiGetPointRequestStreamOf()
	self.ForEach(func(v ApiGetPointRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetPointRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetPointRequestStream) AllMatch(fn func(ApiGetPointRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetPointRequestStream) AnyMatch(fn func(ApiGetPointRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetPointRequestStream) Clone() *ApiGetPointRequestStream {
	temp := make([]ApiGetPointRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetPointRequestStream)(&temp)
}
func (self *ApiGetPointRequestStream) Copy() *ApiGetPointRequestStream {
	return self.Clone()
}
func (self *ApiGetPointRequestStream) Concat(arg []ApiGetPointRequest) *ApiGetPointRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetPointRequestStream) Contains(arg ApiGetPointRequest) bool {
	return self.FindIndex(func(_arg ApiGetPointRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetPointRequestStream) Clean() *ApiGetPointRequestStream {
	*self = ApiGetPointRequestStreamOf()
	return self
}
func (self *ApiGetPointRequestStream) Delete(index int) *ApiGetPointRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetPointRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetPointRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetPointRequestStream) Distinct() *ApiGetPointRequestStream {
	caches := map[string]bool{}
	result := ApiGetPointRequestStreamOf()
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
func (self *ApiGetPointRequestStream) Each(fn func(ApiGetPointRequest)) *ApiGetPointRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetPointRequestStream) EachRight(fn func(ApiGetPointRequest)) *ApiGetPointRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetPointRequestStream) Equals(arr []ApiGetPointRequest) bool {
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
func (self *ApiGetPointRequestStream) Filter(fn func(ApiGetPointRequest, int) bool) *ApiGetPointRequestStream {
	result := ApiGetPointRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetPointRequestStream) FilterSlim(fn func(ApiGetPointRequest, int) bool) *ApiGetPointRequestStream {
	result := ApiGetPointRequestStreamOf()
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
func (self *ApiGetPointRequestStream) Find(fn func(ApiGetPointRequest, int) bool) *ApiGetPointRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetPointRequestStream) FindOr(fn func(ApiGetPointRequest, int) bool, or ApiGetPointRequest) ApiGetPointRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetPointRequestStream) FindIndex(fn func(ApiGetPointRequest, int) bool) int {
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
func (self *ApiGetPointRequestStream) First() *ApiGetPointRequest {
	return self.Get(0)
}
func (self *ApiGetPointRequestStream) FirstOr(arg ApiGetPointRequest) ApiGetPointRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetPointRequestStream) ForEach(fn func(ApiGetPointRequest, int)) *ApiGetPointRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetPointRequestStream) ForEachRight(fn func(ApiGetPointRequest, int)) *ApiGetPointRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetPointRequestStream) GroupBy(fn func(ApiGetPointRequest, int) string) map[string][]ApiGetPointRequest {
	m := map[string][]ApiGetPointRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetPointRequestStream) GroupByValues(fn func(ApiGetPointRequest, int) string) [][]ApiGetPointRequest {
	var tmp [][]ApiGetPointRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetPointRequestStream) IndexOf(arg ApiGetPointRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetPointRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetPointRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetPointRequestStream) Last() *ApiGetPointRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetPointRequestStream) LastOr(arg ApiGetPointRequest) ApiGetPointRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetPointRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetPointRequestStream) Limit(limit int) *ApiGetPointRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetPointRequestStream) Map(fn func(ApiGetPointRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetPointRequestStream) Map2Int(fn func(ApiGetPointRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetPointRequestStream) Map2Int32(fn func(ApiGetPointRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetPointRequestStream) Map2Int64(fn func(ApiGetPointRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetPointRequestStream) Map2Float32(fn func(ApiGetPointRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetPointRequestStream) Map2Float64(fn func(ApiGetPointRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetPointRequestStream) Map2Bool(fn func(ApiGetPointRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetPointRequestStream) Map2Bytes(fn func(ApiGetPointRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetPointRequestStream) Map2String(fn func(ApiGetPointRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetPointRequestStream) Max(fn func(ApiGetPointRequest, int) float64) *ApiGetPointRequest {
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
func (self *ApiGetPointRequestStream) Min(fn func(ApiGetPointRequest, int) float64) *ApiGetPointRequest {
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
func (self *ApiGetPointRequestStream) NoneMatch(fn func(ApiGetPointRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetPointRequestStream) Get(index int) *ApiGetPointRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetPointRequestStream) GetOr(index int, arg ApiGetPointRequest) ApiGetPointRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetPointRequestStream) Peek(fn func(*ApiGetPointRequest, int)) *ApiGetPointRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetPointRequestStream) Reduce(fn func(ApiGetPointRequest, ApiGetPointRequest, int) ApiGetPointRequest) *ApiGetPointRequestStream {
	return self.ReduceInit(fn, ApiGetPointRequest{})
}
func (self *ApiGetPointRequestStream) ReduceInit(fn func(ApiGetPointRequest, ApiGetPointRequest, int) ApiGetPointRequest, initialValue ApiGetPointRequest) *ApiGetPointRequestStream {
	result := ApiGetPointRequestStreamOf()
	self.ForEach(func(v ApiGetPointRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetPointRequestStream) ReduceInterface(fn func(interface{}, ApiGetPointRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetPointRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetPointRequestStream) ReduceString(fn func(string, ApiGetPointRequest, int) string) []string {
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
func (self *ApiGetPointRequestStream) ReduceInt(fn func(int, ApiGetPointRequest, int) int) []int {
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
func (self *ApiGetPointRequestStream) ReduceInt32(fn func(int32, ApiGetPointRequest, int) int32) []int32 {
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
func (self *ApiGetPointRequestStream) ReduceInt64(fn func(int64, ApiGetPointRequest, int) int64) []int64 {
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
func (self *ApiGetPointRequestStream) ReduceFloat32(fn func(float32, ApiGetPointRequest, int) float32) []float32 {
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
func (self *ApiGetPointRequestStream) ReduceFloat64(fn func(float64, ApiGetPointRequest, int) float64) []float64 {
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
func (self *ApiGetPointRequestStream) ReduceBool(fn func(bool, ApiGetPointRequest, int) bool) []bool {
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
func (self *ApiGetPointRequestStream) Reverse() *ApiGetPointRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetPointRequestStream) Replace(fn func(ApiGetPointRequest, int) ApiGetPointRequest) *ApiGetPointRequestStream {
	return self.ForEach(func(v ApiGetPointRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetPointRequestStream) Select(fn func(ApiGetPointRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetPointRequestStream) Set(index int, val ApiGetPointRequest) *ApiGetPointRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetPointRequestStream) Skip(skip int) *ApiGetPointRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetPointRequestStream) SkippingEach(fn func(ApiGetPointRequest, int) int) *ApiGetPointRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetPointRequestStream) Slice(startIndex, n int) *ApiGetPointRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetPointRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetPointRequestStream) Sort(fn func(i, j int) bool) *ApiGetPointRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetPointRequestStream) Tail() *ApiGetPointRequest {
	return self.Last()
}
func (self *ApiGetPointRequestStream) TailOr(arg ApiGetPointRequest) ApiGetPointRequest {
	return self.LastOr(arg)
}
func (self *ApiGetPointRequestStream) ToList() []ApiGetPointRequest {
	return self.Val()
}
func (self *ApiGetPointRequestStream) Unique() *ApiGetPointRequestStream {
	return self.Distinct()
}
func (self *ApiGetPointRequestStream) Val() []ApiGetPointRequest {
	if self == nil {
		return []ApiGetPointRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetPointRequestStream) While(fn func(ApiGetPointRequest, int) bool) *ApiGetPointRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetPointRequestStream) Where(fn func(ApiGetPointRequest) bool) *ApiGetPointRequestStream {
	result := ApiGetPointRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetPointRequestStream) WhereSlim(fn func(ApiGetPointRequest) bool) *ApiGetPointRequestStream {
	result := ApiGetPointRequestStreamOf()
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
