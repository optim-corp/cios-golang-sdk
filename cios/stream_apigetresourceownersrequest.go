/*
 * Collection utility of ApiGetResourceOwnersRequest Struct
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

type ApiGetResourceOwnersRequestStream []ApiGetResourceOwnersRequest

func ApiGetResourceOwnersRequestStreamOf(arg ...ApiGetResourceOwnersRequest) ApiGetResourceOwnersRequestStream {
	return arg
}
func ApiGetResourceOwnersRequestStreamFrom(arg []ApiGetResourceOwnersRequest) ApiGetResourceOwnersRequestStream {
	return arg
}
func CreateApiGetResourceOwnersRequestStream(arg ...ApiGetResourceOwnersRequest) *ApiGetResourceOwnersRequestStream {
	tmp := ApiGetResourceOwnersRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetResourceOwnersRequestStream(arg []ApiGetResourceOwnersRequest) *ApiGetResourceOwnersRequestStream {
	tmp := ApiGetResourceOwnersRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetResourceOwnersRequestStream) Add(arg ApiGetResourceOwnersRequest) *ApiGetResourceOwnersRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetResourceOwnersRequestStream) AddAll(arg ...ApiGetResourceOwnersRequest) *ApiGetResourceOwnersRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetResourceOwnersRequestStream) AddSafe(arg *ApiGetResourceOwnersRequest) *ApiGetResourceOwnersRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetResourceOwnersRequestStream) Aggregate(fn func(ApiGetResourceOwnersRequest, ApiGetResourceOwnersRequest) ApiGetResourceOwnersRequest) *ApiGetResourceOwnersRequestStream {
	result := ApiGetResourceOwnersRequestStreamOf()
	self.ForEach(func(v ApiGetResourceOwnersRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetResourceOwnersRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetResourceOwnersRequestStream) AllMatch(fn func(ApiGetResourceOwnersRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetResourceOwnersRequestStream) AnyMatch(fn func(ApiGetResourceOwnersRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetResourceOwnersRequestStream) Clone() *ApiGetResourceOwnersRequestStream {
	temp := make([]ApiGetResourceOwnersRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetResourceOwnersRequestStream)(&temp)
}
func (self *ApiGetResourceOwnersRequestStream) Copy() *ApiGetResourceOwnersRequestStream {
	return self.Clone()
}
func (self *ApiGetResourceOwnersRequestStream) Concat(arg []ApiGetResourceOwnersRequest) *ApiGetResourceOwnersRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetResourceOwnersRequestStream) Contains(arg ApiGetResourceOwnersRequest) bool {
	return self.FindIndex(func(_arg ApiGetResourceOwnersRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetResourceOwnersRequestStream) Clean() *ApiGetResourceOwnersRequestStream {
	*self = ApiGetResourceOwnersRequestStreamOf()
	return self
}
func (self *ApiGetResourceOwnersRequestStream) Delete(index int) *ApiGetResourceOwnersRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetResourceOwnersRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetResourceOwnersRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetResourceOwnersRequestStream) Distinct() *ApiGetResourceOwnersRequestStream {
	caches := map[string]bool{}
	result := ApiGetResourceOwnersRequestStreamOf()
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
func (self *ApiGetResourceOwnersRequestStream) Each(fn func(ApiGetResourceOwnersRequest)) *ApiGetResourceOwnersRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetResourceOwnersRequestStream) EachRight(fn func(ApiGetResourceOwnersRequest)) *ApiGetResourceOwnersRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetResourceOwnersRequestStream) Equals(arr []ApiGetResourceOwnersRequest) bool {
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
func (self *ApiGetResourceOwnersRequestStream) Filter(fn func(ApiGetResourceOwnersRequest, int) bool) *ApiGetResourceOwnersRequestStream {
	result := ApiGetResourceOwnersRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetResourceOwnersRequestStream) FilterSlim(fn func(ApiGetResourceOwnersRequest, int) bool) *ApiGetResourceOwnersRequestStream {
	result := ApiGetResourceOwnersRequestStreamOf()
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
func (self *ApiGetResourceOwnersRequestStream) Find(fn func(ApiGetResourceOwnersRequest, int) bool) *ApiGetResourceOwnersRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetResourceOwnersRequestStream) FindOr(fn func(ApiGetResourceOwnersRequest, int) bool, or ApiGetResourceOwnersRequest) ApiGetResourceOwnersRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetResourceOwnersRequestStream) FindIndex(fn func(ApiGetResourceOwnersRequest, int) bool) int {
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
func (self *ApiGetResourceOwnersRequestStream) First() *ApiGetResourceOwnersRequest {
	return self.Get(0)
}
func (self *ApiGetResourceOwnersRequestStream) FirstOr(arg ApiGetResourceOwnersRequest) ApiGetResourceOwnersRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetResourceOwnersRequestStream) ForEach(fn func(ApiGetResourceOwnersRequest, int)) *ApiGetResourceOwnersRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetResourceOwnersRequestStream) ForEachRight(fn func(ApiGetResourceOwnersRequest, int)) *ApiGetResourceOwnersRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetResourceOwnersRequestStream) GroupBy(fn func(ApiGetResourceOwnersRequest, int) string) map[string][]ApiGetResourceOwnersRequest {
	m := map[string][]ApiGetResourceOwnersRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetResourceOwnersRequestStream) GroupByValues(fn func(ApiGetResourceOwnersRequest, int) string) [][]ApiGetResourceOwnersRequest {
	var tmp [][]ApiGetResourceOwnersRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetResourceOwnersRequestStream) IndexOf(arg ApiGetResourceOwnersRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetResourceOwnersRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetResourceOwnersRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetResourceOwnersRequestStream) Last() *ApiGetResourceOwnersRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetResourceOwnersRequestStream) LastOr(arg ApiGetResourceOwnersRequest) ApiGetResourceOwnersRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetResourceOwnersRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetResourceOwnersRequestStream) Limit(limit int) *ApiGetResourceOwnersRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetResourceOwnersRequestStream) Map(fn func(ApiGetResourceOwnersRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetResourceOwnersRequestStream) Map2Int(fn func(ApiGetResourceOwnersRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetResourceOwnersRequestStream) Map2Int32(fn func(ApiGetResourceOwnersRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetResourceOwnersRequestStream) Map2Int64(fn func(ApiGetResourceOwnersRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetResourceOwnersRequestStream) Map2Float32(fn func(ApiGetResourceOwnersRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetResourceOwnersRequestStream) Map2Float64(fn func(ApiGetResourceOwnersRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetResourceOwnersRequestStream) Map2Bool(fn func(ApiGetResourceOwnersRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetResourceOwnersRequestStream) Map2Bytes(fn func(ApiGetResourceOwnersRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetResourceOwnersRequestStream) Map2String(fn func(ApiGetResourceOwnersRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetResourceOwnersRequestStream) Max(fn func(ApiGetResourceOwnersRequest, int) float64) *ApiGetResourceOwnersRequest {
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
func (self *ApiGetResourceOwnersRequestStream) Min(fn func(ApiGetResourceOwnersRequest, int) float64) *ApiGetResourceOwnersRequest {
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
func (self *ApiGetResourceOwnersRequestStream) NoneMatch(fn func(ApiGetResourceOwnersRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetResourceOwnersRequestStream) Get(index int) *ApiGetResourceOwnersRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetResourceOwnersRequestStream) GetOr(index int, arg ApiGetResourceOwnersRequest) ApiGetResourceOwnersRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetResourceOwnersRequestStream) Peek(fn func(*ApiGetResourceOwnersRequest, int)) *ApiGetResourceOwnersRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetResourceOwnersRequestStream) Reduce(fn func(ApiGetResourceOwnersRequest, ApiGetResourceOwnersRequest, int) ApiGetResourceOwnersRequest) *ApiGetResourceOwnersRequestStream {
	return self.ReduceInit(fn, ApiGetResourceOwnersRequest{})
}
func (self *ApiGetResourceOwnersRequestStream) ReduceInit(fn func(ApiGetResourceOwnersRequest, ApiGetResourceOwnersRequest, int) ApiGetResourceOwnersRequest, initialValue ApiGetResourceOwnersRequest) *ApiGetResourceOwnersRequestStream {
	result := ApiGetResourceOwnersRequestStreamOf()
	self.ForEach(func(v ApiGetResourceOwnersRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetResourceOwnersRequestStream) ReduceInterface(fn func(interface{}, ApiGetResourceOwnersRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetResourceOwnersRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetResourceOwnersRequestStream) ReduceString(fn func(string, ApiGetResourceOwnersRequest, int) string) []string {
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
func (self *ApiGetResourceOwnersRequestStream) ReduceInt(fn func(int, ApiGetResourceOwnersRequest, int) int) []int {
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
func (self *ApiGetResourceOwnersRequestStream) ReduceInt32(fn func(int32, ApiGetResourceOwnersRequest, int) int32) []int32 {
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
func (self *ApiGetResourceOwnersRequestStream) ReduceInt64(fn func(int64, ApiGetResourceOwnersRequest, int) int64) []int64 {
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
func (self *ApiGetResourceOwnersRequestStream) ReduceFloat32(fn func(float32, ApiGetResourceOwnersRequest, int) float32) []float32 {
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
func (self *ApiGetResourceOwnersRequestStream) ReduceFloat64(fn func(float64, ApiGetResourceOwnersRequest, int) float64) []float64 {
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
func (self *ApiGetResourceOwnersRequestStream) ReduceBool(fn func(bool, ApiGetResourceOwnersRequest, int) bool) []bool {
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
func (self *ApiGetResourceOwnersRequestStream) Reverse() *ApiGetResourceOwnersRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetResourceOwnersRequestStream) Replace(fn func(ApiGetResourceOwnersRequest, int) ApiGetResourceOwnersRequest) *ApiGetResourceOwnersRequestStream {
	return self.ForEach(func(v ApiGetResourceOwnersRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetResourceOwnersRequestStream) Select(fn func(ApiGetResourceOwnersRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetResourceOwnersRequestStream) Set(index int, val ApiGetResourceOwnersRequest) *ApiGetResourceOwnersRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetResourceOwnersRequestStream) Skip(skip int) *ApiGetResourceOwnersRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetResourceOwnersRequestStream) SkippingEach(fn func(ApiGetResourceOwnersRequest, int) int) *ApiGetResourceOwnersRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetResourceOwnersRequestStream) Slice(startIndex, n int) *ApiGetResourceOwnersRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetResourceOwnersRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetResourceOwnersRequestStream) Sort(fn func(i, j int) bool) *ApiGetResourceOwnersRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetResourceOwnersRequestStream) Tail() *ApiGetResourceOwnersRequest {
	return self.Last()
}
func (self *ApiGetResourceOwnersRequestStream) TailOr(arg ApiGetResourceOwnersRequest) ApiGetResourceOwnersRequest {
	return self.LastOr(arg)
}
func (self *ApiGetResourceOwnersRequestStream) ToList() []ApiGetResourceOwnersRequest {
	return self.Val()
}
func (self *ApiGetResourceOwnersRequestStream) Unique() *ApiGetResourceOwnersRequestStream {
	return self.Distinct()
}
func (self *ApiGetResourceOwnersRequestStream) Val() []ApiGetResourceOwnersRequest {
	if self == nil {
		return []ApiGetResourceOwnersRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetResourceOwnersRequestStream) While(fn func(ApiGetResourceOwnersRequest, int) bool) *ApiGetResourceOwnersRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetResourceOwnersRequestStream) Where(fn func(ApiGetResourceOwnersRequest) bool) *ApiGetResourceOwnersRequestStream {
	result := ApiGetResourceOwnersRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetResourceOwnersRequestStream) WhereSlim(fn func(ApiGetResourceOwnersRequest) bool) *ApiGetResourceOwnersRequestStream {
	result := ApiGetResourceOwnersRequestStreamOf()
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
