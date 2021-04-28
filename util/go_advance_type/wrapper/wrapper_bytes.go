package wrapper

import "github.com/optim-corp/cios-golang-sdk/util/go_advance_type/convert"

func NewResource(byts []byte) *Resource {
	ins := AsResource(byts)
	return &ins
}
func AsResource(byts []byte) Resource {
	return byts
}
func EncodeJson(stc interface{}) (Resource, error) {
	byts, err := convert.MarshalJson(stc)
	return AsResource(byts), err
}

func (self Resource) AsWrap() *Resource {
	return &self
}

func (self *Resource) AsString() *String {
	return MakeString(string(self.AsPrimitive()))
}
func (self *Resource) AsPrimitive() []byte {
	return *self
}
func (self *Resource) Copy() (*Resource, error) {
	var tmp interface{}
	err := self.DecodeJson(&tmp)
	if err != nil {
		return nil, err
	}
	ins, err := EncodeJson(tmp)
	if err != nil {
		return nil, err
	}
	return &ins, err
}
func (self *Resource) DecodeJson(stc interface{}) error {
	return convert.UnMarshalJson(self.AsPrimitive(), stc)
}

func (self *Resource) EncodeBase64() string {
	return convert.EncodeBase64(self.AsPrimitive())
}

func (self *Resource) IndentJson() (string, error) {
	return convert.IndentJson(self.AsPrimitive())
}
