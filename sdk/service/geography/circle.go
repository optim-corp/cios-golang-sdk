package srvgeography

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/util"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
)

func (self *CiosGeography) GetCircles(ctx ciosctx.RequestCtx, req cios.ApiGetCirclesRequest) (
	cios.MultipleCircle, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.MultipleCircle{}, nil, err
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

func (self *CiosGeography) CreateCircle(ctx ciosctx.RequestCtx, req cios.Circle) (
	cios.SingleCircle, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.SingleCircle{}, nil, err
	}

	return self.ApiClient.GeographyApi.CreateCircle(self.withHost(ctx)).Circle(req).Execute()
}

func (self *CiosGeography) GetCircle(ctx ciosctx.RequestCtx, pointID string, lang *string, isDev *bool) (
	cios.SingleCircle, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.SingleCircle{}, nil, err
	}

	req := self.ApiClient.GeographyApi.GetCircle(self.withHost(ctx), pointID)

	if lang != nil {
		req = req.Lang(*lang)
	}

	if isDev != nil {
		req = req.IsDev(*isDev)
	}

	return req.Execute()
}

func (self *CiosGeography) UpdateCircle(ctx ciosctx.RequestCtx, pointID string, displayInfo []cios.DisplayInfo) (
	cios.SingleCircle, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.SingleCircle{}, nil, err
	}

	return self.ApiClient.GeographyApi.UpdateCircle(self.withHost(ctx), pointID).DisplayInfo(displayInfo).Execute()
}

func (self *CiosGeography) DeleteCircle(ctx ciosctx.RequestCtx, pointID string) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}

	return self.ApiClient.GeographyApi.DeleteCircle(self.withHost(ctx), pointID).Execute()
}
