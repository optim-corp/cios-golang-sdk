package srvpubsub

import (
	"context"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

type CiosPubSub struct {
	ApiClient      *cios.APIClient
	Url            string
	Host           string
	withHost       func(context.Context) context.Context
	refresh        func() error
	debug          bool
	token          *string
	wsReadTimeout  int64
	wsWriteTimeout int64
}

func (self *CiosPubSub) SetRefresh(refresh func() error) {
	self.refresh = refresh
}

func (self *CiosPubSub) SetWsReadTimeout(wsReadTimeout int64) {
	self.wsReadTimeout = wsReadTimeout
}

func (self *CiosPubSub) SetWsWriteTimeout(wsWriteTimeout int64) {
	self.wsWriteTimeout = wsWriteTimeout
}

func NewCiosPubSub(apiClient *cios.APIClient, url string, withHost func(context.Context) context.Context) *CiosPubSub {
	return &CiosPubSub{ApiClient: apiClient, Url: url, withHost: withHost}
}

func (self *CiosPubSub) SetDebug(isDebug bool) {
	self.debug = isDebug
}

func (self *CiosPubSub) SetToken(token string) {
	if token == "" {
		self.token = nil
	} else {
		self.token = &token
	}
}
