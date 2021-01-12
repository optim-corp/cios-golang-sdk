/*
 * Collection utility of NullableSeriesDataLocation Struct
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

type NullableSeriesDataLocationStream []NullableSeriesDataLocation

func NullableSeriesDataLocationStreamOf(arg ...NullableSeriesDataLocation) NullableSeriesDataLocationStream {
	return arg
}
func NullableSeriesDataLocationStreamFrom(arg []NullableSeriesDataLocation) NullableSeriesDataLocationStream {
	return arg
}
func CreateNullableSeriesDataLocationStream(arg ...NullableSeriesDataLocation) *NullableSeriesDataLocationStream {
	tmp := NullableSeriesDataLocationStreamOf(arg...)
	return &tmp
}
func GenerateNullableSeriesDataLocationStream(arg []NullableSeriesDataLocation) *NullableSeriesDataLocationStream {
	tmp := NullableSeriesDataLocationStreamFrom(arg)
	return &tmp
}

func (self *NullableSeriesDataLocationStream) Add(arg NullableSeriesDataLocation) *NullableSeriesDataLocationStream {
	return self.AddAll(arg)
}
func (self *NullableSeriesDataLocationStream) AddAll(arg ...NullableSeriesDataLocation) *NullableSeriesDataLocationStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableSeriesDataLocationStream) AddSafe(arg *NullableSeriesDataLocation) *NullableSeriesDataLocationStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableSeriesDataLocationStream) Aggregate(fn func(NullableSeriesDataLocation, NullableSeriesDataLocation) NullableSeriesDataLocation) *NullableSeriesDataLocationStream {
	result := NullableSeriesDataLocationStreamOf()
	self.ForEach(func(v NullableSeriesDataLocation, i int) {
		if i == 0 {
			result.Add(fn(NullableSeriesDataLocation{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableSeriesDataLocationStream) AllMatch(fn func(NullableSeriesDataLocation, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableSeriesDataLocationStream) AnyMatch(fn func(NullableSeriesDataLocation, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableSeriesDataLocationStream) Clone() *NullableSeriesDataLocationStream {
	temp := make([]NullableSeriesDataLocation, self.Len())
	copy(temp, *self)
	return (*NullableSeriesDataLocationStream)(&temp)
}
func (self *NullableSeriesDataLocationStream) Copy() *NullableSeriesDataLocationStream {
	return self.Clone()
}
func (self *NullableSeriesDataLocationStream) Concat(arg []NullableSeriesDataLocation) *NullableSeriesDataLocationStream {
	return self.AddAll(arg...)
}
func (self *NullableSeriesDataLocationStream) Contains(arg NullableSeriesDataLocation) bool {
	return self.FindIndex(func(_arg NullableSeriesDataLocation, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableSeriesDataLocationStream) Clean() *NullableSeriesDataLocationStream {
	*self = NullableSeriesDataLocationStreamOf()
	return self
}
func (self *NullableSeriesDataLocationStream) Delete(index int) *NullableSeriesDataLocationStream {
	return self.DeleteRange(index, index)
}
func (self *NullableSeriesDataLocationStream) DeleteRange(startIndex, endIndex int) *NullableSeriesDataLocationStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableSeriesDataLocationStream) Distinct() *NullableSeriesDataLocationStream {
	caches := map[string]bool{}
	result := NullableSeriesDataLocationStreamOf()
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
func (self *NullableSeriesDataLocationStream) Each(fn func(NullableSeriesDataLocation)) *NullableSeriesDataLocationStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableSeriesDataLocationStream) EachRight(fn func(NullableSeriesDataLocation)) *NullableSeriesDataLocationStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableSeriesDataLocationStream) Equals(arr []NullableSeriesDataLocation) bool {
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
func (self *NullableSeriesDataLocationStream) Filter(fn func(NullableSeriesDataLocation, int) bool) *NullableSeriesDataLocationStream {
	result := NullableSeriesDataLocationStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSeriesDataLocationStream) FilterSlim(fn func(NullableSeriesDataLocation, int) bool) *NullableSeriesDataLocationStream {
	result := NullableSeriesDataLocationStreamOf()
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
func (self *NullableSeriesDataLocationStream) Find(fn func(NullableSeriesDataLocation, int) bool) *NullableSeriesDataLocation {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableSeriesDataLocationStream) FindOr(fn func(NullableSeriesDataLocation, int) bool, or NullableSeriesDataLocation) NullableSeriesDataLocation {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableSeriesDataLocationStream) FindIndex(fn func(NullableSeriesDataLocation, int) bool) int {
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
func (self *NullableSeriesDataLocationStream) First() *NullableSeriesDataLocation {
	return self.Get(0)
}
func (self *NullableSeriesDataLocationStream) FirstOr(arg NullableSeriesDataLocation) NullableSeriesDataLocation {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSeriesDataLocationStream) ForEach(fn func(NullableSeriesDataLocation, int)) *NullableSeriesDataLocationStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableSeriesDataLocationStream) ForEachRight(fn func(NullableSeriesDataLocation, int)) *NullableSeriesDataLocationStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableSeriesDataLocationStream) GroupBy(fn func(NullableSeriesDataLocation, int) string) map[string][]NullableSeriesDataLocation {
	m := map[string][]NullableSeriesDataLocation{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableSeriesDataLocationStream) GroupByValues(fn func(NullableSeriesDataLocation, int) string) [][]NullableSeriesDataLocation {
	var tmp [][]NullableSeriesDataLocation
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableSeriesDataLocationStream) IndexOf(arg NullableSeriesDataLocation) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableSeriesDataLocationStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableSeriesDataLocationStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableSeriesDataLocationStream) Last() *NullableSeriesDataLocation {
	return self.Get(self.Len() - 1)
}
func (self *NullableSeriesDataLocationStream) LastOr(arg NullableSeriesDataLocation) NullableSeriesDataLocation {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSeriesDataLocationStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableSeriesDataLocationStream) Limit(limit int) *NullableSeriesDataLocationStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableSeriesDataLocationStream) Map(fn func(NullableSeriesDataLocation, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataLocationStream) Map2Int(fn func(NullableSeriesDataLocation, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataLocationStream) Map2Int32(fn func(NullableSeriesDataLocation, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataLocationStream) Map2Int64(fn func(NullableSeriesDataLocation, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataLocationStream) Map2Float32(fn func(NullableSeriesDataLocation, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataLocationStream) Map2Float64(fn func(NullableSeriesDataLocation, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataLocationStream) Map2Bool(fn func(NullableSeriesDataLocation, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataLocationStream) Map2Bytes(fn func(NullableSeriesDataLocation, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataLocationStream) Map2String(fn func(NullableSeriesDataLocation, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSeriesDataLocationStream) Max(fn func(NullableSeriesDataLocation, int) float64) *NullableSeriesDataLocation {
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
func (self *NullableSeriesDataLocationStream) Min(fn func(NullableSeriesDataLocation, int) float64) *NullableSeriesDataLocation {
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
func (self *NullableSeriesDataLocationStream) NoneMatch(fn func(NullableSeriesDataLocation, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableSeriesDataLocationStream) Get(index int) *NullableSeriesDataLocation {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableSeriesDataLocationStream) GetOr(index int, arg NullableSeriesDataLocation) NullableSeriesDataLocation {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSeriesDataLocationStream) Peek(fn func(*NullableSeriesDataLocation, int)) *NullableSeriesDataLocationStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableSeriesDataLocationStream) Reduce(fn func(NullableSeriesDataLocation, NullableSeriesDataLocation, int) NullableSeriesDataLocation) *NullableSeriesDataLocationStream {
	return self.ReduceInit(fn, NullableSeriesDataLocation{})
}
func (self *NullableSeriesDataLocationStream) ReduceInit(fn func(NullableSeriesDataLocation, NullableSeriesDataLocation, int) NullableSeriesDataLocation, initialValue NullableSeriesDataLocation) *NullableSeriesDataLocationStream {
	result := NullableSeriesDataLocationStreamOf()
	self.ForEach(func(v NullableSeriesDataLocation, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableSeriesDataLocationStream) ReduceInterface(fn func(interface{}, NullableSeriesDataLocation, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableSeriesDataLocation{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableSeriesDataLocationStream) ReduceString(fn func(string, NullableSeriesDataLocation, int) string) []string {
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
func (self *NullableSeriesDataLocationStream) ReduceInt(fn func(int, NullableSeriesDataLocation, int) int) []int {
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
func (self *NullableSeriesDataLocationStream) ReduceInt32(fn func(int32, NullableSeriesDataLocation, int) int32) []int32 {
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
func (self *NullableSeriesDataLocationStream) ReduceInt64(fn func(int64, NullableSeriesDataLocation, int) int64) []int64 {
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
func (self *NullableSeriesDataLocationStream) ReduceFloat32(fn func(float32, NullableSeriesDataLocation, int) float32) []float32 {
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
func (self *NullableSeriesDataLocationStream) ReduceFloat64(fn func(float64, NullableSeriesDataLocation, int) float64) []float64 {
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
func (self *NullableSeriesDataLocationStream) ReduceBool(fn func(bool, NullableSeriesDataLocation, int) bool) []bool {
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
func (self *NullableSeriesDataLocationStream) Reverse() *NullableSeriesDataLocationStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableSeriesDataLocationStream) Replace(fn func(NullableSeriesDataLocation, int) NullableSeriesDataLocation) *NullableSeriesDataLocationStream {
	return self.ForEach(func(v NullableSeriesDataLocation, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableSeriesDataLocationStream) Select(fn func(NullableSeriesDataLocation) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableSeriesDataLocationStream) Set(index int, val NullableSeriesDataLocation) *NullableSeriesDataLocationStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableSeriesDataLocationStream) Skip(skip int) *NullableSeriesDataLocationStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableSeriesDataLocationStream) SkippingEach(fn func(NullableSeriesDataLocation, int) int) *NullableSeriesDataLocationStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableSeriesDataLocationStream) Slice(startIndex, n int) *NullableSeriesDataLocationStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableSeriesDataLocation{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableSeriesDataLocationStream) Sort(fn func(i, j int) bool) *NullableSeriesDataLocationStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableSeriesDataLocationStream) Tail() *NullableSeriesDataLocation {
	return self.Last()
}
func (self *NullableSeriesDataLocationStream) TailOr(arg NullableSeriesDataLocation) NullableSeriesDataLocation {
	return self.LastOr(arg)
}
func (self *NullableSeriesDataLocationStream) ToList() []NullableSeriesDataLocation {
	return self.Val()
}
func (self *NullableSeriesDataLocationStream) Unique() *NullableSeriesDataLocationStream {
	return self.Distinct()
}
func (self *NullableSeriesDataLocationStream) Val() []NullableSeriesDataLocation {
	if self == nil {
		return []NullableSeriesDataLocation{}
	}
	return *self.Copy()
}
func (self *NullableSeriesDataLocationStream) While(fn func(NullableSeriesDataLocation, int) bool) *NullableSeriesDataLocationStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableSeriesDataLocationStream) Where(fn func(NullableSeriesDataLocation) bool) *NullableSeriesDataLocationStream {
	result := NullableSeriesDataLocationStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSeriesDataLocationStream) WhereSlim(fn func(NullableSeriesDataLocation) bool) *NullableSeriesDataLocationStream {
	result := NullableSeriesDataLocationStreamOf()
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
