/*
 * Collection utility of ContractUser Struct
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

type ContractUserStream []ContractUser

func ContractUserStreamOf(arg ...ContractUser) ContractUserStream {
	return arg
}
func ContractUserStreamFrom(arg []ContractUser) ContractUserStream {
	return arg
}
func CreateContractUserStream(arg ...ContractUser) *ContractUserStream {
	tmp := ContractUserStreamOf(arg...)
	return &tmp
}
func GenerateContractUserStream(arg []ContractUser) *ContractUserStream {
	tmp := ContractUserStreamFrom(arg)
	return &tmp
}

func (self *ContractUserStream) Add(arg ContractUser) *ContractUserStream {
	return self.AddAll(arg)
}
func (self *ContractUserStream) AddAll(arg ...ContractUser) *ContractUserStream {
	*self = append(*self, arg...)
	return self
}
func (self *ContractUserStream) AddSafe(arg *ContractUser) *ContractUserStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ContractUserStream) Aggregate(fn func(ContractUser, ContractUser) ContractUser) *ContractUserStream {
	result := ContractUserStreamOf()
	self.ForEach(func(v ContractUser, i int) {
		if i == 0 {
			result.Add(fn(ContractUser{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ContractUserStream) AllMatch(fn func(ContractUser, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ContractUserStream) AnyMatch(fn func(ContractUser, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ContractUserStream) Clone() *ContractUserStream {
	temp := make([]ContractUser, self.Len())
	copy(temp, *self)
	return (*ContractUserStream)(&temp)
}
func (self *ContractUserStream) Copy() *ContractUserStream {
	return self.Clone()
}
func (self *ContractUserStream) Concat(arg []ContractUser) *ContractUserStream {
	return self.AddAll(arg...)
}
func (self *ContractUserStream) Contains(arg ContractUser) bool {
	return self.FindIndex(func(_arg ContractUser, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ContractUserStream) Clean() *ContractUserStream {
	*self = ContractUserStreamOf()
	return self
}
func (self *ContractUserStream) Delete(index int) *ContractUserStream {
	return self.DeleteRange(index, index)
}
func (self *ContractUserStream) DeleteRange(startIndex, endIndex int) *ContractUserStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ContractUserStream) Distinct() *ContractUserStream {
	caches := map[string]bool{}
	result := ContractUserStreamOf()
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
func (self *ContractUserStream) Each(fn func(ContractUser)) *ContractUserStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ContractUserStream) EachRight(fn func(ContractUser)) *ContractUserStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ContractUserStream) Equals(arr []ContractUser) bool {
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
func (self *ContractUserStream) Filter(fn func(ContractUser, int) bool) *ContractUserStream {
	result := ContractUserStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ContractUserStream) FilterSlim(fn func(ContractUser, int) bool) *ContractUserStream {
	result := ContractUserStreamOf()
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
func (self *ContractUserStream) Find(fn func(ContractUser, int) bool) *ContractUser {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ContractUserStream) FindOr(fn func(ContractUser, int) bool, or ContractUser) ContractUser {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ContractUserStream) FindIndex(fn func(ContractUser, int) bool) int {
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
func (self *ContractUserStream) First() *ContractUser {
	return self.Get(0)
}
func (self *ContractUserStream) FirstOr(arg ContractUser) ContractUser {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ContractUserStream) ForEach(fn func(ContractUser, int)) *ContractUserStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ContractUserStream) ForEachRight(fn func(ContractUser, int)) *ContractUserStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ContractUserStream) GroupBy(fn func(ContractUser, int) string) map[string][]ContractUser {
	m := map[string][]ContractUser{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ContractUserStream) GroupByValues(fn func(ContractUser, int) string) [][]ContractUser {
	var tmp [][]ContractUser
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ContractUserStream) IndexOf(arg ContractUser) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ContractUserStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ContractUserStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ContractUserStream) Last() *ContractUser {
	return self.Get(self.Len() - 1)
}
func (self *ContractUserStream) LastOr(arg ContractUser) ContractUser {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ContractUserStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ContractUserStream) Limit(limit int) *ContractUserStream {
	self.Slice(0, limit)
	return self
}

func (self *ContractUserStream) Map(fn func(ContractUser, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ContractUserStream) Map2Int(fn func(ContractUser, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ContractUserStream) Map2Int32(fn func(ContractUser, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ContractUserStream) Map2Int64(fn func(ContractUser, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ContractUserStream) Map2Float32(fn func(ContractUser, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ContractUserStream) Map2Float64(fn func(ContractUser, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ContractUserStream) Map2Bool(fn func(ContractUser, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ContractUserStream) Map2Bytes(fn func(ContractUser, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ContractUserStream) Map2String(fn func(ContractUser, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ContractUserStream) Max(fn func(ContractUser, int) float64) *ContractUser {
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
func (self *ContractUserStream) Min(fn func(ContractUser, int) float64) *ContractUser {
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
func (self *ContractUserStream) NoneMatch(fn func(ContractUser, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ContractUserStream) Get(index int) *ContractUser {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ContractUserStream) GetOr(index int, arg ContractUser) ContractUser {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ContractUserStream) Peek(fn func(*ContractUser, int)) *ContractUserStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ContractUserStream) Reduce(fn func(ContractUser, ContractUser, int) ContractUser) *ContractUserStream {
	return self.ReduceInit(fn, ContractUser{})
}
func (self *ContractUserStream) ReduceInit(fn func(ContractUser, ContractUser, int) ContractUser, initialValue ContractUser) *ContractUserStream {
	result := ContractUserStreamOf()
	self.ForEach(func(v ContractUser, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ContractUserStream) ReduceInterface(fn func(interface{}, ContractUser, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ContractUser{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ContractUserStream) ReduceString(fn func(string, ContractUser, int) string) []string {
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
func (self *ContractUserStream) ReduceInt(fn func(int, ContractUser, int) int) []int {
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
func (self *ContractUserStream) ReduceInt32(fn func(int32, ContractUser, int) int32) []int32 {
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
func (self *ContractUserStream) ReduceInt64(fn func(int64, ContractUser, int) int64) []int64 {
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
func (self *ContractUserStream) ReduceFloat32(fn func(float32, ContractUser, int) float32) []float32 {
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
func (self *ContractUserStream) ReduceFloat64(fn func(float64, ContractUser, int) float64) []float64 {
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
func (self *ContractUserStream) ReduceBool(fn func(bool, ContractUser, int) bool) []bool {
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
func (self *ContractUserStream) Reverse() *ContractUserStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ContractUserStream) Replace(fn func(ContractUser, int) ContractUser) *ContractUserStream {
	return self.ForEach(func(v ContractUser, i int) { self.Set(i, fn(v, i)) })
}
func (self *ContractUserStream) Select(fn func(ContractUser) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ContractUserStream) Set(index int, val ContractUser) *ContractUserStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ContractUserStream) Skip(skip int) *ContractUserStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ContractUserStream) SkippingEach(fn func(ContractUser, int) int) *ContractUserStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ContractUserStream) Slice(startIndex, n int) *ContractUserStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ContractUser{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ContractUserStream) Sort(fn func(i, j int) bool) *ContractUserStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ContractUserStream) Tail() *ContractUser {
	return self.Last()
}
func (self *ContractUserStream) TailOr(arg ContractUser) ContractUser {
	return self.LastOr(arg)
}
func (self *ContractUserStream) ToList() []ContractUser {
	return self.Val()
}
func (self *ContractUserStream) Unique() *ContractUserStream {
	return self.Distinct()
}
func (self *ContractUserStream) Val() []ContractUser {
	if self == nil {
		return []ContractUser{}
	}
	return *self.Copy()
}
func (self *ContractUserStream) While(fn func(ContractUser, int) bool) *ContractUserStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ContractUserStream) Where(fn func(ContractUser) bool) *ContractUserStream {
	result := ContractUserStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ContractUserStream) WhereSlim(fn func(ContractUser) bool) *ContractUserStream {
	result := ContractUserStreamOf()
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