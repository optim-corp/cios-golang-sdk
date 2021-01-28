# MultipleChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** |  | 
**Channels** | [**[]Channel**](Channel.md) |  | 

## Methods

### NewMultipleChannel

`func NewMultipleChannel(total int64, channels []Channel, ) *MultipleChannel`

NewMultipleChannel instantiates a new MultipleChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleChannelWithDefaults

`func NewMultipleChannelWithDefaults() *MultipleChannel`

NewMultipleChannelWithDefaults instantiates a new MultipleChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleChannel) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleChannel) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleChannel) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetChannels

`func (o *MultipleChannel) GetChannels() []Channel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *MultipleChannel) GetChannelsOk() (*[]Channel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *MultipleChannel) SetChannels(v []Channel)`

SetChannels sets Channels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


