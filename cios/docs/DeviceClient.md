# DeviceClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**RsaPublickey** | Pointer to **string** | PEM | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceClient

`func NewDeviceClient() *DeviceClient`

NewDeviceClient instantiates a new DeviceClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceClientWithDefaults

`func NewDeviceClientWithDefaults() *DeviceClient`

NewDeviceClientWithDefaults instantiates a new DeviceClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *DeviceClient) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *DeviceClient) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *DeviceClient) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *DeviceClient) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *DeviceClient) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *DeviceClient) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *DeviceClient) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *DeviceClient) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetRsaPublickey

`func (o *DeviceClient) GetRsaPublickey() string`

GetRsaPublickey returns the RsaPublickey field if non-nil, zero value otherwise.

### GetRsaPublickeyOk

`func (o *DeviceClient) GetRsaPublickeyOk() (*string, bool)`

GetRsaPublickeyOk returns a tuple with the RsaPublickey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaPublickey

`func (o *DeviceClient) SetRsaPublickey(v string)`

SetRsaPublickey sets RsaPublickey field to given value.

### HasRsaPublickey

`func (o *DeviceClient) HasRsaPublickey() bool`

HasRsaPublickey returns a boolean if a field has been set.

### GetDisplayName

`func (o *DeviceClient) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DeviceClient) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DeviceClient) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DeviceClient) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


