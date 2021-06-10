package ciossdk

import (
	"fmt"
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

type CiosClient interface {
	fmt.Stringer
	Auth() Auth
	PubSub() PubSub
	Account() Account
	DeviceAssetManagement() DeviceAssetManagement
	DeviceManagement() DeviceManagement
	FileStorage() FileStorage
	Geography() Geography
	License() License
	Contract() Contract
	Video() VideoStreaming
	Debug(debug bool) CiosClient
	RequestScope(scope string) CiosClient
	TokenExp() int64
	SetTokenExp(tokenExp int64)

	_accessToken(accessToken string) CiosClient
}
type ciosClient struct {
	pubSub                PubSub
	account               Account
	deviceAssetManagement DeviceAssetManagement
	deviceManagement      DeviceManagement
	fileStorage           FileStorage
	geography             Geography
	auth                  Auth
	license               License
	contract              Contract
	video                 VideoStreaming
	tokenExp              int64
	cfg                   *cios.Configuration
}

func (self *ciosClient) SetTokenExp(tokenExp int64) {
	self.tokenExp = tokenExp
}

func (self *ciosClient) TokenExp() int64 {
	return self.tokenExp
}

func NewCiosClient(config CiosClientConfig) CiosClient {
	instance := new(ciosClient)
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
					token, _, _, _, err := instance.auth.GetAccessTokenOnClient()
					if err != nil {
						return err
					}
					instance._accessToken(token)
				case sdkmodel.REFRESH_TYPE:
					token, _, _, _, err := instance.auth.GetAccessTokenByRefreshToken()
					if err != nil {
						return err
					}
					instance._accessToken(token)
				case sdkmodel.DEVICE_TYPE:
					token, _, _, _, err := instance.auth.GetAccessTokenOnDevice()
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

	instance.contract = contract
	instance.license = license
	instance.account = account
	instance.deviceManagement = deviceManagement
	instance.deviceAssetManagement = deviceAssetManagement
	instance.fileStorage = fileStorage
	instance.geography = geography
	instance.auth = auth
	instance.pubSub = pubsub
	instance.video = video

	return instance
}

func (self *ciosClient) PubSub() PubSub {
	return self.pubSub
}

func (self *ciosClient) Account() Account {
	return self.account
}

func (self *ciosClient) DeviceAssetManagement() DeviceAssetManagement {
	return self.deviceAssetManagement
}

func (self *ciosClient) DeviceManagement() DeviceManagement {
	return self.deviceManagement
}

func (self *ciosClient) FileStorage() FileStorage {
	return self.fileStorage
}

func (self *ciosClient) Geography() Geography {
	return self.geography
}

func (self *ciosClient) Auth() Auth {
	return self.auth
}

func (self *ciosClient) License() License {
	return self.license
}

func (self *ciosClient) Contract() Contract {
	return self.contract
}

func (self *ciosClient) Video() VideoStreaming {
	return self.video
}

func (self *ciosClient) Debug(debug bool) CiosClient {
	if !check.IsNil(self.cfg) {
		self.cfg.Debug = debug
	}
	self.pubSub.SetDebug(debug)
	return self
}
func (self *ciosClient) _accessToken(accessToken string) CiosClient {
	accessToken = regexp.MustCompile(`^bearer|Bearer| `).ReplaceAllString(accessToken, "")
	token, _, err := new(jwt.Parser).ParseUnverified(accessToken, jwt.MapClaims{})
	if err == nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			self.tokenExp = cnv.MustInt64(claims["exp"])
		}
	}
	self.pubSub.SetToken(accessToken)
	self.video.SetToken(accessToken)
	if !check.IsNil(self.cfg) {
		self.cfg.AddDefaultHeader("Authorization", ciosutil.ParseAccessToken(accessToken))
	}
	return self
}
func (self *ciosClient) RequestScope(scope string) CiosClient {
	self.auth.SetScope(scope)
	return self
}

func (self ciosClient) String() string {
	// TODO: ToString
	return ""
}
