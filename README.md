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
|Device|✅|
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

### APIs

* [Account](./sdk/service/account/README.md)
  * [Me API](./sdk/service/account/README.md#me-api)
    * [Get Me](./sdk/service/account/README.md#get-me)
  * [Group API](./sdk/service/account/README.md#group-api)
    * [Get a Group](./sdk/service/account/README.md#group-api)
    * [Get a Group by ResourceOwnerID](./sdk/service/account/README.md#get-a-group-by-resource-owner-id)
    * [Get Groups limit 1000](./sdk/service/account/README.md#get-groups-limit-1000)
    * [Get Groups no limit](./sdk/service/account/README.md#get-groups-no-limit)
    * [Get Groups unlimited](./sdk/service/account/README.md#get-groups-unlimited)
    * [Create a Group](./sdk/service/account/README.md#crete-a-group)
    * [Update a Group](./sdk/service/account/README.md#update-a-group)
    * [Delete a Group](./sdk/service/account/README.md#delete-a-group)
  * [ResourceOwner API](./sdk/service/account/README.md#resource-owner-api)
    * [Get a ResourceOwner](./sdk/service/account/README.md#resource-owner-api)
    * [Get a ResourceOwner by Group ID](./sdk/service/account/README.md#get-a-resourceowner-by-group-id)
    * [Get ResourceOwner Map (Key ResourceOwner ID)](./sdk/service/account/README.md#get-resourceowner-map-by-resourceowner-id)
    * [Get ResourceOwner Map (Key Group ID)](./sdk/service/account/README.md#get-resourceowner-map-by-group-id)
    * [Get ResourceOwners](./sdk/service/account/README.md#get-resourceowners)
    * [Get ResourceOwners limit 1000](./sdk/service/account/README.md#get-resourceowners-limit-1000)
    * [Get ResourceOwners no limit](./sdk/service/account/README.md#get-groups-no-limit)
    * [Get ResourceOwners unlimited](./sdk/service/account/README.md#get-resourceowners-unlimited)

## How to Support

If you have any issues or questions, please raise them on [Github issues](https://github.com/optim-corp/cios-golang-sdk/issues).

## LICENSE

[The Apache-2.0 License](https://www.apache.org/licenses/LICENSE-2.0)

Copyright (c) 2021 OPTiM Corporation <https://www.optim.co.jp/>
