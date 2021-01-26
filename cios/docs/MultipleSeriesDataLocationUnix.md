# MultipleSeriesDataLocationUnix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** | 検索条件にマッチする全結果数 | 
**Data** | [**[]SeriesDataLocationUnix**](SeriesDataLocationUnix.md) | 時系列データ配列。データが0件の場合は空配列。 | 

## Methods

### NewMultipleSeriesDataLocationUnix

`func NewMultipleSeriesDataLocationUnix(total int64, data []SeriesDataLocationUnix, ) *MultipleSeriesDataLocationUnix`

NewMultipleSeriesDataLocationUnix instantiates a new MultipleSeriesDataLocationUnix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleSeriesDataLocationUnixWithDefaults

`func NewMultipleSeriesDataLocationUnixWithDefaults() *MultipleSeriesDataLocationUnix`

NewMultipleSeriesDataLocationUnixWithDefaults instantiates a new MultipleSeriesDataLocationUnix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleSeriesDataLocationUnix) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleSeriesDataLocationUnix) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleSeriesDataLocationUnix) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetData

`func (o *MultipleSeriesDataLocationUnix) GetData() []SeriesDataLocationUnix`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MultipleSeriesDataLocationUnix) GetDataOk() (*[]SeriesDataLocationUnix, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MultipleSeriesDataLocationUnix) SetData(v []SeriesDataLocationUnix)`

SetData sets Data field to given value.


### SetDataNil

`func (o *MultipleSeriesDataLocationUnix) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *MultipleSeriesDataLocationUnix) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


