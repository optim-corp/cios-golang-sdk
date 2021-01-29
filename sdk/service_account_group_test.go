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
		query = r.URL.Query()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cios.MultipleGroup{Total: 10})
	})
	ts := httptest.NewServer(responseHandler)
	client := NewCiosClient(CiosClientConfig{Urls: model.CIOSUrl{AccountsUrl: ts.URL}})
	defer ts.Close()
	for _, test := range tests {
		client.Account.GetGroups(test.params, ctx)
		test.test()
	}

	ts.Close()

	//// Auto Refresh Test
	//client = NewCiosClient(
	//	CiosClientConfig{
	//		Urls:        model.CIOSUrl{AccountsUrl: ts.URL},
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
	//refFunc := func() (model.AccessToken, model.Scope, model.TokenType, model.ExpiresIn, error) {
	//	result = "Accept"
	//	return "", "", "", 0, nil
	//}
	//client.Account.refresh = &refFunc
	//if result == "Failed" {
	//	t.Fatal("Cant Refresh", result)
	//}
	////　念のためクローズ
	//ts.Close()
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
	client := NewCiosClient(CiosClientConfig{Urls: model.CIOSUrl{AccountsUrl: ts.URL}})

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
	client := NewCiosClient(CiosClientConfig{Urls: model.CIOSUrl{AccountsUrl: ts.URL}})

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
	client := NewCiosClient(CiosClientConfig{Urls: model.CIOSUrl{AccountsUrl: ts.URL}})
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
	client := NewCiosClient(CiosClientConfig{Urls: model.CIOSUrl{AccountsUrl: ts.URL}})
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
	client := NewCiosClient(CiosClientConfig{Urls: model.CIOSUrl{AccountsUrl: ts.URL}})
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
	client := NewCiosClient(CiosClientConfig{Urls: model.CIOSUrl{AccountsUrl: ts.URL}})
	client.Account.UpdateGroup("id", cios.GroupUpdateRequest{
		Name: convert.StringPtr("name"),
	}, context.Background())
}
