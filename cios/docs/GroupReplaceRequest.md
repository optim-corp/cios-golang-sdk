# GroupReplaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentGroupId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**PhoneNumber2** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**GroupAddress**](Group_address.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGroupReplaceRequest

`func NewGroupReplaceRequest() *GroupReplaceRequest`

NewGroupReplaceRequest instantiates a new GroupReplaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupReplaceRequestWithDefaults

`func NewGroupReplaceRequestWithDefaults() *GroupReplaceRequest`

NewGroupReplaceRequestWithDefaults instantiates a new GroupReplaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentGroupId

`func (o *GroupReplaceRequest) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *GroupReplaceRequest) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *GroupReplaceRequest) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *GroupReplaceRequest) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetName

`func (o *GroupReplaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupReplaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupReplaceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupReplaceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GroupReplaceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupReplaceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupReplaceRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GroupReplaceRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *GroupReplaceRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *GroupReplaceRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *GroupReplaceRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *GroupReplaceRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPhoneNumber2

`func (o *GroupReplaceRequest) GetPhoneNumber2() string`

GetPhoneNumber2 returns the PhoneNumber2 field if non-nil, zero value otherwise.

### GetPhoneNumber2Ok

`func (o *GroupReplaceRequest) GetPhoneNumber2Ok() (*string, bool)`

GetPhoneNumber2Ok returns a tuple with the PhoneNumber2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber2

`func (o *GroupReplaceRequest) SetPhoneNumber2(v string)`

SetPhoneNumber2 sets PhoneNumber2 field to given value.

### HasPhoneNumber2

`func (o *GroupReplaceRequest) HasPhoneNumber2() bool`

HasPhoneNumber2 returns a boolean if a field has been set.

### GetAddress

`func (o *GroupReplaceRequest) GetAddress() GroupAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GroupReplaceRequest) GetAddressOk() (*GroupAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GroupReplaceRequest) SetAddress(v GroupAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GroupReplaceRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTags

`func (o *GroupReplaceRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GroupReplaceRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GroupReplaceRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GroupReplaceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


