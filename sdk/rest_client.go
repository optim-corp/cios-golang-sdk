package ciossdk

import (
	"regexp"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fcfcqloow/go-advance/check"
	cnv "github.com/fcfcqloow/go-advance/convert"
	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
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
	contract := &CiosContract{ApiClient: client, Url: config.Urls.ContractUrl, withHost: getWithHostFunc(sdkmodel.CONTRACT_INDEX)}
	license := &CiosLicense{ApiClient: client, Url: config.Urls.LicenseUrl, withHost: getWithHostFunc(sdkmodel.LICENSE_INDEX)}
	account := &CiosAccount{ApiClient: client, Url: config.Urls.AccountsUrl, withHost: getWithHostFunc(sdkmodel.ACCOUNT_INDEX)}
	deviceManagement := &CiosDeviceManagement{ApiClient: client, Url: config.Urls.DeviceManagementUrl, withHost: getWithHostFunc(sdkmodel.DEVICE_INDEX)}
	deviceAssetManagement := &CiosDeviceAssetManagement{ApiClient: client, Url: config.Urls.DeviceAssetManagementUrl, withHost: getWithHostFunc(sdkmodel.DEVICE_ASSET_INDEX)}
	fileStorage := &CiosFileStorage{ApiClient: client, Url: config.Urls.StorageUrl, withHost: getWithHostFunc(sdkmodel.FILE_STORAGE_INDEX)}
	geography := &CiosGeography{ApiClient: client, Url: config.Urls.LocationUrl, withHost: getWithHostFunc(sdkmodel.LOCATION_INDEX)}
	auth := &CiosAuth{
		_instance: _instance{
			ApiClient: client,
			Url:       config.Urls.AuthUrl,
			withHost:  getWithHostFunc(sdkmodel.AUTH_INDEX),
		},
	}
	pubsub := &CiosPubSub{
		_instance: _instance{
			ApiClient: client,
			Url:       config.Urls.MessagingUrl,
			withHost:  getWithHostFunc(sdkmodel.MESSAGING_INDEX),
			refresh:   nil,
		},
	}
	video := &CiosVideoStreaming{
		_instance: _instance{
			ApiClient: client,
			Url:       config.Urls.VideoStreamingUrl,
			withHost:  getWithHostFunc(sdkmodel.VIDEO_STREAMING_INDEX),
		},
	}
	// Instance

	// AuthConfig
	if config.AuthConfig != nil {
		auth.clientId = config.AuthConfig.ClientID
		auth.clientSecret = config.AuthConfig.ClientSecret
		auth.assertion = config.AuthConfig.Assertion
		auth.ref = config.AuthConfig.RefreshToken
		auth.scope = config.AuthConfig.Scope
	}
	// WebSocketConfig

	if !check.IsNil(config.WebSocketConfig) {
		pubsub.wsReadTimeout = config.WebSocketConfig.ReadTimeoutMilliSec
		pubsub.wsWriteTimeout = config.WebSocketConfig.WriteTimeoutMilliSec
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

	video.refresh = refFunc
	account.refresh = refFunc
	license.refresh = refFunc
	contract.refresh = refFunc
	geography.refresh = refFunc
	fileStorage.refresh = refFunc
	deviceManagement.refresh = refFunc
	deviceAssetManagement.refresh = refFunc
	pubsub.refresh = refFunc

	instance.Contract = contract
	instance.License = license
	instance.Account = account
	instance.DeviceManagement = deviceManagement
	instance.DeviceAssetManagement = deviceAssetManagement
	instance.FileStorage = fileStorage
	instance.Geography = geography
	instance.Auth = auth
	instance.PubSub = pubsub
	instance.Video = video

	return instance
}

func (self *CiosClient) Debug(debug bool) *CiosClient {
	if !check.IsNil(self.cfg) {
		self.cfg.Debug = debug
	}
	self.PubSub.setDebug(debug)
	return self
}
func (self *CiosClient) _accessToken(accessToken string) *CiosClient {
	accessToken = regexp.MustCompile(`^bearer|Bearer| `).ReplaceAllString(accessToken, "")
	token, _, err := new(jwt.Parser).ParseUnverified(accessToken, jwt.MapClaims{})
	if err == nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			self.tokenExp = cnv.MustInt64(claims["exp"])
		}
	}
	self.PubSub.setToken(accessToken)
	self.Video.setToken(accessToken)
	if !check.IsNil(self.cfg) {
		self.cfg.AddDefaultHeader("Authorization", ParseAccessToken(accessToken))
	}
	return self
}
func (self *CiosClient) RequestScope(scope string) *CiosClient {
	self.Auth.SetScope(scope)
	return self
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
