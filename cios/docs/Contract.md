# Contract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**LicenseType** | Pointer to **string** |  | [optional] 
**ContractType** | Pointer to **string** |  | [optional] 
**SpanType** | Pointer to **string** |  | [optional] 
**SpanNumber** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**PurchasedAt** | Pointer to **string** | ISO8601 | [optional] 
**IssuedAt** | Pointer to **string** | ISO8601 | [optional] 
**ExpiresAt** | Pointer to **string** | ISO8601 | [optional] 
**StatusChangeAt** | Pointer to **string** | ISO8601 | [optional] 
**Items** | Pointer to [**[]ContractItem**](ContractItem.md) |  | [optional] 
**Owner** | Pointer to [**ContractOwner**](ContractOwner.md) |  | [optional] 

## Methods

### NewContract

`func NewContract() *Contract`

NewContract instantiates a new Contract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractWithDefaults

`func NewContractWithDefaults() *Contract`

NewContractWithDefaults instantiates a new Contract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Contract) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Contract) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Contract) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Contract) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderId

`func (o *Contract) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *Contract) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *Contract) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *Contract) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetQuantity

`func (o *Contract) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Contract) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Contract) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Contract) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnit

`func (o *Contract) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Contract) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Contract) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Contract) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetLicenseType

`func (o *Contract) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *Contract) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *Contract) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *Contract) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetContractType

`func (o *Contract) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *Contract) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *Contract) SetContractType(v string)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *Contract) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetSpanType

`func (o *Contract) GetSpanType() string`

GetSpanType returns the SpanType field if non-nil, zero value otherwise.

### GetSpanTypeOk

`func (o *Contract) GetSpanTypeOk() (*string, bool)`

GetSpanTypeOk returns a tuple with the SpanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanType

`func (o *Contract) SetSpanType(v string)`

SetSpanType sets SpanType field to given value.

### HasSpanType

`func (o *Contract) HasSpanType() bool`

HasSpanType returns a boolean if a field has been set.

### GetSpanNumber

`func (o *Contract) GetSpanNumber() int32`

GetSpanNumber returns the SpanNumber field if non-nil, zero value otherwise.

### GetSpanNumberOk

`func (o *Contract) GetSpanNumberOk() (*int32, bool)`

GetSpanNumberOk returns a tuple with the SpanNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanNumber

`func (o *Contract) SetSpanNumber(v int32)`

SetSpanNumber sets SpanNumber field to given value.

### HasSpanNumber

`func (o *Contract) HasSpanNumber() bool`

HasSpanNumber returns a boolean if a field has been set.

### GetStatus

`func (o *Contract) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Contract) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Contract) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Contract) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPurchasedAt

`func (o *Contract) GetPurchasedAt() string`

GetPurchasedAt returns the PurchasedAt field if non-nil, zero value otherwise.

### GetPurchasedAtOk

`func (o *Contract) GetPurchasedAtOk() (*string, bool)`

GetPurchasedAtOk returns a tuple with the PurchasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedAt

`func (o *Contract) SetPurchasedAt(v string)`

SetPurchasedAt sets PurchasedAt field to given value.

### HasPurchasedAt

`func (o *Contract) HasPurchasedAt() bool`

HasPurchasedAt returns a boolean if a field has been set.

### GetIssuedAt

`func (o *Contract) GetIssuedAt() string`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *Contract) GetIssuedAtOk() (*string, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *Contract) SetIssuedAt(v string)`

SetIssuedAt sets IssuedAt field to given value.

### HasIssuedAt

`func (o *Contract) HasIssuedAt() bool`

HasIssuedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *Contract) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Contract) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Contract) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Contract) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetStatusChangeAt

`func (o *Contract) GetStatusChangeAt() string`

GetStatusChangeAt returns the StatusChangeAt field if non-nil, zero value otherwise.

### GetStatusChangeAtOk

`func (o *Contract) GetStatusChangeAtOk() (*string, bool)`

GetStatusChangeAtOk returns a tuple with the StatusChangeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChangeAt

`func (o *Contract) SetStatusChangeAt(v string)`

SetStatusChangeAt sets StatusChangeAt field to given value.

### HasStatusChangeAt

`func (o *Contract) HasStatusChangeAt() bool`

HasStatusChangeAt returns a boolean if a field has been set.

### GetItems

`func (o *Contract) GetItems() []ContractItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Contract) GetItemsOk() (*[]ContractItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Contract) SetItems(v []ContractItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *Contract) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOwner

`func (o *Contract) GetOwner() ContractOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Contract) GetOwnerOk() (*ContractOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Contract) SetOwner(v ContractOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Contract) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


