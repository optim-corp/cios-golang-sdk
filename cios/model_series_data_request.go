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

// SeriesDataRequest struct for SeriesDataRequest
type SeriesDataRequest struct {
	Location NullableSeriesDataLocation `json:"location,omitempty"`
	// 計測値。送信可能なプロパティはcollectionによって異なる。 collectionごとの定義は[Available Collections](./docs/available-collections.html)を参照。  以下の場合はリクエスト不正となる。 - collectionに定義されていないプロパティがmeasurementsに含まれる場合 - measurementsのJSON Objectのサイズが350KB以上の場合
	Measurements *map[string]map[string]interface{} `json:"measurements,omitempty"`
}

// NewSeriesDataRequest instantiates a new SeriesDataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesDataRequest() *SeriesDataRequest {
	this := SeriesDataRequest{}
	return &this
}

// NewSeriesDataRequestWithDefaults instantiates a new SeriesDataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesDataRequestWithDefaults() *SeriesDataRequest {
	this := SeriesDataRequest{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeriesDataRequest) GetLocation() SeriesDataLocation {
	if o == nil || o.Location.Get() == nil {
		var ret SeriesDataLocation
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeriesDataRequest) GetLocationOk() (*SeriesDataLocation, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *SeriesDataRequest) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableSeriesDataLocation and assigns it to the Location field.
func (o *SeriesDataRequest) SetLocation(v SeriesDataLocation) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *SeriesDataRequest) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *SeriesDataRequest) UnsetLocation() {
	o.Location.Unset()
}

// GetMeasurements returns the Measurements field value if set, zero value otherwise.
func (o *SeriesDataRequest) GetMeasurements() map[string]map[string]interface{} {
	if o == nil || o.Measurements == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Measurements
}

// GetMeasurementsOk returns a tuple with the Measurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesDataRequest) GetMeasurementsOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Measurements == nil {
		return nil, false
	}
	return o.Measurements, true
}

// HasMeasurements returns a boolean if a field has been set.
func (o *SeriesDataRequest) HasMeasurements() bool {
	if o != nil && o.Measurements != nil {
		return true
	}

	return false
}

// SetMeasurements gets a reference to the given map[string]map[string]interface{} and assigns it to the Measurements field.
func (o *SeriesDataRequest) SetMeasurements(v map[string]map[string]interface{}) {
	o.Measurements = &v
}

func (o SeriesDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Measurements != nil {
		toSerialize["measurements"] = o.Measurements
	}
	return json.Marshal(toSerialize)
}

type NullableSeriesDataRequest struct {
	value *SeriesDataRequest
	isSet bool
}

func (v NullableSeriesDataRequest) Get() *SeriesDataRequest {
	return v.value
}

func (v *NullableSeriesDataRequest) Set(val *SeriesDataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesDataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesDataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesDataRequest(val *SeriesDataRequest) *NullableSeriesDataRequest {
	return &NullableSeriesDataRequest{value: val, isSet: true}
}

func (v NullableSeriesDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesDataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

