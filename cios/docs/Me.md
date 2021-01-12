# Me

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**FamilyName** | Pointer to **string** |  | [optional] 
**GivenName** | Pointer to **string** |  | [optional] 
**PhoneticFamilyName** | Pointer to **string** |  | [optional] 
**PhoneticGivenName** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**PhoneNumber2** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**Emails** | Pointer to **[]string** |  | [optional] 
**Picture** | Pointer to **string** |  | [optional] 
**Corporation** | Pointer to [**GroupChildren**](Group_children.md) |  | [optional] 
**Groups** | Pointer to [**[]MeGroups**](MeGroups.md) |  | [optional] 

## Methods

### NewMe

`func NewMe(id string, type_ string, email string, ) *Me`

NewMe instantiates a new Me object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeWithDefaults

`func NewMeWithDefaults() *Me`

NewMeWithDefaults instantiates a new Me object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Me) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Me) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Me) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Me) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Me) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Me) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *Me) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Me) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Me) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Me) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFamilyName

`func (o *Me) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *Me) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *Me) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *Me) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetGivenName

`func (o *Me) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *Me) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *Me) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *Me) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetPhoneticFamilyName

`func (o *Me) GetPhoneticFamilyName() string`

GetPhoneticFamilyName returns the PhoneticFamilyName field if non-nil, zero value otherwise.

### GetPhoneticFamilyNameOk

`func (o *Me) GetPhoneticFamilyNameOk() (*string, bool)`

GetPhoneticFamilyNameOk returns a tuple with the PhoneticFamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneticFamilyName

`func (o *Me) SetPhoneticFamilyName(v string)`

SetPhoneticFamilyName sets PhoneticFamilyName field to given value.

### HasPhoneticFamilyName

`func (o *Me) HasPhoneticFamilyName() bool`

HasPhoneticFamilyName returns a boolean if a field has been set.

### GetPhoneticGivenName

`func (o *Me) GetPhoneticGivenName() string`

GetPhoneticGivenName returns the PhoneticGivenName field if non-nil, zero value otherwise.

### GetPhoneticGivenNameOk

`func (o *Me) GetPhoneticGivenNameOk() (*string, bool)`

GetPhoneticGivenNameOk returns a tuple with the PhoneticGivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneticGivenName

`func (o *Me) SetPhoneticGivenName(v string)`

SetPhoneticGivenName sets PhoneticGivenName field to given value.

### HasPhoneticGivenName

`func (o *Me) HasPhoneticGivenName() bool`

HasPhoneticGivenName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Me) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Me) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Me) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Me) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPhoneNumber2

`func (o *Me) GetPhoneNumber2() string`

GetPhoneNumber2 returns the PhoneNumber2 field if non-nil, zero value otherwise.

### GetPhoneNumber2Ok

`func (o *Me) GetPhoneNumber2Ok() (*string, bool)`

GetPhoneNumber2Ok returns a tuple with the PhoneNumber2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber2

`func (o *Me) SetPhoneNumber2(v string)`

SetPhoneNumber2 sets PhoneNumber2 field to given value.

### HasPhoneNumber2

`func (o *Me) HasPhoneNumber2() bool`

HasPhoneNumber2 returns a boolean if a field has been set.

### GetLanguage

`func (o *Me) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Me) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Me) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Me) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetEmail

`func (o *Me) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Me) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Me) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEmails

`func (o *Me) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *Me) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *Me) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *Me) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetPicture

`func (o *Me) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *Me) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *Me) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *Me) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetCorporation

`func (o *Me) GetCorporation() GroupChildren`

GetCorporation returns the Corporation field if non-nil, zero value otherwise.

### GetCorporationOk

`func (o *Me) GetCorporationOk() (*GroupChildren, bool)`

GetCorporationOk returns a tuple with the Corporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporation

`func (o *Me) SetCorporation(v GroupChildren)`

SetCorporation sets Corporation field to given value.

### HasCorporation

`func (o *Me) HasCorporation() bool`

HasCorporation returns a boolean if a field has been set.

### GetGroups

`func (o *Me) GetGroups() []MeGroups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Me) GetGroupsOk() (*[]MeGroups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Me) SetGroups(v []MeGroups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *Me) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


