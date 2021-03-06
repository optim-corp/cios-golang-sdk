/*
 * Collection utility of ApplicationClient Struct
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

type ApplicationClientStream []ApplicationClient

func ApplicationClientStreamOf(arg ...ApplicationClient) ApplicationClientStream {
	return arg
}
func ApplicationClientStreamFrom(arg []ApplicationClient) ApplicationClientStream {
	return arg
}
func CreateApplicationClientStream(arg ...ApplicationClient) *ApplicationClientStream {
	tmp := ApplicationClientStreamOf(arg...)
	return &tmp
}
func GenerateApplicationClientStream(arg []ApplicationClient) *ApplicationClientStream {
	tmp := ApplicationClientStreamFrom(arg)
	return &tmp
}

func (self *ApplicationClientStream) Add(arg ApplicationClient) *ApplicationClientStream {
	return self.AddAll(arg)
}
func (self *ApplicationClientStream) AddAll(arg ...ApplicationClient) *ApplicationClientStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApplicationClientStream) AddSafe(arg *ApplicationClient) *ApplicationClientStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApplicationClientStream) Aggregate(fn func(ApplicationClient, ApplicationClient) ApplicationClient) *ApplicationClientStream {
	result := ApplicationClientStreamOf()
	self.ForEach(func(v ApplicationClient, i int) {
		if i == 0 {
			result.Add(fn(ApplicationClient{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApplicationClientStream) AllMatch(fn func(ApplicationClient, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApplicationClientStream) AnyMatch(fn func(ApplicationClient, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApplicationClientStream) Clone() *ApplicationClientStream {
	temp := make([]ApplicationClient, self.Len())
	copy(temp, *self)
	return (*ApplicationClientStream)(&temp)
}
func (self *ApplicationClientStream) Copy() *ApplicationClientStream {
	return self.Clone()
}
func (self *ApplicationClientStream) Concat(arg []ApplicationClient) *ApplicationClientStream {
	return self.AddAll(arg...)
}
func (self *ApplicationClientStream) Contains(arg ApplicationClient) bool {
	return self.FindIndex(func(_arg ApplicationClient, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApplicationClientStream) Clean() *ApplicationClientStream {
	*self = ApplicationClientStreamOf()
	return self
}
func (self *ApplicationClientStream) Delete(index int) *ApplicationClientStream {
	return self.DeleteRange(index, index)
}
func (self *ApplicationClientStream) DeleteRange(startIndex, endIndex int) *ApplicationClientStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApplicationClientStream) Distinct() *ApplicationClientStream {
	caches := map[string]bool{}
	result := ApplicationClientStreamOf()
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
func (self *ApplicationClientStream) Each(fn func(ApplicationClient)) *ApplicationClientStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApplicationClientStream) EachRight(fn func(ApplicationClient)) *ApplicationClientStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApplicationClientStream) Equals(arr []ApplicationClient) bool {
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
func (self *ApplicationClientStream) Filter(fn func(ApplicationClient, int) bool) *ApplicationClientStream {
	result := ApplicationClientStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApplicationClientStream) FilterSlim(fn func(ApplicationClient, int) bool) *ApplicationClientStream {
	result := ApplicationClientStreamOf()
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
func (self *ApplicationClientStream) Find(fn func(ApplicationClient, int) bool) *ApplicationClient {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApplicationClientStream) FindOr(fn func(ApplicationClient, int) bool, or ApplicationClient) ApplicationClient {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApplicationClientStream) FindIndex(fn func(ApplicationClient, int) bool) int {
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
func (self *ApplicationClientStream) First() *ApplicationClient {
	return self.Get(0)
}
func (self *ApplicationClientStream) FirstOr(arg ApplicationClient) ApplicationClient {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApplicationClientStream) ForEach(fn func(ApplicationClient, int)) *ApplicationClientStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApplicationClientStream) ForEachRight(fn func(ApplicationClient, int)) *ApplicationClientStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApplicationClientStream) GroupBy(fn func(ApplicationClient, int) string) map[string][]ApplicationClient {
	m := map[string][]ApplicationClient{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApplicationClientStream) GroupByValues(fn func(ApplicationClient, int) string) [][]ApplicationClient {
	var tmp [][]ApplicationClient
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApplicationClientStream) IndexOf(arg ApplicationClient) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApplicationClientStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApplicationClientStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApplicationClientStream) Last() *ApplicationClient {
	return self.Get(self.Len() - 1)
}
func (self *ApplicationClientStream) LastOr(arg ApplicationClient) ApplicationClient {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApplicationClientStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApplicationClientStream) Limit(limit int) *ApplicationClientStream {
	self.Slice(0, limit)
	return self
}

func (self *ApplicationClientStream) Map(fn func(ApplicationClient, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApplicationClientStream) Map2Int(fn func(ApplicationClient, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApplicationClientStream) Map2Int32(fn func(ApplicationClient, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApplicationClientStream) Map2Int64(fn func(ApplicationClient, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApplicationClientStream) Map2Float32(fn func(ApplicationClient, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApplicationClientStream) Map2Float64(fn func(ApplicationClient, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApplicationClientStream) Map2Bool(fn func(ApplicationClient, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApplicationClientStream) Map2Bytes(fn func(ApplicationClient, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApplicationClientStream) Map2String(fn func(ApplicationClient, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApplicationClientStream) Max(fn func(ApplicationClient, int) float64) *ApplicationClient {
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
func (self *ApplicationClientStream) Min(fn func(ApplicationClient, int) float64) *ApplicationClient {
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
func (self *ApplicationClientStream) NoneMatch(fn func(ApplicationClient, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApplicationClientStream) Get(index int) *ApplicationClient {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApplicationClientStream) GetOr(index int, arg ApplicationClient) ApplicationClient {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApplicationClientStream) Peek(fn func(*ApplicationClient, int)) *ApplicationClientStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApplicationClientStream) Reduce(fn func(ApplicationClient, ApplicationClient, int) ApplicationClient) *ApplicationClientStream {
	return self.ReduceInit(fn, ApplicationClient{})
}
func (self *ApplicationClientStream) ReduceInit(fn func(ApplicationClient, ApplicationClient, int) ApplicationClient, initialValue ApplicationClient) *ApplicationClientStream {
	result := ApplicationClientStreamOf()
	self.ForEach(func(v ApplicationClient, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApplicationClientStream) ReduceInterface(fn func(interface{}, ApplicationClient, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApplicationClient{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApplicationClientStream) ReduceString(fn func(string, ApplicationClient, int) string) []string {
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
func (self *ApplicationClientStream) ReduceInt(fn func(int, ApplicationClient, int) int) []int {
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
func (self *ApplicationClientStream) ReduceInt32(fn func(int32, ApplicationClient, int) int32) []int32 {
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
func (self *ApplicationClientStream) ReduceInt64(fn func(int64, ApplicationClient, int) int64) []int64 {
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
func (self *ApplicationClientStream) ReduceFloat32(fn func(float32, ApplicationClient, int) float32) []float32 {
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
func (self *ApplicationClientStream) ReduceFloat64(fn func(float64, ApplicationClient, int) float64) []float64 {
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
func (self *ApplicationClientStream) ReduceBool(fn func(bool, ApplicationClient, int) bool) []bool {
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
func (self *ApplicationClientStream) Reverse() *ApplicationClientStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApplicationClientStream) Replace(fn func(ApplicationClient, int) ApplicationClient) *ApplicationClientStream {
	return self.ForEach(func(v ApplicationClient, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApplicationClientStream) Select(fn func(ApplicationClient) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApplicationClientStream) Set(index int, val ApplicationClient) *ApplicationClientStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApplicationClientStream) Skip(skip int) *ApplicationClientStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApplicationClientStream) SkippingEach(fn func(ApplicationClient, int) int) *ApplicationClientStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApplicationClientStream) Slice(startIndex, n int) *ApplicationClientStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApplicationClient{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApplicationClientStream) Sort(fn func(i, j int) bool) *ApplicationClientStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApplicationClientStream) Tail() *ApplicationClient {
	return self.Last()
}
func (self *ApplicationClientStream) TailOr(arg ApplicationClient) ApplicationClient {
	return self.LastOr(arg)
}
func (self *ApplicationClientStream) ToList() []ApplicationClient {
	return self.Val()
}
func (self *ApplicationClientStream) Unique() *ApplicationClientStream {
	return self.Distinct()
}
func (self *ApplicationClientStream) Val() []ApplicationClient {
	if self == nil {
		return []ApplicationClient{}
	}
	return *self.Copy()
}
func (self *ApplicationClientStream) While(fn func(ApplicationClient, int) bool) *ApplicationClientStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApplicationClientStream) Where(fn func(ApplicationClient) bool) *ApplicationClientStream {
	result := ApplicationClientStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApplicationClientStream) WhereSlim(fn func(ApplicationClient) bool) *ApplicationClientStream {
	result := ApplicationClientStreamOf()
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
