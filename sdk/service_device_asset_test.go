package ciossdk

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	srvdevice "github.com/optim-corp/cios-golang-sdk/sdk/service/device"

	xmath "github.com/fcfcqloow/go-advance/math"

	cnv "github.com/fcfcqloow/go-advance/convert"

	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

func TestDeviceAssetManagement_GetModels(t *testing.T) {
	var (
		query url.Values
		ctx   context.Context

		tests = []struct {
			params cios.ApiGetDeviceModelsRequest
			test   func()
		}{
			{
				params: srvdevice.MakeGetModelsOpts().Limit(1000),
				test: func() {
					if query.Get("limit") != "1000" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvdevice.MakeGetModelsOpts().Limit(1000).Offset(50),
				test: func() {
					if query.Get("limit") != "1000" || query.Get("offset") != "50" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvdevice.MakeGetModelsOpts().ResourceOwnerId("aaaaa"),
				test: func() {
					if query.Get("resource_owner_id") != "aaaaa" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvdevice.MakeGetModelsOpts().Limit(1000).Offset(50).ResourceOwnerId("aaaaa").Name("name"),
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
				params: srvdevice.MakeGetModelsOpts().OrderBy("created_at"),
				test: func() {
					if query.Get("order_by") != "created_at" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvdevice.MakeGetModelsOpts().OrderBy("").Order("").ResourceOwnerId("").Name(""),
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
			Urls: sdkmodel.CIOSUrl{DeviceAssetManagementUrl: ts.URL},
		},
	)
	defer ts.Close()
	for _, test := range tests {
		client.DeviceAssetManagement.GetModels(ctx, test.params)
		test.test()
	}

	ts.Close()

}

func TestDeviceAssetManagement_GetModelsAll(t *testing.T) {
	var offsets []int
	var limits []int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleDeviceModel{Total: 3500, Models: []cios.DeviceModel{}}
		offset := cnv.MustInt(r.URL.Query().Get("offset"))
		limit := cnv.MustInt(r.URL.Query().Get("limit"))
		offsets = append(offsets, offset)
		limits = append(limits, limit)
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Models = append(response.Models, cios.DeviceModel{Id: cnv.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceAssetManagementUrl: ts.URL}})

	responses, _, _ := client.DeviceAssetManagement.GetModelsAll(nil, srvdevice.MakeGetModelsOpts().Limit(999))
	if len(responses) != 999 || offsets[0] != 0 && limits[0] != 1000 {
		t.Fatal(len(responses))
	}

	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceAssetManagement.GetModelsAll(nil, srvdevice.MakeGetModelsOpts().Limit(1500))
	if len(responses) != 1500 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 {
		t.Fatal(len(responses), limits, offsets)
	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceAssetManagement.GetModelsAll(nil, srvdevice.MakeGetModelsOpts().Limit(2001))
	if len(responses) != 2001 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 || offsets[2] != 2000 || limits[2] != 1 {
		t.Fatal(len(responses), limits, offsets)

	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceAssetManagement.GetModelsAll(nil, srvdevice.MakeGetModelsOpts().Limit(3501))
	if len(responses) != 3500 ||
		offsets[0] != 0 || limits[0] != 1000 ||
		offsets[1] != 1000 && limits[1] != 1000 ||
		offsets[2] != 2000 || limits[2] != 1000 ||
		offsets[3] != 3000 || limits[3] != 501 {
		t.Fatal(len(responses), limits, offsets)
	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceAssetManagement.GetModelsAll(nil, srvdevice.MakeGetModelsOpts().Limit(2001).Offset(20))
	if len(responses) != 2001 || offsets[0] != 20 && limits[0] != 1000 || offsets[1] != 1020 && limits[1] != 1000 || offsets[2] != 2020 || limits[2] != 1 {
		t.Fatal(len(responses), limits, offsets)

	}
}

func TestDeviceAssetManagement_GetModelsUnlimited(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleDeviceModel{Total: 3500, Models: []cios.DeviceModel{}}
		offset := cnv.MustInt(r.URL.Query().Get("offset"))
		limit := cnv.MustInt(r.URL.Query().Get("limit"))
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Models = append(response.Models, cios.DeviceModel{Id: cnv.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceAssetManagementUrl: ts.URL}})

	responses, _, _ := client.DeviceAssetManagement.GetModelsUnlimited(nil, srvdevice.MakeGetModelsOpts().Limit(1))
	if len(responses) != 3500 {
		t.Fatal(len(responses))
	}
}

func TestDeviceAssetManagement_GetModel(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/v2/device_models/test" {
			response := cios.SingleDeviceModel{Model: cios.DeviceModel{
				Id:              "test",
				ResourceOwnerId: "test_resource_owner",
			}}
			json.NewEncoder(w).Encode(response)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceAssetManagementUrl: ts.URL}})
	body, response, err := client.DeviceAssetManagement.GetModel(nil, "test")
	if body.Id != "test" || err != nil || response.StatusCode != 200 {
		t.Fatal(body)
	}
}

func TestDeviceAssetManagement_CreateModel(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/device_models" {
			t.Fatal(r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		body := cios.DeviceModelRequest{}
		if r.Method != "POST" {
			t.Fatal(r.Method)
		}
		byts, _ := ioutil.ReadAll(r.Body)
		cnv.UnMarshalJson(byts, &body)
		if body.Name != "name" || body.ResourceOwnerId != "resource_owner_id" {
			t.Fatal(body)
		}

	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceAssetManagementUrl: ts.URL}})
	_, _, err := client.DeviceAssetManagement.CreateModel(nil, cios.DeviceModelRequest{
		Name:            "name",
		ResourceOwnerId: "resource_owner_id",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestDeviceAssetManagement_DeleteModel(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v2/device_models/device_id" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "DELETE" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceAssetManagementUrl: ts.URL}})
	_, err := client.DeviceAssetManagement.DeleteModel(nil, "device_id")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestDeviceAssetManagement_GetEntities(t *testing.T) {
	var (
		query url.Values
		ctx   context.Context

		tests = []struct {
			params cios.ApiGetDeviceModelsRequest
			test   func()
		}{
			{
				params: srvdevice.MakeGetModelsOpts().Limit(1000),
				test: func() {
					if query.Get("limit") != "1000" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvdevice.MakeGetModelsOpts().Limit(1000).Offset(50),
				test: func() {
					if query.Get("limit") != "1000" || query.Get("offset") != "50" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvdevice.MakeGetModelsOpts().ResourceOwnerId("aaaaa"),
				test: func() {
					if query.Get("resource_owner_id") != "aaaaa" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvdevice.MakeGetModelsOpts().Limit(1000).Offset(50).ResourceOwnerId("aaaaa").Name("name"),
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
				params: srvdevice.MakeGetModelsOpts().OrderBy("created_at"),
				test: func() {
					if query.Get("order_by") != "created_at" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvdevice.MakeGetModelsOpts().OrderBy("").Order("").ResourceOwnerId("").Name(""),
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
			Urls: sdkmodel.CIOSUrl{DeviceAssetManagementUrl: ts.URL},
		},
	)
	defer ts.Close()
	for _, test := range tests {
		client.DeviceAssetManagement.GetModels(ctx, test.params)
		test.test()
	}

	ts.Close()

}

func TestDeviceAssetManagement_GetEntitiesAll(t *testing.T) {
	var offsets []int
	var limits []int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleDeviceModelEntity{Total: 3500, Entities: []cios.DeviceModelsEntity{}}
		offset := cnv.MustInt(r.URL.Query().Get("offset"))
		limit := cnv.MustInt(r.URL.Query().Get("limit"))
		offsets = append(offsets, offset)
		limits = append(limits, limit)
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Entities = append(response.Entities, cios.DeviceModelsEntity{Id: cnv.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceAssetManagementUrl: ts.URL}})

	responses, _, _ := client.DeviceAssetManagement.GetEntitiesAll(nil, srvdevice.MakeGetEntitiesOpts().Limit(999))
	if len(responses) != 999 || offsets[0] != 0 && limits[0] != 1000 {
		t.Fatal(len(responses))
	}

	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceAssetManagement.GetEntitiesAll(nil, srvdevice.MakeGetEntitiesOpts().Limit(1500))
	if len(responses) != 1500 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 {
		t.Fatal(len(responses), limits, offsets)
	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceAssetManagement.GetEntitiesAll(nil, srvdevice.MakeGetEntitiesOpts().Limit(2001))
	if len(responses) != 2001 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 || offsets[2] != 2000 || limits[2] != 1 {
		t.Fatal(len(responses), limits, offsets)

	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceAssetManagement.GetEntitiesAll(nil, srvdevice.MakeGetEntitiesOpts().Limit(3501))
	if len(responses) != 3500 ||
		offsets[0] != 0 || limits[0] != 1000 ||
		offsets[1] != 1000 && limits[1] != 1000 ||
		offsets[2] != 2000 || limits[2] != 1000 ||
		offsets[3] != 3000 || limits[3] != 501 {
		t.Fatal(len(responses), limits, offsets)
	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceAssetManagement.GetEntitiesAll(nil, srvdevice.MakeGetEntitiesOpts().Limit(2001).Offset(20))
	if len(responses) != 2001 || offsets[0] != 20 && limits[0] != 1000 || offsets[1] != 1020 && limits[1] != 1000 || offsets[2] != 2020 || limits[2] != 1 {
		t.Fatal(len(responses), limits, offsets)

	}
}

func TestDeviceAssetManagement_GetGetEntitiesUnlimited(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleDeviceModelEntity{Total: 3500, Entities: []cios.DeviceModelsEntity{}}
		offset := cnv.MustInt(r.URL.Query().Get("offset"))
		limit := cnv.MustInt(r.URL.Query().Get("limit"))
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Entities = append(response.Entities, cios.DeviceModelsEntity{Id: cnv.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceAssetManagementUrl: ts.URL}})

	responses, _, err := client.DeviceAssetManagement.GetEntitiesUnlimited(nil, srvdevice.MakeGetEntitiesOpts().Limit(1))
	if err != nil {
		t.Log(err.Error())
	}
	if len(responses) != 3500 {
		t.Fatal(len(responses))
	}
}

func TestDeviceAssetManagement_GetEntity(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/v2/device_entities/test" {
			response := cios.SingleDeviceModelsEntity{Entity: cios.DeviceModelsEntity{
				Id:              "test",
				ResourceOwnerId: "test_resource_owner",
			}}
			json.NewEncoder(w).Encode(response)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceAssetManagementUrl: ts.URL}})
	body, response, err := client.DeviceAssetManagement.GetEntity(nil, "test")
	if body.Id != "test" || err != nil || response.StatusCode != 200 {
		t.Fatal(body)
	}
}

func TestDeviceAssetManagement_CreateEntity(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/device_models/name/entities" {
			t.Fatal(r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		body := cios.Inventory{}
		if r.Method != "POST" {
			t.Fatal(r.Method)
		}
		cnv.UnMarshalJson(r.Body, &body)
		if *body.SerialNumber != "111" {
			t.Fatal(body)
		}

	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceAssetManagementUrl: ts.URL}})
	_, _, err := client.DeviceAssetManagement.CreateEntity(nil, "name", cios.Inventory{
		SerialNumber: cnv.StrPtr("111"),
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestDeviceAssetManagement_DeleteEntity(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v2/device_entities/device_id" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "DELETE" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceAssetManagementUrl: ts.URL}})
	_, err := client.DeviceAssetManagement.DeleteEntity(nil, "device_id")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestDeviceAssetManagement_GetLifecycles(t *testing.T) {
	var (
		query url.Values
		ctx   context.Context

		tests = []struct {
			params cios.ApiGetDeviceEntitiesLifecyclesRequest
			test   func()
		}{
			{
				params: srvdevice.MakeGetLifecyclesOpts().Limit(1000),
				test: func() {
					if query.Get("limit") != "1000" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvdevice.MakeGetLifecyclesOpts().Limit(1000).Offset(50),
				test: func() {
					if query.Get("limit") != "1000" || query.Get("offset") != "50" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvdevice.MakeGetLifecyclesOpts().ComponentId("aaaaa"),
				test: func() {
					if query.Get("component_id") != "aaaaa" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvdevice.MakeGetLifecyclesOpts().Limit(1000).Offset(50).EndEventAt("aaaaa").StartEventAt(":190219021"),
				test: func() {
					if query.Get("end_event_at") != "aaaaa" ||
						query.Get("start_event_at") != ":190219021" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvdevice.MakeGetLifecyclesOpts().OrderBy("created_at"),
				test: func() {
					if query.Get("order_by") != "created_at" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvdevice.MakeGetLifecyclesOpts().OrderBy("").Order("").AfterId("").
					EndEventAt("").StartEventAt("").BeforeId("").
					EventType("").EventMode("").ComponentId(""),
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
			Urls: sdkmodel.CIOSUrl{DeviceAssetManagementUrl: ts.URL},
		},
	)
	defer ts.Close()
	for _, test := range tests {
		client.DeviceAssetManagement.GetLifecycles(ctx, "key", test.params)
		test.test()
	}

	ts.Close()

}

func TestDeviceAssetManagement_GetLifecyclesAll(t *testing.T) {
	var offsets []int
	var limits []int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v2/device_entities/key/lifecycles" {
			t.Fatal(r.URL.Path)
		}
		response := cios.MultipleLifeCycle{Total: 3500, Lifecycles: []cios.LifeCycle{}}
		offset := cnv.MustInt(r.URL.Query().Get("offset"))
		limit := cnv.MustInt(r.URL.Query().Get("limit"))
		offsets = append(offsets, offset)
		limits = append(limits, limit)
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Lifecycles = append(response.Lifecycles, cios.LifeCycle{Id: cnv.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceAssetManagementUrl: ts.URL}})

	responses, _, _ := client.DeviceAssetManagement.GetLifecyclesAll(nil, "key", srvdevice.MakeGetLifecyclesOpts().Limit(999))
	if len(responses) != 999 || offsets[0] != 0 && limits[0] != 1000 {
		t.Fatal(len(responses))
	}

	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceAssetManagement.GetLifecyclesAll(nil, "key", srvdevice.MakeGetLifecyclesOpts().Limit(1500))
	if len(responses) != 1500 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 {
		t.Fatal(len(responses), limits, offsets)
	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceAssetManagement.GetLifecyclesAll(nil, "key", srvdevice.MakeGetLifecyclesOpts().Limit(2001))
	if len(responses) != 2001 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 || offsets[2] != 2000 || limits[2] != 1 {
		t.Fatal(len(responses), limits, offsets)

	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceAssetManagement.GetLifecyclesAll(nil, "key", srvdevice.MakeGetLifecyclesOpts().Limit(3501))
	if len(responses) != 3500 ||
		offsets[0] != 0 || limits[0] != 1000 ||
		offsets[1] != 1000 && limits[1] != 1000 ||
		offsets[2] != 2000 || limits[2] != 1000 ||
		offsets[3] != 3000 || limits[3] != 501 {
		t.Fatal(len(responses), limits, offsets)
	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.DeviceAssetManagement.GetLifecyclesAll(nil, "key", srvdevice.MakeGetLifecyclesOpts().Limit(2001).Offset(20))
	if len(responses) != 2001 || offsets[0] != 20 && limits[0] != 1000 || offsets[1] != 1020 && limits[1] != 1000 || offsets[2] != 2020 || limits[2] != 1 {
		t.Fatal(len(responses), limits, offsets)

	}
}

func TestDeviceAssetManagement_GetLifecyclesUnlimited(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v2/device_entities/key/lifecycles" {
			t.Fatal(r.URL.Path)
		}
		response := cios.MultipleLifeCycle{Total: 3500, Lifecycles: []cios.LifeCycle{}}
		offset := cnv.MustInt(r.URL.Query().Get("offset"))
		limit := cnv.MustInt(r.URL.Query().Get("limit"))
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Lifecycles = append(response.Lifecycles, cios.LifeCycle{Id: cnv.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceAssetManagementUrl: ts.URL}})

	responses, _, _ := client.DeviceAssetManagement.GetLifecyclesUnlimited(nil, "key", srvdevice.MakeGetLifecyclesOpts().Limit(1))
	if len(responses) != 3500 {
		t.Fatal(len(responses))
	}
}

func TestDeviceAssetManagement_CreateLifecycle(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body := cios.LifeCycleRequest{}
		byts, _ := ioutil.ReadAll(r.Body)
		cnv.UnMarshalJson(byts, &body)
		if body.EventKind != "a" || body.EventMode != "b" || body.EventType != "c" || body.EventAt != "d" {
			t.Fatal(body)
		}
		if r.URL.Path != "/v2/device_entities/device_id/lifecycles" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "POST" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceAssetManagementUrl: ts.URL}})
	_, _, err := client.DeviceAssetManagement.CreateLifecycle(nil, "device_id", cios.LifeCycleRequest{
		EventKind: "a",
		EventMode: "b",
		EventType: "c",
		EventAt:   "d",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestDeviceAssetManagement_DeleteLifecycle(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v2/device_entities/device_id/lifecycles/test_id" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "DELETE" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceAssetManagementUrl: ts.URL}})
	_, err := client.DeviceAssetManagement.DeleteLifecycle(nil, "device_id", "test_id")
	if err != nil {
		t.Fatal(err.Error())
	}
}
