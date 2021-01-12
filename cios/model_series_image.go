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

// SeriesImage struct for SeriesImage
type SeriesImage struct {
	Timestamp string `json:"timestamp"`
	// base64エンコードされた画像データ
	Image string `json:"image"`
	// 画像データのフォーマット
	ImageType string `json:"image_type"`
}

// NewSeriesImage instantiates a new SeriesImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesImage(timestamp string, image string, imageType string, ) *SeriesImage {
	this := SeriesImage{}
	this.Timestamp = timestamp
	this.Image = image
	this.ImageType = imageType
	return &this
}

// NewSeriesImageWithDefaults instantiates a new SeriesImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesImageWithDefaults() *SeriesImage {
	this := SeriesImage{}
	return &this
}

// GetTimestamp returns the Timestamp field value
func (o *SeriesImage) GetTimestamp() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *SeriesImage) GetTimestampOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *SeriesImage) SetTimestamp(v string) {
	o.Timestamp = v
}

// GetImage returns the Image field value
func (o *SeriesImage) GetImage() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *SeriesImage) GetImageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *SeriesImage) SetImage(v string) {
	o.Image = v
}

// GetImageType returns the ImageType field value
func (o *SeriesImage) GetImageType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ImageType
}

// GetImageTypeOk returns a tuple with the ImageType field value
// and a boolean to check if the value has been set.
func (o *SeriesImage) GetImageTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImageType, true
}

// SetImageType sets field value
func (o *SeriesImage) SetImageType(v string) {
	o.ImageType = v
}

func (o SeriesImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["image"] = o.Image
	}
	if true {
		toSerialize["image_type"] = o.ImageType
	}
	return json.Marshal(toSerialize)
}

type NullableSeriesImage struct {
	value *SeriesImage
	isSet bool
}

func (v NullableSeriesImage) Get() *SeriesImage {
	return v.value
}

func (v *NullableSeriesImage) Set(val *SeriesImage) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesImage) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesImage(val *SeriesImage) *NullableSeriesImage {
	return &NullableSeriesImage{value: val, isSet: true}
}

func (v NullableSeriesImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

