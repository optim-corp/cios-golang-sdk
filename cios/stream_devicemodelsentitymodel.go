/*
 * Collection utility of DeviceModelsEntityModel Struct
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

type DeviceModelsEntityModelStream []DeviceModelsEntityModel

func DeviceModelsEntityModelStreamOf(arg ...DeviceModelsEntityModel) DeviceModelsEntityModelStream {
	return arg
}
func DeviceModelsEntityModelStreamFrom(arg []DeviceModelsEntityModel) DeviceModelsEntityModelStream {
	return arg
}
func CreateDeviceModelsEntityModelStream(arg ...DeviceModelsEntityModel) *DeviceModelsEntityModelStream {
	tmp := DeviceModelsEntityModelStreamOf(arg...)
	return &tmp
}
func GenerateDeviceModelsEntityModelStream(arg []DeviceModelsEntityModel) *DeviceModelsEntityModelStream {
	tmp := DeviceModelsEntityModelStreamFrom(arg)
	return &tmp
}

func (self *DeviceModelsEntityModelStream) Add(arg DeviceModelsEntityModel) *DeviceModelsEntityModelStream {
	return self.AddAll(arg)
}
func (self *DeviceModelsEntityModelStream) AddAll(arg ...DeviceModelsEntityModel) *DeviceModelsEntityModelStream {
	*self = append(*self, arg...)
	return self
}
func (self *DeviceModelsEntityModelStream) AddSafe(arg *DeviceModelsEntityModel) *DeviceModelsEntityModelStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *DeviceModelsEntityModelStream) Aggregate(fn func(DeviceModelsEntityModel, DeviceModelsEntityModel) DeviceModelsEntityModel) *DeviceModelsEntityModelStream {
	result := DeviceModelsEntityModelStreamOf()
	self.ForEach(func(v DeviceModelsEntityModel, i int) {
		if i == 0 {
			result.Add(fn(DeviceModelsEntityModel{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *DeviceModelsEntityModelStream) AllMatch(fn func(DeviceModelsEntityModel, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *DeviceModelsEntityModelStream) AnyMatch(fn func(DeviceModelsEntityModel, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *DeviceModelsEntityModelStream) Clone() *DeviceModelsEntityModelStream {
	temp := make([]DeviceModelsEntityModel, self.Len())
	copy(temp, *self)
	return (*DeviceModelsEntityModelStream)(&temp)
}
func (self *DeviceModelsEntityModelStream) Copy() *DeviceModelsEntityModelStream {
	return self.Clone()
}
func (self *DeviceModelsEntityModelStream) Concat(arg []DeviceModelsEntityModel) *DeviceModelsEntityModelStream {
	return self.AddAll(arg...)
}
func (self *DeviceModelsEntityModelStream) Contains(arg DeviceModelsEntityModel) bool {
	return self.FindIndex(func(_arg DeviceModelsEntityModel, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *DeviceModelsEntityModelStream) Clean() *DeviceModelsEntityModelStream {
	*self = DeviceModelsEntityModelStreamOf()
	return self
}
func (self *DeviceModelsEntityModelStream) Delete(index int) *DeviceModelsEntityModelStream {
	return self.DeleteRange(index, index)
}
func (self *DeviceModelsEntityModelStream) DeleteRange(startIndex, endIndex int) *DeviceModelsEntityModelStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *DeviceModelsEntityModelStream) Distinct() *DeviceModelsEntityModelStream {
	caches := map[string]bool{}
	result := DeviceModelsEntityModelStreamOf()
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
func (self *DeviceModelsEntityModelStream) Each(fn func(DeviceModelsEntityModel)) *DeviceModelsEntityModelStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *DeviceModelsEntityModelStream) EachRight(fn func(DeviceModelsEntityModel)) *DeviceModelsEntityModelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *DeviceModelsEntityModelStream) Equals(arr []DeviceModelsEntityModel) bool {
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
func (self *DeviceModelsEntityModelStream) Filter(fn func(DeviceModelsEntityModel, int) bool) *DeviceModelsEntityModelStream {
	result := DeviceModelsEntityModelStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *DeviceModelsEntityModelStream) FilterSlim(fn func(DeviceModelsEntityModel, int) bool) *DeviceModelsEntityModelStream {
	result := DeviceModelsEntityModelStreamOf()
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
func (self *DeviceModelsEntityModelStream) Find(fn func(DeviceModelsEntityModel, int) bool) *DeviceModelsEntityModel {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *DeviceModelsEntityModelStream) FindOr(fn func(DeviceModelsEntityModel, int) bool, or DeviceModelsEntityModel) DeviceModelsEntityModel {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *DeviceModelsEntityModelStream) FindIndex(fn func(DeviceModelsEntityModel, int) bool) int {
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
func (self *DeviceModelsEntityModelStream) First() *DeviceModelsEntityModel {
	return self.Get(0)
}
func (self *DeviceModelsEntityModelStream) FirstOr(arg DeviceModelsEntityModel) DeviceModelsEntityModel {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *DeviceModelsEntityModelStream) ForEach(fn func(DeviceModelsEntityModel, int)) *DeviceModelsEntityModelStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *DeviceModelsEntityModelStream) ForEachRight(fn func(DeviceModelsEntityModel, int)) *DeviceModelsEntityModelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *DeviceModelsEntityModelStream) GroupBy(fn func(DeviceModelsEntityModel, int) string) map[string][]DeviceModelsEntityModel {
	m := map[string][]DeviceModelsEntityModel{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *DeviceModelsEntityModelStream) GroupByValues(fn func(DeviceModelsEntityModel, int) string) [][]DeviceModelsEntityModel {
	var tmp [][]DeviceModelsEntityModel
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *DeviceModelsEntityModelStream) IndexOf(arg DeviceModelsEntityModel) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *DeviceModelsEntityModelStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *DeviceModelsEntityModelStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *DeviceModelsEntityModelStream) Last() *DeviceModelsEntityModel {
	return self.Get(self.Len() - 1)
}
func (self *DeviceModelsEntityModelStream) LastOr(arg DeviceModelsEntityModel) DeviceModelsEntityModel {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *DeviceModelsEntityModelStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *DeviceModelsEntityModelStream) Limit(limit int) *DeviceModelsEntityModelStream {
	self.Slice(0, limit)
	return self
}

func (self *DeviceModelsEntityModelStream) Map(fn func(DeviceModelsEntityModel, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceModelsEntityModelStream) Map2Int(fn func(DeviceModelsEntityModel, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceModelsEntityModelStream) Map2Int32(fn func(DeviceModelsEntityModel, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceModelsEntityModelStream) Map2Int64(fn func(DeviceModelsEntityModel, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceModelsEntityModelStream) Map2Float32(fn func(DeviceModelsEntityModel, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceModelsEntityModelStream) Map2Float64(fn func(DeviceModelsEntityModel, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceModelsEntityModelStream) Map2Bool(fn func(DeviceModelsEntityModel, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceModelsEntityModelStream) Map2Bytes(fn func(DeviceModelsEntityModel, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceModelsEntityModelStream) Map2String(fn func(DeviceModelsEntityModel, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DeviceModelsEntityModelStream) Max(fn func(DeviceModelsEntityModel, int) float64) *DeviceModelsEntityModel {
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
func (self *DeviceModelsEntityModelStream) Min(fn func(DeviceModelsEntityModel, int) float64) *DeviceModelsEntityModel {
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
func (self *DeviceModelsEntityModelStream) NoneMatch(fn func(DeviceModelsEntityModel, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *DeviceModelsEntityModelStream) Get(index int) *DeviceModelsEntityModel {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *DeviceModelsEntityModelStream) GetOr(index int, arg DeviceModelsEntityModel) DeviceModelsEntityModel {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *DeviceModelsEntityModelStream) Peek(fn func(*DeviceModelsEntityModel, int)) *DeviceModelsEntityModelStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *DeviceModelsEntityModelStream) Reduce(fn func(DeviceModelsEntityModel, DeviceModelsEntityModel, int) DeviceModelsEntityModel) *DeviceModelsEntityModelStream {
	return self.ReduceInit(fn, DeviceModelsEntityModel{})
}
func (self *DeviceModelsEntityModelStream) ReduceInit(fn func(DeviceModelsEntityModel, DeviceModelsEntityModel, int) DeviceModelsEntityModel, initialValue DeviceModelsEntityModel) *DeviceModelsEntityModelStream {
	result := DeviceModelsEntityModelStreamOf()
	self.ForEach(func(v DeviceModelsEntityModel, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *DeviceModelsEntityModelStream) ReduceInterface(fn func(interface{}, DeviceModelsEntityModel, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(DeviceModelsEntityModel{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *DeviceModelsEntityModelStream) ReduceString(fn func(string, DeviceModelsEntityModel, int) string) []string {
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
func (self *DeviceModelsEntityModelStream) ReduceInt(fn func(int, DeviceModelsEntityModel, int) int) []int {
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
func (self *DeviceModelsEntityModelStream) ReduceInt32(fn func(int32, DeviceModelsEntityModel, int) int32) []int32 {
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
func (self *DeviceModelsEntityModelStream) ReduceInt64(fn func(int64, DeviceModelsEntityModel, int) int64) []int64 {
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
func (self *DeviceModelsEntityModelStream) ReduceFloat32(fn func(float32, DeviceModelsEntityModel, int) float32) []float32 {
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
func (self *DeviceModelsEntityModelStream) ReduceFloat64(fn func(float64, DeviceModelsEntityModel, int) float64) []float64 {
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
func (self *DeviceModelsEntityModelStream) ReduceBool(fn func(bool, DeviceModelsEntityModel, int) bool) []bool {
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
func (self *DeviceModelsEntityModelStream) Reverse() *DeviceModelsEntityModelStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *DeviceModelsEntityModelStream) Replace(fn func(DeviceModelsEntityModel, int) DeviceModelsEntityModel) *DeviceModelsEntityModelStream {
	return self.ForEach(func(v DeviceModelsEntityModel, i int) { self.Set(i, fn(v, i)) })
}
func (self *DeviceModelsEntityModelStream) Select(fn func(DeviceModelsEntityModel) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *DeviceModelsEntityModelStream) Set(index int, val DeviceModelsEntityModel) *DeviceModelsEntityModelStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *DeviceModelsEntityModelStream) Skip(skip int) *DeviceModelsEntityModelStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *DeviceModelsEntityModelStream) SkippingEach(fn func(DeviceModelsEntityModel, int) int) *DeviceModelsEntityModelStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *DeviceModelsEntityModelStream) Slice(startIndex, n int) *DeviceModelsEntityModelStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []DeviceModelsEntityModel{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *DeviceModelsEntityModelStream) Sort(fn func(i, j int) bool) *DeviceModelsEntityModelStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *DeviceModelsEntityModelStream) Tail() *DeviceModelsEntityModel {
	return self.Last()
}
func (self *DeviceModelsEntityModelStream) TailOr(arg DeviceModelsEntityModel) DeviceModelsEntityModel {
	return self.LastOr(arg)
}
func (self *DeviceModelsEntityModelStream) ToList() []DeviceModelsEntityModel {
	return self.Val()
}
func (self *DeviceModelsEntityModelStream) Unique() *DeviceModelsEntityModelStream {
	return self.Distinct()
}
func (self *DeviceModelsEntityModelStream) Val() []DeviceModelsEntityModel {
	if self == nil {
		return []DeviceModelsEntityModel{}
	}
	return *self.Copy()
}
func (self *DeviceModelsEntityModelStream) While(fn func(DeviceModelsEntityModel, int) bool) *DeviceModelsEntityModelStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *DeviceModelsEntityModelStream) Where(fn func(DeviceModelsEntityModel) bool) *DeviceModelsEntityModelStream {
	result := DeviceModelsEntityModelStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *DeviceModelsEntityModelStream) WhereSlim(fn func(DeviceModelsEntityModel) bool) *DeviceModelsEntityModelStream {
	result := DeviceModelsEntityModelStreamOf()
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
