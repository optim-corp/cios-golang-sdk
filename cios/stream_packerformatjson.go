/*
 * Collection utility of PackerFormatJson Struct
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

type PackerFormatJsonStream []PackerFormatJson

func PackerFormatJsonStreamOf(arg ...PackerFormatJson) PackerFormatJsonStream {
	return arg
}
func PackerFormatJsonStreamFrom(arg []PackerFormatJson) PackerFormatJsonStream {
	return arg
}
func CreatePackerFormatJsonStream(arg ...PackerFormatJson) *PackerFormatJsonStream {
	tmp := PackerFormatJsonStreamOf(arg...)
	return &tmp
}
func GeneratePackerFormatJsonStream(arg []PackerFormatJson) *PackerFormatJsonStream {
	tmp := PackerFormatJsonStreamFrom(arg)
	return &tmp
}

func (self *PackerFormatJsonStream) Add(arg PackerFormatJson) *PackerFormatJsonStream {
	return self.AddAll(arg)
}
func (self *PackerFormatJsonStream) AddAll(arg ...PackerFormatJson) *PackerFormatJsonStream {
	*self = append(*self, arg...)
	return self
}
func (self *PackerFormatJsonStream) AddSafe(arg *PackerFormatJson) *PackerFormatJsonStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *PackerFormatJsonStream) Aggregate(fn func(PackerFormatJson, PackerFormatJson) PackerFormatJson) *PackerFormatJsonStream {
	result := PackerFormatJsonStreamOf()
	self.ForEach(func(v PackerFormatJson, i int) {
		if i == 0 {
			result.Add(fn(PackerFormatJson{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *PackerFormatJsonStream) AllMatch(fn func(PackerFormatJson, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *PackerFormatJsonStream) AnyMatch(fn func(PackerFormatJson, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *PackerFormatJsonStream) Clone() *PackerFormatJsonStream {
	temp := make([]PackerFormatJson, self.Len())
	copy(temp, *self)
	return (*PackerFormatJsonStream)(&temp)
}
func (self *PackerFormatJsonStream) Copy() *PackerFormatJsonStream {
	return self.Clone()
}
func (self *PackerFormatJsonStream) Concat(arg []PackerFormatJson) *PackerFormatJsonStream {
	return self.AddAll(arg...)
}
func (self *PackerFormatJsonStream) Contains(arg PackerFormatJson) bool {
	return self.FindIndex(func(_arg PackerFormatJson, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *PackerFormatJsonStream) Clean() *PackerFormatJsonStream {
	*self = PackerFormatJsonStreamOf()
	return self
}
func (self *PackerFormatJsonStream) Delete(index int) *PackerFormatJsonStream {
	return self.DeleteRange(index, index)
}
func (self *PackerFormatJsonStream) DeleteRange(startIndex, endIndex int) *PackerFormatJsonStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *PackerFormatJsonStream) Distinct() *PackerFormatJsonStream {
	caches := map[string]bool{}
	result := PackerFormatJsonStreamOf()
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
func (self *PackerFormatJsonStream) Each(fn func(PackerFormatJson)) *PackerFormatJsonStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *PackerFormatJsonStream) EachRight(fn func(PackerFormatJson)) *PackerFormatJsonStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *PackerFormatJsonStream) Equals(arr []PackerFormatJson) bool {
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
func (self *PackerFormatJsonStream) Filter(fn func(PackerFormatJson, int) bool) *PackerFormatJsonStream {
	result := PackerFormatJsonStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *PackerFormatJsonStream) FilterSlim(fn func(PackerFormatJson, int) bool) *PackerFormatJsonStream {
	result := PackerFormatJsonStreamOf()
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
func (self *PackerFormatJsonStream) Find(fn func(PackerFormatJson, int) bool) *PackerFormatJson {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *PackerFormatJsonStream) FindOr(fn func(PackerFormatJson, int) bool, or PackerFormatJson) PackerFormatJson {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *PackerFormatJsonStream) FindIndex(fn func(PackerFormatJson, int) bool) int {
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
func (self *PackerFormatJsonStream) First() *PackerFormatJson {
	return self.Get(0)
}
func (self *PackerFormatJsonStream) FirstOr(arg PackerFormatJson) PackerFormatJson {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *PackerFormatJsonStream) ForEach(fn func(PackerFormatJson, int)) *PackerFormatJsonStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *PackerFormatJsonStream) ForEachRight(fn func(PackerFormatJson, int)) *PackerFormatJsonStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *PackerFormatJsonStream) GroupBy(fn func(PackerFormatJson, int) string) map[string][]PackerFormatJson {
	m := map[string][]PackerFormatJson{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *PackerFormatJsonStream) GroupByValues(fn func(PackerFormatJson, int) string) [][]PackerFormatJson {
	var tmp [][]PackerFormatJson
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *PackerFormatJsonStream) IndexOf(arg PackerFormatJson) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *PackerFormatJsonStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *PackerFormatJsonStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *PackerFormatJsonStream) Last() *PackerFormatJson {
	return self.Get(self.Len() - 1)
}
func (self *PackerFormatJsonStream) LastOr(arg PackerFormatJson) PackerFormatJson {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *PackerFormatJsonStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *PackerFormatJsonStream) Limit(limit int) *PackerFormatJsonStream {
	self.Slice(0, limit)
	return self
}

func (self *PackerFormatJsonStream) Map(fn func(PackerFormatJson, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *PackerFormatJsonStream) Map2Int(fn func(PackerFormatJson, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *PackerFormatJsonStream) Map2Int32(fn func(PackerFormatJson, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *PackerFormatJsonStream) Map2Int64(fn func(PackerFormatJson, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *PackerFormatJsonStream) Map2Float32(fn func(PackerFormatJson, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *PackerFormatJsonStream) Map2Float64(fn func(PackerFormatJson, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *PackerFormatJsonStream) Map2Bool(fn func(PackerFormatJson, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *PackerFormatJsonStream) Map2Bytes(fn func(PackerFormatJson, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *PackerFormatJsonStream) Map2String(fn func(PackerFormatJson, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *PackerFormatJsonStream) Max(fn func(PackerFormatJson, int) float64) *PackerFormatJson {
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
func (self *PackerFormatJsonStream) Min(fn func(PackerFormatJson, int) float64) *PackerFormatJson {
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
func (self *PackerFormatJsonStream) NoneMatch(fn func(PackerFormatJson, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *PackerFormatJsonStream) Get(index int) *PackerFormatJson {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *PackerFormatJsonStream) GetOr(index int, arg PackerFormatJson) PackerFormatJson {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *PackerFormatJsonStream) Peek(fn func(*PackerFormatJson, int)) *PackerFormatJsonStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *PackerFormatJsonStream) Reduce(fn func(PackerFormatJson, PackerFormatJson, int) PackerFormatJson) *PackerFormatJsonStream {
	return self.ReduceInit(fn, PackerFormatJson{})
}
func (self *PackerFormatJsonStream) ReduceInit(fn func(PackerFormatJson, PackerFormatJson, int) PackerFormatJson, initialValue PackerFormatJson) *PackerFormatJsonStream {
	result := PackerFormatJsonStreamOf()
	self.ForEach(func(v PackerFormatJson, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *PackerFormatJsonStream) ReduceInterface(fn func(interface{}, PackerFormatJson, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(PackerFormatJson{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *PackerFormatJsonStream) ReduceString(fn func(string, PackerFormatJson, int) string) []string {
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
func (self *PackerFormatJsonStream) ReduceInt(fn func(int, PackerFormatJson, int) int) []int {
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
func (self *PackerFormatJsonStream) ReduceInt32(fn func(int32, PackerFormatJson, int) int32) []int32 {
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
func (self *PackerFormatJsonStream) ReduceInt64(fn func(int64, PackerFormatJson, int) int64) []int64 {
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
func (self *PackerFormatJsonStream) ReduceFloat32(fn func(float32, PackerFormatJson, int) float32) []float32 {
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
func (self *PackerFormatJsonStream) ReduceFloat64(fn func(float64, PackerFormatJson, int) float64) []float64 {
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
func (self *PackerFormatJsonStream) ReduceBool(fn func(bool, PackerFormatJson, int) bool) []bool {
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
func (self *PackerFormatJsonStream) Reverse() *PackerFormatJsonStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *PackerFormatJsonStream) Replace(fn func(PackerFormatJson, int) PackerFormatJson) *PackerFormatJsonStream {
	return self.ForEach(func(v PackerFormatJson, i int) { self.Set(i, fn(v, i)) })
}
func (self *PackerFormatJsonStream) Select(fn func(PackerFormatJson) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *PackerFormatJsonStream) Set(index int, val PackerFormatJson) *PackerFormatJsonStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *PackerFormatJsonStream) Skip(skip int) *PackerFormatJsonStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *PackerFormatJsonStream) SkippingEach(fn func(PackerFormatJson, int) int) *PackerFormatJsonStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *PackerFormatJsonStream) Slice(startIndex, n int) *PackerFormatJsonStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []PackerFormatJson{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *PackerFormatJsonStream) Sort(fn func(i, j int) bool) *PackerFormatJsonStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *PackerFormatJsonStream) Tail() *PackerFormatJson {
	return self.Last()
}
func (self *PackerFormatJsonStream) TailOr(arg PackerFormatJson) PackerFormatJson {
	return self.LastOr(arg)
}
func (self *PackerFormatJsonStream) ToList() []PackerFormatJson {
	return self.Val()
}
func (self *PackerFormatJsonStream) Unique() *PackerFormatJsonStream {
	return self.Distinct()
}
func (self *PackerFormatJsonStream) Val() []PackerFormatJson {
	if self == nil {
		return []PackerFormatJson{}
	}
	return *self.Copy()
}
func (self *PackerFormatJsonStream) While(fn func(PackerFormatJson, int) bool) *PackerFormatJsonStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *PackerFormatJsonStream) Where(fn func(PackerFormatJson) bool) *PackerFormatJsonStream {
	result := PackerFormatJsonStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *PackerFormatJsonStream) WhereSlim(fn func(PackerFormatJson) bool) *PackerFormatJsonStream {
	result := PackerFormatJsonStreamOf()
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