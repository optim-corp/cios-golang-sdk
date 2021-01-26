package ciossdk

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/cios"
	"github.com/optim-corp/cios-golang-sdk/model"
	"github.com/optim-kazuhiro-seida/go-advance-type/check"
)

func (self Account) InviteGroup(groupID string, email string, ctx model.RequestCtx) (response cios.Member, httpResponse *_nethttp.Response, err error) {
	request := self.ApiClient.GroupApi.InviteGroup(
		ctx,
		groupID,
	).GroupInviteRequest(cios.GroupInviteRequest{Email: &email})
	response, httpResponse, err = request.Execute()
	if err != nil && !check.IsNil(self.refresh) {
		if _, _, _, _, err = (*self.refresh)(); err != nil {
			return
		}
		response, httpResponse, err = request.Execute()
	}
	return
}
