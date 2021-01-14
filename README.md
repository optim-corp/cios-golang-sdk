# CIOS Golang SDK

It is a convenient SDK to use [Cloud IoT OS](https://www.optim.cloud/platform/) in Go language.

This SDK uses OpenAPI and is generated in code.
OpenAPI will be split at a later date, but is now included in the code.

## Instration

```shell script
go get -u github.com/optim-corp/cios-golang-sdk
```

## Document

[CIOS Golang SDK Documents](./cios/README.md)
## Usage


### Initialization

```go
client = ciossdk.NewCiosClient(&ciossdk.CiosClientConfig{
    Debug:             true,
    ClientID:          "client ID",
    ClientSecret:      "client secret",
    RequestScope:      "hogehoge.hoge",
    RefreshToken:      "refresh token",
    AutoRefresh:       true,
    Urls: &ciossdk.CIOSUrl{
        MonitoringUrl:            " https://",
        MessagingUrl:             " https://",
        LocationUrl:              " https://",
        AccountsUrl:              " https://",
        StorageUrl:               " https://",
        IamUrl:                   " https://",
        AuthUrl:                  " https://",
        VideoStreamingUrl:        " https://",
        DeviceManagementUrl:      " https://",
        DeviceAssetManagementUrl: " https://",
    },
})
```


### [Examples] Get Buckets

```go
ctx := ciossdk.MakeRequestCtx("Bearer Hogehoge")) // or context.Background()
params := ciossdk.MakeGetBucketsOpts(ctx).
    Name("sample").
    Limit(20)
buckets, httpResponse, err :=  client.FileStorage.GetBucketsUnlimited(params)
```

## LICENSE

[The Apache-2.0 License](https://www.apache.org/licenses/LICENSE-2.0)

Copyright (c) 2021 OPTiM Corporation <https://www.optim.co.jp/>