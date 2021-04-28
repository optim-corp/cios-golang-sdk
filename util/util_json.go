package util

import (
	"strings"

	"github.com/optim-corp/cios-golang-sdk/cios"

	"github.com/optim-corp/cios-golang-sdk/util/go_advance_type/convert"
)

func DataStoreStreamToStruct(str []string, st interface{}) error {
	value := "[" + strings.Join(str, ",") + "]"
	return convert.UnMarshalJson(value, st)
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
	byts, err := convert.MarshalJson(arr)
	if err != nil {
		return nil, err
	}
	return headers, convert.UnMarshalJson(byts, stc)
}
