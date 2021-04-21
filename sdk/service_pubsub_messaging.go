package ciossdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
	"github.com/optim-kazuhiro-seida/go-advance-type/check"

	_nethttp "net/http"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
)

type (
	CiosMessaging struct {
		SubscribeFunc func([]byte) (bool, error)
		CloseFunc     func()
		Connection    *websocket.Conn
		isUpdating    bool
		wsUrl         string
		isDebug       bool
		token         string
		closed        chan bool
		readDeadTime  time.Duration
		writeDeadTime time.Duration
		refresh       *func() sdkmodel.AccessToken
	}

	// Deprecated: should not be used
	ConnectWebSocketOptions struct {
		PackerFormat  *string
		SubscribeFunc *func(body []byte) (bool, error)
		PublishStr    *chan *string
		Setting       *func(*websocket.Conn)
		Context       sdkmodel.RequestCtx
	}
)

var (
	CreateCiosWsConn = func(isDebug bool, url, authorization string) (connection *websocket.Conn, err error) {
		if isDebug {
			log.Printf("Websocket URL: %s\nAuthorization: %s", url, authorization)
		}
		connection, _, err = (&websocket.Dialer{}).Dial(url, http.Header{"Authorization": []string{authorization}})
		return
	}
	CreateCiosWsMessagingURL = func(httpUrl, channelID, mode string, packerFormat *string) string {
		_url, err := url.Parse(strings.Replace(httpUrl, "https", "wss", 1) + "/v2/messaging")
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
)

func (self *PubSub) NewMessaging(channelId, mode, packerFormat string) *CiosMessaging {
	ref := func() (token string) {
		self.refresh()
		if !check.IsNil(self.token) {
			token = *self.token
		}
		return
	}
	if packerFormat == "" {
		packerFormat = "payload_only"
	}
	instance := CiosMessaging{}
	instance.wsUrl = CreateCiosWsMessagingURL(self.Url, channelId, mode, &packerFormat)
	instance.isDebug = self.debug
	instance.refresh = &ref
	instance.CloseFunc = func() {}
	return &instance
}

func (self *CiosMessaging) SetWriteTimeout(t time.Duration) *CiosMessaging {
	self.writeDeadTime = t
	return self
}
func (self *CiosMessaging) SetReadTimeout(t time.Duration) *CiosMessaging {
	self.readDeadTime = t
	return self
}
func (self *CiosMessaging) debug(text ...interface{}) {
	if self.isDebug {
		result := ""
		for _, t := range text {
			result += convert.MustStr(t) + " "
		}
		log.Println(result)
	}
}
func (self *CiosMessaging) OnReceive(arg func([]byte) (bool, error)) error {
	for {
		body, err := self.Receive()
		if err != nil {
			return err
		}
		if ok, err := arg(body); !ok || err != nil {
			return err
		}
	}
}
func (self *CiosMessaging) OnClose(arg func()) {
	self.CloseFunc = arg
}
func (self *CiosMessaging) Send(message []byte) error {
	if check.IsNil(self.Connection) {
		return fmt.Errorf("no connection used Start()")
	}
	defer self.debug("Send: " + string(message))
	if self.writeDeadTime != 0 {
		if err := self.Connection.SetWriteDeadline(time.Now().Add(self.writeDeadTime)); err != nil {
			self.debug("Set Write Timeout", err)
			return err
		}
	}
	return self.Connection.WriteMessage(websocket.TextMessage, message)
}
func (self *CiosMessaging) SendStr(message string) error {
	return self.Send([]byte(message))
}
func (self *CiosMessaging) SendAny(message interface{}) error {
	return self.SendStr(convert.MustStr(message))
}
func (self *CiosMessaging) SendJson(message interface{}) error {
	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return self.Send(msg)
}
func (self *CiosMessaging) Publish(message interface{}) error {
	return self.SendAny(message)
}

func (self *CiosMessaging) receive() (body []byte, err error) {
	var messageType int
	if check.IsNil(self.Connection) {
		err = fmt.Errorf("no connection use Start()")
		return
	}
	if self.readDeadTime != 0 {
		if err = self.Connection.SetReadDeadline(time.Now().Add(self.readDeadTime)); err != nil {
			self.debug("Set Read Timeout", err)
			return
		}
	}
	messageType, body, err = self.Connection.ReadMessage()
	switch {
	case websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived):
		err = nil
		self.debug("Receive close err: ", fmt.Sprintf("%d, %s", messageType, convert.MustStr(err)))
	case websocket.IsUnexpectedCloseError(err):
		self.debug("Receive unexpected close err: ", fmt.Sprintf("%d, %s", messageType, convert.MustStr(err)))
	case messageType == websocket.CloseMessage:
		err = nil
		self.debug("Receive CloseMessage: ", fmt.Sprintf("%d, %s", messageType, convert.MustStr(err)))
	case messageType == websocket.TextMessage:
		self.debug(string(body))
	}
	return
}
func (self *CiosMessaging) Receive() (body []byte, err error) {
AGAIN:
	body, err = self.receive()
	if err != nil && !check.IsNil(self.refresh) {
		self.token = (*self.refresh)()
		self.debug("Receive Refresh")
		if _connection, _err := CreateCiosWsConn(self.isDebug, self.wsUrl, ParseAccessToken(self.token)); _err == nil {
			self.debug("Close err: ", self.Connection.Close())
			self.Connection = _connection
			goto AGAIN
		}
	}
	return
}
func (self *CiosMessaging) ReceiveStr() (string, error) {
	res, err := self.Receive()
	return string(res), err
}
func (self *CiosMessaging) MapReceived(stct interface{}) error {
	res, err := self.Receive()
	if err != nil {
		return err
	}
	return convert.UnMarshalJson(res, stct)
}

func (self *CiosMessaging) Start(ctx sdkmodel.RequestCtx) (err error) {
	self.closed = make(chan bool)
	if _token, ok := ctx.Value(cios.ContextAccessToken).(string); !ok && !check.IsNil(self.refresh) {
		self.token = (*self.refresh)()
	} else {
		self.token = _token
	}
	if self.Connection, err = CreateCiosWsConn(self.isDebug, self.wsUrl, ParseAccessToken(self.token)); err != nil {
		return
	}
	autoRefresh := func() {
	Refresh:
		for {
			self.debug("Registration Refresh Loop")
			select {
			case <-time.After(time.Minute * 55):
				if check.IsNil(self.refresh) {
					break
				}
				self.token = (*self.refresh)()
				self.debug("Auto Refresh")
				if _connection, _err := CreateCiosWsConn(self.isDebug, self.wsUrl, ParseAccessToken(self.token)); _err == nil {
					self.debug("Connection Close: ", self.Connection.Close())
					self.Connection = _connection
					self.debug("Reconnect Websocket")
				}
			case _, _ = <-self.closed:
				self.debug("End Auto Refresh")
				break Refresh
			}
		}
	}
	autoPing := func() {
	Ping:
		for {
			self.debug("Registration Ping Loop")
			select {
			case <-time.After(time.Minute):
				if !check.IsNil(self.Connection) {
					self.debug("Ping")
					if err := self.Connection.WriteMessage(websocket.PingMessage, nil); err != nil {
						self.debug("Ping Err: ", err)
						continue
					}
				}
			case _, _ = <-self.closed:
				self.debug("End Auto Ping")
				break Ping
			}
		}
	}
	go autoRefresh()
	go autoPing()
	return
}
func (self *CiosMessaging) Close() (err error) {
	self.debug("Close")
	defer self.CloseFunc()
	self.closed <- true
	self.closed <- true
	safeCloseChan(self.closed)
	if !check.IsNil(self.Connection) {
		return self.Connection.Close()
	}
	return nil
}

func (self *PubSub) PublishMessage(id string, body interface{}, packerFormat *string, ctx sdkmodel.RequestCtx) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	request := self.ApiClient.PublishSubscribeApi.PublishMessage(self.withHost(ctx)).ChannelId(id).Body(body)
	request.P_packerFormat = packerFormat
	return request.Execute()
}
func (self *PubSub) PublishMessagePackerOnly(id string, body interface{}, ctx sdkmodel.RequestCtx) (*_nethttp.Response, error) {
	return self.PublishMessage(id, &body, nil, ctx)
}
func (self *PubSub) PublishMessageJSON(id string, body cios.PackerFormatJson, ctx sdkmodel.RequestCtx) (*_nethttp.Response, error) {
	return self.PublishMessage(id, &body, convert.StringPtr("json"), ctx)
}

// Deprecated: should not be used
func (self *PubSub) ConnectWebSocket(channelID string, done chan bool, params ConnectWebSocketOptions) (err error) {
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
	if _token == "" {
		_ = self.refresh()
		_token = str(self.token)
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
			_ = self.refresh()
			_connection, _err := self.CreateCIOSWebsocketConnection(wsUrl, ParseAccessToken(str(self.token)))
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
					if !check.IsNil(self.refresh) {
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

// Deprecated: should not be used
func (self *PubSub) CreateMessagingURL(channelID string, mode string, packerFormat *string) string {
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

// Deprecated: should not be used
func (self *PubSub) CreateCIOSWebsocketConnection(url string, authorization string) (connection *websocket.Conn, err error) {
	if self.debug {
		log.Printf("Websocket URL: %s\nAuthorization: %s", url, authorization)
	}
	connection, _, err = (&websocket.Dialer{}).Dial(url, http.Header{"Authorization": []string{authorization}})
	return
}

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
