package ciossdk

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	cnv "github.com/fcfcqloow/go-advance/convert"
	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

func TestAccount_InviteGroup(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body := cios.GroupInviteRequest{}
		byt, _ := ioutil.ReadAll(r.Body)
		cnv.UnMarshalJson(byt, &body)
		if *body.Email != "email@temp.com" {
			t.Fatal(body)
		}
		if r.URL.Path != "/v2/groups/id/invites" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "POST" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{AccountsUrl: ts.URL}})
	_, _, err := client.Account().InviteGroup(context.Background(), "id", "email@temp.com")
	if err != nil {
		t.Fatal(err.Error())
	}
}
