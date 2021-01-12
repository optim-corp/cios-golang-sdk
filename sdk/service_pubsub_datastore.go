package ciossdk

import (
	"errors"
	"log"
	_nethttp "net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	xmath "github.com/optim-kazuhiro-seida/go-advance-type/math"

	"github.com/optim-corp/cios-golang-sdk/cios"

	"github.com/optim-corp/cios-golang-sdk/model"

	"github.com/optim-corp/cios-golang-sdk/util"

	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	"github.com/gorilla/websocket"
)

func MakeGetDataStoreChannelsOpts() cios.ApiGetDataStoreChannelsRequest {
	return cios.ApiGetDataStoreChannelsRequest{}
}

func MakeGetObjectsOpts() cios.ApiGetDataStoreObjectsRequest {
	return cios.ApiGetDataStoreObjectsRequest{}
}
func MakeGetStreamOpts() model.ApiGetStreamRequest {
	return model.ApiGetStreamRequest{}
}
func (self PubSub) GetDataStoreChannels(params cios.ApiGetDataStoreChannelsRequest, ctx model.RequestCtx) (response cios.MultipleDataStoreChannel, httpResponse *_nethttp.Response, err error) {
	params.Ctx = ctx
	params.ApiService = self.ApiClient.PublishSubscribeApi
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_channelProtocolId = util.ToNil(params.P_channelProtocolId)
	response, httpResponse, err = params.Execute()
	if err != nil && self.autoR {
		if _, _, _, _, err = self.refresh(); err != nil {
			return
		}
		response, httpResponse, err = params.Execute()
	}
	return
}
func (self PubSub) GetDataStoreChannel(channelID string, ctx model.RequestCtx) (cios.DataStoreChannel, *_nethttp.Response, error) {
	request := self.ApiClient.PublishSubscribeApi.GetDataStoreChannel(ctx, channelID)
	response, httpResponse, err := request.Execute()
	if err != nil {
		if self.autoR {
			if _, _, _, _, err = self.refresh(); err != nil {
				return cios.DataStoreChannel{}, nil, err
			}
			response, httpResponse, err = request.Execute()
		}
		if err != nil {
			return cios.DataStoreChannel{}, httpResponse, err
		}
	}
	return response.Channel, httpResponse, err
}
func (self PubSub) GetObjects(channelID string, params cios.ApiGetDataStoreObjectsRequest, ctx model.RequestCtx) (response cios.MultipleDataStoreObject, httpResponse *_nethttp.Response, err error) {
	params.ApiService = self.ApiClient.PublishSubscribeApi
	params.Ctx = ctx
	params.P_channelId = channelID
	params.P_label = util.ToNil(params.P_label)
	params.P_location = util.ToNil(params.P_location)
	params.P_locationRange = util.ToNil(params.P_locationRange)
	params.P_sessionId = util.ToNil(params.P_sessionId)
	params.P_timestamp = util.ToNil(params.P_timestamp)
	params.P_timestampRange = util.ToNil(params.P_timestampRange)
	params.P_channelProtocolId = util.ToNil(params.P_channelProtocolId)
	response, httpResponse, err = params.Execute()
	if err != nil && self.autoR {
		if _, _, _, _, err = self.refresh(); err != nil {
			return
		}
		response, httpResponse, err = params.Execute()
	}
	return
}
func (self PubSub) GetObjectsAll(channelID string, params cios.ApiGetDataStoreObjectsRequest, ctx model.RequestCtx) ([]cios.DataStoreObject, *_nethttp.Response, error) {
	var (
		result      []cios.DataStoreObject
		httpRes     *_nethttp.Response
		err         error
		offset      = int64(0)
		_limit      = int64(1000)
		getFunction = func(offset *int64) (cios.MultipleDataStoreObject, *_nethttp.Response, error) {
			if offset != nil {
				params.P_offset = offset
			}
			tlimit := xmath.MinInt64(_limit, 1000)
			params.P_limit = &tlimit
			return self.GetObjects(channelID, params, ctx)
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			res, httpRes, err := getFunction(&offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Objects...)
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
		result = append(result, res.Objects...)
		for offset = int64(1000); offset < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(&offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Objects...)
		}
	}
	return result, httpRes, err
}
func (self PubSub) GetObjectsUnlimited(channelID string, params cios.ApiGetDataStoreObjectsRequest, ctx model.RequestCtx) ([]cios.DataStoreObject, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetObjectsAll(channelID, params, ctx)
}
func (self PubSub) GetObject(channelID string, objectID string, packerFormat *string, ctx model.RequestCtx) (interface{}, *_nethttp.Response, error) {
	request := self.ApiClient.PublishSubscribeApi.GetDataStoreObjectData(ctx, channelID, objectID)
	if packerFormat != nil {
		request = request.PackerFormat(*packerFormat)
	}
	response, httpResponse, err := request.Execute()
	if err != nil && self.autoR {
		return request.Execute()
	}
	return response, httpResponse, err
}
func (self PubSub) GetObjectLatest(channelID string, packerFormat *string, ctx model.RequestCtx) (interface{}, *_nethttp.Response, error) {
	request := self.ApiClient.PublishSubscribeApi.GetDataStoreObjectDataLatest(ctx, channelID)
	if packerFormat != nil {
		request = request.PackerFormat(*packerFormat)
	}
	response, httpResponse, err := request.Execute()
	if err != nil && self.autoR {
		if _, _, _, _, err = self.refresh(); err != nil {
			return nil, nil, err
		}
		return request.Execute()
	}
	return response, httpResponse, err

}
func (self PubSub) MapObjectLatest(channelID string, packerFormat *string, stc interface{}, ctx model.RequestCtx) (*_nethttp.Response, error) {
	response, httpResponse, err := self.GetObjectLatest(channelID, packerFormat, ctx)
	if err != nil {
		return httpResponse, err
	}
	return httpResponse, convert.DeepCopy(response, stc)
}
func (self PubSub) GetMultiObjectLatest(channelIDs []string, ctx model.RequestCtx) (cios.MultipleDataStoreDataLatest, *_nethttp.Response, error) {
	request := self.ApiClient.PublishSubscribeApi.GetDataStoreMultiObjectDataLatest(ctx).Ids(cios.Ids{Ids: &channelIDs})
	response, httpResponse, err := request.Execute()
	if err != nil && self.autoR {
		if _, _, _, _, err = self.refresh(); err != nil {
			return cios.MultipleDataStoreDataLatest{}, nil, err
		}
		return request.Execute()
	}
	return response, httpResponse, err
}
func (self PubSub) GetMultiObjectLatestByChannels(channels []cios.Channel, ctx model.RequestCtx) (cios.MultipleDataStoreDataLatest, *_nethttp.Response, error) {
	var channelIDs []string
	for _, channel := range channels {
		channelIDs = append(channelIDs, channel.Id)
	}
	return self.GetMultiObjectLatest(channelIDs, ctx)
}
func (self PubSub) MapMultiObjectLatestPayload(channelIDs []string, stc interface{}, ctx model.RequestCtx) ([]cios.PackerFormatJsonHeader, *_nethttp.Response, error) {
	var headers []cios.PackerFormatJsonHeader
	objects, httpResponse, err := self.GetMultiObjectLatest(channelIDs, ctx)
	if err == nil {
		headers, err = util.MapPayloads(*objects.Objects, stc)
	}
	return headers, httpResponse, err
}
func (self PubSub) MapMultiObjectLatestPayloadByChannels(channels []cios.Channel, stc interface{}, ctx model.RequestCtx) ([]cios.PackerFormatJsonHeader, *_nethttp.Response, error) {
	var channelIDs []string
	for _, channel := range channels {
		channelIDs = append(channelIDs, channel.Id)
	}
	return self.MapMultiObjectLatestPayload(channelIDs, stc, ctx)
}

func (self PubSub) GetStream(channelID string, params model.ApiGetStreamRequest, ctx model.RequestCtx) ([]string, error) {
	var (
		result                    []string
		channelProtocolVersionStr *string
		ascendingStr              = "false"
		offsetStr                 = "0"
		limitStr                  = "0"
		count                     = int64(0)
		_url, _                   = url.Parse(strings.Replace(self.Url, "https", "wss", 1) + "/v2/datastore/channels/" + channelID + "/object_stream")
	)
	if params.AscendingParam != nil {
		ascendingStr = strconv.FormatBool(*params.AscendingParam)
	}
	if params.P_offset != nil {
		offsetStr = strconv.FormatInt(*params.P_offset, 10)
	}
	if params.ChannelProtocolVersionParam != nil {
		tmp := strconv.FormatInt(int64(*params.ChannelProtocolVersionParam), 10)
		channelProtocolVersionStr = &tmp
	}
	if params.P_limit != nil {
		limitStr = strconv.FormatInt(*params.P_limit, 10)
	}
	_url.RawQuery = util.Query(_url, util.QueryMap{
		"packer_format":            params.PackerFormatParam,
		"ascending":                &ascendingStr,
		"channel_protocol_id":      params.ChannelProtocolIdParam,
		"channel_protocol_version": channelProtocolVersionStr,
		"session_id":               params.SessionIdParam,
		"location":                 params.LocationParam,
		"location_range":           params.LocationRangeParam,
		"timestamp":                params.TimestampParam,
		"timestamp_range":          params.TimestampRangeParam,
		"label":                    params.LabelParam,
		"offset":                   &offsetStr,
		"limit":                    &limitStr,
	})
	if limitStr <= "0" {
		return []string{}, nil
	}

	bf := func(conn *websocket.Conn) {
		if err := conn.WriteMessage(websocket.TextMessage, []byte("load")); err != nil {
			panic(err)
		}
		if params.TimeoutParam != nil {
			if err := conn.SetReadDeadline(time.Now().Add(time.Duration(*params.TimeoutParam) * time.Second)); err != nil {
				panic(err)
			}
		}

	}
	err := self.SubscribeCiosWebSocket(_url.String(), &bf,
		func(body []byte) (bool, error) {
			result = append(result, string(body))
			count++
			if count >= *params.P_limit {
				return true, nil
			}
			return false, nil
		}, params.Context,
	)
	if self.debug {
		log.Println(result)
	}
	return result, err
}
func (self PubSub) GetStreamSafeLimit(channelID string, params model.ApiGetStreamRequest, ctx model.RequestCtx) ([]string, int64, error) {
	_limit := int64(1)
	timeout := 10
	options := MakeGetObjectsOpts()
	options.P_limit = &_limit
	options.P_label = params.LabelParam
	options.P_offset = params.P_offset
	options.P_location = params.LocationRangeParam
	options.P_timestamp = params.TimestampParam
	options.P_sessionId = params.SessionIdParam
	options.P_locationRange = params.LocationRangeParam
	options.P_timestampRange = params.TimestampRangeParam
	options.P_channelProtocolId = params.ChannelProtocolIdParam
	options.P_channelProtocolVersion = params.ChannelProtocolVersionParam
	response, _, err := self.GetObjects(channelID, options, ctx)
	if err != nil {
		return nil, 0, err
	}
	if params.P_limit == nil {
		return nil, 0, errors.New("nothing limit")
	}
	limit := xmath.MinInt64(*params.P_limit, response.Total)
	if params.P_offset != nil {
		limit = xmath.MinInt64(limit, response.Total-*params.P_offset)
	}
	params.P_limit = &limit
	params.TimeoutParam = &timeout
	val, err := self.GetStream(channelID, params, ctx)
	return val, response.Total, err
}
func (self PubSub) GetStreamAll(channelID string, params model.ApiGetStreamRequest, ctx model.RequestCtx) ([]string, error) {
	var (
		result  = []string{}
		total   = int64(0)
		err     error
		_limit  = int64(1000)
		timeout = 10
	)
	options := MakeGetObjectsOpts()
	options.P_limit = &_limit
	options.P_label = params.LabelParam
	options.P_offset = params.P_offset
	options.P_location = params.LocationRangeParam
	options.P_timestamp = params.TimestampParam
	options.P_sessionId = params.SessionIdParam
	options.P_locationRange = params.LocationRangeParam
	options.P_timestampRange = params.TimestampRangeParam
	options.P_channelProtocolId = params.ChannelProtocolIdParam
	options.P_channelProtocolVersion = params.ChannelProtocolVersionParam
	response, _, err := self.GetObjects(channelID, options, ctx)
	if err != nil {
		return nil, err
	}
	total = response.Total
	getFunction := func(offset *int64) ([]string, error) {
		limit := int64(1000)
		if params.P_limit != nil {
			limit = xmath.MinInt64(_limit, int64(1000))
		}
		limitArgs := xmath.MinInt64(limit, response.Total)
		if offset != nil {
			limitArgs = xmath.MinInt64(limit, response.Total-*offset)
		}
		params.P_offset = offset
		params.P_limit = &limitArgs
		params.TimeoutParam = &timeout
		return self.GetStream(channelID, params, ctx)
	}
	offset := int64(0)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			_result, err := getFunction(&offset)
			if err != nil {
				return result, err
			}
			result = append(result, _result...)
			offset += 1000
			_limit -= 1000
			if _limit <= 0 {
				break
			}
		}
	} else {
		result, err = getFunction(&offset)
		if err != nil {
			return result, err
		}
		if total > 1000 {
			for offset = int64(1000); offset < total; offset += 1000 {
				_result, err := getFunction(&offset)
				if err != nil {
					return result, err
				}
				result = append(result, _result...)
			}
		}
	}
	return result, err
}
func (self PubSub) MapStreamAll(channelID string, params model.ApiGetStreamRequest, stc interface{}, ctx model.RequestCtx) error {
	data, err := self.GetStreamAll(channelID, params, ctx)
	if err != nil {
		return err
	}
	return util.DataStoreStreamToStruct(data, stc)
}
func (self PubSub) GetStreamUnlimited(channelID string, params model.ApiGetStreamRequest, ctx model.RequestCtx) ([]string, error) {
	if params.TimestampRangeParam == nil {
		timestampRange := ":" + convert.MustStr(time.Now().UnixNano())
		params.TimestampRangeParam = &timestampRange
	}
	params.P_limit = nil
	return self.GetStreamAll(channelID, params, ctx)
}
func (self PubSub) MapStreamUnlimited(channelID string, params model.ApiGetStreamRequest, stc interface{}, ctx model.RequestCtx) error {
	data, err := self.GetStreamUnlimited(channelID, params, ctx)
	if err != nil {
		return err
	}
	return util.DataStoreStreamToStruct(data, stc)
}
func (self PubSub) GetJsonStreamUnlimited(channelID string, params model.ApiGetStreamRequest, ctx model.RequestCtx) (result []cios.PackerFormatJson, err error) {
	params.PackerFormatParam = convert.StringPtr("json")
	data, _err := self.GetStreamUnlimited(channelID, params, ctx)
	if _err != nil {
		err = _err
	}
	err = util.DataStoreStreamToStruct(data, &result)
	return
}
func (self PubSub) GetStreamFirst(channelID string, params model.ApiGetStreamRequest, ctx model.RequestCtx) (string, error) {
	value, err := self.GetStreamAll(channelID, params.Limit(1), ctx)
	if err != nil {
		if value, err = self.GetStreamAll(channelID, params.Limit(1), ctx); err != nil {
			return "", err
		}
	}
	if len(value) == 0 {
		return "", nil
	}
	return value[0], err
}
func (self PubSub) MapStreamFirst(channelID string, params model.ApiGetStreamRequest, stc interface{}, ctx model.RequestCtx) error {
	value, err := self.GetStreamFirst(channelID, params, ctx)
	if err != nil {
		return err
	}
	return convert.UnMarshalJson([]byte(value), stc)
}
func (self PubSub) DeleteDataByChannel(channelID string, ctx model.RequestCtx) (*_nethttp.Response, error) {
	request := self.ApiClient.PublishSubscribeApi.DeleteDataStoreChannel(ctx, channelID)
	httpResponse, err := request.Execute()
	if err != nil && self.autoR {
		httpResponse, err = request.Execute()
	}
	return httpResponse, err
}
func (self PubSub) DeleteObject(channelID string, objectID string, ctx model.RequestCtx) (*_nethttp.Response, error) {
	request := self.ApiClient.PublishSubscribeApi.DeleteDataStoreObjectData(ctx, channelID, objectID)
	_, hErr, err := request.Execute()
	if err != nil && self.autoR {
		_, hErr, err = request.Execute()
	}
	return hErr, err
}
