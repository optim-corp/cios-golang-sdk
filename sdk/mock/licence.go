package ciossdmock

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
)

type NoImplementLicense struct{}

func (NoImplementLicense) GetLicenses(ctx ciosctx.RequestCtx, request cios.ApiGetMyLicensesRequest) ([]cios.License, *_nethttp.Response, error) {
	panic("implement me")
}
