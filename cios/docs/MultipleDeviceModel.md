# MultipleDeviceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** |  | 
**Models** | [**[]DeviceModel**](DeviceModel.md) |  | 

## Methods

### NewMultipleDeviceModel

`func NewMultipleDeviceModel(total int64, models []DeviceModel, ) *MultipleDeviceModel`

NewMultipleDeviceModel instantiates a new MultipleDeviceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleDeviceModelWithDefaults

`func NewMultipleDeviceModelWithDefaults() *MultipleDeviceModel`

NewMultipleDeviceModelWithDefaults instantiates a new MultipleDeviceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleDeviceModel) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleDeviceModel) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleDeviceModel) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetModels

`func (o *MultipleDeviceModel) GetModels() []DeviceModel`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *MultipleDeviceModel) GetModelsOk() (*[]DeviceModel, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *MultipleDeviceModel) SetModels(v []DeviceModel)`

SetModels sets Models field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


