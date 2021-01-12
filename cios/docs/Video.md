# Video

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**DeviceId** | Pointer to **string** |  | [optional] 
**ResourceOwnerId** | **string** |  | 
**VideoName** | Pointer to **string** |  | [optional] 
**VideoDescription** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewVideo

`func NewVideo(id string, resourceOwnerId string, createdAt string, updatedAt string, ) *Video`

NewVideo instantiates a new Video object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoWithDefaults

`func NewVideoWithDefaults() *Video`

NewVideoWithDefaults instantiates a new Video object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Video) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Video) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Video) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Video) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *Video) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Video) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Video) SetId(v string)`

SetId sets Id field to given value.


### GetDeviceId

`func (o *Video) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *Video) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *Video) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *Video) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetResourceOwnerId

`func (o *Video) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *Video) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *Video) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.


### GetVideoName

`func (o *Video) GetVideoName() string`

GetVideoName returns the VideoName field if non-nil, zero value otherwise.

### GetVideoNameOk

`func (o *Video) GetVideoNameOk() (*string, bool)`

GetVideoNameOk returns a tuple with the VideoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoName

`func (o *Video) SetVideoName(v string)`

SetVideoName sets VideoName field to given value.

### HasVideoName

`func (o *Video) HasVideoName() bool`

HasVideoName returns a boolean if a field has been set.

### GetVideoDescription

`func (o *Video) GetVideoDescription() string`

GetVideoDescription returns the VideoDescription field if non-nil, zero value otherwise.

### GetVideoDescriptionOk

`func (o *Video) GetVideoDescriptionOk() (*string, bool)`

GetVideoDescriptionOk returns a tuple with the VideoDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoDescription

`func (o *Video) SetVideoDescription(v string)`

SetVideoDescription sets VideoDescription field to given value.

### HasVideoDescription

`func (o *Video) HasVideoDescription() bool`

HasVideoDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *Video) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Video) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Video) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Video) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Video) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Video) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Video) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Video) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Video) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Video) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


