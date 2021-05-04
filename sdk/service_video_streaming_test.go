package ciossdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	cnv "github.com/fcfcqloow/go-advance/convert"

	"github.com/optim-corp/cios-golang-sdk/cios"

	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

func TestVideo_Videos(t *testing.T) {
	var (
		query url.Values
		tests = []struct {
			params cios.ApiGetVideoStreamsListRequest
			test   func()
		}{
			{
				params: MakeGetVideosOpts().
					ResourceOwnerId("test").
					DeviceId("test"),
				test: func() {
					if query.Encode() != "device_id=test&resource_owner_id=test" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: MakeGetVideosOpts().
					ResourceOwnerId("test").
					DeviceId(""),
				test: func() {
					if query.Encode() != "resource_owner_id=test" {
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
		json.NewEncoder(w).Encode(cios.MultipleVideo{Total: 10})
	})
	ts := httptest.NewServer(responseHandler)
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{VideoStreamingUrl: ts.URL}})
	defer ts.Close()
	for _, test := range tests {
		client.Video.GetVideoInfos(test.params, nil)
		test.test()
	}

	ts.Close()
}

func TestVideo_GetVideoInfo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/v2/video_streams/test" {
			json.NewEncoder(w).Encode(cios.SingleVideo{
				Video: cios.Video{
					Name:             nil,
					Id:               "test",
					DeviceId:         nil,
					ResourceOwnerId:  "",
					VideoName:        nil,
					VideoDescription: nil,
					Enabled:          nil,
					CreatedAt:        "",
					UpdatedAt:        "",
				},
			})
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{VideoStreamingUrl: ts.URL}})
	responseB, response, err := client.Video.GetVideoInfo("test", nil)
	if responseB.Id != "test" || err != nil || response.StatusCode != 200 {
		t.Fatal(responseB)
	}
}
func TestVideo_GetThumbnail(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/jpeg")
		if r.URL.Path == "/v2/video_streams/test/thumbnail" {
			if _, err := w.Write([]byte("test")); err != nil {
				t.Error(err.Error())
			}

		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{VideoStreamingUrl: ts.URL}})
	responseB, response, err := client.Video.GetThumbnail("test", nil)
	if string(responseB) != "test" || err != nil || response.StatusCode != 200 {
		t.Fatal(responseB)

	}
}
func TestVideo_UpdateVideoInfo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var body cios.VideoUpdateRequest
		if err := cnv.UnMarshalJson(r.Body, &body); err != nil {
			t.Fatal(err.Error())
		}
		if *body.VideoName != "name" {
			t.Fatal(body)
		}
		if *body.VideoDescription != "dp" {
			t.Fatal(body)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{VideoStreamingUrl: ts.URL}})
	_, _, err := client.Video.UpdateVideoInfo("test", "name", "dp", nil)
	if err != nil {
		t.Fatal(err.Error())
	}
}
func TestVideo_Play(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v2/video_streams/id/play" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "POST" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{VideoStreamingUrl: ts.URL}})
	_, _, err := client.Video.Play("id", nil)
	if err != nil {
		t.Fatal(err.Error())
	}
}
func TestVideo_Stop(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v2/video_streams/id/stop" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "POST" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{VideoStreamingUrl: ts.URL}})
	_, err := client.Video.Stop("id", nil)
	if err != nil {
		t.Fatal(err.Error())
	}
}
