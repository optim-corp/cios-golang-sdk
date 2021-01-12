# ChannelProposal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceOwnerId** | **string** |  | 
**ChannelProtocols** | Pointer to [**[]ChannelProtocol**](ChannelProtocol.md) |  | [optional] 
**DisplayInfo** | [**[]DisplayInfo**](DisplayInfo.md) |  | 
**Labels** | Pointer to [**[]Label**](Label.md) |  | [optional] 
**MessagingConfig** | Pointer to [**MessagingConfig**](MessagingConfig.md) |  | [optional] 
**DatastoreConfig** | Pointer to [**DataStoreConfig**](DataStoreConfig.md) |  | [optional] 

## Methods

### NewChannelProposal

`func NewChannelProposal(resourceOwnerId string, displayInfo []DisplayInfo, ) *ChannelProposal`

NewChannelProposal instantiates a new ChannelProposal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelProposalWithDefaults

`func NewChannelProposalWithDefaults() *ChannelProposal`

NewChannelProposalWithDefaults instantiates a new ChannelProposal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceOwnerId

`func (o *ChannelProposal) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *ChannelProposal) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *ChannelProposal) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.


### GetChannelProtocols

`func (o *ChannelProposal) GetChannelProtocols() []ChannelProtocol`

GetChannelProtocols returns the ChannelProtocols field if non-nil, zero value otherwise.

### GetChannelProtocolsOk

`func (o *ChannelProposal) GetChannelProtocolsOk() (*[]ChannelProtocol, bool)`

GetChannelProtocolsOk returns a tuple with the ChannelProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProtocols

`func (o *ChannelProposal) SetChannelProtocols(v []ChannelProtocol)`

SetChannelProtocols sets ChannelProtocols field to given value.

### HasChannelProtocols

`func (o *ChannelProposal) HasChannelProtocols() bool`

HasChannelProtocols returns a boolean if a field has been set.

### GetDisplayInfo

`func (o *ChannelProposal) GetDisplayInfo() []DisplayInfo`

GetDisplayInfo returns the DisplayInfo field if non-nil, zero value otherwise.

### GetDisplayInfoOk

`func (o *ChannelProposal) GetDisplayInfoOk() (*[]DisplayInfo, bool)`

GetDisplayInfoOk returns a tuple with the DisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInfo

`func (o *ChannelProposal) SetDisplayInfo(v []DisplayInfo)`

SetDisplayInfo sets DisplayInfo field to given value.


### GetLabels

`func (o *ChannelProposal) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ChannelProposal) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ChannelProposal) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ChannelProposal) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMessagingConfig

`func (o *ChannelProposal) GetMessagingConfig() MessagingConfig`

GetMessagingConfig returns the MessagingConfig field if non-nil, zero value otherwise.

### GetMessagingConfigOk

`func (o *ChannelProposal) GetMessagingConfigOk() (*MessagingConfig, bool)`

GetMessagingConfigOk returns a tuple with the MessagingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingConfig

`func (o *ChannelProposal) SetMessagingConfig(v MessagingConfig)`

SetMessagingConfig sets MessagingConfig field to given value.

### HasMessagingConfig

`func (o *ChannelProposal) HasMessagingConfig() bool`

HasMessagingConfig returns a boolean if a field has been set.

### GetDatastoreConfig

`func (o *ChannelProposal) GetDatastoreConfig() DataStoreConfig`

GetDatastoreConfig returns the DatastoreConfig field if non-nil, zero value otherwise.

### GetDatastoreConfigOk

`func (o *ChannelProposal) GetDatastoreConfigOk() (*DataStoreConfig, bool)`

GetDatastoreConfigOk returns a tuple with the DatastoreConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreConfig

`func (o *ChannelProposal) SetDatastoreConfig(v DataStoreConfig)`

SetDatastoreConfig sets DatastoreConfig field to given value.

### HasDatastoreConfig

`func (o *ChannelProposal) HasDatastoreConfig() bool`

HasDatastoreConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


