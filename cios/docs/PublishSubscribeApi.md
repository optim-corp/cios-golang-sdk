# \PublishSubscribeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChannel**](PublishSubscribeApi.md#CreateChannel) | **Post** /v2/channels | 
[**CreateDataStoreObject**](PublishSubscribeApi.md#CreateDataStoreObject) | **Post** /v2/datastore/channels/{channel_id}/objects | 
[**DeleteChannel**](PublishSubscribeApi.md#DeleteChannel) | **Delete** /v2/channels/{channel_id} | 
[**DeleteDataStoreChannel**](PublishSubscribeApi.md#DeleteDataStoreChannel) | **Delete** /v2/datastore/channels/{channel_id} | 
[**DeleteDataStoreObjectData**](PublishSubscribeApi.md#DeleteDataStoreObjectData) | **Delete** /v2/datastore/channels/{channel_id}/objects/{object_id} | 
[**DeleteDataStoreSession**](PublishSubscribeApi.md#DeleteDataStoreSession) | **Delete** /v2/datastore/channels/{channel_id}/sessions/{session_id} | 
[**GetChannel**](PublishSubscribeApi.md#GetChannel) | **Get** /v2/channels/{channel_id} | 
[**GetChannels**](PublishSubscribeApi.md#GetChannels) | **Get** /v2/channels | 
[**GetDataStoreChannel**](PublishSubscribeApi.md#GetDataStoreChannel) | **Get** /v2/datastore/channels/{channel_id} | 
[**GetDataStoreChannels**](PublishSubscribeApi.md#GetDataStoreChannels) | **Get** /v2/datastore/channels | 
[**GetDataStoreMultiObjectDataLatest**](PublishSubscribeApi.md#GetDataStoreMultiObjectDataLatest) | **Post** /v2/datastore/objects_latest | 
[**GetDataStoreObjectData**](PublishSubscribeApi.md#GetDataStoreObjectData) | **Get** /v2/datastore/channels/{channel_id}/objects/{object_id} | 
[**GetDataStoreObjectDataLatest**](PublishSubscribeApi.md#GetDataStoreObjectDataLatest) | **Get** /v2/datastore/channels/{channel_id}/object_latest | 
[**GetDataStoreObjects**](PublishSubscribeApi.md#GetDataStoreObjects) | **Get** /v2/datastore/channels/{channel_id}/objects | 
[**GetDataStoreSession**](PublishSubscribeApi.md#GetDataStoreSession) | **Get** /v2/datastore/channels/{channel_id}/sessions/{session_id} | 
[**GetDataStoreSessions**](PublishSubscribeApi.md#GetDataStoreSessions) | **Get** /v2/datastore/channels/{channel_id}/sessions | 
[**PublishMessage**](PublishSubscribeApi.md#PublishMessage) | **Post** /v2/messaging | 
[**SubscribeMessage**](PublishSubscribeApi.md#SubscribeMessage) | **Get** /v2/messaging | 
[**UpdateChannel**](PublishSubscribeApi.md#UpdateChannel) | **Patch** /v2/channels/{channel_id} | 



## CreateChannel

> SingleChannel CreateChannel(ctx).ChannelProposal(channelProposal).Execute()



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
    channelProposal := *openapiclient.NewChannelProposal("ResourceOwnerId_example", []openapiclient.DisplayInfo{*openapiclient.NewDisplayInfo("Name_example", "Language_example", false)}) // ChannelProposal | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublishSubscribeApi.CreateChannel(context.Background()).ChannelProposal(channelProposal).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishSubscribeApi.CreateChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateChannel`: SingleChannel
    fmt.Fprintf(os.Stdout, "Response from `PublishSubscribeApi.CreateChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelProposal** | [**ChannelProposal**](ChannelProposal.md) |  | 

### Return type

[**SingleChannel**](SingleChannel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDataStoreObject

> SingleDataStoreObject CreateDataStoreObject(ctx, channelId).Body(body).ChannelProtocolId(channelProtocolId).ChannelProtocolVersion(channelProtocolVersion).SessionId(sessionId).Location(location).Timestamp(timestamp).Execute()



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
    channelId := "channelId_example" // string | 
    body := os.NewFile(1234, "some_file") // *os.File | 
    channelProtocolId := "channelProtocolId_example" // string |  (optional)
    channelProtocolVersion := int32(56) // int32 |  (optional)
    sessionId := "sessionId_example" // string |  (optional)
    location := "location_example" // string |  (optional)
    timestamp := "timestamp_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublishSubscribeApi.CreateDataStoreObject(context.Background(), channelId).Body(body).ChannelProtocolId(channelProtocolId).ChannelProtocolVersion(channelProtocolVersion).SessionId(sessionId).Location(location).Timestamp(timestamp).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishSubscribeApi.CreateDataStoreObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDataStoreObject`: SingleDataStoreObject
    fmt.Fprintf(os.Stdout, "Response from `PublishSubscribeApi.CreateDataStoreObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataStoreObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** |  | 
 **channelProtocolId** | **string** |  | 
 **channelProtocolVersion** | **int32** |  | 
 **sessionId** | **string** |  | 
 **location** | **string** |  | 
 **timestamp** | **string** |  | 

### Return type

[**SingleDataStoreObject**](SingleDataStoreObject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChannel

> DeleteChannel(ctx, channelId).Execute()



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
    channelId := "channelId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublishSubscribeApi.DeleteChannel(context.Background(), channelId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishSubscribeApi.DeleteChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChannelRequest struct via the builder pattern


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


## DeleteDataStoreChannel

> DeleteDataStoreChannel(ctx, channelId).Execute()



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
    channelId := "channelId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublishSubscribeApi.DeleteDataStoreChannel(context.Background(), channelId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishSubscribeApi.DeleteDataStoreChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataStoreChannelRequest struct via the builder pattern


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


## DeleteDataStoreObjectData

> MultipleDataStoreObject DeleteDataStoreObjectData(ctx, channelId, objectId).Execute()



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
    channelId := "channelId_example" // string | 
    objectId := "objectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublishSubscribeApi.DeleteDataStoreObjectData(context.Background(), channelId, objectId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishSubscribeApi.DeleteDataStoreObjectData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDataStoreObjectData`: MultipleDataStoreObject
    fmt.Fprintf(os.Stdout, "Response from `PublishSubscribeApi.DeleteDataStoreObjectData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**objectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataStoreObjectDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MultipleDataStoreObject**](MultipleDataStoreObject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDataStoreSession

> DeleteDataStoreSession(ctx, channelId, sessionId).Execute()



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
    channelId := "channelId_example" // string | 
    sessionId := "sessionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublishSubscribeApi.DeleteDataStoreSession(context.Background(), channelId, sessionId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishSubscribeApi.DeleteDataStoreSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataStoreSessionRequest struct via the builder pattern


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


## GetChannel

> SingleChannel GetChannel(ctx, channelId).Lang(lang).IsDev(isDev).Execute()



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
    channelId := "channelId_example" // string | 
    lang := "lang_example" // string | 言語指定 (optional)
    isDev := true // bool | 開発者モードの有効・無効 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublishSubscribeApi.GetChannel(context.Background(), channelId).Lang(lang).IsDev(isDev).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishSubscribeApi.GetChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChannel`: SingleChannel
    fmt.Fprintf(os.Stdout, "Response from `PublishSubscribeApi.GetChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lang** | **string** | 言語指定 | 
 **isDev** | **bool** | 開発者モードの有効・無効 | 

### Return type

[**SingleChannel**](SingleChannel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannels

> MultipleChanel GetChannels(ctx).Name(name).ResourceOwnerId(resourceOwnerId).Label(label).ChannelProtocol(channelProtocol).Limit(limit).Offset(offset).Order(order).OrderBy(orderBy).Lang(lang).IsDev(isDev).IsPublic(isPublic).MessagingEnabled(messagingEnabled).DatastoreEnabled(datastoreEnabled).MessagingPersisted(messagingPersisted).Execute()



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
    name := "name_example" // string | 名前での部分一致検索 (optional)
    resourceOwnerId := TODO // string |  (optional)
    label := "label_example" // string |  (optional)
    channelProtocol := "channelProtocol_example" // string |  (optional)
    limit := int64(789) // int64 | ページネーションのうち、上位何件取得するか (optional)
    offset := int64(789) // int64 | ページネーションのうち、何件読み飛ばすか (optional)
    order := "order_example" // string | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 (optional)
    orderBy := "orderBy_example" // string | ソート対象。指定順に高優先でのソートとなる (optional)
    lang := "lang_example" // string | 言語指定 (optional)
    isDev := true // bool | 開発者モードの有効・無効 (optional)
    isPublic := "isPublic_example" // string | デバイスの公開・非公開 (optional)
    messagingEnabled := "messagingEnabled_example" // string |  (optional)
    datastoreEnabled := "datastoreEnabled_example" // string |  (optional)
    messagingPersisted := "messagingPersisted_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublishSubscribeApi.GetChannels(context.Background()).Name(name).ResourceOwnerId(resourceOwnerId).Label(label).ChannelProtocol(channelProtocol).Limit(limit).Offset(offset).Order(order).OrderBy(orderBy).Lang(lang).IsDev(isDev).IsPublic(isPublic).MessagingEnabled(messagingEnabled).DatastoreEnabled(datastoreEnabled).MessagingPersisted(messagingPersisted).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishSubscribeApi.GetChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChannels`: MultipleChanel
    fmt.Fprintf(os.Stdout, "Response from `PublishSubscribeApi.GetChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | 名前での部分一致検索 | 
 **resourceOwnerId** | [**string**](string.md) |  | 
 **label** | **string** |  | 
 **channelProtocol** | **string** |  | 
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 
 **order** | **string** | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 | 
 **orderBy** | **string** | ソート対象。指定順に高優先でのソートとなる | 
 **lang** | **string** | 言語指定 | 
 **isDev** | **bool** | 開発者モードの有効・無効 | 
 **isPublic** | **string** | デバイスの公開・非公開 | 
 **messagingEnabled** | **string** |  | 
 **datastoreEnabled** | **string** |  | 
 **messagingPersisted** | **string** |  | 

### Return type

[**MultipleChanel**](MultipleChanel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataStoreChannel

> SingleDataStoreChannel GetDataStoreChannel(ctx, channelId).Execute()



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
    channelId := "channelId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublishSubscribeApi.GetDataStoreChannel(context.Background(), channelId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishSubscribeApi.GetDataStoreChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataStoreChannel`: SingleDataStoreChannel
    fmt.Fprintf(os.Stdout, "Response from `PublishSubscribeApi.GetDataStoreChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataStoreChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SingleDataStoreChannel**](SingleDataStoreChannel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataStoreChannels

> MultipleDataStoreChannel GetDataStoreChannels(ctx).ChannelProtocolId(channelProtocolId).Limit(limit).Offset(offset).Order(order).OrderBy(orderBy).Execute()



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
    channelProtocolId := "channelProtocolId_example" // string |  (optional)
    limit := int64(789) // int64 | ページネーションのうち、上位何件取得するか (optional)
    offset := int64(789) // int64 | ページネーションのうち、何件読み飛ばすか (optional)
    order := "order_example" // string | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 (optional)
    orderBy := "orderBy_example" // string | ソート対象。指定順に高優先でのソートとなる (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublishSubscribeApi.GetDataStoreChannels(context.Background()).ChannelProtocolId(channelProtocolId).Limit(limit).Offset(offset).Order(order).OrderBy(orderBy).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishSubscribeApi.GetDataStoreChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataStoreChannels`: MultipleDataStoreChannel
    fmt.Fprintf(os.Stdout, "Response from `PublishSubscribeApi.GetDataStoreChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataStoreChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelProtocolId** | **string** |  | 
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 
 **order** | **string** | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 | 
 **orderBy** | **string** | ソート対象。指定順に高優先でのソートとなる | 

### Return type

[**MultipleDataStoreChannel**](MultipleDataStoreChannel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataStoreMultiObjectDataLatest

> MultipleDataStoreDataLatest GetDataStoreMultiObjectDataLatest(ctx).Ids(ids).Execute()



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
    ids := *openapiclient.NewIds() // Ids | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublishSubscribeApi.GetDataStoreMultiObjectDataLatest(context.Background()).Ids(ids).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishSubscribeApi.GetDataStoreMultiObjectDataLatest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataStoreMultiObjectDataLatest`: MultipleDataStoreDataLatest
    fmt.Fprintf(os.Stdout, "Response from `PublishSubscribeApi.GetDataStoreMultiObjectDataLatest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataStoreMultiObjectDataLatestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**Ids**](Ids.md) |  | 

### Return type

[**MultipleDataStoreDataLatest**](MultipleDataStoreDataLatest.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataStoreObjectData

> interface{} GetDataStoreObjectData(ctx, channelId, objectId).PackerFormat(packerFormat).Execute()



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
    channelId := "channelId_example" // string | 
    objectId := "objectId_example" // string | 
    packerFormat := "packerFormat_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublishSubscribeApi.GetDataStoreObjectData(context.Background(), channelId, objectId).PackerFormat(packerFormat).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishSubscribeApi.GetDataStoreObjectData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataStoreObjectData`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `PublishSubscribeApi.GetDataStoreObjectData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**objectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataStoreObjectDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **packerFormat** | **string** |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataStoreObjectDataLatest

> interface{} GetDataStoreObjectDataLatest(ctx, channelId).PackerFormat(packerFormat).Execute()



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
    channelId := "channelId_example" // string | 
    packerFormat := "packerFormat_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublishSubscribeApi.GetDataStoreObjectDataLatest(context.Background(), channelId).PackerFormat(packerFormat).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishSubscribeApi.GetDataStoreObjectDataLatest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataStoreObjectDataLatest`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `PublishSubscribeApi.GetDataStoreObjectDataLatest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataStoreObjectDataLatestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **packerFormat** | **string** |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/octet-stream, application/xml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataStoreObjects

> MultipleDataStoreObject GetDataStoreObjects(ctx, channelId).SessionId(sessionId).ChannelProtocolVersion(channelProtocolVersion).ChannelProtocolId(channelProtocolId).Label(label).Location(location).LocationRange(locationRange).Timestamp(timestamp).TimestampRange(timestampRange).Ascending(ascending).Offset(offset).Limit(limit).Execute()



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
    channelId := "channelId_example" // string | 
    sessionId := "sessionId_example" // string |  (optional)
    channelProtocolVersion := int32(56) // int32 |  (optional)
    channelProtocolId := "channelProtocolId_example" // string |  (optional)
    label := "label_example" // string |  (optional)
    location := "location_example" // string |  (optional)
    locationRange := "locationRange_example" // string |  (optional)
    timestamp := "timestamp_example" // string |  (optional)
    timestampRange := "timestampRange_example" // string |  (optional)
    ascending := true // bool |  (optional)
    offset := int64(789) // int64 | ページネーションのうち、何件読み飛ばすか (optional)
    limit := int64(789) // int64 | ページネーションのうち、上位何件取得するか (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublishSubscribeApi.GetDataStoreObjects(context.Background(), channelId).SessionId(sessionId).ChannelProtocolVersion(channelProtocolVersion).ChannelProtocolId(channelProtocolId).Label(label).Location(location).LocationRange(locationRange).Timestamp(timestamp).TimestampRange(timestampRange).Ascending(ascending).Offset(offset).Limit(limit).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishSubscribeApi.GetDataStoreObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataStoreObjects`: MultipleDataStoreObject
    fmt.Fprintf(os.Stdout, "Response from `PublishSubscribeApi.GetDataStoreObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataStoreObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sessionId** | **string** |  | 
 **channelProtocolVersion** | **int32** |  | 
 **channelProtocolId** | **string** |  | 
 **label** | **string** |  | 
 **location** | **string** |  | 
 **locationRange** | **string** |  | 
 **timestamp** | **string** |  | 
 **timestampRange** | **string** |  | 
 **ascending** | **bool** |  | 
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 

### Return type

[**MultipleDataStoreObject**](MultipleDataStoreObject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataStoreSession

> SingleSession GetDataStoreSession(ctx, channelId, sessionId).Execute()



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
    channelId := "channelId_example" // string | 
    sessionId := "sessionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublishSubscribeApi.GetDataStoreSession(context.Background(), channelId, sessionId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishSubscribeApi.GetDataStoreSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataStoreSession`: SingleSession
    fmt.Fprintf(os.Stdout, "Response from `PublishSubscribeApi.GetDataStoreSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataStoreSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SingleSession**](SingleSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataStoreSessions

> MultipleSession GetDataStoreSessions(ctx, channelId).Timestamp(timestamp).TimestampRange(timestampRange).LocationRange(locationRange).Location(location).Ascending(ascending).Offset(offset).Limit(limit).Execute()



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
    channelId := "channelId_example" // string | 
    timestamp := "timestamp_example" // string |  (optional)
    timestampRange := "timestampRange_example" // string |  (optional)
    locationRange := "locationRange_example" // string |  (optional)
    location := "location_example" // string |  (optional)
    ascending := true // bool |  (optional)
    offset := int64(789) // int64 | ページネーションのうち、何件読み飛ばすか (optional)
    limit := int64(789) // int64 | ページネーションのうち、上位何件取得するか (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublishSubscribeApi.GetDataStoreSessions(context.Background(), channelId).Timestamp(timestamp).TimestampRange(timestampRange).LocationRange(locationRange).Location(location).Ascending(ascending).Offset(offset).Limit(limit).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishSubscribeApi.GetDataStoreSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataStoreSessions`: MultipleSession
    fmt.Fprintf(os.Stdout, "Response from `PublishSubscribeApi.GetDataStoreSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataStoreSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timestamp** | **string** |  | 
 **timestampRange** | **string** |  | 
 **locationRange** | **string** |  | 
 **location** | **string** |  | 
 **ascending** | **bool** |  | 
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 

### Return type

[**MultipleSession**](MultipleSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishMessage

> PublishMessage(ctx).ChannelId(channelId).PackerFormat(packerFormat).Body(body).Execute()



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
    channelId := "channelId_example" // string | 
    packerFormat := "packerFormat_example" // string |  (optional)
    body := interface{}(987) // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublishSubscribeApi.PublishMessage(context.Background()).ChannelId(channelId).PackerFormat(packerFormat).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishSubscribeApi.PublishMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **string** |  | 
 **packerFormat** | **string** |  | 
 **body** | **interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/octet-stream, application/xml, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribeMessage

> *os.File SubscribeMessage(ctx).ChannelId(channelId).PackerFormat(packerFormat).Execute()



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
    channelId := "channelId_example" // string | 
    packerFormat := "packerFormat_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublishSubscribeApi.SubscribeMessage(context.Background()).ChannelId(channelId).PackerFormat(packerFormat).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishSubscribeApi.SubscribeMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscribeMessage`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `PublishSubscribeApi.SubscribeMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **string** |  | 
 **packerFormat** | **string** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannel

> MultipleChanel UpdateChannel(ctx, channelId).ChannelUpdateProposal(channelUpdateProposal).Execute()



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
    channelId := "channelId_example" // string | 
    channelUpdateProposal := *openapiclient.NewChannelUpdateProposal([]openapiclient.DisplayInfo{*openapiclient.NewDisplayInfo("Name_example", "Language_example", false)}) // ChannelUpdateProposal | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublishSubscribeApi.UpdateChannel(context.Background(), channelId).ChannelUpdateProposal(channelUpdateProposal).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PublishSubscribeApi.UpdateChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateChannel`: MultipleChanel
    fmt.Fprintf(os.Stdout, "Response from `PublishSubscribeApi.UpdateChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channelUpdateProposal** | [**ChannelUpdateProposal**](ChannelUpdateProposal.md) |  | 

### Return type

[**MultipleChanel**](MultipleChanel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

