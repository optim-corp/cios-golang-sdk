# DataStoreChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DataExist** | Pointer to **bool** |  | [optional] 
**Stats** | Pointer to [**DataStoreChannelStats**](DataStoreChannel_stats.md) |  | [optional] 

## Methods

### NewDataStoreChannel

`func NewDataStoreChannel() *DataStoreChannel`

NewDataStoreChannel instantiates a new DataStoreChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStoreChannelWithDefaults

`func NewDataStoreChannelWithDefaults() *DataStoreChannel`

NewDataStoreChannelWithDefaults instantiates a new DataStoreChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataStoreChannel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataStoreChannel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataStoreChannel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataStoreChannel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDataExist

`func (o *DataStoreChannel) GetDataExist() bool`

GetDataExist returns the DataExist field if non-nil, zero value otherwise.

### GetDataExistOk

`func (o *DataStoreChannel) GetDataExistOk() (*bool, bool)`

GetDataExistOk returns a tuple with the DataExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataExist

`func (o *DataStoreChannel) SetDataExist(v bool)`

SetDataExist sets DataExist field to given value.

### HasDataExist

`func (o *DataStoreChannel) HasDataExist() bool`

HasDataExist returns a boolean if a field has been set.

### GetStats

`func (o *DataStoreChannel) GetStats() DataStoreChannelStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *DataStoreChannel) GetStatsOk() (*DataStoreChannelStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *DataStoreChannel) SetStats(v DataStoreChannelStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *DataStoreChannel) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


