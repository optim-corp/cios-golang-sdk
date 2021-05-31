package ciossdk

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	srvfilestorage "github.com/optim-corp/cios-golang-sdk/sdk/service/filestorage"

	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"

	xmath "github.com/fcfcqloow/go-advance/math"

	cnv "github.com/fcfcqloow/go-advance/convert"

	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

func TestFileStorage_GetNodes(t *testing.T) {
	var (
		query url.Values
		ctx   context.Context

		tests = []struct {
			params cios.ApiGetNodesRequest
			test   func()
		}{
			{
				params: srvfilestorage.MakeGetNodesOpts().Limit(1000),
				test: func() {
					if query.Get("limit") != "1000" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvfilestorage.MakeGetNodesOpts().Limit(1000).Offset(50),
				test: func() {
					if query.Get("limit") != "1000" || query.Get("offset") != "50" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvfilestorage.MakeGetNodesOpts().IsDirectory(true),
				test: func() {
					if query.Get("is_directory") != "true" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvfilestorage.MakeGetNodesOpts().Limit(1000).Offset(50).IsDirectory(true).Name("name"),
				test: func() {
					if query.Get("is_directory") != "true" ||
						query.Get("name") != "name" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvfilestorage.MakeGetNodesOpts().OrderBy("created_at"),
				test: func() {
					if query.Get("order_by") != "created_at" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvfilestorage.MakeGetNodesOpts().OrderBy("").Order("").Name("").ParentNodeId(""),
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
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cios.MultipleNode{Total: 10})
		if r.URL.Path == "/v2/file_storage/buckets/test/nodes" {
			query = r.URL.Query()
		}
	})
	ts := httptest.NewServer(bucketHandler)
	client := NewCiosClient(
		CiosClientConfig{
			Urls: sdkmodel.CIOSUrl{StorageUrl: ts.URL},
		},
	)
	defer ts.Close()
	for _, test := range tests {
		client.FileStorage.GetNodes(ctx, "test", test.params)
		test.test()
	}

	ts.Close()
}

func TestFileStorage_GetNodesAll(t *testing.T) {
	var offsets []int
	var limits []int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleNode{Total: 3500, Nodes: []cios.Node{}}
		offset := cnv.MustInt(r.URL.Query().Get("offset"))
		limit := cnv.MustInt(r.URL.Query().Get("limit"))
		offsets = append(offsets, offset)
		limits = append(limits, limit)
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Nodes = append(response.Nodes, cios.Node{Id: cnv.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{StorageUrl: ts.URL}})

	response, _, _ := client.FileStorage.GetNodesAll(ciosctx.Background(), "test", srvfilestorage.MakeGetNodesOpts().Limit(999))
	if len(response) != 999 || offsets[0] != 0 && limits[0] != 1000 {
		t.Fatal(len(response))
	}

	offsets = []int{}
	limits = []int{}
	response, _, _ = client.FileStorage.GetNodesAll(ciosctx.Background(), "test", srvfilestorage.MakeGetNodesOpts().Limit(1500))
	if len(response) != 1500 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 {
		t.Fatal(len(response), limits, offsets)
	}
	offsets = []int{}
	limits = []int{}
	response, _, _ = client.FileStorage.GetNodesAll(ciosctx.Background(), "test", srvfilestorage.MakeGetNodesOpts().Limit(2001))
	if len(response) != 2001 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 || offsets[2] != 2000 || limits[2] != 1 {
		t.Fatal(len(response), limits, offsets)

	}
	offsets = []int{}
	limits = []int{}
	response, _, _ = client.FileStorage.GetNodesAll(ciosctx.Background(), "test", srvfilestorage.MakeGetNodesOpts().Limit(3501))
	if len(response) != 3500 ||
		offsets[0] != 0 || limits[0] != 1000 ||
		offsets[1] != 1000 && limits[1] != 1000 ||
		offsets[2] != 2000 || limits[2] != 1000 ||
		offsets[3] != 3000 || limits[3] != 501 {
		t.Fatal(len(response), limits, offsets)
	}
	offsets = []int{}
	limits = []int{}
	response, _, _ = client.FileStorage.GetNodesAll(ciosctx.Background(), "test", srvfilestorage.MakeGetNodesOpts().Limit(2001).Offset(20))
	if len(response) != 2001 || offsets[0] != 20 && limits[0] != 1000 || offsets[1] != 1020 && limits[1] != 1000 || offsets[2] != 2020 || limits[2] != 1 {
		t.Fatal(len(response), limits, offsets)

	}
}

func TestFileStorage_GetNodesUnlimited(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleNode{Total: 3500, Nodes: []cios.Node{}}
		offset := cnv.MustInt(r.URL.Query().Get("offset"))
		limit := cnv.MustInt(r.URL.Query().Get("limit"))
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Nodes = append(response.Nodes, cios.Node{Id: cnv.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{StorageUrl: ts.URL}})

	buckets, _, _ := client.FileStorage.GetNodesUnlimited(ciosctx.Background(), "test", srvfilestorage.MakeGetNodesOpts().Limit(1))
	if len(buckets) != 3500 {
		t.Fatal(len(buckets))
	}
}

func TestFileStorage_GetNode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/v2/file_storage/buckets/test1/nodes/test2" {
			response := cios.SingleNode{Node: cios.Node{
				Id:          "test",
				IsDirectory: true,
				Key:         "",
				File:        nil,
				CreatedBy:   nil,
				UpdatedBy:   nil,
				Name:        "",
			}}
			json.NewEncoder(w).Encode(response)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{StorageUrl: ts.URL}})
	responseBody, response, err := client.FileStorage.GetNode(ciosctx.Background(), "test1", "test2")
	if responseBody.Id != "test" || err != nil || response.StatusCode != 200 {
		t.Fatal(responseBody)
	}
}

func TestFileStorage_CreateNode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body := cios.NodeRequest{}
		if r.Method != "POST" {
			t.Fatal(r.Method)
		}
		byts, _ := ioutil.ReadAll(r.Body)
		cnv.UnMarshalJson(byts, &body)
		if body.Name != "name" || *body.ParentNodeId != "parent" {
			t.Fatal(body)
		}

	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{StorageUrl: ts.URL}})
	_, _, err := client.FileStorage.CreateNode(ciosctx.Background(), "test", "name", cnv.StrPtr("parent"))
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestFileStorage_DeleteNode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v2/file_storage/buckets/bucketid/nodes/delete" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "DELETE" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{StorageUrl: ts.URL}})
	_, err := client.FileStorage.DeleteNode(ciosctx.Background(), "bucketid", "delete")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestFileStorage_RenameNode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		byts, _ := ioutil.ReadAll(r.Body)
		body := cios.NodeName{}
		cnv.UnMarshalJson(byts, &body)
		if body.Name != "name" || r.URL.Path != "/v2/file_storage/buckets/bucketid/nodes/node/rename" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "POST" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{StorageUrl: ts.URL}})
	_, _, err := client.FileStorage.RenameNode(ciosctx.Background(), "bucketid", "node", "name")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestFileStorage_CopyNode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		byts, _ := ioutil.ReadAll(r.Body)
		body := cios.BucketEditBody{}
		cnv.UnMarshalJson(byts, &body)
		if *body.DestBucketId != "destBucket" || *body.ParentNodeId != "destNode" || r.URL.Path != "/v2/file_storage/buckets/bucketid/nodes/node/copy" {
			t.Fatal(r.URL.Path, body)
		}
		if r.Method != "POST" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{StorageUrl: ts.URL}})
	_, _, err := client.FileStorage.CopyNode(ciosctx.Background(), "bucketid", "node", cnv.StrPtr("destBucket"), cnv.StrPtr("destNode"))
	if err != nil {
		t.Fatal(err.Error())
	}
}
func TestFileStorage_MoveNode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		byts, _ := ioutil.ReadAll(r.Body)
		body := cios.BucketEditBody{}
		cnv.UnMarshalJson(byts, &body)
		if *body.DestBucketId != "destBucket" || *body.ParentNodeId != "destNode" || r.URL.Path != "/v2/file_storage/buckets/bucketid/nodes/node/move" {
			t.Fatal(r.URL.Path, body)
		}
		if r.Method != "POST" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{StorageUrl: ts.URL}})
	_, _, err := client.FileStorage.MoveNode(ciosctx.Background(), "bucketid", "node", cnv.StrPtr("destBucket"), cnv.StrPtr("destNode"))
	if err != nil {
		t.Fatal(err.Error())
	}
}
