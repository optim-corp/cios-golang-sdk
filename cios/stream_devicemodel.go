/*
 * Collection utility of DeviceModel Struct
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

type DeviceModelStream []DeviceModel

func DeviceModelStreamOf(arg ...DeviceModel) DeviceModelStream {
	return arg
}
func DeviceModelStreamFrom(arg []DeviceModel) DeviceModelStream {
	return arg
}
func CreateDeviceModelStream(arg ...DeviceModel) *DeviceModelStream {
	tmp := DeviceModelStreamOf(arg...)
	return &tmp
}
func GenerateDeviceModelStream(arg []DeviceModel) *DeviceModelStream {
	tmp := DeviceModelStreamFrom(arg)
	return &tmp
}

func (self *DeviceModelStream) Add(arg DeviceModel) *DeviceModelStream {
	return self.AddAll(arg)
}
func (self *DeviceModelStream) AddAll(arg ...DeviceModel) *DeviceModelStream {
	*self = append(*self, arg...)
	return self
}
func (self *DeviceModelStream) AddSafe(arg *DeviceModel) *DeviceModelStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *DeviceModelStream) Aggregate(fn func(DeviceModel, DeviceModel) DeviceModel) *DeviceModelStream {
	result := DeviceModelStreamOf()
	self.ForEach(func(v DeviceModel, i int) {
		if i == 0 {
			result.Add(fn(DeviceModel{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *DeviceModelStream) AllMatch(fn func(DeviceModel, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *DeviceModelStream) AnyMatch(fn func(DeviceModel, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *DeviceModelStream) Clone() *DeviceModelStream {
	temp := make([]DeviceModel, self.Len())
	copy(temp, *self)
	return (*DeviceModelStream)(&temp)
}
func (self *DeviceModelStream) Copy() *DeviceModelStream {
	return self.Clone()
}
func (self *DeviceModelStream) Concat(arg []DeviceModel) *DeviceModelStream {
	return self.AddAll(arg...)
}
func (self *DeviceModelStream) Contains(arg DeviceModel) bool {
	return self.FindIndex(func(_arg DeviceModel, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *DeviceModelStream) Clean() *DeviceModelStream {
	*self = DeviceModelStreamOf()
	return self
}
func (self *DeviceModelStream) Delete(index int) *DeviceModelStream {
	return self.DeleteRange(index, index)
}
func (self *DeviceModelStream) DeleteRange(startIndex, endIndex int) *DeviceModelStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *DeviceModelStream) Distinct() *DeviceModelStream {
	caches := map[string]bool{}
	result := DeviceModelStreamOf()
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
func (self *DeviceModelStream) Each(fn func(DeviceModel)) *DeviceModelStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *DeviceModelStream) EachRight(fn func(DeviceModel)) *DeviceModelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *DeviceModelStream) Equals(arr []DeviceModel) bool {
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
func (self *DeviceModelStream) Filter(fn func(DeviceModel, int) bool) *DeviceModelStream {
	result := DeviceModelStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *DeviceModelStream) FilterSlim(fn func(DeviceModel, int) bool) *DeviceModelStream {
	result := DeviceModelStreamOf()
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
func (self *DeviceModelStream) Find(fn func(DeviceModel, int) bool) *DeviceModel {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *DeviceModelStream) FindOr(fn func(DeviceModel, int) bool, or DeviceModel) DeviceModel {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *DeviceModelStream) FindIndex(fn func(DeviceModel, int) bool) int {
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
func (self *DeviceModelStream) First() *DeviceModel {
	return self.Get(0)
}
func (self *DeviceModelStream) FirstOr(arg DeviceModel) DeviceModel {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *DeviceModelStream) ForEach(fn func(DeviceModel, int)) *DeviceModelStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *DeviceModelStream) ForEachRight(fn func(DeviceModel, int)) *DeviceModelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *DeviceModelStream) GroupBy(fn func(DeviceModel, int) string) map[string][]DeviceModel {
	m := map[string][]DeviceModel{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *DeviceModelStream) GroupByValues(fn func(DeviceModel, int) string) [][]DeviceModel {
	var tmp [][]DeviceModel
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *DeviceModelStream) IndexOf(arg DeviceModel) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *DeviceModelStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *DeviceModelStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *DeviceModelStream) Last() *DeviceModel {
	return self.Get(self.Len() - 1)
}
func (self *DeviceModelStream) LastOr(arg DeviceModel) DeviceModel {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *DeviceModelStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *DeviceModelStream) Limit(limit int) *DeviceModelStream {
	self.Slice(0, limit)
	return self
}

func (self *DeviceModelStream) Map(fn func(DeviceModel, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceModelStream) Map2Int(fn func(DeviceModel, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceModelStream) Map2Int32(fn func(DeviceModel, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceModelStream) Map2Int64(fn func(DeviceModel, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceModelStream) Map2Float32(fn func(DeviceModel, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceModelStream) Map2Float64(fn func(DeviceModel, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceModelStream) Map2Bool(fn func(DeviceModel, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceModelStream) Map2Bytes(fn func(DeviceModel, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceModelStream) Map2String(fn func(DeviceModel, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceModelStream) Max(fn func(DeviceModel, int) float64) *DeviceModel {
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
func (self *DeviceModelStream) Min(fn func(DeviceModel, int) float64) *DeviceModel {
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
func (self *DeviceModelStream) NoneMatch(fn func(DeviceModel, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *DeviceModelStream) Get(index int) *DeviceModel {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *DeviceModelStream) GetOr(index int, arg DeviceModel) DeviceModel {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *DeviceModelStream) Peek(fn func(*DeviceModel, int)) *DeviceModelStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *DeviceModelStream) Reduce(fn func(DeviceModel, DeviceModel, int) DeviceModel) *DeviceModelStream {
	return self.ReduceInit(fn, DeviceModel{})
}
func (self *DeviceModelStream) ReduceInit(fn func(DeviceModel, DeviceModel, int) DeviceModel, initialValue DeviceModel) *DeviceModelStream {
	result := DeviceModelStreamOf()
	self.ForEach(func(v DeviceModel, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *DeviceModelStream) ReduceInterface(fn func(interface{}, DeviceModel, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(DeviceModel{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *DeviceModelStream) ReduceString(fn func(string, DeviceModel, int) string) []string {
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
func (self *DeviceModelStream) ReduceInt(fn func(int, DeviceModel, int) int) []int {
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
func (self *DeviceModelStream) ReduceInt32(fn func(int32, DeviceModel, int) int32) []int32 {
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
func (self *DeviceModelStream) ReduceInt64(fn func(int64, DeviceModel, int) int64) []int64 {
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
func (self *DeviceModelStream) ReduceFloat32(fn func(float32, DeviceModel, int) float32) []float32 {
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
func (self *DeviceModelStream) ReduceFloat64(fn func(float64, DeviceModel, int) float64) []float64 {
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
func (self *DeviceModelStream) ReduceBool(fn func(bool, DeviceModel, int) bool) []bool {
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
func (self *DeviceModelStream) Reverse() *DeviceModelStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *DeviceModelStream) Replace(fn func(DeviceModel, int) DeviceModel) *DeviceModelStream {
	return self.ForEach(func(v DeviceModel, i int) { self.Set(i, fn(v, i)) })
}
func (self *DeviceModelStream) Select(fn func(DeviceModel) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *DeviceModelStream) Set(index int, val DeviceModel) *DeviceModelStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *DeviceModelStream) Skip(skip int) *DeviceModelStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *DeviceModelStream) SkippingEach(fn func(DeviceModel, int) int) *DeviceModelStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *DeviceModelStream) Slice(startIndex, n int) *DeviceModelStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []DeviceModel{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *DeviceModelStream) Sort(fn func(i, j int) bool) *DeviceModelStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *DeviceModelStream) Tail() *DeviceModel {
	return self.Last()
}
func (self *DeviceModelStream) TailOr(arg DeviceModel) DeviceModel {
	return self.LastOr(arg)
}
func (self *DeviceModelStream) ToList() []DeviceModel {
	return self.Val()
}
func (self *DeviceModelStream) Unique() *DeviceModelStream {
	return self.Distinct()
}
func (self *DeviceModelStream) Val() []DeviceModel {
	if self == nil {
		return []DeviceModel{}
	}
	return *self.Copy()
}
func (self *DeviceModelStream) While(fn func(DeviceModel, int) bool) *DeviceModelStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *DeviceModelStream) Where(fn func(DeviceModel) bool) *DeviceModelStream {
	result := DeviceModelStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *DeviceModelStream) WhereSlim(fn func(DeviceModel) bool) *DeviceModelStream {
	result := DeviceModelStreamOf()
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
