package ciossdk

import (
	_nethttp "net/http"

	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

func (self *CiosAccount) GetMe(ctx ciosctx.RequestCtx) (cios.Me, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Me{}, nil, err
	}
	return self.ApiClient.AccountApi.GetMe(self.withHost(ctx)).Execute()
}
