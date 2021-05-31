package ciossdk_mock_test

import (
	"testing"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	ciossdk_mock "github.com/optim-corp/cios-golang-sdk/sdk/mock"
	"github.com/stretchr/testify/assert"
)

func TestNoImplementGeography_CreatePoint(t *testing.T) {
	mock := ciossdk_mock.NoImplementGeography{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.CreatePoint(ciosctx.Background(), cios.PointRequest{})
}

func TestNoImplementGeography_DeletePoint(t *testing.T) {
	mock := ciossdk_mock.NoImplementGeography{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.DeletePoint(ciosctx.Background(), "")
}

func TestNoImplementGeography_GetPoints(t *testing.T) {
	mock := ciossdk_mock.NoImplementGeography{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetPoints(ciosctx.Background(), cios.ApiGetPointsRequest{})
}
