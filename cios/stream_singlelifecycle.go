/*
 * Collection utility of SingleLifeCycle Struct
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

type SingleLifeCycleStream []SingleLifeCycle

func SingleLifeCycleStreamOf(arg ...SingleLifeCycle) SingleLifeCycleStream {
	return arg
}
func SingleLifeCycleStreamFrom(arg []SingleLifeCycle) SingleLifeCycleStream {
	return arg
}
func CreateSingleLifeCycleStream(arg ...SingleLifeCycle) *SingleLifeCycleStream {
	tmp := SingleLifeCycleStreamOf(arg...)
	return &tmp
}
func GenerateSingleLifeCycleStream(arg []SingleLifeCycle) *SingleLifeCycleStream {
	tmp := SingleLifeCycleStreamFrom(arg)
	return &tmp
}

func (self *SingleLifeCycleStream) Add(arg SingleLifeCycle) *SingleLifeCycleStream {
	return self.AddAll(arg)
}
func (self *SingleLifeCycleStream) AddAll(arg ...SingleLifeCycle) *SingleLifeCycleStream {
	*self = append(*self, arg...)
	return self
}
func (self *SingleLifeCycleStream) AddSafe(arg *SingleLifeCycle) *SingleLifeCycleStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *SingleLifeCycleStream) Aggregate(fn func(SingleLifeCycle, SingleLifeCycle) SingleLifeCycle) *SingleLifeCycleStream {
	result := SingleLifeCycleStreamOf()
	self.ForEach(func(v SingleLifeCycle, i int) {
		if i == 0 {
			result.Add(fn(SingleLifeCycle{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *SingleLifeCycleStream) AllMatch(fn func(SingleLifeCycle, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *SingleLifeCycleStream) AnyMatch(fn func(SingleLifeCycle, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *SingleLifeCycleStream) Clone() *SingleLifeCycleStream {
	temp := make([]SingleLifeCycle, self.Len())
	copy(temp, *self)
	return (*SingleLifeCycleStream)(&temp)
}
func (self *SingleLifeCycleStream) Copy() *SingleLifeCycleStream {
	return self.Clone()
}
func (self *SingleLifeCycleStream) Concat(arg []SingleLifeCycle) *SingleLifeCycleStream {
	return self.AddAll(arg...)
}
func (self *SingleLifeCycleStream) Contains(arg SingleLifeCycle) bool {
	return self.FindIndex(func(_arg SingleLifeCycle, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *SingleLifeCycleStream) Clean() *SingleLifeCycleStream {
	*self = SingleLifeCycleStreamOf()
	return self
}
func (self *SingleLifeCycleStream) Delete(index int) *SingleLifeCycleStream {
	return self.DeleteRange(index, index)
}
func (self *SingleLifeCycleStream) DeleteRange(startIndex, endIndex int) *SingleLifeCycleStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *SingleLifeCycleStream) Distinct() *SingleLifeCycleStream {
	caches := map[string]bool{}
	result := SingleLifeCycleStreamOf()
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
func (self *SingleLifeCycleStream) Each(fn func(SingleLifeCycle)) *SingleLifeCycleStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *SingleLifeCycleStream) EachRight(fn func(SingleLifeCycle)) *SingleLifeCycleStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *SingleLifeCycleStream) Equals(arr []SingleLifeCycle) bool {
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
func (self *SingleLifeCycleStream) Filter(fn func(SingleLifeCycle, int) bool) *SingleLifeCycleStream {
	result := SingleLifeCycleStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *SingleLifeCycleStream) FilterSlim(fn func(SingleLifeCycle, int) bool) *SingleLifeCycleStream {
	result := SingleLifeCycleStreamOf()
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
func (self *SingleLifeCycleStream) Find(fn func(SingleLifeCycle, int) bool) *SingleLifeCycle {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *SingleLifeCycleStream) FindOr(fn func(SingleLifeCycle, int) bool, or SingleLifeCycle) SingleLifeCycle {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *SingleLifeCycleStream) FindIndex(fn func(SingleLifeCycle, int) bool) int {
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
func (self *SingleLifeCycleStream) First() *SingleLifeCycle {
	return self.Get(0)
}
func (self *SingleLifeCycleStream) FirstOr(arg SingleLifeCycle) SingleLifeCycle {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *SingleLifeCycleStream) ForEach(fn func(SingleLifeCycle, int)) *SingleLifeCycleStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *SingleLifeCycleStream) ForEachRight(fn func(SingleLifeCycle, int)) *SingleLifeCycleStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *SingleLifeCycleStream) GroupBy(fn func(SingleLifeCycle, int) string) map[string][]SingleLifeCycle {
	m := map[string][]SingleLifeCycle{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *SingleLifeCycleStream) GroupByValues(fn func(SingleLifeCycle, int) string) [][]SingleLifeCycle {
	var tmp [][]SingleLifeCycle
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *SingleLifeCycleStream) IndexOf(arg SingleLifeCycle) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *SingleLifeCycleStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *SingleLifeCycleStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *SingleLifeCycleStream) Last() *SingleLifeCycle {
	return self.Get(self.Len() - 1)
}
func (self *SingleLifeCycleStream) LastOr(arg SingleLifeCycle) SingleLifeCycle {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *SingleLifeCycleStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *SingleLifeCycleStream) Limit(limit int) *SingleLifeCycleStream {
	self.Slice(0, limit)
	return self
}

func (self *SingleLifeCycleStream) Map(fn func(SingleLifeCycle, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleLifeCycleStream) Map2Int(fn func(SingleLifeCycle, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleLifeCycleStream) Map2Int32(fn func(SingleLifeCycle, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleLifeCycleStream) Map2Int64(fn func(SingleLifeCycle, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleLifeCycleStream) Map2Float32(fn func(SingleLifeCycle, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleLifeCycleStream) Map2Float64(fn func(SingleLifeCycle, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleLifeCycleStream) Map2Bool(fn func(SingleLifeCycle, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleLifeCycleStream) Map2Bytes(fn func(SingleLifeCycle, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleLifeCycleStream) Map2String(fn func(SingleLifeCycle, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleLifeCycleStream) Max(fn func(SingleLifeCycle, int) float64) *SingleLifeCycle {
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
func (self *SingleLifeCycleStream) Min(fn func(SingleLifeCycle, int) float64) *SingleLifeCycle {
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
func (self *SingleLifeCycleStream) NoneMatch(fn func(SingleLifeCycle, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *SingleLifeCycleStream) Get(index int) *SingleLifeCycle {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *SingleLifeCycleStream) GetOr(index int, arg SingleLifeCycle) SingleLifeCycle {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *SingleLifeCycleStream) Peek(fn func(*SingleLifeCycle, int)) *SingleLifeCycleStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *SingleLifeCycleStream) Reduce(fn func(SingleLifeCycle, SingleLifeCycle, int) SingleLifeCycle) *SingleLifeCycleStream {
	return self.ReduceInit(fn, SingleLifeCycle{})
}
func (self *SingleLifeCycleStream) ReduceInit(fn func(SingleLifeCycle, SingleLifeCycle, int) SingleLifeCycle, initialValue SingleLifeCycle) *SingleLifeCycleStream {
	result := SingleLifeCycleStreamOf()
	self.ForEach(func(v SingleLifeCycle, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *SingleLifeCycleStream) ReduceInterface(fn func(interface{}, SingleLifeCycle, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(SingleLifeCycle{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *SingleLifeCycleStream) ReduceString(fn func(string, SingleLifeCycle, int) string) []string {
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
func (self *SingleLifeCycleStream) ReduceInt(fn func(int, SingleLifeCycle, int) int) []int {
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
func (self *SingleLifeCycleStream) ReduceInt32(fn func(int32, SingleLifeCycle, int) int32) []int32 {
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
func (self *SingleLifeCycleStream) ReduceInt64(fn func(int64, SingleLifeCycle, int) int64) []int64 {
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
func (self *SingleLifeCycleStream) ReduceFloat32(fn func(float32, SingleLifeCycle, int) float32) []float32 {
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
func (self *SingleLifeCycleStream) ReduceFloat64(fn func(float64, SingleLifeCycle, int) float64) []float64 {
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
func (self *SingleLifeCycleStream) ReduceBool(fn func(bool, SingleLifeCycle, int) bool) []bool {
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
func (self *SingleLifeCycleStream) Reverse() *SingleLifeCycleStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *SingleLifeCycleStream) Replace(fn func(SingleLifeCycle, int) SingleLifeCycle) *SingleLifeCycleStream {
	return self.ForEach(func(v SingleLifeCycle, i int) { self.Set(i, fn(v, i)) })
}
func (self *SingleLifeCycleStream) Select(fn func(SingleLifeCycle) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *SingleLifeCycleStream) Set(index int, val SingleLifeCycle) *SingleLifeCycleStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *SingleLifeCycleStream) Skip(skip int) *SingleLifeCycleStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *SingleLifeCycleStream) SkippingEach(fn func(SingleLifeCycle, int) int) *SingleLifeCycleStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *SingleLifeCycleStream) Slice(startIndex, n int) *SingleLifeCycleStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []SingleLifeCycle{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *SingleLifeCycleStream) Sort(fn func(i, j int) bool) *SingleLifeCycleStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *SingleLifeCycleStream) Tail() *SingleLifeCycle {
	return self.Last()
}
func (self *SingleLifeCycleStream) TailOr(arg SingleLifeCycle) SingleLifeCycle {
	return self.LastOr(arg)
}
func (self *SingleLifeCycleStream) ToList() []SingleLifeCycle {
	return self.Val()
}
func (self *SingleLifeCycleStream) Unique() *SingleLifeCycleStream {
	return self.Distinct()
}
func (self *SingleLifeCycleStream) Val() []SingleLifeCycle {
	if self == nil {
		return []SingleLifeCycle{}
	}
	return *self.Copy()
}
func (self *SingleLifeCycleStream) While(fn func(SingleLifeCycle, int) bool) *SingleLifeCycleStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *SingleLifeCycleStream) Where(fn func(SingleLifeCycle) bool) *SingleLifeCycleStream {
	result := SingleLifeCycleStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *SingleLifeCycleStream) WhereSlim(fn func(SingleLifeCycle) bool) *SingleLifeCycleStream {
	result := SingleLifeCycleStreamOf()
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
