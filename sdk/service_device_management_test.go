package ciossdk

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	xmath "github.com/fcfcqloow/go-advance/math"

	cnv "github.com/fcfcqloow/go-advance/convert"

	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

func TestDeviceManagement_GetDevices(t *testing.T) {
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
		client.DeviceManagement.GetDevices(test.params, ctx)
		test.test()
	}

	ts.Close()

	//// Auto Refresh Test
	//client = NewCiosClient(
	//	CiosClientConfig{
	//		Urls:        sdkmodel.CIOSUrl{DeviceManagementUrl: ts.URL},
	//		AutoRefresh: true,
	//	},
	//)
	//responseHandler = func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(404)
	//}
	//ts = httptest.NewServer(responseHandler)
	//
	//result := "Failed"
	//refFunc := func() (sdkmodel.AccessToken, sdkmodel.Scope, sdkmodel.TokenType, sdkmodel.ExpiresIn, error) {
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

func TestDeviceManagement_GetDevicesAll(t *testing.T) {
	var offsets []int
	var limits []int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleDevice{Total: 3500, Devices: []cios.Device{}}
		offset := cnv.MustInt(r.URL.Query().Get("offset"))
		limit := cnv.MustInt(r.URL.Query().Get("limit"))
		offsets = append(offsets, offset)
		limits = append(limits, limit)
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Devices = append(response.Devices, cios.Device{Id: cnv.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceManagementUrl: ts.URL}})

	responses, _, _ := client.DeviceManagement.GetDevicesAll(MakeGetDevicesOpts().Limit(999), context.Background())
	if len(responses) != 999 || offsets[0] != 0 && limits[0] != 1000 {
		t.Fatal(len(responses))
	}

	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceManagement.GetDevicesAll(MakeGetDevicesOpts().Limit(1500), context.Background())
	if len(responses) != 1500 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 {
		t.Fatal(len(responses), limits, offsets)
	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceManagement.GetDevicesAll(MakeGetDevicesOpts().Limit(2001), context.Background())
	if len(responses) != 2001 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 || offsets[2] != 2000 || limits[2] != 1 {
		t.Fatal(len(responses), limits, offsets)

	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceManagement.GetDevicesAll(MakeGetDevicesOpts().Limit(3501), context.Background())
	if len(responses) != 3500 ||
		offsets[0] != 0 || limits[0] != 1000 ||
		offsets[1] != 1000 && limits[1] != 1000 ||
		offsets[2] != 2000 || limits[2] != 1000 ||
		offsets[3] != 3000 || limits[3] != 501 {
		t.Fatal(len(responses), limits, offsets)
	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceManagement.GetDevicesAll(MakeGetDevicesOpts().Limit(2001).Offset(20), context.Background())
	if len(responses) != 2001 || offsets[0] != 20 && limits[0] != 1000 || offsets[1] != 1020 && limits[1] != 1000 || offsets[2] != 2020 || limits[2] != 1 {
		t.Fatal(len(responses), limits, offsets)

	}
}

func TestDeviceManagement_GetDevicesUnlimited(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleDevice{Total: 3500, Devices: []cios.Device{}}
		offset := cnv.MustInt(r.URL.Query().Get("offset"))
		limit := cnv.MustInt(r.URL.Query().Get("limit"))
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Devices = append(response.Devices, cios.Device{Id: cnv.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceManagementUrl: ts.URL}})

	responses, _, _ := client.DeviceManagement.GetDevicesUnlimited(MakeGetDevicesOpts().Limit(1), context.Background())
	if len(responses) != 3500 {
		t.Fatal(len(responses))
	}
}

func TestDeviceManagement_GetDevice(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/v2/devices/test" {
			response := cios.SingleDevice{Device: cios.Device{
				Id:              "test",
				ResourceOwnerId: "test_resource_owner",
				Name:            cnv.StrPtr(""),
			}}
			json.NewEncoder(w).Encode(response)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceManagementUrl: ts.URL}})
	body, response, err := client.DeviceManagement.GetDevice("test", nil, nil, context.Background())
	if body.Id != "test" || err != nil || response.StatusCode != 200 {
		t.Fatal(body)
	}
}

func TestDeviceManagement_CreateDevice(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/devices" {
			t.Fatal(r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		body := cios.DeviceInfo{}
		if r.Method != "POST" {
			t.Fatal(r.Method)
		}
		byts, _ := ioutil.ReadAll(r.Body)
		cnv.UnMarshalJson(byts, &body)
		if *body.Name != "name" || body.ResourceOwnerId != "resource_owner_id" {
			t.Fatal(body)
		}

	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceManagementUrl: ts.URL}})
	_, _, err := client.DeviceManagement.CreateDevice(cios.DeviceInfo{
		Name:            cnv.StrPtr("name"),
		ResourceOwnerId: "resource_owner_id",
	}, context.Background())
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestDeviceManagement_DeleteDevice(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v2/devices/device_id" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "DELETE" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceManagementUrl: ts.URL}})
	_, err := client.DeviceManagement.DeleteDevice("device_id", context.Background())
	if err != nil {
		t.Fatal(err.Error())
	}
}
func TestDeviceManagement_UpdateDevice(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body := cios.DeviceUpdateRequest{}
		byts, _ := ioutil.ReadAll(r.Body)
		cnv.UnMarshalJson(byts, &body)
		if *body.Name != "name" || *body.Description != "desc" ||
			*body.IdNumber != "1234" ||
			*body.RsaPublickey != "rsa" {
			t.Fatal(body)
		}
		if r.URL.Path != "/v2/devices/responseid" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "PATCH" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceManagementUrl: ts.URL}})
	_, _, err := client.DeviceManagement.UpdateDevice("responseid", cios.DeviceUpdateRequest{
		Name:         cnv.StrPtr("name"),
		IdNumber:     cnv.StrPtr("1234"),
		Description:  cnv.StrPtr("desc"),
		RsaPublickey: cnv.StrPtr("rsa"),
	}, context.Background())
	if err != nil {
		t.Fatal(err.Error())
	}
}
