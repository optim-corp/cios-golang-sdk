/*
 * Collection utility of ApiGetCollectionStatusRequest Struct
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

type ApiGetCollectionStatusRequestStream []ApiGetCollectionStatusRequest

func ApiGetCollectionStatusRequestStreamOf(arg ...ApiGetCollectionStatusRequest) ApiGetCollectionStatusRequestStream {
	return arg
}
func ApiGetCollectionStatusRequestStreamFrom(arg []ApiGetCollectionStatusRequest) ApiGetCollectionStatusRequestStream {
	return arg
}
func CreateApiGetCollectionStatusRequestStream(arg ...ApiGetCollectionStatusRequest) *ApiGetCollectionStatusRequestStream {
	tmp := ApiGetCollectionStatusRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetCollectionStatusRequestStream(arg []ApiGetCollectionStatusRequest) *ApiGetCollectionStatusRequestStream {
	tmp := ApiGetCollectionStatusRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetCollectionStatusRequestStream) Add(arg ApiGetCollectionStatusRequest) *ApiGetCollectionStatusRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetCollectionStatusRequestStream) AddAll(arg ...ApiGetCollectionStatusRequest) *ApiGetCollectionStatusRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetCollectionStatusRequestStream) AddSafe(arg *ApiGetCollectionStatusRequest) *ApiGetCollectionStatusRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetCollectionStatusRequestStream) Aggregate(fn func(ApiGetCollectionStatusRequest, ApiGetCollectionStatusRequest) ApiGetCollectionStatusRequest) *ApiGetCollectionStatusRequestStream {
	result := ApiGetCollectionStatusRequestStreamOf()
	self.ForEach(func(v ApiGetCollectionStatusRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetCollectionStatusRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetCollectionStatusRequestStream) AllMatch(fn func(ApiGetCollectionStatusRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetCollectionStatusRequestStream) AnyMatch(fn func(ApiGetCollectionStatusRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetCollectionStatusRequestStream) Clone() *ApiGetCollectionStatusRequestStream {
	temp := make([]ApiGetCollectionStatusRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetCollectionStatusRequestStream)(&temp)
}
func (self *ApiGetCollectionStatusRequestStream) Copy() *ApiGetCollectionStatusRequestStream {
	return self.Clone()
}
func (self *ApiGetCollectionStatusRequestStream) Concat(arg []ApiGetCollectionStatusRequest) *ApiGetCollectionStatusRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetCollectionStatusRequestStream) Contains(arg ApiGetCollectionStatusRequest) bool {
	return self.FindIndex(func(_arg ApiGetCollectionStatusRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetCollectionStatusRequestStream) Clean() *ApiGetCollectionStatusRequestStream {
	*self = ApiGetCollectionStatusRequestStreamOf()
	return self
}
func (self *ApiGetCollectionStatusRequestStream) Delete(index int) *ApiGetCollectionStatusRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetCollectionStatusRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetCollectionStatusRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetCollectionStatusRequestStream) Distinct() *ApiGetCollectionStatusRequestStream {
	caches := map[string]bool{}
	result := ApiGetCollectionStatusRequestStreamOf()
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
func (self *ApiGetCollectionStatusRequestStream) Each(fn func(ApiGetCollectionStatusRequest)) *ApiGetCollectionStatusRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetCollectionStatusRequestStream) EachRight(fn func(ApiGetCollectionStatusRequest)) *ApiGetCollectionStatusRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetCollectionStatusRequestStream) Equals(arr []ApiGetCollectionStatusRequest) bool {
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
func (self *ApiGetCollectionStatusRequestStream) Filter(fn func(ApiGetCollectionStatusRequest, int) bool) *ApiGetCollectionStatusRequestStream {
	result := ApiGetCollectionStatusRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetCollectionStatusRequestStream) FilterSlim(fn func(ApiGetCollectionStatusRequest, int) bool) *ApiGetCollectionStatusRequestStream {
	result := ApiGetCollectionStatusRequestStreamOf()
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
func (self *ApiGetCollectionStatusRequestStream) Find(fn func(ApiGetCollectionStatusRequest, int) bool) *ApiGetCollectionStatusRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetCollectionStatusRequestStream) FindOr(fn func(ApiGetCollectionStatusRequest, int) bool, or ApiGetCollectionStatusRequest) ApiGetCollectionStatusRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetCollectionStatusRequestStream) FindIndex(fn func(ApiGetCollectionStatusRequest, int) bool) int {
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
func (self *ApiGetCollectionStatusRequestStream) First() *ApiGetCollectionStatusRequest {
	return self.Get(0)
}
func (self *ApiGetCollectionStatusRequestStream) FirstOr(arg ApiGetCollectionStatusRequest) ApiGetCollectionStatusRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetCollectionStatusRequestStream) ForEach(fn func(ApiGetCollectionStatusRequest, int)) *ApiGetCollectionStatusRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetCollectionStatusRequestStream) ForEachRight(fn func(ApiGetCollectionStatusRequest, int)) *ApiGetCollectionStatusRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetCollectionStatusRequestStream) GroupBy(fn func(ApiGetCollectionStatusRequest, int) string) map[string][]ApiGetCollectionStatusRequest {
	m := map[string][]ApiGetCollectionStatusRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetCollectionStatusRequestStream) GroupByValues(fn func(ApiGetCollectionStatusRequest, int) string) [][]ApiGetCollectionStatusRequest {
	var tmp [][]ApiGetCollectionStatusRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetCollectionStatusRequestStream) IndexOf(arg ApiGetCollectionStatusRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetCollectionStatusRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetCollectionStatusRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetCollectionStatusRequestStream) Last() *ApiGetCollectionStatusRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetCollectionStatusRequestStream) LastOr(arg ApiGetCollectionStatusRequest) ApiGetCollectionStatusRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetCollectionStatusRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetCollectionStatusRequestStream) Limit(limit int) *ApiGetCollectionStatusRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetCollectionStatusRequestStream) Map(fn func(ApiGetCollectionStatusRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCollectionStatusRequestStream) Map2Int(fn func(ApiGetCollectionStatusRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCollectionStatusRequestStream) Map2Int32(fn func(ApiGetCollectionStatusRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCollectionStatusRequestStream) Map2Int64(fn func(ApiGetCollectionStatusRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCollectionStatusRequestStream) Map2Float32(fn func(ApiGetCollectionStatusRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCollectionStatusRequestStream) Map2Float64(fn func(ApiGetCollectionStatusRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCollectionStatusRequestStream) Map2Bool(fn func(ApiGetCollectionStatusRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCollectionStatusRequestStream) Map2Bytes(fn func(ApiGetCollectionStatusRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCollectionStatusRequestStream) Map2String(fn func(ApiGetCollectionStatusRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetCollectionStatusRequestStream) Max(fn func(ApiGetCollectionStatusRequest, int) float64) *ApiGetCollectionStatusRequest {
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
func (self *ApiGetCollectionStatusRequestStream) Min(fn func(ApiGetCollectionStatusRequest, int) float64) *ApiGetCollectionStatusRequest {
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
func (self *ApiGetCollectionStatusRequestStream) NoneMatch(fn func(ApiGetCollectionStatusRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetCollectionStatusRequestStream) Get(index int) *ApiGetCollectionStatusRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetCollectionStatusRequestStream) GetOr(index int, arg ApiGetCollectionStatusRequest) ApiGetCollectionStatusRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetCollectionStatusRequestStream) Peek(fn func(*ApiGetCollectionStatusRequest, int)) *ApiGetCollectionStatusRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetCollectionStatusRequestStream) Reduce(fn func(ApiGetCollectionStatusRequest, ApiGetCollectionStatusRequest, int) ApiGetCollectionStatusRequest) *ApiGetCollectionStatusRequestStream {
	return self.ReduceInit(fn, ApiGetCollectionStatusRequest{})
}
func (self *ApiGetCollectionStatusRequestStream) ReduceInit(fn func(ApiGetCollectionStatusRequest, ApiGetCollectionStatusRequest, int) ApiGetCollectionStatusRequest, initialValue ApiGetCollectionStatusRequest) *ApiGetCollectionStatusRequestStream {
	result := ApiGetCollectionStatusRequestStreamOf()
	self.ForEach(func(v ApiGetCollectionStatusRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetCollectionStatusRequestStream) ReduceInterface(fn func(interface{}, ApiGetCollectionStatusRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetCollectionStatusRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetCollectionStatusRequestStream) ReduceString(fn func(string, ApiGetCollectionStatusRequest, int) string) []string {
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
func (self *ApiGetCollectionStatusRequestStream) ReduceInt(fn func(int, ApiGetCollectionStatusRequest, int) int) []int {
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
func (self *ApiGetCollectionStatusRequestStream) ReduceInt32(fn func(int32, ApiGetCollectionStatusRequest, int) int32) []int32 {
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
func (self *ApiGetCollectionStatusRequestStream) ReduceInt64(fn func(int64, ApiGetCollectionStatusRequest, int) int64) []int64 {
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
func (self *ApiGetCollectionStatusRequestStream) ReduceFloat32(fn func(float32, ApiGetCollectionStatusRequest, int) float32) []float32 {
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
func (self *ApiGetCollectionStatusRequestStream) ReduceFloat64(fn func(float64, ApiGetCollectionStatusRequest, int) float64) []float64 {
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
func (self *ApiGetCollectionStatusRequestStream) ReduceBool(fn func(bool, ApiGetCollectionStatusRequest, int) bool) []bool {
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
func (self *ApiGetCollectionStatusRequestStream) Reverse() *ApiGetCollectionStatusRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetCollectionStatusRequestStream) Replace(fn func(ApiGetCollectionStatusRequest, int) ApiGetCollectionStatusRequest) *ApiGetCollectionStatusRequestStream {
	return self.ForEach(func(v ApiGetCollectionStatusRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetCollectionStatusRequestStream) Select(fn func(ApiGetCollectionStatusRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetCollectionStatusRequestStream) Set(index int, val ApiGetCollectionStatusRequest) *ApiGetCollectionStatusRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetCollectionStatusRequestStream) Skip(skip int) *ApiGetCollectionStatusRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetCollectionStatusRequestStream) SkippingEach(fn func(ApiGetCollectionStatusRequest, int) int) *ApiGetCollectionStatusRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetCollectionStatusRequestStream) Slice(startIndex, n int) *ApiGetCollectionStatusRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetCollectionStatusRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetCollectionStatusRequestStream) Sort(fn func(i, j int) bool) *ApiGetCollectionStatusRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetCollectionStatusRequestStream) Tail() *ApiGetCollectionStatusRequest {
	return self.Last()
}
func (self *ApiGetCollectionStatusRequestStream) TailOr(arg ApiGetCollectionStatusRequest) ApiGetCollectionStatusRequest {
	return self.LastOr(arg)
}
func (self *ApiGetCollectionStatusRequestStream) ToList() []ApiGetCollectionStatusRequest {
	return self.Val()
}
func (self *ApiGetCollectionStatusRequestStream) Unique() *ApiGetCollectionStatusRequestStream {
	return self.Distinct()
}
func (self *ApiGetCollectionStatusRequestStream) Val() []ApiGetCollectionStatusRequest {
	if self == nil {
		return []ApiGetCollectionStatusRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetCollectionStatusRequestStream) While(fn func(ApiGetCollectionStatusRequest, int) bool) *ApiGetCollectionStatusRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetCollectionStatusRequestStream) Where(fn func(ApiGetCollectionStatusRequest) bool) *ApiGetCollectionStatusRequestStream {
	result := ApiGetCollectionStatusRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetCollectionStatusRequestStream) WhereSlim(fn func(ApiGetCollectionStatusRequest) bool) *ApiGetCollectionStatusRequestStream {
	result := ApiGetCollectionStatusRequestStreamOf()
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