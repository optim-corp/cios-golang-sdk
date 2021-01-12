/*
 * Collection utility of GroupChildren Struct
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

type GroupChildrenStream []GroupChildren

func GroupChildrenStreamOf(arg ...GroupChildren) GroupChildrenStream {
	return arg
}
func GroupChildrenStreamFrom(arg []GroupChildren) GroupChildrenStream {
	return arg
}
func CreateGroupChildrenStream(arg ...GroupChildren) *GroupChildrenStream {
	tmp := GroupChildrenStreamOf(arg...)
	return &tmp
}
func GenerateGroupChildrenStream(arg []GroupChildren) *GroupChildrenStream {
	tmp := GroupChildrenStreamFrom(arg)
	return &tmp
}

func (self *GroupChildrenStream) Add(arg GroupChildren) *GroupChildrenStream {
	return self.AddAll(arg)
}
func (self *GroupChildrenStream) AddAll(arg ...GroupChildren) *GroupChildrenStream {
	*self = append(*self, arg...)
	return self
}
func (self *GroupChildrenStream) AddSafe(arg *GroupChildren) *GroupChildrenStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *GroupChildrenStream) Aggregate(fn func(GroupChildren, GroupChildren) GroupChildren) *GroupChildrenStream {
	result := GroupChildrenStreamOf()
	self.ForEach(func(v GroupChildren, i int) {
		if i == 0 {
			result.Add(fn(GroupChildren{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *GroupChildrenStream) AllMatch(fn func(GroupChildren, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *GroupChildrenStream) AnyMatch(fn func(GroupChildren, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *GroupChildrenStream) Clone() *GroupChildrenStream {
	temp := make([]GroupChildren, self.Len())
	copy(temp, *self)
	return (*GroupChildrenStream)(&temp)
}
func (self *GroupChildrenStream) Copy() *GroupChildrenStream {
	return self.Clone()
}
func (self *GroupChildrenStream) Concat(arg []GroupChildren) *GroupChildrenStream {
	return self.AddAll(arg...)
}
func (self *GroupChildrenStream) Contains(arg GroupChildren) bool {
	return self.FindIndex(func(_arg GroupChildren, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *GroupChildrenStream) Clean() *GroupChildrenStream {
	*self = GroupChildrenStreamOf()
	return self
}
func (self *GroupChildrenStream) Delete(index int) *GroupChildrenStream {
	return self.DeleteRange(index, index)
}
func (self *GroupChildrenStream) DeleteRange(startIndex, endIndex int) *GroupChildrenStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *GroupChildrenStream) Distinct() *GroupChildrenStream {
	caches := map[string]bool{}
	result := GroupChildrenStreamOf()
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
func (self *GroupChildrenStream) Each(fn func(GroupChildren)) *GroupChildrenStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *GroupChildrenStream) EachRight(fn func(GroupChildren)) *GroupChildrenStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *GroupChildrenStream) Equals(arr []GroupChildren) bool {
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
func (self *GroupChildrenStream) Filter(fn func(GroupChildren, int) bool) *GroupChildrenStream {
	result := GroupChildrenStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *GroupChildrenStream) FilterSlim(fn func(GroupChildren, int) bool) *GroupChildrenStream {
	result := GroupChildrenStreamOf()
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
func (self *GroupChildrenStream) Find(fn func(GroupChildren, int) bool) *GroupChildren {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *GroupChildrenStream) FindOr(fn func(GroupChildren, int) bool, or GroupChildren) GroupChildren {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *GroupChildrenStream) FindIndex(fn func(GroupChildren, int) bool) int {
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
func (self *GroupChildrenStream) First() *GroupChildren {
	return self.Get(0)
}
func (self *GroupChildrenStream) FirstOr(arg GroupChildren) GroupChildren {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *GroupChildrenStream) ForEach(fn func(GroupChildren, int)) *GroupChildrenStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *GroupChildrenStream) ForEachRight(fn func(GroupChildren, int)) *GroupChildrenStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *GroupChildrenStream) GroupBy(fn func(GroupChildren, int) string) map[string][]GroupChildren {
	m := map[string][]GroupChildren{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *GroupChildrenStream) GroupByValues(fn func(GroupChildren, int) string) [][]GroupChildren {
	var tmp [][]GroupChildren
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *GroupChildrenStream) IndexOf(arg GroupChildren) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *GroupChildrenStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *GroupChildrenStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *GroupChildrenStream) Last() *GroupChildren {
	return self.Get(self.Len() - 1)
}
func (self *GroupChildrenStream) LastOr(arg GroupChildren) GroupChildren {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *GroupChildrenStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *GroupChildrenStream) Limit(limit int) *GroupChildrenStream {
	self.Slice(0, limit)
	return self
}

func (self *GroupChildrenStream) Map(fn func(GroupChildren, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *GroupChildrenStream) Map2Int(fn func(GroupChildren, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *GroupChildrenStream) Map2Int32(fn func(GroupChildren, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *GroupChildrenStream) Map2Int64(fn func(GroupChildren, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *GroupChildrenStream) Map2Float32(fn func(GroupChildren, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *GroupChildrenStream) Map2Float64(fn func(GroupChildren, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *GroupChildrenStream) Map2Bool(fn func(GroupChildren, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *GroupChildrenStream) Map2Bytes(fn func(GroupChildren, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *GroupChildrenStream) Map2String(fn func(GroupChildren, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *GroupChildrenStream) Max(fn func(GroupChildren, int) float64) *GroupChildren {
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
func (self *GroupChildrenStream) Min(fn func(GroupChildren, int) float64) *GroupChildren {
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
func (self *GroupChildrenStream) NoneMatch(fn func(GroupChildren, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *GroupChildrenStream) Get(index int) *GroupChildren {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *GroupChildrenStream) GetOr(index int, arg GroupChildren) GroupChildren {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *GroupChildrenStream) Peek(fn func(*GroupChildren, int)) *GroupChildrenStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *GroupChildrenStream) Reduce(fn func(GroupChildren, GroupChildren, int) GroupChildren) *GroupChildrenStream {
	return self.ReduceInit(fn, GroupChildren{})
}
func (self *GroupChildrenStream) ReduceInit(fn func(GroupChildren, GroupChildren, int) GroupChildren, initialValue GroupChildren) *GroupChildrenStream {
	result := GroupChildrenStreamOf()
	self.ForEach(func(v GroupChildren, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *GroupChildrenStream) ReduceInterface(fn func(interface{}, GroupChildren, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(GroupChildren{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *GroupChildrenStream) ReduceString(fn func(string, GroupChildren, int) string) []string {
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
func (self *GroupChildrenStream) ReduceInt(fn func(int, GroupChildren, int) int) []int {
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
func (self *GroupChildrenStream) ReduceInt32(fn func(int32, GroupChildren, int) int32) []int32 {
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
func (self *GroupChildrenStream) ReduceInt64(fn func(int64, GroupChildren, int) int64) []int64 {
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
func (self *GroupChildrenStream) ReduceFloat32(fn func(float32, GroupChildren, int) float32) []float32 {
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
func (self *GroupChildrenStream) ReduceFloat64(fn func(float64, GroupChildren, int) float64) []float64 {
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
func (self *GroupChildrenStream) ReduceBool(fn func(bool, GroupChildren, int) bool) []bool {
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
func (self *GroupChildrenStream) Reverse() *GroupChildrenStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *GroupChildrenStream) Replace(fn func(GroupChildren, int) GroupChildren) *GroupChildrenStream {
	return self.ForEach(func(v GroupChildren, i int) { self.Set(i, fn(v, i)) })
}
func (self *GroupChildrenStream) Select(fn func(GroupChildren) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *GroupChildrenStream) Set(index int, val GroupChildren) *GroupChildrenStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *GroupChildrenStream) Skip(skip int) *GroupChildrenStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *GroupChildrenStream) SkippingEach(fn func(GroupChildren, int) int) *GroupChildrenStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *GroupChildrenStream) Slice(startIndex, n int) *GroupChildrenStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []GroupChildren{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *GroupChildrenStream) Sort(fn func(i, j int) bool) *GroupChildrenStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *GroupChildrenStream) Tail() *GroupChildren {
	return self.Last()
}
func (self *GroupChildrenStream) TailOr(arg GroupChildren) GroupChildren {
	return self.LastOr(arg)
}
func (self *GroupChildrenStream) ToList() []GroupChildren {
	return self.Val()
}
func (self *GroupChildrenStream) Unique() *GroupChildrenStream {
	return self.Distinct()
}
func (self *GroupChildrenStream) Val() []GroupChildren {
	if self == nil {
		return []GroupChildren{}
	}
	return *self.Copy()
}
func (self *GroupChildrenStream) While(fn func(GroupChildren, int) bool) *GroupChildrenStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *GroupChildrenStream) Where(fn func(GroupChildren) bool) *GroupChildrenStream {
	result := GroupChildrenStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *GroupChildrenStream) WhereSlim(fn func(GroupChildren) bool) *GroupChildrenStream {
	result := GroupChildrenStreamOf()
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
