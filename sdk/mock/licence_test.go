package ciossdmock_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
	ciossdk_mock "github.com/optim-corp/cios-golang-sdk/sdk/mock"
)

func TestMockClient_License(t *testing.T) {
	assert.Implements(t, (*ciossdk.License)(nil), ciossdk_mock.NoImplementLicense{})
}
