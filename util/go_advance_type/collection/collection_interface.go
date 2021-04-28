package collection

import (
	"reflect"
	"sort"
)

func InterfaceStreamOf(arg ...interface{}) InterfaceStream {
	return arg
}
func InterfaceStreamFrom(arg []interface{}) InterfaceStream {
	return arg
}
func CreateInterfaceStream(arg ...interface{}) *InterfaceStream {
	tmp := InterfaceStreamOf(arg...)
	return &tmp
}
func GenerateInterfaceStream(arg []interface{}) *InterfaceStream {
	tmp := InterfaceStreamFrom(arg)
	return &tmp
}

func (self *InterfaceStream) Add(arg interface{}) *InterfaceStream {
	return self.AddAll(arg)
}
func (self *InterfaceStream) AddAll(arg ...interface{}) *InterfaceStream {
	*self = append(*self, arg...)
	return self
}
func (self *InterfaceStream) AddSafe(arg *interface{}) *InterfaceStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *InterfaceStream) AllMatch(fn func(interface{}, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *InterfaceStream) AnyMatch(fn func(interface{}, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *InterfaceStream) Clone() *InterfaceStream {
	temp := make([]interface{}, self.Len())
	copy(temp, *self)
	return (*InterfaceStream)(&temp)
}
func (self *InterfaceStream) Copy() *InterfaceStream {
	return self.Clone()
}
func (self *InterfaceStream) Concat(arg []interface{}) *InterfaceStream {
	return self.AddAll(arg...)
}
func (self *InterfaceStream) Contains(arg interface{}) bool {
	return self.FindIndex(func(_arg interface{}, index int) bool { return _arg == arg }) != -1
}
func (self *InterfaceStream) Clean() *InterfaceStream {
	*self = InterfaceStreamOf()
	return self
}
func (self *InterfaceStream) Delete(index int) *InterfaceStream {
	return self.DeleteRange(index, index)
}
func (self *InterfaceStream) DeleteRange(startIndex, endIndex int) *InterfaceStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *InterfaceStream) Distinct() *InterfaceStream {
	caches := map[interface{}]bool{}
	result := InterfaceStreamOf()
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
func (self *InterfaceStream) Each(fn func(interface{})) *InterfaceStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *InterfaceStream) EachRight(fn func(interface{})) *InterfaceStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *InterfaceStream) Equals(arr []interface{}) bool {
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
func (self *InterfaceStream) Where(fn func(interface{}) bool) *InterfaceStream {
	result := InterfaceStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *InterfaceStream) WhereSlim(fn func(interface{}) bool) *InterfaceStream {
	result := InterfaceStreamOf()
	caches := map[interface{}]bool{}
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
func (self *InterfaceStream) Filter(fn func(interface{}, int) bool) *InterfaceStream {
	result := InterfaceStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *InterfaceStream) FilterSlim(fn func(interface{}, int) bool) *InterfaceStream {
	result := InterfaceStreamOf()
	caches := map[interface{}]bool{}
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
func (self *InterfaceStream) Find(fn func(interface{}, int) bool) *interface{} {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *InterfaceStream) FindOr(fn func(interface{}, int) bool, or interface{}) interface{} {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *InterfaceStream) FindIndex(fn func(interface{}, int) bool) int {
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
func (self *InterfaceStream) First() *interface{} {
	return self.Get(0)
}
func (self *InterfaceStream) FirstOr(arg interface{}) interface{} {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *InterfaceStream) ForEach(fn func(interface{}, int)) *InterfaceStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *InterfaceStream) ForEachRight(fn func(interface{}, int)) *InterfaceStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *InterfaceStream) GroupBy(fn func(interface{}, int) interface{}) map[interface{}][]interface{} {
	m := map[interface{}][]interface{}{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *InterfaceStream) GroupByValues(fn func(interface{}, int) interface{}) [][]interface{} {
	var tmp [][]interface{}
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *InterfaceStream) IndexOf(arg interface{}) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *InterfaceStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *InterfaceStream) IsPreset() bool {
	return !self.IsEmpty()
}

func (self *InterfaceStream) Last() *interface{} {
	return self.Get(self.Len() - 1)
}
func (self *InterfaceStream) LastOr(arg interface{}) interface{} {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *InterfaceStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *InterfaceStream) Limit(limit int) *InterfaceStream {
	self.Slice(0, limit)
	return self
}
func (self *InterfaceStream) Select(fn func(interface{}) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *InterfaceStream) Map(fn func(interface{}, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *InterfaceStream) Map2Int(fn func(interface{}, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *InterfaceStream) Map2Int32(fn func(interface{}, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *InterfaceStream) Map2Int64(fn func(interface{}, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *InterfaceStream) Map2Float32(fn func(interface{}, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *InterfaceStream) Map2Float64(fn func(interface{}, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *InterfaceStream) Map2Bool(fn func(interface{}, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *InterfaceStream) Map2Bytes(fn func(interface{}, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *InterfaceStream) Map2Interface(fn func(interface{}, int) interface{}) []interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *InterfaceStream) NoneMatch(fn func(interface{}, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *InterfaceStream) Get(index int) *interface{} {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *InterfaceStream) GetOr(index int, arg interface{}) interface{} {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *InterfaceStream) Peek(fn func(*interface{}, int)) *InterfaceStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}
func (self *InterfaceStream) Aggregate(fn func(interface{}, interface{}) interface{}) *InterfaceStream {
	result := InterfaceStreamOf()
	self.ForEach(func(v interface{}, i int) {
		if i == 0 {
			result.Add(fn("", v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *InterfaceStream) Reduce(fn func(interface{}, interface{}, int) interface{}) *InterfaceStream {
	return self.ReduceInit(fn, "")
}
func (self *InterfaceStream) ReduceInit(fn func(interface{}, interface{}, int) interface{}, initialValue interface{}) *InterfaceStream {
	result := InterfaceStreamOf()
	self.ForEach(func(v interface{}, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *InterfaceStream) ReduceInterface(fn func(interface{}, interface{}, int) interface{}) (result []interface{}) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn("", v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *InterfaceStream) ReduceInt(fn func(int, interface{}, int) int) (result []int) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *InterfaceStream) ReduceInt32(fn func(int32, interface{}, int) int32) (result []int32) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *InterfaceStream) ReduceInt64(fn func(int64, interface{}, int) int64) (result []int64) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *InterfaceStream) ReduceFloat32(fn func(float32, interface{}, int) float32) (result []float32) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *InterfaceStream) ReduceFloat64(fn func(float64, interface{}, int) float64) (result []float64) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *InterfaceStream) ReduceBool(fn func(bool, interface{}, int) bool) (result []bool) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(false, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *InterfaceStream) Reverse() *InterfaceStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *InterfaceStream) Replace(fn func(interface{}, int) interface{}) *InterfaceStream {
	return self.ForEach(func(v interface{}, i int) { self.Set(i, fn(v, i)) })
}
func (self *InterfaceStream) Set(index int, val interface{}) *InterfaceStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *InterfaceStream) Skip(skip int) *InterfaceStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *InterfaceStream) SkippingEach(fn func(interface{}, int) int) *InterfaceStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *InterfaceStream) Slice(startIndex, n int) *InterfaceStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []interface{}{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *InterfaceStream) Sort(fn func(i, j int) bool) *InterfaceStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *InterfaceStream) Tail() *interface{} {
	return self.Last()
}
func (self *InterfaceStream) TailOr(arg interface{}) interface{} {
	return self.LastOr(arg)
}
func (self *InterfaceStream) ToList() []interface{} {
	return self.Val()
}
func (self *InterfaceStream) Unique() *InterfaceStream {
	return self.Distinct()
}
func (self *InterfaceStream) Val() []interface{} {
	if self == nil {
		return []interface{}{}
	}
	return *self.Copy()
}
func (self *InterfaceStream) While(fn func(interface{}, int) bool) *InterfaceStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}

func (self *InterfaceStream) ContainsSome(arg ...interface{}) bool {
	for _, v := range arg {
		if self.Contains(v) {
			return true
		}
	}
	return false
}
func (self *InterfaceStream) ContainsAll(arg ...interface{}) bool {
	for _, v := range arg {
		if !self.Contains(v) {
			return false
		}
	}
	return true
}
