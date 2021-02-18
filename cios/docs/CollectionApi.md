# \CollectionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSeriesDataBulk**](CollectionApi.md#CreateSeriesDataBulk) | **Put** /v2/collections/{collection_id}/series/{series_id}/data-bulk | 
[**CreateSeriesImage**](CollectionApi.md#CreateSeriesImage) | **Put** /v2/collections/{collection_id}/series/{series_id}/timestamp/{timestamp}/image | 
[**DeleteSeriesLatest**](CollectionApi.md#DeleteSeriesLatest) | **Delete** /v2/collections/{collection_id}/series/{series_id}/latest | 
[**GetCollectionStatus**](CollectionApi.md#GetCollectionStatus) | **Get** /v2/collections/{collection_id}/status | 
[**GetRecordedDates**](CollectionApi.md#GetRecordedDates) | **Get** /v2/collections/{collection_id}/recorded-dates | 
[**GetSeries**](CollectionApi.md#GetSeries) | **Get** /v2/collections/{collection_id}/series/{series_id} | 
[**GetSeriesAggregation**](CollectionApi.md#GetSeriesAggregation) | **Post** /v2/collections/{collection_id}/series/{series_id}/aggregation | 
[**GetSeriesImage**](CollectionApi.md#GetSeriesImage) | **Get** /v2/collections/{collection_id}/series/{series_id}/timestamp/{timestamp}/image | 
[**GetSeriesImages**](CollectionApi.md#GetSeriesImages) | **Get** /v2/collections/{collection_id}/series/{series_id}/images | 
[**GetSeriesThumbnails**](CollectionApi.md#GetSeriesThumbnails) | **Get** /v2/collections/{collection_id}/series/{series_id}/thumbnails | 
[**GetSeriesThumnbnail**](CollectionApi.md#GetSeriesThumnbnail) | **Get** /v2/collections/{collection_id}/series/{series_id}/timestamp/{timestamp}/thumbnail | 
[**GetTimeSeriesData**](CollectionApi.md#GetTimeSeriesData) | **Get** /v2/collections/{collection_id}/series/{series_id}/data | 
[**PostCollectionArchive**](CollectionApi.md#PostCollectionArchive) | **Post** /v2/collections/{collection_id}/archive | 
[**PostCollectionRestore**](CollectionApi.md#PostCollectionRestore) | **Post** /v2/collections/{collection_id}/restore | 
[**PostSearchLatest**](CollectionApi.md#PostSearchLatest) | **Post** /v2/collections/{collection_id}/search-latest | 
[**PostSeries**](CollectionApi.md#PostSeries) | **Post** /v2/collections/{collection_id}/series | 時系列データ保存
[**PostSeriesBulk**](CollectionApi.md#PostSeriesBulk) | **Post** /v2/collections/{collection_id}/series-bulk | 
[**PutSeries**](CollectionApi.md#PutSeries) | **Put** /v2/collections/{collection_id}/series/{series_id}/timestamp/{timestamp}/data | 



## CreateSeriesDataBulk

> CreateSeriesDataBulk(ctx, collectionId, seriesId).SeriesDataBulkRequest(seriesDataBulkRequest).ResourceOwnerId(resourceOwnerId).Execute()





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
    collectionId := "location" // string | コレクション名
    seriesId := "seriesId_example" // string | シリーズID
    seriesDataBulkRequest := *openapiclient.NewSeriesDataBulkRequest([]openapiclient.SeriesDataLocationUnix{*openapiclient.NewSeriesDataLocationUnix()}) // SeriesDataBulkRequest | 
    resourceOwnerId := TODO // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionApi.CreateSeriesDataBulk(context.Background(), collectionId, seriesId).SeriesDataBulkRequest(seriesDataBulkRequest).ResourceOwnerId(resourceOwnerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.CreateSeriesDataBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | コレクション名 | 
**seriesId** | **string** | シリーズID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSeriesDataBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **seriesDataBulkRequest** | [**SeriesDataBulkRequest**](SeriesDataBulkRequest.md) |  | 
 **resourceOwnerId** | [**string**](string.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSeriesImage

> CreateSeriesImage(ctx, collectionId, seriesId, timestamp).Body(body).ResourceOwnerId(resourceOwnerId).IsLatest(isLatest).Execute()





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
    collectionId := "location" // string | コレクション名
    seriesId := "seriesId_example" // string | シリーズID
    timestamp := int64(1518165596566320) // int64 | UNIXタイム(ミリ秒)
    body := string(BYTE_ARRAY_DATA_HERE) // string | 
    resourceOwnerId := TODO // string |  (optional)
    isLatest := false // bool | 最新画像としてアップロードする場合はtrue。保存に失敗した画像を再度アップロードする場合など、最新画像でない場合はfalseを指定する。 (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionApi.CreateSeriesImage(context.Background(), collectionId, seriesId, timestamp).Body(body).ResourceOwnerId(resourceOwnerId).IsLatest(isLatest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.CreateSeriesImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | コレクション名 | 
**seriesId** | **string** | シリーズID | 
**timestamp** | **int64** | UNIXタイム(ミリ秒) | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSeriesImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** |  | 
 **resourceOwnerId** | [**string**](string.md) |  | 
 **isLatest** | **bool** | 最新画像としてアップロードする場合はtrue。保存に失敗した画像を再度アップロードする場合など、最新画像でない場合はfalseを指定する。 | [default to true]

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: image/jpeg, image/png
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSeriesLatest

> DeleteSeriesLatest(ctx, collectionId, seriesId, timestamp).ResourceOwnerId(resourceOwnerId).Execute()





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
    collectionId := "location" // string | コレクション名
    seriesId := "seriesId_example" // string | シリーズID
    timestamp := int64(1518165596566320) // int64 | UNIXタイム(ミリ秒)
    resourceOwnerId := TODO // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionApi.DeleteSeriesLatest(context.Background(), collectionId, seriesId, timestamp).ResourceOwnerId(resourceOwnerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.DeleteSeriesLatest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | コレクション名 | 
**seriesId** | **string** | シリーズID | 
**timestamp** | **int64** | UNIXタイム(ミリ秒) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSeriesLatestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **resourceOwnerId** | [**string**](string.md) |  | 

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


## GetCollectionStatus

> interface{} GetCollectionStatus(ctx, collectionId).ResourceOwnerId(resourceOwnerId).Execute()





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
    collectionId := "location" // string | コレクション名
    resourceOwnerId := TODO // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionApi.GetCollectionStatus(context.Background(), collectionId).ResourceOwnerId(resourceOwnerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.GetCollectionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCollectionStatus`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `CollectionApi.GetCollectionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | コレクション名 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceOwnerId** | [**string**](string.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecordedDates

> RecordedDates GetRecordedDates(ctx, collectionId).Month(month).ResourceOwnerId(resourceOwnerId).GmtOffset(gmtOffset).Execute()





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
    collectionId := "location" // string | コレクション名
    month := "2019-12" // string | 問い合わせ年月(YYYY-MM)。1970-01から9999-12まで
    resourceOwnerId := TODO // string |  (optional)
    gmtOffset := "gmtOffset_example" // string | タイムゾーンをGMTからのオフセットで指定(日本の場合9) hh:mmの形式で指定。時間が一桁の場合は0パディングはなし。時間は[-12, 14]の範囲。分は、00, 15, 30, 45のいずれかのみ許容 (optional) (default to "0:00")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionApi.GetRecordedDates(context.Background(), collectionId).Month(month).ResourceOwnerId(resourceOwnerId).GmtOffset(gmtOffset).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.GetRecordedDates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecordedDates`: RecordedDates
    fmt.Fprintf(os.Stdout, "Response from `CollectionApi.GetRecordedDates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | コレクション名 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordedDatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **month** | **string** | 問い合わせ年月(YYYY-MM)。1970-01から9999-12まで | 
 **resourceOwnerId** | [**string**](string.md) |  | 
 **gmtOffset** | **string** | タイムゾーンをGMTからのオフセットで指定(日本の場合9) hh:mmの形式で指定。時間が一桁の場合は0パディングはなし。時間は[-12, 14]の範囲。分は、00, 15, 30, 45のいずれかのみ許容 | [default to &quot;0:00&quot;]

### Return type

[**RecordedDates**](RecordedDates.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeries

> MultipleSeriesDataLocationUnix GetSeries(ctx, collectionId, seriesId).TimeRange(timeRange).ResourceOwnerId(resourceOwnerId).AcceptEncoding(acceptEncoding).Limit(limit).Offset(offset).Ascending(ascending).Execute()





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
    collectionId := "location" // string | コレクション名
    seriesId := "seriesId_example" // string | シリーズID
    timeRange := "1518165596566:1518165596566" // string | <開始時刻>:<終了時刻>、UNIX時間(ミリ秒)。開始時刻はInclusive、終了時刻はExclusive(開始時刻≦取得範囲＜終了時刻)。最大で取得できる期間は7日間以内とする。開始、終了時刻の指定が7日間を超える場合は、400エラーとなる
    resourceOwnerId := TODO // string |  (optional)
    acceptEncoding := "acceptEncoding_example" // string | レスポンスボディの圧縮形式。gzipのみ対応。 (optional)
    limit := int64(789) // int64 | ページネーションのうち、上位何件取得するか (optional)
    offset := int64(789) // int64 | ページネーションのうち、何件読み飛ばすか (optional)
    ascending := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionApi.GetSeries(context.Background(), collectionId, seriesId).TimeRange(timeRange).ResourceOwnerId(resourceOwnerId).AcceptEncoding(acceptEncoding).Limit(limit).Offset(offset).Ascending(ascending).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.GetSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSeries`: MultipleSeriesDataLocationUnix
    fmt.Fprintf(os.Stdout, "Response from `CollectionApi.GetSeries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | コレクション名 | 
**seriesId** | **string** | シリーズID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timeRange** | **string** | &lt;開始時刻&gt;:&lt;終了時刻&gt;、UNIX時間(ミリ秒)。開始時刻はInclusive、終了時刻はExclusive(開始時刻≦取得範囲＜終了時刻)。最大で取得できる期間は7日間以内とする。開始、終了時刻の指定が7日間を超える場合は、400エラーとなる | 
 **resourceOwnerId** | [**string**](string.md) |  | 
 **acceptEncoding** | **string** | レスポンスボディの圧縮形式。gzipのみ対応。 | 
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 
 **ascending** | **bool** |  | 

### Return type

[**MultipleSeriesDataLocationUnix**](MultipleSeriesDataLocationUnix.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeriesAggregation

> SeriesAggregations GetSeriesAggregation(ctx, collectionId, seriesId).TimeRange(timeRange).Interval(interval).RequestBody(requestBody).ResourceOwnerId(resourceOwnerId).AcceptEncoding(acceptEncoding).Ascending(ascending).Execute()





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
    collectionId := "location" // string | コレクション名
    seriesId := "seriesId_example" // string | シリーズID
    timeRange := "1518165596566:1518165596566" // string | <開始時刻>:<終了時刻>、UNIX時間(ミリ秒)。開始時刻はInclusive、終了時刻はExclusive(開始時刻≦取得範囲＜終了時刻)。最大で取得できる期間は7日間以内とする。開始、終了時刻の指定が7日間を超える場合は、400エラーとなる
    interval := int32(60000) // int32 | 時系列データを集計する単位時間（ミリ秒）を指定する
    requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | 
    resourceOwnerId := TODO // string |  (optional)
    acceptEncoding := "acceptEncoding_example" // string | レスポンスボディの圧縮形式。gzipのみ対応。 (optional)
    ascending := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionApi.GetSeriesAggregation(context.Background(), collectionId, seriesId).TimeRange(timeRange).Interval(interval).RequestBody(requestBody).ResourceOwnerId(resourceOwnerId).AcceptEncoding(acceptEncoding).Ascending(ascending).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.GetSeriesAggregation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSeriesAggregation`: SeriesAggregations
    fmt.Fprintf(os.Stdout, "Response from `CollectionApi.GetSeriesAggregation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | コレクション名 | 
**seriesId** | **string** | シリーズID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSeriesAggregationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timeRange** | **string** | &lt;開始時刻&gt;:&lt;終了時刻&gt;、UNIX時間(ミリ秒)。開始時刻はInclusive、終了時刻はExclusive(開始時刻≦取得範囲＜終了時刻)。最大で取得できる期間は7日間以内とする。開始、終了時刻の指定が7日間を超える場合は、400エラーとなる | 
 **interval** | **int32** | 時系列データを集計する単位時間（ミリ秒）を指定する | 
 **requestBody** | **[]map[string]interface{}** |  | 
 **resourceOwnerId** | [**string**](string.md) |  | 
 **acceptEncoding** | **string** | レスポンスボディの圧縮形式。gzipのみ対応。 | 
 **ascending** | **bool** |  | 

### Return type

[**SeriesAggregations**](SeriesAggregations.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeriesImage

> string GetSeriesImage(ctx, collectionId, seriesId, timestamp).SeriesDataBulkRequest(seriesDataBulkRequest).ResourceOwnerId(resourceOwnerId).Execute()





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
    collectionId := "location" // string | コレクション名
    seriesId := "seriesId_example" // string | シリーズID
    timestamp := int64(1518165596566320) // int64 | UNIXタイム(ミリ秒)
    seriesDataBulkRequest := *openapiclient.NewSeriesDataBulkRequest([]openapiclient.SeriesDataLocationUnix{*openapiclient.NewSeriesDataLocationUnix()}) // SeriesDataBulkRequest | 
    resourceOwnerId := TODO // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionApi.GetSeriesImage(context.Background(), collectionId, seriesId, timestamp).SeriesDataBulkRequest(seriesDataBulkRequest).ResourceOwnerId(resourceOwnerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.GetSeriesImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSeriesImage`: string
    fmt.Fprintf(os.Stdout, "Response from `CollectionApi.GetSeriesImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | コレクション名 | 
**seriesId** | **string** | シリーズID | 
**timestamp** | **int64** | UNIXタイム(ミリ秒) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSeriesImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **seriesDataBulkRequest** | [**SeriesDataBulkRequest**](SeriesDataBulkRequest.md) |  | 
 **resourceOwnerId** | [**string**](string.md) |  | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: image/jpeg, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeriesImages

> MultipleSeriesImage GetSeriesImages(ctx, collectionId, seriesId).TimeRange(timeRange).ResourceOwnerId(resourceOwnerId).AcceptEncoding(acceptEncoding).Limit(limit).Offset(offset).Ascending(ascending).Execute()





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
    collectionId := "location" // string | コレクション名
    seriesId := "seriesId_example" // string | シリーズID
    timeRange := "1518165596566:1518165596566" // string | <開始時刻>:<終了時刻>、UNIX時間(ミリ秒)。開始時刻はInclusive、終了時刻はExclusive(開始時刻≦取得範囲＜終了時刻)。最大で取得できる期間は7日間以内とする。開始、終了時刻の指定が7日間を超える場合は、400エラーとなる
    resourceOwnerId := TODO // string |  (optional)
    acceptEncoding := "acceptEncoding_example" // string | レスポンスボディの圧縮形式。gzipのみ対応。 (optional)
    limit := int64(789) // int64 | ページネーションのうち、上位何件取得するか (optional)
    offset := int64(789) // int64 | ページネーションのうち、何件読み飛ばすか (optional)
    ascending := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionApi.GetSeriesImages(context.Background(), collectionId, seriesId).TimeRange(timeRange).ResourceOwnerId(resourceOwnerId).AcceptEncoding(acceptEncoding).Limit(limit).Offset(offset).Ascending(ascending).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.GetSeriesImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSeriesImages`: MultipleSeriesImage
    fmt.Fprintf(os.Stdout, "Response from `CollectionApi.GetSeriesImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | コレクション名 | 
**seriesId** | **string** | シリーズID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSeriesImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timeRange** | **string** | &lt;開始時刻&gt;:&lt;終了時刻&gt;、UNIX時間(ミリ秒)。開始時刻はInclusive、終了時刻はExclusive(開始時刻≦取得範囲＜終了時刻)。最大で取得できる期間は7日間以内とする。開始、終了時刻の指定が7日間を超える場合は、400エラーとなる | 
 **resourceOwnerId** | [**string**](string.md) |  | 
 **acceptEncoding** | **string** | レスポンスボディの圧縮形式。gzipのみ対応。 | 
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 
 **ascending** | **bool** |  | 

### Return type

[**MultipleSeriesImage**](MultipleSeriesImage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeriesThumbnails

> MultipleSeriesImage GetSeriesThumbnails(ctx, collectionId, seriesId).TimeRange(timeRange).ResourceOwnerId(resourceOwnerId).AcceptEncoding(acceptEncoding).Limit(limit).Offset(offset).Ascending(ascending).Execute()





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
    collectionId := "location" // string | コレクション名
    seriesId := "seriesId_example" // string | シリーズID
    timeRange := "1518165596566:1518165596566" // string | <開始時刻>:<終了時刻>、UNIX時間(ミリ秒)。開始時刻はInclusive、終了時刻はExclusive(開始時刻≦取得範囲＜終了時刻)。最大で取得できる期間は7日間以内とする。開始、終了時刻の指定が7日間を超える場合は、400エラーとなる
    resourceOwnerId := TODO // string |  (optional)
    acceptEncoding := "acceptEncoding_example" // string | レスポンスボディの圧縮形式。gzipのみ対応。 (optional)
    limit := int64(789) // int64 | ページネーションのうち、上位何件取得するか (optional)
    offset := int64(789) // int64 | ページネーションのうち、何件読み飛ばすか (optional)
    ascending := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionApi.GetSeriesThumbnails(context.Background(), collectionId, seriesId).TimeRange(timeRange).ResourceOwnerId(resourceOwnerId).AcceptEncoding(acceptEncoding).Limit(limit).Offset(offset).Ascending(ascending).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.GetSeriesThumbnails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSeriesThumbnails`: MultipleSeriesImage
    fmt.Fprintf(os.Stdout, "Response from `CollectionApi.GetSeriesThumbnails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | コレクション名 | 
**seriesId** | **string** | シリーズID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSeriesThumbnailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timeRange** | **string** | &lt;開始時刻&gt;:&lt;終了時刻&gt;、UNIX時間(ミリ秒)。開始時刻はInclusive、終了時刻はExclusive(開始時刻≦取得範囲＜終了時刻)。最大で取得できる期間は7日間以内とする。開始、終了時刻の指定が7日間を超える場合は、400エラーとなる | 
 **resourceOwnerId** | [**string**](string.md) |  | 
 **acceptEncoding** | **string** | レスポンスボディの圧縮形式。gzipのみ対応。 | 
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 
 **ascending** | **bool** |  | 

### Return type

[**MultipleSeriesImage**](MultipleSeriesImage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeriesThumnbnail

> string GetSeriesThumnbnail(ctx, collectionId, seriesId, timestamp).SeriesDataBulkRequest(seriesDataBulkRequest).ResourceOwnerId(resourceOwnerId).Execute()





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
    collectionId := "location" // string | コレクション名
    seriesId := "seriesId_example" // string | シリーズID
    timestamp := int64(1518165596566320) // int64 | UNIXタイム(ミリ秒)
    seriesDataBulkRequest := *openapiclient.NewSeriesDataBulkRequest([]openapiclient.SeriesDataLocationUnix{*openapiclient.NewSeriesDataLocationUnix()}) // SeriesDataBulkRequest | 
    resourceOwnerId := TODO // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionApi.GetSeriesThumnbnail(context.Background(), collectionId, seriesId, timestamp).SeriesDataBulkRequest(seriesDataBulkRequest).ResourceOwnerId(resourceOwnerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.GetSeriesThumnbnail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSeriesThumnbnail`: string
    fmt.Fprintf(os.Stdout, "Response from `CollectionApi.GetSeriesThumnbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | コレクション名 | 
**seriesId** | **string** | シリーズID | 
**timestamp** | **int64** | UNIXタイム(ミリ秒) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSeriesThumnbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **seriesDataBulkRequest** | [**SeriesDataBulkRequest**](SeriesDataBulkRequest.md) |  | 
 **resourceOwnerId** | [**string**](string.md) |  | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: image/jpeg, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeSeriesData

> MultipleSeriesDataLocationUnix GetTimeSeriesData(ctx, collectionId, seriesId).TimeRange(timeRange).ResourceOwnerId(resourceOwnerId).AcceptEncoding(acceptEncoding).Limit(limit).Offset(offset).Ascending(ascending).Execute()





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
    collectionId := "location" // string | コレクション名
    seriesId := "seriesId_example" // string | シリーズID
    timeRange := "1518165596566:1518165596566" // string | <開始時刻>:<終了時刻>、UNIX時間(ミリ秒)。開始時刻はInclusive、終了時刻はExclusive(開始時刻≦取得範囲＜終了時刻)。最大で取得できる期間は7日間以内とする。開始、終了時刻の指定が7日間を超える場合は、400エラーとなる
    resourceOwnerId := TODO // string |  (optional)
    acceptEncoding := "acceptEncoding_example" // string | レスポンスボディの圧縮形式。gzipのみ対応。 (optional)
    limit := int64(789) // int64 | ページネーションのうち、上位何件取得するか (optional)
    offset := int64(789) // int64 | ページネーションのうち、何件読み飛ばすか (optional)
    ascending := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionApi.GetTimeSeriesData(context.Background(), collectionId, seriesId).TimeRange(timeRange).ResourceOwnerId(resourceOwnerId).AcceptEncoding(acceptEncoding).Limit(limit).Offset(offset).Ascending(ascending).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.GetTimeSeriesData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTimeSeriesData`: MultipleSeriesDataLocationUnix
    fmt.Fprintf(os.Stdout, "Response from `CollectionApi.GetTimeSeriesData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | コレクション名 | 
**seriesId** | **string** | シリーズID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeSeriesDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timeRange** | **string** | &lt;開始時刻&gt;:&lt;終了時刻&gt;、UNIX時間(ミリ秒)。開始時刻はInclusive、終了時刻はExclusive(開始時刻≦取得範囲＜終了時刻)。最大で取得できる期間は7日間以内とする。開始、終了時刻の指定が7日間を超える場合は、400エラーとなる | 
 **resourceOwnerId** | [**string**](string.md) |  | 
 **acceptEncoding** | **string** | レスポンスボディの圧縮形式。gzipのみ対応。 | 
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 
 **ascending** | **bool** |  | 

### Return type

[**MultipleSeriesDataLocationUnix**](MultipleSeriesDataLocationUnix.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCollectionArchive

> PostCollectionArchive(ctx, collectionId).ResourceOwnerId(resourceOwnerId).Execute()





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
    collectionId := "location" // string | コレクション名
    resourceOwnerId := TODO // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionApi.PostCollectionArchive(context.Background(), collectionId).ResourceOwnerId(resourceOwnerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.PostCollectionArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | コレクション名 | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCollectionArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceOwnerId** | [**string**](string.md) |  | 

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


## PostCollectionRestore

> PostCollectionRestore(ctx, collectionId).ResourceOwnerId(resourceOwnerId).Execute()





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
    collectionId := "location" // string | コレクション名
    resourceOwnerId := TODO // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionApi.PostCollectionRestore(context.Background(), collectionId).ResourceOwnerId(resourceOwnerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.PostCollectionRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | コレクション名 | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCollectionRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceOwnerId** | [**string**](string.md) |  | 

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


## PostSearchLatest

> MultipleCollectionLatest PostSearchLatest(ctx, collectionId).CollectionLatestRequest(collectionLatestRequest).ResourceOwnerId(resourceOwnerId).Limit(limit).Offset(offset).Projection(projection).AcceptEncoding(acceptEncoding).Execute()





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
    collectionId := "location" // string | コレクション名
    collectionLatestRequest := *openapiclient.NewCollectionLatestRequest("ResourceOwnerId_example") // CollectionLatestRequest | 
    resourceOwnerId := TODO // string |  (optional)
    limit := int64(789) // int64 | ページネーションのうち、上位何件取得するか (optional)
    offset := int64(789) // int64 | ページネーションのうち、何件読み飛ばすか (optional)
    projection := "projection_example" // string | 取得する対象を以下から選択する  |値|取得対象| |---|---| |data|画像以外の最新データのみ取得。時系列データ保存APIで保存された各シリーズの時系列データで最後に保存されたデータが返却される。時系列データ保存 (bulk upload)で保存されたデータは取得されない。| |thumbnail|最新のサムネイル画像のみ取得。画像保存APIで各シリーズに対してis_latest=trueを指定して最後に保存された画像が返却される。| |image|最新のオリジナル画像のみ取得。画像保存APIで各シリーズに対してis_latest=trueを指定して最後に保存された画像が返却される。| (optional) (default to "data")
    acceptEncoding := "acceptEncoding_example" // string | レスポンスボディの圧縮形式。gzipのみ対応。 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionApi.PostSearchLatest(context.Background(), collectionId).CollectionLatestRequest(collectionLatestRequest).ResourceOwnerId(resourceOwnerId).Limit(limit).Offset(offset).Projection(projection).AcceptEncoding(acceptEncoding).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.PostSearchLatest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSearchLatest`: MultipleCollectionLatest
    fmt.Fprintf(os.Stdout, "Response from `CollectionApi.PostSearchLatest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | コレクション名 | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSearchLatestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collectionLatestRequest** | [**CollectionLatestRequest**](CollectionLatestRequest.md) |  | 
 **resourceOwnerId** | [**string**](string.md) |  | 
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 
 **projection** | **string** | 取得する対象を以下から選択する  |値|取得対象| |---|---| |data|画像以外の最新データのみ取得。時系列データ保存APIで保存された各シリーズの時系列データで最後に保存されたデータが返却される。時系列データ保存 (bulk upload)で保存されたデータは取得されない。| |thumbnail|最新のサムネイル画像のみ取得。画像保存APIで各シリーズに対してis_latest&#x3D;trueを指定して最後に保存された画像が返却される。| |image|最新のオリジナル画像のみ取得。画像保存APIで各シリーズに対してis_latest&#x3D;trueを指定して最後に保存された画像が返却される。| | [default to &quot;data&quot;]
 **acceptEncoding** | **string** | レスポンスボディの圧縮形式。gzipのみ対応。 | 

### Return type

[**MultipleCollectionLatest**](MultipleCollectionLatest.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSeries

> PostSeries(ctx, collectionId).SeriesRequest(seriesRequest).ResourceOwnerId(resourceOwnerId).Execute()

時系列データ保存



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
    collectionId := "location" // string | コレクション名
    seriesRequest := *openapiclient.NewSeriesRequest("SeriesId_example", int64(123)) // SeriesRequest | 
    resourceOwnerId := TODO // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionApi.PostSeries(context.Background(), collectionId).SeriesRequest(seriesRequest).ResourceOwnerId(resourceOwnerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.PostSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | コレクション名 | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seriesRequest** | [**SeriesRequest**](SeriesRequest.md) |  | 
 **resourceOwnerId** | [**string**](string.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSeriesBulk

> PostSeriesBulk(ctx, collectionId).SeriesBulkRequest(seriesBulkRequest).ResourceOwnerId(resourceOwnerId).Execute()





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
    collectionId := "location" // string | コレクション名
    seriesBulkRequest := *openapiclient.NewSeriesBulkRequest([]openapiclient.SeriesRequest{*openapiclient.NewSeriesRequest("SeriesId_example", int64(123))}) // SeriesBulkRequest | 
    resourceOwnerId := TODO // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionApi.PostSeriesBulk(context.Background(), collectionId).SeriesBulkRequest(seriesBulkRequest).ResourceOwnerId(resourceOwnerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.PostSeriesBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | コレクション名 | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSeriesBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seriesBulkRequest** | [**SeriesBulkRequest**](SeriesBulkRequest.md) |  | 
 **resourceOwnerId** | [**string**](string.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSeries

> PutSeries(ctx, collectionId, seriesId, timestamp).SeriesDataRequest(seriesDataRequest).ResourceOwnerId(resourceOwnerId).Execute()





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
    collectionId := "location" // string | コレクション名
    seriesId := "seriesId_example" // string | シリーズID
    timestamp := int64(1518165596566320) // int64 | UNIXタイム(ミリ秒)
    seriesDataRequest := *openapiclient.NewSeriesDataRequest() // SeriesDataRequest | 
    resourceOwnerId := TODO // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CollectionApi.PutSeries(context.Background(), collectionId, seriesId, timestamp).SeriesDataRequest(seriesDataRequest).ResourceOwnerId(resourceOwnerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.PutSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | コレクション名 | 
**seriesId** | **string** | シリーズID | 
**timestamp** | **int64** | UNIXタイム(ミリ秒) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **seriesDataRequest** | [**SeriesDataRequest**](SeriesDataRequest.md) |  | 
 **resourceOwnerId** | [**string**](string.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

