/*
 * Collection utility of MultipleDataStoreObject Struct
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

type MultipleDataStoreObjectStream []MultipleDataStoreObject

func MultipleDataStoreObjectStreamOf(arg ...MultipleDataStoreObject) MultipleDataStoreObjectStream {
	return arg
}
func MultipleDataStoreObjectStreamFrom(arg []MultipleDataStoreObject) MultipleDataStoreObjectStream {
	return arg
}
func CreateMultipleDataStoreObjectStream(arg ...MultipleDataStoreObject) *MultipleDataStoreObjectStream {
	tmp := MultipleDataStoreObjectStreamOf(arg...)
	return &tmp
}
func GenerateMultipleDataStoreObjectStream(arg []MultipleDataStoreObject) *MultipleDataStoreObjectStream {
	tmp := MultipleDataStoreObjectStreamFrom(arg)
	return &tmp
}

func (self *MultipleDataStoreObjectStream) Add(arg MultipleDataStoreObject) *MultipleDataStoreObjectStream {
	return self.AddAll(arg)
}
func (self *MultipleDataStoreObjectStream) AddAll(arg ...MultipleDataStoreObject) *MultipleDataStoreObjectStream {
	*self = append(*self, arg...)
	return self
}
func (self *MultipleDataStoreObjectStream) AddSafe(arg *MultipleDataStoreObject) *MultipleDataStoreObjectStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *MultipleDataStoreObjectStream) Aggregate(fn func(MultipleDataStoreObject, MultipleDataStoreObject) MultipleDataStoreObject) *MultipleDataStoreObjectStream {
	result := MultipleDataStoreObjectStreamOf()
	self.ForEach(func(v MultipleDataStoreObject, i int) {
		if i == 0 {
			result.Add(fn(MultipleDataStoreObject{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *MultipleDataStoreObjectStream) AllMatch(fn func(MultipleDataStoreObject, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *MultipleDataStoreObjectStream) AnyMatch(fn func(MultipleDataStoreObject, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *MultipleDataStoreObjectStream) Clone() *MultipleDataStoreObjectStream {
	temp := make([]MultipleDataStoreObject, self.Len())
	copy(temp, *self)
	return (*MultipleDataStoreObjectStream)(&temp)
}
func (self *MultipleDataStoreObjectStream) Copy() *MultipleDataStoreObjectStream {
	return self.Clone()
}
func (self *MultipleDataStoreObjectStream) Concat(arg []MultipleDataStoreObject) *MultipleDataStoreObjectStream {
	return self.AddAll(arg...)
}
func (self *MultipleDataStoreObjectStream) Contains(arg MultipleDataStoreObject) bool {
	return self.FindIndex(func(_arg MultipleDataStoreObject, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *MultipleDataStoreObjectStream) Clean() *MultipleDataStoreObjectStream {
	*self = MultipleDataStoreObjectStreamOf()
	return self
}
func (self *MultipleDataStoreObjectStream) Delete(index int) *MultipleDataStoreObjectStream {
	return self.DeleteRange(index, index)
}
func (self *MultipleDataStoreObjectStream) DeleteRange(startIndex, endIndex int) *MultipleDataStoreObjectStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *MultipleDataStoreObjectStream) Distinct() *MultipleDataStoreObjectStream {
	caches := map[string]bool{}
	result := MultipleDataStoreObjectStreamOf()
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
func (self *MultipleDataStoreObjectStream) Each(fn func(MultipleDataStoreObject)) *MultipleDataStoreObjectStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *MultipleDataStoreObjectStream) EachRight(fn func(MultipleDataStoreObject)) *MultipleDataStoreObjectStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *MultipleDataStoreObjectStream) Equals(arr []MultipleDataStoreObject) bool {
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
func (self *MultipleDataStoreObjectStream) Filter(fn func(MultipleDataStoreObject, int) bool) *MultipleDataStoreObjectStream {
	result := MultipleDataStoreObjectStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleDataStoreObjectStream) FilterSlim(fn func(MultipleDataStoreObject, int) bool) *MultipleDataStoreObjectStream {
	result := MultipleDataStoreObjectStreamOf()
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
func (self *MultipleDataStoreObjectStream) Find(fn func(MultipleDataStoreObject, int) bool) *MultipleDataStoreObject {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *MultipleDataStoreObjectStream) FindOr(fn func(MultipleDataStoreObject, int) bool, or MultipleDataStoreObject) MultipleDataStoreObject {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *MultipleDataStoreObjectStream) FindIndex(fn func(MultipleDataStoreObject, int) bool) int {
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
func (self *MultipleDataStoreObjectStream) First() *MultipleDataStoreObject {
	return self.Get(0)
}
func (self *MultipleDataStoreObjectStream) FirstOr(arg MultipleDataStoreObject) MultipleDataStoreObject {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleDataStoreObjectStream) ForEach(fn func(MultipleDataStoreObject, int)) *MultipleDataStoreObjectStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *MultipleDataStoreObjectStream) ForEachRight(fn func(MultipleDataStoreObject, int)) *MultipleDataStoreObjectStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *MultipleDataStoreObjectStream) GroupBy(fn func(MultipleDataStoreObject, int) string) map[string][]MultipleDataStoreObject {
	m := map[string][]MultipleDataStoreObject{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *MultipleDataStoreObjectStream) GroupByValues(fn func(MultipleDataStoreObject, int) string) [][]MultipleDataStoreObject {
	var tmp [][]MultipleDataStoreObject
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *MultipleDataStoreObjectStream) IndexOf(arg MultipleDataStoreObject) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *MultipleDataStoreObjectStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *MultipleDataStoreObjectStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *MultipleDataStoreObjectStream) Last() *MultipleDataStoreObject {
	return self.Get(self.Len() - 1)
}
func (self *MultipleDataStoreObjectStream) LastOr(arg MultipleDataStoreObject) MultipleDataStoreObject {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleDataStoreObjectStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *MultipleDataStoreObjectStream) Limit(limit int) *MultipleDataStoreObjectStream {
	self.Slice(0, limit)
	return self
}

func (self *MultipleDataStoreObjectStream) Map(fn func(MultipleDataStoreObject, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreObjectStream) Map2Int(fn func(MultipleDataStoreObject, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreObjectStream) Map2Int32(fn func(MultipleDataStoreObject, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreObjectStream) Map2Int64(fn func(MultipleDataStoreObject, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreObjectStream) Map2Float32(fn func(MultipleDataStoreObject, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreObjectStream) Map2Float64(fn func(MultipleDataStoreObject, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreObjectStream) Map2Bool(fn func(MultipleDataStoreObject, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreObjectStream) Map2Bytes(fn func(MultipleDataStoreObject, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreObjectStream) Map2String(fn func(MultipleDataStoreObject, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDataStoreObjectStream) Max(fn func(MultipleDataStoreObject, int) float64) *MultipleDataStoreObject {
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
func (self *MultipleDataStoreObjectStream) Min(fn func(MultipleDataStoreObject, int) float64) *MultipleDataStoreObject {
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
func (self *MultipleDataStoreObjectStream) NoneMatch(fn func(MultipleDataStoreObject, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *MultipleDataStoreObjectStream) Get(index int) *MultipleDataStoreObject {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *MultipleDataStoreObjectStream) GetOr(index int, arg MultipleDataStoreObject) MultipleDataStoreObject {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleDataStoreObjectStream) Peek(fn func(*MultipleDataStoreObject, int)) *MultipleDataStoreObjectStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *MultipleDataStoreObjectStream) Reduce(fn func(MultipleDataStoreObject, MultipleDataStoreObject, int) MultipleDataStoreObject) *MultipleDataStoreObjectStream {
	return self.ReduceInit(fn, MultipleDataStoreObject{})
}
func (self *MultipleDataStoreObjectStream) ReduceInit(fn func(MultipleDataStoreObject, MultipleDataStoreObject, int) MultipleDataStoreObject, initialValue MultipleDataStoreObject) *MultipleDataStoreObjectStream {
	result := MultipleDataStoreObjectStreamOf()
	self.ForEach(func(v MultipleDataStoreObject, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *MultipleDataStoreObjectStream) ReduceInterface(fn func(interface{}, MultipleDataStoreObject, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(MultipleDataStoreObject{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *MultipleDataStoreObjectStream) ReduceString(fn func(string, MultipleDataStoreObject, int) string) []string {
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
func (self *MultipleDataStoreObjectStream) ReduceInt(fn func(int, MultipleDataStoreObject, int) int) []int {
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
func (self *MultipleDataStoreObjectStream) ReduceInt32(fn func(int32, MultipleDataStoreObject, int) int32) []int32 {
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
func (self *MultipleDataStoreObjectStream) ReduceInt64(fn func(int64, MultipleDataStoreObject, int) int64) []int64 {
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
func (self *MultipleDataStoreObjectStream) ReduceFloat32(fn func(float32, MultipleDataStoreObject, int) float32) []float32 {
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
func (self *MultipleDataStoreObjectStream) ReduceFloat64(fn func(float64, MultipleDataStoreObject, int) float64) []float64 {
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
func (self *MultipleDataStoreObjectStream) ReduceBool(fn func(bool, MultipleDataStoreObject, int) bool) []bool {
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
func (self *MultipleDataStoreObjectStream) Reverse() *MultipleDataStoreObjectStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *MultipleDataStoreObjectStream) Replace(fn func(MultipleDataStoreObject, int) MultipleDataStoreObject) *MultipleDataStoreObjectStream {
	return self.ForEach(func(v MultipleDataStoreObject, i int) { self.Set(i, fn(v, i)) })
}
func (self *MultipleDataStoreObjectStream) Select(fn func(MultipleDataStoreObject) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *MultipleDataStoreObjectStream) Set(index int, val MultipleDataStoreObject) *MultipleDataStoreObjectStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *MultipleDataStoreObjectStream) Skip(skip int) *MultipleDataStoreObjectStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *MultipleDataStoreObjectStream) SkippingEach(fn func(MultipleDataStoreObject, int) int) *MultipleDataStoreObjectStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *MultipleDataStoreObjectStream) Slice(startIndex, n int) *MultipleDataStoreObjectStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []MultipleDataStoreObject{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *MultipleDataStoreObjectStream) Sort(fn func(i, j int) bool) *MultipleDataStoreObjectStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *MultipleDataStoreObjectStream) Tail() *MultipleDataStoreObject {
	return self.Last()
}
func (self *MultipleDataStoreObjectStream) TailOr(arg MultipleDataStoreObject) MultipleDataStoreObject {
	return self.LastOr(arg)
}
func (self *MultipleDataStoreObjectStream) ToList() []MultipleDataStoreObject {
	return self.Val()
}
func (self *MultipleDataStoreObjectStream) Unique() *MultipleDataStoreObjectStream {
	return self.Distinct()
}
func (self *MultipleDataStoreObjectStream) Val() []MultipleDataStoreObject {
	if self == nil {
		return []MultipleDataStoreObject{}
	}
	return *self.Copy()
}
func (self *MultipleDataStoreObjectStream) While(fn func(MultipleDataStoreObject, int) bool) *MultipleDataStoreObjectStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *MultipleDataStoreObjectStream) Where(fn func(MultipleDataStoreObject) bool) *MultipleDataStoreObjectStream {
	result := MultipleDataStoreObjectStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleDataStoreObjectStream) WhereSlim(fn func(MultipleDataStoreObject) bool) *MultipleDataStoreObjectStream {
	result := MultipleDataStoreObjectStreamOf()
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
