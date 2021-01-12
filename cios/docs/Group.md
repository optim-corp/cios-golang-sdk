# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ParentGroupId** | Pointer to **string** |  | [optional] 
**CorporationId** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Domain** | Pointer to **NullableString** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**PhoneNumber2** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**GroupAddress**](Group_address.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Children** | Pointer to [**[]GroupChildren**](GroupChildren.md) |  | [optional] 

## Methods

### NewGroup

`func NewGroup(id string, name string, type_ string, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v string)`

SetId sets Id field to given value.


### GetParentGroupId

`func (o *Group) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *Group) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *Group) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *Group) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetCorporationId

`func (o *Group) GetCorporationId() string`

GetCorporationId returns the CorporationId field if non-nil, zero value otherwise.

### GetCorporationIdOk

`func (o *Group) GetCorporationIdOk() (*string, bool)`

GetCorporationIdOk returns a tuple with the CorporationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporationId

`func (o *Group) SetCorporationId(v string)`

SetCorporationId sets CorporationId field to given value.

### HasCorporationId

`func (o *Group) HasCorporationId() bool`

HasCorporationId returns a boolean if a field has been set.

### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Group) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Group) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Group) SetType(v string)`

SetType sets Type field to given value.


### GetDomain

`func (o *Group) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Group) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Group) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Group) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *Group) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *Group) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetPhoneNumber

`func (o *Group) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Group) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Group) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Group) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPhoneNumber2

`func (o *Group) GetPhoneNumber2() string`

GetPhoneNumber2 returns the PhoneNumber2 field if non-nil, zero value otherwise.

### GetPhoneNumber2Ok

`func (o *Group) GetPhoneNumber2Ok() (*string, bool)`

GetPhoneNumber2Ok returns a tuple with the PhoneNumber2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber2

`func (o *Group) SetPhoneNumber2(v string)`

SetPhoneNumber2 sets PhoneNumber2 field to given value.

### HasPhoneNumber2

`func (o *Group) HasPhoneNumber2() bool`

HasPhoneNumber2 returns a boolean if a field has been set.

### GetAddress

`func (o *Group) GetAddress() GroupAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Group) GetAddressOk() (*GroupAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Group) SetAddress(v GroupAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Group) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTags

`func (o *Group) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Group) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Group) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Group) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Group) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Group) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Group) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Group) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetChildren

`func (o *Group) GetChildren() []GroupChildren`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Group) GetChildrenOk() (*[]GroupChildren, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Group) SetChildren(v []GroupChildren)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Group) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


