# Pub/Sub

## Channel API

### interface

```
GetChannels(ciosctx.RequestCtx, cios.ApiGetChannelsRequest) (cios.MultipleChannel, *_nethttp.Response, error)
GetChannelsAll(ciosctx.RequestCtx, cios.ApiGetChannelsRequest) ([]cios.Channel, *_nethttp.Response, error)
GetChannelsUnlimited(ciosctx.RequestCtx, cios.ApiGetChannelsRequest) ([]cios.Channel, *_nethttp.Response, error)
GetChannel(ciosctx.RequestCtx, string, *bool, *string) (cios.Channel, *_nethttp.Response, error)
GetChannelFirst(ciosctx.RequestCtx, cios.ApiGetChannelsRequest) (cios.Channel, *_nethttp.Response, error)
GetChannelsMapByID(ciosctx.RequestCtx, cios.ApiGetChannelsRequest) (map[string]cios.Channel, *_nethttp.Response, error)
GetChannelsMapByResourceOwnerID(ciosctx.RequestCtx, cios.ApiGetChannelsRequest) (map[string][]cios.Channel, *_nethttp.Response, error)
DeleteChannel(ciosctx.RequestCtx, string) (*_nethttp.Response, error)
GetOrCreateChannel(ciosctx.RequestCtx, cios.ApiGetChannelsRequest, cios.ChannelProposal) (cios.Channel, *_nethttp.Response, error)
CreateChannel(ciosctx.RequestCtx, cios.ChannelProposal) (cios.Channel, *_nethttp.Response, error)
UpdateChannel(ciosctx.RequestCtx, string, cios.ChannelUpdateProposal) (cios.MultipleChannel, *_nethttp.Response, error)
```

### Usage

#### Get Channels max limit 1000

```go
options := srvpubsub.MakeGetChannelsOpts
channels, httpResponse, err := client.PubSub.GetChannels(ctx, options().Limit(500))
```

#### Get Channels no limit

```go
options := srvpubsub.MakeGetChannelsOpts
channels, httpResponse, err := client.PubSub.GetChannelsAll(ctx, options().Limit(500))
```

#### Get Channels unlimited

```go
options := srvpubsub.MakeGetChannelsOpts
channels, httpResponse, err := client.PubSub.GetChannelsUnlimited(ctx, options().Limit(500))
```

#### Get a Channel

```go
options := srvpubsub.MakeGetChannelsOpts
channel, httpResponse, err := client.PubSub.GetChannel(ctx, "channel_id", nil, nil)
```

#### Get a Channel that is first Channels API

```go
options := srvpubsub.MakeGetChannelsOpts
channel, httpResponse, err := client.PubSub.GetChannelFirst(ctx, options().Label("sample=test"))
```

#### Get Channel Map by Channel ID

```go
options := srvpubsub.MakeGetChannelsOpts
channelMap, httpResponse, err := client.PubSub.GetChannelsMapByID(ctx, options().Label("sample=test"))
```

#### Get Channel Map by ResourceOwner ID

```go
options := srvpubsub.MakeGetChannelsOpts
channelMap, httpResponse, err := client.PubSub.GetChannelsMapByResourceOwnerID(ctx, options().Label("sample=test"))
```

#### Get or Create a Channel

```go
options := srvpubsub.MakeGetChannelsOpts
channel, httpResponse, err := client.PubSub.GetOrCreateChannel(ctx, options().Limit(500), cios.ChannelProposal{})
```

#### Create a Channel

```go
channel, httpResponse, err := client.PubSub.CreateChannel(ctx,  cios.ChannelProposal{})
```


#### Update a Channel

```go
_, httpResponse, err := client.PubSub.UpdateChannel(ctx,  cios.ChannelUpdateProposal{})
```

#### Delete a Channel

```go
_, httpResponse, err := client.PubSub.UpdateChannel(ctx,  cios.ChannelUpdateProposal{})
```

## 