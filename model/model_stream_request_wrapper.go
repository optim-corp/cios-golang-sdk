package sdkmodel

import (
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
	OffsetParam                 *int64
	LimitParam                  *int64
	TimeoutParam                *int
	SessionIdParam              *string
	ChannelProtocolIdParam      *string
}

func (r ApiGetStreamRequest) PackerFormat(packerPacker string) ApiGetStreamRequest {
	r.PackerFormatParam = util.ToNilStr(packerPacker)
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
	r.SessionIdParam = util.ToNilStr(sessionID)
	return r
}
func (r ApiGetStreamRequest) ChannelProtocolId(channelProtocolId string) ApiGetStreamRequest {
	r.ChannelProtocolIdParam = util.ToNilStr(channelProtocolId)
	return r
}
func (r ApiGetStreamRequest) Label(label string) ApiGetStreamRequest {
	r.LabelParam = util.ToNilStr(label)
	return r
}
func (r ApiGetStreamRequest) Location(location string) ApiGetStreamRequest {
	r.LocationParam = util.ToNilStr(location)
	return r
}
func (r ApiGetStreamRequest) LocationRange(locationRange string) ApiGetStreamRequest {
	r.LocationRangeParam = util.ToNilStr(locationRange)
	return r
}
func (r ApiGetStreamRequest) Timestamp(timestamp string) ApiGetStreamRequest {
	r.TimestampParam = util.ToNilStr(timestamp)
	return r
}
func (r ApiGetStreamRequest) TimestampRange(timestampRange string) ApiGetStreamRequest {
	r.TimestampRangeParam = util.ToNilStr(timestampRange)
	return r
}
func (r ApiGetStreamRequest) Ascending(ascending bool) ApiGetStreamRequest {
	r.AscendingParam = &ascending
	return r
}
func (r ApiGetStreamRequest) Offset(offset int64) ApiGetStreamRequest {
	r.OffsetParam = &offset
	return r
}
func (r ApiGetStreamRequest) Limit(limit int64) ApiGetStreamRequest {
	r.LimitParam = &limit
	return r
}
