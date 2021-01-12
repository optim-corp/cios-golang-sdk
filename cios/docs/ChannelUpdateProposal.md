# ChannelUpdateProposal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelProtocols** | Pointer to [**[]ChannelProtocol**](ChannelProtocol.md) |  | [optional] 
**DisplayInfo** | [**[]DisplayInfo**](DisplayInfo.md) |  | 
**Labels** | Pointer to [**[]Label**](Label.md) |  | [optional] 
**MessagingConfig** | Pointer to [**MessagingConfig**](MessagingConfig.md) |  | [optional] 
**DatastoreConfig** | Pointer to [**DataStoreConfig**](DataStoreConfig.md) |  | [optional] 

## Methods

### NewChannelUpdateProposal

`func NewChannelUpdateProposal(displayInfo []DisplayInfo, ) *ChannelUpdateProposal`

NewChannelUpdateProposal instantiates a new ChannelUpdateProposal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelUpdateProposalWithDefaults

`func NewChannelUpdateProposalWithDefaults() *ChannelUpdateProposal`

NewChannelUpdateProposalWithDefaults instantiates a new ChannelUpdateProposal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelProtocols

`func (o *ChannelUpdateProposal) GetChannelProtocols() []ChannelProtocol`

GetChannelProtocols returns the ChannelProtocols field if non-nil, zero value otherwise.

### GetChannelProtocolsOk

`func (o *ChannelUpdateProposal) GetChannelProtocolsOk() (*[]ChannelProtocol, bool)`

GetChannelProtocolsOk returns a tuple with the ChannelProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProtocols

`func (o *ChannelUpdateProposal) SetChannelProtocols(v []ChannelProtocol)`

SetChannelProtocols sets ChannelProtocols field to given value.

### HasChannelProtocols

`func (o *ChannelUpdateProposal) HasChannelProtocols() bool`

HasChannelProtocols returns a boolean if a field has been set.

### GetDisplayInfo

`func (o *ChannelUpdateProposal) GetDisplayInfo() []DisplayInfo`

GetDisplayInfo returns the DisplayInfo field if non-nil, zero value otherwise.

### GetDisplayInfoOk

`func (o *ChannelUpdateProposal) GetDisplayInfoOk() (*[]DisplayInfo, bool)`

GetDisplayInfoOk returns a tuple with the DisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInfo

`func (o *ChannelUpdateProposal) SetDisplayInfo(v []DisplayInfo)`

SetDisplayInfo sets DisplayInfo field to given value.


### GetLabels

`func (o *ChannelUpdateProposal) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ChannelUpdateProposal) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ChannelUpdateProposal) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ChannelUpdateProposal) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMessagingConfig

`func (o *ChannelUpdateProposal) GetMessagingConfig() MessagingConfig`

GetMessagingConfig returns the MessagingConfig field if non-nil, zero value otherwise.

### GetMessagingConfigOk

`func (o *ChannelUpdateProposal) GetMessagingConfigOk() (*MessagingConfig, bool)`

GetMessagingConfigOk returns a tuple with the MessagingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingConfig

`func (o *ChannelUpdateProposal) SetMessagingConfig(v MessagingConfig)`

SetMessagingConfig sets MessagingConfig field to given value.

### HasMessagingConfig

`func (o *ChannelUpdateProposal) HasMessagingConfig() bool`

HasMessagingConfig returns a boolean if a field has been set.

### GetDatastoreConfig

`func (o *ChannelUpdateProposal) GetDatastoreConfig() DataStoreConfig`

GetDatastoreConfig returns the DatastoreConfig field if non-nil, zero value otherwise.

### GetDatastoreConfigOk

`func (o *ChannelUpdateProposal) GetDatastoreConfigOk() (*DataStoreConfig, bool)`

GetDatastoreConfigOk returns a tuple with the DatastoreConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreConfig

`func (o *ChannelUpdateProposal) SetDatastoreConfig(v DataStoreConfig)`

SetDatastoreConfig sets DatastoreConfig field to given value.

### HasDatastoreConfig

`func (o *ChannelUpdateProposal) HasDatastoreConfig() bool`

HasDatastoreConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


