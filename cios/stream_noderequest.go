/*
 * Collection utility of NodeRequest Struct
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

type NodeRequestStream []NodeRequest

func NodeRequestStreamOf(arg ...NodeRequest) NodeRequestStream {
	return arg
}
func NodeRequestStreamFrom(arg []NodeRequest) NodeRequestStream {
	return arg
}
func CreateNodeRequestStream(arg ...NodeRequest) *NodeRequestStream {
	tmp := NodeRequestStreamOf(arg...)
	return &tmp
}
func GenerateNodeRequestStream(arg []NodeRequest) *NodeRequestStream {
	tmp := NodeRequestStreamFrom(arg)
	return &tmp
}

func (self *NodeRequestStream) Add(arg NodeRequest) *NodeRequestStream {
	return self.AddAll(arg)
}
func (self *NodeRequestStream) AddAll(arg ...NodeRequest) *NodeRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *NodeRequestStream) AddSafe(arg *NodeRequest) *NodeRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NodeRequestStream) Aggregate(fn func(NodeRequest, NodeRequest) NodeRequest) *NodeRequestStream {
	result := NodeRequestStreamOf()
	self.ForEach(func(v NodeRequest, i int) {
		if i == 0 {
			result.Add(fn(NodeRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NodeRequestStream) AllMatch(fn func(NodeRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NodeRequestStream) AnyMatch(fn func(NodeRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NodeRequestStream) Clone() *NodeRequestStream {
	temp := make([]NodeRequest, self.Len())
	copy(temp, *self)
	return (*NodeRequestStream)(&temp)
}
func (self *NodeRequestStream) Copy() *NodeRequestStream {
	return self.Clone()
}
func (self *NodeRequestStream) Concat(arg []NodeRequest) *NodeRequestStream {
	return self.AddAll(arg...)
}
func (self *NodeRequestStream) Contains(arg NodeRequest) bool {
	return self.FindIndex(func(_arg NodeRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NodeRequestStream) Clean() *NodeRequestStream {
	*self = NodeRequestStreamOf()
	return self
}
func (self *NodeRequestStream) Delete(index int) *NodeRequestStream {
	return self.DeleteRange(index, index)
}
func (self *NodeRequestStream) DeleteRange(startIndex, endIndex int) *NodeRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NodeRequestStream) Distinct() *NodeRequestStream {
	caches := map[string]bool{}
	result := NodeRequestStreamOf()
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
func (self *NodeRequestStream) Each(fn func(NodeRequest)) *NodeRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NodeRequestStream) EachRight(fn func(NodeRequest)) *NodeRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NodeRequestStream) Equals(arr []NodeRequest) bool {
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
func (self *NodeRequestStream) Filter(fn func(NodeRequest, int) bool) *NodeRequestStream {
	result := NodeRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NodeRequestStream) FilterSlim(fn func(NodeRequest, int) bool) *NodeRequestStream {
	result := NodeRequestStreamOf()
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
func (self *NodeRequestStream) Find(fn func(NodeRequest, int) bool) *NodeRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NodeRequestStream) FindOr(fn func(NodeRequest, int) bool, or NodeRequest) NodeRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NodeRequestStream) FindIndex(fn func(NodeRequest, int) bool) int {
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
func (self *NodeRequestStream) First() *NodeRequest {
	return self.Get(0)
}
func (self *NodeRequestStream) FirstOr(arg NodeRequest) NodeRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NodeRequestStream) ForEach(fn func(NodeRequest, int)) *NodeRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NodeRequestStream) ForEachRight(fn func(NodeRequest, int)) *NodeRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NodeRequestStream) GroupBy(fn func(NodeRequest, int) string) map[string][]NodeRequest {
	m := map[string][]NodeRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NodeRequestStream) GroupByValues(fn func(NodeRequest, int) string) [][]NodeRequest {
	var tmp [][]NodeRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NodeRequestStream) IndexOf(arg NodeRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NodeRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NodeRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NodeRequestStream) Last() *NodeRequest {
	return self.Get(self.Len() - 1)
}
func (self *NodeRequestStream) LastOr(arg NodeRequest) NodeRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NodeRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NodeRequestStream) Limit(limit int) *NodeRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *NodeRequestStream) Map(fn func(NodeRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NodeRequestStream) Map2Int(fn func(NodeRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NodeRequestStream) Map2Int32(fn func(NodeRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NodeRequestStream) Map2Int64(fn func(NodeRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NodeRequestStream) Map2Float32(fn func(NodeRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NodeRequestStream) Map2Float64(fn func(NodeRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NodeRequestStream) Map2Bool(fn func(NodeRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NodeRequestStream) Map2Bytes(fn func(NodeRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NodeRequestStream) Map2String(fn func(NodeRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NodeRequestStream) Max(fn func(NodeRequest, int) float64) *NodeRequest {
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
func (self *NodeRequestStream) Min(fn func(NodeRequest, int) float64) *NodeRequest {
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
func (self *NodeRequestStream) NoneMatch(fn func(NodeRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NodeRequestStream) Get(index int) *NodeRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NodeRequestStream) GetOr(index int, arg NodeRequest) NodeRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NodeRequestStream) Peek(fn func(*NodeRequest, int)) *NodeRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NodeRequestStream) Reduce(fn func(NodeRequest, NodeRequest, int) NodeRequest) *NodeRequestStream {
	return self.ReduceInit(fn, NodeRequest{})
}
func (self *NodeRequestStream) ReduceInit(fn func(NodeRequest, NodeRequest, int) NodeRequest, initialValue NodeRequest) *NodeRequestStream {
	result := NodeRequestStreamOf()
	self.ForEach(func(v NodeRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NodeRequestStream) ReduceInterface(fn func(interface{}, NodeRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NodeRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NodeRequestStream) ReduceString(fn func(string, NodeRequest, int) string) []string {
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
func (self *NodeRequestStream) ReduceInt(fn func(int, NodeRequest, int) int) []int {
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
func (self *NodeRequestStream) ReduceInt32(fn func(int32, NodeRequest, int) int32) []int32 {
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
func (self *NodeRequestStream) ReduceInt64(fn func(int64, NodeRequest, int) int64) []int64 {
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
func (self *NodeRequestStream) ReduceFloat32(fn func(float32, NodeRequest, int) float32) []float32 {
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
func (self *NodeRequestStream) ReduceFloat64(fn func(float64, NodeRequest, int) float64) []float64 {
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
func (self *NodeRequestStream) ReduceBool(fn func(bool, NodeRequest, int) bool) []bool {
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
func (self *NodeRequestStream) Reverse() *NodeRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NodeRequestStream) Replace(fn func(NodeRequest, int) NodeRequest) *NodeRequestStream {
	return self.ForEach(func(v NodeRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *NodeRequestStream) Select(fn func(NodeRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NodeRequestStream) Set(index int, val NodeRequest) *NodeRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NodeRequestStream) Skip(skip int) *NodeRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NodeRequestStream) SkippingEach(fn func(NodeRequest, int) int) *NodeRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NodeRequestStream) Slice(startIndex, n int) *NodeRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NodeRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NodeRequestStream) Sort(fn func(i, j int) bool) *NodeRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NodeRequestStream) Tail() *NodeRequest {
	return self.Last()
}
func (self *NodeRequestStream) TailOr(arg NodeRequest) NodeRequest {
	return self.LastOr(arg)
}
func (self *NodeRequestStream) ToList() []NodeRequest {
	return self.Val()
}
func (self *NodeRequestStream) Unique() *NodeRequestStream {
	return self.Distinct()
}
func (self *NodeRequestStream) Val() []NodeRequest {
	if self == nil {
		return []NodeRequest{}
	}
	return *self.Copy()
}
func (self *NodeRequestStream) While(fn func(NodeRequest, int) bool) *NodeRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NodeRequestStream) Where(fn func(NodeRequest) bool) *NodeRequestStream {
	result := NodeRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NodeRequestStream) WhereSlim(fn func(NodeRequest) bool) *NodeRequestStream {
	result := NodeRequestStreamOf()
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
