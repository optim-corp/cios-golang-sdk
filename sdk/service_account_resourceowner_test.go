package ciossdk

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	srvaccount "github.com/optim-corp/cios-golang-sdk/sdk/service/account"

	cnv "github.com/fcfcqloow/go-advance/convert"
	xmath "github.com/fcfcqloow/go-advance/math"
	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

func TestResourceOwner_GetResourceOwners(t *testing.T) {
	var (
		query url.Values
		ctx   context.Context

		tests = []struct {
			params cios.ApiGetResourceOwnersRequest
			test   func()
		}{
			{
				params: srvaccount.MakeGetResourceOwnersOpts().Limit(1000),
				test: func() {
					if query.Get("limit") != "1000" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvaccount.MakeGetResourceOwnersOpts().Limit(1000).Offset(50),
				test: func() {
					if query.Get("limit") != "1000" || query.Get("offset") != "50" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvaccount.MakeGetResourceOwnersOpts().Page("aaaaa"),
				test: func() {
					if query.Get("page") != "aaaaa" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvaccount.MakeGetResourceOwnersOpts().GroupId("aaaaa"),
				test: func() {
					if query.Get("group_id") != "aaaaa" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvaccount.MakeGetResourceOwnersOpts().Page(""),
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
		json.NewEncoder(w).Encode(cios.MultipleResourceOwner{Total: 10})
	})
	ts := httptest.NewServer(responseHandler)
	client := NewCiosClient(
		CiosClientConfig{
			Urls: sdkmodel.CIOSUrl{AccountsUrl: ts.URL},
		},
	)
	defer ts.Close()
	for _, test := range tests {
		client.Account().GetResourceOwners(ctx, test.params)
		test.test()
	}
}

func TestResourceOwner_GetResourceOwnersAll(t *testing.T) {
	var offsets []int
	var limits []int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleResourceOwner{Total: 3500, ResourceOwners: []cios.ResourceOwner{}}
		offset := cnv.MustInt(r.URL.Query().Get("offset"))
		limit := cnv.MustInt(r.URL.Query().Get("limit"))
		offsets = append(offsets, offset)
		limits = append(limits, limit)
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.ResourceOwners = append(response.ResourceOwners, cios.ResourceOwner{Id: cnv.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{AccountsUrl: ts.URL}})

	responses, _, _ := client.Account().GetResourceOwnersAll(nil, srvaccount.MakeGetResourceOwnersOpts().Limit(999))
	if len(responses) != 999 || offsets[0] != 0 && limits[0] != 1000 {
		t.Fatal(len(responses))
	}

	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.Account().GetResourceOwnersAll(nil, srvaccount.MakeGetResourceOwnersOpts().Limit(1500))
	if len(responses) != 1500 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 {
		t.Fatal(len(responses), limits, offsets)
	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.Account().GetResourceOwnersAll(nil, srvaccount.MakeGetResourceOwnersOpts().Limit(2001))
	if len(responses) != 2001 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 || offsets[2] != 2000 || limits[2] != 1 {
		t.Fatal(len(responses), limits, offsets)

	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.Account().GetResourceOwnersAll(nil, srvaccount.MakeGetResourceOwnersOpts().Limit(3501))
	if len(responses) != 3500 ||
		offsets[0] != 0 || limits[0] != 1000 ||
		offsets[1] != 1000 && limits[1] != 1000 ||
		offsets[2] != 2000 || limits[2] != 1000 ||
		offsets[3] != 3000 || limits[3] != 501 {
		t.Fatal(len(responses), limits, offsets)
	}

	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.Account().GetResourceOwnersAll(nil, srvaccount.MakeGetResourceOwnersOpts().Limit(2001).Offset(20))
	if len(responses) != 2001 || offsets[0] != 20 && limits[0] != 1000 || offsets[1] != 1020 && limits[1] != 1000 || offsets[2] != 2020 || limits[2] != 1 {
		t.Fatal(len(responses), limits, offsets)

	}
}

func TestResourceOwner_GetResourceOwnersUnlimited(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleResourceOwner{Total: 3500, ResourceOwners: []cios.ResourceOwner{}}
		offset := cnv.MustInt(r.URL.Query().Get("offset"))
		limit := cnv.MustInt(r.URL.Query().Get("limit"))
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.ResourceOwners = append(response.ResourceOwners, cios.ResourceOwner{Id: cnv.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{AccountsUrl: ts.URL}})

	responses, _, _ := client.Account().GetResourceOwnersUnlimited(nil, srvaccount.MakeGetResourceOwnersOpts().Limit(1))
	if len(responses) != 3500 {
		t.Fatal(len(responses))
	}
}
func TestPubSub_GetResourceOwner(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != "GET" {
			t.Fatal(r.Method)
		}
		if r.URL.Path == "/v2/resource_owners/test" {
			response := cios.ResourceOwner{
				Id:       "test",
				GroupId:  nil,
				UserId:   nil,
				AuthorId: nil,
				Profile:  cios.ResourceOwnerProfile{},
				Type:     "",
			}
			json.NewEncoder(w).Encode(response)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{AccountsUrl: ts.URL}})
	responseB, response, err := client.Account().GetResourceOwner(nil, "test")
	if responseB.Id != "test" || err != nil || response.StatusCode != 200 {
		t.Fatal(responseB)
	}
}
