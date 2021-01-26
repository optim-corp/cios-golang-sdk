# SeriesDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**NullableSeriesDataLocation**](SeriesDataLocation.md) |  | [optional] 
**Measurements** | Pointer to **map[string]map[string]interface{}** | 計測値。送信可能なプロパティはcollectionによって異なる。 collectionごとの定義は[Available Collections](./docs/available-collections.html)を参照。  以下の場合はリクエスト不正となる。 - collectionに定義されていないプロパティがmeasurementsに含まれる場合 - measurementsのJSON Objectのサイズが350KB以上の場合 | [optional] 

## Methods

### NewSeriesDataRequest

`func NewSeriesDataRequest() *SeriesDataRequest`

NewSeriesDataRequest instantiates a new SeriesDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesDataRequestWithDefaults

`func NewSeriesDataRequestWithDefaults() *SeriesDataRequest`

NewSeriesDataRequestWithDefaults instantiates a new SeriesDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *SeriesDataRequest) GetLocation() SeriesDataLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SeriesDataRequest) GetLocationOk() (*SeriesDataLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SeriesDataRequest) SetLocation(v SeriesDataLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SeriesDataRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *SeriesDataRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *SeriesDataRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetMeasurements

`func (o *SeriesDataRequest) GetMeasurements() map[string]map[string]interface{}`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *SeriesDataRequest) GetMeasurementsOk() (*map[string]map[string]interface{}, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *SeriesDataRequest) SetMeasurements(v map[string]map[string]interface{})`

SetMeasurements sets Measurements field to given value.

### HasMeasurements

`func (o *SeriesDataRequest) HasMeasurements() bool`

HasMeasurements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


