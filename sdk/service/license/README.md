# License

## interface

```
GetLicenses(ciosctx.RequestCtx, cios.ApiGetMyLicensesRequest) ([]cios.License, *_nethttp.Response, error)
```

### Usage

#### Get Licenses

```go
licensed, httpResponse, err := client.License.GetLicenses(ctx, srvlicense.MakeGetLicensesOpts())
```