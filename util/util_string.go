package util

import "github.com/optim-kazuhiro-seida/go-advance-type/check"

func ToNil(str *string) *string {
	if check.IsNil(str) || str == nil {
		return nil
	} else if *str == "" {
		return nil
	}
	return str
}
func ToNilStr(str string) *string {
	if str == ""  {
		return nil
	}
	return &str
}
