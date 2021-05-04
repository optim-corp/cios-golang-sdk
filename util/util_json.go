package util

import (
	"strings"

	cnv "github.com/fcfcqloow/go-advance/convert"
	"github.com/optim-corp/cios-golang-sdk/cios"
)

func DataStoreStreamToStruct(str []string, st interface{}) error {
	value := "[" + strings.Join(str, ",") + "]"
	return cnv.UnMarshalJson(value, st)
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
	byts, err := cnv.MarshalJson(arr)
	if err != nil {
		return nil, err
	}
	return headers, cnv.UnMarshalJson(byts, stc)
}
