/*
 * Collection utility of NullableContract Struct
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

type NullableContractStream []NullableContract

func NullableContractStreamOf(arg ...NullableContract) NullableContractStream {
	return arg
}
func NullableContractStreamFrom(arg []NullableContract) NullableContractStream {
	return arg
}
func CreateNullableContractStream(arg ...NullableContract) *NullableContractStream {
	tmp := NullableContractStreamOf(arg...)
	return &tmp
}
func GenerateNullableContractStream(arg []NullableContract) *NullableContractStream {
	tmp := NullableContractStreamFrom(arg)
	return &tmp
}

func (self *NullableContractStream) Add(arg NullableContract) *NullableContractStream {
	return self.AddAll(arg)
}
func (self *NullableContractStream) AddAll(arg ...NullableContract) *NullableContractStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableContractStream) AddSafe(arg *NullableContract) *NullableContractStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableContractStream) Aggregate(fn func(NullableContract, NullableContract) NullableContract) *NullableContractStream {
	result := NullableContractStreamOf()
	self.ForEach(func(v NullableContract, i int) {
		if i == 0 {
			result.Add(fn(NullableContract{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableContractStream) AllMatch(fn func(NullableContract, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableContractStream) AnyMatch(fn func(NullableContract, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableContractStream) Clone() *NullableContractStream {
	temp := make([]NullableContract, self.Len())
	copy(temp, *self)
	return (*NullableContractStream)(&temp)
}
func (self *NullableContractStream) Copy() *NullableContractStream {
	return self.Clone()
}
func (self *NullableContractStream) Concat(arg []NullableContract) *NullableContractStream {
	return self.AddAll(arg...)
}
func (self *NullableContractStream) Contains(arg NullableContract) bool {
	return self.FindIndex(func(_arg NullableContract, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableContractStream) Clean() *NullableContractStream {
	*self = NullableContractStreamOf()
	return self
}
func (self *NullableContractStream) Delete(index int) *NullableContractStream {
	return self.DeleteRange(index, index)
}
func (self *NullableContractStream) DeleteRange(startIndex, endIndex int) *NullableContractStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableContractStream) Distinct() *NullableContractStream {
	caches := map[string]bool{}
	result := NullableContractStreamOf()
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
func (self *NullableContractStream) Each(fn func(NullableContract)) *NullableContractStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableContractStream) EachRight(fn func(NullableContract)) *NullableContractStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableContractStream) Equals(arr []NullableContract) bool {
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
func (self *NullableContractStream) Filter(fn func(NullableContract, int) bool) *NullableContractStream {
	result := NullableContractStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableContractStream) FilterSlim(fn func(NullableContract, int) bool) *NullableContractStream {
	result := NullableContractStreamOf()
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
func (self *NullableContractStream) Find(fn func(NullableContract, int) bool) *NullableContract {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableContractStream) FindOr(fn func(NullableContract, int) bool, or NullableContract) NullableContract {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableContractStream) FindIndex(fn func(NullableContract, int) bool) int {
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
func (self *NullableContractStream) First() *NullableContract {
	return self.Get(0)
}
func (self *NullableContractStream) FirstOr(arg NullableContract) NullableContract {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableContractStream) ForEach(fn func(NullableContract, int)) *NullableContractStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableContractStream) ForEachRight(fn func(NullableContract, int)) *NullableContractStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableContractStream) GroupBy(fn func(NullableContract, int) string) map[string][]NullableContract {
	m := map[string][]NullableContract{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableContractStream) GroupByValues(fn func(NullableContract, int) string) [][]NullableContract {
	var tmp [][]NullableContract
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableContractStream) IndexOf(arg NullableContract) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableContractStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableContractStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableContractStream) Last() *NullableContract {
	return self.Get(self.Len() - 1)
}
func (self *NullableContractStream) LastOr(arg NullableContract) NullableContract {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableContractStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableContractStream) Limit(limit int) *NullableContractStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableContractStream) Map(fn func(NullableContract, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractStream) Map2Int(fn func(NullableContract, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractStream) Map2Int32(fn func(NullableContract, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractStream) Map2Int64(fn func(NullableContract, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractStream) Map2Float32(fn func(NullableContract, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractStream) Map2Float64(fn func(NullableContract, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractStream) Map2Bool(fn func(NullableContract, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractStream) Map2Bytes(fn func(NullableContract, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractStream) Map2String(fn func(NullableContract, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractStream) Max(fn func(NullableContract, int) float64) *NullableContract {
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
func (self *NullableContractStream) Min(fn func(NullableContract, int) float64) *NullableContract {
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
func (self *NullableContractStream) NoneMatch(fn func(NullableContract, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableContractStream) Get(index int) *NullableContract {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableContractStream) GetOr(index int, arg NullableContract) NullableContract {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableContractStream) Peek(fn func(*NullableContract, int)) *NullableContractStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableContractStream) Reduce(fn func(NullableContract, NullableContract, int) NullableContract) *NullableContractStream {
	return self.ReduceInit(fn, NullableContract{})
}
func (self *NullableContractStream) ReduceInit(fn func(NullableContract, NullableContract, int) NullableContract, initialValue NullableContract) *NullableContractStream {
	result := NullableContractStreamOf()
	self.ForEach(func(v NullableContract, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableContractStream) ReduceInterface(fn func(interface{}, NullableContract, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableContract{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableContractStream) ReduceString(fn func(string, NullableContract, int) string) []string {
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
func (self *NullableContractStream) ReduceInt(fn func(int, NullableContract, int) int) []int {
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
func (self *NullableContractStream) ReduceInt32(fn func(int32, NullableContract, int) int32) []int32 {
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
func (self *NullableContractStream) ReduceInt64(fn func(int64, NullableContract, int) int64) []int64 {
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
func (self *NullableContractStream) ReduceFloat32(fn func(float32, NullableContract, int) float32) []float32 {
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
func (self *NullableContractStream) ReduceFloat64(fn func(float64, NullableContract, int) float64) []float64 {
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
func (self *NullableContractStream) ReduceBool(fn func(bool, NullableContract, int) bool) []bool {
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
func (self *NullableContractStream) Reverse() *NullableContractStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableContractStream) Replace(fn func(NullableContract, int) NullableContract) *NullableContractStream {
	return self.ForEach(func(v NullableContract, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableContractStream) Select(fn func(NullableContract) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableContractStream) Set(index int, val NullableContract) *NullableContractStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableContractStream) Skip(skip int) *NullableContractStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableContractStream) SkippingEach(fn func(NullableContract, int) int) *NullableContractStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableContractStream) Slice(startIndex, n int) *NullableContractStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableContract{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableContractStream) Sort(fn func(i, j int) bool) *NullableContractStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableContractStream) Tail() *NullableContract {
	return self.Last()
}
func (self *NullableContractStream) TailOr(arg NullableContract) NullableContract {
	return self.LastOr(arg)
}
func (self *NullableContractStream) ToList() []NullableContract {
	return self.Val()
}
func (self *NullableContractStream) Unique() *NullableContractStream {
	return self.Distinct()
}
func (self *NullableContractStream) Val() []NullableContract {
	if self == nil {
		return []NullableContract{}
	}
	return *self.Copy()
}
func (self *NullableContractStream) While(fn func(NullableContract, int) bool) *NullableContractStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableContractStream) Where(fn func(NullableContract) bool) *NullableContractStream {
	result := NullableContractStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableContractStream) WhereSlim(fn func(NullableContract) bool) *NullableContractStream {
	result := NullableContractStreamOf()
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
