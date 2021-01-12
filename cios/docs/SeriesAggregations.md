# SeriesAggregations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Intervals** | **[]int64** | intervalの開始時刻の配列 | 
**TotalCounts** | **[]int64** | intervalごとの時系列データ件数の配列 | 
**Aggregations** | Pointer to **[]interface{}** | 集計結果の配列。リクエストで指定された集計対象・集計方法の順で並んでいる | [optional] 

## Methods

### NewSeriesAggregations

`func NewSeriesAggregations(intervals []int64, totalCounts []int64, ) *SeriesAggregations`

NewSeriesAggregations instantiates a new SeriesAggregations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesAggregationsWithDefaults

`func NewSeriesAggregationsWithDefaults() *SeriesAggregations`

NewSeriesAggregationsWithDefaults instantiates a new SeriesAggregations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntervals

`func (o *SeriesAggregations) GetIntervals() []int64`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *SeriesAggregations) GetIntervalsOk() (*[]int64, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *SeriesAggregations) SetIntervals(v []int64)`

SetIntervals sets Intervals field to given value.


### GetTotalCounts

`func (o *SeriesAggregations) GetTotalCounts() []int64`

GetTotalCounts returns the TotalCounts field if non-nil, zero value otherwise.

### GetTotalCountsOk

`func (o *SeriesAggregations) GetTotalCountsOk() (*[]int64, bool)`

GetTotalCountsOk returns a tuple with the TotalCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCounts

`func (o *SeriesAggregations) SetTotalCounts(v []int64)`

SetTotalCounts sets TotalCounts field to given value.


### GetAggregations

`func (o *SeriesAggregations) GetAggregations() []interface{}`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *SeriesAggregations) GetAggregationsOk() (*[]interface{}, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *SeriesAggregations) SetAggregations(v []interface{})`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *SeriesAggregations) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


