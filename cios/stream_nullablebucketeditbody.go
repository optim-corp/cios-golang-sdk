/*
 * Collection utility of NullableBucketEditBody Struct
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

type NullableBucketEditBodyStream []NullableBucketEditBody

func NullableBucketEditBodyStreamOf(arg ...NullableBucketEditBody) NullableBucketEditBodyStream {
	return arg
}
func NullableBucketEditBodyStreamFrom(arg []NullableBucketEditBody) NullableBucketEditBodyStream {
	return arg
}
func CreateNullableBucketEditBodyStream(arg ...NullableBucketEditBody) *NullableBucketEditBodyStream {
	tmp := NullableBucketEditBodyStreamOf(arg...)
	return &tmp
}
func GenerateNullableBucketEditBodyStream(arg []NullableBucketEditBody) *NullableBucketEditBodyStream {
	tmp := NullableBucketEditBodyStreamFrom(arg)
	return &tmp
}

func (self *NullableBucketEditBodyStream) Add(arg NullableBucketEditBody) *NullableBucketEditBodyStream {
	return self.AddAll(arg)
}
func (self *NullableBucketEditBodyStream) AddAll(arg ...NullableBucketEditBody) *NullableBucketEditBodyStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableBucketEditBodyStream) AddSafe(arg *NullableBucketEditBody) *NullableBucketEditBodyStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableBucketEditBodyStream) Aggregate(fn func(NullableBucketEditBody, NullableBucketEditBody) NullableBucketEditBody) *NullableBucketEditBodyStream {
	result := NullableBucketEditBodyStreamOf()
	self.ForEach(func(v NullableBucketEditBody, i int) {
		if i == 0 {
			result.Add(fn(NullableBucketEditBody{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableBucketEditBodyStream) AllMatch(fn func(NullableBucketEditBody, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableBucketEditBodyStream) AnyMatch(fn func(NullableBucketEditBody, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableBucketEditBodyStream) Clone() *NullableBucketEditBodyStream {
	temp := make([]NullableBucketEditBody, self.Len())
	copy(temp, *self)
	return (*NullableBucketEditBodyStream)(&temp)
}
func (self *NullableBucketEditBodyStream) Copy() *NullableBucketEditBodyStream {
	return self.Clone()
}
func (self *NullableBucketEditBodyStream) Concat(arg []NullableBucketEditBody) *NullableBucketEditBodyStream {
	return self.AddAll(arg...)
}
func (self *NullableBucketEditBodyStream) Contains(arg NullableBucketEditBody) bool {
	return self.FindIndex(func(_arg NullableBucketEditBody, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableBucketEditBodyStream) Clean() *NullableBucketEditBodyStream {
	*self = NullableBucketEditBodyStreamOf()
	return self
}
func (self *NullableBucketEditBodyStream) Delete(index int) *NullableBucketEditBodyStream {
	return self.DeleteRange(index, index)
}
func (self *NullableBucketEditBodyStream) DeleteRange(startIndex, endIndex int) *NullableBucketEditBodyStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableBucketEditBodyStream) Distinct() *NullableBucketEditBodyStream {
	caches := map[string]bool{}
	result := NullableBucketEditBodyStreamOf()
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
func (self *NullableBucketEditBodyStream) Each(fn func(NullableBucketEditBody)) *NullableBucketEditBodyStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableBucketEditBodyStream) EachRight(fn func(NullableBucketEditBody)) *NullableBucketEditBodyStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableBucketEditBodyStream) Equals(arr []NullableBucketEditBody) bool {
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
func (self *NullableBucketEditBodyStream) Filter(fn func(NullableBucketEditBody, int) bool) *NullableBucketEditBodyStream {
	result := NullableBucketEditBodyStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableBucketEditBodyStream) FilterSlim(fn func(NullableBucketEditBody, int) bool) *NullableBucketEditBodyStream {
	result := NullableBucketEditBodyStreamOf()
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
func (self *NullableBucketEditBodyStream) Find(fn func(NullableBucketEditBody, int) bool) *NullableBucketEditBody {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableBucketEditBodyStream) FindOr(fn func(NullableBucketEditBody, int) bool, or NullableBucketEditBody) NullableBucketEditBody {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableBucketEditBodyStream) FindIndex(fn func(NullableBucketEditBody, int) bool) int {
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
func (self *NullableBucketEditBodyStream) First() *NullableBucketEditBody {
	return self.Get(0)
}
func (self *NullableBucketEditBodyStream) FirstOr(arg NullableBucketEditBody) NullableBucketEditBody {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableBucketEditBodyStream) ForEach(fn func(NullableBucketEditBody, int)) *NullableBucketEditBodyStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableBucketEditBodyStream) ForEachRight(fn func(NullableBucketEditBody, int)) *NullableBucketEditBodyStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableBucketEditBodyStream) GroupBy(fn func(NullableBucketEditBody, int) string) map[string][]NullableBucketEditBody {
	m := map[string][]NullableBucketEditBody{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableBucketEditBodyStream) GroupByValues(fn func(NullableBucketEditBody, int) string) [][]NullableBucketEditBody {
	var tmp [][]NullableBucketEditBody
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableBucketEditBodyStream) IndexOf(arg NullableBucketEditBody) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableBucketEditBodyStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableBucketEditBodyStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableBucketEditBodyStream) Last() *NullableBucketEditBody {
	return self.Get(self.Len() - 1)
}
func (self *NullableBucketEditBodyStream) LastOr(arg NullableBucketEditBody) NullableBucketEditBody {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableBucketEditBodyStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableBucketEditBodyStream) Limit(limit int) *NullableBucketEditBodyStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableBucketEditBodyStream) Map(fn func(NullableBucketEditBody, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableBucketEditBodyStream) Map2Int(fn func(NullableBucketEditBody, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableBucketEditBodyStream) Map2Int32(fn func(NullableBucketEditBody, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableBucketEditBodyStream) Map2Int64(fn func(NullableBucketEditBody, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableBucketEditBodyStream) Map2Float32(fn func(NullableBucketEditBody, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableBucketEditBodyStream) Map2Float64(fn func(NullableBucketEditBody, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableBucketEditBodyStream) Map2Bool(fn func(NullableBucketEditBody, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableBucketEditBodyStream) Map2Bytes(fn func(NullableBucketEditBody, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableBucketEditBodyStream) Map2String(fn func(NullableBucketEditBody, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableBucketEditBodyStream) Max(fn func(NullableBucketEditBody, int) float64) *NullableBucketEditBody {
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
func (self *NullableBucketEditBodyStream) Min(fn func(NullableBucketEditBody, int) float64) *NullableBucketEditBody {
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
func (self *NullableBucketEditBodyStream) NoneMatch(fn func(NullableBucketEditBody, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableBucketEditBodyStream) Get(index int) *NullableBucketEditBody {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableBucketEditBodyStream) GetOr(index int, arg NullableBucketEditBody) NullableBucketEditBody {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableBucketEditBodyStream) Peek(fn func(*NullableBucketEditBody, int)) *NullableBucketEditBodyStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableBucketEditBodyStream) Reduce(fn func(NullableBucketEditBody, NullableBucketEditBody, int) NullableBucketEditBody) *NullableBucketEditBodyStream {
	return self.ReduceInit(fn, NullableBucketEditBody{})
}
func (self *NullableBucketEditBodyStream) ReduceInit(fn func(NullableBucketEditBody, NullableBucketEditBody, int) NullableBucketEditBody, initialValue NullableBucketEditBody) *NullableBucketEditBodyStream {
	result := NullableBucketEditBodyStreamOf()
	self.ForEach(func(v NullableBucketEditBody, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableBucketEditBodyStream) ReduceInterface(fn func(interface{}, NullableBucketEditBody, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableBucketEditBody{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableBucketEditBodyStream) ReduceString(fn func(string, NullableBucketEditBody, int) string) []string {
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
func (self *NullableBucketEditBodyStream) ReduceInt(fn func(int, NullableBucketEditBody, int) int) []int {
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
func (self *NullableBucketEditBodyStream) ReduceInt32(fn func(int32, NullableBucketEditBody, int) int32) []int32 {
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
func (self *NullableBucketEditBodyStream) ReduceInt64(fn func(int64, NullableBucketEditBody, int) int64) []int64 {
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
func (self *NullableBucketEditBodyStream) ReduceFloat32(fn func(float32, NullableBucketEditBody, int) float32) []float32 {
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
func (self *NullableBucketEditBodyStream) ReduceFloat64(fn func(float64, NullableBucketEditBody, int) float64) []float64 {
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
func (self *NullableBucketEditBodyStream) ReduceBool(fn func(bool, NullableBucketEditBody, int) bool) []bool {
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
func (self *NullableBucketEditBodyStream) Reverse() *NullableBucketEditBodyStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableBucketEditBodyStream) Replace(fn func(NullableBucketEditBody, int) NullableBucketEditBody) *NullableBucketEditBodyStream {
	return self.ForEach(func(v NullableBucketEditBody, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableBucketEditBodyStream) Select(fn func(NullableBucketEditBody) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableBucketEditBodyStream) Set(index int, val NullableBucketEditBody) *NullableBucketEditBodyStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableBucketEditBodyStream) Skip(skip int) *NullableBucketEditBodyStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableBucketEditBodyStream) SkippingEach(fn func(NullableBucketEditBody, int) int) *NullableBucketEditBodyStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableBucketEditBodyStream) Slice(startIndex, n int) *NullableBucketEditBodyStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableBucketEditBody{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableBucketEditBodyStream) Sort(fn func(i, j int) bool) *NullableBucketEditBodyStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableBucketEditBodyStream) Tail() *NullableBucketEditBody {
	return self.Last()
}
func (self *NullableBucketEditBodyStream) TailOr(arg NullableBucketEditBody) NullableBucketEditBody {
	return self.LastOr(arg)
}
func (self *NullableBucketEditBodyStream) ToList() []NullableBucketEditBody {
	return self.Val()
}
func (self *NullableBucketEditBodyStream) Unique() *NullableBucketEditBodyStream {
	return self.Distinct()
}
func (self *NullableBucketEditBodyStream) Val() []NullableBucketEditBody {
	if self == nil {
		return []NullableBucketEditBody{}
	}
	return *self.Copy()
}
func (self *NullableBucketEditBodyStream) While(fn func(NullableBucketEditBody, int) bool) *NullableBucketEditBodyStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableBucketEditBodyStream) Where(fn func(NullableBucketEditBody) bool) *NullableBucketEditBodyStream {
	result := NullableBucketEditBodyStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableBucketEditBodyStream) WhereSlim(fn func(NullableBucketEditBody) bool) *NullableBucketEditBodyStream {
	result := NullableBucketEditBodyStreamOf()
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