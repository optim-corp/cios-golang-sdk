package ciosutil_test

import (
	"reflect"
	"testing"

	cnv "github.com/fcfcqloow/go-advance/convert"
	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosutil "github.com/optim-corp/cios-golang-sdk/util/cios"
	"github.com/stretchr/testify/assert"
)

func TestDataStoreStreamToStruct(t *testing.T) {
	type (
		inside struct {
			A string `json:"a,omitempty"`
			B int    `json:"b,omitempty"`
			C bool   `json:"c,omitempty"`
		}
		testStruct struct {
			Test  string `json:"test,omitempty"`
			Test2 int    `json:"test_2,omitempty"`
			Test3 bool   `json:"test_3,omitempty"`
			Test4 inside `json:"test_4"`
		}
		args struct {
			str []string
			st  interface{}
		}
		test struct {
			name    string
			args    args
			wantErr bool
			result  string
		}
	)
	tests := []test{
		{
			name: "Empty Test",
			args: args{
				str: []string{},
				st:  []testStruct{},
			},
			result:  "[]",
			wantErr: false,
		},
		{
			name: "Array Test 1",
			args: args{
				str: []string{cnv.MustCompactJson(testStruct{}), cnv.MustCompactJson(testStruct{}), cnv.MustCompactJson(testStruct{})},
				st:  []testStruct{},
			},
			result:  "[{\"test_4\":{}},{\"test_4\":{}},{\"test_4\":{}}]",
			wantErr: false,
		},
		{
			name: "Array Test 2",
			args: args{
				str: []string{cnv.MustCompactJson(testStruct{
					Test:  "test",
					Test2: 10,
				}), cnv.MustCompactJson(testStruct{}), cnv.MustCompactJson(testStruct{})},
				st: []testStruct{},
			},
			result:  "[{\"test\":\"test\",\"test_2\":10,\"test_4\":{}},{\"test_4\":{}},{\"test_4\":{}}]",
			wantErr: false,
		},
		{
			name: "Array Test 2",
			args: args{
				str: []string{cnv.MustCompactJson(testStruct{
					Test:  "test",
					Test2: 10,
					Test3: true,
				}), cnv.MustCompactJson(testStruct{}), cnv.MustCompactJson(testStruct{})},
				st: []testStruct{},
			},
			result:  "[{\"test\":\"test\",\"test_2\":10,\"test_3\":true,\"test_4\":{}},{\"test_4\":{}},{\"test_4\":{}}]",
			wantErr: false,
		},
		{
			name: "Array Test 2",
			args: args{
				str: []string{cnv.MustCompactJson(testStruct{
					Test:  "test",
					Test2: 10,
					Test3: true,
					Test4: inside{A: "a", B: 2, C: true},
				}), cnv.MustCompactJson(testStruct{}), cnv.MustCompactJson(testStruct{})},
				st: []testStruct{},
			},
			result:  "[{\"test\":\"test\",\"test_2\":10,\"test_3\":true,\"test_4\":{\"a\":\"a\",\"b\":2,\"c\":true}},{\"test_4\":{}},{\"test_4\":{}}]",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ciosutil.DataStoreStreamToStruct(tt.args.str, &tt.args.st); (err != nil) != tt.wantErr {
				t.Errorf("DataStoreStreamToStruct() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Equal(t, tt.result, cnv.MustCompactJson(tt.args.st))
			}
		})
	}
}

func TestMapPayloads(t *testing.T) {
	type (
		args struct {
			objects []cios.PackerFormatJson
			stc     interface{}
		}
		test struct {
			name    string
			args    args
			want    []cios.PackerFormatJsonHeader
			wantErr bool
		}
		testPayload struct {
			Test1 string `json:"test_1,omitempty"`
			Test2 int    `json:"test_2,omitempty"`
		}
	)
	tests := []test{
		{
			name: "Test MapPayloads",
			args: args{
				objects: []cios.PackerFormatJson{{
					Header:  cios.PackerFormatJsonHeader{},
					Payload: testPayload{},
				}},
				stc: []testPayload{},
			},
			want:    []cios.PackerFormatJsonHeader{{}},
			wantErr: false,
		},
		{
			name: "Nil Payload MapPayloads",
			args: args{
				objects: []cios.PackerFormatJson{{
					Header:  cios.PackerFormatJsonHeader{},
					Payload: testPayload{},
				}},
				stc: nil,
			},
			want:    []cios.PackerFormatJsonHeader{{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ciosutil.MapPayloads(tt.args.objects, &tt.args.stc)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapPayloads() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapPayloads() got = %v, want %v", got, tt.want)
			}
		})
	}
}
