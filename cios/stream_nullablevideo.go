/*
 * Collection utility of NullableVideo Struct
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

type NullableVideoStream []NullableVideo

func NullableVideoStreamOf(arg ...NullableVideo) NullableVideoStream {
	return arg
}
func NullableVideoStreamFrom(arg []NullableVideo) NullableVideoStream {
	return arg
}
func CreateNullableVideoStream(arg ...NullableVideo) *NullableVideoStream {
	tmp := NullableVideoStreamOf(arg...)
	return &tmp
}
func GenerateNullableVideoStream(arg []NullableVideo) *NullableVideoStream {
	tmp := NullableVideoStreamFrom(arg)
	return &tmp
}

func (self *NullableVideoStream) Add(arg NullableVideo) *NullableVideoStream {
	return self.AddAll(arg)
}
func (self *NullableVideoStream) AddAll(arg ...NullableVideo) *NullableVideoStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableVideoStream) AddSafe(arg *NullableVideo) *NullableVideoStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableVideoStream) Aggregate(fn func(NullableVideo, NullableVideo) NullableVideo) *NullableVideoStream {
	result := NullableVideoStreamOf()
	self.ForEach(func(v NullableVideo, i int) {
		if i == 0 {
			result.Add(fn(NullableVideo{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableVideoStream) AllMatch(fn func(NullableVideo, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableVideoStream) AnyMatch(fn func(NullableVideo, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableVideoStream) Clone() *NullableVideoStream {
	temp := make([]NullableVideo, self.Len())
	copy(temp, *self)
	return (*NullableVideoStream)(&temp)
}
func (self *NullableVideoStream) Copy() *NullableVideoStream {
	return self.Clone()
}
func (self *NullableVideoStream) Concat(arg []NullableVideo) *NullableVideoStream {
	return self.AddAll(arg...)
}
func (self *NullableVideoStream) Contains(arg NullableVideo) bool {
	return self.FindIndex(func(_arg NullableVideo, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableVideoStream) Clean() *NullableVideoStream {
	*self = NullableVideoStreamOf()
	return self
}
func (self *NullableVideoStream) Delete(index int) *NullableVideoStream {
	return self.DeleteRange(index, index)
}
func (self *NullableVideoStream) DeleteRange(startIndex, endIndex int) *NullableVideoStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableVideoStream) Distinct() *NullableVideoStream {
	caches := map[string]bool{}
	result := NullableVideoStreamOf()
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
func (self *NullableVideoStream) Each(fn func(NullableVideo)) *NullableVideoStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableVideoStream) EachRight(fn func(NullableVideo)) *NullableVideoStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableVideoStream) Equals(arr []NullableVideo) bool {
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
func (self *NullableVideoStream) Filter(fn func(NullableVideo, int) bool) *NullableVideoStream {
	result := NullableVideoStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableVideoStream) FilterSlim(fn func(NullableVideo, int) bool) *NullableVideoStream {
	result := NullableVideoStreamOf()
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
func (self *NullableVideoStream) Find(fn func(NullableVideo, int) bool) *NullableVideo {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableVideoStream) FindOr(fn func(NullableVideo, int) bool, or NullableVideo) NullableVideo {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableVideoStream) FindIndex(fn func(NullableVideo, int) bool) int {
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
func (self *NullableVideoStream) First() *NullableVideo {
	return self.Get(0)
}
func (self *NullableVideoStream) FirstOr(arg NullableVideo) NullableVideo {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableVideoStream) ForEach(fn func(NullableVideo, int)) *NullableVideoStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableVideoStream) ForEachRight(fn func(NullableVideo, int)) *NullableVideoStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableVideoStream) GroupBy(fn func(NullableVideo, int) string) map[string][]NullableVideo {
	m := map[string][]NullableVideo{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableVideoStream) GroupByValues(fn func(NullableVideo, int) string) [][]NullableVideo {
	var tmp [][]NullableVideo
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableVideoStream) IndexOf(arg NullableVideo) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableVideoStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableVideoStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableVideoStream) Last() *NullableVideo {
	return self.Get(self.Len() - 1)
}
func (self *NullableVideoStream) LastOr(arg NullableVideo) NullableVideo {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableVideoStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableVideoStream) Limit(limit int) *NullableVideoStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableVideoStream) Map(fn func(NullableVideo, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableVideoStream) Map2Int(fn func(NullableVideo, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableVideoStream) Map2Int32(fn func(NullableVideo, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableVideoStream) Map2Int64(fn func(NullableVideo, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableVideoStream) Map2Float32(fn func(NullableVideo, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableVideoStream) Map2Float64(fn func(NullableVideo, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableVideoStream) Map2Bool(fn func(NullableVideo, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableVideoStream) Map2Bytes(fn func(NullableVideo, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableVideoStream) Map2String(fn func(NullableVideo, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableVideoStream) Max(fn func(NullableVideo, int) float64) *NullableVideo {
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
func (self *NullableVideoStream) Min(fn func(NullableVideo, int) float64) *NullableVideo {
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
func (self *NullableVideoStream) NoneMatch(fn func(NullableVideo, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableVideoStream) Get(index int) *NullableVideo {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableVideoStream) GetOr(index int, arg NullableVideo) NullableVideo {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableVideoStream) Peek(fn func(*NullableVideo, int)) *NullableVideoStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableVideoStream) Reduce(fn func(NullableVideo, NullableVideo, int) NullableVideo) *NullableVideoStream {
	return self.ReduceInit(fn, NullableVideo{})
}
func (self *NullableVideoStream) ReduceInit(fn func(NullableVideo, NullableVideo, int) NullableVideo, initialValue NullableVideo) *NullableVideoStream {
	result := NullableVideoStreamOf()
	self.ForEach(func(v NullableVideo, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableVideoStream) ReduceInterface(fn func(interface{}, NullableVideo, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableVideo{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableVideoStream) ReduceString(fn func(string, NullableVideo, int) string) []string {
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
func (self *NullableVideoStream) ReduceInt(fn func(int, NullableVideo, int) int) []int {
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
func (self *NullableVideoStream) ReduceInt32(fn func(int32, NullableVideo, int) int32) []int32 {
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
func (self *NullableVideoStream) ReduceInt64(fn func(int64, NullableVideo, int) int64) []int64 {
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
func (self *NullableVideoStream) ReduceFloat32(fn func(float32, NullableVideo, int) float32) []float32 {
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
func (self *NullableVideoStream) ReduceFloat64(fn func(float64, NullableVideo, int) float64) []float64 {
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
func (self *NullableVideoStream) ReduceBool(fn func(bool, NullableVideo, int) bool) []bool {
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
func (self *NullableVideoStream) Reverse() *NullableVideoStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableVideoStream) Replace(fn func(NullableVideo, int) NullableVideo) *NullableVideoStream {
	return self.ForEach(func(v NullableVideo, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableVideoStream) Select(fn func(NullableVideo) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableVideoStream) Set(index int, val NullableVideo) *NullableVideoStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableVideoStream) Skip(skip int) *NullableVideoStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableVideoStream) SkippingEach(fn func(NullableVideo, int) int) *NullableVideoStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableVideoStream) Slice(startIndex, n int) *NullableVideoStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableVideo{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableVideoStream) Sort(fn func(i, j int) bool) *NullableVideoStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableVideoStream) Tail() *NullableVideo {
	return self.Last()
}
func (self *NullableVideoStream) TailOr(arg NullableVideo) NullableVideo {
	return self.LastOr(arg)
}
func (self *NullableVideoStream) ToList() []NullableVideo {
	return self.Val()
}
func (self *NullableVideoStream) Unique() *NullableVideoStream {
	return self.Distinct()
}
func (self *NullableVideoStream) Val() []NullableVideo {
	if self == nil {
		return []NullableVideo{}
	}
	return *self.Copy()
}
func (self *NullableVideoStream) While(fn func(NullableVideo, int) bool) *NullableVideoStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableVideoStream) Where(fn func(NullableVideo) bool) *NullableVideoStream {
	result := NullableVideoStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableVideoStream) WhereSlim(fn func(NullableVideo) bool) *NullableVideoStream {
	result := NullableVideoStreamOf()
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
