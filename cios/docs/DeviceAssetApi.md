# \DeviceAssetApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceEntitiesComponent**](DeviceAssetApi.md#CreateDeviceEntitiesComponent) | **Post** /v2/device_entities/{key}/components | 
[**CreateDeviceEntitiesLifecycle**](DeviceAssetApi.md#CreateDeviceEntitiesLifecycle) | **Post** /v2/device_entities/{key}/lifecycles | 
[**CreateDeviceModel**](DeviceAssetApi.md#CreateDeviceModel) | **Post** /v2/device_models | 
[**CreateInventory**](DeviceAssetApi.md#CreateInventory) | **Post** /v2/device_models/{name}/entities | 
[**DeleteDeviceEntitiesComponent**](DeviceAssetApi.md#DeleteDeviceEntitiesComponent) | **Delete** /v2/device_entities/{key}/components/{component_id} | 
[**DeleteDeviceEntitiesLifecycle**](DeviceAssetApi.md#DeleteDeviceEntitiesLifecycle) | **Delete** /v2/device_entities/{key}/lifecycles/{lifecycle_id} | 
[**DeleteDeviceEntity**](DeviceAssetApi.md#DeleteDeviceEntity) | **Delete** /v2/device_entities/{key} | 
[**DeleteDeviceModel**](DeviceAssetApi.md#DeleteDeviceModel) | **Delete** /v2/device_models/{name} | 
[**GetDeviceEntities**](DeviceAssetApi.md#GetDeviceEntities) | **Get** /v2/device_entities | 
[**GetDeviceEntitiesComponent**](DeviceAssetApi.md#GetDeviceEntitiesComponent) | **Get** /v2/device_entities/{key}/components/{component_id} | 
[**GetDeviceEntitiesComponents**](DeviceAssetApi.md#GetDeviceEntitiesComponents) | **Get** /v2/device_entities/{key}/components | 
[**GetDeviceEntitiesLifecycle**](DeviceAssetApi.md#GetDeviceEntitiesLifecycle) | **Get** /v2/device_entities/{key}/lifecycles/{lifecycle_id} | 
[**GetDeviceEntitiesLifecycles**](DeviceAssetApi.md#GetDeviceEntitiesLifecycles) | **Get** /v2/device_entities/{key}/lifecycles | 
[**GetDeviceEntity**](DeviceAssetApi.md#GetDeviceEntity) | **Get** /v2/device_entities/{key} | 
[**GetDeviceModel**](DeviceAssetApi.md#GetDeviceModel) | **Get** /v2/device_models/{name} | 
[**GetDeviceModels**](DeviceAssetApi.md#GetDeviceModels) | **Get** /v2/device_models | 
[**GetDeviceModelsComponents**](DeviceAssetApi.md#GetDeviceModelsComponents) | **Get** /v2/device_models/{name}/components | 
[**UpdateDeviceEntitiesComponent**](DeviceAssetApi.md#UpdateDeviceEntitiesComponent) | **Patch** /v2/device_entities/{key}/components/{component_id} | 
[**UpdateDeviceEntitiesComponents**](DeviceAssetApi.md#UpdateDeviceEntitiesComponents) | **Patch** /v2/device_entities/{key}/components | 
[**UpdateDeviceEntity**](DeviceAssetApi.md#UpdateDeviceEntity) | **Patch** /v2/device_entities/{key} | 
[**UpdateDeviceModel**](DeviceAssetApi.md#UpdateDeviceModel) | **Patch** /v2/device_models/{name} | 



## CreateDeviceEntitiesComponent

> SingleDeviceEntitiesComponent CreateDeviceEntitiesComponent(ctx, key).DeviceEntitiesComponent(deviceEntitiesComponent).Execute()



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
    key := "key_example" // string | 
    deviceEntitiesComponent := *openapiclient.NewDeviceEntitiesComponent("Id_example", openapiclient.ComponentTypeEnum("product")) // DeviceEntitiesComponent | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.CreateDeviceEntitiesComponent(context.Background(), key).DeviceEntitiesComponent(deviceEntitiesComponent).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.CreateDeviceEntitiesComponent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceEntitiesComponent`: SingleDeviceEntitiesComponent
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssetApi.CreateDeviceEntitiesComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceEntitiesComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceEntitiesComponent** | [**DeviceEntitiesComponent**](DeviceEntitiesComponent.md) |  | 

### Return type

[**SingleDeviceEntitiesComponent**](SingleDeviceEntitiesComponent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceEntitiesLifecycle

> SingleLifeCycle CreateDeviceEntitiesLifecycle(ctx, key).LifeCycleRequest(lifeCycleRequest).Execute()



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
    key := "key_example" // string | 
    lifeCycleRequest := *openapiclient.NewLifeCycleRequest("EventKind_example", "EventMode_example", "EventType_example", "EventAt_example") // LifeCycleRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.CreateDeviceEntitiesLifecycle(context.Background(), key).LifeCycleRequest(lifeCycleRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.CreateDeviceEntitiesLifecycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceEntitiesLifecycle`: SingleLifeCycle
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssetApi.CreateDeviceEntitiesLifecycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceEntitiesLifecycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lifeCycleRequest** | [**LifeCycleRequest**](LifeCycleRequest.md) |  | 

### Return type

[**SingleLifeCycle**](SingleLifeCycle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceModel

> SingleDeviceModel CreateDeviceModel(ctx).DeviceModelRequest(deviceModelRequest).Execute()



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
    deviceModelRequest := *openapiclient.NewDeviceModelRequest("Name_example", "ResourceOwnerId_example") // DeviceModelRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.CreateDeviceModel(context.Background()).DeviceModelRequest(deviceModelRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.CreateDeviceModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceModel`: SingleDeviceModel
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssetApi.CreateDeviceModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceModelRequest** | [**DeviceModelRequest**](DeviceModelRequest.md) |  | 

### Return type

[**SingleDeviceModel**](SingleDeviceModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInventory

> SingleDeviceModelsEntity CreateInventory(ctx, name).Inventory(inventory).Execute()



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
    name := "name_example" // string | 
    inventory := *openapiclient.NewInventory() // Inventory | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.CreateInventory(context.Background(), name).Inventory(inventory).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.CreateInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInventory`: SingleDeviceModelsEntity
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssetApi.CreateInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inventory** | [**Inventory**](Inventory.md) |  | 

### Return type

[**SingleDeviceModelsEntity**](SingleDeviceModelsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceEntitiesComponent

> DeleteDeviceEntitiesComponent(ctx, key, component).IsRecursive(isRecursive).Execute()



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
    key := "key_example" // string | 
    component := "component_example" // string | 
    isRecursive := true // bool | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.DeleteDeviceEntitiesComponent(context.Background(), key, component).IsRecursive(isRecursive).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.DeleteDeviceEntitiesComponent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 
**component** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceEntitiesComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isRecursive** | **bool** |  | 

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


## DeleteDeviceEntitiesLifecycle

> DeleteDeviceEntitiesLifecycle(ctx, key, lifecycleId).Execute()



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
    key := "key_example" // string | 
    lifecycleId := "lifecycleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.DeleteDeviceEntitiesLifecycle(context.Background(), key, lifecycleId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.DeleteDeviceEntitiesLifecycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 
**lifecycleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceEntitiesLifecycleRequest struct via the builder pattern


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


## DeleteDeviceEntity

> DeleteDeviceEntity(ctx, key).Execute()



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
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.DeleteDeviceEntity(context.Background(), key).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.DeleteDeviceEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceEntityRequest struct via the builder pattern


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


## DeleteDeviceModel

> DeleteDeviceModel(ctx, name).Execute()



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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.DeleteDeviceModel(context.Background(), name).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.DeleteDeviceModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceModelRequest struct via the builder pattern


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


## GetDeviceEntities

> MultipleDeviceModelEntity GetDeviceEntities(ctx).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).Key(key).SerialNumber(serialNumber).DeviceId(deviceId).ResourceOwnerId(resourceOwnerId).ComponentKey(componentKey).ComponentValue(componentValue).IsFlat(isFlat).Execute()



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
    key := "key_example" // string |  (optional)
    serialNumber := "serialNumber_example" // string |  (optional)
    deviceId := "deviceId_example" // string |  (optional)
    resourceOwnerId := TODO // string |  (optional)
    componentKey := "componentKey_example" // string |  (optional)
    componentValue := "componentValue_example" // string |  (optional)
    isFlat := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.GetDeviceEntities(context.Background()).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).Key(key).SerialNumber(serialNumber).DeviceId(deviceId).ResourceOwnerId(resourceOwnerId).ComponentKey(componentKey).ComponentValue(componentValue).IsFlat(isFlat).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.GetDeviceEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceEntities`: MultipleDeviceModelEntity
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssetApi.GetDeviceEntities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 
 **order** | **string** | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 | 
 **orderBy** | **string** | ソート対象。指定順に高優先でのソートとなる | 
 **key** | **string** |  | 
 **serialNumber** | **string** |  | 
 **deviceId** | **string** |  | 
 **resourceOwnerId** | [**string**](string.md) |  | 
 **componentKey** | **string** |  | 
 **componentValue** | **string** |  | 
 **isFlat** | **bool** |  | [default to false]

### Return type

[**MultipleDeviceModelEntity**](MultipleDeviceModelEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceEntitiesComponent

> SingleDeviceEntitiesComponent GetDeviceEntitiesComponent(ctx, key, component).Execute()



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
    key := "key_example" // string | 
    component := "component_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.GetDeviceEntitiesComponent(context.Background(), key, component).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.GetDeviceEntitiesComponent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceEntitiesComponent`: SingleDeviceEntitiesComponent
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssetApi.GetDeviceEntitiesComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 
**component** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceEntitiesComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SingleDeviceEntitiesComponent**](SingleDeviceEntitiesComponent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceEntitiesComponents

> MultipleDeviceEntitiesComponent GetDeviceEntitiesComponents(ctx, key).ComponentIdQuery(componentIdQuery).Execute()



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
    key := "key_example" // string | 
    componentIdQuery := "componentIdQuery_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.GetDeviceEntitiesComponents(context.Background(), key).ComponentIdQuery(componentIdQuery).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.GetDeviceEntitiesComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceEntitiesComponents`: MultipleDeviceEntitiesComponent
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssetApi.GetDeviceEntitiesComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceEntitiesComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **componentIdQuery** | **string** |  | 

### Return type

[**MultipleDeviceEntitiesComponent**](MultipleDeviceEntitiesComponent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceEntitiesLifecycle

> SingleLifeCycle GetDeviceEntitiesLifecycle(ctx, key, lifecycleId).Execute()



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
    key := "key_example" // string | 
    lifecycleId := "lifecycleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.GetDeviceEntitiesLifecycle(context.Background(), key, lifecycleId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.GetDeviceEntitiesLifecycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceEntitiesLifecycle`: SingleLifeCycle
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssetApi.GetDeviceEntitiesLifecycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 
**lifecycleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceEntitiesLifecycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SingleLifeCycle**](SingleLifeCycle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceEntitiesLifecycles

> MultipleLifeCycle GetDeviceEntitiesLifecycles(ctx, key).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).EventKind(eventKind).EventMode(eventMode).EventType(eventType).ComponentId(componentId).BeforeId(beforeId).AfterId(afterId).StartEventAt(startEventAt).EndEventAt(endEventAt).Execute()



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
    key := "key_example" // string | 
    offset := int64(789) // int64 | ページネーションのうち、何件読み飛ばすか (optional)
    limit := int64(789) // int64 | ページネーションのうち、上位何件取得するか (optional)
    order := "order_example" // string | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 (optional)
    orderBy := "orderBy_example" // string | ソート対象。指定順に高優先でのソートとなる (optional)
    eventKind := "eventKind_example" // string |  (optional)
    eventMode := "eventMode_example" // string |  (optional)
    eventType := "eventType_example" // string |  (optional)
    componentId := "componentId_example" // string |  (optional)
    beforeId := "beforeId_example" // string |  (optional)
    afterId := "afterId_example" // string |  (optional)
    startEventAt := "startEventAt_example" // string |  (optional)
    endEventAt := "endEventAt_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.GetDeviceEntitiesLifecycles(context.Background(), key).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).EventKind(eventKind).EventMode(eventMode).EventType(eventType).ComponentId(componentId).BeforeId(beforeId).AfterId(afterId).StartEventAt(startEventAt).EndEventAt(endEventAt).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.GetDeviceEntitiesLifecycles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceEntitiesLifecycles`: MultipleLifeCycle
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssetApi.GetDeviceEntitiesLifecycles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceEntitiesLifecyclesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 
 **order** | **string** | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 | 
 **orderBy** | **string** | ソート対象。指定順に高優先でのソートとなる | 
 **eventKind** | **string** |  | 
 **eventMode** | **string** |  | 
 **eventType** | **string** |  | 
 **componentId** | **string** |  | 
 **beforeId** | **string** |  | 
 **afterId** | **string** |  | 
 **startEventAt** | **string** |  | 
 **endEventAt** | **string** |  | 

### Return type

[**MultipleLifeCycle**](MultipleLifeCycle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceEntity

> SingleDeviceModelsEntity GetDeviceEntity(ctx, key).Execute()



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
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.GetDeviceEntity(context.Background(), key).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.GetDeviceEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceEntity`: SingleDeviceModelsEntity
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssetApi.GetDeviceEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SingleDeviceModelsEntity**](SingleDeviceModelsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceModel

> SingleDeviceModel GetDeviceModel(ctx, name).Execute()



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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.GetDeviceModel(context.Background(), name).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.GetDeviceModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceModel`: SingleDeviceModel
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssetApi.GetDeviceModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SingleDeviceModel**](SingleDeviceModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceModels

> MultipleDeviceModel GetDeviceModels(ctx).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).Name(name).EndEventAt(endEventAt).Version(version).ResourceOwnerId(resourceOwnerId).ComponentKey(componentKey).ComponentValue(componentValue).Execute()



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
    name := "name_example" // string | 名前での部分一致検索 (optional)
    endEventAt := "endEventAt_example" // string |  (optional)
    version := "version_example" // string |  (optional)
    resourceOwnerId := TODO // string |  (optional)
    componentKey := "componentKey_example" // string |  (optional)
    componentValue := "componentValue_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.GetDeviceModels(context.Background()).Offset(offset).Limit(limit).Order(order).OrderBy(orderBy).Name(name).EndEventAt(endEventAt).Version(version).ResourceOwnerId(resourceOwnerId).ComponentKey(componentKey).ComponentValue(componentValue).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.GetDeviceModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceModels`: MultipleDeviceModel
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssetApi.GetDeviceModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 
 **order** | **string** | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 | 
 **orderBy** | **string** | ソート対象。指定順に高優先でのソートとなる | 
 **name** | **string** | 名前での部分一致検索 | 
 **endEventAt** | **string** |  | 
 **version** | **string** |  | 
 **resourceOwnerId** | [**string**](string.md) |  | 
 **componentKey** | **string** |  | 
 **componentValue** | **string** |  | 

### Return type

[**MultipleDeviceModel**](MultipleDeviceModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceModelsComponents

> MultipleConstitutionComponent GetDeviceModelsComponents(ctx, name).Execute()



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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.GetDeviceModelsComponents(context.Background(), name).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.GetDeviceModelsComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceModelsComponents`: MultipleConstitutionComponent
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssetApi.GetDeviceModelsComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceModelsComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MultipleConstitutionComponent**](MultipleConstitutionComponent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceEntitiesComponent

> SingleDeviceEntitiesComponent UpdateDeviceEntitiesComponent(ctx, key, component).DeviceEntitiesComponentUpdateRequest(deviceEntitiesComponentUpdateRequest).Execute()



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
    key := "key_example" // string | 
    component := "component_example" // string | 
    deviceEntitiesComponentUpdateRequest := *openapiclient.NewDeviceEntitiesComponentUpdateRequest() // DeviceEntitiesComponentUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.UpdateDeviceEntitiesComponent(context.Background(), key, component).DeviceEntitiesComponentUpdateRequest(deviceEntitiesComponentUpdateRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.UpdateDeviceEntitiesComponent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceEntitiesComponent`: SingleDeviceEntitiesComponent
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssetApi.UpdateDeviceEntitiesComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 
**component** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceEntitiesComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceEntitiesComponentUpdateRequest** | [**DeviceEntitiesComponentUpdateRequest**](DeviceEntitiesComponentUpdateRequest.md) |  | 

### Return type

[**SingleDeviceEntitiesComponent**](SingleDeviceEntitiesComponent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceEntitiesComponents

> UpdateDeviceEntitiesComponents(ctx, key).DeviceEntitiesComponent(deviceEntitiesComponent).Execute()





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
    key := "key_example" // string | 
    deviceEntitiesComponent := []openapiclient.DeviceEntitiesComponent{*openapiclient.NewDeviceEntitiesComponent("Id_example", openapiclient.ComponentTypeEnum("product"))} // []DeviceEntitiesComponent | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.UpdateDeviceEntitiesComponents(context.Background(), key).DeviceEntitiesComponent(deviceEntitiesComponent).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.UpdateDeviceEntitiesComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceEntitiesComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceEntitiesComponent** | [**[]DeviceEntitiesComponent**](DeviceEntitiesComponent.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceEntity

> SingleDeviceModelsEntity UpdateDeviceEntity(ctx, key).DeviceModelsEntity(deviceModelsEntity).Execute()



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
    key := "key_example" // string | 
    deviceModelsEntity := *openapiclient.NewDeviceModelsEntity("Id_example", "Key_example", "DeviceId_example", *openapiclient.NewDeviceModelsEntityModel("Id_example", "Name_example", "CreatedAt_example", "UpdatedAt_example"), "ResourceOwnerId_example", "CreatedAt_example", "UpdatedAt_example") // DeviceModelsEntity | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.UpdateDeviceEntity(context.Background(), key).DeviceModelsEntity(deviceModelsEntity).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.UpdateDeviceEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceEntity`: SingleDeviceModelsEntity
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssetApi.UpdateDeviceEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceModelsEntity** | [**DeviceModelsEntity**](DeviceModelsEntity.md) |  | 

### Return type

[**SingleDeviceModelsEntity**](SingleDeviceModelsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceModel

> SingleDeviceModel UpdateDeviceModel(ctx, name).DeviceModelUpdateRequest(deviceModelUpdateRequest).Execute()





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
    name := "name_example" // string | 
    deviceModelUpdateRequest := *openapiclient.NewDeviceModelUpdateRequest() // DeviceModelUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAssetApi.UpdateDeviceModel(context.Background(), name).DeviceModelUpdateRequest(deviceModelUpdateRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAssetApi.UpdateDeviceModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceModel`: SingleDeviceModel
    fmt.Fprintf(os.Stdout, "Response from `DeviceAssetApi.UpdateDeviceModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceModelUpdateRequest** | [**DeviceModelUpdateRequest**](DeviceModelUpdateRequest.md) |  | 

### Return type

[**SingleDeviceModel**](SingleDeviceModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

