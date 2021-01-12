# ContractOwnerAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | ISO 3166-1 alpha-2 | [optional] 
**ZipCode** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Address1** | Pointer to **string** |  | [optional] 
**Address2** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewContractOwnerAddress

`func NewContractOwnerAddress() *ContractOwnerAddress`

NewContractOwnerAddress instantiates a new ContractOwnerAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractOwnerAddressWithDefaults

`func NewContractOwnerAddressWithDefaults() *ContractOwnerAddress`

NewContractOwnerAddressWithDefaults instantiates a new ContractOwnerAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *ContractOwnerAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ContractOwnerAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ContractOwnerAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ContractOwnerAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetZipCode

`func (o *ContractOwnerAddress) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *ContractOwnerAddress) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *ContractOwnerAddress) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *ContractOwnerAddress) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetState

`func (o *ContractOwnerAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ContractOwnerAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ContractOwnerAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ContractOwnerAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCity

`func (o *ContractOwnerAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ContractOwnerAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ContractOwnerAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ContractOwnerAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetAddress1

`func (o *ContractOwnerAddress) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *ContractOwnerAddress) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *ContractOwnerAddress) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *ContractOwnerAddress) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *ContractOwnerAddress) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *ContractOwnerAddress) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *ContractOwnerAddress) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *ContractOwnerAddress) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetTags

`func (o *ContractOwnerAddress) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ContractOwnerAddress) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ContractOwnerAddress) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ContractOwnerAddress) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *ContractOwnerAddress) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ContractOwnerAddress) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ContractOwnerAddress) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ContractOwnerAddress) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


