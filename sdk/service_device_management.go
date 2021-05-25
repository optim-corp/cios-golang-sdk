package ciossdk

import (
	_nethttp "net/http"

	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"

	cnv "github.com/fcfcqloow/go-advance/convert"

	"github.com/optim-corp/cios-golang-sdk/util"

	xmath "github.com/fcfcqloow/go-advance/math"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

func MakeGetDevicesOpts() cios.ApiGetDevicesRequest {
	return cios.ApiGetDevicesRequest{}
}

func (self DeviceManagement) GetDevices(ctx ciosctx.RequestCtx, params cios.ApiGetDevicesRequest) (response cios.MultipleDevice, httpResponse *_nethttp.Response, err error) {
	if err := self.refresh(); err != nil {
		return cios.MultipleDevice{}, nil, err
	}
	params.ApiService = self.ApiClient.DeviceApi
	params.Ctx = self.withHost(ctx)
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_name = util.ToNil(params.P_name)
	params.P_definitionId = util.ToNil(params.P_definitionId)
	params.P_definitionTag = util.ToNil(params.P_definitionTag)
	params.P_idNumber = util.ToNil(params.P_idNumber)
	params.P_resourceOwnerId = util.ToNil(params.P_resourceOwnerId)
	params.P_isPublic = util.ToNil(params.P_isPublic)
	params.P_lang = util.ToNil(params.P_lang)
	return params.Execute()
}
func (self DeviceManagement) GetDevicesAll(ctx ciosctx.RequestCtx, params cios.ApiGetDevicesRequest) ([]cios.Device, *_nethttp.Response, error) {
	var (
		result       []cios.Device
		httpResponse *_nethttp.Response
		err          error
		_limit       = int64(1000)
		offset       = int64(0)
		getFunction  = func(offset int64) (cios.MultipleDevice, *_nethttp.Response, error) {
			return self.GetDevices(ctx, params.Limit(xmath.MinInt64(_limit, 1000)).Offset(offset+cnv.MustInt64(params.P_offset)))
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			res, httpRes, err := getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Devices...)
			offset += 1000
			_limit -= 1000
			if _limit <= 0 {
				break
			}
		}
	} else {
		res, httpRes, err := getFunction(offset)
		if err != nil {
			return nil, httpRes, err
		}
		result = append(result, res.Devices...)
		for offset = int64(1000); offset+cnv.MustInt64(params.P_offset) < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Devices...)
		}
	}
	return result, httpResponse, err
}
func (self DeviceManagement) GetDevicesUnlimited(ctx ciosctx.RequestCtx, params cios.ApiGetDevicesRequest) ([]cios.Device, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetDevicesAll(ctx, params)
}
func (self DeviceManagement) GetDevice(ctx ciosctx.RequestCtx, deviceID string, lang *string, isDev *bool) (cios.Device, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Device{}, nil, err
	}
	request := self.ApiClient.DeviceApi.GetDevice(self.withHost(ctx), deviceID)
	if lang != nil {
		request = request.Lang(*lang)
	}
	if isDev != nil {
		request = request.IsDev(*isDev)
	}
	response, httpResponse, err := request.Execute()
	if err != nil {
		return cios.Device{}, httpResponse, err
	}
	return response.Device, httpResponse, err
}
func (self DeviceManagement) GetDeviceInventory(ctx ciosctx.RequestCtx, deviceID string) (map[string]interface{}, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, nil, err
	}
	return self.ApiClient.DeviceApi.GetDeviceInventoryLatest(ctx, deviceID).Execute()
}
func (self DeviceManagement) DeleteDevice(ctx ciosctx.RequestCtx, id string) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	return self.ApiClient.DeviceApi.DeleteDevice(self.withHost(ctx), id).Execute()
}
func (self DeviceManagement) CreateDevice(ctx ciosctx.RequestCtx, body cios.DeviceInfo) (cios.Device, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Device{}, nil, err
	}
	response, httpResponse, err := self.ApiClient.DeviceApi.CreateDevice(self.withHost(ctx)).DeviceInfo(body).Execute()
	if err != nil {
		return cios.Device{}, httpResponse, err
	}
	return response.Device, httpResponse, err
}
func (self DeviceManagement) UpdateDevice(ctx ciosctx.RequestCtx, deviceId string, body cios.DeviceUpdateRequest) (cios.Device, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Device{}, nil, err
	}
	response, httpResponse, err := self.ApiClient.DeviceApi.UpdateDevice(self.withHost(ctx), deviceId).DeviceUpdateRequest(body).Execute()
	if err != nil {
		return cios.Device{}, httpResponse, err
	}
	return response.Device, httpResponse, err
}
