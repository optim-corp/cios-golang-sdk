package collection

import (
	"math"
	"reflect"
	"sort"
)

func BoolStreamOf(arg ...bool) BoolStream {
	return arg
}
func BoolStreamFrom(arg []bool) BoolStream {
	return arg
}
func CreateBoolStream(arg ...bool) *BoolStream {
	tmp := BoolStreamOf(arg...)
	return &tmp
}
func GenerateBoolStream(arg []bool) *BoolStream {
	tmp := BoolStreamFrom(arg)
	return &tmp
}

func (self *BoolStream) Add(arg bool) *BoolStream {
	return self.AddAll(arg)
}
func (self *BoolStream) AddAll(arg ...bool) *BoolStream {
	*self = append(*self, arg...)
	return self
}
func (self *BoolStream) AddSafe(arg *bool) *BoolStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *BoolStream) AllMatch(fn func(bool, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *BoolStream) AnyMatch(fn func(bool, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *BoolStream) Clone() *BoolStream {
	temp := make([]bool, self.Len())
	copy(temp, *self)
	return (*BoolStream)(&temp)
}
func (self *BoolStream) Copy() *BoolStream {
	return self.Clone()
}
func (self *BoolStream) Concat(arg []bool) *BoolStream {
	return self.AddAll(arg...)
}
func (self *BoolStream) Contains(arg bool) bool {
	return self.FindIndex(func(_arg bool, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *BoolStream) Clean() *BoolStream {
	*self = BoolStreamOf()
	return self
}
func (self *BoolStream) Delete(index int) *BoolStream {
	return self.DeleteRange(index, index)
}
func (self *BoolStream) DeleteRange(startIndex, endIndex int) *BoolStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *BoolStream) Distinct() *BoolStream {
	caches := map[bool]bool{}
	result := BoolStreamOf()
	for _, v := range *self {
		if f, ok := caches[v]; ok {
			if !f {
				result = append(result, v)
			}
		} else if caches[v] = true; !f {
			result = append(result, v)
		}

	}
	*self = result
	return self
}
func (self *BoolStream) Each(fn func(bool)) *BoolStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *BoolStream) EachRight(fn func(bool)) *BoolStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *BoolStream) Equals(arr []bool) bool {
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
func (self *BoolStream) Filter(fn func(bool, int) bool) *BoolStream {
	result := BoolStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *BoolStream) FilterSlim(fn func(bool, int) bool) *BoolStream {
	result := BoolStreamOf()
	caches := map[bool]bool{}
	for i, v := range *self {
		if f, ok := caches[v]; ok {
			if f {
				result.Add(v)
			}
		} else if caches[v] = fn(v, i); caches[v] {
			result.Add(v)

		}
	}
	*self = result
	return self
}
func (self *BoolStream) Find(fn func(bool, int) bool) *bool {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *BoolStream) FindOr(fn func(bool, int) bool, or bool) bool {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *BoolStream) FindIndex(fn func(bool, int) bool) int {
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
func (self *BoolStream) First() *bool {
	return self.Get(0)
}
func (self *BoolStream) FirstOr(arg bool) bool {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *BoolStream) ForEach(fn func(bool, int)) *BoolStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *BoolStream) ForEachRight(fn func(bool, int)) *BoolStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *BoolStream) GroupBy(fn func(bool, int) string) map[string][]bool {
	m := map[string][]bool{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *BoolStream) GroupByValues(fn func(bool, int) string) [][]bool {
	var tmp [][]bool
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *BoolStream) IndexOf(arg bool) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *BoolStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *BoolStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *BoolStream) Last() *bool {
	return self.Get(self.Len() - 1)
}
func (self *BoolStream) LastOr(arg bool) bool {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *BoolStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *BoolStream) Limit(limit int) *BoolStream {
	self.Slice(0, limit)
	return self
}
func (self *BoolStream) Map(fn func(bool, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BoolStream) Map2Int(fn func(bool, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BoolStream) Map2Int32(fn func(bool, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BoolStream) Map2Int64(fn func(bool, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BoolStream) Map2Float32(fn func(bool, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BoolStream) Map2Float64(fn func(bool, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BoolStream) Map2Bool(fn func(bool, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BoolStream) Map2Bytes(fn func(bool, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BoolStream) Map2String(fn func(bool, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *BoolStream) Max(fn func(bool, int) float64) *bool {
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
func (self *BoolStream) Min(fn func(bool, int) float64) *bool {
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
func (self *BoolStream) NoneMatch(fn func(bool, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *BoolStream) Get(index int) *bool {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *BoolStream) GetOr(index int, arg bool) bool {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *BoolStream) Peek(fn func(*bool, int)) *BoolStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}
func (self *BoolStream) Reduce(fn func(bool, bool, int) bool) *BoolStream {
	return self.ReduceInit(fn, false)
}
func (self *BoolStream) ReduceInit(fn func(bool, bool, int) bool, initialValue bool) *BoolStream {
	result := BoolStreamOf()
	self.ForEach(func(v bool, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *BoolStream) ReduceString(fn func(string, bool, int) string) (result []string) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn("", v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *BoolStream) ReduceInt(fn func(int, bool, int) int) (result []int) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *BoolStream) ReduceInt32(fn func(int32, bool, int) int32) (result []int32) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *BoolStream) ReduceInt64(fn func(int64, bool, int) int64) (result []int64) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *BoolStream) ReduceFloat32(fn func(float32, bool, int) float32) (result []float32) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *BoolStream) ReduceFloat64(fn func(float64, bool, int) float64) (result []float64) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *BoolStream) ReduceBool(fn func(bool, bool, int) bool) (result []bool) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(false, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *BoolStream) Reverse() *BoolStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *BoolStream) Replace(fn func(bool, int) bool) *BoolStream {
	return self.ForEach(func(v bool, i int) { self.Set(i, fn(v, i)) })
}
func (self *BoolStream) Set(index int, val bool) *BoolStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *BoolStream) Skip(skip int) *BoolStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *BoolStream) SkippingEach(fn func(bool, int) int) *BoolStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *BoolStream) Slice(startIndex, n int) *BoolStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []bool{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *BoolStream) Sort(fn func(i, j int) bool) *BoolStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *BoolStream) Tail() *bool {
	return self.Last()
}
func (self *BoolStream) TailOr(arg bool) bool {
	return self.LastOr(arg)
}
func (self *BoolStream) ToList() []bool {
	return self.Val()
}
func (self *BoolStream) Unique() *BoolStream {
	return self.Distinct()
}
func (self *BoolStream) Val() []bool {
	if self == nil {
		return []bool{}
	}
	return *self.Copy()
}
func (self *BoolStream) While(fn func(bool, int) bool) *BoolStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
