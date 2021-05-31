package srvaccount

import (
	"context"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

type CiosAccount struct {
	ApiClient *cios.APIClient
	Url       string
	Host      string
	withHost  func(context.Context) context.Context
	refresh   func() error
}

func (self *CiosAccount) SetWithHost(withHost func(context.Context) context.Context) {
	self.withHost = withHost
}

func (self *CiosAccount) SetRefresh(refresh func() error) {
	self.refresh = refresh
}

func NewCiosAccount(apiClient *cios.APIClient, url string, withHost func(context.Context) context.Context) *CiosAccount {
	return &CiosAccount{ApiClient: apiClient, Url: url, withHost: withHost}
}
