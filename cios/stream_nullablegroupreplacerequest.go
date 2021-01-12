/*
 * Collection utility of NullableGroupReplaceRequest Struct
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

type NullableGroupReplaceRequestStream []NullableGroupReplaceRequest

func NullableGroupReplaceRequestStreamOf(arg ...NullableGroupReplaceRequest) NullableGroupReplaceRequestStream {
	return arg
}
func NullableGroupReplaceRequestStreamFrom(arg []NullableGroupReplaceRequest) NullableGroupReplaceRequestStream {
	return arg
}
func CreateNullableGroupReplaceRequestStream(arg ...NullableGroupReplaceRequest) *NullableGroupReplaceRequestStream {
	tmp := NullableGroupReplaceRequestStreamOf(arg...)
	return &tmp
}
func GenerateNullableGroupReplaceRequestStream(arg []NullableGroupReplaceRequest) *NullableGroupReplaceRequestStream {
	tmp := NullableGroupReplaceRequestStreamFrom(arg)
	return &tmp
}

func (self *NullableGroupReplaceRequestStream) Add(arg NullableGroupReplaceRequest) *NullableGroupReplaceRequestStream {
	return self.AddAll(arg)
}
func (self *NullableGroupReplaceRequestStream) AddAll(arg ...NullableGroupReplaceRequest) *NullableGroupReplaceRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableGroupReplaceRequestStream) AddSafe(arg *NullableGroupReplaceRequest) *NullableGroupReplaceRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableGroupReplaceRequestStream) Aggregate(fn func(NullableGroupReplaceRequest, NullableGroupReplaceRequest) NullableGroupReplaceRequest) *NullableGroupReplaceRequestStream {
	result := NullableGroupReplaceRequestStreamOf()
	self.ForEach(func(v NullableGroupReplaceRequest, i int) {
		if i == 0 {
			result.Add(fn(NullableGroupReplaceRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableGroupReplaceRequestStream) AllMatch(fn func(NullableGroupReplaceRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableGroupReplaceRequestStream) AnyMatch(fn func(NullableGroupReplaceRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableGroupReplaceRequestStream) Clone() *NullableGroupReplaceRequestStream {
	temp := make([]NullableGroupReplaceRequest, self.Len())
	copy(temp, *self)
	return (*NullableGroupReplaceRequestStream)(&temp)
}
func (self *NullableGroupReplaceRequestStream) Copy() *NullableGroupReplaceRequestStream {
	return self.Clone()
}
func (self *NullableGroupReplaceRequestStream) Concat(arg []NullableGroupReplaceRequest) *NullableGroupReplaceRequestStream {
	return self.AddAll(arg...)
}
func (self *NullableGroupReplaceRequestStream) Contains(arg NullableGroupReplaceRequest) bool {
	return self.FindIndex(func(_arg NullableGroupReplaceRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableGroupReplaceRequestStream) Clean() *NullableGroupReplaceRequestStream {
	*self = NullableGroupReplaceRequestStreamOf()
	return self
}
func (self *NullableGroupReplaceRequestStream) Delete(index int) *NullableGroupReplaceRequestStream {
	return self.DeleteRange(index, index)
}
func (self *NullableGroupReplaceRequestStream) DeleteRange(startIndex, endIndex int) *NullableGroupReplaceRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableGroupReplaceRequestStream) Distinct() *NullableGroupReplaceRequestStream {
	caches := map[string]bool{}
	result := NullableGroupReplaceRequestStreamOf()
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
func (self *NullableGroupReplaceRequestStream) Each(fn func(NullableGroupReplaceRequest)) *NullableGroupReplaceRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableGroupReplaceRequestStream) EachRight(fn func(NullableGroupReplaceRequest)) *NullableGroupReplaceRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableGroupReplaceRequestStream) Equals(arr []NullableGroupReplaceRequest) bool {
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
func (self *NullableGroupReplaceRequestStream) Filter(fn func(NullableGroupReplaceRequest, int) bool) *NullableGroupReplaceRequestStream {
	result := NullableGroupReplaceRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableGroupReplaceRequestStream) FilterSlim(fn func(NullableGroupReplaceRequest, int) bool) *NullableGroupReplaceRequestStream {
	result := NullableGroupReplaceRequestStreamOf()
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
func (self *NullableGroupReplaceRequestStream) Find(fn func(NullableGroupReplaceRequest, int) bool) *NullableGroupReplaceRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableGroupReplaceRequestStream) FindOr(fn func(NullableGroupReplaceRequest, int) bool, or NullableGroupReplaceRequest) NullableGroupReplaceRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableGroupReplaceRequestStream) FindIndex(fn func(NullableGroupReplaceRequest, int) bool) int {
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
func (self *NullableGroupReplaceRequestStream) First() *NullableGroupReplaceRequest {
	return self.Get(0)
}
func (self *NullableGroupReplaceRequestStream) FirstOr(arg NullableGroupReplaceRequest) NullableGroupReplaceRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableGroupReplaceRequestStream) ForEach(fn func(NullableGroupReplaceRequest, int)) *NullableGroupReplaceRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableGroupReplaceRequestStream) ForEachRight(fn func(NullableGroupReplaceRequest, int)) *NullableGroupReplaceRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableGroupReplaceRequestStream) GroupBy(fn func(NullableGroupReplaceRequest, int) string) map[string][]NullableGroupReplaceRequest {
	m := map[string][]NullableGroupReplaceRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableGroupReplaceRequestStream) GroupByValues(fn func(NullableGroupReplaceRequest, int) string) [][]NullableGroupReplaceRequest {
	var tmp [][]NullableGroupReplaceRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableGroupReplaceRequestStream) IndexOf(arg NullableGroupReplaceRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableGroupReplaceRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableGroupReplaceRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableGroupReplaceRequestStream) Last() *NullableGroupReplaceRequest {
	return self.Get(self.Len() - 1)
}
func (self *NullableGroupReplaceRequestStream) LastOr(arg NullableGroupReplaceRequest) NullableGroupReplaceRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableGroupReplaceRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableGroupReplaceRequestStream) Limit(limit int) *NullableGroupReplaceRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableGroupReplaceRequestStream) Map(fn func(NullableGroupReplaceRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupReplaceRequestStream) Map2Int(fn func(NullableGroupReplaceRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupReplaceRequestStream) Map2Int32(fn func(NullableGroupReplaceRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupReplaceRequestStream) Map2Int64(fn func(NullableGroupReplaceRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupReplaceRequestStream) Map2Float32(fn func(NullableGroupReplaceRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupReplaceRequestStream) Map2Float64(fn func(NullableGroupReplaceRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupReplaceRequestStream) Map2Bool(fn func(NullableGroupReplaceRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupReplaceRequestStream) Map2Bytes(fn func(NullableGroupReplaceRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupReplaceRequestStream) Map2String(fn func(NullableGroupReplaceRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableGroupReplaceRequestStream) Max(fn func(NullableGroupReplaceRequest, int) float64) *NullableGroupReplaceRequest {
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
func (self *NullableGroupReplaceRequestStream) Min(fn func(NullableGroupReplaceRequest, int) float64) *NullableGroupReplaceRequest {
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
func (self *NullableGroupReplaceRequestStream) NoneMatch(fn func(NullableGroupReplaceRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableGroupReplaceRequestStream) Get(index int) *NullableGroupReplaceRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableGroupReplaceRequestStream) GetOr(index int, arg NullableGroupReplaceRequest) NullableGroupReplaceRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableGroupReplaceRequestStream) Peek(fn func(*NullableGroupReplaceRequest, int)) *NullableGroupReplaceRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableGroupReplaceRequestStream) Reduce(fn func(NullableGroupReplaceRequest, NullableGroupReplaceRequest, int) NullableGroupReplaceRequest) *NullableGroupReplaceRequestStream {
	return self.ReduceInit(fn, NullableGroupReplaceRequest{})
}
func (self *NullableGroupReplaceRequestStream) ReduceInit(fn func(NullableGroupReplaceRequest, NullableGroupReplaceRequest, int) NullableGroupReplaceRequest, initialValue NullableGroupReplaceRequest) *NullableGroupReplaceRequestStream {
	result := NullableGroupReplaceRequestStreamOf()
	self.ForEach(func(v NullableGroupReplaceRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableGroupReplaceRequestStream) ReduceInterface(fn func(interface{}, NullableGroupReplaceRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableGroupReplaceRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableGroupReplaceRequestStream) ReduceString(fn func(string, NullableGroupReplaceRequest, int) string) []string {
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
func (self *NullableGroupReplaceRequestStream) ReduceInt(fn func(int, NullableGroupReplaceRequest, int) int) []int {
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
func (self *NullableGroupReplaceRequestStream) ReduceInt32(fn func(int32, NullableGroupReplaceRequest, int) int32) []int32 {
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
func (self *NullableGroupReplaceRequestStream) ReduceInt64(fn func(int64, NullableGroupReplaceRequest, int) int64) []int64 {
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
func (self *NullableGroupReplaceRequestStream) ReduceFloat32(fn func(float32, NullableGroupReplaceRequest, int) float32) []float32 {
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
func (self *NullableGroupReplaceRequestStream) ReduceFloat64(fn func(float64, NullableGroupReplaceRequest, int) float64) []float64 {
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
func (self *NullableGroupReplaceRequestStream) ReduceBool(fn func(bool, NullableGroupReplaceRequest, int) bool) []bool {
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
func (self *NullableGroupReplaceRequestStream) Reverse() *NullableGroupReplaceRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableGroupReplaceRequestStream) Replace(fn func(NullableGroupReplaceRequest, int) NullableGroupReplaceRequest) *NullableGroupReplaceRequestStream {
	return self.ForEach(func(v NullableGroupReplaceRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableGroupReplaceRequestStream) Select(fn func(NullableGroupReplaceRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableGroupReplaceRequestStream) Set(index int, val NullableGroupReplaceRequest) *NullableGroupReplaceRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableGroupReplaceRequestStream) Skip(skip int) *NullableGroupReplaceRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableGroupReplaceRequestStream) SkippingEach(fn func(NullableGroupReplaceRequest, int) int) *NullableGroupReplaceRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableGroupReplaceRequestStream) Slice(startIndex, n int) *NullableGroupReplaceRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableGroupReplaceRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableGroupReplaceRequestStream) Sort(fn func(i, j int) bool) *NullableGroupReplaceRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableGroupReplaceRequestStream) Tail() *NullableGroupReplaceRequest {
	return self.Last()
}
func (self *NullableGroupReplaceRequestStream) TailOr(arg NullableGroupReplaceRequest) NullableGroupReplaceRequest {
	return self.LastOr(arg)
}
func (self *NullableGroupReplaceRequestStream) ToList() []NullableGroupReplaceRequest {
	return self.Val()
}
func (self *NullableGroupReplaceRequestStream) Unique() *NullableGroupReplaceRequestStream {
	return self.Distinct()
}
func (self *NullableGroupReplaceRequestStream) Val() []NullableGroupReplaceRequest {
	if self == nil {
		return []NullableGroupReplaceRequest{}
	}
	return *self.Copy()
}
func (self *NullableGroupReplaceRequestStream) While(fn func(NullableGroupReplaceRequest, int) bool) *NullableGroupReplaceRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableGroupReplaceRequestStream) Where(fn func(NullableGroupReplaceRequest) bool) *NullableGroupReplaceRequestStream {
	result := NullableGroupReplaceRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableGroupReplaceRequestStream) WhereSlim(fn func(NullableGroupReplaceRequest) bool) *NullableGroupReplaceRequestStream {
	result := NullableGroupReplaceRequestStreamOf()
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
