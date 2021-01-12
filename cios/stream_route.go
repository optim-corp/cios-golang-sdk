/*
 * Collection utility of Route Struct
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

type RouteStream []Route

func RouteStreamOf(arg ...Route) RouteStream {
	return arg
}
func RouteStreamFrom(arg []Route) RouteStream {
	return arg
}
func CreateRouteStream(arg ...Route) *RouteStream {
	tmp := RouteStreamOf(arg...)
	return &tmp
}
func GenerateRouteStream(arg []Route) *RouteStream {
	tmp := RouteStreamFrom(arg)
	return &tmp
}

func (self *RouteStream) Add(arg Route) *RouteStream {
	return self.AddAll(arg)
}
func (self *RouteStream) AddAll(arg ...Route) *RouteStream {
	*self = append(*self, arg...)
	return self
}
func (self *RouteStream) AddSafe(arg *Route) *RouteStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *RouteStream) Aggregate(fn func(Route, Route) Route) *RouteStream {
	result := RouteStreamOf()
	self.ForEach(func(v Route, i int) {
		if i == 0 {
			result.Add(fn(Route{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *RouteStream) AllMatch(fn func(Route, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *RouteStream) AnyMatch(fn func(Route, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *RouteStream) Clone() *RouteStream {
	temp := make([]Route, self.Len())
	copy(temp, *self)
	return (*RouteStream)(&temp)
}
func (self *RouteStream) Copy() *RouteStream {
	return self.Clone()
}
func (self *RouteStream) Concat(arg []Route) *RouteStream {
	return self.AddAll(arg...)
}
func (self *RouteStream) Contains(arg Route) bool {
	return self.FindIndex(func(_arg Route, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *RouteStream) Clean() *RouteStream {
	*self = RouteStreamOf()
	return self
}
func (self *RouteStream) Delete(index int) *RouteStream {
	return self.DeleteRange(index, index)
}
func (self *RouteStream) DeleteRange(startIndex, endIndex int) *RouteStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *RouteStream) Distinct() *RouteStream {
	caches := map[string]bool{}
	result := RouteStreamOf()
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
func (self *RouteStream) Each(fn func(Route)) *RouteStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *RouteStream) EachRight(fn func(Route)) *RouteStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *RouteStream) Equals(arr []Route) bool {
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
func (self *RouteStream) Filter(fn func(Route, int) bool) *RouteStream {
	result := RouteStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *RouteStream) FilterSlim(fn func(Route, int) bool) *RouteStream {
	result := RouteStreamOf()
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
func (self *RouteStream) Find(fn func(Route, int) bool) *Route {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *RouteStream) FindOr(fn func(Route, int) bool, or Route) Route {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *RouteStream) FindIndex(fn func(Route, int) bool) int {
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
func (self *RouteStream) First() *Route {
	return self.Get(0)
}
func (self *RouteStream) FirstOr(arg Route) Route {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *RouteStream) ForEach(fn func(Route, int)) *RouteStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *RouteStream) ForEachRight(fn func(Route, int)) *RouteStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *RouteStream) GroupBy(fn func(Route, int) string) map[string][]Route {
	m := map[string][]Route{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *RouteStream) GroupByValues(fn func(Route, int) string) [][]Route {
	var tmp [][]Route
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *RouteStream) IndexOf(arg Route) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *RouteStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *RouteStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *RouteStream) Last() *Route {
	return self.Get(self.Len() - 1)
}
func (self *RouteStream) LastOr(arg Route) Route {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *RouteStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *RouteStream) Limit(limit int) *RouteStream {
	self.Slice(0, limit)
	return self
}

func (self *RouteStream) Map(fn func(Route, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RouteStream) Map2Int(fn func(Route, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RouteStream) Map2Int32(fn func(Route, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RouteStream) Map2Int64(fn func(Route, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RouteStream) Map2Float32(fn func(Route, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RouteStream) Map2Float64(fn func(Route, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RouteStream) Map2Bool(fn func(Route, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RouteStream) Map2Bytes(fn func(Route, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RouteStream) Map2String(fn func(Route, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *RouteStream) Max(fn func(Route, int) float64) *Route {
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
func (self *RouteStream) Min(fn func(Route, int) float64) *Route {
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
func (self *RouteStream) NoneMatch(fn func(Route, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *RouteStream) Get(index int) *Route {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *RouteStream) GetOr(index int, arg Route) Route {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *RouteStream) Peek(fn func(*Route, int)) *RouteStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *RouteStream) Reduce(fn func(Route, Route, int) Route) *RouteStream {
	return self.ReduceInit(fn, Route{})
}
func (self *RouteStream) ReduceInit(fn func(Route, Route, int) Route, initialValue Route) *RouteStream {
	result := RouteStreamOf()
	self.ForEach(func(v Route, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *RouteStream) ReduceInterface(fn func(interface{}, Route, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(Route{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *RouteStream) ReduceString(fn func(string, Route, int) string) []string {
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
func (self *RouteStream) ReduceInt(fn func(int, Route, int) int) []int {
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
func (self *RouteStream) ReduceInt32(fn func(int32, Route, int) int32) []int32 {
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
func (self *RouteStream) ReduceInt64(fn func(int64, Route, int) int64) []int64 {
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
func (self *RouteStream) ReduceFloat32(fn func(float32, Route, int) float32) []float32 {
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
func (self *RouteStream) ReduceFloat64(fn func(float64, Route, int) float64) []float64 {
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
func (self *RouteStream) ReduceBool(fn func(bool, Route, int) bool) []bool {
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
func (self *RouteStream) Reverse() *RouteStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *RouteStream) Replace(fn func(Route, int) Route) *RouteStream {
	return self.ForEach(func(v Route, i int) { self.Set(i, fn(v, i)) })
}
func (self *RouteStream) Select(fn func(Route) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *RouteStream) Set(index int, val Route) *RouteStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *RouteStream) Skip(skip int) *RouteStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *RouteStream) SkippingEach(fn func(Route, int) int) *RouteStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *RouteStream) Slice(startIndex, n int) *RouteStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []Route{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *RouteStream) Sort(fn func(i, j int) bool) *RouteStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *RouteStream) Tail() *Route {
	return self.Last()
}
func (self *RouteStream) TailOr(arg Route) Route {
	return self.LastOr(arg)
}
func (self *RouteStream) ToList() []Route {
	return self.Val()
}
func (self *RouteStream) Unique() *RouteStream {
	return self.Distinct()
}
func (self *RouteStream) Val() []Route {
	if self == nil {
		return []Route{}
	}
	return *self.Copy()
}
func (self *RouteStream) While(fn func(Route, int) bool) *RouteStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *RouteStream) Where(fn func(Route) bool) *RouteStream {
	result := RouteStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *RouteStream) WhereSlim(fn func(Route) bool) *RouteStream {
	result := RouteStreamOf()
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
