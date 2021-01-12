/*
 * Collection utility of NullableMessagingConfig Struct
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

type NullableMessagingConfigStream []NullableMessagingConfig

func NullableMessagingConfigStreamOf(arg ...NullableMessagingConfig) NullableMessagingConfigStream {
	return arg
}
func NullableMessagingConfigStreamFrom(arg []NullableMessagingConfig) NullableMessagingConfigStream {
	return arg
}
func CreateNullableMessagingConfigStream(arg ...NullableMessagingConfig) *NullableMessagingConfigStream {
	tmp := NullableMessagingConfigStreamOf(arg...)
	return &tmp
}
func GenerateNullableMessagingConfigStream(arg []NullableMessagingConfig) *NullableMessagingConfigStream {
	tmp := NullableMessagingConfigStreamFrom(arg)
	return &tmp
}

func (self *NullableMessagingConfigStream) Add(arg NullableMessagingConfig) *NullableMessagingConfigStream {
	return self.AddAll(arg)
}
func (self *NullableMessagingConfigStream) AddAll(arg ...NullableMessagingConfig) *NullableMessagingConfigStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableMessagingConfigStream) AddSafe(arg *NullableMessagingConfig) *NullableMessagingConfigStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableMessagingConfigStream) Aggregate(fn func(NullableMessagingConfig, NullableMessagingConfig) NullableMessagingConfig) *NullableMessagingConfigStream {
	result := NullableMessagingConfigStreamOf()
	self.ForEach(func(v NullableMessagingConfig, i int) {
		if i == 0 {
			result.Add(fn(NullableMessagingConfig{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableMessagingConfigStream) AllMatch(fn func(NullableMessagingConfig, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableMessagingConfigStream) AnyMatch(fn func(NullableMessagingConfig, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableMessagingConfigStream) Clone() *NullableMessagingConfigStream {
	temp := make([]NullableMessagingConfig, self.Len())
	copy(temp, *self)
	return (*NullableMessagingConfigStream)(&temp)
}
func (self *NullableMessagingConfigStream) Copy() *NullableMessagingConfigStream {
	return self.Clone()
}
func (self *NullableMessagingConfigStream) Concat(arg []NullableMessagingConfig) *NullableMessagingConfigStream {
	return self.AddAll(arg...)
}
func (self *NullableMessagingConfigStream) Contains(arg NullableMessagingConfig) bool {
	return self.FindIndex(func(_arg NullableMessagingConfig, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableMessagingConfigStream) Clean() *NullableMessagingConfigStream {
	*self = NullableMessagingConfigStreamOf()
	return self
}
func (self *NullableMessagingConfigStream) Delete(index int) *NullableMessagingConfigStream {
	return self.DeleteRange(index, index)
}
func (self *NullableMessagingConfigStream) DeleteRange(startIndex, endIndex int) *NullableMessagingConfigStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableMessagingConfigStream) Distinct() *NullableMessagingConfigStream {
	caches := map[string]bool{}
	result := NullableMessagingConfigStreamOf()
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
func (self *NullableMessagingConfigStream) Each(fn func(NullableMessagingConfig)) *NullableMessagingConfigStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableMessagingConfigStream) EachRight(fn func(NullableMessagingConfig)) *NullableMessagingConfigStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableMessagingConfigStream) Equals(arr []NullableMessagingConfig) bool {
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
func (self *NullableMessagingConfigStream) Filter(fn func(NullableMessagingConfig, int) bool) *NullableMessagingConfigStream {
	result := NullableMessagingConfigStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMessagingConfigStream) FilterSlim(fn func(NullableMessagingConfig, int) bool) *NullableMessagingConfigStream {
	result := NullableMessagingConfigStreamOf()
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
func (self *NullableMessagingConfigStream) Find(fn func(NullableMessagingConfig, int) bool) *NullableMessagingConfig {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableMessagingConfigStream) FindOr(fn func(NullableMessagingConfig, int) bool, or NullableMessagingConfig) NullableMessagingConfig {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableMessagingConfigStream) FindIndex(fn func(NullableMessagingConfig, int) bool) int {
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
func (self *NullableMessagingConfigStream) First() *NullableMessagingConfig {
	return self.Get(0)
}
func (self *NullableMessagingConfigStream) FirstOr(arg NullableMessagingConfig) NullableMessagingConfig {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMessagingConfigStream) ForEach(fn func(NullableMessagingConfig, int)) *NullableMessagingConfigStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableMessagingConfigStream) ForEachRight(fn func(NullableMessagingConfig, int)) *NullableMessagingConfigStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableMessagingConfigStream) GroupBy(fn func(NullableMessagingConfig, int) string) map[string][]NullableMessagingConfig {
	m := map[string][]NullableMessagingConfig{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableMessagingConfigStream) GroupByValues(fn func(NullableMessagingConfig, int) string) [][]NullableMessagingConfig {
	var tmp [][]NullableMessagingConfig
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableMessagingConfigStream) IndexOf(arg NullableMessagingConfig) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableMessagingConfigStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableMessagingConfigStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableMessagingConfigStream) Last() *NullableMessagingConfig {
	return self.Get(self.Len() - 1)
}
func (self *NullableMessagingConfigStream) LastOr(arg NullableMessagingConfig) NullableMessagingConfig {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMessagingConfigStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableMessagingConfigStream) Limit(limit int) *NullableMessagingConfigStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableMessagingConfigStream) Map(fn func(NullableMessagingConfig, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMessagingConfigStream) Map2Int(fn func(NullableMessagingConfig, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMessagingConfigStream) Map2Int32(fn func(NullableMessagingConfig, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMessagingConfigStream) Map2Int64(fn func(NullableMessagingConfig, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMessagingConfigStream) Map2Float32(fn func(NullableMessagingConfig, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMessagingConfigStream) Map2Float64(fn func(NullableMessagingConfig, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMessagingConfigStream) Map2Bool(fn func(NullableMessagingConfig, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMessagingConfigStream) Map2Bytes(fn func(NullableMessagingConfig, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMessagingConfigStream) Map2String(fn func(NullableMessagingConfig, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMessagingConfigStream) Max(fn func(NullableMessagingConfig, int) float64) *NullableMessagingConfig {
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
func (self *NullableMessagingConfigStream) Min(fn func(NullableMessagingConfig, int) float64) *NullableMessagingConfig {
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
func (self *NullableMessagingConfigStream) NoneMatch(fn func(NullableMessagingConfig, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableMessagingConfigStream) Get(index int) *NullableMessagingConfig {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableMessagingConfigStream) GetOr(index int, arg NullableMessagingConfig) NullableMessagingConfig {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMessagingConfigStream) Peek(fn func(*NullableMessagingConfig, int)) *NullableMessagingConfigStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableMessagingConfigStream) Reduce(fn func(NullableMessagingConfig, NullableMessagingConfig, int) NullableMessagingConfig) *NullableMessagingConfigStream {
	return self.ReduceInit(fn, NullableMessagingConfig{})
}
func (self *NullableMessagingConfigStream) ReduceInit(fn func(NullableMessagingConfig, NullableMessagingConfig, int) NullableMessagingConfig, initialValue NullableMessagingConfig) *NullableMessagingConfigStream {
	result := NullableMessagingConfigStreamOf()
	self.ForEach(func(v NullableMessagingConfig, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableMessagingConfigStream) ReduceInterface(fn func(interface{}, NullableMessagingConfig, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableMessagingConfig{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMessagingConfigStream) ReduceString(fn func(string, NullableMessagingConfig, int) string) []string {
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
func (self *NullableMessagingConfigStream) ReduceInt(fn func(int, NullableMessagingConfig, int) int) []int {
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
func (self *NullableMessagingConfigStream) ReduceInt32(fn func(int32, NullableMessagingConfig, int) int32) []int32 {
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
func (self *NullableMessagingConfigStream) ReduceInt64(fn func(int64, NullableMessagingConfig, int) int64) []int64 {
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
func (self *NullableMessagingConfigStream) ReduceFloat32(fn func(float32, NullableMessagingConfig, int) float32) []float32 {
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
func (self *NullableMessagingConfigStream) ReduceFloat64(fn func(float64, NullableMessagingConfig, int) float64) []float64 {
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
func (self *NullableMessagingConfigStream) ReduceBool(fn func(bool, NullableMessagingConfig, int) bool) []bool {
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
func (self *NullableMessagingConfigStream) Reverse() *NullableMessagingConfigStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableMessagingConfigStream) Replace(fn func(NullableMessagingConfig, int) NullableMessagingConfig) *NullableMessagingConfigStream {
	return self.ForEach(func(v NullableMessagingConfig, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableMessagingConfigStream) Select(fn func(NullableMessagingConfig) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableMessagingConfigStream) Set(index int, val NullableMessagingConfig) *NullableMessagingConfigStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableMessagingConfigStream) Skip(skip int) *NullableMessagingConfigStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableMessagingConfigStream) SkippingEach(fn func(NullableMessagingConfig, int) int) *NullableMessagingConfigStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableMessagingConfigStream) Slice(startIndex, n int) *NullableMessagingConfigStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableMessagingConfig{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableMessagingConfigStream) Sort(fn func(i, j int) bool) *NullableMessagingConfigStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableMessagingConfigStream) Tail() *NullableMessagingConfig {
	return self.Last()
}
func (self *NullableMessagingConfigStream) TailOr(arg NullableMessagingConfig) NullableMessagingConfig {
	return self.LastOr(arg)
}
func (self *NullableMessagingConfigStream) ToList() []NullableMessagingConfig {
	return self.Val()
}
func (self *NullableMessagingConfigStream) Unique() *NullableMessagingConfigStream {
	return self.Distinct()
}
func (self *NullableMessagingConfigStream) Val() []NullableMessagingConfig {
	if self == nil {
		return []NullableMessagingConfig{}
	}
	return *self.Copy()
}
func (self *NullableMessagingConfigStream) While(fn func(NullableMessagingConfig, int) bool) *NullableMessagingConfigStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableMessagingConfigStream) Where(fn func(NullableMessagingConfig) bool) *NullableMessagingConfigStream {
	result := NullableMessagingConfigStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMessagingConfigStream) WhereSlim(fn func(NullableMessagingConfig) bool) *NullableMessagingConfigStream {
	result := NullableMessagingConfigStreamOf()
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
