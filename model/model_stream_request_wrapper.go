package model

import (
	"context"

	"github.com/optim-corp/cios-golang-sdk/util"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

type ApiGetStreamRequest struct {
	ApiService                  *cios.PublishSubscribeApiService
	ChannelId                   string
	PackerFormatParam           *string
	AscendingParam              *bool
	ChannelProtocolVersionParam *int32
	LocationParam               *string
	LocationRangeParam          *string
	TimestampParam              *string
	TimestampRangeParam         *string
	LabelParam                  *string
	P_offset                    *int64
	P_limit                     *int64
	TimeoutParam                *int
	SessionIdParam              *string
	ChannelProtocolIdParam      *string
	Context                     context.Context
}

func (r ApiGetStreamRequest) PackerFormat(packerPacker string) ApiGetStreamRequest {
	r.PackerFormatParam = util.ToNil(packerPacker)
	return r
}
func (r ApiGetStreamRequest) ChannelProtocolVersion(channelProtocolVersion int32) ApiGetStreamRequest {
	r.ChannelProtocolVersionParam = &channelProtocolVersion
	return r
}
func (r ApiGetStreamRequest) Timeout(timeout int) ApiGetStreamRequest {
	r.TimeoutParam = &timeout
	return r
}
func (r ApiGetStreamRequest) SessionID(sessionID string) ApiGetStreamRequest {
	r.SessionIdParam = util.ToNil(sessionID)
	return r
}
func (r ApiGetStreamRequest) ChannelProtocolId(channelProtocolId string) ApiGetStreamRequest {
	r.ChannelProtocolIdParam = util.ToNil(channelProtocolId)
	return r
}
func (r ApiGetStreamRequest) Label(label string) ApiGetStreamRequest {
	r.LabelParam = util.ToNil(label)
	return r
}
func (r ApiGetStreamRequest) Location(location string) ApiGetStreamRequest {
	r.LocationParam = util.ToNil(location)
	return r
}
func (r ApiGetStreamRequest) LocationRange(locationRange string) ApiGetStreamRequest {
	r.LocationRangeParam = util.ToNil(locationRange)
	return r
}
func (r ApiGetStreamRequest) Timestamp(timestamp string) ApiGetStreamRequest {
	r.TimestampParam = util.ToNil(timestamp)
	return r
}
func (r ApiGetStreamRequest) TimestampRange(timestampRange string) ApiGetStreamRequest {
	r.TimestampRangeParam = util.ToNil(timestampRange)
	return r
}
func (r ApiGetStreamRequest) Ascending(ascending bool) ApiGetStreamRequest {
	r.AscendingParam = &ascending
	return r
}
func (r ApiGetStreamRequest) Offset(offset int64) ApiGetStreamRequest {
	r.P_offset = &offset
	return r
}
func (r ApiGetStreamRequest) Limit(limit int64) ApiGetStreamRequest {
	r.P_limit = &limit
	return r
}
