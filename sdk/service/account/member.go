package srvaccount

import (
	_nethttp "net/http"

	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

func (self *CiosAccount) InviteGroup(ctx ciosctx.RequestCtx, groupID string, email string) (response cios.Member, httpResponse *_nethttp.Response, err error) {
	if err := self.refresh(); err != nil {
		return cios.Member{}, nil, err
	}
	request := self.ApiClient.GroupApi.InviteGroup(
		self.withHost(ctx),
		groupID,
	).GroupInviteRequest(cios.GroupInviteRequest{Email: &email})
	return request.Execute()
}
