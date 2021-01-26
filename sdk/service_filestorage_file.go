package ciossdk

import (
	"crypto/md5"
	"encoding/base64"
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/util"

	"github.com/optim-kazuhiro-seida/go-advance-type/check"
	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	"github.com/optim-corp/cios-golang-sdk/cios"

	"github.com/optim-corp/cios-golang-sdk/model"
)

func MakeUploadFileOpts() cios.ApiUploadFileRequest {
	return cios.ApiUploadFileRequest{}
}

func (self FileStorage) DownloadFile(bucketID string, nodeID string, ctx model.RequestCtx) ([]byte, *_nethttp.Response, error) {
	request := self.ApiClient.FileStorageApi.DownloadFile(ctx, bucketID).NodeId(nodeID)
	f, httpResponse, err := request.Execute()
	if err != nil && !check.IsNil(self.refresh) {
		if _, _, _, _, err = (*self.refresh)(); err == nil {
			f, httpResponse, err = request.Execute()
		}
	}
	return []byte(f), httpResponse, err
}
func (self FileStorage) DownloadFileByKey(bucketID string, key string, ctx model.RequestCtx) ([]byte, *_nethttp.Response, error) {
	request := self.ApiClient.FileStorageApi.DownloadFile(ctx, bucketID).Key(key)
	f, httpResponse, err := request.Execute()
	if err != nil && !check.IsNil(self.refresh) {
		if _, _, _, _, err = (*self.refresh)(); err == nil {
			f, httpResponse, err = request.Execute()
		}
		f, httpResponse, err = request.Execute()
	}
	return []byte(f), httpResponse, err
}

func (self FileStorage) UploadFile(bucketID string, body []byte, params cios.ApiUploadFileRequest, ctx model.RequestCtx) (*_nethttp.Response, error) {
	params.Ctx = ctx
	params.ApiService = self.ApiClient.FileStorageApi
	params.P_bucketId = bucketID
	params.P_name = util.ToNil(params.P_name)
	params.P_nodeId = util.ToNil(params.P_nodeId)
	params.P_parentNodeId = util.ToNil(params.P_parentNodeId)
	params.P_body = convert.StringPtr(string(body))
	params.P_checksum = convert.StringPtr(toCheckSum(body))
	_, httpResponse, err := params.Execute()
	if err != nil && !check.IsNil(self.refresh) {
		if _, _, _, _, err = (*self.refresh)(); err == nil {
			_, httpResponse, err = params.Execute()
		}
	}
	return httpResponse, err
}

func toCheckSum(byts []byte) string {
	h := md5.New()
	h.Write(byts)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
