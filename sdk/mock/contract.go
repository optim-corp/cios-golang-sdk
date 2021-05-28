package ciossdk_mock

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
)

type NoImplementContract struct{}

func (NoImplementContract) GetContracts(ctx ciosctx.RequestCtx, request cios.ApiGetContractsRequest) (cios.MultipleContract, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementContract) GetContractsAll(ctx ciosctx.RequestCtx, request cios.ApiGetContractsRequest) ([]cios.Contract, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementContract) GetContractsUnlimited(ctx ciosctx.RequestCtx, request cios.ApiGetContractsRequest) ([]cios.Contract, *_nethttp.Response, error) {
	panic("implement me")
}
