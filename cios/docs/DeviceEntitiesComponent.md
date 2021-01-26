# DeviceEntitiesComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**ComponentTypeEnum**](ComponentTypeEnum.md) |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**Attribute** | Pointer to **map[string]interface{}** |  | [optional] 
**DisplayInfo** | Pointer to [**[]DisplayInfo**](DisplayInfo.md) |  | [optional] 
**Components** | Pointer to [**[]DeviceEntitiesComponent**](DeviceEntitiesComponent.md) |  | [optional] 

## Methods

### NewDeviceEntitiesComponent

`func NewDeviceEntitiesComponent(id string, type_ ComponentTypeEnum, ) *DeviceEntitiesComponent`

NewDeviceEntitiesComponent instantiates a new DeviceEntitiesComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceEntitiesComponentWithDefaults

`func NewDeviceEntitiesComponentWithDefaults() *DeviceEntitiesComponent`

NewDeviceEntitiesComponentWithDefaults instantiates a new DeviceEntitiesComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceEntitiesComponent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceEntitiesComponent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceEntitiesComponent) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *DeviceEntitiesComponent) GetType() ComponentTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceEntitiesComponent) GetTypeOk() (*ComponentTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceEntitiesComponent) SetType(v ComponentTypeEnum)`

SetType sets Type field to given value.


### GetParentId

`func (o *DeviceEntitiesComponent) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *DeviceEntitiesComponent) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *DeviceEntitiesComponent) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *DeviceEntitiesComponent) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetAttribute

`func (o *DeviceEntitiesComponent) GetAttribute() map[string]interface{}`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *DeviceEntitiesComponent) GetAttributeOk() (*map[string]interface{}, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *DeviceEntitiesComponent) SetAttribute(v map[string]interface{})`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *DeviceEntitiesComponent) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetDisplayInfo

`func (o *DeviceEntitiesComponent) GetDisplayInfo() []DisplayInfo`

GetDisplayInfo returns the DisplayInfo field if non-nil, zero value otherwise.

### GetDisplayInfoOk

`func (o *DeviceEntitiesComponent) GetDisplayInfoOk() (*[]DisplayInfo, bool)`

GetDisplayInfoOk returns a tuple with the DisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInfo

`func (o *DeviceEntitiesComponent) SetDisplayInfo(v []DisplayInfo)`

SetDisplayInfo sets DisplayInfo field to given value.

### HasDisplayInfo

`func (o *DeviceEntitiesComponent) HasDisplayInfo() bool`

HasDisplayInfo returns a boolean if a field has been set.

### GetComponents

`func (o *DeviceEntitiesComponent) GetComponents() []DeviceEntitiesComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *DeviceEntitiesComponent) GetComponentsOk() (*[]DeviceEntitiesComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *DeviceEntitiesComponent) SetComponents(v []DeviceEntitiesComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *DeviceEntitiesComponent) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


