/*
 * Collection utility of NullableMultipleGroup Struct
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

type NullableMultipleGroupStream []NullableMultipleGroup

func NullableMultipleGroupStreamOf(arg ...NullableMultipleGroup) NullableMultipleGroupStream {
	return arg
}
func NullableMultipleGroupStreamFrom(arg []NullableMultipleGroup) NullableMultipleGroupStream {
	return arg
}
func CreateNullableMultipleGroupStream(arg ...NullableMultipleGroup) *NullableMultipleGroupStream {
	tmp := NullableMultipleGroupStreamOf(arg...)
	return &tmp
}
func GenerateNullableMultipleGroupStream(arg []NullableMultipleGroup) *NullableMultipleGroupStream {
	tmp := NullableMultipleGroupStreamFrom(arg)
	return &tmp
}

func (self *NullableMultipleGroupStream) Add(arg NullableMultipleGroup) *NullableMultipleGroupStream {
	return self.AddAll(arg)
}
func (self *NullableMultipleGroupStream) AddAll(arg ...NullableMultipleGroup) *NullableMultipleGroupStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableMultipleGroupStream) AddSafe(arg *NullableMultipleGroup) *NullableMultipleGroupStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableMultipleGroupStream) Aggregate(fn func(NullableMultipleGroup, NullableMultipleGroup) NullableMultipleGroup) *NullableMultipleGroupStream {
	result := NullableMultipleGroupStreamOf()
	self.ForEach(func(v NullableMultipleGroup, i int) {
		if i == 0 {
			result.Add(fn(NullableMultipleGroup{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleGroupStream) AllMatch(fn func(NullableMultipleGroup, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableMultipleGroupStream) AnyMatch(fn func(NullableMultipleGroup, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableMultipleGroupStream) Clone() *NullableMultipleGroupStream {
	temp := make([]NullableMultipleGroup, self.Len())
	copy(temp, *self)
	return (*NullableMultipleGroupStream)(&temp)
}
func (self *NullableMultipleGroupStream) Copy() *NullableMultipleGroupStream {
	return self.Clone()
}
func (self *NullableMultipleGroupStream) Concat(arg []NullableMultipleGroup) *NullableMultipleGroupStream {
	return self.AddAll(arg...)
}
func (self *NullableMultipleGroupStream) Contains(arg NullableMultipleGroup) bool {
	return self.FindIndex(func(_arg NullableMultipleGroup, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableMultipleGroupStream) Clean() *NullableMultipleGroupStream {
	*self = NullableMultipleGroupStreamOf()
	return self
}
func (self *NullableMultipleGroupStream) Delete(index int) *NullableMultipleGroupStream {
	return self.DeleteRange(index, index)
}
func (self *NullableMultipleGroupStream) DeleteRange(startIndex, endIndex int) *NullableMultipleGroupStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableMultipleGroupStream) Distinct() *NullableMultipleGroupStream {
	caches := map[string]bool{}
	result := NullableMultipleGroupStreamOf()
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
func (self *NullableMultipleGroupStream) Each(fn func(NullableMultipleGroup)) *NullableMultipleGroupStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableMultipleGroupStream) EachRight(fn func(NullableMultipleGroup)) *NullableMultipleGroupStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableMultipleGroupStream) Equals(arr []NullableMultipleGroup) bool {
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
func (self *NullableMultipleGroupStream) Filter(fn func(NullableMultipleGroup, int) bool) *NullableMultipleGroupStream {
	result := NullableMultipleGroupStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleGroupStream) FilterSlim(fn func(NullableMultipleGroup, int) bool) *NullableMultipleGroupStream {
	result := NullableMultipleGroupStreamOf()
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
func (self *NullableMultipleGroupStream) Find(fn func(NullableMultipleGroup, int) bool) *NullableMultipleGroup {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleGroupStream) FindOr(fn func(NullableMultipleGroup, int) bool, or NullableMultipleGroup) NullableMultipleGroup {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableMultipleGroupStream) FindIndex(fn func(NullableMultipleGroup, int) bool) int {
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
func (self *NullableMultipleGroupStream) First() *NullableMultipleGroup {
	return self.Get(0)
}
func (self *NullableMultipleGroupStream) FirstOr(arg NullableMultipleGroup) NullableMultipleGroup {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleGroupStream) ForEach(fn func(NullableMultipleGroup, int)) *NullableMultipleGroupStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableMultipleGroupStream) ForEachRight(fn func(NullableMultipleGroup, int)) *NullableMultipleGroupStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableMultipleGroupStream) GroupBy(fn func(NullableMultipleGroup, int) string) map[string][]NullableMultipleGroup {
	m := map[string][]NullableMultipleGroup{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableMultipleGroupStream) GroupByValues(fn func(NullableMultipleGroup, int) string) [][]NullableMultipleGroup {
	var tmp [][]NullableMultipleGroup
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableMultipleGroupStream) IndexOf(arg NullableMultipleGroup) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableMultipleGroupStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableMultipleGroupStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableMultipleGroupStream) Last() *NullableMultipleGroup {
	return self.Get(self.Len() - 1)
}
func (self *NullableMultipleGroupStream) LastOr(arg NullableMultipleGroup) NullableMultipleGroup {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleGroupStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableMultipleGroupStream) Limit(limit int) *NullableMultipleGroupStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableMultipleGroupStream) Map(fn func(NullableMultipleGroup, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleGroupStream) Map2Int(fn func(NullableMultipleGroup, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleGroupStream) Map2Int32(fn func(NullableMultipleGroup, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleGroupStream) Map2Int64(fn func(NullableMultipleGroup, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleGroupStream) Map2Float32(fn func(NullableMultipleGroup, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleGroupStream) Map2Float64(fn func(NullableMultipleGroup, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleGroupStream) Map2Bool(fn func(NullableMultipleGroup, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleGroupStream) Map2Bytes(fn func(NullableMultipleGroup, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleGroupStream) Map2String(fn func(NullableMultipleGroup, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleGroupStream) Max(fn func(NullableMultipleGroup, int) float64) *NullableMultipleGroup {
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
func (self *NullableMultipleGroupStream) Min(fn func(NullableMultipleGroup, int) float64) *NullableMultipleGroup {
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
func (self *NullableMultipleGroupStream) NoneMatch(fn func(NullableMultipleGroup, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableMultipleGroupStream) Get(index int) *NullableMultipleGroup {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleGroupStream) GetOr(index int, arg NullableMultipleGroup) NullableMultipleGroup {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleGroupStream) Peek(fn func(*NullableMultipleGroup, int)) *NullableMultipleGroupStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableMultipleGroupStream) Reduce(fn func(NullableMultipleGroup, NullableMultipleGroup, int) NullableMultipleGroup) *NullableMultipleGroupStream {
	return self.ReduceInit(fn, NullableMultipleGroup{})
}
func (self *NullableMultipleGroupStream) ReduceInit(fn func(NullableMultipleGroup, NullableMultipleGroup, int) NullableMultipleGroup, initialValue NullableMultipleGroup) *NullableMultipleGroupStream {
	result := NullableMultipleGroupStreamOf()
	self.ForEach(func(v NullableMultipleGroup, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleGroupStream) ReduceInterface(fn func(interface{}, NullableMultipleGroup, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableMultipleGroup{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleGroupStream) ReduceString(fn func(string, NullableMultipleGroup, int) string) []string {
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
func (self *NullableMultipleGroupStream) ReduceInt(fn func(int, NullableMultipleGroup, int) int) []int {
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
func (self *NullableMultipleGroupStream) ReduceInt32(fn func(int32, NullableMultipleGroup, int) int32) []int32 {
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
func (self *NullableMultipleGroupStream) ReduceInt64(fn func(int64, NullableMultipleGroup, int) int64) []int64 {
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
func (self *NullableMultipleGroupStream) ReduceFloat32(fn func(float32, NullableMultipleGroup, int) float32) []float32 {
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
func (self *NullableMultipleGroupStream) ReduceFloat64(fn func(float64, NullableMultipleGroup, int) float64) []float64 {
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
func (self *NullableMultipleGroupStream) ReduceBool(fn func(bool, NullableMultipleGroup, int) bool) []bool {
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
func (self *NullableMultipleGroupStream) Reverse() *NullableMultipleGroupStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableMultipleGroupStream) Replace(fn func(NullableMultipleGroup, int) NullableMultipleGroup) *NullableMultipleGroupStream {
	return self.ForEach(func(v NullableMultipleGroup, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableMultipleGroupStream) Select(fn func(NullableMultipleGroup) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableMultipleGroupStream) Set(index int, val NullableMultipleGroup) *NullableMultipleGroupStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableMultipleGroupStream) Skip(skip int) *NullableMultipleGroupStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableMultipleGroupStream) SkippingEach(fn func(NullableMultipleGroup, int) int) *NullableMultipleGroupStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableMultipleGroupStream) Slice(startIndex, n int) *NullableMultipleGroupStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableMultipleGroup{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableMultipleGroupStream) Sort(fn func(i, j int) bool) *NullableMultipleGroupStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableMultipleGroupStream) Tail() *NullableMultipleGroup {
	return self.Last()
}
func (self *NullableMultipleGroupStream) TailOr(arg NullableMultipleGroup) NullableMultipleGroup {
	return self.LastOr(arg)
}
func (self *NullableMultipleGroupStream) ToList() []NullableMultipleGroup {
	return self.Val()
}
func (self *NullableMultipleGroupStream) Unique() *NullableMultipleGroupStream {
	return self.Distinct()
}
func (self *NullableMultipleGroupStream) Val() []NullableMultipleGroup {
	if self == nil {
		return []NullableMultipleGroup{}
	}
	return *self.Copy()
}
func (self *NullableMultipleGroupStream) While(fn func(NullableMultipleGroup, int) bool) *NullableMultipleGroupStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableMultipleGroupStream) Where(fn func(NullableMultipleGroup) bool) *NullableMultipleGroupStream {
	result := NullableMultipleGroupStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleGroupStream) WhereSlim(fn func(NullableMultipleGroup) bool) *NullableMultipleGroupStream {
	result := NullableMultipleGroupStreamOf()
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
