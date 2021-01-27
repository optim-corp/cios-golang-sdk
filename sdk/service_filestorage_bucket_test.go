package ciossdk

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	xmath "github.com/optim-kazuhiro-seida/go-advance-type/math"

	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	"github.com/optim-corp/cios-golang-sdk/cios"

	"github.com/optim-corp/cios-golang-sdk/model"
)

func TestFileStorage_GetBuckets(t *testing.T) {
	var (
		query url.Values
		ctx   context.Context

		tests = []struct {
			params cios.ApiGetBucketsRequest
			test   func()
		}{
			{
				params: MakeGetBucketsOpts().Limit(1000),
				test: func() {
					if query.Get("limit") != "1000" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: MakeGetBucketsOpts().Limit(1000).Offset(50),
				test: func() {
					if query.Get("limit") != "1000" || query.Get("offset") != "50" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: MakeGetBucketsOpts().ResourceOwnerId("aaaaa"),
				test: func() {
					if query.Get("resource_owner_id") != "aaaaa" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: MakeGetBucketsOpts().Limit(1000).Offset(50).ResourceOwnerId("aaaaa").Name("name"),
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
				params: MakeGetBucketsOpts().OrderBy("created_at"),
				test: func() {
					if query.Get("order_by") != "created_at" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: MakeGetBucketsOpts().OrderBy("").Order("").ResourceOwnerId("").Name(""),
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
		json.NewEncoder(w).Encode(cios.MultipleBucket{Total: 10})
	})
	ts := httptest.NewServer(bucketHandler)
	client := NewCiosClient(
		CiosClientConfig{
			Urls: model.CIOSUrl{StorageUrl: ts.URL},
		},
	)
	defer ts.Close()
	for _, test := range tests {
		client.FileStorage.GetBuckets(test.params, ctx)
		test.test()
	}

	ts.Close()

	// Auto Refresh Test
	client = NewCiosClient(
		CiosClientConfig{
			Urls:        model.CIOSUrl{StorageUrl: ts.URL},
			AutoRefresh: true,
		},
	)
	bucketHandler = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
	}
	ts = httptest.NewServer(bucketHandler)

	result := "Failed"
	refFunc := func() (model.AccessToken, model.Scope, model.TokenType, model.ExpiresIn, error) {
		result = "Accept"
		return "", "", "", 0, nil
	}
	client.FileStorage.refresh = &refFunc
	if result == "Failed" {
		t.Fatal("Cant Refresh", result)
	}
	//　念のためクローズ
	ts.Close()
}

func TestFileStorage_GetBucketsAll(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleBucket{Total: 3500, Buckets: []cios.Bucket{}}
		offset := convert.MustInt(r.URL.Query().Get("offset"))
		limit := convert.MustInt(r.URL.Query().Get("limit"))
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Buckets = append(response.Buckets, cios.Bucket{Id: convert.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: model.CIOSUrl{StorageUrl: ts.URL}})

	buckets, _, _ := client.FileStorage.GetBucketsAll(MakeGetBucketsOpts().Limit(999), context.Background())
	if len(buckets) != 999 {
		t.Fatal(len(buckets))
	}
	buckets, _, _ = client.FileStorage.GetBucketsAll(MakeGetBucketsOpts().Limit(1500), context.Background())
	if len(buckets) != 1500 {
		t.Fatal(len(buckets))
	}
	buckets, _, _ = client.FileStorage.GetBucketsAll(MakeGetBucketsOpts().Limit(2001), context.Background())
	if len(buckets) != 2001 {
		t.Fatal(len(buckets))
	}
	buckets, _, _ = client.FileStorage.GetBucketsAll(MakeGetBucketsOpts().Limit(3501), context.Background())
	if len(buckets) != 3500 {
		t.Fatal(len(buckets))
	}
}
func TestFileStorage_GetBucketsUnlimited(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleBucket{Total: 3500, Buckets: []cios.Bucket{}}
		offset := convert.MustInt(r.URL.Query().Get("offset"))
		limit := convert.MustInt(r.URL.Query().Get("limit"))
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Buckets = append(response.Buckets, cios.Bucket{Id: convert.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: model.CIOSUrl{StorageUrl: ts.URL}})

	buckets, _, _ := client.FileStorage.GetBucketsUnlimited(MakeGetBucketsOpts().Limit(1), context.Background())
	if len(buckets) != 3500 {
		t.Fatal(len(buckets))
	}
}
