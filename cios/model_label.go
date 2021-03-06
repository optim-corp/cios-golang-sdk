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

// Label struct for Label
type Label struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

// NewLabel instantiates a new Label object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabel(key string, value string, ) *Label {
	this := Label{}
	this.Key = key
	this.Value = value
	return &this
}

// NewLabelWithDefaults instantiates a new Label object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelWithDefaults() *Label {
	this := Label{}
	return &this
}

// GetKey returns the Key field value
func (o *Label) GetKey() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Label) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *Label) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *Label) GetValue() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Label) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Label) SetValue(v string) {
	o.Value = v
}

func (o Label) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableLabel struct {
	value *Label
	isSet bool
}

func (v NullableLabel) Get() *Label {
	return v.value
}

func (v *NullableLabel) Set(val *Label) {
	v.value = val
	v.isSet = true
}

func (v NullableLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabel(val *Label) *NullableLabel {
	return &NullableLabel{value: val, isSet: true}
}

func (v NullableLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


