/*
 * Collection utility of NullableTime Struct
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

type NullableTimeStream []NullableTime

func NullableTimeStreamOf(arg ...NullableTime) NullableTimeStream {
	return arg
}
func NullableTimeStreamFrom(arg []NullableTime) NullableTimeStream {
	return arg
}
func CreateNullableTimeStream(arg ...NullableTime) *NullableTimeStream {
	tmp := NullableTimeStreamOf(arg...)
	return &tmp
}
func GenerateNullableTimeStream(arg []NullableTime) *NullableTimeStream {
	tmp := NullableTimeStreamFrom(arg)
	return &tmp
}

func (self *NullableTimeStream) Add(arg NullableTime) *NullableTimeStream {
	return self.AddAll(arg)
}
func (self *NullableTimeStream) AddAll(arg ...NullableTime) *NullableTimeStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableTimeStream) AddSafe(arg *NullableTime) *NullableTimeStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableTimeStream) Aggregate(fn func(NullableTime, NullableTime) NullableTime) *NullableTimeStream {
	result := NullableTimeStreamOf()
	self.ForEach(func(v NullableTime, i int) {
		if i == 0 {
			result.Add(fn(NullableTime{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableTimeStream) AllMatch(fn func(NullableTime, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableTimeStream) AnyMatch(fn func(NullableTime, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableTimeStream) Clone() *NullableTimeStream {
	temp := make([]NullableTime, self.Len())
	copy(temp, *self)
	return (*NullableTimeStream)(&temp)
}
func (self *NullableTimeStream) Copy() *NullableTimeStream {
	return self.Clone()
}
func (self *NullableTimeStream) Concat(arg []NullableTime) *NullableTimeStream {
	return self.AddAll(arg...)
}
func (self *NullableTimeStream) Contains(arg NullableTime) bool {
	return self.FindIndex(func(_arg NullableTime, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableTimeStream) Clean() *NullableTimeStream {
	*self = NullableTimeStreamOf()
	return self
}
func (self *NullableTimeStream) Delete(index int) *NullableTimeStream {
	return self.DeleteRange(index, index)
}
func (self *NullableTimeStream) DeleteRange(startIndex, endIndex int) *NullableTimeStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableTimeStream) Distinct() *NullableTimeStream {
	caches := map[string]bool{}
	result := NullableTimeStreamOf()
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
func (self *NullableTimeStream) Each(fn func(NullableTime)) *NullableTimeStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableTimeStream) EachRight(fn func(NullableTime)) *NullableTimeStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableTimeStream) Equals(arr []NullableTime) bool {
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
func (self *NullableTimeStream) Filter(fn func(NullableTime, int) bool) *NullableTimeStream {
	result := NullableTimeStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableTimeStream) FilterSlim(fn func(NullableTime, int) bool) *NullableTimeStream {
	result := NullableTimeStreamOf()
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
func (self *NullableTimeStream) Find(fn func(NullableTime, int) bool) *NullableTime {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableTimeStream) FindOr(fn func(NullableTime, int) bool, or NullableTime) NullableTime {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableTimeStream) FindIndex(fn func(NullableTime, int) bool) int {
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
func (self *NullableTimeStream) First() *NullableTime {
	return self.Get(0)
}
func (self *NullableTimeStream) FirstOr(arg NullableTime) NullableTime {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableTimeStream) ForEach(fn func(NullableTime, int)) *NullableTimeStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableTimeStream) ForEachRight(fn func(NullableTime, int)) *NullableTimeStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableTimeStream) GroupBy(fn func(NullableTime, int) string) map[string][]NullableTime {
	m := map[string][]NullableTime{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableTimeStream) GroupByValues(fn func(NullableTime, int) string) [][]NullableTime {
	var tmp [][]NullableTime
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableTimeStream) IndexOf(arg NullableTime) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableTimeStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableTimeStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableTimeStream) Last() *NullableTime {
	return self.Get(self.Len() - 1)
}
func (self *NullableTimeStream) LastOr(arg NullableTime) NullableTime {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableTimeStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableTimeStream) Limit(limit int) *NullableTimeStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableTimeStream) Map(fn func(NullableTime, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableTimeStream) Map2Int(fn func(NullableTime, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableTimeStream) Map2Int32(fn func(NullableTime, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableTimeStream) Map2Int64(fn func(NullableTime, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableTimeStream) Map2Float32(fn func(NullableTime, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableTimeStream) Map2Float64(fn func(NullableTime, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableTimeStream) Map2Bool(fn func(NullableTime, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableTimeStream) Map2Bytes(fn func(NullableTime, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableTimeStream) Map2String(fn func(NullableTime, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableTimeStream) Max(fn func(NullableTime, int) float64) *NullableTime {
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
func (self *NullableTimeStream) Min(fn func(NullableTime, int) float64) *NullableTime {
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
func (self *NullableTimeStream) NoneMatch(fn func(NullableTime, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableTimeStream) Get(index int) *NullableTime {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableTimeStream) GetOr(index int, arg NullableTime) NullableTime {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableTimeStream) Peek(fn func(*NullableTime, int)) *NullableTimeStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableTimeStream) Reduce(fn func(NullableTime, NullableTime, int) NullableTime) *NullableTimeStream {
	return self.ReduceInit(fn, NullableTime{})
}
func (self *NullableTimeStream) ReduceInit(fn func(NullableTime, NullableTime, int) NullableTime, initialValue NullableTime) *NullableTimeStream {
	result := NullableTimeStreamOf()
	self.ForEach(func(v NullableTime, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableTimeStream) ReduceInterface(fn func(interface{}, NullableTime, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableTime{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableTimeStream) ReduceString(fn func(string, NullableTime, int) string) []string {
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
func (self *NullableTimeStream) ReduceInt(fn func(int, NullableTime, int) int) []int {
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
func (self *NullableTimeStream) ReduceInt32(fn func(int32, NullableTime, int) int32) []int32 {
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
func (self *NullableTimeStream) ReduceInt64(fn func(int64, NullableTime, int) int64) []int64 {
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
func (self *NullableTimeStream) ReduceFloat32(fn func(float32, NullableTime, int) float32) []float32 {
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
func (self *NullableTimeStream) ReduceFloat64(fn func(float64, NullableTime, int) float64) []float64 {
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
func (self *NullableTimeStream) ReduceBool(fn func(bool, NullableTime, int) bool) []bool {
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
func (self *NullableTimeStream) Reverse() *NullableTimeStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableTimeStream) Replace(fn func(NullableTime, int) NullableTime) *NullableTimeStream {
	return self.ForEach(func(v NullableTime, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableTimeStream) Select(fn func(NullableTime) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableTimeStream) Set(index int, val NullableTime) *NullableTimeStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableTimeStream) Skip(skip int) *NullableTimeStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableTimeStream) SkippingEach(fn func(NullableTime, int) int) *NullableTimeStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableTimeStream) Slice(startIndex, n int) *NullableTimeStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableTime{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableTimeStream) Sort(fn func(i, j int) bool) *NullableTimeStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableTimeStream) Tail() *NullableTime {
	return self.Last()
}
func (self *NullableTimeStream) TailOr(arg NullableTime) NullableTime {
	return self.LastOr(arg)
}
func (self *NullableTimeStream) ToList() []NullableTime {
	return self.Val()
}
func (self *NullableTimeStream) Unique() *NullableTimeStream {
	return self.Distinct()
}
func (self *NullableTimeStream) Val() []NullableTime {
	if self == nil {
		return []NullableTime{}
	}
	return *self.Copy()
}
func (self *NullableTimeStream) While(fn func(NullableTime, int) bool) *NullableTimeStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableTimeStream) Where(fn func(NullableTime) bool) *NullableTimeStream {
	result := NullableTimeStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableTimeStream) WhereSlim(fn func(NullableTime) bool) *NullableTimeStream {
	result := NullableTimeStreamOf()
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