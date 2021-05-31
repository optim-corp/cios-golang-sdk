package util

import (
	"reflect"
	"testing"

	cnv "github.com/fcfcqloow/go-advance/convert"
)

func TestToNil(t *testing.T) {
	type args struct {
		str *string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{

		{
			name: "Test Nil",
			args: args{
				str: nil,
			},
			want: nil,
		},
		{
			name: "Test Empty string",
			args: args{
				str: cnv.StrPtr(""),
			},
			want: nil,
		},
		{
			name: "Test Preset string",
			args: args{
				str: cnv.StrPtr("test"),
			},
			want: cnv.StrPtr("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToNil(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToNilStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "Test Nil",
			args: args{
				str: "",
			},
			want: nil,
		},
		{
			name: "Test Preset string",
			args: args{
				str: ("test"),
			},
			want: cnv.StrPtr("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToNilStr(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToNilStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
