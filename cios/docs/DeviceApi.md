# \DeviceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDevice**](DeviceApi.md#CreateDevice) | **Post** /v2/devices | デバイスの登録
[**CreateDeviceClient**](DeviceApi.md#CreateDeviceClient) | **Post** /v2/devices/{device_id}/device_clients | 
[**CreateDevicePolicy**](DeviceApi.md#CreateDevicePolicy) | **Post** /v2/devices/group_policies | 
[**DeleteDevice**](DeviceApi.md#DeleteDevice) | **Delete** /v2/devices/{device_id} | 
[**DeleteDeviceClient**](DeviceApi.md#DeleteDeviceClient) | **Delete** /v2/devices/{device_id}/device_clients/{client_id} | device.write
[**DeletePolicy**](DeviceApi.md#DeletePolicy) | **Delete** /v2/devices/group_policies/{policy_id} | 
[**GetDevice**](DeviceApi.md#GetDevice) | **Get** /v2/devices/{device_id} | 指定したidのデバイスの情報を取得する
[**GetDeviceInventoryLatest**](DeviceApi.md#GetDeviceInventoryLatest) | **Get** /v2/devices/{device_id}/inventory/latest | 
[**GetDeviceMonitoringLatest**](DeviceApi.md#GetDeviceMonitoringLatest) | **Get** /v2/devices/{device_id}/monitoring/latest | 
[**GetDeviceMonitoringsLatest**](DeviceApi.md#GetDeviceMonitoringsLatest) | **Post** /v2/devices/monitoring/latest | 
[**GetDevicePolicies**](DeviceApi.md#GetDevicePolicies) | **Get** /v2/devices/group_policies | 
[**GetDeviceProfile**](DeviceApi.md#GetDeviceProfile) | **Get** /v2/devices/profile | device.profile
[**GetDevices**](DeviceApi.md#GetDevices) | **Get** /v2/devices | デバイスの一覧を取得する
[**UpdateDevice**](DeviceApi.md#UpdateDevice) | **Patch** /v2/devices/{device_id} | 指定したidのデバイス情報を更新する
[**UpdateDeviceClient**](DeviceApi.md#UpdateDeviceClient) | **Patch** /v2/devices/{device_id}/device_clients/{client_id} | device.write



## CreateDevice

> SingleDevice CreateDevice(ctx).DeviceInfo(deviceInfo).Execute()

デバイスの登録



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
    deviceInfo := *openapiclient.NewDeviceInfo("ResourceOwnerId_example") // DeviceInfo | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.CreateDevice(context.Background()).DeviceInfo(deviceInfo).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.CreateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDevice`: SingleDevice
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.CreateDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceInfo** | [**DeviceInfo**](DeviceInfo.md) |  | 

### Return type

[**SingleDevice**](SingleDevice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceClient

> DeviceClientList CreateDeviceClient(ctx, deviceId).DeviceClientRequest(deviceClientRequest).Execute()





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
    deviceId := "deviceId_example" // string | 
    deviceClientRequest := *openapiclient.NewDeviceClientRequest() // DeviceClientRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.CreateDeviceClient(context.Background(), deviceId).DeviceClientRequest(deviceClientRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.CreateDeviceClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceClient`: DeviceClientList
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.CreateDeviceClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceClientRequest** | [**DeviceClientRequest**](DeviceClientRequest.md) |  | 

### Return type

[**DeviceClientList**](DeviceClientList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDevicePolicy

> DevicePolicy CreateDevicePolicy(ctx).DevicePolicyRequest(devicePolicyRequest).Execute()



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
    devicePolicyRequest := *openapiclient.NewDevicePolicyRequest("ResourceOwnerId_example") // DevicePolicyRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.CreateDevicePolicy(context.Background()).DevicePolicyRequest(devicePolicyRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.CreateDevicePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDevicePolicy`: DevicePolicy
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.CreateDevicePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDevicePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **devicePolicyRequest** | [**DevicePolicyRequest**](DevicePolicyRequest.md) |  | 

### Return type

[**DevicePolicy**](DevicePolicy.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDevice

> DeleteDevice(ctx, deviceId).Execute()



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
    deviceId := "deviceId_example" // string | デバイスID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.DeleteDevice(context.Background(), deviceId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeleteDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | デバイスID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceRequest struct via the builder pattern


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


## DeleteDeviceClient

> DeleteDeviceClient(ctx, deviceId, clientId).Execute()

device.write



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
    deviceId := "deviceId_example" // string | 
    clientId := "clientId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.DeleteDeviceClient(context.Background(), deviceId, clientId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeleteDeviceClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 
**clientId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicy

> DeletePolicy(ctx, policyId).Execute()



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
    policyId := "policyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.DeletePolicy(context.Background(), policyId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeletePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern


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


## GetDevice

> SingleDevice GetDevice(ctx, deviceId).Lang(lang).IsDev(isDev).Execute()

指定したidのデバイスの情報を取得する



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
    deviceId := "deviceId_example" // string | デバイスID
    lang := "lang_example" // string | 言語指定 (optional)
    isDev := true // bool | 開発者モードの有効・無効 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.GetDevice(context.Background(), deviceId).Lang(lang).IsDev(isDev).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevice`: SingleDevice
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | デバイスID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lang** | **string** | 言語指定 | 
 **isDev** | **bool** | 開発者モードの有効・無効 | 

### Return type

[**SingleDevice**](SingleDevice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceInventoryLatest

> map[string]interface{} GetDeviceInventoryLatest(ctx, deviceId).Execute()



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
    deviceId := "deviceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.GetDeviceInventoryLatest(context.Background(), deviceId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDeviceInventoryLatest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceInventoryLatest`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDeviceInventoryLatest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceInventoryLatestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceMonitoringLatest

> SingleDeviceMonitoring GetDeviceMonitoringLatest(ctx, deviceId).Execute()



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
    deviceId := "deviceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.GetDeviceMonitoringLatest(context.Background(), deviceId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDeviceMonitoringLatest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceMonitoringLatest`: SingleDeviceMonitoring
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDeviceMonitoringLatest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceMonitoringLatestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SingleDeviceMonitoring**](SingleDeviceMonitoring.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceMonitoringsLatest

> MultipleDeviceMonitoring GetDeviceMonitoringsLatest(ctx).DeviceMonitoringIDsRequest(deviceMonitoringIDsRequest).Execute()



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
    deviceMonitoringIDsRequest := *openapiclient.NewDeviceMonitoringIDsRequest([]string{"DeviceIds_example"}) // DeviceMonitoringIDsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.GetDeviceMonitoringsLatest(context.Background()).DeviceMonitoringIDsRequest(deviceMonitoringIDsRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDeviceMonitoringsLatest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceMonitoringsLatest`: MultipleDeviceMonitoring
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDeviceMonitoringsLatest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceMonitoringsLatestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceMonitoringIDsRequest** | [**DeviceMonitoringIDsRequest**](DeviceMonitoringIDsRequest.md) |  | 

### Return type

[**MultipleDeviceMonitoring**](MultipleDeviceMonitoring.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicePolicies

> MultipleDevicePolicy GetDevicePolicies(ctx).Limit(limit).Offset(offset).Order(order).OrderBy(orderBy).ResourceOwnerId(resourceOwnerId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.GetDevicePolicies(context.Background()).Limit(limit).Offset(offset).Order(order).OrderBy(orderBy).ResourceOwnerId(resourceOwnerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDevicePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevicePolicies`: MultipleDevicePolicy
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDevicePolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 
 **order** | **string** | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 | 
 **orderBy** | **string** | ソート対象。指定順に高優先でのソートとなる | 
 **resourceOwnerId** | [**string**](string.md) |  | 

### Return type

[**MultipleDevicePolicy**](MultipleDevicePolicy.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceProfile

> SingleDevice GetDeviceProfile(ctx).Lang(lang).IsDev(isDev).Execute()

device.profile

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
    lang := "lang_example" // string | 言語指定 (optional)
    isDev := true // bool | 開発者モードの有効・無効 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.GetDeviceProfile(context.Background()).Lang(lang).IsDev(isDev).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDeviceProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceProfile`: SingleDevice
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDeviceProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lang** | **string** | 言語指定 | 
 **isDev** | **bool** | 開発者モードの有効・無効 | 

### Return type

[**SingleDevice**](SingleDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevices

> MultipleDevice GetDevices(ctx).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).Lang(lang).IsDev(isDev).IsPublic(isPublic).ResourceOwnerId(resourceOwnerId).Name(name).IdNumber(idNumber).DefinitionId(definitionId).DefinitionTag(definitionTag).Execute()

デバイスの一覧を取得する



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
    offset := int64(789) // int64 | ページネーションのうち、何件読み飛ばすか (optional)
    limit := int64(789) // int64 | ページネーションのうち、上位何件取得するか (optional)
    order := "order_example" // string | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 (optional)
    orderBy := "orderBy_example" // string | ソート対象。指定順に高優先でのソートとなる (optional)
    lang := "lang_example" // string | 言語指定 (optional)
    isDev := true // bool | 開発者モードの有効・無効 (optional)
    isPublic := "isPublic_example" // string | デバイスの公開・非公開 (optional)
    resourceOwnerId := TODO // string |  (optional)
    name := "name_example" // string | 名前での部分一致検索 (optional)
    idNumber := "idNumber_example" // string | 識別番号での完全一致検索 (optional)
    definitionId := "definitionId_example" // string | デバイス定義ID(custom_inventory.definition_id)での完全一致検索 (optional)
    definitionTag := "definitionTag_example" // string | custominventoryの内部データでの検索時に指定。 key,valueのセットの完全一致での検索とし、本パラメータでvalueを指定し、 inventory_keyでkeyを指定する。 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.GetDevices(context.Background()).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).Lang(lang).IsDev(isDev).IsPublic(isPublic).ResourceOwnerId(resourceOwnerId).Name(name).IdNumber(idNumber).DefinitionId(definitionId).DefinitionTag(definitionTag).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevices`: MultipleDevice
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 
 **order** | **string** | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 | 
 **orderBy** | **string** | ソート対象。指定順に高優先でのソートとなる | 
 **lang** | **string** | 言語指定 | 
 **isDev** | **bool** | 開発者モードの有効・無効 | 
 **isPublic** | **string** | デバイスの公開・非公開 | 
 **resourceOwnerId** | [**string**](string.md) |  | 
 **name** | **string** | 名前での部分一致検索 | 
 **idNumber** | **string** | 識別番号での完全一致検索 | 
 **definitionId** | **string** | デバイス定義ID(custom_inventory.definition_id)での完全一致検索 | 
 **definitionTag** | **string** | custominventoryの内部データでの検索時に指定。 key,valueのセットの完全一致での検索とし、本パラメータでvalueを指定し、 inventory_keyでkeyを指定する。 | 

### Return type

[**MultipleDevice**](MultipleDevice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevice

> SingleDevice UpdateDevice(ctx, deviceId).DeviceUpdateRequest(deviceUpdateRequest).Execute()

指定したidのデバイス情報を更新する



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
    deviceId := "deviceId_example" // string | デバイスID
    deviceUpdateRequest := *openapiclient.NewDeviceUpdateRequest() // DeviceUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.UpdateDevice(context.Background(), deviceId).DeviceUpdateRequest(deviceUpdateRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.UpdateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDevice`: SingleDevice
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.UpdateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | デバイスID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceUpdateRequest** | [**DeviceUpdateRequest**](DeviceUpdateRequest.md) |  | 

### Return type

[**SingleDevice**](SingleDevice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceClient

> SingleDeviceClientList UpdateDeviceClient(ctx, deviceId, clientId).RsaPublicKey(rsaPublicKey).Execute()

device.write



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
    deviceId := "deviceId_example" // string | 
    clientId := "clientId_example" // string | 
    rsaPublicKey := *openapiclient.NewRsaPublicKey("RsaPublickey_example") // RsaPublicKey | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.UpdateDeviceClient(context.Background(), deviceId, clientId).RsaPublicKey(rsaPublicKey).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.UpdateDeviceClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceClient`: SingleDeviceClientList
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.UpdateDeviceClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 
**clientId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rsaPublicKey** | [**RsaPublicKey**](RsaPublicKey.md) |  | 

### Return type

[**SingleDeviceClientList**](SingleDeviceClientList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

