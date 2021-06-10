# Pub/Sub

## Channel API

### interface

```
GetChannels(ctx, cios.ApiGetChannelsRequest) (cios.MultipleChannel, *_nethttp.Response, error)
GetChannelsAll(ctx, cios.ApiGetChannelsRequest) ([]cios.Channel, *_nethttp.Response, error)
GetChannelsUnlimited(ctx, cios.ApiGetChannelsRequest) ([]cios.Channel, *_nethttp.Response, error)
GetChannel(ctx, string, *bool, *string) (cios.Channel, *_nethttp.Response, error)
GetChannelFirst(ctx, cios.ApiGetChannelsRequest) (cios.Channel, *_nethttp.Response, error)
GetChannelsMapByID(ctx, cios.ApiGetChannelsRequest) (map[string]cios.Channel, *_nethttp.Response, error)
GetChannelsMapByResourceOwnerID(ctx, cios.ApiGetChannelsRequest) (map[string][]cios.Channel, *_nethttp.Response, error)
DeleteChannel(ctx, string) (*_nethttp.Response, error)
GetOrCreateChannel(ctx, cios.ApiGetChannelsRequest, cios.ChannelProposal) (cios.Channel, *_nethttp.Response, error)
CreateChannel(ctx, cios.ChannelProposal) (cios.Channel, *_nethttp.Response, error)
UpdateChannel(ctx, string, cios.ChannelUpdateProposal) (cios.MultipleChannel, *_nethttp.Response, error)
```

### Usage

#### Get a Channel

```go
options := srvpubsub.MakeGetChannelsOpts
channel, httpResponse, err := client.PubSub().GetChannel(ctx, "channel_id", nil, nil)
```

#### Get Channels max limit 1000

```go
options := srvpubsub.MakeGetChannelsOpts
channels, httpResponse, err := client.PubSub().GetChannels(ctx, options().Limit(500))
```

#### Get Channels no limit

```go
options := srvpubsub.MakeGetChannelsOpts
channels, httpResponse, err := client.PubSub().GetChannelsAll(ctx, options().Limit(500))
```

#### Get Channels unlimited

```go
options := srvpubsub.MakeGetChannelsOpts
channels, httpResponse, err := client.PubSub().GetChannelsUnlimited(ctx, options().Limit(500))
```


#### Get a Channel that is first Channels API

```go
options := srvpubsub.MakeGetChannelsOpts
channel, httpResponse, err := client.PubSub().GetChannelFirst(ctx, options().Label("sample=test"))
```

#### Get Channel Map by Channel ID

```go
options := srvpubsub.MakeGetChannelsOpts
channelMap, httpResponse, err := client.PubSub().GetChannelsMapByID(ctx, options().Label("sample=test"))
```

#### Get Channel Map by ResourceOwner ID

```go
options := srvpubsub.MakeGetChannelsOpts
channelMap, httpResponse, err := client.PubSub().GetChannelsMapByResourceOwnerID(ctx, options().Label("sample=test"))
```

#### Get or Create a Channel

```go
options := srvpubsub.MakeGetChannelsOpts
channel, httpResponse, err := client.PubSub().GetOrCreateChannel(ctx, options().Limit(500), cios.ChannelProposal{})
```

#### Create a Channel

```go
channel, httpResponse, err := client.PubSub().CreateChannel(ctx,  cios.ChannelProposal{})
```


#### Update a Channel

```go
_, httpResponse, err := client.PubSub().UpdateChannel(ctx,  cios.ChannelUpdateProposal{})
```

#### Delete a Channel

```go
_, httpResponse, err := client.PubSub().UpdateChannel(ctx,  cios.ChannelUpdateProposal{})
```

## Data Store API

### interface 

```
GetDataStoreChannels(ctx, cios.ApiGetDataStoreChannelsRequest) (cios.MultipleDataStoreChannel, *_nethttp.Response, error)
GetDataStoreChannel(ctx, string) (cios.DataStoreChannel, *_nethttp.Response, error)
GetObjects(ctx, string, cios.ApiGetDataStoreObjectsRequest) (cios.MultipleDataStoreObject, *_nethttp.Response, error)
GetObjectsAll(ctx, string, cios.ApiGetDataStoreObjectsRequest) ([]cios.DataStoreObject, *_nethttp.Response, error)
GetObjectsUnlimited(ctx, string, cios.ApiGetDataStoreObjectsRequest) ([]cios.DataStoreObject, *_nethttp.Response, error)
GetObject(ctx, string, string, *string) (interface{}, *_nethttp.Response, error)
GetObjectLatest(ctx, string, *string) (interface{}, *_nethttp.Response, error)
MapObjectLatest(ctx, string, *string, interface{}) (*_nethttp.Response, error)
GetMultiObjectLatest(ctx, []string) (cios.MultipleDataStoreDataLatest, *_nethttp.Response, error)
GetMultiObjectLatestByChannels(ctx, []cios.Channel) (cios.MultipleDataStoreDataLatest, *_nethttp.Response, error)
MapMultiObjectLatestPayload(ctx, []string, interface{}) ([]cios.PackerFormatJsonHeader, *_nethttp.Response, error)
MapMultiObjectLatestPayloadByChannels(ctx, []cios.Channel, interface{}) ([]cios.PackerFormatJsonHeader, *_nethttp.Response, error)
GetStream(ctx, string, sdkmodel.ApiGetStreamRequest) ([]string, error)
GetStreamAll(ctx, string, sdkmodel.ApiGetStreamRequest) ([]string, error)
GetStreamUnlimited(ctx, string, sdkmodel.ApiGetStreamRequest) ([]string, error)
MapStreamAll(ctx, string, sdkmodel.ApiGetStreamRequest, interface{}) error
MapStreamUnlimited(ctx, string, sdkmodel.ApiGetStreamRequest, interface{}) error
GetJsonStreamUnlimited(ctx, string, sdkmodel.ApiGetStreamRequest) ([]cios.PackerFormatJson, error)
GetStreamFirst(ctx, string, sdkmodel.ApiGetStreamRequest) (string, error)
MapStreamFirst(ctx, string, sdkmodel.ApiGetStreamRequest, interface{}) error
DeleteDataByChannel(ctx, string) (*_nethttp.Response, error)
DeleteObject(ctx, string, string) (*_nethttp.Response, error)
```

### Usage

#### Get Channels on Data Store max limit 1000

```go
options := srvpubsub.MakeGetDataStoreChannelsOpts()
channels, httpResponse, err := client.PubSub().GetDataStoreChannels(ctx, options().Limit(1000))
```

#### Get a Channel on Data Store 

```go
channel, httpResponse, err := client.PubSub().GetDataStoreChannel(ctx, "channel_id")
```

#### Get a DataStore Object

※ Fix in the future
```go
packer_format := "payload_only"
object, httpResponse, err := client.PubSub().GetObject(ctx, "channel_id", "object_id", &packer_format)
```

#### Get latest DataStore Object 

※ Fix in the future
```go
packer_format := "payload_only"
object, httpResponse, err := client.PubSub().GetObjectLatest(ctx, "channel_id",  &packer_format)
```

#### Get latest DataStore Object on Channels

```go
objects, httpResponse, err := client.PubSub().GetMultiObjectLatest(ctx, []string{"channel_id1", "channel_id2", "channel_id3"})
```

#### Get latest DataStore Object by Channels

```go
objects, httpResponse, err := client.PubSub().GetMultiObjectLatest(ctx, []cios.Channel{channel1, channel2, channel3})
```

#### Get DataStore Objects max limit 1000

```go
options := srvpubsub.MakeGetObjectsOpts
objects, httpResponse, err := client.PubSub().GetObjects(ctx, "channel_id", options())
```

#### Get DataStore Objects no limit

```go
options := srvpubsub.MakeGetObjectsOpts
objects, httpResponse, err := client.PubSub().GetObjectsAll(ctx, "channel_id", options())
```

#### Get DataStore Objects unlimited

```go
options := srvpubsub.MakeGetObjectsOpts
objects, httpResponse, err := client.PubSub().GetObjectsUnlimited(ctx, "channel_id", options())
```

#### Parse latest for Object

```go
packerFormat := "payload_only"
sample := struct{....}{}
httpResponse, err := client.PubSub().MapObjectLatest(ctx, "channel_id", &packerFormat, &sample)
```

#### Parse latest for Object on channels

※ Fix in the future

```go
packerFormat := "payload_only"
sample := struct{....}{}
httpResponse, err := client.PubSub().MapMultiObjectLatestPayload(ctx, []string{"channel_id", "channel_id"}, &sample)
```

#### Get DataStore Stream max limit 1000 

```go
options := srvpubsub.MakeGetStreamOpts
stream, err := client.PubSub().GetStream(ctx, "channel_id", options())
```

#### Get DataStore Stream no limit

```go
options := srvpubsub.MakeGetStreamOpts
stream, err := client.PubSub().GetStreamAll(ctx, "channel_id", options())
```

#### Get DataStore Stream unlimited

```go
options := srvpubsub.MakeGetStreamOpts
stream, err := client.PubSub().GetStreamUnimited(ctx, "channel_id", options())
```
#### Get a DataStore Stream first

```go
options := srvpubsub.MakeGetStreamOpts
data, err := client.PubSub().GetStreamFirst(ctx, "channel_id", options())
```

#### Parse DataStore Stream no limit

```go
data := []struct{...}{}
options := srvpubsub.MakeGetStreamOpts
err := client.PubSub().MapStreamAll(ctx, "channel_id", options(), &data)
```

#### Parse DataStore Stream unlimited

```go
data := []struct{...}{}
options := srvpubsub.MakeGetStreamOpts
err := client.PubSub().MapStreamUnlimited(ctx, "channel_id", options(), &data)
```

#### Map a DataStore Stream first

```go
options := srvpubsub.MakeGetStreamOpts
data := struct{}{}
err := client.PubSub().GetStreamFirst(ctx, "channel_id", options(), &data)
```

#### Delete DataStore Objects

```go
httpResponse, err := client.PubSub().DeleteDataByChannel(ctx, "channel_id")
```

#### Delete a DataStore Object

```go
httpResponse, err := client.PubSub().DeleteObject(ctx, "channel_id", "object_id")
```

## Messaging API

### interface

```
NewMessaging(string, enum.MessagingMode, enum.PackerFormat) *srvpubsub.CiosMessaging
PublishMessage(ciosctx.RequestCtx, string, interface{}, *string) (*_nethttp.Response, error)
PublishMessagePackerOnly(ciosctx.RequestCtx, string, interface{}) (*_nethttp.Response, error)
PublishMessageJSON(ciosctx.RequestCtx, string, cios.PackerFormatJson) (*_nethttp.Response, error)
```

### Usage

#### Publish Message

※ Fix in the future

```go
packerFormat := "payload_only"
data := struct{}{}
httpResponse, err := client.PubSub().PublishMessage(ctx, "channel_id", data, &packetFormat)
```


#### Publish payload only format

※ Fix in the future
```go
data := struct{}{}
httpResponse, err := client.PubSub().PublishMessagePackerOnly(ctx, "channel_id", data)
```

#### Publish json format

```go
data := struct{}{}
httpResponse, err := client.PubSub().PublishMessageJSON(ctx, "channel_id", data)
```


#### Create Messaging instance

```go
ms := client.PubSub().NewMessaging("channel_id", enum.PubSub, enum.Json)
err := ms.Start(ctx)
...
ms.Close()
```

#### About Messaging Instance

※ Fix in the future

What doesn't exist interface yet.

* Send a Message 
    ```
    Send(message []byte) (err error)
    SendStr(message string) error
    SendAny(message interface{}) error
    SendJson(message interface{}) error
    Publish(message interface{}) error
    ```
    ```go
    err := ms.Send([]byte{})
    err := ms.SendStr("sample")
    err := ms.SendJson(struct{}{})
    ```

* Receive a Message

    ```
    Receive() (body []byte, err error)
    ReceiveStr() (string, error)
    MapReceived(stct interface{})
    ```
    ```go
    byts, err := ms.Receive()
    str, err := ms.ReceiveStr()
    data struct{}
    err := ms.MapReceived(&data)
    ```