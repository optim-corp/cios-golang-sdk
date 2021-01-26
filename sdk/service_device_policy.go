package ciossdk

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/util"
	"github.com/optim-kazuhiro-seida/go-advance-type/check"

	"github.com/optim-corp/cios-golang-sdk/cios"

	"github.com/optim-corp/cios-golang-sdk/model"
)

func MakeGetPoliciesOpts() cios.ApiGetDevicePoliciesRequest {
	return cios.ApiGetDevicePoliciesRequest{}
}

func (self DeviceManagement) GetPolicies(params cios.ApiGetDevicePoliciesRequest, ctx model.RequestCtx) (response cios.MultipleDevicePolicy, httpResponse *_nethttp.Response, err error) {
	params.Ctx = ctx
	params.ApiService = self.ApiClient.DeviceApi
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_resourceOwnerId = util.ToNil(params.P_resourceOwnerId)
	response, httpResponse, err = params.Execute()
	if err != nil && !check.IsNil(self.refresh) {
		if _, _, _, _, err = (*self.refresh)(); err != nil {
			return
		}
		response, httpResponse, err = params.Execute()
	}
	return
}

func (self DeviceManagement) DeletePolicy(id string, ctx model.RequestCtx) (*_nethttp.Response, error) {
	return self.ApiClient.DeviceApi.DeletePolicy(ctx, id).Execute()
}
func (self DeviceManagement) CreatePolicy(resourceOwnerID string, ctx model.RequestCtx) (cios.DevicePolicy, *_nethttp.Response, error) {
	return self.ApiClient.DeviceApi.CreateDevicePolicy(ctx).DevicePolicyRequest(cios.DevicePolicyRequest{ResourceOwnerId: &resourceOwnerID}).Execute()
}
