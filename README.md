# Go CIOS Client SDK!!!!

## Go Get!!!

```shell script
go get -u github.com/optim-corp/cios-golang-sdk
```

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


### Get Buckets

```go
ctx := ciossdk.MakeRequestCtx("Bearer Hogehoge")) // or context.Background()
params := ciossdk.MakeGetBucketsOpts(ctx).
    Name("sample").
    Limit(20)
buckets, httpResponse, err :=  client.FileStorage.GetBucketsUnlimited(params)
```
