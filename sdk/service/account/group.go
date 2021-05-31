package srvaccount

import (
	"errors"
	_nethttp "net/http"

	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"

	cnv "github.com/fcfcqloow/go-advance/convert"

	"github.com/optim-corp/cios-golang-sdk/util"

	xmath "github.com/fcfcqloow/go-advance/math"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

func MakeGetGroupsOpts() cios.ApiGetGroupsRequest {
	return cios.ApiGetGroupsRequest{}
}

func (self *CiosAccount) GetGroups(ctx ciosctx.RequestCtx, params cios.ApiGetGroupsRequest) (response cios.MultipleGroup, httpResponse *_nethttp.Response, err error) {
	if err := self.refresh(); err != nil {
		return cios.MultipleGroup{}, nil, err
	}
	params.Ctx = self.withHost(ctx)
	params.ApiService = self.ApiClient.GroupApi
	params.P_name = util.ToNil(params.P_name)
	params.P_label = util.ToNil(params.P_label)
	params.P_city = util.ToNil(params.P_city)
	params.P_domain = util.ToNil(params.P_domain)
	params.P_address1 = util.ToNil(params.P_address1)
	params.P_address2 = util.ToNil(params.P_address2)
	params.P_includes = util.ToNil(params.P_includes)
	params.P_parentGroupId = util.ToNil(params.P_parentGroupId)
	params.P_page = util.ToNil(params.P_page)
	params.P_state = util.ToNil(params.P_state)
	params.P_tags = util.ToNil(params.P_tags)
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_type_ = util.ToNil(params.P_type_)
	return params.Execute()
}
func (self *CiosAccount) GetGroupsAll(ctx ciosctx.RequestCtx, params cios.ApiGetGroupsRequest) ([]cios.Group, *_nethttp.Response, error) {
	var (
		result      []cios.Group
		httpRes     *_nethttp.Response
		err         error
		offset      = int64(0)
		_limit      = int64(1000)
		getFunction = func(offset int64) (cios.MultipleGroup, *_nethttp.Response, error) {
			return self.GetGroups(ctx, params.Limit(xmath.MinInt64(_limit, 1000)).Offset(offset+cnv.MustInt64(params.P_offset)))
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			res, httpRes, err := getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Groups...)
			offset += 1000
			_limit -= 1000
			if _limit <= 0 {
				break
			}
		}
	} else {
		res, httpRes, err := getFunction(offset)
		if err != nil {
			return nil, httpRes, err
		}
		result = append(result, res.Groups...)
		for offset = int64(1000); offset+cnv.MustInt64(params.P_offset) < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Groups...)
		}
	}
	return result, httpRes, err
}
func (self *CiosAccount) GetGroupsUnlimited(ctx ciosctx.RequestCtx, params cios.ApiGetGroupsRequest) ([]cios.Group, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetGroupsAll(ctx, params)
}

func (self *CiosAccount) GetGroup(ctx ciosctx.RequestCtx, groupId string, includes *string) (cios.Group, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Group{}, nil, err
	}
	req := self.ApiClient.GroupApi.GetGroup(self.withHost(ctx), groupId)
	if includes != nil {
		req = req.Includes(*includes)
	}
	return req.Execute()
}
func (self *CiosAccount) GetGroupByResourceOwnerId(ctx ciosctx.RequestCtx, resourceOwnerID string, includes *string) (cios.Group, *_nethttp.Response, error) {
	resourceOwner, httpResponse, err := self.GetResourceOwner(ctx, resourceOwnerID)
	if err != nil {
		return cios.Group{}, httpResponse, err
	}
	if resourceOwner.GroupId != nil {
		return self.GetGroup(ctx, *resourceOwner.GroupId, includes)
	}
	return cios.Group{}, nil, errors.New("No Group")
}

func (self *CiosAccount) GetGroupMapByResourceOwner(ctx ciosctx.RequestCtx, p1 cios.ApiGetGroupsRequest, p2 cios.ApiGetResourceOwnersRequest) (map[string]cios.Group, *_nethttp.Response, error) {
	ros, httpResponse, err := self.GetResourceOwnersUnlimited(ctx, p2)
	if err != nil {
		return nil, httpResponse, err
	}
	gps, httpResponse, err := self.GetGroupsUnlimited(ctx, p1)
	if err != nil {
		return nil, httpResponse, err
	}
	result := map[string]cios.Group{}
	for _, ro := range ros {
		for _, gp := range gps {
			if ro.GroupId != nil && gp.Id == *ro.GroupId {
				result[ro.Id] = gp
			}
		}
	}
	return result, nil, nil
}

func (self *CiosAccount) DeleteGroup(ctx ciosctx.RequestCtx, groupID string) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	return self.ApiClient.GroupApi.DeleteGroup(self.withHost(ctx), groupID).Execute()
}

func (self *CiosAccount) CreateGroup(ctx ciosctx.RequestCtx, body cios.GroupCreateRequest) (cios.Group, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Group{}, nil, err
	}
	return self.ApiClient.GroupApi.CreateGroup(self.withHost(ctx)).GroupCreateRequest(body).Execute()
}

func (self *CiosAccount) UpdateGroup(ctx ciosctx.RequestCtx, groupID string, body cios.GroupUpdateRequest) (cios.Group, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Group{}, nil, err
	}
	return self.ApiClient.GroupApi.UpdateGroup(self.withHost(ctx), groupID).GroupUpdateRequest(body).Execute()
}
