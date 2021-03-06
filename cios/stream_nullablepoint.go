/*
 * Collection utility of NullablePoint Struct
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

type NullablePointStream []NullablePoint

func NullablePointStreamOf(arg ...NullablePoint) NullablePointStream {
	return arg
}
func NullablePointStreamFrom(arg []NullablePoint) NullablePointStream {
	return arg
}
func CreateNullablePointStream(arg ...NullablePoint) *NullablePointStream {
	tmp := NullablePointStreamOf(arg...)
	return &tmp
}
func GenerateNullablePointStream(arg []NullablePoint) *NullablePointStream {
	tmp := NullablePointStreamFrom(arg)
	return &tmp
}

func (self *NullablePointStream) Add(arg NullablePoint) *NullablePointStream {
	return self.AddAll(arg)
}
func (self *NullablePointStream) AddAll(arg ...NullablePoint) *NullablePointStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullablePointStream) AddSafe(arg *NullablePoint) *NullablePointStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullablePointStream) Aggregate(fn func(NullablePoint, NullablePoint) NullablePoint) *NullablePointStream {
	result := NullablePointStreamOf()
	self.ForEach(func(v NullablePoint, i int) {
		if i == 0 {
			result.Add(fn(NullablePoint{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullablePointStream) AllMatch(fn func(NullablePoint, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullablePointStream) AnyMatch(fn func(NullablePoint, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullablePointStream) Clone() *NullablePointStream {
	temp := make([]NullablePoint, self.Len())
	copy(temp, *self)
	return (*NullablePointStream)(&temp)
}
func (self *NullablePointStream) Copy() *NullablePointStream {
	return self.Clone()
}
func (self *NullablePointStream) Concat(arg []NullablePoint) *NullablePointStream {
	return self.AddAll(arg...)
}
func (self *NullablePointStream) Contains(arg NullablePoint) bool {
	return self.FindIndex(func(_arg NullablePoint, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullablePointStream) Clean() *NullablePointStream {
	*self = NullablePointStreamOf()
	return self
}
func (self *NullablePointStream) Delete(index int) *NullablePointStream {
	return self.DeleteRange(index, index)
}
func (self *NullablePointStream) DeleteRange(startIndex, endIndex int) *NullablePointStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullablePointStream) Distinct() *NullablePointStream {
	caches := map[string]bool{}
	result := NullablePointStreamOf()
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
func (self *NullablePointStream) Each(fn func(NullablePoint)) *NullablePointStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullablePointStream) EachRight(fn func(NullablePoint)) *NullablePointStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullablePointStream) Equals(arr []NullablePoint) bool {
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
func (self *NullablePointStream) Filter(fn func(NullablePoint, int) bool) *NullablePointStream {
	result := NullablePointStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullablePointStream) FilterSlim(fn func(NullablePoint, int) bool) *NullablePointStream {
	result := NullablePointStreamOf()
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
func (self *NullablePointStream) Find(fn func(NullablePoint, int) bool) *NullablePoint {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullablePointStream) FindOr(fn func(NullablePoint, int) bool, or NullablePoint) NullablePoint {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullablePointStream) FindIndex(fn func(NullablePoint, int) bool) int {
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
func (self *NullablePointStream) First() *NullablePoint {
	return self.Get(0)
}
func (self *NullablePointStream) FirstOr(arg NullablePoint) NullablePoint {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullablePointStream) ForEach(fn func(NullablePoint, int)) *NullablePointStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullablePointStream) ForEachRight(fn func(NullablePoint, int)) *NullablePointStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullablePointStream) GroupBy(fn func(NullablePoint, int) string) map[string][]NullablePoint {
	m := map[string][]NullablePoint{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullablePointStream) GroupByValues(fn func(NullablePoint, int) string) [][]NullablePoint {
	var tmp [][]NullablePoint
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullablePointStream) IndexOf(arg NullablePoint) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullablePointStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullablePointStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullablePointStream) Last() *NullablePoint {
	return self.Get(self.Len() - 1)
}
func (self *NullablePointStream) LastOr(arg NullablePoint) NullablePoint {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullablePointStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullablePointStream) Limit(limit int) *NullablePointStream {
	self.Slice(0, limit)
	return self
}

func (self *NullablePointStream) Map(fn func(NullablePoint, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullablePointStream) Map2Int(fn func(NullablePoint, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullablePointStream) Map2Int32(fn func(NullablePoint, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullablePointStream) Map2Int64(fn func(NullablePoint, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullablePointStream) Map2Float32(fn func(NullablePoint, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullablePointStream) Map2Float64(fn func(NullablePoint, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullablePointStream) Map2Bool(fn func(NullablePoint, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullablePointStream) Map2Bytes(fn func(NullablePoint, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullablePointStream) Map2String(fn func(NullablePoint, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullablePointStream) Max(fn func(NullablePoint, int) float64) *NullablePoint {
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
func (self *NullablePointStream) Min(fn func(NullablePoint, int) float64) *NullablePoint {
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
func (self *NullablePointStream) NoneMatch(fn func(NullablePoint, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullablePointStream) Get(index int) *NullablePoint {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullablePointStream) GetOr(index int, arg NullablePoint) NullablePoint {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullablePointStream) Peek(fn func(*NullablePoint, int)) *NullablePointStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullablePointStream) Reduce(fn func(NullablePoint, NullablePoint, int) NullablePoint) *NullablePointStream {
	return self.ReduceInit(fn, NullablePoint{})
}
func (self *NullablePointStream) ReduceInit(fn func(NullablePoint, NullablePoint, int) NullablePoint, initialValue NullablePoint) *NullablePointStream {
	result := NullablePointStreamOf()
	self.ForEach(func(v NullablePoint, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullablePointStream) ReduceInterface(fn func(interface{}, NullablePoint, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullablePoint{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullablePointStream) ReduceString(fn func(string, NullablePoint, int) string) []string {
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
func (self *NullablePointStream) ReduceInt(fn func(int, NullablePoint, int) int) []int {
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
func (self *NullablePointStream) ReduceInt32(fn func(int32, NullablePoint, int) int32) []int32 {
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
func (self *NullablePointStream) ReduceInt64(fn func(int64, NullablePoint, int) int64) []int64 {
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
func (self *NullablePointStream) ReduceFloat32(fn func(float32, NullablePoint, int) float32) []float32 {
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
func (self *NullablePointStream) ReduceFloat64(fn func(float64, NullablePoint, int) float64) []float64 {
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
func (self *NullablePointStream) ReduceBool(fn func(bool, NullablePoint, int) bool) []bool {
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
func (self *NullablePointStream) Reverse() *NullablePointStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullablePointStream) Replace(fn func(NullablePoint, int) NullablePoint) *NullablePointStream {
	return self.ForEach(func(v NullablePoint, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullablePointStream) Select(fn func(NullablePoint) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullablePointStream) Set(index int, val NullablePoint) *NullablePointStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullablePointStream) Skip(skip int) *NullablePointStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullablePointStream) SkippingEach(fn func(NullablePoint, int) int) *NullablePointStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullablePointStream) Slice(startIndex, n int) *NullablePointStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullablePoint{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullablePointStream) Sort(fn func(i, j int) bool) *NullablePointStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullablePointStream) Tail() *NullablePoint {
	return self.Last()
}
func (self *NullablePointStream) TailOr(arg NullablePoint) NullablePoint {
	return self.LastOr(arg)
}
func (self *NullablePointStream) ToList() []NullablePoint {
	return self.Val()
}
func (self *NullablePointStream) Unique() *NullablePointStream {
	return self.Distinct()
}
func (self *NullablePointStream) Val() []NullablePoint {
	if self == nil {
		return []NullablePoint{}
	}
	return *self.Copy()
}
func (self *NullablePointStream) While(fn func(NullablePoint, int) bool) *NullablePointStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullablePointStream) Where(fn func(NullablePoint) bool) *NullablePointStream {
	result := NullablePointStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullablePointStream) WhereSlim(fn func(NullablePoint) bool) *NullablePointStream {
	result := NullablePointStreamOf()
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
