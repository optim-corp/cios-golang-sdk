/*
 * Collection utility of NullableMultipleMember Struct
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

type NullableMultipleMemberStream []NullableMultipleMember

func NullableMultipleMemberStreamOf(arg ...NullableMultipleMember) NullableMultipleMemberStream {
	return arg
}
func NullableMultipleMemberStreamFrom(arg []NullableMultipleMember) NullableMultipleMemberStream {
	return arg
}
func CreateNullableMultipleMemberStream(arg ...NullableMultipleMember) *NullableMultipleMemberStream {
	tmp := NullableMultipleMemberStreamOf(arg...)
	return &tmp
}
func GenerateNullableMultipleMemberStream(arg []NullableMultipleMember) *NullableMultipleMemberStream {
	tmp := NullableMultipleMemberStreamFrom(arg)
	return &tmp
}

func (self *NullableMultipleMemberStream) Add(arg NullableMultipleMember) *NullableMultipleMemberStream {
	return self.AddAll(arg)
}
func (self *NullableMultipleMemberStream) AddAll(arg ...NullableMultipleMember) *NullableMultipleMemberStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableMultipleMemberStream) AddSafe(arg *NullableMultipleMember) *NullableMultipleMemberStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableMultipleMemberStream) Aggregate(fn func(NullableMultipleMember, NullableMultipleMember) NullableMultipleMember) *NullableMultipleMemberStream {
	result := NullableMultipleMemberStreamOf()
	self.ForEach(func(v NullableMultipleMember, i int) {
		if i == 0 {
			result.Add(fn(NullableMultipleMember{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleMemberStream) AllMatch(fn func(NullableMultipleMember, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableMultipleMemberStream) AnyMatch(fn func(NullableMultipleMember, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableMultipleMemberStream) Clone() *NullableMultipleMemberStream {
	temp := make([]NullableMultipleMember, self.Len())
	copy(temp, *self)
	return (*NullableMultipleMemberStream)(&temp)
}
func (self *NullableMultipleMemberStream) Copy() *NullableMultipleMemberStream {
	return self.Clone()
}
func (self *NullableMultipleMemberStream) Concat(arg []NullableMultipleMember) *NullableMultipleMemberStream {
	return self.AddAll(arg...)
}
func (self *NullableMultipleMemberStream) Contains(arg NullableMultipleMember) bool {
	return self.FindIndex(func(_arg NullableMultipleMember, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableMultipleMemberStream) Clean() *NullableMultipleMemberStream {
	*self = NullableMultipleMemberStreamOf()
	return self
}
func (self *NullableMultipleMemberStream) Delete(index int) *NullableMultipleMemberStream {
	return self.DeleteRange(index, index)
}
func (self *NullableMultipleMemberStream) DeleteRange(startIndex, endIndex int) *NullableMultipleMemberStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableMultipleMemberStream) Distinct() *NullableMultipleMemberStream {
	caches := map[string]bool{}
	result := NullableMultipleMemberStreamOf()
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
func (self *NullableMultipleMemberStream) Each(fn func(NullableMultipleMember)) *NullableMultipleMemberStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableMultipleMemberStream) EachRight(fn func(NullableMultipleMember)) *NullableMultipleMemberStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableMultipleMemberStream) Equals(arr []NullableMultipleMember) bool {
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
func (self *NullableMultipleMemberStream) Filter(fn func(NullableMultipleMember, int) bool) *NullableMultipleMemberStream {
	result := NullableMultipleMemberStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleMemberStream) FilterSlim(fn func(NullableMultipleMember, int) bool) *NullableMultipleMemberStream {
	result := NullableMultipleMemberStreamOf()
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
func (self *NullableMultipleMemberStream) Find(fn func(NullableMultipleMember, int) bool) *NullableMultipleMember {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleMemberStream) FindOr(fn func(NullableMultipleMember, int) bool, or NullableMultipleMember) NullableMultipleMember {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableMultipleMemberStream) FindIndex(fn func(NullableMultipleMember, int) bool) int {
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
func (self *NullableMultipleMemberStream) First() *NullableMultipleMember {
	return self.Get(0)
}
func (self *NullableMultipleMemberStream) FirstOr(arg NullableMultipleMember) NullableMultipleMember {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleMemberStream) ForEach(fn func(NullableMultipleMember, int)) *NullableMultipleMemberStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableMultipleMemberStream) ForEachRight(fn func(NullableMultipleMember, int)) *NullableMultipleMemberStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableMultipleMemberStream) GroupBy(fn func(NullableMultipleMember, int) string) map[string][]NullableMultipleMember {
	m := map[string][]NullableMultipleMember{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableMultipleMemberStream) GroupByValues(fn func(NullableMultipleMember, int) string) [][]NullableMultipleMember {
	var tmp [][]NullableMultipleMember
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableMultipleMemberStream) IndexOf(arg NullableMultipleMember) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableMultipleMemberStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableMultipleMemberStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableMultipleMemberStream) Last() *NullableMultipleMember {
	return self.Get(self.Len() - 1)
}
func (self *NullableMultipleMemberStream) LastOr(arg NullableMultipleMember) NullableMultipleMember {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleMemberStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableMultipleMemberStream) Limit(limit int) *NullableMultipleMemberStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableMultipleMemberStream) Map(fn func(NullableMultipleMember, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleMemberStream) Map2Int(fn func(NullableMultipleMember, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleMemberStream) Map2Int32(fn func(NullableMultipleMember, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleMemberStream) Map2Int64(fn func(NullableMultipleMember, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleMemberStream) Map2Float32(fn func(NullableMultipleMember, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleMemberStream) Map2Float64(fn func(NullableMultipleMember, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleMemberStream) Map2Bool(fn func(NullableMultipleMember, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleMemberStream) Map2Bytes(fn func(NullableMultipleMember, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleMemberStream) Map2String(fn func(NullableMultipleMember, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleMemberStream) Max(fn func(NullableMultipleMember, int) float64) *NullableMultipleMember {
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
func (self *NullableMultipleMemberStream) Min(fn func(NullableMultipleMember, int) float64) *NullableMultipleMember {
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
func (self *NullableMultipleMemberStream) NoneMatch(fn func(NullableMultipleMember, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableMultipleMemberStream) Get(index int) *NullableMultipleMember {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleMemberStream) GetOr(index int, arg NullableMultipleMember) NullableMultipleMember {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleMemberStream) Peek(fn func(*NullableMultipleMember, int)) *NullableMultipleMemberStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableMultipleMemberStream) Reduce(fn func(NullableMultipleMember, NullableMultipleMember, int) NullableMultipleMember) *NullableMultipleMemberStream {
	return self.ReduceInit(fn, NullableMultipleMember{})
}
func (self *NullableMultipleMemberStream) ReduceInit(fn func(NullableMultipleMember, NullableMultipleMember, int) NullableMultipleMember, initialValue NullableMultipleMember) *NullableMultipleMemberStream {
	result := NullableMultipleMemberStreamOf()
	self.ForEach(func(v NullableMultipleMember, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleMemberStream) ReduceInterface(fn func(interface{}, NullableMultipleMember, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableMultipleMember{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleMemberStream) ReduceString(fn func(string, NullableMultipleMember, int) string) []string {
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
func (self *NullableMultipleMemberStream) ReduceInt(fn func(int, NullableMultipleMember, int) int) []int {
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
func (self *NullableMultipleMemberStream) ReduceInt32(fn func(int32, NullableMultipleMember, int) int32) []int32 {
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
func (self *NullableMultipleMemberStream) ReduceInt64(fn func(int64, NullableMultipleMember, int) int64) []int64 {
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
func (self *NullableMultipleMemberStream) ReduceFloat32(fn func(float32, NullableMultipleMember, int) float32) []float32 {
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
func (self *NullableMultipleMemberStream) ReduceFloat64(fn func(float64, NullableMultipleMember, int) float64) []float64 {
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
func (self *NullableMultipleMemberStream) ReduceBool(fn func(bool, NullableMultipleMember, int) bool) []bool {
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
func (self *NullableMultipleMemberStream) Reverse() *NullableMultipleMemberStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableMultipleMemberStream) Replace(fn func(NullableMultipleMember, int) NullableMultipleMember) *NullableMultipleMemberStream {
	return self.ForEach(func(v NullableMultipleMember, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableMultipleMemberStream) Select(fn func(NullableMultipleMember) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableMultipleMemberStream) Set(index int, val NullableMultipleMember) *NullableMultipleMemberStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableMultipleMemberStream) Skip(skip int) *NullableMultipleMemberStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableMultipleMemberStream) SkippingEach(fn func(NullableMultipleMember, int) int) *NullableMultipleMemberStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableMultipleMemberStream) Slice(startIndex, n int) *NullableMultipleMemberStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableMultipleMember{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableMultipleMemberStream) Sort(fn func(i, j int) bool) *NullableMultipleMemberStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableMultipleMemberStream) Tail() *NullableMultipleMember {
	return self.Last()
}
func (self *NullableMultipleMemberStream) TailOr(arg NullableMultipleMember) NullableMultipleMember {
	return self.LastOr(arg)
}
func (self *NullableMultipleMemberStream) ToList() []NullableMultipleMember {
	return self.Val()
}
func (self *NullableMultipleMemberStream) Unique() *NullableMultipleMemberStream {
	return self.Distinct()
}
func (self *NullableMultipleMemberStream) Val() []NullableMultipleMember {
	if self == nil {
		return []NullableMultipleMember{}
	}
	return *self.Copy()
}
func (self *NullableMultipleMemberStream) While(fn func(NullableMultipleMember, int) bool) *NullableMultipleMemberStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableMultipleMemberStream) Where(fn func(NullableMultipleMember) bool) *NullableMultipleMemberStream {
	result := NullableMultipleMemberStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleMemberStream) WhereSlim(fn func(NullableMultipleMember) bool) *NullableMultipleMemberStream {
	result := NullableMultipleMemberStreamOf()
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
