package ciossdk_mock_test

import (
	"testing"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	ciossdk_mock "github.com/optim-corp/cios-golang-sdk/sdk/mock"
	"github.com/stretchr/testify/assert"
)

func TestNoImplementAccount_GetGroups(t *testing.T) {
	mock := ciossdk_mock.NoImplementAccount{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetGroups(ciosctx.Background(), cios.ApiGetGroupsRequest{})
}
func TestNoImplementAccount_GetGroupsAll(t *testing.T) {
	mock := ciossdk_mock.NoImplementAccount{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetGroupsAll(ciosctx.Background(), cios.ApiGetGroupsRequest{})
}
func TestNoImplementAccount_GetGroupsUnlimited(t *testing.T) {
	mock := ciossdk_mock.NoImplementAccount{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetGroupsUnlimited(ciosctx.Background(), cios.ApiGetGroupsRequest{})
}
func TestNoImplementAccount_GetResourceOwnerByGroupId(t *testing.T) {
	mock := ciossdk_mock.NoImplementAccount{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetResourceOwnerByGroupId(ciosctx.Background(), "")
}
func TestNoImplementAccount_GetResourceOwnersMapByID(t *testing.T) {
	mock := ciossdk_mock.NoImplementAccount{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetResourceOwnersMapByID(ciosctx.Background())
}
func TestNoImplementAccount_GetGroup(t *testing.T) {
	mock := ciossdk_mock.NoImplementAccount{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.DeleteGroup(ciosctx.Background(), "")
}

func TestNoImplementAccount_CreateGroup(t *testing.T) {
	mock := ciossdk_mock.NoImplementAccount{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.CreateGroup(ciosctx.Background(), cios.GroupCreateRequest{})
}

func TestNoImplementAccount_DeleteGroup(t *testing.T) {
	mock := ciossdk_mock.NoImplementAccount{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.DeleteGroup(ciosctx.Background(), "")

}

func TestNoImplementAccount_UpdateGroup(t *testing.T) {
	mock := ciossdk_mock.NoImplementAccount{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.UpdateGroup(ciosctx.Background(), "", cios.GroupUpdateRequest{})
}
func TestNoImplementAccount_GetMe(t *testing.T) {
	mock := ciossdk_mock.NoImplementAccount{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetMe(ciosctx.Background())
}

func TestNoImplementAccount_InviteGroup(t *testing.T) {
	mock := ciossdk_mock.NoImplementAccount{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.InviteGroup(ciosctx.Background(), "", "")
}
func TestNoImplementAccount_GetResourceOwners(t *testing.T) {
	mock := ciossdk_mock.NoImplementAccount{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetResourceOwners(ciosctx.Background(), cios.ApiGetResourceOwnersRequest{})
}
func TestNoImplementAccount_GetResourceOwnersAll(t *testing.T) {
	mock := ciossdk_mock.NoImplementAccount{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetResourceOwnersAll(ciosctx.Background(), cios.ApiGetResourceOwnersRequest{})
}
func TestNoImplementAccount_GetResourceOwnersUnlimited(t *testing.T) {
	mock := ciossdk_mock.NoImplementAccount{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetResourceOwnersUnlimited(ciosctx.Background(), cios.ApiGetResourceOwnersRequest{})
}
func TestNoImplementAccount_GetResourceOwner(t *testing.T) {
	mock := ciossdk_mock.NoImplementAccount{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetResourceOwner(ciosctx.Background(), "")
}
