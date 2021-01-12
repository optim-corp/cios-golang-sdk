# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Owner** | Pointer to [**SubscriptionOwner**](SubscriptionOwner.md) |  | [optional] 
**Items** | Pointer to [**[]SubscriptionItem**](SubscriptionItem.md) |  | [optional] 
**ProviderId** | Pointer to **string** | UUID形式 | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**LicenseType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**PurchasedAt** | Pointer to **string** | ISO8601 | [optional] 
**IssuedAt** | Pointer to **string** | ISO8601 | [optional] 
**ExpiresAt** | Pointer to **string** | ISO8601 | [optional] 

## Methods

### NewSubscription

`func NewSubscription(id string, ) *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v string)`

SetId sets Id field to given value.


### GetOwner

`func (o *Subscription) GetOwner() SubscriptionOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Subscription) GetOwnerOk() (*SubscriptionOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Subscription) SetOwner(v SubscriptionOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Subscription) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetItems

`func (o *Subscription) GetItems() []SubscriptionItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Subscription) GetItemsOk() (*[]SubscriptionItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Subscription) SetItems(v []SubscriptionItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *Subscription) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetProviderId

`func (o *Subscription) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *Subscription) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *Subscription) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *Subscription) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetQuantity

`func (o *Subscription) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Subscription) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Subscription) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Subscription) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnit

`func (o *Subscription) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Subscription) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Subscription) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Subscription) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetLicenseType

`func (o *Subscription) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *Subscription) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *Subscription) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *Subscription) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetStatus

`func (o *Subscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Subscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Subscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Subscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPurchasedAt

`func (o *Subscription) GetPurchasedAt() string`

GetPurchasedAt returns the PurchasedAt field if non-nil, zero value otherwise.

### GetPurchasedAtOk

`func (o *Subscription) GetPurchasedAtOk() (*string, bool)`

GetPurchasedAtOk returns a tuple with the PurchasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedAt

`func (o *Subscription) SetPurchasedAt(v string)`

SetPurchasedAt sets PurchasedAt field to given value.

### HasPurchasedAt

`func (o *Subscription) HasPurchasedAt() bool`

HasPurchasedAt returns a boolean if a field has been set.

### GetIssuedAt

`func (o *Subscription) GetIssuedAt() string`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *Subscription) GetIssuedAtOk() (*string, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *Subscription) SetIssuedAt(v string)`

SetIssuedAt sets IssuedAt field to given value.

### HasIssuedAt

`func (o *Subscription) HasIssuedAt() bool`

HasIssuedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *Subscription) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Subscription) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Subscription) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Subscription) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


