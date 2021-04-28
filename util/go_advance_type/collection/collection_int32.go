package collection

import (
	"reflect"
	"sort"

	xmath "github.com/optim-corp/cios-golang-sdk/util/go_advance_type/math"
)

func Int32StreamOf(arg ...int32) Int32Stream {
	return arg
}
func Int32StreamFrom(arg []int32) Int32Stream {
	return arg
}
func CreateInt32Stream(arg ...int32) *Int32Stream {
	tmp := Int32StreamOf(arg...)
	return &tmp
}
func GenerateInt32Stream(arg []int32) *Int32Stream {
	tmp := Int32StreamFrom(arg)
	return &tmp
}

func (self *Int32Stream) Add(arg int32) *Int32Stream {
	return self.AddAll(arg)
}
func (self *Int32Stream) AddAll(arg ...int32) *Int32Stream {
	*self = append(*self, arg...)
	return self
}
func (self *Int32Stream) AddSafe(arg *int32) *Int32Stream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *Int32Stream) AllMatch(fn func(int32, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *Int32Stream) AnyMatch(fn func(int32, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *Int32Stream) Clone() *Int32Stream {
	temp := make([]int32, self.Len())
	copy(temp, *self)
	return (*Int32Stream)(&temp)
}
func (self *Int32Stream) Copy() *Int32Stream {
	return self.Clone()
}
func (self *Int32Stream) Concat(arg []int32) *Int32Stream {
	return self.AddAll(arg...)
}
func (self *Int32Stream) Contains(arg int32) bool {
	return self.FindIndex(func(_arg int32, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *Int32Stream) Clean() *Int32Stream {
	*self = Int32StreamOf()
	return self
}
func (self *Int32Stream) Delete(index int) *Int32Stream {
	return self.DeleteRange(index, index)
}
func (self *Int32Stream) DeleteRange(startIndex, endIndex int) *Int32Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *Int32Stream) Distinct() *Int32Stream {
	caches := map[int32]bool{}
	result := Int32StreamOf()
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
func (self *Int32Stream) Each(fn func(int32)) *Int32Stream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *Int32Stream) EachRight(fn func(int32)) *Int32Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *Int32Stream) Equals(arr []int32) bool {
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
func (self *Int32Stream) Where(fn func(int32) bool) *Int32Stream {
	result := Int32StreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *Int32Stream) WhereSlim(fn func(int32) bool) *Int32Stream {
	result := Int32StreamOf()
	caches := map[int32]bool{}
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
func (self *Int32Stream) Filter(fn func(int32, int) bool) *Int32Stream {
	result := Int32StreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *Int32Stream) FilterSlim(fn func(int32, int) bool) *Int32Stream {
	result := Int32StreamOf()
	caches := map[int32]bool{}
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
func (self *Int32Stream) Find(fn func(int32, int) bool) *int32 {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *Int32Stream) FindOr(fn func(int32, int) bool, or int32) int32 {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *Int32Stream) FindIndex(fn func(int32, int) bool) int {
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
func (self *Int32Stream) First() *int32 {
	return self.Get(0)
}
func (self *Int32Stream) FirstOr(arg int32) int32 {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *Int32Stream) ForEach(fn func(int32, int)) *Int32Stream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *Int32Stream) ForEachRight(fn func(int32, int)) *Int32Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *Int32Stream) GroupBy(fn func(int32, int) string) map[string][]int32 {
	m := map[string][]int32{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *Int32Stream) GroupByValues(fn func(int32, int) string) [][]int32 {
	var tmp [][]int32
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *Int32Stream) IndexOf(arg int32) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *Int32Stream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *Int32Stream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *Int32Stream) Last() *int32 {
	return self.Get(self.Len() - 1)
}
func (self *Int32Stream) LastOr(arg int32) int32 {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *Int32Stream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *Int32Stream) Limit(limit int) *Int32Stream {
	self.Slice(0, limit)
	return self
}
func (self *Int32Stream) Select(fn func(int32) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *Int32Stream) Map(fn func(int32, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Int32Stream) Map2Int(fn func(int32, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Int32Stream) Map2Int32(fn func(int32, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Int32Stream) Map2Int64(fn func(int32, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Int32Stream) Map2Float32(fn func(int32, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Int32Stream) Map2Float64(fn func(int32, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Int32Stream) Map2Bool(fn func(int32, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Int32Stream) Map2Bytes(fn func(int32, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Int32Stream) Map2String(fn func(int32, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *Int32Stream) NoneMatch(fn func(int32, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *Int32Stream) Get(index int) *int32 {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *Int32Stream) GetOr(index int, arg int32) int32 {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *Int32Stream) Peek(fn func(*int32, int)) *Int32Stream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}

func (self *Int32Stream) Aggregate(fn func(int32, int32) int32) *Int32Stream {
	result := Int32StreamOf()
	self.ForEach(func(v int32, i int) {
		if i == 0 {
			result.Add(fn(0, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *Int32Stream) Reduce(fn func(int32, int32, int) int32) *Int32Stream {
	return self.ReduceInit(fn, 0)
}
func (self *Int32Stream) ReduceInit(fn func(int32, int32, int) int32, initialValue int32) *Int32Stream {
	result := Int32StreamOf()
	self.ForEach(func(v int32, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *Int32Stream) ReduceInterface(fn func(interface{}, int32, int) interface{}) (result []interface{}) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(struct{}{}, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Int32Stream) ReduceString(fn func(string, int32, int) string) (result []string) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn("", v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Int32Stream) ReduceInt(fn func(int, int32, int) int) (result []int) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Int32Stream) ReduceInt32(fn func(int32, int32, int) int32) (result []int32) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Int32Stream) ReduceInt64(fn func(int64, int32, int) int64) (result []int64) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Int32Stream) ReduceFloat32(fn func(float32, int32, int) float32) (result []float32) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Int32Stream) ReduceFloat64(fn func(float64, int32, int) float64) (result []float64) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Int32Stream) ReduceBool(fn func(bool, int32, int) bool) (result []bool) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(false, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *Int32Stream) Reverse() *Int32Stream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *Int32Stream) Replace(fn func(int32, int) int32) *Int32Stream {
	return self.ForEach(func(v int32, i int) { self.Set(i, fn(v, i)) })
}
func (self *Int32Stream) Set(index int, val int32) *Int32Stream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *Int32Stream) Skip(skip int) *Int32Stream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *Int32Stream) SkippingEach(fn func(int32, int) int) *Int32Stream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *Int32Stream) Slice(startIndex, n int) *Int32Stream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []int32{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *Int32Stream) Sort(fn func(i, j int) bool) *Int32Stream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *Int32Stream) Tail() *int32 {
	return self.Last()
}
func (self *Int32Stream) TailOr(arg int32) int32 {
	return self.LastOr(arg)
}
func (self *Int32Stream) ToList() []int32 {
	return self.Val()
}
func (self *Int32Stream) Unique() *Int32Stream {
	return self.Distinct()
}
func (self *Int32Stream) Val() []int32 {
	if self == nil {
		return []int32{}
	}
	return *self.Copy()
}
func (self *Int32Stream) While(fn func(int32, int) bool) *Int32Stream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}

func (self *Int32Stream) Min() int32 {
	return xmath.MinInt32(self.Val()...)
}

func (self *Int32Stream) Max() int32 {
	return xmath.MaxInt32(self.Val()...)
}

func (self *Int32Stream) Sum() int32 {
	return xmath.SumInt32(self.Val()...)
}

func (self *Int32Stream) Average() int32 {
	return xmath.AverageInt32(self.Val()...)
}
func (self *Int32Stream) Median() *int32 {
	return xmath.MedianInt32(self.Val()...)
}
func (self *Int32Stream) Mode() *int32 {
	return xmath.ModeInt32(self.Val()...)
}
func (self *Int32Stream) Range() *int32 {
	return xmath.RangeInt32(self.Val()...)
}
func (self *Int32Stream) AsFloat64() *Float64Stream {
	result := self.Map2Float64(func(arg int32, index int) float64 { return float64(arg) })
	stream := Float64StreamFrom(result)
	return &stream
}

func (self *Int32Stream) AsInt() *IntegerStream {
	result := self.Map2Int(func(arg int32, index int) int { return int(arg) })
	stream := IntegerStreamFrom(result)
	return &stream
}
func (self *Int32Stream) AsInt32() *Int32Stream {
	result := self.Map2Int32(func(arg int32, index int) int32 { return int32(arg) })
	stream := Int32StreamFrom(result)
	return &stream
}

func (self *Int32Stream) AsInt64() *Int64Stream {
	result := self.Map2Int64(func(arg int32, index int) int64 { return int64(arg) })
	stream := Int64StreamFrom(result)
	return &stream
}
func (self *Int32Stream) ContainsSome(arg ...int32) bool {
	for _, v := range arg {
		if self.Contains(v) {
			return true
		}
	}
	return false
}
func (self *Int32Stream) ContainsAll(arg ...int32) bool {
	for _, v := range arg {
		if !self.Contains(v) {
			return false
		}
	}
	return true
}
