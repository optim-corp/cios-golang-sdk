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

// SinglePoint struct for SinglePoint
type SinglePoint struct {
	Point Point `json:"point"`
}

// NewSinglePoint instantiates a new SinglePoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSinglePoint(point Point, ) *SinglePoint {
	this := SinglePoint{}
	this.Point = point
	return &this
}

// NewSinglePointWithDefaults instantiates a new SinglePoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSinglePointWithDefaults() *SinglePoint {
	this := SinglePoint{}
	return &this
}

// GetPoint returns the Point field value
func (o *SinglePoint) GetPoint() Point {
	if o == nil  {
		var ret Point
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *SinglePoint) GetPointOk() (*Point, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *SinglePoint) SetPoint(v Point) {
	o.Point = v
}

func (o SinglePoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["point"] = o.Point
	}
	return json.Marshal(toSerialize)
}

type NullableSinglePoint struct {
	value *SinglePoint
	isSet bool
}

func (v NullableSinglePoint) Get() *SinglePoint {
	return v.value
}

func (v *NullableSinglePoint) Set(val *SinglePoint) {
	v.value = val
	v.isSet = true
}

func (v NullableSinglePoint) IsSet() bool {
	return v.isSet
}

func (v *NullableSinglePoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSinglePoint(val *SinglePoint) *NullableSinglePoint {
	return &NullableSinglePoint{value: val, isSet: true}
}

func (v NullableSinglePoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSinglePoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

