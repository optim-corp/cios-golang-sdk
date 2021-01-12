/*
 * Collection utility of NullableMember Struct
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

type NullableMemberStream []NullableMember

func NullableMemberStreamOf(arg ...NullableMember) NullableMemberStream {
	return arg
}
func NullableMemberStreamFrom(arg []NullableMember) NullableMemberStream {
	return arg
}
func CreateNullableMemberStream(arg ...NullableMember) *NullableMemberStream {
	tmp := NullableMemberStreamOf(arg...)
	return &tmp
}
func GenerateNullableMemberStream(arg []NullableMember) *NullableMemberStream {
	tmp := NullableMemberStreamFrom(arg)
	return &tmp
}

func (self *NullableMemberStream) Add(arg NullableMember) *NullableMemberStream {
	return self.AddAll(arg)
}
func (self *NullableMemberStream) AddAll(arg ...NullableMember) *NullableMemberStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableMemberStream) AddSafe(arg *NullableMember) *NullableMemberStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableMemberStream) Aggregate(fn func(NullableMember, NullableMember) NullableMember) *NullableMemberStream {
	result := NullableMemberStreamOf()
	self.ForEach(func(v NullableMember, i int) {
		if i == 0 {
			result.Add(fn(NullableMember{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableMemberStream) AllMatch(fn func(NullableMember, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableMemberStream) AnyMatch(fn func(NullableMember, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableMemberStream) Clone() *NullableMemberStream {
	temp := make([]NullableMember, self.Len())
	copy(temp, *self)
	return (*NullableMemberStream)(&temp)
}
func (self *NullableMemberStream) Copy() *NullableMemberStream {
	return self.Clone()
}
func (self *NullableMemberStream) Concat(arg []NullableMember) *NullableMemberStream {
	return self.AddAll(arg...)
}
func (self *NullableMemberStream) Contains(arg NullableMember) bool {
	return self.FindIndex(func(_arg NullableMember, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableMemberStream) Clean() *NullableMemberStream {
	*self = NullableMemberStreamOf()
	return self
}
func (self *NullableMemberStream) Delete(index int) *NullableMemberStream {
	return self.DeleteRange(index, index)
}
func (self *NullableMemberStream) DeleteRange(startIndex, endIndex int) *NullableMemberStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableMemberStream) Distinct() *NullableMemberStream {
	caches := map[string]bool{}
	result := NullableMemberStreamOf()
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
func (self *NullableMemberStream) Each(fn func(NullableMember)) *NullableMemberStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableMemberStream) EachRight(fn func(NullableMember)) *NullableMemberStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableMemberStream) Equals(arr []NullableMember) bool {
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
func (self *NullableMemberStream) Filter(fn func(NullableMember, int) bool) *NullableMemberStream {
	result := NullableMemberStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMemberStream) FilterSlim(fn func(NullableMember, int) bool) *NullableMemberStream {
	result := NullableMemberStreamOf()
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
func (self *NullableMemberStream) Find(fn func(NullableMember, int) bool) *NullableMember {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableMemberStream) FindOr(fn func(NullableMember, int) bool, or NullableMember) NullableMember {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableMemberStream) FindIndex(fn func(NullableMember, int) bool) int {
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
func (self *NullableMemberStream) First() *NullableMember {
	return self.Get(0)
}
func (self *NullableMemberStream) FirstOr(arg NullableMember) NullableMember {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMemberStream) ForEach(fn func(NullableMember, int)) *NullableMemberStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableMemberStream) ForEachRight(fn func(NullableMember, int)) *NullableMemberStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableMemberStream) GroupBy(fn func(NullableMember, int) string) map[string][]NullableMember {
	m := map[string][]NullableMember{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableMemberStream) GroupByValues(fn func(NullableMember, int) string) [][]NullableMember {
	var tmp [][]NullableMember
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableMemberStream) IndexOf(arg NullableMember) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableMemberStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableMemberStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableMemberStream) Last() *NullableMember {
	return self.Get(self.Len() - 1)
}
func (self *NullableMemberStream) LastOr(arg NullableMember) NullableMember {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMemberStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableMemberStream) Limit(limit int) *NullableMemberStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableMemberStream) Map(fn func(NullableMember, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMemberStream) Map2Int(fn func(NullableMember, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMemberStream) Map2Int32(fn func(NullableMember, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMemberStream) Map2Int64(fn func(NullableMember, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMemberStream) Map2Float32(fn func(NullableMember, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMemberStream) Map2Float64(fn func(NullableMember, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMemberStream) Map2Bool(fn func(NullableMember, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMemberStream) Map2Bytes(fn func(NullableMember, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMemberStream) Map2String(fn func(NullableMember, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMemberStream) Max(fn func(NullableMember, int) float64) *NullableMember {
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
func (self *NullableMemberStream) Min(fn func(NullableMember, int) float64) *NullableMember {
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
func (self *NullableMemberStream) NoneMatch(fn func(NullableMember, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableMemberStream) Get(index int) *NullableMember {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableMemberStream) GetOr(index int, arg NullableMember) NullableMember {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMemberStream) Peek(fn func(*NullableMember, int)) *NullableMemberStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableMemberStream) Reduce(fn func(NullableMember, NullableMember, int) NullableMember) *NullableMemberStream {
	return self.ReduceInit(fn, NullableMember{})
}
func (self *NullableMemberStream) ReduceInit(fn func(NullableMember, NullableMember, int) NullableMember, initialValue NullableMember) *NullableMemberStream {
	result := NullableMemberStreamOf()
	self.ForEach(func(v NullableMember, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableMemberStream) ReduceInterface(fn func(interface{}, NullableMember, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableMember{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMemberStream) ReduceString(fn func(string, NullableMember, int) string) []string {
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
func (self *NullableMemberStream) ReduceInt(fn func(int, NullableMember, int) int) []int {
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
func (self *NullableMemberStream) ReduceInt32(fn func(int32, NullableMember, int) int32) []int32 {
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
func (self *NullableMemberStream) ReduceInt64(fn func(int64, NullableMember, int) int64) []int64 {
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
func (self *NullableMemberStream) ReduceFloat32(fn func(float32, NullableMember, int) float32) []float32 {
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
func (self *NullableMemberStream) ReduceFloat64(fn func(float64, NullableMember, int) float64) []float64 {
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
func (self *NullableMemberStream) ReduceBool(fn func(bool, NullableMember, int) bool) []bool {
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
func (self *NullableMemberStream) Reverse() *NullableMemberStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableMemberStream) Replace(fn func(NullableMember, int) NullableMember) *NullableMemberStream {
	return self.ForEach(func(v NullableMember, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableMemberStream) Select(fn func(NullableMember) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableMemberStream) Set(index int, val NullableMember) *NullableMemberStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableMemberStream) Skip(skip int) *NullableMemberStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableMemberStream) SkippingEach(fn func(NullableMember, int) int) *NullableMemberStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableMemberStream) Slice(startIndex, n int) *NullableMemberStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableMember{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableMemberStream) Sort(fn func(i, j int) bool) *NullableMemberStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableMemberStream) Tail() *NullableMember {
	return self.Last()
}
func (self *NullableMemberStream) TailOr(arg NullableMember) NullableMember {
	return self.LastOr(arg)
}
func (self *NullableMemberStream) ToList() []NullableMember {
	return self.Val()
}
func (self *NullableMemberStream) Unique() *NullableMemberStream {
	return self.Distinct()
}
func (self *NullableMemberStream) Val() []NullableMember {
	if self == nil {
		return []NullableMember{}
	}
	return *self.Copy()
}
func (self *NullableMemberStream) While(fn func(NullableMember, int) bool) *NullableMemberStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableMemberStream) Where(fn func(NullableMember) bool) *NullableMemberStream {
	result := NullableMemberStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMemberStream) WhereSlim(fn func(NullableMember) bool) *NullableMemberStream {
	result := NullableMemberStreamOf()
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