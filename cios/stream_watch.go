/*
 * Collection utility of Watch Struct
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

type WatchStream []Watch

func WatchStreamOf(arg ...Watch) WatchStream {
	return arg
}
func WatchStreamFrom(arg []Watch) WatchStream {
	return arg
}
func CreateWatchStream(arg ...Watch) *WatchStream {
	tmp := WatchStreamOf(arg...)
	return &tmp
}
func GenerateWatchStream(arg []Watch) *WatchStream {
	tmp := WatchStreamFrom(arg)
	return &tmp
}

func (self *WatchStream) Add(arg Watch) *WatchStream {
	return self.AddAll(arg)
}
func (self *WatchStream) AddAll(arg ...Watch) *WatchStream {
	*self = append(*self, arg...)
	return self
}
func (self *WatchStream) AddSafe(arg *Watch) *WatchStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *WatchStream) Aggregate(fn func(Watch, Watch) Watch) *WatchStream {
	result := WatchStreamOf()
	self.ForEach(func(v Watch, i int) {
		if i == 0 {
			result.Add(fn(Watch{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *WatchStream) AllMatch(fn func(Watch, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *WatchStream) AnyMatch(fn func(Watch, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *WatchStream) Clone() *WatchStream {
	temp := make([]Watch, self.Len())
	copy(temp, *self)
	return (*WatchStream)(&temp)
}
func (self *WatchStream) Copy() *WatchStream {
	return self.Clone()
}
func (self *WatchStream) Concat(arg []Watch) *WatchStream {
	return self.AddAll(arg...)
}
func (self *WatchStream) Contains(arg Watch) bool {
	return self.FindIndex(func(_arg Watch, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *WatchStream) Clean() *WatchStream {
	*self = WatchStreamOf()
	return self
}
func (self *WatchStream) Delete(index int) *WatchStream {
	return self.DeleteRange(index, index)
}
func (self *WatchStream) DeleteRange(startIndex, endIndex int) *WatchStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *WatchStream) Distinct() *WatchStream {
	caches := map[string]bool{}
	result := WatchStreamOf()
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
func (self *WatchStream) Each(fn func(Watch)) *WatchStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *WatchStream) EachRight(fn func(Watch)) *WatchStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *WatchStream) Equals(arr []Watch) bool {
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
func (self *WatchStream) Filter(fn func(Watch, int) bool) *WatchStream {
	result := WatchStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *WatchStream) FilterSlim(fn func(Watch, int) bool) *WatchStream {
	result := WatchStreamOf()
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
func (self *WatchStream) Find(fn func(Watch, int) bool) *Watch {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *WatchStream) FindOr(fn func(Watch, int) bool, or Watch) Watch {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *WatchStream) FindIndex(fn func(Watch, int) bool) int {
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
func (self *WatchStream) First() *Watch {
	return self.Get(0)
}
func (self *WatchStream) FirstOr(arg Watch) Watch {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *WatchStream) ForEach(fn func(Watch, int)) *WatchStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *WatchStream) ForEachRight(fn func(Watch, int)) *WatchStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *WatchStream) GroupBy(fn func(Watch, int) string) map[string][]Watch {
	m := map[string][]Watch{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *WatchStream) GroupByValues(fn func(Watch, int) string) [][]Watch {
	var tmp [][]Watch
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *WatchStream) IndexOf(arg Watch) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *WatchStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *WatchStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *WatchStream) Last() *Watch {
	return self.Get(self.Len() - 1)
}
func (self *WatchStream) LastOr(arg Watch) Watch {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *WatchStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *WatchStream) Limit(limit int) *WatchStream {
	self.Slice(0, limit)
	return self
}

func (self *WatchStream) Map(fn func(Watch, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *WatchStream) Map2Int(fn func(Watch, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *WatchStream) Map2Int32(fn func(Watch, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *WatchStream) Map2Int64(fn func(Watch, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *WatchStream) Map2Float32(fn func(Watch, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *WatchStream) Map2Float64(fn func(Watch, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *WatchStream) Map2Bool(fn func(Watch, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *WatchStream) Map2Bytes(fn func(Watch, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *WatchStream) Map2String(fn func(Watch, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *WatchStream) Max(fn func(Watch, int) float64) *Watch {
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
func (self *WatchStream) Min(fn func(Watch, int) float64) *Watch {
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
func (self *WatchStream) NoneMatch(fn func(Watch, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *WatchStream) Get(index int) *Watch {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *WatchStream) GetOr(index int, arg Watch) Watch {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *WatchStream) Peek(fn func(*Watch, int)) *WatchStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *WatchStream) Reduce(fn func(Watch, Watch, int) Watch) *WatchStream {
	return self.ReduceInit(fn, Watch{})
}
func (self *WatchStream) ReduceInit(fn func(Watch, Watch, int) Watch, initialValue Watch) *WatchStream {
	result := WatchStreamOf()
	self.ForEach(func(v Watch, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *WatchStream) ReduceInterface(fn func(interface{}, Watch, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(Watch{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *WatchStream) ReduceString(fn func(string, Watch, int) string) []string {
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
func (self *WatchStream) ReduceInt(fn func(int, Watch, int) int) []int {
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
func (self *WatchStream) ReduceInt32(fn func(int32, Watch, int) int32) []int32 {
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
func (self *WatchStream) ReduceInt64(fn func(int64, Watch, int) int64) []int64 {
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
func (self *WatchStream) ReduceFloat32(fn func(float32, Watch, int) float32) []float32 {
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
func (self *WatchStream) ReduceFloat64(fn func(float64, Watch, int) float64) []float64 {
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
func (self *WatchStream) ReduceBool(fn func(bool, Watch, int) bool) []bool {
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
func (self *WatchStream) Reverse() *WatchStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *WatchStream) Replace(fn func(Watch, int) Watch) *WatchStream {
	return self.ForEach(func(v Watch, i int) { self.Set(i, fn(v, i)) })
}
func (self *WatchStream) Select(fn func(Watch) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *WatchStream) Set(index int, val Watch) *WatchStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *WatchStream) Skip(skip int) *WatchStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *WatchStream) SkippingEach(fn func(Watch, int) int) *WatchStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *WatchStream) Slice(startIndex, n int) *WatchStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []Watch{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *WatchStream) Sort(fn func(i, j int) bool) *WatchStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *WatchStream) Tail() *Watch {
	return self.Last()
}
func (self *WatchStream) TailOr(arg Watch) Watch {
	return self.LastOr(arg)
}
func (self *WatchStream) ToList() []Watch {
	return self.Val()
}
func (self *WatchStream) Unique() *WatchStream {
	return self.Distinct()
}
func (self *WatchStream) Val() []Watch {
	if self == nil {
		return []Watch{}
	}
	return *self.Copy()
}
func (self *WatchStream) While(fn func(Watch, int) bool) *WatchStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *WatchStream) Where(fn func(Watch) bool) *WatchStream {
	result := WatchStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *WatchStream) WhereSlim(fn func(Watch) bool) *WatchStream {
	result := WatchStreamOf()
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