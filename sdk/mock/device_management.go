package ciossdk_mock

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
)

type NoImplementDeviceManagement struct{}

func (NoImplementDeviceManagement) GetDevices(ctx ciosctx.RequestCtx, request cios.ApiGetDevicesRequest) (cios.MultipleDevice, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceManagement) GetDevicesAll(ctx ciosctx.RequestCtx, request cios.ApiGetDevicesRequest) ([]cios.Device, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceManagement) GetDevicesUnlimited(ctx ciosctx.RequestCtx, request cios.ApiGetDevicesRequest) ([]cios.Device, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceManagement) GetDevice(ctx ciosctx.RequestCtx, s string, s2 *string, b *bool) (cios.Device, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceManagement) GetDeviceInventory(ctx ciosctx.RequestCtx, s string) (map[string]interface{}, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceManagement) DeleteDevice(ctx ciosctx.RequestCtx, s string) (*_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceManagement) CreateDevice(ctx ciosctx.RequestCtx, info cios.DeviceInfo) (cios.Device, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceManagement) UpdateDevice(ctx ciosctx.RequestCtx, s string, request cios.DeviceUpdateRequest) (cios.Device, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceManagement) GetMonitoringLatestList(ctx ciosctx.RequestCtx, strings []string) ([]cios.DeviceMonitoring, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceManagement) GetMonitoring(ctx ciosctx.RequestCtx, s string) (cios.DeviceMonitoring, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceManagement) GetPolicies(ctx ciosctx.RequestCtx, request cios.ApiGetDevicePoliciesRequest) (cios.MultipleDevicePolicy, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceManagement) GetPoliciesAll(ctx ciosctx.RequestCtx, request cios.ApiGetDevicePoliciesRequest) ([]cios.DevicePolicy, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceManagement) GetPoliciesUnlimited(ctx ciosctx.RequestCtx, request cios.ApiGetDevicePoliciesRequest) ([]cios.DevicePolicy, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceManagement) DeletePolicy(ctx ciosctx.RequestCtx, s string) (*_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceManagement) CreatePolicy(ctx ciosctx.RequestCtx, s string) (cios.DevicePolicy, *_nethttp.Response, error) {
	panic("implement me")
}
