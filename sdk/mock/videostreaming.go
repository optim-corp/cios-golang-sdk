package ciossdmock

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
)

type NoImplementVideoStreaming struct{}

func (NoImplementVideoStreaming) SetToken(s2 string) {
	panic("implement me")
}

func (NoImplementVideoStreaming) GetVideoInfos(ctx ciosctx.RequestCtx, request cios.ApiGetVideoStreamsListRequest) ([]cios.Video, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementVideoStreaming) GetVideoInfo(ctx ciosctx.RequestCtx, s string) (cios.Video, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementVideoStreaming) UpdateVideoInfo(ctx ciosctx.RequestCtx, s string, s2 string, s3 string) (cios.Video, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementVideoStreaming) GetThumbnail(ctx ciosctx.RequestCtx, s string) ([]byte, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementVideoStreaming) Play(ctx ciosctx.RequestCtx, s string) (cios.Room, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementVideoStreaming) Stop(ctx ciosctx.RequestCtx, s string) (*_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementVideoStreaming) setToken(s string) {
	panic("implement me")
}
