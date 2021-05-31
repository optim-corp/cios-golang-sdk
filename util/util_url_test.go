package util_test

import (
	"net/url"
	"testing"

	cnv "github.com/fcfcqloow/go-advance/convert"

	"github.com/optim-corp/cios-golang-sdk/util"
)

func TestQuery(t *testing.T) {
	type args struct {
		_url *url.URL
		args util.QueryMap
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Query1",
			args: args{
				_url: func() *url.URL {
					_url, _ := url.Parse("http://test.example.com")
					return _url
				}(),
				args: util.QueryMap{
					"sample": nil,
					"test1":  nil,
				},
			},
			want: "",
		},
		{
			name: "Test Query2",
			args: args{
				_url: func() *url.URL {
					_url, _ := url.Parse("http://test.example.com")
					return _url
				}(),
				args: util.QueryMap{
					"sample": cnv.StrPtr("a"),
					"test1":  nil,
				},
			},
			want: "sample=a",
		},
		{
			name: "Test Query3",
			args: args{
				_url: func() *url.URL {
					_url, _ := url.Parse("http://test.example.com")
					return _url
				}(),
				args: util.QueryMap{
					"sample": cnv.StrPtr("a"),
					"test1":  cnv.StrPtr("b"),
				},
			},
			want: "sample=a&test1=b",
		},
		{
			name: "Test Query4",
			args: args{
				_url: func() *url.URL {
					_url, _ := url.Parse("http://test.example.com?test2=b")
					return _url
				}(),
				args: util.QueryMap{
					"sample": cnv.StrPtr("a"),
					"test1":  cnv.StrPtr("b"),
				},
			},
			want: "test2=b&sample=a&test1=b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.Query(tt.args._url, tt.args.args); got != tt.want {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}
