# SeriesImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **string** |  | 
**Image** | **string** | base64エンコードされた画像データ | 
**ImageType** | **string** | 画像データのフォーマット | 

## Methods

### NewSeriesImage

`func NewSeriesImage(timestamp string, image string, imageType string, ) *SeriesImage`

NewSeriesImage instantiates a new SeriesImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesImageWithDefaults

`func NewSeriesImageWithDefaults() *SeriesImage`

NewSeriesImageWithDefaults instantiates a new SeriesImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *SeriesImage) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SeriesImage) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SeriesImage) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetImage

`func (o *SeriesImage) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *SeriesImage) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *SeriesImage) SetImage(v string)`

SetImage sets Image field to given value.


### GetImageType

`func (o *SeriesImage) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *SeriesImage) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *SeriesImage) SetImageType(v string)`

SetImageType sets ImageType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


