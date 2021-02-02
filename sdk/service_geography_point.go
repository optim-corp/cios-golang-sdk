package ciossdk

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

func MakeGetPointsOpts() cios.ApiGetPointsRequest {
	return cios.ApiGetPointsRequest{}
}

func (self Geography) GetPoints(params cios.ApiGetPointsRequest, ctx sdkmodel.RequestCtx) (response cios.MultiplePoint, httpResponse *_nethttp.Response, err error) {
	if err = self.refresh(); err != nil {
		return
	}
	params.ApiService = self.ApiClient.GeographyApi
	params.Ctx = ctx
	return params.Execute()

}
func (self Geography) CreatePoint(body cios.PointRequest, ctx sdkmodel.RequestCtx) (cios.Point, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Point{}, nil, err
	}
	request := self.ApiClient.GeographyApi.CreatePoint(ctx).PointRequest(body)
	response, httpResponse, err := request.Execute()
	if err != nil {
		return cios.Point{}, httpResponse, err
	}
	return response.Point, httpResponse, err
}

func (self Geography) DeletePoint(pointID string, ctx sdkmodel.RequestCtx) (cios.Point, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Point{}, nil, err
	}
	request := self.ApiClient.GeographyApi.DeletePoint(ctx, pointID)
	response, httpResponse, err := request.Execute()
	if err != nil {
		return cios.Point{}, httpResponse, err
	}
	return response.Point, httpResponse, err
}
