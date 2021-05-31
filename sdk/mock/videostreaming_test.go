package ciossdmock_test

import (
	"testing"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	ciossdk_mock "github.com/optim-corp/cios-golang-sdk/sdk/mock"
	"github.com/stretchr/testify/assert"
)

func TestNoImplementVideoStreaming_GetThumbnail(t *testing.T) {
	mock := ciossdk_mock.NoImplementVideoStreaming{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetThumbnail(ciosctx.Background(), "")
}

func TestNoImplementVideoStreaming_GetVideoInfo(t *testing.T) {
	mock := ciossdk_mock.NoImplementVideoStreaming{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetVideoInfo(ciosctx.Background(), "")
}

func TestNoImplementVideoStreaming_GetVideoInfos(t *testing.T) {
	mock := ciossdk_mock.NoImplementVideoStreaming{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetVideoInfos(ciosctx.Background(), cios.ApiGetVideoStreamsListRequest{})
}

func TestNoImplementVideoStreaming_Play(t *testing.T) {
	mock := ciossdk_mock.NoImplementVideoStreaming{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.Play(ciosctx.Background(), "")
}

func TestNoImplementVideoStreaming_Stop(t *testing.T) {
	mock := ciossdk_mock.NoImplementVideoStreaming{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.Stop(ciosctx.Background(), "")
}

func TestNoImplementVideoStreaming_UpdateVideoInfo(t *testing.T) {
	mock := ciossdk_mock.NoImplementVideoStreaming{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.UpdateVideoInfo(ciosctx.Background(), "", "", "")
}
