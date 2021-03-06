package ciossdk

import (
	"context"
	"net/http"

	"github.com/fcfcqloow/go-advance/check"
	cnv "github.com/fcfcqloow/go-advance/convert"

	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

var (
	createClient = func(client *http.Client, servers cios.ServerConfigurations, debug bool) *cios.APIClient {
		config := cios.NewConfiguration()
		config.Debug = debug
		config.UserAgent = "OPTiM Cloud IoT OS Golang SDK"
		config.HTTPClient = client
		config.Servers = servers
		return cios.NewAPIClient(config)
	}
	getWithHostFunc = func(index int) func(ctx context.Context) context.Context {
		return func(ctx context.Context) context.Context {
			if check.IsNil(ctx) {
				return context.WithValue(context.Background(), cios.ContextServerIndex, index)
			}
			return context.WithValue(ctx, cios.ContextServerIndex, index)
		}
	}
	str = cnv.MustStr
)

type (
	CiosClientConfig struct {
		AutoRefresh     bool
		Debug           bool
		Urls            sdkmodel.CIOSUrl
		CustomClient    *http.Client
		AuthConfig      *AuthConfig
		WebSocketConfig *WebSocketConfig
	}
	WebSocketConfig struct {
		ReadTimeoutMilliSec  int64
		WriteTimeoutMilliSec int64
	}
	AuthConfig struct {
		sdkmodel.ClientID
		sdkmodel.ClientSecret
		sdkmodel.RefreshToken
		sdkmodel.Assertion
		sdkmodel.Scope
		_type string
	}
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
