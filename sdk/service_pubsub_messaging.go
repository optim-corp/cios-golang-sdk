package ciossdk

import (
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/optim-corp/cios-golang-sdk/cios"

	_nethttp "net/http"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/optim-corp/cios-golang-sdk/model"
)

func (self PubSub) PublishMessage(id string, body interface{}, packerFormat *string, ctx model.RequestCtx) (*_nethttp.Response, error) {
	request := self.ApiClient.PublishSubscribeApi.PublishMessage(ctx).ChannelId(id).Body(body)
	if packerFormat != nil {
		request = request.PackerFormat(*packerFormat)
	}
	httpResponse, err := request.Execute()
	if err != nil && self.autoR {
		if _, _, _, _, err := self.refresh(); err != nil {
			return nil, err
		}
		return request.Execute()
	}
	return httpResponse, err
}
func (self PubSub) PublishMessagePackerOnly(id string, body interface{}, ctx model.RequestCtx) (*_nethttp.Response, error) {
	return self.PublishMessage(id, &body, nil, ctx)
}
func (self PubSub) PublishMessageJSON(id string, body cios.PackerFormatJson, ctx model.RequestCtx) (*_nethttp.Response, error) {
	format := "json"
	return self.PublishMessage(id, &body, &format, ctx)
}

type (
	ConnectWebSocketOptions struct {
		PackerFormat  *string
		SubscribeFunc *func(body []byte) (bool, error)
		PublishStr    *chan *string
		Setting       *func(*websocket.Conn)
		Context       model.RequestCtx
	}
)

func safeCloseChan(ch chan bool) {
	isClosed := func() bool {
		select {
		case _, ok := <-ch:
			return !ok
		default:
			return false
		}
	}
	if !isClosed() {
		close(ch)
	}
}
func safeSendChan(ch chan bool, v bool) {
	isClosed := func() bool {
		select {
		case _, ok := <-ch:
			return !ok
		default:
			return false
		}
	}
	if !isClosed() {
		ch <- v
	}
}
func (self PubSub) ConnectWebSocket(channelID string, done chan bool, params ConnectWebSocketOptions) (err error) {
	if params.SubscribeFunc == nil && params.PublishStr == nil {
		return errors.New("no publish str and subscribe func")
	}
	mode, _token, ok := "subscribe", "", true
	if params.SubscribeFunc == nil && params.PublishStr != nil {
		mode = "publish"
	} else if params.SubscribeFunc != nil && params.PublishStr != nil {
		mode = "pubsub"
	}
	if params.Context != nil {
		_token, ok = params.Context.Value(cios.ContextAccessToken).(string)
	}
	if _token == "" && !ok && self.autoR {
		_token, _, _, _, _ = self.refresh()

	}

	var (
		wsUrl                     = self.CreateMessagingURL(channelID, mode, params.PackerFormat)
		connection, connectionErr = self.CreateCIOSWebsocketConnection(wsUrl, ParseAccessToken(_token))
		wg                        sync.WaitGroup
		finish                    = false
		beUpdating                = false
		beCompletedS              = make(chan bool, 1)
		beCompletedP              = make(chan bool, 1)
		debug                     = func(text ...interface{}) {
			if self.debug {
				log.Println(text...)
			}
		}
		closeConnection = func() {
			if connection != nil {
				if _err := connection.Close(); _err != nil {
					debug("\nClosed Connection\n")
				}
			}
		}
		cleanBeCompleted = func() {
			for {
				select {
				case <-beCompletedP:
					debug("ignore publish update")
					continue
				case <-beCompletedS:
					debug("ignore subscribe update")
					continue
				default:
				}
				break
			}
		}
		reconnection = func() error {
			token, _, _, _, _ := self.refresh()
			_connection, _err := self.CreateCIOSWebsocketConnection(wsUrl, ParseAccessToken(token))
			closeConnection()
			if _err != nil {
				return _err
			}
			connection = _connection
			return nil
		}
		closeLogic = func() {
			if params.PublishStr != nil {
				isClosed := func() (f bool) {
					select {
					case _, ok = <-*params.PublishStr:
						f = !ok
					default:
					}
					return
				}
				if !isClosed() {
					close(*params.PublishStr)
				}
			}
			beUpdating = false
			finish = true
			closeConnection()
			safeSendChan(done, false)
			safeCloseChan(done)
			safeCloseChan(beCompletedS)
			safeCloseChan(beCompletedP)
		}
		updateConnection = func() {
			for {
				select {
				case <-time.After(time.Second * 270):
					if self.autoR {
						cleanBeCompleted()
						beUpdating = true
						debug("Updating access token and connection")
						if err = reconnection(); err != nil {
							closeLogic()
							beUpdating = false
							break
						}
						beCompletedS <- true
						beCompletedP <- true
						debug("Completed reconnection")
						beUpdating = false
					}
					break
				case isClosing := <-done:
					if isClosing {
						closeLogic()
					}
					break
				}
				if finish {
					break
				}
			}
		}
		subscribe = func() {
			for {
				if finish {
					break
				}
				if _, body, _err := connection.ReadMessage(); _err != nil {
					if beUpdating {
						if isConnecting, ok := <-beCompletedS; ok && isConnecting {
							continue
						}
					}
					err = errors.New("close connection")
					break
				} else if finish, err = (*params.SubscribeFunc)(body); err != nil || finish {
					break
				} else {
					debug("Receive text: " + string(body))
				}
			}
			closeLogic()
		}
		publish = func() {
			for {
				text, ok := <-(*params.PublishStr)
				if text == nil || !ok {
					break
				}
				if beUpdating {
					<-beCompletedP
				}
				if _err := connection.WriteMessage(websocket.TextMessage, []byte(*text)); _err != nil {
					if _err = connection.WriteMessage(websocket.TextMessage, []byte(*text)); _err != nil {
						err = errors.New("close connection")
						break
					}
				}
				debug("Publish text: " + *text)
			}
			closeLogic()
		}
	)
	if err = connectionErr; err != nil {
		safeCloseChan(beCompletedS)
		safeCloseChan(beCompletedP)
		return err
	}
	if params.Setting != nil {
		(*params.Setting)(connection)
	}
	wg.Add(1)
	go func() {
		updateConnection()
		wg.Done()
	}()
	if params.SubscribeFunc != nil {
		wg.Add(1)
		go func() {
			subscribe()
			wg.Done()
		}()
	}
	if params.PublishStr != nil {
		wg.Add(1)
		go func() {
			publish()
			wg.Done()
		}()
	}
	wg.Wait()
	debug(err)
	return err
}
func (self PubSub) SubscribeCiosWebSocket(_url string, beforeFunc *func(*websocket.Conn), logic func(body []byte) (bool, error), ctx model.RequestCtx) error {
	_token, ok := "", true
	if ctx != nil {
		_token, ok = ctx.Value(cios.ContextAccessToken).(string)
	}
	if _token == "" && !ok && self.autoR {
		_token, _, _, _, _ = self.refresh()

	}
	connection, err := self.CreateCIOSWebsocketConnection(_url, ParseAccessToken(_token))
	if err != nil {
		return err
	}
	defer connection.Close()
	if beforeFunc != nil {
		(*beforeFunc)(connection)
	}
	for {
		_, body, err := connection.ReadMessage()
		if err != nil {
			connection, err = self.CreateCIOSWebsocketConnection(_url, _token)
			if err != nil {
				return err
			}
			defer connection.Close()
			continue
		}
		done, err := logic(body)
		if err != nil {
			return err
		}
		if done {
			return nil
		}
	}
}

func (self PubSub) CreateMessagingURL(channelID string, mode string, packerFormat *string) string {
	_url, err := url.Parse(strings.Replace(self.Url, "https", "wss", 1) + "/v2/messaging")
	if err != nil {
		return ""
	}
	q := _url.Query()
	q.Set("channel_id", channelID)
	q.Set("mode", mode)
	if packerFormat != nil {
		q.Set("packer_format", *packerFormat)
	}
	_url.RawQuery = q.Encode()
	return _url.String()
}

func (self PubSub) CreateCIOSWebsocketConnection(url string, authorization string) (*websocket.Conn, error) {
	if self.debug {
		log.Printf("Websocket URL: %s\nAuthorization: %s", url, authorization)
	}
	dialer := websocket.Dialer{}
	headers := http.Header{"Authorization": []string{authorization}}
	connection, _, err := (&dialer).Dial(url, headers)
	if err != nil {
		return nil, err
	}
	return connection, nil
}
