package ciossdk

import (
	"errors"
	_nethttp "net/http"

	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	"github.com/optim-corp/cios-golang-sdk/util"

	xmath "github.com/optim-kazuhiro-seida/go-advance-type/math"

	"github.com/optim-corp/cios-golang-sdk/model"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

func MakeGetBucketsOpts() cios.ApiGetBucketsRequest {
	return cios.ApiGetBucketsRequest{}
}

func (self FileStorage) GetBuckets(params cios.ApiGetBucketsRequest, ctx model.RequestCtx) (response cios.MultipleBucket, httpResponse *_nethttp.Response, err error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	params.ApiService = self.ApiClient.FileStorageApi
	params.Ctx = ctx
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_name = util.ToNil(params.P_name)
	params.P_resourceOwnerId = util.ToNil(params.P_resourceOwnerId)
	return params.Execute()
}
func (self FileStorage) GetBucketsAll(params cios.ApiGetBucketsRequest, ctx model.RequestCtx) ([]cios.Bucket, *_nethttp.Response, error) {
	var (
		result      []cios.Bucket
		httpRes     *_nethttp.Response
		err         error
		offset      = int64(0)
		_limit      = int64(1000)
		getFunction = func(offset int64) (cios.MultipleBucket, *_nethttp.Response, error) {
			return self.GetBuckets(
				params.
					Limit(xmath.MinInt64(_limit, 1000)).
					Offset(convert.MustInt64(offset)),
				ctx)
		}
	)
	if params.P_limit != nil {
		for _limit = *params.P_limit; true; {
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
		for offset = int64(1000); offset < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Buckets...)
		}
	}
	return result, httpRes, err
}
func (self FileStorage) GetBucketsUnlimited(params cios.ApiGetBucketsRequest, ctx model.RequestCtx) ([]cios.Bucket, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetBucketsAll(params, ctx)
}

func (self FileStorage) GetBucket(bucketID string, ctx model.RequestCtx) (cios.Bucket, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	request := self.ApiClient.FileStorageApi.GetBucket(ctx, bucketID)
	response, httpResponse, err := request.Execute()
	if err != nil {
		return cios.Bucket{}, httpResponse, err
	}
	return response.Bucket, httpResponse, err
}
func (self FileStorage) GetBucketByResourceOwnerIDAndName(resourceOwnerID string, name string, ctx model.RequestCtx) (cios.Bucket, *_nethttp.Response, error) {
	buckets, httpResponse, err := self.GetBucketsUnlimited(MakeGetBucketsOpts().ResourceOwnerId(resourceOwnerID).Name(name), ctx)
	if len(buckets) == 0 {
		return cios.Bucket{}, nil, errors.New("No Bucket")
	}
	if err != nil {
		return cios.Bucket{}, httpResponse, err
	}
	return buckets[0], httpResponse, err
}
func (self FileStorage) GetOrCreateBucket(resourceOwnerID string, name string, ctx model.RequestCtx) (cios.Bucket, *_nethttp.Response, error) {
	res, httpResponse, err := self.GetBucketByResourceOwnerIDAndName(resourceOwnerID, name, ctx)
	if err != nil || res.Id == "" {
		return self.CreateBucket(resourceOwnerID, name, ctx)
	}
	return res, httpResponse, err
}
func (self FileStorage) CreateBucket(resourceOwnerID string, name string, ctx model.RequestCtx) (cios.Bucket, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	request := self.ApiClient.FileStorageApi.CreateBucket(ctx).BucketRequest(cios.BucketRequest{ResourceOwnerId: resourceOwnerID, Name: name})
	response, httpResponse, err := request.Execute()
	if err != nil {
		return cios.Bucket{}, httpResponse, err
	}
	return response.Bucket, httpResponse, err
}
func (self FileStorage) DeleteBucket(bucketID string, ctx model.RequestCtx) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	return self.ApiClient.FileStorageApi.DeleteBucket(ctx, bucketID).Execute()
}
func (self FileStorage) UpdateBucket(bucketID string, name string, ctx model.RequestCtx) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	return self.ApiClient.FileStorageApi.UpdateBucket(ctx, bucketID).BucketName(cios.BucketName{Name: name}).Execute()
}
