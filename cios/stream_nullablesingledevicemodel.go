/*
 * Collection utility of NullableSingleDeviceModel Struct
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

type NullableSingleDeviceModelStream []NullableSingleDeviceModel

func NullableSingleDeviceModelStreamOf(arg ...NullableSingleDeviceModel) NullableSingleDeviceModelStream {
	return arg
}
func NullableSingleDeviceModelStreamFrom(arg []NullableSingleDeviceModel) NullableSingleDeviceModelStream {
	return arg
}
func CreateNullableSingleDeviceModelStream(arg ...NullableSingleDeviceModel) *NullableSingleDeviceModelStream {
	tmp := NullableSingleDeviceModelStreamOf(arg...)
	return &tmp
}
func GenerateNullableSingleDeviceModelStream(arg []NullableSingleDeviceModel) *NullableSingleDeviceModelStream {
	tmp := NullableSingleDeviceModelStreamFrom(arg)
	return &tmp
}

func (self *NullableSingleDeviceModelStream) Add(arg NullableSingleDeviceModel) *NullableSingleDeviceModelStream {
	return self.AddAll(arg)
}
func (self *NullableSingleDeviceModelStream) AddAll(arg ...NullableSingleDeviceModel) *NullableSingleDeviceModelStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableSingleDeviceModelStream) AddSafe(arg *NullableSingleDeviceModel) *NullableSingleDeviceModelStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableSingleDeviceModelStream) Aggregate(fn func(NullableSingleDeviceModel, NullableSingleDeviceModel) NullableSingleDeviceModel) *NullableSingleDeviceModelStream {
	result := NullableSingleDeviceModelStreamOf()
	self.ForEach(func(v NullableSingleDeviceModel, i int) {
		if i == 0 {
			result.Add(fn(NullableSingleDeviceModel{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableSingleDeviceModelStream) AllMatch(fn func(NullableSingleDeviceModel, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableSingleDeviceModelStream) AnyMatch(fn func(NullableSingleDeviceModel, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableSingleDeviceModelStream) Clone() *NullableSingleDeviceModelStream {
	temp := make([]NullableSingleDeviceModel, self.Len())
	copy(temp, *self)
	return (*NullableSingleDeviceModelStream)(&temp)
}
func (self *NullableSingleDeviceModelStream) Copy() *NullableSingleDeviceModelStream {
	return self.Clone()
}
func (self *NullableSingleDeviceModelStream) Concat(arg []NullableSingleDeviceModel) *NullableSingleDeviceModelStream {
	return self.AddAll(arg...)
}
func (self *NullableSingleDeviceModelStream) Contains(arg NullableSingleDeviceModel) bool {
	return self.FindIndex(func(_arg NullableSingleDeviceModel, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableSingleDeviceModelStream) Clean() *NullableSingleDeviceModelStream {
	*self = NullableSingleDeviceModelStreamOf()
	return self
}
func (self *NullableSingleDeviceModelStream) Delete(index int) *NullableSingleDeviceModelStream {
	return self.DeleteRange(index, index)
}
func (self *NullableSingleDeviceModelStream) DeleteRange(startIndex, endIndex int) *NullableSingleDeviceModelStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableSingleDeviceModelStream) Distinct() *NullableSingleDeviceModelStream {
	caches := map[string]bool{}
	result := NullableSingleDeviceModelStreamOf()
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
func (self *NullableSingleDeviceModelStream) Each(fn func(NullableSingleDeviceModel)) *NullableSingleDeviceModelStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableSingleDeviceModelStream) EachRight(fn func(NullableSingleDeviceModel)) *NullableSingleDeviceModelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableSingleDeviceModelStream) Equals(arr []NullableSingleDeviceModel) bool {
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
func (self *NullableSingleDeviceModelStream) Filter(fn func(NullableSingleDeviceModel, int) bool) *NullableSingleDeviceModelStream {
	result := NullableSingleDeviceModelStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSingleDeviceModelStream) FilterSlim(fn func(NullableSingleDeviceModel, int) bool) *NullableSingleDeviceModelStream {
	result := NullableSingleDeviceModelStreamOf()
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
func (self *NullableSingleDeviceModelStream) Find(fn func(NullableSingleDeviceModel, int) bool) *NullableSingleDeviceModel {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableSingleDeviceModelStream) FindOr(fn func(NullableSingleDeviceModel, int) bool, or NullableSingleDeviceModel) NullableSingleDeviceModel {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableSingleDeviceModelStream) FindIndex(fn func(NullableSingleDeviceModel, int) bool) int {
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
func (self *NullableSingleDeviceModelStream) First() *NullableSingleDeviceModel {
	return self.Get(0)
}
func (self *NullableSingleDeviceModelStream) FirstOr(arg NullableSingleDeviceModel) NullableSingleDeviceModel {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleDeviceModelStream) ForEach(fn func(NullableSingleDeviceModel, int)) *NullableSingleDeviceModelStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableSingleDeviceModelStream) ForEachRight(fn func(NullableSingleDeviceModel, int)) *NullableSingleDeviceModelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableSingleDeviceModelStream) GroupBy(fn func(NullableSingleDeviceModel, int) string) map[string][]NullableSingleDeviceModel {
	m := map[string][]NullableSingleDeviceModel{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableSingleDeviceModelStream) GroupByValues(fn func(NullableSingleDeviceModel, int) string) [][]NullableSingleDeviceModel {
	var tmp [][]NullableSingleDeviceModel
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableSingleDeviceModelStream) IndexOf(arg NullableSingleDeviceModel) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableSingleDeviceModelStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableSingleDeviceModelStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableSingleDeviceModelStream) Last() *NullableSingleDeviceModel {
	return self.Get(self.Len() - 1)
}
func (self *NullableSingleDeviceModelStream) LastOr(arg NullableSingleDeviceModel) NullableSingleDeviceModel {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleDeviceModelStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableSingleDeviceModelStream) Limit(limit int) *NullableSingleDeviceModelStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableSingleDeviceModelStream) Map(fn func(NullableSingleDeviceModel, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceModelStream) Map2Int(fn func(NullableSingleDeviceModel, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceModelStream) Map2Int32(fn func(NullableSingleDeviceModel, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceModelStream) Map2Int64(fn func(NullableSingleDeviceModel, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceModelStream) Map2Float32(fn func(NullableSingleDeviceModel, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceModelStream) Map2Float64(fn func(NullableSingleDeviceModel, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceModelStream) Map2Bool(fn func(NullableSingleDeviceModel, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceModelStream) Map2Bytes(fn func(NullableSingleDeviceModel, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceModelStream) Map2String(fn func(NullableSingleDeviceModel, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceModelStream) Max(fn func(NullableSingleDeviceModel, int) float64) *NullableSingleDeviceModel {
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
func (self *NullableSingleDeviceModelStream) Min(fn func(NullableSingleDeviceModel, int) float64) *NullableSingleDeviceModel {
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
func (self *NullableSingleDeviceModelStream) NoneMatch(fn func(NullableSingleDeviceModel, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableSingleDeviceModelStream) Get(index int) *NullableSingleDeviceModel {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableSingleDeviceModelStream) GetOr(index int, arg NullableSingleDeviceModel) NullableSingleDeviceModel {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleDeviceModelStream) Peek(fn func(*NullableSingleDeviceModel, int)) *NullableSingleDeviceModelStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableSingleDeviceModelStream) Reduce(fn func(NullableSingleDeviceModel, NullableSingleDeviceModel, int) NullableSingleDeviceModel) *NullableSingleDeviceModelStream {
	return self.ReduceInit(fn, NullableSingleDeviceModel{})
}
func (self *NullableSingleDeviceModelStream) ReduceInit(fn func(NullableSingleDeviceModel, NullableSingleDeviceModel, int) NullableSingleDeviceModel, initialValue NullableSingleDeviceModel) *NullableSingleDeviceModelStream {
	result := NullableSingleDeviceModelStreamOf()
	self.ForEach(func(v NullableSingleDeviceModel, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableSingleDeviceModelStream) ReduceInterface(fn func(interface{}, NullableSingleDeviceModel, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableSingleDeviceModel{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableSingleDeviceModelStream) ReduceString(fn func(string, NullableSingleDeviceModel, int) string) []string {
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
func (self *NullableSingleDeviceModelStream) ReduceInt(fn func(int, NullableSingleDeviceModel, int) int) []int {
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
func (self *NullableSingleDeviceModelStream) ReduceInt32(fn func(int32, NullableSingleDeviceModel, int) int32) []int32 {
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
func (self *NullableSingleDeviceModelStream) ReduceInt64(fn func(int64, NullableSingleDeviceModel, int) int64) []int64 {
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
func (self *NullableSingleDeviceModelStream) ReduceFloat32(fn func(float32, NullableSingleDeviceModel, int) float32) []float32 {
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
func (self *NullableSingleDeviceModelStream) ReduceFloat64(fn func(float64, NullableSingleDeviceModel, int) float64) []float64 {
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
func (self *NullableSingleDeviceModelStream) ReduceBool(fn func(bool, NullableSingleDeviceModel, int) bool) []bool {
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
func (self *NullableSingleDeviceModelStream) Reverse() *NullableSingleDeviceModelStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableSingleDeviceModelStream) Replace(fn func(NullableSingleDeviceModel, int) NullableSingleDeviceModel) *NullableSingleDeviceModelStream {
	return self.ForEach(func(v NullableSingleDeviceModel, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableSingleDeviceModelStream) Select(fn func(NullableSingleDeviceModel) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableSingleDeviceModelStream) Set(index int, val NullableSingleDeviceModel) *NullableSingleDeviceModelStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableSingleDeviceModelStream) Skip(skip int) *NullableSingleDeviceModelStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableSingleDeviceModelStream) SkippingEach(fn func(NullableSingleDeviceModel, int) int) *NullableSingleDeviceModelStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableSingleDeviceModelStream) Slice(startIndex, n int) *NullableSingleDeviceModelStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableSingleDeviceModel{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableSingleDeviceModelStream) Sort(fn func(i, j int) bool) *NullableSingleDeviceModelStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableSingleDeviceModelStream) Tail() *NullableSingleDeviceModel {
	return self.Last()
}
func (self *NullableSingleDeviceModelStream) TailOr(arg NullableSingleDeviceModel) NullableSingleDeviceModel {
	return self.LastOr(arg)
}
func (self *NullableSingleDeviceModelStream) ToList() []NullableSingleDeviceModel {
	return self.Val()
}
func (self *NullableSingleDeviceModelStream) Unique() *NullableSingleDeviceModelStream {
	return self.Distinct()
}
func (self *NullableSingleDeviceModelStream) Val() []NullableSingleDeviceModel {
	if self == nil {
		return []NullableSingleDeviceModel{}
	}
	return *self.Copy()
}
func (self *NullableSingleDeviceModelStream) While(fn func(NullableSingleDeviceModel, int) bool) *NullableSingleDeviceModelStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableSingleDeviceModelStream) Where(fn func(NullableSingleDeviceModel) bool) *NullableSingleDeviceModelStream {
	result := NullableSingleDeviceModelStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSingleDeviceModelStream) WhereSlim(fn func(NullableSingleDeviceModel) bool) *NullableSingleDeviceModelStream {
	result := NullableSingleDeviceModelStreamOf()
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
