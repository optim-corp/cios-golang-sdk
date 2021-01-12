package ciossdk

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/util"

	xmath "github.com/optim-kazuhiro-seida/go-advance-type/math"

	"github.com/optim-corp/cios-golang-sdk/cios"

	"github.com/optim-corp/cios-golang-sdk/model"
)

func MakeGetNodesOpts() cios.ApiGetNodesRequest {
	return cios.ApiGetNodesRequest{}
}
func (self FileStorage) GetNodes(bucketID string, params cios.ApiGetNodesRequest, ctx model.RequestCtx) (response cios.MultipleNode, httpResponse *_nethttp.Response, err error) {
	params.Ctx = ctx
	params.ApiService = self.ApiClient.FileStorageApi
	params.P_bucketId = bucketID
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_name = util.ToNil(params.P_name)
	params.P_parentNodeId = util.ToNil(params.P_parentNodeId)
	params.P_key = util.ToNil(params.P_key)
	response, httpResponse, err = params.Execute()
	if err != nil && self.autoR {
		if _, _, _, _, err = self.refresh(); err != nil {
			return
		}
		response, httpResponse, err = params.Execute()
	}
	return
}
func (self FileStorage) GetNodesAll(bucketID string, params cios.ApiGetNodesRequest, ctx model.RequestCtx) ([]cios.Node, *_nethttp.Response, error) {
	var (
		result      []cios.Node
		httpRes     *_nethttp.Response
		err         error
		offset      = int64(0)
		_limit      = int64(1000)
		getFunction = func(offset *int64) (cios.MultipleNode, *_nethttp.Response, error) {
			if offset != nil {
				params.P_offset = offset
			}
			tlimit := xmath.MinInt64(_limit, 1000)
			params.P_limit = &tlimit
			return self.GetNodes(bucketID, params, ctx)
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			res, httpRes, err := getFunction(&offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Nodes...)
			offset += 1000
			_limit -= 1000
			if _limit <= 0 {
				break
			}
		}
	} else {
		res, httpRes, err := getFunction(&offset)
		if err != nil {
			return nil, httpRes, err
		}
		result = append(result, res.Nodes...)
		for offset = int64(1000); offset < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(&offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Nodes...)
		}
	}
	return result, httpRes, err
}
func (self FileStorage) GetNodesUnlimited(bucketID string, params cios.ApiGetNodesRequest, ctx model.RequestCtx) ([]cios.Node, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetNodesAll(bucketID, params, ctx)
}

func (self FileStorage) GetNode(bucketID string, nodeID string, ctx model.RequestCtx) (cios.Node, *_nethttp.Response, error) {
	request := self.ApiClient.FileStorageApi.GetNode(ctx, bucketID, nodeID)
	response, httpResponse, err := request.Execute()
	if err != nil {
		if self.autoR {
			if _, _, _, _, err = self.refresh(); err != nil {
				return cios.Node{}, httpResponse, err
			}
			response, httpResponse, err = request.Execute()
		}
		if err != nil {
			return cios.Node{}, httpResponse, err
		}
	}
	return response.Node, httpResponse, err
}
func (self FileStorage) CreateNode(bucketID string, name string, parentNodeID *string, ctx model.RequestCtx) (cios.Node, *_nethttp.Response, error) {
	return self.CreateNodeOnNodeID(bucketID, cios.NodeRequest{
		Name:         name,
		ParentNodeId: parentNodeID,
	}, ctx)
}
func (self FileStorage) CreateNodeOnNodeID(bucketID string, body cios.NodeRequest, ctx model.RequestCtx) (cios.Node, *_nethttp.Response, error) {
	request := self.ApiClient.FileStorageApi.CreateDirectory(ctx, bucketID).NodeRequest(body)
	response, httpResponse, err := request.Execute()
	if err != nil {
		if self.autoR {
			if _, _, _, _, err = self.refresh(); err != nil {
				return cios.Node{}, httpResponse, err
			}
			response, httpResponse, err = request.Execute()
		}
		if err != nil {
			return cios.Node{}, httpResponse, err
		}
	}
	return response.Node, httpResponse, err
}

func (self FileStorage) DeleteNode(bucketID string, nodeID string, ctx model.RequestCtx) (*_nethttp.Response, error) {
	request := self.ApiClient.FileStorageApi.DeleteNode(ctx, bucketID, nodeID)
	httpResponse, err := request.Execute()
	if err != nil && self.autoR {
		if _, _, _, _, err = self.refresh(); err != nil {
			return httpResponse, err
		}
		return request.Execute()
	}
	return httpResponse, err
}

func (self FileStorage) CopyNode(bucketID string, nodeID string, destBucketID *string, destParentNodeID *string, ctx model.RequestCtx) (cios.Node, *_nethttp.Response, error) {
	request := self.ApiClient.FileStorageApi.CopyNode(ctx, bucketID, nodeID).BucketEditBody(cios.BucketEditBody{
		DestBucketId: destBucketID,
		ParentNodeId: destParentNodeID,
	})
	response, httpResponse, err := request.Execute()
	if err != nil {
		if self.autoR {
			if _, _, _, _, err = self.refresh(); err != nil {
				return cios.Node{}, httpResponse, err
			}
			response, httpResponse, err = request.Execute()
		}
		if err != nil {
			return cios.Node{}, httpResponse, err
		}
	}
	return response.Node, httpResponse, err
}

func (self FileStorage) MoveNode(bucketID string, nodeID string, destBucketID *string, destParentNodeID *string, ctx model.RequestCtx) (cios.Node, *_nethttp.Response, error) {
	request := self.ApiClient.FileStorageApi.MoveNode(ctx, bucketID, nodeID).BucketEditBody(cios.BucketEditBody{
		DestBucketId: destBucketID,
		ParentNodeId: destParentNodeID,
	})
	response, httpResponse, err := request.Execute()
	if err != nil {
		if self.autoR {
			if _, _, _, _, err = self.refresh(); err != nil {
				return cios.Node{}, httpResponse, err
			}
			response, httpResponse, err = request.Execute()
		}
		if err != nil {
			return cios.Node{}, httpResponse, err
		}
	}
	return response.Node, httpResponse, err
}

func (self FileStorage) RenameNode(bucketID string, nodeID string, name string, ctx model.RequestCtx) (cios.Node, *_nethttp.Response, error) {
	request := self.ApiClient.FileStorageApi.RenameNode(ctx, bucketID, nodeID).BucketName(cios.BucketName{Name: name})
	response, httpResponse, err := request.Execute()
	if err != nil {
		if self.autoR {
			self.refresh()
			response, httpResponse, err = request.Execute()
		}
		if err != nil {
			return cios.Node{}, httpResponse, err
		}
	}
	return response.Node, httpResponse, err

}
