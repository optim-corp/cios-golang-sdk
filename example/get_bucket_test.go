package example

import (
	_nethttp "net/http"
	"testing"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	ciossdk_mock "github.com/optim-corp/cios-golang-sdk/sdk/mock"
)

type MockBucket struct {
	ciossdk_mock.NoImplementFileStorage
}

func (*MockBucket) GetBuckets(ctx ciosctx.RequestCtx, request cios.ApiGetBucketsRequest) (cios.MultipleBucket, *_nethttp.Response, error) {
	return cios.MultipleBucket{
		Total:   1,
		Buckets: []cios.Bucket{{Id: "test-bucket"}},
	}, nil, nil
}

func Test_sampleGetBucket(t *testing.T) {
	mock := &ciossdk_mock.MockClient{}
	mock.SetFileStorage(&MockBucket{})
	client = mock
	t.Run("test sample bucket", func(t *testing.T) {
		buckets := sampleGetBucket()
		if len(buckets) == 0 || buckets[0].Id != "test-bucket" {
			t.Fatal("failed test get bucket sample")
		}
	})
}
