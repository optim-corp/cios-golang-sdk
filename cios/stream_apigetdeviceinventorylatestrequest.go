/*
 * Collection utility of ApiGetDeviceInventoryLatestRequest Struct
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

type ApiGetDeviceInventoryLatestRequestStream []ApiGetDeviceInventoryLatestRequest

func ApiGetDeviceInventoryLatestRequestStreamOf(arg ...ApiGetDeviceInventoryLatestRequest) ApiGetDeviceInventoryLatestRequestStream {
	return arg
}
func ApiGetDeviceInventoryLatestRequestStreamFrom(arg []ApiGetDeviceInventoryLatestRequest) ApiGetDeviceInventoryLatestRequestStream {
	return arg
}
func CreateApiGetDeviceInventoryLatestRequestStream(arg ...ApiGetDeviceInventoryLatestRequest) *ApiGetDeviceInventoryLatestRequestStream {
	tmp := ApiGetDeviceInventoryLatestRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetDeviceInventoryLatestRequestStream(arg []ApiGetDeviceInventoryLatestRequest) *ApiGetDeviceInventoryLatestRequestStream {
	tmp := ApiGetDeviceInventoryLatestRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetDeviceInventoryLatestRequestStream) Add(arg ApiGetDeviceInventoryLatestRequest) *ApiGetDeviceInventoryLatestRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetDeviceInventoryLatestRequestStream) AddAll(arg ...ApiGetDeviceInventoryLatestRequest) *ApiGetDeviceInventoryLatestRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetDeviceInventoryLatestRequestStream) AddSafe(arg *ApiGetDeviceInventoryLatestRequest) *ApiGetDeviceInventoryLatestRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Aggregate(fn func(ApiGetDeviceInventoryLatestRequest, ApiGetDeviceInventoryLatestRequest) ApiGetDeviceInventoryLatestRequest) *ApiGetDeviceInventoryLatestRequestStream {
	result := ApiGetDeviceInventoryLatestRequestStreamOf()
	self.ForEach(func(v ApiGetDeviceInventoryLatestRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetDeviceInventoryLatestRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDeviceInventoryLatestRequestStream) AllMatch(fn func(ApiGetDeviceInventoryLatestRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetDeviceInventoryLatestRequestStream) AnyMatch(fn func(ApiGetDeviceInventoryLatestRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Clone() *ApiGetDeviceInventoryLatestRequestStream {
	temp := make([]ApiGetDeviceInventoryLatestRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetDeviceInventoryLatestRequestStream)(&temp)
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Copy() *ApiGetDeviceInventoryLatestRequestStream {
	return self.Clone()
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Concat(arg []ApiGetDeviceInventoryLatestRequest) *ApiGetDeviceInventoryLatestRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Contains(arg ApiGetDeviceInventoryLatestRequest) bool {
	return self.FindIndex(func(_arg ApiGetDeviceInventoryLatestRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Clean() *ApiGetDeviceInventoryLatestRequestStream {
	*self = ApiGetDeviceInventoryLatestRequestStreamOf()
	return self
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Delete(index int) *ApiGetDeviceInventoryLatestRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetDeviceInventoryLatestRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetDeviceInventoryLatestRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Distinct() *ApiGetDeviceInventoryLatestRequestStream {
	caches := map[string]bool{}
	result := ApiGetDeviceInventoryLatestRequestStreamOf()
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
func (self *ApiGetDeviceInventoryLatestRequestStream) Each(fn func(ApiGetDeviceInventoryLatestRequest)) *ApiGetDeviceInventoryLatestRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetDeviceInventoryLatestRequestStream) EachRight(fn func(ApiGetDeviceInventoryLatestRequest)) *ApiGetDeviceInventoryLatestRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Equals(arr []ApiGetDeviceInventoryLatestRequest) bool {
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
func (self *ApiGetDeviceInventoryLatestRequestStream) Filter(fn func(ApiGetDeviceInventoryLatestRequest, int) bool) *ApiGetDeviceInventoryLatestRequestStream {
	result := ApiGetDeviceInventoryLatestRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDeviceInventoryLatestRequestStream) FilterSlim(fn func(ApiGetDeviceInventoryLatestRequest, int) bool) *ApiGetDeviceInventoryLatestRequestStream {
	result := ApiGetDeviceInventoryLatestRequestStreamOf()
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
func (self *ApiGetDeviceInventoryLatestRequestStream) Find(fn func(ApiGetDeviceInventoryLatestRequest, int) bool) *ApiGetDeviceInventoryLatestRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetDeviceInventoryLatestRequestStream) FindOr(fn func(ApiGetDeviceInventoryLatestRequest, int) bool, or ApiGetDeviceInventoryLatestRequest) ApiGetDeviceInventoryLatestRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetDeviceInventoryLatestRequestStream) FindIndex(fn func(ApiGetDeviceInventoryLatestRequest, int) bool) int {
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
func (self *ApiGetDeviceInventoryLatestRequestStream) First() *ApiGetDeviceInventoryLatestRequest {
	return self.Get(0)
}
func (self *ApiGetDeviceInventoryLatestRequestStream) FirstOr(arg ApiGetDeviceInventoryLatestRequest) ApiGetDeviceInventoryLatestRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDeviceInventoryLatestRequestStream) ForEach(fn func(ApiGetDeviceInventoryLatestRequest, int)) *ApiGetDeviceInventoryLatestRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetDeviceInventoryLatestRequestStream) ForEachRight(fn func(ApiGetDeviceInventoryLatestRequest, int)) *ApiGetDeviceInventoryLatestRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetDeviceInventoryLatestRequestStream) GroupBy(fn func(ApiGetDeviceInventoryLatestRequest, int) string) map[string][]ApiGetDeviceInventoryLatestRequest {
	m := map[string][]ApiGetDeviceInventoryLatestRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetDeviceInventoryLatestRequestStream) GroupByValues(fn func(ApiGetDeviceInventoryLatestRequest, int) string) [][]ApiGetDeviceInventoryLatestRequest {
	var tmp [][]ApiGetDeviceInventoryLatestRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetDeviceInventoryLatestRequestStream) IndexOf(arg ApiGetDeviceInventoryLatestRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetDeviceInventoryLatestRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetDeviceInventoryLatestRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Last() *ApiGetDeviceInventoryLatestRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetDeviceInventoryLatestRequestStream) LastOr(arg ApiGetDeviceInventoryLatestRequest) ApiGetDeviceInventoryLatestRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Limit(limit int) *ApiGetDeviceInventoryLatestRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetDeviceInventoryLatestRequestStream) Map(fn func(ApiGetDeviceInventoryLatestRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Map2Int(fn func(ApiGetDeviceInventoryLatestRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Map2Int32(fn func(ApiGetDeviceInventoryLatestRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Map2Int64(fn func(ApiGetDeviceInventoryLatestRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Map2Float32(fn func(ApiGetDeviceInventoryLatestRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Map2Float64(fn func(ApiGetDeviceInventoryLatestRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Map2Bool(fn func(ApiGetDeviceInventoryLatestRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Map2Bytes(fn func(ApiGetDeviceInventoryLatestRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Map2String(fn func(ApiGetDeviceInventoryLatestRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Max(fn func(ApiGetDeviceInventoryLatestRequest, int) float64) *ApiGetDeviceInventoryLatestRequest {
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
func (self *ApiGetDeviceInventoryLatestRequestStream) Min(fn func(ApiGetDeviceInventoryLatestRequest, int) float64) *ApiGetDeviceInventoryLatestRequest {
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
func (self *ApiGetDeviceInventoryLatestRequestStream) NoneMatch(fn func(ApiGetDeviceInventoryLatestRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Get(index int) *ApiGetDeviceInventoryLatestRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetDeviceInventoryLatestRequestStream) GetOr(index int, arg ApiGetDeviceInventoryLatestRequest) ApiGetDeviceInventoryLatestRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Peek(fn func(*ApiGetDeviceInventoryLatestRequest, int)) *ApiGetDeviceInventoryLatestRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetDeviceInventoryLatestRequestStream) Reduce(fn func(ApiGetDeviceInventoryLatestRequest, ApiGetDeviceInventoryLatestRequest, int) ApiGetDeviceInventoryLatestRequest) *ApiGetDeviceInventoryLatestRequestStream {
	return self.ReduceInit(fn, ApiGetDeviceInventoryLatestRequest{})
}
func (self *ApiGetDeviceInventoryLatestRequestStream) ReduceInit(fn func(ApiGetDeviceInventoryLatestRequest, ApiGetDeviceInventoryLatestRequest, int) ApiGetDeviceInventoryLatestRequest, initialValue ApiGetDeviceInventoryLatestRequest) *ApiGetDeviceInventoryLatestRequestStream {
	result := ApiGetDeviceInventoryLatestRequestStreamOf()
	self.ForEach(func(v ApiGetDeviceInventoryLatestRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDeviceInventoryLatestRequestStream) ReduceInterface(fn func(interface{}, ApiGetDeviceInventoryLatestRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetDeviceInventoryLatestRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetDeviceInventoryLatestRequestStream) ReduceString(fn func(string, ApiGetDeviceInventoryLatestRequest, int) string) []string {
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
func (self *ApiGetDeviceInventoryLatestRequestStream) ReduceInt(fn func(int, ApiGetDeviceInventoryLatestRequest, int) int) []int {
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
func (self *ApiGetDeviceInventoryLatestRequestStream) ReduceInt32(fn func(int32, ApiGetDeviceInventoryLatestRequest, int) int32) []int32 {
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
func (self *ApiGetDeviceInventoryLatestRequestStream) ReduceInt64(fn func(int64, ApiGetDeviceInventoryLatestRequest, int) int64) []int64 {
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
func (self *ApiGetDeviceInventoryLatestRequestStream) ReduceFloat32(fn func(float32, ApiGetDeviceInventoryLatestRequest, int) float32) []float32 {
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
func (self *ApiGetDeviceInventoryLatestRequestStream) ReduceFloat64(fn func(float64, ApiGetDeviceInventoryLatestRequest, int) float64) []float64 {
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
func (self *ApiGetDeviceInventoryLatestRequestStream) ReduceBool(fn func(bool, ApiGetDeviceInventoryLatestRequest, int) bool) []bool {
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
func (self *ApiGetDeviceInventoryLatestRequestStream) Reverse() *ApiGetDeviceInventoryLatestRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Replace(fn func(ApiGetDeviceInventoryLatestRequest, int) ApiGetDeviceInventoryLatestRequest) *ApiGetDeviceInventoryLatestRequestStream {
	return self.ForEach(func(v ApiGetDeviceInventoryLatestRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Select(fn func(ApiGetDeviceInventoryLatestRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Set(index int, val ApiGetDeviceInventoryLatestRequest) *ApiGetDeviceInventoryLatestRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Skip(skip int) *ApiGetDeviceInventoryLatestRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetDeviceInventoryLatestRequestStream) SkippingEach(fn func(ApiGetDeviceInventoryLatestRequest, int) int) *ApiGetDeviceInventoryLatestRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Slice(startIndex, n int) *ApiGetDeviceInventoryLatestRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetDeviceInventoryLatestRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Sort(fn func(i, j int) bool) *ApiGetDeviceInventoryLatestRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetDeviceInventoryLatestRequestStream) Tail() *ApiGetDeviceInventoryLatestRequest {
	return self.Last()
}
func (self *ApiGetDeviceInventoryLatestRequestStream) TailOr(arg ApiGetDeviceInventoryLatestRequest) ApiGetDeviceInventoryLatestRequest {
	return self.LastOr(arg)
}
func (self *ApiGetDeviceInventoryLatestRequestStream) ToList() []ApiGetDeviceInventoryLatestRequest {
	return self.Val()
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Unique() *ApiGetDeviceInventoryLatestRequestStream {
	return self.Distinct()
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Val() []ApiGetDeviceInventoryLatestRequest {
	if self == nil {
		return []ApiGetDeviceInventoryLatestRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetDeviceInventoryLatestRequestStream) While(fn func(ApiGetDeviceInventoryLatestRequest, int) bool) *ApiGetDeviceInventoryLatestRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetDeviceInventoryLatestRequestStream) Where(fn func(ApiGetDeviceInventoryLatestRequest) bool) *ApiGetDeviceInventoryLatestRequestStream {
	result := ApiGetDeviceInventoryLatestRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDeviceInventoryLatestRequestStream) WhereSlim(fn func(ApiGetDeviceInventoryLatestRequest) bool) *ApiGetDeviceInventoryLatestRequestStream {
	result := ApiGetDeviceInventoryLatestRequestStreamOf()
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
