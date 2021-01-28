package ciossdk

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/cios"
	"github.com/optim-corp/cios-golang-sdk/model"
)

func (self DeviceManagement) GetMonitorings(deviceIDs []string, ctx model.RequestCtx) (cios.MultipleDeviceMonitoring, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.MultipleDeviceMonitoring{}, nil, err
	}
	return self.ApiClient.DeviceApi.GetDeviceMonitoringsLatest(ctx).DeviceMonitoringIDsRequest(cios.DeviceMonitoringIDsRequest{DeviceIds: &deviceIDs}).Execute()
}
func (self DeviceManagement) GetMonitoring(deviceID string, ctx model.RequestCtx) (cios.DeviceMonitoring, *_nethttp.Response, error) {
	response, httpResponse, err := self.ApiClient.DeviceApi.GetDeviceMonitoringLatest(ctx, deviceID).Execute()
	if err != nil {
		return cios.DeviceMonitoring{}, httpResponse, err
	}
	return response.Monitoring, httpResponse, err
}
