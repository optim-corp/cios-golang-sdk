package sdkmodel

import (
	"reflect"
	"testing"

	cnv "github.com/fcfcqloow/go-advance/convert"

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
		{
			name:   "Test ascending true",
			fields: fields{},
			args:   args{false},
			want: ApiGetStreamRequest{
				AscendingParam: cnv.BoolPtr(false),
			},
		},
		{
			name:   "Test ascending true",
			fields: fields{},
			args:   args{true},
			want: ApiGetStreamRequest{
				AscendingParam: cnv.BoolPtr(true),
			},
		},
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
		{
			name:   "Test channelProtocolId true",
			fields: fields{},
			args:   args{""},
			want: ApiGetStreamRequest{
				ChannelProtocolIdParam: nil,
			},
		},
		{
			name:   "Test channelProtocolId true",
			fields: fields{},
			args:   args{"test"},
			want: ApiGetStreamRequest{
				ChannelProtocolIdParam: cnv.StrPtr("test"),
			},
		},
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
		{
			name:   "Test channelProtocolVersion true",
			fields: fields{},
			args:   args{1},
			want: ApiGetStreamRequest{
				ChannelProtocolVersionParam: cnv.Int32Ptr(1),
			},
		},
		{
			name:   "Test channelProtocolVersion true",
			fields: fields{},
			args:   args{12},
			want: ApiGetStreamRequest{
				ChannelProtocolVersionParam: cnv.Int32Ptr(12),
			},
		},
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
		{
			name:   "Test label true",
			fields: fields{},
			args:   args{""},
			want: ApiGetStreamRequest{
				LabelParam: nil,
			},
		},
		{
			name:   "Test label true",
			fields: fields{},
			args:   args{"test"},
			want: ApiGetStreamRequest{
				LabelParam: cnv.StrPtr("test"),
			},
		},
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
		{
			name:   "Test limit true",
			fields: fields{},
			args:   args{21},
			want: ApiGetStreamRequest{
				LimitParam: cnv.Int64Ptr(21),
			},
		},
		{
			name:   "Test limit true",
			fields: fields{},
			args:   args{0},
			want: ApiGetStreamRequest{
				LimitParam: cnv.Int64Ptr(0),
			},
		},
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
		{
			name:   "Test location true",
			fields: fields{},
			args:   args{""},
			want: ApiGetStreamRequest{
				LocationParam: nil,
			},
		},
		{
			name:   "Test location true",
			fields: fields{},
			args:   args{"test"},
			want: ApiGetStreamRequest{
				LocationParam: cnv.StrPtr("test"),
			},
		},
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
		{
			name:   "Test locationRange true",
			fields: fields{},
			args:   args{""},
			want: ApiGetStreamRequest{
				LocationRangeParam: nil,
			},
		},
		{
			name:   "Test locationRange true",
			fields: fields{},
			args:   args{"test"},
			want: ApiGetStreamRequest{
				LocationRangeParam: cnv.StrPtr("test"),
			},
		},
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
		{
			name:   "Test offset true",
			fields: fields{},
			args:   args{0},
			want: ApiGetStreamRequest{
				OffsetParam: cnv.Int64Ptr(0),
			},
		},
		{
			name:   "Test offset true",
			fields: fields{},
			args:   args{21},
			want: ApiGetStreamRequest{
				OffsetParam: cnv.Int64Ptr(21),
			},
		},
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
		{
			name:   "Test packerPacker true",
			fields: fields{},
			args:   args{""},
			want: ApiGetStreamRequest{
				PackerFormatParam: nil,
			},
		},
		{
			name:   "Test packerPacker true",
			fields: fields{},
			args:   args{"test"},
			want: ApiGetStreamRequest{
				PackerFormatParam: cnv.StrPtr("test"),
			},
		},
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
		{
			name:   "Test sessionID true",
			fields: fields{},
			args:   args{""},
			want: ApiGetStreamRequest{
				SessionIdParam: nil,
			},
		},
		{
			name:   "Test sessionID true",
			fields: fields{},
			args:   args{"test"},
			want: ApiGetStreamRequest{
				SessionIdParam: cnv.StrPtr("test"),
			},
		},
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
		{
			name:   "Test timeout true",
			fields: fields{},
			args:   args{0},
			want: ApiGetStreamRequest{
				TimeoutParam: cnv.IntPtr(0),
			},
		},
		{
			name:   "Test timeout true",
			fields: fields{},
			args:   args{12},
			want: ApiGetStreamRequest{
				TimeoutParam: cnv.IntPtr(12),
			},
		},
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
		{
			name:   "Test timestamp true",
			fields: fields{},
			args:   args{""},
			want: ApiGetStreamRequest{
				TimestampParam: nil,
			},
		},
		{
			name:   "Test timestamp true",
			fields: fields{},
			args:   args{"test"},
			want: ApiGetStreamRequest{
				TimestampParam: cnv.StrPtr("test"),
			},
		},
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
		{
			name:   "Test timestampRange true",
			fields: fields{},
			args:   args{""},
			want: ApiGetStreamRequest{
				TimestampRangeParam: nil,
			},
		},
		{
			name:   "Test timestampRange true",
			fields: fields{},
			args:   args{"test"},
			want: ApiGetStreamRequest{
				TimestampRangeParam: cnv.StrPtr("test"),
			},
		},
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
