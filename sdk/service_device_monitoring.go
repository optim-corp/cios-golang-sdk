package ciossdk

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/cios"
	"github.com/optim-corp/cios-golang-sdk/model"
)

func (self DeviceManagement) GetMonitorings(deviceIDs []string, ctx model.RequestCtx) (cios.MultipleDeviceMonitoring, *_nethttp.Response, error) {
	request := self.ApiClient.DeviceApi.GetDeviceMonitoringsLatest(ctx).DeviceMonitoringIDsRequest(cios.DeviceMonitoringIDsRequest{DeviceIds: &deviceIDs})
	response, hErr, err := request.Execute()
	if err != nil && self.autoR {
		if _, _, _, _, err = self.refresh(); err != nil {
			return cios.MultipleDeviceMonitoring{}, hErr, err
		}
		response, hErr, err = request.Execute()
	}
	return response, hErr, err
}
func (self DeviceManagement) GetMonitoring(deviceID string, ctx model.RequestCtx) (cios.DeviceMonitoring, *_nethttp.Response, error) {
	request := self.ApiClient.DeviceApi.GetDeviceMonitoringLatest(ctx, deviceID)
	response, httpResponse, err := request.Execute()
	if err != nil {
		if self.autoR {
			if _, _, _, _, err = self.refresh(); err != nil {
				return cios.DeviceMonitoring{}, httpResponse, err
			}
			response, httpResponse, err = request.Execute()
		}
		if err != nil {
			return cios.DeviceMonitoring{}, httpResponse, err
		}
	}
	return response.Monitoring, httpResponse, err
}
