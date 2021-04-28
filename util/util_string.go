package util

import "github.com/optim-corp/cios-golang-sdk/util/go_advance_type/check"

func ToNil(str *string) *string {
	if check.IsNil(str) || str == nil {
		return nil
	} else if *str == "" {
		return nil
	}
	return str
}
func ToNilStr(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}
