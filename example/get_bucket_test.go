package example

import (
	_nethttp "net/http"
	"testing"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
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
	client = &ciossdk.CiosClient{}
	client.FileStorage = &MockBucket{}
	t.Run("test sample bucket", func(t *testing.T) {
		buckets := sampleGetBucket()
		if len(buckets) == 0 || buckets[0].Id != "test-bucket" {
			t.Fatal("failed test get bucket sample")
		}
	})
}
