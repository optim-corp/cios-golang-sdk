# MemberInfoProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**FamilyName** | Pointer to **string** |  | [optional] 
**GivenName** | Pointer to **string** |  | [optional] 
**PhoneticFamilyName** | Pointer to **string** |  | [optional] 
**PhoneticGivenName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Emails** | Pointer to **[]string** |  | [optional] 
**Picture** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**PhoneNumber2** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**GroupAddress**](Group_address.md) |  | [optional] 

## Methods

### NewMemberInfoProfile

`func NewMemberInfoProfile() *MemberInfoProfile`

NewMemberInfoProfile instantiates a new MemberInfoProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberInfoProfileWithDefaults

`func NewMemberInfoProfileWithDefaults() *MemberInfoProfile`

NewMemberInfoProfileWithDefaults instantiates a new MemberInfoProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MemberInfoProfile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MemberInfoProfile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MemberInfoProfile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MemberInfoProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFamilyName

`func (o *MemberInfoProfile) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *MemberInfoProfile) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *MemberInfoProfile) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *MemberInfoProfile) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetGivenName

`func (o *MemberInfoProfile) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *MemberInfoProfile) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *MemberInfoProfile) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *MemberInfoProfile) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetPhoneticFamilyName

`func (o *MemberInfoProfile) GetPhoneticFamilyName() string`

GetPhoneticFamilyName returns the PhoneticFamilyName field if non-nil, zero value otherwise.

### GetPhoneticFamilyNameOk

`func (o *MemberInfoProfile) GetPhoneticFamilyNameOk() (*string, bool)`

GetPhoneticFamilyNameOk returns a tuple with the PhoneticFamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneticFamilyName

`func (o *MemberInfoProfile) SetPhoneticFamilyName(v string)`

SetPhoneticFamilyName sets PhoneticFamilyName field to given value.

### HasPhoneticFamilyName

`func (o *MemberInfoProfile) HasPhoneticFamilyName() bool`

HasPhoneticFamilyName returns a boolean if a field has been set.

### GetPhoneticGivenName

`func (o *MemberInfoProfile) GetPhoneticGivenName() string`

GetPhoneticGivenName returns the PhoneticGivenName field if non-nil, zero value otherwise.

### GetPhoneticGivenNameOk

`func (o *MemberInfoProfile) GetPhoneticGivenNameOk() (*string, bool)`

GetPhoneticGivenNameOk returns a tuple with the PhoneticGivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneticGivenName

`func (o *MemberInfoProfile) SetPhoneticGivenName(v string)`

SetPhoneticGivenName sets PhoneticGivenName field to given value.

### HasPhoneticGivenName

`func (o *MemberInfoProfile) HasPhoneticGivenName() bool`

HasPhoneticGivenName returns a boolean if a field has been set.

### GetEmail

`func (o *MemberInfoProfile) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MemberInfoProfile) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MemberInfoProfile) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MemberInfoProfile) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmails

`func (o *MemberInfoProfile) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *MemberInfoProfile) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *MemberInfoProfile) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *MemberInfoProfile) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetPicture

`func (o *MemberInfoProfile) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *MemberInfoProfile) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *MemberInfoProfile) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *MemberInfoProfile) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *MemberInfoProfile) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *MemberInfoProfile) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *MemberInfoProfile) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *MemberInfoProfile) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPhoneNumber2

`func (o *MemberInfoProfile) GetPhoneNumber2() string`

GetPhoneNumber2 returns the PhoneNumber2 field if non-nil, zero value otherwise.

### GetPhoneNumber2Ok

`func (o *MemberInfoProfile) GetPhoneNumber2Ok() (*string, bool)`

GetPhoneNumber2Ok returns a tuple with the PhoneNumber2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber2

`func (o *MemberInfoProfile) SetPhoneNumber2(v string)`

SetPhoneNumber2 sets PhoneNumber2 field to given value.

### HasPhoneNumber2

`func (o *MemberInfoProfile) HasPhoneNumber2() bool`

HasPhoneNumber2 returns a boolean if a field has been set.

### GetLanguage

`func (o *MemberInfoProfile) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MemberInfoProfile) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MemberInfoProfile) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MemberInfoProfile) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetAddress

`func (o *MemberInfoProfile) GetAddress() GroupAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MemberInfoProfile) GetAddressOk() (*GroupAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MemberInfoProfile) SetAddress(v GroupAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MemberInfoProfile) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


