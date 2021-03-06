/*
 * Collection utility of NullableContractUserLicense Struct
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

type NullableContractUserLicenseStream []NullableContractUserLicense

func NullableContractUserLicenseStreamOf(arg ...NullableContractUserLicense) NullableContractUserLicenseStream {
	return arg
}
func NullableContractUserLicenseStreamFrom(arg []NullableContractUserLicense) NullableContractUserLicenseStream {
	return arg
}
func CreateNullableContractUserLicenseStream(arg ...NullableContractUserLicense) *NullableContractUserLicenseStream {
	tmp := NullableContractUserLicenseStreamOf(arg...)
	return &tmp
}
func GenerateNullableContractUserLicenseStream(arg []NullableContractUserLicense) *NullableContractUserLicenseStream {
	tmp := NullableContractUserLicenseStreamFrom(arg)
	return &tmp
}

func (self *NullableContractUserLicenseStream) Add(arg NullableContractUserLicense) *NullableContractUserLicenseStream {
	return self.AddAll(arg)
}
func (self *NullableContractUserLicenseStream) AddAll(arg ...NullableContractUserLicense) *NullableContractUserLicenseStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableContractUserLicenseStream) AddSafe(arg *NullableContractUserLicense) *NullableContractUserLicenseStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableContractUserLicenseStream) Aggregate(fn func(NullableContractUserLicense, NullableContractUserLicense) NullableContractUserLicense) *NullableContractUserLicenseStream {
	result := NullableContractUserLicenseStreamOf()
	self.ForEach(func(v NullableContractUserLicense, i int) {
		if i == 0 {
			result.Add(fn(NullableContractUserLicense{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableContractUserLicenseStream) AllMatch(fn func(NullableContractUserLicense, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableContractUserLicenseStream) AnyMatch(fn func(NullableContractUserLicense, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableContractUserLicenseStream) Clone() *NullableContractUserLicenseStream {
	temp := make([]NullableContractUserLicense, self.Len())
	copy(temp, *self)
	return (*NullableContractUserLicenseStream)(&temp)
}
func (self *NullableContractUserLicenseStream) Copy() *NullableContractUserLicenseStream {
	return self.Clone()
}
func (self *NullableContractUserLicenseStream) Concat(arg []NullableContractUserLicense) *NullableContractUserLicenseStream {
	return self.AddAll(arg...)
}
func (self *NullableContractUserLicenseStream) Contains(arg NullableContractUserLicense) bool {
	return self.FindIndex(func(_arg NullableContractUserLicense, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableContractUserLicenseStream) Clean() *NullableContractUserLicenseStream {
	*self = NullableContractUserLicenseStreamOf()
	return self
}
func (self *NullableContractUserLicenseStream) Delete(index int) *NullableContractUserLicenseStream {
	return self.DeleteRange(index, index)
}
func (self *NullableContractUserLicenseStream) DeleteRange(startIndex, endIndex int) *NullableContractUserLicenseStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableContractUserLicenseStream) Distinct() *NullableContractUserLicenseStream {
	caches := map[string]bool{}
	result := NullableContractUserLicenseStreamOf()
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
func (self *NullableContractUserLicenseStream) Each(fn func(NullableContractUserLicense)) *NullableContractUserLicenseStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableContractUserLicenseStream) EachRight(fn func(NullableContractUserLicense)) *NullableContractUserLicenseStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableContractUserLicenseStream) Equals(arr []NullableContractUserLicense) bool {
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
func (self *NullableContractUserLicenseStream) Filter(fn func(NullableContractUserLicense, int) bool) *NullableContractUserLicenseStream {
	result := NullableContractUserLicenseStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableContractUserLicenseStream) FilterSlim(fn func(NullableContractUserLicense, int) bool) *NullableContractUserLicenseStream {
	result := NullableContractUserLicenseStreamOf()
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
func (self *NullableContractUserLicenseStream) Find(fn func(NullableContractUserLicense, int) bool) *NullableContractUserLicense {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableContractUserLicenseStream) FindOr(fn func(NullableContractUserLicense, int) bool, or NullableContractUserLicense) NullableContractUserLicense {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableContractUserLicenseStream) FindIndex(fn func(NullableContractUserLicense, int) bool) int {
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
func (self *NullableContractUserLicenseStream) First() *NullableContractUserLicense {
	return self.Get(0)
}
func (self *NullableContractUserLicenseStream) FirstOr(arg NullableContractUserLicense) NullableContractUserLicense {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableContractUserLicenseStream) ForEach(fn func(NullableContractUserLicense, int)) *NullableContractUserLicenseStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableContractUserLicenseStream) ForEachRight(fn func(NullableContractUserLicense, int)) *NullableContractUserLicenseStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableContractUserLicenseStream) GroupBy(fn func(NullableContractUserLicense, int) string) map[string][]NullableContractUserLicense {
	m := map[string][]NullableContractUserLicense{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableContractUserLicenseStream) GroupByValues(fn func(NullableContractUserLicense, int) string) [][]NullableContractUserLicense {
	var tmp [][]NullableContractUserLicense
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableContractUserLicenseStream) IndexOf(arg NullableContractUserLicense) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableContractUserLicenseStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableContractUserLicenseStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableContractUserLicenseStream) Last() *NullableContractUserLicense {
	return self.Get(self.Len() - 1)
}
func (self *NullableContractUserLicenseStream) LastOr(arg NullableContractUserLicense) NullableContractUserLicense {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableContractUserLicenseStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableContractUserLicenseStream) Limit(limit int) *NullableContractUserLicenseStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableContractUserLicenseStream) Map(fn func(NullableContractUserLicense, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractUserLicenseStream) Map2Int(fn func(NullableContractUserLicense, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractUserLicenseStream) Map2Int32(fn func(NullableContractUserLicense, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractUserLicenseStream) Map2Int64(fn func(NullableContractUserLicense, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractUserLicenseStream) Map2Float32(fn func(NullableContractUserLicense, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractUserLicenseStream) Map2Float64(fn func(NullableContractUserLicense, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractUserLicenseStream) Map2Bool(fn func(NullableContractUserLicense, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractUserLicenseStream) Map2Bytes(fn func(NullableContractUserLicense, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractUserLicenseStream) Map2String(fn func(NullableContractUserLicense, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableContractUserLicenseStream) Max(fn func(NullableContractUserLicense, int) float64) *NullableContractUserLicense {
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
func (self *NullableContractUserLicenseStream) Min(fn func(NullableContractUserLicense, int) float64) *NullableContractUserLicense {
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
func (self *NullableContractUserLicenseStream) NoneMatch(fn func(NullableContractUserLicense, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableContractUserLicenseStream) Get(index int) *NullableContractUserLicense {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableContractUserLicenseStream) GetOr(index int, arg NullableContractUserLicense) NullableContractUserLicense {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableContractUserLicenseStream) Peek(fn func(*NullableContractUserLicense, int)) *NullableContractUserLicenseStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableContractUserLicenseStream) Reduce(fn func(NullableContractUserLicense, NullableContractUserLicense, int) NullableContractUserLicense) *NullableContractUserLicenseStream {
	return self.ReduceInit(fn, NullableContractUserLicense{})
}
func (self *NullableContractUserLicenseStream) ReduceInit(fn func(NullableContractUserLicense, NullableContractUserLicense, int) NullableContractUserLicense, initialValue NullableContractUserLicense) *NullableContractUserLicenseStream {
	result := NullableContractUserLicenseStreamOf()
	self.ForEach(func(v NullableContractUserLicense, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableContractUserLicenseStream) ReduceInterface(fn func(interface{}, NullableContractUserLicense, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableContractUserLicense{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableContractUserLicenseStream) ReduceString(fn func(string, NullableContractUserLicense, int) string) []string {
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
func (self *NullableContractUserLicenseStream) ReduceInt(fn func(int, NullableContractUserLicense, int) int) []int {
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
func (self *NullableContractUserLicenseStream) ReduceInt32(fn func(int32, NullableContractUserLicense, int) int32) []int32 {
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
func (self *NullableContractUserLicenseStream) ReduceInt64(fn func(int64, NullableContractUserLicense, int) int64) []int64 {
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
func (self *NullableContractUserLicenseStream) ReduceFloat32(fn func(float32, NullableContractUserLicense, int) float32) []float32 {
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
func (self *NullableContractUserLicenseStream) ReduceFloat64(fn func(float64, NullableContractUserLicense, int) float64) []float64 {
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
func (self *NullableContractUserLicenseStream) ReduceBool(fn func(bool, NullableContractUserLicense, int) bool) []bool {
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
func (self *NullableContractUserLicenseStream) Reverse() *NullableContractUserLicenseStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableContractUserLicenseStream) Replace(fn func(NullableContractUserLicense, int) NullableContractUserLicense) *NullableContractUserLicenseStream {
	return self.ForEach(func(v NullableContractUserLicense, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableContractUserLicenseStream) Select(fn func(NullableContractUserLicense) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableContractUserLicenseStream) Set(index int, val NullableContractUserLicense) *NullableContractUserLicenseStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableContractUserLicenseStream) Skip(skip int) *NullableContractUserLicenseStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableContractUserLicenseStream) SkippingEach(fn func(NullableContractUserLicense, int) int) *NullableContractUserLicenseStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableContractUserLicenseStream) Slice(startIndex, n int) *NullableContractUserLicenseStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableContractUserLicense{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableContractUserLicenseStream) Sort(fn func(i, j int) bool) *NullableContractUserLicenseStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableContractUserLicenseStream) Tail() *NullableContractUserLicense {
	return self.Last()
}
func (self *NullableContractUserLicenseStream) TailOr(arg NullableContractUserLicense) NullableContractUserLicense {
	return self.LastOr(arg)
}
func (self *NullableContractUserLicenseStream) ToList() []NullableContractUserLicense {
	return self.Val()
}
func (self *NullableContractUserLicenseStream) Unique() *NullableContractUserLicenseStream {
	return self.Distinct()
}
func (self *NullableContractUserLicenseStream) Val() []NullableContractUserLicense {
	if self == nil {
		return []NullableContractUserLicense{}
	}
	return *self.Copy()
}
func (self *NullableContractUserLicenseStream) While(fn func(NullableContractUserLicense, int) bool) *NullableContractUserLicenseStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableContractUserLicenseStream) Where(fn func(NullableContractUserLicense) bool) *NullableContractUserLicenseStream {
	result := NullableContractUserLicenseStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableContractUserLicenseStream) WhereSlim(fn func(NullableContractUserLicense) bool) *NullableContractUserLicenseStream {
	result := NullableContractUserLicenseStreamOf()
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
