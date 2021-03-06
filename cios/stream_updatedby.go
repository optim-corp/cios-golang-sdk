/*
 * Collection utility of UpdatedBy Struct
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

type UpdatedByStream []UpdatedBy

func UpdatedByStreamOf(arg ...UpdatedBy) UpdatedByStream {
	return arg
}
func UpdatedByStreamFrom(arg []UpdatedBy) UpdatedByStream {
	return arg
}
func CreateUpdatedByStream(arg ...UpdatedBy) *UpdatedByStream {
	tmp := UpdatedByStreamOf(arg...)
	return &tmp
}
func GenerateUpdatedByStream(arg []UpdatedBy) *UpdatedByStream {
	tmp := UpdatedByStreamFrom(arg)
	return &tmp
}

func (self *UpdatedByStream) Add(arg UpdatedBy) *UpdatedByStream {
	return self.AddAll(arg)
}
func (self *UpdatedByStream) AddAll(arg ...UpdatedBy) *UpdatedByStream {
	*self = append(*self, arg...)
	return self
}
func (self *UpdatedByStream) AddSafe(arg *UpdatedBy) *UpdatedByStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *UpdatedByStream) Aggregate(fn func(UpdatedBy, UpdatedBy) UpdatedBy) *UpdatedByStream {
	result := UpdatedByStreamOf()
	self.ForEach(func(v UpdatedBy, i int) {
		if i == 0 {
			result.Add(fn(UpdatedBy{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *UpdatedByStream) AllMatch(fn func(UpdatedBy, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *UpdatedByStream) AnyMatch(fn func(UpdatedBy, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *UpdatedByStream) Clone() *UpdatedByStream {
	temp := make([]UpdatedBy, self.Len())
	copy(temp, *self)
	return (*UpdatedByStream)(&temp)
}
func (self *UpdatedByStream) Copy() *UpdatedByStream {
	return self.Clone()
}
func (self *UpdatedByStream) Concat(arg []UpdatedBy) *UpdatedByStream {
	return self.AddAll(arg...)
}
func (self *UpdatedByStream) Contains(arg UpdatedBy) bool {
	return self.FindIndex(func(_arg UpdatedBy, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *UpdatedByStream) Clean() *UpdatedByStream {
	*self = UpdatedByStreamOf()
	return self
}
func (self *UpdatedByStream) Delete(index int) *UpdatedByStream {
	return self.DeleteRange(index, index)
}
func (self *UpdatedByStream) DeleteRange(startIndex, endIndex int) *UpdatedByStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *UpdatedByStream) Distinct() *UpdatedByStream {
	caches := map[string]bool{}
	result := UpdatedByStreamOf()
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
func (self *UpdatedByStream) Each(fn func(UpdatedBy)) *UpdatedByStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *UpdatedByStream) EachRight(fn func(UpdatedBy)) *UpdatedByStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *UpdatedByStream) Equals(arr []UpdatedBy) bool {
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
func (self *UpdatedByStream) Filter(fn func(UpdatedBy, int) bool) *UpdatedByStream {
	result := UpdatedByStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *UpdatedByStream) FilterSlim(fn func(UpdatedBy, int) bool) *UpdatedByStream {
	result := UpdatedByStreamOf()
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
func (self *UpdatedByStream) Find(fn func(UpdatedBy, int) bool) *UpdatedBy {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *UpdatedByStream) FindOr(fn func(UpdatedBy, int) bool, or UpdatedBy) UpdatedBy {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *UpdatedByStream) FindIndex(fn func(UpdatedBy, int) bool) int {
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
func (self *UpdatedByStream) First() *UpdatedBy {
	return self.Get(0)
}
func (self *UpdatedByStream) FirstOr(arg UpdatedBy) UpdatedBy {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *UpdatedByStream) ForEach(fn func(UpdatedBy, int)) *UpdatedByStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *UpdatedByStream) ForEachRight(fn func(UpdatedBy, int)) *UpdatedByStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *UpdatedByStream) GroupBy(fn func(UpdatedBy, int) string) map[string][]UpdatedBy {
	m := map[string][]UpdatedBy{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *UpdatedByStream) GroupByValues(fn func(UpdatedBy, int) string) [][]UpdatedBy {
	var tmp [][]UpdatedBy
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *UpdatedByStream) IndexOf(arg UpdatedBy) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *UpdatedByStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *UpdatedByStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *UpdatedByStream) Last() *UpdatedBy {
	return self.Get(self.Len() - 1)
}
func (self *UpdatedByStream) LastOr(arg UpdatedBy) UpdatedBy {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *UpdatedByStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *UpdatedByStream) Limit(limit int) *UpdatedByStream {
	self.Slice(0, limit)
	return self
}

func (self *UpdatedByStream) Map(fn func(UpdatedBy, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *UpdatedByStream) Map2Int(fn func(UpdatedBy, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *UpdatedByStream) Map2Int32(fn func(UpdatedBy, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *UpdatedByStream) Map2Int64(fn func(UpdatedBy, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *UpdatedByStream) Map2Float32(fn func(UpdatedBy, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *UpdatedByStream) Map2Float64(fn func(UpdatedBy, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *UpdatedByStream) Map2Bool(fn func(UpdatedBy, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *UpdatedByStream) Map2Bytes(fn func(UpdatedBy, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *UpdatedByStream) Map2String(fn func(UpdatedBy, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *UpdatedByStream) Max(fn func(UpdatedBy, int) float64) *UpdatedBy {
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
func (self *UpdatedByStream) Min(fn func(UpdatedBy, int) float64) *UpdatedBy {
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
func (self *UpdatedByStream) NoneMatch(fn func(UpdatedBy, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *UpdatedByStream) Get(index int) *UpdatedBy {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *UpdatedByStream) GetOr(index int, arg UpdatedBy) UpdatedBy {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *UpdatedByStream) Peek(fn func(*UpdatedBy, int)) *UpdatedByStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *UpdatedByStream) Reduce(fn func(UpdatedBy, UpdatedBy, int) UpdatedBy) *UpdatedByStream {
	return self.ReduceInit(fn, UpdatedBy{})
}
func (self *UpdatedByStream) ReduceInit(fn func(UpdatedBy, UpdatedBy, int) UpdatedBy, initialValue UpdatedBy) *UpdatedByStream {
	result := UpdatedByStreamOf()
	self.ForEach(func(v UpdatedBy, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *UpdatedByStream) ReduceInterface(fn func(interface{}, UpdatedBy, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(UpdatedBy{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *UpdatedByStream) ReduceString(fn func(string, UpdatedBy, int) string) []string {
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
func (self *UpdatedByStream) ReduceInt(fn func(int, UpdatedBy, int) int) []int {
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
func (self *UpdatedByStream) ReduceInt32(fn func(int32, UpdatedBy, int) int32) []int32 {
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
func (self *UpdatedByStream) ReduceInt64(fn func(int64, UpdatedBy, int) int64) []int64 {
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
func (self *UpdatedByStream) ReduceFloat32(fn func(float32, UpdatedBy, int) float32) []float32 {
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
func (self *UpdatedByStream) ReduceFloat64(fn func(float64, UpdatedBy, int) float64) []float64 {
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
func (self *UpdatedByStream) ReduceBool(fn func(bool, UpdatedBy, int) bool) []bool {
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
func (self *UpdatedByStream) Reverse() *UpdatedByStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *UpdatedByStream) Replace(fn func(UpdatedBy, int) UpdatedBy) *UpdatedByStream {
	return self.ForEach(func(v UpdatedBy, i int) { self.Set(i, fn(v, i)) })
}
func (self *UpdatedByStream) Select(fn func(UpdatedBy) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *UpdatedByStream) Set(index int, val UpdatedBy) *UpdatedByStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *UpdatedByStream) Skip(skip int) *UpdatedByStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *UpdatedByStream) SkippingEach(fn func(UpdatedBy, int) int) *UpdatedByStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *UpdatedByStream) Slice(startIndex, n int) *UpdatedByStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []UpdatedBy{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *UpdatedByStream) Sort(fn func(i, j int) bool) *UpdatedByStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *UpdatedByStream) Tail() *UpdatedBy {
	return self.Last()
}
func (self *UpdatedByStream) TailOr(arg UpdatedBy) UpdatedBy {
	return self.LastOr(arg)
}
func (self *UpdatedByStream) ToList() []UpdatedBy {
	return self.Val()
}
func (self *UpdatedByStream) Unique() *UpdatedByStream {
	return self.Distinct()
}
func (self *UpdatedByStream) Val() []UpdatedBy {
	if self == nil {
		return []UpdatedBy{}
	}
	return *self.Copy()
}
func (self *UpdatedByStream) While(fn func(UpdatedBy, int) bool) *UpdatedByStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *UpdatedByStream) Where(fn func(UpdatedBy) bool) *UpdatedByStream {
	result := UpdatedByStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *UpdatedByStream) WhereSlim(fn func(UpdatedBy) bool) *UpdatedByStream {
	result := UpdatedByStreamOf()
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
