package ciossdk

import (
	"github.com/optim-corp/cios-golang-sdk/cios"
	"github.com/optim-corp/cios-golang-sdk/model"
)

type (
	_instance struct {
		ApiClient *cios.APIClient
		Url       string
		refresh   *func() (model.AccessToken, model.Scope, model.TokenType, model.ExpiresIn, error)
	}
	PubSub struct {
		ApiClient *cios.APIClient
		Url       string
		refresh   *func() (model.AccessToken, model.Scope, model.TokenType, model.ExpiresIn, error)
		debug     bool
	}
	Auth struct {
		ApiClient    *cios.APIClient
		Url          string
		debug        bool
		scope        string
		ref          string
		clientId     string
		clientSecret string
	}
	Account               _instance
	Contract              _instance
	DeviceAssetManagement _instance
	DeviceManagement      _instance
	FileStorage           _instance
	Geography             _instance
	License               _instance
)
