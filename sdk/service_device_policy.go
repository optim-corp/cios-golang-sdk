package ciossdk

import (
	_nethttp "net/http"

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
	return self.ApiClient.DeviceApi.CreateDevicePolicy(ctx).DevicePolicyRequest(cios.DevicePolicyRequest{ResourceOwnerId: &resourceOwnerID}).Execute()
}
