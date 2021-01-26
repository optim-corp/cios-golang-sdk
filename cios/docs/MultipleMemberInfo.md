# MultipleMemberInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Members** | Pointer to [**[]MemberInfo**](MemberInfo.md) |  | [optional] 

## Methods

### NewMultipleMemberInfo

`func NewMultipleMemberInfo() *MultipleMemberInfo`

NewMultipleMemberInfo instantiates a new MultipleMemberInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleMemberInfoWithDefaults

`func NewMultipleMemberInfoWithDefaults() *MultipleMemberInfo`

NewMultipleMemberInfoWithDefaults instantiates a new MultipleMemberInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleMemberInfo) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleMemberInfo) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleMemberInfo) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MultipleMemberInfo) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetMembers

`func (o *MultipleMemberInfo) GetMembers() []MemberInfo`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MultipleMemberInfo) GetMembersOk() (*[]MemberInfo, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *MultipleMemberInfo) SetMembers(v []MemberInfo)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *MultipleMemberInfo) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


