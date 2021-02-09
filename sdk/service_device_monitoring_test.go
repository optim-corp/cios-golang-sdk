package ciossdk

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/optim-corp/cios-golang-sdk/cios"

	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

func TestDeviceManagement_GetMonitoringLatestList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var body cios.DeviceMonitoringIDsRequest
		byts, _ := ioutil.ReadAll(r.Body)
		convert.UnMarshalJson(byts, &body)
		if body.DeviceIds[0] != "1" || body.DeviceIds[1] != "2" || body.DeviceIds[2] != "3" {
			t.Fatal(body)
		}
		if r.URL.Path != "/v2/devices/monitoring/latest" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "POST" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceManagementUrl: ts.URL}})

	_, _, err := client.DeviceManagement.GetMonitoringLatestList([]string{"1", "2", "3"}, context.Background())
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestDeviceManagement_GetMonitoring(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v2/devices/id/monitoring/latest" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "GET" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{DeviceManagementUrl: ts.URL}})

	_, _, err := client.DeviceManagement.GetMonitoring("id", context.Background())
	if err != nil {
		t.Fatal(err.Error())
	}
}
