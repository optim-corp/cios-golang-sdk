# Video Streaming

## Video Streaming API

```
GetVideoInfos(ciosctx.RequestCtx, cios.ApiGetVideoStreamsListRequest) ([]cios.Video, *_nethttp.Response, error)
GetVideoInfo(ciosctx.RequestCtx, string) (cios.Video, *_nethttp.Response, error)
UpdateVideoInfo(ciosctx.RequestCtx, string, string, string) (cios.Video, *_nethttp.Response, error)
GetThumbnail(ciosctx.RequestCtx, string) ([]byte, *_nethttp.Response, error)
Play(ciosctx.RequestCtx, string) (cios.Room, *_nethttp.Response, error)
Stop(ciosctx.RequestCtx, string) (*_nethttp.Response, error)
```

#### Get a Video information

```go
videoInfo, httpResponse, err := client.Video.GetVideoInfo(ctx, "video_id")
```

#### Get Videos information

```go
options := srvvideo.MakeGetVideosOpts()
videosInfo, httpResponse, err := client.Video.GetVideoInfo(ctx, options())
```

#### Update a Video information

```go
videosInfo, httpResponse, err := client.Video.UpdateVideoInfo(ctx, "video_id", "name", "description")
```

#### Get a Thumbnail

```go
byts, httpResponse, err := client.Video.GetThumbnail(ctx, "video_id")
```

#### Play Streaming

```go
room, httpResponse, err := client.Video.Play(ctx, "video_id")
```

#### Stop Streaming

```go
httpResponse, err := client.Video.Stop(ctx, "video_id")
```
