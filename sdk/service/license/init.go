package srvlicense

import (
	"context"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

type CiosLicense struct {
	ApiClient *cios.APIClient
	Url       string
	Host      string
	withHost  func(context.Context) context.Context
	refresh   func() error
}

func (self *CiosLicense) SetWithHost(withHost func(context.Context) context.Context) {
	self.withHost = withHost
}

func (self *CiosLicense) SetRefresh(refresh func() error) {
	self.refresh = refresh
}

func NewCiosLicense(apiClient *cios.APIClient, host string, withHost func(context.Context) context.Context) *CiosLicense {
	return &CiosLicense{ApiClient: apiClient, Host: host, withHost: withHost}
}
