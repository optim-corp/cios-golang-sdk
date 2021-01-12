/*
 * Collection utility of Subscription Struct
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

type SubscriptionStream []Subscription

func SubscriptionStreamOf(arg ...Subscription) SubscriptionStream {
	return arg
}
func SubscriptionStreamFrom(arg []Subscription) SubscriptionStream {
	return arg
}
func CreateSubscriptionStream(arg ...Subscription) *SubscriptionStream {
	tmp := SubscriptionStreamOf(arg...)
	return &tmp
}
func GenerateSubscriptionStream(arg []Subscription) *SubscriptionStream {
	tmp := SubscriptionStreamFrom(arg)
	return &tmp
}

func (self *SubscriptionStream) Add(arg Subscription) *SubscriptionStream {
	return self.AddAll(arg)
}
func (self *SubscriptionStream) AddAll(arg ...Subscription) *SubscriptionStream {
	*self = append(*self, arg...)
	return self
}
func (self *SubscriptionStream) AddSafe(arg *Subscription) *SubscriptionStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *SubscriptionStream) Aggregate(fn func(Subscription, Subscription) Subscription) *SubscriptionStream {
	result := SubscriptionStreamOf()
	self.ForEach(func(v Subscription, i int) {
		if i == 0 {
			result.Add(fn(Subscription{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *SubscriptionStream) AllMatch(fn func(Subscription, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *SubscriptionStream) AnyMatch(fn func(Subscription, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *SubscriptionStream) Clone() *SubscriptionStream {
	temp := make([]Subscription, self.Len())
	copy(temp, *self)
	return (*SubscriptionStream)(&temp)
}
func (self *SubscriptionStream) Copy() *SubscriptionStream {
	return self.Clone()
}
func (self *SubscriptionStream) Concat(arg []Subscription) *SubscriptionStream {
	return self.AddAll(arg...)
}
func (self *SubscriptionStream) Contains(arg Subscription) bool {
	return self.FindIndex(func(_arg Subscription, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *SubscriptionStream) Clean() *SubscriptionStream {
	*self = SubscriptionStreamOf()
	return self
}
func (self *SubscriptionStream) Delete(index int) *SubscriptionStream {
	return self.DeleteRange(index, index)
}
func (self *SubscriptionStream) DeleteRange(startIndex, endIndex int) *SubscriptionStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *SubscriptionStream) Distinct() *SubscriptionStream {
	caches := map[string]bool{}
	result := SubscriptionStreamOf()
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
func (self *SubscriptionStream) Each(fn func(Subscription)) *SubscriptionStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *SubscriptionStream) EachRight(fn func(Subscription)) *SubscriptionStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *SubscriptionStream) Equals(arr []Subscription) bool {
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
func (self *SubscriptionStream) Filter(fn func(Subscription, int) bool) *SubscriptionStream {
	result := SubscriptionStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *SubscriptionStream) FilterSlim(fn func(Subscription, int) bool) *SubscriptionStream {
	result := SubscriptionStreamOf()
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
func (self *SubscriptionStream) Find(fn func(Subscription, int) bool) *Subscription {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *SubscriptionStream) FindOr(fn func(Subscription, int) bool, or Subscription) Subscription {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *SubscriptionStream) FindIndex(fn func(Subscription, int) bool) int {
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
func (self *SubscriptionStream) First() *Subscription {
	return self.Get(0)
}
func (self *SubscriptionStream) FirstOr(arg Subscription) Subscription {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *SubscriptionStream) ForEach(fn func(Subscription, int)) *SubscriptionStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *SubscriptionStream) ForEachRight(fn func(Subscription, int)) *SubscriptionStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *SubscriptionStream) GroupBy(fn func(Subscription, int) string) map[string][]Subscription {
	m := map[string][]Subscription{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *SubscriptionStream) GroupByValues(fn func(Subscription, int) string) [][]Subscription {
	var tmp [][]Subscription
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *SubscriptionStream) IndexOf(arg Subscription) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *SubscriptionStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *SubscriptionStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *SubscriptionStream) Last() *Subscription {
	return self.Get(self.Len() - 1)
}
func (self *SubscriptionStream) LastOr(arg Subscription) Subscription {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *SubscriptionStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *SubscriptionStream) Limit(limit int) *SubscriptionStream {
	self.Slice(0, limit)
	return self
}

func (self *SubscriptionStream) Map(fn func(Subscription, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SubscriptionStream) Map2Int(fn func(Subscription, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SubscriptionStream) Map2Int32(fn func(Subscription, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SubscriptionStream) Map2Int64(fn func(Subscription, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SubscriptionStream) Map2Float32(fn func(Subscription, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SubscriptionStream) Map2Float64(fn func(Subscription, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SubscriptionStream) Map2Bool(fn func(Subscription, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SubscriptionStream) Map2Bytes(fn func(Subscription, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SubscriptionStream) Map2String(fn func(Subscription, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SubscriptionStream) Max(fn func(Subscription, int) float64) *Subscription {
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
func (self *SubscriptionStream) Min(fn func(Subscription, int) float64) *Subscription {
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
func (self *SubscriptionStream) NoneMatch(fn func(Subscription, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *SubscriptionStream) Get(index int) *Subscription {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *SubscriptionStream) GetOr(index int, arg Subscription) Subscription {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *SubscriptionStream) Peek(fn func(*Subscription, int)) *SubscriptionStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *SubscriptionStream) Reduce(fn func(Subscription, Subscription, int) Subscription) *SubscriptionStream {
	return self.ReduceInit(fn, Subscription{})
}
func (self *SubscriptionStream) ReduceInit(fn func(Subscription, Subscription, int) Subscription, initialValue Subscription) *SubscriptionStream {
	result := SubscriptionStreamOf()
	self.ForEach(func(v Subscription, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *SubscriptionStream) ReduceInterface(fn func(interface{}, Subscription, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(Subscription{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *SubscriptionStream) ReduceString(fn func(string, Subscription, int) string) []string {
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
func (self *SubscriptionStream) ReduceInt(fn func(int, Subscription, int) int) []int {
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
func (self *SubscriptionStream) ReduceInt32(fn func(int32, Subscription, int) int32) []int32 {
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
func (self *SubscriptionStream) ReduceInt64(fn func(int64, Subscription, int) int64) []int64 {
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
func (self *SubscriptionStream) ReduceFloat32(fn func(float32, Subscription, int) float32) []float32 {
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
func (self *SubscriptionStream) ReduceFloat64(fn func(float64, Subscription, int) float64) []float64 {
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
func (self *SubscriptionStream) ReduceBool(fn func(bool, Subscription, int) bool) []bool {
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
func (self *SubscriptionStream) Reverse() *SubscriptionStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *SubscriptionStream) Replace(fn func(Subscription, int) Subscription) *SubscriptionStream {
	return self.ForEach(func(v Subscription, i int) { self.Set(i, fn(v, i)) })
}
func (self *SubscriptionStream) Select(fn func(Subscription) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *SubscriptionStream) Set(index int, val Subscription) *SubscriptionStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *SubscriptionStream) Skip(skip int) *SubscriptionStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *SubscriptionStream) SkippingEach(fn func(Subscription, int) int) *SubscriptionStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *SubscriptionStream) Slice(startIndex, n int) *SubscriptionStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []Subscription{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *SubscriptionStream) Sort(fn func(i, j int) bool) *SubscriptionStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *SubscriptionStream) Tail() *Subscription {
	return self.Last()
}
func (self *SubscriptionStream) TailOr(arg Subscription) Subscription {
	return self.LastOr(arg)
}
func (self *SubscriptionStream) ToList() []Subscription {
	return self.Val()
}
func (self *SubscriptionStream) Unique() *SubscriptionStream {
	return self.Distinct()
}
func (self *SubscriptionStream) Val() []Subscription {
	if self == nil {
		return []Subscription{}
	}
	return *self.Copy()
}
func (self *SubscriptionStream) While(fn func(Subscription, int) bool) *SubscriptionStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *SubscriptionStream) Where(fn func(Subscription) bool) *SubscriptionStream {
	result := SubscriptionStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *SubscriptionStream) WhereSlim(fn func(Subscription) bool) *SubscriptionStream {
	result := SubscriptionStreamOf()
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