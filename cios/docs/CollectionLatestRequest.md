# CollectionLatestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceOwnerId** | **string** |  | 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Projection** | Pointer to **string** |  | [optional] [default to "data"]

## Methods

### NewCollectionLatestRequest

`func NewCollectionLatestRequest(resourceOwnerId string, ) *CollectionLatestRequest`

NewCollectionLatestRequest instantiates a new CollectionLatestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionLatestRequestWithDefaults

`func NewCollectionLatestRequestWithDefaults() *CollectionLatestRequest`

NewCollectionLatestRequestWithDefaults instantiates a new CollectionLatestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceOwnerId

`func (o *CollectionLatestRequest) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *CollectionLatestRequest) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *CollectionLatestRequest) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.


### GetLimit

`func (o *CollectionLatestRequest) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CollectionLatestRequest) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CollectionLatestRequest) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CollectionLatestRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *CollectionLatestRequest) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *CollectionLatestRequest) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *CollectionLatestRequest) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *CollectionLatestRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetProjection

`func (o *CollectionLatestRequest) GetProjection() string`

GetProjection returns the Projection field if non-nil, zero value otherwise.

### GetProjectionOk

`func (o *CollectionLatestRequest) GetProjectionOk() (*string, bool)`

GetProjectionOk returns a tuple with the Projection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjection

`func (o *CollectionLatestRequest) SetProjection(v string)`

SetProjection sets Projection field to given value.

### HasProjection

`func (o *CollectionLatestRequest) HasProjection() bool`

HasProjection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


