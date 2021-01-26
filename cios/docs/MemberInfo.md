# MemberInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Invited** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Profile** | Pointer to [**MemberInfoProfile**](MemberInfo_profile.md) |  | [optional] 

## Methods

### NewMemberInfo

`func NewMemberInfo() *MemberInfo`

NewMemberInfo instantiates a new MemberInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberInfoWithDefaults

`func NewMemberInfoWithDefaults() *MemberInfo`

NewMemberInfoWithDefaults instantiates a new MemberInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemberInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MemberInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *MemberInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MemberInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MemberInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MemberInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetRole

`func (o *MemberInfo) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MemberInfo) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MemberInfo) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *MemberInfo) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetCategory

`func (o *MemberInfo) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MemberInfo) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MemberInfo) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MemberInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetInvited

`func (o *MemberInfo) GetInvited() bool`

GetInvited returns the Invited field if non-nil, zero value otherwise.

### GetInvitedOk

`func (o *MemberInfo) GetInvitedOk() (*bool, bool)`

GetInvitedOk returns a tuple with the Invited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvited

`func (o *MemberInfo) SetInvited(v bool)`

SetInvited sets Invited field to given value.

### HasInvited

`func (o *MemberInfo) HasInvited() bool`

HasInvited returns a boolean if a field has been set.

### GetTags

`func (o *MemberInfo) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MemberInfo) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MemberInfo) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MemberInfo) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProfile

`func (o *MemberInfo) GetProfile() MemberInfoProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *MemberInfo) GetProfileOk() (*MemberInfoProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *MemberInfo) SetProfile(v MemberInfoProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *MemberInfo) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


