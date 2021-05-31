package ciossdk

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/sdk/enum"

	srvpubsub "github.com/optim-corp/cios-golang-sdk/sdk/service/pubsub"

	"github.com/gorilla/websocket"
	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
)

type (
	PubSub interface {
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
		GetDataStoreChannels(ciosctx.RequestCtx, cios.ApiGetDataStoreChannelsRequest) (cios.MultipleDataStoreChannel, *_nethttp.Response, error)
		GetDataStoreChannel(ciosctx.RequestCtx, string) (cios.DataStoreChannel, *_nethttp.Response, error)
		GetObjects(ciosctx.RequestCtx, string, cios.ApiGetDataStoreObjectsRequest) (cios.MultipleDataStoreObject, *_nethttp.Response, error)
		GetObjectsAll(ciosctx.RequestCtx, string, cios.ApiGetDataStoreObjectsRequest) ([]cios.DataStoreObject, *_nethttp.Response, error)
		GetObjectsUnlimited(ciosctx.RequestCtx, string, cios.ApiGetDataStoreObjectsRequest) ([]cios.DataStoreObject, *_nethttp.Response, error)
		GetObject(ciosctx.RequestCtx, string, string, *string) (interface{}, *_nethttp.Response, error)
		GetObjectLatest(ciosctx.RequestCtx, string, *string) (interface{}, *_nethttp.Response, error)
		MapObjectLatest(ciosctx.RequestCtx, string, *string, interface{}) (*_nethttp.Response, error)
		GetMultiObjectLatest(ciosctx.RequestCtx, []string) (cios.MultipleDataStoreDataLatest, *_nethttp.Response, error)
		GetMultiObjectLatestByChannels(ciosctx.RequestCtx, []cios.Channel) (cios.MultipleDataStoreDataLatest, *_nethttp.Response, error)
		MapMultiObjectLatestPayload(ciosctx.RequestCtx, []string, interface{}) ([]cios.PackerFormatJsonHeader, *_nethttp.Response, error)
		MapMultiObjectLatestPayloadByChannels(ciosctx.RequestCtx, []cios.Channel, interface{}) ([]cios.PackerFormatJsonHeader, *_nethttp.Response, error)
		GetStream(ciosctx.RequestCtx, string, sdkmodel.ApiGetStreamRequest) ([]string, error)
		GetStreamAll(ciosctx.RequestCtx, string, sdkmodel.ApiGetStreamRequest) ([]string, error)
		GetStreamUnlimited(ciosctx.RequestCtx, string, sdkmodel.ApiGetStreamRequest) ([]string, error)
		MapStreamAll(ciosctx.RequestCtx, string, sdkmodel.ApiGetStreamRequest, interface{}) error
		MapStreamUnlimited(ciosctx.RequestCtx, string, sdkmodel.ApiGetStreamRequest, interface{}) error
		GetJsonStreamUnlimited(ciosctx.RequestCtx, string, sdkmodel.ApiGetStreamRequest) ([]cios.PackerFormatJson, error)
		GetStreamFirst(ciosctx.RequestCtx, string, sdkmodel.ApiGetStreamRequest) (string, error)
		MapStreamFirst(ciosctx.RequestCtx, string, sdkmodel.ApiGetStreamRequest, interface{}) error
		DeleteDataByChannel(ciosctx.RequestCtx, string) (*_nethttp.Response, error)
		DeleteObject(ciosctx.RequestCtx, string, string) (*_nethttp.Response, error)
		NewMessaging(string, enum.MessagingMode, enum.PackerFormat) *srvpubsub.CiosMessaging
		PublishMessage(ciosctx.RequestCtx, string, interface{}, *string) (*_nethttp.Response, error)
		PublishMessagePackerOnly(ciosctx.RequestCtx, string, interface{}) (*_nethttp.Response, error)
		PublishMessageJSON(ciosctx.RequestCtx, string, cios.PackerFormatJson) (*_nethttp.Response, error)
		// Deprecated: should not be used
		ConnectWebSocket(string, chan bool, srvpubsub.ConnectWebSocketOptions) error
		// Deprecated: should not be used
		CreateMessagingURL(string, string, *string) string
		// Deprecated: should not be used
		CreateCIOSWebsocketConnection(string, string) (*websocket.Conn, error)
		SetDebug(bool)
		SetToken(string)
	}
	Account interface {
		GetGroups(ciosctx.RequestCtx, cios.ApiGetGroupsRequest) (cios.MultipleGroup, *_nethttp.Response, error)
		GetGroupsAll(ciosctx.RequestCtx, cios.ApiGetGroupsRequest) ([]cios.Group, *_nethttp.Response, error)
		GetGroupsUnlimited(ciosctx.RequestCtx, cios.ApiGetGroupsRequest) ([]cios.Group, *_nethttp.Response, error)
		GetGroup(ciosctx.RequestCtx, string, *string) (cios.Group, *_nethttp.Response, error)
		GetGroupByResourceOwnerId(ciosctx.RequestCtx, string, *string) (cios.Group, *_nethttp.Response, error)
		GetGroupMapByResourceOwner(ciosctx.RequestCtx, cios.ApiGetGroupsRequest, cios.ApiGetResourceOwnersRequest) (map[string]cios.Group, *_nethttp.Response, error)
		DeleteGroup(ciosctx.RequestCtx, string) (*_nethttp.Response, error)
		CreateGroup(ciosctx.RequestCtx, cios.GroupCreateRequest) (cios.Group, *_nethttp.Response, error)
		UpdateGroup(ciosctx.RequestCtx, string, cios.GroupUpdateRequest) (cios.Group, *_nethttp.Response, error)
		GetMe(ciosctx.RequestCtx) (cios.Me, *_nethttp.Response, error)
		InviteGroup(ciosctx.RequestCtx, string, string) (cios.Member, *_nethttp.Response, error)
		GetResourceOwners(ciosctx.RequestCtx, cios.ApiGetResourceOwnersRequest) (cios.MultipleResourceOwner, *_nethttp.Response, error)
		GetResourceOwnersAll(ciosctx.RequestCtx, cios.ApiGetResourceOwnersRequest) ([]cios.ResourceOwner, *_nethttp.Response, error)
		GetResourceOwnersUnlimited(ciosctx.RequestCtx, cios.ApiGetResourceOwnersRequest) ([]cios.ResourceOwner, *_nethttp.Response, error)
		GetResourceOwner(ciosctx.RequestCtx, string) (cios.ResourceOwner, *_nethttp.Response, error)
		GetResourceOwnerByGroupId(ciosctx.RequestCtx, string) (cios.ResourceOwner, *_nethttp.Response, error)
		GetResourceOwnersMapByID(ciosctx.RequestCtx) (map[string]cios.ResourceOwner, *_nethttp.Response, error)
		GetResourceOwnersMapByGroupID(ciosctx.RequestCtx) (map[string]cios.ResourceOwner, error)
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
		GetContracts(ciosctx.RequestCtx, cios.ApiGetContractsRequest) (cios.MultipleContract, *_nethttp.Response, error)
		GetContractsAll(ciosctx.RequestCtx, cios.ApiGetContractsRequest) ([]cios.Contract, *_nethttp.Response, error)
		GetContractsUnlimited(ciosctx.RequestCtx, cios.ApiGetContractsRequest) ([]cios.Contract, *_nethttp.Response, error)
	}
	DeviceAssetManagement interface {
		GetModels(ciosctx.RequestCtx, cios.ApiGetDeviceModelsRequest) (cios.MultipleDeviceModel, *_nethttp.Response, error)
		GetModelsAll(ciosctx.RequestCtx, cios.ApiGetDeviceModelsRequest) ([]cios.DeviceModel, *_nethttp.Response, error)
		GetModelsUnlimited(ciosctx.RequestCtx, cios.ApiGetDeviceModelsRequest) ([]cios.DeviceModel, *_nethttp.Response, error)
		GetModelsMapByID(ciosctx.RequestCtx, cios.ApiGetDeviceModelsRequest) (map[string]cios.DeviceModel, error)
		GetModel(ciosctx.RequestCtx, string) (cios.DeviceModel, *_nethttp.Response, error)
		CreateModel(ciosctx.RequestCtx, cios.DeviceModelRequest) (cios.DeviceModel, *_nethttp.Response, error)
		DeleteModel(ciosctx.RequestCtx, string) (*_nethttp.Response, error)
		GetEntities(ciosctx.RequestCtx, cios.ApiGetDeviceEntitiesRequest) (cios.MultipleDeviceModelEntity, *_nethttp.Response, error)
		GetEntitiesAll(ciosctx.RequestCtx, cios.ApiGetDeviceEntitiesRequest) ([]cios.DeviceModelsEntity, *_nethttp.Response, error)
		GetEntitiesUnlimited(ciosctx.RequestCtx, cios.ApiGetDeviceEntitiesRequest) ([]cios.DeviceModelsEntity, *_nethttp.Response, error)
		GetEntitiesMapByID(ciosctx.RequestCtx, cios.ApiGetDeviceEntitiesRequest) (map[string]cios.DeviceModelsEntity, error)
		GetEntity(ciosctx.RequestCtx, string) (cios.DeviceModelsEntity, *_nethttp.Response, error)
		DeleteEntity(ciosctx.RequestCtx, string) (*_nethttp.Response, error)
		CreateEntity(ciosctx.RequestCtx, string, cios.Inventory) (cios.DeviceModelsEntity, *_nethttp.Response, error)
		GetLifecycles(ciosctx.RequestCtx, string, cios.ApiGetDeviceEntitiesLifecyclesRequest) (cios.MultipleLifeCycle, *_nethttp.Response, error)
		GetLifecyclesAll(ciosctx.RequestCtx, string, cios.ApiGetDeviceEntitiesLifecyclesRequest) ([]cios.LifeCycle, *_nethttp.Response, error)
		GetLifecyclesUnlimitedByEntities(ciosctx.RequestCtx, []cios.DeviceModelsEntity, cios.ApiGetDeviceEntitiesLifecyclesRequest) ([][]cios.LifeCycle, *_nethttp.Response, error)
		GetLifecyclesUnlimited(ciosctx.RequestCtx, string, cios.ApiGetDeviceEntitiesLifecyclesRequest) ([]cios.LifeCycle, *_nethttp.Response, error)
		CreateLifecycle(ciosctx.RequestCtx, string, cios.LifeCycleRequest) (cios.LifeCycle, *_nethttp.Response, error)
		DeleteLifecycle(ciosctx.RequestCtx, string, string) (*_nethttp.Response, error)
	}
	DeviceManagement interface {
		GetDevices(ciosctx.RequestCtx, cios.ApiGetDevicesRequest) (cios.MultipleDevice, *_nethttp.Response, error)
		GetDevicesAll(ciosctx.RequestCtx, cios.ApiGetDevicesRequest) ([]cios.Device, *_nethttp.Response, error)
		GetDevicesUnlimited(ciosctx.RequestCtx, cios.ApiGetDevicesRequest) ([]cios.Device, *_nethttp.Response, error)
		GetDevice(ciosctx.RequestCtx, string, *string, *bool) (cios.Device, *_nethttp.Response, error)
		GetDeviceInventory(ciosctx.RequestCtx, string) (map[string]interface{}, *_nethttp.Response, error)
		DeleteDevice(ciosctx.RequestCtx, string) (*_nethttp.Response, error)
		CreateDevice(ciosctx.RequestCtx, cios.DeviceInfo) (cios.Device, *_nethttp.Response, error)
		UpdateDevice(ciosctx.RequestCtx, string, cios.DeviceUpdateRequest) (cios.Device, *_nethttp.Response, error)
		GetMonitoringLatestList(ciosctx.RequestCtx, []string) ([]cios.DeviceMonitoring, *_nethttp.Response, error)
		GetMonitoring(ciosctx.RequestCtx, string) (cios.DeviceMonitoring, *_nethttp.Response, error)
		GetPolicies(ciosctx.RequestCtx, cios.ApiGetDevicePoliciesRequest) (cios.MultipleDevicePolicy, *_nethttp.Response, error)
		GetPoliciesAll(ciosctx.RequestCtx, cios.ApiGetDevicePoliciesRequest) ([]cios.DevicePolicy, *_nethttp.Response, error)
		GetPoliciesUnlimited(ciosctx.RequestCtx, cios.ApiGetDevicePoliciesRequest) ([]cios.DevicePolicy, *_nethttp.Response, error)
		DeletePolicy(ciosctx.RequestCtx, string) (*_nethttp.Response, error)
		CreatePolicy(ciosctx.RequestCtx, string) (cios.DevicePolicy, *_nethttp.Response, error)
	}
	FileStorage interface {
		GetBuckets(ciosctx.RequestCtx, cios.ApiGetBucketsRequest) (cios.MultipleBucket, *_nethttp.Response, error)
		GetBucketsAll(ciosctx.RequestCtx, cios.ApiGetBucketsRequest) ([]cios.Bucket, *_nethttp.Response, error)
		GetBucketsUnlimited(ciosctx.RequestCtx, cios.ApiGetBucketsRequest) ([]cios.Bucket, *_nethttp.Response, error)
		GetBucket(ciosctx.RequestCtx, string) (cios.Bucket, *_nethttp.Response, error)
		GetBucketByResourceOwnerIDAndName(ciosctx.RequestCtx, string, string) (cios.Bucket, *_nethttp.Response, error)
		GetOrCreateBucket(ciosctx.RequestCtx, string, string) (cios.Bucket, *_nethttp.Response, error)
		CreateBucket(ciosctx.RequestCtx, string, string) (cios.Bucket, *_nethttp.Response, error)
		DeleteBucket(ciosctx.RequestCtx, string) (*_nethttp.Response, error)
		UpdateBucket(ciosctx.RequestCtx, string, string) (*_nethttp.Response, error)
		DownloadFile(ciosctx.RequestCtx, string, string) ([]byte, *_nethttp.Response, error)
		DownloadFileByKey(ciosctx.RequestCtx, string, string) ([]byte, *_nethttp.Response, error)
		UploadFile(ciosctx.RequestCtx, string, []byte, cios.ApiUploadFileRequest) (*_nethttp.Response, error)
		GetNodes(ciosctx.RequestCtx, string, cios.ApiGetNodesRequest) (cios.MultipleNode, *_nethttp.Response, error)
		GetNodesAll(ciosctx.RequestCtx, string, cios.ApiGetNodesRequest) ([]cios.Node, *_nethttp.Response, error)
		GetNodesUnlimited(ciosctx.RequestCtx, string, cios.ApiGetNodesRequest) ([]cios.Node, *_nethttp.Response, error)
		GetNode(ciosctx.RequestCtx, string, string) (cios.Node, *_nethttp.Response, error)
		CreateNode(ciosctx.RequestCtx, string, string, *string) (cios.Node, *_nethttp.Response, error)
		CreateNodeOnNodeID(ciosctx.RequestCtx, string, cios.NodeRequest) (cios.Node, *_nethttp.Response, error)
		DeleteNode(ciosctx.RequestCtx, string, string) (*_nethttp.Response, error)
		CopyNode(ciosctx.RequestCtx, string, string, *string, *string) (cios.Node, *_nethttp.Response, error)
		MoveNode(ciosctx.RequestCtx, string, string, *string, *string) (cios.Node, *_nethttp.Response, error)
		RenameNode(ciosctx.RequestCtx, string, string, string) (cios.Node, *_nethttp.Response, error)
	}

	Geography interface {
		GetPoints(ciosctx.RequestCtx, cios.ApiGetPointsRequest) (cios.MultiplePoint, *_nethttp.Response, error)
		CreatePoint(ciosctx.RequestCtx, cios.PointRequest) (cios.Point, *_nethttp.Response, error)
		DeletePoint(ciosctx.RequestCtx, string) (cios.Point, *_nethttp.Response, error)
	}
	VideoStreaming interface {
		GetVideoInfos(ciosctx.RequestCtx, cios.ApiGetVideoStreamsListRequest) ([]cios.Video, *_nethttp.Response, error)
		GetVideoInfo(ciosctx.RequestCtx, string) (cios.Video, *_nethttp.Response, error)
		UpdateVideoInfo(ciosctx.RequestCtx, string, string, string) (cios.Video, *_nethttp.Response, error)
		GetThumbnail(ciosctx.RequestCtx, string) ([]byte, *_nethttp.Response, error)
		Play(ciosctx.RequestCtx, string) (cios.Room, *_nethttp.Response, error)
		Stop(ciosctx.RequestCtx, string) (*_nethttp.Response, error)
		SetToken(string)
	}
	License interface {
		GetLicenses(ciosctx.RequestCtx, cios.ApiGetMyLicensesRequest) ([]cios.License, *_nethttp.Response, error)
	}
)
