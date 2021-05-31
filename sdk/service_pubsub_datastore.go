package ciossdk

import (
	"errors"
	"log"
	_nethttp "net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	cios_util "github.com/optim-corp/cios-golang-sdk/util/cios"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
	"github.com/optim-corp/cios-golang-sdk/util"

	"github.com/fcfcqloow/go-advance/check"
	cnv "github.com/fcfcqloow/go-advance/convert"
	xmath "github.com/fcfcqloow/go-advance/math"

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
func (self *CiosPubSub) GetDataStoreChannels(ctx ciosctx.RequestCtx, params cios.ApiGetDataStoreChannelsRequest) (response cios.MultipleDataStoreChannel, httpResponse *_nethttp.Response, err error) {
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
func (self *CiosPubSub) GetDataStoreChannel(ctx ciosctx.RequestCtx, channelID string) (cios.DataStoreChannel, *_nethttp.Response, error) {
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
func (self *CiosPubSub) GetObjects(ctx ciosctx.RequestCtx, channelID string, params cios.ApiGetDataStoreObjectsRequest) (response cios.MultipleDataStoreObject, httpResponse *_nethttp.Response, err error) {
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
func (self *CiosPubSub) GetObjectsAll(ctx ciosctx.RequestCtx, channelID string, params cios.ApiGetDataStoreObjectsRequest) ([]cios.DataStoreObject, *_nethttp.Response, error) {
	var (
		result      []cios.DataStoreObject
		httpRes     *_nethttp.Response
		err         error
		offset      = int64(0)
		_limit      = int64(1000)
		getFunction = func(offset int64) (cios.MultipleDataStoreObject, *_nethttp.Response, error) {
			return self.GetObjects(ctx, channelID, params.Limit(xmath.MinInt64(_limit, 1000)).Offset(offset+cnv.MustInt64(params.P_offset)))
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
		for offset = int64(1000); offset+cnv.MustInt64(params.P_offset) < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Objects...)
		}
	}
	return result, httpRes, err
}
func (self *CiosPubSub) GetObjectsUnlimited(ctx ciosctx.RequestCtx, channelID string, params cios.ApiGetDataStoreObjectsRequest) ([]cios.DataStoreObject, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetObjectsAll(ctx, channelID, params)
}
func (self *CiosPubSub) GetObject(ctx ciosctx.RequestCtx, channelID string, objectID string, packerFormat *string) (interface{}, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return map[string]interface{}{}, nil, err
	}
	request := self.ApiClient.PublishSubscribeApi.GetDataStoreObjectData(self.withHost(ctx), channelID, objectID)
	if packerFormat != nil {
		request = request.PackerFormat(*packerFormat)
	}
	return request.Execute()
}
func (self *CiosPubSub) GetObjectLatest(ctx ciosctx.RequestCtx, channelID string, packerFormat *string) (interface{}, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return map[string]interface{}{}, nil, err
	}
	request := self.ApiClient.PublishSubscribeApi.GetDataStoreObjectDataLatest(self.withHost(ctx), channelID)
	if packerFormat != nil {
		request = request.PackerFormat(*packerFormat)
	}
	return request.Execute()

}
func (self *CiosPubSub) MapObjectLatest(ctx ciosctx.RequestCtx, channelID string, packerFormat *string, stc interface{}) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	response, httpResponse, err := self.GetObjectLatest(ctx, channelID, packerFormat)
	if err != nil {
		return httpResponse, err
	}
	return httpResponse, cnv.DeepCopy(response, stc)
}
func (self *CiosPubSub) GetMultiObjectLatest(ctx ciosctx.RequestCtx, channelIDs []string) (cios.MultipleDataStoreDataLatest, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.MultipleDataStoreDataLatest{}, nil, err
	}
	request := self.ApiClient.PublishSubscribeApi.GetDataStoreMultiObjectDataLatest(self.withHost(ctx)).Ids(cios.Ids{Ids: &channelIDs})
	return request.Execute()
}
func (self *CiosPubSub) GetMultiObjectLatestByChannels(ctx ciosctx.RequestCtx, channels []cios.Channel) (cios.MultipleDataStoreDataLatest, *_nethttp.Response, error) {
	var channelIDs []string
	for _, channel := range channels {
		channelIDs = append(channelIDs, channel.Id)
	}
	return self.GetMultiObjectLatest(ctx, channelIDs)
}
func (self *CiosPubSub) MapMultiObjectLatestPayload(ctx ciosctx.RequestCtx, channelIDs []string, stc interface{}) ([]cios.PackerFormatJsonHeader, *_nethttp.Response, error) {
	var headers []cios.PackerFormatJsonHeader
	objects, httpResponse, err := self.GetMultiObjectLatest(ctx, channelIDs)
	if err == nil {
		headers, err = cios_util.MapPayloads(*objects.Objects, stc)
	}
	return headers, httpResponse, err
}
func (self *CiosPubSub) MapMultiObjectLatestPayloadByChannels(ctx ciosctx.RequestCtx, channels []cios.Channel, stc interface{}) ([]cios.PackerFormatJsonHeader, *_nethttp.Response, error) {
	var channelIDs []string
	for _, channel := range channels {
		channelIDs = append(channelIDs, channel.Id)
	}
	return self.MapMultiObjectLatestPayload(ctx, channelIDs, stc)
}
func (self *CiosPubSub) subscribeCiosWebSocket(ctx ciosctx.RequestCtx, _url string, beforeFunc *func(*websocket.Conn), logic func(body []byte) (bool, error), wsR, wsW int64) error {
	if ctx != nil {
		if _token, ok := ctx.Value(cios.ContextAccessToken).(string); ok {
			self.token = &_token
		}
	}
	if cnv.MustStr(self.token) == "" {
		if err := self.refresh(); err != nil {
			return err
		}
	}
	connection, err := self.CreateCIOSWebsocketConnection(_url, ParseAccessToken(cnv.MustStr(self.token)))
	if err != nil {
		return err
	}
	defer connection.Close()
	if beforeFunc != nil {
		(*beforeFunc)(connection)
	}
	for {
		if wsR != 0 {
			if err := connection.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(wsR))); err != nil {
				return err
			}
		}
		messageType, body, err := connection.ReadMessage()
		switch {
		case websocket.IsCloseError(err, websocket.CloseNormalClosure):
			return nil
		case websocket.IsUnexpectedCloseError(err):
			if _err := self.refresh(); _err != nil {
				return _err
			}
			connection.Close()
			if connection, err = self.CreateCIOSWebsocketConnection(_url, ParseAccessToken(cnv.MustStr(self.token))); err != nil {
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

func (self *CiosPubSub) GetStream(ctx ciosctx.RequestCtx, channelID string, params sdkmodel.ApiGetStreamRequest) ([]string, error) {
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
		limitStr = cnv.StrPtr(cnv.MustStr(params.LimitParam))
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
	err := self.subscribeCiosWebSocket(ctx, _url.String(), &bf,
		func(body []byte) (bool, error) {
			result = append(result, string(body))
			return false, nil
		}, self.wsReadTimeout, self.wsWriteTimeout,
	)
	if self.debug {
		log.Println(result)
	}
	return result, err
}
func (self *CiosPubSub) GetStreamAll(ctx ciosctx.RequestCtx, channelID string, params sdkmodel.ApiGetStreamRequest) ([]string, error) {
	var (
		result      []string
		err         error
		offset      = int64(0)
		_limit      = int64(1000)
		getFunction = func(offset int64) ([]string, error) {
			return self.GetStream(ctx, channelID, params.Limit(xmath.MinInt64(_limit, 1000)).Offset(offset+cnv.MustInt64(params.OffsetParam)))
		}
	)
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
		for {
			res, err := getFunction(offset)
			if err != nil {
				break
			}

			result = append(result, res...)
			if len(res) < 1000 {
				break
			}
			offset += 1000
		}
	}
	return result, err
}
func (self *CiosPubSub) GetStreamUnlimited(ctx ciosctx.RequestCtx, channelID string, params sdkmodel.ApiGetStreamRequest) ([]string, error) {
	if params.TimestampRangeParam == nil {
		timestampRange := ":" + cnv.MustStr(time.Now().UnixNano())
		params.TimestampRangeParam = &timestampRange
	}
	params.LimitParam = nil
	return self.GetStreamAll(ctx, channelID, params)
}
func (self *CiosPubSub) MapStreamAll(ctx ciosctx.RequestCtx, channelID string, params sdkmodel.ApiGetStreamRequest, stc interface{}) error {
	data, err := self.GetStreamAll(ctx, channelID, params)
	if err != nil {
		return err
	}
	return cios_util.DataStoreStreamToStruct(data, stc)
}
func (self *CiosPubSub) MapStreamUnlimited(ctx ciosctx.RequestCtx, channelID string, params sdkmodel.ApiGetStreamRequest, stc interface{}) error {
	data, err := self.GetStreamUnlimited(ctx, channelID, params)
	if err != nil {
		return err
	}
	return cios_util.DataStoreStreamToStruct(data, stc)
}
func (self *CiosPubSub) GetJsonStreamUnlimited(ctx ciosctx.RequestCtx, channelID string, params sdkmodel.ApiGetStreamRequest) (result []cios.PackerFormatJson, err error) {
	params.PackerFormatParam = cnv.StrPtr("json")
	data, _err := self.GetStreamUnlimited(ctx, channelID, params)
	if _err != nil {
		err = _err
	}
	err = cios_util.DataStoreStreamToStruct(data, &result)
	return
}
func (self *CiosPubSub) GetStreamFirst(ctx ciosctx.RequestCtx, channelID string, params sdkmodel.ApiGetStreamRequest) (string, error) {
	value, err := self.GetStreamAll(ctx, channelID, params.Limit(1))
	if err != nil {
		if value, err = self.GetStreamAll(ctx, channelID, params.Limit(1)); err != nil {
			return "", err
		}
	}
	if len(value) == 0 {
		return "", nil
	}
	return value[0], err
}
func (self *CiosPubSub) MapStreamFirst(ctx ciosctx.RequestCtx, channelID string, params sdkmodel.ApiGetStreamRequest, stc interface{}) error {
	value, err := self.GetStreamFirst(ctx, channelID, params)
	if err != nil {
		return err
	}
	return cnv.UnMarshalJson([]byte(value), stc)
}
func (self *CiosPubSub) DeleteDataByChannel(ctx ciosctx.RequestCtx, channelID string) (*_nethttp.Response, error) {
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
func (self *CiosPubSub) DeleteObject(ctx ciosctx.RequestCtx, channelID string, objectID string) (*_nethttp.Response, error) {
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
