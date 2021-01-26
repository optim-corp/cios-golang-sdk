# MultipleDeviceModelEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** |  | 
**Entities** | [**[]DeviceModelsEntity**](DeviceModelsEntity.md) |  | 

## Methods

### NewMultipleDeviceModelEntity

`func NewMultipleDeviceModelEntity(total int64, entities []DeviceModelsEntity, ) *MultipleDeviceModelEntity`

NewMultipleDeviceModelEntity instantiates a new MultipleDeviceModelEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleDeviceModelEntityWithDefaults

`func NewMultipleDeviceModelEntityWithDefaults() *MultipleDeviceModelEntity`

NewMultipleDeviceModelEntityWithDefaults instantiates a new MultipleDeviceModelEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleDeviceModelEntity) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleDeviceModelEntity) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleDeviceModelEntity) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetEntities

`func (o *MultipleDeviceModelEntity) GetEntities() []DeviceModelsEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *MultipleDeviceModelEntity) GetEntitiesOk() (*[]DeviceModelsEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *MultipleDeviceModelEntity) SetEntities(v []DeviceModelsEntity)`

SetEntities sets Entities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


