package ciossdk

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

func (self Account) GetMe(ctx sdkmodel.RequestCtx) (cios.Me, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Me{}, nil, err
	}
	return self.ApiClient.AccountApi.GetMe(ctx).Execute()
}
