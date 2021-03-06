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

// SingleDataStoreChannel struct for SingleDataStoreChannel
type SingleDataStoreChannel struct {
	Channel DataStoreChannel `json:"channel"`
}

// NewSingleDataStoreChannel instantiates a new SingleDataStoreChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleDataStoreChannel(channel DataStoreChannel, ) *SingleDataStoreChannel {
	this := SingleDataStoreChannel{}
	this.Channel = channel
	return &this
}

// NewSingleDataStoreChannelWithDefaults instantiates a new SingleDataStoreChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleDataStoreChannelWithDefaults() *SingleDataStoreChannel {
	this := SingleDataStoreChannel{}
	return &this
}

// GetChannel returns the Channel field value
func (o *SingleDataStoreChannel) GetChannel() DataStoreChannel {
	if o == nil  {
		var ret DataStoreChannel
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *SingleDataStoreChannel) GetChannelOk() (*DataStoreChannel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *SingleDataStoreChannel) SetChannel(v DataStoreChannel) {
	o.Channel = v
}

func (o SingleDataStoreChannel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["channel"] = o.Channel
	}
	return json.Marshal(toSerialize)
}

type NullableSingleDataStoreChannel struct {
	value *SingleDataStoreChannel
	isSet bool
}

func (v NullableSingleDataStoreChannel) Get() *SingleDataStoreChannel {
	return v.value
}

func (v *NullableSingleDataStoreChannel) Set(val *SingleDataStoreChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleDataStoreChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleDataStoreChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleDataStoreChannel(val *SingleDataStoreChannel) *NullableSingleDataStoreChannel {
	return &NullableSingleDataStoreChannel{value: val, isSet: true}
}

func (v NullableSingleDataStoreChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleDataStoreChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


