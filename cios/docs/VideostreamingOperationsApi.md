# \VideostreamingOperationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVideoStreamsPlay**](VideostreamingOperationsApi.md#CreateVideoStreamsPlay) | **Post** /v2/video_streams/{video_id}/play | 
[**CreateVideoStreamsStop**](VideostreamingOperationsApi.md#CreateVideoStreamsStop) | **Post** /v2/video_streams/{video_id}/stop | 
[**GetVideoStreams**](VideostreamingOperationsApi.md#GetVideoStreams) | **Get** /v2/video_streams/{video_id} | 
[**GetVideoStreamsList**](VideostreamingOperationsApi.md#GetVideoStreamsList) | **Get** /v2/video_streams | 
[**GetVideoStreamsThumbnail**](VideostreamingOperationsApi.md#GetVideoStreamsThumbnail) | **Get** /v2/video_streams/{video_id}/thumbnail | 
[**UpdateStreams**](VideostreamingOperationsApi.md#UpdateStreams) | **Patch** /v2/video_streams/{video_id} | 



## CreateVideoStreamsPlay

> SingleRoom CreateVideoStreamsPlay(ctx, videoId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    videoId := "videoId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideostreamingOperationsApi.CreateVideoStreamsPlay(context.Background(), videoId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VideostreamingOperationsApi.CreateVideoStreamsPlay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVideoStreamsPlay`: SingleRoom
    fmt.Fprintf(os.Stdout, "Response from `VideostreamingOperationsApi.CreateVideoStreamsPlay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVideoStreamsPlayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SingleRoom**](SingleRoom.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVideoStreamsStop

> CreateVideoStreamsStop(ctx, videoId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    videoId := "videoId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideostreamingOperationsApi.CreateVideoStreamsStop(context.Background(), videoId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VideostreamingOperationsApi.CreateVideoStreamsStop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVideoStreamsStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideoStreams

> SingleVideo GetVideoStreams(ctx, videoId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    videoId := "videoId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideostreamingOperationsApi.GetVideoStreams(context.Background(), videoId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VideostreamingOperationsApi.GetVideoStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVideoStreams`: SingleVideo
    fmt.Fprintf(os.Stdout, "Response from `VideostreamingOperationsApi.GetVideoStreams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SingleVideo**](SingleVideo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideoStreamsList

> MultipleVideo GetVideoStreamsList(ctx).ResourceOwnerId(resourceOwnerId).DeviceId(deviceId).IsEnable(isEnable).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    resourceOwnerId := "resourceOwnerId_example" // string | 
    deviceId := "deviceId_example" // string |  (optional)
    isEnable := "isEnable_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideostreamingOperationsApi.GetVideoStreamsList(context.Background()).ResourceOwnerId(resourceOwnerId).DeviceId(deviceId).IsEnable(isEnable).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VideostreamingOperationsApi.GetVideoStreamsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVideoStreamsList`: MultipleVideo
    fmt.Fprintf(os.Stdout, "Response from `VideostreamingOperationsApi.GetVideoStreamsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoStreamsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceOwnerId** | **string** |  | 
 **deviceId** | **string** |  | 
 **isEnable** | **string** |  | 

### Return type

[**MultipleVideo**](MultipleVideo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideoStreamsThumbnail

> *os.File GetVideoStreamsThumbnail(ctx, videoId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    videoId := "videoId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideostreamingOperationsApi.GetVideoStreamsThumbnail(context.Background(), videoId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VideostreamingOperationsApi.GetVideoStreamsThumbnail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVideoStreamsThumbnail`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `VideostreamingOperationsApi.GetVideoStreamsThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoStreamsThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStreams

> SingleVideo UpdateStreams(ctx, videoId).VideoUpdateRequest(videoUpdateRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    videoId := "videoId_example" // string | 
    videoUpdateRequest := *openapiclient.NewVideoUpdateRequest() // VideoUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideostreamingOperationsApi.UpdateStreams(context.Background(), videoId).VideoUpdateRequest(videoUpdateRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VideostreamingOperationsApi.UpdateStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStreams`: SingleVideo
    fmt.Fprintf(os.Stdout, "Response from `VideostreamingOperationsApi.UpdateStreams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **videoUpdateRequest** | [**VideoUpdateRequest**](VideoUpdateRequest.md) |  | 

### Return type

[**SingleVideo**](SingleVideo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

