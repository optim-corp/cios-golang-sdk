package ciossdk

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	cnv "github.com/fcfcqloow/go-advance/convert"

	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

type CiosAuth struct {
	_instance
	debug        bool
	scope        string
	assertion    string
	ref          string
	clientId     string
	clientSecret string
}

func (self *CiosAuth) SetScope(scope string) {
	self.scope = scope
}

func (self *CiosAuth) SetClientSecret(clientSecret string) {
	self.clientSecret = clientSecret
}

func (self *CiosAuth) SetClientId(clientId string) {
	self.clientId = clientId
}

func (self *CiosAuth) SetRef(ref string) {
	self.ref = ref
}

func (self *CiosAuth) SetAssertion(assertion string) {
	self.assertion = assertion
}

func (self *CiosAuth) SetDebug(debug bool) {
	self.debug = debug
}

func (self *CiosAuth) GetAccessTokenByRefreshToken() (sdkmodel.AccessToken, sdkmodel.Scope, sdkmodel.TokenType, sdkmodel.ExpiresIn, error) {
	if self.debug {
		log.Printf("%s", "Refresh AccessToken.")
	}
	debug := self.ApiClient.GetConfig().Debug
	self.ApiClient.GetConfig().Debug = false
	response, _, err := self.ApiClient.AuthApi.
		RefreshToken(self.withHost(context.Background())).
		GrantType("refresh_token").
		RefreshToken(self.ref).
		ClientId(self.clientId).
		ClientSecret(self.clientSecret).
		Scope(self.scope).
		Execute()
	if err != nil {
		return "", "", "", 0, err
	}
	self.ApiClient.GetConfig().Debug = debug
	return response.AccessToken, response.Scope, response.TokenType, int(response.ExpiresIn), nil

}

func (self *CiosAuth) GetAccessTokenOnClient() (sdkmodel.AccessToken, sdkmodel.Scope, sdkmodel.TokenType, sdkmodel.ExpiresIn, error) {
	if self.debug {
		log.Printf("%s", "Refresh AccessToken.")
	}
	responseData := struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
	}{}
	values := url.Values{}
	values.Add("grant_type", "client_credentials")
	values.Add("client_id", self.clientId)
	values.Add("client_secret", self.clientSecret)
	values.Add("scope", self.scope)
	resp, _ := http.Post(
		self.Url+"/connect/token",
		"application/x-www-form-urlencoded",
		strings.NewReader(values.Encode()),
	)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", 0, err
	}

	err = cnv.UnMarshalJson(body, &responseData)
	if err != nil {
		return "", "", "", 0, err
	}
	return responseData.AccessToken, responseData.Scope, responseData.TokenType, responseData.ExpiresIn, nil
}

func (self *CiosAuth) GetAccessTokenOnDevice() (sdkmodel.AccessToken, sdkmodel.Scope, sdkmodel.TokenType, sdkmodel.ExpiresIn, error) {
	if self.debug {
		log.Printf("%s", "Refresh AccessToken.")
	}
	responseBody := struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
	}{}
	values := url.Values{
		"grant_type":    []string{"urn:ietf:params:oauth:grant-type:jwt-bearer"},
		"client_id":     []string{self.clientId},
		"client_secret": []string{self.clientSecret},
		"scope":         []string{self.scope},
	}

	response, err := http.Post(self.Url+"/connect/token", "application/x-www-form-urlencoded", strings.NewReader(values.Encode()))
	if err != nil {
		return "", "", "", 0, err
	}
	defer response.Body.Close()
	if err := cnv.UnMarshalJson(response.Body, &responseBody); err != nil {
		return "", "", "", 0, err
	}
	return responseBody.AccessToken, responseBody.Scope, responseBody.TokenType, responseBody.ExpiresIn, nil
}
