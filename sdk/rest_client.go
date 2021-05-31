package ciossdk

import (
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fcfcqloow/go-advance/check"
	cnv "github.com/fcfcqloow/go-advance/convert"
	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
	srvaccount "github.com/optim-corp/cios-golang-sdk/sdk/service/account"
	srvauth "github.com/optim-corp/cios-golang-sdk/sdk/service/authorization"
	srvcontract "github.com/optim-corp/cios-golang-sdk/sdk/service/contract"
	srvdevice "github.com/optim-corp/cios-golang-sdk/sdk/service/device"
	srvfilestorage "github.com/optim-corp/cios-golang-sdk/sdk/service/filestorage"
	srvgeography "github.com/optim-corp/cios-golang-sdk/sdk/service/geography"
	srvlicense "github.com/optim-corp/cios-golang-sdk/sdk/service/license"
	srvpubsub "github.com/optim-corp/cios-golang-sdk/sdk/service/pubsub"
	srvvideo "github.com/optim-corp/cios-golang-sdk/sdk/service/video"
	ciosutil "github.com/optim-corp/cios-golang-sdk/util/cios"
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
	account := srvaccount.NewCiosAccount(client, config.Urls.AccountsUrl, getWithHostFunc(sdkmodel.ACCOUNT_INDEX))
	contract := srvcontract.NewCiosContract(client, config.Urls.ContractUrl, getWithHostFunc(sdkmodel.CONTRACT_INDEX))
	license := srvlicense.NewCiosLicense(client, config.Urls.LicenseUrl, getWithHostFunc(sdkmodel.LICENSE_INDEX))
	deviceManagement := srvdevice.NewCiosDeviceManagement(client, config.Urls.DeviceManagementUrl, getWithHostFunc(sdkmodel.DEVICE_INDEX))
	deviceAssetManagement := srvdevice.NewCiosDeviceAssetManagement(client, config.Urls.DeviceAssetManagementUrl, getWithHostFunc(sdkmodel.DEVICE_ASSET_INDEX))
	fileStorage := srvfilestorage.NewCiosFileStorage(client, config.Urls.StorageUrl, getWithHostFunc(sdkmodel.FILE_STORAGE_INDEX))
	geography := srvgeography.NewCiosGeography(client, config.Urls.LocationUrl, getWithHostFunc(sdkmodel.LOCATION_INDEX))
	auth := srvauth.NewCiosAuth(client, config.Urls.AuthUrl, getWithHostFunc(sdkmodel.AUTH_INDEX))
	pubsub := srvpubsub.NewCiosPubSub(client, config.Urls.MessagingUrl, getWithHostFunc(sdkmodel.MESSAGING_INDEX))
	video := srvvideo.NewCiosVideoStreaming(client, config.Urls.VideoStreamingUrl, getWithHostFunc(sdkmodel.VIDEO_STREAMING_INDEX))

	// AuthConfig
	if config.AuthConfig != nil {
		auth.SetClientId(config.AuthConfig.ClientID)
		auth.SetClientSecret(config.AuthConfig.ClientSecret)
		auth.SetAssertion(config.AuthConfig.Assertion)
		auth.SetRef(config.AuthConfig.RefreshToken)
		auth.SetScope(config.AuthConfig.Scope)
	}
	// WebSocketConfig

	if !check.IsNil(config.WebSocketConfig) {
		pubsub.SetWsReadTimeout(config.WebSocketConfig.ReadTimeoutMilliSec)
		pubsub.SetWsWriteTimeout(config.WebSocketConfig.WriteTimeoutMilliSec)
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

	video.SetRefresh(refFunc)
	account.SetRefresh(refFunc)
	license.SetRefresh(refFunc)
	contract.SetRefresh(refFunc)
	geography.SetRefresh(refFunc)
	fileStorage.SetRefresh(refFunc)
	deviceManagement.SetRefresh(refFunc)
	deviceAssetManagement.SetRefresh(refFunc)
	pubsub.SetRefresh(refFunc)

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
	self.PubSub.SetDebug(debug)
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
	self.PubSub.SetToken(accessToken)
	self.Video.SetToken(accessToken)
	if !check.IsNil(self.cfg) {
		self.cfg.AddDefaultHeader("Authorization", ciosutil.ParseAccessToken(accessToken))
	}
	return self
}
func (self *CiosClient) RequestScope(scope string) *CiosClient {
	self.Auth.SetScope(scope)
	return self
}

func (self CiosClient) String() string {
	// TODO: ToString
	return ""
}
