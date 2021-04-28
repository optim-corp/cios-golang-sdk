package collection

import (
	"reflect"
	"sort"

	xmath "github.com/optim-corp/cios-golang-sdk/util/go_advance_type/math"
)

func Float32StreamOf(arg ...float32) Float32Stream {
	return arg
}
func Float32StreamFrom(arg []float32) Float32Stream {
	return arg
}
func CreateFloat32Stream(arg ...float32) *Float32Stream {
	tmp := Float32StreamOf(arg...)
	return &tmp
}
func GenerateFloat32Stream(arg []float32) *Float32Stream {
	tmp := Float32StreamFrom(arg)
	return &tmp
}

func (self *Float32Stream) Add(arg float32) *Float32Stream {
	return self.AddAll(arg)
}
func (self *Float32Stream) AddAll(arg ...float32) *Float32Stream {
	*self = append(*self, arg...)
	return self
}
func (self *Float32Stream) AddSafe(arg *float32) *Float32Stream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *Float32Stream) AllMatch(fn func(float32, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *Float32Stream) AnyMatch(fn func(float32, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *Float32Stream) Clone() *Float32Stream {
	temp := make([]float32, self.Len())
	copy(temp, *self)
	return (*Float32Stream)(&temp)
}
func (self *Float32Stream) Copy() *Float32Stream {
	return self.Clone()
}
func (self *Float32Stream) Concat(arg []float32) *Float32Stream {
	return self.AddAll(arg...)
}
func (self *Float32Stream) Contains(arg float32) bool {
	return self.FindIndex(func(_arg float32, index int) bool { return _arg == arg }) != -1
}
func (self *Float32Stream) Clean() *Float32Stream {
	*self = Float32StreamOf()
	return self
}
func (self *Float32Stream) Delete(index int) *Float32Stream {
	return self.DeleteRange(index, index)
}
func (self *Float32Stream) DeleteRange(startIndex, endIndex int) *Float32Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *Float32Stream) Distinct() *Float32Stream {
	caches := map[float32]bool{}
	result := Float32StreamOf()
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
func (self *Float32Stream) Each(fn func(float32)) *Float32Stream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *Float32Stream) EachRight(fn func(float32)) *Float32Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *Float32Stream) Equals(arr []float32) bool {
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
func (self *Float32Stream) Filter(fn func(float32, int) bool) *Float32Stream {
	result := Float32StreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *Float32Stream) FilterSlim(fn func(float32, int) bool) *Float32Stream {
	result := Float32StreamOf()
	caches := map[float32]bool{}
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
func (self *Float32Stream) Find(fn func(float32, int) bool) *float32 {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *Float32Stream) FindOr(fn func(float32, int) bool, or float32) float32 {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *Float32Stream) FindIndex(fn func(float32, int) bool) int {
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
func (self *Float32Stream) First() *float32 {
	return self.Get(0)
}
func (self *Float32Stream) FirstOr(arg float32) float32 {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *Float32Stream) ForEach(fn func(float32, int)) *Float32Stream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *Float32Stream) ForEachRight(fn func(float32, int)) *Float32Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *Float32Stream) GroupBy(fn func(float32, int) string) map[string][]float32 {
	m := map[string][]float32{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *Float32Stream) GroupByValues(fn func(float32, int) string) [][]float32 {
	var tmp [][]float32
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *Float32Stream) IndexOf(arg float32) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *Float32Stream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *Float32Stream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *Float32Stream) Last() *float32 {
	return self.Get(self.Len() - 1)
}
func (self *Float32Stream) LastOr(arg float32) float32 {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *Float32Stream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *Float32Stream) Limit(limit int) *Float32Stream {
	self.Slice(0, limit)
	return self
}
func (self *Float32Stream) Map(fn func(float32, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Float32Stream) Map2Int(fn func(float32, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Float32Stream) Map2Int32(fn func(float32, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Float32Stream) Map2Int64(fn func(float32, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Float32Stream) Map2Float32(fn func(float32, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Float32Stream) Map2Float64(fn func(float32, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Float32Stream) Map2Bool(fn func(float32, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Float32Stream) Map2Bytes(fn func(float32, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Float32Stream) Map2String(fn func(float32, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Float32Stream) NoneMatch(fn func(float32, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *Float32Stream) Get(index int) *float32 {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *Float32Stream) GetOr(index int, arg float32) float32 {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *Float32Stream) Peek(fn func(*float32, int)) *Float32Stream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}
func (self *Float32Stream) Reduce(fn func(float32, float32, int) float32) *Float32Stream {
	return self.ReduceInit(fn, 0)
}
func (self *Float32Stream) ReduceInit(fn func(float32, float32, int) float32, initialValue float32) *Float32Stream {
	result := Float32StreamOf()
	self.ForEach(func(v float32, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *Float32Stream) ReduceString(fn func(string, float32, int) string) (result []string) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn("", v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Float32Stream) ReduceInt(fn func(int, float32, int) int) (result []int) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Float32Stream) ReduceInt32(fn func(int32, float32, int) int32) (result []int32) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Float32Stream) ReduceInt64(fn func(int64, float32, int) int64) (result []int64) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Float32Stream) ReduceFloat32(fn func(float32, float32, int) float32) (result []float32) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Float32Stream) ReduceFloat64(fn func(float64, float32, int) float64) (result []float64) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Float32Stream) ReduceBool(fn func(bool, float32, int) bool) (result []bool) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(false, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Float32Stream) Reverse() *Float32Stream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *Float32Stream) Replace(fn func(float32, int) float32) *Float32Stream {
	return self.ForEach(func(v float32, i int) { self.Set(i, fn(v, i)) })
}
func (self *Float32Stream) Set(index int, val float32) *Float32Stream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *Float32Stream) Skip(skip int) *Float32Stream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *Float32Stream) SkippingEach(fn func(float32, int) int) *Float32Stream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *Float32Stream) Slice(startIndex, n int) *Float32Stream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []float32{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *Float32Stream) Sort(fn func(i, j int) bool) *Float32Stream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *Float32Stream) Tail() *float32 {
	return self.Last()
}
func (self *Float32Stream) TailOr(arg float32) float32 {
	return self.LastOr(arg)
}
func (self *Float32Stream) ToList() []float32 {
	return self.Val()
}
func (self *Float32Stream) Unique() *Float32Stream {
	return self.Distinct()
}
func (self *Float32Stream) Val() []float32 {
	if self == nil {
		return []float32{}
	}
	return *self.Copy()
}
func (self *Float32Stream) While(fn func(float32, int) bool) *Float32Stream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}

func (self *Float32Stream) Min() float32 {
	return xmath.MinFloat32(self.Val()...)
}

func (self *Float32Stream) Max() float32 {
	return xmath.MaxFloat32(self.Val()...)
}

func (self *Float32Stream) Sum() float32 {
	return xmath.SumFloat32(self.Val()...)
}

func (self *Float32Stream) Average() float32 {
	return xmath.AverageFloat32(self.Val()...)
}
func (self *Float32Stream) Median() *float32 {
	return xmath.MedianFloat32(self.Val()...)
}
func (self *Float32Stream) Mode() *float32 {
	return xmath.ModeFloat32(self.Val()...)
}
func (self *Float32Stream) Range() *float32 {
	return xmath.RangeFloat32(self.Val()...)
}
func (self *Float32Stream) AsFloat64() *Float64Stream {
	result := self.Map2Float64(func(arg float32, index int) float64 { return float64(arg) })
	stream := Float64StreamFrom(result)
	return &stream
}

func (self *Float32Stream) AsInt() *IntegerStream {
	result := self.Map2Int(func(arg float32, index int) int { return int(arg) })
	stream := IntegerStreamFrom(result)
	return &stream
}
func (self *Float32Stream) AsInt32() *Int32Stream {
	result := self.Map2Int32(func(arg float32, index int) int32 { return int32(arg) })
	stream := Int32StreamFrom(result)
	return &stream
}

func (self *Float32Stream) AsInt64() *Int64Stream {
	result := self.Map2Int64(func(arg float32, index int) int64 { return int64(arg) })
	stream := Int64StreamFrom(result)
	return &stream
}
func (self *Float32Stream) ContainsSome(arg ...float32) bool {
	for _, v := range arg {
		if self.Contains(v) {
			return true
		}
	}
	return false
}
func (self *Float32Stream) ContainsAll(arg ...float32) bool {
	for _, v := range arg {
		if !self.Contains(v) {
			return false
		}
	}
	return true
}
