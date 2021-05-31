package srvvideo

import (
	"context"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

type CiosVideoStreaming struct {
	ApiClient *cios.APIClient
	Url       string
	Host      string
	withHost  func(context.Context) context.Context
	refresh   func() error
	token     *string
}

func (self *CiosVideoStreaming) SetToken(token string) {
	if token == "" {
		self.token = nil
	} else {
		self.token = &token
	}
}
func (self *CiosVideoStreaming) SetRefresh(refresh func() error) {
	self.refresh = refresh
}

func NewCiosVideoStreaming(apiClient *cios.APIClient, url string, withHost func(context.Context) context.Context) *CiosVideoStreaming {
	return &CiosVideoStreaming{ApiClient: apiClient, Url: url, withHost: withHost}
}
