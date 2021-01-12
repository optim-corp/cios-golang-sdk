/*
 * Collection utility of NullableResourceOwnerProfile Struct
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

type NullableResourceOwnerProfileStream []NullableResourceOwnerProfile

func NullableResourceOwnerProfileStreamOf(arg ...NullableResourceOwnerProfile) NullableResourceOwnerProfileStream {
	return arg
}
func NullableResourceOwnerProfileStreamFrom(arg []NullableResourceOwnerProfile) NullableResourceOwnerProfileStream {
	return arg
}
func CreateNullableResourceOwnerProfileStream(arg ...NullableResourceOwnerProfile) *NullableResourceOwnerProfileStream {
	tmp := NullableResourceOwnerProfileStreamOf(arg...)
	return &tmp
}
func GenerateNullableResourceOwnerProfileStream(arg []NullableResourceOwnerProfile) *NullableResourceOwnerProfileStream {
	tmp := NullableResourceOwnerProfileStreamFrom(arg)
	return &tmp
}

func (self *NullableResourceOwnerProfileStream) Add(arg NullableResourceOwnerProfile) *NullableResourceOwnerProfileStream {
	return self.AddAll(arg)
}
func (self *NullableResourceOwnerProfileStream) AddAll(arg ...NullableResourceOwnerProfile) *NullableResourceOwnerProfileStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableResourceOwnerProfileStream) AddSafe(arg *NullableResourceOwnerProfile) *NullableResourceOwnerProfileStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableResourceOwnerProfileStream) Aggregate(fn func(NullableResourceOwnerProfile, NullableResourceOwnerProfile) NullableResourceOwnerProfile) *NullableResourceOwnerProfileStream {
	result := NullableResourceOwnerProfileStreamOf()
	self.ForEach(func(v NullableResourceOwnerProfile, i int) {
		if i == 0 {
			result.Add(fn(NullableResourceOwnerProfile{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableResourceOwnerProfileStream) AllMatch(fn func(NullableResourceOwnerProfile, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableResourceOwnerProfileStream) AnyMatch(fn func(NullableResourceOwnerProfile, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableResourceOwnerProfileStream) Clone() *NullableResourceOwnerProfileStream {
	temp := make([]NullableResourceOwnerProfile, self.Len())
	copy(temp, *self)
	return (*NullableResourceOwnerProfileStream)(&temp)
}
func (self *NullableResourceOwnerProfileStream) Copy() *NullableResourceOwnerProfileStream {
	return self.Clone()
}
func (self *NullableResourceOwnerProfileStream) Concat(arg []NullableResourceOwnerProfile) *NullableResourceOwnerProfileStream {
	return self.AddAll(arg...)
}
func (self *NullableResourceOwnerProfileStream) Contains(arg NullableResourceOwnerProfile) bool {
	return self.FindIndex(func(_arg NullableResourceOwnerProfile, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableResourceOwnerProfileStream) Clean() *NullableResourceOwnerProfileStream {
	*self = NullableResourceOwnerProfileStreamOf()
	return self
}
func (self *NullableResourceOwnerProfileStream) Delete(index int) *NullableResourceOwnerProfileStream {
	return self.DeleteRange(index, index)
}
func (self *NullableResourceOwnerProfileStream) DeleteRange(startIndex, endIndex int) *NullableResourceOwnerProfileStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableResourceOwnerProfileStream) Distinct() *NullableResourceOwnerProfileStream {
	caches := map[string]bool{}
	result := NullableResourceOwnerProfileStreamOf()
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
func (self *NullableResourceOwnerProfileStream) Each(fn func(NullableResourceOwnerProfile)) *NullableResourceOwnerProfileStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableResourceOwnerProfileStream) EachRight(fn func(NullableResourceOwnerProfile)) *NullableResourceOwnerProfileStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableResourceOwnerProfileStream) Equals(arr []NullableResourceOwnerProfile) bool {
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
func (self *NullableResourceOwnerProfileStream) Filter(fn func(NullableResourceOwnerProfile, int) bool) *NullableResourceOwnerProfileStream {
	result := NullableResourceOwnerProfileStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableResourceOwnerProfileStream) FilterSlim(fn func(NullableResourceOwnerProfile, int) bool) *NullableResourceOwnerProfileStream {
	result := NullableResourceOwnerProfileStreamOf()
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
func (self *NullableResourceOwnerProfileStream) Find(fn func(NullableResourceOwnerProfile, int) bool) *NullableResourceOwnerProfile {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableResourceOwnerProfileStream) FindOr(fn func(NullableResourceOwnerProfile, int) bool, or NullableResourceOwnerProfile) NullableResourceOwnerProfile {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableResourceOwnerProfileStream) FindIndex(fn func(NullableResourceOwnerProfile, int) bool) int {
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
func (self *NullableResourceOwnerProfileStream) First() *NullableResourceOwnerProfile {
	return self.Get(0)
}
func (self *NullableResourceOwnerProfileStream) FirstOr(arg NullableResourceOwnerProfile) NullableResourceOwnerProfile {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableResourceOwnerProfileStream) ForEach(fn func(NullableResourceOwnerProfile, int)) *NullableResourceOwnerProfileStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableResourceOwnerProfileStream) ForEachRight(fn func(NullableResourceOwnerProfile, int)) *NullableResourceOwnerProfileStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableResourceOwnerProfileStream) GroupBy(fn func(NullableResourceOwnerProfile, int) string) map[string][]NullableResourceOwnerProfile {
	m := map[string][]NullableResourceOwnerProfile{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableResourceOwnerProfileStream) GroupByValues(fn func(NullableResourceOwnerProfile, int) string) [][]NullableResourceOwnerProfile {
	var tmp [][]NullableResourceOwnerProfile
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableResourceOwnerProfileStream) IndexOf(arg NullableResourceOwnerProfile) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableResourceOwnerProfileStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableResourceOwnerProfileStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableResourceOwnerProfileStream) Last() *NullableResourceOwnerProfile {
	return self.Get(self.Len() - 1)
}
func (self *NullableResourceOwnerProfileStream) LastOr(arg NullableResourceOwnerProfile) NullableResourceOwnerProfile {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableResourceOwnerProfileStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableResourceOwnerProfileStream) Limit(limit int) *NullableResourceOwnerProfileStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableResourceOwnerProfileStream) Map(fn func(NullableResourceOwnerProfile, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableResourceOwnerProfileStream) Map2Int(fn func(NullableResourceOwnerProfile, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableResourceOwnerProfileStream) Map2Int32(fn func(NullableResourceOwnerProfile, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableResourceOwnerProfileStream) Map2Int64(fn func(NullableResourceOwnerProfile, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableResourceOwnerProfileStream) Map2Float32(fn func(NullableResourceOwnerProfile, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableResourceOwnerProfileStream) Map2Float64(fn func(NullableResourceOwnerProfile, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableResourceOwnerProfileStream) Map2Bool(fn func(NullableResourceOwnerProfile, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableResourceOwnerProfileStream) Map2Bytes(fn func(NullableResourceOwnerProfile, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableResourceOwnerProfileStream) Map2String(fn func(NullableResourceOwnerProfile, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableResourceOwnerProfileStream) Max(fn func(NullableResourceOwnerProfile, int) float64) *NullableResourceOwnerProfile {
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
func (self *NullableResourceOwnerProfileStream) Min(fn func(NullableResourceOwnerProfile, int) float64) *NullableResourceOwnerProfile {
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
func (self *NullableResourceOwnerProfileStream) NoneMatch(fn func(NullableResourceOwnerProfile, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableResourceOwnerProfileStream) Get(index int) *NullableResourceOwnerProfile {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableResourceOwnerProfileStream) GetOr(index int, arg NullableResourceOwnerProfile) NullableResourceOwnerProfile {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableResourceOwnerProfileStream) Peek(fn func(*NullableResourceOwnerProfile, int)) *NullableResourceOwnerProfileStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableResourceOwnerProfileStream) Reduce(fn func(NullableResourceOwnerProfile, NullableResourceOwnerProfile, int) NullableResourceOwnerProfile) *NullableResourceOwnerProfileStream {
	return self.ReduceInit(fn, NullableResourceOwnerProfile{})
}
func (self *NullableResourceOwnerProfileStream) ReduceInit(fn func(NullableResourceOwnerProfile, NullableResourceOwnerProfile, int) NullableResourceOwnerProfile, initialValue NullableResourceOwnerProfile) *NullableResourceOwnerProfileStream {
	result := NullableResourceOwnerProfileStreamOf()
	self.ForEach(func(v NullableResourceOwnerProfile, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableResourceOwnerProfileStream) ReduceInterface(fn func(interface{}, NullableResourceOwnerProfile, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableResourceOwnerProfile{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableResourceOwnerProfileStream) ReduceString(fn func(string, NullableResourceOwnerProfile, int) string) []string {
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
func (self *NullableResourceOwnerProfileStream) ReduceInt(fn func(int, NullableResourceOwnerProfile, int) int) []int {
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
func (self *NullableResourceOwnerProfileStream) ReduceInt32(fn func(int32, NullableResourceOwnerProfile, int) int32) []int32 {
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
func (self *NullableResourceOwnerProfileStream) ReduceInt64(fn func(int64, NullableResourceOwnerProfile, int) int64) []int64 {
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
func (self *NullableResourceOwnerProfileStream) ReduceFloat32(fn func(float32, NullableResourceOwnerProfile, int) float32) []float32 {
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
func (self *NullableResourceOwnerProfileStream) ReduceFloat64(fn func(float64, NullableResourceOwnerProfile, int) float64) []float64 {
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
func (self *NullableResourceOwnerProfileStream) ReduceBool(fn func(bool, NullableResourceOwnerProfile, int) bool) []bool {
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
func (self *NullableResourceOwnerProfileStream) Reverse() *NullableResourceOwnerProfileStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableResourceOwnerProfileStream) Replace(fn func(NullableResourceOwnerProfile, int) NullableResourceOwnerProfile) *NullableResourceOwnerProfileStream {
	return self.ForEach(func(v NullableResourceOwnerProfile, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableResourceOwnerProfileStream) Select(fn func(NullableResourceOwnerProfile) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableResourceOwnerProfileStream) Set(index int, val NullableResourceOwnerProfile) *NullableResourceOwnerProfileStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableResourceOwnerProfileStream) Skip(skip int) *NullableResourceOwnerProfileStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableResourceOwnerProfileStream) SkippingEach(fn func(NullableResourceOwnerProfile, int) int) *NullableResourceOwnerProfileStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableResourceOwnerProfileStream) Slice(startIndex, n int) *NullableResourceOwnerProfileStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableResourceOwnerProfile{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableResourceOwnerProfileStream) Sort(fn func(i, j int) bool) *NullableResourceOwnerProfileStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableResourceOwnerProfileStream) Tail() *NullableResourceOwnerProfile {
	return self.Last()
}
func (self *NullableResourceOwnerProfileStream) TailOr(arg NullableResourceOwnerProfile) NullableResourceOwnerProfile {
	return self.LastOr(arg)
}
func (self *NullableResourceOwnerProfileStream) ToList() []NullableResourceOwnerProfile {
	return self.Val()
}
func (self *NullableResourceOwnerProfileStream) Unique() *NullableResourceOwnerProfileStream {
	return self.Distinct()
}
func (self *NullableResourceOwnerProfileStream) Val() []NullableResourceOwnerProfile {
	if self == nil {
		return []NullableResourceOwnerProfile{}
	}
	return *self.Copy()
}
func (self *NullableResourceOwnerProfileStream) While(fn func(NullableResourceOwnerProfile, int) bool) *NullableResourceOwnerProfileStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableResourceOwnerProfileStream) Where(fn func(NullableResourceOwnerProfile) bool) *NullableResourceOwnerProfileStream {
	result := NullableResourceOwnerProfileStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableResourceOwnerProfileStream) WhereSlim(fn func(NullableResourceOwnerProfile) bool) *NullableResourceOwnerProfileStream {
	result := NullableResourceOwnerProfileStreamOf()
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