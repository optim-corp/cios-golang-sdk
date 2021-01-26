# PackerFormatJsonHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | **string** |  | 
**ChannelProtocols** | Pointer to [**[]ChannelProtocol**](ChannelProtocol.md) |  | [optional] 
**Timestamp** | **string** |  | 
**SessionId** | Pointer to **string** |  | [optional] 
**MimeType** | **string** |  | 
**Labels** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPackerFormatJsonHeader

`func NewPackerFormatJsonHeader(channelId string, timestamp string, mimeType string, ) *PackerFormatJsonHeader`

NewPackerFormatJsonHeader instantiates a new PackerFormatJsonHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackerFormatJsonHeaderWithDefaults

`func NewPackerFormatJsonHeaderWithDefaults() *PackerFormatJsonHeader`

NewPackerFormatJsonHeaderWithDefaults instantiates a new PackerFormatJsonHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *PackerFormatJsonHeader) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *PackerFormatJsonHeader) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *PackerFormatJsonHeader) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetChannelProtocols

`func (o *PackerFormatJsonHeader) GetChannelProtocols() []ChannelProtocol`

GetChannelProtocols returns the ChannelProtocols field if non-nil, zero value otherwise.

### GetChannelProtocolsOk

`func (o *PackerFormatJsonHeader) GetChannelProtocolsOk() (*[]ChannelProtocol, bool)`

GetChannelProtocolsOk returns a tuple with the ChannelProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProtocols

`func (o *PackerFormatJsonHeader) SetChannelProtocols(v []ChannelProtocol)`

SetChannelProtocols sets ChannelProtocols field to given value.

### HasChannelProtocols

`func (o *PackerFormatJsonHeader) HasChannelProtocols() bool`

HasChannelProtocols returns a boolean if a field has been set.

### GetTimestamp

`func (o *PackerFormatJsonHeader) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PackerFormatJsonHeader) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PackerFormatJsonHeader) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetSessionId

`func (o *PackerFormatJsonHeader) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *PackerFormatJsonHeader) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *PackerFormatJsonHeader) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *PackerFormatJsonHeader) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetMimeType

`func (o *PackerFormatJsonHeader) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *PackerFormatJsonHeader) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *PackerFormatJsonHeader) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetLabels

`func (o *PackerFormatJsonHeader) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PackerFormatJsonHeader) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PackerFormatJsonHeader) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PackerFormatJsonHeader) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


