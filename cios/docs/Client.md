# Client

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationClient** | Pointer to [**ApplicationClient**](ApplicationClient.md) |  | [optional] 
**DeviceClient** | Pointer to [**DeviceClient**](DeviceClient.md) |  | [optional] 

## Methods

### NewClient

`func NewClient() *Client`

NewClient instantiates a new Client object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientWithDefaults

`func NewClientWithDefaults() *Client`

NewClientWithDefaults instantiates a new Client object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationClient

`func (o *Client) GetApplicationClient() ApplicationClient`

GetApplicationClient returns the ApplicationClient field if non-nil, zero value otherwise.

### GetApplicationClientOk

`func (o *Client) GetApplicationClientOk() (*ApplicationClient, bool)`

GetApplicationClientOk returns a tuple with the ApplicationClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationClient

`func (o *Client) SetApplicationClient(v ApplicationClient)`

SetApplicationClient sets ApplicationClient field to given value.

### HasApplicationClient

`func (o *Client) HasApplicationClient() bool`

HasApplicationClient returns a boolean if a field has been set.

### GetDeviceClient

`func (o *Client) GetDeviceClient() DeviceClient`

GetDeviceClient returns the DeviceClient field if non-nil, zero value otherwise.

### GetDeviceClientOk

`func (o *Client) GetDeviceClientOk() (*DeviceClient, bool)`

GetDeviceClientOk returns a tuple with the DeviceClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClient

`func (o *Client) SetDeviceClient(v DeviceClient)`

SetDeviceClient sets DeviceClient field to given value.

### HasDeviceClient

`func (o *Client) HasDeviceClient() bool`

HasDeviceClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


