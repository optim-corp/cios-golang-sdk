/*
 * Collection utility of NullableSeriesRequest Struct
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

type NullableSeriesRequestStream []NullableSeriesRequest

func NullableSeriesRequestStreamOf(arg ...NullableSeriesRequest) NullableSeriesRequestStream {
	return arg
}
func NullableSeriesRequestStreamFrom(arg []NullableSeriesRequest) NullableSeriesRequestStream {
	return arg
}
func CreateNullableSeriesRequestStream(arg ...NullableSeriesRequest) *NullableSeriesRequestStream {
	tmp := NullableSeriesRequestStreamOf(arg...)
	return &tmp
}
func GenerateNullableSeriesRequestStream(arg []NullableSeriesRequest) *NullableSeriesRequestStream {
	tmp := NullableSeriesRequestStreamFrom(arg)
	return &tmp
}

func (self *NullableSeriesRequestStream) Add(arg NullableSeriesRequest) *NullableSeriesRequestStream {
	return self.AddAll(arg)
}
func (self *NullableSeriesRequestStream) AddAll(arg ...NullableSeriesRequest) *NullableSeriesRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableSeriesRequestStream) AddSafe(arg *NullableSeriesRequest) *NullableSeriesRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableSeriesRequestStream) Aggregate(fn func(NullableSeriesRequest, NullableSeriesRequest) NullableSeriesRequest) *NullableSeriesRequestStream {
	result := NullableSeriesRequestStreamOf()
	self.ForEach(func(v NullableSeriesRequest, i int) {
		if i == 0 {
			result.Add(fn(NullableSeriesRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableSeriesRequestStream) AllMatch(fn func(NullableSeriesRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableSeriesRequestStream) AnyMatch(fn func(NullableSeriesRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableSeriesRequestStream) Clone() *NullableSeriesRequestStream {
	temp := make([]NullableSeriesRequest, self.Len())
	copy(temp, *self)
	return (*NullableSeriesRequestStream)(&temp)
}
func (self *NullableSeriesRequestStream) Copy() *NullableSeriesRequestStream {
	return self.Clone()
}
func (self *NullableSeriesRequestStream) Concat(arg []NullableSeriesRequest) *NullableSeriesRequestStream {
	return self.AddAll(arg...)
}
func (self *NullableSeriesRequestStream) Contains(arg NullableSeriesRequest) bool {
	return self.FindIndex(func(_arg NullableSeriesRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableSeriesRequestStream) Clean() *NullableSeriesRequestStream {
	*self = NullableSeriesRequestStreamOf()
	return self
}
func (self *NullableSeriesRequestStream) Delete(index int) *NullableSeriesRequestStream {
	return self.DeleteRange(index, index)
}
func (self *NullableSeriesRequestStream) DeleteRange(startIndex, endIndex int) *NullableSeriesRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableSeriesRequestStream) Distinct() *NullableSeriesRequestStream {
	caches := map[string]bool{}
	result := NullableSeriesRequestStreamOf()
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
func (self *NullableSeriesRequestStream) Each(fn func(NullableSeriesRequest)) *NullableSeriesRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableSeriesRequestStream) EachRight(fn func(NullableSeriesRequest)) *NullableSeriesRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableSeriesRequestStream) Equals(arr []NullableSeriesRequest) bool {
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
func (self *NullableSeriesRequestStream) Filter(fn func(NullableSeriesRequest, int) bool) *NullableSeriesRequestStream {
	result := NullableSeriesRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSeriesRequestStream) FilterSlim(fn func(NullableSeriesRequest, int) bool) *NullableSeriesRequestStream {
	result := NullableSeriesRequestStreamOf()
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
func (self *NullableSeriesRequestStream) Find(fn func(NullableSeriesRequest, int) bool) *NullableSeriesRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableSeriesRequestStream) FindOr(fn func(NullableSeriesRequest, int) bool, or NullableSeriesRequest) NullableSeriesRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableSeriesRequestStream) FindIndex(fn func(NullableSeriesRequest, int) bool) int {
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
func (self *NullableSeriesRequestStream) First() *NullableSeriesRequest {
	return self.Get(0)
}
func (self *NullableSeriesRequestStream) FirstOr(arg NullableSeriesRequest) NullableSeriesRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSeriesRequestStream) ForEach(fn func(NullableSeriesRequest, int)) *NullableSeriesRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableSeriesRequestStream) ForEachRight(fn func(NullableSeriesRequest, int)) *NullableSeriesRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableSeriesRequestStream) GroupBy(fn func(NullableSeriesRequest, int) string) map[string][]NullableSeriesRequest {
	m := map[string][]NullableSeriesRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableSeriesRequestStream) GroupByValues(fn func(NullableSeriesRequest, int) string) [][]NullableSeriesRequest {
	var tmp [][]NullableSeriesRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableSeriesRequestStream) IndexOf(arg NullableSeriesRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableSeriesRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableSeriesRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableSeriesRequestStream) Last() *NullableSeriesRequest {
	return self.Get(self.Len() - 1)
}
func (self *NullableSeriesRequestStream) LastOr(arg NullableSeriesRequest) NullableSeriesRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSeriesRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableSeriesRequestStream) Limit(limit int) *NullableSeriesRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableSeriesRequestStream) Map(fn func(NullableSeriesRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesRequestStream) Map2Int(fn func(NullableSeriesRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesRequestStream) Map2Int32(fn func(NullableSeriesRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesRequestStream) Map2Int64(fn func(NullableSeriesRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesRequestStream) Map2Float32(fn func(NullableSeriesRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesRequestStream) Map2Float64(fn func(NullableSeriesRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesRequestStream) Map2Bool(fn func(NullableSeriesRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesRequestStream) Map2Bytes(fn func(NullableSeriesRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesRequestStream) Map2String(fn func(NullableSeriesRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesRequestStream) Max(fn func(NullableSeriesRequest, int) float64) *NullableSeriesRequest {
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
func (self *NullableSeriesRequestStream) Min(fn func(NullableSeriesRequest, int) float64) *NullableSeriesRequest {
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
func (self *NullableSeriesRequestStream) NoneMatch(fn func(NullableSeriesRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableSeriesRequestStream) Get(index int) *NullableSeriesRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableSeriesRequestStream) GetOr(index int, arg NullableSeriesRequest) NullableSeriesRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSeriesRequestStream) Peek(fn func(*NullableSeriesRequest, int)) *NullableSeriesRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableSeriesRequestStream) Reduce(fn func(NullableSeriesRequest, NullableSeriesRequest, int) NullableSeriesRequest) *NullableSeriesRequestStream {
	return self.ReduceInit(fn, NullableSeriesRequest{})
}
func (self *NullableSeriesRequestStream) ReduceInit(fn func(NullableSeriesRequest, NullableSeriesRequest, int) NullableSeriesRequest, initialValue NullableSeriesRequest) *NullableSeriesRequestStream {
	result := NullableSeriesRequestStreamOf()
	self.ForEach(func(v NullableSeriesRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableSeriesRequestStream) ReduceInterface(fn func(interface{}, NullableSeriesRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableSeriesRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableSeriesRequestStream) ReduceString(fn func(string, NullableSeriesRequest, int) string) []string {
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
func (self *NullableSeriesRequestStream) ReduceInt(fn func(int, NullableSeriesRequest, int) int) []int {
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
func (self *NullableSeriesRequestStream) ReduceInt32(fn func(int32, NullableSeriesRequest, int) int32) []int32 {
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
func (self *NullableSeriesRequestStream) ReduceInt64(fn func(int64, NullableSeriesRequest, int) int64) []int64 {
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
func (self *NullableSeriesRequestStream) ReduceFloat32(fn func(float32, NullableSeriesRequest, int) float32) []float32 {
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
func (self *NullableSeriesRequestStream) ReduceFloat64(fn func(float64, NullableSeriesRequest, int) float64) []float64 {
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
func (self *NullableSeriesRequestStream) ReduceBool(fn func(bool, NullableSeriesRequest, int) bool) []bool {
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
func (self *NullableSeriesRequestStream) Reverse() *NullableSeriesRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableSeriesRequestStream) Replace(fn func(NullableSeriesRequest, int) NullableSeriesRequest) *NullableSeriesRequestStream {
	return self.ForEach(func(v NullableSeriesRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableSeriesRequestStream) Select(fn func(NullableSeriesRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableSeriesRequestStream) Set(index int, val NullableSeriesRequest) *NullableSeriesRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableSeriesRequestStream) Skip(skip int) *NullableSeriesRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableSeriesRequestStream) SkippingEach(fn func(NullableSeriesRequest, int) int) *NullableSeriesRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableSeriesRequestStream) Slice(startIndex, n int) *NullableSeriesRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableSeriesRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableSeriesRequestStream) Sort(fn func(i, j int) bool) *NullableSeriesRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableSeriesRequestStream) Tail() *NullableSeriesRequest {
	return self.Last()
}
func (self *NullableSeriesRequestStream) TailOr(arg NullableSeriesRequest) NullableSeriesRequest {
	return self.LastOr(arg)
}
func (self *NullableSeriesRequestStream) ToList() []NullableSeriesRequest {
	return self.Val()
}
func (self *NullableSeriesRequestStream) Unique() *NullableSeriesRequestStream {
	return self.Distinct()
}
func (self *NullableSeriesRequestStream) Val() []NullableSeriesRequest {
	if self == nil {
		return []NullableSeriesRequest{}
	}
	return *self.Copy()
}
func (self *NullableSeriesRequestStream) While(fn func(NullableSeriesRequest, int) bool) *NullableSeriesRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableSeriesRequestStream) Where(fn func(NullableSeriesRequest) bool) *NullableSeriesRequestStream {
	result := NullableSeriesRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSeriesRequestStream) WhereSlim(fn func(NullableSeriesRequest) bool) *NullableSeriesRequestStream {
	result := NullableSeriesRequestStreamOf()
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
