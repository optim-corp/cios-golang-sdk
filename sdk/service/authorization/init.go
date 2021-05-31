package srvauth

import (
	"context"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

type CiosAuth struct {
	ApiClient    *cios.APIClient
	Url          string
	Host         string
	withHost     func(context.Context) context.Context
	refresh      func() error
	debug        bool
	scope        string
	assertion    string
	ref          string
	clientId     string
	clientSecret string
}

func NewCiosAuth(apiClient *cios.APIClient, url string, withHost func(context.Context) context.Context) *CiosAuth {
	return &CiosAuth{ApiClient: apiClient, Url: url, withHost: withHost}
}

func (self *CiosAuth) SetScope(scope string) {
	self.scope = scope
}

func (self *CiosAuth) SetClientSecret(clientSecret string) {
	self.clientSecret = clientSecret
}

func (self *CiosAuth) SetClientId(clientId string) {
	self.clientId = clientId
}

func (self *CiosAuth) SetRef(ref string) {
	self.ref = ref
}

func (self *CiosAuth) SetAssertion(assertion string) {
	self.assertion = assertion
}

func (self *CiosAuth) SetDebug(debug bool) {
	self.debug = debug
}
