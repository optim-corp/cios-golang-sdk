package ciossdk_mock_test

import (
	"testing"

	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	ciossdk_mock "github.com/optim-corp/cios-golang-sdk/sdk/mock"
	"github.com/stretchr/testify/assert"
)

func TestNoImplementFileStorage_CopyNode(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.CopyNode(ciosctx.Background(), "", "", nil, nil)
}

func TestNoImplementFileStorage_CreateBucket(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.CreateBucket(ciosctx.Background(), "", "")
}

func TestNoImplementFileStorage_CreateNode(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.CreateNode(ciosctx.Background(), "", "", nil)
}

func TestNoImplementFileStorage_CreateNodeOnNodeID(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.CreateNodeOnNodeID(ciosctx.Background(), "", cios.NodeRequest{})
}

func TestNoImplementFileStorage_DeleteBucket(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.DeleteBucket(ciosctx.Background(), "")
}

func TestNoImplementFileStorage_DeleteNode(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.DeleteNode(ciosctx.Background(), "", "")
}

func TestNoImplementFileStorage_DownloadFile(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.DownloadFile(ciosctx.Background(), "", "")
}

func TestNoImplementFileStorage_DownloadFileByKey(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.DownloadFileByKey(ciosctx.Background(), "", "")
}

func TestNoImplementFileStorage_GetBucket(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetBucket(ciosctx.Background(), "")
}

func TestNoImplementFileStorage_GetBucketByResourceOwnerIDAndName(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetBucketByResourceOwnerIDAndName(ciosctx.Background(), "", "")
}

func TestNoImplementFileStorage_GetBuckets(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetBuckets(ciosctx.Background(), ciossdk.MakeGetBucketsOpts())
}

func TestNoImplementFileStorage_GetBucketsAll(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetBucketsAll(ciosctx.Background(), ciossdk.MakeGetBucketsOpts())
}

func TestNoImplementFileStorage_GetBucketsUnlimited(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetBucketsUnlimited(ciosctx.Background(), ciossdk.MakeGetBucketsOpts())
}

func TestNoImplementFileStorage_GetNode(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetNode(ciosctx.Background(), "", "")
}

func TestNoImplementFileStorage_GetNodes(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetNodes(ciosctx.Background(), "", cios.ApiGetNodesRequest{})
}

func TestNoImplementFileStorage_GetNodesAll(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetNodesAll(ciosctx.Background(), "", cios.ApiGetNodesRequest{})
}

func TestNoImplementFileStorage_GetNodesUnlimited(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetNodesUnlimited(ciosctx.Background(), "", cios.ApiGetNodesRequest{})
}

func TestNoImplementFileStorage_GetOrCreateBucket(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.GetOrCreateBucket(ciosctx.Background(), "", "")
}

func TestNoImplementFileStorage_MoveNode(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.MoveNode(ciosctx.Background(), "", "", nil, nil)
}

func TestNoImplementFileStorage_RenameNode(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.RenameNode(ciosctx.Background(), "", "", "")
}

func TestNoImplementFileStorage_UpdateBucket(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.UpdateBucket(ciosctx.Background(), "", "")
}

func TestNoImplementFileStorage_UploadFile(t *testing.T) {
	mock := ciossdk_mock.NoImplementFileStorage{}
	defer func() {
		err := recover()
		t.Log("Panic Err", err)
		assert.Equal(t, "implement me", err)
	}()
	mock.UploadFile(ciosctx.Background(), "", nil, cios.ApiUploadFileRequest{})
}
