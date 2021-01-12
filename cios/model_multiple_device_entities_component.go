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

// MultipleDeviceEntitiesComponent struct for MultipleDeviceEntitiesComponent
type MultipleDeviceEntitiesComponent struct {
	Total int64 `json:"total"`
	Components []DeviceEntitiesComponent `json:"components"`
}

// NewMultipleDeviceEntitiesComponent instantiates a new MultipleDeviceEntitiesComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipleDeviceEntitiesComponent(total int64, components []DeviceEntitiesComponent, ) *MultipleDeviceEntitiesComponent {
	this := MultipleDeviceEntitiesComponent{}
	this.Total = total
	this.Components = components
	return &this
}

// NewMultipleDeviceEntitiesComponentWithDefaults instantiates a new MultipleDeviceEntitiesComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipleDeviceEntitiesComponentWithDefaults() *MultipleDeviceEntitiesComponent {
	this := MultipleDeviceEntitiesComponent{}
	return &this
}

// GetTotal returns the Total field value
func (o *MultipleDeviceEntitiesComponent) GetTotal() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *MultipleDeviceEntitiesComponent) GetTotalOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *MultipleDeviceEntitiesComponent) SetTotal(v int64) {
	o.Total = v
}

// GetComponents returns the Components field value
func (o *MultipleDeviceEntitiesComponent) GetComponents() []DeviceEntitiesComponent {
	if o == nil  {
		var ret []DeviceEntitiesComponent
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *MultipleDeviceEntitiesComponent) GetComponentsOk() (*[]DeviceEntitiesComponent, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Components, true
}

// SetComponents sets field value
func (o *MultipleDeviceEntitiesComponent) SetComponents(v []DeviceEntitiesComponent) {
	o.Components = v
}

func (o MultipleDeviceEntitiesComponent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["components"] = o.Components
	}
	return json.Marshal(toSerialize)
}

type NullableMultipleDeviceEntitiesComponent struct {
	value *MultipleDeviceEntitiesComponent
	isSet bool
}

func (v NullableMultipleDeviceEntitiesComponent) Get() *MultipleDeviceEntitiesComponent {
	return v.value
}

func (v *NullableMultipleDeviceEntitiesComponent) Set(val *MultipleDeviceEntitiesComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipleDeviceEntitiesComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipleDeviceEntitiesComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipleDeviceEntitiesComponent(val *MultipleDeviceEntitiesComponent) *NullableMultipleDeviceEntitiesComponent {
	return &NullableMultipleDeviceEntitiesComponent{value: val, isSet: true}
}

func (v NullableMultipleDeviceEntitiesComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipleDeviceEntitiesComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

