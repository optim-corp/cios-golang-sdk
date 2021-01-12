/*
 * Collection utility of ApiCreateGroupRequest Struct
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

type ApiCreateGroupRequestStream []ApiCreateGroupRequest

func ApiCreateGroupRequestStreamOf(arg ...ApiCreateGroupRequest) ApiCreateGroupRequestStream {
	return arg
}
func ApiCreateGroupRequestStreamFrom(arg []ApiCreateGroupRequest) ApiCreateGroupRequestStream {
	return arg
}
func CreateApiCreateGroupRequestStream(arg ...ApiCreateGroupRequest) *ApiCreateGroupRequestStream {
	tmp := ApiCreateGroupRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiCreateGroupRequestStream(arg []ApiCreateGroupRequest) *ApiCreateGroupRequestStream {
	tmp := ApiCreateGroupRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiCreateGroupRequestStream) Add(arg ApiCreateGroupRequest) *ApiCreateGroupRequestStream {
	return self.AddAll(arg)
}
func (self *ApiCreateGroupRequestStream) AddAll(arg ...ApiCreateGroupRequest) *ApiCreateGroupRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiCreateGroupRequestStream) AddSafe(arg *ApiCreateGroupRequest) *ApiCreateGroupRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiCreateGroupRequestStream) Aggregate(fn func(ApiCreateGroupRequest, ApiCreateGroupRequest) ApiCreateGroupRequest) *ApiCreateGroupRequestStream {
	result := ApiCreateGroupRequestStreamOf()
	self.ForEach(func(v ApiCreateGroupRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiCreateGroupRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiCreateGroupRequestStream) AllMatch(fn func(ApiCreateGroupRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiCreateGroupRequestStream) AnyMatch(fn func(ApiCreateGroupRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiCreateGroupRequestStream) Clone() *ApiCreateGroupRequestStream {
	temp := make([]ApiCreateGroupRequest, self.Len())
	copy(temp, *self)
	return (*ApiCreateGroupRequestStream)(&temp)
}
func (self *ApiCreateGroupRequestStream) Copy() *ApiCreateGroupRequestStream {
	return self.Clone()
}
func (self *ApiCreateGroupRequestStream) Concat(arg []ApiCreateGroupRequest) *ApiCreateGroupRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiCreateGroupRequestStream) Contains(arg ApiCreateGroupRequest) bool {
	return self.FindIndex(func(_arg ApiCreateGroupRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiCreateGroupRequestStream) Clean() *ApiCreateGroupRequestStream {
	*self = ApiCreateGroupRequestStreamOf()
	return self
}
func (self *ApiCreateGroupRequestStream) Delete(index int) *ApiCreateGroupRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiCreateGroupRequestStream) DeleteRange(startIndex, endIndex int) *ApiCreateGroupRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiCreateGroupRequestStream) Distinct() *ApiCreateGroupRequestStream {
	caches := map[string]bool{}
	result := ApiCreateGroupRequestStreamOf()
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
func (self *ApiCreateGroupRequestStream) Each(fn func(ApiCreateGroupRequest)) *ApiCreateGroupRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiCreateGroupRequestStream) EachRight(fn func(ApiCreateGroupRequest)) *ApiCreateGroupRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiCreateGroupRequestStream) Equals(arr []ApiCreateGroupRequest) bool {
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
func (self *ApiCreateGroupRequestStream) Filter(fn func(ApiCreateGroupRequest, int) bool) *ApiCreateGroupRequestStream {
	result := ApiCreateGroupRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiCreateGroupRequestStream) FilterSlim(fn func(ApiCreateGroupRequest, int) bool) *ApiCreateGroupRequestStream {
	result := ApiCreateGroupRequestStreamOf()
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
func (self *ApiCreateGroupRequestStream) Find(fn func(ApiCreateGroupRequest, int) bool) *ApiCreateGroupRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiCreateGroupRequestStream) FindOr(fn func(ApiCreateGroupRequest, int) bool, or ApiCreateGroupRequest) ApiCreateGroupRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiCreateGroupRequestStream) FindIndex(fn func(ApiCreateGroupRequest, int) bool) int {
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
func (self *ApiCreateGroupRequestStream) First() *ApiCreateGroupRequest {
	return self.Get(0)
}
func (self *ApiCreateGroupRequestStream) FirstOr(arg ApiCreateGroupRequest) ApiCreateGroupRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiCreateGroupRequestStream) ForEach(fn func(ApiCreateGroupRequest, int)) *ApiCreateGroupRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiCreateGroupRequestStream) ForEachRight(fn func(ApiCreateGroupRequest, int)) *ApiCreateGroupRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiCreateGroupRequestStream) GroupBy(fn func(ApiCreateGroupRequest, int) string) map[string][]ApiCreateGroupRequest {
	m := map[string][]ApiCreateGroupRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiCreateGroupRequestStream) GroupByValues(fn func(ApiCreateGroupRequest, int) string) [][]ApiCreateGroupRequest {
	var tmp [][]ApiCreateGroupRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiCreateGroupRequestStream) IndexOf(arg ApiCreateGroupRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiCreateGroupRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiCreateGroupRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiCreateGroupRequestStream) Last() *ApiCreateGroupRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiCreateGroupRequestStream) LastOr(arg ApiCreateGroupRequest) ApiCreateGroupRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiCreateGroupRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiCreateGroupRequestStream) Limit(limit int) *ApiCreateGroupRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiCreateGroupRequestStream) Map(fn func(ApiCreateGroupRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateGroupRequestStream) Map2Int(fn func(ApiCreateGroupRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateGroupRequestStream) Map2Int32(fn func(ApiCreateGroupRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateGroupRequestStream) Map2Int64(fn func(ApiCreateGroupRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateGroupRequestStream) Map2Float32(fn func(ApiCreateGroupRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateGroupRequestStream) Map2Float64(fn func(ApiCreateGroupRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateGroupRequestStream) Map2Bool(fn func(ApiCreateGroupRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateGroupRequestStream) Map2Bytes(fn func(ApiCreateGroupRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateGroupRequestStream) Map2String(fn func(ApiCreateGroupRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiCreateGroupRequestStream) Max(fn func(ApiCreateGroupRequest, int) float64) *ApiCreateGroupRequest {
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
func (self *ApiCreateGroupRequestStream) Min(fn func(ApiCreateGroupRequest, int) float64) *ApiCreateGroupRequest {
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
func (self *ApiCreateGroupRequestStream) NoneMatch(fn func(ApiCreateGroupRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiCreateGroupRequestStream) Get(index int) *ApiCreateGroupRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiCreateGroupRequestStream) GetOr(index int, arg ApiCreateGroupRequest) ApiCreateGroupRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiCreateGroupRequestStream) Peek(fn func(*ApiCreateGroupRequest, int)) *ApiCreateGroupRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiCreateGroupRequestStream) Reduce(fn func(ApiCreateGroupRequest, ApiCreateGroupRequest, int) ApiCreateGroupRequest) *ApiCreateGroupRequestStream {
	return self.ReduceInit(fn, ApiCreateGroupRequest{})
}
func (self *ApiCreateGroupRequestStream) ReduceInit(fn func(ApiCreateGroupRequest, ApiCreateGroupRequest, int) ApiCreateGroupRequest, initialValue ApiCreateGroupRequest) *ApiCreateGroupRequestStream {
	result := ApiCreateGroupRequestStreamOf()
	self.ForEach(func(v ApiCreateGroupRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiCreateGroupRequestStream) ReduceInterface(fn func(interface{}, ApiCreateGroupRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiCreateGroupRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiCreateGroupRequestStream) ReduceString(fn func(string, ApiCreateGroupRequest, int) string) []string {
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
func (self *ApiCreateGroupRequestStream) ReduceInt(fn func(int, ApiCreateGroupRequest, int) int) []int {
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
func (self *ApiCreateGroupRequestStream) ReduceInt32(fn func(int32, ApiCreateGroupRequest, int) int32) []int32 {
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
func (self *ApiCreateGroupRequestStream) ReduceInt64(fn func(int64, ApiCreateGroupRequest, int) int64) []int64 {
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
func (self *ApiCreateGroupRequestStream) ReduceFloat32(fn func(float32, ApiCreateGroupRequest, int) float32) []float32 {
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
func (self *ApiCreateGroupRequestStream) ReduceFloat64(fn func(float64, ApiCreateGroupRequest, int) float64) []float64 {
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
func (self *ApiCreateGroupRequestStream) ReduceBool(fn func(bool, ApiCreateGroupRequest, int) bool) []bool {
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
func (self *ApiCreateGroupRequestStream) Reverse() *ApiCreateGroupRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiCreateGroupRequestStream) Replace(fn func(ApiCreateGroupRequest, int) ApiCreateGroupRequest) *ApiCreateGroupRequestStream {
	return self.ForEach(func(v ApiCreateGroupRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiCreateGroupRequestStream) Select(fn func(ApiCreateGroupRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiCreateGroupRequestStream) Set(index int, val ApiCreateGroupRequest) *ApiCreateGroupRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiCreateGroupRequestStream) Skip(skip int) *ApiCreateGroupRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiCreateGroupRequestStream) SkippingEach(fn func(ApiCreateGroupRequest, int) int) *ApiCreateGroupRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiCreateGroupRequestStream) Slice(startIndex, n int) *ApiCreateGroupRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiCreateGroupRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiCreateGroupRequestStream) Sort(fn func(i, j int) bool) *ApiCreateGroupRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiCreateGroupRequestStream) Tail() *ApiCreateGroupRequest {
	return self.Last()
}
func (self *ApiCreateGroupRequestStream) TailOr(arg ApiCreateGroupRequest) ApiCreateGroupRequest {
	return self.LastOr(arg)
}
func (self *ApiCreateGroupRequestStream) ToList() []ApiCreateGroupRequest {
	return self.Val()
}
func (self *ApiCreateGroupRequestStream) Unique() *ApiCreateGroupRequestStream {
	return self.Distinct()
}
func (self *ApiCreateGroupRequestStream) Val() []ApiCreateGroupRequest {
	if self == nil {
		return []ApiCreateGroupRequest{}
	}
	return *self.Copy()
}
func (self *ApiCreateGroupRequestStream) While(fn func(ApiCreateGroupRequest, int) bool) *ApiCreateGroupRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiCreateGroupRequestStream) Where(fn func(ApiCreateGroupRequest) bool) *ApiCreateGroupRequestStream {
	result := ApiCreateGroupRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiCreateGroupRequestStream) WhereSlim(fn func(ApiCreateGroupRequest) bool) *ApiCreateGroupRequestStream {
	result := ApiCreateGroupRequestStreamOf()
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
