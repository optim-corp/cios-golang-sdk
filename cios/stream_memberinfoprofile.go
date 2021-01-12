/*
 * Collection utility of MemberInfoProfile Struct
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

type MemberInfoProfileStream []MemberInfoProfile

func MemberInfoProfileStreamOf(arg ...MemberInfoProfile) MemberInfoProfileStream {
	return arg
}
func MemberInfoProfileStreamFrom(arg []MemberInfoProfile) MemberInfoProfileStream {
	return arg
}
func CreateMemberInfoProfileStream(arg ...MemberInfoProfile) *MemberInfoProfileStream {
	tmp := MemberInfoProfileStreamOf(arg...)
	return &tmp
}
func GenerateMemberInfoProfileStream(arg []MemberInfoProfile) *MemberInfoProfileStream {
	tmp := MemberInfoProfileStreamFrom(arg)
	return &tmp
}

func (self *MemberInfoProfileStream) Add(arg MemberInfoProfile) *MemberInfoProfileStream {
	return self.AddAll(arg)
}
func (self *MemberInfoProfileStream) AddAll(arg ...MemberInfoProfile) *MemberInfoProfileStream {
	*self = append(*self, arg...)
	return self
}
func (self *MemberInfoProfileStream) AddSafe(arg *MemberInfoProfile) *MemberInfoProfileStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *MemberInfoProfileStream) Aggregate(fn func(MemberInfoProfile, MemberInfoProfile) MemberInfoProfile) *MemberInfoProfileStream {
	result := MemberInfoProfileStreamOf()
	self.ForEach(func(v MemberInfoProfile, i int) {
		if i == 0 {
			result.Add(fn(MemberInfoProfile{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *MemberInfoProfileStream) AllMatch(fn func(MemberInfoProfile, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *MemberInfoProfileStream) AnyMatch(fn func(MemberInfoProfile, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *MemberInfoProfileStream) Clone() *MemberInfoProfileStream {
	temp := make([]MemberInfoProfile, self.Len())
	copy(temp, *self)
	return (*MemberInfoProfileStream)(&temp)
}
func (self *MemberInfoProfileStream) Copy() *MemberInfoProfileStream {
	return self.Clone()
}
func (self *MemberInfoProfileStream) Concat(arg []MemberInfoProfile) *MemberInfoProfileStream {
	return self.AddAll(arg...)
}
func (self *MemberInfoProfileStream) Contains(arg MemberInfoProfile) bool {
	return self.FindIndex(func(_arg MemberInfoProfile, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *MemberInfoProfileStream) Clean() *MemberInfoProfileStream {
	*self = MemberInfoProfileStreamOf()
	return self
}
func (self *MemberInfoProfileStream) Delete(index int) *MemberInfoProfileStream {
	return self.DeleteRange(index, index)
}
func (self *MemberInfoProfileStream) DeleteRange(startIndex, endIndex int) *MemberInfoProfileStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *MemberInfoProfileStream) Distinct() *MemberInfoProfileStream {
	caches := map[string]bool{}
	result := MemberInfoProfileStreamOf()
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
func (self *MemberInfoProfileStream) Each(fn func(MemberInfoProfile)) *MemberInfoProfileStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *MemberInfoProfileStream) EachRight(fn func(MemberInfoProfile)) *MemberInfoProfileStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *MemberInfoProfileStream) Equals(arr []MemberInfoProfile) bool {
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
func (self *MemberInfoProfileStream) Filter(fn func(MemberInfoProfile, int) bool) *MemberInfoProfileStream {
	result := MemberInfoProfileStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MemberInfoProfileStream) FilterSlim(fn func(MemberInfoProfile, int) bool) *MemberInfoProfileStream {
	result := MemberInfoProfileStreamOf()
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
func (self *MemberInfoProfileStream) Find(fn func(MemberInfoProfile, int) bool) *MemberInfoProfile {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *MemberInfoProfileStream) FindOr(fn func(MemberInfoProfile, int) bool, or MemberInfoProfile) MemberInfoProfile {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *MemberInfoProfileStream) FindIndex(fn func(MemberInfoProfile, int) bool) int {
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
func (self *MemberInfoProfileStream) First() *MemberInfoProfile {
	return self.Get(0)
}
func (self *MemberInfoProfileStream) FirstOr(arg MemberInfoProfile) MemberInfoProfile {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *MemberInfoProfileStream) ForEach(fn func(MemberInfoProfile, int)) *MemberInfoProfileStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *MemberInfoProfileStream) ForEachRight(fn func(MemberInfoProfile, int)) *MemberInfoProfileStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *MemberInfoProfileStream) GroupBy(fn func(MemberInfoProfile, int) string) map[string][]MemberInfoProfile {
	m := map[string][]MemberInfoProfile{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *MemberInfoProfileStream) GroupByValues(fn func(MemberInfoProfile, int) string) [][]MemberInfoProfile {
	var tmp [][]MemberInfoProfile
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *MemberInfoProfileStream) IndexOf(arg MemberInfoProfile) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *MemberInfoProfileStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *MemberInfoProfileStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *MemberInfoProfileStream) Last() *MemberInfoProfile {
	return self.Get(self.Len() - 1)
}
func (self *MemberInfoProfileStream) LastOr(arg MemberInfoProfile) MemberInfoProfile {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *MemberInfoProfileStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *MemberInfoProfileStream) Limit(limit int) *MemberInfoProfileStream {
	self.Slice(0, limit)
	return self
}

func (self *MemberInfoProfileStream) Map(fn func(MemberInfoProfile, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MemberInfoProfileStream) Map2Int(fn func(MemberInfoProfile, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MemberInfoProfileStream) Map2Int32(fn func(MemberInfoProfile, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MemberInfoProfileStream) Map2Int64(fn func(MemberInfoProfile, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MemberInfoProfileStream) Map2Float32(fn func(MemberInfoProfile, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MemberInfoProfileStream) Map2Float64(fn func(MemberInfoProfile, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MemberInfoProfileStream) Map2Bool(fn func(MemberInfoProfile, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MemberInfoProfileStream) Map2Bytes(fn func(MemberInfoProfile, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MemberInfoProfileStream) Map2String(fn func(MemberInfoProfile, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *MemberInfoProfileStream) Max(fn func(MemberInfoProfile, int) float64) *MemberInfoProfile {
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
func (self *MemberInfoProfileStream) Min(fn func(MemberInfoProfile, int) float64) *MemberInfoProfile {
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
func (self *MemberInfoProfileStream) NoneMatch(fn func(MemberInfoProfile, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *MemberInfoProfileStream) Get(index int) *MemberInfoProfile {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *MemberInfoProfileStream) GetOr(index int, arg MemberInfoProfile) MemberInfoProfile {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *MemberInfoProfileStream) Peek(fn func(*MemberInfoProfile, int)) *MemberInfoProfileStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *MemberInfoProfileStream) Reduce(fn func(MemberInfoProfile, MemberInfoProfile, int) MemberInfoProfile) *MemberInfoProfileStream {
	return self.ReduceInit(fn, MemberInfoProfile{})
}
func (self *MemberInfoProfileStream) ReduceInit(fn func(MemberInfoProfile, MemberInfoProfile, int) MemberInfoProfile, initialValue MemberInfoProfile) *MemberInfoProfileStream {
	result := MemberInfoProfileStreamOf()
	self.ForEach(func(v MemberInfoProfile, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *MemberInfoProfileStream) ReduceInterface(fn func(interface{}, MemberInfoProfile, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(MemberInfoProfile{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *MemberInfoProfileStream) ReduceString(fn func(string, MemberInfoProfile, int) string) []string {
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
func (self *MemberInfoProfileStream) ReduceInt(fn func(int, MemberInfoProfile, int) int) []int {
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
func (self *MemberInfoProfileStream) ReduceInt32(fn func(int32, MemberInfoProfile, int) int32) []int32 {
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
func (self *MemberInfoProfileStream) ReduceInt64(fn func(int64, MemberInfoProfile, int) int64) []int64 {
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
func (self *MemberInfoProfileStream) ReduceFloat32(fn func(float32, MemberInfoProfile, int) float32) []float32 {
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
func (self *MemberInfoProfileStream) ReduceFloat64(fn func(float64, MemberInfoProfile, int) float64) []float64 {
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
func (self *MemberInfoProfileStream) ReduceBool(fn func(bool, MemberInfoProfile, int) bool) []bool {
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
func (self *MemberInfoProfileStream) Reverse() *MemberInfoProfileStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *MemberInfoProfileStream) Replace(fn func(MemberInfoProfile, int) MemberInfoProfile) *MemberInfoProfileStream {
	return self.ForEach(func(v MemberInfoProfile, i int) { self.Set(i, fn(v, i)) })
}
func (self *MemberInfoProfileStream) Select(fn func(MemberInfoProfile) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *MemberInfoProfileStream) Set(index int, val MemberInfoProfile) *MemberInfoProfileStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *MemberInfoProfileStream) Skip(skip int) *MemberInfoProfileStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *MemberInfoProfileStream) SkippingEach(fn func(MemberInfoProfile, int) int) *MemberInfoProfileStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *MemberInfoProfileStream) Slice(startIndex, n int) *MemberInfoProfileStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []MemberInfoProfile{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *MemberInfoProfileStream) Sort(fn func(i, j int) bool) *MemberInfoProfileStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *MemberInfoProfileStream) Tail() *MemberInfoProfile {
	return self.Last()
}
func (self *MemberInfoProfileStream) TailOr(arg MemberInfoProfile) MemberInfoProfile {
	return self.LastOr(arg)
}
func (self *MemberInfoProfileStream) ToList() []MemberInfoProfile {
	return self.Val()
}
func (self *MemberInfoProfileStream) Unique() *MemberInfoProfileStream {
	return self.Distinct()
}
func (self *MemberInfoProfileStream) Val() []MemberInfoProfile {
	if self == nil {
		return []MemberInfoProfile{}
	}
	return *self.Copy()
}
func (self *MemberInfoProfileStream) While(fn func(MemberInfoProfile, int) bool) *MemberInfoProfileStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *MemberInfoProfileStream) Where(fn func(MemberInfoProfile) bool) *MemberInfoProfileStream {
	result := MemberInfoProfileStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *MemberInfoProfileStream) WhereSlim(fn func(MemberInfoProfile) bool) *MemberInfoProfileStream {
	result := MemberInfoProfileStreamOf()
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