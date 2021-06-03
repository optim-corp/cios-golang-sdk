# Geography

## Point API

### interface

```
GetPoints(ciosctx.RequestCtx, cios.ApiGetPointsRequest) (cios.MultiplePoint, *_nethttp.Response, error)
CreatePoint(ciosctx.RequestCtx, cios.PointRequest) (cios.Point, *_nethttp.Response, error)
DeletePoint(ciosctx.RequestCtx, string) (cios.Point, *_nethttp.Response, error)
```

### Usage

#### Get Points max limit 1000

```go
options := srvgeography.MakeGetPointsOpts
points, httpResponse, err := client.Geography.GetPoints(ctx, options())
```


#### Create a Point

```go
options := srvgeography.MakeGetPointsOpts
points, httpResponse, err := client.Geography.GetPoints(ctx, options())
```


#### Delete a Point

```go
options := srvgeography.MakeGetPointsOpts
points, httpResponse, err := client.Geography.GetPoints(ctx, options())
```

## Cycle API 

WIP

## Route API 

WIP