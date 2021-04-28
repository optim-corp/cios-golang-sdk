package wrapper

import (
	"crypto/md5"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/optim-corp/cios-golang-sdk/util/go_advance_type/collection"
	"github.com/optim-corp/cios-golang-sdk/util/go_advance_type/convert"

	"github.com/iancoleman/strcase"
)

func AsString(str string) String {
	return String(str)
}
func MakeString(str string) *String {
	ins := AsString(str)
	return &ins
}
func NewString() *String {
	return MakeString("")
}
func (self String) AsWrap() *String {
	return &self
}

func (self *String) Add(str string) *String {
	return self.AddAll(str)
}
func (self *String) AddHead(str string) *String {
	return self.AddAllHead(str)
}
func (self *String) AddAllHead(arg ...string) *String {
	*self = String(strings.Join(collection.CreateStringStream(arg...).Reverse().Add(self.AsPrimitive()).ToList(), ""))
	return self
}
func (self *String) AddAll(arg ...string) *String {
	*self = String(strings.Join(collection.CreateStringStream(self.AsPrimitive()).AddAll(arg...).ToList(), ""))
	return self
}
func (self *String) AsPrimitive() string {
	return string(*self)
}
func (self *String) AsRunes() []rune {
	return []rune(*self)
}
func (self *String) AsBytes() []byte {
	return []byte(*self)
}
func (self *String) Copy() *String {
	return MakeString(self.AsPrimitive())
}
func (self *String) ChatAt(index int) Char {
	return Char(self.AsPrimitive()[index])
}
func (self *String) Default(str string) *String {
	if self.IsEmpty() {
		*self = String(str)
	}
	return self
}
func (self *String) DeleteDuplicate(str string) *String {
	*self = String(regexp.MustCompile(`\s+`).ReplaceAllString(self.AsPrimitive(), str))
	return self
}
func (self *String) DeleteDupSpace() *String {
	*self = String(regexp.MustCompile(`\s+`).ReplaceAllString(self.AsPrimitive(), " "))
	return self
}
func (self *String) Delete(str string, n int) *String {
	if n <= -1 {
		return self.DeleteAll(str)
	} else if MakeString(str).IsPreset() && self.IsPreset() {
		count := 0
		for index := strings.Index(self.AsPrimitive(), str); index > -1 && count < n; index = strings.Index(self.AsPrimitive(), str) {
			*self = String(self.AsPrimitive()[:index] + self.AsPrimitive()[index+len(str):])
			count++
		}
	}
	return self
}
func (self *String) DeleteAll(str string) *String {
	*self = String(strings.ReplaceAll(self.AsPrimitive(), str, ""))
	return self
}
func (self *String) DeletePattern(str string) *String {
	*self = String(regexp.MustCompile(str).ReplaceAllString(self.AsPrimitive(), ""))
	return self
}
func (self *String) Equals(target string) bool {
	return self.AsPrimitive() == target
}
func (self *String) Clean() *String {
	*self = ""
	return self
}
func (self *String) Concat(arg *String) *String {
	return self.AddAll(arg.AsPrimitive())
}
func (self *String) ConcatHead(arg *String) *String {
	return self.AddAllHead(arg.AsPrimitive())
}
func (self *String) Contains(str string) bool {
	return strings.Contains(self.AsPrimitive(), str)
}
func (self *String) ContainsSome(str ...string) bool {
	for _, v := range str {
		if self.Contains(v) {
			return true
		}
	}
	return false
}
func (self *String) ContainsAll(str ...string) bool {
	for _, v := range str {
		if !self.Contains(v) {
			return false
		}

	}
	return true
}
func (self *String) Encode() Resource {
	return convert.DecodeBase64(self.AsPrimitive())
}
func (self *String) EndWith(str string) bool {
	return strings.HasSuffix(self.AsPrimitive(), str)
}
func (self *String) IsEmpty() bool {
	if self == nil {
		return false
	}
	return len(*self) == 0

}
func (self *String) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *String) IsBlank() bool {
	return regexp.MustCompile(`^\s+$`).MatchString(self.AsPrimitive())
}
func (self *String) IsNumeric() bool {
	_, err := strconv.ParseFloat(self.AsPrimitive(), 64)
	return err == nil
}
func (self *String) Len() int {
	return len(*self)
}

func (self *String) LenUtf8() int {
	return utf8.RuneCountInString(self.AsPrimitive())
}

func (self *String) MathPattern(pattern string) bool {
	return regexp.MustCompile(pattern).MatchString(self.AsPrimitive())
}
func (self *String) Merge(arg String) *String {
	return self.AddAll(string(arg))
}
func (self *String) MergeHead(arg String) *String {
	return self.AddAllHead(string(arg))
}
func (self *String) Replace(src, target string, count int) *String {
	if count < -1 {
		count = -1
	}
	*self = String(strings.Replace(self.AsPrimitive(), src, target, count))
	return self
}

func (self *String) ReplaceAll(src, target string) *String {
	*self = String(strings.ReplaceAll(self.AsPrimitive(), src, target))
	return self
}
func (self *String) Reverse() *String {
	r := []rune(self.AsPrimitive())
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	*self = String(r)
	return self
}
func (self *String) Str() string {
	return self.AsPrimitive()
}
func (self *String) String() string {
	return self.AsPrimitive()
}
func (self *String) Split(str string) *collection.StringStream {
	return collection.CreateStringStream(strings.Split(self.AsPrimitive(), str)...)
}
func (self *String) Slice(startIndex, n int) *String {
	last := startIndex + n
	if self.Len()-1 < startIndex {
		*self = ""
	} else if self.Len() < last {
		*self = String(self.AsPrimitive()[startIndex:len(*self)])
	} else {
		*self = String(self.AsPrimitive()[startIndex:last])
	}
	return self
}
func (self *String) SliceMatched(pattern string) *collection.StringStream {
	return collection.CreateStringStream(regexp.MustCompile(pattern).FindAllString(self.AsPrimitive(), -1)...)
}
func (self *String) StartWith(str string) bool {
	return strings.HasPrefix(self.AsPrimitive(), str)
}
func (self *String) ToLowerCase() *String {
	*self = String(strings.ToLower(self.AsPrimitive()))
	return self
}
func (self *String) ToUpperCase() *String {
	*self = String(strings.ToUpper(self.AsPrimitive()))
	return self
}
func (self *String) ToCamel() *String {
	*self = String(strcase.ToCamel(self.AsPrimitive()))
	return self
}
func (self *String) ToSnake() *String {
	*self = String(strcase.ToSnake(self.AsPrimitive()))
	return self
}
func (self *String) ToKebab() *String {
	*self = String(strcase.ToKebab(self.AsPrimitive()))
	return self
}
func (self *String) ToLowerCamel() *String {
	*self = String(strcase.ToLowerCamel(self.AsPrimitive()))
	return self
}
func (self *String) Trim() *String {
	*self = String(strings.Trim(self.AsPrimitive(), " "))
	return self
}
func (self *String) ToMd5() []byte {
	h := md5.New()
	h.Write([]byte(self.AsPrimitive()))
	return h.Sum(nil)
}
