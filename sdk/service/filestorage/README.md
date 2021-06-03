# File Storage

## Bucket API

### interface

```
GetBuckets(ciosctx.RequestCtx, cios.ApiGetBucketsRequest) (cios.MultipleBucket, *_nethttp.Response, error)
GetBucketsAll(ciosctx.RequestCtx, cios.ApiGetBucketsRequest) ([]cios.Bucket, *_nethttp.Response, error)
GetBucketsUnlimited(ciosctx.RequestCtx, cios.ApiGetBucketsRequest) ([]cios.Bucket, *_nethttp.Response, error)
GetBucket(ciosctx.RequestCtx, string) (cios.Bucket, *_nethttp.Response, error)
GetBucketByResourceOwnerIDAndName(ciosctx.RequestCtx, string, string) (cios.Bucket, *_nethttp.Response, error)
GetOrCreateBucket(ciosctx.RequestCtx, string, string) (cios.Bucket, *_nethttp.Response, error)
CreateBucket(ciosctx.RequestCtx, string, string) (cios.Bucket, *_nethttp.Response, error)
DeleteBucket(ciosctx.RequestCtx, string) (*_nethttp.Response, error)
UpdateBucket(ciosctx.RequestCtx, string, string) (*_nethttp.Response, error)
```

### Usage

#### Get a Bucket

```go
bucket, httpResponse, err := client.FileStorage.GetBucketsUnlimited(ctx, "bucket_id")
```

#### Get a Bucket by ResourceOwnerID + Bucket Name

```go
bucket, httpResponse, err := client.FileStorage.GetBucketByResourceOwnerIDAndName(ctx, "resource_owner_id", "name")
```

#### Get Buckets max limit 1000

```go
options := srvfilestorage.MakeGetBucketsOpts
buckets, httpResponse, err := client.FileStorage.GetBuckets(ctx, options().Name("test"))
```

#### Get Buckets no limit 

```go
options := srvfilestorage.MakeGetBucketsOpts
buckets, httpResponse, err := client.FileStorage.GetBucketsAll(ctx, options().Name("test").Limit(2000))
```

#### Get Buckets unlimited

```go
options := srvfilestorage.MakeGetBucketsOpts
buckets, httpResponse, err := client.FileStorage.GetBucketsUnlimited(ctx, options().Name("test"))
```


#### Get or create a Bucket

```go
bucket, httpResponse, err := client.FileStorage.GetOrCreateBucket(ctx, "resource_owner_id", "name")
```

#### Create a Bucket

```go
bucket, httpResponse, err := client.FileStorage.CreateBucket(ctx, "resource_owner_id", "name")
```

#### Update a Bucket

```go
httpResponse, err := client.FileStorage.UpdateBucket(ctx, "bucket_id", "name")
```

#### Delete a Bucket

```go
httpResponse, err := client.FileStorage.DeleteBucket(ctx, "bucket_id")
```


## Node API

### interface

```
GetNodes(ciosctx.RequestCtx, string, cios.ApiGetNodesRequest) (cios.MultipleNode, *_nethttp.Response, error)
GetNodesAll(ciosctx.RequestCtx, string, cios.ApiGetNodesRequest) ([]cios.Node, *_nethttp.Response, error)
GetNodesUnlimited(ciosctx.RequestCtx, string, cios.ApiGetNodesRequest) ([]cios.Node, *_nethttp.Response, error)
GetNode(ciosctx.RequestCtx, string, string) (cios.Node, *_nethttp.Response, error)
CreateNode(ciosctx.RequestCtx, string, string, *string) (cios.Node, *_nethttp.Response, error)
CreateNodeOnNodeID(ciosctx.RequestCtx, string, cios.NodeRequest) (cios.Node, *_nethttp.Response, error)
DeleteNode(ciosctx.RequestCtx, string, string) (*_nethttp.Response, error)
CopyNode(ciosctx.RequestCtx, string, string, *string, *string) (cios.Node, *_nethttp.Response, error)
MoveNode(ciosctx.RequestCtx, string, string, *string, *string) (cios.Node, *_nethttp.Response, error)
RenameNode(ciosctx.RequestCtx, string, string, string) (cios.Node, *_nethttp.Response, error)
```

### Usage

#### Get a Node

```go
node, httpResponse, err := client.FileStorage.GetNode(ctx, "bucket_id", "node_id")
```


#### Get Nodes max limit 1000

```go
options := srvfilestorage.MakeGetNodesOpts
nodes, httpResponse, err := client.FileStorage.GetNodes(ctx, "bucket_id", options().Limit(200).Name("sample"))
```

#### Get Nodes no limit 

```go
options := srvfilestorage.MakeGetNodesOpts
nodes, httpResponse, err := client.FileStorage.GetNodesAll(ctx, "bucket_id", options().Limit(20000).Name("sample"))
```

#### Get Nodes unlimited

```go
options := srvfilestorage.MakeGetNodesOpts
nodes, httpResponse, err := client.FileStorage.GetNodesUnlimited(ctx, "bucket_id", options().Name("sample"))
```


#### Create a Directory

```go
parentNodeId := "parent_node_id"
node, httpResponse, err := CreateNode(ciosctx.RequestCtx, "bucket_id", "node_name", &parentNodeId)
```

#### Delete a Node

```go
httpResponse, err := client.FileStorage.DeleteNode(ctx, "bucket_id", "node_id")
```

#### Copy a Node

```go
destBucketId := "dest_bucket_id"
destNodeId := "dest_node_id"
node, httpResponse, err := client.FileStorage.CopyNode(ctx, "bucket_id", "node_id", &destBucketId, &destNodeId)
```

#### Move a Node

```go
destBucketId := "dest_bucket_id"
destNodeId := "dest_node_id"
node, httpResponse, err := client.FileStorage.MoveNode(ctx, "bucket_id", "node_id", &destBucketId, &destNodeId)
```

#### Rename a Node

```go
node, httpResponse, err := client.FileStorage.MoveNode(ctx, "bucket_id", "node_id", "name")
```

## File API

### interface


```
DownloadFile(ciosctx.RequestCtx, string, string) ([]byte, *_nethttp.Response, error)
DownloadFileByKey(ciosctx.RequestCtx, string, string) ([]byte, *_nethttp.Response, error)
UploadFile(ciosctx.RequestCtx, string, []byte, cios.ApiUploadFileRequest) (*_nethttp.Response, error)
```

### Usage

#### Download a File

```go
byts, httpResponse, err := client.FileStorage.DownloadFile(ctx, "bucket_id", "node_id")
```

#### Upload a File

```go
httpResponse, err := client.FileStorage.UploadFile(ctx, "bucket_id", []byte{}, srvfilestorage.MakeUploadFileOpts().NodeId("node_id"))
```