package ciossdmock_test

import (
	"testing"

	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
	ciossdmock "github.com/optim-corp/cios-golang-sdk/sdk/mock"
	"github.com/stretchr/testify/assert"
)

func Test_PubSub_Mock(t *testing.T) {
	assert.Implements(t, (*ciossdk.PubSub)(nil), ciossdmock.NoImplementPubsub{})
}
