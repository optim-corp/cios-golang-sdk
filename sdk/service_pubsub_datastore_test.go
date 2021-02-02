package ciossdk

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

var upgrader = websocket.Upgrader{}

func TestPubSub_GetStream(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		for {
			_, message, _ := c.ReadMessage()
			fmt.Println(string(message))
			c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			break
		}
	}
	s := httptest.NewServer(http.HandlerFunc(handler))
	url := "ws" + strings.TrimPrefix(s.URL, "http")
	client := NewCiosClient(CiosClientConfig{Urls: sdkmodel.CIOSUrl{MessagingUrl: url}})
	client.PubSub.GetStream("id", MakeGetStreamOpts(), context.Background())
}
