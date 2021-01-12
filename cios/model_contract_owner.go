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

// ContractOwner struct for ContractOwner
type ContractOwner struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	PhoneNumber2 *string `json:"phone_number_2,omitempty"`
	Address *ContractOwnerAddress `json:"address,omitempty"`
}

// NewContractOwner instantiates a new ContractOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractOwner() *ContractOwner {
	this := ContractOwner{}
	return &this
}

// NewContractOwnerWithDefaults instantiates a new ContractOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractOwnerWithDefaults() *ContractOwner {
	this := ContractOwner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContractOwner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractOwner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContractOwner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ContractOwner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContractOwner) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractOwner) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContractOwner) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContractOwner) SetName(v string) {
	o.Name = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *ContractOwner) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractOwner) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *ContractOwner) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *ContractOwner) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetPhoneNumber2 returns the PhoneNumber2 field value if set, zero value otherwise.
func (o *ContractOwner) GetPhoneNumber2() string {
	if o == nil || o.PhoneNumber2 == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber2
}

// GetPhoneNumber2Ok returns a tuple with the PhoneNumber2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractOwner) GetPhoneNumber2Ok() (*string, bool) {
	if o == nil || o.PhoneNumber2 == nil {
		return nil, false
	}
	return o.PhoneNumber2, true
}

// HasPhoneNumber2 returns a boolean if a field has been set.
func (o *ContractOwner) HasPhoneNumber2() bool {
	if o != nil && o.PhoneNumber2 != nil {
		return true
	}

	return false
}

// SetPhoneNumber2 gets a reference to the given string and assigns it to the PhoneNumber2 field.
func (o *ContractOwner) SetPhoneNumber2(v string) {
	o.PhoneNumber2 = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ContractOwner) GetAddress() ContractOwnerAddress {
	if o == nil || o.Address == nil {
		var ret ContractOwnerAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractOwner) GetAddressOk() (*ContractOwnerAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ContractOwner) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given ContractOwnerAddress and assigns it to the Address field.
func (o *ContractOwner) SetAddress(v ContractOwnerAddress) {
	o.Address = &v
}

func (o ContractOwner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PhoneNumber != nil {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if o.PhoneNumber2 != nil {
		toSerialize["phone_number_2"] = o.PhoneNumber2
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableContractOwner struct {
	value *ContractOwner
	isSet bool
}

func (v NullableContractOwner) Get() *ContractOwner {
	return v.value
}

func (v *NullableContractOwner) Set(val *ContractOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableContractOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableContractOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractOwner(val *ContractOwner) *NullableContractOwner {
	return &NullableContractOwner{value: val, isSet: true}
}

func (v NullableContractOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

