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

	"github.com/optim-corp/cios-golang-sdk/model"
)

func TestDeviceManagement_GetPolicies(t *testing.T) {
	var (
		query url.Values
		ctx   context.Context

		tests = []struct {
			params cios.ApiGetDevicesRequest
			test   func()
		}{
			{
				params: MakeGetDevicesOpts().Limit(1000),
				test: func() {
					if query.Get("limit") != "1000" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: MakeGetDevicesOpts().Limit(1000).Offset(50),
				test: func() {
					if query.Get("limit") != "1000" || query.Get("offset") != "50" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: MakeGetDevicesOpts().ResourceOwnerId("aaaaa"),
				test: func() {
					if query.Get("resource_owner_id") != "aaaaa" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: MakeGetDevicesOpts().Limit(1000).Offset(50).ResourceOwnerId("aaaaa").Name("name"),
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
				params: MakeGetDevicesOpts().OrderBy("created_at"),
				test: func() {
					if query.Get("order_by") != "created_at" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: MakeGetDevicesOpts().OrderBy("").Order("").ResourceOwnerId("").Name(""),
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
	bucketHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query = r.URL.Query()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cios.MultipleDevice{Total: 10})
	})
	ts := httptest.NewServer(bucketHandler)
	client := NewCiosClient(
		CiosClientConfig{
			Urls: model.CIOSUrl{DeviceManagementUrl: ts.URL},
		},
	)
	defer ts.Close()
	for _, test := range tests {
		client.DeviceManagement.GetDevices(test.params, ctx)
		test.test()
	}

	ts.Close()

	//// Auto Refresh Test
	//client = NewCiosClient(
	//	CiosClientConfig{
	//		Urls:        model.CIOSUrl{DeviceManagementUrl: ts.URL},
	//		AutoRefresh: true,
	//	},
	//)
	//bucketHandler = func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(404)
	//}
	//ts = httptest.NewServer(bucketHandler)
	//
	//result := "Failed"
	//refFunc := func() (model.AccessToken, model.Scope, model.TokenType, model.ExpiresIn, error) {
	//	result = "Accept"
	//	return "", "", "", 0, nil
	//}
	//client.DeviceManagement.refresh = &refFunc
	//if result == "Failed" {
	//	t.Fatal("Cant Refresh", result)
	//}
	////　念のためクローズ
	//ts.Close()
}

func TestDeviceManagement_GetPoliciesAll(t *testing.T) {
	var offsets []int
	var limits []int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleDevice{Total: 3500, Devices: []cios.Device{}}
		offset := convert.MustInt(r.URL.Query().Get("offset"))
		limit := convert.MustInt(r.URL.Query().Get("limit"))
		offsets = append(offsets, offset)
		limits = append(limits, limit)
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Devices = append(response.Devices, cios.Device{Id: convert.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: model.CIOSUrl{DeviceManagementUrl: ts.URL}})

	buckets, _, _ := client.DeviceManagement.GetDevicesAll(MakeGetDevicesOpts().Limit(999), context.Background())
	if len(buckets) != 999 || offsets[0] != 0 && limits[0] != 1000 {
		t.Fatal(len(buckets))
	}

	offsets = []int{}
	limits = []int{}
	buckets, _, _ = client.DeviceManagement.GetDevicesAll(MakeGetDevicesOpts().Limit(1500), context.Background())
	if len(buckets) != 1500 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 {
		t.Fatal(len(buckets), limits, offsets)
	}
	offsets = []int{}
	limits = []int{}
	buckets, _, _ = client.DeviceManagement.GetDevicesAll(MakeGetDevicesOpts().Limit(2001), context.Background())
	if len(buckets) != 2001 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 || offsets[2] != 2000 || limits[2] != 1 {
		t.Fatal(len(buckets), limits, offsets)

	}
	offsets = []int{}
	limits = []int{}
	buckets, _, _ = client.DeviceManagement.GetDevicesAll(MakeGetDevicesOpts().Limit(3501), context.Background())
	if len(buckets) != 3500 ||
		offsets[0] != 0 || limits[0] != 1000 ||
		offsets[1] != 1000 && limits[1] != 1000 ||
		offsets[2] != 2000 || limits[2] != 1000 ||
		offsets[3] != 3000 || limits[3] != 501 {
		t.Fatal(len(buckets), limits, offsets)
	}
}

func TestDeviceManagement_GetPoliciesUnlimited(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleDevice{Total: 3500, Devices: []cios.Device{}}
		offset := convert.MustInt(r.URL.Query().Get("offset"))
		limit := convert.MustInt(r.URL.Query().Get("limit"))
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Devices = append(response.Devices, cios.Device{Id: convert.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: model.CIOSUrl{DeviceManagementUrl: ts.URL}})

	buckets, _, _ := client.DeviceManagement.GetDevicesUnlimited(MakeGetDevicesOpts().Limit(1), context.Background())
	if len(buckets) != 3500 {
		t.Fatal(len(buckets))
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
	client := NewCiosClient(CiosClientConfig{Urls: model.CIOSUrl{DeviceManagementUrl: ts.URL}})
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
		if r.URL.Path != "/v2/devices/bucketid" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "PATCH" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: model.CIOSUrl{DeviceManagementUrl: ts.URL}})
	client.DeviceManagement.CreatePolicy("resource_owner_id", context.Background())
}