package srvfilestorage

import (
	"context"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

type CiosFileStorage struct {
	ApiClient *cios.APIClient
	Url       string
	Host      string
	withHost  func(context.Context) context.Context
	refresh   func() error
}

func (self *CiosFileStorage) SetWithHost(withHost func(context.Context) context.Context) {
	self.withHost = withHost
}

func (self *CiosFileStorage) SetRefresh(refresh func() error) {
	self.refresh = refresh
}

func NewCiosFileStorage(apiClient *cios.APIClient, url string, withHost func(context.Context) context.Context) *CiosFileStorage {
	return &CiosFileStorage{ApiClient: apiClient, Url: url, withHost: withHost}
}
