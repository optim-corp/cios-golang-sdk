package ciosctx

import (
	"reflect"
	"testing"
)

func TestBackground(t *testing.T) {
	tests := []struct {
		name string
		want RequestCtx
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Background(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Background() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToken(t *testing.T) {
	type args struct {
		ctx RequestCtx
	}
	tests := []struct {
		name      string
		args      args
		wantToken *string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotToken := Token(tt.args.ctx); !reflect.DeepEqual(gotToken, tt.wantToken) {
				t.Errorf("Token() = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}

func TestWithToken(t *testing.T) {
	type args struct {
		ctx   RequestCtx
		token string
	}
	tests := []struct {
		name string
		args args
		want RequestCtx
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithToken(tt.args.ctx, tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
