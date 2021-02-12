package ciossdk

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
	"github.com/optim-kazuhiro-seida/go-advance-type/check"
	"github.com/optim-kazuhiro-seida/go-advance-type/convert"
)

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
		{URL: config.Urls.VideoStreamingUrl},
	}, config.Debug)

	instance.cfg = client.GetConfig()
	// Instance
	instance.Contract = &Contract{ApiClient: client, Url: config.Urls.ContractUrl, withHost: getWithHostFunc(sdkmodel.CONTRACT_INDEX)}
	instance.License = &License{ApiClient: client, Url: config.Urls.LicenseUrl, withHost: getWithHostFunc(sdkmodel.LICENSE_INDEX)}
	instance.Account = &Account{ApiClient: client, Url: config.Urls.AccountsUrl, withHost: getWithHostFunc(sdkmodel.ACCOUNT_INDEX)}
	instance.DeviceManagement = &DeviceManagement{ApiClient: client, Url: config.Urls.DeviceManagementUrl, withHost: getWithHostFunc(sdkmodel.DEVICE_INDEX)}
	instance.DeviceAssetManagement = &DeviceAssetManagement{ApiClient: client, Url: config.Urls.DeviceAssetManagementUrl, withHost: getWithHostFunc(sdkmodel.DEVICE_ASSET_INDEX)}
	instance.FileStorage = &FileStorage{ApiClient: client, Url: config.Urls.StorageUrl, withHost: getWithHostFunc(sdkmodel.FILE_STORAGE_INDEX)}
	instance.Geography = &Geography{ApiClient: client, Url: config.Urls.LocationUrl, withHost: getWithHostFunc(sdkmodel.LOCATION_INDEX)}
	instance.Auth = &Auth{}
	instance.Auth.ApiClient = client
	instance.Auth.Url = config.Urls.AuthUrl
	instance.Auth.withHost = getWithHostFunc(sdkmodel.AUTH_INDEX)
	instance.PubSub = &PubSub{}
	instance.PubSub.ApiClient = client
	instance.PubSub.Url = config.Urls.MessagingUrl
	instance.PubSub.withHost = getWithHostFunc(sdkmodel.MESSAGING_INDEX)
	instance.Video = &VideoStreaming{}
	instance.Video.ApiClient = client
	instance.Video.Url = config.Urls.VideoStreamingUrl
	instance.Video.withHost = getWithHostFunc(sdkmodel.VIDEO_STREAMING_INDEX)

	// AuthConfig
	if config.AuthConfig != nil {
		instance.Auth.clientId = config.AuthConfig.ClientID
		instance.Auth.clientSecret = config.AuthConfig.ClientSecret
		instance.Auth.assertion = config.AuthConfig.Assertion
		instance.Auth.ref = config.AuthConfig.RefreshToken
		instance.Auth.scope = config.AuthConfig.Scope
	}
	// WebSocketConfig

	if !check.IsNil(config.WebSocketConfig) {
		instance.PubSub.wsReadTimeout = config.WebSocketConfig.ReadTimeoutMilliSec
		instance.PubSub.wsWriteTimeout = config.WebSocketConfig.WriteTimeoutMilliSec
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
	if !check.IsNil(self.cfg) {
		self.cfg.Debug = debug
	}
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
	if !check.IsNil(self.cfg) {
		self.cfg.AddDefaultHeader("Authorization", ParseAccessToken(accessToken))
	}
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

func GetTokenFromCtx(ctx sdkmodel.RequestCtx) *string {
	_token, ok := ctx.Value(cios.ContextAccessToken).(string)
	if ok {
		return &_token
	}
	return nil
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
