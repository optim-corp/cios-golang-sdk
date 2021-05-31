package sdkmodel

import (
	"reflect"
	"testing"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

func TestApiGetStreamRequest_Ascending(t *testing.T) {
	type fields struct {
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
	type args struct {
		ascending bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiGetStreamRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetStreamRequest{
				ApiService:                  tt.fields.ApiService,
				ChannelId:                   tt.fields.ChannelId,
				PackerFormatParam:           tt.fields.PackerFormatParam,
				AscendingParam:              tt.fields.AscendingParam,
				ChannelProtocolVersionParam: tt.fields.ChannelProtocolVersionParam,
				LocationParam:               tt.fields.LocationParam,
				LocationRangeParam:          tt.fields.LocationRangeParam,
				TimestampParam:              tt.fields.TimestampParam,
				TimestampRangeParam:         tt.fields.TimestampRangeParam,
				LabelParam:                  tt.fields.LabelParam,
				OffsetParam:                 tt.fields.OffsetParam,
				LimitParam:                  tt.fields.LimitParam,
				TimeoutParam:                tt.fields.TimeoutParam,
				SessionIdParam:              tt.fields.SessionIdParam,
				ChannelProtocolIdParam:      tt.fields.ChannelProtocolIdParam,
			}
			if got := r.Ascending(tt.args.ascending); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ascending() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiGetStreamRequest_ChannelProtocolId(t *testing.T) {
	type fields struct {
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
	type args struct {
		channelProtocolId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiGetStreamRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetStreamRequest{
				ApiService:                  tt.fields.ApiService,
				ChannelId:                   tt.fields.ChannelId,
				PackerFormatParam:           tt.fields.PackerFormatParam,
				AscendingParam:              tt.fields.AscendingParam,
				ChannelProtocolVersionParam: tt.fields.ChannelProtocolVersionParam,
				LocationParam:               tt.fields.LocationParam,
				LocationRangeParam:          tt.fields.LocationRangeParam,
				TimestampParam:              tt.fields.TimestampParam,
				TimestampRangeParam:         tt.fields.TimestampRangeParam,
				LabelParam:                  tt.fields.LabelParam,
				OffsetParam:                 tt.fields.OffsetParam,
				LimitParam:                  tt.fields.LimitParam,
				TimeoutParam:                tt.fields.TimeoutParam,
				SessionIdParam:              tt.fields.SessionIdParam,
				ChannelProtocolIdParam:      tt.fields.ChannelProtocolIdParam,
			}
			if got := r.ChannelProtocolId(tt.args.channelProtocolId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChannelProtocolId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiGetStreamRequest_ChannelProtocolVersion(t *testing.T) {
	type fields struct {
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
	type args struct {
		channelProtocolVersion int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiGetStreamRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetStreamRequest{
				ApiService:                  tt.fields.ApiService,
				ChannelId:                   tt.fields.ChannelId,
				PackerFormatParam:           tt.fields.PackerFormatParam,
				AscendingParam:              tt.fields.AscendingParam,
				ChannelProtocolVersionParam: tt.fields.ChannelProtocolVersionParam,
				LocationParam:               tt.fields.LocationParam,
				LocationRangeParam:          tt.fields.LocationRangeParam,
				TimestampParam:              tt.fields.TimestampParam,
				TimestampRangeParam:         tt.fields.TimestampRangeParam,
				LabelParam:                  tt.fields.LabelParam,
				OffsetParam:                 tt.fields.OffsetParam,
				LimitParam:                  tt.fields.LimitParam,
				TimeoutParam:                tt.fields.TimeoutParam,
				SessionIdParam:              tt.fields.SessionIdParam,
				ChannelProtocolIdParam:      tt.fields.ChannelProtocolIdParam,
			}
			if got := r.ChannelProtocolVersion(tt.args.channelProtocolVersion); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChannelProtocolVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiGetStreamRequest_Label(t *testing.T) {
	type fields struct {
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
	type args struct {
		label string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiGetStreamRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetStreamRequest{
				ApiService:                  tt.fields.ApiService,
				ChannelId:                   tt.fields.ChannelId,
				PackerFormatParam:           tt.fields.PackerFormatParam,
				AscendingParam:              tt.fields.AscendingParam,
				ChannelProtocolVersionParam: tt.fields.ChannelProtocolVersionParam,
				LocationParam:               tt.fields.LocationParam,
				LocationRangeParam:          tt.fields.LocationRangeParam,
				TimestampParam:              tt.fields.TimestampParam,
				TimestampRangeParam:         tt.fields.TimestampRangeParam,
				LabelParam:                  tt.fields.LabelParam,
				OffsetParam:                 tt.fields.OffsetParam,
				LimitParam:                  tt.fields.LimitParam,
				TimeoutParam:                tt.fields.TimeoutParam,
				SessionIdParam:              tt.fields.SessionIdParam,
				ChannelProtocolIdParam:      tt.fields.ChannelProtocolIdParam,
			}
			if got := r.Label(tt.args.label); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Label() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiGetStreamRequest_Limit(t *testing.T) {
	type fields struct {
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
	type args struct {
		limit int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiGetStreamRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetStreamRequest{
				ApiService:                  tt.fields.ApiService,
				ChannelId:                   tt.fields.ChannelId,
				PackerFormatParam:           tt.fields.PackerFormatParam,
				AscendingParam:              tt.fields.AscendingParam,
				ChannelProtocolVersionParam: tt.fields.ChannelProtocolVersionParam,
				LocationParam:               tt.fields.LocationParam,
				LocationRangeParam:          tt.fields.LocationRangeParam,
				TimestampParam:              tt.fields.TimestampParam,
				TimestampRangeParam:         tt.fields.TimestampRangeParam,
				LabelParam:                  tt.fields.LabelParam,
				OffsetParam:                 tt.fields.OffsetParam,
				LimitParam:                  tt.fields.LimitParam,
				TimeoutParam:                tt.fields.TimeoutParam,
				SessionIdParam:              tt.fields.SessionIdParam,
				ChannelProtocolIdParam:      tt.fields.ChannelProtocolIdParam,
			}
			if got := r.Limit(tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Limit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiGetStreamRequest_Location(t *testing.T) {
	type fields struct {
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
	type args struct {
		location string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiGetStreamRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetStreamRequest{
				ApiService:                  tt.fields.ApiService,
				ChannelId:                   tt.fields.ChannelId,
				PackerFormatParam:           tt.fields.PackerFormatParam,
				AscendingParam:              tt.fields.AscendingParam,
				ChannelProtocolVersionParam: tt.fields.ChannelProtocolVersionParam,
				LocationParam:               tt.fields.LocationParam,
				LocationRangeParam:          tt.fields.LocationRangeParam,
				TimestampParam:              tt.fields.TimestampParam,
				TimestampRangeParam:         tt.fields.TimestampRangeParam,
				LabelParam:                  tt.fields.LabelParam,
				OffsetParam:                 tt.fields.OffsetParam,
				LimitParam:                  tt.fields.LimitParam,
				TimeoutParam:                tt.fields.TimeoutParam,
				SessionIdParam:              tt.fields.SessionIdParam,
				ChannelProtocolIdParam:      tt.fields.ChannelProtocolIdParam,
			}
			if got := r.Location(tt.args.location); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Location() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiGetStreamRequest_LocationRange(t *testing.T) {
	type fields struct {
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
	type args struct {
		locationRange string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiGetStreamRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetStreamRequest{
				ApiService:                  tt.fields.ApiService,
				ChannelId:                   tt.fields.ChannelId,
				PackerFormatParam:           tt.fields.PackerFormatParam,
				AscendingParam:              tt.fields.AscendingParam,
				ChannelProtocolVersionParam: tt.fields.ChannelProtocolVersionParam,
				LocationParam:               tt.fields.LocationParam,
				LocationRangeParam:          tt.fields.LocationRangeParam,
				TimestampParam:              tt.fields.TimestampParam,
				TimestampRangeParam:         tt.fields.TimestampRangeParam,
				LabelParam:                  tt.fields.LabelParam,
				OffsetParam:                 tt.fields.OffsetParam,
				LimitParam:                  tt.fields.LimitParam,
				TimeoutParam:                tt.fields.TimeoutParam,
				SessionIdParam:              tt.fields.SessionIdParam,
				ChannelProtocolIdParam:      tt.fields.ChannelProtocolIdParam,
			}
			if got := r.LocationRange(tt.args.locationRange); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LocationRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiGetStreamRequest_Offset(t *testing.T) {
	type fields struct {
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
	type args struct {
		offset int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiGetStreamRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetStreamRequest{
				ApiService:                  tt.fields.ApiService,
				ChannelId:                   tt.fields.ChannelId,
				PackerFormatParam:           tt.fields.PackerFormatParam,
				AscendingParam:              tt.fields.AscendingParam,
				ChannelProtocolVersionParam: tt.fields.ChannelProtocolVersionParam,
				LocationParam:               tt.fields.LocationParam,
				LocationRangeParam:          tt.fields.LocationRangeParam,
				TimestampParam:              tt.fields.TimestampParam,
				TimestampRangeParam:         tt.fields.TimestampRangeParam,
				LabelParam:                  tt.fields.LabelParam,
				OffsetParam:                 tt.fields.OffsetParam,
				LimitParam:                  tt.fields.LimitParam,
				TimeoutParam:                tt.fields.TimeoutParam,
				SessionIdParam:              tt.fields.SessionIdParam,
				ChannelProtocolIdParam:      tt.fields.ChannelProtocolIdParam,
			}
			if got := r.Offset(tt.args.offset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Offset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiGetStreamRequest_PackerFormat(t *testing.T) {
	type fields struct {
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
	type args struct {
		packerPacker string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiGetStreamRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetStreamRequest{
				ApiService:                  tt.fields.ApiService,
				ChannelId:                   tt.fields.ChannelId,
				PackerFormatParam:           tt.fields.PackerFormatParam,
				AscendingParam:              tt.fields.AscendingParam,
				ChannelProtocolVersionParam: tt.fields.ChannelProtocolVersionParam,
				LocationParam:               tt.fields.LocationParam,
				LocationRangeParam:          tt.fields.LocationRangeParam,
				TimestampParam:              tt.fields.TimestampParam,
				TimestampRangeParam:         tt.fields.TimestampRangeParam,
				LabelParam:                  tt.fields.LabelParam,
				OffsetParam:                 tt.fields.OffsetParam,
				LimitParam:                  tt.fields.LimitParam,
				TimeoutParam:                tt.fields.TimeoutParam,
				SessionIdParam:              tt.fields.SessionIdParam,
				ChannelProtocolIdParam:      tt.fields.ChannelProtocolIdParam,
			}
			if got := r.PackerFormat(tt.args.packerPacker); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PackerFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiGetStreamRequest_SessionID(t *testing.T) {
	type fields struct {
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
	type args struct {
		sessionID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiGetStreamRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetStreamRequest{
				ApiService:                  tt.fields.ApiService,
				ChannelId:                   tt.fields.ChannelId,
				PackerFormatParam:           tt.fields.PackerFormatParam,
				AscendingParam:              tt.fields.AscendingParam,
				ChannelProtocolVersionParam: tt.fields.ChannelProtocolVersionParam,
				LocationParam:               tt.fields.LocationParam,
				LocationRangeParam:          tt.fields.LocationRangeParam,
				TimestampParam:              tt.fields.TimestampParam,
				TimestampRangeParam:         tt.fields.TimestampRangeParam,
				LabelParam:                  tt.fields.LabelParam,
				OffsetParam:                 tt.fields.OffsetParam,
				LimitParam:                  tt.fields.LimitParam,
				TimeoutParam:                tt.fields.TimeoutParam,
				SessionIdParam:              tt.fields.SessionIdParam,
				ChannelProtocolIdParam:      tt.fields.ChannelProtocolIdParam,
			}
			if got := r.SessionID(tt.args.sessionID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SessionID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiGetStreamRequest_Timeout(t *testing.T) {
	type fields struct {
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
	type args struct {
		timeout int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiGetStreamRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetStreamRequest{
				ApiService:                  tt.fields.ApiService,
				ChannelId:                   tt.fields.ChannelId,
				PackerFormatParam:           tt.fields.PackerFormatParam,
				AscendingParam:              tt.fields.AscendingParam,
				ChannelProtocolVersionParam: tt.fields.ChannelProtocolVersionParam,
				LocationParam:               tt.fields.LocationParam,
				LocationRangeParam:          tt.fields.LocationRangeParam,
				TimestampParam:              tt.fields.TimestampParam,
				TimestampRangeParam:         tt.fields.TimestampRangeParam,
				LabelParam:                  tt.fields.LabelParam,
				OffsetParam:                 tt.fields.OffsetParam,
				LimitParam:                  tt.fields.LimitParam,
				TimeoutParam:                tt.fields.TimeoutParam,
				SessionIdParam:              tt.fields.SessionIdParam,
				ChannelProtocolIdParam:      tt.fields.ChannelProtocolIdParam,
			}
			if got := r.Timeout(tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Timeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiGetStreamRequest_Timestamp(t *testing.T) {
	type fields struct {
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
	type args struct {
		timestamp string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiGetStreamRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetStreamRequest{
				ApiService:                  tt.fields.ApiService,
				ChannelId:                   tt.fields.ChannelId,
				PackerFormatParam:           tt.fields.PackerFormatParam,
				AscendingParam:              tt.fields.AscendingParam,
				ChannelProtocolVersionParam: tt.fields.ChannelProtocolVersionParam,
				LocationParam:               tt.fields.LocationParam,
				LocationRangeParam:          tt.fields.LocationRangeParam,
				TimestampParam:              tt.fields.TimestampParam,
				TimestampRangeParam:         tt.fields.TimestampRangeParam,
				LabelParam:                  tt.fields.LabelParam,
				OffsetParam:                 tt.fields.OffsetParam,
				LimitParam:                  tt.fields.LimitParam,
				TimeoutParam:                tt.fields.TimeoutParam,
				SessionIdParam:              tt.fields.SessionIdParam,
				ChannelProtocolIdParam:      tt.fields.ChannelProtocolIdParam,
			}
			if got := r.Timestamp(tt.args.timestamp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Timestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiGetStreamRequest_TimestampRange(t *testing.T) {
	type fields struct {
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
	type args struct {
		timestampRange string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ApiGetStreamRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ApiGetStreamRequest{
				ApiService:                  tt.fields.ApiService,
				ChannelId:                   tt.fields.ChannelId,
				PackerFormatParam:           tt.fields.PackerFormatParam,
				AscendingParam:              tt.fields.AscendingParam,
				ChannelProtocolVersionParam: tt.fields.ChannelProtocolVersionParam,
				LocationParam:               tt.fields.LocationParam,
				LocationRangeParam:          tt.fields.LocationRangeParam,
				TimestampParam:              tt.fields.TimestampParam,
				TimestampRangeParam:         tt.fields.TimestampRangeParam,
				LabelParam:                  tt.fields.LabelParam,
				OffsetParam:                 tt.fields.OffsetParam,
				LimitParam:                  tt.fields.LimitParam,
				TimeoutParam:                tt.fields.TimeoutParam,
				SessionIdParam:              tt.fields.SessionIdParam,
				ChannelProtocolIdParam:      tt.fields.ChannelProtocolIdParam,
			}
			if got := r.TimestampRange(tt.args.timestampRange); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimestampRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
