package ciosctx_test

import (
	"context"
	"reflect"
	"testing"

	cnv "github.com/fcfcqloow/go-advance/convert"

	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
)

func TestBackground(t *testing.T) {
	tests := []struct {
		name string
		want ciosctx.RequestCtx
	}{
		{
			name: "Test Background",
			want: context.Background(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ciosctx.Background(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Background() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToken(t *testing.T) {
	type args struct {
		ctx ciosctx.RequestCtx
	}
	tests := []struct {
		name      string
		args      args
		wantToken *string
	}{
		{
			name: "Test CTX Token",
			args: args{
				ctx: context.Background(),
			},
			wantToken: nil,
		},
		{
			name: "Test CTX Token",
			args: args{
				ctx: func() ciosctx.RequestCtx {
					return ciosctx.WithToken(ciosctx.Background(), "testtest")
				}(),
			},
			wantToken: cnv.StrPtr("testtest"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotToken := ciosctx.Token(tt.args.ctx); !reflect.DeepEqual(gotToken, tt.wantToken) {
				t.Errorf("Token() = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}

func TestWithToken(t *testing.T) {
	type args struct {
		ctx   ciosctx.RequestCtx
		token string
	}
	tests := []struct {
		name string
		args args
		want ciosctx.RequestCtx
	}{
		{
			name: "Test With Token",
			args: args{
				ctx:   ciosctx.Background(),
				token: "Test Token",
			},
			want: func() ciosctx.RequestCtx {
				return ciosctx.WithToken(context.Background(), "Test Token")
				// TODO: 考える
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ciosctx.WithToken(tt.args.ctx, tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
