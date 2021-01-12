# PackerFormatJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | [**PackerFormatJsonHeader**](PackerFormatJsonHeader.md) |  | 
**Payload** | **interface{}** |  | 

## Methods

### NewPackerFormatJson

`func NewPackerFormatJson(header PackerFormatJsonHeader, payload interface{}, ) *PackerFormatJson`

NewPackerFormatJson instantiates a new PackerFormatJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackerFormatJsonWithDefaults

`func NewPackerFormatJsonWithDefaults() *PackerFormatJson`

NewPackerFormatJsonWithDefaults instantiates a new PackerFormatJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *PackerFormatJson) GetHeader() PackerFormatJsonHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *PackerFormatJson) GetHeaderOk() (*PackerFormatJsonHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *PackerFormatJson) SetHeader(v PackerFormatJsonHeader)`

SetHeader sets Header field to given value.


### GetPayload

`func (o *PackerFormatJson) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *PackerFormatJson) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *PackerFormatJson) SetPayload(v interface{})`

SetPayload sets Payload field to given value.


### SetPayloadNil

`func (o *PackerFormatJson) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *PackerFormatJson) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


