package ciossdk

import (
	"context"

	"github.com/optim-corp/cios-golang-sdk/cios"
	"github.com/optim-kazuhiro-seida/go-advance-type/convert"
)

type (
	_instance struct {
		ApiClient *cios.APIClient
		Url       string
		Host      string
		withHost  func(context.Context) context.Context
		refresh   func() error
	}
	PubSub struct {
		_instance
		debug bool
		token *string
	}
	Auth struct {
		_instance
		debug        bool
		scope        string
		assertion    string
		ref          string
		clientId     string
		clientSecret string
	}
	Contract              _instance
	DeviceAssetManagement _instance
	DeviceManagement      _instance
	FileStorage           _instance
	Geography             _instance
	License               _instance
	Account               _instance
)

var (
	str = convert.MustStr
)
