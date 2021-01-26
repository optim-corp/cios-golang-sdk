# MultipleBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** |  | 
**Buckets** | [**[]Bucket**](Bucket.md) |  | 

## Methods

### NewMultipleBucket

`func NewMultipleBucket(total int64, buckets []Bucket, ) *MultipleBucket`

NewMultipleBucket instantiates a new MultipleBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleBucketWithDefaults

`func NewMultipleBucketWithDefaults() *MultipleBucket`

NewMultipleBucketWithDefaults instantiates a new MultipleBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleBucket) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleBucket) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleBucket) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetBuckets

`func (o *MultipleBucket) GetBuckets() []Bucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *MultipleBucket) GetBucketsOk() (*[]Bucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *MultipleBucket) SetBuckets(v []Bucket)`

SetBuckets sets Buckets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


