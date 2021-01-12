/*
 * Collection utility of MultipleMemberInfo Struct
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

type MultipleMemberInfoStream []MultipleMemberInfo

func MultipleMemberInfoStreamOf(arg ...MultipleMemberInfo) MultipleMemberInfoStream {
	return arg
}
func MultipleMemberInfoStreamFrom(arg []MultipleMemberInfo) MultipleMemberInfoStream {
	return arg
}
func CreateMultipleMemberInfoStream(arg ...MultipleMemberInfo) *MultipleMemberInfoStream {
	tmp := MultipleMemberInfoStreamOf(arg...)
	return &tmp
}
func GenerateMultipleMemberInfoStream(arg []MultipleMemberInfo) *MultipleMemberInfoStream {
	tmp := MultipleMemberInfoStreamFrom(arg)
	return &tmp
}

func (self *MultipleMemberInfoStream) Add(arg MultipleMemberInfo) *MultipleMemberInfoStream {
	return self.AddAll(arg)
}
func (self *MultipleMemberInfoStream) AddAll(arg ...MultipleMemberInfo) *MultipleMemberInfoStream {
	*self = append(*self, arg...)
	return self
}
func (self *MultipleMemberInfoStream) AddSafe(arg *MultipleMemberInfo) *MultipleMemberInfoStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *MultipleMemberInfoStream) Aggregate(fn func(MultipleMemberInfo, MultipleMemberInfo) MultipleMemberInfo) *MultipleMemberInfoStream {
	result := MultipleMemberInfoStreamOf()
	self.ForEach(func(v MultipleMemberInfo, i int) {
		if i == 0 {
			result.Add(fn(MultipleMemberInfo{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *MultipleMemberInfoStream) AllMatch(fn func(MultipleMemberInfo, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *MultipleMemberInfoStream) AnyMatch(fn func(MultipleMemberInfo, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *MultipleMemberInfoStream) Clone() *MultipleMemberInfoStream {
	temp := make([]MultipleMemberInfo, self.Len())
	copy(temp, *self)
	return (*MultipleMemberInfoStream)(&temp)
}
func (self *MultipleMemberInfoStream) Copy() *MultipleMemberInfoStream {
	return self.Clone()
}
func (self *MultipleMemberInfoStream) Concat(arg []MultipleMemberInfo) *MultipleMemberInfoStream {
	return self.AddAll(arg...)
}
func (self *MultipleMemberInfoStream) Contains(arg MultipleMemberInfo) bool {
	return self.FindIndex(func(_arg MultipleMemberInfo, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *MultipleMemberInfoStream) Clean() *MultipleMemberInfoStream {
	*self = MultipleMemberInfoStreamOf()
	return self
}
func (self *MultipleMemberInfoStream) Delete(index int) *MultipleMemberInfoStream {
	return self.DeleteRange(index, index)
}
func (self *MultipleMemberInfoStream) DeleteRange(startIndex, endIndex int) *MultipleMemberInfoStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *MultipleMemberInfoStream) Distinct() *MultipleMemberInfoStream {
	caches := map[string]bool{}
	result := MultipleMemberInfoStreamOf()
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
func (self *MultipleMemberInfoStream) Each(fn func(MultipleMemberInfo)) *MultipleMemberInfoStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *MultipleMemberInfoStream) EachRight(fn func(MultipleMemberInfo)) *MultipleMemberInfoStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *MultipleMemberInfoStream) Equals(arr []MultipleMemberInfo) bool {
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
func (self *MultipleMemberInfoStream) Filter(fn func(MultipleMemberInfo, int) bool) *MultipleMemberInfoStream {
	result := MultipleMemberInfoStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleMemberInfoStream) FilterSlim(fn func(MultipleMemberInfo, int) bool) *MultipleMemberInfoStream {
	result := MultipleMemberInfoStreamOf()
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
func (self *MultipleMemberInfoStream) Find(fn func(MultipleMemberInfo, int) bool) *MultipleMemberInfo {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *MultipleMemberInfoStream) FindOr(fn func(MultipleMemberInfo, int) bool, or MultipleMemberInfo) MultipleMemberInfo {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *MultipleMemberInfoStream) FindIndex(fn func(MultipleMemberInfo, int) bool) int {
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
func (self *MultipleMemberInfoStream) First() *MultipleMemberInfo {
	return self.Get(0)
}
func (self *MultipleMemberInfoStream) FirstOr(arg MultipleMemberInfo) MultipleMemberInfo {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleMemberInfoStream) ForEach(fn func(MultipleMemberInfo, int)) *MultipleMemberInfoStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *MultipleMemberInfoStream) ForEachRight(fn func(MultipleMemberInfo, int)) *MultipleMemberInfoStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *MultipleMemberInfoStream) GroupBy(fn func(MultipleMemberInfo, int) string) map[string][]MultipleMemberInfo {
	m := map[string][]MultipleMemberInfo{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *MultipleMemberInfoStream) GroupByValues(fn func(MultipleMemberInfo, int) string) [][]MultipleMemberInfo {
	var tmp [][]MultipleMemberInfo
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *MultipleMemberInfoStream) IndexOf(arg MultipleMemberInfo) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *MultipleMemberInfoStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *MultipleMemberInfoStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *MultipleMemberInfoStream) Last() *MultipleMemberInfo {
	return self.Get(self.Len() - 1)
}
func (self *MultipleMemberInfoStream) LastOr(arg MultipleMemberInfo) MultipleMemberInfo {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleMemberInfoStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *MultipleMemberInfoStream) Limit(limit int) *MultipleMemberInfoStream {
	self.Slice(0, limit)
	return self
}

func (self *MultipleMemberInfoStream) Map(fn func(MultipleMemberInfo, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleMemberInfoStream) Map2Int(fn func(MultipleMemberInfo, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleMemberInfoStream) Map2Int32(fn func(MultipleMemberInfo, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleMemberInfoStream) Map2Int64(fn func(MultipleMemberInfo, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleMemberInfoStream) Map2Float32(fn func(MultipleMemberInfo, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleMemberInfoStream) Map2Float64(fn func(MultipleMemberInfo, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleMemberInfoStream) Map2Bool(fn func(MultipleMemberInfo, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleMemberInfoStream) Map2Bytes(fn func(MultipleMemberInfo, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleMemberInfoStream) Map2String(fn func(MultipleMemberInfo, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleMemberInfoStream) Max(fn func(MultipleMemberInfo, int) float64) *MultipleMemberInfo {
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
func (self *MultipleMemberInfoStream) Min(fn func(MultipleMemberInfo, int) float64) *MultipleMemberInfo {
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
func (self *MultipleMemberInfoStream) NoneMatch(fn func(MultipleMemberInfo, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *MultipleMemberInfoStream) Get(index int) *MultipleMemberInfo {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *MultipleMemberInfoStream) GetOr(index int, arg MultipleMemberInfo) MultipleMemberInfo {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleMemberInfoStream) Peek(fn func(*MultipleMemberInfo, int)) *MultipleMemberInfoStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *MultipleMemberInfoStream) Reduce(fn func(MultipleMemberInfo, MultipleMemberInfo, int) MultipleMemberInfo) *MultipleMemberInfoStream {
	return self.ReduceInit(fn, MultipleMemberInfo{})
}
func (self *MultipleMemberInfoStream) ReduceInit(fn func(MultipleMemberInfo, MultipleMemberInfo, int) MultipleMemberInfo, initialValue MultipleMemberInfo) *MultipleMemberInfoStream {
	result := MultipleMemberInfoStreamOf()
	self.ForEach(func(v MultipleMemberInfo, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *MultipleMemberInfoStream) ReduceInterface(fn func(interface{}, MultipleMemberInfo, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(MultipleMemberInfo{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *MultipleMemberInfoStream) ReduceString(fn func(string, MultipleMemberInfo, int) string) []string {
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
func (self *MultipleMemberInfoStream) ReduceInt(fn func(int, MultipleMemberInfo, int) int) []int {
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
func (self *MultipleMemberInfoStream) ReduceInt32(fn func(int32, MultipleMemberInfo, int) int32) []int32 {
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
func (self *MultipleMemberInfoStream) ReduceInt64(fn func(int64, MultipleMemberInfo, int) int64) []int64 {
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
func (self *MultipleMemberInfoStream) ReduceFloat32(fn func(float32, MultipleMemberInfo, int) float32) []float32 {
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
func (self *MultipleMemberInfoStream) ReduceFloat64(fn func(float64, MultipleMemberInfo, int) float64) []float64 {
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
func (self *MultipleMemberInfoStream) ReduceBool(fn func(bool, MultipleMemberInfo, int) bool) []bool {
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
func (self *MultipleMemberInfoStream) Reverse() *MultipleMemberInfoStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *MultipleMemberInfoStream) Replace(fn func(MultipleMemberInfo, int) MultipleMemberInfo) *MultipleMemberInfoStream {
	return self.ForEach(func(v MultipleMemberInfo, i int) { self.Set(i, fn(v, i)) })
}
func (self *MultipleMemberInfoStream) Select(fn func(MultipleMemberInfo) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *MultipleMemberInfoStream) Set(index int, val MultipleMemberInfo) *MultipleMemberInfoStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *MultipleMemberInfoStream) Skip(skip int) *MultipleMemberInfoStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *MultipleMemberInfoStream) SkippingEach(fn func(MultipleMemberInfo, int) int) *MultipleMemberInfoStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *MultipleMemberInfoStream) Slice(startIndex, n int) *MultipleMemberInfoStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []MultipleMemberInfo{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *MultipleMemberInfoStream) Sort(fn func(i, j int) bool) *MultipleMemberInfoStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *MultipleMemberInfoStream) Tail() *MultipleMemberInfo {
	return self.Last()
}
func (self *MultipleMemberInfoStream) TailOr(arg MultipleMemberInfo) MultipleMemberInfo {
	return self.LastOr(arg)
}
func (self *MultipleMemberInfoStream) ToList() []MultipleMemberInfo {
	return self.Val()
}
func (self *MultipleMemberInfoStream) Unique() *MultipleMemberInfoStream {
	return self.Distinct()
}
func (self *MultipleMemberInfoStream) Val() []MultipleMemberInfo {
	if self == nil {
		return []MultipleMemberInfo{}
	}
	return *self.Copy()
}
func (self *MultipleMemberInfoStream) While(fn func(MultipleMemberInfo, int) bool) *MultipleMemberInfoStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *MultipleMemberInfoStream) Where(fn func(MultipleMemberInfo) bool) *MultipleMemberInfoStream {
	result := MultipleMemberInfoStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleMemberInfoStream) WhereSlim(fn func(MultipleMemberInfo) bool) *MultipleMemberInfoStream {
	result := MultipleMemberInfoStreamOf()
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
