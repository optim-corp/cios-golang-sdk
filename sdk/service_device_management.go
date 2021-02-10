package ciossdk

import (
	_nethttp "net/http"

	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	"github.com/optim-corp/cios-golang-sdk/util"

	xmath "github.com/optim-kazuhiro-seida/go-advance-type/math"

	"github.com/optim-corp/cios-golang-sdk/cios"

	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

func MakeGetDevicesOpts() cios.ApiGetDevicesRequest {
	return cios.ApiGetDevicesRequest{}
}

func (self DeviceManagement) GetDevices(params cios.ApiGetDevicesRequest, ctx sdkmodel.RequestCtx) (response cios.MultipleDevice, httpResponse *_nethttp.Response, err error) {
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
func (self DeviceManagement) GetDevicesAll(params cios.ApiGetDevicesRequest, ctx sdkmodel.RequestCtx) ([]cios.Device, *_nethttp.Response, error) {
	var (
		result       []cios.Device
		httpResponse *_nethttp.Response
		err          error
		_limit       = int64(1000)
		offset       = int64(0)
		getFunction  = func(offset int64) (cios.MultipleDevice, *_nethttp.Response, error) {
			return self.GetDevices(params.Limit(xmath.MinInt64(_limit, 1000)).Offset(offset+convert.MustInt64(params.P_offset)), ctx)
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
		for offset = int64(1000); offset+convert.MustInt64(params.P_offset) < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Devices...)
		}
	}
	return result, httpResponse, err
}
func (self DeviceManagement) GetDevicesUnlimited(params cios.ApiGetDevicesRequest, ctx sdkmodel.RequestCtx) ([]cios.Device, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetDevicesAll(params, ctx)
}
func (self DeviceManagement) GetDevice(deviceID string, lang *string, isDev *bool, ctx sdkmodel.RequestCtx) (cios.Device, *_nethttp.Response, error) {
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
func (self DeviceManagement) GetDeviceInventory(deviceID string, ctx sdkmodel.RequestCtx) (map[string]interface{}, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, nil, err
	}
	return self.ApiClient.DeviceApi.GetDeviceInventoryLatest(ctx, deviceID).Execute()
}
func (self DeviceManagement) DeleteDevice(id string, ctx sdkmodel.RequestCtx) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	return self.ApiClient.DeviceApi.DeleteDevice(self.withHost(ctx), id).Execute()
}
func (self DeviceManagement) CreateDevice(body cios.DeviceInfo, ctx sdkmodel.RequestCtx) (cios.Device, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Device{}, nil, err
	}
	response, httpResponse, err := self.ApiClient.DeviceApi.CreateDevice(self.withHost(ctx)).DeviceInfo(body).Execute()
	if err != nil {
		return cios.Device{}, httpResponse, err
	}
	return response.Device, httpResponse, err
}
func (self DeviceManagement) UpdateDevice(deviceId string, body cios.DeviceUpdateRequest, ctx sdkmodel.RequestCtx) (cios.Device, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Device{}, nil, err
	}
	response, httpResponse, err := self.ApiClient.DeviceApi.UpdateDevice(self.withHost(ctx), deviceId).DeviceUpdateRequest(body).Execute()
	if err != nil {
		return cios.Device{}, httpResponse, err
	}
	return response.Device, httpResponse, err
}
