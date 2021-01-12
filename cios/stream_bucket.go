/*
 * Collection utility of Bucket Struct
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

type BucketStream []Bucket

func BucketStreamOf(arg ...Bucket) BucketStream {
	return arg
}
func BucketStreamFrom(arg []Bucket) BucketStream {
	return arg
}
func CreateBucketStream(arg ...Bucket) *BucketStream {
	tmp := BucketStreamOf(arg...)
	return &tmp
}
func GenerateBucketStream(arg []Bucket) *BucketStream {
	tmp := BucketStreamFrom(arg)
	return &tmp
}

func (self *BucketStream) Add(arg Bucket) *BucketStream {
	return self.AddAll(arg)
}
func (self *BucketStream) AddAll(arg ...Bucket) *BucketStream {
	*self = append(*self, arg...)
	return self
}
func (self *BucketStream) AddSafe(arg *Bucket) *BucketStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *BucketStream) Aggregate(fn func(Bucket, Bucket) Bucket) *BucketStream {
	result := BucketStreamOf()
	self.ForEach(func(v Bucket, i int) {
		if i == 0 {
			result.Add(fn(Bucket{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *BucketStream) AllMatch(fn func(Bucket, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *BucketStream) AnyMatch(fn func(Bucket, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *BucketStream) Clone() *BucketStream {
	temp := make([]Bucket, self.Len())
	copy(temp, *self)
	return (*BucketStream)(&temp)
}
func (self *BucketStream) Copy() *BucketStream {
	return self.Clone()
}
func (self *BucketStream) Concat(arg []Bucket) *BucketStream {
	return self.AddAll(arg...)
}
func (self *BucketStream) Contains(arg Bucket) bool {
	return self.FindIndex(func(_arg Bucket, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *BucketStream) Clean() *BucketStream {
	*self = BucketStreamOf()
	return self
}
func (self *BucketStream) Delete(index int) *BucketStream {
	return self.DeleteRange(index, index)
}
func (self *BucketStream) DeleteRange(startIndex, endIndex int) *BucketStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *BucketStream) Distinct() *BucketStream {
	caches := map[string]bool{}
	result := BucketStreamOf()
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
func (self *BucketStream) Each(fn func(Bucket)) *BucketStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *BucketStream) EachRight(fn func(Bucket)) *BucketStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *BucketStream) Equals(arr []Bucket) bool {
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
func (self *BucketStream) Filter(fn func(Bucket, int) bool) *BucketStream {
	result := BucketStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *BucketStream) FilterSlim(fn func(Bucket, int) bool) *BucketStream {
	result := BucketStreamOf()
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
func (self *BucketStream) Find(fn func(Bucket, int) bool) *Bucket {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *BucketStream) FindOr(fn func(Bucket, int) bool, or Bucket) Bucket {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *BucketStream) FindIndex(fn func(Bucket, int) bool) int {
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
func (self *BucketStream) First() *Bucket {
	return self.Get(0)
}
func (self *BucketStream) FirstOr(arg Bucket) Bucket {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *BucketStream) ForEach(fn func(Bucket, int)) *BucketStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *BucketStream) ForEachRight(fn func(Bucket, int)) *BucketStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *BucketStream) GroupBy(fn func(Bucket, int) string) map[string][]Bucket {
	m := map[string][]Bucket{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *BucketStream) GroupByValues(fn func(Bucket, int) string) [][]Bucket {
	var tmp [][]Bucket
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *BucketStream) IndexOf(arg Bucket) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *BucketStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *BucketStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *BucketStream) Last() *Bucket {
	return self.Get(self.Len() - 1)
}
func (self *BucketStream) LastOr(arg Bucket) Bucket {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *BucketStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *BucketStream) Limit(limit int) *BucketStream {
	self.Slice(0, limit)
	return self
}

func (self *BucketStream) Map(fn func(Bucket, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketStream) Map2Int(fn func(Bucket, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketStream) Map2Int32(fn func(Bucket, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketStream) Map2Int64(fn func(Bucket, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketStream) Map2Float32(fn func(Bucket, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketStream) Map2Float64(fn func(Bucket, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketStream) Map2Bool(fn func(Bucket, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketStream) Map2Bytes(fn func(Bucket, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketStream) Map2String(fn func(Bucket, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BucketStream) Max(fn func(Bucket, int) float64) *Bucket {
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
func (self *BucketStream) Min(fn func(Bucket, int) float64) *Bucket {
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
func (self *BucketStream) NoneMatch(fn func(Bucket, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *BucketStream) Get(index int) *Bucket {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *BucketStream) GetOr(index int, arg Bucket) Bucket {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *BucketStream) Peek(fn func(*Bucket, int)) *BucketStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *BucketStream) Reduce(fn func(Bucket, Bucket, int) Bucket) *BucketStream {
	return self.ReduceInit(fn, Bucket{})
}
func (self *BucketStream) ReduceInit(fn func(Bucket, Bucket, int) Bucket, initialValue Bucket) *BucketStream {
	result := BucketStreamOf()
	self.ForEach(func(v Bucket, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *BucketStream) ReduceInterface(fn func(interface{}, Bucket, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(Bucket{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *BucketStream) ReduceString(fn func(string, Bucket, int) string) []string {
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
func (self *BucketStream) ReduceInt(fn func(int, Bucket, int) int) []int {
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
func (self *BucketStream) ReduceInt32(fn func(int32, Bucket, int) int32) []int32 {
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
func (self *BucketStream) ReduceInt64(fn func(int64, Bucket, int) int64) []int64 {
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
func (self *BucketStream) ReduceFloat32(fn func(float32, Bucket, int) float32) []float32 {
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
func (self *BucketStream) ReduceFloat64(fn func(float64, Bucket, int) float64) []float64 {
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
func (self *BucketStream) ReduceBool(fn func(bool, Bucket, int) bool) []bool {
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
func (self *BucketStream) Reverse() *BucketStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *BucketStream) Replace(fn func(Bucket, int) Bucket) *BucketStream {
	return self.ForEach(func(v Bucket, i int) { self.Set(i, fn(v, i)) })
}
func (self *BucketStream) Select(fn func(Bucket) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *BucketStream) Set(index int, val Bucket) *BucketStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *BucketStream) Skip(skip int) *BucketStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *BucketStream) SkippingEach(fn func(Bucket, int) int) *BucketStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *BucketStream) Slice(startIndex, n int) *BucketStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []Bucket{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *BucketStream) Sort(fn func(i, j int) bool) *BucketStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *BucketStream) Tail() *Bucket {
	return self.Last()
}
func (self *BucketStream) TailOr(arg Bucket) Bucket {
	return self.LastOr(arg)
}
func (self *BucketStream) ToList() []Bucket {
	return self.Val()
}
func (self *BucketStream) Unique() *BucketStream {
	return self.Distinct()
}
func (self *BucketStream) Val() []Bucket {
	if self == nil {
		return []Bucket{}
	}
	return *self.Copy()
}
func (self *BucketStream) While(fn func(Bucket, int) bool) *BucketStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *BucketStream) Where(fn func(Bucket) bool) *BucketStream {
	result := BucketStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *BucketStream) WhereSlim(fn func(Bucket) bool) *BucketStream {
	result := BucketStreamOf()
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
