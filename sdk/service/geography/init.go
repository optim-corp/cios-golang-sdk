package srvgeography

import (
	"context"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

type CiosGeography struct {
	ApiClient *cios.APIClient
	Url       string
	Host      string
	withHost  func(context.Context) context.Context
	refresh   func() error
}

func (self *CiosGeography) SetWithHost(withHost func(context.Context) context.Context) {
	self.withHost = withHost
}

func (self *CiosGeography) SetRefresh(refresh func() error) {
	self.refresh = refresh
}

func NewCiosGeography(apiClient *cios.APIClient, url string, withHost func(context.Context) context.Context) *CiosGeography {
	return &CiosGeography{ApiClient: apiClient, Url: url, withHost: withHost}
}
