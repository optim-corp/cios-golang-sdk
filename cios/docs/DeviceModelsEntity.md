# DeviceModelsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Key** | **string** |  | 
**DeviceId** | **string** |  | 
**Model** | [**DeviceModelsEntityModel**](DeviceModelsEntity_model.md) |  | 
**SerialNumber** | Pointer to **string** |  | [optional] 
**StartAt** | Pointer to **NullableString** | ナノ秒 | [optional] 
**CustomInventory** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ResourceOwnerId** | **string** |  | 
**Watch** | Pointer to [**NullableWatch**](Watch.md) |  | [optional] 
**Components** | Pointer to [**NullableDeviceEntitiesComponent**](DeviceEntitiesComponent.md) |  | [optional] 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewDeviceModelsEntity

`func NewDeviceModelsEntity(id string, key string, deviceId string, model DeviceModelsEntityModel, resourceOwnerId string, createdAt string, updatedAt string, ) *DeviceModelsEntity`

NewDeviceModelsEntity instantiates a new DeviceModelsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceModelsEntityWithDefaults

`func NewDeviceModelsEntityWithDefaults() *DeviceModelsEntity`

NewDeviceModelsEntityWithDefaults instantiates a new DeviceModelsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceModelsEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceModelsEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceModelsEntity) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *DeviceModelsEntity) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DeviceModelsEntity) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DeviceModelsEntity) SetKey(v string)`

SetKey sets Key field to given value.


### GetDeviceId

`func (o *DeviceModelsEntity) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceModelsEntity) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceModelsEntity) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetModel

`func (o *DeviceModelsEntity) GetModel() DeviceModelsEntityModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceModelsEntity) GetModelOk() (*DeviceModelsEntityModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceModelsEntity) SetModel(v DeviceModelsEntityModel)`

SetModel sets Model field to given value.


### GetSerialNumber

`func (o *DeviceModelsEntity) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DeviceModelsEntity) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DeviceModelsEntity) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *DeviceModelsEntity) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetStartAt

`func (o *DeviceModelsEntity) GetStartAt() string`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *DeviceModelsEntity) GetStartAtOk() (*string, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *DeviceModelsEntity) SetStartAt(v string)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *DeviceModelsEntity) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### SetStartAtNil

`func (o *DeviceModelsEntity) SetStartAtNil(b bool)`

 SetStartAtNil sets the value for StartAt to be an explicit nil

### UnsetStartAt
`func (o *DeviceModelsEntity) UnsetStartAt()`

UnsetStartAt ensures that no value is present for StartAt, not even an explicit nil
### GetCustomInventory

`func (o *DeviceModelsEntity) GetCustomInventory() []map[string]interface{}`

GetCustomInventory returns the CustomInventory field if non-nil, zero value otherwise.

### GetCustomInventoryOk

`func (o *DeviceModelsEntity) GetCustomInventoryOk() (*[]map[string]interface{}, bool)`

GetCustomInventoryOk returns a tuple with the CustomInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomInventory

`func (o *DeviceModelsEntity) SetCustomInventory(v []map[string]interface{})`

SetCustomInventory sets CustomInventory field to given value.

### HasCustomInventory

`func (o *DeviceModelsEntity) HasCustomInventory() bool`

HasCustomInventory returns a boolean if a field has been set.

### GetResourceOwnerId

`func (o *DeviceModelsEntity) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *DeviceModelsEntity) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *DeviceModelsEntity) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.


### GetWatch

`func (o *DeviceModelsEntity) GetWatch() Watch`

GetWatch returns the Watch field if non-nil, zero value otherwise.

### GetWatchOk

`func (o *DeviceModelsEntity) GetWatchOk() (*Watch, bool)`

GetWatchOk returns a tuple with the Watch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatch

`func (o *DeviceModelsEntity) SetWatch(v Watch)`

SetWatch sets Watch field to given value.

### HasWatch

`func (o *DeviceModelsEntity) HasWatch() bool`

HasWatch returns a boolean if a field has been set.

### SetWatchNil

`func (o *DeviceModelsEntity) SetWatchNil(b bool)`

 SetWatchNil sets the value for Watch to be an explicit nil

### UnsetWatch
`func (o *DeviceModelsEntity) UnsetWatch()`

UnsetWatch ensures that no value is present for Watch, not even an explicit nil
### GetComponents

`func (o *DeviceModelsEntity) GetComponents() DeviceEntitiesComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *DeviceModelsEntity) GetComponentsOk() (*DeviceEntitiesComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *DeviceModelsEntity) SetComponents(v DeviceEntitiesComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *DeviceModelsEntity) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### SetComponentsNil

`func (o *DeviceModelsEntity) SetComponentsNil(b bool)`

 SetComponentsNil sets the value for Components to be an explicit nil

### UnsetComponents
`func (o *DeviceModelsEntity) UnsetComponents()`

UnsetComponents ensures that no value is present for Components, not even an explicit nil
### GetCreatedAt

`func (o *DeviceModelsEntity) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeviceModelsEntity) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeviceModelsEntity) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DeviceModelsEntity) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeviceModelsEntity) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeviceModelsEntity) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


