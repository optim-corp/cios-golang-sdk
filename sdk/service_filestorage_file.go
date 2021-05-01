package ciossdk

import (
	"crypto/md5"
	"encoding/base64"
	_nethttp "net/http"

	cnv "github.com/fcfcqloow/go-advance/convert"
	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
	"github.com/optim-corp/cios-golang-sdk/util"
)

func MakeUploadFileOpts() cios.ApiUploadFileRequest {
	return cios.ApiUploadFileRequest{}
}

func (self *FileStorage) DownloadFile(bucketID string, nodeID string, ctx sdkmodel.RequestCtx) ([]byte, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, nil, err
	}
	request := self.ApiClient.FileStorageApi.DownloadFile(self.withHost(ctx), bucketID).NodeId(nodeID)
	f, httpResponse, err := request.Execute()
	return []byte(f), httpResponse, err
}
func (self *FileStorage) DownloadFileByKey(bucketID string, key string, ctx sdkmodel.RequestCtx) ([]byte, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, nil, err
	}
	request := self.ApiClient.FileStorageApi.DownloadFile(self.withHost(ctx), bucketID).Key(key)
	f, httpResponse, err := request.Execute()
	return []byte(f), httpResponse, err
}
func (self *FileStorage) UploadFile(bucketID string, body []byte, params cios.ApiUploadFileRequest, ctx sdkmodel.RequestCtx) (*_nethttp.Response, error) {
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
