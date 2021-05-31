package srvdevice

import (
	"context"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

type (
	CiosDeviceManagement struct {
		ApiClient *cios.APIClient
		Url       string
		Host      string
		withHost  func(context.Context) context.Context
		refresh   func() error
	}
	CiosDeviceAssetManagement struct {
		ApiClient *cios.APIClient
		Url       string
		Host      string
		withHost  func(context.Context) context.Context
		refresh   func() error
	}
)

func (self *CiosDeviceAssetManagement) SetWithHost(withHost func(context.Context) context.Context) {
	self.withHost = withHost
}

func (self *CiosDeviceAssetManagement) SetRefresh(refresh func() error) {
	self.refresh = refresh
}

func (self *CiosDeviceManagement) SetWithHost(withHost func(context.Context) context.Context) {
	self.withHost = withHost
}

func (self *CiosDeviceManagement) SetRefresh(refresh func() error) {
	self.refresh = refresh
}

func NewCiosDeviceManagement(apiClient *cios.APIClient, url string, withHost func(context.Context) context.Context) *CiosDeviceManagement {
	return &CiosDeviceManagement{ApiClient: apiClient, Url: url, withHost: withHost}
}

func NewCiosDeviceAssetManagement(apiClient *cios.APIClient, url string, withHost func(context.Context) context.Context) *CiosDeviceAssetManagement {
	return &CiosDeviceAssetManagement{ApiClient: apiClient, Url: url, withHost: withHost}
}
