package ciossdk

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	srvpubsub "github.com/optim-corp/cios-golang-sdk/sdk/service/pubsub"

	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"

	xmath "github.com/fcfcqloow/go-advance/math"

	cnv "github.com/fcfcqloow/go-advance/convert"

	"github.com/optim-corp/cios-golang-sdk/cios"

	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

func TestPubSub_Channels(t *testing.T) {
	var (
		query url.Values
		ctx   context.Context

		tests = []struct {
			params cios.ApiGetChannelsRequest
			test   func()
		}{
			{
				params: srvpubsub.MakeGetChannelsOpts().Limit(1000),
				test: func() {
					if query.Get("limit") != "1000" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvpubsub.MakeGetChannelsOpts().Limit(1000).Offset(50),
				test: func() {
					if query.Get("limit") != "1000" || query.Get("offset") != "50" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvpubsub.MakeGetChannelsOpts().ResourceOwnerId("aaaaa"),
				test: func() {
					if query.Get("resource_owner_id") != "aaaaa" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvpubsub.MakeGetChannelsOpts().Limit(1000).Offset(50).ResourceOwnerId("aaaaa").Name("name"),
				test: func() {
					if query.Get("resource_owner_id") != "aaaaa" ||
						query.Get("name") != "name" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvpubsub.MakeGetChannelsOpts().OrderBy("created_at"),
				test: func() {
					if query.Get("order_by") != "created_at" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvpubsub.MakeGetChannelsOpts().
					OrderBy("").
					Order("").
					ResourceOwnerId("").
					Name("").
					Label("").
					Lang("").
					MessagingEnabled("").
					DatastoreEnabled("").
					MessagingPersisted(""),
				test: func() {
					if query.Encode() != "" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
		}
	)

	// Query Test
	responseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query = r.URL.Query()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cios.MultipleChannel{Total: 10})
	})
	ts := httptest.NewServer(responseHandler)
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{MessagingUrl: ts.URL}})
	defer ts.Close()
	for _, test := range tests {
		client.PubSub().GetChannels(ctx, test.params)
		test.test()
	}

	ts.Close()
}

func TestPubSub_GetChannelsAll(t *testing.T) {
	var offsets []int
	var limits []int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleChannel{Total: 3500, Channels: []cios.Channel{}}
		offset := cnv.MustInt(r.URL.Query().Get("offset"))
		limit := cnv.MustInt(r.URL.Query().Get("limit"))
		offsets = append(offsets, offset)
		limits = append(limits, limit)
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Channels = append(response.Channels, cios.Channel{Id: cnv.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{MessagingUrl: ts.URL}})

	responses, _, _ := client.PubSub().GetChannelsAll(ciosctx.Background(), srvpubsub.MakeGetChannelsOpts().Limit(999))
	if len(responses) != 999 || offsets[0] != 0 && limits[0] != 1000 {
		t.Fatal(len(responses))
	}

	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.PubSub().GetChannelsAll(ciosctx.Background(), srvpubsub.MakeGetChannelsOpts().Limit(1500))
	if len(responses) != 1500 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 {
		t.Fatal(len(responses), limits, offsets)
	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.PubSub().GetChannelsAll(ciosctx.Background(), srvpubsub.MakeGetChannelsOpts().Limit(2001))
	if len(responses) != 2001 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 || offsets[2] != 2000 || limits[2] != 1 {
		t.Fatal(len(responses), limits, offsets)

	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.PubSub().GetChannelsAll(ciosctx.Background(), srvpubsub.MakeGetChannelsOpts().Limit(3501))
	if len(responses) != 3500 ||
		offsets[0] != 0 || limits[0] != 1000 ||
		offsets[1] != 1000 && limits[1] != 1000 ||
		offsets[2] != 2000 || limits[2] != 1000 ||
		offsets[3] != 3000 || limits[3] != 501 {
		t.Fatal(len(responses), limits, offsets)
	}

	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.PubSub().GetChannelsAll(ciosctx.Background(), srvpubsub.MakeGetChannelsOpts().Limit(2001).Offset(20))
	if len(responses) != 2001 || offsets[0] != 20 && limits[0] != 1000 || offsets[1] != 1020 && limits[1] != 1000 || offsets[2] != 2020 || limits[2] != 1 {
		t.Fatal(len(responses), limits, offsets)

	}
}

func TestPubSub_GetChannelsUnlimited(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleChannel{Total: 3500, Channels: []cios.Channel{}}
		offset := cnv.MustInt(r.URL.Query().Get("offset"))
		limit := cnv.MustInt(r.URL.Query().Get("limit"))
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Channels = append(response.Channels, cios.Channel{Id: cnv.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{MessagingUrl: ts.URL}})

	response, _, _ := client.PubSub().GetChannelsUnlimited(ciosctx.Background(), srvpubsub.MakeGetChannelsOpts().Limit(1))
	if len(response) != 3500 {
		t.Fatal(len(response))
	}
}

func TestPubSub_GetChannel(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/v2/channels/test" {
			response := cios.SingleChannel{Channel: cios.Channel{
				Id:              "test",
				ResourceOwnerId: "test_resource_owner",
				CreatedAt:       "",
				UpdatedAt:       "",
				Name:            "",
			}}
			json.NewEncoder(w).Encode(response)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{MessagingUrl: ts.URL}})
	responseB, response, err := client.PubSub().GetChannel(ciosctx.Background(), "test", nil, nil)
	if responseB.Id != "test" || err != nil || response.StatusCode != 200 {
		t.Fatal(responseB)
	}
}

func TestPubSub_CreateChannel(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body := cios.ChannelProposal{}
		cnv.UnMarshalJson(r.Body, &body)
		if body.ResourceOwnerId != "resource_owner_id" ||
			*body.DatastoreConfig.Enabled != true ||
			*body.DatastoreConfig.MaxCount != "222" ||
			*body.DatastoreConfig.MaxSize != "1000" ||
			*body.MessagingConfig.Enabled != true ||
			*body.MessagingConfig.Persisted != true ||
			(*body.Labels)[0].Key != "test1" ||
			(*body.Labels)[0].Value != "test1" ||
			(*body.Labels)[1].Key != "test2" ||
			(*body.Labels)[1].Value != "test2" ||
			body.DisplayInfo[0].Name != "test" ||
			body.DisplayInfo[0].Language != "ja" {
			t.Fatal(body, *body.Labels, *body.MessagingConfig.Enabled, *body.DatastoreConfig.MaxSize)
		}
		if r.Method != "POST" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{MessagingUrl: ts.URL}})
	_, _, err := client.PubSub().CreateChannel(ciosctx.Background(), cios.ChannelProposal{
		ResourceOwnerId:  "resource_owner_id",
		ChannelProtocols: nil,
		DisplayInfo: cios.DisplayInfoStream{
			{
				Name:      "test",
				Language:  "ja",
				IsDefault: true,
			},
		},
		Labels: &[]cios.Label{
			{
				Key:   "test1",
				Value: "test1",
			},
			{
				Key:   "test2",
				Value: "test2",
			},
		},
		MessagingConfig: &cios.MessagingConfig{
			Enabled:   cnv.BoolPtr(true),
			Persisted: cnv.BoolPtr(true),
		},
		DatastoreConfig: &cios.DataStoreConfig{
			Enabled:  cnv.BoolPtr(true),
			MaxSize:  cnv.StrPtr("1000"),
			MaxCount: cnv.StrPtr("222"),
		},
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestPubSub_DeleteChannel(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v2/channels/id" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "DELETE" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{MessagingUrl: ts.URL}})
	_, err := client.PubSub().DeleteChannel(ciosctx.Background(), "id")
	if err != nil {
		t.Fatal(err.Error())
	}
}
func TestPubSub_UpdateChannel(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v2/channels/id" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "PATCH" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{MessagingUrl: ts.URL}})
	_, _, err := client.PubSub().UpdateChannel(ciosctx.Background(), "id", cios.ChannelUpdateProposal{
		ChannelProtocols: nil,
		DisplayInfo:      nil,
		Labels:           nil,
		MessagingConfig:  nil,
		DatastoreConfig:  nil,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}
