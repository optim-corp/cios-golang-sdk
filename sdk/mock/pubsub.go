package ciossdmock

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/sdk/enum"

	srvpubsub "github.com/optim-corp/cios-golang-sdk/sdk/service/pubsub"

	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"

	"github.com/gorilla/websocket"
	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
)

type NoImplementPubsub struct{}

func (NoImplementPubsub) GetChannels(ctx ciosctx.RequestCtx, request cios.ApiGetChannelsRequest) (cios.MultipleChannel, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) GetChannelsAll(ctx ciosctx.RequestCtx, request cios.ApiGetChannelsRequest) ([]cios.Channel, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) GetChannelsUnlimited(ctx ciosctx.RequestCtx, request cios.ApiGetChannelsRequest) ([]cios.Channel, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) GetChannel(ctx ciosctx.RequestCtx, s string, b *bool, s2 *string) (cios.Channel, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) GetChannelFirst(ctx ciosctx.RequestCtx, request cios.ApiGetChannelsRequest) (cios.Channel, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) GetChannelsMapByID(ctx ciosctx.RequestCtx, request cios.ApiGetChannelsRequest) (map[string]cios.Channel, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) GetChannelsMapByResourceOwnerID(ctx ciosctx.RequestCtx, request cios.ApiGetChannelsRequest) (map[string][]cios.Channel, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) DeleteChannel(ctx ciosctx.RequestCtx, s string) (*_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) GetOrCreateChannel(ctx ciosctx.RequestCtx, request cios.ApiGetChannelsRequest, proposal cios.ChannelProposal) (cios.Channel, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) CreateChannel(ctx ciosctx.RequestCtx, proposal cios.ChannelProposal) (cios.Channel, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) UpdateChannel(ctx ciosctx.RequestCtx, s string, proposal cios.ChannelUpdateProposal) (cios.MultipleChannel, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) GetDataStoreChannels(ctx ciosctx.RequestCtx, request cios.ApiGetDataStoreChannelsRequest) (cios.MultipleDataStoreChannel, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) GetDataStoreChannel(ctx ciosctx.RequestCtx, s string) (cios.DataStoreChannel, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) GetObjects(ctx ciosctx.RequestCtx, s string, request cios.ApiGetDataStoreObjectsRequest) (cios.MultipleDataStoreObject, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) GetObjectsAll(ctx ciosctx.RequestCtx, s string, request cios.ApiGetDataStoreObjectsRequest) ([]cios.DataStoreObject, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) GetObjectsUnlimited(ctx ciosctx.RequestCtx, s string, request cios.ApiGetDataStoreObjectsRequest) ([]cios.DataStoreObject, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) GetObject(ctx ciosctx.RequestCtx, s string, s2 string, s3 *string) (interface{}, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) GetObjectLatest(ctx ciosctx.RequestCtx, s string, s2 *string) (interface{}, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) MapObjectLatest(ctx ciosctx.RequestCtx, s string, s2 *string, i interface{}) (*_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) GetMultiObjectLatest(ctx ciosctx.RequestCtx, strings []string) (cios.MultipleDataStoreDataLatest, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) GetMultiObjectLatestByChannels(ctx ciosctx.RequestCtx, channels []cios.Channel) (cios.MultipleDataStoreDataLatest, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) MapMultiObjectLatestPayload(ctx ciosctx.RequestCtx, strings []string, i interface{}) ([]cios.PackerFormatJsonHeader, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) MapMultiObjectLatestPayloadByChannels(ctx ciosctx.RequestCtx, channels []cios.Channel, i interface{}) ([]cios.PackerFormatJsonHeader, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) subscribeCiosWebSocket(ctx ciosctx.RequestCtx, s string, f *func(closeError *websocket.Conn), f2 func([]byte) (bool, error), i int64, i2 int64) error {
	panic("implement me")
}

func (NoImplementPubsub) GetStream(ctx ciosctx.RequestCtx, s string, request sdkmodel.ApiGetStreamRequest) ([]string, error) {
	panic("implement me")
}

func (NoImplementPubsub) GetStreamAll(ctx ciosctx.RequestCtx, s string, request sdkmodel.ApiGetStreamRequest) ([]string, error) {
	panic("implement me")
}

func (NoImplementPubsub) GetStreamUnlimited(ctx ciosctx.RequestCtx, s string, request sdkmodel.ApiGetStreamRequest) ([]string, error) {
	panic("implement me")
}

func (NoImplementPubsub) MapStreamAll(ctx ciosctx.RequestCtx, s string, request sdkmodel.ApiGetStreamRequest, i interface{}) error {
	panic("implement me")
}

func (NoImplementPubsub) MapStreamUnlimited(ctx ciosctx.RequestCtx, s string, request sdkmodel.ApiGetStreamRequest, i interface{}) error {
	panic("implement me")
}

func (NoImplementPubsub) GetJsonStreamUnlimited(ctx ciosctx.RequestCtx, s string, request sdkmodel.ApiGetStreamRequest) ([]cios.PackerFormatJson, error) {
	panic("implement me")
}

func (NoImplementPubsub) GetStreamFirst(ctx ciosctx.RequestCtx, s string, request sdkmodel.ApiGetStreamRequest) (string, error) {
	panic("implement me")
}

func (NoImplementPubsub) MapStreamFirst(ctx ciosctx.RequestCtx, s string, request sdkmodel.ApiGetStreamRequest, i interface{}) error {
	panic("implement me")
}

func (NoImplementPubsub) DeleteDataByChannel(ctx ciosctx.RequestCtx, s string) (*_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) DeleteObject(ctx ciosctx.RequestCtx, s string, s2 string) (*_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) NewMessaging(s string, mode enum.MessagingMode, format enum.PackerFormat) *cios.MessagingConfig {
	panic("implement me")
}

func (NoImplementPubsub) PublishMessage(ctx ciosctx.RequestCtx, s string, i interface{}, s2 *string) (*_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) PublishMessagePackerOnly(ctx ciosctx.RequestCtx, s string, i interface{}) (*_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) PublishMessageJSON(ctx ciosctx.RequestCtx, s string, json cios.PackerFormatJson) (*_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementPubsub) ConnectWebSocket(s string, bools chan bool, options srvpubsub.ConnectWebSocketOptions) error {
	panic("implement me")
}

func (NoImplementPubsub) CreateMessagingURL(s string, s2 string, s3 *string) string {
	panic("implement me")
}

func (NoImplementPubsub) CreateCIOSWebsocketConnection(s string, s2 string) (*websocket.Conn, error) {
	panic("implement me")
}

func (NoImplementPubsub) setDebug(b bool) {
	panic("implement me")
}

func (NoImplementPubsub) setToken(s string) {
	panic("implement me")
}
