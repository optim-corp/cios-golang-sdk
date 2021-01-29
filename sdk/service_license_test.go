package ciossdk

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/optim-corp/cios-golang-sdk/model"
)

func TestPubSub_GetLicenses(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v2/licenses/me" {
			t.Fatal(r.URL.Path)
		}
		if r.Method != "GET" {
			t.Fatal(r.Method)
		}
	}))
	defer ts.Close()
	client := NewCiosClient(CiosClientConfig{Urls: model.CIOSUrl{LicenseUrl: ts.URL}})
	client.License.GetLicenses(MakeGetLicensesOpts(), context.Background())
}
