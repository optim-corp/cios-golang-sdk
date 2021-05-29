package ciossdk_mock

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
)

type NoImplementAccount struct{}

func (NoImplementAccount) GetGroups(ctx ciosctx.RequestCtx, request cios.ApiGetGroupsRequest) (cios.MultipleGroup, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementAccount) GetGroupsAll(ctx ciosctx.RequestCtx, request cios.ApiGetGroupsRequest) ([]cios.Group, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementAccount) GetGroupsUnlimited(ctx ciosctx.RequestCtx, request cios.ApiGetGroupsRequest) ([]cios.Group, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementAccount) GetGroup(ctx ciosctx.RequestCtx, s string, s2 *string) (cios.Group, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementAccount) GetGroupByResourceOwnerId(ctx ciosctx.RequestCtx, s string, s2 *string) (cios.Group, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementAccount) GetGroupMapByResourceOwner(ctx ciosctx.RequestCtx, request cios.ApiGetGroupsRequest, request2 cios.ApiGetResourceOwnersRequest) (map[string]cios.Group, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementAccount) DeleteGroup(ctx ciosctx.RequestCtx, s string) (*_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementAccount) CreateGroup(ctx ciosctx.RequestCtx, request cios.GroupCreateRequest) (cios.Group, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementAccount) UpdateGroup(ctx ciosctx.RequestCtx, s string, request cios.GroupUpdateRequest) (cios.Group, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementAccount) GetMe(ctx ciosctx.RequestCtx) (cios.Me, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementAccount) InviteGroup(ctx ciosctx.RequestCtx, s string, s2 string) (cios.Member, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementAccount) GetResourceOwners(ctx ciosctx.RequestCtx, request cios.ApiGetResourceOwnersRequest) (cios.MultipleResourceOwner, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementAccount) GetResourceOwnersAll(ctx ciosctx.RequestCtx, request cios.ApiGetResourceOwnersRequest) ([]cios.ResourceOwner, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementAccount) GetResourceOwnersUnlimited(ctx ciosctx.RequestCtx, request cios.ApiGetResourceOwnersRequest) ([]cios.ResourceOwner, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementAccount) GetResourceOwner(ctx ciosctx.RequestCtx, s string) (cios.ResourceOwner, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementAccount) GetResourceOwnerByGroupId(ctx ciosctx.RequestCtx, s string) (cios.ResourceOwner, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementAccount) GetResourceOwnersMapByID(ctx ciosctx.RequestCtx) (map[string]cios.ResourceOwner, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementAccount) GetResourceOwnersMapByGroupID(ctx ciosctx.RequestCtx) (map[string]cios.ResourceOwner, error) {
	panic("implement me")
}
