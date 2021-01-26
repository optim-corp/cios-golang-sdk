# GroupCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentGroupId** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**PhoneNumber2** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**GroupAddress**](Group_address.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | **string** |  | 
**CustomFields** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewGroupCreateRequest

`func NewGroupCreateRequest(name string, type_ string, ) *GroupCreateRequest`

NewGroupCreateRequest instantiates a new GroupCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupCreateRequestWithDefaults

`func NewGroupCreateRequestWithDefaults() *GroupCreateRequest`

NewGroupCreateRequestWithDefaults instantiates a new GroupCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentGroupId

`func (o *GroupCreateRequest) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *GroupCreateRequest) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *GroupCreateRequest) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *GroupCreateRequest) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetName

`func (o *GroupCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPhoneNumber

`func (o *GroupCreateRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *GroupCreateRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *GroupCreateRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *GroupCreateRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPhoneNumber2

`func (o *GroupCreateRequest) GetPhoneNumber2() string`

GetPhoneNumber2 returns the PhoneNumber2 field if non-nil, zero value otherwise.

### GetPhoneNumber2Ok

`func (o *GroupCreateRequest) GetPhoneNumber2Ok() (*string, bool)`

GetPhoneNumber2Ok returns a tuple with the PhoneNumber2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber2

`func (o *GroupCreateRequest) SetPhoneNumber2(v string)`

SetPhoneNumber2 sets PhoneNumber2 field to given value.

### HasPhoneNumber2

`func (o *GroupCreateRequest) HasPhoneNumber2() bool`

HasPhoneNumber2 returns a boolean if a field has been set.

### GetAddress

`func (o *GroupCreateRequest) GetAddress() GroupAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GroupCreateRequest) GetAddressOk() (*GroupAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GroupCreateRequest) SetAddress(v GroupAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GroupCreateRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTags

`func (o *GroupCreateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GroupCreateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GroupCreateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GroupCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *GroupCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupCreateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCustomFields

`func (o *GroupCreateRequest) GetCustomFields() interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *GroupCreateRequest) GetCustomFieldsOk() (*interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *GroupCreateRequest) SetCustomFields(v interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *GroupCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *GroupCreateRequest) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *GroupCreateRequest) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


