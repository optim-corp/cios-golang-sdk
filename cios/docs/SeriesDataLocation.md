# SeriesDataLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Point固定 | 
**Coordinates** | **[]float32** |  | 

## Methods

### NewSeriesDataLocation

`func NewSeriesDataLocation(type_ string, coordinates []float32, ) *SeriesDataLocation`

NewSeriesDataLocation instantiates a new SeriesDataLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesDataLocationWithDefaults

`func NewSeriesDataLocationWithDefaults() *SeriesDataLocation`

NewSeriesDataLocationWithDefaults instantiates a new SeriesDataLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SeriesDataLocation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SeriesDataLocation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SeriesDataLocation) SetType(v string)`

SetType sets Type field to given value.


### GetCoordinates

`func (o *SeriesDataLocation) GetCoordinates() []float32`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *SeriesDataLocation) GetCoordinatesOk() (*[]float32, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *SeriesDataLocation) SetCoordinates(v []float32)`

SetCoordinates sets Coordinates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


