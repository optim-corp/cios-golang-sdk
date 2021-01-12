/*
 * Collection utility of NullableMultiplePoint Struct
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

type NullableMultiplePointStream []NullableMultiplePoint

func NullableMultiplePointStreamOf(arg ...NullableMultiplePoint) NullableMultiplePointStream {
	return arg
}
func NullableMultiplePointStreamFrom(arg []NullableMultiplePoint) NullableMultiplePointStream {
	return arg
}
func CreateNullableMultiplePointStream(arg ...NullableMultiplePoint) *NullableMultiplePointStream {
	tmp := NullableMultiplePointStreamOf(arg...)
	return &tmp
}
func GenerateNullableMultiplePointStream(arg []NullableMultiplePoint) *NullableMultiplePointStream {
	tmp := NullableMultiplePointStreamFrom(arg)
	return &tmp
}

func (self *NullableMultiplePointStream) Add(arg NullableMultiplePoint) *NullableMultiplePointStream {
	return self.AddAll(arg)
}
func (self *NullableMultiplePointStream) AddAll(arg ...NullableMultiplePoint) *NullableMultiplePointStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableMultiplePointStream) AddSafe(arg *NullableMultiplePoint) *NullableMultiplePointStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableMultiplePointStream) Aggregate(fn func(NullableMultiplePoint, NullableMultiplePoint) NullableMultiplePoint) *NullableMultiplePointStream {
	result := NullableMultiplePointStreamOf()
	self.ForEach(func(v NullableMultiplePoint, i int) {
		if i == 0 {
			result.Add(fn(NullableMultiplePoint{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableMultiplePointStream) AllMatch(fn func(NullableMultiplePoint, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableMultiplePointStream) AnyMatch(fn func(NullableMultiplePoint, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableMultiplePointStream) Clone() *NullableMultiplePointStream {
	temp := make([]NullableMultiplePoint, self.Len())
	copy(temp, *self)
	return (*NullableMultiplePointStream)(&temp)
}
func (self *NullableMultiplePointStream) Copy() *NullableMultiplePointStream {
	return self.Clone()
}
func (self *NullableMultiplePointStream) Concat(arg []NullableMultiplePoint) *NullableMultiplePointStream {
	return self.AddAll(arg...)
}
func (self *NullableMultiplePointStream) Contains(arg NullableMultiplePoint) bool {
	return self.FindIndex(func(_arg NullableMultiplePoint, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableMultiplePointStream) Clean() *NullableMultiplePointStream {
	*self = NullableMultiplePointStreamOf()
	return self
}
func (self *NullableMultiplePointStream) Delete(index int) *NullableMultiplePointStream {
	return self.DeleteRange(index, index)
}
func (self *NullableMultiplePointStream) DeleteRange(startIndex, endIndex int) *NullableMultiplePointStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableMultiplePointStream) Distinct() *NullableMultiplePointStream {
	caches := map[string]bool{}
	result := NullableMultiplePointStreamOf()
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
func (self *NullableMultiplePointStream) Each(fn func(NullableMultiplePoint)) *NullableMultiplePointStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableMultiplePointStream) EachRight(fn func(NullableMultiplePoint)) *NullableMultiplePointStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableMultiplePointStream) Equals(arr []NullableMultiplePoint) bool {
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
func (self *NullableMultiplePointStream) Filter(fn func(NullableMultiplePoint, int) bool) *NullableMultiplePointStream {
	result := NullableMultiplePointStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultiplePointStream) FilterSlim(fn func(NullableMultiplePoint, int) bool) *NullableMultiplePointStream {
	result := NullableMultiplePointStreamOf()
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
func (self *NullableMultiplePointStream) Find(fn func(NullableMultiplePoint, int) bool) *NullableMultiplePoint {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableMultiplePointStream) FindOr(fn func(NullableMultiplePoint, int) bool, or NullableMultiplePoint) NullableMultiplePoint {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableMultiplePointStream) FindIndex(fn func(NullableMultiplePoint, int) bool) int {
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
func (self *NullableMultiplePointStream) First() *NullableMultiplePoint {
	return self.Get(0)
}
func (self *NullableMultiplePointStream) FirstOr(arg NullableMultiplePoint) NullableMultiplePoint {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultiplePointStream) ForEach(fn func(NullableMultiplePoint, int)) *NullableMultiplePointStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableMultiplePointStream) ForEachRight(fn func(NullableMultiplePoint, int)) *NullableMultiplePointStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableMultiplePointStream) GroupBy(fn func(NullableMultiplePoint, int) string) map[string][]NullableMultiplePoint {
	m := map[string][]NullableMultiplePoint{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableMultiplePointStream) GroupByValues(fn func(NullableMultiplePoint, int) string) [][]NullableMultiplePoint {
	var tmp [][]NullableMultiplePoint
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableMultiplePointStream) IndexOf(arg NullableMultiplePoint) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableMultiplePointStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableMultiplePointStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableMultiplePointStream) Last() *NullableMultiplePoint {
	return self.Get(self.Len() - 1)
}
func (self *NullableMultiplePointStream) LastOr(arg NullableMultiplePoint) NullableMultiplePoint {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultiplePointStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableMultiplePointStream) Limit(limit int) *NullableMultiplePointStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableMultiplePointStream) Map(fn func(NullableMultiplePoint, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultiplePointStream) Map2Int(fn func(NullableMultiplePoint, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultiplePointStream) Map2Int32(fn func(NullableMultiplePoint, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultiplePointStream) Map2Int64(fn func(NullableMultiplePoint, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultiplePointStream) Map2Float32(fn func(NullableMultiplePoint, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultiplePointStream) Map2Float64(fn func(NullableMultiplePoint, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultiplePointStream) Map2Bool(fn func(NullableMultiplePoint, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultiplePointStream) Map2Bytes(fn func(NullableMultiplePoint, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultiplePointStream) Map2String(fn func(NullableMultiplePoint, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultiplePointStream) Max(fn func(NullableMultiplePoint, int) float64) *NullableMultiplePoint {
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
func (self *NullableMultiplePointStream) Min(fn func(NullableMultiplePoint, int) float64) *NullableMultiplePoint {
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
func (self *NullableMultiplePointStream) NoneMatch(fn func(NullableMultiplePoint, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableMultiplePointStream) Get(index int) *NullableMultiplePoint {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableMultiplePointStream) GetOr(index int, arg NullableMultiplePoint) NullableMultiplePoint {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultiplePointStream) Peek(fn func(*NullableMultiplePoint, int)) *NullableMultiplePointStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableMultiplePointStream) Reduce(fn func(NullableMultiplePoint, NullableMultiplePoint, int) NullableMultiplePoint) *NullableMultiplePointStream {
	return self.ReduceInit(fn, NullableMultiplePoint{})
}
func (self *NullableMultiplePointStream) ReduceInit(fn func(NullableMultiplePoint, NullableMultiplePoint, int) NullableMultiplePoint, initialValue NullableMultiplePoint) *NullableMultiplePointStream {
	result := NullableMultiplePointStreamOf()
	self.ForEach(func(v NullableMultiplePoint, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableMultiplePointStream) ReduceInterface(fn func(interface{}, NullableMultiplePoint, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableMultiplePoint{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultiplePointStream) ReduceString(fn func(string, NullableMultiplePoint, int) string) []string {
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
func (self *NullableMultiplePointStream) ReduceInt(fn func(int, NullableMultiplePoint, int) int) []int {
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
func (self *NullableMultiplePointStream) ReduceInt32(fn func(int32, NullableMultiplePoint, int) int32) []int32 {
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
func (self *NullableMultiplePointStream) ReduceInt64(fn func(int64, NullableMultiplePoint, int) int64) []int64 {
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
func (self *NullableMultiplePointStream) ReduceFloat32(fn func(float32, NullableMultiplePoint, int) float32) []float32 {
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
func (self *NullableMultiplePointStream) ReduceFloat64(fn func(float64, NullableMultiplePoint, int) float64) []float64 {
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
func (self *NullableMultiplePointStream) ReduceBool(fn func(bool, NullableMultiplePoint, int) bool) []bool {
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
func (self *NullableMultiplePointStream) Reverse() *NullableMultiplePointStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableMultiplePointStream) Replace(fn func(NullableMultiplePoint, int) NullableMultiplePoint) *NullableMultiplePointStream {
	return self.ForEach(func(v NullableMultiplePoint, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableMultiplePointStream) Select(fn func(NullableMultiplePoint) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableMultiplePointStream) Set(index int, val NullableMultiplePoint) *NullableMultiplePointStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableMultiplePointStream) Skip(skip int) *NullableMultiplePointStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableMultiplePointStream) SkippingEach(fn func(NullableMultiplePoint, int) int) *NullableMultiplePointStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableMultiplePointStream) Slice(startIndex, n int) *NullableMultiplePointStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableMultiplePoint{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableMultiplePointStream) Sort(fn func(i, j int) bool) *NullableMultiplePointStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableMultiplePointStream) Tail() *NullableMultiplePoint {
	return self.Last()
}
func (self *NullableMultiplePointStream) TailOr(arg NullableMultiplePoint) NullableMultiplePoint {
	return self.LastOr(arg)
}
func (self *NullableMultiplePointStream) ToList() []NullableMultiplePoint {
	return self.Val()
}
func (self *NullableMultiplePointStream) Unique() *NullableMultiplePointStream {
	return self.Distinct()
}
func (self *NullableMultiplePointStream) Val() []NullableMultiplePoint {
	if self == nil {
		return []NullableMultiplePoint{}
	}
	return *self.Copy()
}
func (self *NullableMultiplePointStream) While(fn func(NullableMultiplePoint, int) bool) *NullableMultiplePointStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableMultiplePointStream) Where(fn func(NullableMultiplePoint) bool) *NullableMultiplePointStream {
	result := NullableMultiplePointStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultiplePointStream) WhereSlim(fn func(NullableMultiplePoint) bool) *NullableMultiplePointStream {
	result := NullableMultiplePointStreamOf()
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