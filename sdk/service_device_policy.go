package ciossdk

import (
	_nethttp "net/http"

	xmath "github.com/optim-kazuhiro-seida/go-advance-type/math"

	"github.com/optim-corp/cios-golang-sdk/cios"
	"github.com/optim-corp/cios-golang-sdk/util"

	"github.com/optim-corp/cios-golang-sdk/model"
)

func MakeGetPoliciesOpts() cios.ApiGetDevicePoliciesRequest {
	return cios.ApiGetDevicePoliciesRequest{}
}

func (self DeviceManagement) GetPolicies(params cios.ApiGetDevicePoliciesRequest, ctx model.RequestCtx) (response cios.MultipleDevicePolicy, httpResponse *_nethttp.Response, err error) {
	if err := self.refresh(); err != nil {
		return cios.MultipleDevicePolicy{}, nil, err
	}
	params.Ctx = ctx
	params.ApiService = self.ApiClient.DeviceApi
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_resourceOwnerId = util.ToNil(params.P_resourceOwnerId)
	return params.Execute()
}

func (self DeviceManagement) GetPoliciesAll(params cios.ApiGetDevicePoliciesRequest, ctx model.RequestCtx) ([]cios.DevicePolicy, *_nethttp.Response, error) {
	var (
		result       []cios.DevicePolicy
		httpResponse *_nethttp.Response
		err          error
		_limit       = int64(1000)
		offset       = int64(0)
		getFunction  = func(offset int64) (cios.MultipleDevicePolicy, *_nethttp.Response, error) {
			return self.GetPolicies(params.Limit(xmath.MinInt64(_limit, 1000)).Offset(offset), ctx)
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			res, httpRes, err := getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Policies...)
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
		result = append(result, res.Policies...)
		for offset = int64(1000); offset < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Policies...)
		}
	}
	return result, httpResponse, err
}
func (self DeviceManagement) GetPoliciesUnlimited(params cios.ApiGetDevicePoliciesRequest, ctx model.RequestCtx) ([]cios.DevicePolicy, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetPoliciesAll(params, ctx)
}
func (self DeviceManagement) DeletePolicy(id string, ctx model.RequestCtx) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	return self.ApiClient.DeviceApi.DeletePolicy(ctx, id).Execute()
}
func (self DeviceManagement) CreatePolicy(resourceOwnerID string, ctx model.RequestCtx) (cios.DevicePolicy, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.DevicePolicy{}, nil, err
	}
	return self.ApiClient.DeviceApi.CreateDevicePolicy(ctx).DevicePolicyRequest(cios.DevicePolicyRequest{ResourceOwnerId: resourceOwnerID}).Execute()
}
