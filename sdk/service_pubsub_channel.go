package ciossdk

import (
	"fmt"
	_nethttp "net/http"

	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
	"github.com/optim-corp/cios-golang-sdk/util"

	xmath "github.com/optim-kazuhiro-seida/go-advance-type/math"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

func MakeGetChannelsOpts() cios.ApiGetChannelsRequest {
	return cios.ApiGetChannelsRequest{}
}

func (self *PubSub) GetChannels(params cios.ApiGetChannelsRequest, ctx sdkmodel.RequestCtx) (response cios.MultipleChannel, httpResponse *_nethttp.Response, err error) {
	if err = self.refresh(); err != nil {
		return
	}
	params.Ctx = ctx
	params.ApiService = self.ApiClient.PublishSubscribeApi
	params.P_name = util.ToNil(params.P_name)
	params.P_label = util.ToNil(params.P_label)
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_resourceOwnerId = util.ToNil(params.P_resourceOwnerId)
	params.P_isPublic = util.ToNil(params.P_isPublic)
	params.P_lang = util.ToNil(params.P_lang)
	params.P_channelProtocol = util.ToNil(params.P_channelProtocol)
	params.P_datastoreEnabled = util.ToNil(params.P_datastoreEnabled)
	params.P_messagingEnabled = util.ToNil(params.P_messagingEnabled)
	params.P_messagingPersisted = util.ToNil(params.P_messagingPersisted)
	return params.Execute()
}
func (self *PubSub) GetChannelsAll(params cios.ApiGetChannelsRequest, ctx sdkmodel.RequestCtx) ([]cios.Channel, *_nethttp.Response, error) {
	var (
		result      []cios.Channel
		httpRes     *_nethttp.Response
		err         error
		offset      = int64(0)
		_limit      = int64(1000)
		getFunction = func(offset int64) (cios.MultipleChannel, *_nethttp.Response, error) {
			return self.GetChannels(params.Limit(xmath.MinInt64(_limit, 1000)).Offset(offset+convert.MustInt64(params.P_offset)), ctx)
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			res, httpRes, err := getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Channels...)
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
		result = append(result, res.Channels...)
		for offset = int64(1000); offset+convert.MustInt64(params.P_offset) < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Channels...)
		}
	}
	return result, httpRes, err
}
func (self *PubSub) GetChannelsUnlimited(params cios.ApiGetChannelsRequest, ctx sdkmodel.RequestCtx) ([]cios.Channel, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetChannelsAll(params, ctx)
}
func (self *PubSub) GetChannel(channelID string, isDev *bool, lang *string, ctx sdkmodel.RequestCtx) (cios.Channel, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Channel{}, nil, err
	}
	request := self.ApiClient.PublishSubscribeApi.GetChannel(ctx, channelID)
	if isDev != nil {
		request = request.IsDev(*isDev)
	}
	if lang != nil {
		request = request.Lang(*lang)
	}
	response, httpResponse, err := request.Execute()
	if err != nil {
		return cios.Channel{}, httpResponse, err
	}
	return response.Channel, httpResponse, err
}
func (self *PubSub) GetChannelFirst(params cios.ApiGetChannelsRequest, ctx sdkmodel.RequestCtx) (cios.Channel, *_nethttp.Response, error) {
	response, httpResponse, err := self.GetChannels(params.Limit(1), ctx)
	if err != nil || len(response.Channels) == 0 {
		return cios.Channel{}, httpResponse, fmt.Errorf("not found")
	}
	return response.Channels[0], httpResponse, err
}
func (self *PubSub) GetChannelsMapByID(params cios.ApiGetChannelsRequest, ctx sdkmodel.RequestCtx) (map[string]cios.Channel, *_nethttp.Response, error) {
	channels, httpResponse, err := self.GetChannelsUnlimited(params, ctx)
	if err != nil {
		return nil, httpResponse, err
	}
	channelsMap := map[string]cios.Channel{}
	for _, channel := range channels {
		channelsMap[channel.Id] = channel
	}
	return channelsMap, httpResponse, nil
}
func (self *PubSub) GetChannelsMapByResourceOwnerID(params cios.ApiGetChannelsRequest, ctx sdkmodel.RequestCtx) (map[string][]cios.Channel, *_nethttp.Response, error) {
	channels, httpResponse, err := self.GetChannelsUnlimited(params, ctx)
	if err != nil {
		return nil, httpResponse, err
	}
	channelsMap := map[string][]cios.Channel{}
	for _, channel := range channels {
		channelsMap[channel.ResourceOwnerId] = append(channelsMap[channel.Id], channel)
	}
	return channelsMap, httpResponse, nil
}
func (self *PubSub) DeleteChannel(channelID string, ctx sdkmodel.RequestCtx) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	request := self.ApiClient.PublishSubscribeApi.
		DeleteChannel(
			ctx,
			channelID,
		)
	httpResponse, err := request.Execute()
	return httpResponse, err
}
func (self *PubSub) GetOrCreateChannel(params cios.ApiGetChannelsRequest, body cios.ChannelProposal, ctx sdkmodel.RequestCtx) (cios.Channel, *_nethttp.Response, error) {
	channels, httpResponse, err := self.GetChannelsUnlimited(params, ctx)
	if len(channels) == 0 {
		return self.CreateChannel(body, ctx)
	}
	return channels[0], httpResponse, err
}
func (self *PubSub) CreateChannel(body cios.ChannelProposal, ctx sdkmodel.RequestCtx) (cios.Channel, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Channel{}, nil, err
	}
	request := self.ApiClient.PublishSubscribeApi.CreateChannel(ctx).ChannelProposal(body)
	response, httpResponse, err := request.Execute()
	if err != nil {
		return cios.Channel{}, httpResponse, err
	}
	return response.Channel, httpResponse, err
}
func (self *PubSub) UpdateChannel(channelID string, body cios.ChannelUpdateProposal, ctx sdkmodel.RequestCtx) (cios.MultipleChannel, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.MultipleChannel{}, nil, err
	}
	return self.ApiClient.PublishSubscribeApi.UpdateChannel(ctx, channelID).ChannelUpdateProposal(body).Execute()
}
func GetDefaultDisplayInfo(displayInfos []cios.DisplayInfo) *cios.DisplayInfo {
	for _, v := range displayInfos {
		if v.IsDefault == true {
			return &v
		}
	}
	return nil
}
