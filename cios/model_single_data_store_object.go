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

// SingleDataStoreObject struct for SingleDataStoreObject
type SingleDataStoreObject struct {
	Object DataStoreObject `json:"object"`
}

// NewSingleDataStoreObject instantiates a new SingleDataStoreObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleDataStoreObject(object DataStoreObject, ) *SingleDataStoreObject {
	this := SingleDataStoreObject{}
	this.Object = object
	return &this
}

// NewSingleDataStoreObjectWithDefaults instantiates a new SingleDataStoreObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleDataStoreObjectWithDefaults() *SingleDataStoreObject {
	this := SingleDataStoreObject{}
	return &this
}

// GetObject returns the Object field value
func (o *SingleDataStoreObject) GetObject() DataStoreObject {
	if o == nil  {
		var ret DataStoreObject
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *SingleDataStoreObject) GetObjectOk() (*DataStoreObject, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *SingleDataStoreObject) SetObject(v DataStoreObject) {
	o.Object = v
}

func (o SingleDataStoreObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["object"] = o.Object
	}
	return json.Marshal(toSerialize)
}

type NullableSingleDataStoreObject struct {
	value *SingleDataStoreObject
	isSet bool
}

func (v NullableSingleDataStoreObject) Get() *SingleDataStoreObject {
	return v.value
}

func (v *NullableSingleDataStoreObject) Set(val *SingleDataStoreObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleDataStoreObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleDataStoreObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleDataStoreObject(val *SingleDataStoreObject) *NullableSingleDataStoreObject {
	return &NullableSingleDataStoreObject{value: val, isSet: true}
}

func (v NullableSingleDataStoreObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleDataStoreObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


