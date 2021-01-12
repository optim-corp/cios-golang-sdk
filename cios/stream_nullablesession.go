/*
 * Collection utility of NullableSession Struct
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

type NullableSessionStream []NullableSession

func NullableSessionStreamOf(arg ...NullableSession) NullableSessionStream {
	return arg
}
func NullableSessionStreamFrom(arg []NullableSession) NullableSessionStream {
	return arg
}
func CreateNullableSessionStream(arg ...NullableSession) *NullableSessionStream {
	tmp := NullableSessionStreamOf(arg...)
	return &tmp
}
func GenerateNullableSessionStream(arg []NullableSession) *NullableSessionStream {
	tmp := NullableSessionStreamFrom(arg)
	return &tmp
}

func (self *NullableSessionStream) Add(arg NullableSession) *NullableSessionStream {
	return self.AddAll(arg)
}
func (self *NullableSessionStream) AddAll(arg ...NullableSession) *NullableSessionStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableSessionStream) AddSafe(arg *NullableSession) *NullableSessionStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableSessionStream) Aggregate(fn func(NullableSession, NullableSession) NullableSession) *NullableSessionStream {
	result := NullableSessionStreamOf()
	self.ForEach(func(v NullableSession, i int) {
		if i == 0 {
			result.Add(fn(NullableSession{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableSessionStream) AllMatch(fn func(NullableSession, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableSessionStream) AnyMatch(fn func(NullableSession, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableSessionStream) Clone() *NullableSessionStream {
	temp := make([]NullableSession, self.Len())
	copy(temp, *self)
	return (*NullableSessionStream)(&temp)
}
func (self *NullableSessionStream) Copy() *NullableSessionStream {
	return self.Clone()
}
func (self *NullableSessionStream) Concat(arg []NullableSession) *NullableSessionStream {
	return self.AddAll(arg...)
}
func (self *NullableSessionStream) Contains(arg NullableSession) bool {
	return self.FindIndex(func(_arg NullableSession, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableSessionStream) Clean() *NullableSessionStream {
	*self = NullableSessionStreamOf()
	return self
}
func (self *NullableSessionStream) Delete(index int) *NullableSessionStream {
	return self.DeleteRange(index, index)
}
func (self *NullableSessionStream) DeleteRange(startIndex, endIndex int) *NullableSessionStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableSessionStream) Distinct() *NullableSessionStream {
	caches := map[string]bool{}
	result := NullableSessionStreamOf()
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
func (self *NullableSessionStream) Each(fn func(NullableSession)) *NullableSessionStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableSessionStream) EachRight(fn func(NullableSession)) *NullableSessionStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableSessionStream) Equals(arr []NullableSession) bool {
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
func (self *NullableSessionStream) Filter(fn func(NullableSession, int) bool) *NullableSessionStream {
	result := NullableSessionStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSessionStream) FilterSlim(fn func(NullableSession, int) bool) *NullableSessionStream {
	result := NullableSessionStreamOf()
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
func (self *NullableSessionStream) Find(fn func(NullableSession, int) bool) *NullableSession {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableSessionStream) FindOr(fn func(NullableSession, int) bool, or NullableSession) NullableSession {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableSessionStream) FindIndex(fn func(NullableSession, int) bool) int {
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
func (self *NullableSessionStream) First() *NullableSession {
	return self.Get(0)
}
func (self *NullableSessionStream) FirstOr(arg NullableSession) NullableSession {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSessionStream) ForEach(fn func(NullableSession, int)) *NullableSessionStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableSessionStream) ForEachRight(fn func(NullableSession, int)) *NullableSessionStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableSessionStream) GroupBy(fn func(NullableSession, int) string) map[string][]NullableSession {
	m := map[string][]NullableSession{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableSessionStream) GroupByValues(fn func(NullableSession, int) string) [][]NullableSession {
	var tmp [][]NullableSession
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableSessionStream) IndexOf(arg NullableSession) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableSessionStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableSessionStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableSessionStream) Last() *NullableSession {
	return self.Get(self.Len() - 1)
}
func (self *NullableSessionStream) LastOr(arg NullableSession) NullableSession {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSessionStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableSessionStream) Limit(limit int) *NullableSessionStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableSessionStream) Map(fn func(NullableSession, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSessionStream) Map2Int(fn func(NullableSession, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSessionStream) Map2Int32(fn func(NullableSession, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSessionStream) Map2Int64(fn func(NullableSession, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSessionStream) Map2Float32(fn func(NullableSession, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSessionStream) Map2Float64(fn func(NullableSession, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSessionStream) Map2Bool(fn func(NullableSession, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSessionStream) Map2Bytes(fn func(NullableSession, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSessionStream) Map2String(fn func(NullableSession, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSessionStream) Max(fn func(NullableSession, int) float64) *NullableSession {
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
func (self *NullableSessionStream) Min(fn func(NullableSession, int) float64) *NullableSession {
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
func (self *NullableSessionStream) NoneMatch(fn func(NullableSession, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableSessionStream) Get(index int) *NullableSession {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableSessionStream) GetOr(index int, arg NullableSession) NullableSession {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSessionStream) Peek(fn func(*NullableSession, int)) *NullableSessionStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableSessionStream) Reduce(fn func(NullableSession, NullableSession, int) NullableSession) *NullableSessionStream {
	return self.ReduceInit(fn, NullableSession{})
}
func (self *NullableSessionStream) ReduceInit(fn func(NullableSession, NullableSession, int) NullableSession, initialValue NullableSession) *NullableSessionStream {
	result := NullableSessionStreamOf()
	self.ForEach(func(v NullableSession, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableSessionStream) ReduceInterface(fn func(interface{}, NullableSession, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableSession{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableSessionStream) ReduceString(fn func(string, NullableSession, int) string) []string {
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
func (self *NullableSessionStream) ReduceInt(fn func(int, NullableSession, int) int) []int {
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
func (self *NullableSessionStream) ReduceInt32(fn func(int32, NullableSession, int) int32) []int32 {
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
func (self *NullableSessionStream) ReduceInt64(fn func(int64, NullableSession, int) int64) []int64 {
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
func (self *NullableSessionStream) ReduceFloat32(fn func(float32, NullableSession, int) float32) []float32 {
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
func (self *NullableSessionStream) ReduceFloat64(fn func(float64, NullableSession, int) float64) []float64 {
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
func (self *NullableSessionStream) ReduceBool(fn func(bool, NullableSession, int) bool) []bool {
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
func (self *NullableSessionStream) Reverse() *NullableSessionStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableSessionStream) Replace(fn func(NullableSession, int) NullableSession) *NullableSessionStream {
	return self.ForEach(func(v NullableSession, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableSessionStream) Select(fn func(NullableSession) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableSessionStream) Set(index int, val NullableSession) *NullableSessionStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableSessionStream) Skip(skip int) *NullableSessionStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableSessionStream) SkippingEach(fn func(NullableSession, int) int) *NullableSessionStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableSessionStream) Slice(startIndex, n int) *NullableSessionStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableSession{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableSessionStream) Sort(fn func(i, j int) bool) *NullableSessionStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableSessionStream) Tail() *NullableSession {
	return self.Last()
}
func (self *NullableSessionStream) TailOr(arg NullableSession) NullableSession {
	return self.LastOr(arg)
}
func (self *NullableSessionStream) ToList() []NullableSession {
	return self.Val()
}
func (self *NullableSessionStream) Unique() *NullableSessionStream {
	return self.Distinct()
}
func (self *NullableSessionStream) Val() []NullableSession {
	if self == nil {
		return []NullableSession{}
	}
	return *self.Copy()
}
func (self *NullableSessionStream) While(fn func(NullableSession, int) bool) *NullableSessionStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableSessionStream) Where(fn func(NullableSession) bool) *NullableSessionStream {
	result := NullableSessionStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSessionStream) WhereSlim(fn func(NullableSession) bool) *NullableSessionStream {
	result := NullableSessionStreamOf()
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
