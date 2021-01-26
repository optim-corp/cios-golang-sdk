# \ContractApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContract**](ContractApi.md#GetContract) | **Get** /v2/contracts/{contract_id} | 契約IDの契約情報を取得する
[**GetContractUser**](ContractApi.md#GetContractUser) | **Get** /v2/contracts/{contract_id}/users/{user_id} | 契約IDの契約情報を取得する
[**GetContractUsers**](ContractApi.md#GetContractUsers) | **Get** /v2/contracts/{contract_id}/users | 契約IDの契約情報を取得する
[**GetContracts**](ContractApi.md#GetContracts) | **Get** /v2/contracts | 契約一覧情報を取得する



## GetContract

> Contract GetContract(ctx).Page(page).Limit(limit).Offset(offset).Execute()

契約IDの契約情報を取得する



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
    page := "page_example" // string |  (optional)
    limit := int64(789) // int64 | ページネーションのうち、上位何件取得するか (optional)
    offset := int64(789) // int64 | ページネーションのうち、何件読み飛ばすか (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractApi.GetContract(context.Background()).Page(page).Limit(limit).Offset(offset).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.GetContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContract`: Contract
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.GetContract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** |  | 
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 

### Return type

[**Contract**](Contract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractUser

> ContractUser GetContractUser(ctx, contractId, userId).Execute()

契約IDの契約情報を取得する



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
    contractId := "contractId_example" // string | 
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractApi.GetContractUser(context.Background(), contractId, userId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.GetContractUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractUser`: ContractUser
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.GetContractUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ContractUser**](ContractUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractUsers

> MultipleContractUser GetContractUsers(ctx, contractId).Execute()

契約IDの契約情報を取得する



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
    contractId := "contractId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractApi.GetContractUsers(context.Background(), contractId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.GetContractUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractUsers`: MultipleContractUser
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.GetContractUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MultipleContractUser**](MultipleContractUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContracts

> MultipleContract GetContracts(ctx).Page(page).Limit(limit).Offset(offset).Execute()

契約一覧情報を取得する



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
    page := "page_example" // string |  (optional)
    limit := int64(789) // int64 | ページネーションのうち、上位何件取得するか (optional)
    offset := int64(789) // int64 | ページネーションのうち、何件読み飛ばすか (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractApi.GetContracts(context.Background()).Page(page).Limit(limit).Offset(offset).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.GetContracts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContracts`: MultipleContract
    fmt.Fprintf(os.Stdout, "Response from `ContractApi.GetContracts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContractsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** |  | 
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 

### Return type

[**MultipleContract**](MultipleContract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

