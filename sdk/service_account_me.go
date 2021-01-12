package ciossdk

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/cios"
	"github.com/optim-corp/cios-golang-sdk/model"
)

func (self Account) GetMe(ctx model.RequestCtx) (cios.Me, *_nethttp.Response, error) {
	request := self.ApiClient.AccountApi.GetMe(ctx)
	response, hErr, err := request.Execute()
	if err != nil && self.autoR {
		if _, _, _, _, err = self.refresh(); err != nil {
			return response, hErr, err
		}
		return request.Execute()
	}
	return response, hErr, err
}
