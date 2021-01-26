package ciossdk

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/cios"
	"github.com/optim-corp/cios-golang-sdk/model"
	"github.com/optim-kazuhiro-seida/go-advance-type/check"
)

func MakeGetPointsOpts() cios.ApiGetPointsRequest {
	return cios.ApiGetPointsRequest{}
}

func (self Geography) GetPoints(params cios.ApiGetPointsRequest, ctx model.RequestCtx) (response cios.MultiplePoint, httpResponse *_nethttp.Response, err error) {
	params.ApiService = self.ApiClient.GeographyApi
	params.Ctx = ctx
	response, httpResponse, err = params.Execute()
	if err != nil && !check.IsNil(self.refresh) {
		if _, _, _, _, err = (*self.refresh)(); err != nil {
			return
		}
		response, httpResponse, err = params.Execute()
	}
	return
}
func (self Geography) CreatePoint(body cios.PointRequest, ctx model.RequestCtx) (cios.Point, *_nethttp.Response, error) {
	request := self.ApiClient.GeographyApi.CreatePoint(ctx).PointRequest(body)
	response, httpResponse, err := request.Execute()
	if err != nil {
		return cios.Point{}, httpResponse, err
	}
	return response.Point, httpResponse, err
}

func (self Geography) DeletePoint(pointID string, ctx model.RequestCtx) (cios.Point, *_nethttp.Response, error) {
	request := self.ApiClient.GeographyApi.DeletePoint(ctx, pointID)
	response, httpResponse, err := request.Execute()
	if err != nil {
		return cios.Point{}, httpResponse, err
	}
	return response.Point, httpResponse, err
}
