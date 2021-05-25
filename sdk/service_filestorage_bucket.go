package ciossdk

import (
	"errors"
	_nethttp "net/http"

	cnv "github.com/fcfcqloow/go-advance/convert"

	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	"github.com/optim-corp/cios-golang-sdk/util"

	xmath "github.com/fcfcqloow/go-advance/math"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

func MakeGetBucketsOpts() cios.ApiGetBucketsRequest {
	return cios.ApiGetBucketsRequest{}
}

func (self *FileStorage) GetBuckets(ctx ciosctx.RequestCtx, params cios.ApiGetBucketsRequest) (response cios.MultipleBucket, httpResponse *_nethttp.Response, err error) {
	if err = self.refresh(); err != nil {
		return
	}
	params.ApiService = self.ApiClient.FileStorageApi
	params.Ctx = self.withHost(ctx)
	params.P_name = util.ToNil(params.P_name)
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_resourceOwnerId = util.ToNil(params.P_resourceOwnerId)
	return params.Execute()
}
func (self *FileStorage) GetBucketsAll(ctx ciosctx.RequestCtx, params cios.ApiGetBucketsRequest) ([]cios.Bucket, *_nethttp.Response, error) {
	var (
		result      []cios.Bucket
		httpRes     *_nethttp.Response
		err         error
		offset      = int64(0)
		_limit      = int64(1000)
		getFunction = func(offset int64) (cios.MultipleBucket, *_nethttp.Response, error) {
			return self.GetBuckets(ctx, params.Limit(xmath.MinInt64(_limit, 1000)).Offset(offset+cnv.MustInt64(params.P_offset)))
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			res, httpRes, err := getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Buckets...)
			offset += 1000
			_limit -= 1000
			if _limit <= 0 {
				break
			}
		}
	} else {
		res, httpRes, err := getFunction(offset)
		if err != nil {
			return nil, httpRes, err
		}
		result = append(result, res.Buckets...)
		for offset = int64(1000); offset+cnv.MustInt64(params.P_offset) < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Buckets...)
		}
	}
	return result, httpRes, err
}
func (self *FileStorage) GetBucketsUnlimited(ctx ciosctx.RequestCtx, params cios.ApiGetBucketsRequest) ([]cios.Bucket, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetBucketsAll(ctx, params)
}

func (self *FileStorage) GetBucket(ctx ciosctx.RequestCtx, bucketID string) (cios.Bucket, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Bucket{}, nil, err
	}
	response, httpResponse, err := self.ApiClient.FileStorageApi.GetBucket(self.withHost(ctx), bucketID).Execute()
	if err != nil {
		return cios.Bucket{}, httpResponse, err
	}
	return response.Bucket, httpResponse, err
}
func (self *FileStorage) GetBucketByResourceOwnerIDAndName(ctx ciosctx.RequestCtx, resourceOwnerID string, name string) (cios.Bucket, *_nethttp.Response, error) {
	buckets, httpResponse, err := self.GetBucketsUnlimited(ctx, MakeGetBucketsOpts().ResourceOwnerId(resourceOwnerID).Name(name))
	if len(buckets) == 0 {
		return cios.Bucket{}, nil, errors.New("No Bucket")
	}
	if err != nil {
		return cios.Bucket{}, httpResponse, err
	}
	return buckets[0], httpResponse, err
}
func (self *FileStorage) GetOrCreateBucket(ctx ciosctx.RequestCtx, resourceOwnerID string, name string) (cios.Bucket, *_nethttp.Response, error) {
	res, httpResponse, err := self.GetBucketByResourceOwnerIDAndName(ctx, resourceOwnerID, name)
	if err != nil || res.Id == "" {
		return self.CreateBucket(ctx, resourceOwnerID, name)
	}
	return res, httpResponse, err
}
func (self *FileStorage) CreateBucket(ctx ciosctx.RequestCtx, resourceOwnerID string, name string) (cios.Bucket, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Bucket{}, nil, err
	}
	request := self.ApiClient.FileStorageApi.CreateBucket(self.withHost(ctx)).BucketRequest(cios.BucketRequest{ResourceOwnerId: resourceOwnerID, Name: name})
	response, httpResponse, err := request.Execute()
	if err != nil {
		return cios.Bucket{}, httpResponse, err
	}
	return response.Bucket, httpResponse, err
}
func (self *FileStorage) DeleteBucket(ctx ciosctx.RequestCtx, bucketID string) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	return self.ApiClient.FileStorageApi.DeleteBucket(self.withHost(ctx), bucketID).Execute()
}
func (self *FileStorage) UpdateBucket(ctx ciosctx.RequestCtx, bucketID string, name string) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	return self.ApiClient.FileStorageApi.UpdateBucket(self.withHost(ctx), bucketID).BucketName(cios.BucketName{Name: name}).Execute()
}
