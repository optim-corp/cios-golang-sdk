package collection

import (
	"reflect"
	"sort"
	"strings"

	xmath "github.com/optim-corp/cios-golang-sdk/util/go_advance_type/math"
)

func StringStreamOf(arg ...string) StringStream {
	return arg
}
func StringStreamFrom(arg []string) StringStream {
	return arg
}
func CreateStringStream(arg ...string) *StringStream {
	tmp := StringStreamOf(arg...)
	return &tmp
}
func GenerateStringStream(arg []string) *StringStream {
	tmp := StringStreamFrom(arg)
	return &tmp
}

func (self *StringStream) Add(arg string) *StringStream {
	return self.AddAll(arg)
}
func (self *StringStream) AddAll(arg ...string) *StringStream {
	*self = append(*self, arg...)
	return self
}
func (self *StringStream) AddSafe(arg *string) *StringStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *StringStream) AllMatch(fn func(string, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *StringStream) AnyMatch(fn func(string, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *StringStream) Clone() *StringStream {
	temp := make([]string, self.Len())
	copy(temp, *self)
	return (*StringStream)(&temp)
}
func (self *StringStream) Copy() *StringStream {
	return self.Clone()
}
func (self *StringStream) Concat(arg []string) *StringStream {
	return self.AddAll(arg...)
}
func (self *StringStream) Contains(arg string) bool {
	return self.FindIndex(func(_arg string, index int) bool { return _arg == arg }) != -1
}
func (self *StringStream) Clean() *StringStream {
	*self = StringStreamOf()
	return self
}
func (self *StringStream) Delete(index int) *StringStream {
	return self.DeleteRange(index, index)
}
func (self *StringStream) DeleteRange(startIndex, endIndex int) *StringStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *StringStream) Distinct() *StringStream {
	caches := map[string]bool{}
	result := StringStreamOf()
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
func (self *StringStream) Each(fn func(string)) *StringStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *StringStream) EachRight(fn func(string)) *StringStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *StringStream) Equals(arr []string) bool {
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
func (self *StringStream) Where(fn func(string) bool) *StringStream {
	result := StringStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *StringStream) WhereSlim(fn func(string) bool) *StringStream {
	result := StringStreamOf()
	caches := map[string]bool{}
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
func (self *StringStream) Filter(fn func(string, int) bool) *StringStream {
	result := StringStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *StringStream) FilterSlim(fn func(string, int) bool) *StringStream {
	result := StringStreamOf()
	caches := map[string]bool{}
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
func (self *StringStream) Find(fn func(string, int) bool) *string {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *StringStream) FindOr(fn func(string, int) bool, or string) string {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *StringStream) FindIndex(fn func(string, int) bool) int {
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
func (self *StringStream) First() *string {
	return self.Get(0)
}
func (self *StringStream) FirstOr(arg string) string {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *StringStream) ForEach(fn func(string, int)) *StringStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *StringStream) ForEachRight(fn func(string, int)) *StringStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *StringStream) GroupBy(fn func(string, int) string) map[string][]string {
	m := map[string][]string{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *StringStream) GroupByValues(fn func(string, int) string) [][]string {
	var tmp [][]string
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *StringStream) IndexOf(arg string) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *StringStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *StringStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *StringStream) Join(arg []string, sep string) string {
	self.Concat(arg)
	return strings.Join(*self, sep)
}
func (self *StringStream) Last() *string {
	return self.Get(self.Len() - 1)
}
func (self *StringStream) LastOr(arg string) string {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *StringStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *StringStream) Limit(limit int) *StringStream {
	self.Slice(0, limit)
	return self
}
func (self *StringStream) Select(fn func(string) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *StringStream) Map(fn func(string, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *StringStream) Map2Int(fn func(string, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *StringStream) Map2Int32(fn func(string, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *StringStream) Map2Int64(fn func(string, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *StringStream) Map2Float32(fn func(string, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *StringStream) Map2Float64(fn func(string, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *StringStream) Map2Bool(fn func(string, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *StringStream) Map2Bytes(fn func(string, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *StringStream) Map2String(fn func(string, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *StringStream) Max(fn func(string, int) float64) string {
	return xmath.MaxStr(self.Val()...)
}
func (self *StringStream) Min(fn func(string, int) float64) string {
	return xmath.MinStr(self.Val()...)
}
func (self *StringStream) NoneMatch(fn func(string, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *StringStream) Get(index int) *string {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *StringStream) GetOr(index int, arg string) string {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *StringStream) Peek(fn func(*string, int)) *StringStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}
func (self *StringStream) Aggregate(fn func(string, string) string) *StringStream {
	result := StringStreamOf()
	self.ForEach(func(v string, i int) {
		if i == 0 {
			result.Add(fn("", v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *StringStream) Reduce(fn func(string, string, int) string) *StringStream {
	return self.ReduceInit(fn, "")
}
func (self *StringStream) ReduceInit(fn func(string, string, int) string, initialValue string) *StringStream {
	result := StringStreamOf()
	self.ForEach(func(v string, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *StringStream) ReduceString(fn func(string, string, int) string) (result []string) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn("", v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *StringStream) ReduceInt(fn func(int, string, int) int) (result []int) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *StringStream) ReduceInt32(fn func(int32, string, int) int32) (result []int32) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *StringStream) ReduceInt64(fn func(int64, string, int) int64) (result []int64) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *StringStream) ReduceFloat32(fn func(float32, string, int) float32) (result []float32) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *StringStream) ReduceFloat64(fn func(float64, string, int) float64) (result []float64) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *StringStream) ReduceBool(fn func(bool, string, int) bool) (result []bool) {
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(false, v, i))
		} else if len(result) > 0 {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return
}
func (self *StringStream) Reverse() *StringStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *StringStream) Replace(fn func(string, int) string) *StringStream {
	return self.ForEach(func(v string, i int) { self.Set(i, fn(v, i)) })
}
func (self *StringStream) Set(index int, val string) *StringStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *StringStream) Skip(skip int) *StringStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *StringStream) SkippingEach(fn func(string, int) int) *StringStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *StringStream) Slice(startIndex, n int) *StringStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []string{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *StringStream) Sort(fn func(i, j int) bool) *StringStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *StringStream) Tail() *string {
	return self.Last()
}
func (self *StringStream) TailOr(arg string) string {
	return self.LastOr(arg)
}
func (self *StringStream) ToList() []string {
	return self.Val()
}
func (self *StringStream) Unique() *StringStream {
	return self.Distinct()
}
func (self *StringStream) Val() []string {
	if self == nil {
		return []string{}
	}
	return *self.Copy()
}
func (self *StringStream) While(fn func(string, int) bool) *StringStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}

func (self *StringStream) Sum() string {
	return self.Join(nil, "")
}

func (self *StringStream) ContainsSome(arg ...string) bool {
	for _, v := range arg {
		if self.Contains(v) {
			return true
		}
	}
	return false
}
func (self *StringStream) ContainsAll(arg ...string) bool {
	for _, v := range arg {
		if !self.Contains(v) {
			return false
		}
	}
	return true
}
