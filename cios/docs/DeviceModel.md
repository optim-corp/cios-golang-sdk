# DeviceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**ResourceOwnerId** | **string** |  | 
**MakerId** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**Watch** | Pointer to [**NullableWatch**](Watch.md) |  | [optional] 
**Components** | Pointer to [**ConstitutionComponent**](ConstitutionComponent.md) |  | [optional] 
**CreatedAt** | **string** | ナノ秒 | 
**UpdatedAt** | **string** | ナノ秒 | 

## Methods

### NewDeviceModel

`func NewDeviceModel(id string, name string, resourceOwnerId string, createdAt string, updatedAt string, ) *DeviceModel`

NewDeviceModel instantiates a new DeviceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceModelWithDefaults

`func NewDeviceModelWithDefaults() *DeviceModel`

NewDeviceModelWithDefaults instantiates a new DeviceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DeviceModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceModel) SetName(v string)`

SetName sets Name field to given value.


### GetResourceOwnerId

`func (o *DeviceModel) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *DeviceModel) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *DeviceModel) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.


### GetMakerId

`func (o *DeviceModel) GetMakerId() string`

GetMakerId returns the MakerId field if non-nil, zero value otherwise.

### GetMakerIdOk

`func (o *DeviceModel) GetMakerIdOk() (*string, bool)`

GetMakerIdOk returns a tuple with the MakerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerId

`func (o *DeviceModel) SetMakerId(v string)`

SetMakerId sets MakerId field to given value.

### HasMakerId

`func (o *DeviceModel) HasMakerId() bool`

HasMakerId returns a boolean if a field has been set.

### SetMakerIdNil

`func (o *DeviceModel) SetMakerIdNil(b bool)`

 SetMakerIdNil sets the value for MakerId to be an explicit nil

### UnsetMakerId
`func (o *DeviceModel) UnsetMakerId()`

UnsetMakerId ensures that no value is present for MakerId, not even an explicit nil
### GetVersion

`func (o *DeviceModel) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceModel) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeviceModel) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeviceModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *DeviceModel) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *DeviceModel) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetWatch

`func (o *DeviceModel) GetWatch() Watch`

GetWatch returns the Watch field if non-nil, zero value otherwise.

### GetWatchOk

`func (o *DeviceModel) GetWatchOk() (*Watch, bool)`

GetWatchOk returns a tuple with the Watch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatch

`func (o *DeviceModel) SetWatch(v Watch)`

SetWatch sets Watch field to given value.

### HasWatch

`func (o *DeviceModel) HasWatch() bool`

HasWatch returns a boolean if a field has been set.

### SetWatchNil

`func (o *DeviceModel) SetWatchNil(b bool)`

 SetWatchNil sets the value for Watch to be an explicit nil

### UnsetWatch
`func (o *DeviceModel) UnsetWatch()`

UnsetWatch ensures that no value is present for Watch, not even an explicit nil
### GetComponents

`func (o *DeviceModel) GetComponents() ConstitutionComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *DeviceModel) GetComponentsOk() (*ConstitutionComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *DeviceModel) SetComponents(v ConstitutionComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *DeviceModel) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DeviceModel) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeviceModel) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeviceModel) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DeviceModel) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeviceModel) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeviceModel) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


