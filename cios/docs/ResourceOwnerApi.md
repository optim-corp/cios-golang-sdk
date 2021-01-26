# \ResourceOwnerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetResourceOwner**](ResourceOwnerApi.md#GetResourceOwner) | **Get** /v2/resource_owners/{resource_owner_id} | 
[**GetResourceOwners**](ResourceOwnerApi.md#GetResourceOwners) | **Get** /v2/resource_owners | 



## GetResourceOwner

> ResourceOwner GetResourceOwner(ctx, resourceOwnerId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceOwnerApi.GetResourceOwner(context.Background(), resourceOwnerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceOwnerApi.GetResourceOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceOwner`: ResourceOwner
    fmt.Fprintf(os.Stdout, "Response from `ResourceOwnerApi.GetResourceOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceOwnerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourceOwner**](ResourceOwner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceOwners

> MultipleResourceOwner GetResourceOwners(ctx).GroupId(groupId).UserId(userId).Type_(type_).Page(page).Offset(offset).Limit(limit).OrderBy(orderBy).Order(order).Execute()



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
    groupId := "groupId_example" // string |  (optional)
    userId := TODO // string |  (optional)
    type_ := "type__example" // string |  (optional)
    page := "page_example" // string |  (optional)
    offset := int64(789) // int64 | ページネーションのうち、何件読み飛ばすか (optional)
    limit := int64(789) // int64 | ページネーションのうち、上位何件取得するか (optional)
    orderBy := "orderBy_example" // string | ソート対象。指定順に高優先でのソートとなる (optional)
    order := "order_example" // string | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceOwnerApi.GetResourceOwners(context.Background()).GroupId(groupId).UserId(userId).Type_(type_).Page(page).Offset(offset).Limit(limit).OrderBy(orderBy).Order(order).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceOwnerApi.GetResourceOwners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceOwners`: MultipleResourceOwner
    fmt.Fprintf(os.Stdout, "Response from `ResourceOwnerApi.GetResourceOwners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceOwnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string** |  | 
 **userId** | [**string**](string.md) |  | 
 **type_** | **string** |  | 
 **page** | **string** |  | 
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 
 **orderBy** | **string** | ソート対象。指定順に高優先でのソートとなる | 
 **order** | **string** | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 | 

### Return type

[**MultipleResourceOwner**](MultipleResourceOwner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

