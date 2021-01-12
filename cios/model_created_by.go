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

// CreatedBy struct for CreatedBy
type CreatedBy struct {
	User *CreatedByUser `json:"user,omitempty"`
	Client *CreatedByClient `json:"client,omitempty"`
}

// NewCreatedBy instantiates a new CreatedBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatedBy() *CreatedBy {
	this := CreatedBy{}
	return &this
}

// NewCreatedByWithDefaults instantiates a new CreatedBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatedByWithDefaults() *CreatedBy {
	this := CreatedBy{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *CreatedBy) GetUser() CreatedByUser {
	if o == nil || o.User == nil {
		var ret CreatedByUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedBy) GetUserOk() (*CreatedByUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *CreatedBy) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given CreatedByUser and assigns it to the User field.
func (o *CreatedBy) SetUser(v CreatedByUser) {
	o.User = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *CreatedBy) GetClient() CreatedByClient {
	if o == nil || o.Client == nil {
		var ret CreatedByClient
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedBy) GetClientOk() (*CreatedByClient, bool) {
	if o == nil || o.Client == nil {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *CreatedBy) HasClient() bool {
	if o != nil && o.Client != nil {
		return true
	}

	return false
}

// SetClient gets a reference to the given CreatedByClient and assigns it to the Client field.
func (o *CreatedBy) SetClient(v CreatedByClient) {
	o.Client = &v
}

func (o CreatedBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Client != nil {
		toSerialize["client"] = o.Client
	}
	return json.Marshal(toSerialize)
}

type NullableCreatedBy struct {
	value *CreatedBy
	isSet bool
}

func (v NullableCreatedBy) Get() *CreatedBy {
	return v.value
}

func (v *NullableCreatedBy) Set(val *CreatedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatedBy(val *CreatedBy) *NullableCreatedBy {
	return &NullableCreatedBy{value: val, isSet: true}
}

func (v NullableCreatedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


