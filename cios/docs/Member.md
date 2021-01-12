# Member

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Invited** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**CustomFields** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewMember

`func NewMember() *Member`

NewMember instantiates a new Member object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberWithDefaults

`func NewMemberWithDefaults() *Member`

NewMemberWithDefaults instantiates a new Member object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Member) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Member) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Member) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Member) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *Member) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Member) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Member) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Member) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetRole

`func (o *Member) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Member) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Member) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Member) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetCategory

`func (o *Member) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Member) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Member) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Member) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetInvited

`func (o *Member) GetInvited() bool`

GetInvited returns the Invited field if non-nil, zero value otherwise.

### GetInvitedOk

`func (o *Member) GetInvitedOk() (*bool, bool)`

GetInvitedOk returns a tuple with the Invited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvited

`func (o *Member) SetInvited(v bool)`

SetInvited sets Invited field to given value.

### HasInvited

`func (o *Member) HasInvited() bool`

HasInvited returns a boolean if a field has been set.

### GetTags

`func (o *Member) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Member) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Member) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Member) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Member) GetCustomFields() interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Member) GetCustomFieldsOk() (*interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Member) SetCustomFields(v interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Member) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *Member) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *Member) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


