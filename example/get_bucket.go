package example

import (
	"fmt"

	"github.com/optim-corp/cios-golang-sdk/cios"

	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	srvfilestorage "github.com/optim-corp/cios-golang-sdk/sdk/service/filestorage"
)

func sampleGetBucket() []cios.Bucket {
	// usage
	response, httpResponse, err := client.FileStorage.GetBuckets(ciosctx.Background(), srvfilestorage.MakeGetBucketsOpts().Limit(30))
	if err != nil {
		println(err)
	}
	fmt.Println(httpResponse)
	return response.Buckets
}
