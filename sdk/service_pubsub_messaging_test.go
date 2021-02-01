package ciossdk

//import (
//	"net/http"
//	"net/http/httptest"
//	"strings"
//	"testing"
//
//	"github.com/optim-corp/cios-golang-sdk/model"
//
//	"github.com/gorilla/websocket"
//)
//
//var upgrader = websocket.Upgrader{}
//
//func TestPubSub_subscribeCiosWebSocket(t *testing.T) {
//	handler := func(w http.ResponseWriter, r *http.Request) {
//		c, err := upgrader.Upgrade(w, r, nil)
//		if err != nil {
//			return
//		}
//		defer c.Close()
//		for {
//			mt, message, err := c.ReadMessage()
//			if err != nil {
//				break
//			}
//
//			err = c.WriteMessage(mt, message)
//			if err != nil {
//				break
//			}
//		}
//	}
//	s := httptest.NewServer(http.HandlerFunc(handler))
//	url := "ws" + strings.TrimPrefix(s.URL, "http")
//	client := NewCiosClient(CiosClientConfig{Urls: model.CIOSUrl{MessagingUrl: url}})
//	client.PubSub.subscribeCiosWebSocket("")
//}
