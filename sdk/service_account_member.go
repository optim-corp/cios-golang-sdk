package ciossdk

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

func (self Account) InviteGroup(groupID string, email string, ctx sdkmodel.RequestCtx) (response cios.Member, httpResponse *_nethttp.Response, err error) {
	if err := self.refresh(); err != nil {
		return cios.Member{}, nil, err
	}
	request := self.ApiClient.GroupApi.InviteGroup(
		ctx,
		groupID,
	).GroupInviteRequest(cios.GroupInviteRequest{Email: &email})
	return request.Execute()
}
