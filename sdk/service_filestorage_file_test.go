package ciossdk

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	srvfilestorage "github.com/optim-corp/cios-golang-sdk/sdk/service/filestorage"

	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

func TestFileStorage_DownloadFile(t *testing.T) {
	requested := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requested = true
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v2/file_storage/buckets/bucketid/download" || r.URL.Query().Get("node_id") != "node" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "GET" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{StorageUrl: ts.URL}})
	client.FileStorage().DownloadFile(context.Background(), "bucketid", "node")
	if !requested {
		t.Fatal(ts.URL)
	}
}

func TestFileStorage_DownloadFileByKey(t *testing.T) {
	requested := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requested = true
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v2/file_storage/buckets/bucketid/download" || r.URL.Query().Get("key") != "node" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "GET" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{StorageUrl: ts.URL}})
	client.FileStorage().DownloadFileByKey(context.Background(), "bucketid", "node")
	if !requested {
		t.Fatal(ts.URL)
	}
}
func TestFileStorage_UploadFile(t *testing.T) {
	key, nodeId := "", ""
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		key = r.URL.Query().Get("key")
		nodeId = r.URL.Query().Get("node_id")
		if r.URL.Query().Get("checksum") != "NsRTaZbKVhXc+ZEfBoeG3A==" {
			t.Fatal(r.URL.Query().Get("checksum"))
		}
		if r.URL.Path != "/v2/file_storage/buckets/bucketid/upload" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "PUT" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{StorageUrl: ts.URL}})
	client.FileStorage().UploadFile(context.Background(), "bucketid", []byte("node"),
		srvfilestorage.MakeUploadFileOpts().Key("").NodeId(""))
	if nodeId != "" || key != "" {
		t.Fatal(nodeId, key)
	}
	client.FileStorage().UploadFile(context.Background(), "bucketid", []byte("node"),
		srvfilestorage.MakeUploadFileOpts().Key("test").NodeId("test"))
	if nodeId != "test" || key != "test" {
		t.Fatal(nodeId, key)
	}
}
