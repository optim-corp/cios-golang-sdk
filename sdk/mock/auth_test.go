package ciossdk_mock_test

import (
	"testing"

	ciossdk_mock "github.com/optim-corp/cios-golang-sdk/sdk/mock"
	"github.com/stretchr/testify/assert"
)

func TestNoImplementAuth_GetAccessTokenByRefreshToken(t *testing.T) {
	mock := ciossdk_mock.NoImplementAuth{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetAccessTokenByRefreshToken()
}

func TestNoImplementAuth_GetAccessTokenOnClient(t *testing.T) {
	mock := ciossdk_mock.NoImplementAuth{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetAccessTokenOnClient()

}

func TestNoImplementAuth_GetAccessTokenOnDevice(t *testing.T) {
	mock := ciossdk_mock.NoImplementAuth{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetAccessTokenOnDevice()
}

func TestNoImplementAuth_SetAssertion(t *testing.T) {
	mock := ciossdk_mock.NoImplementAuth{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.SetAssertion("")
}

func TestNoImplementAuth_SetClientId(t *testing.T) {
	mock := ciossdk_mock.NoImplementAuth{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.SetClientId("")
}

func TestNoImplementAuth_SetClientSecret(t *testing.T) {
	mock := ciossdk_mock.NoImplementAuth{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.SetClientSecret("")
}

func TestNoImplementAuth_SetDebug(t *testing.T) {
	mock := ciossdk_mock.NoImplementAuth{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.SetDebug(false)
}

func TestNoImplementAuth_SetRef(t *testing.T) {
	mock := ciossdk_mock.NoImplementAuth{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.SetRef("")
}

func TestNoImplementAuth_SetScope(t *testing.T) {
	mock := ciossdk_mock.NoImplementAuth{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.SetScope("")
}
