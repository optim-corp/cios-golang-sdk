# MultipleResourceOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** |  | 
**ResourceOwners** | [**[]ResourceOwner**](ResourceOwner.md) |  | 

## Methods

### NewMultipleResourceOwner

`func NewMultipleResourceOwner(total int64, resourceOwners []ResourceOwner, ) *MultipleResourceOwner`

NewMultipleResourceOwner instantiates a new MultipleResourceOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleResourceOwnerWithDefaults

`func NewMultipleResourceOwnerWithDefaults() *MultipleResourceOwner`

NewMultipleResourceOwnerWithDefaults instantiates a new MultipleResourceOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleResourceOwner) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleResourceOwner) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleResourceOwner) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetResourceOwners

`func (o *MultipleResourceOwner) GetResourceOwners() []ResourceOwner`

GetResourceOwners returns the ResourceOwners field if non-nil, zero value otherwise.

### GetResourceOwnersOk

`func (o *MultipleResourceOwner) GetResourceOwnersOk() (*[]ResourceOwner, bool)`

GetResourceOwnersOk returns a tuple with the ResourceOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwners

`func (o *MultipleResourceOwner) SetResourceOwners(v []ResourceOwner)`

SetResourceOwners sets ResourceOwners field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


