# DeviceModelUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ResourceOwnerId** | Pointer to **string** |  | [optional] 
**MakerId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Watch** | Pointer to [**DeviceModelUpdateRequestWatch**](DeviceModelUpdateRequest_watch.md) |  | [optional] 
**Components** | Pointer to [**[]ConstitutionComponent**](ConstitutionComponent.md) |  | [optional] 

## Methods

### NewDeviceModelUpdateRequest

`func NewDeviceModelUpdateRequest() *DeviceModelUpdateRequest`

NewDeviceModelUpdateRequest instantiates a new DeviceModelUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceModelUpdateRequestWithDefaults

`func NewDeviceModelUpdateRequestWithDefaults() *DeviceModelUpdateRequest`

NewDeviceModelUpdateRequestWithDefaults instantiates a new DeviceModelUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeviceModelUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceModelUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceModelUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceModelUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceOwnerId

`func (o *DeviceModelUpdateRequest) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *DeviceModelUpdateRequest) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *DeviceModelUpdateRequest) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.

### HasResourceOwnerId

`func (o *DeviceModelUpdateRequest) HasResourceOwnerId() bool`

HasResourceOwnerId returns a boolean if a field has been set.

### GetMakerId

`func (o *DeviceModelUpdateRequest) GetMakerId() string`

GetMakerId returns the MakerId field if non-nil, zero value otherwise.

### GetMakerIdOk

`func (o *DeviceModelUpdateRequest) GetMakerIdOk() (*string, bool)`

GetMakerIdOk returns a tuple with the MakerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerId

`func (o *DeviceModelUpdateRequest) SetMakerId(v string)`

SetMakerId sets MakerId field to given value.

### HasMakerId

`func (o *DeviceModelUpdateRequest) HasMakerId() bool`

HasMakerId returns a boolean if a field has been set.

### GetVersion

`func (o *DeviceModelUpdateRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceModelUpdateRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeviceModelUpdateRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeviceModelUpdateRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWatch

`func (o *DeviceModelUpdateRequest) GetWatch() DeviceModelUpdateRequestWatch`

GetWatch returns the Watch field if non-nil, zero value otherwise.

### GetWatchOk

`func (o *DeviceModelUpdateRequest) GetWatchOk() (*DeviceModelUpdateRequestWatch, bool)`

GetWatchOk returns a tuple with the Watch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatch

`func (o *DeviceModelUpdateRequest) SetWatch(v DeviceModelUpdateRequestWatch)`

SetWatch sets Watch field to given value.

### HasWatch

`func (o *DeviceModelUpdateRequest) HasWatch() bool`

HasWatch returns a boolean if a field has been set.

### GetComponents

`func (o *DeviceModelUpdateRequest) GetComponents() []ConstitutionComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *DeviceModelUpdateRequest) GetComponentsOk() (*[]ConstitutionComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *DeviceModelUpdateRequest) SetComponents(v []ConstitutionComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *DeviceModelUpdateRequest) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


