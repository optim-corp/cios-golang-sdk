package ciossdk

import (
	"github.com/optim-corp/cios-golang-sdk/cios"
	"github.com/optim-kazuhiro-seida/go-advance-type/convert"
)

type (
	_instance struct {
		ApiClient *cios.APIClient
		Url       string
		refresh   func() error
	}
	PubSub struct {
		ApiClient *cios.APIClient
		Url       string
		refresh   func() error
		debug     bool
		token     *string
	}
	Auth struct {
		ApiClient    *cios.APIClient
		Url          string
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
