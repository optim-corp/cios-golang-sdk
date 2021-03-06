/*
 * Collection utility of ApiGetDataStoreObjectDataLatestRequest Struct
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

type ApiGetDataStoreObjectDataLatestRequestStream []ApiGetDataStoreObjectDataLatestRequest

func ApiGetDataStoreObjectDataLatestRequestStreamOf(arg ...ApiGetDataStoreObjectDataLatestRequest) ApiGetDataStoreObjectDataLatestRequestStream {
	return arg
}
func ApiGetDataStoreObjectDataLatestRequestStreamFrom(arg []ApiGetDataStoreObjectDataLatestRequest) ApiGetDataStoreObjectDataLatestRequestStream {
	return arg
}
func CreateApiGetDataStoreObjectDataLatestRequestStream(arg ...ApiGetDataStoreObjectDataLatestRequest) *ApiGetDataStoreObjectDataLatestRequestStream {
	tmp := ApiGetDataStoreObjectDataLatestRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetDataStoreObjectDataLatestRequestStream(arg []ApiGetDataStoreObjectDataLatestRequest) *ApiGetDataStoreObjectDataLatestRequestStream {
	tmp := ApiGetDataStoreObjectDataLatestRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetDataStoreObjectDataLatestRequestStream) Add(arg ApiGetDataStoreObjectDataLatestRequest) *ApiGetDataStoreObjectDataLatestRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) AddAll(arg ...ApiGetDataStoreObjectDataLatestRequest) *ApiGetDataStoreObjectDataLatestRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) AddSafe(arg *ApiGetDataStoreObjectDataLatestRequest) *ApiGetDataStoreObjectDataLatestRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Aggregate(fn func(ApiGetDataStoreObjectDataLatestRequest, ApiGetDataStoreObjectDataLatestRequest) ApiGetDataStoreObjectDataLatestRequest) *ApiGetDataStoreObjectDataLatestRequestStream {
	result := ApiGetDataStoreObjectDataLatestRequestStreamOf()
	self.ForEach(func(v ApiGetDataStoreObjectDataLatestRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetDataStoreObjectDataLatestRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) AllMatch(fn func(ApiGetDataStoreObjectDataLatestRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) AnyMatch(fn func(ApiGetDataStoreObjectDataLatestRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Clone() *ApiGetDataStoreObjectDataLatestRequestStream {
	temp := make([]ApiGetDataStoreObjectDataLatestRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetDataStoreObjectDataLatestRequestStream)(&temp)
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Copy() *ApiGetDataStoreObjectDataLatestRequestStream {
	return self.Clone()
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Concat(arg []ApiGetDataStoreObjectDataLatestRequest) *ApiGetDataStoreObjectDataLatestRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Contains(arg ApiGetDataStoreObjectDataLatestRequest) bool {
	return self.FindIndex(func(_arg ApiGetDataStoreObjectDataLatestRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Clean() *ApiGetDataStoreObjectDataLatestRequestStream {
	*self = ApiGetDataStoreObjectDataLatestRequestStreamOf()
	return self
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Delete(index int) *ApiGetDataStoreObjectDataLatestRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetDataStoreObjectDataLatestRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Distinct() *ApiGetDataStoreObjectDataLatestRequestStream {
	caches := map[string]bool{}
	result := ApiGetDataStoreObjectDataLatestRequestStreamOf()
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
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Each(fn func(ApiGetDataStoreObjectDataLatestRequest)) *ApiGetDataStoreObjectDataLatestRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) EachRight(fn func(ApiGetDataStoreObjectDataLatestRequest)) *ApiGetDataStoreObjectDataLatestRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Equals(arr []ApiGetDataStoreObjectDataLatestRequest) bool {
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
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Filter(fn func(ApiGetDataStoreObjectDataLatestRequest, int) bool) *ApiGetDataStoreObjectDataLatestRequestStream {
	result := ApiGetDataStoreObjectDataLatestRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) FilterSlim(fn func(ApiGetDataStoreObjectDataLatestRequest, int) bool) *ApiGetDataStoreObjectDataLatestRequestStream {
	result := ApiGetDataStoreObjectDataLatestRequestStreamOf()
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
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Find(fn func(ApiGetDataStoreObjectDataLatestRequest, int) bool) *ApiGetDataStoreObjectDataLatestRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) FindOr(fn func(ApiGetDataStoreObjectDataLatestRequest, int) bool, or ApiGetDataStoreObjectDataLatestRequest) ApiGetDataStoreObjectDataLatestRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) FindIndex(fn func(ApiGetDataStoreObjectDataLatestRequest, int) bool) int {
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
func (self *ApiGetDataStoreObjectDataLatestRequestStream) First() *ApiGetDataStoreObjectDataLatestRequest {
	return self.Get(0)
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) FirstOr(arg ApiGetDataStoreObjectDataLatestRequest) ApiGetDataStoreObjectDataLatestRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) ForEach(fn func(ApiGetDataStoreObjectDataLatestRequest, int)) *ApiGetDataStoreObjectDataLatestRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) ForEachRight(fn func(ApiGetDataStoreObjectDataLatestRequest, int)) *ApiGetDataStoreObjectDataLatestRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) GroupBy(fn func(ApiGetDataStoreObjectDataLatestRequest, int) string) map[string][]ApiGetDataStoreObjectDataLatestRequest {
	m := map[string][]ApiGetDataStoreObjectDataLatestRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) GroupByValues(fn func(ApiGetDataStoreObjectDataLatestRequest, int) string) [][]ApiGetDataStoreObjectDataLatestRequest {
	var tmp [][]ApiGetDataStoreObjectDataLatestRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) IndexOf(arg ApiGetDataStoreObjectDataLatestRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Last() *ApiGetDataStoreObjectDataLatestRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) LastOr(arg ApiGetDataStoreObjectDataLatestRequest) ApiGetDataStoreObjectDataLatestRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Limit(limit int) *ApiGetDataStoreObjectDataLatestRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetDataStoreObjectDataLatestRequestStream) Map(fn func(ApiGetDataStoreObjectDataLatestRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Map2Int(fn func(ApiGetDataStoreObjectDataLatestRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Map2Int32(fn func(ApiGetDataStoreObjectDataLatestRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Map2Int64(fn func(ApiGetDataStoreObjectDataLatestRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Map2Float32(fn func(ApiGetDataStoreObjectDataLatestRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Map2Float64(fn func(ApiGetDataStoreObjectDataLatestRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Map2Bool(fn func(ApiGetDataStoreObjectDataLatestRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Map2Bytes(fn func(ApiGetDataStoreObjectDataLatestRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Map2String(fn func(ApiGetDataStoreObjectDataLatestRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Max(fn func(ApiGetDataStoreObjectDataLatestRequest, int) float64) *ApiGetDataStoreObjectDataLatestRequest {
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
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Min(fn func(ApiGetDataStoreObjectDataLatestRequest, int) float64) *ApiGetDataStoreObjectDataLatestRequest {
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
func (self *ApiGetDataStoreObjectDataLatestRequestStream) NoneMatch(fn func(ApiGetDataStoreObjectDataLatestRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Get(index int) *ApiGetDataStoreObjectDataLatestRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) GetOr(index int, arg ApiGetDataStoreObjectDataLatestRequest) ApiGetDataStoreObjectDataLatestRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Peek(fn func(*ApiGetDataStoreObjectDataLatestRequest, int)) *ApiGetDataStoreObjectDataLatestRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetDataStoreObjectDataLatestRequestStream) Reduce(fn func(ApiGetDataStoreObjectDataLatestRequest, ApiGetDataStoreObjectDataLatestRequest, int) ApiGetDataStoreObjectDataLatestRequest) *ApiGetDataStoreObjectDataLatestRequestStream {
	return self.ReduceInit(fn, ApiGetDataStoreObjectDataLatestRequest{})
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) ReduceInit(fn func(ApiGetDataStoreObjectDataLatestRequest, ApiGetDataStoreObjectDataLatestRequest, int) ApiGetDataStoreObjectDataLatestRequest, initialValue ApiGetDataStoreObjectDataLatestRequest) *ApiGetDataStoreObjectDataLatestRequestStream {
	result := ApiGetDataStoreObjectDataLatestRequestStreamOf()
	self.ForEach(func(v ApiGetDataStoreObjectDataLatestRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) ReduceInterface(fn func(interface{}, ApiGetDataStoreObjectDataLatestRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetDataStoreObjectDataLatestRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) ReduceString(fn func(string, ApiGetDataStoreObjectDataLatestRequest, int) string) []string {
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
func (self *ApiGetDataStoreObjectDataLatestRequestStream) ReduceInt(fn func(int, ApiGetDataStoreObjectDataLatestRequest, int) int) []int {
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
func (self *ApiGetDataStoreObjectDataLatestRequestStream) ReduceInt32(fn func(int32, ApiGetDataStoreObjectDataLatestRequest, int) int32) []int32 {
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
func (self *ApiGetDataStoreObjectDataLatestRequestStream) ReduceInt64(fn func(int64, ApiGetDataStoreObjectDataLatestRequest, int) int64) []int64 {
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
func (self *ApiGetDataStoreObjectDataLatestRequestStream) ReduceFloat32(fn func(float32, ApiGetDataStoreObjectDataLatestRequest, int) float32) []float32 {
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
func (self *ApiGetDataStoreObjectDataLatestRequestStream) ReduceFloat64(fn func(float64, ApiGetDataStoreObjectDataLatestRequest, int) float64) []float64 {
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
func (self *ApiGetDataStoreObjectDataLatestRequestStream) ReduceBool(fn func(bool, ApiGetDataStoreObjectDataLatestRequest, int) bool) []bool {
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
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Reverse() *ApiGetDataStoreObjectDataLatestRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Replace(fn func(ApiGetDataStoreObjectDataLatestRequest, int) ApiGetDataStoreObjectDataLatestRequest) *ApiGetDataStoreObjectDataLatestRequestStream {
	return self.ForEach(func(v ApiGetDataStoreObjectDataLatestRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Select(fn func(ApiGetDataStoreObjectDataLatestRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Set(index int, val ApiGetDataStoreObjectDataLatestRequest) *ApiGetDataStoreObjectDataLatestRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Skip(skip int) *ApiGetDataStoreObjectDataLatestRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) SkippingEach(fn func(ApiGetDataStoreObjectDataLatestRequest, int) int) *ApiGetDataStoreObjectDataLatestRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Slice(startIndex, n int) *ApiGetDataStoreObjectDataLatestRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetDataStoreObjectDataLatestRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Sort(fn func(i, j int) bool) *ApiGetDataStoreObjectDataLatestRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetDataStoreObjectDataLatestRequestStream) Tail() *ApiGetDataStoreObjectDataLatestRequest {
	return self.Last()
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) TailOr(arg ApiGetDataStoreObjectDataLatestRequest) ApiGetDataStoreObjectDataLatestRequest {
	return self.LastOr(arg)
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) ToList() []ApiGetDataStoreObjectDataLatestRequest {
	return self.Val()
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Unique() *ApiGetDataStoreObjectDataLatestRequestStream {
	return self.Distinct()
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Val() []ApiGetDataStoreObjectDataLatestRequest {
	if self == nil {
		return []ApiGetDataStoreObjectDataLatestRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) While(fn func(ApiGetDataStoreObjectDataLatestRequest, int) bool) *ApiGetDataStoreObjectDataLatestRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) Where(fn func(ApiGetDataStoreObjectDataLatestRequest) bool) *ApiGetDataStoreObjectDataLatestRequestStream {
	result := ApiGetDataStoreObjectDataLatestRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDataStoreObjectDataLatestRequestStream) WhereSlim(fn func(ApiGetDataStoreObjectDataLatestRequest) bool) *ApiGetDataStoreObjectDataLatestRequestStream {
	result := ApiGetDataStoreObjectDataLatestRequestStreamOf()
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
