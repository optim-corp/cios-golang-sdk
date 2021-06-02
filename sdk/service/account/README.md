# Account API

## Me API

### Interface

```
GetMe(ciosctx.RequestCtx) (cios.Me, *_nethttp.Response, error)
```
### Usage

#### Get Me

```go
me, httpResponse, err := client.Account.GetMe(ctx)
```


## Group API

### Interface

```
GetGroups(ciosctx.RequestCtx, cios.ApiGetGroupsRequest) (cios.MultipleGroup, *_nethttp.Response, error)
GetGroupsAll(ciosctx.RequestCtx, cios.ApiGetGroupsRequest) ([]cios.Group, *_nethttp.Response, error)
GetGroupsUnlimited(ciosctx.RequestCtx, cios.ApiGetGroupsRequest) ([]cios.Group, *_nethttp.Response, error)
GetGroup(ciosctx.RequestCtx, string, *string) (cios.Group, *_nethttp.Response, error)
GetGroupByResourceOwnerId(ciosctx.RequestCtx, string, *string) (cios.Group, *_nethttp.Response, error)
GetGroupMapByResourceOwner(ciosctx.RequestCtx, cios.ApiGetGroupsRequest, cios.ApiGetResourceOwnersRequest) (map[string]cios.Group, *_nethttp.Response, error)
DeleteGroup(ciosctx.RequestCtx, string) (*_nethttp.Response, error)
CreateGroup(ciosctx.RequestCtx, cios.GroupCreateRequest) (cios.Group, *_nethttp.Response, error)
UpdateGroup(ciosctx.RequestCtx, string, cios.GroupUpdateRequest) (cios.Group, *_nethttp.Response, error)
```

### Usage


#### Get a Group
```go
includes := "children"
group, httpResponse, err := client.Account.GetGroup(ctx, "groupID", &includes)
```

#### Get Groups limit 1000

```go
options := srvaccount.MakeGetGroupsOpts
groups, httpResponse, err := client.Account.GetGroups(ctx, options().Tags("sample_test").Limit(10))
```

#### Get Groups no limit

```go
options := srvaccount.MakeGetGroupsOpts
groups, httpResponse, err := client.Account.GetGroupsAll(ctx, options().Tags("sample_test").Limit(2000))
```

#### Get Groups unlimited

```go
options := srvaccount.MakeGetGroupsOpts
groups, httpResponse, err := client.Account.GetGroupsUnlimited(ctx, options().Tags("sample_test"))
```

#### Get a Group by Resource Owner ID

```go
includes := "children"
group, httpResponse, err := client.Account.GetGroupByResourceOwnerId(ctx, "resource_owner_id", &includes)
```

#### Crete a Group

```go
group, httpResponse, err := client.Account.CreateGroup(ctx, cios.GroupCreateRequest{
    Name: "",
    Type: "",
})
```

#### Update a Group

```go
parentGroupId := "parent_group_id"
groupName := "group_name"
group, httpResponse, err := client.Account.UpdateGroup(ctx, "group_id",
    cios.GroupUpdateRequest{
        ParentGroupId: &parentGroupId,
        Name:          &groupName,
    })
```

#### Delete a Group

```go
httpResponse, err := client.Account.DeleteGroup(ctx, "group_id")
```

## Resource Owner API


### interface

```
GetResourceOwners(ciosctx.RequestCtx, cios.ApiGetResourceOwnersRequest) (cios.MultipleResourceOwner, *_nethttp.Response, error)
GetResourceOwnersAll(ciosctx.RequestCtx, cios.ApiGetResourceOwnersRequest) ([]cios.ResourceOwner, *_nethttp.Response, error)
GetResourceOwnersUnlimited(ciosctx.RequestCtx, cios.ApiGetResourceOwnersRequest) ([]cios.ResourceOwner, *_nethttp.Response, error)
GetResourceOwner(ciosctx.RequestCtx, string) (cios.ResourceOwner, *_nethttp.Response, error)
GetResourceOwnerByGroupId(ciosctx.RequestCtx, string) (cios.ResourceOwner, *_nethttp.Response, error)
GetResourceOwnersMapByID(ciosctx.RequestCtx) (map[string]cios.ResourceOwner, *_nethttp.Response, error)
GetResourceOwnersMapByGroupID(ciosctx.RequestCtx) (map[string]cios.ResourceOwner, error)
```

### Usage

#### Get a ResourceOwner

```go
resourceOwner, httpResponse, err := client.Account.GetResourceOwner(ctx, "resource_owner_id")
```

#### Get a ResourceOwner by Group ID

```go
resourceOwner, httpResponse, err := client.Account.GetResourceOwnerByGroupId(ctx, "group_id")
```

#### Get ResourceOwner Map by ResourceOwner ID

```go
resourceOwner, httpResponse, err := client.Account.GetResourceOwnersMapByID(ctx)
```

#### Get ResourceOwner Map by Group ID

```go
resourceOwner, httpResponse, err := client.Account.GetResourceOwnersMapByGroupID(ctx)
```

#### Get ResourceOwners limit 1000

```go
options := srvaccount.MakeGetResourceOwnersOpts
resourceOwners, httpResponse, err := client.Account.GetResourceOwners(ctx, options().Limit(20).Offset(1000))
```

#### Get ResourceOwners no limit

```go
options := srvaccount.MakeGetResourceOwnersOpts
resourceOwners, httpResponse, err := client.Account.GetResourceOwnersAll(ctx, options().Limit(20000).Offset(1000))
```

#### Get ResourceOwners unlimited

```go
options := srvaccount.MakeGetResourceOwnersOpts
resourceOwners, httpResponse, err := client.Account.GetResourceOwnersUnlimited(ctx, options().Offset(1000))
```