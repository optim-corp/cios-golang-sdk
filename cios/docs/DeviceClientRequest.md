# DeviceClientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIdList** | Pointer to [**[]DeviceClientRsaKey**](DeviceClientRsaKey.md) |  | [optional] 

## Methods

### NewDeviceClientRequest

`func NewDeviceClientRequest() *DeviceClientRequest`

NewDeviceClientRequest instantiates a new DeviceClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceClientRequestWithDefaults

`func NewDeviceClientRequestWithDefaults() *DeviceClientRequest`

NewDeviceClientRequestWithDefaults instantiates a new DeviceClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIdList

`func (o *DeviceClientRequest) GetClientIdList() []DeviceClientRsaKey`

GetClientIdList returns the ClientIdList field if non-nil, zero value otherwise.

### GetClientIdListOk

`func (o *DeviceClientRequest) GetClientIdListOk() (*[]DeviceClientRsaKey, bool)`

GetClientIdListOk returns a tuple with the ClientIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdList

`func (o *DeviceClientRequest) SetClientIdList(v []DeviceClientRsaKey)`

SetClientIdList sets ClientIdList field to given value.

### HasClientIdList

`func (o *DeviceClientRequest) HasClientIdList() bool`

HasClientIdList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


