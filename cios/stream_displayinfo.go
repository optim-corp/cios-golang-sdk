/*
 * Collection utility of DisplayInfo Struct
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

type DisplayInfoStream []DisplayInfo

func DisplayInfoStreamOf(arg ...DisplayInfo) DisplayInfoStream {
	return arg
}
func DisplayInfoStreamFrom(arg []DisplayInfo) DisplayInfoStream {
	return arg
}
func CreateDisplayInfoStream(arg ...DisplayInfo) *DisplayInfoStream {
	tmp := DisplayInfoStreamOf(arg...)
	return &tmp
}
func GenerateDisplayInfoStream(arg []DisplayInfo) *DisplayInfoStream {
	tmp := DisplayInfoStreamFrom(arg)
	return &tmp
}

func (self *DisplayInfoStream) Add(arg DisplayInfo) *DisplayInfoStream {
	return self.AddAll(arg)
}
func (self *DisplayInfoStream) AddAll(arg ...DisplayInfo) *DisplayInfoStream {
	*self = append(*self, arg...)
	return self
}
func (self *DisplayInfoStream) AddSafe(arg *DisplayInfo) *DisplayInfoStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *DisplayInfoStream) Aggregate(fn func(DisplayInfo, DisplayInfo) DisplayInfo) *DisplayInfoStream {
	result := DisplayInfoStreamOf()
	self.ForEach(func(v DisplayInfo, i int) {
		if i == 0 {
			result.Add(fn(DisplayInfo{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *DisplayInfoStream) AllMatch(fn func(DisplayInfo, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *DisplayInfoStream) AnyMatch(fn func(DisplayInfo, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *DisplayInfoStream) Clone() *DisplayInfoStream {
	temp := make([]DisplayInfo, self.Len())
	copy(temp, *self)
	return (*DisplayInfoStream)(&temp)
}
func (self *DisplayInfoStream) Copy() *DisplayInfoStream {
	return self.Clone()
}
func (self *DisplayInfoStream) Concat(arg []DisplayInfo) *DisplayInfoStream {
	return self.AddAll(arg...)
}
func (self *DisplayInfoStream) Contains(arg DisplayInfo) bool {
	return self.FindIndex(func(_arg DisplayInfo, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *DisplayInfoStream) Clean() *DisplayInfoStream {
	*self = DisplayInfoStreamOf()
	return self
}
func (self *DisplayInfoStream) Delete(index int) *DisplayInfoStream {
	return self.DeleteRange(index, index)
}
func (self *DisplayInfoStream) DeleteRange(startIndex, endIndex int) *DisplayInfoStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *DisplayInfoStream) Distinct() *DisplayInfoStream {
	caches := map[string]bool{}
	result := DisplayInfoStreamOf()
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
func (self *DisplayInfoStream) Each(fn func(DisplayInfo)) *DisplayInfoStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *DisplayInfoStream) EachRight(fn func(DisplayInfo)) *DisplayInfoStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *DisplayInfoStream) Equals(arr []DisplayInfo) bool {
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
func (self *DisplayInfoStream) Filter(fn func(DisplayInfo, int) bool) *DisplayInfoStream {
	result := DisplayInfoStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *DisplayInfoStream) FilterSlim(fn func(DisplayInfo, int) bool) *DisplayInfoStream {
	result := DisplayInfoStreamOf()
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
func (self *DisplayInfoStream) Find(fn func(DisplayInfo, int) bool) *DisplayInfo {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *DisplayInfoStream) FindOr(fn func(DisplayInfo, int) bool, or DisplayInfo) DisplayInfo {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *DisplayInfoStream) FindIndex(fn func(DisplayInfo, int) bool) int {
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
func (self *DisplayInfoStream) First() *DisplayInfo {
	return self.Get(0)
}
func (self *DisplayInfoStream) FirstOr(arg DisplayInfo) DisplayInfo {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *DisplayInfoStream) ForEach(fn func(DisplayInfo, int)) *DisplayInfoStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *DisplayInfoStream) ForEachRight(fn func(DisplayInfo, int)) *DisplayInfoStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *DisplayInfoStream) GroupBy(fn func(DisplayInfo, int) string) map[string][]DisplayInfo {
	m := map[string][]DisplayInfo{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *DisplayInfoStream) GroupByValues(fn func(DisplayInfo, int) string) [][]DisplayInfo {
	var tmp [][]DisplayInfo
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *DisplayInfoStream) IndexOf(arg DisplayInfo) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *DisplayInfoStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *DisplayInfoStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *DisplayInfoStream) Last() *DisplayInfo {
	return self.Get(self.Len() - 1)
}
func (self *DisplayInfoStream) LastOr(arg DisplayInfo) DisplayInfo {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *DisplayInfoStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *DisplayInfoStream) Limit(limit int) *DisplayInfoStream {
	self.Slice(0, limit)
	return self
}

func (self *DisplayInfoStream) Map(fn func(DisplayInfo, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DisplayInfoStream) Map2Int(fn func(DisplayInfo, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DisplayInfoStream) Map2Int32(fn func(DisplayInfo, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DisplayInfoStream) Map2Int64(fn func(DisplayInfo, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DisplayInfoStream) Map2Float32(fn func(DisplayInfo, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DisplayInfoStream) Map2Float64(fn func(DisplayInfo, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DisplayInfoStream) Map2Bool(fn func(DisplayInfo, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DisplayInfoStream) Map2Bytes(fn func(DisplayInfo, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DisplayInfoStream) Map2String(fn func(DisplayInfo, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DisplayInfoStream) Max(fn func(DisplayInfo, int) float64) *DisplayInfo {
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
func (self *DisplayInfoStream) Min(fn func(DisplayInfo, int) float64) *DisplayInfo {
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
func (self *DisplayInfoStream) NoneMatch(fn func(DisplayInfo, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *DisplayInfoStream) Get(index int) *DisplayInfo {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *DisplayInfoStream) GetOr(index int, arg DisplayInfo) DisplayInfo {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *DisplayInfoStream) Peek(fn func(*DisplayInfo, int)) *DisplayInfoStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *DisplayInfoStream) Reduce(fn func(DisplayInfo, DisplayInfo, int) DisplayInfo) *DisplayInfoStream {
	return self.ReduceInit(fn, DisplayInfo{})
}
func (self *DisplayInfoStream) ReduceInit(fn func(DisplayInfo, DisplayInfo, int) DisplayInfo, initialValue DisplayInfo) *DisplayInfoStream {
	result := DisplayInfoStreamOf()
	self.ForEach(func(v DisplayInfo, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *DisplayInfoStream) ReduceInterface(fn func(interface{}, DisplayInfo, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(DisplayInfo{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *DisplayInfoStream) ReduceString(fn func(string, DisplayInfo, int) string) []string {
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
func (self *DisplayInfoStream) ReduceInt(fn func(int, DisplayInfo, int) int) []int {
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
func (self *DisplayInfoStream) ReduceInt32(fn func(int32, DisplayInfo, int) int32) []int32 {
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
func (self *DisplayInfoStream) ReduceInt64(fn func(int64, DisplayInfo, int) int64) []int64 {
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
func (self *DisplayInfoStream) ReduceFloat32(fn func(float32, DisplayInfo, int) float32) []float32 {
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
func (self *DisplayInfoStream) ReduceFloat64(fn func(float64, DisplayInfo, int) float64) []float64 {
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
func (self *DisplayInfoStream) ReduceBool(fn func(bool, DisplayInfo, int) bool) []bool {
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
func (self *DisplayInfoStream) Reverse() *DisplayInfoStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *DisplayInfoStream) Replace(fn func(DisplayInfo, int) DisplayInfo) *DisplayInfoStream {
	return self.ForEach(func(v DisplayInfo, i int) { self.Set(i, fn(v, i)) })
}
func (self *DisplayInfoStream) Select(fn func(DisplayInfo) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *DisplayInfoStream) Set(index int, val DisplayInfo) *DisplayInfoStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *DisplayInfoStream) Skip(skip int) *DisplayInfoStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *DisplayInfoStream) SkippingEach(fn func(DisplayInfo, int) int) *DisplayInfoStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *DisplayInfoStream) Slice(startIndex, n int) *DisplayInfoStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []DisplayInfo{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *DisplayInfoStream) Sort(fn func(i, j int) bool) *DisplayInfoStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *DisplayInfoStream) Tail() *DisplayInfo {
	return self.Last()
}
func (self *DisplayInfoStream) TailOr(arg DisplayInfo) DisplayInfo {
	return self.LastOr(arg)
}
func (self *DisplayInfoStream) ToList() []DisplayInfo {
	return self.Val()
}
func (self *DisplayInfoStream) Unique() *DisplayInfoStream {
	return self.Distinct()
}
func (self *DisplayInfoStream) Val() []DisplayInfo {
	if self == nil {
		return []DisplayInfo{}
	}
	return *self.Copy()
}
func (self *DisplayInfoStream) While(fn func(DisplayInfo, int) bool) *DisplayInfoStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *DisplayInfoStream) Where(fn func(DisplayInfo) bool) *DisplayInfoStream {
	result := DisplayInfoStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *DisplayInfoStream) WhereSlim(fn func(DisplayInfo) bool) *DisplayInfoStream {
	result := DisplayInfoStreamOf()
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
