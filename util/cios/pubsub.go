package ciosutil

import (
	"fmt"
	"strings"

	cnv "github.com/fcfcqloow/go-advance/convert"
	"github.com/optim-corp/cios-golang-sdk/cios"
)

func DataStoreStreamToStruct(str []string, st interface{}) error {
	return cnv.UnMarshalJson(fmt.Sprintf("[%s]", strings.Join(str, ",")), st)
}

func MapPayloads(objects []cios.PackerFormatJson, stc interface{}) ([]cios.PackerFormatJsonHeader, error) {
	var (
		arr     []interface{}
		headers []cios.PackerFormatJsonHeader
	)
	for _, obj := range objects {
		if obj.Payload != nil {
			headers = append(headers, obj.Header)
			arr = append(arr, obj.Payload)
		}
	}
	return headers, cnv.DeepCopy(arr, stc)
}
