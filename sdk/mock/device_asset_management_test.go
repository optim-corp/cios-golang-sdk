package ciossdk_mock_test

import (
	"testing"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	ciossdk_mock "github.com/optim-corp/cios-golang-sdk/sdk/mock"
	"github.com/stretchr/testify/assert"
)

func TestNoImplementDeviceAssetManagement_CreateEntity(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.CreateEntity(ciosctx.Background(), "", cios.Inventory{})
}

func TestNoImplementDeviceAssetManagement_CreateLifecycle(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.CreateLifecycle(ciosctx.Background(), "", cios.LifeCycleRequest{})
}

func TestNoImplementDeviceAssetManagement_CreateModel(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.CreateModel(ciosctx.Background(), cios.DeviceModelRequest{})
}

func TestNoImplementDeviceAssetManagement_DeleteEntity(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.DeleteEntity(ciosctx.Background(), "")
}

func TestNoImplementDeviceAssetManagement_DeleteLifecycle(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.DeleteLifecycle(ciosctx.Background(), "", "")
}

func TestNoImplementDeviceAssetManagement_DeleteModel(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.DeleteModel(ciosctx.Background(), "")
}

func TestNoImplementDeviceAssetManagement_GetEntities(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetEntities(ciosctx.Background(), cios.ApiGetDeviceEntitiesRequest{})
}

func TestNoImplementDeviceAssetManagement_GetEntitiesAll(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetEntitiesAll(ciosctx.Background(), cios.ApiGetDeviceEntitiesRequest{})
}

func TestNoImplementDeviceAssetManagement_GetEntitiesMapByID(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetEntitiesMapByID(ciosctx.Background(), cios.ApiGetDeviceEntitiesRequest{})
}

func TestNoImplementDeviceAssetManagement_GetEntitiesUnlimited(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetEntitiesUnlimited(ciosctx.Background(), cios.ApiGetDeviceEntitiesRequest{})
}

func TestNoImplementDeviceAssetManagement_GetEntity(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetEntity(ciosctx.Background(), "")
}

func TestNoImplementDeviceAssetManagement_GetLifecycles(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetLifecycles(ciosctx.Background(), "", cios.ApiGetDeviceEntitiesLifecyclesRequest{})
}

func TestNoImplementDeviceAssetManagement_GetLifecyclesAll(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetLifecyclesAll(ciosctx.Background(), "", cios.ApiGetDeviceEntitiesLifecyclesRequest{})
}

func TestNoImplementDeviceAssetManagement_GetLifecyclesUnlimited(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetLifecyclesUnlimited(ciosctx.Background(), "", cios.ApiGetDeviceEntitiesLifecyclesRequest{})
}

func TestNoImplementDeviceAssetManagement_GetLifecyclesUnlimitedByEntities(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetLifecyclesUnlimitedByEntities(ciosctx.Background(), nil, cios.ApiGetDeviceEntitiesLifecyclesRequest{})
}

func TestNoImplementDeviceAssetManagement_GetModel(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetModel(ciosctx.Background(), "")
}

func TestNoImplementDeviceAssetManagement_GetModels(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetModels(ciosctx.Background(), cios.ApiGetDeviceModelsRequest{})
}

func TestNoImplementDeviceAssetManagement_GetModelsAll(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetModelsAll(ciosctx.Background(), cios.ApiGetDeviceModelsRequest{})
}

func TestNoImplementDeviceAssetManagement_GetModelsMapByID(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetModelsMapByID(ciosctx.Background(), cios.ApiGetDeviceModelsRequest{})
}

func TestNoImplementDeviceAssetManagement_GetModelsUnlimited(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceAssetManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetModelsUnlimited(ciosctx.Background(), cios.ApiGetDeviceModelsRequest{})
}
