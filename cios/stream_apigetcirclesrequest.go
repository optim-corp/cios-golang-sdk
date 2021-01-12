/*
 * Collection utility of ApiGetCirclesRequest Struct
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

type ApiGetCirclesRequestStream []ApiGetCirclesRequest

func ApiGetCirclesRequestStreamOf(arg ...ApiGetCirclesRequest) ApiGetCirclesRequestStream {
	return arg
}
func ApiGetCirclesRequestStreamFrom(arg []ApiGetCirclesRequest) ApiGetCirclesRequestStream {
	return arg
}
func CreateApiGetCirclesRequestStream(arg ...ApiGetCirclesRequest) *ApiGetCirclesRequestStream {
	tmp := ApiGetCirclesRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetCirclesRequestStream(arg []ApiGetCirclesRequest) *ApiGetCirclesRequestStream {
	tmp := ApiGetCirclesRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetCirclesRequestStream) Add(arg ApiGetCirclesRequest) *ApiGetCirclesRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetCirclesRequestStream) AddAll(arg ...ApiGetCirclesRequest) *ApiGetCirclesRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetCirclesRequestStream) AddSafe(arg *ApiGetCirclesRequest) *ApiGetCirclesRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetCirclesRequestStream) Aggregate(fn func(ApiGetCirclesRequest, ApiGetCirclesRequest) ApiGetCirclesRequest) *ApiGetCirclesRequestStream {
	result := ApiGetCirclesRequestStreamOf()
	self.ForEach(func(v ApiGetCirclesRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetCirclesRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetCirclesRequestStream) AllMatch(fn func(ApiGetCirclesRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetCirclesRequestStream) AnyMatch(fn func(ApiGetCirclesRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetCirclesRequestStream) Clone() *ApiGetCirclesRequestStream {
	temp := make([]ApiGetCirclesRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetCirclesRequestStream)(&temp)
}
func (self *ApiGetCirclesRequestStream) Copy() *ApiGetCirclesRequestStream {
	return self.Clone()
}
func (self *ApiGetCirclesRequestStream) Concat(arg []ApiGetCirclesRequest) *ApiGetCirclesRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetCirclesRequestStream) Contains(arg ApiGetCirclesRequest) bool {
	return self.FindIndex(func(_arg ApiGetCirclesRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetCirclesRequestStream) Clean() *ApiGetCirclesRequestStream {
	*self = ApiGetCirclesRequestStreamOf()
	return self
}
func (self *ApiGetCirclesRequestStream) Delete(index int) *ApiGetCirclesRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetCirclesRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetCirclesRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetCirclesRequestStream) Distinct() *ApiGetCirclesRequestStream {
	caches := map[string]bool{}
	result := ApiGetCirclesRequestStreamOf()
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
func (self *ApiGetCirclesRequestStream) Each(fn func(ApiGetCirclesRequest)) *ApiGetCirclesRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetCirclesRequestStream) EachRight(fn func(ApiGetCirclesRequest)) *ApiGetCirclesRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetCirclesRequestStream) Equals(arr []ApiGetCirclesRequest) bool {
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
func (self *ApiGetCirclesRequestStream) Filter(fn func(ApiGetCirclesRequest, int) bool) *ApiGetCirclesRequestStream {
	result := ApiGetCirclesRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetCirclesRequestStream) FilterSlim(fn func(ApiGetCirclesRequest, int) bool) *ApiGetCirclesRequestStream {
	result := ApiGetCirclesRequestStreamOf()
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
func (self *ApiGetCirclesRequestStream) Find(fn func(ApiGetCirclesRequest, int) bool) *ApiGetCirclesRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetCirclesRequestStream) FindOr(fn func(ApiGetCirclesRequest, int) bool, or ApiGetCirclesRequest) ApiGetCirclesRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetCirclesRequestStream) FindIndex(fn func(ApiGetCirclesRequest, int) bool) int {
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
func (self *ApiGetCirclesRequestStream) First() *ApiGetCirclesRequest {
	return self.Get(0)
}
func (self *ApiGetCirclesRequestStream) FirstOr(arg ApiGetCirclesRequest) ApiGetCirclesRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetCirclesRequestStream) ForEach(fn func(ApiGetCirclesRequest, int)) *ApiGetCirclesRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetCirclesRequestStream) ForEachRight(fn func(ApiGetCirclesRequest, int)) *ApiGetCirclesRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetCirclesRequestStream) GroupBy(fn func(ApiGetCirclesRequest, int) string) map[string][]ApiGetCirclesRequest {
	m := map[string][]ApiGetCirclesRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetCirclesRequestStream) GroupByValues(fn func(ApiGetCirclesRequest, int) string) [][]ApiGetCirclesRequest {
	var tmp [][]ApiGetCirclesRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetCirclesRequestStream) IndexOf(arg ApiGetCirclesRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetCirclesRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetCirclesRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetCirclesRequestStream) Last() *ApiGetCirclesRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetCirclesRequestStream) LastOr(arg ApiGetCirclesRequest) ApiGetCirclesRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetCirclesRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetCirclesRequestStream) Limit(limit int) *ApiGetCirclesRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetCirclesRequestStream) Map(fn func(ApiGetCirclesRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCirclesRequestStream) Map2Int(fn func(ApiGetCirclesRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCirclesRequestStream) Map2Int32(fn func(ApiGetCirclesRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCirclesRequestStream) Map2Int64(fn func(ApiGetCirclesRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCirclesRequestStream) Map2Float32(fn func(ApiGetCirclesRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCirclesRequestStream) Map2Float64(fn func(ApiGetCirclesRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCirclesRequestStream) Map2Bool(fn func(ApiGetCirclesRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCirclesRequestStream) Map2Bytes(fn func(ApiGetCirclesRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCirclesRequestStream) Map2String(fn func(ApiGetCirclesRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCirclesRequestStream) Max(fn func(ApiGetCirclesRequest, int) float64) *ApiGetCirclesRequest {
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
func (self *ApiGetCirclesRequestStream) Min(fn func(ApiGetCirclesRequest, int) float64) *ApiGetCirclesRequest {
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
func (self *ApiGetCirclesRequestStream) NoneMatch(fn func(ApiGetCirclesRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetCirclesRequestStream) Get(index int) *ApiGetCirclesRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetCirclesRequestStream) GetOr(index int, arg ApiGetCirclesRequest) ApiGetCirclesRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetCirclesRequestStream) Peek(fn func(*ApiGetCirclesRequest, int)) *ApiGetCirclesRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetCirclesRequestStream) Reduce(fn func(ApiGetCirclesRequest, ApiGetCirclesRequest, int) ApiGetCirclesRequest) *ApiGetCirclesRequestStream {
	return self.ReduceInit(fn, ApiGetCirclesRequest{})
}
func (self *ApiGetCirclesRequestStream) ReduceInit(fn func(ApiGetCirclesRequest, ApiGetCirclesRequest, int) ApiGetCirclesRequest, initialValue ApiGetCirclesRequest) *ApiGetCirclesRequestStream {
	result := ApiGetCirclesRequestStreamOf()
	self.ForEach(func(v ApiGetCirclesRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetCirclesRequestStream) ReduceInterface(fn func(interface{}, ApiGetCirclesRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetCirclesRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetCirclesRequestStream) ReduceString(fn func(string, ApiGetCirclesRequest, int) string) []string {
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
func (self *ApiGetCirclesRequestStream) ReduceInt(fn func(int, ApiGetCirclesRequest, int) int) []int {
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
func (self *ApiGetCirclesRequestStream) ReduceInt32(fn func(int32, ApiGetCirclesRequest, int) int32) []int32 {
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
func (self *ApiGetCirclesRequestStream) ReduceInt64(fn func(int64, ApiGetCirclesRequest, int) int64) []int64 {
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
func (self *ApiGetCirclesRequestStream) ReduceFloat32(fn func(float32, ApiGetCirclesRequest, int) float32) []float32 {
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
func (self *ApiGetCirclesRequestStream) ReduceFloat64(fn func(float64, ApiGetCirclesRequest, int) float64) []float64 {
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
func (self *ApiGetCirclesRequestStream) ReduceBool(fn func(bool, ApiGetCirclesRequest, int) bool) []bool {
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
func (self *ApiGetCirclesRequestStream) Reverse() *ApiGetCirclesRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetCirclesRequestStream) Replace(fn func(ApiGetCirclesRequest, int) ApiGetCirclesRequest) *ApiGetCirclesRequestStream {
	return self.ForEach(func(v ApiGetCirclesRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetCirclesRequestStream) Select(fn func(ApiGetCirclesRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetCirclesRequestStream) Set(index int, val ApiGetCirclesRequest) *ApiGetCirclesRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetCirclesRequestStream) Skip(skip int) *ApiGetCirclesRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetCirclesRequestStream) SkippingEach(fn func(ApiGetCirclesRequest, int) int) *ApiGetCirclesRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetCirclesRequestStream) Slice(startIndex, n int) *ApiGetCirclesRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetCirclesRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetCirclesRequestStream) Sort(fn func(i, j int) bool) *ApiGetCirclesRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetCirclesRequestStream) Tail() *ApiGetCirclesRequest {
	return self.Last()
}
func (self *ApiGetCirclesRequestStream) TailOr(arg ApiGetCirclesRequest) ApiGetCirclesRequest {
	return self.LastOr(arg)
}
func (self *ApiGetCirclesRequestStream) ToList() []ApiGetCirclesRequest {
	return self.Val()
}
func (self *ApiGetCirclesRequestStream) Unique() *ApiGetCirclesRequestStream {
	return self.Distinct()
}
func (self *ApiGetCirclesRequestStream) Val() []ApiGetCirclesRequest {
	if self == nil {
		return []ApiGetCirclesRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetCirclesRequestStream) While(fn func(ApiGetCirclesRequest, int) bool) *ApiGetCirclesRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetCirclesRequestStream) Where(fn func(ApiGetCirclesRequest) bool) *ApiGetCirclesRequestStream {
	result := ApiGetCirclesRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetCirclesRequestStream) WhereSlim(fn func(ApiGetCirclesRequest) bool) *ApiGetCirclesRequestStream {
	result := ApiGetCirclesRequestStreamOf()
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