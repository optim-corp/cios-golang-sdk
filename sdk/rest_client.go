package ciossdk

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	"github.com/optim-corp/cios-golang-sdk/model"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

var (
	createClient = func(basePath string) *cios.APIClient {
		c := cios.NewAPIClient(cios.NewConfiguration())
		sp := strings.Split(basePath, "://")
		if len(sp) >= 2 {
			c.GetConfig().Scheme = sp[0]
			c.GetConfig().Host = sp[1]
		}
		return c
	}
)

type (
	CiosClient struct {
		PubSub                   *PubSub
		Account                  *Account
		DeviceAssetManagement    *DeviceAssetManagement
		DeviceManagement         *DeviceManagement
		FileStorage              *FileStorage
		Geography                *Geography
		Auth                     *Auth
		License                  *License
		Contract                 *Contract
		ContractUrl              string
		MonitoringUrl            string
		MessagingUrl             string
		LocationUrl              string
		AccountsUrl              string
		StorageUrl               string
		IamUrl                   string
		AuthUrl                  string
		VideoStreamingUrl        string
		DeviceManagementUrl      string
		DeviceAssetManagementUrl string
		LicenseUrl               string
		refreshToken             string
		AutoRefresh              bool
		debug                    bool
		authType                 model.AuthType
		accessToken              string
		clientID                 string
		clientSecret             string
		requestScope             string
	}
	CIOSUrl struct {
		LicenseUrl               string
		ContractUrl              string
		MessagingUrl             string
		LocationUrl              string
		AccountsUrl              string
		StorageUrl               string
		IamUrl                   string
		AuthUrl                  string
		VideoStreamingUrl        string
		DeviceManagementUrl      string
		DeviceMonitoringUrl      string
		DeviceAssetManagementUrl string
	}
	CiosClientConfig struct {
		authType     model.AuthType
		AutoRefresh  bool
		Debug        bool
		Urls         CIOSUrl
		RefreshToken string
		ClientID     string
		ClientSecret string
		RequestScope string
	}
	Auth struct {
		ApiClient *cios.APIClient
		Url       string
		refresh   func() (model.AccessToken, model.Scope, model.TokenType, model.ExpiresIn, error)
		autoR     bool
	}
)

func NewCiosClient(config CiosClientConfig) *CiosClient {
	instance := new(CiosClient)
	if config.ClientID != "" && config.ClientSecret != "" {
		instance.clientID = config.ClientID
		instance.clientSecret = config.ClientSecret
		instance.requestScope = config.RequestScope
		instance.AutoRefresh = config.AutoRefresh
		if config.RefreshToken != "" {
			instance.refreshToken = config.RefreshToken
			instance.authType = model.REFRESH_TYPE
		} else if false {
			instance.authType = model.DEVICE_TYPE
		} else {
			instance.authType = model.CLIENT_TYPE
		}
	} else {
		instance.authType = model.NONE_TYPE
	}

	instance.LicenseUrl = config.Urls.LicenseUrl
	instance.IamUrl = config.Urls.IamUrl
	instance.AuthUrl = config.Urls.AuthUrl
	instance.StorageUrl = config.Urls.StorageUrl
	instance.AccountsUrl = config.Urls.AccountsUrl
	instance.LocationUrl = config.Urls.LocationUrl
	instance.MessagingUrl = config.Urls.MessagingUrl
	instance.MonitoringUrl = config.Urls.DeviceMonitoringUrl
	instance.VideoStreamingUrl = config.Urls.VideoStreamingUrl
	instance.DeviceManagementUrl = config.Urls.DeviceManagementUrl
	instance.DeviceAssetManagementUrl = config.Urls.DeviceAssetManagementUrl
	instance.ContractUrl = config.Urls.ContractUrl

	instance.Contract = &Contract{ApiClient: createClient(config.Urls.ContractUrl), Url: config.Urls.ContractUrl}
	instance.License = &License{ApiClient: createClient(config.Urls.LicenseUrl), Url: config.Urls.LicenseUrl}
	instance.Auth = &Auth{ApiClient: createClient(config.Urls.AuthUrl), Url: config.Urls.AuthUrl}
	instance.PubSub = &PubSub{ApiClient: createClient(config.Urls.MessagingUrl), Url: config.Urls.MessagingUrl}
	instance.Account = &Account{ApiClient: createClient(config.Urls.AccountsUrl), Url: config.Urls.AccountsUrl}
	instance.DeviceManagement = &DeviceManagement{ApiClient: createClient(config.Urls.DeviceManagementUrl), Url: config.Urls.DeviceManagementUrl}
	instance.DeviceAssetManagement = &DeviceAssetManagement{ApiClient: createClient(config.Urls.DeviceAssetManagementUrl), Url: config.Urls.DeviceAssetManagementUrl}
	instance.FileStorage = &FileStorage{ApiClient: createClient(config.Urls.StorageUrl), Url: config.Urls.StorageUrl}
	instance.Geography = &Geography{ApiClient: createClient(config.Urls.LocationUrl), Url: config.Urls.LocationUrl}
	instance.Auth.autoR = instance.AutoRefresh
	instance.PubSub.autoR = instance.AutoRefresh
	instance.Account.autoR = instance.AutoRefresh
	instance.Geography.autoR = instance.AutoRefresh
	instance.FileStorage.autoR = instance.AutoRefresh
	instance.DeviceManagement.autoR = instance.AutoRefresh
	instance.DeviceAssetManagement.autoR = instance.AutoRefresh
	instance.License.autoR = instance.AutoRefresh
	instance.Contract.autoR = instance.AutoRefresh
	if instance.AutoRefresh {
		refFunc := func() (model.AccessToken, model.Scope, model.TokenType, model.ExpiresIn, error) {
			if instance.debug {
				log.Printf("%s", "Refresh AccessToken.")
			}
			if !instance.AutoRefresh {
				return "", "", "", 0, fmt.Errorf("%s", "No AutoRefresh")
			}
			instance.Debug(false)
			defer func() { instance.Debug(instance.debug) }()
			switch instance.authType {
			case model.CLIENT_TYPE:
				responseData := struct {
					AccessToken string `json:"access_token"`
					TokenType   string `json:"token_type"`
					ExpiresIn   int    `json:"expires_in"`
					Scope       string `json:"scope"`
				}{}
				values := url.Values{}
				values.Add("grant_type", "client_credentials")
				values.Add("client_id", instance.clientID)
				values.Add("client_secret", instance.clientSecret)
				values.Add("scope", instance.requestScope)
				resp, _ := http.Post(
					instance.AuthUrl+"/connect/token",
					"application/x-www-form-urlencoded",
					strings.NewReader(values.Encode()),
				)
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return "", "", "", 0, err
				}

				err = convert.UnMarshalJson(body, &responseData)
				if err != nil {
					return "", "", "", 0, err
				}
				instance._accessToken(responseData.AccessToken)
				return responseData.AccessToken, responseData.Scope, responseData.TokenType, responseData.ExpiresIn, nil
			case model.REFRESH_TYPE:
				response, _, err := instance.Auth.ApiClient.AuthApi.RefreshToken(context.Background()).
					GrantType("refresh_token").RefreshToken(instance.refreshToken).
					ClientId(instance.clientID).ClientSecret(instance.clientSecret).
					Scope(instance.requestScope).Execute()
				if err != nil {
					return "", "", "", 0, err
				}
				instance._accessToken(response.AccessToken)
				return response.AccessToken, response.Scope, response.TokenType, int(response.ExpiresIn), nil
			case model.DEVICE_TYPE:
				return "", "", "", 0, fmt.Errorf("%s", "No Implement Device Auth Type")
			default:
				return "", "", "", 0, fmt.Errorf("%s", "No AccessToken")
			}
		}
		instance.Auth.refresh = refFunc
		instance.PubSub.refresh = refFunc
		instance.Account.refresh = refFunc
		instance.Geography.refresh = refFunc
		instance.FileStorage.refresh = refFunc
		instance.DeviceManagement.refresh = refFunc
		instance.DeviceAssetManagement.refresh = refFunc
		instance.License.refresh = refFunc
		instance.Contract.refresh = refFunc
		select {
		case <-time.After(3 * time.Second):
		default:
			// TODO: 考える
			instance.RefreshAccessToken()
		}
	} else {
		refFunc := func() (model.AccessToken, model.Scope, model.TokenType, model.ExpiresIn, error){
			return "", "", "", 0, errors.New("no refresh client")
		}
		instance.Auth.refresh = refFunc
		instance.PubSub.refresh = refFunc
		instance.Account.refresh = refFunc
		instance.Geography.refresh = refFunc
		instance.FileStorage.refresh = refFunc
		instance.DeviceManagement.refresh = refFunc
		instance.DeviceAssetManagement.refresh = refFunc
		instance.License.refresh = refFunc
		instance.Contract.refresh = refFunc
	}
	instance.Debug(config.Debug)

	return instance
}

func (self *CiosClient) Debug(debug bool) *CiosClient {
	self.Auth.ApiClient.GetConfig().Debug = false
	self.PubSub.ApiClient.GetConfig().Debug = debug
	self.DeviceManagement.ApiClient.GetConfig().Debug = debug
	self.Account.ApiClient.GetConfig().Debug = debug
	self.Geography.ApiClient.GetConfig().Debug = debug
	self.DeviceAssetManagement.ApiClient.GetConfig().Debug = debug
	self.FileStorage.ApiClient.GetConfig().Debug = debug
	self.PubSub.debug = debug
	return self
}
func (self *CiosClient) _accessToken(accessToken string) *CiosClient {
	accessToken = "Bearer " + regexp.MustCompile(`^bearer|Bearer| `).ReplaceAllString(accessToken, "")
	self.accessToken = accessToken
	self.Auth.ApiClient.GetConfig().AddDefaultHeader("Authorization", self.accessToken)
	self.PubSub.ApiClient.GetConfig().AddDefaultHeader("Authorization", self.accessToken)
	self.DeviceManagement.ApiClient.GetConfig().AddDefaultHeader("Authorization", self.accessToken)
	self.Account.ApiClient.GetConfig().AddDefaultHeader("Authorization", self.accessToken)
	self.Geography.ApiClient.GetConfig().AddDefaultHeader("Authorization", self.accessToken)
	self.DeviceAssetManagement.ApiClient.GetConfig().AddDefaultHeader("Authorization", self.accessToken)
	self.FileStorage.ApiClient.GetConfig().AddDefaultHeader("Authorization", self.accessToken)
	return self
}
func (self *CiosClient) RequestScope(scope string) *CiosClient {
	self.requestScope = scope
	return self
}
func (self *CiosClient) RefreshAccessToken() (model.AccessToken, model.Scope, model.TokenType, model.ExpiresIn, error) {
	if self.debug {
		log.Printf("%s", "Refresh AccessToken.")
	}
	if !self.AutoRefresh {
		return "", "", "", 0, fmt.Errorf("%s", "No AutoRefresh")
	}
	self.Debug(false)
	defer func() { self.Debug(self.debug) }()
	switch self.authType {
	case model.CLIENT_TYPE:
		responseData := struct {
			AccessToken string `json:"access_token"`
			TokenType   string `json:"token_type"`
			ExpiresIn   int    `json:"expires_in"`
			Scope       string `json:"scope"`
		}{}
		values := url.Values{}
		values.Add("grant_type", "client_credentials")
		values.Add("client_id", self.clientID)
		values.Add("client_secret", self.clientSecret)
		values.Add("scope", self.requestScope)
		resp, _ := http.Post(
			self.AuthUrl+"/connect/token",
			"application/x-www-form-urlencoded",
			strings.NewReader(values.Encode()),
		)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", "", "", 0, err
		}

		err = convert.UnMarshalJson(body, &responseData)
		if err != nil {
			return "", "", "", 0, err
		}
		self._accessToken(responseData.AccessToken)
		return responseData.AccessToken, responseData.Scope, responseData.TokenType, responseData.ExpiresIn, nil
	case model.REFRESH_TYPE:
		response, _, err := self.Auth.ApiClient.AuthApi.RefreshToken(context.Background()).
			GrantType("refresh_token").RefreshToken(self.refreshToken).
			ClientId(self.clientID).ClientSecret(self.clientSecret).
			Scope(self.requestScope).Execute()
		if err != nil {
			return "", "", "", 0, err
		}
		self._accessToken(response.AccessToken)
		return response.AccessToken, response.Scope, response.TokenType, int(response.ExpiresIn), nil
	case model.DEVICE_TYPE:
		return "", "", "", 0, fmt.Errorf("%s", "No Implement Device Auth Type")
	default:
		return "", "", "", 0, fmt.Errorf("%s", "No AccessToken")
	}
}

func MakeRequestCtx(token string) model.RequestCtx {
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
