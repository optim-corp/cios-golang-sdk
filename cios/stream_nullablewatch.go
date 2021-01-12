/*
 * Collection utility of NullableWatch Struct
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

type NullableWatchStream []NullableWatch

func NullableWatchStreamOf(arg ...NullableWatch) NullableWatchStream {
	return arg
}
func NullableWatchStreamFrom(arg []NullableWatch) NullableWatchStream {
	return arg
}
func CreateNullableWatchStream(arg ...NullableWatch) *NullableWatchStream {
	tmp := NullableWatchStreamOf(arg...)
	return &tmp
}
func GenerateNullableWatchStream(arg []NullableWatch) *NullableWatchStream {
	tmp := NullableWatchStreamFrom(arg)
	return &tmp
}

func (self *NullableWatchStream) Add(arg NullableWatch) *NullableWatchStream {
	return self.AddAll(arg)
}
func (self *NullableWatchStream) AddAll(arg ...NullableWatch) *NullableWatchStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableWatchStream) AddSafe(arg *NullableWatch) *NullableWatchStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableWatchStream) Aggregate(fn func(NullableWatch, NullableWatch) NullableWatch) *NullableWatchStream {
	result := NullableWatchStreamOf()
	self.ForEach(func(v NullableWatch, i int) {
		if i == 0 {
			result.Add(fn(NullableWatch{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableWatchStream) AllMatch(fn func(NullableWatch, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableWatchStream) AnyMatch(fn func(NullableWatch, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableWatchStream) Clone() *NullableWatchStream {
	temp := make([]NullableWatch, self.Len())
	copy(temp, *self)
	return (*NullableWatchStream)(&temp)
}
func (self *NullableWatchStream) Copy() *NullableWatchStream {
	return self.Clone()
}
func (self *NullableWatchStream) Concat(arg []NullableWatch) *NullableWatchStream {
	return self.AddAll(arg...)
}
func (self *NullableWatchStream) Contains(arg NullableWatch) bool {
	return self.FindIndex(func(_arg NullableWatch, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableWatchStream) Clean() *NullableWatchStream {
	*self = NullableWatchStreamOf()
	return self
}
func (self *NullableWatchStream) Delete(index int) *NullableWatchStream {
	return self.DeleteRange(index, index)
}
func (self *NullableWatchStream) DeleteRange(startIndex, endIndex int) *NullableWatchStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableWatchStream) Distinct() *NullableWatchStream {
	caches := map[string]bool{}
	result := NullableWatchStreamOf()
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
func (self *NullableWatchStream) Each(fn func(NullableWatch)) *NullableWatchStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableWatchStream) EachRight(fn func(NullableWatch)) *NullableWatchStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableWatchStream) Equals(arr []NullableWatch) bool {
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
func (self *NullableWatchStream) Filter(fn func(NullableWatch, int) bool) *NullableWatchStream {
	result := NullableWatchStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableWatchStream) FilterSlim(fn func(NullableWatch, int) bool) *NullableWatchStream {
	result := NullableWatchStreamOf()
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
func (self *NullableWatchStream) Find(fn func(NullableWatch, int) bool) *NullableWatch {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableWatchStream) FindOr(fn func(NullableWatch, int) bool, or NullableWatch) NullableWatch {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableWatchStream) FindIndex(fn func(NullableWatch, int) bool) int {
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
func (self *NullableWatchStream) First() *NullableWatch {
	return self.Get(0)
}
func (self *NullableWatchStream) FirstOr(arg NullableWatch) NullableWatch {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableWatchStream) ForEach(fn func(NullableWatch, int)) *NullableWatchStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableWatchStream) ForEachRight(fn func(NullableWatch, int)) *NullableWatchStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableWatchStream) GroupBy(fn func(NullableWatch, int) string) map[string][]NullableWatch {
	m := map[string][]NullableWatch{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableWatchStream) GroupByValues(fn func(NullableWatch, int) string) [][]NullableWatch {
	var tmp [][]NullableWatch
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableWatchStream) IndexOf(arg NullableWatch) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableWatchStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableWatchStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableWatchStream) Last() *NullableWatch {
	return self.Get(self.Len() - 1)
}
func (self *NullableWatchStream) LastOr(arg NullableWatch) NullableWatch {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableWatchStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableWatchStream) Limit(limit int) *NullableWatchStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableWatchStream) Map(fn func(NullableWatch, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableWatchStream) Map2Int(fn func(NullableWatch, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableWatchStream) Map2Int32(fn func(NullableWatch, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableWatchStream) Map2Int64(fn func(NullableWatch, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableWatchStream) Map2Float32(fn func(NullableWatch, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableWatchStream) Map2Float64(fn func(NullableWatch, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableWatchStream) Map2Bool(fn func(NullableWatch, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableWatchStream) Map2Bytes(fn func(NullableWatch, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableWatchStream) Map2String(fn func(NullableWatch, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableWatchStream) Max(fn func(NullableWatch, int) float64) *NullableWatch {
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
func (self *NullableWatchStream) Min(fn func(NullableWatch, int) float64) *NullableWatch {
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
func (self *NullableWatchStream) NoneMatch(fn func(NullableWatch, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableWatchStream) Get(index int) *NullableWatch {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableWatchStream) GetOr(index int, arg NullableWatch) NullableWatch {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableWatchStream) Peek(fn func(*NullableWatch, int)) *NullableWatchStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableWatchStream) Reduce(fn func(NullableWatch, NullableWatch, int) NullableWatch) *NullableWatchStream {
	return self.ReduceInit(fn, NullableWatch{})
}
func (self *NullableWatchStream) ReduceInit(fn func(NullableWatch, NullableWatch, int) NullableWatch, initialValue NullableWatch) *NullableWatchStream {
	result := NullableWatchStreamOf()
	self.ForEach(func(v NullableWatch, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableWatchStream) ReduceInterface(fn func(interface{}, NullableWatch, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableWatch{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableWatchStream) ReduceString(fn func(string, NullableWatch, int) string) []string {
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
func (self *NullableWatchStream) ReduceInt(fn func(int, NullableWatch, int) int) []int {
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
func (self *NullableWatchStream) ReduceInt32(fn func(int32, NullableWatch, int) int32) []int32 {
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
func (self *NullableWatchStream) ReduceInt64(fn func(int64, NullableWatch, int) int64) []int64 {
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
func (self *NullableWatchStream) ReduceFloat32(fn func(float32, NullableWatch, int) float32) []float32 {
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
func (self *NullableWatchStream) ReduceFloat64(fn func(float64, NullableWatch, int) float64) []float64 {
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
func (self *NullableWatchStream) ReduceBool(fn func(bool, NullableWatch, int) bool) []bool {
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
func (self *NullableWatchStream) Reverse() *NullableWatchStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableWatchStream) Replace(fn func(NullableWatch, int) NullableWatch) *NullableWatchStream {
	return self.ForEach(func(v NullableWatch, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableWatchStream) Select(fn func(NullableWatch) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableWatchStream) Set(index int, val NullableWatch) *NullableWatchStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableWatchStream) Skip(skip int) *NullableWatchStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableWatchStream) SkippingEach(fn func(NullableWatch, int) int) *NullableWatchStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableWatchStream) Slice(startIndex, n int) *NullableWatchStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableWatch{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableWatchStream) Sort(fn func(i, j int) bool) *NullableWatchStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableWatchStream) Tail() *NullableWatch {
	return self.Last()
}
func (self *NullableWatchStream) TailOr(arg NullableWatch) NullableWatch {
	return self.LastOr(arg)
}
func (self *NullableWatchStream) ToList() []NullableWatch {
	return self.Val()
}
func (self *NullableWatchStream) Unique() *NullableWatchStream {
	return self.Distinct()
}
func (self *NullableWatchStream) Val() []NullableWatch {
	if self == nil {
		return []NullableWatch{}
	}
	return *self.Copy()
}
func (self *NullableWatchStream) While(fn func(NullableWatch, int) bool) *NullableWatchStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableWatchStream) Where(fn func(NullableWatch) bool) *NullableWatchStream {
	result := NullableWatchStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableWatchStream) WhereSlim(fn func(NullableWatch) bool) *NullableWatchStream {
	result := NullableWatchStreamOf()
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
