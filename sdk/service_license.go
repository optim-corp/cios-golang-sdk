package ciossdk

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/util"
	"github.com/optim-kazuhiro-seida/go-advance-type/check"

	"github.com/optim-corp/cios-golang-sdk/cios"
	"github.com/optim-corp/cios-golang-sdk/model"
)

func MakeGetLicensesOpts() cios.ApiGetMyLicensesRequest {
	return cios.ApiGetMyLicensesRequest{}
}

func (self License) GetLicenses(params cios.ApiGetMyLicensesRequest, ctx model.RequestCtx) (response []cios.License, httpResponse *_nethttp.Response, err error) {
	params.ApiService = self.ApiClient.LicenseApi
	params.Ctx = ctx
	params.P_status = util.ToNil(params.P_status)
	var temp cios.MultipleLicense
	temp, httpResponse, err = params.Execute()
	if err != nil && !check.IsNil(self.refresh) {
		if _, _, _, _, err = (*self.refresh)(); err != nil {
			return
		}
		temp, httpResponse, err = params.Execute()
	}
	if err == nil {
		response = temp.Licenses
	}
	return
}
