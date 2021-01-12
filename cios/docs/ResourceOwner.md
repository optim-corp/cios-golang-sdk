# ResourceOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**GroupId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**AuthorId** | Pointer to **string** |  | [optional] 
**Profile** | [**ResourceOwnerProfile**](ResourceOwner_profile.md) |  | 
**Type** | **string** |  | 

## Methods

### NewResourceOwner

`func NewResourceOwner(id string, profile ResourceOwnerProfile, type_ string, ) *ResourceOwner`

NewResourceOwner instantiates a new ResourceOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceOwnerWithDefaults

`func NewResourceOwnerWithDefaults() *ResourceOwner`

NewResourceOwnerWithDefaults instantiates a new ResourceOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceOwner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceOwner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceOwner) SetId(v string)`

SetId sets Id field to given value.


### GetGroupId

`func (o *ResourceOwner) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ResourceOwner) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ResourceOwner) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ResourceOwner) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetUserId

`func (o *ResourceOwner) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResourceOwner) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResourceOwner) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ResourceOwner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAuthorId

`func (o *ResourceOwner) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *ResourceOwner) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *ResourceOwner) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *ResourceOwner) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetProfile

`func (o *ResourceOwner) GetProfile() ResourceOwnerProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ResourceOwner) GetProfileOk() (*ResourceOwnerProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ResourceOwner) SetProfile(v ResourceOwnerProfile)`

SetProfile sets Profile field to given value.


### GetType

`func (o *ResourceOwner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceOwner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceOwner) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


