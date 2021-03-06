/*
 * Collection utility of NullableSubscription Struct
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

type NullableSubscriptionStream []NullableSubscription

func NullableSubscriptionStreamOf(arg ...NullableSubscription) NullableSubscriptionStream {
	return arg
}
func NullableSubscriptionStreamFrom(arg []NullableSubscription) NullableSubscriptionStream {
	return arg
}
func CreateNullableSubscriptionStream(arg ...NullableSubscription) *NullableSubscriptionStream {
	tmp := NullableSubscriptionStreamOf(arg...)
	return &tmp
}
func GenerateNullableSubscriptionStream(arg []NullableSubscription) *NullableSubscriptionStream {
	tmp := NullableSubscriptionStreamFrom(arg)
	return &tmp
}

func (self *NullableSubscriptionStream) Add(arg NullableSubscription) *NullableSubscriptionStream {
	return self.AddAll(arg)
}
func (self *NullableSubscriptionStream) AddAll(arg ...NullableSubscription) *NullableSubscriptionStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableSubscriptionStream) AddSafe(arg *NullableSubscription) *NullableSubscriptionStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableSubscriptionStream) Aggregate(fn func(NullableSubscription, NullableSubscription) NullableSubscription) *NullableSubscriptionStream {
	result := NullableSubscriptionStreamOf()
	self.ForEach(func(v NullableSubscription, i int) {
		if i == 0 {
			result.Add(fn(NullableSubscription{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableSubscriptionStream) AllMatch(fn func(NullableSubscription, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableSubscriptionStream) AnyMatch(fn func(NullableSubscription, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableSubscriptionStream) Clone() *NullableSubscriptionStream {
	temp := make([]NullableSubscription, self.Len())
	copy(temp, *self)
	return (*NullableSubscriptionStream)(&temp)
}
func (self *NullableSubscriptionStream) Copy() *NullableSubscriptionStream {
	return self.Clone()
}
func (self *NullableSubscriptionStream) Concat(arg []NullableSubscription) *NullableSubscriptionStream {
	return self.AddAll(arg...)
}
func (self *NullableSubscriptionStream) Contains(arg NullableSubscription) bool {
	return self.FindIndex(func(_arg NullableSubscription, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableSubscriptionStream) Clean() *NullableSubscriptionStream {
	*self = NullableSubscriptionStreamOf()
	return self
}
func (self *NullableSubscriptionStream) Delete(index int) *NullableSubscriptionStream {
	return self.DeleteRange(index, index)
}
func (self *NullableSubscriptionStream) DeleteRange(startIndex, endIndex int) *NullableSubscriptionStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableSubscriptionStream) Distinct() *NullableSubscriptionStream {
	caches := map[string]bool{}
	result := NullableSubscriptionStreamOf()
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
func (self *NullableSubscriptionStream) Each(fn func(NullableSubscription)) *NullableSubscriptionStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableSubscriptionStream) EachRight(fn func(NullableSubscription)) *NullableSubscriptionStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableSubscriptionStream) Equals(arr []NullableSubscription) bool {
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
func (self *NullableSubscriptionStream) Filter(fn func(NullableSubscription, int) bool) *NullableSubscriptionStream {
	result := NullableSubscriptionStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSubscriptionStream) FilterSlim(fn func(NullableSubscription, int) bool) *NullableSubscriptionStream {
	result := NullableSubscriptionStreamOf()
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
func (self *NullableSubscriptionStream) Find(fn func(NullableSubscription, int) bool) *NullableSubscription {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableSubscriptionStream) FindOr(fn func(NullableSubscription, int) bool, or NullableSubscription) NullableSubscription {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableSubscriptionStream) FindIndex(fn func(NullableSubscription, int) bool) int {
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
func (self *NullableSubscriptionStream) First() *NullableSubscription {
	return self.Get(0)
}
func (self *NullableSubscriptionStream) FirstOr(arg NullableSubscription) NullableSubscription {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSubscriptionStream) ForEach(fn func(NullableSubscription, int)) *NullableSubscriptionStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableSubscriptionStream) ForEachRight(fn func(NullableSubscription, int)) *NullableSubscriptionStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableSubscriptionStream) GroupBy(fn func(NullableSubscription, int) string) map[string][]NullableSubscription {
	m := map[string][]NullableSubscription{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableSubscriptionStream) GroupByValues(fn func(NullableSubscription, int) string) [][]NullableSubscription {
	var tmp [][]NullableSubscription
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableSubscriptionStream) IndexOf(arg NullableSubscription) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableSubscriptionStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableSubscriptionStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableSubscriptionStream) Last() *NullableSubscription {
	return self.Get(self.Len() - 1)
}
func (self *NullableSubscriptionStream) LastOr(arg NullableSubscription) NullableSubscription {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSubscriptionStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableSubscriptionStream) Limit(limit int) *NullableSubscriptionStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableSubscriptionStream) Map(fn func(NullableSubscription, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSubscriptionStream) Map2Int(fn func(NullableSubscription, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSubscriptionStream) Map2Int32(fn func(NullableSubscription, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSubscriptionStream) Map2Int64(fn func(NullableSubscription, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSubscriptionStream) Map2Float32(fn func(NullableSubscription, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSubscriptionStream) Map2Float64(fn func(NullableSubscription, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSubscriptionStream) Map2Bool(fn func(NullableSubscription, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSubscriptionStream) Map2Bytes(fn func(NullableSubscription, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSubscriptionStream) Map2String(fn func(NullableSubscription, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSubscriptionStream) Max(fn func(NullableSubscription, int) float64) *NullableSubscription {
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
func (self *NullableSubscriptionStream) Min(fn func(NullableSubscription, int) float64) *NullableSubscription {
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
func (self *NullableSubscriptionStream) NoneMatch(fn func(NullableSubscription, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableSubscriptionStream) Get(index int) *NullableSubscription {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableSubscriptionStream) GetOr(index int, arg NullableSubscription) NullableSubscription {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSubscriptionStream) Peek(fn func(*NullableSubscription, int)) *NullableSubscriptionStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableSubscriptionStream) Reduce(fn func(NullableSubscription, NullableSubscription, int) NullableSubscription) *NullableSubscriptionStream {
	return self.ReduceInit(fn, NullableSubscription{})
}
func (self *NullableSubscriptionStream) ReduceInit(fn func(NullableSubscription, NullableSubscription, int) NullableSubscription, initialValue NullableSubscription) *NullableSubscriptionStream {
	result := NullableSubscriptionStreamOf()
	self.ForEach(func(v NullableSubscription, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableSubscriptionStream) ReduceInterface(fn func(interface{}, NullableSubscription, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableSubscription{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableSubscriptionStream) ReduceString(fn func(string, NullableSubscription, int) string) []string {
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
func (self *NullableSubscriptionStream) ReduceInt(fn func(int, NullableSubscription, int) int) []int {
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
func (self *NullableSubscriptionStream) ReduceInt32(fn func(int32, NullableSubscription, int) int32) []int32 {
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
func (self *NullableSubscriptionStream) ReduceInt64(fn func(int64, NullableSubscription, int) int64) []int64 {
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
func (self *NullableSubscriptionStream) ReduceFloat32(fn func(float32, NullableSubscription, int) float32) []float32 {
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
func (self *NullableSubscriptionStream) ReduceFloat64(fn func(float64, NullableSubscription, int) float64) []float64 {
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
func (self *NullableSubscriptionStream) ReduceBool(fn func(bool, NullableSubscription, int) bool) []bool {
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
func (self *NullableSubscriptionStream) Reverse() *NullableSubscriptionStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableSubscriptionStream) Replace(fn func(NullableSubscription, int) NullableSubscription) *NullableSubscriptionStream {
	return self.ForEach(func(v NullableSubscription, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableSubscriptionStream) Select(fn func(NullableSubscription) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableSubscriptionStream) Set(index int, val NullableSubscription) *NullableSubscriptionStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableSubscriptionStream) Skip(skip int) *NullableSubscriptionStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableSubscriptionStream) SkippingEach(fn func(NullableSubscription, int) int) *NullableSubscriptionStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableSubscriptionStream) Slice(startIndex, n int) *NullableSubscriptionStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableSubscription{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableSubscriptionStream) Sort(fn func(i, j int) bool) *NullableSubscriptionStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableSubscriptionStream) Tail() *NullableSubscription {
	return self.Last()
}
func (self *NullableSubscriptionStream) TailOr(arg NullableSubscription) NullableSubscription {
	return self.LastOr(arg)
}
func (self *NullableSubscriptionStream) ToList() []NullableSubscription {
	return self.Val()
}
func (self *NullableSubscriptionStream) Unique() *NullableSubscriptionStream {
	return self.Distinct()
}
func (self *NullableSubscriptionStream) Val() []NullableSubscription {
	if self == nil {
		return []NullableSubscription{}
	}
	return *self.Copy()
}
func (self *NullableSubscriptionStream) While(fn func(NullableSubscription, int) bool) *NullableSubscriptionStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableSubscriptionStream) Where(fn func(NullableSubscription) bool) *NullableSubscriptionStream {
	result := NullableSubscriptionStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSubscriptionStream) WhereSlim(fn func(NullableSubscription) bool) *NullableSubscriptionStream {
	result := NullableSubscriptionStreamOf()
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
