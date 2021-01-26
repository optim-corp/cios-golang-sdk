# DevicePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyId** | **string** |  | 
**ActivationKey** | Pointer to **string** |  | [optional] 
**ResourceOwnerId** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewDevicePolicy

`func NewDevicePolicy(policyId string, resourceOwnerId string, createdAt string, updatedAt string, ) *DevicePolicy`

NewDevicePolicy instantiates a new DevicePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicePolicyWithDefaults

`func NewDevicePolicyWithDefaults() *DevicePolicy`

NewDevicePolicyWithDefaults instantiates a new DevicePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyId

`func (o *DevicePolicy) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *DevicePolicy) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *DevicePolicy) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.


### GetActivationKey

`func (o *DevicePolicy) GetActivationKey() string`

GetActivationKey returns the ActivationKey field if non-nil, zero value otherwise.

### GetActivationKeyOk

`func (o *DevicePolicy) GetActivationKeyOk() (*string, bool)`

GetActivationKeyOk returns a tuple with the ActivationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationKey

`func (o *DevicePolicy) SetActivationKey(v string)`

SetActivationKey sets ActivationKey field to given value.

### HasActivationKey

`func (o *DevicePolicy) HasActivationKey() bool`

HasActivationKey returns a boolean if a field has been set.

### GetResourceOwnerId

`func (o *DevicePolicy) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *DevicePolicy) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *DevicePolicy) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.


### GetCreatedAt

`func (o *DevicePolicy) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DevicePolicy) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DevicePolicy) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DevicePolicy) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DevicePolicy) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DevicePolicy) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


