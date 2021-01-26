# MultiplePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** |  | 
**Points** | [**[]Point**](Point.md) |  | 

## Methods

### NewMultiplePoint

`func NewMultiplePoint(total int64, points []Point, ) *MultiplePoint`

NewMultiplePoint instantiates a new MultiplePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiplePointWithDefaults

`func NewMultiplePointWithDefaults() *MultiplePoint`

NewMultiplePointWithDefaults instantiates a new MultiplePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultiplePoint) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultiplePoint) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultiplePoint) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetPoints

`func (o *MultiplePoint) GetPoints() []Point`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *MultiplePoint) GetPointsOk() (*[]Point, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *MultiplePoint) SetPoints(v []Point)`

SetPoints sets Points field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


