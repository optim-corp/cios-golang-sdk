package ciossdk_mock_test

import (
	"testing"

	"github.com/optim-corp/cios-golang-sdk/cios"

	"github.com/stretchr/testify/assert"

	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	ciossdk_mock "github.com/optim-corp/cios-golang-sdk/sdk/mock"
)

func TestNoImplementLicense_GetLicenses(t *testing.T) {
	mock := ciossdk_mock.NoImplementLicense{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetLicenses(ciosctx.Background(), cios.ApiGetMyLicensesRequest{})
}
