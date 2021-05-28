package ciossdk

import (
	_nethttp "net/http"

	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	"github.com/optim-corp/cios-golang-sdk/util"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

type CiosLicense _instance

func MakeGetLicensesOpts() cios.ApiGetMyLicensesRequest {
	return cios.ApiGetMyLicensesRequest{}
}

func (self *CiosLicense) GetLicenses(ctx ciosctx.RequestCtx, params cios.ApiGetMyLicensesRequest) (response []cios.License, httpResponse *_nethttp.Response, err error) {
	if err = self.refresh(); err != nil {
		return
	}
	params.ApiService = self.ApiClient.LicenseApi
	params.Ctx = self.withHost(ctx)
	params.P_status = util.ToNil(params.P_status)
	var temp cios.MultipleLicense
	temp, httpResponse, err = params.Execute()
	if err == nil {
		response = temp.Licenses
	}
	return
}
