/*
 * Collection utility of ApiGetVideoStreamsThumbnailRequest Struct
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

type ApiGetVideoStreamsThumbnailRequestStream []ApiGetVideoStreamsThumbnailRequest

func ApiGetVideoStreamsThumbnailRequestStreamOf(arg ...ApiGetVideoStreamsThumbnailRequest) ApiGetVideoStreamsThumbnailRequestStream {
	return arg
}
func ApiGetVideoStreamsThumbnailRequestStreamFrom(arg []ApiGetVideoStreamsThumbnailRequest) ApiGetVideoStreamsThumbnailRequestStream {
	return arg
}
func CreateApiGetVideoStreamsThumbnailRequestStream(arg ...ApiGetVideoStreamsThumbnailRequest) *ApiGetVideoStreamsThumbnailRequestStream {
	tmp := ApiGetVideoStreamsThumbnailRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetVideoStreamsThumbnailRequestStream(arg []ApiGetVideoStreamsThumbnailRequest) *ApiGetVideoStreamsThumbnailRequestStream {
	tmp := ApiGetVideoStreamsThumbnailRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetVideoStreamsThumbnailRequestStream) Add(arg ApiGetVideoStreamsThumbnailRequest) *ApiGetVideoStreamsThumbnailRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) AddAll(arg ...ApiGetVideoStreamsThumbnailRequest) *ApiGetVideoStreamsThumbnailRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) AddSafe(arg *ApiGetVideoStreamsThumbnailRequest) *ApiGetVideoStreamsThumbnailRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Aggregate(fn func(ApiGetVideoStreamsThumbnailRequest, ApiGetVideoStreamsThumbnailRequest) ApiGetVideoStreamsThumbnailRequest) *ApiGetVideoStreamsThumbnailRequestStream {
	result := ApiGetVideoStreamsThumbnailRequestStreamOf()
	self.ForEach(func(v ApiGetVideoStreamsThumbnailRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetVideoStreamsThumbnailRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) AllMatch(fn func(ApiGetVideoStreamsThumbnailRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) AnyMatch(fn func(ApiGetVideoStreamsThumbnailRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Clone() *ApiGetVideoStreamsThumbnailRequestStream {
	temp := make([]ApiGetVideoStreamsThumbnailRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetVideoStreamsThumbnailRequestStream)(&temp)
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Copy() *ApiGetVideoStreamsThumbnailRequestStream {
	return self.Clone()
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Concat(arg []ApiGetVideoStreamsThumbnailRequest) *ApiGetVideoStreamsThumbnailRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Contains(arg ApiGetVideoStreamsThumbnailRequest) bool {
	return self.FindIndex(func(_arg ApiGetVideoStreamsThumbnailRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Clean() *ApiGetVideoStreamsThumbnailRequestStream {
	*self = ApiGetVideoStreamsThumbnailRequestStreamOf()
	return self
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Delete(index int) *ApiGetVideoStreamsThumbnailRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetVideoStreamsThumbnailRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Distinct() *ApiGetVideoStreamsThumbnailRequestStream {
	caches := map[string]bool{}
	result := ApiGetVideoStreamsThumbnailRequestStreamOf()
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
func (self *ApiGetVideoStreamsThumbnailRequestStream) Each(fn func(ApiGetVideoStreamsThumbnailRequest)) *ApiGetVideoStreamsThumbnailRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) EachRight(fn func(ApiGetVideoStreamsThumbnailRequest)) *ApiGetVideoStreamsThumbnailRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Equals(arr []ApiGetVideoStreamsThumbnailRequest) bool {
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
func (self *ApiGetVideoStreamsThumbnailRequestStream) Filter(fn func(ApiGetVideoStreamsThumbnailRequest, int) bool) *ApiGetVideoStreamsThumbnailRequestStream {
	result := ApiGetVideoStreamsThumbnailRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) FilterSlim(fn func(ApiGetVideoStreamsThumbnailRequest, int) bool) *ApiGetVideoStreamsThumbnailRequestStream {
	result := ApiGetVideoStreamsThumbnailRequestStreamOf()
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
func (self *ApiGetVideoStreamsThumbnailRequestStream) Find(fn func(ApiGetVideoStreamsThumbnailRequest, int) bool) *ApiGetVideoStreamsThumbnailRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) FindOr(fn func(ApiGetVideoStreamsThumbnailRequest, int) bool, or ApiGetVideoStreamsThumbnailRequest) ApiGetVideoStreamsThumbnailRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) FindIndex(fn func(ApiGetVideoStreamsThumbnailRequest, int) bool) int {
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
func (self *ApiGetVideoStreamsThumbnailRequestStream) First() *ApiGetVideoStreamsThumbnailRequest {
	return self.Get(0)
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) FirstOr(arg ApiGetVideoStreamsThumbnailRequest) ApiGetVideoStreamsThumbnailRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) ForEach(fn func(ApiGetVideoStreamsThumbnailRequest, int)) *ApiGetVideoStreamsThumbnailRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) ForEachRight(fn func(ApiGetVideoStreamsThumbnailRequest, int)) *ApiGetVideoStreamsThumbnailRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) GroupBy(fn func(ApiGetVideoStreamsThumbnailRequest, int) string) map[string][]ApiGetVideoStreamsThumbnailRequest {
	m := map[string][]ApiGetVideoStreamsThumbnailRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) GroupByValues(fn func(ApiGetVideoStreamsThumbnailRequest, int) string) [][]ApiGetVideoStreamsThumbnailRequest {
	var tmp [][]ApiGetVideoStreamsThumbnailRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) IndexOf(arg ApiGetVideoStreamsThumbnailRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Last() *ApiGetVideoStreamsThumbnailRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) LastOr(arg ApiGetVideoStreamsThumbnailRequest) ApiGetVideoStreamsThumbnailRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Limit(limit int) *ApiGetVideoStreamsThumbnailRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetVideoStreamsThumbnailRequestStream) Map(fn func(ApiGetVideoStreamsThumbnailRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Map2Int(fn func(ApiGetVideoStreamsThumbnailRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Map2Int32(fn func(ApiGetVideoStreamsThumbnailRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Map2Int64(fn func(ApiGetVideoStreamsThumbnailRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Map2Float32(fn func(ApiGetVideoStreamsThumbnailRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Map2Float64(fn func(ApiGetVideoStreamsThumbnailRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Map2Bool(fn func(ApiGetVideoStreamsThumbnailRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Map2Bytes(fn func(ApiGetVideoStreamsThumbnailRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Map2String(fn func(ApiGetVideoStreamsThumbnailRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Max(fn func(ApiGetVideoStreamsThumbnailRequest, int) float64) *ApiGetVideoStreamsThumbnailRequest {
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
func (self *ApiGetVideoStreamsThumbnailRequestStream) Min(fn func(ApiGetVideoStreamsThumbnailRequest, int) float64) *ApiGetVideoStreamsThumbnailRequest {
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
func (self *ApiGetVideoStreamsThumbnailRequestStream) NoneMatch(fn func(ApiGetVideoStreamsThumbnailRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Get(index int) *ApiGetVideoStreamsThumbnailRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) GetOr(index int, arg ApiGetVideoStreamsThumbnailRequest) ApiGetVideoStreamsThumbnailRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Peek(fn func(*ApiGetVideoStreamsThumbnailRequest, int)) *ApiGetVideoStreamsThumbnailRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetVideoStreamsThumbnailRequestStream) Reduce(fn func(ApiGetVideoStreamsThumbnailRequest, ApiGetVideoStreamsThumbnailRequest, int) ApiGetVideoStreamsThumbnailRequest) *ApiGetVideoStreamsThumbnailRequestStream {
	return self.ReduceInit(fn, ApiGetVideoStreamsThumbnailRequest{})
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) ReduceInit(fn func(ApiGetVideoStreamsThumbnailRequest, ApiGetVideoStreamsThumbnailRequest, int) ApiGetVideoStreamsThumbnailRequest, initialValue ApiGetVideoStreamsThumbnailRequest) *ApiGetVideoStreamsThumbnailRequestStream {
	result := ApiGetVideoStreamsThumbnailRequestStreamOf()
	self.ForEach(func(v ApiGetVideoStreamsThumbnailRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) ReduceInterface(fn func(interface{}, ApiGetVideoStreamsThumbnailRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetVideoStreamsThumbnailRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) ReduceString(fn func(string, ApiGetVideoStreamsThumbnailRequest, int) string) []string {
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
func (self *ApiGetVideoStreamsThumbnailRequestStream) ReduceInt(fn func(int, ApiGetVideoStreamsThumbnailRequest, int) int) []int {
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
func (self *ApiGetVideoStreamsThumbnailRequestStream) ReduceInt32(fn func(int32, ApiGetVideoStreamsThumbnailRequest, int) int32) []int32 {
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
func (self *ApiGetVideoStreamsThumbnailRequestStream) ReduceInt64(fn func(int64, ApiGetVideoStreamsThumbnailRequest, int) int64) []int64 {
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
func (self *ApiGetVideoStreamsThumbnailRequestStream) ReduceFloat32(fn func(float32, ApiGetVideoStreamsThumbnailRequest, int) float32) []float32 {
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
func (self *ApiGetVideoStreamsThumbnailRequestStream) ReduceFloat64(fn func(float64, ApiGetVideoStreamsThumbnailRequest, int) float64) []float64 {
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
func (self *ApiGetVideoStreamsThumbnailRequestStream) ReduceBool(fn func(bool, ApiGetVideoStreamsThumbnailRequest, int) bool) []bool {
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
func (self *ApiGetVideoStreamsThumbnailRequestStream) Reverse() *ApiGetVideoStreamsThumbnailRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Replace(fn func(ApiGetVideoStreamsThumbnailRequest, int) ApiGetVideoStreamsThumbnailRequest) *ApiGetVideoStreamsThumbnailRequestStream {
	return self.ForEach(func(v ApiGetVideoStreamsThumbnailRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Select(fn func(ApiGetVideoStreamsThumbnailRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Set(index int, val ApiGetVideoStreamsThumbnailRequest) *ApiGetVideoStreamsThumbnailRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Skip(skip int) *ApiGetVideoStreamsThumbnailRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) SkippingEach(fn func(ApiGetVideoStreamsThumbnailRequest, int) int) *ApiGetVideoStreamsThumbnailRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Slice(startIndex, n int) *ApiGetVideoStreamsThumbnailRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetVideoStreamsThumbnailRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Sort(fn func(i, j int) bool) *ApiGetVideoStreamsThumbnailRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetVideoStreamsThumbnailRequestStream) Tail() *ApiGetVideoStreamsThumbnailRequest {
	return self.Last()
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) TailOr(arg ApiGetVideoStreamsThumbnailRequest) ApiGetVideoStreamsThumbnailRequest {
	return self.LastOr(arg)
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) ToList() []ApiGetVideoStreamsThumbnailRequest {
	return self.Val()
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Unique() *ApiGetVideoStreamsThumbnailRequestStream {
	return self.Distinct()
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Val() []ApiGetVideoStreamsThumbnailRequest {
	if self == nil {
		return []ApiGetVideoStreamsThumbnailRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) While(fn func(ApiGetVideoStreamsThumbnailRequest, int) bool) *ApiGetVideoStreamsThumbnailRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) Where(fn func(ApiGetVideoStreamsThumbnailRequest) bool) *ApiGetVideoStreamsThumbnailRequestStream {
	result := ApiGetVideoStreamsThumbnailRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetVideoStreamsThumbnailRequestStream) WhereSlim(fn func(ApiGetVideoStreamsThumbnailRequest) bool) *ApiGetVideoStreamsThumbnailRequestStream {
	result := ApiGetVideoStreamsThumbnailRequestStreamOf()
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