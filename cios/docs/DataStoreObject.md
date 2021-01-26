# DataStoreObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ChannelProtocols** | Pointer to [**[]ChannelProtocol**](ChannelProtocol.md) |  | [optional] 
**MimeType** | **string** |  | 
**Timestamp** | **string** |  | 
**Location** | Pointer to [**DataStoreObjectLocation**](DataStoreObject_location.md) |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 

## Methods

### NewDataStoreObject

`func NewDataStoreObject(id string, mimeType string, timestamp string, ) *DataStoreObject`

NewDataStoreObject instantiates a new DataStoreObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStoreObjectWithDefaults

`func NewDataStoreObjectWithDefaults() *DataStoreObject`

NewDataStoreObjectWithDefaults instantiates a new DataStoreObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataStoreObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataStoreObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataStoreObject) SetId(v string)`

SetId sets Id field to given value.


### GetChannelProtocols

`func (o *DataStoreObject) GetChannelProtocols() []ChannelProtocol`

GetChannelProtocols returns the ChannelProtocols field if non-nil, zero value otherwise.

### GetChannelProtocolsOk

`func (o *DataStoreObject) GetChannelProtocolsOk() (*[]ChannelProtocol, bool)`

GetChannelProtocolsOk returns a tuple with the ChannelProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProtocols

`func (o *DataStoreObject) SetChannelProtocols(v []ChannelProtocol)`

SetChannelProtocols sets ChannelProtocols field to given value.

### HasChannelProtocols

`func (o *DataStoreObject) HasChannelProtocols() bool`

HasChannelProtocols returns a boolean if a field has been set.

### GetMimeType

`func (o *DataStoreObject) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *DataStoreObject) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *DataStoreObject) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetTimestamp

`func (o *DataStoreObject) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DataStoreObject) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DataStoreObject) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetLocation

`func (o *DataStoreObject) GetLocation() DataStoreObjectLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DataStoreObject) GetLocationOk() (*DataStoreObjectLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DataStoreObject) SetLocation(v DataStoreObjectLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DataStoreObject) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetSessionId

`func (o *DataStoreObject) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *DataStoreObject) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *DataStoreObject) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *DataStoreObject) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


