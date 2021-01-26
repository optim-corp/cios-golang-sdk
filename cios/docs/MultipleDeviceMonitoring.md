# MultipleDeviceMonitoring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** |  | 
**Monitorings** | [**[]DeviceMonitoring**](DeviceMonitoring.md) |  | 

## Methods

### NewMultipleDeviceMonitoring

`func NewMultipleDeviceMonitoring(total int64, monitorings []DeviceMonitoring, ) *MultipleDeviceMonitoring`

NewMultipleDeviceMonitoring instantiates a new MultipleDeviceMonitoring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleDeviceMonitoringWithDefaults

`func NewMultipleDeviceMonitoringWithDefaults() *MultipleDeviceMonitoring`

NewMultipleDeviceMonitoringWithDefaults instantiates a new MultipleDeviceMonitoring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleDeviceMonitoring) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleDeviceMonitoring) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleDeviceMonitoring) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetMonitorings

`func (o *MultipleDeviceMonitoring) GetMonitorings() []DeviceMonitoring`

GetMonitorings returns the Monitorings field if non-nil, zero value otherwise.

### GetMonitoringsOk

`func (o *MultipleDeviceMonitoring) GetMonitoringsOk() (*[]DeviceMonitoring, bool)`

GetMonitoringsOk returns a tuple with the Monitorings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorings

`func (o *MultipleDeviceMonitoring) SetMonitorings(v []DeviceMonitoring)`

SetMonitorings sets Monitorings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


