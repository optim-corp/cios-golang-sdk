package ciossdk

import (
	_nethttp "net/http"

	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"

	cnv "github.com/fcfcqloow/go-advance/convert"

	xmath "github.com/fcfcqloow/go-advance/math"

	"github.com/optim-corp/cios-golang-sdk/cios"
	"github.com/optim-corp/cios-golang-sdk/util"
)

func MakeGetPoliciesOpts() cios.ApiGetDevicePoliciesRequest {
	return cios.ApiGetDevicePoliciesRequest{}
}

func (self *CiosDeviceManagement) GetPolicies(ctx ciosctx.RequestCtx, params cios.ApiGetDevicePoliciesRequest) (response cios.MultipleDevicePolicy, httpResponse *_nethttp.Response, err error) {
	if err := self.refresh(); err != nil {
		return cios.MultipleDevicePolicy{}, nil, err
	}
	params.Ctx = self.withHost(ctx)
	params.ApiService = self.ApiClient.DeviceApi
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_resourceOwnerId = util.ToNil(params.P_resourceOwnerId)
	return params.Execute()
}

func (self *CiosDeviceManagement) GetPoliciesAll(ctx ciosctx.RequestCtx, params cios.ApiGetDevicePoliciesRequest) ([]cios.DevicePolicy, *_nethttp.Response, error) {
	var (
		result       []cios.DevicePolicy
		httpResponse *_nethttp.Response
		err          error
		_limit       = int64(1000)
		offset       = int64(0)
		getFunction  = func(offset int64) (cios.MultipleDevicePolicy, *_nethttp.Response, error) {
			return self.GetPolicies(ctx, params.Limit(xmath.MinInt64(_limit, 1000)).Offset(offset+cnv.MustInt64(params.P_offset)))
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
		for offset = int64(1000); offset+cnv.MustInt64(params.P_offset) < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Policies...)
		}
	}
	return result, httpResponse, err
}
func (self *CiosDeviceManagement) GetPoliciesUnlimited(ctx ciosctx.RequestCtx, params cios.ApiGetDevicePoliciesRequest) ([]cios.DevicePolicy, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetPoliciesAll(ctx, params)
}
func (self *CiosDeviceManagement) DeletePolicy(ctx ciosctx.RequestCtx, id string) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	return self.ApiClient.DeviceApi.DeletePolicy(self.withHost(ctx), id).Execute()
}
func (self *CiosDeviceManagement) CreatePolicy(ctx ciosctx.RequestCtx, resourceOwnerID string) (cios.DevicePolicy, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.DevicePolicy{}, nil, err
	}
	return self.ApiClient.DeviceApi.CreateDevicePolicy(self.withHost(ctx)).DevicePolicyRequest(cios.DevicePolicyRequest{ResourceOwnerId: resourceOwnerID}).Execute()
}
