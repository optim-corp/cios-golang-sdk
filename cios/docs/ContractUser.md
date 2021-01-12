# ContractUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**FamilyName** | Pointer to **string** |  | [optional] 
**GivenName** | Pointer to **string** |  | [optional] 
**PhoneticFamilyName** | Pointer to **string** |  | [optional] 
**PhoneticGivenName** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**PhoneNumber2** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** | ISO639 | [optional] 
**Picture** | Pointer to **string** | uri | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Emails** | Pointer to **[]string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**ContractOwnerAddress**](ContractOwnerAddress.md) |  | [optional] 
**License** | Pointer to [**ContractUserLicense**](ContractUserLicense.md) |  | [optional] 

## Methods

### NewContractUser

`func NewContractUser() *ContractUser`

NewContractUser instantiates a new ContractUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractUserWithDefaults

`func NewContractUserWithDefaults() *ContractUser`

NewContractUserWithDefaults instantiates a new ContractUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContractUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContractUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContractUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContractUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ContractUser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContractUser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContractUser) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContractUser) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ContractUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContractUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContractUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContractUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFamilyName

`func (o *ContractUser) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *ContractUser) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *ContractUser) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *ContractUser) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetGivenName

`func (o *ContractUser) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *ContractUser) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *ContractUser) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *ContractUser) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetPhoneticFamilyName

`func (o *ContractUser) GetPhoneticFamilyName() string`

GetPhoneticFamilyName returns the PhoneticFamilyName field if non-nil, zero value otherwise.

### GetPhoneticFamilyNameOk

`func (o *ContractUser) GetPhoneticFamilyNameOk() (*string, bool)`

GetPhoneticFamilyNameOk returns a tuple with the PhoneticFamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneticFamilyName

`func (o *ContractUser) SetPhoneticFamilyName(v string)`

SetPhoneticFamilyName sets PhoneticFamilyName field to given value.

### HasPhoneticFamilyName

`func (o *ContractUser) HasPhoneticFamilyName() bool`

HasPhoneticFamilyName returns a boolean if a field has been set.

### GetPhoneticGivenName

`func (o *ContractUser) GetPhoneticGivenName() string`

GetPhoneticGivenName returns the PhoneticGivenName field if non-nil, zero value otherwise.

### GetPhoneticGivenNameOk

`func (o *ContractUser) GetPhoneticGivenNameOk() (*string, bool)`

GetPhoneticGivenNameOk returns a tuple with the PhoneticGivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneticGivenName

`func (o *ContractUser) SetPhoneticGivenName(v string)`

SetPhoneticGivenName sets PhoneticGivenName field to given value.

### HasPhoneticGivenName

`func (o *ContractUser) HasPhoneticGivenName() bool`

HasPhoneticGivenName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *ContractUser) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ContractUser) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ContractUser) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ContractUser) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPhoneNumber2

`func (o *ContractUser) GetPhoneNumber2() string`

GetPhoneNumber2 returns the PhoneNumber2 field if non-nil, zero value otherwise.

### GetPhoneNumber2Ok

`func (o *ContractUser) GetPhoneNumber2Ok() (*string, bool)`

GetPhoneNumber2Ok returns a tuple with the PhoneNumber2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber2

`func (o *ContractUser) SetPhoneNumber2(v string)`

SetPhoneNumber2 sets PhoneNumber2 field to given value.

### HasPhoneNumber2

`func (o *ContractUser) HasPhoneNumber2() bool`

HasPhoneNumber2 returns a boolean if a field has been set.

### GetLanguage

`func (o *ContractUser) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ContractUser) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ContractUser) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ContractUser) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPicture

`func (o *ContractUser) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *ContractUser) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *ContractUser) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *ContractUser) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetEmail

`func (o *ContractUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContractUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContractUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ContractUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmails

`func (o *ContractUser) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ContractUser) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ContractUser) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *ContractUser) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetRole

`func (o *ContractUser) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ContractUser) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ContractUser) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ContractUser) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetCategory

`func (o *ContractUser) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ContractUser) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ContractUser) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ContractUser) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetAddress

`func (o *ContractUser) GetAddress() ContractOwnerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContractUser) GetAddressOk() (*ContractOwnerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContractUser) SetAddress(v ContractOwnerAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ContractUser) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLicense

`func (o *ContractUser) GetLicense() ContractUserLicense`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ContractUser) GetLicenseOk() (*ContractUserLicense, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ContractUser) SetLicense(v ContractUserLicense)`

SetLicense sets License field to given value.

### HasLicense

`func (o *ContractUser) HasLicense() bool`

HasLicense returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


