# \FileStorageApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyNode**](FileStorageApi.md#CopyNode) | **Post** /v2/file_storage/buckets/{bucket_id}/nodes/{node_id}/copy | 
[**CreateBucket**](FileStorageApi.md#CreateBucket) | **Post** /v2/file_storage/buckets | 
[**CreateDirectory**](FileStorageApi.md#CreateDirectory) | **Post** /v2/file_storage/buckets/{bucket_id}/create_directory | 
[**DeleteBucket**](FileStorageApi.md#DeleteBucket) | **Delete** /v2/file_storage/buckets/{bucket_id} | 
[**DeleteNode**](FileStorageApi.md#DeleteNode) | **Delete** /v2/file_storage/buckets/{bucket_id}/nodes/{node_id} | 
[**DownloadFile**](FileStorageApi.md#DownloadFile) | **Get** /v2/file_storage/buckets/{bucket_id}/download | 
[**GetBucket**](FileStorageApi.md#GetBucket) | **Get** /v2/file_storage/buckets/{bucket_id} | 
[**GetBuckets**](FileStorageApi.md#GetBuckets) | **Get** /v2/file_storage/buckets | 
[**GetNode**](FileStorageApi.md#GetNode) | **Get** /v2/file_storage/buckets/{bucket_id}/nodes/{node_id} | 
[**GetNodes**](FileStorageApi.md#GetNodes) | **Get** /v2/file_storage/buckets/{bucket_id}/nodes | 
[**MoveNode**](FileStorageApi.md#MoveNode) | **Post** /v2/file_storage/buckets/{bucket_id}/nodes/{node_id}/move | 
[**RenameNode**](FileStorageApi.md#RenameNode) | **Post** /v2/file_storage/buckets/{bucket_id}/nodes/{node_id}/rename | 
[**UpdateBucket**](FileStorageApi.md#UpdateBucket) | **Patch** /v2/file_storage/buckets/{bucket_id} | 
[**UploadFile**](FileStorageApi.md#UploadFile) | **Put** /v2/file_storage/buckets/{bucket_id}/upload | 



## CopyNode

> SingleNode CopyNode(ctx, bucketId, nodeId).BucketEditBody(bucketEditBody).Execute()



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
    bucketId := "bucketId_example" // string | 
    nodeId := "nodeId_example" // string | 
    bucketEditBody := *openapiclient.NewBucketEditBody() // BucketEditBody | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileStorageApi.CopyNode(context.Background(), bucketId, nodeId).BucketEditBody(bucketEditBody).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FileStorageApi.CopyNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopyNode`: SingleNode
    fmt.Fprintf(os.Stdout, "Response from `FileStorageApi.CopyNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **string** |  | 
**nodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bucketEditBody** | [**BucketEditBody**](BucketEditBody.md) |  | 

### Return type

[**SingleNode**](SingleNode.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBucket

> SingleBucket CreateBucket(ctx).BucketRequest(bucketRequest).Execute()





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
    bucketRequest := *openapiclient.NewBucketRequest("ResourceOwnerId_example", "Name_example") // BucketRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileStorageApi.CreateBucket(context.Background()).BucketRequest(bucketRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FileStorageApi.CreateBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBucket`: SingleBucket
    fmt.Fprintf(os.Stdout, "Response from `FileStorageApi.CreateBucket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketRequest** | [**BucketRequest**](BucketRequest.md) |  | 

### Return type

[**SingleBucket**](SingleBucket.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDirectory

> SingleNode CreateDirectory(ctx, bucketId).NodeRequest(nodeRequest).Execute()





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
    bucketId := "bucketId_example" // string | 
    nodeRequest := *openapiclient.NewNodeRequest("Name_example") // NodeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileStorageApi.CreateDirectory(context.Background(), bucketId).NodeRequest(nodeRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FileStorageApi.CreateDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDirectory`: SingleNode
    fmt.Fprintf(os.Stdout, "Response from `FileStorageApi.CreateDirectory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeRequest** | [**NodeRequest**](NodeRequest.md) |  | 

### Return type

[**SingleNode**](SingleNode.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBucket

> DeleteBucket(ctx, bucketId).Execute()





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
    bucketId := "bucketId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileStorageApi.DeleteBucket(context.Background(), bucketId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FileStorageApi.DeleteBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBucketRequest struct via the builder pattern


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


## DeleteNode

> DeleteNode(ctx, bucketId, nodeId).DirectorySize(directorySize).Execute()





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
    bucketId := "bucketId_example" // string | 
    nodeId := "nodeId_example" // string | 
    directorySize := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileStorageApi.DeleteNode(context.Background(), bucketId, nodeId).DirectorySize(directorySize).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FileStorageApi.DeleteNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **string** |  | 
**nodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **directorySize** | **bool** |  | 

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


## DownloadFile

> string DownloadFile(ctx, bucketId).NodeId(nodeId).Key(key).Execute()





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
    bucketId := "bucketId_example" // string | 
    nodeId := "nodeId_example" // string |  (optional)
    key := "key_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileStorageApi.DownloadFile(context.Background(), bucketId).NodeId(nodeId).Key(key).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FileStorageApi.DownloadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadFile`: string
    fmt.Fprintf(os.Stdout, "Response from `FileStorageApi.DownloadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeId** | **string** |  | 
 **key** | **string** |  | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: /_*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBucket

> SingleBucket GetBucket(ctx, bucketId).Execute()





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
    bucketId := "bucketId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileStorageApi.GetBucket(context.Background(), bucketId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FileStorageApi.GetBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBucket`: SingleBucket
    fmt.Fprintf(os.Stdout, "Response from `FileStorageApi.GetBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SingleBucket**](SingleBucket.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuckets

> MultipleBucket GetBuckets(ctx).Limit(limit).Offset(offset).Order(order).OrderBy(orderBy).ResourceOwnerId(resourceOwnerId).Name(name).Execute()





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
    limit := int64(789) // int64 | ページネーションのうち、上位何件取得するか (optional)
    offset := int64(789) // int64 | ページネーションのうち、何件読み飛ばすか (optional)
    order := "order_example" // string | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 (optional)
    orderBy := "orderBy_example" // string | ソート対象。指定順に高優先でのソートとなる (optional)
    resourceOwnerId := TODO // string |  (optional)
    name := "name_example" // string | 名前での部分一致検索 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileStorageApi.GetBuckets(context.Background()).Limit(limit).Offset(offset).Order(order).OrderBy(orderBy).ResourceOwnerId(resourceOwnerId).Name(name).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FileStorageApi.GetBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuckets`: MultipleBucket
    fmt.Fprintf(os.Stdout, "Response from `FileStorageApi.GetBuckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 
 **order** | **string** | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 | 
 **orderBy** | **string** | ソート対象。指定順に高優先でのソートとなる | 
 **resourceOwnerId** | [**string**](string.md) |  | 
 **name** | **string** | 名前での部分一致検索 | 

### Return type

[**MultipleBucket**](MultipleBucket.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNode

> SingleNode GetNode(ctx, bucketId, nodeId).DirectorySize(directorySize).Execute()





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
    bucketId := "bucketId_example" // string | 
    nodeId := "nodeId_example" // string | 
    directorySize := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileStorageApi.GetNode(context.Background(), bucketId, nodeId).DirectorySize(directorySize).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FileStorageApi.GetNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNode`: SingleNode
    fmt.Fprintf(os.Stdout, "Response from `FileStorageApi.GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **string** |  | 
**nodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **directorySize** | **bool** |  | 

### Return type

[**SingleNode**](SingleNode.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodes

> MultipleNode GetNodes(ctx, bucketId).ParentNodeId(parentNodeId).Recursive(recursive).Name(name).Key(key).IsDirectory(isDirectory).Limit(limit).Offset(offset).OrderBy(orderBy).Order(order).Execute()





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
    bucketId := "bucketId_example" // string | 
    parentNodeId := "parentNodeId_example" // string |  (optional)
    recursive := true // bool |  (optional)
    name := "name_example" // string | 名前での部分一致検索 (optional)
    key := "key_example" // string |  (optional)
    isDirectory := true // bool |  (optional)
    limit := int64(789) // int64 | ページネーションのうち、上位何件取得するか (optional)
    offset := int64(789) // int64 | ページネーションのうち、何件読み飛ばすか (optional)
    orderBy := "orderBy_example" // string |  (optional)
    order := "order_example" // string | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileStorageApi.GetNodes(context.Background(), bucketId).ParentNodeId(parentNodeId).Recursive(recursive).Name(name).Key(key).IsDirectory(isDirectory).Limit(limit).Offset(offset).OrderBy(orderBy).Order(order).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FileStorageApi.GetNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodes`: MultipleNode
    fmt.Fprintf(os.Stdout, "Response from `FileStorageApi.GetNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentNodeId** | **string** |  | 
 **recursive** | **bool** |  | 
 **name** | **string** | 名前での部分一致検索 | 
 **key** | **string** |  | 
 **isDirectory** | **bool** |  | 
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 
 **orderBy** | **string** |  | 
 **order** | **string** | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 | 

### Return type

[**MultipleNode**](MultipleNode.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveNode

> SingleNode MoveNode(ctx, bucketId, nodeId).BucketEditBody(bucketEditBody).Execute()



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
    bucketId := "bucketId_example" // string | 
    nodeId := "nodeId_example" // string | 
    bucketEditBody := *openapiclient.NewBucketEditBody() // BucketEditBody | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileStorageApi.MoveNode(context.Background(), bucketId, nodeId).BucketEditBody(bucketEditBody).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FileStorageApi.MoveNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveNode`: SingleNode
    fmt.Fprintf(os.Stdout, "Response from `FileStorageApi.MoveNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **string** |  | 
**nodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bucketEditBody** | [**BucketEditBody**](BucketEditBody.md) |  | 

### Return type

[**SingleNode**](SingleNode.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameNode

> SingleNode RenameNode(ctx, bucketId, nodeId).NodeName(nodeName).Execute()



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
    bucketId := "bucketId_example" // string | 
    nodeId := "nodeId_example" // string | 
    nodeName := *openapiclient.NewNodeName("Name_example") // NodeName | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileStorageApi.RenameNode(context.Background(), bucketId, nodeId).NodeName(nodeName).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FileStorageApi.RenameNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenameNode`: SingleNode
    fmt.Fprintf(os.Stdout, "Response from `FileStorageApi.RenameNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **string** |  | 
**nodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nodeName** | [**NodeName**](NodeName.md) |  | 

### Return type

[**SingleNode**](SingleNode.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBucket

> UpdateBucket(ctx, bucketId).BucketName(bucketName).Execute()



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
    bucketId := "bucketId_example" // string | 
    bucketName := *openapiclient.NewBucketName("Name_example") // BucketName | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileStorageApi.UpdateBucket(context.Background(), bucketId).BucketName(bucketName).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FileStorageApi.UpdateBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketName** | [**BucketName**](BucketName.md) |  | 

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


## UploadFile

> SingleBucket UploadFile(ctx, bucketId).Checksum(checksum).Body(body).Name(name).NodeId(nodeId).Key(key).ParentNodeId(parentNodeId).ParentKey(parentKey).Force(force).Execute()





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
    bucketId := "bucketId_example" // string | 
    checksum := "checksum_example" // string | 
    body := string(BYTE_ARRAY_DATA_HERE) // string | 
    name := "name_example" // string | 名前での部分一致検索 (optional)
    nodeId := "nodeId_example" // string |  (optional)
    key := "key_example" // string |  (optional)
    parentNodeId := "parentNodeId_example" // string |  (optional)
    parentKey := "parentKey_example" // string |  (optional)
    force := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileStorageApi.UploadFile(context.Background(), bucketId).Checksum(checksum).Body(body).Name(name).NodeId(nodeId).Key(key).ParentNodeId(parentNodeId).ParentKey(parentKey).Force(force).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `FileStorageApi.UploadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFile`: SingleBucket
    fmt.Fprintf(os.Stdout, "Response from `FileStorageApi.UploadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checksum** | **string** |  | 
 **body** | **string** |  | 
 **name** | **string** | 名前での部分一致検索 | 
 **nodeId** | **string** |  | 
 **key** | **string** |  | 
 **parentNodeId** | **string** |  | 
 **parentKey** | **string** |  | 
 **force** | **bool** |  | 

### Return type

[**SingleBucket**](SingleBucket.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

