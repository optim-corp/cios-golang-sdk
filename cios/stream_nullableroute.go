/*
 * Collection utility of NullableRoute Struct
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

type NullableRouteStream []NullableRoute

func NullableRouteStreamOf(arg ...NullableRoute) NullableRouteStream {
	return arg
}
func NullableRouteStreamFrom(arg []NullableRoute) NullableRouteStream {
	return arg
}
func CreateNullableRouteStream(arg ...NullableRoute) *NullableRouteStream {
	tmp := NullableRouteStreamOf(arg...)
	return &tmp
}
func GenerateNullableRouteStream(arg []NullableRoute) *NullableRouteStream {
	tmp := NullableRouteStreamFrom(arg)
	return &tmp
}

func (self *NullableRouteStream) Add(arg NullableRoute) *NullableRouteStream {
	return self.AddAll(arg)
}
func (self *NullableRouteStream) AddAll(arg ...NullableRoute) *NullableRouteStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableRouteStream) AddSafe(arg *NullableRoute) *NullableRouteStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableRouteStream) Aggregate(fn func(NullableRoute, NullableRoute) NullableRoute) *NullableRouteStream {
	result := NullableRouteStreamOf()
	self.ForEach(func(v NullableRoute, i int) {
		if i == 0 {
			result.Add(fn(NullableRoute{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableRouteStream) AllMatch(fn func(NullableRoute, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableRouteStream) AnyMatch(fn func(NullableRoute, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableRouteStream) Clone() *NullableRouteStream {
	temp := make([]NullableRoute, self.Len())
	copy(temp, *self)
	return (*NullableRouteStream)(&temp)
}
func (self *NullableRouteStream) Copy() *NullableRouteStream {
	return self.Clone()
}
func (self *NullableRouteStream) Concat(arg []NullableRoute) *NullableRouteStream {
	return self.AddAll(arg...)
}
func (self *NullableRouteStream) Contains(arg NullableRoute) bool {
	return self.FindIndex(func(_arg NullableRoute, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableRouteStream) Clean() *NullableRouteStream {
	*self = NullableRouteStreamOf()
	return self
}
func (self *NullableRouteStream) Delete(index int) *NullableRouteStream {
	return self.DeleteRange(index, index)
}
func (self *NullableRouteStream) DeleteRange(startIndex, endIndex int) *NullableRouteStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableRouteStream) Distinct() *NullableRouteStream {
	caches := map[string]bool{}
	result := NullableRouteStreamOf()
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
func (self *NullableRouteStream) Each(fn func(NullableRoute)) *NullableRouteStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableRouteStream) EachRight(fn func(NullableRoute)) *NullableRouteStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableRouteStream) Equals(arr []NullableRoute) bool {
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
func (self *NullableRouteStream) Filter(fn func(NullableRoute, int) bool) *NullableRouteStream {
	result := NullableRouteStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableRouteStream) FilterSlim(fn func(NullableRoute, int) bool) *NullableRouteStream {
	result := NullableRouteStreamOf()
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
func (self *NullableRouteStream) Find(fn func(NullableRoute, int) bool) *NullableRoute {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableRouteStream) FindOr(fn func(NullableRoute, int) bool, or NullableRoute) NullableRoute {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableRouteStream) FindIndex(fn func(NullableRoute, int) bool) int {
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
func (self *NullableRouteStream) First() *NullableRoute {
	return self.Get(0)
}
func (self *NullableRouteStream) FirstOr(arg NullableRoute) NullableRoute {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableRouteStream) ForEach(fn func(NullableRoute, int)) *NullableRouteStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableRouteStream) ForEachRight(fn func(NullableRoute, int)) *NullableRouteStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableRouteStream) GroupBy(fn func(NullableRoute, int) string) map[string][]NullableRoute {
	m := map[string][]NullableRoute{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableRouteStream) GroupByValues(fn func(NullableRoute, int) string) [][]NullableRoute {
	var tmp [][]NullableRoute
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableRouteStream) IndexOf(arg NullableRoute) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableRouteStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableRouteStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableRouteStream) Last() *NullableRoute {
	return self.Get(self.Len() - 1)
}
func (self *NullableRouteStream) LastOr(arg NullableRoute) NullableRoute {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableRouteStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableRouteStream) Limit(limit int) *NullableRouteStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableRouteStream) Map(fn func(NullableRoute, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableRouteStream) Map2Int(fn func(NullableRoute, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableRouteStream) Map2Int32(fn func(NullableRoute, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableRouteStream) Map2Int64(fn func(NullableRoute, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableRouteStream) Map2Float32(fn func(NullableRoute, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableRouteStream) Map2Float64(fn func(NullableRoute, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableRouteStream) Map2Bool(fn func(NullableRoute, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableRouteStream) Map2Bytes(fn func(NullableRoute, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableRouteStream) Map2String(fn func(NullableRoute, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableRouteStream) Max(fn func(NullableRoute, int) float64) *NullableRoute {
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
func (self *NullableRouteStream) Min(fn func(NullableRoute, int) float64) *NullableRoute {
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
func (self *NullableRouteStream) NoneMatch(fn func(NullableRoute, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableRouteStream) Get(index int) *NullableRoute {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableRouteStream) GetOr(index int, arg NullableRoute) NullableRoute {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableRouteStream) Peek(fn func(*NullableRoute, int)) *NullableRouteStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableRouteStream) Reduce(fn func(NullableRoute, NullableRoute, int) NullableRoute) *NullableRouteStream {
	return self.ReduceInit(fn, NullableRoute{})
}
func (self *NullableRouteStream) ReduceInit(fn func(NullableRoute, NullableRoute, int) NullableRoute, initialValue NullableRoute) *NullableRouteStream {
	result := NullableRouteStreamOf()
	self.ForEach(func(v NullableRoute, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableRouteStream) ReduceInterface(fn func(interface{}, NullableRoute, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableRoute{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableRouteStream) ReduceString(fn func(string, NullableRoute, int) string) []string {
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
func (self *NullableRouteStream) ReduceInt(fn func(int, NullableRoute, int) int) []int {
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
func (self *NullableRouteStream) ReduceInt32(fn func(int32, NullableRoute, int) int32) []int32 {
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
func (self *NullableRouteStream) ReduceInt64(fn func(int64, NullableRoute, int) int64) []int64 {
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
func (self *NullableRouteStream) ReduceFloat32(fn func(float32, NullableRoute, int) float32) []float32 {
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
func (self *NullableRouteStream) ReduceFloat64(fn func(float64, NullableRoute, int) float64) []float64 {
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
func (self *NullableRouteStream) ReduceBool(fn func(bool, NullableRoute, int) bool) []bool {
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
func (self *NullableRouteStream) Reverse() *NullableRouteStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableRouteStream) Replace(fn func(NullableRoute, int) NullableRoute) *NullableRouteStream {
	return self.ForEach(func(v NullableRoute, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableRouteStream) Select(fn func(NullableRoute) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableRouteStream) Set(index int, val NullableRoute) *NullableRouteStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableRouteStream) Skip(skip int) *NullableRouteStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableRouteStream) SkippingEach(fn func(NullableRoute, int) int) *NullableRouteStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableRouteStream) Slice(startIndex, n int) *NullableRouteStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableRoute{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableRouteStream) Sort(fn func(i, j int) bool) *NullableRouteStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableRouteStream) Tail() *NullableRoute {
	return self.Last()
}
func (self *NullableRouteStream) TailOr(arg NullableRoute) NullableRoute {
	return self.LastOr(arg)
}
func (self *NullableRouteStream) ToList() []NullableRoute {
	return self.Val()
}
func (self *NullableRouteStream) Unique() *NullableRouteStream {
	return self.Distinct()
}
func (self *NullableRouteStream) Val() []NullableRoute {
	if self == nil {
		return []NullableRoute{}
	}
	return *self.Copy()
}
func (self *NullableRouteStream) While(fn func(NullableRoute, int) bool) *NullableRouteStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableRouteStream) Where(fn func(NullableRoute) bool) *NullableRouteStream {
	result := NullableRouteStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableRouteStream) WhereSlim(fn func(NullableRoute) bool) *NullableRouteStream {
	result := NullableRouteStreamOf()
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