# SeriesDataLocationUnix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**NullableSeriesDataLocation**](SeriesDataLocation.md) |  | [optional] 
**Measurements** | Pointer to **map[string]map[string]interface{}** | 計測値。送信可能なプロパティはcollectionによって異なる。 collectionごとの定義は[Available Collections](./docs/available-collections.html)を参照。  以下の場合はリクエスト不正となる。 - collectionに定義されていないプロパティがmeasurementsに含まれる場合 - measurementsのJSON Objectのサイズが350KB以上の場合 | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 

## Methods

### NewSeriesDataLocationUnix

`func NewSeriesDataLocationUnix() *SeriesDataLocationUnix`

NewSeriesDataLocationUnix instantiates a new SeriesDataLocationUnix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesDataLocationUnixWithDefaults

`func NewSeriesDataLocationUnixWithDefaults() *SeriesDataLocationUnix`

NewSeriesDataLocationUnixWithDefaults instantiates a new SeriesDataLocationUnix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *SeriesDataLocationUnix) GetLocation() SeriesDataLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SeriesDataLocationUnix) GetLocationOk() (*SeriesDataLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SeriesDataLocationUnix) SetLocation(v SeriesDataLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SeriesDataLocationUnix) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *SeriesDataLocationUnix) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *SeriesDataLocationUnix) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetMeasurements

`func (o *SeriesDataLocationUnix) GetMeasurements() map[string]map[string]interface{}`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *SeriesDataLocationUnix) GetMeasurementsOk() (*map[string]map[string]interface{}, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *SeriesDataLocationUnix) SetMeasurements(v map[string]map[string]interface{})`

SetMeasurements sets Measurements field to given value.

### HasMeasurements

`func (o *SeriesDataLocationUnix) HasMeasurements() bool`

HasMeasurements returns a boolean if a field has been set.

### GetTimestamp

`func (o *SeriesDataLocationUnix) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SeriesDataLocationUnix) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SeriesDataLocationUnix) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SeriesDataLocationUnix) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


