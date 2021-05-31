package ciossdmock_test

import (
	"testing"

	srvpubsub "github.com/optim-corp/cios-golang-sdk/sdk/service/pubsub"

	ciossdk_mock "github.com/optim-corp/cios-golang-sdk/sdk/mock"
	"github.com/stretchr/testify/assert"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

func TestNoImplementPubsub_ConnectWebSocket(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.ConnectWebSocket("", make(chan bool), srvpubsub.ConnectWebSocketOptions{})
}

func TestNoImplementPubsub_CreateCIOSWebsocketConnection(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.CreateCIOSWebsocketConnection("", "")
}

func TestNoImplementPubsub_CreateChannel(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.CreateChannel(ciosctx.Background(), cios.ChannelProposal{})
}

func TestNoImplementPubsub_CreateMessagingURL(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.CreateMessagingURL("", "", nil)
}

func TestNoImplementPubsub_DeleteChannel(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.DeleteChannel(ciosctx.Background(), "")
}

func TestNoImplementPubsub_DeleteDataByChannel(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.DeleteDataByChannel(ciosctx.Background(), "")
}

func TestNoImplementPubsub_DeleteObject(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.DeleteObject(ciosctx.Background(), "", "")
}

func TestNoImplementPubsub_GetChannel(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetChannel(ciosctx.Background(), "", nil, nil)
}

func TestNoImplementPubsub_GetChannelFirst(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetChannelFirst(ciosctx.Background(), cios.ApiGetChannelsRequest{})
}

func TestNoImplementPubsub_GetChannels(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetChannels(ciosctx.Background(), cios.ApiGetChannelsRequest{})
}

func TestNoImplementPubsub_GetChannelsAll(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetChannelsAll(ciosctx.Background(), cios.ApiGetChannelsRequest{})
}

func TestNoImplementPubsub_GetChannelsMapByID(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetChannelsUnlimited(ciosctx.Background(), cios.ApiGetChannelsRequest{})
}

func TestNoImplementPubsub_GetChannelsMapByResourceOwnerID(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetChannelsMapByResourceOwnerID(ciosctx.Background(), cios.ApiGetChannelsRequest{})
}

func TestNoImplementPubsub_GetChannelsUnlimited(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetChannelsUnlimited(ciosctx.Background(), cios.ApiGetChannelsRequest{})
}

func TestNoImplementPubsub_GetDataStoreChannel(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetDataStoreChannel(ciosctx.Background(), "")
}

func TestNoImplementPubsub_GetDataStoreChannels(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetDataStoreChannels(ciosctx.Background(), cios.ApiGetDataStoreChannelsRequest{})
}

func TestNoImplementPubsub_GetJsonStreamUnlimited(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetJsonStreamUnlimited(ciosctx.Background(), "", sdkmodel.ApiGetStreamRequest{})
}

func TestNoImplementPubsub_GetMultiObjectLatest(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetMultiObjectLatest(ciosctx.Background(), nil)
}

func TestNoImplementPubsub_GetMultiObjectLatestByChannels(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetMultiObjectLatestByChannels(ciosctx.Background(), nil)
}

func TestNoImplementPubsub_GetObject(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetObject(ciosctx.Background(), "", "", nil)
}

func TestNoImplementPubsub_GetObjectLatest(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetObjectLatest(ciosctx.Background(), "", nil)
}

func TestNoImplementPubsub_GetObjects(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetObjects(ciosctx.Background(), "", cios.ApiGetDataStoreObjectsRequest{})
}

func TestNoImplementPubsub_GetObjectsAll(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetObjectsAll(ciosctx.Background(), "", cios.ApiGetDataStoreObjectsRequest{})
}

func TestNoImplementPubsub_GetObjectsUnlimited(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetObjectsUnlimited(ciosctx.Background(), "", cios.ApiGetDataStoreObjectsRequest{})
}

func TestNoImplementPubsub_GetOrCreateChannel(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetOrCreateChannel(ciosctx.Background(), cios.ApiGetChannelsRequest{}, cios.ChannelProposal{})
}

func TestNoImplementPubsub_GetStream(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetStream(ciosctx.Background(), "", sdkmodel.ApiGetStreamRequest{})
}

func TestNoImplementPubsub_GetStreamAll(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetStreamAll(ciosctx.Background(), "", sdkmodel.ApiGetStreamRequest{})
}

func TestNoImplementPubsub_GetStreamFirst(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetStreamFirst(ciosctx.Background(), "", sdkmodel.ApiGetStreamRequest{})
}

func TestNoImplementPubsub_GetStreamUnlimited(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetStreamUnlimited(ciosctx.Background(), "", sdkmodel.ApiGetStreamRequest{})
}

func TestNoImplementPubsub_MapMultiObjectLatestPayload(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.MapMultiObjectLatestPayload(ciosctx.Background(), nil, struct{}{})
}

func TestNoImplementPubsub_MapMultiObjectLatestPayloadByChannels(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.MapMultiObjectLatestPayloadByChannels(ciosctx.Background(), nil, struct{}{})
}

func TestNoImplementPubsub_MapObjectLatest(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.MapObjectLatest(ciosctx.Background(), "", nil, struct{}{})
}

func TestNoImplementPubsub_MapStreamAll(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.MapStreamAll(ciosctx.Background(), "", sdkmodel.ApiGetStreamRequest{}, struct{}{})
}

func TestNoImplementPubsub_MapStreamFirst(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.MapStreamFirst(ciosctx.Background(), "", sdkmodel.ApiGetStreamRequest{}, struct{}{})
}

func TestNoImplementPubsub_MapStreamUnlimited(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.MapStreamUnlimited(ciosctx.Background(), "", sdkmodel.ApiGetStreamRequest{}, struct{}{})
}

func TestNoImplementPubsub_NewMessaging(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.NewMessaging("", "", "")
}

func TestNoImplementPubsub_PublishMessage(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.PublishMessage(ciosctx.Background(), "", "", nil)
}

func TestNoImplementPubsub_PublishMessageJSON(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.PublishMessageJSON(ciosctx.Background(), "", cios.PackerFormatJson{})
}

func TestNoImplementPubsub_PublishMessagePackerOnly(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.PublishMessagePackerOnly(ciosctx.Background(), "", struct{}{})
}

func TestNoImplementPubsub_UpdateChannel(t *testing.T) {
	mock := ciossdk_mock.NoImplementPubsub{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.UpdateChannel(ciosctx.Background(), "", cios.ChannelUpdateProposal{})
}
