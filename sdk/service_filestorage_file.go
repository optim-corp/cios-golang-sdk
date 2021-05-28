package ciossdk

import (
	"crypto/md5"
	"encoding/base64"
	_nethttp "net/http"

	cnv "github.com/fcfcqloow/go-advance/convert"
	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	"github.com/optim-corp/cios-golang-sdk/util"
)

func MakeUploadFileOpts() cios.ApiUploadFileRequest {
	return cios.ApiUploadFileRequest{}
}

func (self *CiosFileStorage) DownloadFile(ctx ciosctx.RequestCtx, bucketID string, nodeID string) ([]byte, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, nil, err
	}
	request := self.ApiClient.FileStorageApi.DownloadFile(self.withHost(ctx), bucketID).NodeId(nodeID)
	f, httpResponse, err := request.Execute()
	return []byte(f), httpResponse, err
}
func (self *CiosFileStorage) DownloadFileByKey(ctx ciosctx.RequestCtx, bucketID string, key string) ([]byte, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, nil, err
	}
	request := self.ApiClient.FileStorageApi.DownloadFile(self.withHost(ctx), bucketID).Key(key)
	f, httpResponse, err := request.Execute()
	return []byte(f), httpResponse, err
}
func (self *CiosFileStorage) UploadFile(ctx ciosctx.RequestCtx, bucketID string, body []byte, params cios.ApiUploadFileRequest) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	params.Ctx = self.withHost(ctx)
	params.ApiService = self.ApiClient.FileStorageApi
	params.P_bucketId = bucketID
	params.P_name = util.ToNil(params.P_name)
	params.P_key = util.ToNil(params.P_key)
	params.P_nodeId = util.ToNil(params.P_nodeId)
	params.P_parentNodeId = util.ToNil(params.P_parentNodeId)
	params.P_body = cnv.StrPtr(string(body))
	params.P_checksum = cnv.StrPtr(toCheckSum(body))
	_, httpResponse, err := params.Execute()
	return httpResponse, err
}

func toCheckSum(byts []byte) string {
	h := md5.New()
	h.Write(byts)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
