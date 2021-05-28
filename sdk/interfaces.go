package ciossdk

import (
	_nethttp "net/http"

	"github.com/gorilla/websocket"
	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

type (
	PubSub interface {
		GetChannels(ctx ciosctx.RequestCtx, params cios.ApiGetChannelsRequest) (response cios.MultipleChannel, httpResponse *_nethttp.Response, err error)
		GetChannelsAll(ctx ciosctx.RequestCtx, params cios.ApiGetChannelsRequest) ([]cios.Channel, *_nethttp.Response, error)
		GetChannelsUnlimited(ctx ciosctx.RequestCtx, params cios.ApiGetChannelsRequest) ([]cios.Channel, *_nethttp.Response, error)
		GetChannel(ctx ciosctx.RequestCtx, channelID string, isDev *bool, lang *string) (cios.Channel, *_nethttp.Response, error)
		GetChannelFirst(ctx ciosctx.RequestCtx, params cios.ApiGetChannelsRequest) (cios.Channel, *_nethttp.Response, error)
		GetChannelsMapByID(ctx ciosctx.RequestCtx, params cios.ApiGetChannelsRequest) (map[string]cios.Channel, *_nethttp.Response, error)
		GetChannelsMapByResourceOwnerID(ctx ciosctx.RequestCtx, params cios.ApiGetChannelsRequest) (map[string][]cios.Channel, *_nethttp.Response, error)
		DeleteChannel(ctx ciosctx.RequestCtx, channelID string) (*_nethttp.Response, error)
		GetOrCreateChannel(ctx ciosctx.RequestCtx, params cios.ApiGetChannelsRequest, body cios.ChannelProposal) (cios.Channel, *_nethttp.Response, error)
		CreateChannel(ctx ciosctx.RequestCtx, body cios.ChannelProposal) (cios.Channel, *_nethttp.Response, error)
		UpdateChannel(ctx ciosctx.RequestCtx, channelID string, body cios.ChannelUpdateProposal) (cios.MultipleChannel, *_nethttp.Response, error)
		GetDataStoreChannels(ctx ciosctx.RequestCtx, params cios.ApiGetDataStoreChannelsRequest) (response cios.MultipleDataStoreChannel, httpResponse *_nethttp.Response, err error)
		GetDataStoreChannel(ctx ciosctx.RequestCtx, channelID string) (cios.DataStoreChannel, *_nethttp.Response, error)
		GetObjects(ctx ciosctx.RequestCtx, channelID string, params cios.ApiGetDataStoreObjectsRequest) (response cios.MultipleDataStoreObject, httpResponse *_nethttp.Response, err error)
		GetObjectsAll(ctx ciosctx.RequestCtx, channelID string, params cios.ApiGetDataStoreObjectsRequest) ([]cios.DataStoreObject, *_nethttp.Response, error)
		GetObjectsUnlimited(ctx ciosctx.RequestCtx, channelID string, params cios.ApiGetDataStoreObjectsRequest) ([]cios.DataStoreObject, *_nethttp.Response, error)
		GetObject(ctx ciosctx.RequestCtx, channelID string, objectID string, packerFormat *string) (interface{}, *_nethttp.Response, error)
		GetObjectLatest(ctx ciosctx.RequestCtx, channelID string, packerFormat *string) (interface{}, *_nethttp.Response, error)
		MapObjectLatest(ctx ciosctx.RequestCtx, channelID string, packerFormat *string, stc interface{}) (*_nethttp.Response, error)
		GetMultiObjectLatest(ctx ciosctx.RequestCtx, channelIDs []string) (cios.MultipleDataStoreDataLatest, *_nethttp.Response, error)
		GetMultiObjectLatestByChannels(ctx ciosctx.RequestCtx, channels []cios.Channel) (cios.MultipleDataStoreDataLatest, *_nethttp.Response, error)
		MapMultiObjectLatestPayload(ctx ciosctx.RequestCtx, channelIDs []string, stc interface{}) ([]cios.PackerFormatJsonHeader, *_nethttp.Response, error)
		MapMultiObjectLatestPayloadByChannels(ctx ciosctx.RequestCtx, channels []cios.Channel, stc interface{}) ([]cios.PackerFormatJsonHeader, *_nethttp.Response, error)
		subscribeCiosWebSocket(ctx ciosctx.RequestCtx, _url string, beforeFunc *func(closeError *websocket.Conn), logic func(body []byte) (bool, error), wsR, wsW int64) error
		GetStream(ctx ciosctx.RequestCtx, channelID string, params sdkmodel.ApiGetStreamRequest) ([]string, error)
		GetStreamAll(ctx ciosctx.RequestCtx, channelID string, params sdkmodel.ApiGetStreamRequest) ([]string, error)
		GetStreamUnlimited(ctx ciosctx.RequestCtx, channelID string, params sdkmodel.ApiGetStreamRequest) ([]string, error)
		MapStreamAll(ctx ciosctx.RequestCtx, channelID string, params sdkmodel.ApiGetStreamRequest, stc interface{}) error
		MapStreamUnlimited(ctx ciosctx.RequestCtx, channelID string, params sdkmodel.ApiGetStreamRequest, stc interface{}) error
		GetJsonStreamUnlimited(ctx ciosctx.RequestCtx, channelID string, params sdkmodel.ApiGetStreamRequest) (result []cios.PackerFormatJson, err error)
		GetStreamFirst(ctx ciosctx.RequestCtx, channelID string, params sdkmodel.ApiGetStreamRequest) (string, error)
		MapStreamFirst(ctx ciosctx.RequestCtx, channelID string, params sdkmodel.ApiGetStreamRequest, stc interface{}) error
		DeleteDataByChannel(ctx ciosctx.RequestCtx, channelID string) (*_nethttp.Response, error)
		DeleteObject(ctx ciosctx.RequestCtx, channelID string, objectID string) (*_nethttp.Response, error)
		NewMessaging(channelId, mode, packerFormat string) *CiosMessaging
		PublishMessage(ctx ciosctx.RequestCtx, id string, body interface{}, packerFormat *string) (*_nethttp.Response, error)
		PublishMessagePackerOnly(ctx ciosctx.RequestCtx, id string, body interface{}) (*_nethttp.Response, error)
		PublishMessageJSON(ctx ciosctx.RequestCtx, id string, body cios.PackerFormatJson) (*_nethttp.Response, error)
		ConnectWebSocket(channelID string, done chan bool, params ConnectWebSocketOptions) (err error)
		CreateMessagingURL(channelID string, mode string, packerFormat *string) string
		CreateCIOSWebsocketConnection(url string, authorization string) (connection *websocket.Conn, err error)
		setDebug(bool)
		setToken(string)
	}
	Account interface {
		GetGroups(ctx ciosctx.RequestCtx, params cios.ApiGetGroupsRequest) (response cios.MultipleGroup, httpResponse *_nethttp.Response, err error)
		GetGroupsAll(ctx ciosctx.RequestCtx, params cios.ApiGetGroupsRequest) ([]cios.Group, *_nethttp.Response, error)
		GetGroupsUnlimited(ctx ciosctx.RequestCtx, params cios.ApiGetGroupsRequest) ([]cios.Group, *_nethttp.Response, error)
		GetGroup(ctx ciosctx.RequestCtx, groupId string, includes *string) (cios.Group, *_nethttp.Response, error)
		GetGroupByResourceOwnerId(ctx ciosctx.RequestCtx, resourceOwnerID string, includes *string) (cios.Group, *_nethttp.Response, error)
		GetGroupMapByResourceOwner(ctx ciosctx.RequestCtx, p1 cios.ApiGetGroupsRequest, p2 cios.ApiGetResourceOwnersRequest) (map[string]cios.Group, *_nethttp.Response, error)
		DeleteGroup(ctx ciosctx.RequestCtx, groupID string) (*_nethttp.Response, error)
		CreateGroup(ctx ciosctx.RequestCtx, body cios.GroupCreateRequest) (cios.Group, *_nethttp.Response, error)
		UpdateGroup(ctx ciosctx.RequestCtx, groupID string, body cios.GroupUpdateRequest) (cios.Group, *_nethttp.Response, error)
		GetMe(ctx ciosctx.RequestCtx) (cios.Me, *_nethttp.Response, error)
		InviteGroup(ctx ciosctx.RequestCtx, groupID string, email string) (response cios.Member, httpResponse *_nethttp.Response, err error)
		GetResourceOwners(ctx ciosctx.RequestCtx, params cios.ApiGetResourceOwnersRequest) (response cios.MultipleResourceOwner, httpResponse *_nethttp.Response, err error)
		GetResourceOwnersAll(ctx ciosctx.RequestCtx, params cios.ApiGetResourceOwnersRequest) ([]cios.ResourceOwner, *_nethttp.Response, error)
		GetResourceOwnersUnlimited(ctx ciosctx.RequestCtx, params cios.ApiGetResourceOwnersRequest) ([]cios.ResourceOwner, *_nethttp.Response, error)
		GetResourceOwner(ctx ciosctx.RequestCtx, id string) (cios.ResourceOwner, *_nethttp.Response, error)
		GetResourceOwnerByGroupId(ctx ciosctx.RequestCtx, groupID string) (cios.ResourceOwner, *_nethttp.Response, error)
		GetResourceOwnersMapByID(ctx ciosctx.RequestCtx) (map[string]cios.ResourceOwner, *_nethttp.Response, error)
		GetResourceOwnersMapByGroupID(ctx ciosctx.RequestCtx) (map[string]cios.ResourceOwner, error)
	}
	Auth interface {
		GetAccessTokenByRefreshToken() (sdkmodel.AccessToken, sdkmodel.Scope, sdkmodel.TokenType, sdkmodel.ExpiresIn, error)
		GetAccessTokenOnClient() (sdkmodel.AccessToken, sdkmodel.Scope, sdkmodel.TokenType, sdkmodel.ExpiresIn, error)
		GetAccessTokenOnDevice() (sdkmodel.AccessToken, sdkmodel.Scope, sdkmodel.TokenType, sdkmodel.ExpiresIn, error)

		SetClientSecret(clientSecret string)
		SetClientId(clientId string)
		SetRef(ref string)
		SetAssertion(assertion string)
		SetDebug(debug bool)
		SetScope(scope string)
	}
	Contract interface {
		GetContracts(ctx ciosctx.RequestCtx, params cios.ApiGetContractsRequest) (response cios.MultipleContract, httpResponse *_nethttp.Response, err error)
		GetContractsAll(ctx ciosctx.RequestCtx, params cios.ApiGetContractsRequest) ([]cios.Contract, *_nethttp.Response, error)
		GetContractsUnlimited(ctx ciosctx.RequestCtx, params cios.ApiGetContractsRequest) ([]cios.Contract, *_nethttp.Response, error)
	}
	DeviceAssetManagement interface {
		GetModels(ctx ciosctx.RequestCtx, params cios.ApiGetDeviceModelsRequest) (response cios.MultipleDeviceModel, httpResponse *_nethttp.Response, err error)
		GetModelsAll(ctx ciosctx.RequestCtx, params cios.ApiGetDeviceModelsRequest) ([]cios.DeviceModel, *_nethttp.Response, error)
		GetModelsUnlimited(ctx ciosctx.RequestCtx, params cios.ApiGetDeviceModelsRequest) ([]cios.DeviceModel, *_nethttp.Response, error)
		GetModelsMapByID(ctx ciosctx.RequestCtx, params cios.ApiGetDeviceModelsRequest) (map[string]cios.DeviceModel, error)
		GetModel(ctx ciosctx.RequestCtx, name string) (cios.DeviceModel, *_nethttp.Response, error)
		CreateModel(ctx ciosctx.RequestCtx, body cios.DeviceModelRequest) (cios.DeviceModel, *_nethttp.Response, error)
		DeleteModel(ctx ciosctx.RequestCtx, name string) (*_nethttp.Response, error)
		GetEntities(ctx ciosctx.RequestCtx, params cios.ApiGetDeviceEntitiesRequest) (response cios.MultipleDeviceModelEntity, httpResponse *_nethttp.Response, err error)
		GetEntitiesAll(ctx ciosctx.RequestCtx, params cios.ApiGetDeviceEntitiesRequest) ([]cios.DeviceModelsEntity, *_nethttp.Response, error)
		GetEntitiesUnlimited(ctx ciosctx.RequestCtx, params cios.ApiGetDeviceEntitiesRequest) ([]cios.DeviceModelsEntity, *_nethttp.Response, error)
		GetEntitiesMapByID(ctx ciosctx.RequestCtx, params cios.ApiGetDeviceEntitiesRequest) (map[string]cios.DeviceModelsEntity, error)
		GetEntity(ctx ciosctx.RequestCtx, key string) (cios.DeviceModelsEntity, *_nethttp.Response, error)
		DeleteEntity(ctx ciosctx.RequestCtx, key string) (*_nethttp.Response, error)
		CreateEntity(ctx ciosctx.RequestCtx, name string, body cios.Inventory) (cios.DeviceModelsEntity, *_nethttp.Response, error)
		GetLifecycles(ctx ciosctx.RequestCtx, key string, params cios.ApiGetDeviceEntitiesLifecyclesRequest) (response cios.MultipleLifeCycle, httpResponse *_nethttp.Response, err error)
		GetLifecyclesAll(ctx ciosctx.RequestCtx, key string, params cios.ApiGetDeviceEntitiesLifecyclesRequest) ([]cios.LifeCycle, *_nethttp.Response, error)
		GetLifecyclesUnlimitedByEntities(ctx ciosctx.RequestCtx, entities []cios.DeviceModelsEntity, params cios.ApiGetDeviceEntitiesLifecyclesRequest) ([][]cios.LifeCycle, *_nethttp.Response, error)
		GetLifecyclesUnlimited(ctx ciosctx.RequestCtx, key string, params cios.ApiGetDeviceEntitiesLifecyclesRequest) ([]cios.LifeCycle, *_nethttp.Response, error)
		CreateLifecycle(ctx ciosctx.RequestCtx, key string, body cios.LifeCycleRequest) (cios.LifeCycle, *_nethttp.Response, error)
		DeleteLifecycle(ctx ciosctx.RequestCtx, key string, id string) (*_nethttp.Response, error)
	}
	DeviceManagement interface {
		GetDevices(ctx ciosctx.RequestCtx, params cios.ApiGetDevicesRequest) (response cios.MultipleDevice, httpResponse *_nethttp.Response, err error)
		GetDevicesAll(ctx ciosctx.RequestCtx, params cios.ApiGetDevicesRequest) ([]cios.Device, *_nethttp.Response, error)
		GetDevicesUnlimited(ctx ciosctx.RequestCtx, params cios.ApiGetDevicesRequest) ([]cios.Device, *_nethttp.Response, error)
		GetDevice(ctx ciosctx.RequestCtx, deviceID string, lang *string, isDev *bool) (cios.Device, *_nethttp.Response, error)
		GetDeviceInventory(ctx ciosctx.RequestCtx, deviceID string) (map[string]interface{}, *_nethttp.Response, error)
		DeleteDevice(ctx ciosctx.RequestCtx, id string) (*_nethttp.Response, error)
		CreateDevice(ctx ciosctx.RequestCtx, body cios.DeviceInfo) (cios.Device, *_nethttp.Response, error)
		UpdateDevice(ctx ciosctx.RequestCtx, deviceId string, body cios.DeviceUpdateRequest) (cios.Device, *_nethttp.Response, error)
		GetMonitoringLatestList(ctx ciosctx.RequestCtx, deviceIDs []string) ([]cios.DeviceMonitoring, *_nethttp.Response, error)
		GetMonitoring(ctx ciosctx.RequestCtx, deviceID string) (cios.DeviceMonitoring, *_nethttp.Response, error)
		GetPolicies(ctx ciosctx.RequestCtx, params cios.ApiGetDevicePoliciesRequest) (response cios.MultipleDevicePolicy, httpResponse *_nethttp.Response, err error)
		GetPoliciesAll(ctx ciosctx.RequestCtx, params cios.ApiGetDevicePoliciesRequest) ([]cios.DevicePolicy, *_nethttp.Response, error)
		GetPoliciesUnlimited(ctx ciosctx.RequestCtx, params cios.ApiGetDevicePoliciesRequest) ([]cios.DevicePolicy, *_nethttp.Response, error)
		DeletePolicy(ctx ciosctx.RequestCtx, id string) (*_nethttp.Response, error)
		CreatePolicy(ctx ciosctx.RequestCtx, resourceOwnerID string) (cios.DevicePolicy, *_nethttp.Response, error)
	}
	FileStorage interface {
		GetBuckets(ctx ciosctx.RequestCtx, params cios.ApiGetBucketsRequest) (response cios.MultipleBucket, httpResponse *_nethttp.Response, err error)
		GetBucketsAll(ctx ciosctx.RequestCtx, params cios.ApiGetBucketsRequest) ([]cios.Bucket, *_nethttp.Response, error)
		GetBucketsUnlimited(ctx ciosctx.RequestCtx, params cios.ApiGetBucketsRequest) ([]cios.Bucket, *_nethttp.Response, error)
		GetBucket(ctx ciosctx.RequestCtx, bucketID string) (cios.Bucket, *_nethttp.Response, error)
		GetBucketByResourceOwnerIDAndName(ctx ciosctx.RequestCtx, resourceOwnerID string, name string) (cios.Bucket, *_nethttp.Response, error)
		GetOrCreateBucket(ctx ciosctx.RequestCtx, resourceOwnerID string, name string) (cios.Bucket, *_nethttp.Response, error)
		CreateBucket(ctx ciosctx.RequestCtx, resourceOwnerID string, name string) (cios.Bucket, *_nethttp.Response, error)
		DeleteBucket(ctx ciosctx.RequestCtx, bucketID string) (*_nethttp.Response, error)
		UpdateBucket(ctx ciosctx.RequestCtx, bucketID string, name string) (*_nethttp.Response, error)
		DownloadFile(ctx ciosctx.RequestCtx, bucketID string, nodeID string) ([]byte, *_nethttp.Response, error)
		DownloadFileByKey(ctx ciosctx.RequestCtx, bucketID string, key string) ([]byte, *_nethttp.Response, error)
		UploadFile(ctx ciosctx.RequestCtx, bucketID string, body []byte, params cios.ApiUploadFileRequest) (*_nethttp.Response, error)
		GetNodes(ctx ciosctx.RequestCtx, bucketID string, params cios.ApiGetNodesRequest) (response cios.MultipleNode, httpResponse *_nethttp.Response, err error)
		GetNodesAll(ctx ciosctx.RequestCtx, bucketID string, params cios.ApiGetNodesRequest) ([]cios.Node, *_nethttp.Response, error)
		GetNodesUnlimited(ctx ciosctx.RequestCtx, bucketID string, params cios.ApiGetNodesRequest) ([]cios.Node, *_nethttp.Response, error)
		GetNode(ctx ciosctx.RequestCtx, bucketID string, nodeID string) (cios.Node, *_nethttp.Response, error)
		CreateNode(ctx ciosctx.RequestCtx, bucketID string, name string, parentNodeID *string) (cios.Node, *_nethttp.Response, error)
		CreateNodeOnNodeID(ctx ciosctx.RequestCtx, bucketID string, body cios.NodeRequest) (cios.Node, *_nethttp.Response, error)
		DeleteNode(ctx ciosctx.RequestCtx, bucketID string, nodeID string) (*_nethttp.Response, error)
		CopyNode(ctx ciosctx.RequestCtx, bucketID string, nodeID string, destBucketID *string, destParentNodeID *string) (cios.Node, *_nethttp.Response, error)
		MoveNode(ctx ciosctx.RequestCtx, bucketID string, nodeID string, destBucketID *string, destParentNodeID *string) (cios.Node, *_nethttp.Response, error)
		RenameNode(ctx ciosctx.RequestCtx, bucketID string, nodeID string, name string) (cios.Node, *_nethttp.Response, error)
	}

	Geography interface {
		GetPoints(ctx ciosctx.RequestCtx, params cios.ApiGetPointsRequest) (response cios.MultiplePoint, httpResponse *_nethttp.Response, err error)
		CreatePoint(ctx ciosctx.RequestCtx, body cios.PointRequest) (cios.Point, *_nethttp.Response, error)
		DeletePoint(ctx ciosctx.RequestCtx, pointID string) (cios.Point, *_nethttp.Response, error)
	}
	VideoStreaming interface {
		GetVideoInfos(ctx ciosctx.RequestCtx, params cios.ApiGetVideoStreamsListRequest) ([]cios.Video, *_nethttp.Response, error)
		GetVideoInfo(ctx ciosctx.RequestCtx, videoID string) (cios.Video, *_nethttp.Response, error)
		UpdateVideoInfo(ctx ciosctx.RequestCtx, videoID string, name, description string) (cios.Video, *_nethttp.Response, error)
		GetThumbnail(ctx ciosctx.RequestCtx, videoID string) (img []byte, httpResponse *_nethttp.Response, err error)
		Play(ctx ciosctx.RequestCtx, videoID string) (cios.Room, *_nethttp.Response, error)
		Stop(ctx ciosctx.RequestCtx, videoID string) (*_nethttp.Response, error)
		setToken(string)
	}
	License interface {
		GetLicenses(ctx ciosctx.RequestCtx, params cios.ApiGetMyLicensesRequest) (response []cios.License, httpResponse *_nethttp.Response, err error)
	}
)
