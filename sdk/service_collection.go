package ciossdk

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

func (self *Collection) UploadData(collectionId, seriesId, resourceOwnerId string, timestamp int64, body cios.SeriesDataRequest, ctx sdkmodel.RequestCtx) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	return self.ApiClient.CollectionApi.PutSeries(self.withHost(ctx), collectionId, seriesId, timestamp).ResourceOwnerId(resourceOwnerId).SeriesDataRequest(body).Execute()
}

func (self *Collection) UploadDataBulk(collectionId, resourceOwnerId string, body cios.SeriesBulkRequest, ctx sdkmodel.RequestCtx) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	return self.ApiClient.CollectionApi.PostSeriesBulk(self.withHost(ctx), collectionId).ResourceOwnerId(resourceOwnerId).SeriesBulkRequest(body).Execute()
}

func (self *Collection) UploadImage(collectionId, seriesId, resourceOwnerId string, timestamp int64, image []byte, isLatest *bool, ctx sdkmodel.RequestCtx) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	req := self.ApiClient.CollectionApi.CreateSeriesImage(self.withHost(ctx), collectionId, seriesId, timestamp).ResourceOwnerId(resourceOwnerId).Body(string(image))
	req.P_isLatest = isLatest
	return req.Execute()
}

func (self *Collection) GetImage(collectionId, seriesId, resourceOwnerId string, timestamp int64, ctx sdkmodel.RequestCtx) ([]byte, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, nil, err
	}
	response, httpResponse, err := self.ApiClient.CollectionApi.GetSeriesImage(self.withHost(ctx), collectionId, seriesId, timestamp).ResourceOwnerId(resourceOwnerId).Execute()
	if err != nil {
		return nil, httpResponse, err
	}
	return []byte(response), httpResponse, err
}

func (self *Collection) GetData(params cios.ApiGetTimeSeriesDataRequest, collectionId, seriesId, resourceOwnerId string, ctx sdkmodel.RequestCtx) ([]byte, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, nil, err
	}
	params.Ctx = self.withHost(ctx)
	params.ApiService = self.ApiClient.CollectionApi
	params.P_collectionId = collectionId
	params.P_seriesId = seriesId
	return params.Execute()
}
