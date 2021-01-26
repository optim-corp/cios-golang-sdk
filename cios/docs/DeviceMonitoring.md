# DeviceMonitoring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **string** |  | 
**Timestamp** | Pointer to **string** |  | [optional] 
**ConnectionStatus** | Pointer to **string** |  | [optional] 
**MonitoringStatus** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Perfomance** | Pointer to [**DevicePerformance**](DevicePerformance.md) |  | [optional] 
**ConnectedDeviceResults** | Pointer to [**[]ConnectedDeviceResult**](ConnectedDeviceResult.md) |  | [optional] 

## Methods

### NewDeviceMonitoring

`func NewDeviceMonitoring(deviceId string, ) *DeviceMonitoring`

NewDeviceMonitoring instantiates a new DeviceMonitoring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceMonitoringWithDefaults

`func NewDeviceMonitoringWithDefaults() *DeviceMonitoring`

NewDeviceMonitoringWithDefaults instantiates a new DeviceMonitoring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *DeviceMonitoring) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceMonitoring) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceMonitoring) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetTimestamp

`func (o *DeviceMonitoring) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DeviceMonitoring) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DeviceMonitoring) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DeviceMonitoring) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *DeviceMonitoring) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *DeviceMonitoring) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *DeviceMonitoring) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *DeviceMonitoring) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetMonitoringStatus

`func (o *DeviceMonitoring) GetMonitoringStatus() string`

GetMonitoringStatus returns the MonitoringStatus field if non-nil, zero value otherwise.

### GetMonitoringStatusOk

`func (o *DeviceMonitoring) GetMonitoringStatusOk() (*string, bool)`

GetMonitoringStatusOk returns a tuple with the MonitoringStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringStatus

`func (o *DeviceMonitoring) SetMonitoringStatus(v string)`

SetMonitoringStatus sets MonitoringStatus field to given value.

### HasMonitoringStatus

`func (o *DeviceMonitoring) HasMonitoringStatus() bool`

HasMonitoringStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *DeviceMonitoring) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *DeviceMonitoring) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *DeviceMonitoring) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *DeviceMonitoring) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetPerfomance

`func (o *DeviceMonitoring) GetPerfomance() DevicePerformance`

GetPerfomance returns the Perfomance field if non-nil, zero value otherwise.

### GetPerfomanceOk

`func (o *DeviceMonitoring) GetPerfomanceOk() (*DevicePerformance, bool)`

GetPerfomanceOk returns a tuple with the Perfomance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfomance

`func (o *DeviceMonitoring) SetPerfomance(v DevicePerformance)`

SetPerfomance sets Perfomance field to given value.

### HasPerfomance

`func (o *DeviceMonitoring) HasPerfomance() bool`

HasPerfomance returns a boolean if a field has been set.

### GetConnectedDeviceResults

`func (o *DeviceMonitoring) GetConnectedDeviceResults() []ConnectedDeviceResult`

GetConnectedDeviceResults returns the ConnectedDeviceResults field if non-nil, zero value otherwise.

### GetConnectedDeviceResultsOk

`func (o *DeviceMonitoring) GetConnectedDeviceResultsOk() (*[]ConnectedDeviceResult, bool)`

GetConnectedDeviceResultsOk returns a tuple with the ConnectedDeviceResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedDeviceResults

`func (o *DeviceMonitoring) SetConnectedDeviceResults(v []ConnectedDeviceResult)`

SetConnectedDeviceResults sets ConnectedDeviceResults field to given value.

### HasConnectedDeviceResults

`func (o *DeviceMonitoring) HasConnectedDeviceResults() bool`

HasConnectedDeviceResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


