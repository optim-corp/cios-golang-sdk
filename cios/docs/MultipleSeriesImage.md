# MultipleSeriesImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** | 検索条件にマッチする全結果数 | 
**Data** | [**[]SeriesImage**](SeriesImage.md) | 画像データとタイムスタンプの配列。データが0件の場合は空配列。 | 
**ImageMap** | **map[string]int32** | timestampがキーとなってdataのindexが格納されているhashmap | 

## Methods

### NewMultipleSeriesImage

`func NewMultipleSeriesImage(total int64, data []SeriesImage, imageMap map[string]int32, ) *MultipleSeriesImage`

NewMultipleSeriesImage instantiates a new MultipleSeriesImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleSeriesImageWithDefaults

`func NewMultipleSeriesImageWithDefaults() *MultipleSeriesImage`

NewMultipleSeriesImageWithDefaults instantiates a new MultipleSeriesImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleSeriesImage) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleSeriesImage) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleSeriesImage) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetData

`func (o *MultipleSeriesImage) GetData() []SeriesImage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MultipleSeriesImage) GetDataOk() (*[]SeriesImage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MultipleSeriesImage) SetData(v []SeriesImage)`

SetData sets Data field to given value.


### GetImageMap

`func (o *MultipleSeriesImage) GetImageMap() map[string]int32`

GetImageMap returns the ImageMap field if non-nil, zero value otherwise.

### GetImageMapOk

`func (o *MultipleSeriesImage) GetImageMapOk() (*map[string]int32, bool)`

GetImageMapOk returns a tuple with the ImageMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageMap

`func (o *MultipleSeriesImage) SetImageMap(v map[string]int32)`

SetImageMap sets ImageMap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


