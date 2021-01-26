# DeviceEntitiesComponentUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Attribute** | Pointer to **map[string]interface{}** |  | [optional] 
**DisplayInfo** | Pointer to [**DisplayInfo**](DisplayInfo.md) |  | [optional] 

## Methods

### NewDeviceEntitiesComponentUpdateRequest

`func NewDeviceEntitiesComponentUpdateRequest() *DeviceEntitiesComponentUpdateRequest`

NewDeviceEntitiesComponentUpdateRequest instantiates a new DeviceEntitiesComponentUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceEntitiesComponentUpdateRequestWithDefaults

`func NewDeviceEntitiesComponentUpdateRequestWithDefaults() *DeviceEntitiesComponentUpdateRequest`

NewDeviceEntitiesComponentUpdateRequestWithDefaults instantiates a new DeviceEntitiesComponentUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeviceEntitiesComponentUpdateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceEntitiesComponentUpdateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceEntitiesComponentUpdateRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeviceEntitiesComponentUpdateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetParentId

`func (o *DeviceEntitiesComponentUpdateRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *DeviceEntitiesComponentUpdateRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *DeviceEntitiesComponentUpdateRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *DeviceEntitiesComponentUpdateRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetAttribute

`func (o *DeviceEntitiesComponentUpdateRequest) GetAttribute() map[string]interface{}`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *DeviceEntitiesComponentUpdateRequest) GetAttributeOk() (*map[string]interface{}, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *DeviceEntitiesComponentUpdateRequest) SetAttribute(v map[string]interface{})`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *DeviceEntitiesComponentUpdateRequest) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetDisplayInfo

`func (o *DeviceEntitiesComponentUpdateRequest) GetDisplayInfo() DisplayInfo`

GetDisplayInfo returns the DisplayInfo field if non-nil, zero value otherwise.

### GetDisplayInfoOk

`func (o *DeviceEntitiesComponentUpdateRequest) GetDisplayInfoOk() (*DisplayInfo, bool)`

GetDisplayInfoOk returns a tuple with the DisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInfo

`func (o *DeviceEntitiesComponentUpdateRequest) SetDisplayInfo(v DisplayInfo)`

SetDisplayInfo sets DisplayInfo field to given value.

### HasDisplayInfo

`func (o *DeviceEntitiesComponentUpdateRequest) HasDisplayInfo() bool`

HasDisplayInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


