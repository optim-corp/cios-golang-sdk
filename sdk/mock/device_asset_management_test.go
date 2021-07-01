package ciossdmock_test

import (
	"testing"

	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"

	ciossdk_mock "github.com/optim-corp/cios-golang-sdk/sdk/mock"
	"github.com/stretchr/testify/assert"
)

func TestMockClient_DeviceAssetManagement(t *testing.T) {
	assert.Implements(t, (*ciossdk.DeviceAssetManagement)(nil), ciossdk_mock.NoImplementDeviceAssetManagement{})
}
