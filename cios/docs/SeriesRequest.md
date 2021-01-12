# SeriesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**CollectionLocation**](CollectionLocation.md) |  | [optional] 
**Measurements** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**SeriesId** | **string** |  | 
**Timestamp** | **int64** |  | 

## Methods

### NewSeriesRequest

`func NewSeriesRequest(seriesId string, timestamp int64, ) *SeriesRequest`

NewSeriesRequest instantiates a new SeriesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesRequestWithDefaults

`func NewSeriesRequestWithDefaults() *SeriesRequest`

NewSeriesRequestWithDefaults instantiates a new SeriesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *SeriesRequest) GetLocation() CollectionLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SeriesRequest) GetLocationOk() (*CollectionLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SeriesRequest) SetLocation(v CollectionLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SeriesRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMeasurements

`func (o *SeriesRequest) GetMeasurements() map[string]map[string]interface{}`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *SeriesRequest) GetMeasurementsOk() (*map[string]map[string]interface{}, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *SeriesRequest) SetMeasurements(v map[string]map[string]interface{})`

SetMeasurements sets Measurements field to given value.

### HasMeasurements

`func (o *SeriesRequest) HasMeasurements() bool`

HasMeasurements returns a boolean if a field has been set.

### GetSeriesId

`func (o *SeriesRequest) GetSeriesId() string`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *SeriesRequest) GetSeriesIdOk() (*string, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *SeriesRequest) SetSeriesId(v string)`

SetSeriesId sets SeriesId field to given value.


### GetTimestamp

`func (o *SeriesRequest) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SeriesRequest) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SeriesRequest) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


