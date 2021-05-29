package ciossdk_mock

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
)

type NoImplementGeography struct{}

func (NoImplementGeography) GetPoints(ctx ciosctx.RequestCtx, request cios.ApiGetPointsRequest) (cios.MultiplePoint, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementGeography) CreatePoint(ctx ciosctx.RequestCtx, request cios.PointRequest) (cios.Point, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementGeography) DeletePoint(ctx ciosctx.RequestCtx, s string) (cios.Point, *_nethttp.Response, error) {
	panic("implement me")
}
