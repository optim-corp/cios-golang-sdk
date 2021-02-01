# OPTiM Cloud IoT OS Golang SDK

It is a convenient SDK to use [OPTiM Cloud IoT OS(CIOS)](https://www.optim.cloud/platform/) in Go language.

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

* If you want to use Prod Client

```go
client = ciossdk.NewCiosClient(ciossdk.CiosClientConfig{
    Debug:             true,
    AutoRefresh:       true,
    Urls:              sdkmodel.ProdUrls(),
})
```

* If you want to use Pre Client

```go
client = ciossdk.NewCiosClient(ciossdk.CiosClientConfig{
    Debug:             true,
    AutoRefresh:       true,
    Urls:              sdkmodel.PreUrls(),
})
```

* If you want to use Custom URL

```go
client = ciossdk.NewCiosClient(ciossdk.CiosClientConfig{
    Debug:             true,
    AutoRefresh:       true,
    Urls: ciossdk.CIOSUrl{
        MessagingUrl:             " https://",
        LocationUrl:              " https://",
        AccountsUrl:              " https://",
        StorageUrl:               " https://",
        IamUrl:                   " https://",
        AuthUrl:                  " https://",
        VideoStreamingUrl:        " https://",
        DeviceManagementUrl:      " https://",
        DeviceMonitoringUrl:      " https://",
        DeviceAssetManagementUrl: " https://",
    },
})
```
#### Authentication

|OAuthClientType|is Suppoted|
|---|---|
|Web|✅|
|Client|✅|
|Device|https://github.com/optim-corp/cios-golang-sdk/issues/2|
|Native|❌|


* Client Auth 
    * client id
    * client secret
    * request scope
```go
client = ciossdk.NewCiosClient(ciossdk.CiosClientConfig{
    AuthConfig: ClientAuthConf(
           "clientID",
           "clientSecret",
           "scope",
   ),
})
```

* Refresh Token Auth
    * client id
    * client secret
    * refresh token
    * request scope
```go
client = ciossdk.NewCiosClient(ciossdk.CiosClientConfig{
    AuthConfig: RefreshTokenAuth(
           "clientID",
           "clientSecret",
           "refreshToken",
           "scope",
   ),
})
```

* Device Auth
    * client id
    * client secret
    * assertion
    * request scope
```go
client = ciossdk.NewCiosClient(ciossdk.CiosClientConfig{
    AuthConfig: DeviceAuthConf(
           "clientID",
           "clientSecret",
           "assertion",
           "scope",
   ),
})
```

The refresh token will only be executed once if the request fails.

### [Examples] Get Buckets

```go
ctx := ciossdk.MakeRequestCtx("Bearer Hogehoge")) // or context.Background()
params := ciossdk.MakeGetBucketsOpts().
    Name("sample").
    Limit(20)
buckets, httpResponse, err :=  client.FileStorage.GetBucketsUnlimited(params, ctx)
```

## How to Support

If you have any issues or questions, please raise them on [Github issues](https://github.com/optim-corp/cios-golang-sdk/issues).

## LICENSE

[The Apache-2.0 License](https://www.apache.org/licenses/LICENSE-2.0)

Copyright (c) 2021 OPTiM Corporation <https://www.optim.co.jp/>
