package ciossdmock_test

import (
	"testing"

	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
	ciossdk_mock "github.com/optim-corp/cios-golang-sdk/sdk/mock"
	"github.com/stretchr/testify/assert"
)

func TestMockClient_DeviceManagement(t *testing.T) {
	assert.Implements(t, (*ciossdk.DeviceManagement)(nil), ciossdk_mock.NoImplementDeviceManagement{})
}
