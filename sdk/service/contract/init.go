package srvcontract

import (
	"context"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

type CiosContract struct {
	ApiClient *cios.APIClient
	Url       string
	Host      string
	withHost  func(context.Context) context.Context
	refresh   func() error
}

func (self *CiosContract) SetWithHost(withHost func(context.Context) context.Context) {
	self.withHost = withHost
}

func (self *CiosContract) SetRefresh(refresh func() error) {
	self.refresh = refresh
}

func NewCiosContract(apiClient *cios.APIClient, url string, withHost func(context.Context) context.Context) *CiosContract {
	return &CiosContract{ApiClient: apiClient, Url: url, withHost: withHost}
}
