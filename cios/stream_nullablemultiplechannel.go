/*
 * Collection utility of NullableMultipleChannel Struct
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

type NullableMultipleChannelStream []NullableMultipleChannel

func NullableMultipleChannelStreamOf(arg ...NullableMultipleChannel) NullableMultipleChannelStream {
	return arg
}
func NullableMultipleChannelStreamFrom(arg []NullableMultipleChannel) NullableMultipleChannelStream {
	return arg
}
func CreateNullableMultipleChannelStream(arg ...NullableMultipleChannel) *NullableMultipleChannelStream {
	tmp := NullableMultipleChannelStreamOf(arg...)
	return &tmp
}
func GenerateNullableMultipleChannelStream(arg []NullableMultipleChannel) *NullableMultipleChannelStream {
	tmp := NullableMultipleChannelStreamFrom(arg)
	return &tmp
}

func (self *NullableMultipleChannelStream) Add(arg NullableMultipleChannel) *NullableMultipleChannelStream {
	return self.AddAll(arg)
}
func (self *NullableMultipleChannelStream) AddAll(arg ...NullableMultipleChannel) *NullableMultipleChannelStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableMultipleChannelStream) AddSafe(arg *NullableMultipleChannel) *NullableMultipleChannelStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableMultipleChannelStream) Aggregate(fn func(NullableMultipleChannel, NullableMultipleChannel) NullableMultipleChannel) *NullableMultipleChannelStream {
	result := NullableMultipleChannelStreamOf()
	self.ForEach(func(v NullableMultipleChannel, i int) {
		if i == 0 {
			result.Add(fn(NullableMultipleChannel{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleChannelStream) AllMatch(fn func(NullableMultipleChannel, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableMultipleChannelStream) AnyMatch(fn func(NullableMultipleChannel, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableMultipleChannelStream) Clone() *NullableMultipleChannelStream {
	temp := make([]NullableMultipleChannel, self.Len())
	copy(temp, *self)
	return (*NullableMultipleChannelStream)(&temp)
}
func (self *NullableMultipleChannelStream) Copy() *NullableMultipleChannelStream {
	return self.Clone()
}
func (self *NullableMultipleChannelStream) Concat(arg []NullableMultipleChannel) *NullableMultipleChannelStream {
	return self.AddAll(arg...)
}
func (self *NullableMultipleChannelStream) Contains(arg NullableMultipleChannel) bool {
	return self.FindIndex(func(_arg NullableMultipleChannel, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableMultipleChannelStream) Clean() *NullableMultipleChannelStream {
	*self = NullableMultipleChannelStreamOf()
	return self
}
func (self *NullableMultipleChannelStream) Delete(index int) *NullableMultipleChannelStream {
	return self.DeleteRange(index, index)
}
func (self *NullableMultipleChannelStream) DeleteRange(startIndex, endIndex int) *NullableMultipleChannelStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableMultipleChannelStream) Distinct() *NullableMultipleChannelStream {
	caches := map[string]bool{}
	result := NullableMultipleChannelStreamOf()
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
func (self *NullableMultipleChannelStream) Each(fn func(NullableMultipleChannel)) *NullableMultipleChannelStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableMultipleChannelStream) EachRight(fn func(NullableMultipleChannel)) *NullableMultipleChannelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableMultipleChannelStream) Equals(arr []NullableMultipleChannel) bool {
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
func (self *NullableMultipleChannelStream) Filter(fn func(NullableMultipleChannel, int) bool) *NullableMultipleChannelStream {
	result := NullableMultipleChannelStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleChannelStream) FilterSlim(fn func(NullableMultipleChannel, int) bool) *NullableMultipleChannelStream {
	result := NullableMultipleChannelStreamOf()
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
func (self *NullableMultipleChannelStream) Find(fn func(NullableMultipleChannel, int) bool) *NullableMultipleChannel {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleChannelStream) FindOr(fn func(NullableMultipleChannel, int) bool, or NullableMultipleChannel) NullableMultipleChannel {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableMultipleChannelStream) FindIndex(fn func(NullableMultipleChannel, int) bool) int {
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
func (self *NullableMultipleChannelStream) First() *NullableMultipleChannel {
	return self.Get(0)
}
func (self *NullableMultipleChannelStream) FirstOr(arg NullableMultipleChannel) NullableMultipleChannel {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleChannelStream) ForEach(fn func(NullableMultipleChannel, int)) *NullableMultipleChannelStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableMultipleChannelStream) ForEachRight(fn func(NullableMultipleChannel, int)) *NullableMultipleChannelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableMultipleChannelStream) GroupBy(fn func(NullableMultipleChannel, int) string) map[string][]NullableMultipleChannel {
	m := map[string][]NullableMultipleChannel{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableMultipleChannelStream) GroupByValues(fn func(NullableMultipleChannel, int) string) [][]NullableMultipleChannel {
	var tmp [][]NullableMultipleChannel
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableMultipleChannelStream) IndexOf(arg NullableMultipleChannel) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableMultipleChannelStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableMultipleChannelStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableMultipleChannelStream) Last() *NullableMultipleChannel {
	return self.Get(self.Len() - 1)
}
func (self *NullableMultipleChannelStream) LastOr(arg NullableMultipleChannel) NullableMultipleChannel {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleChannelStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableMultipleChannelStream) Limit(limit int) *NullableMultipleChannelStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableMultipleChannelStream) Map(fn func(NullableMultipleChannel, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleChannelStream) Map2Int(fn func(NullableMultipleChannel, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleChannelStream) Map2Int32(fn func(NullableMultipleChannel, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleChannelStream) Map2Int64(fn func(NullableMultipleChannel, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleChannelStream) Map2Float32(fn func(NullableMultipleChannel, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleChannelStream) Map2Float64(fn func(NullableMultipleChannel, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleChannelStream) Map2Bool(fn func(NullableMultipleChannel, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleChannelStream) Map2Bytes(fn func(NullableMultipleChannel, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleChannelStream) Map2String(fn func(NullableMultipleChannel, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleChannelStream) Max(fn func(NullableMultipleChannel, int) float64) *NullableMultipleChannel {
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
func (self *NullableMultipleChannelStream) Min(fn func(NullableMultipleChannel, int) float64) *NullableMultipleChannel {
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
func (self *NullableMultipleChannelStream) NoneMatch(fn func(NullableMultipleChannel, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableMultipleChannelStream) Get(index int) *NullableMultipleChannel {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleChannelStream) GetOr(index int, arg NullableMultipleChannel) NullableMultipleChannel {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleChannelStream) Peek(fn func(*NullableMultipleChannel, int)) *NullableMultipleChannelStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableMultipleChannelStream) Reduce(fn func(NullableMultipleChannel, NullableMultipleChannel, int) NullableMultipleChannel) *NullableMultipleChannelStream {
	return self.ReduceInit(fn, NullableMultipleChannel{})
}
func (self *NullableMultipleChannelStream) ReduceInit(fn func(NullableMultipleChannel, NullableMultipleChannel, int) NullableMultipleChannel, initialValue NullableMultipleChannel) *NullableMultipleChannelStream {
	result := NullableMultipleChannelStreamOf()
	self.ForEach(func(v NullableMultipleChannel, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleChannelStream) ReduceInterface(fn func(interface{}, NullableMultipleChannel, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableMultipleChannel{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleChannelStream) ReduceString(fn func(string, NullableMultipleChannel, int) string) []string {
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
func (self *NullableMultipleChannelStream) ReduceInt(fn func(int, NullableMultipleChannel, int) int) []int {
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
func (self *NullableMultipleChannelStream) ReduceInt32(fn func(int32, NullableMultipleChannel, int) int32) []int32 {
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
func (self *NullableMultipleChannelStream) ReduceInt64(fn func(int64, NullableMultipleChannel, int) int64) []int64 {
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
func (self *NullableMultipleChannelStream) ReduceFloat32(fn func(float32, NullableMultipleChannel, int) float32) []float32 {
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
func (self *NullableMultipleChannelStream) ReduceFloat64(fn func(float64, NullableMultipleChannel, int) float64) []float64 {
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
func (self *NullableMultipleChannelStream) ReduceBool(fn func(bool, NullableMultipleChannel, int) bool) []bool {
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
func (self *NullableMultipleChannelStream) Reverse() *NullableMultipleChannelStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableMultipleChannelStream) Replace(fn func(NullableMultipleChannel, int) NullableMultipleChannel) *NullableMultipleChannelStream {
	return self.ForEach(func(v NullableMultipleChannel, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableMultipleChannelStream) Select(fn func(NullableMultipleChannel) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableMultipleChannelStream) Set(index int, val NullableMultipleChannel) *NullableMultipleChannelStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableMultipleChannelStream) Skip(skip int) *NullableMultipleChannelStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableMultipleChannelStream) SkippingEach(fn func(NullableMultipleChannel, int) int) *NullableMultipleChannelStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableMultipleChannelStream) Slice(startIndex, n int) *NullableMultipleChannelStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableMultipleChannel{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableMultipleChannelStream) Sort(fn func(i, j int) bool) *NullableMultipleChannelStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableMultipleChannelStream) Tail() *NullableMultipleChannel {
	return self.Last()
}
func (self *NullableMultipleChannelStream) TailOr(arg NullableMultipleChannel) NullableMultipleChannel {
	return self.LastOr(arg)
}
func (self *NullableMultipleChannelStream) ToList() []NullableMultipleChannel {
	return self.Val()
}
func (self *NullableMultipleChannelStream) Unique() *NullableMultipleChannelStream {
	return self.Distinct()
}
func (self *NullableMultipleChannelStream) Val() []NullableMultipleChannel {
	if self == nil {
		return []NullableMultipleChannel{}
	}
	return *self.Copy()
}
func (self *NullableMultipleChannelStream) While(fn func(NullableMultipleChannel, int) bool) *NullableMultipleChannelStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableMultipleChannelStream) Where(fn func(NullableMultipleChannel) bool) *NullableMultipleChannelStream {
	result := NullableMultipleChannelStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleChannelStream) WhereSlim(fn func(NullableMultipleChannel) bool) *NullableMultipleChannelStream {
	result := NullableMultipleChannelStreamOf()
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
