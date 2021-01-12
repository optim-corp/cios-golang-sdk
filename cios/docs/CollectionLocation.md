# CollectionLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "Point"]
**Coordinates** | **[]float32** |  | 

## Methods

### NewCollectionLocation

`func NewCollectionLocation(type_ string, coordinates []float32, ) *CollectionLocation`

NewCollectionLocation instantiates a new CollectionLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionLocationWithDefaults

`func NewCollectionLocationWithDefaults() *CollectionLocation`

NewCollectionLocationWithDefaults instantiates a new CollectionLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CollectionLocation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CollectionLocation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CollectionLocation) SetType(v string)`

SetType sets Type field to given value.


### GetCoordinates

`func (o *CollectionLocation) GetCoordinates() []float32`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *CollectionLocation) GetCoordinatesOk() (*[]float32, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *CollectionLocation) SetCoordinates(v []float32)`

SetCoordinates sets Coordinates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


