/*
 * Collection utility of NullableConnectTokenResponse Struct
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

type NullableConnectTokenResponseStream []NullableConnectTokenResponse

func NullableConnectTokenResponseStreamOf(arg ...NullableConnectTokenResponse) NullableConnectTokenResponseStream {
	return arg
}
func NullableConnectTokenResponseStreamFrom(arg []NullableConnectTokenResponse) NullableConnectTokenResponseStream {
	return arg
}
func CreateNullableConnectTokenResponseStream(arg ...NullableConnectTokenResponse) *NullableConnectTokenResponseStream {
	tmp := NullableConnectTokenResponseStreamOf(arg...)
	return &tmp
}
func GenerateNullableConnectTokenResponseStream(arg []NullableConnectTokenResponse) *NullableConnectTokenResponseStream {
	tmp := NullableConnectTokenResponseStreamFrom(arg)
	return &tmp
}

func (self *NullableConnectTokenResponseStream) Add(arg NullableConnectTokenResponse) *NullableConnectTokenResponseStream {
	return self.AddAll(arg)
}
func (self *NullableConnectTokenResponseStream) AddAll(arg ...NullableConnectTokenResponse) *NullableConnectTokenResponseStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableConnectTokenResponseStream) AddSafe(arg *NullableConnectTokenResponse) *NullableConnectTokenResponseStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableConnectTokenResponseStream) Aggregate(fn func(NullableConnectTokenResponse, NullableConnectTokenResponse) NullableConnectTokenResponse) *NullableConnectTokenResponseStream {
	result := NullableConnectTokenResponseStreamOf()
	self.ForEach(func(v NullableConnectTokenResponse, i int) {
		if i == 0 {
			result.Add(fn(NullableConnectTokenResponse{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableConnectTokenResponseStream) AllMatch(fn func(NullableConnectTokenResponse, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableConnectTokenResponseStream) AnyMatch(fn func(NullableConnectTokenResponse, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableConnectTokenResponseStream) Clone() *NullableConnectTokenResponseStream {
	temp := make([]NullableConnectTokenResponse, self.Len())
	copy(temp, *self)
	return (*NullableConnectTokenResponseStream)(&temp)
}
func (self *NullableConnectTokenResponseStream) Copy() *NullableConnectTokenResponseStream {
	return self.Clone()
}
func (self *NullableConnectTokenResponseStream) Concat(arg []NullableConnectTokenResponse) *NullableConnectTokenResponseStream {
	return self.AddAll(arg...)
}
func (self *NullableConnectTokenResponseStream) Contains(arg NullableConnectTokenResponse) bool {
	return self.FindIndex(func(_arg NullableConnectTokenResponse, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableConnectTokenResponseStream) Clean() *NullableConnectTokenResponseStream {
	*self = NullableConnectTokenResponseStreamOf()
	return self
}
func (self *NullableConnectTokenResponseStream) Delete(index int) *NullableConnectTokenResponseStream {
	return self.DeleteRange(index, index)
}
func (self *NullableConnectTokenResponseStream) DeleteRange(startIndex, endIndex int) *NullableConnectTokenResponseStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableConnectTokenResponseStream) Distinct() *NullableConnectTokenResponseStream {
	caches := map[string]bool{}
	result := NullableConnectTokenResponseStreamOf()
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
func (self *NullableConnectTokenResponseStream) Each(fn func(NullableConnectTokenResponse)) *NullableConnectTokenResponseStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableConnectTokenResponseStream) EachRight(fn func(NullableConnectTokenResponse)) *NullableConnectTokenResponseStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableConnectTokenResponseStream) Equals(arr []NullableConnectTokenResponse) bool {
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
func (self *NullableConnectTokenResponseStream) Filter(fn func(NullableConnectTokenResponse, int) bool) *NullableConnectTokenResponseStream {
	result := NullableConnectTokenResponseStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableConnectTokenResponseStream) FilterSlim(fn func(NullableConnectTokenResponse, int) bool) *NullableConnectTokenResponseStream {
	result := NullableConnectTokenResponseStreamOf()
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
func (self *NullableConnectTokenResponseStream) Find(fn func(NullableConnectTokenResponse, int) bool) *NullableConnectTokenResponse {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableConnectTokenResponseStream) FindOr(fn func(NullableConnectTokenResponse, int) bool, or NullableConnectTokenResponse) NullableConnectTokenResponse {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableConnectTokenResponseStream) FindIndex(fn func(NullableConnectTokenResponse, int) bool) int {
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
func (self *NullableConnectTokenResponseStream) First() *NullableConnectTokenResponse {
	return self.Get(0)
}
func (self *NullableConnectTokenResponseStream) FirstOr(arg NullableConnectTokenResponse) NullableConnectTokenResponse {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableConnectTokenResponseStream) ForEach(fn func(NullableConnectTokenResponse, int)) *NullableConnectTokenResponseStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableConnectTokenResponseStream) ForEachRight(fn func(NullableConnectTokenResponse, int)) *NullableConnectTokenResponseStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableConnectTokenResponseStream) GroupBy(fn func(NullableConnectTokenResponse, int) string) map[string][]NullableConnectTokenResponse {
	m := map[string][]NullableConnectTokenResponse{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableConnectTokenResponseStream) GroupByValues(fn func(NullableConnectTokenResponse, int) string) [][]NullableConnectTokenResponse {
	var tmp [][]NullableConnectTokenResponse
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableConnectTokenResponseStream) IndexOf(arg NullableConnectTokenResponse) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableConnectTokenResponseStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableConnectTokenResponseStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableConnectTokenResponseStream) Last() *NullableConnectTokenResponse {
	return self.Get(self.Len() - 1)
}
func (self *NullableConnectTokenResponseStream) LastOr(arg NullableConnectTokenResponse) NullableConnectTokenResponse {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableConnectTokenResponseStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableConnectTokenResponseStream) Limit(limit int) *NullableConnectTokenResponseStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableConnectTokenResponseStream) Map(fn func(NullableConnectTokenResponse, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableConnectTokenResponseStream) Map2Int(fn func(NullableConnectTokenResponse, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableConnectTokenResponseStream) Map2Int32(fn func(NullableConnectTokenResponse, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableConnectTokenResponseStream) Map2Int64(fn func(NullableConnectTokenResponse, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableConnectTokenResponseStream) Map2Float32(fn func(NullableConnectTokenResponse, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableConnectTokenResponseStream) Map2Float64(fn func(NullableConnectTokenResponse, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableConnectTokenResponseStream) Map2Bool(fn func(NullableConnectTokenResponse, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableConnectTokenResponseStream) Map2Bytes(fn func(NullableConnectTokenResponse, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableConnectTokenResponseStream) Map2String(fn func(NullableConnectTokenResponse, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableConnectTokenResponseStream) Max(fn func(NullableConnectTokenResponse, int) float64) *NullableConnectTokenResponse {
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
func (self *NullableConnectTokenResponseStream) Min(fn func(NullableConnectTokenResponse, int) float64) *NullableConnectTokenResponse {
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
func (self *NullableConnectTokenResponseStream) NoneMatch(fn func(NullableConnectTokenResponse, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableConnectTokenResponseStream) Get(index int) *NullableConnectTokenResponse {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableConnectTokenResponseStream) GetOr(index int, arg NullableConnectTokenResponse) NullableConnectTokenResponse {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableConnectTokenResponseStream) Peek(fn func(*NullableConnectTokenResponse, int)) *NullableConnectTokenResponseStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableConnectTokenResponseStream) Reduce(fn func(NullableConnectTokenResponse, NullableConnectTokenResponse, int) NullableConnectTokenResponse) *NullableConnectTokenResponseStream {
	return self.ReduceInit(fn, NullableConnectTokenResponse{})
}
func (self *NullableConnectTokenResponseStream) ReduceInit(fn func(NullableConnectTokenResponse, NullableConnectTokenResponse, int) NullableConnectTokenResponse, initialValue NullableConnectTokenResponse) *NullableConnectTokenResponseStream {
	result := NullableConnectTokenResponseStreamOf()
	self.ForEach(func(v NullableConnectTokenResponse, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableConnectTokenResponseStream) ReduceInterface(fn func(interface{}, NullableConnectTokenResponse, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableConnectTokenResponse{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableConnectTokenResponseStream) ReduceString(fn func(string, NullableConnectTokenResponse, int) string) []string {
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
func (self *NullableConnectTokenResponseStream) ReduceInt(fn func(int, NullableConnectTokenResponse, int) int) []int {
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
func (self *NullableConnectTokenResponseStream) ReduceInt32(fn func(int32, NullableConnectTokenResponse, int) int32) []int32 {
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
func (self *NullableConnectTokenResponseStream) ReduceInt64(fn func(int64, NullableConnectTokenResponse, int) int64) []int64 {
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
func (self *NullableConnectTokenResponseStream) ReduceFloat32(fn func(float32, NullableConnectTokenResponse, int) float32) []float32 {
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
func (self *NullableConnectTokenResponseStream) ReduceFloat64(fn func(float64, NullableConnectTokenResponse, int) float64) []float64 {
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
func (self *NullableConnectTokenResponseStream) ReduceBool(fn func(bool, NullableConnectTokenResponse, int) bool) []bool {
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
func (self *NullableConnectTokenResponseStream) Reverse() *NullableConnectTokenResponseStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableConnectTokenResponseStream) Replace(fn func(NullableConnectTokenResponse, int) NullableConnectTokenResponse) *NullableConnectTokenResponseStream {
	return self.ForEach(func(v NullableConnectTokenResponse, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableConnectTokenResponseStream) Select(fn func(NullableConnectTokenResponse) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableConnectTokenResponseStream) Set(index int, val NullableConnectTokenResponse) *NullableConnectTokenResponseStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableConnectTokenResponseStream) Skip(skip int) *NullableConnectTokenResponseStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableConnectTokenResponseStream) SkippingEach(fn func(NullableConnectTokenResponse, int) int) *NullableConnectTokenResponseStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableConnectTokenResponseStream) Slice(startIndex, n int) *NullableConnectTokenResponseStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableConnectTokenResponse{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableConnectTokenResponseStream) Sort(fn func(i, j int) bool) *NullableConnectTokenResponseStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableConnectTokenResponseStream) Tail() *NullableConnectTokenResponse {
	return self.Last()
}
func (self *NullableConnectTokenResponseStream) TailOr(arg NullableConnectTokenResponse) NullableConnectTokenResponse {
	return self.LastOr(arg)
}
func (self *NullableConnectTokenResponseStream) ToList() []NullableConnectTokenResponse {
	return self.Val()
}
func (self *NullableConnectTokenResponseStream) Unique() *NullableConnectTokenResponseStream {
	return self.Distinct()
}
func (self *NullableConnectTokenResponseStream) Val() []NullableConnectTokenResponse {
	if self == nil {
		return []NullableConnectTokenResponse{}
	}
	return *self.Copy()
}
func (self *NullableConnectTokenResponseStream) While(fn func(NullableConnectTokenResponse, int) bool) *NullableConnectTokenResponseStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableConnectTokenResponseStream) Where(fn func(NullableConnectTokenResponse) bool) *NullableConnectTokenResponseStream {
	result := NullableConnectTokenResponseStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableConnectTokenResponseStream) WhereSlim(fn func(NullableConnectTokenResponse) bool) *NullableConnectTokenResponseStream {
	result := NullableConnectTokenResponseStreamOf()
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
