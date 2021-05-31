package ciossdmock

import sdkmodel "github.com/optim-corp/cios-golang-sdk/model"

type NoImplementAuth struct{}

func (NoImplementAuth) GetAccessTokenByRefreshToken() (sdkmodel.AccessToken, sdkmodel.Scope, sdkmodel.TokenType, sdkmodel.ExpiresIn, error) {
	panic("implement me")
}

func (NoImplementAuth) GetAccessTokenOnClient() (sdkmodel.AccessToken, sdkmodel.Scope, sdkmodel.TokenType, sdkmodel.ExpiresIn, error) {
	panic("implement me")
}

func (NoImplementAuth) GetAccessTokenOnDevice() (sdkmodel.AccessToken, sdkmodel.Scope, sdkmodel.TokenType, sdkmodel.ExpiresIn, error) {
	panic("implement me")
}

func (NoImplementAuth) SetClientSecret(clientSecret string) {
	panic("implement me")
}

func (NoImplementAuth) SetClientId(clientId string) {
	panic("implement me")
}

func (NoImplementAuth) SetRef(ref string) {
	panic("implement me")
}

func (NoImplementAuth) SetAssertion(assertion string) {
	panic("implement me")
}

func (NoImplementAuth) SetDebug(debug bool) {
	panic("implement me")
}

func (NoImplementAuth) SetScope(scope string) {
	panic("implement me")
}
