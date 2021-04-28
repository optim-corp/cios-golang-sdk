package collection

import (
	"reflect"
	"sort"

	xmath "github.com/optim-corp/cios-golang-sdk/util/go_advance_type/math"
)

func IntegerStreamOf(arg ...int) IntegerStream {
	return arg
}
func IntegerStreamFrom(arg []int) IntegerStream {
	return arg
}
func CreateIntegerStream(arg ...int) *IntegerStream {
	tmp := IntegerStreamOf(arg...)
	return &tmp
}
func GenerateIntegerStream(arg []int) *IntegerStream {
	tmp := IntegerStreamFrom(arg)
	return &tmp
}

func (self *IntegerStream) Add(arg int) *IntegerStream {
	return self.AddAll(arg)
}
func (self *IntegerStream) AddAll(arg ...int) *IntegerStream {
	*self = append(*self, arg...)
	return self
}
func (self *IntegerStream) AddSafe(arg *int) *IntegerStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *IntegerStream) AllMatch(fn func(int, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *IntegerStream) AnyMatch(fn func(int, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *IntegerStream) Clone() *IntegerStream {
	temp := make([]int, self.Len())
	copy(temp, *self)
	return (*IntegerStream)(&temp)
}
func (self *IntegerStream) Copy() *IntegerStream {
	return self.Clone()
}
func (self *IntegerStream) Concat(arg []int) *IntegerStream {
	return self.AddAll(arg...)
}
func (self *IntegerStream) Contains(arg int) bool {
	return self.FindIndex(func(_arg int, index int) bool { return _arg == arg }) != -1
}
func (self *IntegerStream) Clean() *IntegerStream {
	*self = IntegerStreamOf()
	return self
}
func (self *IntegerStream) Delete(index int) *IntegerStream {
	return self.DeleteRange(index, index)
}
func (self *IntegerStream) DeleteRange(startIndex, endIndex int) *IntegerStream {
	if self != nil {
		*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	}
	return self
}
func (self *IntegerStream) Distinct() *IntegerStream {
	if self != nil {
		caches := map[int]bool{}
		result := IntegerStreamOf()
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
	}
	return self
}
func (self *IntegerStream) Each(fn func(int)) *IntegerStream {
	if self != nil {
		for _, v := range *self {
			fn(v)
		}
	}
	return self
}
func (self *IntegerStream) EachRight(fn func(int)) *IntegerStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *IntegerStream) Equals(arr []int) bool {
	if self != nil {
		if (*self == nil) != (arr == nil) || len(*self) != len(arr) {
			return false
		}
		for i := range *self {
			if !reflect.DeepEqual((*self)[i], arr[i]) {
				return false
			}
		}
	}
	return true
}
func (self *IntegerStream) Where(fn func(int) bool) *IntegerStream {
	if self != nil {
		result := IntegerStreamOf()
		for _, v := range *self {
			if fn(v) {
				result.Add(v)
			}
		}
		*self = result
	}
	return self
}
func (self *IntegerStream) WhereSlim(fn func(int) bool) *IntegerStream {
	if self != nil {
		result := IntegerStreamOf()
		caches := map[int]bool{}
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
	}
	return self
}
func (self *IntegerStream) Filter(fn func(int, int) bool) *IntegerStream {
	if self != nil {
		result := IntegerStreamOf()
		for i, v := range *self {
			if fn(v, i) {
				result.Add(v)
			}
		}
		*self = result
	}
	return self
}
func (self *IntegerStream) FilterSlim(fn func(int, int) bool) *IntegerStream {
	if self != nil {
		result := IntegerStreamOf()
		caches := map[int]bool{}
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
	}
	return self
}
func (self *IntegerStream) Find(fn func(int, int) bool) *int {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *IntegerStream) FindOr(fn func(int, int) bool, or int) int {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *IntegerStream) FindIndex(fn func(int, int) bool) int {
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
func (self *IntegerStream) First() *int {
	return self.Get(0)
}
func (self *IntegerStream) FirstOr(arg int) int {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *IntegerStream) ForEach(fn func(int, int)) *IntegerStream {
	if self != nil {
		for i, v := range *self {
			fn(v, i)
		}
	}
	return self
}
func (self *IntegerStream) ForEachRight(fn func(int, int)) *IntegerStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *IntegerStream) GroupBy(fn func(int, int) string) map[string][]int {
	m := map[string][]int{}
	if self == nil {
		return nil
	}
	for i, v := range *self {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *IntegerStream) GroupByValues(fn func(int, int) string) [][]int {
	var tmp [][]int
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *IntegerStream) IndexOf(arg int) int {
	if self != nil {
		for index, _arg := range *self {
			if _arg == arg {
				return index
			}
		}
	}
	return -1
}
func (self *IntegerStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *IntegerStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *IntegerStream) Last() *int {
	return self.Get(self.Len() - 1)
}
func (self *IntegerStream) LastOr(arg int) int {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *IntegerStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *IntegerStream) Limit(limit int) *IntegerStream {
	self.Slice(0, limit)
	return self
}
func (self *IntegerStream) Select(fn func(int) interface{}) interface{} {
	if self == nil {
		return struct{}{}
	}
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *IntegerStream) Map(fn func(int, int) interface{}) interface{} {
	if self == nil {
		return struct{}{}
	}
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *IntegerStream) Map2Int(fn func(int, int) int) []int {
	if self == nil {
		return nil
	}
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *IntegerStream) Map2Int32(fn func(int, int) int32) []int32 {
	if self == nil {
		return nil
	}
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *IntegerStream) Map2Int64(fn func(int, int) int64) []int64 {
	if self == nil {
		return nil
	}
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *IntegerStream) Map2Float32(fn func(int, int) float32) []float32 {
	if self == nil {
		return nil
	}
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *IntegerStream) Map2Float64(fn func(int, int) float64) []float64 {
	if self == nil {
		return nil
	}
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *IntegerStream) Map2Bool(fn func(int, int) bool) []bool {
	if self == nil {
		return nil
	}
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *IntegerStream) Map2Bytes(fn func(int, int) []byte) [][]byte {
	if self == nil {
		return nil
	}
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *IntegerStream) Map2String(fn func(int, int) string) []string {
	if self == nil {
		return nil
	}
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *IntegerStream) NoneMatch(fn func(int, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *IntegerStream) Get(index int) *int {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *IntegerStream) GetOr(index int, arg int) int {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *IntegerStream) Peek(fn func(*int, int)) *IntegerStream {
	if self != nil {
		for i, v := range *self {
			fn(&v, i)
			self.Set(i, v)
		}
	}
	return self
}

func (self *IntegerStream) Aggregate(fn func(int, int) int) *IntegerStream {
	if self != nil {
		result := IntegerStreamOf()
		self.ForEach(func(v int, i int) {
			if i == 0 {
				result.Add(fn(0, v))
			} else {
				result.Add(fn(result[i-1], v))
			}
		})
		*self = result
	}
	return self
}
func (self *IntegerStream) Reduce(fn func(int, int, int) int) *IntegerStream {
	return self.ReduceInit(fn, 0)
}
func (self *IntegerStream) ReduceInit(fn func(int, int, int) int, initialValue int) *IntegerStream {
	if self != nil {
		result := IntegerStreamOf()
		self.ForEach(func(v int, i int) {
			if i == 0 {
				result.Add(fn(initialValue, v, i))
			} else {
				result.Add(fn(result[i-1], v, i))
			}
		})
		*self = result
	}
	return self
}
func (self *IntegerStream) ReduceInterface(fn func(interface{}, int, int) interface{}) (result []interface{}) {
	if self == nil {
		return nil
	}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(struct{}{}, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *IntegerStream) ReduceString(fn func(string, int, int) string) (result []string) {
	if self == nil {
		return nil
	}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn("", v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *IntegerStream) ReduceInt(fn func(int, int, int) int) (result []int) {
	if self == nil {
		return nil
	}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *IntegerStream) ReduceInt32(fn func(int32, int, int) int32) (result []int32) {
	if self == nil {
		return nil
	}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *IntegerStream) ReduceInt64(fn func(int64, int, int) int64) (result []int64) {
	if self == nil {
		return nil
	}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *IntegerStream) ReduceFloat32(fn func(float32, int, int) float32) (result []float32) {
	if self == nil {
		return nil
	}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *IntegerStream) ReduceFloat64(fn func(float64, int, int) float64) (result []float64) {
	if self == nil {
		return nil
	}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *IntegerStream) ReduceBool(fn func(bool, int, int) bool) (result []bool) {
	if self == nil {
		return nil
	}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(false, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *IntegerStream) Reverse() *IntegerStream {
	if self != nil {
		for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
			(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
		}
	}
	return self
}
func (self *IntegerStream) Replace(fn func(int, int) int) *IntegerStream {
	return self.ForEach(func(v int, i int) { self.Set(i, fn(v, i)) })
}
func (self *IntegerStream) Set(index int, val int) *IntegerStream {
	if self != nil {
		if len(*self) > index && index >= 0 {
			(*self)[index] = val
		}
	}
	return self
}
func (self *IntegerStream) Skip(skip int) *IntegerStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *IntegerStream) SkippingEach(fn func(int, int) int) *IntegerStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *IntegerStream) Slice(startIndex, n int) *IntegerStream {
	if self != nil {
		if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
			*self = []int{}
		} else if len(*self) < last {
			*self = (*self)[startIndex:len(*self)]
		} else {
			*self = (*self)[startIndex:last]
		}
	}
	return self
}
func (self *IntegerStream) Sort(fn func(i, j int) bool) *IntegerStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *IntegerStream) Tail() *int {
	return self.Last()
}
func (self *IntegerStream) TailOr(arg int) int {
	return self.LastOr(arg)
}
func (self *IntegerStream) ToList() []int {
	return self.Val()
}
func (self *IntegerStream) Unique() *IntegerStream {
	return self.Distinct()
}
func (self *IntegerStream) Val() []int {
	if self == nil {
		return []int{}
	}
	return *self.Copy()
}
func (self *IntegerStream) While(fn func(int, int) bool) *IntegerStream {
	if self != nil {
		for i, v := range *self {
			if !fn(v, i) {
				break
			}
		}
	}
	return self
}

func (self *IntegerStream) Min() int {
	return xmath.MinInt(self.Val()...)
}

func (self *IntegerStream) Max() int {
	return xmath.MaxInt(self.Val()...)
}

func (self *IntegerStream) Sum() int {
	return xmath.SumInt(self.Val()...)
}

func (self *IntegerStream) Average() int {
	return xmath.AverageInt(self.Val()...)
}
func (self *IntegerStream) Median() *int {
	return xmath.MedianInt(self.Val()...)
}
func (self *IntegerStream) Mode() *int {
	return xmath.ModeInt(self.Val()...)
}
func (self *IntegerStream) Range() *int {
	return xmath.RangeInt(self.Val()...)
}

func (self *IntegerStream) AsFloat32() *Float32Stream {
	result := self.Map2Float32(func(arg int, index int) float32 { return float32(arg) })
	stream := Float32StreamFrom(result)
	return &stream
}
func (self *IntegerStream) AsFloat64() *Float64Stream {
	result := self.Map2Float64(func(arg int, index int) float64 { return float64(arg) })
	stream := Float64StreamFrom(result)
	return &stream
}
func (self *IntegerStream) AsInt32() *Int32Stream {
	result := self.Map2Int32(func(arg int, index int) int32 { return int32(arg) })
	stream := Int32StreamFrom(result)
	return &stream
}

func (self *IntegerStream) AsInt64() *Int64Stream {
	result := self.Map2Int64(func(arg int, index int) int64 { return int64(arg) })
	stream := Int64StreamFrom(result)
	return &stream
}

func (self *IntegerStream) ContainsSome(arg ...int) bool {
	for _, v := range arg {
		if self.Contains(v) {
			return true
		}
	}
	return false
}
func (self *IntegerStream) ContainsAll(arg ...int) bool {
	for _, v := range arg {
		if !self.Contains(v) {
			return false
		}
	}
	return true
}
