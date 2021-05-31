package ciossdk

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	srvfilestorage "github.com/optim-corp/cios-golang-sdk/sdk/service/filestorage"

	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"

	xmath "github.com/fcfcqloow/go-advance/math"

	cnv "github.com/fcfcqloow/go-advance/convert"

	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

func Test_RefreshBucket(t *testing.T) {
	count, _token, ctx := 0, "", context.Background()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_token = r.Header.Get("Authorization")
		if r.URL.Path == "/connect/token" {
			count += 1
			json.NewEncoder(w).Encode(cios.ConnectTokenResponse{
				AccessToken:  sampleToken,
				TokenType:    "",
				RefreshToken: "",
				ExpiresIn:    0,
				Scope:        "",
			})
		}
	})
	ts := httptest.NewServer(handler)
	client := NewCiosClient(CiosClientConfig{
		AutoRefresh: true,
		Urls:        sdkmodel.CIOSUrl{StorageUrl: ts.URL, AuthUrl: ts.URL},
		AuthConfig: RefreshTokenAuth(
			"clientID",
			"clientSecret",
			"assertion",
			"scope",
		),
	})
	funcs := []func(){
		func() { client.FileStorage.GetBuckets(ctx, srvfilestorage.MakeGetBucketsOpts()) },
		func() { client.FileStorage.CreateBucket(ctx, "", "") },
		func() { client.FileStorage.GetBucket(ctx, "") },
		func() { client.FileStorage.UpdateBucket(ctx, "", "") },
		func() { client.FileStorage.DeleteBucket(ctx, "") },
	}
	for i, fnc := range funcs {
		fnc()
		if _token != "Bearer "+sampleToken {
			t.Fatal(_token)
		}
		if count != i+1 {
			t.Fatal(count)
		}
		_token = ""
	}
	for _, fnc := range funcs {
		client.tokenExp = time.Now().Unix() + 10000
		fnc()
		if _token != "Bearer "+sampleToken {
			t.Fatal(_token)
		}
		if count != len(funcs) {
			t.Fatal(count)
		}
		_token = ""
	}

	ctx = ciosctx.WithToken(ctx, "AAA")
	count = 0
	_token = ""
	client = NewCiosClient(CiosClientConfig{
		AutoRefresh: true,
		Urls:        sdkmodel.CIOSUrl{StorageUrl: ts.URL, AuthUrl: ts.URL},
		AuthConfig: RefreshTokenAuth(
			"clientID",
			"clientSecret",
			"assertion",
			"scope",
		),
	})
	for i, fnc := range funcs {
		fnc()
		if _token != "Bearer AAA" {
			t.Fatal(_token)
		}
		if count != i+1 {
			t.Fatal(count)
		}
		_token = ""
	}
}

func TestFileStorage_GetBuckets(t *testing.T) {
	var (
		query url.Values
		ctx   context.Context

		tests = []struct {
			params cios.ApiGetBucketsRequest
			test   func()
		}{
			{
				params: srvfilestorage.MakeGetBucketsOpts().Limit(1000),
				test: func() {
					if query.Get("limit") != "1000" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvfilestorage.MakeGetBucketsOpts().Limit(1000).Offset(50),
				test: func() {
					if query.Get("limit") != "1000" || query.Get("offset") != "50" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvfilestorage.MakeGetBucketsOpts().ResourceOwnerId("aaaaa"),
				test: func() {
					if query.Get("resource_owner_id") != "aaaaa" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvfilestorage.MakeGetBucketsOpts().Limit(1000).Offset(50).ResourceOwnerId("aaaaa").Name("name"),
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
				params: srvfilestorage.MakeGetBucketsOpts().OrderBy("created_at"),
				test: func() {
					if query.Get("order_by") != "created_at" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: srvfilestorage.MakeGetBucketsOpts().OrderBy("").Order("").ResourceOwnerId("").Name(""),
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
			Urls: sdkmodel.CIOSUrl{StorageUrl: ts.URL},
		},
	)
	defer ts.Close()
	for _, test := range tests {
		client.FileStorage.GetBuckets(ctx, test.params)
		test.test()
	}

	ts.Close()
}

func TestFileStorage_GetBucketsAll(t *testing.T) {
	var offsets []int
	var limits []int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleBucket{Total: 3500, Buckets: []cios.Bucket{}}
		offset := cnv.MustInt(r.URL.Query().Get("offset"))
		limit := cnv.MustInt(r.URL.Query().Get("limit"))
		offsets = append(offsets, offset)
		limits = append(limits, limit)
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Buckets = append(response.Buckets, cios.Bucket{Id: cnv.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{StorageUrl: ts.URL}})

	buckets, _, _ := client.FileStorage.GetBucketsAll(nil, srvfilestorage.MakeGetBucketsOpts().Limit(999))
	if len(buckets) != 999 || offsets[0] != 0 && limits[0] != 1000 {
		t.Fatal(len(buckets))
	}

	offsets = []int{}
	limits = []int{}
	buckets, _, _ = client.FileStorage.GetBucketsAll(nil, srvfilestorage.MakeGetBucketsOpts().Limit(1500))
	if len(buckets) != 1500 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 {
		t.Fatal(len(buckets), limits, offsets)
	}
	offsets = []int{}
	limits = []int{}
	buckets, _, _ = client.FileStorage.GetBucketsAll(nil, srvfilestorage.MakeGetBucketsOpts().Limit(2001))
	if len(buckets) != 2001 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 || offsets[2] != 2000 || limits[2] != 1 {
		t.Fatal(len(buckets), limits, offsets)

	}
	offsets = []int{}
	limits = []int{}
	buckets, _, _ = client.FileStorage.GetBucketsAll(nil, srvfilestorage.MakeGetBucketsOpts().Limit(3501))
	if len(buckets) != 3500 ||
		offsets[0] != 0 || limits[0] != 1000 ||
		offsets[1] != 1000 && limits[1] != 1000 ||
		offsets[2] != 2000 || limits[2] != 1000 ||
		offsets[3] != 3000 || limits[3] != 501 {
		t.Fatal(len(buckets), limits, offsets)
	}
	offsets = []int{}
	limits = []int{}
	buckets, _, _ = client.FileStorage.GetBucketsAll(nil, srvfilestorage.MakeGetBucketsOpts().Limit(2001).Offset(20))
	if len(buckets) != 2001 || offsets[0] != 20 && limits[0] != 1000 || offsets[1] != 1020 && limits[1] != 1000 || offsets[2] != 2020 || limits[2] != 1 {
		t.Fatal(len(buckets), limits, offsets)

	}
}

func TestFileStorage_GetBucketsUnlimited(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleBucket{Total: 3500, Buckets: []cios.Bucket{}}
		offset := cnv.MustInt(r.URL.Query().Get("offset"))
		limit := cnv.MustInt(r.URL.Query().Get("limit"))
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Buckets = append(response.Buckets, cios.Bucket{Id: cnv.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{StorageUrl: ts.URL}})

	buckets, _, _ := client.FileStorage.GetBucketsUnlimited(nil, srvfilestorage.MakeGetBucketsOpts().Limit(1))
	if len(buckets) != 3500 {
		t.Fatal(len(buckets))
	}
}

func TestFileStorage_GetBucket(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/v2/file_storage/buckets/test" {
			response := cios.SingleBucket{Bucket: cios.Bucket{
				Id:              "test",
				ResourceOwnerId: "test_resource_owner",
				CreatedAt:       nil,
				CreatedBy:       nil,
				UpdatedAt:       nil,
				UpdatedBy:       nil,
				Name:            "",
				Files:           nil,
			}}
			json.NewEncoder(w).Encode(response)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{StorageUrl: ts.URL}})
	bucket, response, err := client.FileStorage.GetBucket(nil, "test")
	if bucket.Id != "test" || err != nil || response.StatusCode != 200 {
		t.Fatal(bucket)
	}
}

func TestFileStorage_CreateBucket(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body := cios.BucketRequest{}
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
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{StorageUrl: ts.URL}})
	_, _, err := client.FileStorage.CreateBucket(nil, "resource_owner_id", "name")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestFileStorage_DeleteBucket(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v2/file_storage/buckets/bucketid" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "DELETE" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{StorageUrl: ts.URL}})
	_, err := client.FileStorage.DeleteBucket(nil, "bucketid")
	if err != nil {
		t.Fatal(err.Error())
	}
}
func TestFileStorage_UpdateBucket(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body := cios.BucketName{}
		byts, _ := ioutil.ReadAll(r.Body)
		cnv.UnMarshalJson(byts, &body)
		if r.URL.Path != "/v2/file_storage/buckets/bucketid" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "PATCH" {
			t.Fatal(r.Method)
		}
		if body.Name != "test" {
			t.Fatal(body)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{StorageUrl: ts.URL}})
	_, err := client.FileStorage.UpdateBucket(nil, "bucketid", "test")
	if err != nil {
		t.Fatal(err.Error())
	}
}
