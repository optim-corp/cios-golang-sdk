package srvdevice

import (
	_nethttp "net/http"

	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

func (self *CiosDeviceManagement) GetMonitoringLatestList(ctx ciosctx.RequestCtx, deviceIDs []string) ([]cios.DeviceMonitoring, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return []cios.DeviceMonitoring{}, nil, err
	}
	response, httpResponse, err := self.ApiClient.DeviceApi.GetDeviceMonitoringsLatest(self.withHost(ctx)).DeviceMonitoringIDsRequest(cios.DeviceMonitoringIDsRequest{DeviceIds: deviceIDs}).Execute()
	if err != nil {
		return []cios.DeviceMonitoring{}, httpResponse, err
	}
	return response.Monitorings, httpResponse, err
}
func (self *CiosDeviceManagement) GetMonitoring(ctx ciosctx.RequestCtx, deviceID string) (cios.DeviceMonitoring, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.DeviceMonitoring{}, nil, err
	}
	response, httpResponse, err := self.ApiClient.DeviceApi.GetDeviceMonitoringLatest(self.withHost(ctx), deviceID).Execute()
	if err != nil {
		return cios.DeviceMonitoring{}, httpResponse, err
	}
	return response.Monitoring, httpResponse, err
}
