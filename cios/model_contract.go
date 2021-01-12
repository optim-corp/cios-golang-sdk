/*
 * Cios Openapi
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cios

import (
	"encoding/json"
)

// Contract struct for Contract
type Contract struct {
	Id *string `json:"id,omitempty"`
	ProviderId *string `json:"provider_id,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
	Unit *string `json:"unit,omitempty"`
	LicenseType *string `json:"license_type,omitempty"`
	ContractType *string `json:"contract_type,omitempty"`
	SpanType *string `json:"span_type,omitempty"`
	SpanNumber *int32 `json:"span_number,omitempty"`
	Status *string `json:"status,omitempty"`
	// ISO8601
	PurchasedAt *string `json:"purchased_at,omitempty"`
	// ISO8601
	IssuedAt *string `json:"issued_at,omitempty"`
	// ISO8601
	ExpiresAt *string `json:"expires_at,omitempty"`
	// ISO8601
	StatusChangeAt *string `json:"status_change_at,omitempty"`
	Items *[]ContractItem `json:"items,omitempty"`
	Owner *ContractOwner `json:"owner,omitempty"`
}

// NewContract instantiates a new Contract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContract() *Contract {
	this := Contract{}
	return &this
}

// NewContractWithDefaults instantiates a new Contract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractWithDefaults() *Contract {
	this := Contract{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Contract) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Contract) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Contract) SetId(v string) {
	o.Id = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *Contract) GetProviderId() string {
	if o == nil || o.ProviderId == nil {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetProviderIdOk() (*string, bool) {
	if o == nil || o.ProviderId == nil {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *Contract) HasProviderId() bool {
	if o != nil && o.ProviderId != nil {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *Contract) SetProviderId(v string) {
	o.ProviderId = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *Contract) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *Contract) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *Contract) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Contract) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Contract) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *Contract) SetUnit(v string) {
	o.Unit = &v
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *Contract) GetLicenseType() string {
	if o == nil || o.LicenseType == nil {
		var ret string
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetLicenseTypeOk() (*string, bool) {
	if o == nil || o.LicenseType == nil {
		return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *Contract) HasLicenseType() bool {
	if o != nil && o.LicenseType != nil {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given string and assigns it to the LicenseType field.
func (o *Contract) SetLicenseType(v string) {
	o.LicenseType = &v
}

// GetContractType returns the ContractType field value if set, zero value otherwise.
func (o *Contract) GetContractType() string {
	if o == nil || o.ContractType == nil {
		var ret string
		return ret
	}
	return *o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetContractTypeOk() (*string, bool) {
	if o == nil || o.ContractType == nil {
		return nil, false
	}
	return o.ContractType, true
}

// HasContractType returns a boolean if a field has been set.
func (o *Contract) HasContractType() bool {
	if o != nil && o.ContractType != nil {
		return true
	}

	return false
}

// SetContractType gets a reference to the given string and assigns it to the ContractType field.
func (o *Contract) SetContractType(v string) {
	o.ContractType = &v
}

// GetSpanType returns the SpanType field value if set, zero value otherwise.
func (o *Contract) GetSpanType() string {
	if o == nil || o.SpanType == nil {
		var ret string
		return ret
	}
	return *o.SpanType
}

// GetSpanTypeOk returns a tuple with the SpanType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetSpanTypeOk() (*string, bool) {
	if o == nil || o.SpanType == nil {
		return nil, false
	}
	return o.SpanType, true
}

// HasSpanType returns a boolean if a field has been set.
func (o *Contract) HasSpanType() bool {
	if o != nil && o.SpanType != nil {
		return true
	}

	return false
}

// SetSpanType gets a reference to the given string and assigns it to the SpanType field.
func (o *Contract) SetSpanType(v string) {
	o.SpanType = &v
}

// GetSpanNumber returns the SpanNumber field value if set, zero value otherwise.
func (o *Contract) GetSpanNumber() int32 {
	if o == nil || o.SpanNumber == nil {
		var ret int32
		return ret
	}
	return *o.SpanNumber
}

// GetSpanNumberOk returns a tuple with the SpanNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetSpanNumberOk() (*int32, bool) {
	if o == nil || o.SpanNumber == nil {
		return nil, false
	}
	return o.SpanNumber, true
}

// HasSpanNumber returns a boolean if a field has been set.
func (o *Contract) HasSpanNumber() bool {
	if o != nil && o.SpanNumber != nil {
		return true
	}

	return false
}

// SetSpanNumber gets a reference to the given int32 and assigns it to the SpanNumber field.
func (o *Contract) SetSpanNumber(v int32) {
	o.SpanNumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Contract) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Contract) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Contract) SetStatus(v string) {
	o.Status = &v
}

// GetPurchasedAt returns the PurchasedAt field value if set, zero value otherwise.
func (o *Contract) GetPurchasedAt() string {
	if o == nil || o.PurchasedAt == nil {
		var ret string
		return ret
	}
	return *o.PurchasedAt
}

// GetPurchasedAtOk returns a tuple with the PurchasedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetPurchasedAtOk() (*string, bool) {
	if o == nil || o.PurchasedAt == nil {
		return nil, false
	}
	return o.PurchasedAt, true
}

// HasPurchasedAt returns a boolean if a field has been set.
func (o *Contract) HasPurchasedAt() bool {
	if o != nil && o.PurchasedAt != nil {
		return true
	}

	return false
}

// SetPurchasedAt gets a reference to the given string and assigns it to the PurchasedAt field.
func (o *Contract) SetPurchasedAt(v string) {
	o.PurchasedAt = &v
}

// GetIssuedAt returns the IssuedAt field value if set, zero value otherwise.
func (o *Contract) GetIssuedAt() string {
	if o == nil || o.IssuedAt == nil {
		var ret string
		return ret
	}
	return *o.IssuedAt
}

// GetIssuedAtOk returns a tuple with the IssuedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetIssuedAtOk() (*string, bool) {
	if o == nil || o.IssuedAt == nil {
		return nil, false
	}
	return o.IssuedAt, true
}

// HasIssuedAt returns a boolean if a field has been set.
func (o *Contract) HasIssuedAt() bool {
	if o != nil && o.IssuedAt != nil {
		return true
	}

	return false
}

// SetIssuedAt gets a reference to the given string and assigns it to the IssuedAt field.
func (o *Contract) SetIssuedAt(v string) {
	o.IssuedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *Contract) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *Contract) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *Contract) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetStatusChangeAt returns the StatusChangeAt field value if set, zero value otherwise.
func (o *Contract) GetStatusChangeAt() string {
	if o == nil || o.StatusChangeAt == nil {
		var ret string
		return ret
	}
	return *o.StatusChangeAt
}

// GetStatusChangeAtOk returns a tuple with the StatusChangeAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetStatusChangeAtOk() (*string, bool) {
	if o == nil || o.StatusChangeAt == nil {
		return nil, false
	}
	return o.StatusChangeAt, true
}

// HasStatusChangeAt returns a boolean if a field has been set.
func (o *Contract) HasStatusChangeAt() bool {
	if o != nil && o.StatusChangeAt != nil {
		return true
	}

	return false
}

// SetStatusChangeAt gets a reference to the given string and assigns it to the StatusChangeAt field.
func (o *Contract) SetStatusChangeAt(v string) {
	o.StatusChangeAt = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Contract) GetItems() []ContractItem {
	if o == nil || o.Items == nil {
		var ret []ContractItem
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetItemsOk() (*[]ContractItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Contract) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ContractItem and assigns it to the Items field.
func (o *Contract) SetItems(v []ContractItem) {
	o.Items = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Contract) GetOwner() ContractOwner {
	if o == nil || o.Owner == nil {
		var ret ContractOwner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetOwnerOk() (*ContractOwner, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Contract) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given ContractOwner and assigns it to the Owner field.
func (o *Contract) SetOwner(v ContractOwner) {
	o.Owner = &v
}

func (o Contract) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ProviderId != nil {
		toSerialize["provider_id"] = o.ProviderId
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.LicenseType != nil {
		toSerialize["license_type"] = o.LicenseType
	}
	if o.ContractType != nil {
		toSerialize["contract_type"] = o.ContractType
	}
	if o.SpanType != nil {
		toSerialize["span_type"] = o.SpanType
	}
	if o.SpanNumber != nil {
		toSerialize["span_number"] = o.SpanNumber
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.PurchasedAt != nil {
		toSerialize["purchased_at"] = o.PurchasedAt
	}
	if o.IssuedAt != nil {
		toSerialize["issued_at"] = o.IssuedAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.StatusChangeAt != nil {
		toSerialize["status_change_at"] = o.StatusChangeAt
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	return json.Marshal(toSerialize)
}

type NullableContract struct {
	value *Contract
	isSet bool
}

func (v NullableContract) Get() *Contract {
	return v.value
}

func (v *NullableContract) Set(val *Contract) {
	v.value = val
	v.isSet = true
}

func (v NullableContract) IsSet() bool {
	return v.isSet
}

func (v *NullableContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContract(val *Contract) *NullableContract {
	return &NullableContract{value: val, isSet: true}
}

func (v NullableContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

