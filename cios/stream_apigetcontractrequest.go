/*
 * Collection utility of ApiGetContractRequest Struct
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

type ApiGetContractRequestStream []ApiGetContractRequest

func ApiGetContractRequestStreamOf(arg ...ApiGetContractRequest) ApiGetContractRequestStream {
	return arg
}
func ApiGetContractRequestStreamFrom(arg []ApiGetContractRequest) ApiGetContractRequestStream {
	return arg
}
func CreateApiGetContractRequestStream(arg ...ApiGetContractRequest) *ApiGetContractRequestStream {
	tmp := ApiGetContractRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetContractRequestStream(arg []ApiGetContractRequest) *ApiGetContractRequestStream {
	tmp := ApiGetContractRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetContractRequestStream) Add(arg ApiGetContractRequest) *ApiGetContractRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetContractRequestStream) AddAll(arg ...ApiGetContractRequest) *ApiGetContractRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetContractRequestStream) AddSafe(arg *ApiGetContractRequest) *ApiGetContractRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetContractRequestStream) Aggregate(fn func(ApiGetContractRequest, ApiGetContractRequest) ApiGetContractRequest) *ApiGetContractRequestStream {
	result := ApiGetContractRequestStreamOf()
	self.ForEach(func(v ApiGetContractRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetContractRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetContractRequestStream) AllMatch(fn func(ApiGetContractRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetContractRequestStream) AnyMatch(fn func(ApiGetContractRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetContractRequestStream) Clone() *ApiGetContractRequestStream {
	temp := make([]ApiGetContractRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetContractRequestStream)(&temp)
}
func (self *ApiGetContractRequestStream) Copy() *ApiGetContractRequestStream {
	return self.Clone()
}
func (self *ApiGetContractRequestStream) Concat(arg []ApiGetContractRequest) *ApiGetContractRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetContractRequestStream) Contains(arg ApiGetContractRequest) bool {
	return self.FindIndex(func(_arg ApiGetContractRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetContractRequestStream) Clean() *ApiGetContractRequestStream {
	*self = ApiGetContractRequestStreamOf()
	return self
}
func (self *ApiGetContractRequestStream) Delete(index int) *ApiGetContractRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetContractRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetContractRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetContractRequestStream) Distinct() *ApiGetContractRequestStream {
	caches := map[string]bool{}
	result := ApiGetContractRequestStreamOf()
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
func (self *ApiGetContractRequestStream) Each(fn func(ApiGetContractRequest)) *ApiGetContractRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetContractRequestStream) EachRight(fn func(ApiGetContractRequest)) *ApiGetContractRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetContractRequestStream) Equals(arr []ApiGetContractRequest) bool {
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
func (self *ApiGetContractRequestStream) Filter(fn func(ApiGetContractRequest, int) bool) *ApiGetContractRequestStream {
	result := ApiGetContractRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetContractRequestStream) FilterSlim(fn func(ApiGetContractRequest, int) bool) *ApiGetContractRequestStream {
	result := ApiGetContractRequestStreamOf()
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
func (self *ApiGetContractRequestStream) Find(fn func(ApiGetContractRequest, int) bool) *ApiGetContractRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetContractRequestStream) FindOr(fn func(ApiGetContractRequest, int) bool, or ApiGetContractRequest) ApiGetContractRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetContractRequestStream) FindIndex(fn func(ApiGetContractRequest, int) bool) int {
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
func (self *ApiGetContractRequestStream) First() *ApiGetContractRequest {
	return self.Get(0)
}
func (self *ApiGetContractRequestStream) FirstOr(arg ApiGetContractRequest) ApiGetContractRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetContractRequestStream) ForEach(fn func(ApiGetContractRequest, int)) *ApiGetContractRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetContractRequestStream) ForEachRight(fn func(ApiGetContractRequest, int)) *ApiGetContractRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetContractRequestStream) GroupBy(fn func(ApiGetContractRequest, int) string) map[string][]ApiGetContractRequest {
	m := map[string][]ApiGetContractRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetContractRequestStream) GroupByValues(fn func(ApiGetContractRequest, int) string) [][]ApiGetContractRequest {
	var tmp [][]ApiGetContractRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetContractRequestStream) IndexOf(arg ApiGetContractRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetContractRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetContractRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetContractRequestStream) Last() *ApiGetContractRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetContractRequestStream) LastOr(arg ApiGetContractRequest) ApiGetContractRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetContractRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetContractRequestStream) Limit(limit int) *ApiGetContractRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetContractRequestStream) Map(fn func(ApiGetContractRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetContractRequestStream) Map2Int(fn func(ApiGetContractRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetContractRequestStream) Map2Int32(fn func(ApiGetContractRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetContractRequestStream) Map2Int64(fn func(ApiGetContractRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetContractRequestStream) Map2Float32(fn func(ApiGetContractRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetContractRequestStream) Map2Float64(fn func(ApiGetContractRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetContractRequestStream) Map2Bool(fn func(ApiGetContractRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetContractRequestStream) Map2Bytes(fn func(ApiGetContractRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetContractRequestStream) Map2String(fn func(ApiGetContractRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetContractRequestStream) Max(fn func(ApiGetContractRequest, int) float64) *ApiGetContractRequest {
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
func (self *ApiGetContractRequestStream) Min(fn func(ApiGetContractRequest, int) float64) *ApiGetContractRequest {
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
func (self *ApiGetContractRequestStream) NoneMatch(fn func(ApiGetContractRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetContractRequestStream) Get(index int) *ApiGetContractRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetContractRequestStream) GetOr(index int, arg ApiGetContractRequest) ApiGetContractRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetContractRequestStream) Peek(fn func(*ApiGetContractRequest, int)) *ApiGetContractRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetContractRequestStream) Reduce(fn func(ApiGetContractRequest, ApiGetContractRequest, int) ApiGetContractRequest) *ApiGetContractRequestStream {
	return self.ReduceInit(fn, ApiGetContractRequest{})
}
func (self *ApiGetContractRequestStream) ReduceInit(fn func(ApiGetContractRequest, ApiGetContractRequest, int) ApiGetContractRequest, initialValue ApiGetContractRequest) *ApiGetContractRequestStream {
	result := ApiGetContractRequestStreamOf()
	self.ForEach(func(v ApiGetContractRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetContractRequestStream) ReduceInterface(fn func(interface{}, ApiGetContractRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetContractRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetContractRequestStream) ReduceString(fn func(string, ApiGetContractRequest, int) string) []string {
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
func (self *ApiGetContractRequestStream) ReduceInt(fn func(int, ApiGetContractRequest, int) int) []int {
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
func (self *ApiGetContractRequestStream) ReduceInt32(fn func(int32, ApiGetContractRequest, int) int32) []int32 {
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
func (self *ApiGetContractRequestStream) ReduceInt64(fn func(int64, ApiGetContractRequest, int) int64) []int64 {
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
func (self *ApiGetContractRequestStream) ReduceFloat32(fn func(float32, ApiGetContractRequest, int) float32) []float32 {
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
func (self *ApiGetContractRequestStream) ReduceFloat64(fn func(float64, ApiGetContractRequest, int) float64) []float64 {
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
func (self *ApiGetContractRequestStream) ReduceBool(fn func(bool, ApiGetContractRequest, int) bool) []bool {
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
func (self *ApiGetContractRequestStream) Reverse() *ApiGetContractRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetContractRequestStream) Replace(fn func(ApiGetContractRequest, int) ApiGetContractRequest) *ApiGetContractRequestStream {
	return self.ForEach(func(v ApiGetContractRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetContractRequestStream) Select(fn func(ApiGetContractRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetContractRequestStream) Set(index int, val ApiGetContractRequest) *ApiGetContractRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetContractRequestStream) Skip(skip int) *ApiGetContractRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetContractRequestStream) SkippingEach(fn func(ApiGetContractRequest, int) int) *ApiGetContractRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetContractRequestStream) Slice(startIndex, n int) *ApiGetContractRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetContractRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetContractRequestStream) Sort(fn func(i, j int) bool) *ApiGetContractRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetContractRequestStream) Tail() *ApiGetContractRequest {
	return self.Last()
}
func (self *ApiGetContractRequestStream) TailOr(arg ApiGetContractRequest) ApiGetContractRequest {
	return self.LastOr(arg)
}
func (self *ApiGetContractRequestStream) ToList() []ApiGetContractRequest {
	return self.Val()
}
func (self *ApiGetContractRequestStream) Unique() *ApiGetContractRequestStream {
	return self.Distinct()
}
func (self *ApiGetContractRequestStream) Val() []ApiGetContractRequest {
	if self == nil {
		return []ApiGetContractRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetContractRequestStream) While(fn func(ApiGetContractRequest, int) bool) *ApiGetContractRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetContractRequestStream) Where(fn func(ApiGetContractRequest) bool) *ApiGetContractRequestStream {
	result := ApiGetContractRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetContractRequestStream) WhereSlim(fn func(ApiGetContractRequest) bool) *ApiGetContractRequestStream {
	result := ApiGetContractRequestStreamOf()
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
