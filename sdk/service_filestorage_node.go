package ciossdk

import (
	_nethttp "net/http"

	cnv "github.com/fcfcqloow/go-advance/convert"

	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	"github.com/optim-corp/cios-golang-sdk/util"

	xmath "github.com/fcfcqloow/go-advance/math"

	"github.com/optim-corp/cios-golang-sdk/cios"
)

func MakeGetNodesOpts() cios.ApiGetNodesRequest {
	return cios.ApiGetNodesRequest{}
}
func (self *FileStorage) GetNodes(ctx ciosctx.RequestCtx, bucketID string, params cios.ApiGetNodesRequest) (response cios.MultipleNode, httpResponse *_nethttp.Response, err error) {
	if err := self.refresh(); err != nil {
		return cios.MultipleNode{}, nil, err
	}
	params.Ctx = self.withHost(ctx)
	params.ApiService = self.ApiClient.FileStorageApi
	params.P_bucketId = bucketID
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_name = util.ToNil(params.P_name)
	params.P_parentNodeId = util.ToNil(params.P_parentNodeId)
	params.P_key = util.ToNil(params.P_key)
	return params.Execute()
}
func (self *FileStorage) GetNodesAll(ctx ciosctx.RequestCtx, bucketID string, params cios.ApiGetNodesRequest) ([]cios.Node, *_nethttp.Response, error) {
	var (
		result      []cios.Node
		httpRes     *_nethttp.Response
		err         error
		offset      = int64(0)
		_limit      = int64(1000)
		getFunction = func(offset int64) (cios.MultipleNode, *_nethttp.Response, error) {
			return self.GetNodes(ctx, bucketID, params.Limit(xmath.MinInt64(_limit, 1000)).Offset(offset+cnv.MustInt64(params.P_offset)))
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			res, httpRes, err := getFunction(offset)
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
		res, httpRes, err := getFunction(offset)
		if err != nil {
			return nil, httpRes, err
		}
		result = append(result, res.Nodes...)
		for offset = int64(1000); offset+cnv.MustInt64(params.P_offset) < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Nodes...)
		}
	}
	return result, httpRes, err
}
func (self *FileStorage) GetNodesUnlimited(ctx ciosctx.RequestCtx, bucketID string, params cios.ApiGetNodesRequest) ([]cios.Node, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetNodesAll(ctx, bucketID, params)
}

func (self *FileStorage) GetNode(ctx ciosctx.RequestCtx, bucketID string, nodeID string) (cios.Node, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Node{}, nil, err
	}
	request := self.ApiClient.FileStorageApi.GetNode(self.withHost(ctx), bucketID, nodeID)
	response, httpResponse, err := request.Execute()
	if err != nil {
		return cios.Node{}, httpResponse, err
	}
	return response.Node, httpResponse, err
}
func (self *FileStorage) CreateNode(ctx ciosctx.RequestCtx, bucketID string, name string, parentNodeID *string) (cios.Node, *_nethttp.Response, error) {
	return self.CreateNodeOnNodeID(ctx, bucketID, cios.NodeRequest{
		Name:         name,
		ParentNodeId: parentNodeID,
	})
}
func (self *FileStorage) CreateNodeOnNodeID(ctx ciosctx.RequestCtx, bucketID string, body cios.NodeRequest) (cios.Node, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Node{}, nil, err
	}
	request := self.ApiClient.FileStorageApi.CreateDirectory(self.withHost(ctx), bucketID).NodeRequest(body)
	response, httpResponse, err := request.Execute()
	if err != nil {
		return cios.Node{}, httpResponse, err
	}
	return response.Node, httpResponse, err
}

func (self *FileStorage) DeleteNode(ctx ciosctx.RequestCtx, bucketID string, nodeID string) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	return self.ApiClient.FileStorageApi.DeleteNode(self.withHost(ctx), bucketID, nodeID).Execute()
}

func (self *FileStorage) CopyNode(ctx ciosctx.RequestCtx, bucketID string, nodeID string, destBucketID *string, destParentNodeID *string) (cios.Node, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Node{}, nil, err
	}
	request := self.ApiClient.FileStorageApi.CopyNode(self.withHost(ctx), bucketID, nodeID).BucketEditBody(cios.BucketEditBody{
		DestBucketId: destBucketID,
		ParentNodeId: destParentNodeID,
	})
	response, httpResponse, err := request.Execute()
	if err != nil {
		return cios.Node{}, httpResponse, err
	}
	return response.Node, httpResponse, err
}

func (self *FileStorage) MoveNode(ctx ciosctx.RequestCtx, bucketID string, nodeID string, destBucketID *string, destParentNodeID *string) (cios.Node, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Node{}, nil, err
	}
	request := self.ApiClient.FileStorageApi.MoveNode(self.withHost(ctx), bucketID, nodeID).BucketEditBody(cios.BucketEditBody{
		DestBucketId: destBucketID,
		ParentNodeId: destParentNodeID,
	})
	response, httpResponse, err := request.Execute()
	if err != nil {
		return cios.Node{}, httpResponse, err
	}
	return response.Node, httpResponse, err
}

func (self *FileStorage) RenameNode(ctx ciosctx.RequestCtx, bucketID string, nodeID string, name string) (cios.Node, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Node{}, nil, err
	}
	request := self.ApiClient.FileStorageApi.RenameNode(self.withHost(ctx), bucketID, nodeID).NodeName(cios.NodeName{Name: name})
	response, httpResponse, err := request.Execute()
	if err != nil {
		return cios.Node{}, httpResponse, err
	}
	return response.Node, httpResponse, err

}
