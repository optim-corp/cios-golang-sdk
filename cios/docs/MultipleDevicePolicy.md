# MultipleDevicePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** |  | 
**Policies** | [**[]DevicePolicy**](DevicePolicy.md) |  | 

## Methods

### NewMultipleDevicePolicy

`func NewMultipleDevicePolicy(total int64, policies []DevicePolicy, ) *MultipleDevicePolicy`

NewMultipleDevicePolicy instantiates a new MultipleDevicePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleDevicePolicyWithDefaults

`func NewMultipleDevicePolicyWithDefaults() *MultipleDevicePolicy`

NewMultipleDevicePolicyWithDefaults instantiates a new MultipleDevicePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleDevicePolicy) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleDevicePolicy) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleDevicePolicy) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetPolicies

`func (o *MultipleDevicePolicy) GetPolicies() []DevicePolicy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *MultipleDevicePolicy) GetPoliciesOk() (*[]DevicePolicy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *MultipleDevicePolicy) SetPolicies(v []DevicePolicy)`

SetPolicies sets Policies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


