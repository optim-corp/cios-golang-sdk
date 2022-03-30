package ciossdmock

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

func (NoImplementGeography) GetPoint(ctx ciosctx.RequestCtx, pointID string, lang *string, isDev *bool) (
	cios.SinglePoint, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementGeography) UpdatePoint(ctx ciosctx.RequestCtx, request cios.ApiUpdatePointRequest) (cios.SinglePoint, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementGeography) DeletePoint(ctx ciosctx.RequestCtx, s string) (*_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementGeography) GetCircles(ctx ciosctx.RequestCtx, req cios.ApiGetCirclesRequest) (
	cios.MultipleCircle, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementGeography) CreateCircle(ctx ciosctx.RequestCtx, req cios.Circle) (
	cios.SingleCircle, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementGeography) GetCircle(ctx ciosctx.RequestCtx, pointID string, lang *string, isDev *bool) (
	cios.SingleCircle, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementGeography) UpdateCircle(ctx ciosctx.RequestCtx, pointID string, displayInfo []cios.DisplayInfo) (
	cios.SingleCircle, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementGeography) DeleteCircle(ctx ciosctx.RequestCtx, pointID string) (*_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementGeography) GetRoutes(ctx ciosctx.RequestCtx, req cios.ApiGetRoutesRequest) (
	cios.MultipleRoute, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementGeography) CreateRoute(ctx ciosctx.RequestCtx, req cios.Route) (
	cios.SingleRoute, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementGeography) GetRoute(ctx ciosctx.RequestCtx, routeID string, lang *string, isDev *bool) (
	cios.SingleRoute, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementGeography) UpdateRoute(ctx ciosctx.RequestCtx, routeID string, displayInfo []cios.DisplayInfo) (
	cios.SingleRoute, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementGeography) DeleteRoute(ctx ciosctx.RequestCtx, routeID string) (*_nethttp.Response, error) {
	panic("implement me")
}
