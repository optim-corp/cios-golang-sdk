package collection

import (
	"reflect"
	"sort"

	xmath "github.com/optim-corp/cios-golang-sdk/util/go_advance_type/math"
)

func Int64StreamOf(arg ...int64) Int64Stream {
	return arg
}
func Int64StreamFrom(arg []int64) Int64Stream {
	return arg
}
func CreateInt64Stream(arg ...int64) *Int64Stream {
	tmp := Int64StreamOf(arg...)
	return &tmp
}
func GenerateInt64Stream(arg []int64) *Int64Stream {
	tmp := Int64StreamFrom(arg)
	return &tmp
}

func (self *Int64Stream) Add(arg int64) *Int64Stream {
	return self.AddAll(arg)
}
func (self *Int64Stream) AddAll(arg ...int64) *Int64Stream {
	*self = append(*self, arg...)
	return self
}
func (self *Int64Stream) AddSafe(arg *int64) *Int64Stream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *Int64Stream) AllMatch(fn func(int64, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *Int64Stream) AnyMatch(fn func(int64, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *Int64Stream) Clone() *Int64Stream {
	temp := make([]int64, self.Len())
	copy(temp, *self)
	return (*Int64Stream)(&temp)
}
func (self *Int64Stream) Copy() *Int64Stream {
	return self.Clone()
}
func (self *Int64Stream) Concat(arg []int64) *Int64Stream {
	return self.AddAll(arg...)
}
func (self *Int64Stream) Contains(arg int64) bool {
	return self.FindIndex(func(_arg int64, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *Int64Stream) Clean() *Int64Stream {
	*self = Int64StreamOf()
	return self
}
func (self *Int64Stream) Delete(index int) *Int64Stream {
	return self.DeleteRange(index, index)
}
func (self *Int64Stream) DeleteRange(startIndex, endIndex int) *Int64Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *Int64Stream) Distinct() *Int64Stream {
	caches := map[int64]bool{}
	result := Int64StreamOf()
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
func (self *Int64Stream) Each(fn func(int64)) *Int64Stream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *Int64Stream) EachRight(fn func(int64)) *Int64Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *Int64Stream) Equals(arr []int64) bool {
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
func (self *Int64Stream) Where(fn func(int64) bool) *Int64Stream {
	result := Int64StreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *Int64Stream) WhereSlim(fn func(int64) bool) *Int64Stream {
	result := Int64StreamOf()
	caches := map[int64]bool{}
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
func (self *Int64Stream) Filter(fn func(int64, int) bool) *Int64Stream {
	result := Int64StreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *Int64Stream) FilterSlim(fn func(int64, int) bool) *Int64Stream {
	result := Int64StreamOf()
	caches := map[int64]bool{}
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
func (self *Int64Stream) Find(fn func(int64, int) bool) *int64 {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *Int64Stream) FindOr(fn func(int64, int) bool, or int64) int64 {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *Int64Stream) FindIndex(fn func(int64, int) bool) int {
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
func (self *Int64Stream) First() *int64 {
	return self.Get(0)
}
func (self *Int64Stream) FirstOr(arg int64) int64 {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *Int64Stream) ForEach(fn func(int64, int)) *Int64Stream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *Int64Stream) ForEachRight(fn func(int64, int)) *Int64Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *Int64Stream) GroupBy(fn func(int64, int) string) map[string][]int64 {
	m := map[string][]int64{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *Int64Stream) GroupByValues(fn func(int64, int) string) [][]int64 {
	var tmp [][]int64
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *Int64Stream) IndexOf(arg int64) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *Int64Stream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *Int64Stream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *Int64Stream) Last() *int64 {
	return self.Get(self.Len() - 1)
}
func (self *Int64Stream) LastOr(arg int64) int64 {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *Int64Stream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *Int64Stream) Limit(limit int) *Int64Stream {
	self.Slice(0, limit)
	return self
}
func (self *Int64Stream) Select(fn func(int64) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *Int64Stream) Map(fn func(int64, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Int64Stream) Map2Int(fn func(int64, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Int64Stream) Map2Int32(fn func(int64, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Int64Stream) Map2Int64(fn func(int64, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Int64Stream) Map2Float32(fn func(int64, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Int64Stream) Map2Float64(fn func(int64, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Int64Stream) Map2Bool(fn func(int64, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Int64Stream) Map2Bytes(fn func(int64, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Int64Stream) Map2String(fn func(int64, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Int64Stream) NoneMatch(fn func(int64, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *Int64Stream) Get(index int) *int64 {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *Int64Stream) GetOr(index int, arg int64) int64 {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *Int64Stream) Peek(fn func(*int64, int)) *Int64Stream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *Int64Stream) Aggregate(fn func(int64, int64) int64) *Int64Stream {
	result := Int64StreamOf()
	self.ForEach(func(v int64, i int) {
		if i == 0 {
			result.Add(fn(0, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *Int64Stream) Reduce(fn func(int64, int64, int) int64) *Int64Stream {
	return self.ReduceInit(fn, 0)
}
func (self *Int64Stream) ReduceInit(fn func(int64, int64, int) int64, initialValue int64) *Int64Stream {
	result := Int64StreamOf()
	self.ForEach(func(v int64, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *Int64Stream) ReduceInterface(fn func(interface{}, int64, int) interface{}) (result []interface{}) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(struct{}{}, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Int64Stream) ReduceString(fn func(string, int64, int) string) (result []string) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn("", v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Int64Stream) ReduceInt(fn func(int, int64, int) int) (result []int) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Int64Stream) ReduceInt32(fn func(int32, int64, int) int32) (result []int32) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Int64Stream) ReduceInt64(fn func(int64, int64, int) int64) (result []int64) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Int64Stream) ReduceFloat32(fn func(float32, int64, int) float32) (result []float32) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Int64Stream) ReduceFloat64(fn func(float64, int64, int) float64) (result []float64) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Int64Stream) ReduceBool(fn func(bool, int64, int) bool) (result []bool) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(false, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Int64Stream) Reverse() *Int64Stream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *Int64Stream) Replace(fn func(int64, int) int64) *Int64Stream {
	return self.ForEach(func(v int64, i int) { self.Set(i, fn(v, i)) })
}
func (self *Int64Stream) Set(index int, val int64) *Int64Stream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *Int64Stream) Skip(skip int) *Int64Stream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *Int64Stream) SkippingEach(fn func(int64, int) int) *Int64Stream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *Int64Stream) Slice(startIndex, n int) *Int64Stream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []int64{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *Int64Stream) Sort(fn func(i, j int) bool) *Int64Stream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *Int64Stream) Tail() *int64 {
	return self.Last()
}
func (self *Int64Stream) TailOr(arg int64) int64 {
	return self.LastOr(arg)
}
func (self *Int64Stream) ToList() []int64 {
	return self.Val()
}
func (self *Int64Stream) Unique() *Int64Stream {
	return self.Distinct()
}
func (self *Int64Stream) Val() []int64 {
	if self == nil {
		return []int64{}
	}
	return *self.Copy()
}
func (self *Int64Stream) While(fn func(int64, int) bool) *Int64Stream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}

func (self *Int64Stream) Min() int64 {
	return xmath.MinInt64(self.Val()...)
}

func (self *Int64Stream) Max() int64 {
	return xmath.MaxInt64(self.Val()...)
}

func (self *Int64Stream) Sum() int64 {
	return xmath.SumInt64(self.Val()...)
}

func (self *Int64Stream) Average() int64 {
	return xmath.AverageInt64(self.Val()...)
}
func (self *Int64Stream) Median() *int64 {
	return xmath.MedianInt64(self.Val()...)
}
func (self *Int64Stream) Mode() *int64 {
	return xmath.ModeInt64(self.Val()...)
}
func (self *Int64Stream) Range() *int64 {
	return xmath.RangeInt64(self.Val()...)
}

func (self *Int64Stream) AsFloat32() *Float32Stream {
	result := self.Map2Float32(func(arg int64, index int) float32 { return float32(arg) })
	stream := Float32StreamFrom(result)
	return &stream
}
func (self *Int64Stream) AsFloat64() *Float64Stream {
	result := self.Map2Float64(func(arg int64, index int) float64 { return float64(arg) })
	stream := Float64StreamFrom(result)
	return &stream
}
func (self *Int64Stream) AsInt() *IntegerStream {
	result := self.Map2Int(func(arg int64, index int) int { return int(arg) })
	stream := IntegerStreamFrom(result)
	return &stream
}
func (self *Int64Stream) AsInt32() *Int32Stream {
	result := self.Map2Int32(func(arg int64, index int) int32 { return int32(arg) })
	stream := Int32StreamFrom(result)
	return &stream
}
func (self *Int64Stream) ContainsSome(arg ...int64) bool {
	for _, v := range arg {
		if self.Contains(v) {
			return true
		}
	}
	return false
}
func (self *Int64Stream) ContainsAll(arg ...int64) bool {
	for _, v := range arg {
		if !self.Contains(v) {
			return false
		}
	}
	return true
}
