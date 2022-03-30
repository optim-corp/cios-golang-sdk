package srvgeography

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/util"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
)

func MakeGetPointsOpts() cios.ApiGetPointsRequest {
	return cios.ApiGetPointsRequest{}
}

func (self *CiosGeography) GetPoints(ctx ciosctx.RequestCtx, params cios.ApiGetPointsRequest) (response cios.MultiplePoint, httpResponse *_nethttp.Response, err error) {
	if err = self.refresh(); err != nil {
		return
	}
	params.ApiService = self.ApiClient.GeographyApi
	params.Ctx = self.withHost(ctx)
	params.P_name = util.ToNil(params.P_name)
	params.P_label = util.ToNil(params.P_label)
	params.P_isPublic = util.ToNil(params.P_isPublic)
	params.P_resourceOwnerId = util.ToNil(params.P_resourceOwnerId)
	params.P_sort = util.ToNil(params.P_sort)
	params.P_lang = util.ToNil(params.P_lang)
	return params.Execute()

}
func (self *CiosGeography) CreatePoint(ctx ciosctx.RequestCtx, body cios.PointRequest) (cios.SinglePoint, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.SinglePoint{}, nil, err
	}
	return self.ApiClient.GeographyApi.CreatePoint(self.withHost(ctx)).PointRequest(body).Execute()
}

func (self *CiosGeography) GetPoint(ctx ciosctx.RequestCtx, pointID string, lang *string, isDev *bool) (
	cios.SinglePoint, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.SinglePoint{}, nil, err
	}

	req := self.ApiClient.GeographyApi.GetPoint(self.withHost(ctx), pointID)

	if lang != nil {
		req = req.Lang(*lang)
	}

	if isDev != nil {
		req = req.IsDev(*isDev)
	}

	return req.Execute()
}

func (self *CiosGeography) UpdatePoint(ctx ciosctx.RequestCtx, pointID string, body []cios.DisplayInfo) (cios.SinglePoint, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.SinglePoint{}, nil, err
	}
	return self.ApiClient.GeographyApi.UpdatePoint(self.withHost(ctx), pointID).DisplayInfo(body).Execute()
}

func (self *CiosGeography) DeletePoint(ctx ciosctx.RequestCtx, pointID string) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}

	_, httpStatus, err := self.ApiClient.GeographyApi.DeletePoint(self.withHost(ctx), pointID).Execute()

	return httpStatus, err
}
