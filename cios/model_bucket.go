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

// Bucket struct for Bucket
type Bucket struct {
	Id string `json:"id"`
	ResourceOwnerId string `json:"resource_owner_id"`
	CreatedAt *string `json:"created_at,omitempty"`
	CreatedBy *CreatedBy `json:"created_by,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	UpdatedBy *UpdatedBy `json:"updated_by,omitempty"`
	Name string `json:"name"`
	Files *BucketFiles `json:"files,omitempty"`
}

// NewBucket instantiates a new Bucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucket(id string, resourceOwnerId string, name string, ) *Bucket {
	this := Bucket{}
	this.Id = id
	this.ResourceOwnerId = resourceOwnerId
	this.Name = name
	return &this
}

// NewBucketWithDefaults instantiates a new Bucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketWithDefaults() *Bucket {
	this := Bucket{}
	return &this
}

// GetId returns the Id field value
func (o *Bucket) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Bucket) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Bucket) SetId(v string) {
	o.Id = v
}

// GetResourceOwnerId returns the ResourceOwnerId field value
func (o *Bucket) GetResourceOwnerId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ResourceOwnerId
}

// GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field value
// and a boolean to check if the value has been set.
func (o *Bucket) GetResourceOwnerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResourceOwnerId, true
}

// SetResourceOwnerId sets field value
func (o *Bucket) SetResourceOwnerId(v string) {
	o.ResourceOwnerId = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Bucket) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bucket) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Bucket) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Bucket) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Bucket) GetCreatedBy() CreatedBy {
	if o == nil || o.CreatedBy == nil {
		var ret CreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bucket) GetCreatedByOk() (*CreatedBy, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Bucket) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given CreatedBy and assigns it to the CreatedBy field.
func (o *Bucket) SetCreatedBy(v CreatedBy) {
	o.CreatedBy = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Bucket) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bucket) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Bucket) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Bucket) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *Bucket) GetUpdatedBy() UpdatedBy {
	if o == nil || o.UpdatedBy == nil {
		var ret UpdatedBy
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bucket) GetUpdatedByOk() (*UpdatedBy, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *Bucket) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given UpdatedBy and assigns it to the UpdatedBy field.
func (o *Bucket) SetUpdatedBy(v UpdatedBy) {
	o.UpdatedBy = &v
}

// GetName returns the Name field value
func (o *Bucket) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Bucket) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Bucket) SetName(v string) {
	o.Name = v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *Bucket) GetFiles() BucketFiles {
	if o == nil || o.Files == nil {
		var ret BucketFiles
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bucket) GetFilesOk() (*BucketFiles, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *Bucket) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given BucketFiles and assigns it to the Files field.
func (o *Bucket) SetFiles(v BucketFiles) {
	o.Files = &v
}

func (o Bucket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["resource_owner_id"] = o.ResourceOwnerId
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	return json.Marshal(toSerialize)
}

type NullableBucket struct {
	value *Bucket
	isSet bool
}

func (v NullableBucket) Get() *Bucket {
	return v.value
}

func (v *NullableBucket) Set(val *Bucket) {
	v.value = val
	v.isSet = true
}

func (v NullableBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucket(val *Bucket) *NullableBucket {
	return &NullableBucket{value: val, isSet: true}
}

func (v NullableBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


