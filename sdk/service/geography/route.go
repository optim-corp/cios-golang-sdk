package srvgeography

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	"github.com/optim-corp/cios-golang-sdk/util"
)

func (self *CiosGeography) GetRoutes(ctx ciosctx.RequestCtx, req cios.ApiGetRoutesRequest) (
	cios.MultipleRoute, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.MultipleRoute{}, nil, err
	}
	req.ApiService = self.ApiClient.GeographyApi
	req.Ctx = self.withHost(ctx)
	req.P_name = util.ToNil(req.P_name)
	req.P_label = util.ToNil(req.P_label)
	req.P_isPublic = util.ToNil(req.P_isPublic)
	req.P_resourceOwnerId = util.ToNil(req.P_resourceOwnerId)
	req.P_sort = util.ToNil(req.P_sort)
	req.P_lang = util.ToNil(req.P_lang)
	return req.Execute()
}

func (self *CiosGeography) CreateRoute(ctx ciosctx.RequestCtx, req cios.Route) (
	cios.SingleRoute, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.SingleRoute{}, nil, err
	}

	return self.ApiClient.GeographyApi.CreateRoute(self.withHost(ctx)).Route(req).Execute()
}

func (self *CiosGeography) GetRoute(ctx ciosctx.RequestCtx, routeID string, lang *string, isDev *bool) (
	cios.SingleRoute, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.SingleRoute{}, nil, err
	}

	req := self.ApiClient.GeographyApi.GetRoute(self.withHost(ctx), routeID)

	if lang != nil {
		req = req.Lang(*lang)
	}

	if isDev != nil {
		req = req.IsDev(*isDev)
	}

	return req.Execute()
}

func (self *CiosGeography) UpdateRoute(ctx ciosctx.RequestCtx, routeID string, displayInfo []cios.DisplayInfo) (
	cios.SingleRoute, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.SingleRoute{}, nil, err
	}

	return self.ApiClient.GeographyApi.UpdateRoute(self.withHost(ctx), routeID).DisplayInfo(displayInfo).Execute()
}

func (self *CiosGeography) DeleteRoute(ctx ciosctx.RequestCtx, routeID string) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}

	return self.ApiClient.GeographyApi.DeleteRoute(self.withHost(ctx), routeID).Execute()
}
