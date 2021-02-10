package ciossdk

import (
	"errors"
	"log"
	_nethttp "net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
	"github.com/optim-corp/cios-golang-sdk/util"

	"github.com/optim-kazuhiro-seida/go-advance-type/check"
	xmath "github.com/optim-kazuhiro-seida/go-advance-type/math"

	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	"github.com/gorilla/websocket"
)

func MakeGetDataStoreChannelsOpts() cios.ApiGetDataStoreChannelsRequest {
	return cios.ApiGetDataStoreChannelsRequest{}
}

func MakeGetObjectsOpts() cios.ApiGetDataStoreObjectsRequest {
	return cios.ApiGetDataStoreObjectsRequest{}
}
func MakeGetStreamOpts() sdkmodel.ApiGetStreamRequest {
	return sdkmodel.ApiGetStreamRequest{}
}
func (self *PubSub) GetDataStoreChannels(params cios.ApiGetDataStoreChannelsRequest, ctx sdkmodel.RequestCtx) (response cios.MultipleDataStoreChannel, httpResponse *_nethttp.Response, err error) {
	if err = self.refresh(); err != nil {
		return
	}
	params.Ctx = self.withHost(ctx)
	params.ApiService = self.ApiClient.PublishSubscribeApi
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_channelProtocolId = util.ToNil(params.P_channelProtocolId)
	return params.Execute()
}
func (self *PubSub) GetDataStoreChannel(channelID string, ctx sdkmodel.RequestCtx) (cios.DataStoreChannel, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.DataStoreChannel{}, nil, err
	}
	request := self.ApiClient.PublishSubscribeApi.GetDataStoreChannel(self.withHost(ctx), channelID)
	response, httpResponse, err := request.Execute()
	if err != nil {
		return cios.DataStoreChannel{}, httpResponse, err
	}

	return response.Channel, httpResponse, err
}
func (self *PubSub) GetObjects(channelID string, params cios.ApiGetDataStoreObjectsRequest, ctx sdkmodel.RequestCtx) (response cios.MultipleDataStoreObject, httpResponse *_nethttp.Response, err error) {
	if err = self.refresh(); err != nil {
		return
	}
	params.ApiService = self.ApiClient.PublishSubscribeApi
	params.Ctx = self.withHost(ctx)
	params.P_channelId = channelID
	params.P_label = util.ToNil(params.P_label)
	params.P_location = util.ToNil(params.P_location)
	params.P_locationRange = util.ToNil(params.P_locationRange)
	params.P_sessionId = util.ToNil(params.P_sessionId)
	params.P_timestamp = util.ToNil(params.P_timestamp)
	params.P_timestampRange = util.ToNil(params.P_timestampRange)
	params.P_channelProtocolId = util.ToNil(params.P_channelProtocolId)
	return params.Execute()
}
func (self *PubSub) GetObjectsAll(channelID string, params cios.ApiGetDataStoreObjectsRequest, ctx sdkmodel.RequestCtx) ([]cios.DataStoreObject, *_nethttp.Response, error) {
	var (
		result      []cios.DataStoreObject
		httpRes     *_nethttp.Response
		err         error
		offset      = int64(0)
		_limit      = int64(1000)
		getFunction = func(offset int64) (cios.MultipleDataStoreObject, *_nethttp.Response, error) {
			return self.GetObjects(channelID, params.Limit(xmath.MinInt64(_limit, 1000)).Offset(offset+convert.MustInt64(params.P_offset)), ctx)
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			res, httpRes, err := getFunction(offset)
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
		res, httpRes, err := getFunction(offset)
		if err != nil {
			return nil, httpRes, err
		}
		result = append(result, res.Objects...)
		for offset = int64(1000); offset+convert.MustInt64(params.P_offset) < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Objects...)
		}
	}
	return result, httpRes, err
}
func (self *PubSub) GetObjectsUnlimited(channelID string, params cios.ApiGetDataStoreObjectsRequest, ctx sdkmodel.RequestCtx) ([]cios.DataStoreObject, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetObjectsAll(channelID, params, ctx)
}
func (self *PubSub) GetObject(channelID string, objectID string, packerFormat *string, ctx sdkmodel.RequestCtx) (interface{}, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return map[string]interface{}{}, nil, err
	}
	request := self.ApiClient.PublishSubscribeApi.GetDataStoreObjectData(self.withHost(ctx), channelID, objectID)
	if packerFormat != nil {
		request = request.PackerFormat(*packerFormat)
	}
	return request.Execute()
}
func (self *PubSub) GetObjectLatest(channelID string, packerFormat *string, ctx sdkmodel.RequestCtx) (interface{}, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return map[string]interface{}{}, nil, err
	}
	request := self.ApiClient.PublishSubscribeApi.GetDataStoreObjectDataLatest(self.withHost(ctx), channelID)
	if packerFormat != nil {
		request = request.PackerFormat(*packerFormat)
	}
	return request.Execute()

}
func (self *PubSub) MapObjectLatest(channelID string, packerFormat *string, stc interface{}, ctx sdkmodel.RequestCtx) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	response, httpResponse, err := self.GetObjectLatest(channelID, packerFormat, ctx)
	if err != nil {
		return httpResponse, err
	}
	return httpResponse, convert.DeepCopy(response, stc)
}
func (self *PubSub) GetMultiObjectLatest(channelIDs []string, ctx sdkmodel.RequestCtx) (cios.MultipleDataStoreDataLatest, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.MultipleDataStoreDataLatest{}, nil, err
	}
	request := self.ApiClient.PublishSubscribeApi.GetDataStoreMultiObjectDataLatest(self.withHost(ctx)).Ids(cios.Ids{Ids: &channelIDs})
	return request.Execute()
}
func (self *PubSub) GetMultiObjectLatestByChannels(channels []cios.Channel, ctx sdkmodel.RequestCtx) (cios.MultipleDataStoreDataLatest, *_nethttp.Response, error) {
	var channelIDs []string
	for _, channel := range channels {
		channelIDs = append(channelIDs, channel.Id)
	}
	return self.GetMultiObjectLatest(channelIDs, ctx)
}
func (self *PubSub) MapMultiObjectLatestPayload(channelIDs []string, stc interface{}, ctx sdkmodel.RequestCtx) ([]cios.PackerFormatJsonHeader, *_nethttp.Response, error) {
	var headers []cios.PackerFormatJsonHeader
	objects, httpResponse, err := self.GetMultiObjectLatest(channelIDs, ctx)
	if err == nil {
		headers, err = util.MapPayloads(*objects.Objects, stc)
	}
	return headers, httpResponse, err
}
func (self *PubSub) MapMultiObjectLatestPayloadByChannels(channels []cios.Channel, stc interface{}, ctx sdkmodel.RequestCtx) ([]cios.PackerFormatJsonHeader, *_nethttp.Response, error) {
	var channelIDs []string
	for _, channel := range channels {
		channelIDs = append(channelIDs, channel.Id)
	}
	return self.MapMultiObjectLatestPayload(channelIDs, stc, ctx)
}
func (self *PubSub) subscribeCiosWebSocket(_url string, beforeFunc *func(*websocket.Conn), logic func(body []byte) (bool, error), ctx sdkmodel.RequestCtx) error {
	if ctx != nil {
		if _token, ok := ctx.Value(cios.ContextAccessToken).(string); ok {
			self.token = &_token
		}
	}
	if str(self.token) == "" {
		if err := self.refresh(); err != nil {
			return err
		}
	}
	connection, err := self.CreateCIOSWebsocketConnection(_url, ParseAccessToken(str(self.token)))
	if err != nil {
		return err
	}
	defer connection.Close()
	if beforeFunc != nil {
		(*beforeFunc)(connection)
	}
	for {
		messageType, body, err := connection.ReadMessage()
		switch {
		case websocket.IsCloseError(err, websocket.CloseNormalClosure):
			return nil
		case websocket.IsUnexpectedCloseError(err):
			if _err := self.refresh(); _err != nil {
				return _err
			}
			connection.Close()
			if connection, err = self.CreateCIOSWebsocketConnection(_url, ParseAccessToken(str(self.token))); err != nil {
				return err
			}
		case messageType == websocket.CloseMessage:
			return errors.New(string(body))
		case messageType == websocket.TextMessage:
			if done, err := logic(body); err != nil || done {
				return err
			}
		}
	}
}

func (self *PubSub) GetStream(channelID string, params sdkmodel.ApiGetStreamRequest, ctx sdkmodel.RequestCtx) ([]string, error) {
	var (
		result                    []string
		channelProtocolVersionStr *string
		ascendingStr              = "false"
		offsetStr                 = "0"
		limitStr                  *string
		_url, _                   = url.Parse(strings.Replace(self.Url, "https", "wss", 1) + "/v2/datastore/channels/" + channelID + "/object_stream")
	)
	if params.AscendingParam != nil {
		ascendingStr = strconv.FormatBool(*params.AscendingParam)
	}
	if params.OffsetParam != nil {
		offsetStr = strconv.FormatInt(*params.OffsetParam, 10)
	}
	if params.ChannelProtocolVersionParam != nil {
		tmp := strconv.FormatInt(int64(*params.ChannelProtocolVersionParam), 10)
		channelProtocolVersionStr = &tmp
	}
	if params.LimitParam != nil {
		limitStr = convert.StringPtr(convert.MustStr(params.LimitParam))
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
		"limit":                    limitStr,
	})
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
	err := self.subscribeCiosWebSocket(_url.String(), &bf,
		func(body []byte) (bool, error) {
			result = append(result, string(body))
			return false, nil
		}, ctx,
	)
	if self.debug {
		log.Println(result)
	}
	return result, err
}
func (self *PubSub) GetStreamAll(channelID string, params sdkmodel.ApiGetStreamRequest, ctx sdkmodel.RequestCtx) ([]string, error) {
	var (
		result      []string
		err         error
		offset      = int64(0)
		_limit      = int64(1000)
		getFunction = func(offset int64) ([]string, error) {
			return self.GetStream(channelID, params.Limit(xmath.MinInt64(_limit, 1000)).Offset(offset+convert.MustInt64(params.OffsetParam)), ctx)
		}
	)
	options := MakeGetObjectsOpts()
	options.P_limit = convert.Int64Ptr(1)
	options.P_label = params.LabelParam
	options.P_offset = params.OffsetParam
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
	total := response.Total

	if params.LimitParam != nil {
		_limit = *params.LimitParam
		for {
			res, err := getFunction(offset)
			if err != nil {
				return nil, err
			}
			result = append(result, res...)
			offset += 1000
			_limit -= 1000
			if _limit <= 0 {
				break
			}
		}
	} else {
		res, err := getFunction(offset)
		if err != nil {
			return nil, err
		}
		result = append(result, res...)
		for offset = int64(1000); offset+convert.MustInt64(params.OffsetParam) < total; offset += 1000 {
			res, err = getFunction(offset)
			if err != nil {
				return nil, err
			}
			result = append(result, res...)
		}
	}
	return result, err
}
func (self *PubSub) MapStreamAll(channelID string, params sdkmodel.ApiGetStreamRequest, stc interface{}, ctx sdkmodel.RequestCtx) error {
	data, err := self.GetStreamAll(channelID, params, ctx)
	if err != nil {
		return err
	}
	return util.DataStoreStreamToStruct(data, stc)
}
func (self *PubSub) GetStreamUnlimited(channelID string, params sdkmodel.ApiGetStreamRequest, ctx sdkmodel.RequestCtx) ([]string, error) {
	if params.TimestampRangeParam == nil {
		timestampRange := ":" + convert.MustStr(time.Now().UnixNano())
		params.TimestampRangeParam = &timestampRange
	}
	params.LimitParam = nil
	return self.GetStreamAll(channelID, params, ctx)
}
func (self *PubSub) MapStreamUnlimited(channelID string, params sdkmodel.ApiGetStreamRequest, stc interface{}, ctx sdkmodel.RequestCtx) error {
	data, err := self.GetStreamUnlimited(channelID, params, ctx)
	if err != nil {
		return err
	}
	return util.DataStoreStreamToStruct(data, stc)
}
func (self *PubSub) GetJsonStreamUnlimited(channelID string, params sdkmodel.ApiGetStreamRequest, ctx sdkmodel.RequestCtx) (result []cios.PackerFormatJson, err error) {
	params.PackerFormatParam = convert.StringPtr("json")
	data, _err := self.GetStreamUnlimited(channelID, params, ctx)
	if _err != nil {
		err = _err
	}
	err = util.DataStoreStreamToStruct(data, &result)
	return
}
func (self *PubSub) GetStreamFirst(channelID string, params sdkmodel.ApiGetStreamRequest, ctx sdkmodel.RequestCtx) (string, error) {
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
func (self *PubSub) MapStreamFirst(channelID string, params sdkmodel.ApiGetStreamRequest, stc interface{}, ctx sdkmodel.RequestCtx) error {
	value, err := self.GetStreamFirst(channelID, params, ctx)
	if err != nil {
		return err
	}
	return convert.UnMarshalJson([]byte(value), stc)
}
func (self *PubSub) DeleteDataByChannel(channelID string, ctx sdkmodel.RequestCtx) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	request := self.ApiClient.PublishSubscribeApi.DeleteDataStoreChannel(self.withHost(ctx), channelID)
	httpResponse, err := request.Execute()
	if err != nil && !check.IsNil(self.refresh) {
		httpResponse, err = request.Execute()
	}
	return httpResponse, err
}
func (self *PubSub) DeleteObject(channelID string, objectID string, ctx sdkmodel.RequestCtx) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	request := self.ApiClient.PublishSubscribeApi.DeleteDataStoreObjectData(self.withHost(ctx), channelID, objectID)
	_, hErr, err := request.Execute()
	if err != nil && !check.IsNil(self.refresh) {
		_, hErr, err = request.Execute()
	}
	return hErr, err
}
