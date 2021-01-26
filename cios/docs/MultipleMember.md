# MultipleMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Members** | Pointer to [**[]Member**](Member.md) |  | [optional] 

## Methods

### NewMultipleMember

`func NewMultipleMember() *MultipleMember`

NewMultipleMember instantiates a new MultipleMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleMemberWithDefaults

`func NewMultipleMemberWithDefaults() *MultipleMember`

NewMultipleMemberWithDefaults instantiates a new MultipleMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleMember) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleMember) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleMember) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MultipleMember) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetMembers

`func (o *MultipleMember) GetMembers() []Member`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MultipleMember) GetMembersOk() (*[]Member, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *MultipleMember) SetMembers(v []Member)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *MultipleMember) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


