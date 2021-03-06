/*
 * Collection utility of ApiGetDataStoreChannelRequest Struct
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

type ApiGetDataStoreChannelRequestStream []ApiGetDataStoreChannelRequest

func ApiGetDataStoreChannelRequestStreamOf(arg ...ApiGetDataStoreChannelRequest) ApiGetDataStoreChannelRequestStream {
	return arg
}
func ApiGetDataStoreChannelRequestStreamFrom(arg []ApiGetDataStoreChannelRequest) ApiGetDataStoreChannelRequestStream {
	return arg
}
func CreateApiGetDataStoreChannelRequestStream(arg ...ApiGetDataStoreChannelRequest) *ApiGetDataStoreChannelRequestStream {
	tmp := ApiGetDataStoreChannelRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetDataStoreChannelRequestStream(arg []ApiGetDataStoreChannelRequest) *ApiGetDataStoreChannelRequestStream {
	tmp := ApiGetDataStoreChannelRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetDataStoreChannelRequestStream) Add(arg ApiGetDataStoreChannelRequest) *ApiGetDataStoreChannelRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetDataStoreChannelRequestStream) AddAll(arg ...ApiGetDataStoreChannelRequest) *ApiGetDataStoreChannelRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetDataStoreChannelRequestStream) AddSafe(arg *ApiGetDataStoreChannelRequest) *ApiGetDataStoreChannelRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetDataStoreChannelRequestStream) Aggregate(fn func(ApiGetDataStoreChannelRequest, ApiGetDataStoreChannelRequest) ApiGetDataStoreChannelRequest) *ApiGetDataStoreChannelRequestStream {
	result := ApiGetDataStoreChannelRequestStreamOf()
	self.ForEach(func(v ApiGetDataStoreChannelRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetDataStoreChannelRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDataStoreChannelRequestStream) AllMatch(fn func(ApiGetDataStoreChannelRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetDataStoreChannelRequestStream) AnyMatch(fn func(ApiGetDataStoreChannelRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetDataStoreChannelRequestStream) Clone() *ApiGetDataStoreChannelRequestStream {
	temp := make([]ApiGetDataStoreChannelRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetDataStoreChannelRequestStream)(&temp)
}
func (self *ApiGetDataStoreChannelRequestStream) Copy() *ApiGetDataStoreChannelRequestStream {
	return self.Clone()
}
func (self *ApiGetDataStoreChannelRequestStream) Concat(arg []ApiGetDataStoreChannelRequest) *ApiGetDataStoreChannelRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetDataStoreChannelRequestStream) Contains(arg ApiGetDataStoreChannelRequest) bool {
	return self.FindIndex(func(_arg ApiGetDataStoreChannelRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetDataStoreChannelRequestStream) Clean() *ApiGetDataStoreChannelRequestStream {
	*self = ApiGetDataStoreChannelRequestStreamOf()
	return self
}
func (self *ApiGetDataStoreChannelRequestStream) Delete(index int) *ApiGetDataStoreChannelRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetDataStoreChannelRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetDataStoreChannelRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetDataStoreChannelRequestStream) Distinct() *ApiGetDataStoreChannelRequestStream {
	caches := map[string]bool{}
	result := ApiGetDataStoreChannelRequestStreamOf()
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
func (self *ApiGetDataStoreChannelRequestStream) Each(fn func(ApiGetDataStoreChannelRequest)) *ApiGetDataStoreChannelRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetDataStoreChannelRequestStream) EachRight(fn func(ApiGetDataStoreChannelRequest)) *ApiGetDataStoreChannelRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetDataStoreChannelRequestStream) Equals(arr []ApiGetDataStoreChannelRequest) bool {
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
func (self *ApiGetDataStoreChannelRequestStream) Filter(fn func(ApiGetDataStoreChannelRequest, int) bool) *ApiGetDataStoreChannelRequestStream {
	result := ApiGetDataStoreChannelRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDataStoreChannelRequestStream) FilterSlim(fn func(ApiGetDataStoreChannelRequest, int) bool) *ApiGetDataStoreChannelRequestStream {
	result := ApiGetDataStoreChannelRequestStreamOf()
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
func (self *ApiGetDataStoreChannelRequestStream) Find(fn func(ApiGetDataStoreChannelRequest, int) bool) *ApiGetDataStoreChannelRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetDataStoreChannelRequestStream) FindOr(fn func(ApiGetDataStoreChannelRequest, int) bool, or ApiGetDataStoreChannelRequest) ApiGetDataStoreChannelRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetDataStoreChannelRequestStream) FindIndex(fn func(ApiGetDataStoreChannelRequest, int) bool) int {
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
func (self *ApiGetDataStoreChannelRequestStream) First() *ApiGetDataStoreChannelRequest {
	return self.Get(0)
}
func (self *ApiGetDataStoreChannelRequestStream) FirstOr(arg ApiGetDataStoreChannelRequest) ApiGetDataStoreChannelRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDataStoreChannelRequestStream) ForEach(fn func(ApiGetDataStoreChannelRequest, int)) *ApiGetDataStoreChannelRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetDataStoreChannelRequestStream) ForEachRight(fn func(ApiGetDataStoreChannelRequest, int)) *ApiGetDataStoreChannelRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetDataStoreChannelRequestStream) GroupBy(fn func(ApiGetDataStoreChannelRequest, int) string) map[string][]ApiGetDataStoreChannelRequest {
	m := map[string][]ApiGetDataStoreChannelRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetDataStoreChannelRequestStream) GroupByValues(fn func(ApiGetDataStoreChannelRequest, int) string) [][]ApiGetDataStoreChannelRequest {
	var tmp [][]ApiGetDataStoreChannelRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetDataStoreChannelRequestStream) IndexOf(arg ApiGetDataStoreChannelRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetDataStoreChannelRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetDataStoreChannelRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetDataStoreChannelRequestStream) Last() *ApiGetDataStoreChannelRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetDataStoreChannelRequestStream) LastOr(arg ApiGetDataStoreChannelRequest) ApiGetDataStoreChannelRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDataStoreChannelRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetDataStoreChannelRequestStream) Limit(limit int) *ApiGetDataStoreChannelRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetDataStoreChannelRequestStream) Map(fn func(ApiGetDataStoreChannelRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreChannelRequestStream) Map2Int(fn func(ApiGetDataStoreChannelRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreChannelRequestStream) Map2Int32(fn func(ApiGetDataStoreChannelRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreChannelRequestStream) Map2Int64(fn func(ApiGetDataStoreChannelRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreChannelRequestStream) Map2Float32(fn func(ApiGetDataStoreChannelRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreChannelRequestStream) Map2Float64(fn func(ApiGetDataStoreChannelRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreChannelRequestStream) Map2Bool(fn func(ApiGetDataStoreChannelRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreChannelRequestStream) Map2Bytes(fn func(ApiGetDataStoreChannelRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreChannelRequestStream) Map2String(fn func(ApiGetDataStoreChannelRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreChannelRequestStream) Max(fn func(ApiGetDataStoreChannelRequest, int) float64) *ApiGetDataStoreChannelRequest {
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
func (self *ApiGetDataStoreChannelRequestStream) Min(fn func(ApiGetDataStoreChannelRequest, int) float64) *ApiGetDataStoreChannelRequest {
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
func (self *ApiGetDataStoreChannelRequestStream) NoneMatch(fn func(ApiGetDataStoreChannelRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetDataStoreChannelRequestStream) Get(index int) *ApiGetDataStoreChannelRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetDataStoreChannelRequestStream) GetOr(index int, arg ApiGetDataStoreChannelRequest) ApiGetDataStoreChannelRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDataStoreChannelRequestStream) Peek(fn func(*ApiGetDataStoreChannelRequest, int)) *ApiGetDataStoreChannelRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetDataStoreChannelRequestStream) Reduce(fn func(ApiGetDataStoreChannelRequest, ApiGetDataStoreChannelRequest, int) ApiGetDataStoreChannelRequest) *ApiGetDataStoreChannelRequestStream {
	return self.ReduceInit(fn, ApiGetDataStoreChannelRequest{})
}
func (self *ApiGetDataStoreChannelRequestStream) ReduceInit(fn func(ApiGetDataStoreChannelRequest, ApiGetDataStoreChannelRequest, int) ApiGetDataStoreChannelRequest, initialValue ApiGetDataStoreChannelRequest) *ApiGetDataStoreChannelRequestStream {
	result := ApiGetDataStoreChannelRequestStreamOf()
	self.ForEach(func(v ApiGetDataStoreChannelRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDataStoreChannelRequestStream) ReduceInterface(fn func(interface{}, ApiGetDataStoreChannelRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetDataStoreChannelRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetDataStoreChannelRequestStream) ReduceString(fn func(string, ApiGetDataStoreChannelRequest, int) string) []string {
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
func (self *ApiGetDataStoreChannelRequestStream) ReduceInt(fn func(int, ApiGetDataStoreChannelRequest, int) int) []int {
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
func (self *ApiGetDataStoreChannelRequestStream) ReduceInt32(fn func(int32, ApiGetDataStoreChannelRequest, int) int32) []int32 {
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
func (self *ApiGetDataStoreChannelRequestStream) ReduceInt64(fn func(int64, ApiGetDataStoreChannelRequest, int) int64) []int64 {
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
func (self *ApiGetDataStoreChannelRequestStream) ReduceFloat32(fn func(float32, ApiGetDataStoreChannelRequest, int) float32) []float32 {
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
func (self *ApiGetDataStoreChannelRequestStream) ReduceFloat64(fn func(float64, ApiGetDataStoreChannelRequest, int) float64) []float64 {
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
func (self *ApiGetDataStoreChannelRequestStream) ReduceBool(fn func(bool, ApiGetDataStoreChannelRequest, int) bool) []bool {
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
func (self *ApiGetDataStoreChannelRequestStream) Reverse() *ApiGetDataStoreChannelRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetDataStoreChannelRequestStream) Replace(fn func(ApiGetDataStoreChannelRequest, int) ApiGetDataStoreChannelRequest) *ApiGetDataStoreChannelRequestStream {
	return self.ForEach(func(v ApiGetDataStoreChannelRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetDataStoreChannelRequestStream) Select(fn func(ApiGetDataStoreChannelRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetDataStoreChannelRequestStream) Set(index int, val ApiGetDataStoreChannelRequest) *ApiGetDataStoreChannelRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetDataStoreChannelRequestStream) Skip(skip int) *ApiGetDataStoreChannelRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetDataStoreChannelRequestStream) SkippingEach(fn func(ApiGetDataStoreChannelRequest, int) int) *ApiGetDataStoreChannelRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetDataStoreChannelRequestStream) Slice(startIndex, n int) *ApiGetDataStoreChannelRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetDataStoreChannelRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetDataStoreChannelRequestStream) Sort(fn func(i, j int) bool) *ApiGetDataStoreChannelRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetDataStoreChannelRequestStream) Tail() *ApiGetDataStoreChannelRequest {
	return self.Last()
}
func (self *ApiGetDataStoreChannelRequestStream) TailOr(arg ApiGetDataStoreChannelRequest) ApiGetDataStoreChannelRequest {
	return self.LastOr(arg)
}
func (self *ApiGetDataStoreChannelRequestStream) ToList() []ApiGetDataStoreChannelRequest {
	return self.Val()
}
func (self *ApiGetDataStoreChannelRequestStream) Unique() *ApiGetDataStoreChannelRequestStream {
	return self.Distinct()
}
func (self *ApiGetDataStoreChannelRequestStream) Val() []ApiGetDataStoreChannelRequest {
	if self == nil {
		return []ApiGetDataStoreChannelRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetDataStoreChannelRequestStream) While(fn func(ApiGetDataStoreChannelRequest, int) bool) *ApiGetDataStoreChannelRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetDataStoreChannelRequestStream) Where(fn func(ApiGetDataStoreChannelRequest) bool) *ApiGetDataStoreChannelRequestStream {
	result := ApiGetDataStoreChannelRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDataStoreChannelRequestStream) WhereSlim(fn func(ApiGetDataStoreChannelRequest) bool) *ApiGetDataStoreChannelRequestStream {
	result := ApiGetDataStoreChannelRequestStreamOf()
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
