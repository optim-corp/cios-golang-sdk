package ciossdk

import (
	"context"
	"net/http"

	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

var (
	createClient = func(client *http.Client, servers cios.ServerConfigurations, debug bool) *cios.APIClient {
		config := cios.NewConfiguration()
		config.Debug = debug
		config.UserAgent = "Cloud IoT OS Golang SDK"
		config.HTTPClient = client
		config.Servers = servers
		return cios.NewAPIClient(config)
	}
	getWithHostFunc = func(index int) func(ctx context.Context) context.Context {
		return func(ctx context.Context) context.Context {
			if ctx == nil {
				return context.WithValue(context.Background(), cios.ContextServerIndex, index)
			}
			return context.WithValue(ctx, cios.ContextServerIndex, index)
		}
	}
	str = convert.MustStr
)

type (
	CiosClient struct {
		*PubSub
		*Account
		*DeviceAssetManagement
		*DeviceManagement
		*FileStorage
		*Geography
		*Auth
		*License
		*Contract
		tokenExp int64
		cfg      *cios.Configuration
	}
	CiosClientConfig struct {
		AutoRefresh  bool
		Debug        bool
		Urls         sdkmodel.CIOSUrl
		CustomClient *http.Client
		*AuthConfig
		WebSocketConfig
	}
	WebSocketConfig struct {
		ReadTimeout  int64
		WriteTimeout int64
	}
	AuthConfig struct {
		sdkmodel.ClientID
		sdkmodel.ClientSecret
		sdkmodel.RefreshToken
		sdkmodel.Assertion
		sdkmodel.Scope
		_type string
	}
	_instance struct {
		ApiClient *cios.APIClient
		Url       string
		Host      string
		withHost  func(context.Context) context.Context
		refresh   func() error
	}
	PubSub struct {
		_instance
		debug bool
		token *string
	}
	Auth struct {
		_instance
		debug        bool
		scope        string
		assertion    string
		ref          string
		clientId     string
		clientSecret string
	}
	Contract              _instance
	DeviceAssetManagement _instance
	DeviceManagement      _instance
	FileStorage           _instance
	Geography             _instance
	License               _instance
	Account               _instance
)

func ClientAuthConf(clientId, clientSecret, scope string) *AuthConfig {
	return &AuthConfig{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scope:        scope,
		_type:        sdkmodel.CLIENT_TYPE,
	}
}
func RefreshTokenAuth(clientId, clientSecret, refreshToken, scope string) *AuthConfig {
	return &AuthConfig{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RefreshToken: refreshToken,
		Scope:        scope,
		_type:        sdkmodel.REFRESH_TYPE,
	}
}
func DeviceAuthConf(clientId, clientSecret, assertion, scope string) *AuthConfig {
	return &AuthConfig{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Assertion:    assertion,
		Scope:        scope,
		_type:        sdkmodel.DEVICE_TYPE,
	}
}
