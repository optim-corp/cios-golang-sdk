/*
 * Collection utility of Circle Struct
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

type CircleStream []Circle

func CircleStreamOf(arg ...Circle) CircleStream {
	return arg
}
func CircleStreamFrom(arg []Circle) CircleStream {
	return arg
}
func CreateCircleStream(arg ...Circle) *CircleStream {
	tmp := CircleStreamOf(arg...)
	return &tmp
}
func GenerateCircleStream(arg []Circle) *CircleStream {
	tmp := CircleStreamFrom(arg)
	return &tmp
}

func (self *CircleStream) Add(arg Circle) *CircleStream {
	return self.AddAll(arg)
}
func (self *CircleStream) AddAll(arg ...Circle) *CircleStream {
	*self = append(*self, arg...)
	return self
}
func (self *CircleStream) AddSafe(arg *Circle) *CircleStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *CircleStream) Aggregate(fn func(Circle, Circle) Circle) *CircleStream {
	result := CircleStreamOf()
	self.ForEach(func(v Circle, i int) {
		if i == 0 {
			result.Add(fn(Circle{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *CircleStream) AllMatch(fn func(Circle, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *CircleStream) AnyMatch(fn func(Circle, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *CircleStream) Clone() *CircleStream {
	temp := make([]Circle, self.Len())
	copy(temp, *self)
	return (*CircleStream)(&temp)
}
func (self *CircleStream) Copy() *CircleStream {
	return self.Clone()
}
func (self *CircleStream) Concat(arg []Circle) *CircleStream {
	return self.AddAll(arg...)
}
func (self *CircleStream) Contains(arg Circle) bool {
	return self.FindIndex(func(_arg Circle, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *CircleStream) Clean() *CircleStream {
	*self = CircleStreamOf()
	return self
}
func (self *CircleStream) Delete(index int) *CircleStream {
	return self.DeleteRange(index, index)
}
func (self *CircleStream) DeleteRange(startIndex, endIndex int) *CircleStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *CircleStream) Distinct() *CircleStream {
	caches := map[string]bool{}
	result := CircleStreamOf()
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
func (self *CircleStream) Each(fn func(Circle)) *CircleStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *CircleStream) EachRight(fn func(Circle)) *CircleStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *CircleStream) Equals(arr []Circle) bool {
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
func (self *CircleStream) Filter(fn func(Circle, int) bool) *CircleStream {
	result := CircleStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *CircleStream) FilterSlim(fn func(Circle, int) bool) *CircleStream {
	result := CircleStreamOf()
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
func (self *CircleStream) Find(fn func(Circle, int) bool) *Circle {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *CircleStream) FindOr(fn func(Circle, int) bool, or Circle) Circle {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *CircleStream) FindIndex(fn func(Circle, int) bool) int {
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
func (self *CircleStream) First() *Circle {
	return self.Get(0)
}
func (self *CircleStream) FirstOr(arg Circle) Circle {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *CircleStream) ForEach(fn func(Circle, int)) *CircleStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *CircleStream) ForEachRight(fn func(Circle, int)) *CircleStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *CircleStream) GroupBy(fn func(Circle, int) string) map[string][]Circle {
	m := map[string][]Circle{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *CircleStream) GroupByValues(fn func(Circle, int) string) [][]Circle {
	var tmp [][]Circle
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *CircleStream) IndexOf(arg Circle) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *CircleStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *CircleStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *CircleStream) Last() *Circle {
	return self.Get(self.Len() - 1)
}
func (self *CircleStream) LastOr(arg Circle) Circle {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *CircleStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *CircleStream) Limit(limit int) *CircleStream {
	self.Slice(0, limit)
	return self
}

func (self *CircleStream) Map(fn func(Circle, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CircleStream) Map2Int(fn func(Circle, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CircleStream) Map2Int32(fn func(Circle, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CircleStream) Map2Int64(fn func(Circle, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CircleStream) Map2Float32(fn func(Circle, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CircleStream) Map2Float64(fn func(Circle, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CircleStream) Map2Bool(fn func(Circle, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CircleStream) Map2Bytes(fn func(Circle, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CircleStream) Map2String(fn func(Circle, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CircleStream) Max(fn func(Circle, int) float64) *Circle {
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
func (self *CircleStream) Min(fn func(Circle, int) float64) *Circle {
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
func (self *CircleStream) NoneMatch(fn func(Circle, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *CircleStream) Get(index int) *Circle {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *CircleStream) GetOr(index int, arg Circle) Circle {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *CircleStream) Peek(fn func(*Circle, int)) *CircleStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *CircleStream) Reduce(fn func(Circle, Circle, int) Circle) *CircleStream {
	return self.ReduceInit(fn, Circle{})
}
func (self *CircleStream) ReduceInit(fn func(Circle, Circle, int) Circle, initialValue Circle) *CircleStream {
	result := CircleStreamOf()
	self.ForEach(func(v Circle, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *CircleStream) ReduceInterface(fn func(interface{}, Circle, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(Circle{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *CircleStream) ReduceString(fn func(string, Circle, int) string) []string {
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
func (self *CircleStream) ReduceInt(fn func(int, Circle, int) int) []int {
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
func (self *CircleStream) ReduceInt32(fn func(int32, Circle, int) int32) []int32 {
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
func (self *CircleStream) ReduceInt64(fn func(int64, Circle, int) int64) []int64 {
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
func (self *CircleStream) ReduceFloat32(fn func(float32, Circle, int) float32) []float32 {
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
func (self *CircleStream) ReduceFloat64(fn func(float64, Circle, int) float64) []float64 {
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
func (self *CircleStream) ReduceBool(fn func(bool, Circle, int) bool) []bool {
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
func (self *CircleStream) Reverse() *CircleStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *CircleStream) Replace(fn func(Circle, int) Circle) *CircleStream {
	return self.ForEach(func(v Circle, i int) { self.Set(i, fn(v, i)) })
}
func (self *CircleStream) Select(fn func(Circle) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *CircleStream) Set(index int, val Circle) *CircleStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *CircleStream) Skip(skip int) *CircleStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *CircleStream) SkippingEach(fn func(Circle, int) int) *CircleStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *CircleStream) Slice(startIndex, n int) *CircleStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []Circle{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *CircleStream) Sort(fn func(i, j int) bool) *CircleStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *CircleStream) Tail() *Circle {
	return self.Last()
}
func (self *CircleStream) TailOr(arg Circle) Circle {
	return self.LastOr(arg)
}
func (self *CircleStream) ToList() []Circle {
	return self.Val()
}
func (self *CircleStream) Unique() *CircleStream {
	return self.Distinct()
}
func (self *CircleStream) Val() []Circle {
	if self == nil {
		return []Circle{}
	}
	return *self.Copy()
}
func (self *CircleStream) While(fn func(Circle, int) bool) *CircleStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *CircleStream) Where(fn func(Circle) bool) *CircleStream {
	result := CircleStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *CircleStream) WhereSlim(fn func(Circle) bool) *CircleStream {
	result := CircleStreamOf()
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
