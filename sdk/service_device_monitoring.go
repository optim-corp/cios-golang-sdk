package ciossdk

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

func (self *DeviceManagement) GetMonitoringLatestList(deviceIDs []string, ctx sdkmodel.RequestCtx) ([]cios.DeviceMonitoring, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return []cios.DeviceMonitoring{}, nil, err
	}
	response, httpResponse, err := self.ApiClient.DeviceApi.GetDeviceMonitoringsLatest(self.withHost(ctx)).DeviceMonitoringIDsRequest(cios.DeviceMonitoringIDsRequest{DeviceIds: deviceIDs}).Execute()
	if err != nil {
		return []cios.DeviceMonitoring{}, httpResponse, err
	}
	return response.Monitorings, httpResponse, err
}
func (self *DeviceManagement) GetMonitoring(deviceID string, ctx sdkmodel.RequestCtx) (cios.DeviceMonitoring, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.DeviceMonitoring{}, nil, err
	}
	response, httpResponse, err := self.ApiClient.DeviceApi.GetDeviceMonitoringLatest(self.withHost(ctx), deviceID).Execute()
	if err != nil {
		return cios.DeviceMonitoring{}, httpResponse, err
	}
	return response.Monitoring, httpResponse, err
}
