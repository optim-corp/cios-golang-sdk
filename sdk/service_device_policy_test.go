package ciossdk

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	xmath "github.com/optim-kazuhiro-seida/go-advance-type/math"

	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	"github.com/optim-corp/cios-golang-sdk/cios"

	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

func TestDeviceManagement_GetPolicies(t *testing.T) {
	var (
		query url.Values
		ctx   context.Context

		tests = []struct {
			params cios.ApiGetDevicePoliciesRequest
			test   func()
		}{
			{
				params: MakeGetPoliciesOpts().Limit(1000),
				test: func() {
					if query.Get("limit") != "1000" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: MakeGetPoliciesOpts().Limit(1000).Offset(50),
				test: func() {
					if query.Get("limit") != "1000" || query.Get("offset") != "50" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: MakeGetPoliciesOpts().ResourceOwnerId("aaaaa"),
				test: func() {
					if query.Get("resource_owner_id") != "aaaaa" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: MakeGetPoliciesOpts().Limit(1000).Offset(50).ResourceOwnerId("aaaaa"),
				test: func() {
					if query.Get("resource_owner_id") != "aaaaa" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: MakeGetPoliciesOpts().OrderBy("created_at"),
				test: func() {
					if query.Get("order_by") != "created_at" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: MakeGetPoliciesOpts().OrderBy("").Order("").ResourceOwnerId(""),
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
		json.NewEncoder(w).Encode(cios.MultipleDevice{Total: 10})
	})
	ts := httptest.NewServer(responseHandler)
	client := NewCiosClient(
		CiosClientConfig{
			Urls: sdkmodel.CIOSUrl{DeviceManagementUrl: ts.URL},
		},
	)
	defer ts.Close()
	for _, test := range tests {
		client.DeviceManagement.GetPolicies(test.params, ctx)
		test.test()
	}

	ts.Close()

}

func TestDeviceManagement_GetPoliciesAll(t *testing.T) {
	var offsets []int
	var limits []int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleDevicePolicy{Total: 3500}
		offset := convert.MustInt(r.URL.Query().Get("offset"))
		limit := convert.MustInt(r.URL.Query().Get("limit"))
		offsets = append(offsets, offset)
		limits = append(limits, limit)
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Policies = append(response.Policies, cios.DevicePolicy{ResourceOwnerId: convert.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceManagementUrl: ts.URL}})

	responses, _, _ := client.DeviceManagement.GetPoliciesAll(MakeGetPoliciesOpts().Limit(999), context.Background())
	if len(responses) != 999 || offsets[0] != 0 && limits[0] != 1000 {
		t.Fatal(len(responses))
	}

	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceManagement.GetPoliciesAll(MakeGetPoliciesOpts().Limit(1500), context.Background())
	if len(responses) != 1500 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 {
		t.Fatal(len(responses), limits, offsets)
	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceManagement.GetPoliciesAll(MakeGetPoliciesOpts().Limit(2001), context.Background())
	if len(responses) != 2001 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 || offsets[2] != 2000 || limits[2] != 1 {
		t.Fatal(len(responses), limits, offsets)

	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceManagement.GetPoliciesAll(MakeGetPoliciesOpts().Limit(3501), context.Background())
	if len(responses) != 3500 ||
		offsets[0] != 0 || limits[0] != 1000 ||
		offsets[1] != 1000 && limits[1] != 1000 ||
		offsets[2] != 2000 || limits[2] != 1000 ||
		offsets[3] != 3000 || limits[3] != 501 {
		t.Fatal(len(responses), limits, offsets)
	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceManagement.GetPoliciesAll(MakeGetPoliciesOpts().Limit(2001).Offset(20), context.Background())
	if len(responses) != 2001 || offsets[0] != 20 && limits[0] != 1000 || offsets[1] != 1020 && limits[1] != 1000 || offsets[2] != 2020 || limits[2] != 1 {
		t.Fatal(len(responses), limits, offsets)

	}
}

func TestDeviceManagement_GetPoliciesUnlimited(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleDevicePolicy{Total: 3500, Policies: []cios.DevicePolicy{}}
		offset := convert.MustInt(r.URL.Query().Get("offset"))
		limit := convert.MustInt(r.URL.Query().Get("limit"))
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Policies = append(response.Policies, cios.DevicePolicy{ResourceOwnerId: convert.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceManagementUrl: ts.URL}})

	responses, _, _ := client.DeviceManagement.GetPoliciesUnlimited(MakeGetPoliciesOpts().Limit(1), context.Background())
	if len(responses) != 3500 {
		t.Fatal(len(responses))
	}
}

func TestDeviceManagement_DeletePolicy(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/devices/group_policies/id" {
			t.Fatal(r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		if r.Method != "DELETE" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceManagementUrl: ts.URL}})
	client.DeviceManagement.DeletePolicy("id", context.Background())
}

func TestDeviceManagement_CreatePolicy(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body := cios.DevicePolicyRequest{}
		byts, _ := ioutil.ReadAll(r.Body)
		convert.UnMarshalJson(byts, &body)
		if body.ResourceOwnerId != "resource_owner_id" {
			t.Fatal(body)
		}
		if r.URL.Path != "/v2/devices/group_policies" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "POST" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceManagementUrl: ts.URL}})
	client.DeviceManagement.CreatePolicy("resource_owner_id", context.Background())
}
