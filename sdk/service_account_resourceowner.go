package ciossdk

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/util/go_advance_type/convert"

	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
	"github.com/optim-corp/cios-golang-sdk/util"

	"github.com/optim-corp/cios-golang-sdk/cios"
	xmath "github.com/optim-corp/cios-golang-sdk/util/go_advance_type/math"
)

func MakeGetResourceOwnersOpts() cios.ApiGetResourceOwnersRequest {
	return cios.ApiGetResourceOwnersRequest{}
}

func (self *Account) GetResourceOwners(params cios.ApiGetResourceOwnersRequest, ctx sdkmodel.RequestCtx) (response cios.MultipleResourceOwner, httpResponse *_nethttp.Response, err error) {
	if err := self.refresh(); err != nil {
		return cios.MultipleResourceOwner{}, nil, err
	}
	params.ApiService = self.ApiClient.ResourceOwnerApi
	params.Ctx = self.withHost(ctx)
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_page = util.ToNil(params.P_page)
	params.P_type_ = util.ToNil(params.P_type_)
	params.P_userId = util.ToNil(params.P_userId)
	return params.Execute()
}
func (self *Account) GetResourceOwnersAll(params cios.ApiGetResourceOwnersRequest, ctx sdkmodel.RequestCtx) ([]cios.ResourceOwner, *_nethttp.Response, error) {
	var (
		result      []cios.ResourceOwner
		httpRes     *_nethttp.Response
		err         error
		_limit      = int64(1000)
		offset      = int64(0)
		getFunction = func(offset int64) (cios.MultipleResourceOwner, *_nethttp.Response, error) {
			return self.GetResourceOwners(params.Limit(xmath.MinInt64(_limit, 1000)).Offset(offset+convert.MustInt64(params.P_offset)), ctx)
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			response, httpResponse, err := getFunction(offset)
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
		response, httpResponse, err := getFunction(offset)
		if err != nil {
			return []cios.ResourceOwner{}, httpResponse, err
		}
		result = append(result, response.ResourceOwners...)
		for offset = int64(1000); offset+convert.MustInt64(params.P_offset) < response.Total; offset += 1000 {
			response, httpResponse, err := getFunction(offset)
			if err != nil {
				return []cios.ResourceOwner{}, httpResponse, err
			}
			result = append(result, response.ResourceOwners...)
		}
	}
	return result, httpRes, err
}
func (self *Account) GetResourceOwnersUnlimited(params cios.ApiGetResourceOwnersRequest, ctx sdkmodel.RequestCtx) ([]cios.ResourceOwner, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetResourceOwnersAll(params, ctx)
}

func (self *Account) GetResourceOwner(id string, ctx sdkmodel.RequestCtx) (cios.ResourceOwner, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.ResourceOwner{}, nil, err
	}
	return self.ApiClient.ResourceOwnerApi.GetResourceOwner(self.withHost(ctx), id).Execute()

}
func (self *Account) GetResourceOwnerByGroupId(groupID string, ctx sdkmodel.RequestCtx) (cios.ResourceOwner, *_nethttp.Response, error) {
	resourceOwners, httpResponse, err := self.GetResourceOwners(cios.ApiGetResourceOwnersRequest{P_groupId: &groupID}, ctx)
	if err != nil {
		return cios.ResourceOwner{}, httpResponse, err
	}
	return resourceOwners.ResourceOwners[0], httpResponse, err
}

func (self *Account) GetResourceOwnersMapByID(ctx sdkmodel.RequestCtx) (map[string]cios.ResourceOwner, *_nethttp.Response, error) {
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

func (self *Account) GetResourceOwnersMapByGroupID(options sdkmodel.RequestCtx) (map[string]cios.ResourceOwner, error) {
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
