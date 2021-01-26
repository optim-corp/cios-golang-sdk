# MultipleDataStoreObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** |  | 
**Objects** | [**[]DataStoreObject**](DataStoreObject.md) |  | 

## Methods

### NewMultipleDataStoreObject

`func NewMultipleDataStoreObject(total int64, objects []DataStoreObject, ) *MultipleDataStoreObject`

NewMultipleDataStoreObject instantiates a new MultipleDataStoreObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleDataStoreObjectWithDefaults

`func NewMultipleDataStoreObjectWithDefaults() *MultipleDataStoreObject`

NewMultipleDataStoreObjectWithDefaults instantiates a new MultipleDataStoreObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleDataStoreObject) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleDataStoreObject) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleDataStoreObject) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetObjects

`func (o *MultipleDataStoreObject) GetObjects() []DataStoreObject`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *MultipleDataStoreObject) GetObjectsOk() (*[]DataStoreObject, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *MultipleDataStoreObject) SetObjects(v []DataStoreObject)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


