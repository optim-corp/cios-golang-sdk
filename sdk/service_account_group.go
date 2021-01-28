package ciossdk

import (
	"errors"
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/util"

	xmath "github.com/optim-kazuhiro-seida/go-advance-type/math"

	"github.com/optim-corp/cios-golang-sdk/cios"
	"github.com/optim-corp/cios-golang-sdk/model"
)

func MakeGetGroupsOpts() cios.ApiGetGroupsRequest {
	return cios.ApiGetGroupsRequest{}
}

func (self Account) GetGroups(params cios.ApiGetGroupsRequest, ctx model.RequestCtx) (response cios.MultipleGroup, httpResponse *_nethttp.Response, err error) {
	if err := self.refresh(); err != nil {
		return cios.MultipleGroup{}, nil, err
	}
	params.Ctx = ctx
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
func (self Account) GetGroupsAll(params cios.ApiGetGroupsRequest, ctx model.RequestCtx) ([]cios.Group, *_nethttp.Response, error) {
	var (
		result      []cios.Group
		httpRes     *_nethttp.Response
		err         error
		offset      = int64(0)
		_limit      = int64(1000)
		getFunction = func(offset *int64) (cios.MultipleGroup, *_nethttp.Response, error) {
			if offset != nil {
				params.P_offset = offset
			}
			tlimit := xmath.MinInt64(_limit, 1000)
			params.P_limit = &tlimit
			return self.GetGroups(params, ctx)
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			res, httpRes, err := getFunction(&offset)
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
		res, httpRes, err := getFunction(&offset)
		if err != nil {
			return nil, httpRes, err
		}
		result = append(result, res.Groups...)
		for offset = int64(1000); offset < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(&offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Groups...)
		}
	}
	return result, httpRes, err
}
func (self Account) GetGroupsUnlimited(params cios.ApiGetGroupsRequest, ctx model.RequestCtx) ([]cios.Group, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetGroupsAll(params, ctx)
}

func (self Account) GetGroup(groupId string, includes *string, ctx model.RequestCtx) (cios.Group, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Group{}, nil, err
	}
	req := self.ApiClient.GroupApi.GetGroup(ctx, groupId)
	if includes != nil {
		req = req.Includes(*includes)
	}
	return req.Execute()
}
func (self Account) GetGroupByResourceOwnerId(resourceOwnerID string, includes *string, ctx model.RequestCtx) (cios.Group, *_nethttp.Response, error) {
	resourceOwner, httpResponse, err := self.GetResourceOwner(resourceOwnerID, ctx)
	if err != nil {
		return cios.Group{}, httpResponse, err
	}
	if resourceOwner.GroupId != nil {
		return self.GetGroup(*resourceOwner.GroupId, includes, ctx)
	}
	return cios.Group{}, nil, errors.New("No Group")
}

func (self Account) GetGroupMapByResourceOwner(p1 cios.ApiGetGroupsRequest, p2 cios.ApiGetResourceOwnersRequest, ctx model.RequestCtx) (map[string]cios.Group, *_nethttp.Response, error) {
	ros, httpResponse, err := self.GetResourceOwnersUnlimited(p2, ctx)
	if err != nil {
		return nil, httpResponse, err
	}
	gps, httpResponse, err := self.GetGroupsUnlimited(p1, ctx)
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

func (self Account) DeleteGroup(groupID string, ctx model.RequestCtx) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	return self.ApiClient.GroupApi.DeleteGroup(ctx, groupID).Execute()
}

func (self Account) CreateGroup(body cios.GroupCreateRequest, ctx model.RequestCtx) (cios.Group, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Group{}, nil, err
	}
	return self.ApiClient.GroupApi.CreateGroup(ctx).GroupCreateRequest(body).Execute()
}

func (self Account) UpdateGroup(groupID string, body cios.GroupUpdateRequest, ctx model.RequestCtx) (cios.Group, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Group{}, nil, err
	}
	return self.ApiClient.GroupApi.UpdateGroup(ctx, groupID).GroupUpdateRequest(body).Execute()
}
