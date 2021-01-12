/*
 * Collection utility of CollectionLocation Struct
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

type CollectionLocationStream []CollectionLocation

func CollectionLocationStreamOf(arg ...CollectionLocation) CollectionLocationStream {
	return arg
}
func CollectionLocationStreamFrom(arg []CollectionLocation) CollectionLocationStream {
	return arg
}
func CreateCollectionLocationStream(arg ...CollectionLocation) *CollectionLocationStream {
	tmp := CollectionLocationStreamOf(arg...)
	return &tmp
}
func GenerateCollectionLocationStream(arg []CollectionLocation) *CollectionLocationStream {
	tmp := CollectionLocationStreamFrom(arg)
	return &tmp
}

func (self *CollectionLocationStream) Add(arg CollectionLocation) *CollectionLocationStream {
	return self.AddAll(arg)
}
func (self *CollectionLocationStream) AddAll(arg ...CollectionLocation) *CollectionLocationStream {
	*self = append(*self, arg...)
	return self
}
func (self *CollectionLocationStream) AddSafe(arg *CollectionLocation) *CollectionLocationStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *CollectionLocationStream) Aggregate(fn func(CollectionLocation, CollectionLocation) CollectionLocation) *CollectionLocationStream {
	result := CollectionLocationStreamOf()
	self.ForEach(func(v CollectionLocation, i int) {
		if i == 0 {
			result.Add(fn(CollectionLocation{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *CollectionLocationStream) AllMatch(fn func(CollectionLocation, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *CollectionLocationStream) AnyMatch(fn func(CollectionLocation, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *CollectionLocationStream) Clone() *CollectionLocationStream {
	temp := make([]CollectionLocation, self.Len())
	copy(temp, *self)
	return (*CollectionLocationStream)(&temp)
}
func (self *CollectionLocationStream) Copy() *CollectionLocationStream {
	return self.Clone()
}
func (self *CollectionLocationStream) Concat(arg []CollectionLocation) *CollectionLocationStream {
	return self.AddAll(arg...)
}
func (self *CollectionLocationStream) Contains(arg CollectionLocation) bool {
	return self.FindIndex(func(_arg CollectionLocation, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *CollectionLocationStream) Clean() *CollectionLocationStream {
	*self = CollectionLocationStreamOf()
	return self
}
func (self *CollectionLocationStream) Delete(index int) *CollectionLocationStream {
	return self.DeleteRange(index, index)
}
func (self *CollectionLocationStream) DeleteRange(startIndex, endIndex int) *CollectionLocationStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *CollectionLocationStream) Distinct() *CollectionLocationStream {
	caches := map[string]bool{}
	result := CollectionLocationStreamOf()
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
func (self *CollectionLocationStream) Each(fn func(CollectionLocation)) *CollectionLocationStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *CollectionLocationStream) EachRight(fn func(CollectionLocation)) *CollectionLocationStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *CollectionLocationStream) Equals(arr []CollectionLocation) bool {
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
func (self *CollectionLocationStream) Filter(fn func(CollectionLocation, int) bool) *CollectionLocationStream {
	result := CollectionLocationStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *CollectionLocationStream) FilterSlim(fn func(CollectionLocation, int) bool) *CollectionLocationStream {
	result := CollectionLocationStreamOf()
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
func (self *CollectionLocationStream) Find(fn func(CollectionLocation, int) bool) *CollectionLocation {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *CollectionLocationStream) FindOr(fn func(CollectionLocation, int) bool, or CollectionLocation) CollectionLocation {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *CollectionLocationStream) FindIndex(fn func(CollectionLocation, int) bool) int {
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
func (self *CollectionLocationStream) First() *CollectionLocation {
	return self.Get(0)
}
func (self *CollectionLocationStream) FirstOr(arg CollectionLocation) CollectionLocation {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *CollectionLocationStream) ForEach(fn func(CollectionLocation, int)) *CollectionLocationStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *CollectionLocationStream) ForEachRight(fn func(CollectionLocation, int)) *CollectionLocationStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *CollectionLocationStream) GroupBy(fn func(CollectionLocation, int) string) map[string][]CollectionLocation {
	m := map[string][]CollectionLocation{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *CollectionLocationStream) GroupByValues(fn func(CollectionLocation, int) string) [][]CollectionLocation {
	var tmp [][]CollectionLocation
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *CollectionLocationStream) IndexOf(arg CollectionLocation) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *CollectionLocationStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *CollectionLocationStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *CollectionLocationStream) Last() *CollectionLocation {
	return self.Get(self.Len() - 1)
}
func (self *CollectionLocationStream) LastOr(arg CollectionLocation) CollectionLocation {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *CollectionLocationStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *CollectionLocationStream) Limit(limit int) *CollectionLocationStream {
	self.Slice(0, limit)
	return self
}

func (self *CollectionLocationStream) Map(fn func(CollectionLocation, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CollectionLocationStream) Map2Int(fn func(CollectionLocation, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CollectionLocationStream) Map2Int32(fn func(CollectionLocation, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CollectionLocationStream) Map2Int64(fn func(CollectionLocation, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CollectionLocationStream) Map2Float32(fn func(CollectionLocation, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CollectionLocationStream) Map2Float64(fn func(CollectionLocation, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CollectionLocationStream) Map2Bool(fn func(CollectionLocation, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CollectionLocationStream) Map2Bytes(fn func(CollectionLocation, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CollectionLocationStream) Map2String(fn func(CollectionLocation, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CollectionLocationStream) Max(fn func(CollectionLocation, int) float64) *CollectionLocation {
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
func (self *CollectionLocationStream) Min(fn func(CollectionLocation, int) float64) *CollectionLocation {
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
func (self *CollectionLocationStream) NoneMatch(fn func(CollectionLocation, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *CollectionLocationStream) Get(index int) *CollectionLocation {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *CollectionLocationStream) GetOr(index int, arg CollectionLocation) CollectionLocation {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *CollectionLocationStream) Peek(fn func(*CollectionLocation, int)) *CollectionLocationStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *CollectionLocationStream) Reduce(fn func(CollectionLocation, CollectionLocation, int) CollectionLocation) *CollectionLocationStream {
	return self.ReduceInit(fn, CollectionLocation{})
}
func (self *CollectionLocationStream) ReduceInit(fn func(CollectionLocation, CollectionLocation, int) CollectionLocation, initialValue CollectionLocation) *CollectionLocationStream {
	result := CollectionLocationStreamOf()
	self.ForEach(func(v CollectionLocation, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *CollectionLocationStream) ReduceInterface(fn func(interface{}, CollectionLocation, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(CollectionLocation{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *CollectionLocationStream) ReduceString(fn func(string, CollectionLocation, int) string) []string {
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
func (self *CollectionLocationStream) ReduceInt(fn func(int, CollectionLocation, int) int) []int {
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
func (self *CollectionLocationStream) ReduceInt32(fn func(int32, CollectionLocation, int) int32) []int32 {
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
func (self *CollectionLocationStream) ReduceInt64(fn func(int64, CollectionLocation, int) int64) []int64 {
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
func (self *CollectionLocationStream) ReduceFloat32(fn func(float32, CollectionLocation, int) float32) []float32 {
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
func (self *CollectionLocationStream) ReduceFloat64(fn func(float64, CollectionLocation, int) float64) []float64 {
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
func (self *CollectionLocationStream) ReduceBool(fn func(bool, CollectionLocation, int) bool) []bool {
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
func (self *CollectionLocationStream) Reverse() *CollectionLocationStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *CollectionLocationStream) Replace(fn func(CollectionLocation, int) CollectionLocation) *CollectionLocationStream {
	return self.ForEach(func(v CollectionLocation, i int) { self.Set(i, fn(v, i)) })
}
func (self *CollectionLocationStream) Select(fn func(CollectionLocation) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *CollectionLocationStream) Set(index int, val CollectionLocation) *CollectionLocationStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *CollectionLocationStream) Skip(skip int) *CollectionLocationStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *CollectionLocationStream) SkippingEach(fn func(CollectionLocation, int) int) *CollectionLocationStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *CollectionLocationStream) Slice(startIndex, n int) *CollectionLocationStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []CollectionLocation{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *CollectionLocationStream) Sort(fn func(i, j int) bool) *CollectionLocationStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *CollectionLocationStream) Tail() *CollectionLocation {
	return self.Last()
}
func (self *CollectionLocationStream) TailOr(arg CollectionLocation) CollectionLocation {
	return self.LastOr(arg)
}
func (self *CollectionLocationStream) ToList() []CollectionLocation {
	return self.Val()
}
func (self *CollectionLocationStream) Unique() *CollectionLocationStream {
	return self.Distinct()
}
func (self *CollectionLocationStream) Val() []CollectionLocation {
	if self == nil {
		return []CollectionLocation{}
	}
	return *self.Copy()
}
func (self *CollectionLocationStream) While(fn func(CollectionLocation, int) bool) *CollectionLocationStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *CollectionLocationStream) Where(fn func(CollectionLocation) bool) *CollectionLocationStream {
	result := CollectionLocationStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *CollectionLocationStream) WhereSlim(fn func(CollectionLocation) bool) *CollectionLocationStream {
	result := CollectionLocationStreamOf()
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
