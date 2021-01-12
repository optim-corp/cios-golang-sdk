/*
 * Collection utility of MultipleDeviceEntitiesComponent Struct
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

type MultipleDeviceEntitiesComponentStream []MultipleDeviceEntitiesComponent

func MultipleDeviceEntitiesComponentStreamOf(arg ...MultipleDeviceEntitiesComponent) MultipleDeviceEntitiesComponentStream {
	return arg
}
func MultipleDeviceEntitiesComponentStreamFrom(arg []MultipleDeviceEntitiesComponent) MultipleDeviceEntitiesComponentStream {
	return arg
}
func CreateMultipleDeviceEntitiesComponentStream(arg ...MultipleDeviceEntitiesComponent) *MultipleDeviceEntitiesComponentStream {
	tmp := MultipleDeviceEntitiesComponentStreamOf(arg...)
	return &tmp
}
func GenerateMultipleDeviceEntitiesComponentStream(arg []MultipleDeviceEntitiesComponent) *MultipleDeviceEntitiesComponentStream {
	tmp := MultipleDeviceEntitiesComponentStreamFrom(arg)
	return &tmp
}

func (self *MultipleDeviceEntitiesComponentStream) Add(arg MultipleDeviceEntitiesComponent) *MultipleDeviceEntitiesComponentStream {
	return self.AddAll(arg)
}
func (self *MultipleDeviceEntitiesComponentStream) AddAll(arg ...MultipleDeviceEntitiesComponent) *MultipleDeviceEntitiesComponentStream {
	*self = append(*self, arg...)
	return self
}
func (self *MultipleDeviceEntitiesComponentStream) AddSafe(arg *MultipleDeviceEntitiesComponent) *MultipleDeviceEntitiesComponentStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *MultipleDeviceEntitiesComponentStream) Aggregate(fn func(MultipleDeviceEntitiesComponent, MultipleDeviceEntitiesComponent) MultipleDeviceEntitiesComponent) *MultipleDeviceEntitiesComponentStream {
	result := MultipleDeviceEntitiesComponentStreamOf()
	self.ForEach(func(v MultipleDeviceEntitiesComponent, i int) {
		if i == 0 {
			result.Add(fn(MultipleDeviceEntitiesComponent{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *MultipleDeviceEntitiesComponentStream) AllMatch(fn func(MultipleDeviceEntitiesComponent, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *MultipleDeviceEntitiesComponentStream) AnyMatch(fn func(MultipleDeviceEntitiesComponent, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *MultipleDeviceEntitiesComponentStream) Clone() *MultipleDeviceEntitiesComponentStream {
	temp := make([]MultipleDeviceEntitiesComponent, self.Len())
	copy(temp, *self)
	return (*MultipleDeviceEntitiesComponentStream)(&temp)
}
func (self *MultipleDeviceEntitiesComponentStream) Copy() *MultipleDeviceEntitiesComponentStream {
	return self.Clone()
}
func (self *MultipleDeviceEntitiesComponentStream) Concat(arg []MultipleDeviceEntitiesComponent) *MultipleDeviceEntitiesComponentStream {
	return self.AddAll(arg...)
}
func (self *MultipleDeviceEntitiesComponentStream) Contains(arg MultipleDeviceEntitiesComponent) bool {
	return self.FindIndex(func(_arg MultipleDeviceEntitiesComponent, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *MultipleDeviceEntitiesComponentStream) Clean() *MultipleDeviceEntitiesComponentStream {
	*self = MultipleDeviceEntitiesComponentStreamOf()
	return self
}
func (self *MultipleDeviceEntitiesComponentStream) Delete(index int) *MultipleDeviceEntitiesComponentStream {
	return self.DeleteRange(index, index)
}
func (self *MultipleDeviceEntitiesComponentStream) DeleteRange(startIndex, endIndex int) *MultipleDeviceEntitiesComponentStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *MultipleDeviceEntitiesComponentStream) Distinct() *MultipleDeviceEntitiesComponentStream {
	caches := map[string]bool{}
	result := MultipleDeviceEntitiesComponentStreamOf()
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
func (self *MultipleDeviceEntitiesComponentStream) Each(fn func(MultipleDeviceEntitiesComponent)) *MultipleDeviceEntitiesComponentStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *MultipleDeviceEntitiesComponentStream) EachRight(fn func(MultipleDeviceEntitiesComponent)) *MultipleDeviceEntitiesComponentStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *MultipleDeviceEntitiesComponentStream) Equals(arr []MultipleDeviceEntitiesComponent) bool {
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
func (self *MultipleDeviceEntitiesComponentStream) Filter(fn func(MultipleDeviceEntitiesComponent, int) bool) *MultipleDeviceEntitiesComponentStream {
	result := MultipleDeviceEntitiesComponentStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleDeviceEntitiesComponentStream) FilterSlim(fn func(MultipleDeviceEntitiesComponent, int) bool) *MultipleDeviceEntitiesComponentStream {
	result := MultipleDeviceEntitiesComponentStreamOf()
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
func (self *MultipleDeviceEntitiesComponentStream) Find(fn func(MultipleDeviceEntitiesComponent, int) bool) *MultipleDeviceEntitiesComponent {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *MultipleDeviceEntitiesComponentStream) FindOr(fn func(MultipleDeviceEntitiesComponent, int) bool, or MultipleDeviceEntitiesComponent) MultipleDeviceEntitiesComponent {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *MultipleDeviceEntitiesComponentStream) FindIndex(fn func(MultipleDeviceEntitiesComponent, int) bool) int {
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
func (self *MultipleDeviceEntitiesComponentStream) First() *MultipleDeviceEntitiesComponent {
	return self.Get(0)
}
func (self *MultipleDeviceEntitiesComponentStream) FirstOr(arg MultipleDeviceEntitiesComponent) MultipleDeviceEntitiesComponent {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleDeviceEntitiesComponentStream) ForEach(fn func(MultipleDeviceEntitiesComponent, int)) *MultipleDeviceEntitiesComponentStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *MultipleDeviceEntitiesComponentStream) ForEachRight(fn func(MultipleDeviceEntitiesComponent, int)) *MultipleDeviceEntitiesComponentStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *MultipleDeviceEntitiesComponentStream) GroupBy(fn func(MultipleDeviceEntitiesComponent, int) string) map[string][]MultipleDeviceEntitiesComponent {
	m := map[string][]MultipleDeviceEntitiesComponent{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *MultipleDeviceEntitiesComponentStream) GroupByValues(fn func(MultipleDeviceEntitiesComponent, int) string) [][]MultipleDeviceEntitiesComponent {
	var tmp [][]MultipleDeviceEntitiesComponent
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *MultipleDeviceEntitiesComponentStream) IndexOf(arg MultipleDeviceEntitiesComponent) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *MultipleDeviceEntitiesComponentStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *MultipleDeviceEntitiesComponentStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *MultipleDeviceEntitiesComponentStream) Last() *MultipleDeviceEntitiesComponent {
	return self.Get(self.Len() - 1)
}
func (self *MultipleDeviceEntitiesComponentStream) LastOr(arg MultipleDeviceEntitiesComponent) MultipleDeviceEntitiesComponent {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleDeviceEntitiesComponentStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *MultipleDeviceEntitiesComponentStream) Limit(limit int) *MultipleDeviceEntitiesComponentStream {
	self.Slice(0, limit)
	return self
}

func (self *MultipleDeviceEntitiesComponentStream) Map(fn func(MultipleDeviceEntitiesComponent, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceEntitiesComponentStream) Map2Int(fn func(MultipleDeviceEntitiesComponent, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceEntitiesComponentStream) Map2Int32(fn func(MultipleDeviceEntitiesComponent, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceEntitiesComponentStream) Map2Int64(fn func(MultipleDeviceEntitiesComponent, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceEntitiesComponentStream) Map2Float32(fn func(MultipleDeviceEntitiesComponent, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceEntitiesComponentStream) Map2Float64(fn func(MultipleDeviceEntitiesComponent, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceEntitiesComponentStream) Map2Bool(fn func(MultipleDeviceEntitiesComponent, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceEntitiesComponentStream) Map2Bytes(fn func(MultipleDeviceEntitiesComponent, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceEntitiesComponentStream) Map2String(fn func(MultipleDeviceEntitiesComponent, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MultipleDeviceEntitiesComponentStream) Max(fn func(MultipleDeviceEntitiesComponent, int) float64) *MultipleDeviceEntitiesComponent {
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
func (self *MultipleDeviceEntitiesComponentStream) Min(fn func(MultipleDeviceEntitiesComponent, int) float64) *MultipleDeviceEntitiesComponent {
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
func (self *MultipleDeviceEntitiesComponentStream) NoneMatch(fn func(MultipleDeviceEntitiesComponent, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *MultipleDeviceEntitiesComponentStream) Get(index int) *MultipleDeviceEntitiesComponent {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *MultipleDeviceEntitiesComponentStream) GetOr(index int, arg MultipleDeviceEntitiesComponent) MultipleDeviceEntitiesComponent {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *MultipleDeviceEntitiesComponentStream) Peek(fn func(*MultipleDeviceEntitiesComponent, int)) *MultipleDeviceEntitiesComponentStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *MultipleDeviceEntitiesComponentStream) Reduce(fn func(MultipleDeviceEntitiesComponent, MultipleDeviceEntitiesComponent, int) MultipleDeviceEntitiesComponent) *MultipleDeviceEntitiesComponentStream {
	return self.ReduceInit(fn, MultipleDeviceEntitiesComponent{})
}
func (self *MultipleDeviceEntitiesComponentStream) ReduceInit(fn func(MultipleDeviceEntitiesComponent, MultipleDeviceEntitiesComponent, int) MultipleDeviceEntitiesComponent, initialValue MultipleDeviceEntitiesComponent) *MultipleDeviceEntitiesComponentStream {
	result := MultipleDeviceEntitiesComponentStreamOf()
	self.ForEach(func(v MultipleDeviceEntitiesComponent, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *MultipleDeviceEntitiesComponentStream) ReduceInterface(fn func(interface{}, MultipleDeviceEntitiesComponent, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(MultipleDeviceEntitiesComponent{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *MultipleDeviceEntitiesComponentStream) ReduceString(fn func(string, MultipleDeviceEntitiesComponent, int) string) []string {
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
func (self *MultipleDeviceEntitiesComponentStream) ReduceInt(fn func(int, MultipleDeviceEntitiesComponent, int) int) []int {
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
func (self *MultipleDeviceEntitiesComponentStream) ReduceInt32(fn func(int32, MultipleDeviceEntitiesComponent, int) int32) []int32 {
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
func (self *MultipleDeviceEntitiesComponentStream) ReduceInt64(fn func(int64, MultipleDeviceEntitiesComponent, int) int64) []int64 {
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
func (self *MultipleDeviceEntitiesComponentStream) ReduceFloat32(fn func(float32, MultipleDeviceEntitiesComponent, int) float32) []float32 {
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
func (self *MultipleDeviceEntitiesComponentStream) ReduceFloat64(fn func(float64, MultipleDeviceEntitiesComponent, int) float64) []float64 {
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
func (self *MultipleDeviceEntitiesComponentStream) ReduceBool(fn func(bool, MultipleDeviceEntitiesComponent, int) bool) []bool {
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
func (self *MultipleDeviceEntitiesComponentStream) Reverse() *MultipleDeviceEntitiesComponentStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *MultipleDeviceEntitiesComponentStream) Replace(fn func(MultipleDeviceEntitiesComponent, int) MultipleDeviceEntitiesComponent) *MultipleDeviceEntitiesComponentStream {
	return self.ForEach(func(v MultipleDeviceEntitiesComponent, i int) { self.Set(i, fn(v, i)) })
}
func (self *MultipleDeviceEntitiesComponentStream) Select(fn func(MultipleDeviceEntitiesComponent) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *MultipleDeviceEntitiesComponentStream) Set(index int, val MultipleDeviceEntitiesComponent) *MultipleDeviceEntitiesComponentStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *MultipleDeviceEntitiesComponentStream) Skip(skip int) *MultipleDeviceEntitiesComponentStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *MultipleDeviceEntitiesComponentStream) SkippingEach(fn func(MultipleDeviceEntitiesComponent, int) int) *MultipleDeviceEntitiesComponentStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *MultipleDeviceEntitiesComponentStream) Slice(startIndex, n int) *MultipleDeviceEntitiesComponentStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []MultipleDeviceEntitiesComponent{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *MultipleDeviceEntitiesComponentStream) Sort(fn func(i, j int) bool) *MultipleDeviceEntitiesComponentStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *MultipleDeviceEntitiesComponentStream) Tail() *MultipleDeviceEntitiesComponent {
	return self.Last()
}
func (self *MultipleDeviceEntitiesComponentStream) TailOr(arg MultipleDeviceEntitiesComponent) MultipleDeviceEntitiesComponent {
	return self.LastOr(arg)
}
func (self *MultipleDeviceEntitiesComponentStream) ToList() []MultipleDeviceEntitiesComponent {
	return self.Val()
}
func (self *MultipleDeviceEntitiesComponentStream) Unique() *MultipleDeviceEntitiesComponentStream {
	return self.Distinct()
}
func (self *MultipleDeviceEntitiesComponentStream) Val() []MultipleDeviceEntitiesComponent {
	if self == nil {
		return []MultipleDeviceEntitiesComponent{}
	}
	return *self.Copy()
}
func (self *MultipleDeviceEntitiesComponentStream) While(fn func(MultipleDeviceEntitiesComponent, int) bool) *MultipleDeviceEntitiesComponentStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *MultipleDeviceEntitiesComponentStream) Where(fn func(MultipleDeviceEntitiesComponent) bool) *MultipleDeviceEntitiesComponentStream {
	result := MultipleDeviceEntitiesComponentStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MultipleDeviceEntitiesComponentStream) WhereSlim(fn func(MultipleDeviceEntitiesComponent) bool) *MultipleDeviceEntitiesComponentStream {
	result := MultipleDeviceEntitiesComponentStreamOf()
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
