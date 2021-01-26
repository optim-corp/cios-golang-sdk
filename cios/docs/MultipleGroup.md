# MultipleGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** |  | 
**Groups** | [**[]Group**](Group.md) |  | 

## Methods

### NewMultipleGroup

`func NewMultipleGroup(total int64, groups []Group, ) *MultipleGroup`

NewMultipleGroup instantiates a new MultipleGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleGroupWithDefaults

`func NewMultipleGroupWithDefaults() *MultipleGroup`

NewMultipleGroupWithDefaults instantiates a new MultipleGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleGroup) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleGroup) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleGroup) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetGroups

`func (o *MultipleGroup) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *MultipleGroup) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *MultipleGroup) SetGroups(v []Group)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


