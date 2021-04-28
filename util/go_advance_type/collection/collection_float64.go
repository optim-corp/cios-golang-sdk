package collection

import (
	"reflect"
	"sort"

	xmath "github.com/optim-corp/cios-golang-sdk/util/go_advance_type/math"
)

func Float64StreamOf(arg ...float64) Float64Stream {
	return arg
}
func Float64StreamFrom(arg []float64) Float64Stream {
	return arg
}
func CreateFloat64Stream(arg ...float64) *Float64Stream {
	tmp := Float64StreamOf(arg...)
	return &tmp
}
func GenerateFloat64Stream(arg []float64) *Float64Stream {
	tmp := Float64StreamFrom(arg)
	return &tmp
}

func (self *Float64Stream) Add(arg float64) *Float64Stream {
	return self.AddAll(arg)
}
func (self *Float64Stream) AddAll(arg ...float64) *Float64Stream {
	*self = append(*self, arg...)
	return self
}
func (self *Float64Stream) AddSafe(arg *float64) *Float64Stream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *Float64Stream) AllMatch(fn func(float64, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *Float64Stream) AnyMatch(fn func(float64, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *Float64Stream) Clone() *Float64Stream {
	temp := make([]float64, self.Len())
	copy(temp, *self)
	return (*Float64Stream)(&temp)
}
func (self *Float64Stream) Copy() *Float64Stream {
	return self.Clone()
}
func (self *Float64Stream) Concat(arg []float64) *Float64Stream {
	return self.AddAll(arg...)
}
func (self *Float64Stream) Contains(arg float64) bool {
	return self.FindIndex(func(_arg float64, index int) bool { return _arg == arg }) != -1
}
func (self *Float64Stream) Clean() *Float64Stream {
	*self = Float64StreamOf()
	return self
}
func (self *Float64Stream) Delete(index int) *Float64Stream {
	return self.DeleteRange(index, index)
}
func (self *Float64Stream) DeleteRange(startIndex, endIndex int) *Float64Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *Float64Stream) Distinct() *Float64Stream {
	caches := map[float64]bool{}
	result := Float64StreamOf()
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
func (self *Float64Stream) Each(fn func(float64)) *Float64Stream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *Float64Stream) EachRight(fn func(float64)) *Float64Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *Float64Stream) Equals(arr []float64) bool {
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
func (self *Float64Stream) Where(fn func(float64) bool) *Float64Stream {
	result := Float64StreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *Float64Stream) WhereSlim(fn func(float64) bool) *Float64Stream {
	result := Float64StreamOf()
	caches := map[float64]bool{}
	for _, v := range *self {
		if f, ok := caches[v]; ok {
			if f {
				result.Add(v)
			}
		} else if caches[v] = fn(v); caches[v] {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *Float64Stream) Filter(fn func(float64, int) bool) *Float64Stream {
	result := Float64StreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *Float64Stream) FilterSlim(fn func(float64, int) bool) *Float64Stream {
	result := Float64StreamOf()
	caches := map[float64]bool{}
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
func (self *Float64Stream) Find(fn func(float64, int) bool) *float64 {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *Float64Stream) FindOr(fn func(float64, int) bool, or float64) float64 {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *Float64Stream) FindIndex(fn func(float64, int) bool) int {
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
func (self *Float64Stream) First() *float64 {
	return self.Get(0)
}
func (self *Float64Stream) FirstOr(arg float64) float64 {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *Float64Stream) ForEach(fn func(float64, int)) *Float64Stream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *Float64Stream) ForEachRight(fn func(float64, int)) *Float64Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *Float64Stream) GroupBy(fn func(float64, int) string) map[string][]float64 {
	m := map[string][]float64{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *Float64Stream) GroupByValues(fn func(float64, int) string) [][]float64 {
	var tmp [][]float64
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *Float64Stream) IndexOf(arg float64) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *Float64Stream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *Float64Stream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *Float64Stream) Last() *float64 {
	return self.Get(self.Len() - 1)
}
func (self *Float64Stream) LastOr(arg float64) float64 {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *Float64Stream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *Float64Stream) Limit(limit int) *Float64Stream {
	self.Slice(0, limit)
	return self
}
func (self *Float64Stream) Select(fn func(float64) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *Float64Stream) Map(fn func(float64, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Float64Stream) Map2Int(fn func(float64, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Float64Stream) Map2Int32(fn func(float64, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Float64Stream) Map2Int64(fn func(float64, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Float64Stream) Map2Float32(fn func(float64, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Float64Stream) Map2Bool(fn func(float64, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Float64Stream) Map2Bytes(fn func(float64, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Float64Stream) Map2String(fn func(float64, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Float64Stream) NoneMatch(fn func(float64, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *Float64Stream) Get(index int) *float64 {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *Float64Stream) GetOr(index int, arg float64) float64 {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *Float64Stream) Peek(fn func(*float64, int)) *Float64Stream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *Float64Stream) Aggregate(fn func(float64, float64) float64) *Float64Stream {
	result := Float64StreamOf()
	self.ForEach(func(v float64, i int) {
		if i == 0 {
			result.Add(fn(0, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *Float64Stream) Reduce(fn func(float64, float64, int) float64) *Float64Stream {
	return self.ReduceInit(fn, 0)
}
func (self *Float64Stream) ReduceInit(fn func(float64, float64, int) float64, initialValue float64) *Float64Stream {
	result := Float64StreamOf()
	self.ForEach(func(v float64, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *Float64Stream) ReduceString(fn func(string, float64, int) string) (result []string) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn("", v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Float64Stream) ReduceInt(fn func(int, float64, int) int) (result []int) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Float64Stream) ReduceInt32(fn func(int32, float64, int) int32) (result []int32) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Float64Stream) ReduceInt64(fn func(int64, float64, int) int64) (result []int64) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Float64Stream) ReduceFloat32(fn func(float32, float64, int) float32) (result []float32) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Float64Stream) ReduceFloat64(fn func(float64, float64, int) float64) (result []float64) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Float64Stream) ReduceBool(fn func(bool, float64, int) bool) (result []bool) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(false, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Float64Stream) Reverse() *Float64Stream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *Float64Stream) Replace(fn func(float64, int) float64) *Float64Stream {
	return self.ForEach(func(v float64, i int) { self.Set(i, fn(v, i)) })
}
func (self *Float64Stream) Set(index int, val float64) *Float64Stream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *Float64Stream) Skip(skip int) *Float64Stream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *Float64Stream) SkippingEach(fn func(float64, int) int) *Float64Stream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *Float64Stream) Slice(startIndex, n int) *Float64Stream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []float64{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *Float64Stream) Sort(fn func(i, j int) bool) *Float64Stream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *Float64Stream) Tail() *float64 {
	return self.Last()
}
func (self *Float64Stream) TailOr(arg float64) float64 {
	return self.LastOr(arg)
}
func (self *Float64Stream) ToList() []float64 {
	return self.Val()
}
func (self *Float64Stream) Unique() *Float64Stream {
	return self.Distinct()
}
func (self *Float64Stream) Val() []float64 {
	if self == nil {
		return []float64{}
	}
	return *self.Copy()
}
func (self *Float64Stream) While(fn func(float64, int) bool) *Float64Stream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}

func (self *Float64Stream) Min() float64 {
	return xmath.MinFloat64(self.Val()...)
}

func (self *Float64Stream) Max() float64 {
	return xmath.MaxFloat64(self.Val()...)
}

func (self *Float64Stream) Sum() float64 {
	return xmath.SumFloat64(self.Val()...)
}

func (self *Float64Stream) Average() float64 {
	return xmath.AverageFloat64(self.Val()...)
}

func (self *Float64Stream) Median() *float64 {
	return xmath.MedianFloat64(self.Val()...)
}
func (self *Float64Stream) Mode() *float64 {
	return xmath.ModeFloat64(self.Val()...)
}
func (self *Float64Stream) Range() *float64 {
	return xmath.RangeFloat64(self.Val()...)
}
func (self *Float64Stream) AsFloat32() *Float32Stream {
	result := self.Map2Float32(func(arg float64, index int) float32 { return float32(arg) })
	stream := Float32StreamFrom(result)
	return &stream
}
func (self *Float64Stream) AsInt() *IntegerStream {
	result := self.Map2Int(func(arg float64, index int) int { return int(arg) })
	stream := IntegerStreamFrom(result)
	return &stream
}
func (self *Float64Stream) AsInt32() *Int32Stream {
	result := self.Map2Int32(func(arg float64, index int) int32 { return int32(arg) })
	stream := Int32StreamFrom(result)
	return &stream
}

func (self *Float64Stream) AsInt64() *Int64Stream {
	result := self.Map2Int64(func(arg float64, index int) int64 { return int64(arg) })
	stream := Int64StreamFrom(result)
	return &stream
}

func (self *Float64Stream) ContainsSome(arg ...float64) bool {
	for _, v := range arg {
		if self.Contains(v) {
			return true
		}
	}
	return false
}
func (self *Float64Stream) ContainsAll(arg ...float64) bool {
	for _, v := range arg {
		if !self.Contains(v) {
			return false
		}
	}
	return true
}
