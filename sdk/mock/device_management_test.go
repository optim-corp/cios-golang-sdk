package ciossdk_mock_test

import (
	"testing"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	ciossdk_mock "github.com/optim-corp/cios-golang-sdk/sdk/mock"
	"github.com/stretchr/testify/assert"
)

func TestNoImplementDeviceManagement_CreateDevice(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.CreateDevice(ciosctx.Background(), cios.DeviceInfo{})
}

func TestNoImplementDeviceManagement_CreatePolicy(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.CreatePolicy(ciosctx.Background(), "")
}

func TestNoImplementDeviceManagement_DeleteDevice(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.DeleteDevice(ciosctx.Background(), "")
}

func TestNoImplementDeviceManagement_DeletePolicy(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.DeletePolicy(ciosctx.Background(), "")
}

func TestNoImplementDeviceManagement_GetDevice(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetDevice(ciosctx.Background(), "", nil, nil)
}

func TestNoImplementDeviceManagement_GetDeviceInventory(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetDeviceInventory(ciosctx.Background(), "")
}

func TestNoImplementDeviceManagement_GetDevices(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetDevices(ciosctx.Background(), cios.ApiGetDevicesRequest{})
}

func TestNoImplementDeviceManagement_GetDevicesAll(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetDevicesAll(ciosctx.Background(), cios.ApiGetDevicesRequest{})
}

func TestNoImplementDeviceManagement_GetDevicesUnlimited(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetDevicesUnlimited(ciosctx.Background(), cios.ApiGetDevicesRequest{})
}

func TestNoImplementDeviceManagement_GetMonitoring(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetMonitoring(ciosctx.Background(), "")
}

func TestNoImplementDeviceManagement_GetMonitoringLatestList(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetMonitoringLatestList(ciosctx.Background(), nil)
}

func TestNoImplementDeviceManagement_GetPolicies(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetPolicies(ciosctx.Background(), cios.ApiGetDevicePoliciesRequest{})
}

func TestNoImplementDeviceManagement_GetPoliciesAll(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetPoliciesAll(ciosctx.Background(), cios.ApiGetDevicePoliciesRequest{})
}

func TestNoImplementDeviceManagement_GetPoliciesUnlimited(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetPoliciesUnlimited(ciosctx.Background(), cios.ApiGetDevicePoliciesRequest{})
}

func TestNoImplementDeviceManagement_UpdateDevice(t *testing.T) {
	mock := ciossdk_mock.NoImplementDeviceManagement{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.UpdateDevice(ciosctx.Background(), "", cios.DeviceUpdateRequest{})
}
