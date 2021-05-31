package ciossdmock_test

import (
	"testing"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	ciossdk_mock "github.com/optim-corp/cios-golang-sdk/sdk/mock"
	"github.com/stretchr/testify/assert"
)

func TestNoImplementContract_GetContracts(t *testing.T) {
	mock := ciossdk_mock.NoImplementContract{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetContracts(ciosctx.Background(), cios.ApiGetContractsRequest{})
}

func TestNoImplementContract_GetContractsAll(t *testing.T) {
	mock := ciossdk_mock.NoImplementContract{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetContractsAll(ciosctx.Background(), cios.ApiGetContractsRequest{})
}

func TestNoImplementContract_GetContractsUnlimited(t *testing.T) {
	mock := ciossdk_mock.NoImplementContract{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetContractsUnlimited(ciosctx.Background(), cios.ApiGetContractsRequest{})
}
