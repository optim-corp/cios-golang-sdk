/*
 * Collection utility of CreatedByUser Struct
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

type CreatedByUserStream []CreatedByUser

func CreatedByUserStreamOf(arg ...CreatedByUser) CreatedByUserStream {
	return arg
}
func CreatedByUserStreamFrom(arg []CreatedByUser) CreatedByUserStream {
	return arg
}
func CreateCreatedByUserStream(arg ...CreatedByUser) *CreatedByUserStream {
	tmp := CreatedByUserStreamOf(arg...)
	return &tmp
}
func GenerateCreatedByUserStream(arg []CreatedByUser) *CreatedByUserStream {
	tmp := CreatedByUserStreamFrom(arg)
	return &tmp
}

func (self *CreatedByUserStream) Add(arg CreatedByUser) *CreatedByUserStream {
	return self.AddAll(arg)
}
func (self *CreatedByUserStream) AddAll(arg ...CreatedByUser) *CreatedByUserStream {
	*self = append(*self, arg...)
	return self
}
func (self *CreatedByUserStream) AddSafe(arg *CreatedByUser) *CreatedByUserStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *CreatedByUserStream) Aggregate(fn func(CreatedByUser, CreatedByUser) CreatedByUser) *CreatedByUserStream {
	result := CreatedByUserStreamOf()
	self.ForEach(func(v CreatedByUser, i int) {
		if i == 0 {
			result.Add(fn(CreatedByUser{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *CreatedByUserStream) AllMatch(fn func(CreatedByUser, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *CreatedByUserStream) AnyMatch(fn func(CreatedByUser, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *CreatedByUserStream) Clone() *CreatedByUserStream {
	temp := make([]CreatedByUser, self.Len())
	copy(temp, *self)
	return (*CreatedByUserStream)(&temp)
}
func (self *CreatedByUserStream) Copy() *CreatedByUserStream {
	return self.Clone()
}
func (self *CreatedByUserStream) Concat(arg []CreatedByUser) *CreatedByUserStream {
	return self.AddAll(arg...)
}
func (self *CreatedByUserStream) Contains(arg CreatedByUser) bool {
	return self.FindIndex(func(_arg CreatedByUser, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *CreatedByUserStream) Clean() *CreatedByUserStream {
	*self = CreatedByUserStreamOf()
	return self
}
func (self *CreatedByUserStream) Delete(index int) *CreatedByUserStream {
	return self.DeleteRange(index, index)
}
func (self *CreatedByUserStream) DeleteRange(startIndex, endIndex int) *CreatedByUserStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *CreatedByUserStream) Distinct() *CreatedByUserStream {
	caches := map[string]bool{}
	result := CreatedByUserStreamOf()
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
func (self *CreatedByUserStream) Each(fn func(CreatedByUser)) *CreatedByUserStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *CreatedByUserStream) EachRight(fn func(CreatedByUser)) *CreatedByUserStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *CreatedByUserStream) Equals(arr []CreatedByUser) bool {
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
func (self *CreatedByUserStream) Filter(fn func(CreatedByUser, int) bool) *CreatedByUserStream {
	result := CreatedByUserStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *CreatedByUserStream) FilterSlim(fn func(CreatedByUser, int) bool) *CreatedByUserStream {
	result := CreatedByUserStreamOf()
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
func (self *CreatedByUserStream) Find(fn func(CreatedByUser, int) bool) *CreatedByUser {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *CreatedByUserStream) FindOr(fn func(CreatedByUser, int) bool, or CreatedByUser) CreatedByUser {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *CreatedByUserStream) FindIndex(fn func(CreatedByUser, int) bool) int {
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
func (self *CreatedByUserStream) First() *CreatedByUser {
	return self.Get(0)
}
func (self *CreatedByUserStream) FirstOr(arg CreatedByUser) CreatedByUser {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *CreatedByUserStream) ForEach(fn func(CreatedByUser, int)) *CreatedByUserStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *CreatedByUserStream) ForEachRight(fn func(CreatedByUser, int)) *CreatedByUserStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *CreatedByUserStream) GroupBy(fn func(CreatedByUser, int) string) map[string][]CreatedByUser {
	m := map[string][]CreatedByUser{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *CreatedByUserStream) GroupByValues(fn func(CreatedByUser, int) string) [][]CreatedByUser {
	var tmp [][]CreatedByUser
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *CreatedByUserStream) IndexOf(arg CreatedByUser) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *CreatedByUserStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *CreatedByUserStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *CreatedByUserStream) Last() *CreatedByUser {
	return self.Get(self.Len() - 1)
}
func (self *CreatedByUserStream) LastOr(arg CreatedByUser) CreatedByUser {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *CreatedByUserStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *CreatedByUserStream) Limit(limit int) *CreatedByUserStream {
	self.Slice(0, limit)
	return self
}

func (self *CreatedByUserStream) Map(fn func(CreatedByUser, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CreatedByUserStream) Map2Int(fn func(CreatedByUser, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CreatedByUserStream) Map2Int32(fn func(CreatedByUser, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CreatedByUserStream) Map2Int64(fn func(CreatedByUser, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CreatedByUserStream) Map2Float32(fn func(CreatedByUser, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CreatedByUserStream) Map2Float64(fn func(CreatedByUser, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CreatedByUserStream) Map2Bool(fn func(CreatedByUser, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CreatedByUserStream) Map2Bytes(fn func(CreatedByUser, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CreatedByUserStream) Map2String(fn func(CreatedByUser, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *CreatedByUserStream) Max(fn func(CreatedByUser, int) float64) *CreatedByUser {
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
func (self *CreatedByUserStream) Min(fn func(CreatedByUser, int) float64) *CreatedByUser {
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
func (self *CreatedByUserStream) NoneMatch(fn func(CreatedByUser, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *CreatedByUserStream) Get(index int) *CreatedByUser {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *CreatedByUserStream) GetOr(index int, arg CreatedByUser) CreatedByUser {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *CreatedByUserStream) Peek(fn func(*CreatedByUser, int)) *CreatedByUserStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *CreatedByUserStream) Reduce(fn func(CreatedByUser, CreatedByUser, int) CreatedByUser) *CreatedByUserStream {
	return self.ReduceInit(fn, CreatedByUser{})
}
func (self *CreatedByUserStream) ReduceInit(fn func(CreatedByUser, CreatedByUser, int) CreatedByUser, initialValue CreatedByUser) *CreatedByUserStream {
	result := CreatedByUserStreamOf()
	self.ForEach(func(v CreatedByUser, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *CreatedByUserStream) ReduceInterface(fn func(interface{}, CreatedByUser, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(CreatedByUser{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *CreatedByUserStream) ReduceString(fn func(string, CreatedByUser, int) string) []string {
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
func (self *CreatedByUserStream) ReduceInt(fn func(int, CreatedByUser, int) int) []int {
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
func (self *CreatedByUserStream) ReduceInt32(fn func(int32, CreatedByUser, int) int32) []int32 {
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
func (self *CreatedByUserStream) ReduceInt64(fn func(int64, CreatedByUser, int) int64) []int64 {
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
func (self *CreatedByUserStream) ReduceFloat32(fn func(float32, CreatedByUser, int) float32) []float32 {
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
func (self *CreatedByUserStream) ReduceFloat64(fn func(float64, CreatedByUser, int) float64) []float64 {
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
func (self *CreatedByUserStream) ReduceBool(fn func(bool, CreatedByUser, int) bool) []bool {
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
func (self *CreatedByUserStream) Reverse() *CreatedByUserStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *CreatedByUserStream) Replace(fn func(CreatedByUser, int) CreatedByUser) *CreatedByUserStream {
	return self.ForEach(func(v CreatedByUser, i int) { self.Set(i, fn(v, i)) })
}
func (self *CreatedByUserStream) Select(fn func(CreatedByUser) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *CreatedByUserStream) Set(index int, val CreatedByUser) *CreatedByUserStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *CreatedByUserStream) Skip(skip int) *CreatedByUserStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *CreatedByUserStream) SkippingEach(fn func(CreatedByUser, int) int) *CreatedByUserStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *CreatedByUserStream) Slice(startIndex, n int) *CreatedByUserStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []CreatedByUser{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *CreatedByUserStream) Sort(fn func(i, j int) bool) *CreatedByUserStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *CreatedByUserStream) Tail() *CreatedByUser {
	return self.Last()
}
func (self *CreatedByUserStream) TailOr(arg CreatedByUser) CreatedByUser {
	return self.LastOr(arg)
}
func (self *CreatedByUserStream) ToList() []CreatedByUser {
	return self.Val()
}
func (self *CreatedByUserStream) Unique() *CreatedByUserStream {
	return self.Distinct()
}
func (self *CreatedByUserStream) Val() []CreatedByUser {
	if self == nil {
		return []CreatedByUser{}
	}
	return *self.Copy()
}
func (self *CreatedByUserStream) While(fn func(CreatedByUser, int) bool) *CreatedByUserStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *CreatedByUserStream) Where(fn func(CreatedByUser) bool) *CreatedByUserStream {
	result := CreatedByUserStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *CreatedByUserStream) WhereSlim(fn func(CreatedByUser) bool) *CreatedByUserStream {
	result := CreatedByUserStreamOf()
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
