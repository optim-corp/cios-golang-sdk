package ciosctx

import (
	"context"
	"regexp"

	"github.com/fcfcqloow/go-advance/check"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

type RequestCtx context.Context

func Background() RequestCtx {
	return context.Background()
}

func WithToken(ctx RequestCtx, token string) RequestCtx {
	if check.IsNil(ctx) {
		ctx = Background()
	}
	if token == "" {
		return ctx
	}
	return context.WithValue(ctx, cios.ContextAccessToken, regexp.MustCompile(`^bearer|Bearer| `).ReplaceAllString(token, ""))
}

func Token(ctx RequestCtx) (token *string) {
	if !check.IsNil(ctx) {
		if _token, ok := ctx.Value(cios.ContextAccessToken).(string); ok {
			token = &_token
		}
	}
	return
}
