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

// SeriesDataLocation デバイス位置情報。コレクションの定義に依存。schema.location_typeが”POINT“の場合必須、”NONE”の場合はlocationをリクエストに含めることはできない
type SeriesDataLocation struct {
	// Point固定
	Type string `json:"type"`
	Coordinates []float32 `json:"coordinates"`
}

// NewSeriesDataLocation instantiates a new SeriesDataLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesDataLocation(type_ string, coordinates []float32, ) *SeriesDataLocation {
	this := SeriesDataLocation{}
	this.Type = type_
	this.Coordinates = coordinates
	return &this
}

// NewSeriesDataLocationWithDefaults instantiates a new SeriesDataLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesDataLocationWithDefaults() *SeriesDataLocation {
	this := SeriesDataLocation{}
	return &this
}

// GetType returns the Type field value
func (o *SeriesDataLocation) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SeriesDataLocation) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SeriesDataLocation) SetType(v string) {
	o.Type = v
}

// GetCoordinates returns the Coordinates field value
func (o *SeriesDataLocation) GetCoordinates() []float32 {
	if o == nil  {
		var ret []float32
		return ret
	}

	return o.Coordinates
}

// GetCoordinatesOk returns a tuple with the Coordinates field value
// and a boolean to check if the value has been set.
func (o *SeriesDataLocation) GetCoordinatesOk() (*[]float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Coordinates, true
}

// SetCoordinates sets field value
func (o *SeriesDataLocation) SetCoordinates(v []float32) {
	o.Coordinates = v
}

func (o SeriesDataLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["coordinates"] = o.Coordinates
	}
	return json.Marshal(toSerialize)
}

type NullableSeriesDataLocation struct {
	value *SeriesDataLocation
	isSet bool
}

func (v NullableSeriesDataLocation) Get() *SeriesDataLocation {
	return v.value
}

func (v *NullableSeriesDataLocation) Set(val *SeriesDataLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesDataLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesDataLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesDataLocation(val *SeriesDataLocation) *NullableSeriesDataLocation {
	return &NullableSeriesDataLocation{value: val, isSet: true}
}

func (v NullableSeriesDataLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesDataLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


