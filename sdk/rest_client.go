package ciossdk

import (
	"context"
	"net/http"
	"regexp"
	"strings"
	"time"

	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"

	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	"github.com/dgrijalva/jwt-go"

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
)

const (
	AUTH_INDEX = iota
	ACCOUNT_INDEX
	FILE_STORAGE_INDEX
	MESSAGING_INDEX
	DEVICE_INDEX
	DEVICE_ASSET_INDEX
	CONTRACT_INDEX
	LICENSE_INDEX
	LOCATION_INDEX
)

type (
	CiosClient struct {
		PubSub                *PubSub
		Account               *Account
		DeviceAssetManagement *DeviceAssetManagement
		DeviceManagement      *DeviceManagement
		FileStorage           *FileStorage
		Geography             *Geography
		Auth                  *Auth
		License               *License
		Contract              *Contract
		tokenExp              int64
	}
	CiosClientConfig struct {
		AutoRefresh  bool
		Debug        bool
		Urls         sdkmodel.CIOSUrl
		AuthConfig   *AuthConfig
		CustomClient *http.Client
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
func NewCiosClient(config CiosClientConfig) *CiosClient {
	instance := new(CiosClient)
	client := createClient(config.CustomClient, cios.ServerConfigurations{
		{URL: config.Urls.AuthUrl},
		{URL: config.Urls.AccountsUrl},
		{URL: config.Urls.StorageUrl},
		{URL: config.Urls.MessagingUrl},
		{URL: config.Urls.DeviceManagementUrl},
		{URL: config.Urls.DeviceAssetManagementUrl},
		{URL: config.Urls.ContractUrl},
		{URL: config.Urls.LicenseUrl},
		{URL: config.Urls.LocationUrl},
	}, config.Debug)

	// Instance
	instance.Contract = &Contract{ApiClient: client, Url: config.Urls.ContractUrl, withHost: getWithHostFunc(CONTRACT_INDEX)}
	instance.License = &License{ApiClient: client, Url: config.Urls.LicenseUrl, withHost: getWithHostFunc(LICENSE_INDEX)}
	instance.Account = &Account{ApiClient: client, Url: config.Urls.AccountsUrl, withHost: getWithHostFunc(ACCOUNT_INDEX)}
	instance.DeviceManagement = &DeviceManagement{ApiClient: client, Url: config.Urls.DeviceManagementUrl, withHost: getWithHostFunc(DEVICE_INDEX)}
	instance.DeviceAssetManagement = &DeviceAssetManagement{ApiClient: client, Url: config.Urls.DeviceAssetManagementUrl, withHost: getWithHostFunc(DEVICE_ASSET_INDEX)}
	instance.FileStorage = &FileStorage{ApiClient: client, Url: config.Urls.StorageUrl, withHost: getWithHostFunc(FILE_STORAGE_INDEX)}
	instance.Geography = &Geography{ApiClient: client, Url: config.Urls.LocationUrl, withHost: getWithHostFunc(LOCATION_INDEX)}
	instance.Auth = &Auth{}
	instance.Auth.ApiClient = client
	instance.Auth.Url = config.Urls.AuthUrl
	instance.Auth.withHost = getWithHostFunc(AUTH_INDEX)
	instance.PubSub = &PubSub{}
	instance.PubSub.ApiClient = client
	instance.PubSub.Url = config.Urls.MessagingUrl
	instance.PubSub.withHost = getWithHostFunc(MESSAGING_INDEX)

	// AuthConfig
	if config.AuthConfig != nil {
		instance.Auth.clientId = config.AuthConfig.ClientID
		instance.Auth.clientSecret = config.AuthConfig.ClientSecret
		instance.Auth.assertion = config.AuthConfig.Assertion
		instance.Auth.ref = config.AuthConfig.RefreshToken
		instance.Auth.scope = config.AuthConfig.Scope
	}

	refFunc := func() error {
		if config.AutoRefresh && config.AuthConfig != nil {
			if instance.tokenExp == 0 || instance.tokenExp-60 <= time.Now().Unix() {
				switch config.AuthConfig._type {
				case sdkmodel.CLIENT_TYPE:
					token, _, _, _, err := instance.Auth.GetAccessTokenOnClient()
					if err != nil {
						return err
					}
					instance._accessToken(token)
				case sdkmodel.REFRESH_TYPE:
					token, _, _, _, err := instance.Auth.GetAccessTokenByRefreshToken()
					if err != nil {
						return err
					}
					instance._accessToken(token)
				case sdkmodel.DEVICE_TYPE:
					token, _, _, _, err := instance.Auth.GetAccessTokenOnDevice()
					if err != nil {
						return err
					}
					instance._accessToken(token)
				default:
				}
			}
		}
		return nil
	}
	instance.PubSub.refresh = refFunc
	instance.Account.refresh = refFunc
	instance.License.refresh = refFunc
	instance.Contract.refresh = refFunc
	instance.Geography.refresh = refFunc
	instance.FileStorage.refresh = refFunc
	instance.DeviceManagement.refresh = refFunc
	instance.DeviceAssetManagement.refresh = refFunc
	return instance
}

func (self *CiosClient) Debug(debug bool) *CiosClient {
	self.PubSub.ApiClient.GetConfig().Debug = debug
	self.PubSub.debug = debug
	return self
}
func (self *CiosClient) _accessToken(accessToken string) *CiosClient {
	accessToken = regexp.MustCompile(`^bearer|Bearer| `).ReplaceAllString(accessToken, "")
	token, _, err := new(jwt.Parser).ParseUnverified(accessToken, jwt.MapClaims{})
	if err == nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			self.tokenExp = convert.MustInt64(claims["exp"])
		}
	}
	self.PubSub.token = &accessToken
	self.Auth.ApiClient.GetConfig().AddDefaultHeader("Authorization", ParseAccessToken(accessToken))
	return self
}
func (self *CiosClient) RequestScope(scope string) *CiosClient {
	self.Auth.scope = scope
	return self
}

func MakeRequestCtx(token string) sdkmodel.RequestCtx {
	if token == "" {
		return context.Background()
	}
	return context.WithValue(context.Background(), cios.ContextAccessToken, regexp.MustCompile(`^bearer|Bearer| `).ReplaceAllString(token, ""))
}
func ParseAccessToken(accessToken string) string {
	if !strings.Contains(accessToken, "Bearer ") {
		return "Bearer " + accessToken
	} else {
		return accessToken
	}
}

func (self CiosClient) String() string {
	// TODO: ToString
	return ""
}
