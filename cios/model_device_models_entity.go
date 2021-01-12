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

// DeviceModelsEntity struct for DeviceModelsEntity
type DeviceModelsEntity struct {
	Id string `json:"id"`
	Key string `json:"key"`
	DeviceId *string `json:"device_id,omitempty"`
	Model DeviceModelsEntityModel `json:"model"`
	SerialNumber *string `json:"serial_number,omitempty"`
	// ナノ秒
	StartAt string `json:"start_at"`
	CustomInventory *[]map[string]interface{} `json:"custom_inventory,omitempty"`
	ResourceOwnerId string `json:"resource_owner_id"`
	Watch Watch `json:"watch"`
	Components DeviceEntitiesComponent `json:"components"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewDeviceModelsEntity instantiates a new DeviceModelsEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceModelsEntity(id string, key string, model DeviceModelsEntityModel, startAt string, resourceOwnerId string, watch Watch, components DeviceEntitiesComponent, ) *DeviceModelsEntity {
	this := DeviceModelsEntity{}
	this.Id = id
	this.Key = key
	this.Model = model
	this.StartAt = startAt
	this.ResourceOwnerId = resourceOwnerId
	this.Watch = watch
	this.Components = components
	return &this
}

// NewDeviceModelsEntityWithDefaults instantiates a new DeviceModelsEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceModelsEntityWithDefaults() *DeviceModelsEntity {
	this := DeviceModelsEntity{}
	return &this
}

// GetId returns the Id field value
func (o *DeviceModelsEntity) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeviceModelsEntity) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeviceModelsEntity) SetId(v string) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *DeviceModelsEntity) GetKey() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *DeviceModelsEntity) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *DeviceModelsEntity) SetKey(v string) {
	o.Key = v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *DeviceModelsEntity) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceModelsEntity) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *DeviceModelsEntity) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *DeviceModelsEntity) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetModel returns the Model field value
func (o *DeviceModelsEntity) GetModel() DeviceModelsEntityModel {
	if o == nil  {
		var ret DeviceModelsEntityModel
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *DeviceModelsEntity) GetModelOk() (*DeviceModelsEntityModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *DeviceModelsEntity) SetModel(v DeviceModelsEntityModel) {
	o.Model = v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *DeviceModelsEntity) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceModelsEntity) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *DeviceModelsEntity) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *DeviceModelsEntity) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetStartAt returns the StartAt field value
func (o *DeviceModelsEntity) GetStartAt() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.StartAt
}

// GetStartAtOk returns a tuple with the StartAt field value
// and a boolean to check if the value has been set.
func (o *DeviceModelsEntity) GetStartAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartAt, true
}

// SetStartAt sets field value
func (o *DeviceModelsEntity) SetStartAt(v string) {
	o.StartAt = v
}

// GetCustomInventory returns the CustomInventory field value if set, zero value otherwise.
func (o *DeviceModelsEntity) GetCustomInventory() []map[string]interface{} {
	if o == nil || o.CustomInventory == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.CustomInventory
}

// GetCustomInventoryOk returns a tuple with the CustomInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceModelsEntity) GetCustomInventoryOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.CustomInventory == nil {
		return nil, false
	}
	return o.CustomInventory, true
}

// HasCustomInventory returns a boolean if a field has been set.
func (o *DeviceModelsEntity) HasCustomInventory() bool {
	if o != nil && o.CustomInventory != nil {
		return true
	}

	return false
}

// SetCustomInventory gets a reference to the given []map[string]interface{} and assigns it to the CustomInventory field.
func (o *DeviceModelsEntity) SetCustomInventory(v []map[string]interface{}) {
	o.CustomInventory = &v
}

// GetResourceOwnerId returns the ResourceOwnerId field value
func (o *DeviceModelsEntity) GetResourceOwnerId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ResourceOwnerId
}

// GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field value
// and a boolean to check if the value has been set.
func (o *DeviceModelsEntity) GetResourceOwnerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResourceOwnerId, true
}

// SetResourceOwnerId sets field value
func (o *DeviceModelsEntity) SetResourceOwnerId(v string) {
	o.ResourceOwnerId = v
}

// GetWatch returns the Watch field value
func (o *DeviceModelsEntity) GetWatch() Watch {
	if o == nil  {
		var ret Watch
		return ret
	}

	return o.Watch
}

// GetWatchOk returns a tuple with the Watch field value
// and a boolean to check if the value has been set.
func (o *DeviceModelsEntity) GetWatchOk() (*Watch, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Watch, true
}

// SetWatch sets field value
func (o *DeviceModelsEntity) SetWatch(v Watch) {
	o.Watch = v
}

// GetComponents returns the Components field value
func (o *DeviceModelsEntity) GetComponents() DeviceEntitiesComponent {
	if o == nil  {
		var ret DeviceEntitiesComponent
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *DeviceModelsEntity) GetComponentsOk() (*DeviceEntitiesComponent, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Components, true
}

// SetComponents sets field value
func (o *DeviceModelsEntity) SetComponents(v DeviceEntitiesComponent) {
	o.Components = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DeviceModelsEntity) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceModelsEntity) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DeviceModelsEntity) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *DeviceModelsEntity) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DeviceModelsEntity) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceModelsEntity) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DeviceModelsEntity) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *DeviceModelsEntity) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o DeviceModelsEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if o.DeviceId != nil {
		toSerialize["device_id"] = o.DeviceId
	}
	if true {
		toSerialize["model"] = o.Model
	}
	if o.SerialNumber != nil {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if true {
		toSerialize["start_at"] = o.StartAt
	}
	if o.CustomInventory != nil {
		toSerialize["custom_inventory"] = o.CustomInventory
	}
	if true {
		toSerialize["resource_owner_id"] = o.ResourceOwnerId
	}
	if true {
		toSerialize["watch"] = o.Watch
	}
	if true {
		toSerialize["components"] = o.Components
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceModelsEntity struct {
	value *DeviceModelsEntity
	isSet bool
}

func (v NullableDeviceModelsEntity) Get() *DeviceModelsEntity {
	return v.value
}

func (v *NullableDeviceModelsEntity) Set(val *DeviceModelsEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceModelsEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceModelsEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceModelsEntity(val *DeviceModelsEntity) *NullableDeviceModelsEntity {
	return &NullableDeviceModelsEntity{value: val, isSet: true}
}

func (v NullableDeviceModelsEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceModelsEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

