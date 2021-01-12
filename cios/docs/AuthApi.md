# \AuthApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RefreshToken**](AuthApi.md#RefreshToken) | **Post** /connect/token | 



## RefreshToken

> ConnectTokenResponse RefreshToken(ctx).GrantType(grantType).RefreshToken(refreshToken).ClientId(clientId).ClientSecret(clientSecret).Scope(scope).Execute()



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
    grantType := "grantType_example" // string |  (optional)
    refreshToken := "refreshToken_example" // string |  (optional)
    clientId := "clientId_example" // string |  (optional)
    clientSecret := "clientSecret_example" // string |  (optional)
    scope := "scope_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.RefreshToken(context.Background()).GrantType(grantType).RefreshToken(refreshToken).ClientId(clientId).ClientSecret(clientSecret).Scope(scope).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.RefreshToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshToken`: ConnectTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.RefreshToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** |  | 
 **refreshToken** | **string** |  | 
 **clientId** | **string** |  | 
 **clientSecret** | **string** |  | 
 **scope** | **string** |  | 

### Return type

[**ConnectTokenResponse**](ConnectTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

