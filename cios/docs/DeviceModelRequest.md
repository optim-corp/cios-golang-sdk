# DeviceModelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ResourceOwnerId** | **string** |  | 
**MakerId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Watch** | Pointer to [**Watch**](Watch.md) |  | [optional] 
**Components** | Pointer to [**ConstitutionComponent**](ConstitutionComponent.md) |  | [optional] 

## Methods

### NewDeviceModelRequest

`func NewDeviceModelRequest(name string, resourceOwnerId string, ) *DeviceModelRequest`

NewDeviceModelRequest instantiates a new DeviceModelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceModelRequestWithDefaults

`func NewDeviceModelRequestWithDefaults() *DeviceModelRequest`

NewDeviceModelRequestWithDefaults instantiates a new DeviceModelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeviceModelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceModelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceModelRequest) SetName(v string)`

SetName sets Name field to given value.


### GetResourceOwnerId

`func (o *DeviceModelRequest) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *DeviceModelRequest) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *DeviceModelRequest) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.


### GetMakerId

`func (o *DeviceModelRequest) GetMakerId() string`

GetMakerId returns the MakerId field if non-nil, zero value otherwise.

### GetMakerIdOk

`func (o *DeviceModelRequest) GetMakerIdOk() (*string, bool)`

GetMakerIdOk returns a tuple with the MakerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerId

`func (o *DeviceModelRequest) SetMakerId(v string)`

SetMakerId sets MakerId field to given value.

### HasMakerId

`func (o *DeviceModelRequest) HasMakerId() bool`

HasMakerId returns a boolean if a field has been set.

### GetVersion

`func (o *DeviceModelRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceModelRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeviceModelRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeviceModelRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWatch

`func (o *DeviceModelRequest) GetWatch() Watch`

GetWatch returns the Watch field if non-nil, zero value otherwise.

### GetWatchOk

`func (o *DeviceModelRequest) GetWatchOk() (*Watch, bool)`

GetWatchOk returns a tuple with the Watch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatch

`func (o *DeviceModelRequest) SetWatch(v Watch)`

SetWatch sets Watch field to given value.

### HasWatch

`func (o *DeviceModelRequest) HasWatch() bool`

HasWatch returns a boolean if a field has been set.

### GetComponents

`func (o *DeviceModelRequest) GetComponents() ConstitutionComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *DeviceModelRequest) GetComponentsOk() (*ConstitutionComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *DeviceModelRequest) SetComponents(v ConstitutionComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *DeviceModelRequest) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


