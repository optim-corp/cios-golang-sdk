# \GeographyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCircle**](GeographyApi.md#CreateCircle) | **Post** /v2/geo/circles | 
[**CreatePoint**](GeographyApi.md#CreatePoint) | **Post** /v2/geo/points | 
[**CreateRoute**](GeographyApi.md#CreateRoute) | **Post** /v2/geo/routes | 
[**DeleteCircle**](GeographyApi.md#DeleteCircle) | **Delete** /v2/geo/circles/{circles_id} | 
[**DeletePoint**](GeographyApi.md#DeletePoint) | **Delete** /v2/geo/points/{point_id} | 
[**DeleteRoute**](GeographyApi.md#DeleteRoute) | **Delete** /v2/geo/routes/{routes_id} | 
[**GetCircle**](GeographyApi.md#GetCircle) | **Get** /v2/geo/circles/{circles_id} | 
[**GetCircles**](GeographyApi.md#GetCircles) | **Get** /v2/geo/circles | 
[**GetPoint**](GeographyApi.md#GetPoint) | **Get** /v2/geo/points/{point_id} | 
[**GetPoints**](GeographyApi.md#GetPoints) | **Get** /v2/geo/points | 
[**GetRoute**](GeographyApi.md#GetRoute) | **Get** /v2/geo/routes/{routes_id} | 
[**GetRoutes**](GeographyApi.md#GetRoutes) | **Get** /v2/geo/routes | 
[**UpdateCircle**](GeographyApi.md#UpdateCircle) | **Patch** /v2/geo/circles/{circles_id} | 
[**UpdatePoint**](GeographyApi.md#UpdatePoint) | **Patch** /v2/geo/points/{point_id} | 
[**UpdateRoute**](GeographyApi.md#UpdateRoute) | **Patch** /v2/geo/routes/{routes_id} | 



## CreateCircle

> SingleCircle CreateCircle(ctx).Circle(circle).Execute()



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
    circle := *openapiclient.NewCircle() // Circle | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeographyApi.CreateCircle(context.Background()).Circle(circle).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographyApi.CreateCircle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCircle`: SingleCircle
    fmt.Fprintf(os.Stdout, "Response from `GeographyApi.CreateCircle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCircleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **circle** | [**Circle**](Circle.md) |  | 

### Return type

[**SingleCircle**](SingleCircle.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePoint

> SinglePoint CreatePoint(ctx).PointRequest(pointRequest).Execute()



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
    pointRequest := *openapiclient.NewPointRequest("ResourceOwnerId_example") // PointRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeographyApi.CreatePoint(context.Background()).PointRequest(pointRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographyApi.CreatePoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePoint`: SinglePoint
    fmt.Fprintf(os.Stdout, "Response from `GeographyApi.CreatePoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pointRequest** | [**PointRequest**](PointRequest.md) |  | 

### Return type

[**SinglePoint**](SinglePoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoute

> SingleRoute CreateRoute(ctx).Route(route).Execute()



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
    route := *openapiclient.NewRoute() // Route | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeographyApi.CreateRoute(context.Background()).Route(route).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographyApi.CreateRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRoute`: SingleRoute
    fmt.Fprintf(os.Stdout, "Response from `GeographyApi.CreateRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **route** | [**Route**](Route.md) |  | 

### Return type

[**SingleRoute**](SingleRoute.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCircle

> DeleteCircle(ctx, circleId).Execute()



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
    circleId := "circleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeographyApi.DeleteCircle(context.Background(), circleId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographyApi.DeleteCircle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**circleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCircleRequest struct via the builder pattern


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


## DeletePoint

> SinglePoint DeletePoint(ctx, pointId).Execute()



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
    pointId := "pointId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeographyApi.DeletePoint(context.Background(), pointId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographyApi.DeletePoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePoint`: SinglePoint
    fmt.Fprintf(os.Stdout, "Response from `GeographyApi.DeletePoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pointId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SinglePoint**](SinglePoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoute

> DeleteRoute(ctx, routeId).Execute()



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
    routeId := "routeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeographyApi.DeleteRoute(context.Background(), routeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographyApi.DeleteRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteRequest struct via the builder pattern


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


## GetCircle

> SingleCircle GetCircle(ctx, circleId).Lang(lang).IsDev(isDev).Execute()



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
    circleId := "circleId_example" // string | 
    lang := "lang_example" // string | 言語指定 (optional)
    isDev := true // bool | 開発者モードの有効・無効 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeographyApi.GetCircle(context.Background(), circleId).Lang(lang).IsDev(isDev).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographyApi.GetCircle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCircle`: SingleCircle
    fmt.Fprintf(os.Stdout, "Response from `GeographyApi.GetCircle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**circleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCircleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lang** | **string** | 言語指定 | 
 **isDev** | **bool** | 開発者モードの有効・無効 | 

### Return type

[**SingleCircle**](SingleCircle.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCircles

> MultipleCircle GetCircles(ctx).Name(name).Label(label).IsPublic(isPublic).ResourceOwnerId(resourceOwnerId).Sort(sort).Top(top).Skip(skip).Lang(lang).IsDev(isDev).Execute()



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
    label := "label_example" // string |  (optional)
    isPublic := "isPublic_example" // string | デバイスの公開・非公開 (optional)
    resourceOwnerId := TODO // string |  (optional)
    sort := "sort_example" // string |  (optional)
    top := int32(56) // int32 |  (optional)
    skip := int32(56) // int32 |  (optional)
    lang := "lang_example" // string | 言語指定 (optional)
    isDev := true // bool | 開発者モードの有効・無効 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeographyApi.GetCircles(context.Background()).Name(name).Label(label).IsPublic(isPublic).ResourceOwnerId(resourceOwnerId).Sort(sort).Top(top).Skip(skip).Lang(lang).IsDev(isDev).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographyApi.GetCircles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCircles`: MultipleCircle
    fmt.Fprintf(os.Stdout, "Response from `GeographyApi.GetCircles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCirclesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | 名前での部分一致検索 | 
 **label** | **string** |  | 
 **isPublic** | **string** | デバイスの公開・非公開 | 
 **resourceOwnerId** | [**string**](string.md) |  | 
 **sort** | **string** |  | 
 **top** | **int32** |  | 
 **skip** | **int32** |  | 
 **lang** | **string** | 言語指定 | 
 **isDev** | **bool** | 開発者モードの有効・無効 | 

### Return type

[**MultipleCircle**](MultipleCircle.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoint

> SinglePoint GetPoint(ctx, pointId).Lang(lang).IsDev(isDev).Execute()



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
    pointId := "pointId_example" // string | 
    lang := "lang_example" // string | 言語指定 (optional)
    isDev := true // bool | 開発者モードの有効・無効 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeographyApi.GetPoint(context.Background(), pointId).Lang(lang).IsDev(isDev).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographyApi.GetPoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPoint`: SinglePoint
    fmt.Fprintf(os.Stdout, "Response from `GeographyApi.GetPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pointId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lang** | **string** | 言語指定 | 
 **isDev** | **bool** | 開発者モードの有効・無効 | 

### Return type

[**SinglePoint**](SinglePoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoints

> MultiplePoint GetPoints(ctx).Name(name).Label(label).IsPublic(isPublic).ResourceOwnerId(resourceOwnerId).Sort(sort).Top(top).Skip(skip).Lang(lang).IsDev(isDev).Execute()



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
    label := "label_example" // string |  (optional)
    isPublic := "isPublic_example" // string | デバイスの公開・非公開 (optional)
    resourceOwnerId := TODO // string |  (optional)
    sort := "sort_example" // string |  (optional)
    top := int32(56) // int32 |  (optional)
    skip := int32(56) // int32 |  (optional)
    lang := "lang_example" // string | 言語指定 (optional)
    isDev := true // bool | 開発者モードの有効・無効 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeographyApi.GetPoints(context.Background()).Name(name).Label(label).IsPublic(isPublic).ResourceOwnerId(resourceOwnerId).Sort(sort).Top(top).Skip(skip).Lang(lang).IsDev(isDev).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographyApi.GetPoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPoints`: MultiplePoint
    fmt.Fprintf(os.Stdout, "Response from `GeographyApi.GetPoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | 名前での部分一致検索 | 
 **label** | **string** |  | 
 **isPublic** | **string** | デバイスの公開・非公開 | 
 **resourceOwnerId** | [**string**](string.md) |  | 
 **sort** | **string** |  | 
 **top** | **int32** |  | 
 **skip** | **int32** |  | 
 **lang** | **string** | 言語指定 | 
 **isDev** | **bool** | 開発者モードの有効・無効 | 

### Return type

[**MultiplePoint**](MultiplePoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoute

> SingleRoute GetRoute(ctx, routeId).Lang(lang).IsDev(isDev).Execute()



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
    routeId := "routeId_example" // string | 
    lang := "lang_example" // string | 言語指定 (optional)
    isDev := true // bool | 開発者モードの有効・無効 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeographyApi.GetRoute(context.Background(), routeId).Lang(lang).IsDev(isDev).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographyApi.GetRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoute`: SingleRoute
    fmt.Fprintf(os.Stdout, "Response from `GeographyApi.GetRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lang** | **string** | 言語指定 | 
 **isDev** | **bool** | 開発者モードの有効・無効 | 

### Return type

[**SingleRoute**](SingleRoute.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoutes

> MultipleRoute GetRoutes(ctx).Name(name).Label(label).IsPublic(isPublic).ResourceOwnerId(resourceOwnerId).Sort(sort).Top(top).Skip(skip).Lang(lang).IsDev(isDev).Execute()



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
    label := "label_example" // string |  (optional)
    isPublic := "isPublic_example" // string | デバイスの公開・非公開 (optional)
    resourceOwnerId := TODO // string |  (optional)
    sort := "sort_example" // string |  (optional)
    top := int32(56) // int32 |  (optional)
    skip := int32(56) // int32 |  (optional)
    lang := "lang_example" // string | 言語指定 (optional)
    isDev := true // bool | 開発者モードの有効・無効 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeographyApi.GetRoutes(context.Background()).Name(name).Label(label).IsPublic(isPublic).ResourceOwnerId(resourceOwnerId).Sort(sort).Top(top).Skip(skip).Lang(lang).IsDev(isDev).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographyApi.GetRoutes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoutes`: MultipleRoute
    fmt.Fprintf(os.Stdout, "Response from `GeographyApi.GetRoutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | 名前での部分一致検索 | 
 **label** | **string** |  | 
 **isPublic** | **string** | デバイスの公開・非公開 | 
 **resourceOwnerId** | [**string**](string.md) |  | 
 **sort** | **string** |  | 
 **top** | **int32** |  | 
 **skip** | **int32** |  | 
 **lang** | **string** | 言語指定 | 
 **isDev** | **bool** | 開発者モードの有効・無効 | 

### Return type

[**MultipleRoute**](MultipleRoute.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCircle

> SingleCircle UpdateCircle(ctx, circleId).DisplayInfo(displayInfo).Execute()



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
    circleId := "circleId_example" // string | 
    displayInfo := []openapiclient.DisplayInfo{*openapiclient.NewDisplayInfo("Name_example", "Language_example", false)} // []DisplayInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeographyApi.UpdateCircle(context.Background(), circleId).DisplayInfo(displayInfo).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographyApi.UpdateCircle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCircle`: SingleCircle
    fmt.Fprintf(os.Stdout, "Response from `GeographyApi.UpdateCircle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**circleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCircleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **displayInfo** | [**[]DisplayInfo**](DisplayInfo.md) |  | 

### Return type

[**SingleCircle**](SingleCircle.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePoint

> SinglePoint UpdatePoint(ctx, pointId).DisplayInfo(displayInfo).Execute()



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
    pointId := "pointId_example" // string | 
    displayInfo := []openapiclient.DisplayInfo{*openapiclient.NewDisplayInfo("Name_example", "Language_example", false)} // []DisplayInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeographyApi.UpdatePoint(context.Background(), pointId).DisplayInfo(displayInfo).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographyApi.UpdatePoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePoint`: SinglePoint
    fmt.Fprintf(os.Stdout, "Response from `GeographyApi.UpdatePoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pointId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **displayInfo** | [**[]DisplayInfo**](DisplayInfo.md) |  | 

### Return type

[**SinglePoint**](SinglePoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoute

> SingleRoute UpdateRoute(ctx, routeId).DisplayInfo(displayInfo).Execute()



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
    routeId := "routeId_example" // string | 
    displayInfo := []openapiclient.DisplayInfo{*openapiclient.NewDisplayInfo("Name_example", "Language_example", false)} // []DisplayInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeographyApi.UpdateRoute(context.Background(), routeId).DisplayInfo(displayInfo).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographyApi.UpdateRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoute`: SingleRoute
    fmt.Fprintf(os.Stdout, "Response from `GeographyApi.UpdateRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **displayInfo** | [**[]DisplayInfo**](DisplayInfo.md) |  | 

### Return type

[**SingleRoute**](SingleRoute.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

