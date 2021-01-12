/*
 * Collection utility of NullableMultipleContractUser Struct
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

type NullableMultipleContractUserStream []NullableMultipleContractUser

func NullableMultipleContractUserStreamOf(arg ...NullableMultipleContractUser) NullableMultipleContractUserStream {
	return arg
}
func NullableMultipleContractUserStreamFrom(arg []NullableMultipleContractUser) NullableMultipleContractUserStream {
	return arg
}
func CreateNullableMultipleContractUserStream(arg ...NullableMultipleContractUser) *NullableMultipleContractUserStream {
	tmp := NullableMultipleContractUserStreamOf(arg...)
	return &tmp
}
func GenerateNullableMultipleContractUserStream(arg []NullableMultipleContractUser) *NullableMultipleContractUserStream {
	tmp := NullableMultipleContractUserStreamFrom(arg)
	return &tmp
}

func (self *NullableMultipleContractUserStream) Add(arg NullableMultipleContractUser) *NullableMultipleContractUserStream {
	return self.AddAll(arg)
}
func (self *NullableMultipleContractUserStream) AddAll(arg ...NullableMultipleContractUser) *NullableMultipleContractUserStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableMultipleContractUserStream) AddSafe(arg *NullableMultipleContractUser) *NullableMultipleContractUserStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableMultipleContractUserStream) Aggregate(fn func(NullableMultipleContractUser, NullableMultipleContractUser) NullableMultipleContractUser) *NullableMultipleContractUserStream {
	result := NullableMultipleContractUserStreamOf()
	self.ForEach(func(v NullableMultipleContractUser, i int) {
		if i == 0 {
			result.Add(fn(NullableMultipleContractUser{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleContractUserStream) AllMatch(fn func(NullableMultipleContractUser, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableMultipleContractUserStream) AnyMatch(fn func(NullableMultipleContractUser, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableMultipleContractUserStream) Clone() *NullableMultipleContractUserStream {
	temp := make([]NullableMultipleContractUser, self.Len())
	copy(temp, *self)
	return (*NullableMultipleContractUserStream)(&temp)
}
func (self *NullableMultipleContractUserStream) Copy() *NullableMultipleContractUserStream {
	return self.Clone()
}
func (self *NullableMultipleContractUserStream) Concat(arg []NullableMultipleContractUser) *NullableMultipleContractUserStream {
	return self.AddAll(arg...)
}
func (self *NullableMultipleContractUserStream) Contains(arg NullableMultipleContractUser) bool {
	return self.FindIndex(func(_arg NullableMultipleContractUser, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableMultipleContractUserStream) Clean() *NullableMultipleContractUserStream {
	*self = NullableMultipleContractUserStreamOf()
	return self
}
func (self *NullableMultipleContractUserStream) Delete(index int) *NullableMultipleContractUserStream {
	return self.DeleteRange(index, index)
}
func (self *NullableMultipleContractUserStream) DeleteRange(startIndex, endIndex int) *NullableMultipleContractUserStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableMultipleContractUserStream) Distinct() *NullableMultipleContractUserStream {
	caches := map[string]bool{}
	result := NullableMultipleContractUserStreamOf()
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
func (self *NullableMultipleContractUserStream) Each(fn func(NullableMultipleContractUser)) *NullableMultipleContractUserStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableMultipleContractUserStream) EachRight(fn func(NullableMultipleContractUser)) *NullableMultipleContractUserStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableMultipleContractUserStream) Equals(arr []NullableMultipleContractUser) bool {
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
func (self *NullableMultipleContractUserStream) Filter(fn func(NullableMultipleContractUser, int) bool) *NullableMultipleContractUserStream {
	result := NullableMultipleContractUserStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleContractUserStream) FilterSlim(fn func(NullableMultipleContractUser, int) bool) *NullableMultipleContractUserStream {
	result := NullableMultipleContractUserStreamOf()
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
func (self *NullableMultipleContractUserStream) Find(fn func(NullableMultipleContractUser, int) bool) *NullableMultipleContractUser {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleContractUserStream) FindOr(fn func(NullableMultipleContractUser, int) bool, or NullableMultipleContractUser) NullableMultipleContractUser {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableMultipleContractUserStream) FindIndex(fn func(NullableMultipleContractUser, int) bool) int {
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
func (self *NullableMultipleContractUserStream) First() *NullableMultipleContractUser {
	return self.Get(0)
}
func (self *NullableMultipleContractUserStream) FirstOr(arg NullableMultipleContractUser) NullableMultipleContractUser {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleContractUserStream) ForEach(fn func(NullableMultipleContractUser, int)) *NullableMultipleContractUserStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableMultipleContractUserStream) ForEachRight(fn func(NullableMultipleContractUser, int)) *NullableMultipleContractUserStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableMultipleContractUserStream) GroupBy(fn func(NullableMultipleContractUser, int) string) map[string][]NullableMultipleContractUser {
	m := map[string][]NullableMultipleContractUser{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableMultipleContractUserStream) GroupByValues(fn func(NullableMultipleContractUser, int) string) [][]NullableMultipleContractUser {
	var tmp [][]NullableMultipleContractUser
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableMultipleContractUserStream) IndexOf(arg NullableMultipleContractUser) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableMultipleContractUserStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableMultipleContractUserStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableMultipleContractUserStream) Last() *NullableMultipleContractUser {
	return self.Get(self.Len() - 1)
}
func (self *NullableMultipleContractUserStream) LastOr(arg NullableMultipleContractUser) NullableMultipleContractUser {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleContractUserStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableMultipleContractUserStream) Limit(limit int) *NullableMultipleContractUserStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableMultipleContractUserStream) Map(fn func(NullableMultipleContractUser, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleContractUserStream) Map2Int(fn func(NullableMultipleContractUser, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleContractUserStream) Map2Int32(fn func(NullableMultipleContractUser, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleContractUserStream) Map2Int64(fn func(NullableMultipleContractUser, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleContractUserStream) Map2Float32(fn func(NullableMultipleContractUser, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleContractUserStream) Map2Float64(fn func(NullableMultipleContractUser, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleContractUserStream) Map2Bool(fn func(NullableMultipleContractUser, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleContractUserStream) Map2Bytes(fn func(NullableMultipleContractUser, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleContractUserStream) Map2String(fn func(NullableMultipleContractUser, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleContractUserStream) Max(fn func(NullableMultipleContractUser, int) float64) *NullableMultipleContractUser {
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
func (self *NullableMultipleContractUserStream) Min(fn func(NullableMultipleContractUser, int) float64) *NullableMultipleContractUser {
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
func (self *NullableMultipleContractUserStream) NoneMatch(fn func(NullableMultipleContractUser, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableMultipleContractUserStream) Get(index int) *NullableMultipleContractUser {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleContractUserStream) GetOr(index int, arg NullableMultipleContractUser) NullableMultipleContractUser {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleContractUserStream) Peek(fn func(*NullableMultipleContractUser, int)) *NullableMultipleContractUserStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableMultipleContractUserStream) Reduce(fn func(NullableMultipleContractUser, NullableMultipleContractUser, int) NullableMultipleContractUser) *NullableMultipleContractUserStream {
	return self.ReduceInit(fn, NullableMultipleContractUser{})
}
func (self *NullableMultipleContractUserStream) ReduceInit(fn func(NullableMultipleContractUser, NullableMultipleContractUser, int) NullableMultipleContractUser, initialValue NullableMultipleContractUser) *NullableMultipleContractUserStream {
	result := NullableMultipleContractUserStreamOf()
	self.ForEach(func(v NullableMultipleContractUser, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleContractUserStream) ReduceInterface(fn func(interface{}, NullableMultipleContractUser, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableMultipleContractUser{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleContractUserStream) ReduceString(fn func(string, NullableMultipleContractUser, int) string) []string {
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
func (self *NullableMultipleContractUserStream) ReduceInt(fn func(int, NullableMultipleContractUser, int) int) []int {
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
func (self *NullableMultipleContractUserStream) ReduceInt32(fn func(int32, NullableMultipleContractUser, int) int32) []int32 {
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
func (self *NullableMultipleContractUserStream) ReduceInt64(fn func(int64, NullableMultipleContractUser, int) int64) []int64 {
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
func (self *NullableMultipleContractUserStream) ReduceFloat32(fn func(float32, NullableMultipleContractUser, int) float32) []float32 {
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
func (self *NullableMultipleContractUserStream) ReduceFloat64(fn func(float64, NullableMultipleContractUser, int) float64) []float64 {
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
func (self *NullableMultipleContractUserStream) ReduceBool(fn func(bool, NullableMultipleContractUser, int) bool) []bool {
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
func (self *NullableMultipleContractUserStream) Reverse() *NullableMultipleContractUserStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableMultipleContractUserStream) Replace(fn func(NullableMultipleContractUser, int) NullableMultipleContractUser) *NullableMultipleContractUserStream {
	return self.ForEach(func(v NullableMultipleContractUser, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableMultipleContractUserStream) Select(fn func(NullableMultipleContractUser) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableMultipleContractUserStream) Set(index int, val NullableMultipleContractUser) *NullableMultipleContractUserStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableMultipleContractUserStream) Skip(skip int) *NullableMultipleContractUserStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableMultipleContractUserStream) SkippingEach(fn func(NullableMultipleContractUser, int) int) *NullableMultipleContractUserStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableMultipleContractUserStream) Slice(startIndex, n int) *NullableMultipleContractUserStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableMultipleContractUser{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableMultipleContractUserStream) Sort(fn func(i, j int) bool) *NullableMultipleContractUserStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableMultipleContractUserStream) Tail() *NullableMultipleContractUser {
	return self.Last()
}
func (self *NullableMultipleContractUserStream) TailOr(arg NullableMultipleContractUser) NullableMultipleContractUser {
	return self.LastOr(arg)
}
func (self *NullableMultipleContractUserStream) ToList() []NullableMultipleContractUser {
	return self.Val()
}
func (self *NullableMultipleContractUserStream) Unique() *NullableMultipleContractUserStream {
	return self.Distinct()
}
func (self *NullableMultipleContractUserStream) Val() []NullableMultipleContractUser {
	if self == nil {
		return []NullableMultipleContractUser{}
	}
	return *self.Copy()
}
func (self *NullableMultipleContractUserStream) While(fn func(NullableMultipleContractUser, int) bool) *NullableMultipleContractUserStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableMultipleContractUserStream) Where(fn func(NullableMultipleContractUser) bool) *NullableMultipleContractUserStream {
	result := NullableMultipleContractUserStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleContractUserStream) WhereSlim(fn func(NullableMultipleContractUser) bool) *NullableMultipleContractUserStream {
	result := NullableMultipleContractUserStreamOf()
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
