package ciossdk_mock

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
)

type NoImplementFileStorage struct{}

func (NoImplementFileStorage) GetBuckets(ctx ciosctx.RequestCtx, request cios.ApiGetBucketsRequest) (cios.MultipleBucket, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) GetBucketsAll(ctx ciosctx.RequestCtx, request cios.ApiGetBucketsRequest) ([]cios.Bucket, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) GetBucketsUnlimited(ctx ciosctx.RequestCtx, request cios.ApiGetBucketsRequest) ([]cios.Bucket, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) GetBucket(ctx ciosctx.RequestCtx, s string) (cios.Bucket, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) GetBucketByResourceOwnerIDAndName(ctx ciosctx.RequestCtx, s string, s2 string) (cios.Bucket, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) GetOrCreateBucket(ctx ciosctx.RequestCtx, s string, s2 string) (cios.Bucket, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) CreateBucket(ctx ciosctx.RequestCtx, s string, s2 string) (cios.Bucket, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) DeleteBucket(ctx ciosctx.RequestCtx, s string) (*_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) UpdateBucket(ctx ciosctx.RequestCtx, s string, s2 string) (*_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) DownloadFile(ctx ciosctx.RequestCtx, s string, s2 string) ([]byte, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) DownloadFileByKey(ctx ciosctx.RequestCtx, s string, s2 string) ([]byte, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) UploadFile(ctx ciosctx.RequestCtx, s string, bytes []byte, request cios.ApiUploadFileRequest) (*_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) GetNodes(ctx ciosctx.RequestCtx, s string, request cios.ApiGetNodesRequest) (cios.MultipleNode, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) GetNodesAll(ctx ciosctx.RequestCtx, s string, request cios.ApiGetNodesRequest) ([]cios.Node, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) GetNodesUnlimited(ctx ciosctx.RequestCtx, s string, request cios.ApiGetNodesRequest) ([]cios.Node, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) GetNode(ctx ciosctx.RequestCtx, s string, s2 string) (cios.Node, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) CreateNode(ctx ciosctx.RequestCtx, s string, s2 string, s3 *string) (cios.Node, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) CreateNodeOnNodeID(ctx ciosctx.RequestCtx, s string, request cios.NodeRequest) (cios.Node, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) DeleteNode(ctx ciosctx.RequestCtx, s string, s2 string) (*_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) CopyNode(ctx ciosctx.RequestCtx, s string, s2 string, s3 *string, s4 *string) (cios.Node, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) MoveNode(ctx ciosctx.RequestCtx, s string, s2 string, s3 *string, s4 *string) (cios.Node, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementFileStorage) RenameNode(ctx ciosctx.RequestCtx, s string, s2 string, s3 string) (cios.Node, *_nethttp.Response, error) {
	panic("implement me")
}
