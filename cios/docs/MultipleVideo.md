# MultipleVideo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** |  | 
**Videos** | [**[]Video**](Video.md) |  | 

## Methods

### NewMultipleVideo

`func NewMultipleVideo(total int64, videos []Video, ) *MultipleVideo`

NewMultipleVideo instantiates a new MultipleVideo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleVideoWithDefaults

`func NewMultipleVideoWithDefaults() *MultipleVideo`

NewMultipleVideoWithDefaults instantiates a new MultipleVideo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleVideo) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleVideo) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleVideo) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetVideos

`func (o *MultipleVideo) GetVideos() []Video`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *MultipleVideo) GetVideosOk() (*[]Video, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *MultipleVideo) SetVideos(v []Video)`

SetVideos sets Videos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


