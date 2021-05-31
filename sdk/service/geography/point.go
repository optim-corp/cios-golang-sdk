package srvgeography

import (
	_nethttp "net/http"

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
	return params.Execute()

}
func (self *CiosGeography) CreatePoint(ctx ciosctx.RequestCtx, body cios.PointRequest) (cios.Point, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Point{}, nil, err
	}
	request := self.ApiClient.GeographyApi.CreatePoint(self.withHost(ctx)).PointRequest(body)
	response, httpResponse, err := request.Execute()
	if err != nil {
		return cios.Point{}, httpResponse, err
	}
	return response.Point, httpResponse, err
}

func (self *CiosGeography) DeletePoint(ctx ciosctx.RequestCtx, pointID string) (cios.Point, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Point{}, nil, err
	}
	request := self.ApiClient.GeographyApi.DeletePoint(self.withHost(ctx), pointID)
	response, httpResponse, err := request.Execute()
	if err != nil {
		return cios.Point{}, httpResponse, err
	}
	return response.Point, httpResponse, err
}
