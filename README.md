# OPTiM Cloud IoT OS Golang SDK (0.4.2)


It is a convenient SDK to use [OPTiM Cloud IoT OS(CIOS)](https://www.optim.cloud/platform/) in Go language.

This SDK uses OpenAPI and is generated in code.
OpenAPI will be split at a later date, but is now included in the code.
It is imperfect.

## Instration

```shell script
go get -u github.com/optim-corp/cios-golang-sdk
```

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

## Document

### [Authorization](./sdk/service/authorization/README.md)

* [How to set AccessToken in context](./sdk/service/authorization/README.md#how-to-set-token)
* [Auto Refresh AccessToken for instance](./sdk/service/authorization/README.md#how-to-auto-refresh-auth)

### [Account](./sdk/service/account/README.md)

* [Me API](./sdk/service/account/README.md#me-api)
  * [Get Me](./sdk/service/account/README.md#get-me)
* [Group API](./sdk/service/account/README.md#group-api)
  * [Get a Group](./sdk/service/account/README.md#group-api)
  * [Get a Group by ResourceOwnerID](./sdk/service/account/README.md#get-a-group-by-resource-owner-id)
  * [Get Groups (max limit 1000)](./sdk/service/account/README.md#get-groups-limit-1000)
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
  * [Get ResourceOwners (max limit 1000)](./sdk/service/account/README.md#get-resourceowners-limit-1000)
  * [Get ResourceOwners no limit](./sdk/service/account/README.md#get-groups-no-limit)
  * [Get ResourceOwners unlimited](./sdk/service/account/README.md#get-resourceowners-unlimited)

### [FileStorage](./sdk/service/filestorage/README.md)

* [Bucket API](./sdk/service/filestorage/README.md#bucket-api)
  * [Get a Bucket](./sdk/service/filestorage/README.md#get-a-bucket)
  * [Get a Bucket by ResourceOwnerID and  Name](./sdk/service/filestorage/README.md#get-a-bucket-by-resourceownerid--bucket-name)
  * [Get or Create a Bucket](./sdk/service/filestorage/README.md#get-or-create-a-bucket)
  * [Get Buckets (max limit 1000)](./sdk/service/filestorage/README.md#get-buckets-limit-1000)
  * [Get Buckets no limit](./sdk/service/filestorage/README.md#get-buckets-no-limit)
  * [Get Buckets unlimited](./sdk/service/filestorage/README.md#get-buckets-unlimited)
  * [Create a Bucket](./sdk/service/filestorage/README.md#create-a-bucket)
  * [Update a Bucket](./sdk/service/filestorage/README.md#update-a-bucket)
  * [Delete a Bucket](./sdk/service/filestorage/README.md#delete-a-bucket)
* [Node API](./sdk/service/filestorage/README.md#node-api)
  * [Get a Node](./sdk/service/filestorage/README.md#get-a-node)
  * [Get Nodes (max limit 1000)](./sdk/service/filestorage/README.md#get-nodes-limit-1000)
  * [Get Nodes no limit](./sdk/service/filestorage/README.md#get-nodes-no-limit)
  * [Get Nodes unlimited](./sdk/service/filestorage/README.md#get-nodes-unlimited)
  * [Create a Directory](./sdk/service/filestorage/README.md#create-a-directory)
  * [Copy a Node](./sdk/service/filestorage/README.md#copy-a-node)
  * [Move a Node](./sdk/service/filestorage/README.md#move-a-node)
  * [Rename a Node](./sdk/service/filestorage/README.md#rename-a-node)
  * [Delete a Node](./sdk/service/filestorage/README.md#delete-a-node)
* [File API](./sdk/service/filestorage/README.md#file-api)
  * [Download a File](./sdk/service/filestorage/README.md#download-a-file)
  * [Upload a File](./sdk/service/filestorage/README.md#upload-a-file)

### [Pub/Sub](./sdk/service/pubsub/README.md)

* [Channel API](./sdk/service/pubsub/README.md#channel-api)
  * [Get a Channel](./sdk/service/pubsub/README.md#get-a-channel)
  * [Get a Channel first on channel list response](./sdk/service/pubsub/README.md#get-a-channel-that-is-first-channels-api)
  * [Get or Create Channel](./sdk/service/pubsub/README.md#get-or-create-a-channel)
  * [Get Channels (max limit 1000)](./sdk/service/pubsub/README.md#get-channels-max-limit-1000)
  * [Get Channels no limit](./sdk/service/pubsub/README.md#get-channels-no-limit)
  * [Get Channels unlimited](./sdk/service/pubsub/README.md#get-channels-unlimited)
  * [Get Channel map by Channel ID key](./sdk/service/pubsub/README.md#get-channel-map-by-channel-id)
  * [Get Channel map by ResourceOwner ID key](./sdk/service/pubsub/README.md#get-channel-map-by-resourceowner-id)
  * [Create a Channel](./sdk/service/pubsub/README.md#create-a-channel)
  * [Update a Channel](./sdk/service/pubsub/README.md#update-a-channel)
  * [Delete a Channel](./sdk/service/pubsub/README.md#delete-a-channel)
* [DataStore API](./sdk/service/pubsub/README.md#data-store-api)
  * [Get a Channel for DataStore](./sdk/service/pubsub/README.md#get-a-channel-on-data-store)
  * [Get Channels for DataStore](./sdk/service/pubsub/README.md#get-channels-on-data-store-max-limit-1000)
  * [Get a Object](./sdk/service/pubsub/README.md#get-a-datastore-object)
  * [Get latest Object](./sdk/service/pubsub/README.md#get-latest-datastore-object)
  * [Get latest Object on Channels](./sdk/service/pubsub/README.md#get-latest-datastore-object-on-channels)
  * [Get Objects info (max limit 1000)](./sdk/service/pubsub/README.md#get-datastore-objects-max-limit-1000)
  * [Get Objects info no limit](./sdk/service/pubsub/README.md#get-datastore-objects-no-limit)
  * [Get Objects info unlimited](./sdk/service/pubsub/README.md#get-datastore-objects-unlimited)
  * [Parse latest Object](./sdk/service/pubsub/README.md#parse-latest-for-object)
  * [Parse latest Object on channels](./sdk/service/pubsub/README.md#parse-latest-for-object-on-channels)
  * [Get Stream first (websocket)](./sdk/service/pubsub/README.md#get-a-datastore-stream-first)
  * [Get Stream (max limit 1000) (websocket)](./sdk/service/pubsub/README.md#get-datastore-stream-max-limit-1000)
  * [Get Stream no limit (websocket)](./sdk/service/pubsub/README.md#get-datastore-stream-no-limit)
  * [Get Stream unlimited (websocket)](./sdk/service/pubsub/README.md#get-datastore-stream-unlimited)
  * [Parse Stream no limit (websocket)](./sdk/service/pubsub/README.md#parse-datastore-stream-no-limit)
  * [Parse Stream unlimited (websocket)](./sdk/service/pubsub/README.md#parse-datastore-stream-unlimited)
  * [Parse Stream first (websocket)](./sdk/service/pubsub/README.md#map-a-datastore-stream-first)
  * [Delete a Object](./sdk/service/pubsub/README.md#delete-a-datastore-object)
  * [Delete Objects on Channel](./sdk/service/pubsub/README.md#delete-datastore-objects)
* [Messaging API](./sdk/service/pubsub/README.md#messaging-api)
  * [Publish a message](./sdk/service/pubsub/README.md#publish-message)
  * [Publish a message (payload only format)](./sdk/service/pubsub/README.md#publish-payload-only-format)
  * [Publish a message (json format)](./sdk/service/pubsub/README.md#publish-json-format)
  * [Create Messaging instance (WIP)](./sdk/service/pubsub/README.md#create-messaging-instance)
    * [How to use messaging instance (WIP)](./sdk/service/pubsub/README.md#about-messaging-instance)

### [Device](./sdk/service/device/README.md)

* [Device Management API](./sdk/service/device/README.md#device-management-api)
  * [Get a Device](./sdk/service/device/README.md#get-a-device)
  * [Get a Device Inventory](./sdk/service/device/README.md#get-a-device-inventory)
  * [Get Devices (max limit 1000)](./sdk/service/device/README.md#get-devices-max-limit-1000)
  * [Get Devices no limit](./sdk/service/device/README.md#get-devices-max-no-limit)
  * [Get Devices unlimited](./sdk/service/device/README.md#get-devices-max-unlimited)
  * [Create a Device](./sdk/service/device/README.md#create-a-device)
  * [Update a Device](./sdk/service/device/README.md#update-a-device)
  * [Device a Device](./sdk/service/device/README.md#delete-a-device)
* [Device Monitoring API](./sdk/service/device/README.md#device-monitoring-api)
  * [Get a Device Monitoring](./sdk/service/device/README.md#get-a-device-monitoring-info)
  * [Get Devices Monitoring](./sdk/service/device/README.md#get-devices-monitoring-info)
* [Device Policy](./sdk/service/device/README.md#device-policy-api)
  * [Get Policies (max limit 1000)](./sdk/service/device/README.md#get-policies-max-limit-1000)
  * [Get Policies no limit](./sdk/service/device/README.md#get-policies-max-no-limit)
  * [Get Policies unlimited](./sdk/service/device/README.md#get-policies-unlimited)
  * [Create a Policy](./sdk/service/device/README.md#create-a-policy)
  * [Delete a Policy](./sdk/service/device/README.md#delete-a-policy)

### [Video](./sdk/service/video/README.md)

* [Video Streaming API](./sdk/service/video/README.md#video-streaming-api)
  * [Get a Video information](./sdk/service/video/README.md#get-a-video-information)
  * [Get Videos information](./sdk/service/video/README.md#get-videos-information)
  * [Get a Thumbnail](./sdk/service/video/README.md#get-a-thumbnail)
  * [Play Streaming](./sdk/service/video/README.md#play-streaming)
  * [Stop Streaming](./sdk/service/video/README.md#stop-streaming)

### [Geography](./sdk/service/geography/README.md)

* [Point API (WIP)](./sdk/service/geography/README.md#point-api)
  * [Get Points (max limit 1000)](./sdk/service/geography/README.md#get-points-max-limit-1000)
  * [Create a Point](./sdk/service/geography/README.md#create-a-point)
  * [Delete a Point](./sdk/service/geography/README.md#delete-a-point)
* [Route API (WIP)](./sdk/service/geography/README.md#route-api)
* [Cycle API (WIP)](./sdk/service/geography/README.md#cycle-api)

### [Collection (WIP)](./sdk/service/collection/README.md)

### [License](./sdk/service/license/README.md)

* [License API](./sdk/service/license/README.md#license)
  * [Get Licenses](./sdk/service/license/README.md#get-licenses)

### [Contract](./sdk/service/contract/README.md)

  * [Contract API](./sdk/service/contract/README.md#contract-api)
    * [Get Contracts (max limit 1000)](./sdk/service/contract/README.md#get-contracts-max-limit-1000)
    * [Get Contracts no limit](./sdk/service/contract/README.md#get-contracts-no-limit)
    * [Get Contracts unlimited](./sdk/service/contract/README.md#get-contracts-unlimited)

## How to Unit test

package ciossdk_mock has mock interface

### Usage

```go
type MockAccount struct {
	ciossdk_mock.NoImplementAccount
}

func (MockAccount) GetMe(ctx ciosctx.RequestCtx) (cios.Me, *http.Response, error) {
	return testMe, nil, nil
}
func (MockAccount) GetResourceOwnersMapByGroupID(ciosctx.RequestCtx) (map[string]cios.ResourceOwner, error) {
	return nil, nil
}
func TestExample(t *testing.T) {
    client := &ciossdk.CiosClient{}
    client.Account = MockAccount{}
}
```

## How to Support

If you have any issues or questions, please raise them on [Github issues](https://github.com/optim-corp/cios-golang-sdk/issues).

## LICENSE

[The Apache-2.0 License](https://www.apache.org/licenses/LICENSE-2.0)

Copyright (c) 2021 OPTiM Corporation <https://www.optim.co.jp/>
