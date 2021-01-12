# MultipleDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** |  | 
**Devices** | [**[]Device**](Device.md) |  | 

## Methods

### NewMultipleDevice

`func NewMultipleDevice(total int64, devices []Device, ) *MultipleDevice`

NewMultipleDevice instantiates a new MultipleDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleDeviceWithDefaults

`func NewMultipleDeviceWithDefaults() *MultipleDevice`

NewMultipleDeviceWithDefaults instantiates a new MultipleDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleDevice) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleDevice) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleDevice) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetDevices

`func (o *MultipleDevice) GetDevices() []Device`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *MultipleDevice) GetDevicesOk() (*[]Device, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *MultipleDevice) SetDevices(v []Device)`

SetDevices sets Devices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


