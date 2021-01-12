# Channel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**ResourceOwnerId** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ChannelProtocols** | Pointer to [**[]ChannelProtocol**](ChannelProtocol.md) |  | [optional] 
**DisplayInfo** | Pointer to [**[]DisplayInfo**](DisplayInfo.md) |  | [optional] 
**Labels** | Pointer to [**[]Label**](Label.md) |  | [optional] 
**MessagingConfig** | Pointer to [**MessagingConfig**](MessagingConfig.md) |  | [optional] 
**DatastoreConfig** | Pointer to [**DataStoreConfig**](DataStoreConfig.md) |  | [optional] 

## Methods

### NewChannel

`func NewChannel(id string, createdAt string, updatedAt string, resourceOwnerId string, name string, ) *Channel`

NewChannel instantiates a new Channel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelWithDefaults

`func NewChannelWithDefaults() *Channel`

NewChannelWithDefaults instantiates a new Channel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Channel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Channel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Channel) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Channel) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Channel) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Channel) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Channel) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Channel) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Channel) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetResourceOwnerId

`func (o *Channel) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *Channel) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *Channel) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.


### GetName

`func (o *Channel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Channel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Channel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Channel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Channel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Channel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Channel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetChannelProtocols

`func (o *Channel) GetChannelProtocols() []ChannelProtocol`

GetChannelProtocols returns the ChannelProtocols field if non-nil, zero value otherwise.

### GetChannelProtocolsOk

`func (o *Channel) GetChannelProtocolsOk() (*[]ChannelProtocol, bool)`

GetChannelProtocolsOk returns a tuple with the ChannelProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProtocols

`func (o *Channel) SetChannelProtocols(v []ChannelProtocol)`

SetChannelProtocols sets ChannelProtocols field to given value.

### HasChannelProtocols

`func (o *Channel) HasChannelProtocols() bool`

HasChannelProtocols returns a boolean if a field has been set.

### GetDisplayInfo

`func (o *Channel) GetDisplayInfo() []DisplayInfo`

GetDisplayInfo returns the DisplayInfo field if non-nil, zero value otherwise.

### GetDisplayInfoOk

`func (o *Channel) GetDisplayInfoOk() (*[]DisplayInfo, bool)`

GetDisplayInfoOk returns a tuple with the DisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInfo

`func (o *Channel) SetDisplayInfo(v []DisplayInfo)`

SetDisplayInfo sets DisplayInfo field to given value.

### HasDisplayInfo

`func (o *Channel) HasDisplayInfo() bool`

HasDisplayInfo returns a boolean if a field has been set.

### GetLabels

`func (o *Channel) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Channel) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Channel) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Channel) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMessagingConfig

`func (o *Channel) GetMessagingConfig() MessagingConfig`

GetMessagingConfig returns the MessagingConfig field if non-nil, zero value otherwise.

### GetMessagingConfigOk

`func (o *Channel) GetMessagingConfigOk() (*MessagingConfig, bool)`

GetMessagingConfigOk returns a tuple with the MessagingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingConfig

`func (o *Channel) SetMessagingConfig(v MessagingConfig)`

SetMessagingConfig sets MessagingConfig field to given value.

### HasMessagingConfig

`func (o *Channel) HasMessagingConfig() bool`

HasMessagingConfig returns a boolean if a field has been set.

### GetDatastoreConfig

`func (o *Channel) GetDatastoreConfig() DataStoreConfig`

GetDatastoreConfig returns the DatastoreConfig field if non-nil, zero value otherwise.

### GetDatastoreConfigOk

`func (o *Channel) GetDatastoreConfigOk() (*DataStoreConfig, bool)`

GetDatastoreConfigOk returns a tuple with the DatastoreConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreConfig

`func (o *Channel) SetDatastoreConfig(v DataStoreConfig)`

SetDatastoreConfig sets DatastoreConfig field to given value.

### HasDatastoreConfig

`func (o *Channel) HasDatastoreConfig() bool`

HasDatastoreConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


