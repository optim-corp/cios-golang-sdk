# SeriesAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | **string** | リクエストボディで指定された集計対象項目 | 
**Method** | **string** | リクエストボディで指定された集計方法 | 
**Counts** | **[]int64** | intervalごとの集計対象項目が有効だったデータ件数の配列 | 
**Values** | [**[]AnyOfintegernumberstringbooleanarray**](AnyOfintegernumberstringbooleanarray.md) | intervalごとの集計結果の配列 | 

## Methods

### NewSeriesAggregation

`func NewSeriesAggregation(target string, method string, counts []int64, values []AnyOfintegernumberstringbooleanarray, ) *SeriesAggregation`

NewSeriesAggregation instantiates a new SeriesAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesAggregationWithDefaults

`func NewSeriesAggregationWithDefaults() *SeriesAggregation`

NewSeriesAggregationWithDefaults instantiates a new SeriesAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *SeriesAggregation) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SeriesAggregation) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SeriesAggregation) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetMethod

`func (o *SeriesAggregation) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SeriesAggregation) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SeriesAggregation) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetCounts

`func (o *SeriesAggregation) GetCounts() []int64`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *SeriesAggregation) GetCountsOk() (*[]int64, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *SeriesAggregation) SetCounts(v []int64)`

SetCounts sets Counts field to given value.


### GetValues

`func (o *SeriesAggregation) GetValues() []AnyOfintegernumberstringbooleanarray`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SeriesAggregation) GetValuesOk() (*[]AnyOfintegernumberstringbooleanarray, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SeriesAggregation) SetValues(v []AnyOfintegernumberstringbooleanarray)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


