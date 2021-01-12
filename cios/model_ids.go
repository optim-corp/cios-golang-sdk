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

// Ids struct for Ids
type Ids struct {
	Ids *[]string `json:"ids,omitempty"`
}

// NewIds instantiates a new Ids object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIds() *Ids {
	this := Ids{}
	return &this
}

// NewIdsWithDefaults instantiates a new Ids object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdsWithDefaults() *Ids {
	this := Ids{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *Ids) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ids) GetIdsOk() (*[]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *Ids) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *Ids) SetIds(v []string) {
	o.Ids = &v
}

func (o Ids) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableIds struct {
	value *Ids
	isSet bool
}

func (v NullableIds) Get() *Ids {
	return v.value
}

func (v *NullableIds) Set(val *Ids) {
	v.value = val
	v.isSet = true
}

func (v NullableIds) IsSet() bool {
	return v.isSet
}

func (v *NullableIds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIds(val *Ids) *NullableIds {
	return &NullableIds{value: val, isSet: true}
}

func (v NullableIds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

