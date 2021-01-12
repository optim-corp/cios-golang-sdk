package ciossdk

import (
	_nethttp "net/http"
	"os"

	"github.com/optim-corp/cios-golang-sdk/util"

	xmath "github.com/optim-kazuhiro-seida/go-advance-type/math"
	"github.com/optim-corp/cios-golang-sdk/cios"
	"github.com/optim-corp/cios-golang-sdk/model"
)

func MakeGetResourceOwnersOpts() cios.ApiGetResourceOwnersRequest {
	return cios.ApiGetResourceOwnersRequest{}
}

func (self Account) GetResourceOwners(params cios.ApiGetResourceOwnersRequest, ctx model.RequestCtx) (response cios.MultipleResourceOwner, httpResponse *_nethttp.Response, err error) {
	params.ApiService = self.ApiClient.ResourceOwnerApi
	params.Ctx = ctx
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_page = util.ToNil(params.P_page)
	params.P_type_ = util.ToNil(params.P_type_)
	params.P_userId = util.ToNil(params.P_userId)
	response, httpResponse, err = params.Execute()
	if err != nil && self.autoR {
		if _, _, _, _, err = self.refresh(); err != nil {
			return
		}
		response, httpResponse, err = params.Execute()
	}
	return

}
func (self Account) GetResourceOwnersAll(params cios.ApiGetResourceOwnersRequest, ctx model.RequestCtx) ([]cios.ResourceOwner, *_nethttp.Response, error) {
	var (
		result      []cios.ResourceOwner
		httpRes     *_nethttp.Response
		err         error
		_limit      = int64(1000)
		offset      = int64(0)
		getFunction = func(offset *int64) (cios.MultipleResourceOwner, *_nethttp.Response, error) {
			limit := xmath.MinInt64(_limit, 1000)
			params.P_limit = &limit
			params.P_offset = offset
			return self.GetResourceOwners(params, ctx)
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			response, httpResponse, err := getFunction(&offset)
			if err != nil {
				return []cios.ResourceOwner{}, httpResponse, err
			}
			result = append(result, response.ResourceOwners...)
			offset += 1000
			_limit -= 1000
			if _limit <= 0 {
				break
			}
		}
	} else {
		response, httpResponse, err := getFunction(&offset)
		if err != nil {
			return []cios.ResourceOwner{}, httpResponse, err
		}
		result = append(result, response.ResourceOwners...)
		for offset = int64(1000); offset < response.Total; offset += 1000 {
			response, httpResponse, err := getFunction(&offset)
			if err != nil {
				return []cios.ResourceOwner{}, httpResponse, err
			}
			result = append(result, response.ResourceOwners...)
		}
	}
	return result, httpRes, err
}
func (self Account) GetResourceOwnersUnlimited(params cios.ApiGetResourceOwnersRequest, ctx model.RequestCtx) ([]cios.ResourceOwner, *_nethttp.Response, error) {
	params.P_offset = nil
	return self.GetResourceOwnersAll(params, ctx)
}

func (self Account) GetResourceOwner(id string, ctx model.RequestCtx) (cios.ResourceOwner, *_nethttp.Response, error) {
	request := self.ApiClient.ResourceOwnerApi.GetResourceOwner(ctx, id)
	response, httpResponse, err := request.Execute()
	if err != nil && self.autoR {
		if _, _, _, _, err = self.refresh(); err != nil {
			return response, httpResponse, err
		}
		return request.Execute()
	}
	return response, httpResponse, err

}
func (self Account) GetResourceOwnerByGroupId(groupID string, ctx model.RequestCtx) (cios.ResourceOwner, *_nethttp.Response, error) {
	group, httpResponse, err := self.GetGroup(groupID, nil, ctx)
	if err != nil {
		return cios.ResourceOwner{}, httpResponse, err
	}
	resourceOwners, httpResponse, err := self.GetResourceOwners(cios.ApiGetResourceOwnersRequest{P_groupId: &group.Id}, ctx)
	if err != nil {
		return cios.ResourceOwner{}, httpResponse, err
	}
	return resourceOwners.ResourceOwners[0], httpResponse, err
}

func (self Account) GetMyResourceOwnerId(ctx model.RequestCtx) string {
	me, _, _ := self.GetMe(ctx)
	value, _, _ := self.GetResourceOwners(MakeGetResourceOwnersOpts().UserId(me.Id), ctx)

	if value.Total != 0 {
		return value.ResourceOwners[0].Id
	} else {
		return os.Getenv("RESOURCE_OWNER_ID")
	}
}

func (self Account) GetResourceOwnersMapByID(ctx model.RequestCtx) (map[string]cios.ResourceOwner, *_nethttp.Response, error) {
	resourceOwnerMap := map[string]cios.ResourceOwner{}
	resourceOwners, httpResponse, err := self.GetResourceOwners(MakeGetResourceOwnersOpts(), ctx)
	if err != nil {
		return nil, httpResponse, err
	}
	for _, resourceOwner := range resourceOwners.ResourceOwners {
		resourceOwnerMap[resourceOwner.Id] = resourceOwner
	}
	return resourceOwnerMap, httpResponse, nil

}

func (self Account) GetResourceOwnersMapByGroupID(options model.RequestCtx) (map[string]cios.ResourceOwner, error) {
	resourceOwnerMap := map[string]cios.ResourceOwner{}
	resourceOwners, _, err := self.GetResourceOwners(MakeGetResourceOwnersOpts(), options)
	if err != nil {
		return nil, err
	}
	for _, resourceOwner := range resourceOwners.ResourceOwners {
		if resourceOwner.GroupId != nil {
			resourceOwnerMap[*resourceOwner.GroupId] = resourceOwner
		}
	}
	return resourceOwnerMap, nil

}

func ChannelsMapByResourceOwnerID(channels []cios.Channel) map[string][]cios.Channel {
	result := map[string][]cios.Channel{}
	for _, channel := range channels {
		result[channel.ResourceOwnerId] = append(result[channel.ResourceOwnerId], channel)
	}
	return result
}

func GroupMapByResourceOwnerID(groups []cios.Group, resourceOwners []cios.ResourceOwner) map[string]cios.Group {
	result := map[string]cios.Group{}
	for _, ro := range resourceOwners {
		for _, gp := range groups {
			if ro.GroupId != nil {
				if gp.Id == *ro.GroupId {
					result[ro.Id] = gp
				}
			}
		}
	}
	return result
}
func ResourceOwnerMapByGroupID(resourceOwners []cios.ResourceOwner, groups []cios.Group) map[string]cios.ResourceOwner {
	result := map[string]cios.ResourceOwner{}
	for _, ro := range resourceOwners {
		for _, gp := range groups {
			if ro.GroupId != nil {
				if gp.Id == *ro.GroupId {
					result[gp.Id] = ro
				}
			}
		}
	}
	return result
}
func ResourceOwnerIDMapByChannelID(channels []cios.Channel) map[string]string {
	result := map[string]string{}
	for _, c := range channels {
		result[c.Id] = c.ResourceOwnerId
	}
	return result
}

func BucketsMapByResourceOwnerID(buckets []cios.Bucket) map[string][]cios.Bucket {
	result := map[string][]cios.Bucket{}
	for _, bucket := range buckets {
		result[bucket.ResourceOwnerId] = append(result[bucket.ResourceOwnerId], bucket)
	}
	return result
}
