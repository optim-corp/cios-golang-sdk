package ciossdk

import (
	"fmt"
	"io/ioutil"
	"log"
	_http "net/http"
	"net/http/httputil"

	"github.com/optim-kazuhiro-seida/go-advance-type/check"
	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
	"github.com/optim-corp/cios-golang-sdk/util"
)

func MakeGetVideosOpts() cios.ApiGetVideoStreamsListRequest {
	return cios.ApiGetVideoStreamsListRequest{}
}

func (self *VideoStreaming) GetVideoInfos(params cios.ApiGetVideoStreamsListRequest, ctx sdkmodel.RequestCtx) ([]cios.Video, *_http.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, nil, err
	}
	params.Ctx = self.withHost(ctx)
	params.ApiService = self.ApiClient.VideostreamingOperationsApi
	params.P_deviceId = util.ToNil(params.P_deviceId)
	params.P_resourceOwnerId = util.ToNil(params.P_resourceOwnerId)
	r, h, e := params.Execute()
	if e != nil {
		return nil, nil, e
	}
	return r.Videos, h, e
}
func (self *VideoStreaming) GetVideoInfo(videoID string, ctx sdkmodel.RequestCtx) (cios.Video, *_http.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Video{}, nil, err
	}
	response, httpResponse, err := self.ApiClient.VideostreamingOperationsApi.GetVideoStreams(self.withHost(ctx), videoID).Execute()
	if err != nil {
		return cios.Video{}, httpResponse, err
	}
	return response.Video, httpResponse, err
}

func (self *VideoStreaming) UpdateVideoInfo(videoID string, name, description string, ctx sdkmodel.RequestCtx) (cios.Video, *_http.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Video{}, nil, err
	}
	response, httpResponse, err := self.ApiClient.VideostreamingOperationsApi.
		UpdateStreams(self.withHost(ctx), videoID).
		VideoUpdateRequest(cios.VideoUpdateRequest{
			VideoName:        util.ToNilStr(name),
			VideoDescription: util.ToNilStr(description),
		}).Execute()
	if err != nil {
		return cios.Video{}, httpResponse, err
	}
	return response.Video, httpResponse, err
}

// MEMO: No OpenAPI
func (self *VideoStreaming) GetThumbnail(videoID string, ctx sdkmodel.RequestCtx) (img []byte, httpResponse *_http.Response, err error) {
	if err := self.refresh(); err != nil {
		return nil, nil, err
	}

	req, err := _http.NewRequest(_http.MethodGet, fmt.Sprintf("%s/v2/video_streams/%s/thumbnail", self.Url, videoID), nil)
	if err != nil {
		return
	}

	token := GetTokenFromCtx(ctx)
	if check.IsNil(token) {
		token = self.token
	}

	req.Header.Add("Authentication", convert.MustStr(token))

	if self.ApiClient.GetConfig().Debug {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return nil, nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	if httpResponse, err = self.ApiClient.GetConfig().HTTPClient.Do(req); err == nil {
		defer httpResponse.Body.Close()
		img, err = ioutil.ReadAll(httpResponse.Body)
	}

	if err != nil {
		return
	}

	if self.ApiClient.GetConfig().Debug {
		dump, err := httputil.DumpResponse(httpResponse, true)
		if err != nil {
			return nil, httpResponse, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return
}

func (self *VideoStreaming) Play(videoID string, ctx sdkmodel.RequestCtx) (cios.Room, *_http.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Room{}, nil, err
	}
	room, httpResponse, err := self.ApiClient.VideostreamingOperationsApi.CreateVideoStreamsPlay(self.withHost(ctx), videoID).Execute()
	if err != nil {
		return cios.Room{}, httpResponse, err
	}
	return room.Room, httpResponse, err
}
func (self *VideoStreaming) Stop(videoID string, ctx sdkmodel.RequestCtx) (*_http.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	return self.ApiClient.VideostreamingOperationsApi.CreateVideoStreamsStop(self.withHost(ctx), videoID).Execute()
}
