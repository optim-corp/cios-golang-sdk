/*
 * Collection utility of NullableDeviceClientRsaKey Struct
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

type NullableDeviceClientRsaKeyStream []NullableDeviceClientRsaKey

func NullableDeviceClientRsaKeyStreamOf(arg ...NullableDeviceClientRsaKey) NullableDeviceClientRsaKeyStream {
	return arg
}
func NullableDeviceClientRsaKeyStreamFrom(arg []NullableDeviceClientRsaKey) NullableDeviceClientRsaKeyStream {
	return arg
}
func CreateNullableDeviceClientRsaKeyStream(arg ...NullableDeviceClientRsaKey) *NullableDeviceClientRsaKeyStream {
	tmp := NullableDeviceClientRsaKeyStreamOf(arg...)
	return &tmp
}
func GenerateNullableDeviceClientRsaKeyStream(arg []NullableDeviceClientRsaKey) *NullableDeviceClientRsaKeyStream {
	tmp := NullableDeviceClientRsaKeyStreamFrom(arg)
	return &tmp
}

func (self *NullableDeviceClientRsaKeyStream) Add(arg NullableDeviceClientRsaKey) *NullableDeviceClientRsaKeyStream {
	return self.AddAll(arg)
}
func (self *NullableDeviceClientRsaKeyStream) AddAll(arg ...NullableDeviceClientRsaKey) *NullableDeviceClientRsaKeyStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableDeviceClientRsaKeyStream) AddSafe(arg *NullableDeviceClientRsaKey) *NullableDeviceClientRsaKeyStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableDeviceClientRsaKeyStream) Aggregate(fn func(NullableDeviceClientRsaKey, NullableDeviceClientRsaKey) NullableDeviceClientRsaKey) *NullableDeviceClientRsaKeyStream {
	result := NullableDeviceClientRsaKeyStreamOf()
	self.ForEach(func(v NullableDeviceClientRsaKey, i int) {
		if i == 0 {
			result.Add(fn(NullableDeviceClientRsaKey{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableDeviceClientRsaKeyStream) AllMatch(fn func(NullableDeviceClientRsaKey, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableDeviceClientRsaKeyStream) AnyMatch(fn func(NullableDeviceClientRsaKey, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableDeviceClientRsaKeyStream) Clone() *NullableDeviceClientRsaKeyStream {
	temp := make([]NullableDeviceClientRsaKey, self.Len())
	copy(temp, *self)
	return (*NullableDeviceClientRsaKeyStream)(&temp)
}
func (self *NullableDeviceClientRsaKeyStream) Copy() *NullableDeviceClientRsaKeyStream {
	return self.Clone()
}
func (self *NullableDeviceClientRsaKeyStream) Concat(arg []NullableDeviceClientRsaKey) *NullableDeviceClientRsaKeyStream {
	return self.AddAll(arg...)
}
func (self *NullableDeviceClientRsaKeyStream) Contains(arg NullableDeviceClientRsaKey) bool {
	return self.FindIndex(func(_arg NullableDeviceClientRsaKey, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableDeviceClientRsaKeyStream) Clean() *NullableDeviceClientRsaKeyStream {
	*self = NullableDeviceClientRsaKeyStreamOf()
	return self
}
func (self *NullableDeviceClientRsaKeyStream) Delete(index int) *NullableDeviceClientRsaKeyStream {
	return self.DeleteRange(index, index)
}
func (self *NullableDeviceClientRsaKeyStream) DeleteRange(startIndex, endIndex int) *NullableDeviceClientRsaKeyStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableDeviceClientRsaKeyStream) Distinct() *NullableDeviceClientRsaKeyStream {
	caches := map[string]bool{}
	result := NullableDeviceClientRsaKeyStreamOf()
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
func (self *NullableDeviceClientRsaKeyStream) Each(fn func(NullableDeviceClientRsaKey)) *NullableDeviceClientRsaKeyStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableDeviceClientRsaKeyStream) EachRight(fn func(NullableDeviceClientRsaKey)) *NullableDeviceClientRsaKeyStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableDeviceClientRsaKeyStream) Equals(arr []NullableDeviceClientRsaKey) bool {
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
func (self *NullableDeviceClientRsaKeyStream) Filter(fn func(NullableDeviceClientRsaKey, int) bool) *NullableDeviceClientRsaKeyStream {
	result := NullableDeviceClientRsaKeyStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableDeviceClientRsaKeyStream) FilterSlim(fn func(NullableDeviceClientRsaKey, int) bool) *NullableDeviceClientRsaKeyStream {
	result := NullableDeviceClientRsaKeyStreamOf()
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
func (self *NullableDeviceClientRsaKeyStream) Find(fn func(NullableDeviceClientRsaKey, int) bool) *NullableDeviceClientRsaKey {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableDeviceClientRsaKeyStream) FindOr(fn func(NullableDeviceClientRsaKey, int) bool, or NullableDeviceClientRsaKey) NullableDeviceClientRsaKey {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableDeviceClientRsaKeyStream) FindIndex(fn func(NullableDeviceClientRsaKey, int) bool) int {
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
func (self *NullableDeviceClientRsaKeyStream) First() *NullableDeviceClientRsaKey {
	return self.Get(0)
}
func (self *NullableDeviceClientRsaKeyStream) FirstOr(arg NullableDeviceClientRsaKey) NullableDeviceClientRsaKey {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDeviceClientRsaKeyStream) ForEach(fn func(NullableDeviceClientRsaKey, int)) *NullableDeviceClientRsaKeyStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableDeviceClientRsaKeyStream) ForEachRight(fn func(NullableDeviceClientRsaKey, int)) *NullableDeviceClientRsaKeyStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableDeviceClientRsaKeyStream) GroupBy(fn func(NullableDeviceClientRsaKey, int) string) map[string][]NullableDeviceClientRsaKey {
	m := map[string][]NullableDeviceClientRsaKey{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableDeviceClientRsaKeyStream) GroupByValues(fn func(NullableDeviceClientRsaKey, int) string) [][]NullableDeviceClientRsaKey {
	var tmp [][]NullableDeviceClientRsaKey
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableDeviceClientRsaKeyStream) IndexOf(arg NullableDeviceClientRsaKey) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableDeviceClientRsaKeyStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableDeviceClientRsaKeyStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableDeviceClientRsaKeyStream) Last() *NullableDeviceClientRsaKey {
	return self.Get(self.Len() - 1)
}
func (self *NullableDeviceClientRsaKeyStream) LastOr(arg NullableDeviceClientRsaKey) NullableDeviceClientRsaKey {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDeviceClientRsaKeyStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableDeviceClientRsaKeyStream) Limit(limit int) *NullableDeviceClientRsaKeyStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableDeviceClientRsaKeyStream) Map(fn func(NullableDeviceClientRsaKey, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceClientRsaKeyStream) Map2Int(fn func(NullableDeviceClientRsaKey, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceClientRsaKeyStream) Map2Int32(fn func(NullableDeviceClientRsaKey, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceClientRsaKeyStream) Map2Int64(fn func(NullableDeviceClientRsaKey, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceClientRsaKeyStream) Map2Float32(fn func(NullableDeviceClientRsaKey, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceClientRsaKeyStream) Map2Float64(fn func(NullableDeviceClientRsaKey, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceClientRsaKeyStream) Map2Bool(fn func(NullableDeviceClientRsaKey, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceClientRsaKeyStream) Map2Bytes(fn func(NullableDeviceClientRsaKey, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceClientRsaKeyStream) Map2String(fn func(NullableDeviceClientRsaKey, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceClientRsaKeyStream) Max(fn func(NullableDeviceClientRsaKey, int) float64) *NullableDeviceClientRsaKey {
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
func (self *NullableDeviceClientRsaKeyStream) Min(fn func(NullableDeviceClientRsaKey, int) float64) *NullableDeviceClientRsaKey {
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
func (self *NullableDeviceClientRsaKeyStream) NoneMatch(fn func(NullableDeviceClientRsaKey, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableDeviceClientRsaKeyStream) Get(index int) *NullableDeviceClientRsaKey {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableDeviceClientRsaKeyStream) GetOr(index int, arg NullableDeviceClientRsaKey) NullableDeviceClientRsaKey {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDeviceClientRsaKeyStream) Peek(fn func(*NullableDeviceClientRsaKey, int)) *NullableDeviceClientRsaKeyStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableDeviceClientRsaKeyStream) Reduce(fn func(NullableDeviceClientRsaKey, NullableDeviceClientRsaKey, int) NullableDeviceClientRsaKey) *NullableDeviceClientRsaKeyStream {
	return self.ReduceInit(fn, NullableDeviceClientRsaKey{})
}
func (self *NullableDeviceClientRsaKeyStream) ReduceInit(fn func(NullableDeviceClientRsaKey, NullableDeviceClientRsaKey, int) NullableDeviceClientRsaKey, initialValue NullableDeviceClientRsaKey) *NullableDeviceClientRsaKeyStream {
	result := NullableDeviceClientRsaKeyStreamOf()
	self.ForEach(func(v NullableDeviceClientRsaKey, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableDeviceClientRsaKeyStream) ReduceInterface(fn func(interface{}, NullableDeviceClientRsaKey, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableDeviceClientRsaKey{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableDeviceClientRsaKeyStream) ReduceString(fn func(string, NullableDeviceClientRsaKey, int) string) []string {
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
func (self *NullableDeviceClientRsaKeyStream) ReduceInt(fn func(int, NullableDeviceClientRsaKey, int) int) []int {
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
func (self *NullableDeviceClientRsaKeyStream) ReduceInt32(fn func(int32, NullableDeviceClientRsaKey, int) int32) []int32 {
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
func (self *NullableDeviceClientRsaKeyStream) ReduceInt64(fn func(int64, NullableDeviceClientRsaKey, int) int64) []int64 {
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
func (self *NullableDeviceClientRsaKeyStream) ReduceFloat32(fn func(float32, NullableDeviceClientRsaKey, int) float32) []float32 {
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
func (self *NullableDeviceClientRsaKeyStream) ReduceFloat64(fn func(float64, NullableDeviceClientRsaKey, int) float64) []float64 {
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
func (self *NullableDeviceClientRsaKeyStream) ReduceBool(fn func(bool, NullableDeviceClientRsaKey, int) bool) []bool {
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
func (self *NullableDeviceClientRsaKeyStream) Reverse() *NullableDeviceClientRsaKeyStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableDeviceClientRsaKeyStream) Replace(fn func(NullableDeviceClientRsaKey, int) NullableDeviceClientRsaKey) *NullableDeviceClientRsaKeyStream {
	return self.ForEach(func(v NullableDeviceClientRsaKey, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableDeviceClientRsaKeyStream) Select(fn func(NullableDeviceClientRsaKey) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableDeviceClientRsaKeyStream) Set(index int, val NullableDeviceClientRsaKey) *NullableDeviceClientRsaKeyStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableDeviceClientRsaKeyStream) Skip(skip int) *NullableDeviceClientRsaKeyStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableDeviceClientRsaKeyStream) SkippingEach(fn func(NullableDeviceClientRsaKey, int) int) *NullableDeviceClientRsaKeyStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableDeviceClientRsaKeyStream) Slice(startIndex, n int) *NullableDeviceClientRsaKeyStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableDeviceClientRsaKey{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableDeviceClientRsaKeyStream) Sort(fn func(i, j int) bool) *NullableDeviceClientRsaKeyStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableDeviceClientRsaKeyStream) Tail() *NullableDeviceClientRsaKey {
	return self.Last()
}
func (self *NullableDeviceClientRsaKeyStream) TailOr(arg NullableDeviceClientRsaKey) NullableDeviceClientRsaKey {
	return self.LastOr(arg)
}
func (self *NullableDeviceClientRsaKeyStream) ToList() []NullableDeviceClientRsaKey {
	return self.Val()
}
func (self *NullableDeviceClientRsaKeyStream) Unique() *NullableDeviceClientRsaKeyStream {
	return self.Distinct()
}
func (self *NullableDeviceClientRsaKeyStream) Val() []NullableDeviceClientRsaKey {
	if self == nil {
		return []NullableDeviceClientRsaKey{}
	}
	return *self.Copy()
}
func (self *NullableDeviceClientRsaKeyStream) While(fn func(NullableDeviceClientRsaKey, int) bool) *NullableDeviceClientRsaKeyStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableDeviceClientRsaKeyStream) Where(fn func(NullableDeviceClientRsaKey) bool) *NullableDeviceClientRsaKeyStream {
	result := NullableDeviceClientRsaKeyStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableDeviceClientRsaKeyStream) WhereSlim(fn func(NullableDeviceClientRsaKey) bool) *NullableDeviceClientRsaKeyStream {
	result := NullableDeviceClientRsaKeyStreamOf()
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
