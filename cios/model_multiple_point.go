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

// MultiplePoint struct for MultiplePoint
type MultiplePoint struct {
	Total int64 `json:"total"`
	Points []Point `json:"points"`
}

// NewMultiplePoint instantiates a new MultiplePoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiplePoint(total int64, points []Point, ) *MultiplePoint {
	this := MultiplePoint{}
	this.Total = total
	this.Points = points
	return &this
}

// NewMultiplePointWithDefaults instantiates a new MultiplePoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiplePointWithDefaults() *MultiplePoint {
	this := MultiplePoint{}
	return &this
}

// GetTotal returns the Total field value
func (o *MultiplePoint) GetTotal() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *MultiplePoint) GetTotalOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *MultiplePoint) SetTotal(v int64) {
	o.Total = v
}

// GetPoints returns the Points field value
func (o *MultiplePoint) GetPoints() []Point {
	if o == nil  {
		var ret []Point
		return ret
	}

	return o.Points
}

// GetPointsOk returns a tuple with the Points field value
// and a boolean to check if the value has been set.
func (o *MultiplePoint) GetPointsOk() (*[]Point, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Points, true
}

// SetPoints sets field value
func (o *MultiplePoint) SetPoints(v []Point) {
	o.Points = v
}

func (o MultiplePoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["points"] = o.Points
	}
	return json.Marshal(toSerialize)
}

type NullableMultiplePoint struct {
	value *MultiplePoint
	isSet bool
}

func (v NullableMultiplePoint) Get() *MultiplePoint {
	return v.value
}

func (v *NullableMultiplePoint) Set(val *MultiplePoint) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiplePoint) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiplePoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiplePoint(val *MultiplePoint) *NullableMultiplePoint {
	return &NullableMultiplePoint{value: val, isSet: true}
}

func (v NullableMultiplePoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiplePoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

