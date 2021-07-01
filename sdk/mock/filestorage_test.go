package ciossdmock_test

import (
	"testing"

	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
	ciossdk_mock "github.com/optim-corp/cios-golang-sdk/sdk/mock"
	"github.com/stretchr/testify/assert"
)

func TestMockClient_FileStorage(t *testing.T) {
	assert.Implements(t, (*ciossdk.FileStorage)(nil), ciossdk_mock.NoImplementFileStorage{})
}
