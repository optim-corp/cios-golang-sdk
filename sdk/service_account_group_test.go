package ciossdk

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"
	"time"

	xmath "github.com/optim-kazuhiro-seida/go-advance-type/math"

	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	"github.com/optim-corp/cios-golang-sdk/cios"

	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

//　https://auth0.com/docs/tokens/access-tokens/use-access-tokens
var sampleToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2V4YW1wbGUuYXV0aDAuY29tLyIsImF1ZCI6Imh0dHBzOi8vYXBpLmV4YW1wbGUuY29tL2NhbGFuZGFyL3YxLyIsInN1YiI6InVzcl8xMjMiLCJpYXQiOjE0NTg3ODU3OTYsImV4cCI6MTQ1ODg3MjE5Nn0.CA7eaHjIHz5NxeIJoFK9krqaeZrPLwmMmgI_XiQiIkQ"

func Test_RefreshGroup(t *testing.T) {
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
		if r.URL.Path == "/v2/groups" && r.Method == http.MethodGet {
			time.Sleep(time.Microsecond * time.Duration(rand.Int63n(1000)))
			json.NewEncoder(w).Encode(cios.MultipleGroup{
				Groups: []cios.Group{
					{
						Id: r.URL.Query().Get("name") + ":" + r.Header.Get("Authorization"),
					},
				},
			})
		}
	})
	ts := httptest.NewServer(handler)
	client := NewCiosClient(CiosClientConfig{
		AutoRefresh: true,
		Urls:        sdkmodel.CIOSUrl{AccountsUrl: ts.URL, AuthUrl: ts.URL},
		AuthConfig: RefreshTokenAuth(
			"clientID",
			"clientSecret",
			"assertion",
			"scope",
		),
	})
	funcs := []func(){
		func() { client.Account.GetGroups(MakeGetGroupsOpts(), ctx) },
		func() { client.Account.CreateGroup(cios.GroupCreateRequest{}, ctx) },
		func() { client.Account.GetGroup("", nil, ctx) },
		func() { client.Account.UpdateGroup("", cios.GroupUpdateRequest{}, ctx) },
		func() { client.Account.DeleteGroup("", ctx) },
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

	ctx = MakeRequestCtx("AAA")
	count = 0
	_token = ""
	client = NewCiosClient(CiosClientConfig{
		AutoRefresh: true,
		Urls:        sdkmodel.CIOSUrl{AccountsUrl: ts.URL, AuthUrl: ts.URL},
		AuthConfig: RefreshTokenAuth(
			"clientID",
			"clientSecret",
			"assertion",
			"scope",
		),
	})

	// リクエストごとのトークンを識別しているかテスト
	wg := sync.WaitGroup{}
	asyncTestNumber := 10 //　非同期にリクエストする回数
	wg.Add(asyncTestNumber)
	for i := 0; i < asyncTestNumber; i++ {
		go func(i int) {
			time.Sleep(time.Microsecond * time.Duration(rand.Int63n(100)))
			response, _, _ := client.Account.GetGroups(MakeGetGroupsOpts().Name("async "+convert.MustStr(i)), MakeRequestCtx(convert.MustStr(i)))
			if response.Groups[0].Id != "async "+convert.MustStr(i)+":Bearer "+convert.MustStr(i) {
				t.Fatal(response.Groups)
			}
			t.Log(response)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
func TestAccount_Groups(t *testing.T) {
	var (
		query url.Values
		ctx   context.Context

		tests = []struct {
			params cios.ApiGetGroupsRequest
			test   func()
		}{
			{
				params: MakeGetGroupsOpts().Limit(1000),
				test: func() {
					if query.Get("limit") != "1000" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: MakeGetGroupsOpts().Limit(1000).Offset(50),
				test: func() {
					if query.Get("limit") != "1000" || query.Get("offset") != "50" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: MakeGetGroupsOpts().Limit(1000).Offset(50).Name("name"),
				test: func() {
					if query.Get("name") != "name" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: MakeGetGroupsOpts().OrderBy("created_at"),
				test: func() {
					if query.Get("order_by") != "created_at" {
						t.Fatal("Missing Query", query.Encode())
					} else {
						t.Log(query.Encode())
					}
				},
			},
			{
				params: MakeGetGroupsOpts().
					OrderBy("").
					Order("").
					Name("").
					Label("").
					Page("").
					Address1("").
					Address2("").
					City("").
					Tags("").Domain("").Includes("").State("").ParentGroupId("").Type_(""),
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
		w.Header().Set("Content-Type", "application/json")
		query = r.URL.Query()
		json.NewEncoder(w).Encode(cios.MultipleGroup{Total: 10})
	})
	ts := httptest.NewServer(responseHandler)

	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{AccountsUrl: ts.URL}})

	defer ts.Close()
	for _, test := range tests {
		client.Account.GetGroups(test.params, ctx)
		test.test()
	}

	ts.Close()
}

func TestAccount_GetGroupsAll(t *testing.T) {
	var offsets []int
	var limits []int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleGroup{Total: 3500, Groups: []cios.Group{}}
		offset := convert.MustInt(r.URL.Query().Get("offset"))
		limit := convert.MustInt(r.URL.Query().Get("limit"))
		offsets = append(offsets, offset)
		limits = append(limits, limit)
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Groups = append(response.Groups, cios.Group{Id: convert.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{AccountsUrl: ts.URL}})

	responses, _, _ := client.Account.GetGroupsAll(MakeGetGroupsOpts().Limit(999), context.Background())
	if len(responses) != 999 || offsets[0] != 0 && limits[0] != 1000 {
		t.Fatal(len(responses))
	}

	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.Account.GetGroupsAll(MakeGetGroupsOpts().Limit(1500), context.Background())
	if len(responses) != 1500 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 {
		t.Fatal(len(responses), limits, offsets)
	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.Account.GetGroupsAll(MakeGetGroupsOpts().Limit(2001), context.Background())
	if len(responses) != 2001 || offsets[0] != 0 && limits[0] != 1000 || offsets[1] != 1000 && limits[1] != 1000 || offsets[2] != 2000 || limits[2] != 1 {
		t.Fatal(len(responses), limits, offsets)

	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.Account.GetGroupsAll(MakeGetGroupsOpts().Limit(3501), context.Background())
	if len(responses) != 3500 ||
		offsets[0] != 0 || limits[0] != 1000 ||
		offsets[1] != 1000 && limits[1] != 1000 ||
		offsets[2] != 2000 || limits[2] != 1000 ||
		offsets[3] != 3000 || limits[3] != 501 {
		t.Fatal(len(responses), limits, offsets)
	}
	offsets = []int{}
	limits = []int{}
	responses, _, _ = client.Account.GetGroupsAll(MakeGetGroupsOpts().Limit(2001).Offset(20), context.Background())
	if len(responses) != 2001 || offsets[0] != 20 && limits[0] != 1000 || offsets[1] != 1020 && limits[1] != 1000 || offsets[2] != 2020 || limits[2] != 1 {
		t.Fatal(len(responses), limits, offsets)

	}
}

func TestAccount_GetGroupsUnlimited(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := cios.MultipleGroup{Total: 3500, Groups: []cios.Group{}}
		offset := convert.MustInt(r.URL.Query().Get("offset"))
		limit := convert.MustInt(r.URL.Query().Get("limit"))
		for i := 0; i < xmath.MinInt(3500-offset, 1000, limit); i++ {
			response.Groups = append(response.Groups, cios.Group{Id: convert.MustStr(i)})
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{AccountsUrl: ts.URL}})

	response, _, _ := client.Account.GetGroupsUnlimited(MakeGetGroupsOpts().Limit(1), context.Background())
	if len(response) != 3500 {
		t.Fatal(len(response))
	}
}

func TestAccount_GetGroup(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != "GET" {
			t.Fatal(r.Method)
		}
		if r.URL.Path == "/v2/groups/test" {
			response := cios.Group{
				Id:   "test",
				Name: "",
			}
			json.NewEncoder(w).Encode(response)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{AccountsUrl: ts.URL}})
	responseB, response, err := client.Account.GetGroup("test", nil, context.Background())
	if responseB.Id != "test" || err != nil || response.StatusCode != 200 {
		t.Fatal(responseB)
	}
}

func TestAccount_CreateGroup(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v2/groups" {
			t.Fatal(r.URL.Path)
		}
		body := cios.GroupCreateRequest{}
		if r.Method != "POST" {
			t.Fatal(r.Method)
		}
		byts, _ := ioutil.ReadAll(r.Body)
		convert.UnMarshalJson(byts, &body)
		if body.Name != "name" ||
			*body.ParentGroupId != "parent" ||
			(*body.Tags)[0] != "key=value" {
			t.Fatal(body)
		}

	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{AccountsUrl: ts.URL}})
	client.Account.CreateGroup(cios.GroupCreateRequest{
		ParentGroupId: convert.StringPtr("parent"),
		Name:          "name",
		Tags:          &[]string{"key=value", "a=b"},
	}, context.Background())
}

func TestAccount_DeleteGroup(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v2/groups/id" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "DELETE" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{AccountsUrl: ts.URL}})
	client.Account.DeleteGroup("id", context.Background())
}
func TestAccount_UpdateGroup(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v2/groups/id" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "PATCH" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{AccountsUrl: ts.URL}})
	client.Account.UpdateGroup("id", cios.GroupUpdateRequest{
		Name: convert.StringPtr("name"),
	}, context.Background())
}
