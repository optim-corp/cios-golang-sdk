# MultipleDataStoreChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** |  | 
**Channels** | [**[]DataStoreChannel**](DataStoreChannel.md) |  | 

## Methods

### NewMultipleDataStoreChannel

`func NewMultipleDataStoreChannel(total int64, channels []DataStoreChannel, ) *MultipleDataStoreChannel`

NewMultipleDataStoreChannel instantiates a new MultipleDataStoreChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleDataStoreChannelWithDefaults

`func NewMultipleDataStoreChannelWithDefaults() *MultipleDataStoreChannel`

NewMultipleDataStoreChannelWithDefaults instantiates a new MultipleDataStoreChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleDataStoreChannel) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleDataStoreChannel) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleDataStoreChannel) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetChannels

`func (o *MultipleDataStoreChannel) GetChannels() []DataStoreChannel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *MultipleDataStoreChannel) GetChannelsOk() (*[]DataStoreChannel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *MultipleDataStoreChannel) SetChannels(v []DataStoreChannel)`

SetChannels sets Channels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


