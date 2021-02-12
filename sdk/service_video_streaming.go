package ciossdk

import (
	"fmt"
	_http "net/http"

	"github.com/optim-kazuhiro-seida/go-advance-type/check"

	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
	"github.com/optim-corp/cios-golang-sdk/util"
)

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
	response, httpResponse, err := self.ApiClient.VideostreamingOperationsApi.GetVideoStreams(ctx, videoID).Execute()
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
		UpdateStreams(ctx, videoID).
		VideoUpdateRequest(cios.VideoUpdateRequest{
			VideoName:        util.ToNilStr(name),
			VideoDescription: util.ToNilStr(description),
		}).Execute()
	if err != nil {
		return cios.Video{}, httpResponse, err
	}
	return response.Video, httpResponse, err
}

func (self *VideoStreaming) GetThumbnail(videoID string, ctx sdkmodel.RequestCtx) (img []byte, httpResponse *_http.Response, err error) {
	if err := self.refresh(); err != nil {
		return nil, nil, err
	}

	token := GetTokenFromCtx(ctx)
	if check.IsNil(token) {
		token = self.token
	}

	req, err := _http.NewRequest(_http.MethodGet, fmt.Sprintf("%s/v2/video_streams/%s/thumbnail", self.Url, videoID), nil)
	if err != nil {
		return
	}
	req.Header.Add("Authentication", convert.MustStr(token))

	if httpResponse, err = self.ApiClient.GetConfig().HTTPClient.Do(req); err == nil {
		err = convert.UnMarshalJson(&img, httpResponse.Body)
	}
	return
}

func (self *VideoStreaming) Play(videoID string, ctx sdkmodel.RequestCtx) (cios.Room, *_http.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.Room{}, nil, err
	}
	room, httpResponse, err := self.ApiClient.VideostreamingOperationsApi.CreateVideoStreamsPlay(ctx, videoID).Execute()
	if err != nil {
		return cios.Room{}, httpResponse, err
	}
	return room.Room, httpResponse, err
}
func (self *VideoStreaming) Stop(videoID string, ctx sdkmodel.RequestCtx) (*_http.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	return self.ApiClient.VideostreamingOperationsApi.CreateVideoStreamsStop(ctx, videoID).Execute()
}
