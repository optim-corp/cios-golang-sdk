# \GroupApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroup**](GroupApi.md#CreateGroup) | **Post** /v2/groups | 
[**DeleteGroup**](GroupApi.md#DeleteGroup) | **Delete** /v2/groups/{group_id} | 
[**DeleteMember**](GroupApi.md#DeleteMember) | **Delete** /v2/groups/{group_id}/members | 
[**GetGroup**](GroupApi.md#GetGroup) | **Get** /v2/groups/{group_id} | 
[**GetGroups**](GroupApi.md#GetGroups) | **Get** /v2/groups | 
[**GetMembers**](GroupApi.md#GetMembers) | **Get** /v2/groups/{group_id}/members | 
[**InviteGroup**](GroupApi.md#InviteGroup) | **Post** /v2/groups/{group_id}/invites | 
[**ReplaceGroup**](GroupApi.md#ReplaceGroup) | **Put** /v2/groups/{group_id} | 
[**SetMember**](GroupApi.md#SetMember) | **Patch** /v2/groups/{group_id}/members | 
[**UpdateGroup**](GroupApi.md#UpdateGroup) | **Patch** /v2/groups/{group_id} | 
[**UpdateMember**](GroupApi.md#UpdateMember) | **Put** /v2/groups/{group_id}/members | 



## CreateGroup

> Group CreateGroup(ctx).GroupCreateRequest(groupCreateRequest).Execute()





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
    groupCreateRequest := *openapiclient.NewGroupCreateRequest("Name_example", "Type_example") // GroupCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.CreateGroup(context.Background()).GroupCreateRequest(groupCreateRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.CreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupCreateRequest** | [**GroupCreateRequest**](GroupCreateRequest.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> DeleteGroup(ctx, groupId).Execute()





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
    groupId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.DeleteGroup(context.Background(), groupId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.DeleteGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


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


## DeleteMember

> DeleteMember(ctx, groupId).Member(member).Execute()



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
    groupId := TODO // string | 
    member := []openapiclient.Member{*openapiclient.NewMember()} // []Member | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.DeleteMember(context.Background(), groupId).Member(member).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.DeleteMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **member** | [**[]Member**](Member.md) |  | 

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


## GetGroup

> Group GetGroup(ctx, groupId).Includes(includes).Execute()





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
    groupId := TODO // string | 
    includes := "includes_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.GetGroup(context.Background(), groupId).Includes(includes).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.GetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includes** | **string** |  | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroups

> MultipleGroup GetGroups(ctx).Name(name).ParentGroupId(parentGroupId).State(state).City(city).Address1(address1).Address2(address2).Type_(type_).Tags(tags).Label(label).Domain(domain).Offset(offset).Limit(limit).OrderBy(orderBy).Order(order).Includes(includes).Page(page).Execute()





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
    parentGroupId := TODO // string |  (optional)
    state := "state_example" // string |  (optional)
    city := "city_example" // string |  (optional)
    address1 := "address1_example" // string |  (optional)
    address2 := "address2_example" // string |  (optional)
    type_ := "type__example" // string |  (optional)
    tags := "tags_example" // string |  (optional)
    label := "label_example" // string |  (optional)
    domain := "domain_example" // string |  (optional)
    offset := int64(789) // int64 | ページネーションのうち、何件読み飛ばすか (optional)
    limit := int64(789) // int64 | ページネーションのうち、上位何件取得するか (optional)
    orderBy := "orderBy_example" // string |  (optional)
    order := "order_example" // string | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 (optional)
    includes := "includes_example" // string |  (optional)
    page := "page_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.GetGroups(context.Background()).Name(name).ParentGroupId(parentGroupId).State(state).City(city).Address1(address1).Address2(address2).Type_(type_).Tags(tags).Label(label).Domain(domain).Offset(offset).Limit(limit).OrderBy(orderBy).Order(order).Includes(includes).Page(page).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.GetGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroups`: MultipleGroup
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.GetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | 名前での部分一致検索 | 
 **parentGroupId** | [**string**](string.md) |  | 
 **state** | **string** |  | 
 **city** | **string** |  | 
 **address1** | **string** |  | 
 **address2** | **string** |  | 
 **type_** | **string** |  | 
 **tags** | **string** |  | 
 **label** | **string** |  | 
 **domain** | **string** |  | 
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 
 **orderBy** | **string** |  | 
 **order** | **string** | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 | 
 **includes** | **string** |  | 
 **page** | **string** |  | 

### Return type

[**MultipleGroup**](MultipleGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMembers

> MultipleMemberInfo GetMembers(ctx, groupId).UserId(userId).Name(name).Email(email).Primary(primary).PhoneticFamilyName(phoneticFamilyName).PhoneticGivenName(phoneticGivenName).Role(role).Category(category).Tag(tag).Includes(includes).Page(page).Offset(offset).Limit(limit).OrderBy(orderBy).Order(order).Execute()



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
    groupId := TODO // string | 
    userId := TODO // string |  (optional)
    name := "name_example" // string | 名前での部分一致検索 (optional)
    email := "email_example" // string |  (optional)
    primary := true // bool |  (optional)
    phoneticFamilyName := "phoneticFamilyName_example" // string |  (optional)
    phoneticGivenName := "phoneticGivenName_example" // string |  (optional)
    role := "role_example" // string |  (optional)
    category := "category_example" // string |  (optional)
    tag := "tag_example" // string |  (optional)
    includes := "includes_example" // string |  (optional)
    page := "page_example" // string |  (optional)
    offset := int64(789) // int64 | ページネーションのうち、何件読み飛ばすか (optional)
    limit := int64(789) // int64 | ページネーションのうち、上位何件取得するか (optional)
    orderBy := "orderBy_example" // string |  (optional)
    order := "order_example" // string | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.GetMembers(context.Background(), groupId).UserId(userId).Name(name).Email(email).Primary(primary).PhoneticFamilyName(phoneticFamilyName).PhoneticGivenName(phoneticGivenName).Role(role).Category(category).Tag(tag).Includes(includes).Page(page).Offset(offset).Limit(limit).OrderBy(orderBy).Order(order).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.GetMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMembers`: MultipleMemberInfo
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.GetMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**string**](string.md) |  | 
 **name** | **string** | 名前での部分一致検索 | 
 **email** | **string** |  | 
 **primary** | **bool** |  | 
 **phoneticFamilyName** | **string** |  | 
 **phoneticGivenName** | **string** |  | 
 **role** | **string** |  | 
 **category** | **string** |  | 
 **tag** | **string** |  | 
 **includes** | **string** |  | 
 **page** | **string** |  | 
 **offset** | **int64** | ページネーションのうち、何件読み飛ばすか | 
 **limit** | **int64** | ページネーションのうち、上位何件取得するか | 
 **orderBy** | **string** |  | 
 **order** | **string** | ソート順序。 order_byで指定したkeyそれぞれに対して昇順、降順を配列で指定 | 

### Return type

[**MultipleMemberInfo**](MultipleMemberInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteGroup

> Member InviteGroup(ctx, groupId).GroupInviteRequest(groupInviteRequest).Execute()





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
    groupId := TODO // string | 
    groupInviteRequest := *openapiclient.NewGroupInviteRequest() // GroupInviteRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.InviteGroup(context.Background(), groupId).GroupInviteRequest(groupInviteRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.InviteGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InviteGroup`: Member
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.InviteGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupInviteRequest** | [**GroupInviteRequest**](GroupInviteRequest.md) |  | 

### Return type

[**Member**](Member.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceGroup

> Group ReplaceGroup(ctx, groupId).GroupReplaceRequest(groupReplaceRequest).Execute()





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
    groupId := TODO // string | 
    groupReplaceRequest := *openapiclient.NewGroupReplaceRequest() // GroupReplaceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.ReplaceGroup(context.Background(), groupId).GroupReplaceRequest(groupReplaceRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.ReplaceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.ReplaceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupReplaceRequest** | [**GroupReplaceRequest**](GroupReplaceRequest.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMember

> Member SetMember(ctx, groupId).Member(member).Execute()





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
    groupId := TODO // string | 
    member := *openapiclient.NewMember() // Member | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.SetMember(context.Background(), groupId).Member(member).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.SetMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetMember`: Member
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.SetMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **member** | [**Member**](Member.md) |  | 

### Return type

[**Member**](Member.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> Group UpdateGroup(ctx, groupId).GroupUpdateRequest(groupUpdateRequest).Execute()





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
    groupId := TODO // string | 
    groupUpdateRequest := *openapiclient.NewGroupUpdateRequest() // GroupUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.UpdateGroup(context.Background(), groupId).GroupUpdateRequest(groupUpdateRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.UpdateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.UpdateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupUpdateRequest** | [**GroupUpdateRequest**](GroupUpdateRequest.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMember

> MultipleMember UpdateMember(ctx, groupId).MultipleMember(multipleMember).Execute()





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
    groupId := TODO // string | 
    multipleMember := *openapiclient.NewMultipleMember() // MultipleMember | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.UpdateMember(context.Background(), groupId).MultipleMember(multipleMember).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.UpdateMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMember`: MultipleMember
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.UpdateMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **multipleMember** | [**MultipleMember**](MultipleMember.md) |  | 

### Return type

[**MultipleMember**](MultipleMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

