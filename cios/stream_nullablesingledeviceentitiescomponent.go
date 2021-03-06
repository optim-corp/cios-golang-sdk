/*
 * Collection utility of NullableSingleDeviceEntitiesComponent Struct
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

type NullableSingleDeviceEntitiesComponentStream []NullableSingleDeviceEntitiesComponent

func NullableSingleDeviceEntitiesComponentStreamOf(arg ...NullableSingleDeviceEntitiesComponent) NullableSingleDeviceEntitiesComponentStream {
	return arg
}
func NullableSingleDeviceEntitiesComponentStreamFrom(arg []NullableSingleDeviceEntitiesComponent) NullableSingleDeviceEntitiesComponentStream {
	return arg
}
func CreateNullableSingleDeviceEntitiesComponentStream(arg ...NullableSingleDeviceEntitiesComponent) *NullableSingleDeviceEntitiesComponentStream {
	tmp := NullableSingleDeviceEntitiesComponentStreamOf(arg...)
	return &tmp
}
func GenerateNullableSingleDeviceEntitiesComponentStream(arg []NullableSingleDeviceEntitiesComponent) *NullableSingleDeviceEntitiesComponentStream {
	tmp := NullableSingleDeviceEntitiesComponentStreamFrom(arg)
	return &tmp
}

func (self *NullableSingleDeviceEntitiesComponentStream) Add(arg NullableSingleDeviceEntitiesComponent) *NullableSingleDeviceEntitiesComponentStream {
	return self.AddAll(arg)
}
func (self *NullableSingleDeviceEntitiesComponentStream) AddAll(arg ...NullableSingleDeviceEntitiesComponent) *NullableSingleDeviceEntitiesComponentStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableSingleDeviceEntitiesComponentStream) AddSafe(arg *NullableSingleDeviceEntitiesComponent) *NullableSingleDeviceEntitiesComponentStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableSingleDeviceEntitiesComponentStream) Aggregate(fn func(NullableSingleDeviceEntitiesComponent, NullableSingleDeviceEntitiesComponent) NullableSingleDeviceEntitiesComponent) *NullableSingleDeviceEntitiesComponentStream {
	result := NullableSingleDeviceEntitiesComponentStreamOf()
	self.ForEach(func(v NullableSingleDeviceEntitiesComponent, i int) {
		if i == 0 {
			result.Add(fn(NullableSingleDeviceEntitiesComponent{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableSingleDeviceEntitiesComponentStream) AllMatch(fn func(NullableSingleDeviceEntitiesComponent, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableSingleDeviceEntitiesComponentStream) AnyMatch(fn func(NullableSingleDeviceEntitiesComponent, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableSingleDeviceEntitiesComponentStream) Clone() *NullableSingleDeviceEntitiesComponentStream {
	temp := make([]NullableSingleDeviceEntitiesComponent, self.Len())
	copy(temp, *self)
	return (*NullableSingleDeviceEntitiesComponentStream)(&temp)
}
func (self *NullableSingleDeviceEntitiesComponentStream) Copy() *NullableSingleDeviceEntitiesComponentStream {
	return self.Clone()
}
func (self *NullableSingleDeviceEntitiesComponentStream) Concat(arg []NullableSingleDeviceEntitiesComponent) *NullableSingleDeviceEntitiesComponentStream {
	return self.AddAll(arg...)
}
func (self *NullableSingleDeviceEntitiesComponentStream) Contains(arg NullableSingleDeviceEntitiesComponent) bool {
	return self.FindIndex(func(_arg NullableSingleDeviceEntitiesComponent, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableSingleDeviceEntitiesComponentStream) Clean() *NullableSingleDeviceEntitiesComponentStream {
	*self = NullableSingleDeviceEntitiesComponentStreamOf()
	return self
}
func (self *NullableSingleDeviceEntitiesComponentStream) Delete(index int) *NullableSingleDeviceEntitiesComponentStream {
	return self.DeleteRange(index, index)
}
func (self *NullableSingleDeviceEntitiesComponentStream) DeleteRange(startIndex, endIndex int) *NullableSingleDeviceEntitiesComponentStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableSingleDeviceEntitiesComponentStream) Distinct() *NullableSingleDeviceEntitiesComponentStream {
	caches := map[string]bool{}
	result := NullableSingleDeviceEntitiesComponentStreamOf()
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
func (self *NullableSingleDeviceEntitiesComponentStream) Each(fn func(NullableSingleDeviceEntitiesComponent)) *NullableSingleDeviceEntitiesComponentStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableSingleDeviceEntitiesComponentStream) EachRight(fn func(NullableSingleDeviceEntitiesComponent)) *NullableSingleDeviceEntitiesComponentStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableSingleDeviceEntitiesComponentStream) Equals(arr []NullableSingleDeviceEntitiesComponent) bool {
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
func (self *NullableSingleDeviceEntitiesComponentStream) Filter(fn func(NullableSingleDeviceEntitiesComponent, int) bool) *NullableSingleDeviceEntitiesComponentStream {
	result := NullableSingleDeviceEntitiesComponentStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSingleDeviceEntitiesComponentStream) FilterSlim(fn func(NullableSingleDeviceEntitiesComponent, int) bool) *NullableSingleDeviceEntitiesComponentStream {
	result := NullableSingleDeviceEntitiesComponentStreamOf()
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
func (self *NullableSingleDeviceEntitiesComponentStream) Find(fn func(NullableSingleDeviceEntitiesComponent, int) bool) *NullableSingleDeviceEntitiesComponent {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableSingleDeviceEntitiesComponentStream) FindOr(fn func(NullableSingleDeviceEntitiesComponent, int) bool, or NullableSingleDeviceEntitiesComponent) NullableSingleDeviceEntitiesComponent {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableSingleDeviceEntitiesComponentStream) FindIndex(fn func(NullableSingleDeviceEntitiesComponent, int) bool) int {
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
func (self *NullableSingleDeviceEntitiesComponentStream) First() *NullableSingleDeviceEntitiesComponent {
	return self.Get(0)
}
func (self *NullableSingleDeviceEntitiesComponentStream) FirstOr(arg NullableSingleDeviceEntitiesComponent) NullableSingleDeviceEntitiesComponent {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleDeviceEntitiesComponentStream) ForEach(fn func(NullableSingleDeviceEntitiesComponent, int)) *NullableSingleDeviceEntitiesComponentStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableSingleDeviceEntitiesComponentStream) ForEachRight(fn func(NullableSingleDeviceEntitiesComponent, int)) *NullableSingleDeviceEntitiesComponentStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableSingleDeviceEntitiesComponentStream) GroupBy(fn func(NullableSingleDeviceEntitiesComponent, int) string) map[string][]NullableSingleDeviceEntitiesComponent {
	m := map[string][]NullableSingleDeviceEntitiesComponent{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableSingleDeviceEntitiesComponentStream) GroupByValues(fn func(NullableSingleDeviceEntitiesComponent, int) string) [][]NullableSingleDeviceEntitiesComponent {
	var tmp [][]NullableSingleDeviceEntitiesComponent
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableSingleDeviceEntitiesComponentStream) IndexOf(arg NullableSingleDeviceEntitiesComponent) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableSingleDeviceEntitiesComponentStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableSingleDeviceEntitiesComponentStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableSingleDeviceEntitiesComponentStream) Last() *NullableSingleDeviceEntitiesComponent {
	return self.Get(self.Len() - 1)
}
func (self *NullableSingleDeviceEntitiesComponentStream) LastOr(arg NullableSingleDeviceEntitiesComponent) NullableSingleDeviceEntitiesComponent {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleDeviceEntitiesComponentStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableSingleDeviceEntitiesComponentStream) Limit(limit int) *NullableSingleDeviceEntitiesComponentStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableSingleDeviceEntitiesComponentStream) Map(fn func(NullableSingleDeviceEntitiesComponent, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceEntitiesComponentStream) Map2Int(fn func(NullableSingleDeviceEntitiesComponent, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceEntitiesComponentStream) Map2Int32(fn func(NullableSingleDeviceEntitiesComponent, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceEntitiesComponentStream) Map2Int64(fn func(NullableSingleDeviceEntitiesComponent, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceEntitiesComponentStream) Map2Float32(fn func(NullableSingleDeviceEntitiesComponent, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceEntitiesComponentStream) Map2Float64(fn func(NullableSingleDeviceEntitiesComponent, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceEntitiesComponentStream) Map2Bool(fn func(NullableSingleDeviceEntitiesComponent, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceEntitiesComponentStream) Map2Bytes(fn func(NullableSingleDeviceEntitiesComponent, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceEntitiesComponentStream) Map2String(fn func(NullableSingleDeviceEntitiesComponent, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableSingleDeviceEntitiesComponentStream) Max(fn func(NullableSingleDeviceEntitiesComponent, int) float64) *NullableSingleDeviceEntitiesComponent {
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
func (self *NullableSingleDeviceEntitiesComponentStream) Min(fn func(NullableSingleDeviceEntitiesComponent, int) float64) *NullableSingleDeviceEntitiesComponent {
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
func (self *NullableSingleDeviceEntitiesComponentStream) NoneMatch(fn func(NullableSingleDeviceEntitiesComponent, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableSingleDeviceEntitiesComponentStream) Get(index int) *NullableSingleDeviceEntitiesComponent {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableSingleDeviceEntitiesComponentStream) GetOr(index int, arg NullableSingleDeviceEntitiesComponent) NullableSingleDeviceEntitiesComponent {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableSingleDeviceEntitiesComponentStream) Peek(fn func(*NullableSingleDeviceEntitiesComponent, int)) *NullableSingleDeviceEntitiesComponentStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableSingleDeviceEntitiesComponentStream) Reduce(fn func(NullableSingleDeviceEntitiesComponent, NullableSingleDeviceEntitiesComponent, int) NullableSingleDeviceEntitiesComponent) *NullableSingleDeviceEntitiesComponentStream {
	return self.ReduceInit(fn, NullableSingleDeviceEntitiesComponent{})
}
func (self *NullableSingleDeviceEntitiesComponentStream) ReduceInit(fn func(NullableSingleDeviceEntitiesComponent, NullableSingleDeviceEntitiesComponent, int) NullableSingleDeviceEntitiesComponent, initialValue NullableSingleDeviceEntitiesComponent) *NullableSingleDeviceEntitiesComponentStream {
	result := NullableSingleDeviceEntitiesComponentStreamOf()
	self.ForEach(func(v NullableSingleDeviceEntitiesComponent, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableSingleDeviceEntitiesComponentStream) ReduceInterface(fn func(interface{}, NullableSingleDeviceEntitiesComponent, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableSingleDeviceEntitiesComponent{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableSingleDeviceEntitiesComponentStream) ReduceString(fn func(string, NullableSingleDeviceEntitiesComponent, int) string) []string {
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
func (self *NullableSingleDeviceEntitiesComponentStream) ReduceInt(fn func(int, NullableSingleDeviceEntitiesComponent, int) int) []int {
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
func (self *NullableSingleDeviceEntitiesComponentStream) ReduceInt32(fn func(int32, NullableSingleDeviceEntitiesComponent, int) int32) []int32 {
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
func (self *NullableSingleDeviceEntitiesComponentStream) ReduceInt64(fn func(int64, NullableSingleDeviceEntitiesComponent, int) int64) []int64 {
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
func (self *NullableSingleDeviceEntitiesComponentStream) ReduceFloat32(fn func(float32, NullableSingleDeviceEntitiesComponent, int) float32) []float32 {
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
func (self *NullableSingleDeviceEntitiesComponentStream) ReduceFloat64(fn func(float64, NullableSingleDeviceEntitiesComponent, int) float64) []float64 {
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
func (self *NullableSingleDeviceEntitiesComponentStream) ReduceBool(fn func(bool, NullableSingleDeviceEntitiesComponent, int) bool) []bool {
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
func (self *NullableSingleDeviceEntitiesComponentStream) Reverse() *NullableSingleDeviceEntitiesComponentStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableSingleDeviceEntitiesComponentStream) Replace(fn func(NullableSingleDeviceEntitiesComponent, int) NullableSingleDeviceEntitiesComponent) *NullableSingleDeviceEntitiesComponentStream {
	return self.ForEach(func(v NullableSingleDeviceEntitiesComponent, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableSingleDeviceEntitiesComponentStream) Select(fn func(NullableSingleDeviceEntitiesComponent) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableSingleDeviceEntitiesComponentStream) Set(index int, val NullableSingleDeviceEntitiesComponent) *NullableSingleDeviceEntitiesComponentStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableSingleDeviceEntitiesComponentStream) Skip(skip int) *NullableSingleDeviceEntitiesComponentStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableSingleDeviceEntitiesComponentStream) SkippingEach(fn func(NullableSingleDeviceEntitiesComponent, int) int) *NullableSingleDeviceEntitiesComponentStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableSingleDeviceEntitiesComponentStream) Slice(startIndex, n int) *NullableSingleDeviceEntitiesComponentStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableSingleDeviceEntitiesComponent{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableSingleDeviceEntitiesComponentStream) Sort(fn func(i, j int) bool) *NullableSingleDeviceEntitiesComponentStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableSingleDeviceEntitiesComponentStream) Tail() *NullableSingleDeviceEntitiesComponent {
	return self.Last()
}
func (self *NullableSingleDeviceEntitiesComponentStream) TailOr(arg NullableSingleDeviceEntitiesComponent) NullableSingleDeviceEntitiesComponent {
	return self.LastOr(arg)
}
func (self *NullableSingleDeviceEntitiesComponentStream) ToList() []NullableSingleDeviceEntitiesComponent {
	return self.Val()
}
func (self *NullableSingleDeviceEntitiesComponentStream) Unique() *NullableSingleDeviceEntitiesComponentStream {
	return self.Distinct()
}
func (self *NullableSingleDeviceEntitiesComponentStream) Val() []NullableSingleDeviceEntitiesComponent {
	if self == nil {
		return []NullableSingleDeviceEntitiesComponent{}
	}
	return *self.Copy()
}
func (self *NullableSingleDeviceEntitiesComponentStream) While(fn func(NullableSingleDeviceEntitiesComponent, int) bool) *NullableSingleDeviceEntitiesComponentStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableSingleDeviceEntitiesComponentStream) Where(fn func(NullableSingleDeviceEntitiesComponent) bool) *NullableSingleDeviceEntitiesComponentStream {
	result := NullableSingleDeviceEntitiesComponentStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableSingleDeviceEntitiesComponentStream) WhereSlim(fn func(NullableSingleDeviceEntitiesComponent) bool) *NullableSingleDeviceEntitiesComponentStream {
	result := NullableSingleDeviceEntitiesComponentStreamOf()
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
